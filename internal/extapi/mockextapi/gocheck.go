// Code generated by MockGen. DO NOT EDIT.
// Source: gocheck.go
//
// Generated by this command:
//
//	mockgen -source=gocheck.go -destination=mockextapi/gocheck.go -package=mockextapi
//

// Package mockextapi is a generated GoMock package.
package mockextapi

import (
	context "context"
	reflect "reflect"

	pbgocheck "github.com/Hidayathamir/protobuf/gocheck"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockIGocheck is a mock of IGocheck interface.
type MockIGocheck struct {
	ctrl     *gomock.Controller
	recorder *MockIGocheckMockRecorder
}

// MockIGocheckMockRecorder is the mock recorder for MockIGocheck.
type MockIGocheckMockRecorder struct {
	mock *MockIGocheck
}

// NewMockIGocheck creates a new mock instance.
func NewMockIGocheck(ctrl *gomock.Controller) *MockIGocheck {
	mock := &MockIGocheck{ctrl: ctrl}
	mock.recorder = &MockIGocheckMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIGocheck) EXPECT() *MockIGocheckMockRecorder {
	return m.recorder
}

// Transfer mocks base method.
func (m *MockIGocheck) Transfer(ctx context.Context, in *pbgocheck.ReqDigitalWalletTransfer, opts ...grpc.CallOption) (*pbgocheck.ResDigitalWalletTransfer, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Transfer", varargs...)
	ret0, _ := ret[0].(*pbgocheck.ResDigitalWalletTransfer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Transfer indicates an expected call of Transfer.
func (mr *MockIGocheckMockRecorder) Transfer(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Transfer", reflect.TypeOf((*MockIGocheck)(nil).Transfer), varargs...)
}
