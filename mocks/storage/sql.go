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

// InsertTransaction mocks base method.
func (m *MockSQLClient) InsertTransaction(ctx context.Context, entry models.TransactionEntry) (*models.TransactionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertTransaction", ctx, entry)
	ret0, _ := ret[0].(*models.TransactionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertTransaction indicates an expected call of InsertTransaction.
func (mr *MockSQLClientMockRecorder) InsertTransaction(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertTransaction", reflect.TypeOf((*MockSQLClient)(nil).InsertTransaction), ctx, entry)
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

// SelectBalance mocks base method.
func (m *MockSQLClient) SelectBalance(ctx context.Context, userID, currencyID string) (*models.BalanceEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectBalance", ctx, userID, currencyID)
	ret0, _ := ret[0].(*models.BalanceEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectBalance indicates an expected call of SelectBalance.
func (mr *MockSQLClientMockRecorder) SelectBalance(ctx, userID, currencyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectBalance", reflect.TypeOf((*MockSQLClient)(nil).SelectBalance), ctx, userID, currencyID)
}

// SelectBulkBalance mocks base method.
func (m *MockSQLClient) SelectBulkBalance(ctx context.Context, userID string) ([]models.BalanceEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectBulkBalance", ctx, userID)
	ret0, _ := ret[0].([]models.BalanceEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectBulkBalance indicates an expected call of SelectBulkBalance.
func (mr *MockSQLClientMockRecorder) SelectBulkBalance(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectBulkBalance", reflect.TypeOf((*MockSQLClient)(nil).SelectBulkBalance), ctx, userID)
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

// SelectTradeSession mocks base method.
func (m *MockSQLClient) SelectTradeSession(ctx context.Context, userID, sessionID string) (*models.TradeSessionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTradeSession", ctx, userID, sessionID)
	ret0, _ := ret[0].(*models.TradeSessionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTradeSession indicates an expected call of SelectTradeSession.
func (mr *MockSQLClientMockRecorder) SelectTradeSession(ctx, userID, sessionID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTradeSession", reflect.TypeOf((*MockSQLClient)(nil).SelectTradeSession), ctx, userID, sessionID)
}

// SelectTradeSessions mocks base method.
func (m *MockSQLClient) SelectTradeSessions(ctx context.Context, params models.GetTradeSessionsParams) ([]models.TradeSessionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTradeSessions", ctx, params)
	ret0, _ := ret[0].([]models.TradeSessionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTradeSessions indicates an expected call of SelectTradeSessions.
func (mr *MockSQLClientMockRecorder) SelectTradeSessions(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTradeSessions", reflect.TypeOf((*MockSQLClient)(nil).SelectTradeSessions), ctx, params)
}

// SelectTransaction mocks base method.
func (m *MockSQLClient) SelectTransaction(ctx context.Context, id, userID string) (*models.TransactionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTransaction", ctx, id, userID)
	ret0, _ := ret[0].(*models.TransactionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTransaction indicates an expected call of SelectTransaction.
func (mr *MockSQLClientMockRecorder) SelectTransaction(ctx, id, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTransaction", reflect.TypeOf((*MockSQLClient)(nil).SelectTransaction), ctx, id, userID)
}

// SelectTransactions mocks base method.
func (m *MockSQLClient) SelectTransactions(ctx context.Context, params models.GetTransactionsParams) ([]models.TransactionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectTransactions", ctx, params)
	ret0, _ := ret[0].([]models.TransactionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectTransactions indicates an expected call of SelectTransactions.
func (mr *MockSQLClientMockRecorder) SelectTransactions(ctx, params interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectTransactions", reflect.TypeOf((*MockSQLClient)(nil).SelectTransactions), ctx, params)
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

// UpsertBalance mocks base method.
func (m *MockSQLClient) UpsertBalance(ctx context.Context, entry models.BalanceEntry) (*models.BalanceEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertBalance", ctx, entry)
	ret0, _ := ret[0].(*models.BalanceEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertBalance indicates an expected call of UpsertBalance.
func (mr *MockSQLClientMockRecorder) UpsertBalance(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertBalance", reflect.TypeOf((*MockSQLClient)(nil).UpsertBalance), ctx, entry)
}

// UpsertTradeSession mocks base method.
func (m *MockSQLClient) UpsertTradeSession(ctx context.Context, entry models.TradeSessionEntry) (*models.TradeSessionEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertTradeSession", ctx, entry)
	ret0, _ := ret[0].(*models.TradeSessionEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertTradeSession indicates an expected call of UpsertTradeSession.
func (mr *MockSQLClientMockRecorder) UpsertTradeSession(ctx, entry interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertTradeSession", reflect.TypeOf((*MockSQLClient)(nil).UpsertTradeSession), ctx, entry)
}
