// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import client "github.com/lyft/flyteplugins/go/tasks/v1/qubole_collection/client"
import context "context"
import mock "github.com/stretchr/testify/mock"

// QuboleClient is an autogenerated mock type for the QuboleClient type
type QuboleClient struct {
	mock.Mock
}

// ExecuteHiveCommand provides a mock function with given fields: ctx, commandStr, timeoutVal, clusterLabel, accountKey, tags
func (_m *QuboleClient) ExecuteHiveCommand(ctx context.Context, commandStr string, timeoutVal uint32, clusterLabel string, accountKey string, tags []string) (*client.QuboleCommandDetails, error) {
	ret := _m.Called(ctx, commandStr, timeoutVal, clusterLabel, accountKey, tags)

	var r0 *client.QuboleCommandDetails
	if rf, ok := ret.Get(0).(func(context.Context, string, uint32, string, string, []string) *client.QuboleCommandDetails); ok {
		r0 = rf(ctx, commandStr, timeoutVal, clusterLabel, accountKey, tags)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*client.QuboleCommandDetails)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint32, string, string, []string) error); ok {
		r1 = rf(ctx, commandStr, timeoutVal, clusterLabel, accountKey, tags)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCommandStatus provides a mock function with given fields: ctx, commandID, accountKey
func (_m *QuboleClient) GetCommandStatus(ctx context.Context, commandID string, accountKey string) (client.QuboleStatus, error) {
	ret := _m.Called(ctx, commandID, accountKey)

	var r0 client.QuboleStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string) client.QuboleStatus); ok {
		r0 = rf(ctx, commandID, accountKey)
	} else {
		r0 = ret.Get(0).(client.QuboleStatus)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, commandID, accountKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// KillCommand provides a mock function with given fields: ctx, commandID, accountKey
func (_m *QuboleClient) KillCommand(ctx context.Context, commandID string, accountKey string) error {
	ret := _m.Called(ctx, commandID, accountKey)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, commandID, accountKey)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
