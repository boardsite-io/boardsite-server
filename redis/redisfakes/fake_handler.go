// Code generated by counterfeiter. DO NOT EDIT.
package redisfakes

import (
	"context"
	"sync"
	"time"

	"github.com/heat1q/boardsite/redis"
)

type FakeHandler struct {
	AddPageStub        func(context.Context, string, string, int, any) error
	addPageMutex       sync.RWMutex
	addPageArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 int
		arg5 any
	}
	addPageReturns struct {
		result1 error
	}
	addPageReturnsOnCall map[int]struct {
		result1 error
	}
	ClearPageStub        func(context.Context, string, string) error
	clearPageMutex       sync.RWMutex
	clearPageArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	clearPageReturns struct {
		result1 error
	}
	clearPageReturnsOnCall map[int]struct {
		result1 error
	}
	ClearSessionStub        func(context.Context, string) error
	clearSessionMutex       sync.RWMutex
	clearSessionArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	clearSessionReturns struct {
		result1 error
	}
	clearSessionReturnsOnCall map[int]struct {
		result1 error
	}
	ClosePoolStub        func() error
	closePoolMutex       sync.RWMutex
	closePoolArgsForCall []struct {
	}
	closePoolReturns struct {
		result1 error
	}
	closePoolReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(context.Context, string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	DeletePageStub        func(context.Context, string, string) error
	deletePageMutex       sync.RWMutex
	deletePageArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	deletePageReturns struct {
		result1 error
	}
	deletePageReturnsOnCall map[int]struct {
		result1 error
	}
	GetStub        func(context.Context, string) (any, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getReturns struct {
		result1 any
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 any
		result2 error
	}
	GetPageMetaStub        func(context.Context, string, string, any) error
	getPageMetaMutex       sync.RWMutex
	getPageMetaArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 any
	}
	getPageMetaReturns struct {
		result1 error
	}
	getPageMetaReturnsOnCall map[int]struct {
		result1 error
	}
	GetPageRankStub        func(context.Context, string) ([]string, error)
	getPageRankMutex       sync.RWMutex
	getPageRankArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getPageRankReturns struct {
		result1 []string
		result2 error
	}
	getPageRankReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetPageStrokesStub        func(context.Context, string, string) ([][]byte, error)
	getPageStrokesMutex       sync.RWMutex
	getPageStrokesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	getPageStrokesReturns struct {
		result1 [][]byte
		result2 error
	}
	getPageStrokesReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	PutStub        func(context.Context, string, any, time.Duration) error
	putMutex       sync.RWMutex
	putArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 any
		arg4 time.Duration
	}
	putReturns struct {
		result1 error
	}
	putReturnsOnCall map[int]struct {
		result1 error
	}
	SetPageMetaStub        func(context.Context, string, string, any) error
	setPageMetaMutex       sync.RWMutex
	setPageMetaArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 any
	}
	setPageMetaReturns struct {
		result1 error
	}
	setPageMetaReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateStrokesStub        func(context.Context, string, ...redis.Stroke) error
	updateStrokesMutex       sync.RWMutex
	updateStrokesArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 []redis.Stroke
	}
	updateStrokesReturns struct {
		result1 error
	}
	updateStrokesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]any
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandler) AddPage(arg1 context.Context, arg2 string, arg3 string, arg4 int, arg5 any) error {
	fake.addPageMutex.Lock()
	ret, specificReturn := fake.addPageReturnsOnCall[len(fake.addPageArgsForCall)]
	fake.addPageArgsForCall = append(fake.addPageArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 int
		arg5 any
	}{arg1, arg2, arg3, arg4, arg5})
	stub := fake.AddPageStub
	fakeReturns := fake.addPageReturns
	fake.recordInvocation("AddPage", []any{arg1, arg2, arg3, arg4, arg5})
	fake.addPageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) AddPageCallCount() int {
	fake.addPageMutex.RLock()
	defer fake.addPageMutex.RUnlock()
	return len(fake.addPageArgsForCall)
}

func (fake *FakeHandler) AddPageCalls(stub func(context.Context, string, string, int, any) error) {
	fake.addPageMutex.Lock()
	defer fake.addPageMutex.Unlock()
	fake.AddPageStub = stub
}

func (fake *FakeHandler) AddPageArgsForCall(i int) (context.Context, string, string, int, any) {
	fake.addPageMutex.RLock()
	defer fake.addPageMutex.RUnlock()
	argsForCall := fake.addPageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *FakeHandler) AddPageReturns(result1 error) {
	fake.addPageMutex.Lock()
	defer fake.addPageMutex.Unlock()
	fake.AddPageStub = nil
	fake.addPageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) AddPageReturnsOnCall(i int, result1 error) {
	fake.addPageMutex.Lock()
	defer fake.addPageMutex.Unlock()
	fake.AddPageStub = nil
	if fake.addPageReturnsOnCall == nil {
		fake.addPageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addPageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClearPage(arg1 context.Context, arg2 string, arg3 string) error {
	fake.clearPageMutex.Lock()
	ret, specificReturn := fake.clearPageReturnsOnCall[len(fake.clearPageArgsForCall)]
	fake.clearPageArgsForCall = append(fake.clearPageArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ClearPageStub
	fakeReturns := fake.clearPageReturns
	fake.recordInvocation("ClearPage", []any{arg1, arg2, arg3})
	fake.clearPageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) ClearPageCallCount() int {
	fake.clearPageMutex.RLock()
	defer fake.clearPageMutex.RUnlock()
	return len(fake.clearPageArgsForCall)
}

func (fake *FakeHandler) ClearPageCalls(stub func(context.Context, string, string) error) {
	fake.clearPageMutex.Lock()
	defer fake.clearPageMutex.Unlock()
	fake.ClearPageStub = stub
}

func (fake *FakeHandler) ClearPageArgsForCall(i int) (context.Context, string, string) {
	fake.clearPageMutex.RLock()
	defer fake.clearPageMutex.RUnlock()
	argsForCall := fake.clearPageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHandler) ClearPageReturns(result1 error) {
	fake.clearPageMutex.Lock()
	defer fake.clearPageMutex.Unlock()
	fake.ClearPageStub = nil
	fake.clearPageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClearPageReturnsOnCall(i int, result1 error) {
	fake.clearPageMutex.Lock()
	defer fake.clearPageMutex.Unlock()
	fake.ClearPageStub = nil
	if fake.clearPageReturnsOnCall == nil {
		fake.clearPageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearPageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClearSession(arg1 context.Context, arg2 string) error {
	fake.clearSessionMutex.Lock()
	ret, specificReturn := fake.clearSessionReturnsOnCall[len(fake.clearSessionArgsForCall)]
	fake.clearSessionArgsForCall = append(fake.clearSessionArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.ClearSessionStub
	fakeReturns := fake.clearSessionReturns
	fake.recordInvocation("ClearSession", []any{arg1, arg2})
	fake.clearSessionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) ClearSessionCallCount() int {
	fake.clearSessionMutex.RLock()
	defer fake.clearSessionMutex.RUnlock()
	return len(fake.clearSessionArgsForCall)
}

func (fake *FakeHandler) ClearSessionCalls(stub func(context.Context, string) error) {
	fake.clearSessionMutex.Lock()
	defer fake.clearSessionMutex.Unlock()
	fake.ClearSessionStub = stub
}

func (fake *FakeHandler) ClearSessionArgsForCall(i int) (context.Context, string) {
	fake.clearSessionMutex.RLock()
	defer fake.clearSessionMutex.RUnlock()
	argsForCall := fake.clearSessionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) ClearSessionReturns(result1 error) {
	fake.clearSessionMutex.Lock()
	defer fake.clearSessionMutex.Unlock()
	fake.ClearSessionStub = nil
	fake.clearSessionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClearSessionReturnsOnCall(i int, result1 error) {
	fake.clearSessionMutex.Lock()
	defer fake.clearSessionMutex.Unlock()
	fake.ClearSessionStub = nil
	if fake.clearSessionReturnsOnCall == nil {
		fake.clearSessionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.clearSessionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClosePool() error {
	fake.closePoolMutex.Lock()
	ret, specificReturn := fake.closePoolReturnsOnCall[len(fake.closePoolArgsForCall)]
	fake.closePoolArgsForCall = append(fake.closePoolArgsForCall, struct {
	}{})
	stub := fake.ClosePoolStub
	fakeReturns := fake.closePoolReturns
	fake.recordInvocation("ClosePool", []any{})
	fake.closePoolMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) ClosePoolCallCount() int {
	fake.closePoolMutex.RLock()
	defer fake.closePoolMutex.RUnlock()
	return len(fake.closePoolArgsForCall)
}

func (fake *FakeHandler) ClosePoolCalls(stub func() error) {
	fake.closePoolMutex.Lock()
	defer fake.closePoolMutex.Unlock()
	fake.ClosePoolStub = stub
}

func (fake *FakeHandler) ClosePoolReturns(result1 error) {
	fake.closePoolMutex.Lock()
	defer fake.closePoolMutex.Unlock()
	fake.ClosePoolStub = nil
	fake.closePoolReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) ClosePoolReturnsOnCall(i int, result1 error) {
	fake.closePoolMutex.Lock()
	defer fake.closePoolMutex.Unlock()
	fake.ClosePoolStub = nil
	if fake.closePoolReturnsOnCall == nil {
		fake.closePoolReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closePoolReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) Delete(arg1 context.Context, arg2 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []any{arg1, arg2})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeHandler) DeleteCalls(stub func(context.Context, string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeHandler) DeleteArgsForCall(i int) (context.Context, string) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) DeletePage(arg1 context.Context, arg2 string, arg3 string) error {
	fake.deletePageMutex.Lock()
	ret, specificReturn := fake.deletePageReturnsOnCall[len(fake.deletePageArgsForCall)]
	fake.deletePageArgsForCall = append(fake.deletePageArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.DeletePageStub
	fakeReturns := fake.deletePageReturns
	fake.recordInvocation("DeletePage", []any{arg1, arg2, arg3})
	fake.deletePageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) DeletePageCallCount() int {
	fake.deletePageMutex.RLock()
	defer fake.deletePageMutex.RUnlock()
	return len(fake.deletePageArgsForCall)
}

func (fake *FakeHandler) DeletePageCalls(stub func(context.Context, string, string) error) {
	fake.deletePageMutex.Lock()
	defer fake.deletePageMutex.Unlock()
	fake.DeletePageStub = stub
}

func (fake *FakeHandler) DeletePageArgsForCall(i int) (context.Context, string, string) {
	fake.deletePageMutex.RLock()
	defer fake.deletePageMutex.RUnlock()
	argsForCall := fake.deletePageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHandler) DeletePageReturns(result1 error) {
	fake.deletePageMutex.Lock()
	defer fake.deletePageMutex.Unlock()
	fake.DeletePageStub = nil
	fake.deletePageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) DeletePageReturnsOnCall(i int, result1 error) {
	fake.deletePageMutex.Lock()
	defer fake.deletePageMutex.Unlock()
	fake.DeletePageStub = nil
	if fake.deletePageReturnsOnCall == nil {
		fake.deletePageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deletePageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) Get(arg1 context.Context, arg2 string) (any, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []any{arg1, arg2})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHandler) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeHandler) GetCalls(stub func(context.Context, string) (any, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeHandler) GetArgsForCall(i int) (context.Context, string) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) GetReturns(result1 any, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 any
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) GetReturnsOnCall(i int, result1 any, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 any
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 any
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) GetPageMeta(arg1 context.Context, arg2 string, arg3 string, arg4 any) error {
	fake.getPageMetaMutex.Lock()
	ret, specificReturn := fake.getPageMetaReturnsOnCall[len(fake.getPageMetaArgsForCall)]
	fake.getPageMetaArgsForCall = append(fake.getPageMetaArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 any
	}{arg1, arg2, arg3, arg4})
	stub := fake.GetPageMetaStub
	fakeReturns := fake.getPageMetaReturns
	fake.recordInvocation("GetPageMeta", []any{arg1, arg2, arg3, arg4})
	fake.getPageMetaMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) GetPageMetaCallCount() int {
	fake.getPageMetaMutex.RLock()
	defer fake.getPageMetaMutex.RUnlock()
	return len(fake.getPageMetaArgsForCall)
}

func (fake *FakeHandler) GetPageMetaCalls(stub func(context.Context, string, string, any) error) {
	fake.getPageMetaMutex.Lock()
	defer fake.getPageMetaMutex.Unlock()
	fake.GetPageMetaStub = stub
}

func (fake *FakeHandler) GetPageMetaArgsForCall(i int) (context.Context, string, string, any) {
	fake.getPageMetaMutex.RLock()
	defer fake.getPageMetaMutex.RUnlock()
	argsForCall := fake.getPageMetaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeHandler) GetPageMetaReturns(result1 error) {
	fake.getPageMetaMutex.Lock()
	defer fake.getPageMetaMutex.Unlock()
	fake.GetPageMetaStub = nil
	fake.getPageMetaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) GetPageMetaReturnsOnCall(i int, result1 error) {
	fake.getPageMetaMutex.Lock()
	defer fake.getPageMetaMutex.Unlock()
	fake.GetPageMetaStub = nil
	if fake.getPageMetaReturnsOnCall == nil {
		fake.getPageMetaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.getPageMetaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) GetPageRank(arg1 context.Context, arg2 string) ([]string, error) {
	fake.getPageRankMutex.Lock()
	ret, specificReturn := fake.getPageRankReturnsOnCall[len(fake.getPageRankArgsForCall)]
	fake.getPageRankArgsForCall = append(fake.getPageRankArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetPageRankStub
	fakeReturns := fake.getPageRankReturns
	fake.recordInvocation("GetPageRank", []any{arg1, arg2})
	fake.getPageRankMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHandler) GetPageRankCallCount() int {
	fake.getPageRankMutex.RLock()
	defer fake.getPageRankMutex.RUnlock()
	return len(fake.getPageRankArgsForCall)
}

func (fake *FakeHandler) GetPageRankCalls(stub func(context.Context, string) ([]string, error)) {
	fake.getPageRankMutex.Lock()
	defer fake.getPageRankMutex.Unlock()
	fake.GetPageRankStub = stub
}

func (fake *FakeHandler) GetPageRankArgsForCall(i int) (context.Context, string) {
	fake.getPageRankMutex.RLock()
	defer fake.getPageRankMutex.RUnlock()
	argsForCall := fake.getPageRankArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeHandler) GetPageRankReturns(result1 []string, result2 error) {
	fake.getPageRankMutex.Lock()
	defer fake.getPageRankMutex.Unlock()
	fake.GetPageRankStub = nil
	fake.getPageRankReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) GetPageRankReturnsOnCall(i int, result1 []string, result2 error) {
	fake.getPageRankMutex.Lock()
	defer fake.getPageRankMutex.Unlock()
	fake.GetPageRankStub = nil
	if fake.getPageRankReturnsOnCall == nil {
		fake.getPageRankReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getPageRankReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) GetPageStrokes(arg1 context.Context, arg2 string, arg3 string) ([][]byte, error) {
	fake.getPageStrokesMutex.Lock()
	ret, specificReturn := fake.getPageStrokesReturnsOnCall[len(fake.getPageStrokesArgsForCall)]
	fake.getPageStrokesArgsForCall = append(fake.getPageStrokesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.GetPageStrokesStub
	fakeReturns := fake.getPageStrokesReturns
	fake.recordInvocation("GetPageStrokes", []any{arg1, arg2, arg3})
	fake.getPageStrokesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeHandler) GetPageStrokesCallCount() int {
	fake.getPageStrokesMutex.RLock()
	defer fake.getPageStrokesMutex.RUnlock()
	return len(fake.getPageStrokesArgsForCall)
}

func (fake *FakeHandler) GetPageStrokesCalls(stub func(context.Context, string, string) ([][]byte, error)) {
	fake.getPageStrokesMutex.Lock()
	defer fake.getPageStrokesMutex.Unlock()
	fake.GetPageStrokesStub = stub
}

func (fake *FakeHandler) GetPageStrokesArgsForCall(i int) (context.Context, string, string) {
	fake.getPageStrokesMutex.RLock()
	defer fake.getPageStrokesMutex.RUnlock()
	argsForCall := fake.getPageStrokesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHandler) GetPageStrokesReturns(result1 [][]byte, result2 error) {
	fake.getPageStrokesMutex.Lock()
	defer fake.getPageStrokesMutex.Unlock()
	fake.GetPageStrokesStub = nil
	fake.getPageStrokesReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) GetPageStrokesReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.getPageStrokesMutex.Lock()
	defer fake.getPageStrokesMutex.Unlock()
	fake.GetPageStrokesStub = nil
	if fake.getPageStrokesReturnsOnCall == nil {
		fake.getPageStrokesReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.getPageStrokesReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *FakeHandler) Put(arg1 context.Context, arg2 string, arg3 any, arg4 time.Duration) error {
	fake.putMutex.Lock()
	ret, specificReturn := fake.putReturnsOnCall[len(fake.putArgsForCall)]
	fake.putArgsForCall = append(fake.putArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 any
		arg4 time.Duration
	}{arg1, arg2, arg3, arg4})
	stub := fake.PutStub
	fakeReturns := fake.putReturns
	fake.recordInvocation("Put", []any{arg1, arg2, arg3, arg4})
	fake.putMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) PutCallCount() int {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	return len(fake.putArgsForCall)
}

func (fake *FakeHandler) PutCalls(stub func(context.Context, string, any, time.Duration) error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = stub
}

func (fake *FakeHandler) PutArgsForCall(i int) (context.Context, string, any, time.Duration) {
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	argsForCall := fake.putArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeHandler) PutReturns(result1 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	fake.putReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) PutReturnsOnCall(i int, result1 error) {
	fake.putMutex.Lock()
	defer fake.putMutex.Unlock()
	fake.PutStub = nil
	if fake.putReturnsOnCall == nil {
		fake.putReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.putReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) SetPageMeta(arg1 context.Context, arg2 string, arg3 string, arg4 any) error {
	fake.setPageMetaMutex.Lock()
	ret, specificReturn := fake.setPageMetaReturnsOnCall[len(fake.setPageMetaArgsForCall)]
	fake.setPageMetaArgsForCall = append(fake.setPageMetaArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 any
	}{arg1, arg2, arg3, arg4})
	stub := fake.SetPageMetaStub
	fakeReturns := fake.setPageMetaReturns
	fake.recordInvocation("SetPageMeta", []any{arg1, arg2, arg3, arg4})
	fake.setPageMetaMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) SetPageMetaCallCount() int {
	fake.setPageMetaMutex.RLock()
	defer fake.setPageMetaMutex.RUnlock()
	return len(fake.setPageMetaArgsForCall)
}

func (fake *FakeHandler) SetPageMetaCalls(stub func(context.Context, string, string, any) error) {
	fake.setPageMetaMutex.Lock()
	defer fake.setPageMetaMutex.Unlock()
	fake.SetPageMetaStub = stub
}

func (fake *FakeHandler) SetPageMetaArgsForCall(i int) (context.Context, string, string, any) {
	fake.setPageMetaMutex.RLock()
	defer fake.setPageMetaMutex.RUnlock()
	argsForCall := fake.setPageMetaArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeHandler) SetPageMetaReturns(result1 error) {
	fake.setPageMetaMutex.Lock()
	defer fake.setPageMetaMutex.Unlock()
	fake.SetPageMetaStub = nil
	fake.setPageMetaReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) SetPageMetaReturnsOnCall(i int, result1 error) {
	fake.setPageMetaMutex.Lock()
	defer fake.setPageMetaMutex.Unlock()
	fake.SetPageMetaStub = nil
	if fake.setPageMetaReturnsOnCall == nil {
		fake.setPageMetaReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setPageMetaReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) UpdateStrokes(arg1 context.Context, arg2 string, arg3 ...redis.Stroke) error {
	fake.updateStrokesMutex.Lock()
	ret, specificReturn := fake.updateStrokesReturnsOnCall[len(fake.updateStrokesArgsForCall)]
	fake.updateStrokesArgsForCall = append(fake.updateStrokesArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 []redis.Stroke
	}{arg1, arg2, arg3})
	stub := fake.UpdateStrokesStub
	fakeReturns := fake.updateStrokesReturns
	fake.recordInvocation("UpdateStrokes", []any{arg1, arg2, arg3})
	fake.updateStrokesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeHandler) UpdateStrokesCallCount() int {
	fake.updateStrokesMutex.RLock()
	defer fake.updateStrokesMutex.RUnlock()
	return len(fake.updateStrokesArgsForCall)
}

func (fake *FakeHandler) UpdateStrokesCalls(stub func(context.Context, string, ...redis.Stroke) error) {
	fake.updateStrokesMutex.Lock()
	defer fake.updateStrokesMutex.Unlock()
	fake.UpdateStrokesStub = stub
}

func (fake *FakeHandler) UpdateStrokesArgsForCall(i int) (context.Context, string, []redis.Stroke) {
	fake.updateStrokesMutex.RLock()
	defer fake.updateStrokesMutex.RUnlock()
	argsForCall := fake.updateStrokesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeHandler) UpdateStrokesReturns(result1 error) {
	fake.updateStrokesMutex.Lock()
	defer fake.updateStrokesMutex.Unlock()
	fake.UpdateStrokesStub = nil
	fake.updateStrokesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) UpdateStrokesReturnsOnCall(i int, result1 error) {
	fake.updateStrokesMutex.Lock()
	defer fake.updateStrokesMutex.Unlock()
	fake.UpdateStrokesStub = nil
	if fake.updateStrokesReturnsOnCall == nil {
		fake.updateStrokesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateStrokesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeHandler) Invocations() map[string][][]any {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addPageMutex.RLock()
	defer fake.addPageMutex.RUnlock()
	fake.clearPageMutex.RLock()
	defer fake.clearPageMutex.RUnlock()
	fake.clearSessionMutex.RLock()
	defer fake.clearSessionMutex.RUnlock()
	fake.closePoolMutex.RLock()
	defer fake.closePoolMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.deletePageMutex.RLock()
	defer fake.deletePageMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	fake.getPageMetaMutex.RLock()
	defer fake.getPageMetaMutex.RUnlock()
	fake.getPageRankMutex.RLock()
	defer fake.getPageRankMutex.RUnlock()
	fake.getPageStrokesMutex.RLock()
	defer fake.getPageStrokesMutex.RUnlock()
	fake.putMutex.RLock()
	defer fake.putMutex.RUnlock()
	fake.setPageMetaMutex.RLock()
	defer fake.setPageMetaMutex.RUnlock()
	fake.updateStrokesMutex.RLock()
	defer fake.updateStrokesMutex.RUnlock()
	copiedInvocations := map[string][][]any{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandler) recordInvocation(key string, args []any) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]any{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]any{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ redis.Handler = new(FakeHandler)
