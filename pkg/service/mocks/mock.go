package mock_service

import (
	"context"
	"reflect"
	"tech_task"

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
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockBalanceInfoMockRecorder) BalanceInfoUser(ctx context.Context, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BalanceInfo", reflect.TypeOf((*MockBalanceInfo)(nil).BalanceInfoUser), id)
}

type MockUpBalance struct {
	ctrl     *gomock.Controller
	recorder *MockUpBalanceMockRecorder
}

type MockUpBalanceMockRecorder struct {
	mock *MockUpBalance
}

func NewMockUpBalanceUser(ctrl *gomock.Controller) *MockUpBalance {
	mock := &MockUpBalance{ctrl: ctrl}
	mock.recorder = &MockUpBalanceMockRecorder{mock}
	return mock
}

func (m *MockUpBalance) EXPECT() *MockUpBalanceMockRecorder {
	return m.recorder
}

func (m *MockUpBalance) UpBalanceUser(ctx context.Context, id int64, amount float64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpBalance", id, amount)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockUpBalanceMockRecorder) UpBalanceUser(ctx context.Context, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpBalance", reflect.TypeOf((*MockUpBalance)(nil).UpBalanceUser), id, amount)
}

type MockWritingOff struct {
	ctrl     *gomock.Controller
	recorder *MockWritingOffMockRecorder
}

type MockWritingOffMockRecorder struct {
	mock *MockWritingOff
}

func NewMockWritingOffUser(ctrl *gomock.Controller) *MockWritingOff {
	mock := &MockWritingOff{ctrl: ctrl}
	mock.recorder = &MockWritingOffMockRecorder{mock}
	return mock
}

func (m *MockWritingOff) EXPECT() *MockWritingOffMockRecorder {
	return m.recorder
}

func (m *MockWritingOff) WritingOffUser(ctx context.Context, id int64, amount float64) (int64, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WritingOff", id, amount)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockWritingOffMockRecorder) WritingOffUser(ctx context.Context, id, amount interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WritingOff", reflect.TypeOf((*MockWritingOff)(nil).WritingOffUser), id, amount)
}

type MockAddDescription struct {
	ctrl     *gomock.Controller
	recorder *MockAddDescriptionMockRecorder
}

type MockAddDescriptionMockRecorder struct {
	mock *MockAddDescription
}

func NewMockAddDescriptionUser(ctrl *gomock.Controller) *MockAddDescription {
	mock := &MockAddDescription{ctrl: ctrl}
	mock.recorder = &MockAddDescriptionMockRecorder{mock}
	return mock
}

func (m *MockAddDescription) EXPECT() *MockAddDescriptionMockRecorder {
	return m.recorder
}

func (m *MockAddDescription) AddDescriptionUser(ctx context.Context, id int64, balanceAtMoment, corectAmount float64, refill, description, senderReceiver string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDescription", id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockAddDescriptionMockRecorder) AddDescriptionUser(ctx context.Context, id, balanceAtMoment, corectAmount, refill, description, senderReceiver interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDescription", reflect.TypeOf((*MockAddDescription)(nil).AddDescriptionUser), id, balanceAtMoment, corectAmount, refill, description, senderReceiver)
}

type MockGetDescriptions struct {
	ctrl     *gomock.Controller
	recorder *MockGetDescriptionsMockRecorder
}

type MockGetDescriptionsMockRecorder struct {
	mock *MockGetDescriptions
}

func NewMocketGetDescriptions(ctrl *gomock.Controller) *MockGetDescriptions {
	mock := &MockGetDescriptions{ctrl: ctrl}
	mock.recorder = &MockGetDescriptionsMockRecorder{mock}
	return mock
}

func (m *MockGetDescriptions) EXPECT() *MockGetDescriptionsMockRecorder {
	return m.recorder
}

func (m *MockGetDescriptions) GetDescriptionsUsers(ctx context.Context, id int64, sortBy, orderBy, sqlOrderBy string) ([]tech_task.Description, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescriptions", id, sortBy, orderBy, sqlOrderBy)
	ret0, _ := ret[0].([]tech_task.Description)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockGetDescriptionsMockRecorder) GetDescriptionsUsers(ctx context.Context, id, sortBy, orderBy, sqlOrderBy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescriptions", reflect.TypeOf((*MockGetDescriptions)(nil).GetDescriptionsUsers), id, sortBy, orderBy, sqlOrderBy)
}
