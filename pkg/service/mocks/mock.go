package mock_service

import (
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

type MockBalanceInfo struct {
	ctrl     *gomock.Controller
	recorder *MockBalanceInfoMockRecorder
}

type MockBalanceInfoMockRecorder struct {
	mock *MockBalanceInfo
}

func NewMockBalanceInfoUser(ctrl *gomock.Controller) *MockBalanceInfo {
	mock := &MockBalanceInfo{ctrl: ctrl}
	mock.recorder = &MockBalanceInfoMockRecorder{mock}
	return mock
}

func (m *MockBalanceInfo) EXPECT() *MockBalanceInfoMockRecorder {
	return m.recorder
}

func (m *MockBalanceInfo) BalanceInfoUser(ctx context.Context, id int64) (int64, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BalanceInfo", id)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(float64)
	return ret0, ret1, nil
}

func (mr *MockBalanceInfoMockRecorder) BalanceInfoUser(ctx context.Context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceInfo", reflect.TypeOf((*MockBalanceInfo)(nil).BalanceInfoUser), id)
}
