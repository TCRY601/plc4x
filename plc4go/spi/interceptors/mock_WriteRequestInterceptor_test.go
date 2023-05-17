/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

// Code generated by mockery v2.27.1. DO NOT EDIT.

package interceptors

import (
	context "context"

	model "github.com/apache/plc4x/plc4go/pkg/api/model"
	mock "github.com/stretchr/testify/mock"
)

// MockWriteRequestInterceptor is an autogenerated mock type for the WriteRequestInterceptor type
type MockWriteRequestInterceptor struct {
	mock.Mock
}

type MockWriteRequestInterceptor_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWriteRequestInterceptor) EXPECT() *MockWriteRequestInterceptor_Expecter {
	return &MockWriteRequestInterceptor_Expecter{mock: &_m.Mock}
}

// InterceptWriteRequest provides a mock function with given fields: ctx, writeRequest
func (_m *MockWriteRequestInterceptor) InterceptWriteRequest(ctx context.Context, writeRequest model.PlcWriteRequest) []model.PlcWriteRequest {
	ret := _m.Called(ctx, writeRequest)

	var r0 []model.PlcWriteRequest
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcWriteRequest) []model.PlcWriteRequest); ok {
		r0 = rf(ctx, writeRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.PlcWriteRequest)
		}
	}

	return r0
}

// MockWriteRequestInterceptor_InterceptWriteRequest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'InterceptWriteRequest'
type MockWriteRequestInterceptor_InterceptWriteRequest_Call struct {
	*mock.Call
}

// InterceptWriteRequest is a helper method to define mock.On call
//   - ctx context.Context
//   - writeRequest model.PlcWriteRequest
func (_e *MockWriteRequestInterceptor_Expecter) InterceptWriteRequest(ctx interface{}, writeRequest interface{}) *MockWriteRequestInterceptor_InterceptWriteRequest_Call {
	return &MockWriteRequestInterceptor_InterceptWriteRequest_Call{Call: _e.mock.On("InterceptWriteRequest", ctx, writeRequest)}
}

func (_c *MockWriteRequestInterceptor_InterceptWriteRequest_Call) Run(run func(ctx context.Context, writeRequest model.PlcWriteRequest)) *MockWriteRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcWriteRequest))
	})
	return _c
}

func (_c *MockWriteRequestInterceptor_InterceptWriteRequest_Call) Return(_a0 []model.PlcWriteRequest) *MockWriteRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWriteRequestInterceptor_InterceptWriteRequest_Call) RunAndReturn(run func(context.Context, model.PlcWriteRequest) []model.PlcWriteRequest) *MockWriteRequestInterceptor_InterceptWriteRequest_Call {
	_c.Call.Return(run)
	return _c
}

// ProcessWriteResponses provides a mock function with given fields: ctx, writeRequest, writeResults
func (_m *MockWriteRequestInterceptor) ProcessWriteResponses(ctx context.Context, writeRequest model.PlcWriteRequest, writeResults []model.PlcWriteRequestResult) model.PlcWriteRequestResult {
	ret := _m.Called(ctx, writeRequest, writeResults)

	var r0 model.PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func(context.Context, model.PlcWriteRequest, []model.PlcWriteRequestResult) model.PlcWriteRequestResult); ok {
		r0 = rf(ctx, writeRequest, writeResults)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcWriteRequestResult)
		}
	}

	return r0
}

// MockWriteRequestInterceptor_ProcessWriteResponses_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ProcessWriteResponses'
type MockWriteRequestInterceptor_ProcessWriteResponses_Call struct {
	*mock.Call
}

// ProcessWriteResponses is a helper method to define mock.On call
//   - ctx context.Context
//   - writeRequest model.PlcWriteRequest
//   - writeResults []model.PlcWriteRequestResult
func (_e *MockWriteRequestInterceptor_Expecter) ProcessWriteResponses(ctx interface{}, writeRequest interface{}, writeResults interface{}) *MockWriteRequestInterceptor_ProcessWriteResponses_Call {
	return &MockWriteRequestInterceptor_ProcessWriteResponses_Call{Call: _e.mock.On("ProcessWriteResponses", ctx, writeRequest, writeResults)}
}

func (_c *MockWriteRequestInterceptor_ProcessWriteResponses_Call) Run(run func(ctx context.Context, writeRequest model.PlcWriteRequest, writeResults []model.PlcWriteRequestResult)) *MockWriteRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PlcWriteRequest), args[2].([]model.PlcWriteRequestResult))
	})
	return _c
}

func (_c *MockWriteRequestInterceptor_ProcessWriteResponses_Call) Return(_a0 model.PlcWriteRequestResult) *MockWriteRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockWriteRequestInterceptor_ProcessWriteResponses_Call) RunAndReturn(run func(context.Context, model.PlcWriteRequest, []model.PlcWriteRequestResult) model.PlcWriteRequestResult) *MockWriteRequestInterceptor_ProcessWriteResponses_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockWriteRequestInterceptor interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockWriteRequestInterceptor creates a new instance of MockWriteRequestInterceptor. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockWriteRequestInterceptor(t mockConstructorTestingTNewMockWriteRequestInterceptor) *MockWriteRequestInterceptor {
	mock := &MockWriteRequestInterceptor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}