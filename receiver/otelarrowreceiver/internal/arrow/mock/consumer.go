// Code generated by MockGen. DO NOT EDIT.
// Source: go.opentelemetry.io/collector/consumer (interfaces: Traces,Metrics,Logs)
//
// Generated by this command:
//
//	mockgen -package mock go.opentelemetry.io/collector/consumer Traces,Metrics,Logs
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	consumer "go.opentelemetry.io/collector/consumer"
	plog "go.opentelemetry.io/collector/pdata/plog"
	pmetric "go.opentelemetry.io/collector/pdata/pmetric"
	ptrace "go.opentelemetry.io/collector/pdata/ptrace"
	gomock "go.uber.org/mock/gomock"
)

// MockTraces is a mock of Traces interface.
type MockTraces struct {
	ctrl     *gomock.Controller
	recorder *MockTracesMockRecorder
}

// MockTracesMockRecorder is the mock recorder for MockTraces.
type MockTracesMockRecorder struct {
	mock *MockTraces
}

// NewMockTraces creates a new mock instance.
func NewMockTraces(ctrl *gomock.Controller) *MockTraces {
	mock := &MockTraces{ctrl: ctrl}
	mock.recorder = &MockTracesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTraces) EXPECT() *MockTracesMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockTraces) Capabilities() consumer.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(consumer.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockTracesMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockTraces)(nil).Capabilities))
}

// ConsumeTraces mocks base method.
func (m *MockTraces) ConsumeTraces(arg0 context.Context, arg1 ptrace.Traces) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeTraces", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeTraces indicates an expected call of ConsumeTraces.
func (mr *MockTracesMockRecorder) ConsumeTraces(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeTraces", reflect.TypeOf((*MockTraces)(nil).ConsumeTraces), arg0, arg1)
}

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockMetrics) Capabilities() consumer.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(consumer.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockMetricsMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockMetrics)(nil).Capabilities))
}

// ConsumeMetrics mocks base method.
func (m *MockMetrics) ConsumeMetrics(arg0 context.Context, arg1 pmetric.Metrics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeMetrics", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeMetrics indicates an expected call of ConsumeMetrics.
func (mr *MockMetricsMockRecorder) ConsumeMetrics(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeMetrics", reflect.TypeOf((*MockMetrics)(nil).ConsumeMetrics), arg0, arg1)
}

// MockLogs is a mock of Logs interface.
type MockLogs struct {
	ctrl     *gomock.Controller
	recorder *MockLogsMockRecorder
}

// MockLogsMockRecorder is the mock recorder for MockLogs.
type MockLogsMockRecorder struct {
	mock *MockLogs
}

// NewMockLogs creates a new mock instance.
func NewMockLogs(ctrl *gomock.Controller) *MockLogs {
	mock := &MockLogs{ctrl: ctrl}
	mock.recorder = &MockLogsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLogs) EXPECT() *MockLogsMockRecorder {
	return m.recorder
}

// Capabilities mocks base method.
func (m *MockLogs) Capabilities() consumer.Capabilities {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].(consumer.Capabilities)
	return ret0
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockLogsMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockLogs)(nil).Capabilities))
}

// ConsumeLogs mocks base method.
func (m *MockLogs) ConsumeLogs(arg0 context.Context, arg1 plog.Logs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsumeLogs", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ConsumeLogs indicates an expected call of ConsumeLogs.
func (mr *MockLogsMockRecorder) ConsumeLogs(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsumeLogs", reflect.TypeOf((*MockLogs)(nil).ConsumeLogs), arg0, arg1)
}
