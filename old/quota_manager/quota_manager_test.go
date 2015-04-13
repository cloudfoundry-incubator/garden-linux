package quota_manager_test

import (
	"errors"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/pivotal-golang/lager/lagertest"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-linux/old/quota_manager"
	"github.com/cloudfoundry/gunk/command_runner/fake_command_runner"
	. "github.com/cloudfoundry/gunk/command_runner/fake_command_runner/matchers"
)

var _ = Describe("Linux Quota manager", func() {
	var fakeRunner *fake_command_runner.FakeCommandRunner
	var logger *lagertest.TestLogger
	var quotaManager *quota_manager.LinuxQuotaManager

	BeforeEach(func() {
		fakeRunner = fake_command_runner.New()
		logger = lagertest.NewTestLogger("test")
		quotaManager = quota_manager.New(fakeRunner, "/some/mount/point", "/root/path")
	})

	Describe("setting quotas", func() {
		limits := garden.DiskLimits{
			BlockSoft: 1,
			BlockHard: 2,

			InodeSoft: 11,
			InodeHard: 12,
		}

		It("executes setquota on the container depo's mount point", func() {
			err := quotaManager.SetLimits(logger, 1234, limits)

			Expect(err).ToNot(HaveOccurred())

			Expect(fakeRunner).To(HaveExecutedSerially(
				fake_command_runner.CommandSpec{
					Path: "setquota",
					Args: []string{
						"-u", "1234",
						"1", "2", "11", "12",
						"/some/mount/point",
					},
				},
			))
		})

		Context("when bytes are given", func() {
			limits := garden.DiskLimits{
				InodeSoft: 11,
				InodeHard: 12,

				ByteSoft: 102401,
				ByteHard: 204801,
			}

			It("executes setquota with them converted to blocks", func() {
				err := quotaManager.SetLimits(logger, 1234, limits)

				Expect(err).ToNot(HaveOccurred())

				Expect(fakeRunner).To(HaveExecutedSerially(
					fake_command_runner.CommandSpec{
						Path: "setquota",
						Args: []string{
							"-u", "1234",
							"101", "201", "11", "12",
							"/some/mount/point",
						},
					},
				))
			})
		})

		Context("when setquota fails", func() {
			nastyError := errors.New("oh no!")

			BeforeEach(func() {
				fakeRunner.WhenRunning(
					fake_command_runner.CommandSpec{
						Path: "setquota",
					}, func(*exec.Cmd) error {
						return nastyError
					},
				)
			})

			It("returns the error", func() {
				err := quotaManager.SetLimits(logger, 1234, limits)
				Expect(err).To(Equal(nastyError))
			})
		})

		Context("when quotas are disabled", func() {
			BeforeEach(func() {
				quotaManager.Disable()
			})

			It("runs nothing", func() {
				err := quotaManager.SetLimits(logger, 1234, limits)

				Expect(err).ToNot(HaveOccurred())

				Expect(fakeRunner).ToNot(HaveExecutedSerially(
					fake_command_runner.CommandSpec{
						Path: "setquota",
					},
				))
			})
		})
	})

	Describe("getting quotas limits", func() {
		It("executes repquota in the root path", func() {
			fakeRunner.WhenRunning(
				fake_command_runner.CommandSpec{
					Path: "/root/path/repquota",
					Args: []string{"/some/mount/point", "1234"},
				}, func(cmd *exec.Cmd) error {
					cmd.Stdout.Write([]byte("1234 111 222 333 444 555 666 777 888\n"))

					return nil
				},
			)

			limits, err := quotaManager.GetLimits(logger, 1234)
			Expect(err).ToNot(HaveOccurred())

			Expect(limits.BlockSoft).To(Equal(uint64(222)))
			Expect(limits.BlockHard).To(Equal(uint64(333)))

			Expect(limits.InodeSoft).To(Equal(uint64(666)))
			Expect(limits.InodeHard).To(Equal(uint64(777)))
		})

		Context("when repquota fails", func() {
			disaster := errors.New("oh no!")

			BeforeEach(func() {
				fakeRunner.WhenRunning(
					fake_command_runner.CommandSpec{
						Path: "/root/path/repquota",
						Args: []string{"/some/mount/point", "1234"},
					}, func(cmd *exec.Cmd) error {
						return disaster
					},
				)
			})

			It("returns the error", func() {
				_, err := quotaManager.GetLimits(logger, 1234)
				Expect(err).To(Equal(disaster))
			})
		})

		Context("when the output of repquota is malformed", func() {
			It("returns an error", func() {
				fakeRunner.WhenRunning(
					fake_command_runner.CommandSpec{
						Path: "/root/path/repquota",
						Args: []string{"/some/mount/point", "1234"},
					}, func(cmd *exec.Cmd) error {
						cmd.Stdout.Write([]byte("abc\n"))

						return nil
					},
				)

				_, err := quotaManager.GetLimits(logger, 1234)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when quotas are disabled", func() {
			BeforeEach(func() {
				quotaManager.Disable()
			})

			It("runs nothing", func() {
				limits, err := quotaManager.GetLimits(logger, 1234)
				Expect(err).ToNot(HaveOccurred())

				Expect(limits).To(BeZero())

				Expect(fakeRunner).ToNot(HaveExecutedSerially(
					fake_command_runner.CommandSpec{
						Path: "/root/path/bin/repquota",
					},
				))
			})
		})
	})

	Describe("getting usage", func() {
		It("executes repquota in the root path", func() {
			fakeRunner.WhenRunning(
				fake_command_runner.CommandSpec{
					Path: "/root/path/repquota",
					Args: []string{"/some/mount/point", "1234"},
				}, func(cmd *exec.Cmd) error {
					cmd.Stdout.Write([]byte("1234 111 222 333 444 555 666 777 888\n"))

					return nil
				},
			)

			limits, err := quotaManager.GetUsage(logger, 1234)
			Expect(err).ToNot(HaveOccurred())

			Expect(limits.BytesUsed).To(Equal(uint64(111)))
			Expect(limits.InodesUsed).To(Equal(uint64(555)))
		})

		Context("when repquota fails", func() {
			disaster := errors.New("oh no!")

			BeforeEach(func() {
				fakeRunner.WhenRunning(
					fake_command_runner.CommandSpec{
						Path: "/root/path/repquota",
						Args: []string{"/some/mount/point", "1234"},
					}, func(cmd *exec.Cmd) error {
						return disaster
					},
				)
			})

			It("returns the error", func() {
				_, err := quotaManager.GetUsage(logger, 1234)
				Expect(err).To(Equal(disaster))
			})
		})

		Context("when the output of repquota is malformed", func() {
			It("returns an error", func() {
				fakeRunner.WhenRunning(
					fake_command_runner.CommandSpec{
						Path: "/root/path/repquota",
						Args: []string{"/some/mount/point", "1234"},
					}, func(cmd *exec.Cmd) error {
						cmd.Stdout.Write([]byte("abc\n"))

						return nil
					},
				)

				_, err := quotaManager.GetUsage(logger, 1234)
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when quotas are disabled", func() {
			BeforeEach(func() {
				quotaManager.Disable()
			})

			It("runs nothing", func() {
				usage, err := quotaManager.GetUsage(logger, 1234)
				Expect(err).ToNot(HaveOccurred())

				Expect(usage).To(BeZero())

				Expect(fakeRunner).ToNot(HaveExecutedSerially(
					fake_command_runner.CommandSpec{
						Path: "/root/path/repquota",
					},
				))
			})
		})
	})

	Describe("getting the mount point", func() {
		It("returns the mount point of the container depot", func() {
			Expect(quotaManager.MountPoint()).To(Equal("/some/mount/point"))
		})
	})
})
