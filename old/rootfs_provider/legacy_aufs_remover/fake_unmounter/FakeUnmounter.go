// This file was generated by counterfeiter
package fake_unmounter

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/old/rootfs_provider/legacy_aufs_remover"
)

type FakeUnmounter struct {
	UnmountStub        func(dir string) error
	unmountMutex       sync.RWMutex
	unmountArgsForCall []struct {
		dir string
	}
	unmountReturns struct {
		result1 error
	}
}

func (fake *FakeUnmounter) Unmount(dir string) error {
	fake.unmountMutex.Lock()
	fake.unmountArgsForCall = append(fake.unmountArgsForCall, struct {
		dir string
	}{dir})
	fake.unmountMutex.Unlock()
	if fake.UnmountStub != nil {
		return fake.UnmountStub(dir)
	} else {
		return fake.unmountReturns.result1
	}
}

func (fake *FakeUnmounter) UnmountCallCount() int {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return len(fake.unmountArgsForCall)
}

func (fake *FakeUnmounter) UnmountArgsForCall(i int) string {
	fake.unmountMutex.RLock()
	defer fake.unmountMutex.RUnlock()
	return fake.unmountArgsForCall[i].dir
}

func (fake *FakeUnmounter) UnmountReturns(result1 error) {
	fake.UnmountStub = nil
	fake.unmountReturns = struct {
		result1 error
	}{result1}
}

var _ legacy_aufs_remover.Unmounter = new(FakeUnmounter)
