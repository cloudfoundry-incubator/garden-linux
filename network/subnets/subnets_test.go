package subnets_test

import (
	"net"
	"runtime"

	"code.cloudfoundry.org/garden-linux/linux_backend"
	"code.cloudfoundry.org/garden-linux/network/subnets"
	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/lager/lagertest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Subnet Pool", func() {
	var subnetpool subnets.Subnets
	var defaultSubnetPool *net.IPNet
	var logger lager.Logger

	JustBeforeEach(func() {
		var err error
		logger = lagertest.NewTestLogger("test")
		subnetpool, err = subnets.NewSubnets(defaultSubnetPool)
		Expect(err).ToNot(HaveOccurred())
	})

	Describe("Capacity", func() {
		Context("when the dynamic allocation net is empty", func() {
			BeforeEach(func() {
				defaultSubnetPool = subnetPool("10.2.3.0/32")
			})

			It("returns zero", func() {
				Expect(subnetpool.Capacity()).To(Equal(0))
			})
		})

		Context("when the dynamic allocation net is non-empty", func() {
			BeforeEach(func() {
				defaultSubnetPool = subnetPool("10.2.3.0/27")
			})

			It("returns the correct number of subnets initially and repeatedly", func() {
				Expect(subnetpool.Capacity()).To(Equal(8))
				Expect(subnetpool.Capacity()).To(Equal(8))
			})

			It("returns the correct capacity after allocating subnets", func() {
				cap := subnetpool.Capacity()

				_, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
				Expect(err).ToNot(HaveOccurred())

				Expect(subnetpool.Capacity()).To(Equal(cap))

				_, err = subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
				Expect(err).ToNot(HaveOccurred())

				Expect(subnetpool.Capacity()).To(Equal(cap))
			})
		})
	})

	Describe("Allocating and Releasing", func() {
		Describe("Static Subnet Allocation", func() {
			Context("when the requested subnet is within the dynamic allocation range", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.0/29")
				})

				It("returns an appropriate error", func() {
					_, static := networkParms("10.2.3.4/30")

					_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
					Expect(err).To(MatchError("the requested subnet (10.2.3.4/30) overlaps the dynamic allocation range (10.2.3.0/29)"))
				})
			})

			Context("when the requested subnet subsumes the dynamic allocation range", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.4/30")
				})

				It("returns an appropriate error", func() {
					_, static := networkParms("10.2.3.0/24")

					_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
					Expect(err).To(HaveOccurred())
					Expect(err).To(MatchError("the requested subnet (10.2.3.0/24) overlaps the dynamic allocation range (10.2.3.4/30)"))
				})
			})

			Context("when the requested subnet is not within the dynamic allocation range", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.0/29")
				})

				Context("allocating a static subnet", func() {
					Context("and a static IP", func() {
						It("returns an error if the IP is not inside the subnet", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("9.0.0.1")
							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).To(Equal(subnets.ErrInvalidIP))
						})

						It("returns the same subnet and IP if the IP is inside the subnet", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("11.0.0.2")
							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())

							Expect(network.Subnet).To(Equal(static))
							Expect(network.IP).To(Equal(ip))
						})

						It("does not allow the same IP to be requested twice", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("11.0.0.2")
							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())

							_, static = networkParms("11.0.0.0/8") // make sure we get a new pointer
							_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).To(Equal(subnets.ErrIPAlreadyAcquired))
						})

						It("allows two IPs to be serially requested in the same subnet", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("11.0.0.2")
							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.Subnet).To(Equal(static))
							Expect(network.IP).To(Equal(ip))

							ip2 := net.ParseIP("11.0.0.3")

							_, static = networkParms("11.0.0.0/8") // make sure we get a new pointer
							network2, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip2}, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network2.Subnet).To(Equal(static))
							Expect(network2.IP).To(Equal(ip2))
						})

						It("when an IP is allocated from a subnet but released in between, it should be treated as new both times", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("11.0.0.2")
							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.Subnet).To(Equal(static))
							Expect(network.IP).To(Equal(ip))

							err = subnetpool.Release(network, logger)
							Expect(err).ToNot(HaveOccurred())

							_, static = networkParms("11.0.0.0/8") // make sure we get a new pointer
							network, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.Subnet).To(Equal(static))
							Expect(network.IP).To(Equal(ip))
						})

						It("prevents dynamic allocation of the same IP", func() {
							_, static := networkParms("11.0.0.0/8")

							ip := net.ParseIP("11.0.0.3")
							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
							Expect(err).ToNot(HaveOccurred())

							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.IP.String()).To(Equal("11.0.0.2"))

							network, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.IP.String()).To(Equal("11.0.0.4"))
						})

						Describe("errors", func() {
							It("fails if a static subnet is requested specifying an IP address which clashes with the gateway IP address", func() {
								_, static := networkParms("11.0.0.0/8")
								gateway := net.ParseIP("11.0.0.1")
								_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{gateway}, logger)
								Expect(err).To(MatchError(subnets.ErrIPEqualsGateway))
							})

							It("fails if a static subnet is requested specifying an IP address which clashes with the broadcast IP address", func() {
								_, static := networkParms("11.0.0.0/8")
								max := net.ParseIP("11.255.255.255")
								_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{max}, logger)
								Expect(err).To(MatchError(subnets.ErrIPEqualsBroadcast))
							})
						})
					})

					Context("and a dynamic IP", func() {
						It("does not return an error", func() {
							_, static := networkParms("11.0.0.0/8")

							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
						})

						It("returns the first available IP", func() {
							_, static := networkParms("11.0.0.0/8")

							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())

							Expect(network.IP.String()).To(Equal("11.0.0.2"))
						})

						It("returns distinct IPs", func() {
							_, static := networkParms("11.0.0.0/22")

							seen := make(map[string]bool)
							var err error
							for err == nil {
								network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)

								if err != nil {
									Expect(err).To(Equal(subnets.ErrInsufficientIPs))
									break
								}

								Expect(seen).ToNot(HaveKey(network.IP.String()))
								seen[network.IP.String()] = true
							}
						})

						It("returns all IPs except gateway, minimum and broadcast", func() {
							_, static := networkParms("11.0.0.0/23")

							var err error
							count := 0
							for err == nil {
								if _, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger); err != nil {
									Expect(err).To(Equal(subnets.ErrInsufficientIPs))
								}

								count++
							}

							Expect(count).To(Equal(510))
						})

						It("causes static alocation to fail if it tries to allocate the same IP afterwards", func() {
							_, static := networkParms("11.0.0.0/8")

							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())

							_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{network.IP}, logger)
							Expect(err).To(Equal(subnets.ErrIPAlreadyAcquired))
						})
					})
				})

				Context("after all IPs are allocated from a subnet, a subsequent request for the same subnet", func() {
					var (
						static *net.IPNet
						ips    [5]net.IP
					)

					JustBeforeEach(func() {
						var err error
						_, static, err = net.ParseCIDR("10.9.3.0/29")
						Expect(err).ToNot(HaveOccurred())

						for i := 0; i < 5; i++ {
							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())

							ips[i] = network.IP
						}
					})

					It("returns an appropriate error", func() {
						_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
						Expect(err).To(HaveOccurred())
						Expect(err).To(Equal(subnets.ErrInsufficientIPs))
					})

					Context("but after it is released", func() {
						It("dynamically allocates the released IP again", func() {
							err := subnetpool.Release(&linux_backend.Network{static, ips[3]}, logger)
							Expect(err).ToNot(HaveOccurred())

							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
							Expect(network.IP).To(Equal(ips[3]))
						})

						It("allows static allocation again", func() {
							err := subnetpool.Release(&linux_backend.Network{static, ips[3]}, logger)
							Expect(err).ToNot(HaveOccurred())

							_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ips[3]}, logger)
							Expect(err).ToNot(HaveOccurred())
						})
					})
				})

				Context("after a subnet has been allocated, a subsequent request for an overlapping subnet which begins on the same ip", func() {
					var (
						firstSubnetPool  *net.IPNet
						secondSubnetPool *net.IPNet
					)

					JustBeforeEach(func() {
						_, firstSubnetPool = networkParms("10.9.3.0/30")
						_, secondSubnetPool = networkParms("10.9.3.0/29")

						_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{firstSubnetPool}, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())
					})

					It("returns an appropriate error", func() {
						_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{secondSubnetPool}, subnets.DynamicIPSelector, logger)
						Expect(err).To(MatchError("the requested subnet (10.9.3.0/29) overlaps an existing subnet (10.9.3.0/30)"))
					})
				})

				Context("after a subnet has been allocated, a subsequent request for an overlapping subnet", func() {
					var (
						firstSubnetPool  *net.IPNet
						firstContainerIP net.IP
						secondSubnetPool *net.IPNet
					)

					JustBeforeEach(func() {
						var err error
						firstContainerIP, firstSubnetPool = networkParms("10.9.3.4/30")
						Expect(err).ToNot(HaveOccurred())

						_, secondSubnetPool = networkParms("10.9.3.0/29")
						Expect(err).ToNot(HaveOccurred())

						_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{firstSubnetPool}, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())
					})

					It("returns an appropriate error", func() {
						_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{secondSubnetPool}, subnets.DynamicIPSelector, logger)
						Expect(err).To(MatchError("the requested subnet (10.9.3.0/29) overlaps an existing subnet (10.9.3.4/30)"))
					})

					Context("but after it is released", func() {
						It("allows allocation again", func() {
							err := subnetpool.Release(&linux_backend.Network{firstSubnetPool, firstContainerIP}, logger)
							Expect(err).ToNot(HaveOccurred())

							_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{secondSubnetPool}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
						})
					})
				})

				Context("requesting a specific IP address in a static subnet", func() {
					It("does not return an error", func() {
						_, static := networkParms("10.9.3.6/29")

						_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())
					})
				})

			})
		})

		Describe("Dynamic /30 Subnet Allocation", func() {
			Context("when the pool does not have sufficient IPs to allocate a subnet", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.0/31")
				})

				It("the first request returns an error", func() {
					_, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("when the pool has sufficient IPs to allocate a single subnet", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.0/30")
				})

				Context("the first request", func() {
					It("succeeds, and returns a /30 network within the subnet", func() {
						network, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						Expect(network.Subnet).ToNot(BeNil())
						Expect(network.Subnet.String()).To(Equal("10.2.3.0/30"))
					})
				})

				Context("subsequent requests", func() {
					It("fails, and return an err", func() {
						_, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						_, err = subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).To(HaveOccurred())
					})
				})

				Context("when an allocated network is released", func() {
					It("a subsequent allocation succeeds, and returns the first network again", func() {
						// first
						network, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						// second - will fail (sanity check)
						_, err = subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).To(HaveOccurred())

						// release
						err = subnetpool.Release(&linux_backend.Network{network.Subnet, network.IP}, logger)
						Expect(err).ToNot(HaveOccurred())

						// third - should work now because of release
						network2, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						Expect(network2.Subnet).ToNot(BeNil())
						Expect(network2.Subnet.String()).To(Equal(network.Subnet.String()))
					})

					Context("and it is not the last IP in the subnet", func() {
						It("returns gone=false", func() {
							_, static := networkParms("10.3.3.0/29")

							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())

							network, err := subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())

							err = subnetpool.Release(&linux_backend.Network{network.Subnet, network.IP}, logger)
							Expect(err).ToNot(HaveOccurred())
						})
					})
				})

				Context("when a network is released twice", func() {
					It("returns an error", func() {
						// first
						network, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						// release
						err = subnetpool.Release(&linux_backend.Network{network.Subnet, network.IP}, logger)
						Expect(err).ToNot(HaveOccurred())

						// release again
						err = subnetpool.Release(&linux_backend.Network{network.Subnet, network.IP}, logger)
						Expect(err).To(HaveOccurred())
						Expect(err).To(Equal(subnets.ErrReleasedUnallocatedSubnet))
					})
				})
			})

			Context("when the pool has sufficient IPs to allocate two /30 subnets", func() {
				BeforeEach(func() {
					defaultSubnetPool = subnetPool("10.2.3.0/29")
				})

				Context("the second request", func() {
					It("succeeds", func() {
						_, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						_, err = subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())
					})

					It("returns the second /30 network within the subnet", func() {
						_, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						network, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						Expect(network.Subnet).ToNot(BeNil())
						Expect(network.Subnet.String()).To(Equal("10.2.3.4/30"))
					})
				})

				It("allocates distinct networks concurrently", func() {
					prev := runtime.GOMAXPROCS(2)
					defer runtime.GOMAXPROCS(prev)

					Consistently(func() bool {
						_, network, err := net.ParseCIDR("10.0.0.0/29")
						Expect(err).ToNot(HaveOccurred())

						subnetpool, err := subnets.NewSubnets(network)
						Expect(err).ToNot(HaveOccurred())

						out := make(chan *net.IPNet)
						go func(out chan *net.IPNet) {
							defer GinkgoRecover()
							n1, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
							out <- n1.Subnet
						}(out)

						go func(out chan *net.IPNet) {
							defer GinkgoRecover()
							n1, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
							Expect(err).ToNot(HaveOccurred())
							out <- n1.Subnet
						}(out)

						a := <-out
						b := <-out
						return a.IP.Equal(b.IP)
					}, "100ms", "2ms").ShouldNot(BeTrue())
				})

				It("correctly handles concurrent release of the same network", func() {
					prev := runtime.GOMAXPROCS(2)
					defer runtime.GOMAXPROCS(prev)

					Consistently(func() bool {
						_, network, err := net.ParseCIDR("10.0.0.0/29")
						Expect(err).ToNot(HaveOccurred())

						subnetpool, err := subnets.NewSubnets(network)
						Expect(err).ToNot(HaveOccurred())

						acquired, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.DynamicIPSelector, logger)
						Expect(err).ToNot(HaveOccurred())

						out := make(chan error)
						go func(out chan error) {
							defer GinkgoRecover()
							err := subnetpool.Release(&linux_backend.Network{acquired.Subnet, acquired.IP}, logger)
							out <- err
						}(out)

						go func(out chan error) {
							defer GinkgoRecover()
							err := subnetpool.Release(&linux_backend.Network{acquired.Subnet, acquired.IP}, logger)
							out <- err
						}(out)

						a := <-out
						b := <-out
						return (a == nil) != (b == nil)
					}, "200ms", "2ms").Should(BeTrue())
				})

				It("correctly handles concurrent allocation of the same network", func() {
					prev := runtime.GOMAXPROCS(2)
					defer runtime.GOMAXPROCS(prev)

					Consistently(func() bool {
						network := subnetPool("10.0.0.0/29")

						subnetpool, err := subnets.NewSubnets(network)
						Expect(err).ToNot(HaveOccurred())

						ip, n1 := networkParms("10.1.0.0/30")

						out := make(chan error)
						go func(out chan error) {
							defer GinkgoRecover()
							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{n1}, subnets.StaticIPSelector{ip}, logger)
							out <- err
						}(out)

						go func(out chan error) {
							defer GinkgoRecover()
							_, err := subnetpool.Acquire(subnets.StaticSubnetSelector{n1}, subnets.StaticIPSelector{ip}, logger)
							out <- err
						}(out)

						a := <-out
						b := <-out
						return (a == nil) != (b == nil)
					}, "200ms", "2ms").Should(BeTrue())
				})
			})
		})

		Describe("Removeing", func() {
			BeforeEach(func() {
				defaultSubnetPool = subnetPool("10.2.3.0/29")
			})

			Context("an allocation outside the dynamic allocation net", func() {
				It("recovers the first time", func() {
					_, static := networkParms("10.9.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.9.3.5")}, logger)
					Expect(err).ToNot(HaveOccurred())
				})

				It("does not allow recovering twice", func() {
					_, static := networkParms("10.9.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.9.3.5")}, logger)
					Expect(err).ToNot(HaveOccurred())

					err = subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.9.3.5")}, logger)
					Expect(err).To(HaveOccurred())
				})

				It("does not allow allocating after recovery", func() {
					_, static := networkParms("10.9.3.4/30")

					ip := net.ParseIP("10.9.3.5")
					err := subnetpool.Remove(&linux_backend.Network{static, ip}, logger)
					Expect(err).ToNot(HaveOccurred())

					_, err = subnetpool.Acquire(subnets.StaticSubnetSelector{static}, subnets.StaticIPSelector{ip}, logger)
					Expect(err).To(HaveOccurred())
				})

				It("does not allow recovering without an explicit IP", func() {
					_, static := networkParms("10.9.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, nil}, logger)
					Expect(err).To(HaveOccurred())
				})
			})

			Context("an allocation inside the dynamic allocation net", func() {
				It("recovers the first time", func() {
					_, static := networkParms("10.2.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.2.3.5")}, logger)
					Expect(err).ToNot(HaveOccurred())
				})

				It("does not allow recovering twice", func() {
					_, static := networkParms("10.2.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.2.3.5")}, logger)
					Expect(err).ToNot(HaveOccurred())

					err = subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.2.3.5")}, logger)
					Expect(err).To(HaveOccurred())
				})

				It("does not dynamically allocate a recovered network", func() {
					_, static := networkParms("10.2.3.4/30")

					err := subnetpool.Remove(&linux_backend.Network{static, net.ParseIP("10.2.3.1")}, logger)
					Expect(err).ToNot(HaveOccurred())

					network, err := subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.StaticIPSelector{net.ParseIP("10.2.3.2")}, logger)
					Expect(err).ToNot(HaveOccurred())
					Expect(network.Subnet.String()).To(Equal("10.2.3.0/30"))

					_, err = subnetpool.Acquire(subnets.DynamicSubnetSelector, subnets.StaticIPSelector{net.ParseIP("10.2.3.2")}, logger)
					Expect(err).To(Equal(subnets.ErrInsufficientSubnets))
				})
			})

		})

	})
})

func subnetPool(networkString string) *net.IPNet {
	_, subnetPool := networkParms(networkString)
	return subnetPool
}

func networkParms(networkString string) (net.IP, *net.IPNet) {
	containerIP, subnet, err := net.ParseCIDR(networkString)
	Expect(err).ToNot(HaveOccurred())
	gatewayIP := nextIP(subnet.IP)

	if containerIP.Equal(subnet.IP) {
		containerIP = nextIP(containerIP)
	}
	if containerIP.Equal(gatewayIP) {
		containerIP = nextIP(containerIP)
	}

	return containerIP, subnet
}

func nextIP(ip net.IP) net.IP {
	next := net.ParseIP(ip.String())
	inc(next)
	return next
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
