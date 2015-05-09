package container_daemon_test

import (
	"fmt"

	"github.com/cloudfoundry-incubator/garden"
	"github.com/cloudfoundry-incubator/garden-linux/container_daemon"

	"os/exec"
	"os/user"

	"github.com/cloudfoundry-incubator/garden-linux/containerizer/system/fake_user"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Preparing a command to run", func() {
	var (
		users    *fake_user.FakeUser
		preparer *container_daemon.ProcessSpecPreparer
	)

	etcPasswd := map[string]*user.User{
		"a-user":       &user.User{Uid: "66", Gid: "99"},
		"another-user": &user.User{Uid: "77", Gid: "88", HomeDir: "/the/home/dir"},
		"a-root-user":  &user.User{},
	}

	BeforeEach(func() {
		users = new(fake_user.FakeUser)

		users.LookupStub = func(name string) (*user.User, error) {
			return etcPasswd[name], nil
		}

		preparer = &container_daemon.ProcessSpecPreparer{
			Users: users,
		}
	})

	Describe("Spawn", func() {
		var spec garden.ProcessSpec

		BeforeEach(func() {
			spec = garden.ProcessSpec{
				User: "another-user",
				Dir:  "some-dir",
				Path: "fishfinger",
				Args: []string{
					"foo", "bar",
				},
				Env: []string{
					"foo=bar",
					"baz=barry",
				},
			}
		})

		Describe("the prepared process", func() {
			var thePreparedCmd *exec.Cmd
			var theReturnedError error

			JustBeforeEach(func() {
				thePreparedCmd, theReturnedError = preparer.PrepareCmd(spec)
			})

			It("has the correct path and args", func() {
				Expect(theReturnedError).To(BeNil())
				Expect(thePreparedCmd.Path).To(Equal("fishfinger"))
				Expect(thePreparedCmd.Args).To(Equal([]string{"fishfinger", "foo", "bar"}))
			})

			It("has the correct uid based on the /etc/passwd file", func() {
				Expect(thePreparedCmd.SysProcAttr).ToNot(BeNil())
				Expect(thePreparedCmd.SysProcAttr.Credential).ToNot(BeNil())
				Expect(thePreparedCmd.SysProcAttr.Credential.Uid).To(Equal(uint32(77)))
				Expect(thePreparedCmd.SysProcAttr.Credential.Gid).To(Equal(uint32(88)))
			})

			Context("when the process spec names a user which does not exist in /etc/passwd", func() {
				BeforeEach(func() {
					spec.User = "not-a-user"
				})

				It("returns an informative error", func() {
					Expect(theReturnedError).To(MatchError("container_daemon: failed to lookup user not-a-user"))
				})
			})

			It("has the supplied env vars", func() {
				Expect(thePreparedCmd.Env).To(ContainElement("foo=bar"))
				Expect(thePreparedCmd.Env).To(ContainElement("baz=barry"))
			})

			It("sets the USER environment variable", func() {
				Expect(thePreparedCmd.Env).To(ContainElement("USER=another-user"))
			})

			It("sets the HOME environment variable to the home dir in /etc/passwd", func() {
				Expect(thePreparedCmd.Env).To(ContainElement("HOME=/the/home/dir"))
			})

			Context("when the ENV does not contain a PATH", func() {
				Context("and the uid is not 0", func() {
					It("appends the DefaultUserPATH to the environment", func() {
						Expect(thePreparedCmd.Env).To(ContainElement(fmt.Sprintf("PATH=%s", container_daemon.DefaultUserPath)))
					})
				})

				Context("and the uid is 0", func() {
					BeforeEach(func() {
						spec.User = "a-root-user"
					})

					It("appends the DefaultRootPATH to the environment", func() {
						Expect(thePreparedCmd.Env).To(ContainElement(fmt.Sprintf("PATH=%s", container_daemon.DefaultRootPATH)))
					})
				})

				Context("when the ENV already contains a PATH", func() {
					BeforeEach(func() {
						spec.Env = []string{"PATH=cake"}
					})

					It("is not overridden", func() {
						Expect(thePreparedCmd.Env).To(ContainElement("PATH=cake"))
						Expect(thePreparedCmd.Env).NotTo(ContainElement(fmt.Sprintf("PATH=%s", container_daemon.DefaultUserPath)))
					})
				})
			})

			It("has the supplied dir", func() {
				Expect(thePreparedCmd.Dir).To(Equal("some-dir"))
			})
		})
	})
})
