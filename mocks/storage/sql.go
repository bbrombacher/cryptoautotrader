// Code generated by MockGen. DO NOT EDIT.
// Source: sql.go

// Package mocksql is a generated GoMock package.
package mocksql

import (
	models "bbrombacher/cryptoautotrader/storage/models"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSQLClient is a mock of SQLClient interface.
type MockSQLClient struct {
	ctrl     *gomock.Controller
	recorder *MockSQLClientMockRecorder
}

// MockSQLClientMockRecorder is the mock recorder for MockSQLClient.
type MockSQLClientMockRecorder struct {
	mock *MockSQLClient
}

// NewMockSQLClient creates a new mock instance.
func NewMockSQLClient(ctrl *gomock.Controller) *MockSQLClient {
	mock := &MockSQLClient{ctrl: ctrl}
	mock.recorder = &MockSQLClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLClient) EXPECT() *MockSQLClientMockRecorder {
	return m.recorder
}

// DeleteCurrency mocks base method.
func (m *MockSQLClient) DeleteCurrency(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCurrency", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCurrency indicates an expected call of DeleteCurrency.
func (mr *MockSQLClientMockRecorder) DeleteCurrency(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCurrency", reflect.TypeOf((*MockSQLClient)(nil).DeleteCurrency), ctx, id)
}

// DeleteUser mocks base method.
func (m *MockSQLClient) DeleteUser(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser.
func (mr *MockSQLClientMockRecorder) DeleteUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockSQLClient)(nil).DeleteUser), ctx, id)
}

// InsertCurrency mocks base method.
func (m *MockSQLClient) InsertCurrency(ctx context.Context, entry models.CurrencyEntry) (*models.CurrencyEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertCurrency", ctx, entry)
	ret0, _ := ret[0].(*models.CurrencyEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertCurrency indicates an expected call of InsertCurrency.
func (mr *MockSQLClientMockRecorder) InsertCurrency(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertCurrency", reflect.TypeOf((*MockSQLClient)(nil).InsertCurrency), ctx, entry)
}

// InsertUser mocks base method.
func (m *MockSQLClient) InsertUser(ctx context.Context, entry models.UserEntry) (*models.UserEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx, entry)
	ret0, _ := ret[0].(*models.UserEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertUser indicates an expected call of InsertUser.
func (mr *MockSQLClientMockRecorder) InsertUser(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockSQLClient)(nil).InsertUser), ctx, entry)
}

// SelectCurrencies mocks base method.
func (m *MockSQLClient) SelectCurrencies(ctx context.Context, params models.GetCurrenciesParams) ([]models.CurrencyEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectCurrencies", ctx, params)
	ret0, _ := ret[0].([]models.CurrencyEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectCurrencies indicates an expected call of SelectCurrencies.
func (mr *MockSQLClientMockRecorder) SelectCurrencies(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectCurrencies", reflect.TypeOf((*MockSQLClient)(nil).SelectCurrencies), ctx, params)
}

// SelectCurrency mocks base method.
func (m *MockSQLClient) SelectCurrency(ctx context.Context, id string) (*models.CurrencyEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectCurrency", ctx, id)
	ret0, _ := ret[0].(*models.CurrencyEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectCurrency indicates an expected call of SelectCurrency.
func (mr *MockSQLClientMockRecorder) SelectCurrency(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectCurrency", reflect.TypeOf((*MockSQLClient)(nil).SelectCurrency), ctx, id)
}

// SelectUser mocks base method.
func (m *MockSQLClient) SelectUser(ctx context.Context, id string) (*models.UserEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectUser", ctx, id)
	ret0, _ := ret[0].(*models.UserEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectUser indicates an expected call of SelectUser.
func (mr *MockSQLClientMockRecorder) SelectUser(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectUser", reflect.TypeOf((*MockSQLClient)(nil).SelectUser), ctx, id)
}

// UpdateCurrency mocks base method.
func (m *MockSQLClient) UpdateCurrency(ctx context.Context, entry models.CurrencyEntry, updateColumns []string) (*models.CurrencyEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCurrency", ctx, entry, updateColumns)
	ret0, _ := ret[0].(*models.CurrencyEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateCurrency indicates an expected call of UpdateCurrency.
func (mr *MockSQLClientMockRecorder) UpdateCurrency(ctx, entry, updateColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCurrency", reflect.TypeOf((*MockSQLClient)(nil).UpdateCurrency), ctx, entry, updateColumns)
}

// UpdateUser mocks base method.
func (m *MockSQLClient) UpdateUser(ctx context.Context, entry models.UserEntry, updateColumns []string) (*models.UserEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx, entry, updateColumns)
	ret0, _ := ret[0].(*models.UserEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockSQLClientMockRecorder) UpdateUser(ctx, entry, updateColumns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockSQLClient)(nil).UpdateUser), ctx, entry, updateColumns)
}
