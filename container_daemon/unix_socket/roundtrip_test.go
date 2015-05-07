package unix_socket_test

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/garden-linux/container_daemon/unix_socket"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon/unix_socket/fake_connection_handler"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unix socket", func() {
	var (
		listener          *unix_socket.Listener
		connector         *unix_socket.Connector
		connectionHandler *fake_connection_handler.FakeConnectionHandler
		socketPath        string
		sentPid           int
		sentError         error
		sentErrorMutex    sync.Mutex
	)

	BeforeEach(func() {
		tmpDir, err := ioutil.TempDir("", "")
		Expect(err).ToNot(HaveOccurred())
		socketPath = path.Join(tmpDir, "the_socket_file.sock")

		sentErrorMutex.Lock()
		defer sentErrorMutex.Unlock()
		sentError = nil
		sentPid = 0
		connectionHandler = &fake_connection_handler.FakeConnectionHandler{}
	})

	JustBeforeEach(func() {
		connector = &unix_socket.Connector{
			SocketPath: socketPath,
		}

		listener = &unix_socket.Listener{
			SocketPath: socketPath,
		}
	})

	Describe("Listener.Init", func() {
		It("creates a unix socket for the given socket path", func() {
			err := listener.Init()
			Expect(err).ToNot(HaveOccurred())

			stat, err := os.Stat(socketPath)
			Expect(err).ToNot(HaveOccurred())

			Expect(stat.Mode() & os.ModeSocket).ToNot(Equal(0))
		})

		Context("when the socket cannot be created", func() {
			BeforeEach(func() {
				socketPath = "somewhere/that/does/not/exist"
			})

			It("returns an error", func() {
				err := listener.Init()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	PDescribe("Listener.Stop", func() {
		It("should ", func() {

		})
	})

	Describe("Connect", func() {
		Context("when the server is not running", func() {
			It("fails to connect", func() {
				_, _, err := connector.Connect(nil)
				Expect(err).To(MatchError(ContainSubstring("unix_socket: connect to server socket")))
			})
		})

		Context("when the server is running", func() {
			var recvMsg map[string]string
			var sentFiles []*os.File
			var stubDone chan bool

			JustBeforeEach(func() {
				Expect(listener.Init()).To(Succeed())

				f1, _ := ioutil.TempFile("", "")
				f2, _ := ioutil.TempFile("", "")
				sentFiles = []*os.File{f1, f2}
				sentPid = 123

				stubDone = make(chan bool, 1)

				connectionHandler.HandleStub = func(decoder *json.Decoder) ([]*os.File, int, error) {
					defer GinkgoRecover()
					err := decoder.Decode(&recvMsg)
					Expect(err).ToNot(HaveOccurred())
					//stubDone <- true

					time.Sleep(time.Millisecond * 5)

					sentErrorMutex.Lock()
					defer sentErrorMutex.Unlock()
					return sentFiles, sentPid, sentError
				}

				go listener.Listen(connectionHandler)
			})

			AfterEach(func() {
				Expect(listener.Stop()).To(Succeed())
			})

			It("calls the handler with the sent message", func() {
				sentMsg := map[string]string{"fruit": "apple"}
				_, _, err := connector.Connect(sentMsg)
				Expect(err).ToNot(HaveOccurred())

				Eventually(stubDone).Should(Receive())
				Expect(recvMsg).To(Equal(sentMsg))
			})

			It("gets back the stream the handler provided", func() {
				sentMsg := map[string]string{"fruit": "apple"}
				streams, _, err := connector.Connect(sentMsg)
				Expect(err).ToNot(HaveOccurred())

				Expect(stubDone).To(Receive())
				Expect(streams).To(HaveLen(2))

				_, err = streams[0].Write([]byte("potato potato"))
				Expect(err).NotTo(HaveOccurred())
				sentFiles[0].Seek(0, 0)
				Expect(ioutil.ReadAll(sentFiles[0])).Should(Equal([]byte("potato potato")))

				_, err = sentFiles[1].Write([]byte("brocoli brocoli"))
				Expect(err).NotTo(HaveOccurred())
				sentFiles[1].Seek(0, 0)
				Expect(ioutil.ReadAll(streams[1])).Should(Equal([]byte("brocoli brocoli")))
			})

			FIt("gets back the pid the handler provided", func() {
				for i := 0; i < 1000; i++ {
					sentMsg := map[string]string{"fruit": "apple"}
					_, pid, err := connector.Connect(sentMsg)
					Expect(err).ToNot(HaveOccurred())

					//Expect(stubDone).To(Receive())
					Expect(pid).To(Equal(sentPid))
				}
			})

			Context("when the handler fails", func() {
				BeforeEach(func() {
					sentErrorMutex.Lock()
					defer sentErrorMutex.Unlock()
					sentError = errors.New("no cake")
				})

				It("sends back the error from the handler", func() {
					sentMsg := map[string]string{"fruit": "apple"}
					_, _, err := connector.Connect(sentMsg)
					Expect(err).To(MatchError("no cake"))
				})
			})
		})
	})

	Describe("Listener.Run", func() {
		Context("when the listener is not initialized", func() {
			It("returns an error", func() {
				err := listener.Listen(connectionHandler)
				Expect(err).To(MatchError("unix_socket: listener is not initialized"))
			})
		})
	})
})
