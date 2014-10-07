// This file was generated by counterfeiter
package fakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry-incubator/garden/api"
)

type FakeContainer struct {
	HandleStub        func() string
	handleMutex       sync.RWMutex
	handleArgsForCall []struct{}
	handleReturns struct {
		result1 string
	}
	StopStub        func(kill bool) error
	stopMutex       sync.RWMutex
	stopArgsForCall []struct {
		kill bool
	}
	stopReturns struct {
		result1 error
	}
	InfoStub        func() (api.ContainerInfo, error)
	infoMutex       sync.RWMutex
	infoArgsForCall []struct{}
	infoReturns struct {
		result1 api.ContainerInfo
		result2 error
	}
	StreamInStub        func(dstPath string, tarStream io.Reader) error
	streamInMutex       sync.RWMutex
	streamInArgsForCall []struct {
		dstPath   string
		tarStream io.Reader
	}
	streamInReturns struct {
		result1 error
	}
	StreamOutStub        func(srcPath string) (io.ReadCloser, error)
	streamOutMutex       sync.RWMutex
	streamOutArgsForCall []struct {
		srcPath string
	}
	streamOutReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	LimitBandwidthStub        func(limits api.BandwidthLimits) error
	limitBandwidthMutex       sync.RWMutex
	limitBandwidthArgsForCall []struct {
		limits api.BandwidthLimits
	}
	limitBandwidthReturns struct {
		result1 error
	}
	CurrentBandwidthLimitsStub        func() (api.BandwidthLimits, error)
	currentBandwidthLimitsMutex       sync.RWMutex
	currentBandwidthLimitsArgsForCall []struct{}
	currentBandwidthLimitsReturns struct {
		result1 api.BandwidthLimits
		result2 error
	}
	LimitCPUStub        func(limits api.CPULimits) error
	limitCPUMutex       sync.RWMutex
	limitCPUArgsForCall []struct {
		limits api.CPULimits
	}
	limitCPUReturns struct {
		result1 error
	}
	CurrentCPULimitsStub        func() (api.CPULimits, error)
	currentCPULimitsMutex       sync.RWMutex
	currentCPULimitsArgsForCall []struct{}
	currentCPULimitsReturns struct {
		result1 api.CPULimits
		result2 error
	}
	LimitDiskStub        func(limits api.DiskLimits) error
	limitDiskMutex       sync.RWMutex
	limitDiskArgsForCall []struct {
		limits api.DiskLimits
	}
	limitDiskReturns struct {
		result1 error
	}
	CurrentDiskLimitsStub        func() (api.DiskLimits, error)
	currentDiskLimitsMutex       sync.RWMutex
	currentDiskLimitsArgsForCall []struct{}
	currentDiskLimitsReturns struct {
		result1 api.DiskLimits
		result2 error
	}
	LimitMemoryStub        func(limits api.MemoryLimits) error
	limitMemoryMutex       sync.RWMutex
	limitMemoryArgsForCall []struct {
		limits api.MemoryLimits
	}
	limitMemoryReturns struct {
		result1 error
	}
	CurrentMemoryLimitsStub        func() (api.MemoryLimits, error)
	currentMemoryLimitsMutex       sync.RWMutex
	currentMemoryLimitsArgsForCall []struct{}
	currentMemoryLimitsReturns struct {
		result1 api.MemoryLimits
		result2 error
	}
	NetInStub        func(hostPort, containerPort uint32) (uint32, uint32, error)
	netInMutex       sync.RWMutex
	netInArgsForCall []struct {
		hostPort      uint32
		containerPort uint32
	}
	netInReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	NetOutStub        func(network string, port uint32) error
	netOutMutex       sync.RWMutex
	netOutArgsForCall []struct {
		network string
		port    uint32
	}
	netOutReturns struct {
		result1 error
	}
	RunStub        func(api.ProcessSpec, api.ProcessIO) (api.Process, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 api.ProcessSpec
		arg2 api.ProcessIO
	}
	runReturns struct {
		result1 api.Process
		result2 error
	}
	AttachStub        func(uint32, api.ProcessIO) (api.Process, error)
	attachMutex       sync.RWMutex
	attachArgsForCall []struct {
		arg1 uint32
		arg2 api.ProcessIO
	}
	attachReturns struct {
		result1 api.Process
		result2 error
	}
	BindVolumeStub        func(api.Volume, api.VolumeBinding) error
	bindVolumeMutex       sync.RWMutex
	bindVolumeArgsForCall []struct {
		arg1 api.Volume
		arg2 api.VolumeBinding
	}
	bindVolumeReturns struct {
		result1 error
	}
	UnbindVolumeStub        func(api.Volume) error
	unbindVolumeMutex       sync.RWMutex
	unbindVolumeArgsForCall []struct {
		arg1 api.Volume
	}
	unbindVolumeReturns struct {
		result1 error
	}
}

func (fake *FakeContainer) Handle() string {
	fake.handleMutex.Lock()
	fake.handleArgsForCall = append(fake.handleArgsForCall, struct{}{})
	fake.handleMutex.Unlock()
	if fake.HandleStub != nil {
		return fake.HandleStub()
	} else {
		return fake.handleReturns.result1
	}
}

func (fake *FakeContainer) HandleCallCount() int {
	fake.handleMutex.RLock()
	defer fake.handleMutex.RUnlock()
	return len(fake.handleArgsForCall)
}

func (fake *FakeContainer) HandleReturns(result1 string) {
	fake.HandleStub = nil
	fake.handleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeContainer) Stop(kill bool) error {
	fake.stopMutex.Lock()
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct {
		kill bool
	}{kill})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub(kill)
	} else {
		return fake.stopReturns.result1
	}
}

func (fake *FakeContainer) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeContainer) StopArgsForCall(i int) bool {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return fake.stopArgsForCall[i].kill
}

func (fake *FakeContainer) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Info() (api.ContainerInfo, error) {
	fake.infoMutex.Lock()
	fake.infoArgsForCall = append(fake.infoArgsForCall, struct{}{})
	fake.infoMutex.Unlock()
	if fake.InfoStub != nil {
		return fake.InfoStub()
	} else {
		return fake.infoReturns.result1, fake.infoReturns.result2
	}
}

func (fake *FakeContainer) InfoCallCount() int {
	fake.infoMutex.RLock()
	defer fake.infoMutex.RUnlock()
	return len(fake.infoArgsForCall)
}

func (fake *FakeContainer) InfoReturns(result1 api.ContainerInfo, result2 error) {
	fake.InfoStub = nil
	fake.infoReturns = struct {
		result1 api.ContainerInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) StreamIn(dstPath string, tarStream io.Reader) error {
	fake.streamInMutex.Lock()
	fake.streamInArgsForCall = append(fake.streamInArgsForCall, struct {
		dstPath   string
		tarStream io.Reader
	}{dstPath, tarStream})
	fake.streamInMutex.Unlock()
	if fake.StreamInStub != nil {
		return fake.StreamInStub(dstPath, tarStream)
	} else {
		return fake.streamInReturns.result1
	}
}

func (fake *FakeContainer) StreamInCallCount() int {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return len(fake.streamInArgsForCall)
}

func (fake *FakeContainer) StreamInArgsForCall(i int) (string, io.Reader) {
	fake.streamInMutex.RLock()
	defer fake.streamInMutex.RUnlock()
	return fake.streamInArgsForCall[i].dstPath, fake.streamInArgsForCall[i].tarStream
}

func (fake *FakeContainer) StreamInReturns(result1 error) {
	fake.StreamInStub = nil
	fake.streamInReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) StreamOut(srcPath string) (io.ReadCloser, error) {
	fake.streamOutMutex.Lock()
	fake.streamOutArgsForCall = append(fake.streamOutArgsForCall, struct {
		srcPath string
	}{srcPath})
	fake.streamOutMutex.Unlock()
	if fake.StreamOutStub != nil {
		return fake.StreamOutStub(srcPath)
	} else {
		return fake.streamOutReturns.result1, fake.streamOutReturns.result2
	}
}

func (fake *FakeContainer) StreamOutCallCount() int {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return len(fake.streamOutArgsForCall)
}

func (fake *FakeContainer) StreamOutArgsForCall(i int) string {
	fake.streamOutMutex.RLock()
	defer fake.streamOutMutex.RUnlock()
	return fake.streamOutArgsForCall[i].srcPath
}

func (fake *FakeContainer) StreamOutReturns(result1 io.ReadCloser, result2 error) {
	fake.StreamOutStub = nil
	fake.streamOutReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitBandwidth(limits api.BandwidthLimits) error {
	fake.limitBandwidthMutex.Lock()
	fake.limitBandwidthArgsForCall = append(fake.limitBandwidthArgsForCall, struct {
		limits api.BandwidthLimits
	}{limits})
	fake.limitBandwidthMutex.Unlock()
	if fake.LimitBandwidthStub != nil {
		return fake.LimitBandwidthStub(limits)
	} else {
		return fake.limitBandwidthReturns.result1
	}
}

func (fake *FakeContainer) LimitBandwidthCallCount() int {
	fake.limitBandwidthMutex.RLock()
	defer fake.limitBandwidthMutex.RUnlock()
	return len(fake.limitBandwidthArgsForCall)
}

func (fake *FakeContainer) LimitBandwidthArgsForCall(i int) api.BandwidthLimits {
	fake.limitBandwidthMutex.RLock()
	defer fake.limitBandwidthMutex.RUnlock()
	return fake.limitBandwidthArgsForCall[i].limits
}

func (fake *FakeContainer) LimitBandwidthReturns(result1 error) {
	fake.LimitBandwidthStub = nil
	fake.limitBandwidthReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentBandwidthLimits() (api.BandwidthLimits, error) {
	fake.currentBandwidthLimitsMutex.Lock()
	fake.currentBandwidthLimitsArgsForCall = append(fake.currentBandwidthLimitsArgsForCall, struct{}{})
	fake.currentBandwidthLimitsMutex.Unlock()
	if fake.CurrentBandwidthLimitsStub != nil {
		return fake.CurrentBandwidthLimitsStub()
	} else {
		return fake.currentBandwidthLimitsReturns.result1, fake.currentBandwidthLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentBandwidthLimitsCallCount() int {
	fake.currentBandwidthLimitsMutex.RLock()
	defer fake.currentBandwidthLimitsMutex.RUnlock()
	return len(fake.currentBandwidthLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentBandwidthLimitsReturns(result1 api.BandwidthLimits, result2 error) {
	fake.CurrentBandwidthLimitsStub = nil
	fake.currentBandwidthLimitsReturns = struct {
		result1 api.BandwidthLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitCPU(limits api.CPULimits) error {
	fake.limitCPUMutex.Lock()
	fake.limitCPUArgsForCall = append(fake.limitCPUArgsForCall, struct {
		limits api.CPULimits
	}{limits})
	fake.limitCPUMutex.Unlock()
	if fake.LimitCPUStub != nil {
		return fake.LimitCPUStub(limits)
	} else {
		return fake.limitCPUReturns.result1
	}
}

func (fake *FakeContainer) LimitCPUCallCount() int {
	fake.limitCPUMutex.RLock()
	defer fake.limitCPUMutex.RUnlock()
	return len(fake.limitCPUArgsForCall)
}

func (fake *FakeContainer) LimitCPUArgsForCall(i int) api.CPULimits {
	fake.limitCPUMutex.RLock()
	defer fake.limitCPUMutex.RUnlock()
	return fake.limitCPUArgsForCall[i].limits
}

func (fake *FakeContainer) LimitCPUReturns(result1 error) {
	fake.LimitCPUStub = nil
	fake.limitCPUReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentCPULimits() (api.CPULimits, error) {
	fake.currentCPULimitsMutex.Lock()
	fake.currentCPULimitsArgsForCall = append(fake.currentCPULimitsArgsForCall, struct{}{})
	fake.currentCPULimitsMutex.Unlock()
	if fake.CurrentCPULimitsStub != nil {
		return fake.CurrentCPULimitsStub()
	} else {
		return fake.currentCPULimitsReturns.result1, fake.currentCPULimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentCPULimitsCallCount() int {
	fake.currentCPULimitsMutex.RLock()
	defer fake.currentCPULimitsMutex.RUnlock()
	return len(fake.currentCPULimitsArgsForCall)
}

func (fake *FakeContainer) CurrentCPULimitsReturns(result1 api.CPULimits, result2 error) {
	fake.CurrentCPULimitsStub = nil
	fake.currentCPULimitsReturns = struct {
		result1 api.CPULimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitDisk(limits api.DiskLimits) error {
	fake.limitDiskMutex.Lock()
	fake.limitDiskArgsForCall = append(fake.limitDiskArgsForCall, struct {
		limits api.DiskLimits
	}{limits})
	fake.limitDiskMutex.Unlock()
	if fake.LimitDiskStub != nil {
		return fake.LimitDiskStub(limits)
	} else {
		return fake.limitDiskReturns.result1
	}
}

func (fake *FakeContainer) LimitDiskCallCount() int {
	fake.limitDiskMutex.RLock()
	defer fake.limitDiskMutex.RUnlock()
	return len(fake.limitDiskArgsForCall)
}

func (fake *FakeContainer) LimitDiskArgsForCall(i int) api.DiskLimits {
	fake.limitDiskMutex.RLock()
	defer fake.limitDiskMutex.RUnlock()
	return fake.limitDiskArgsForCall[i].limits
}

func (fake *FakeContainer) LimitDiskReturns(result1 error) {
	fake.LimitDiskStub = nil
	fake.limitDiskReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentDiskLimits() (api.DiskLimits, error) {
	fake.currentDiskLimitsMutex.Lock()
	fake.currentDiskLimitsArgsForCall = append(fake.currentDiskLimitsArgsForCall, struct{}{})
	fake.currentDiskLimitsMutex.Unlock()
	if fake.CurrentDiskLimitsStub != nil {
		return fake.CurrentDiskLimitsStub()
	} else {
		return fake.currentDiskLimitsReturns.result1, fake.currentDiskLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentDiskLimitsCallCount() int {
	fake.currentDiskLimitsMutex.RLock()
	defer fake.currentDiskLimitsMutex.RUnlock()
	return len(fake.currentDiskLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentDiskLimitsReturns(result1 api.DiskLimits, result2 error) {
	fake.CurrentDiskLimitsStub = nil
	fake.currentDiskLimitsReturns = struct {
		result1 api.DiskLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) LimitMemory(limits api.MemoryLimits) error {
	fake.limitMemoryMutex.Lock()
	fake.limitMemoryArgsForCall = append(fake.limitMemoryArgsForCall, struct {
		limits api.MemoryLimits
	}{limits})
	fake.limitMemoryMutex.Unlock()
	if fake.LimitMemoryStub != nil {
		return fake.LimitMemoryStub(limits)
	} else {
		return fake.limitMemoryReturns.result1
	}
}

func (fake *FakeContainer) LimitMemoryCallCount() int {
	fake.limitMemoryMutex.RLock()
	defer fake.limitMemoryMutex.RUnlock()
	return len(fake.limitMemoryArgsForCall)
}

func (fake *FakeContainer) LimitMemoryArgsForCall(i int) api.MemoryLimits {
	fake.limitMemoryMutex.RLock()
	defer fake.limitMemoryMutex.RUnlock()
	return fake.limitMemoryArgsForCall[i].limits
}

func (fake *FakeContainer) LimitMemoryReturns(result1 error) {
	fake.LimitMemoryStub = nil
	fake.limitMemoryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) CurrentMemoryLimits() (api.MemoryLimits, error) {
	fake.currentMemoryLimitsMutex.Lock()
	fake.currentMemoryLimitsArgsForCall = append(fake.currentMemoryLimitsArgsForCall, struct{}{})
	fake.currentMemoryLimitsMutex.Unlock()
	if fake.CurrentMemoryLimitsStub != nil {
		return fake.CurrentMemoryLimitsStub()
	} else {
		return fake.currentMemoryLimitsReturns.result1, fake.currentMemoryLimitsReturns.result2
	}
}

func (fake *FakeContainer) CurrentMemoryLimitsCallCount() int {
	fake.currentMemoryLimitsMutex.RLock()
	defer fake.currentMemoryLimitsMutex.RUnlock()
	return len(fake.currentMemoryLimitsArgsForCall)
}

func (fake *FakeContainer) CurrentMemoryLimitsReturns(result1 api.MemoryLimits, result2 error) {
	fake.CurrentMemoryLimitsStub = nil
	fake.currentMemoryLimitsReturns = struct {
		result1 api.MemoryLimits
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) NetIn(hostPort uint32, containerPort uint32) (uint32, uint32, error) {
	fake.netInMutex.Lock()
	fake.netInArgsForCall = append(fake.netInArgsForCall, struct {
		hostPort      uint32
		containerPort uint32
	}{hostPort, containerPort})
	fake.netInMutex.Unlock()
	if fake.NetInStub != nil {
		return fake.NetInStub(hostPort, containerPort)
	} else {
		return fake.netInReturns.result1, fake.netInReturns.result2, fake.netInReturns.result3
	}
}

func (fake *FakeContainer) NetInCallCount() int {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return len(fake.netInArgsForCall)
}

func (fake *FakeContainer) NetInArgsForCall(i int) (uint32, uint32) {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return fake.netInArgsForCall[i].hostPort, fake.netInArgsForCall[i].containerPort
}

func (fake *FakeContainer) NetInReturns(result1 uint32, result2 uint32, result3 error) {
	fake.NetInStub = nil
	fake.netInReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeContainer) NetOut(network string, port uint32) error {
	fake.netOutMutex.Lock()
	fake.netOutArgsForCall = append(fake.netOutArgsForCall, struct {
		network string
		port    uint32
	}{network, port})
	fake.netOutMutex.Unlock()
	if fake.NetOutStub != nil {
		return fake.NetOutStub(network, port)
	} else {
		return fake.netOutReturns.result1
	}
}

func (fake *FakeContainer) NetOutCallCount() int {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return len(fake.netOutArgsForCall)
}

func (fake *FakeContainer) NetOutArgsForCall(i int) (string, uint32) {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return fake.netOutArgsForCall[i].network, fake.netOutArgsForCall[i].port
}

func (fake *FakeContainer) NetOutReturns(result1 error) {
	fake.NetOutStub = nil
	fake.netOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) Run(arg1 api.ProcessSpec, arg2 api.ProcessIO) (api.Process, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 api.ProcessSpec
		arg2 api.ProcessIO
	}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *FakeContainer) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeContainer) RunArgsForCall(i int) (api.ProcessSpec, api.ProcessIO) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2
}

func (fake *FakeContainer) RunReturns(result1 api.Process, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 api.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) Attach(arg1 uint32, arg2 api.ProcessIO) (api.Process, error) {
	fake.attachMutex.Lock()
	fake.attachArgsForCall = append(fake.attachArgsForCall, struct {
		arg1 uint32
		arg2 api.ProcessIO
	}{arg1, arg2})
	fake.attachMutex.Unlock()
	if fake.AttachStub != nil {
		return fake.AttachStub(arg1, arg2)
	} else {
		return fake.attachReturns.result1, fake.attachReturns.result2
	}
}

func (fake *FakeContainer) AttachCallCount() int {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return len(fake.attachArgsForCall)
}

func (fake *FakeContainer) AttachArgsForCall(i int) (uint32, api.ProcessIO) {
	fake.attachMutex.RLock()
	defer fake.attachMutex.RUnlock()
	return fake.attachArgsForCall[i].arg1, fake.attachArgsForCall[i].arg2
}

func (fake *FakeContainer) AttachReturns(result1 api.Process, result2 error) {
	fake.AttachStub = nil
	fake.attachReturns = struct {
		result1 api.Process
		result2 error
	}{result1, result2}
}

func (fake *FakeContainer) BindVolume(arg1 api.Volume, arg2 api.VolumeBinding) error {
	fake.bindVolumeMutex.Lock()
	fake.bindVolumeArgsForCall = append(fake.bindVolumeArgsForCall, struct {
		arg1 api.Volume
		arg2 api.VolumeBinding
	}{arg1, arg2})
	fake.bindVolumeMutex.Unlock()
	if fake.BindVolumeStub != nil {
		return fake.BindVolumeStub(arg1, arg2)
	} else {
		return fake.bindVolumeReturns.result1
	}
}

func (fake *FakeContainer) BindVolumeCallCount() int {
	fake.bindVolumeMutex.RLock()
	defer fake.bindVolumeMutex.RUnlock()
	return len(fake.bindVolumeArgsForCall)
}

func (fake *FakeContainer) BindVolumeArgsForCall(i int) (api.Volume, api.VolumeBinding) {
	fake.bindVolumeMutex.RLock()
	defer fake.bindVolumeMutex.RUnlock()
	return fake.bindVolumeArgsForCall[i].arg1, fake.bindVolumeArgsForCall[i].arg2
}

func (fake *FakeContainer) BindVolumeReturns(result1 error) {
	fake.BindVolumeStub = nil
	fake.bindVolumeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeContainer) UnbindVolume(arg1 api.Volume) error {
	fake.unbindVolumeMutex.Lock()
	fake.unbindVolumeArgsForCall = append(fake.unbindVolumeArgsForCall, struct {
		arg1 api.Volume
	}{arg1})
	fake.unbindVolumeMutex.Unlock()
	if fake.UnbindVolumeStub != nil {
		return fake.UnbindVolumeStub(arg1)
	} else {
		return fake.unbindVolumeReturns.result1
	}
}

func (fake *FakeContainer) UnbindVolumeCallCount() int {
	fake.unbindVolumeMutex.RLock()
	defer fake.unbindVolumeMutex.RUnlock()
	return len(fake.unbindVolumeArgsForCall)
}

func (fake *FakeContainer) UnbindVolumeArgsForCall(i int) api.Volume {
	fake.unbindVolumeMutex.RLock()
	defer fake.unbindVolumeMutex.RUnlock()
	return fake.unbindVolumeArgsForCall[i].arg1
}

func (fake *FakeContainer) UnbindVolumeReturns(result1 error) {
	fake.UnbindVolumeStub = nil
	fake.unbindVolumeReturns = struct {
		result1 error
	}{result1}
}

var _ api.Container = new(FakeContainer)
