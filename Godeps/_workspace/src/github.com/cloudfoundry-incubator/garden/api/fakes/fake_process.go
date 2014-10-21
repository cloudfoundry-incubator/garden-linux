// This file was generated by counterfeiter
package fakes

import (
	"github.com/cloudfoundry-incubator/garden/api"
	"sync"
)

type FakeProcess struct {
	IDStub        func() uint32
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 uint32
	}
	WaitStub        func() (int, error)
	waitMutex       sync.RWMutex
	waitArgsForCall []struct{}
	waitReturns     struct {
		result1 int
		result2 error
	}
	SetTTYStub        func(api.TTYSpec) error
	setTTYMutex       sync.RWMutex
	setTTYArgsForCall []struct {
		arg1 api.TTYSpec
	}
	setTTYReturns struct {
		result1 error
	}
}

func (fake *FakeProcess) ID() uint32 {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	if fake.IDStub != nil {
		return fake.IDStub()
	} else {
		return fake.iDReturns.result1
	}
}

func (fake *FakeProcess) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeProcess) IDReturns(result1 uint32) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeProcess) Wait() (int, error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct{}{})
	if fake.WaitStub != nil {
		return fake.WaitStub()
	} else {
		return fake.waitReturns.result1, fake.waitReturns.result2
	}
}

func (fake *FakeProcess) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeProcess) WaitReturns(result1 int, result2 error) {
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeProcess) SetTTY(arg1 api.TTYSpec) error {
	fake.setTTYMutex.Lock()
	defer fake.setTTYMutex.Unlock()
	fake.setTTYArgsForCall = append(fake.setTTYArgsForCall, struct {
		arg1 api.TTYSpec
	}{arg1})
	if fake.SetTTYStub != nil {
		return fake.SetTTYStub(arg1)
	} else {
		return fake.setTTYReturns.result1
	}
}

func (fake *FakeProcess) SetTTYCallCount() int {
	fake.setTTYMutex.RLock()
	defer fake.setTTYMutex.RUnlock()
	return len(fake.setTTYArgsForCall)
}

func (fake *FakeProcess) SetTTYArgsForCall(i int) api.TTYSpec {
	fake.setTTYMutex.RLock()
	defer fake.setTTYMutex.RUnlock()
	return fake.setTTYArgsForCall[i].arg1
}

func (fake *FakeProcess) SetTTYReturns(result1 error) {
	fake.SetTTYStub = nil
	fake.setTTYReturns = struct {
		result1 error
	}{result1}
}

var _ api.Process = new(FakeProcess)