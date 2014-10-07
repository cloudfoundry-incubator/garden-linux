// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/cloudfoundry-incubator/garden/api"
)

type FakeBackend struct {
	PingStub        func() error
	pingMutex       sync.RWMutex
	pingArgsForCall []struct{}
	pingReturns struct {
		result1 error
	}
	CapacityStub        func() (api.Capacity, error)
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct{}
	capacityReturns struct {
		result1 api.Capacity
		result2 error
	}
	CreateStub        func(api.ContainerSpec) (api.Container, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 api.ContainerSpec
	}
	createReturns struct {
		result1 api.Container
		result2 error
	}
	DestroyStub        func(handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	ContainersStub        func(api.Properties) ([]api.Container, error)
	containersMutex       sync.RWMutex
	containersArgsForCall []struct {
		arg1 api.Properties
	}
	containersReturns struct {
		result1 []api.Container
		result2 error
	}
	LookupStub        func(handle string) (api.Container, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		handle string
	}
	lookupReturns struct {
		result1 api.Container
		result2 error
	}
	CreateVolumeStub        func(api.VolumeSpec) (api.Volume, error)
	createVolumeMutex       sync.RWMutex
	createVolumeArgsForCall []struct {
		arg1 api.VolumeSpec
	}
	createVolumeReturns struct {
		result1 api.Volume
		result2 error
	}
	DestroyVolumeStub        func(handle string) error
	destroyVolumeMutex       sync.RWMutex
	destroyVolumeArgsForCall []struct {
		handle string
	}
	destroyVolumeReturns struct {
		result1 error
	}
	LookupVolumeStub        func(handle string) (api.Volume, error)
	lookupVolumeMutex       sync.RWMutex
	lookupVolumeArgsForCall []struct {
		handle string
	}
	lookupVolumeReturns struct {
		result1 api.Volume
		result2 error
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns struct {
		result1 error
	}
	StopStub        func()
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	GraceTimeStub        func(api.Container) time.Duration
	graceTimeMutex       sync.RWMutex
	graceTimeArgsForCall []struct {
		arg1 api.Container
	}
	graceTimeReturns struct {
		result1 time.Duration
	}
}

func (fake *FakeBackend) Ping() error {
	fake.pingMutex.Lock()
	fake.pingArgsForCall = append(fake.pingArgsForCall, struct{}{})
	fake.pingMutex.Unlock()
	if fake.PingStub != nil {
		return fake.PingStub()
	} else {
		return fake.pingReturns.result1
	}
}

func (fake *FakeBackend) PingCallCount() int {
	fake.pingMutex.RLock()
	defer fake.pingMutex.RUnlock()
	return len(fake.pingArgsForCall)
}

func (fake *FakeBackend) PingReturns(result1 error) {
	fake.PingStub = nil
	fake.pingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Capacity() (api.Capacity, error) {
	fake.capacityMutex.Lock()
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	} else {
		return fake.capacityReturns.result1, fake.capacityReturns.result2
	}
}

func (fake *FakeBackend) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakeBackend) CapacityReturns(result1 api.Capacity, result2 error) {
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 api.Capacity
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Create(arg1 api.ContainerSpec) (api.Container, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 api.ContainerSpec
	}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeBackend) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBackend) CreateArgsForCall(i int) api.ContainerSpec {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeBackend) CreateReturns(result1 api.Container, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Destroy(handle string) error {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		handle string
	}{handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(handle)
	} else {
		return fake.destroyReturns.result1
	}
}

func (fake *FakeBackend) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeBackend) DestroyArgsForCall(i int) string {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].handle
}

func (fake *FakeBackend) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Containers(arg1 api.Properties) ([]api.Container, error) {
	fake.containersMutex.Lock()
	fake.containersArgsForCall = append(fake.containersArgsForCall, struct {
		arg1 api.Properties
	}{arg1})
	fake.containersMutex.Unlock()
	if fake.ContainersStub != nil {
		return fake.ContainersStub(arg1)
	} else {
		return fake.containersReturns.result1, fake.containersReturns.result2
	}
}

func (fake *FakeBackend) ContainersCallCount() int {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return len(fake.containersArgsForCall)
}

func (fake *FakeBackend) ContainersArgsForCall(i int) api.Properties {
	fake.containersMutex.RLock()
	defer fake.containersMutex.RUnlock()
	return fake.containersArgsForCall[i].arg1
}

func (fake *FakeBackend) ContainersReturns(result1 []api.Container, result2 error) {
	fake.ContainersStub = nil
	fake.containersReturns = struct {
		result1 []api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Lookup(handle string) (api.Container, error) {
	fake.lookupMutex.Lock()
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		handle string
	}{handle})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(handle)
	} else {
		return fake.lookupReturns.result1, fake.lookupReturns.result2
	}
}

func (fake *FakeBackend) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeBackend) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].handle
}

func (fake *FakeBackend) LookupReturns(result1 api.Container, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 api.Container
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) CreateVolume(arg1 api.VolumeSpec) (api.Volume, error) {
	fake.createVolumeMutex.Lock()
	fake.createVolumeArgsForCall = append(fake.createVolumeArgsForCall, struct {
		arg1 api.VolumeSpec
	}{arg1})
	fake.createVolumeMutex.Unlock()
	if fake.CreateVolumeStub != nil {
		return fake.CreateVolumeStub(arg1)
	} else {
		return fake.createVolumeReturns.result1, fake.createVolumeReturns.result2
	}
}

func (fake *FakeBackend) CreateVolumeCallCount() int {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return len(fake.createVolumeArgsForCall)
}

func (fake *FakeBackend) CreateVolumeArgsForCall(i int) api.VolumeSpec {
	fake.createVolumeMutex.RLock()
	defer fake.createVolumeMutex.RUnlock()
	return fake.createVolumeArgsForCall[i].arg1
}

func (fake *FakeBackend) CreateVolumeReturns(result1 api.Volume, result2 error) {
	fake.CreateVolumeStub = nil
	fake.createVolumeReturns = struct {
		result1 api.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) DestroyVolume(handle string) error {
	fake.destroyVolumeMutex.Lock()
	fake.destroyVolumeArgsForCall = append(fake.destroyVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.destroyVolumeMutex.Unlock()
	if fake.DestroyVolumeStub != nil {
		return fake.DestroyVolumeStub(handle)
	} else {
		return fake.destroyVolumeReturns.result1
	}
}

func (fake *FakeBackend) DestroyVolumeCallCount() int {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return len(fake.destroyVolumeArgsForCall)
}

func (fake *FakeBackend) DestroyVolumeArgsForCall(i int) string {
	fake.destroyVolumeMutex.RLock()
	defer fake.destroyVolumeMutex.RUnlock()
	return fake.destroyVolumeArgsForCall[i].handle
}

func (fake *FakeBackend) DestroyVolumeReturns(result1 error) {
	fake.DestroyVolumeStub = nil
	fake.destroyVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) LookupVolume(handle string) (api.Volume, error) {
	fake.lookupVolumeMutex.Lock()
	fake.lookupVolumeArgsForCall = append(fake.lookupVolumeArgsForCall, struct {
		handle string
	}{handle})
	fake.lookupVolumeMutex.Unlock()
	if fake.LookupVolumeStub != nil {
		return fake.LookupVolumeStub(handle)
	} else {
		return fake.lookupVolumeReturns.result1, fake.lookupVolumeReturns.result2
	}
}

func (fake *FakeBackend) LookupVolumeCallCount() int {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return len(fake.lookupVolumeArgsForCall)
}

func (fake *FakeBackend) LookupVolumeArgsForCall(i int) string {
	fake.lookupVolumeMutex.RLock()
	defer fake.lookupVolumeMutex.RUnlock()
	return fake.lookupVolumeArgsForCall[i].handle
}

func (fake *FakeBackend) LookupVolumeReturns(result1 api.Volume, result2 error) {
	fake.LookupVolumeStub = nil
	fake.lookupVolumeReturns = struct {
		result1 api.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeBackend) Start() error {
	fake.startMutex.Lock()
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	} else {
		return fake.startReturns.result1
	}
}

func (fake *FakeBackend) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeBackend) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackend) Stop() {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		fake.StopStub()
	}
}

func (fake *FakeBackend) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeBackend) GraceTime(arg1 api.Container) time.Duration {
	fake.graceTimeMutex.Lock()
	fake.graceTimeArgsForCall = append(fake.graceTimeArgsForCall, struct {
		arg1 api.Container
	}{arg1})
	fake.graceTimeMutex.Unlock()
	if fake.GraceTimeStub != nil {
		return fake.GraceTimeStub(arg1)
	} else {
		return fake.graceTimeReturns.result1
	}
}

func (fake *FakeBackend) GraceTimeCallCount() int {
	fake.graceTimeMutex.RLock()
	defer fake.graceTimeMutex.RUnlock()
	return len(fake.graceTimeArgsForCall)
}

func (fake *FakeBackend) GraceTimeArgsForCall(i int) api.Container {
	fake.graceTimeMutex.RLock()
	defer fake.graceTimeMutex.RUnlock()
	return fake.graceTimeArgsForCall[i].arg1
}

func (fake *FakeBackend) GraceTimeReturns(result1 time.Duration) {
	fake.GraceTimeStub = nil
	fake.graceTimeReturns = struct {
		result1 time.Duration
	}{result1}
}

var _ api.Backend = new(FakeBackend)
