// This file was generated by counterfeiter
package v2fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actors/configactions"
	"code.cloudfoundry.org/cli/commands/v2"
)

type FakeAPIConfigActor struct {
	ClearTargetStub        func()
	clearTargetMutex       sync.RWMutex
	clearTargetArgsForCall []struct{}
	SetTargetStub          func() (configactions.Warnings, error)
	setTargetMutex         sync.RWMutex
	setTargetArgsForCall   []struct{}
	setTargetReturns       struct {
		result1 configactions.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPIConfigActor) ClearTarget() {
	fake.clearTargetMutex.Lock()
	fake.clearTargetArgsForCall = append(fake.clearTargetArgsForCall, struct{}{})
	fake.recordInvocation("ClearTarget", []interface{}{})
	fake.clearTargetMutex.Unlock()
	if fake.ClearTargetStub != nil {
		fake.ClearTargetStub()
	}
}

func (fake *FakeAPIConfigActor) ClearTargetCallCount() int {
	fake.clearTargetMutex.RLock()
	defer fake.clearTargetMutex.RUnlock()
	return len(fake.clearTargetArgsForCall)
}

func (fake *FakeAPIConfigActor) SetTarget() (configactions.Warnings, error) {
	fake.setTargetMutex.Lock()
	fake.setTargetArgsForCall = append(fake.setTargetArgsForCall, struct{}{})
	fake.recordInvocation("SetTarget", []interface{}{})
	fake.setTargetMutex.Unlock()
	if fake.SetTargetStub != nil {
		return fake.SetTargetStub()
	} else {
		return fake.setTargetReturns.result1, fake.setTargetReturns.result2
	}
}

func (fake *FakeAPIConfigActor) SetTargetCallCount() int {
	fake.setTargetMutex.RLock()
	defer fake.setTargetMutex.RUnlock()
	return len(fake.setTargetArgsForCall)
}

func (fake *FakeAPIConfigActor) SetTargetReturns(result1 configactions.Warnings, result2 error) {
	fake.SetTargetStub = nil
	fake.setTargetReturns = struct {
		result1 configactions.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeAPIConfigActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clearTargetMutex.RLock()
	defer fake.clearTargetMutex.RUnlock()
	fake.setTargetMutex.RLock()
	defer fake.setTargetMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAPIConfigActor) recordInvocation(key string, args []interface{}) {
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

var _ v2.APIConfigActor = new(FakeAPIConfigActor)
