// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quasarlabs/quasarnode/x/qoracle/types (interfaces: ICS4Wrapper)

// Package mock is a generated GoMock package.
package mock

import (
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/capability/types"
	exported "github.com/cosmos/ibc-go/v4/modules/core/exported"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockICS4Wrapper is a mock of ICS4Wrapper interface
type MockICS4Wrapper struct {
	ctrl     *gomock.Controller
	recorder *MockICS4WrapperMockRecorder
}

// MockICS4WrapperMockRecorder is the mock recorder for MockICS4Wrapper
type MockICS4WrapperMockRecorder struct {
	mock *MockICS4Wrapper
}

// NewMockICS4Wrapper creates a new mock instance
func NewMockICS4Wrapper(ctrl *gomock.Controller) *MockICS4Wrapper {
	mock := &MockICS4Wrapper{ctrl: ctrl}
	mock.recorder = &MockICS4WrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockICS4Wrapper) EXPECT() *MockICS4WrapperMockRecorder {
	return m.recorder
}

// SendPacket mocks base method
func (m *MockICS4Wrapper) SendPacket(arg0 types.Context, arg1 *types0.Capability, arg2 exported.PacketI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPacket", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPacket indicates an expected call of SendPacket
func (mr *MockICS4WrapperMockRecorder) SendPacket(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPacket", reflect.TypeOf((*MockICS4Wrapper)(nil).SendPacket), arg0, arg1, arg2)
}
