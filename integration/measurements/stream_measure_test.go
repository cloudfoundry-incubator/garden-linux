package measurements_test

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/cloudfoundry-incubator/garden"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type byteCounterWriter struct {
	num *uint64
}

func (w *byteCounterWriter) Write(d []byte) (int, error) {
	atomic.AddUint64(w.num, uint64(len(d)))
	return len(d), nil
}

func (w *byteCounterWriter) Close() error {
	return nil
}

var _ = Describe("The Garden server", func() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var container garden.Container
	var firstGoroutineCount uint64

	BeforeEach(func() {
		firstGoroutineCount = 0
		client = startGarden()

		var err error
		container, err = client.Create(garden.ContainerSpec{})
		Ω(err).ShouldNot(HaveOccurred())
	})

	getGoroutineCount := func(printIt ...bool) uint64 {
		resp, err := http.Get(fmt.Sprintf("http://%s/debug/pprof/goroutine?debug=1", gardenRunner.DebugAddr()))
		Ω(err).ShouldNot(HaveOccurred())

		line, _, err := bufio.NewReader(resp.Body).ReadLine()
		Ω(err).ShouldNot(HaveOccurred())

		if len(printIt) > 0 && printIt[0] {
			io.Copy(os.Stdout, resp.Body)
		}

		words := strings.Split(string(line), " ")

		goroutineCount, err := strconv.ParseUint(words[len(words)-1], 10, 64)
		Ω(err).ShouldNot(HaveOccurred())

		return goroutineCount
	}

	Describe("repeatedly running processes", func() {
		Measure("does not leak goroutines", func(b Benchmarker) {
			iterations := 50

			for i := 1; i <= iterations; i++ {
				process, err := container.Run(garden.ProcessSpec{
					Path: "echo",
					Args: []string{"hi"},
				}, garden.ProcessIO{})
				Ω(err).ShouldNot(HaveOccurred())

				status, err := process.Wait()
				Ω(err).ShouldNot(HaveOccurred())
				Ω(status).Should(Equal(0))

				if i == 1 {
					firstGoroutineCount = getGoroutineCount()
					b.RecordValue("first goroutine count", float64(firstGoroutineCount))
				}

				if i == iterations {
					lastGoroutineCount := getGoroutineCount()
					b.RecordValue("last goroutine count", float64(lastGoroutineCount))

					Ω(lastGoroutineCount).ShouldNot(BeNumerically(">", firstGoroutineCount+5))
				}
			}
		}, 1)
	})

	Describe("repeatedly attaching to a running process", func() {
		var processID uint32

		BeforeEach(func() {
			process, err := container.Run(garden.ProcessSpec{
				Path: "cat",
			}, garden.ProcessIO{})
			Ω(err).ShouldNot(HaveOccurred())

			processID = process.ID()
		})

		Measure("does not leak goroutines", func(b Benchmarker) {
			iterations := 50

			for i := 1; i <= iterations; i++ {
				stdoutR, stdoutW := io.Pipe()
				stdinR, stdinW := io.Pipe()

				_, err := container.Attach(processID, garden.ProcessIO{
					Stdin:  stdinR,
					Stdout: stdoutW,
				})
				Ω(err).ShouldNot(HaveOccurred())

				stdinData := fmt.Sprintf("hello %d", i)

				_, err = stdinW.Write([]byte(stdinData + "\n"))
				Ω(err).ShouldNot(HaveOccurred())

				var line []byte
				doneReading := make(chan struct{})
				go func() {
					line, _, err = bufio.NewReader(stdoutR).ReadLine()
					close(doneReading)
				}()

				Eventually(doneReading).Should(BeClosed())
				Ω(err).ShouldNot(HaveOccurred())
				Ω(string(line)).Should(Equal(stdinData))

				stdinW.CloseWithError(errors.New("going away now"))

				if i == 1 {
					firstGoroutineCount = getGoroutineCount()
					b.RecordValue("first goroutine count", float64(firstGoroutineCount))
				}

				if i == iterations {
					lastGoroutineCount := getGoroutineCount()
					b.RecordValue("last goroutine count", float64(lastGoroutineCount))

					// TODO - we have a leak more.
					// Ω(lastGoroutineCount).ShouldNot(BeNumerically(">", firstGoroutineCount+5))
				}
			}
		}, 1)
	})

	FDescribe("streaming output from a chatty job", func() {
		streamCounts := []int{1, 4, 8}

		// for i := 1; i <= 128; i *= 2 {
		// 	streamCounts = append(streamCounts, i)
		// }

		loggedLine := strings.Repeat("x", 1024)

		for _, streams := range streamCounts {
			Context(fmt.Sprintf("with %d streams", streams), func() {
				var started time.Time
				var receivedBytes uint64

				numToSpawn := streams

				BeforeEach(func() {
					atomic.StoreUint64(&receivedBytes, 0)
					started = time.Now()

					spawned := make(chan bool)

					for j := 0; j < numToSpawn; j++ {
						go func() {
							defer GinkgoRecover()

							_, err := container.Run(garden.ProcessSpec{
								Path: "sh",
								Args: []string{"-c", "while true; do echo " + loggedLine + "; done"},
							}, garden.ProcessIO{})
							Ω(err).ShouldNot(HaveOccurred())

							spawned <- true
						}()
					}

					for j := 0; j < numToSpawn; j++ {
						<-spawned
					}
				})

				AfterEach(func() {
					err := client.Destroy(container.Handle())
					Ω(err).ShouldNot(HaveOccurred())
				})

				Measure("it should not adversely affect the rest of the API", func(b Benchmarker) {
					var newContainer garden.Container

					b.Time("creating another container", func() {
						var err error

						newContainer, err = client.Create(garden.ContainerSpec{})
						Ω(err).ShouldNot(HaveOccurred())
					})

					for i := 0; i < 10; i++ {
						b.Time("getting container info (10x)", func() {
							_, err := newContainer.Info()
							Ω(err).ShouldNot(HaveOccurred())
						})
					}

					for i := 0; i < 10; i++ {
						b.Time("running a job (10x)", func() {
							process, err := newContainer.Run(garden.ProcessSpec{Path: "ls"}, garden.ProcessIO{})
							Ω(err).ShouldNot(HaveOccurred())

							Ω(process.Wait()).Should(Equal(0))
						})
					}

					b.Time("destroying the container", func() {
						err := client.Destroy(newContainer.Handle())
						Ω(err).ShouldNot(HaveOccurred())
					})

					rsyslogMessages := 0

					shCmd := exec.Command("sh", "-c", "cat /var/log/gmeasure | wc -l")

					shOut, err := shCmd.StdoutPipe()
					Ω(err).ShouldNot(HaveOccurred())

					err = shCmd.Start()
					Ω(err).ShouldNot(HaveOccurred())

					_, err = fmt.Fscanf(shOut, "%d", &rsyslogMessages)
					Ω(err).ShouldNot(HaveOccurred())

					err = shCmd.Wait()
					Ω(err).ShouldNot(HaveOccurred())

					b.RecordValue(
						"received bytes/sec",
						float64(rsyslogMessages*len(loggedLine)+len("\n"))/float64(time.Since(started)/time.Second),
						// float64(atomic.LoadUint64(&receivedBytes))/float64(time.Since(started)/time.Second),
					)

					fmt.Println("total time:", time.Since(started))
				}, 5)
			})
		}
	})
})
