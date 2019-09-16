// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import "context"
import "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/core"
import flyteidlcore "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
import "github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"
import "github.com/stretchr/testify/mock"

// CatalogClient is an autogenerated mock type for the CatalogClient type
type CatalogClient struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, tr, input
func (_m *CatalogClient) Get(ctx context.Context, tr core.TaskReader, input io.InputReader) (*flyteidlcore.LiteralMap, error) {
	ret := _m.Called(ctx, tr, input)

	var r0 *flyteidlcore.LiteralMap
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskReader, io.InputReader) *flyteidlcore.LiteralMap); ok {
		r0 = rf(ctx, tr, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*flyteidlcore.LiteralMap)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, core.TaskReader, io.InputReader) error); ok {
		r1 = rf(ctx, tr, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: ctx, tCtx, op
func (_m *CatalogClient) Put(ctx context.Context, tCtx core.TaskExecutionContext, op io.OutputReader) error {
	ret := _m.Called(ctx, tCtx, op)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, core.TaskExecutionContext, io.OutputReader) error); ok {
		r0 = rf(ctx, tCtx, op)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}