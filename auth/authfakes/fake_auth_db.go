// This file was generated by counterfeiter
package authfakes

import (
	"sync"

	"github.com/concourse/atc/auth"
	"github.com/concourse/atc/db"
)

type FakeAuthDB struct {
	GetTeamByNameStub        func(teamName string) (db.SavedTeam, bool, error)
	getTeamByNameMutex       sync.RWMutex
	getTeamByNameArgsForCall []struct {
		teamName string
	}
	getTeamByNameReturns struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAuthDB) GetTeamByName(teamName string) (db.SavedTeam, bool, error) {
	fake.getTeamByNameMutex.Lock()
	fake.getTeamByNameArgsForCall = append(fake.getTeamByNameArgsForCall, struct {
		teamName string
	}{teamName})
	fake.recordInvocation("GetTeamByName", []interface{}{teamName})
	fake.getTeamByNameMutex.Unlock()
	if fake.GetTeamByNameStub != nil {
		return fake.GetTeamByNameStub(teamName)
	} else {
		return fake.getTeamByNameReturns.result1, fake.getTeamByNameReturns.result2, fake.getTeamByNameReturns.result3
	}
}

func (fake *FakeAuthDB) GetTeamByNameCallCount() int {
	fake.getTeamByNameMutex.RLock()
	defer fake.getTeamByNameMutex.RUnlock()
	return len(fake.getTeamByNameArgsForCall)
}

func (fake *FakeAuthDB) GetTeamByNameArgsForCall(i int) string {
	fake.getTeamByNameMutex.RLock()
	defer fake.getTeamByNameMutex.RUnlock()
	return fake.getTeamByNameArgsForCall[i].teamName
}

func (fake *FakeAuthDB) GetTeamByNameReturns(result1 db.SavedTeam, result2 bool, result3 error) {
	fake.GetTeamByNameStub = nil
	fake.getTeamByNameReturns = struct {
		result1 db.SavedTeam
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAuthDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamByNameMutex.RLock()
	defer fake.getTeamByNameMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeAuthDB) recordInvocation(key string, args []interface{}) {
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

var _ auth.AuthDB = new(FakeAuthDB)
