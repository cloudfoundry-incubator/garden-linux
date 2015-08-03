// This file was generated by counterfeiter
package fake_signaller

import (
	"sync"
	"syscall"

	"github.com/cloudfoundry-incubator/garden-linux/container_daemon"
)

type FakeSignaller struct {
	SignalStub        func(pid int, signal syscall.Signal) error
	signalMutex       sync.RWMutex
	signalArgsForCall []struct {
		pid    int
		signal syscall.Signal
	}
	signalReturns struct {
		result1 error
	}
}

func (fake *FakeSignaller) Signal(pid int, signal syscall.Signal) error {
	fake.signalMutex.Lock()
	fake.signalArgsForCall = append(fake.signalArgsForCall, struct {
		pid    int
		signal syscall.Signal
	}{pid, signal})
	fake.signalMutex.Unlock()
	if fake.SignalStub != nil {
		return fake.SignalStub(pid, signal)
	} else {
		return fake.signalReturns.result1
	}
}

func (fake *FakeSignaller) SignalCallCount() int {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return len(fake.signalArgsForCall)
}

func (fake *FakeSignaller) SignalArgsForCall(i int) (int, syscall.Signal) {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return fake.signalArgsForCall[i].pid, fake.signalArgsForCall[i].signal
}

func (fake *FakeSignaller) SignalReturns(result1 error) {
	fake.SignalStub = nil
	fake.signalReturns = struct {
		result1 error
	}{result1}
}

var _ container_daemon.Signaller = new(FakeSignaller)