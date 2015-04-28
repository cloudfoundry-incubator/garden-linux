// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-linux/old/repository_fetcher"
)

type FakeRegistryProvider struct {
	ProvideRegistryStub        func(hostname string) (repository_fetcher.Registry, error)
	provideRegistryMutex       sync.RWMutex
	provideRegistryArgsForCall []struct {
		hostname string
	}
	provideRegistryReturns struct {
		result1 repository_fetcher.Registry
		result2 error
	}
	ApplyDefaultHostnameStub        func(hostname string) string
	applyDefaultHostnameMutex       sync.RWMutex
	applyDefaultHostnameArgsForCall []struct {
		hostname string
	}
	applyDefaultHostnameReturns struct {
		result1 string
	}
}

func (fake *FakeRegistryProvider) ProvideRegistry(hostname string) (repository_fetcher.Registry, error) {
	fake.provideRegistryMutex.Lock()
	fake.provideRegistryArgsForCall = append(fake.provideRegistryArgsForCall, struct {
		hostname string
	}{hostname})
	fake.provideRegistryMutex.Unlock()
	if fake.ProvideRegistryStub != nil {
		return fake.ProvideRegistryStub(hostname)
	} else {
		return fake.provideRegistryReturns.result1, fake.provideRegistryReturns.result2
	}
}

func (fake *FakeRegistryProvider) ProvideRegistryCallCount() int {
	fake.provideRegistryMutex.RLock()
	defer fake.provideRegistryMutex.RUnlock()
	return len(fake.provideRegistryArgsForCall)
}

func (fake *FakeRegistryProvider) ProvideRegistryArgsForCall(i int) string {
	fake.provideRegistryMutex.RLock()
	defer fake.provideRegistryMutex.RUnlock()
	return fake.provideRegistryArgsForCall[i].hostname
}

func (fake *FakeRegistryProvider) ProvideRegistryReturns(result1 repository_fetcher.Registry, result2 error) {
	fake.ProvideRegistryStub = nil
	fake.provideRegistryReturns = struct {
		result1 repository_fetcher.Registry
		result2 error
	}{result1, result2}
}

func (fake *FakeRegistryProvider) ApplyDefaultHostname(hostname string) string {
	fake.applyDefaultHostnameMutex.Lock()
	fake.applyDefaultHostnameArgsForCall = append(fake.applyDefaultHostnameArgsForCall, struct {
		hostname string
	}{hostname})
	fake.applyDefaultHostnameMutex.Unlock()
	if fake.ApplyDefaultHostnameStub != nil {
		return fake.ApplyDefaultHostnameStub(hostname)
	} else {
		return fake.applyDefaultHostnameReturns.result1
	}
}

func (fake *FakeRegistryProvider) ApplyDefaultHostnameCallCount() int {
	fake.applyDefaultHostnameMutex.RLock()
	defer fake.applyDefaultHostnameMutex.RUnlock()
	return len(fake.applyDefaultHostnameArgsForCall)
}

func (fake *FakeRegistryProvider) ApplyDefaultHostnameArgsForCall(i int) string {
	fake.applyDefaultHostnameMutex.RLock()
	defer fake.applyDefaultHostnameMutex.RUnlock()
	return fake.applyDefaultHostnameArgsForCall[i].hostname
}

func (fake *FakeRegistryProvider) ApplyDefaultHostnameReturns(result1 string) {
	fake.ApplyDefaultHostnameStub = nil
	fake.applyDefaultHostnameReturns = struct {
		result1 string
	}{result1}
}

var _ repository_fetcher.RegistryProvider = new(FakeRegistryProvider)
