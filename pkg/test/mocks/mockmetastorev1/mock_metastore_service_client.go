// Code generated by mockery. DO NOT EDIT.

package mockmetastorev1

import (
	context "context"

	grpc "google.golang.org/grpc"

	metastorev1 "github.com/grafana/pyroscope/api/gen/proto/go/metastore/v1"

	mock "github.com/stretchr/testify/mock"

	typesv1 "github.com/grafana/pyroscope/api/gen/proto/go/types/v1"
)

// MockMetastoreServiceClient is an autogenerated mock type for the MetastoreServiceClient type
type MockMetastoreServiceClient struct {
	mock.Mock
}

type MockMetastoreServiceClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockMetastoreServiceClient) EXPECT() *MockMetastoreServiceClient_Expecter {
	return &MockMetastoreServiceClient_Expecter{mock: &_m.Mock}
}

// AddBlock provides a mock function with given fields: ctx, in, opts
func (_m *MockMetastoreServiceClient) AddBlock(ctx context.Context, in *metastorev1.AddBlockRequest, opts ...grpc.CallOption) (*metastorev1.AddBlockResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for AddBlock")
	}

	var r0 *metastorev1.AddBlockResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.AddBlockRequest, ...grpc.CallOption) (*metastorev1.AddBlockResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.AddBlockRequest, ...grpc.CallOption) *metastorev1.AddBlockResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.AddBlockResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.AddBlockRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoreServiceClient_AddBlock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddBlock'
type MockMetastoreServiceClient_AddBlock_Call struct {
	*mock.Call
}

// AddBlock is a helper method to define mock.On call
//   - ctx context.Context
//   - in *metastorev1.AddBlockRequest
//   - opts ...grpc.CallOption
func (_e *MockMetastoreServiceClient_Expecter) AddBlock(ctx interface{}, in interface{}, opts ...interface{}) *MockMetastoreServiceClient_AddBlock_Call {
	return &MockMetastoreServiceClient_AddBlock_Call{Call: _e.mock.On("AddBlock",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetastoreServiceClient_AddBlock_Call) Run(run func(ctx context.Context, in *metastorev1.AddBlockRequest, opts ...grpc.CallOption)) *MockMetastoreServiceClient_AddBlock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*metastorev1.AddBlockRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetastoreServiceClient_AddBlock_Call) Return(_a0 *metastorev1.AddBlockResponse, _a1 error) *MockMetastoreServiceClient_AddBlock_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoreServiceClient_AddBlock_Call) RunAndReturn(run func(context.Context, *metastorev1.AddBlockRequest, ...grpc.CallOption) (*metastorev1.AddBlockResponse, error)) *MockMetastoreServiceClient_AddBlock_Call {
	_c.Call.Return(run)
	return _c
}

// GetProfileStats provides a mock function with given fields: ctx, in, opts
func (_m *MockMetastoreServiceClient) GetProfileStats(ctx context.Context, in *metastorev1.GetProfileStatsRequest, opts ...grpc.CallOption) (*typesv1.GetProfileStatsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for GetProfileStats")
	}

	var r0 *typesv1.GetProfileStatsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.GetProfileStatsRequest, ...grpc.CallOption) (*typesv1.GetProfileStatsResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.GetProfileStatsRequest, ...grpc.CallOption) *typesv1.GetProfileStatsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*typesv1.GetProfileStatsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.GetProfileStatsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoreServiceClient_GetProfileStats_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProfileStats'
type MockMetastoreServiceClient_GetProfileStats_Call struct {
	*mock.Call
}

// GetProfileStats is a helper method to define mock.On call
//   - ctx context.Context
//   - in *metastorev1.GetProfileStatsRequest
//   - opts ...grpc.CallOption
func (_e *MockMetastoreServiceClient_Expecter) GetProfileStats(ctx interface{}, in interface{}, opts ...interface{}) *MockMetastoreServiceClient_GetProfileStats_Call {
	return &MockMetastoreServiceClient_GetProfileStats_Call{Call: _e.mock.On("GetProfileStats",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetastoreServiceClient_GetProfileStats_Call) Run(run func(ctx context.Context, in *metastorev1.GetProfileStatsRequest, opts ...grpc.CallOption)) *MockMetastoreServiceClient_GetProfileStats_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*metastorev1.GetProfileStatsRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetastoreServiceClient_GetProfileStats_Call) Return(_a0 *typesv1.GetProfileStatsResponse, _a1 error) *MockMetastoreServiceClient_GetProfileStats_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoreServiceClient_GetProfileStats_Call) RunAndReturn(run func(context.Context, *metastorev1.GetProfileStatsRequest, ...grpc.CallOption) (*typesv1.GetProfileStatsResponse, error)) *MockMetastoreServiceClient_GetProfileStats_Call {
	_c.Call.Return(run)
	return _c
}

// QueryMetadata provides a mock function with given fields: ctx, in, opts
func (_m *MockMetastoreServiceClient) QueryMetadata(ctx context.Context, in *metastorev1.QueryMetadataRequest, opts ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for QueryMetadata")
	}

	var r0 *metastorev1.QueryMetadataResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) *metastorev1.QueryMetadataResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.QueryMetadataResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoreServiceClient_QueryMetadata_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'QueryMetadata'
type MockMetastoreServiceClient_QueryMetadata_Call struct {
	*mock.Call
}

// QueryMetadata is a helper method to define mock.On call
//   - ctx context.Context
//   - in *metastorev1.QueryMetadataRequest
//   - opts ...grpc.CallOption
func (_e *MockMetastoreServiceClient_Expecter) QueryMetadata(ctx interface{}, in interface{}, opts ...interface{}) *MockMetastoreServiceClient_QueryMetadata_Call {
	return &MockMetastoreServiceClient_QueryMetadata_Call{Call: _e.mock.On("QueryMetadata",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetastoreServiceClient_QueryMetadata_Call) Run(run func(ctx context.Context, in *metastorev1.QueryMetadataRequest, opts ...grpc.CallOption)) *MockMetastoreServiceClient_QueryMetadata_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*metastorev1.QueryMetadataRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetastoreServiceClient_QueryMetadata_Call) Return(_a0 *metastorev1.QueryMetadataResponse, _a1 error) *MockMetastoreServiceClient_QueryMetadata_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoreServiceClient_QueryMetadata_Call) RunAndReturn(run func(context.Context, *metastorev1.QueryMetadataRequest, ...grpc.CallOption) (*metastorev1.QueryMetadataResponse, error)) *MockMetastoreServiceClient_QueryMetadata_Call {
	_c.Call.Return(run)
	return _c
}

// ReadIndex provides a mock function with given fields: ctx, in, opts
func (_m *MockMetastoreServiceClient) ReadIndex(ctx context.Context, in *metastorev1.ReadIndexRequest, opts ...grpc.CallOption) (*metastorev1.ReadIndexResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReadIndex")
	}

	var r0 *metastorev1.ReadIndexResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.ReadIndexRequest, ...grpc.CallOption) (*metastorev1.ReadIndexResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *metastorev1.ReadIndexRequest, ...grpc.CallOption) *metastorev1.ReadIndexResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*metastorev1.ReadIndexResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *metastorev1.ReadIndexRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockMetastoreServiceClient_ReadIndex_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReadIndex'
type MockMetastoreServiceClient_ReadIndex_Call struct {
	*mock.Call
}

// ReadIndex is a helper method to define mock.On call
//   - ctx context.Context
//   - in *metastorev1.ReadIndexRequest
//   - opts ...grpc.CallOption
func (_e *MockMetastoreServiceClient_Expecter) ReadIndex(ctx interface{}, in interface{}, opts ...interface{}) *MockMetastoreServiceClient_ReadIndex_Call {
	return &MockMetastoreServiceClient_ReadIndex_Call{Call: _e.mock.On("ReadIndex",
		append([]interface{}{ctx, in}, opts...)...)}
}

func (_c *MockMetastoreServiceClient_ReadIndex_Call) Run(run func(ctx context.Context, in *metastorev1.ReadIndexRequest, opts ...grpc.CallOption)) *MockMetastoreServiceClient_ReadIndex_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]grpc.CallOption, len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(grpc.CallOption)
			}
		}
		run(args[0].(context.Context), args[1].(*metastorev1.ReadIndexRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockMetastoreServiceClient_ReadIndex_Call) Return(_a0 *metastorev1.ReadIndexResponse, _a1 error) *MockMetastoreServiceClient_ReadIndex_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockMetastoreServiceClient_ReadIndex_Call) RunAndReturn(run func(context.Context, *metastorev1.ReadIndexRequest, ...grpc.CallOption) (*metastorev1.ReadIndexResponse, error)) *MockMetastoreServiceClient_ReadIndex_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockMetastoreServiceClient creates a new instance of MockMetastoreServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMetastoreServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMetastoreServiceClient {
	mock := &MockMetastoreServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
