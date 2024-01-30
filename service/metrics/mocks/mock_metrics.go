// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	metrics "github.com/skip-mev/slinky/service/metrics"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Metrics is an autogenerated mock type for the Metrics type
type Metrics struct {
	mock.Mock
}

// AddABCIRequest provides a mock function with given fields: method, status
func (_m *Metrics) AddABCIRequest(method metrics.ABCIMethod, status metrics.Labeller) {
	_m.Called(method, status)
}

// AddOracleResponse provides a mock function with given fields: status
func (_m *Metrics) AddOracleResponse(status metrics.Labeller) {
	_m.Called(status)
}

// AddTickerInclusionStatus provides a mock function with given fields: ticker, included
func (_m *Metrics) AddTickerInclusionStatus(ticker string, included bool) {
	_m.Called(ticker, included)
}

// AddVoteIncludedInLastCommit provides a mock function with given fields: included
func (_m *Metrics) AddVoteIncludedInLastCommit(included bool) {
	_m.Called(included)
}

// ObserveABCIMethodLatency provides a mock function with given fields: method, duration
func (_m *Metrics) ObserveABCIMethodLatency(method metrics.ABCIMethod, duration time.Duration) {
	_m.Called(method, duration)
}

// ObserveOracleResponseLatency provides a mock function with given fields: duration
func (_m *Metrics) ObserveOracleResponseLatency(duration time.Duration) {
	_m.Called(duration)
}

// NewMetrics creates a new instance of Metrics. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetrics(t interface {
	mock.TestingT
	Cleanup(func())
}) *Metrics {
	mock := &Metrics{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
