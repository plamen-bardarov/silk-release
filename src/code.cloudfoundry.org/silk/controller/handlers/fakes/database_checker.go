// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type DatabaseChecker struct {
	CheckDatabaseStub        func() error
	checkDatabaseMutex       sync.RWMutex
	checkDatabaseArgsForCall []struct {
	}
	checkDatabaseReturns struct {
		result1 error
	}
	checkDatabaseReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DatabaseChecker) CheckDatabase() error {
	fake.checkDatabaseMutex.Lock()
	ret, specificReturn := fake.checkDatabaseReturnsOnCall[len(fake.checkDatabaseArgsForCall)]
	fake.checkDatabaseArgsForCall = append(fake.checkDatabaseArgsForCall, struct {
	}{})
	stub := fake.CheckDatabaseStub
	fakeReturns := fake.checkDatabaseReturns
	fake.recordInvocation("CheckDatabase", []interface{}{})
	fake.checkDatabaseMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *DatabaseChecker) CheckDatabaseCallCount() int {
	fake.checkDatabaseMutex.RLock()
	defer fake.checkDatabaseMutex.RUnlock()
	return len(fake.checkDatabaseArgsForCall)
}

func (fake *DatabaseChecker) CheckDatabaseCalls(stub func() error) {
	fake.checkDatabaseMutex.Lock()
	defer fake.checkDatabaseMutex.Unlock()
	fake.CheckDatabaseStub = stub
}

func (fake *DatabaseChecker) CheckDatabaseReturns(result1 error) {
	fake.checkDatabaseMutex.Lock()
	defer fake.checkDatabaseMutex.Unlock()
	fake.CheckDatabaseStub = nil
	fake.checkDatabaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseChecker) CheckDatabaseReturnsOnCall(i int, result1 error) {
	fake.checkDatabaseMutex.Lock()
	defer fake.checkDatabaseMutex.Unlock()
	fake.CheckDatabaseStub = nil
	if fake.checkDatabaseReturnsOnCall == nil {
		fake.checkDatabaseReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkDatabaseReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseChecker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.checkDatabaseMutex.RLock()
	defer fake.checkDatabaseMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DatabaseChecker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
