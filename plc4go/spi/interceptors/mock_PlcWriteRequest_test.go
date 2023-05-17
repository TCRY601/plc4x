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

	spi "github.com/apache/plc4x/plc4go/spi"

	values "github.com/apache/plc4x/plc4go/pkg/api/values"
)

// MockPlcWriteRequest is an autogenerated mock type for the PlcWriteRequest type
type MockPlcWriteRequest struct {
	mock.Mock
}

type MockPlcWriteRequest_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcWriteRequest) EXPECT() *MockPlcWriteRequest_Expecter {
	return &MockPlcWriteRequest_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields:
func (_m *MockPlcWriteRequest) Execute() <-chan model.PlcWriteRequestResult {
	ret := _m.Called()

	var r0 <-chan model.PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func() <-chan model.PlcWriteRequestResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcWriteRequestResult)
		}
	}

	return r0
}

// MockPlcWriteRequest_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPlcWriteRequest_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) Execute() *MockPlcWriteRequest_Execute_Call {
	return &MockPlcWriteRequest_Execute_Call{Call: _e.mock.On("Execute")}
}

func (_c *MockPlcWriteRequest_Execute_Call) Run(run func()) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_Execute_Call) Return(_a0 <-chan model.PlcWriteRequestResult) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_Execute_Call) RunAndReturn(run func() <-chan model.PlcWriteRequestResult) *MockPlcWriteRequest_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteWithContext provides a mock function with given fields: ctx
func (_m *MockPlcWriteRequest) ExecuteWithContext(ctx context.Context) <-chan model.PlcWriteRequestResult {
	ret := _m.Called(ctx)

	var r0 <-chan model.PlcWriteRequestResult
	if rf, ok := ret.Get(0).(func(context.Context) <-chan model.PlcWriteRequestResult); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan model.PlcWriteRequestResult)
		}
	}

	return r0
}

// MockPlcWriteRequest_ExecuteWithContext_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteWithContext'
type MockPlcWriteRequest_ExecuteWithContext_Call struct {
	*mock.Call
}

// ExecuteWithContext is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockPlcWriteRequest_Expecter) ExecuteWithContext(ctx interface{}) *MockPlcWriteRequest_ExecuteWithContext_Call {
	return &MockPlcWriteRequest_ExecuteWithContext_Call{Call: _e.mock.On("ExecuteWithContext", ctx)}
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) Run(run func(ctx context.Context)) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) Return(_a0 <-chan model.PlcWriteRequestResult) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_ExecuteWithContext_Call) RunAndReturn(run func(context.Context) <-chan model.PlcWriteRequestResult) *MockPlcWriteRequest_ExecuteWithContext_Call {
	_c.Call.Return(run)
	return _c
}

// GetTag provides a mock function with given fields: tagName
func (_m *MockPlcWriteRequest) GetTag(tagName string) model.PlcTag {
	ret := _m.Called(tagName)

	var r0 model.PlcTag
	if rf, ok := ret.Get(0).(func(string) model.PlcTag); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(model.PlcTag)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetTag_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTag'
type MockPlcWriteRequest_GetTag_Call struct {
	*mock.Call
}

// GetTag is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcWriteRequest_Expecter) GetTag(tagName interface{}) *MockPlcWriteRequest_GetTag_Call {
	return &MockPlcWriteRequest_GetTag_Call{Call: _e.mock.On("GetTag", tagName)}
}

func (_c *MockPlcWriteRequest_GetTag_Call) Run(run func(tagName string)) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetTag_Call) Return(_a0 model.PlcTag) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetTag_Call) RunAndReturn(run func(string) model.PlcTag) *MockPlcWriteRequest_GetTag_Call {
	_c.Call.Return(run)
	return _c
}

// GetTagNames provides a mock function with given fields:
func (_m *MockPlcWriteRequest) GetTagNames() []string {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetTagNames_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTagNames'
type MockPlcWriteRequest_GetTagNames_Call struct {
	*mock.Call
}

// GetTagNames is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) GetTagNames() *MockPlcWriteRequest_GetTagNames_Call {
	return &MockPlcWriteRequest_GetTagNames_Call{Call: _e.mock.On("GetTagNames")}
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) Run(run func()) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) Return(_a0 []string) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetTagNames_Call) RunAndReturn(run func() []string) *MockPlcWriteRequest_GetTagNames_Call {
	_c.Call.Return(run)
	return _c
}

// GetValue provides a mock function with given fields: tagName
func (_m *MockPlcWriteRequest) GetValue(tagName string) values.PlcValue {
	ret := _m.Called(tagName)

	var r0 values.PlcValue
	if rf, ok := ret.Get(0).(func(string) values.PlcValue); ok {
		r0 = rf(tagName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(values.PlcValue)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type MockPlcWriteRequest_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - tagName string
func (_e *MockPlcWriteRequest_Expecter) GetValue(tagName interface{}) *MockPlcWriteRequest_GetValue_Call {
	return &MockPlcWriteRequest_GetValue_Call{Call: _e.mock.On("GetValue", tagName)}
}

func (_c *MockPlcWriteRequest_GetValue_Call) Run(run func(tagName string)) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetValue_Call) Return(_a0 values.PlcValue) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetValue_Call) RunAndReturn(run func(string) values.PlcValue) *MockPlcWriteRequest_GetValue_Call {
	_c.Call.Return(run)
	return _c
}

// GetWriteRequestInterceptor provides a mock function with given fields:
func (_m *MockPlcWriteRequest) GetWriteRequestInterceptor() WriteRequestInterceptor {
	ret := _m.Called()

	var r0 WriteRequestInterceptor
	if rf, ok := ret.Get(0).(func() WriteRequestInterceptor); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(WriteRequestInterceptor)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetWriteRequestInterceptor_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWriteRequestInterceptor'
type MockPlcWriteRequest_GetWriteRequestInterceptor_Call struct {
	*mock.Call
}

// GetWriteRequestInterceptor is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) GetWriteRequestInterceptor() *MockPlcWriteRequest_GetWriteRequestInterceptor_Call {
	return &MockPlcWriteRequest_GetWriteRequestInterceptor_Call{Call: _e.mock.On("GetWriteRequestInterceptor")}
}

func (_c *MockPlcWriteRequest_GetWriteRequestInterceptor_Call) Run(run func()) *MockPlcWriteRequest_GetWriteRequestInterceptor_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetWriteRequestInterceptor_Call) Return(_a0 WriteRequestInterceptor) *MockPlcWriteRequest_GetWriteRequestInterceptor_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetWriteRequestInterceptor_Call) RunAndReturn(run func() WriteRequestInterceptor) *MockPlcWriteRequest_GetWriteRequestInterceptor_Call {
	_c.Call.Return(run)
	return _c
}

// GetWriter provides a mock function with given fields:
func (_m *MockPlcWriteRequest) GetWriter() spi.PlcWriter {
	ret := _m.Called()

	var r0 spi.PlcWriter
	if rf, ok := ret.Get(0).(func() spi.PlcWriter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(spi.PlcWriter)
		}
	}

	return r0
}

// MockPlcWriteRequest_GetWriter_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetWriter'
type MockPlcWriteRequest_GetWriter_Call struct {
	*mock.Call
}

// GetWriter is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) GetWriter() *MockPlcWriteRequest_GetWriter_Call {
	return &MockPlcWriteRequest_GetWriter_Call{Call: _e.mock.On("GetWriter")}
}

func (_c *MockPlcWriteRequest_GetWriter_Call) Run(run func()) *MockPlcWriteRequest_GetWriter_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_GetWriter_Call) Return(_a0 spi.PlcWriter) *MockPlcWriteRequest_GetWriter_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_GetWriter_Call) RunAndReturn(run func() spi.PlcWriter) *MockPlcWriteRequest_GetWriter_Call {
	_c.Call.Return(run)
	return _c
}

// IsAPlcMessage provides a mock function with given fields:
func (_m *MockPlcWriteRequest) IsAPlcMessage() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockPlcWriteRequest_IsAPlcMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IsAPlcMessage'
type MockPlcWriteRequest_IsAPlcMessage_Call struct {
	*mock.Call
}

// IsAPlcMessage is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) IsAPlcMessage() *MockPlcWriteRequest_IsAPlcMessage_Call {
	return &MockPlcWriteRequest_IsAPlcMessage_Call{Call: _e.mock.On("IsAPlcMessage")}
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) Run(run func()) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) Return(_a0 bool) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_IsAPlcMessage_Call) RunAndReturn(run func() bool) *MockPlcWriteRequest_IsAPlcMessage_Call {
	_c.Call.Return(run)
	return _c
}

// String provides a mock function with given fields:
func (_m *MockPlcWriteRequest) String() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// MockPlcWriteRequest_String_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'String'
type MockPlcWriteRequest_String_Call struct {
	*mock.Call
}

// String is a helper method to define mock.On call
func (_e *MockPlcWriteRequest_Expecter) String() *MockPlcWriteRequest_String_Call {
	return &MockPlcWriteRequest_String_Call{Call: _e.mock.On("String")}
}

func (_c *MockPlcWriteRequest_String_Call) Run(run func()) *MockPlcWriteRequest_String_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPlcWriteRequest_String_Call) Return(_a0 string) *MockPlcWriteRequest_String_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPlcWriteRequest_String_Call) RunAndReturn(run func() string) *MockPlcWriteRequest_String_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcWriteRequest interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcWriteRequest creates a new instance of MockPlcWriteRequest. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcWriteRequest(t mockConstructorTestingTNewMockPlcWriteRequest) *MockPlcWriteRequest {
	mock := &MockPlcWriteRequest{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}