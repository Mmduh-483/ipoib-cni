// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import current "github.com/containernetworking/cni/pkg/types/current"
import mock "github.com/stretchr/testify/mock"
import ns "github.com/containernetworking/plugins/pkg/ns"
import types "github.com/Mellanox/ipoib-cni/pkg/types"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// CreateIpoibLink provides a mock function with given fields: conf, ifName, netns
func (_m *Manager) CreateIpoibLink(conf *types.NetConf, ifName string, netns ns.NetNS) (*current.Interface, error) {
	ret := _m.Called(conf, ifName, netns)

	var r0 *current.Interface
	if rf, ok := ret.Get(0).(func(*types.NetConf, string, ns.NetNS) *current.Interface); ok {
		r0 = rf(conf, ifName, netns)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*current.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*types.NetConf, string, ns.NetNS) error); ok {
		r1 = rf(conf, ifName, netns)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveIpoibLink provides a mock function with given fields: ifName, netnsPath
func (_m *Manager) RemoveIpoibLink(ifName string, netnsPath string) error {
	ret := _m.Called(ifName, netnsPath)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(ifName, netnsPath)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
