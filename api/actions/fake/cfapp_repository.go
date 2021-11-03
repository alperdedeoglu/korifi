// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/actions"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CFAppRepository struct {
	AppExistsWithNameAndSpaceStub        func(context.Context, client.Client, string, string) (bool, error)
	appExistsWithNameAndSpaceMutex       sync.RWMutex
	appExistsWithNameAndSpaceArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
		arg4 string
	}
	appExistsWithNameAndSpaceReturns struct {
		result1 bool
		result2 error
	}
	appExistsWithNameAndSpaceReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	CreateAppStub        func(context.Context, client.Client, repositories.AppCreateMessage) (repositories.AppRecord, error)
	createAppMutex       sync.RWMutex
	createAppArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.AppCreateMessage
	}
	createAppReturns struct {
		result1 repositories.AppRecord
		result2 error
	}
	createAppReturnsOnCall map[int]struct {
		result1 repositories.AppRecord
		result2 error
	}
	CreateAppEnvironmentVariablesStub        func(context.Context, client.Client, repositories.AppEnvVarsRecord) (repositories.AppEnvVarsRecord, error)
	createAppEnvironmentVariablesMutex       sync.RWMutex
	createAppEnvironmentVariablesArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.AppEnvVarsRecord
	}
	createAppEnvironmentVariablesReturns struct {
		result1 repositories.AppEnvVarsRecord
		result2 error
	}
	createAppEnvironmentVariablesReturnsOnCall map[int]struct {
		result1 repositories.AppEnvVarsRecord
		result2 error
	}
	FetchAppStub        func(context.Context, client.Client, string) (repositories.AppRecord, error)
	fetchAppMutex       sync.RWMutex
	fetchAppArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}
	fetchAppReturns struct {
		result1 repositories.AppRecord
		result2 error
	}
	fetchAppReturnsOnCall map[int]struct {
		result1 repositories.AppRecord
		result2 error
	}
	FetchAppListStub        func(context.Context, client.Client) ([]repositories.AppRecord, error)
	fetchAppListMutex       sync.RWMutex
	fetchAppListArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
	}
	fetchAppListReturns struct {
		result1 []repositories.AppRecord
		result2 error
	}
	fetchAppListReturnsOnCall map[int]struct {
		result1 []repositories.AppRecord
		result2 error
	}
	FetchNamespaceStub        func(context.Context, client.Client, string) (repositories.SpaceRecord, error)
	fetchNamespaceMutex       sync.RWMutex
	fetchNamespaceArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}
	fetchNamespaceReturns struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	fetchNamespaceReturnsOnCall map[int]struct {
		result1 repositories.SpaceRecord
		result2 error
	}
	SetAppDesiredStateStub        func(context.Context, client.Client, repositories.SetAppDesiredStateMessage) (repositories.AppRecord, error)
	setAppDesiredStateMutex       sync.RWMutex
	setAppDesiredStateArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.SetAppDesiredStateMessage
	}
	setAppDesiredStateReturns struct {
		result1 repositories.AppRecord
		result2 error
	}
	setAppDesiredStateReturnsOnCall map[int]struct {
		result1 repositories.AppRecord
		result2 error
	}
	SetCurrentDropletStub        func(context.Context, client.Client, repositories.SetCurrentDropletMessage) (repositories.CurrentDropletRecord, error)
	setCurrentDropletMutex       sync.RWMutex
	setCurrentDropletArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.SetCurrentDropletMessage
	}
	setCurrentDropletReturns struct {
		result1 repositories.CurrentDropletRecord
		result2 error
	}
	setCurrentDropletReturnsOnCall map[int]struct {
		result1 repositories.CurrentDropletRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFAppRepository) AppExistsWithNameAndSpace(arg1 context.Context, arg2 client.Client, arg3 string, arg4 string) (bool, error) {
	fake.appExistsWithNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.appExistsWithNameAndSpaceReturnsOnCall[len(fake.appExistsWithNameAndSpaceArgsForCall)]
	fake.appExistsWithNameAndSpaceArgsForCall = append(fake.appExistsWithNameAndSpaceArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	stub := fake.AppExistsWithNameAndSpaceStub
	fakeReturns := fake.appExistsWithNameAndSpaceReturns
	fake.recordInvocation("AppExistsWithNameAndSpace", []interface{}{arg1, arg2, arg3, arg4})
	fake.appExistsWithNameAndSpaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) AppExistsWithNameAndSpaceCallCount() int {
	fake.appExistsWithNameAndSpaceMutex.RLock()
	defer fake.appExistsWithNameAndSpaceMutex.RUnlock()
	return len(fake.appExistsWithNameAndSpaceArgsForCall)
}

func (fake *CFAppRepository) AppExistsWithNameAndSpaceCalls(stub func(context.Context, client.Client, string, string) (bool, error)) {
	fake.appExistsWithNameAndSpaceMutex.Lock()
	defer fake.appExistsWithNameAndSpaceMutex.Unlock()
	fake.AppExistsWithNameAndSpaceStub = stub
}

func (fake *CFAppRepository) AppExistsWithNameAndSpaceArgsForCall(i int) (context.Context, client.Client, string, string) {
	fake.appExistsWithNameAndSpaceMutex.RLock()
	defer fake.appExistsWithNameAndSpaceMutex.RUnlock()
	argsForCall := fake.appExistsWithNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *CFAppRepository) AppExistsWithNameAndSpaceReturns(result1 bool, result2 error) {
	fake.appExistsWithNameAndSpaceMutex.Lock()
	defer fake.appExistsWithNameAndSpaceMutex.Unlock()
	fake.AppExistsWithNameAndSpaceStub = nil
	fake.appExistsWithNameAndSpaceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) AppExistsWithNameAndSpaceReturnsOnCall(i int, result1 bool, result2 error) {
	fake.appExistsWithNameAndSpaceMutex.Lock()
	defer fake.appExistsWithNameAndSpaceMutex.Unlock()
	fake.AppExistsWithNameAndSpaceStub = nil
	if fake.appExistsWithNameAndSpaceReturnsOnCall == nil {
		fake.appExistsWithNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.appExistsWithNameAndSpaceReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) CreateApp(arg1 context.Context, arg2 client.Client, arg3 repositories.AppCreateMessage) (repositories.AppRecord, error) {
	fake.createAppMutex.Lock()
	ret, specificReturn := fake.createAppReturnsOnCall[len(fake.createAppArgsForCall)]
	fake.createAppArgsForCall = append(fake.createAppArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.AppCreateMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateAppStub
	fakeReturns := fake.createAppReturns
	fake.recordInvocation("CreateApp", []interface{}{arg1, arg2, arg3})
	fake.createAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) CreateAppCallCount() int {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	return len(fake.createAppArgsForCall)
}

func (fake *CFAppRepository) CreateAppCalls(stub func(context.Context, client.Client, repositories.AppCreateMessage) (repositories.AppRecord, error)) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = stub
}

func (fake *CFAppRepository) CreateAppArgsForCall(i int) (context.Context, client.Client, repositories.AppCreateMessage) {
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	argsForCall := fake.createAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) CreateAppReturns(result1 repositories.AppRecord, result2 error) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = nil
	fake.createAppReturns = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) CreateAppReturnsOnCall(i int, result1 repositories.AppRecord, result2 error) {
	fake.createAppMutex.Lock()
	defer fake.createAppMutex.Unlock()
	fake.CreateAppStub = nil
	if fake.createAppReturnsOnCall == nil {
		fake.createAppReturnsOnCall = make(map[int]struct {
			result1 repositories.AppRecord
			result2 error
		})
	}
	fake.createAppReturnsOnCall[i] = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) CreateAppEnvironmentVariables(arg1 context.Context, arg2 client.Client, arg3 repositories.AppEnvVarsRecord) (repositories.AppEnvVarsRecord, error) {
	fake.createAppEnvironmentVariablesMutex.Lock()
	ret, specificReturn := fake.createAppEnvironmentVariablesReturnsOnCall[len(fake.createAppEnvironmentVariablesArgsForCall)]
	fake.createAppEnvironmentVariablesArgsForCall = append(fake.createAppEnvironmentVariablesArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.AppEnvVarsRecord
	}{arg1, arg2, arg3})
	stub := fake.CreateAppEnvironmentVariablesStub
	fakeReturns := fake.createAppEnvironmentVariablesReturns
	fake.recordInvocation("CreateAppEnvironmentVariables", []interface{}{arg1, arg2, arg3})
	fake.createAppEnvironmentVariablesMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) CreateAppEnvironmentVariablesCallCount() int {
	fake.createAppEnvironmentVariablesMutex.RLock()
	defer fake.createAppEnvironmentVariablesMutex.RUnlock()
	return len(fake.createAppEnvironmentVariablesArgsForCall)
}

func (fake *CFAppRepository) CreateAppEnvironmentVariablesCalls(stub func(context.Context, client.Client, repositories.AppEnvVarsRecord) (repositories.AppEnvVarsRecord, error)) {
	fake.createAppEnvironmentVariablesMutex.Lock()
	defer fake.createAppEnvironmentVariablesMutex.Unlock()
	fake.CreateAppEnvironmentVariablesStub = stub
}

func (fake *CFAppRepository) CreateAppEnvironmentVariablesArgsForCall(i int) (context.Context, client.Client, repositories.AppEnvVarsRecord) {
	fake.createAppEnvironmentVariablesMutex.RLock()
	defer fake.createAppEnvironmentVariablesMutex.RUnlock()
	argsForCall := fake.createAppEnvironmentVariablesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) CreateAppEnvironmentVariablesReturns(result1 repositories.AppEnvVarsRecord, result2 error) {
	fake.createAppEnvironmentVariablesMutex.Lock()
	defer fake.createAppEnvironmentVariablesMutex.Unlock()
	fake.CreateAppEnvironmentVariablesStub = nil
	fake.createAppEnvironmentVariablesReturns = struct {
		result1 repositories.AppEnvVarsRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) CreateAppEnvironmentVariablesReturnsOnCall(i int, result1 repositories.AppEnvVarsRecord, result2 error) {
	fake.createAppEnvironmentVariablesMutex.Lock()
	defer fake.createAppEnvironmentVariablesMutex.Unlock()
	fake.CreateAppEnvironmentVariablesStub = nil
	if fake.createAppEnvironmentVariablesReturnsOnCall == nil {
		fake.createAppEnvironmentVariablesReturnsOnCall = make(map[int]struct {
			result1 repositories.AppEnvVarsRecord
			result2 error
		})
	}
	fake.createAppEnvironmentVariablesReturnsOnCall[i] = struct {
		result1 repositories.AppEnvVarsRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchApp(arg1 context.Context, arg2 client.Client, arg3 string) (repositories.AppRecord, error) {
	fake.fetchAppMutex.Lock()
	ret, specificReturn := fake.fetchAppReturnsOnCall[len(fake.fetchAppArgsForCall)]
	fake.fetchAppArgsForCall = append(fake.fetchAppArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FetchAppStub
	fakeReturns := fake.fetchAppReturns
	fake.recordInvocation("FetchApp", []interface{}{arg1, arg2, arg3})
	fake.fetchAppMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) FetchAppCallCount() int {
	fake.fetchAppMutex.RLock()
	defer fake.fetchAppMutex.RUnlock()
	return len(fake.fetchAppArgsForCall)
}

func (fake *CFAppRepository) FetchAppCalls(stub func(context.Context, client.Client, string) (repositories.AppRecord, error)) {
	fake.fetchAppMutex.Lock()
	defer fake.fetchAppMutex.Unlock()
	fake.FetchAppStub = stub
}

func (fake *CFAppRepository) FetchAppArgsForCall(i int) (context.Context, client.Client, string) {
	fake.fetchAppMutex.RLock()
	defer fake.fetchAppMutex.RUnlock()
	argsForCall := fake.fetchAppArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) FetchAppReturns(result1 repositories.AppRecord, result2 error) {
	fake.fetchAppMutex.Lock()
	defer fake.fetchAppMutex.Unlock()
	fake.FetchAppStub = nil
	fake.fetchAppReturns = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchAppReturnsOnCall(i int, result1 repositories.AppRecord, result2 error) {
	fake.fetchAppMutex.Lock()
	defer fake.fetchAppMutex.Unlock()
	fake.FetchAppStub = nil
	if fake.fetchAppReturnsOnCall == nil {
		fake.fetchAppReturnsOnCall = make(map[int]struct {
			result1 repositories.AppRecord
			result2 error
		})
	}
	fake.fetchAppReturnsOnCall[i] = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchAppList(arg1 context.Context, arg2 client.Client) ([]repositories.AppRecord, error) {
	fake.fetchAppListMutex.Lock()
	ret, specificReturn := fake.fetchAppListReturnsOnCall[len(fake.fetchAppListArgsForCall)]
	fake.fetchAppListArgsForCall = append(fake.fetchAppListArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
	}{arg1, arg2})
	stub := fake.FetchAppListStub
	fakeReturns := fake.fetchAppListReturns
	fake.recordInvocation("FetchAppList", []interface{}{arg1, arg2})
	fake.fetchAppListMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) FetchAppListCallCount() int {
	fake.fetchAppListMutex.RLock()
	defer fake.fetchAppListMutex.RUnlock()
	return len(fake.fetchAppListArgsForCall)
}

func (fake *CFAppRepository) FetchAppListCalls(stub func(context.Context, client.Client) ([]repositories.AppRecord, error)) {
	fake.fetchAppListMutex.Lock()
	defer fake.fetchAppListMutex.Unlock()
	fake.FetchAppListStub = stub
}

func (fake *CFAppRepository) FetchAppListArgsForCall(i int) (context.Context, client.Client) {
	fake.fetchAppListMutex.RLock()
	defer fake.fetchAppListMutex.RUnlock()
	argsForCall := fake.fetchAppListArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *CFAppRepository) FetchAppListReturns(result1 []repositories.AppRecord, result2 error) {
	fake.fetchAppListMutex.Lock()
	defer fake.fetchAppListMutex.Unlock()
	fake.FetchAppListStub = nil
	fake.fetchAppListReturns = struct {
		result1 []repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchAppListReturnsOnCall(i int, result1 []repositories.AppRecord, result2 error) {
	fake.fetchAppListMutex.Lock()
	defer fake.fetchAppListMutex.Unlock()
	fake.FetchAppListStub = nil
	if fake.fetchAppListReturnsOnCall == nil {
		fake.fetchAppListReturnsOnCall = make(map[int]struct {
			result1 []repositories.AppRecord
			result2 error
		})
	}
	fake.fetchAppListReturnsOnCall[i] = struct {
		result1 []repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchNamespace(arg1 context.Context, arg2 client.Client, arg3 string) (repositories.SpaceRecord, error) {
	fake.fetchNamespaceMutex.Lock()
	ret, specificReturn := fake.fetchNamespaceReturnsOnCall[len(fake.fetchNamespaceArgsForCall)]
	fake.fetchNamespaceArgsForCall = append(fake.fetchNamespaceArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FetchNamespaceStub
	fakeReturns := fake.fetchNamespaceReturns
	fake.recordInvocation("FetchNamespace", []interface{}{arg1, arg2, arg3})
	fake.fetchNamespaceMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) FetchNamespaceCallCount() int {
	fake.fetchNamespaceMutex.RLock()
	defer fake.fetchNamespaceMutex.RUnlock()
	return len(fake.fetchNamespaceArgsForCall)
}

func (fake *CFAppRepository) FetchNamespaceCalls(stub func(context.Context, client.Client, string) (repositories.SpaceRecord, error)) {
	fake.fetchNamespaceMutex.Lock()
	defer fake.fetchNamespaceMutex.Unlock()
	fake.FetchNamespaceStub = stub
}

func (fake *CFAppRepository) FetchNamespaceArgsForCall(i int) (context.Context, client.Client, string) {
	fake.fetchNamespaceMutex.RLock()
	defer fake.fetchNamespaceMutex.RUnlock()
	argsForCall := fake.fetchNamespaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) FetchNamespaceReturns(result1 repositories.SpaceRecord, result2 error) {
	fake.fetchNamespaceMutex.Lock()
	defer fake.fetchNamespaceMutex.Unlock()
	fake.FetchNamespaceStub = nil
	fake.fetchNamespaceReturns = struct {
		result1 repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) FetchNamespaceReturnsOnCall(i int, result1 repositories.SpaceRecord, result2 error) {
	fake.fetchNamespaceMutex.Lock()
	defer fake.fetchNamespaceMutex.Unlock()
	fake.FetchNamespaceStub = nil
	if fake.fetchNamespaceReturnsOnCall == nil {
		fake.fetchNamespaceReturnsOnCall = make(map[int]struct {
			result1 repositories.SpaceRecord
			result2 error
		})
	}
	fake.fetchNamespaceReturnsOnCall[i] = struct {
		result1 repositories.SpaceRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) SetAppDesiredState(arg1 context.Context, arg2 client.Client, arg3 repositories.SetAppDesiredStateMessage) (repositories.AppRecord, error) {
	fake.setAppDesiredStateMutex.Lock()
	ret, specificReturn := fake.setAppDesiredStateReturnsOnCall[len(fake.setAppDesiredStateArgsForCall)]
	fake.setAppDesiredStateArgsForCall = append(fake.setAppDesiredStateArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.SetAppDesiredStateMessage
	}{arg1, arg2, arg3})
	stub := fake.SetAppDesiredStateStub
	fakeReturns := fake.setAppDesiredStateReturns
	fake.recordInvocation("SetAppDesiredState", []interface{}{arg1, arg2, arg3})
	fake.setAppDesiredStateMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) SetAppDesiredStateCallCount() int {
	fake.setAppDesiredStateMutex.RLock()
	defer fake.setAppDesiredStateMutex.RUnlock()
	return len(fake.setAppDesiredStateArgsForCall)
}

func (fake *CFAppRepository) SetAppDesiredStateCalls(stub func(context.Context, client.Client, repositories.SetAppDesiredStateMessage) (repositories.AppRecord, error)) {
	fake.setAppDesiredStateMutex.Lock()
	defer fake.setAppDesiredStateMutex.Unlock()
	fake.SetAppDesiredStateStub = stub
}

func (fake *CFAppRepository) SetAppDesiredStateArgsForCall(i int) (context.Context, client.Client, repositories.SetAppDesiredStateMessage) {
	fake.setAppDesiredStateMutex.RLock()
	defer fake.setAppDesiredStateMutex.RUnlock()
	argsForCall := fake.setAppDesiredStateArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) SetAppDesiredStateReturns(result1 repositories.AppRecord, result2 error) {
	fake.setAppDesiredStateMutex.Lock()
	defer fake.setAppDesiredStateMutex.Unlock()
	fake.SetAppDesiredStateStub = nil
	fake.setAppDesiredStateReturns = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) SetAppDesiredStateReturnsOnCall(i int, result1 repositories.AppRecord, result2 error) {
	fake.setAppDesiredStateMutex.Lock()
	defer fake.setAppDesiredStateMutex.Unlock()
	fake.SetAppDesiredStateStub = nil
	if fake.setAppDesiredStateReturnsOnCall == nil {
		fake.setAppDesiredStateReturnsOnCall = make(map[int]struct {
			result1 repositories.AppRecord
			result2 error
		})
	}
	fake.setAppDesiredStateReturnsOnCall[i] = struct {
		result1 repositories.AppRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) SetCurrentDroplet(arg1 context.Context, arg2 client.Client, arg3 repositories.SetCurrentDropletMessage) (repositories.CurrentDropletRecord, error) {
	fake.setCurrentDropletMutex.Lock()
	ret, specificReturn := fake.setCurrentDropletReturnsOnCall[len(fake.setCurrentDropletArgsForCall)]
	fake.setCurrentDropletArgsForCall = append(fake.setCurrentDropletArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.SetCurrentDropletMessage
	}{arg1, arg2, arg3})
	stub := fake.SetCurrentDropletStub
	fakeReturns := fake.setCurrentDropletReturns
	fake.recordInvocation("SetCurrentDroplet", []interface{}{arg1, arg2, arg3})
	fake.setCurrentDropletMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFAppRepository) SetCurrentDropletCallCount() int {
	fake.setCurrentDropletMutex.RLock()
	defer fake.setCurrentDropletMutex.RUnlock()
	return len(fake.setCurrentDropletArgsForCall)
}

func (fake *CFAppRepository) SetCurrentDropletCalls(stub func(context.Context, client.Client, repositories.SetCurrentDropletMessage) (repositories.CurrentDropletRecord, error)) {
	fake.setCurrentDropletMutex.Lock()
	defer fake.setCurrentDropletMutex.Unlock()
	fake.SetCurrentDropletStub = stub
}

func (fake *CFAppRepository) SetCurrentDropletArgsForCall(i int) (context.Context, client.Client, repositories.SetCurrentDropletMessage) {
	fake.setCurrentDropletMutex.RLock()
	defer fake.setCurrentDropletMutex.RUnlock()
	argsForCall := fake.setCurrentDropletArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFAppRepository) SetCurrentDropletReturns(result1 repositories.CurrentDropletRecord, result2 error) {
	fake.setCurrentDropletMutex.Lock()
	defer fake.setCurrentDropletMutex.Unlock()
	fake.SetCurrentDropletStub = nil
	fake.setCurrentDropletReturns = struct {
		result1 repositories.CurrentDropletRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) SetCurrentDropletReturnsOnCall(i int, result1 repositories.CurrentDropletRecord, result2 error) {
	fake.setCurrentDropletMutex.Lock()
	defer fake.setCurrentDropletMutex.Unlock()
	fake.SetCurrentDropletStub = nil
	if fake.setCurrentDropletReturnsOnCall == nil {
		fake.setCurrentDropletReturnsOnCall = make(map[int]struct {
			result1 repositories.CurrentDropletRecord
			result2 error
		})
	}
	fake.setCurrentDropletReturnsOnCall[i] = struct {
		result1 repositories.CurrentDropletRecord
		result2 error
	}{result1, result2}
}

func (fake *CFAppRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appExistsWithNameAndSpaceMutex.RLock()
	defer fake.appExistsWithNameAndSpaceMutex.RUnlock()
	fake.createAppMutex.RLock()
	defer fake.createAppMutex.RUnlock()
	fake.createAppEnvironmentVariablesMutex.RLock()
	defer fake.createAppEnvironmentVariablesMutex.RUnlock()
	fake.fetchAppMutex.RLock()
	defer fake.fetchAppMutex.RUnlock()
	fake.fetchAppListMutex.RLock()
	defer fake.fetchAppListMutex.RUnlock()
	fake.fetchNamespaceMutex.RLock()
	defer fake.fetchNamespaceMutex.RUnlock()
	fake.setAppDesiredStateMutex.RLock()
	defer fake.setAppDesiredStateMutex.RUnlock()
	fake.setCurrentDropletMutex.RLock()
	defer fake.setCurrentDropletMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFAppRepository) recordInvocation(key string, args []interface{}) {
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

var _ actions.CFAppRepository = new(CFAppRepository)
