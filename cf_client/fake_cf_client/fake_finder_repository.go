// Code generated by counterfeiter. DO NOT EDIT.
package fake_cf_client

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/models"
	"github.com/orange-cloudfoundry/terraform-provider-cloudfoundry/cf_client"
)

type FakeFinderRepository struct {
	GetDomainFromCfStub        func(domain models.DomainFields) (models.DomainFields, error)
	getDomainFromCfMutex       sync.RWMutex
	getDomainFromCfArgsForCall []struct {
		domain models.DomainFields
	}
	getDomainFromCfReturns struct {
		result1 models.DomainFields
		result2 error
	}
	getDomainFromCfReturnsOnCall map[int]struct {
		result1 models.DomainFields
		result2 error
	}
	GetBuildpackFromCfStub        func(bpGuid string) (models.Buildpack, error)
	getBuildpackFromCfMutex       sync.RWMutex
	getBuildpackFromCfArgsForCall []struct {
		bpGuid string
	}
	getBuildpackFromCfReturns struct {
		result1 models.Buildpack
		result2 error
	}
	getBuildpackFromCfReturnsOnCall map[int]struct {
		result1 models.Buildpack
		result2 error
	}
	GetQuotaFromCfStub        func(quotaGuid string, isOrgQuota bool) (interface{}, error)
	getQuotaFromCfMutex       sync.RWMutex
	getQuotaFromCfArgsForCall []struct {
		quotaGuid  string
		isOrgQuota bool
	}
	getQuotaFromCfReturns struct {
		result1 interface{}
		result2 error
	}
	getQuotaFromCfReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	GetRouteFromCfStub        func(routeGuid string) (models.Route, error)
	getRouteFromCfMutex       sync.RWMutex
	getRouteFromCfArgsForCall []struct {
		routeGuid string
	}
	getRouteFromCfReturns struct {
		result1 models.Route
		result2 error
	}
	getRouteFromCfReturnsOnCall map[int]struct {
		result1 models.Route
		result2 error
	}
	GetSecGroupFromCfStub        func(secGroupId string) (models.SecurityGroup, error)
	getSecGroupFromCfMutex       sync.RWMutex
	getSecGroupFromCfArgsForCall []struct {
		secGroupId string
	}
	getSecGroupFromCfReturns struct {
		result1 models.SecurityGroup
		result2 error
	}
	getSecGroupFromCfReturnsOnCall map[int]struct {
		result1 models.SecurityGroup
		result2 error
	}
	GetServiceFromCfStub        func(svcGuid string) (models.ServiceInstance, error)
	getServiceFromCfMutex       sync.RWMutex
	getServiceFromCfArgsForCall []struct {
		svcGuid string
	}
	getServiceFromCfReturns struct {
		result1 models.ServiceInstance
		result2 error
	}
	getServiceFromCfReturnsOnCall map[int]struct {
		result1 models.ServiceInstance
		result2 error
	}
	GetSpaceFromCfStub        func(spaceGuid string) (models.Space, error)
	getSpaceFromCfMutex       sync.RWMutex
	getSpaceFromCfArgsForCall []struct {
		spaceGuid string
	}
	getSpaceFromCfReturns struct {
		result1 models.Space
		result2 error
	}
	getSpaceFromCfReturnsOnCall map[int]struct {
		result1 models.Space
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFinderRepository) GetDomainFromCf(domain models.DomainFields) (models.DomainFields, error) {
	fake.getDomainFromCfMutex.Lock()
	ret, specificReturn := fake.getDomainFromCfReturnsOnCall[len(fake.getDomainFromCfArgsForCall)]
	fake.getDomainFromCfArgsForCall = append(fake.getDomainFromCfArgsForCall, struct {
		domain models.DomainFields
	}{domain})
	fake.recordInvocation("GetDomainFromCf", []interface{}{domain})
	fake.getDomainFromCfMutex.Unlock()
	if fake.GetDomainFromCfStub != nil {
		return fake.GetDomainFromCfStub(domain)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getDomainFromCfReturns.result1, fake.getDomainFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetDomainFromCfCallCount() int {
	fake.getDomainFromCfMutex.RLock()
	defer fake.getDomainFromCfMutex.RUnlock()
	return len(fake.getDomainFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetDomainFromCfArgsForCall(i int) models.DomainFields {
	fake.getDomainFromCfMutex.RLock()
	defer fake.getDomainFromCfMutex.RUnlock()
	return fake.getDomainFromCfArgsForCall[i].domain
}

func (fake *FakeFinderRepository) GetDomainFromCfReturns(result1 models.DomainFields, result2 error) {
	fake.GetDomainFromCfStub = nil
	fake.getDomainFromCfReturns = struct {
		result1 models.DomainFields
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetDomainFromCfReturnsOnCall(i int, result1 models.DomainFields, result2 error) {
	fake.GetDomainFromCfStub = nil
	if fake.getDomainFromCfReturnsOnCall == nil {
		fake.getDomainFromCfReturnsOnCall = make(map[int]struct {
			result1 models.DomainFields
			result2 error
		})
	}
	fake.getDomainFromCfReturnsOnCall[i] = struct {
		result1 models.DomainFields
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetBuildpackFromCf(bpGuid string) (models.Buildpack, error) {
	fake.getBuildpackFromCfMutex.Lock()
	ret, specificReturn := fake.getBuildpackFromCfReturnsOnCall[len(fake.getBuildpackFromCfArgsForCall)]
	fake.getBuildpackFromCfArgsForCall = append(fake.getBuildpackFromCfArgsForCall, struct {
		bpGuid string
	}{bpGuid})
	fake.recordInvocation("GetBuildpackFromCf", []interface{}{bpGuid})
	fake.getBuildpackFromCfMutex.Unlock()
	if fake.GetBuildpackFromCfStub != nil {
		return fake.GetBuildpackFromCfStub(bpGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getBuildpackFromCfReturns.result1, fake.getBuildpackFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetBuildpackFromCfCallCount() int {
	fake.getBuildpackFromCfMutex.RLock()
	defer fake.getBuildpackFromCfMutex.RUnlock()
	return len(fake.getBuildpackFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetBuildpackFromCfArgsForCall(i int) string {
	fake.getBuildpackFromCfMutex.RLock()
	defer fake.getBuildpackFromCfMutex.RUnlock()
	return fake.getBuildpackFromCfArgsForCall[i].bpGuid
}

func (fake *FakeFinderRepository) GetBuildpackFromCfReturns(result1 models.Buildpack, result2 error) {
	fake.GetBuildpackFromCfStub = nil
	fake.getBuildpackFromCfReturns = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetBuildpackFromCfReturnsOnCall(i int, result1 models.Buildpack, result2 error) {
	fake.GetBuildpackFromCfStub = nil
	if fake.getBuildpackFromCfReturnsOnCall == nil {
		fake.getBuildpackFromCfReturnsOnCall = make(map[int]struct {
			result1 models.Buildpack
			result2 error
		})
	}
	fake.getBuildpackFromCfReturnsOnCall[i] = struct {
		result1 models.Buildpack
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetQuotaFromCf(quotaGuid string, isOrgQuota bool) (interface{}, error) {
	fake.getQuotaFromCfMutex.Lock()
	ret, specificReturn := fake.getQuotaFromCfReturnsOnCall[len(fake.getQuotaFromCfArgsForCall)]
	fake.getQuotaFromCfArgsForCall = append(fake.getQuotaFromCfArgsForCall, struct {
		quotaGuid  string
		isOrgQuota bool
	}{quotaGuid, isOrgQuota})
	fake.recordInvocation("GetQuotaFromCf", []interface{}{quotaGuid, isOrgQuota})
	fake.getQuotaFromCfMutex.Unlock()
	if fake.GetQuotaFromCfStub != nil {
		return fake.GetQuotaFromCfStub(quotaGuid, isOrgQuota)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getQuotaFromCfReturns.result1, fake.getQuotaFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetQuotaFromCfCallCount() int {
	fake.getQuotaFromCfMutex.RLock()
	defer fake.getQuotaFromCfMutex.RUnlock()
	return len(fake.getQuotaFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetQuotaFromCfArgsForCall(i int) (string, bool) {
	fake.getQuotaFromCfMutex.RLock()
	defer fake.getQuotaFromCfMutex.RUnlock()
	return fake.getQuotaFromCfArgsForCall[i].quotaGuid, fake.getQuotaFromCfArgsForCall[i].isOrgQuota
}

func (fake *FakeFinderRepository) GetQuotaFromCfReturns(result1 interface{}, result2 error) {
	fake.GetQuotaFromCfStub = nil
	fake.getQuotaFromCfReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetQuotaFromCfReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.GetQuotaFromCfStub = nil
	if fake.getQuotaFromCfReturnsOnCall == nil {
		fake.getQuotaFromCfReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.getQuotaFromCfReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetRouteFromCf(routeGuid string) (models.Route, error) {
	fake.getRouteFromCfMutex.Lock()
	ret, specificReturn := fake.getRouteFromCfReturnsOnCall[len(fake.getRouteFromCfArgsForCall)]
	fake.getRouteFromCfArgsForCall = append(fake.getRouteFromCfArgsForCall, struct {
		routeGuid string
	}{routeGuid})
	fake.recordInvocation("GetRouteFromCf", []interface{}{routeGuid})
	fake.getRouteFromCfMutex.Unlock()
	if fake.GetRouteFromCfStub != nil {
		return fake.GetRouteFromCfStub(routeGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getRouteFromCfReturns.result1, fake.getRouteFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetRouteFromCfCallCount() int {
	fake.getRouteFromCfMutex.RLock()
	defer fake.getRouteFromCfMutex.RUnlock()
	return len(fake.getRouteFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetRouteFromCfArgsForCall(i int) string {
	fake.getRouteFromCfMutex.RLock()
	defer fake.getRouteFromCfMutex.RUnlock()
	return fake.getRouteFromCfArgsForCall[i].routeGuid
}

func (fake *FakeFinderRepository) GetRouteFromCfReturns(result1 models.Route, result2 error) {
	fake.GetRouteFromCfStub = nil
	fake.getRouteFromCfReturns = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetRouteFromCfReturnsOnCall(i int, result1 models.Route, result2 error) {
	fake.GetRouteFromCfStub = nil
	if fake.getRouteFromCfReturnsOnCall == nil {
		fake.getRouteFromCfReturnsOnCall = make(map[int]struct {
			result1 models.Route
			result2 error
		})
	}
	fake.getRouteFromCfReturnsOnCall[i] = struct {
		result1 models.Route
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetSecGroupFromCf(secGroupId string) (models.SecurityGroup, error) {
	fake.getSecGroupFromCfMutex.Lock()
	ret, specificReturn := fake.getSecGroupFromCfReturnsOnCall[len(fake.getSecGroupFromCfArgsForCall)]
	fake.getSecGroupFromCfArgsForCall = append(fake.getSecGroupFromCfArgsForCall, struct {
		secGroupId string
	}{secGroupId})
	fake.recordInvocation("GetSecGroupFromCf", []interface{}{secGroupId})
	fake.getSecGroupFromCfMutex.Unlock()
	if fake.GetSecGroupFromCfStub != nil {
		return fake.GetSecGroupFromCfStub(secGroupId)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSecGroupFromCfReturns.result1, fake.getSecGroupFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetSecGroupFromCfCallCount() int {
	fake.getSecGroupFromCfMutex.RLock()
	defer fake.getSecGroupFromCfMutex.RUnlock()
	return len(fake.getSecGroupFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetSecGroupFromCfArgsForCall(i int) string {
	fake.getSecGroupFromCfMutex.RLock()
	defer fake.getSecGroupFromCfMutex.RUnlock()
	return fake.getSecGroupFromCfArgsForCall[i].secGroupId
}

func (fake *FakeFinderRepository) GetSecGroupFromCfReturns(result1 models.SecurityGroup, result2 error) {
	fake.GetSecGroupFromCfStub = nil
	fake.getSecGroupFromCfReturns = struct {
		result1 models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetSecGroupFromCfReturnsOnCall(i int, result1 models.SecurityGroup, result2 error) {
	fake.GetSecGroupFromCfStub = nil
	if fake.getSecGroupFromCfReturnsOnCall == nil {
		fake.getSecGroupFromCfReturnsOnCall = make(map[int]struct {
			result1 models.SecurityGroup
			result2 error
		})
	}
	fake.getSecGroupFromCfReturnsOnCall[i] = struct {
		result1 models.SecurityGroup
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetServiceFromCf(svcGuid string) (models.ServiceInstance, error) {
	fake.getServiceFromCfMutex.Lock()
	ret, specificReturn := fake.getServiceFromCfReturnsOnCall[len(fake.getServiceFromCfArgsForCall)]
	fake.getServiceFromCfArgsForCall = append(fake.getServiceFromCfArgsForCall, struct {
		svcGuid string
	}{svcGuid})
	fake.recordInvocation("GetServiceFromCf", []interface{}{svcGuid})
	fake.getServiceFromCfMutex.Unlock()
	if fake.GetServiceFromCfStub != nil {
		return fake.GetServiceFromCfStub(svcGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getServiceFromCfReturns.result1, fake.getServiceFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetServiceFromCfCallCount() int {
	fake.getServiceFromCfMutex.RLock()
	defer fake.getServiceFromCfMutex.RUnlock()
	return len(fake.getServiceFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetServiceFromCfArgsForCall(i int) string {
	fake.getServiceFromCfMutex.RLock()
	defer fake.getServiceFromCfMutex.RUnlock()
	return fake.getServiceFromCfArgsForCall[i].svcGuid
}

func (fake *FakeFinderRepository) GetServiceFromCfReturns(result1 models.ServiceInstance, result2 error) {
	fake.GetServiceFromCfStub = nil
	fake.getServiceFromCfReturns = struct {
		result1 models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetServiceFromCfReturnsOnCall(i int, result1 models.ServiceInstance, result2 error) {
	fake.GetServiceFromCfStub = nil
	if fake.getServiceFromCfReturnsOnCall == nil {
		fake.getServiceFromCfReturnsOnCall = make(map[int]struct {
			result1 models.ServiceInstance
			result2 error
		})
	}
	fake.getServiceFromCfReturnsOnCall[i] = struct {
		result1 models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetSpaceFromCf(spaceGuid string) (models.Space, error) {
	fake.getSpaceFromCfMutex.Lock()
	ret, specificReturn := fake.getSpaceFromCfReturnsOnCall[len(fake.getSpaceFromCfArgsForCall)]
	fake.getSpaceFromCfArgsForCall = append(fake.getSpaceFromCfArgsForCall, struct {
		spaceGuid string
	}{spaceGuid})
	fake.recordInvocation("GetSpaceFromCf", []interface{}{spaceGuid})
	fake.getSpaceFromCfMutex.Unlock()
	if fake.GetSpaceFromCfStub != nil {
		return fake.GetSpaceFromCfStub(spaceGuid)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceFromCfReturns.result1, fake.getSpaceFromCfReturns.result2
}

func (fake *FakeFinderRepository) GetSpaceFromCfCallCount() int {
	fake.getSpaceFromCfMutex.RLock()
	defer fake.getSpaceFromCfMutex.RUnlock()
	return len(fake.getSpaceFromCfArgsForCall)
}

func (fake *FakeFinderRepository) GetSpaceFromCfArgsForCall(i int) string {
	fake.getSpaceFromCfMutex.RLock()
	defer fake.getSpaceFromCfMutex.RUnlock()
	return fake.getSpaceFromCfArgsForCall[i].spaceGuid
}

func (fake *FakeFinderRepository) GetSpaceFromCfReturns(result1 models.Space, result2 error) {
	fake.GetSpaceFromCfStub = nil
	fake.getSpaceFromCfReturns = struct {
		result1 models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) GetSpaceFromCfReturnsOnCall(i int, result1 models.Space, result2 error) {
	fake.GetSpaceFromCfStub = nil
	if fake.getSpaceFromCfReturnsOnCall == nil {
		fake.getSpaceFromCfReturnsOnCall = make(map[int]struct {
			result1 models.Space
			result2 error
		})
	}
	fake.getSpaceFromCfReturnsOnCall[i] = struct {
		result1 models.Space
		result2 error
	}{result1, result2}
}

func (fake *FakeFinderRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getDomainFromCfMutex.RLock()
	defer fake.getDomainFromCfMutex.RUnlock()
	fake.getBuildpackFromCfMutex.RLock()
	defer fake.getBuildpackFromCfMutex.RUnlock()
	fake.getQuotaFromCfMutex.RLock()
	defer fake.getQuotaFromCfMutex.RUnlock()
	fake.getRouteFromCfMutex.RLock()
	defer fake.getRouteFromCfMutex.RUnlock()
	fake.getSecGroupFromCfMutex.RLock()
	defer fake.getSecGroupFromCfMutex.RUnlock()
	fake.getServiceFromCfMutex.RLock()
	defer fake.getServiceFromCfMutex.RUnlock()
	fake.getSpaceFromCfMutex.RLock()
	defer fake.getSpaceFromCfMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFinderRepository) recordInvocation(key string, args []interface{}) {
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

var _ cf_client.FinderRepository = new(FakeFinderRepository)
