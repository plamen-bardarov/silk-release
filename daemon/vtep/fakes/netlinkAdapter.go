// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"

	"github.com/vishvananda/netlink"
)

type NetlinkAdapter struct {
	LinkSetUpStub        func(netlink.Link) error
	linkSetUpMutex       sync.RWMutex
	linkSetUpArgsForCall []struct {
		arg1 netlink.Link
	}
	linkSetUpReturns struct {
		result1 error
	}
	linkSetUpReturnsOnCall map[int]struct {
		result1 error
	}
	LinkAddStub        func(netlink.Link) error
	linkAddMutex       sync.RWMutex
	linkAddArgsForCall []struct {
		arg1 netlink.Link
	}
	linkAddReturns struct {
		result1 error
	}
	linkAddReturnsOnCall map[int]struct {
		result1 error
	}
	LinkByNameStub        func(string) (netlink.Link, error)
	linkByNameMutex       sync.RWMutex
	linkByNameArgsForCall []struct {
		arg1 string
	}
	linkByNameReturns struct {
		result1 netlink.Link
		result2 error
	}
	linkByNameReturnsOnCall map[int]struct {
		result1 netlink.Link
		result2 error
	}
	LinkSetHardwareAddrStub        func(netlink.Link, net.HardwareAddr) error
	linkSetHardwareAddrMutex       sync.RWMutex
	linkSetHardwareAddrArgsForCall []struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}
	linkSetHardwareAddrReturns struct {
		result1 error
	}
	linkSetHardwareAddrReturnsOnCall map[int]struct {
		result1 error
	}
	AddrAddScopeLinkStub        func(link netlink.Link, addr *netlink.Addr) error
	addrAddScopeLinkMutex       sync.RWMutex
	addrAddScopeLinkArgsForCall []struct {
		link netlink.Link
		addr *netlink.Addr
	}
	addrAddScopeLinkReturns struct {
		result1 error
	}
	addrAddScopeLinkReturnsOnCall map[int]struct {
		result1 error
	}
	AddrListStub        func(link netlink.Link, family int) ([]netlink.Addr, error)
	addrListMutex       sync.RWMutex
	addrListArgsForCall []struct {
		link   netlink.Link
		family int
	}
	addrListReturns struct {
		result1 []netlink.Addr
		result2 error
	}
	addrListReturnsOnCall map[int]struct {
		result1 []netlink.Addr
		result2 error
	}
	RouteAddStub        func(*netlink.Route) error
	routeAddMutex       sync.RWMutex
	routeAddArgsForCall []struct {
		arg1 *netlink.Route
	}
	routeAddReturns struct {
		result1 error
	}
	routeAddReturnsOnCall map[int]struct {
		result1 error
	}
	LinkDelStub        func(netlink.Link) error
	linkDelMutex       sync.RWMutex
	linkDelArgsForCall []struct {
		arg1 netlink.Link
	}
	linkDelReturns struct {
		result1 error
	}
	linkDelReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetlinkAdapter) LinkSetUp(arg1 netlink.Link) error {
	fake.linkSetUpMutex.Lock()
	ret, specificReturn := fake.linkSetUpReturnsOnCall[len(fake.linkSetUpArgsForCall)]
	fake.linkSetUpArgsForCall = append(fake.linkSetUpArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkSetUp", []interface{}{arg1})
	fake.linkSetUpMutex.Unlock()
	if fake.LinkSetUpStub != nil {
		return fake.LinkSetUpStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetUpReturns.result1
}

func (fake *NetlinkAdapter) LinkSetUpCallCount() int {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return len(fake.linkSetUpArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetUpArgsForCall(i int) netlink.Link {
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	return fake.linkSetUpArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkSetUpReturns(result1 error) {
	fake.LinkSetUpStub = nil
	fake.linkSetUpReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetUpReturnsOnCall(i int, result1 error) {
	fake.LinkSetUpStub = nil
	if fake.linkSetUpReturnsOnCall == nil {
		fake.linkSetUpReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetUpReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAdd(arg1 netlink.Link) error {
	fake.linkAddMutex.Lock()
	ret, specificReturn := fake.linkAddReturnsOnCall[len(fake.linkAddArgsForCall)]
	fake.linkAddArgsForCall = append(fake.linkAddArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkAdd", []interface{}{arg1})
	fake.linkAddMutex.Unlock()
	if fake.LinkAddStub != nil {
		return fake.LinkAddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkAddReturns.result1
}

func (fake *NetlinkAdapter) LinkAddCallCount() int {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return len(fake.linkAddArgsForCall)
}

func (fake *NetlinkAdapter) LinkAddArgsForCall(i int) netlink.Link {
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	return fake.linkAddArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkAddReturns(result1 error) {
	fake.LinkAddStub = nil
	fake.linkAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkAddReturnsOnCall(i int, result1 error) {
	fake.LinkAddStub = nil
	if fake.linkAddReturnsOnCall == nil {
		fake.linkAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkByName(arg1 string) (netlink.Link, error) {
	fake.linkByNameMutex.Lock()
	ret, specificReturn := fake.linkByNameReturnsOnCall[len(fake.linkByNameArgsForCall)]
	fake.linkByNameArgsForCall = append(fake.linkByNameArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LinkByName", []interface{}{arg1})
	fake.linkByNameMutex.Unlock()
	if fake.LinkByNameStub != nil {
		return fake.LinkByNameStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.linkByNameReturns.result1, fake.linkByNameReturns.result2
}

func (fake *NetlinkAdapter) LinkByNameCallCount() int {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return len(fake.linkByNameArgsForCall)
}

func (fake *NetlinkAdapter) LinkByNameArgsForCall(i int) string {
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	return fake.linkByNameArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkByNameReturns(result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	fake.linkByNameReturns = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) LinkByNameReturnsOnCall(i int, result1 netlink.Link, result2 error) {
	fake.LinkByNameStub = nil
	if fake.linkByNameReturnsOnCall == nil {
		fake.linkByNameReturnsOnCall = make(map[int]struct {
			result1 netlink.Link
			result2 error
		})
	}
	fake.linkByNameReturnsOnCall[i] = struct {
		result1 netlink.Link
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddr(arg1 netlink.Link, arg2 net.HardwareAddr) error {
	fake.linkSetHardwareAddrMutex.Lock()
	ret, specificReturn := fake.linkSetHardwareAddrReturnsOnCall[len(fake.linkSetHardwareAddrArgsForCall)]
	fake.linkSetHardwareAddrArgsForCall = append(fake.linkSetHardwareAddrArgsForCall, struct {
		arg1 netlink.Link
		arg2 net.HardwareAddr
	}{arg1, arg2})
	fake.recordInvocation("LinkSetHardwareAddr", []interface{}{arg1, arg2})
	fake.linkSetHardwareAddrMutex.Unlock()
	if fake.LinkSetHardwareAddrStub != nil {
		return fake.LinkSetHardwareAddrStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkSetHardwareAddrReturns.result1
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrCallCount() int {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return len(fake.linkSetHardwareAddrArgsForCall)
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrArgsForCall(i int) (netlink.Link, net.HardwareAddr) {
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	return fake.linkSetHardwareAddrArgsForCall[i].arg1, fake.linkSetHardwareAddrArgsForCall[i].arg2
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturns(result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	fake.linkSetHardwareAddrReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkSetHardwareAddrReturnsOnCall(i int, result1 error) {
	fake.LinkSetHardwareAddrStub = nil
	if fake.linkSetHardwareAddrReturnsOnCall == nil {
		fake.linkSetHardwareAddrReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkSetHardwareAddrReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddScopeLink(link netlink.Link, addr *netlink.Addr) error {
	fake.addrAddScopeLinkMutex.Lock()
	ret, specificReturn := fake.addrAddScopeLinkReturnsOnCall[len(fake.addrAddScopeLinkArgsForCall)]
	fake.addrAddScopeLinkArgsForCall = append(fake.addrAddScopeLinkArgsForCall, struct {
		link netlink.Link
		addr *netlink.Addr
	}{link, addr})
	fake.recordInvocation("AddrAddScopeLink", []interface{}{link, addr})
	fake.addrAddScopeLinkMutex.Unlock()
	if fake.AddrAddScopeLinkStub != nil {
		return fake.AddrAddScopeLinkStub(link, addr)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addrAddScopeLinkReturns.result1
}

func (fake *NetlinkAdapter) AddrAddScopeLinkCallCount() int {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return len(fake.addrAddScopeLinkArgsForCall)
}

func (fake *NetlinkAdapter) AddrAddScopeLinkArgsForCall(i int) (netlink.Link, *netlink.Addr) {
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	return fake.addrAddScopeLinkArgsForCall[i].link, fake.addrAddScopeLinkArgsForCall[i].addr
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturns(result1 error) {
	fake.AddrAddScopeLinkStub = nil
	fake.addrAddScopeLinkReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrAddScopeLinkReturnsOnCall(i int, result1 error) {
	fake.AddrAddScopeLinkStub = nil
	if fake.addrAddScopeLinkReturnsOnCall == nil {
		fake.addrAddScopeLinkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addrAddScopeLinkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) AddrList(link netlink.Link, family int) ([]netlink.Addr, error) {
	fake.addrListMutex.Lock()
	ret, specificReturn := fake.addrListReturnsOnCall[len(fake.addrListArgsForCall)]
	fake.addrListArgsForCall = append(fake.addrListArgsForCall, struct {
		link   netlink.Link
		family int
	}{link, family})
	fake.recordInvocation("AddrList", []interface{}{link, family})
	fake.addrListMutex.Unlock()
	if fake.AddrListStub != nil {
		return fake.AddrListStub(link, family)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.addrListReturns.result1, fake.addrListReturns.result2
}

func (fake *NetlinkAdapter) AddrListCallCount() int {
	fake.addrListMutex.RLock()
	defer fake.addrListMutex.RUnlock()
	return len(fake.addrListArgsForCall)
}

func (fake *NetlinkAdapter) AddrListArgsForCall(i int) (netlink.Link, int) {
	fake.addrListMutex.RLock()
	defer fake.addrListMutex.RUnlock()
	return fake.addrListArgsForCall[i].link, fake.addrListArgsForCall[i].family
}

func (fake *NetlinkAdapter) AddrListReturns(result1 []netlink.Addr, result2 error) {
	fake.AddrListStub = nil
	fake.addrListReturns = struct {
		result1 []netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) AddrListReturnsOnCall(i int, result1 []netlink.Addr, result2 error) {
	fake.AddrListStub = nil
	if fake.addrListReturnsOnCall == nil {
		fake.addrListReturnsOnCall = make(map[int]struct {
			result1 []netlink.Addr
			result2 error
		})
	}
	fake.addrListReturnsOnCall[i] = struct {
		result1 []netlink.Addr
		result2 error
	}{result1, result2}
}

func (fake *NetlinkAdapter) RouteAdd(arg1 *netlink.Route) error {
	fake.routeAddMutex.Lock()
	ret, specificReturn := fake.routeAddReturnsOnCall[len(fake.routeAddArgsForCall)]
	fake.routeAddArgsForCall = append(fake.routeAddArgsForCall, struct {
		arg1 *netlink.Route
	}{arg1})
	fake.recordInvocation("RouteAdd", []interface{}{arg1})
	fake.routeAddMutex.Unlock()
	if fake.RouteAddStub != nil {
		return fake.RouteAddStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.routeAddReturns.result1
}

func (fake *NetlinkAdapter) RouteAddCallCount() int {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return len(fake.routeAddArgsForCall)
}

func (fake *NetlinkAdapter) RouteAddArgsForCall(i int) *netlink.Route {
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	return fake.routeAddArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) RouteAddReturns(result1 error) {
	fake.RouteAddStub = nil
	fake.routeAddReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) RouteAddReturnsOnCall(i int, result1 error) {
	fake.RouteAddStub = nil
	if fake.routeAddReturnsOnCall == nil {
		fake.routeAddReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.routeAddReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkDel(arg1 netlink.Link) error {
	fake.linkDelMutex.Lock()
	ret, specificReturn := fake.linkDelReturnsOnCall[len(fake.linkDelArgsForCall)]
	fake.linkDelArgsForCall = append(fake.linkDelArgsForCall, struct {
		arg1 netlink.Link
	}{arg1})
	fake.recordInvocation("LinkDel", []interface{}{arg1})
	fake.linkDelMutex.Unlock()
	if fake.LinkDelStub != nil {
		return fake.LinkDelStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.linkDelReturns.result1
}

func (fake *NetlinkAdapter) LinkDelCallCount() int {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return len(fake.linkDelArgsForCall)
}

func (fake *NetlinkAdapter) LinkDelArgsForCall(i int) netlink.Link {
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return fake.linkDelArgsForCall[i].arg1
}

func (fake *NetlinkAdapter) LinkDelReturns(result1 error) {
	fake.LinkDelStub = nil
	fake.linkDelReturns = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) LinkDelReturnsOnCall(i int, result1 error) {
	fake.LinkDelStub = nil
	if fake.linkDelReturnsOnCall == nil {
		fake.linkDelReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.linkDelReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *NetlinkAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.linkSetUpMutex.RLock()
	defer fake.linkSetUpMutex.RUnlock()
	fake.linkAddMutex.RLock()
	defer fake.linkAddMutex.RUnlock()
	fake.linkByNameMutex.RLock()
	defer fake.linkByNameMutex.RUnlock()
	fake.linkSetHardwareAddrMutex.RLock()
	defer fake.linkSetHardwareAddrMutex.RUnlock()
	fake.addrAddScopeLinkMutex.RLock()
	defer fake.addrAddScopeLinkMutex.RUnlock()
	fake.addrListMutex.RLock()
	defer fake.addrListMutex.RUnlock()
	fake.routeAddMutex.RLock()
	defer fake.routeAddMutex.RUnlock()
	fake.linkDelMutex.RLock()
	defer fake.linkDelMutex.RUnlock()
	return fake.invocations
}

func (fake *NetlinkAdapter) recordInvocation(key string, args []interface{}) {
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
