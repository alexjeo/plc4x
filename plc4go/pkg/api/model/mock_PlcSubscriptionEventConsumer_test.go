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

// Code generated by mockery v2.26.1. DO NOT EDIT.

package model

import mock "github.com/stretchr/testify/mock"

// MockPlcSubscriptionEventConsumer is an autogenerated mock type for the PlcSubscriptionEventConsumer type
type MockPlcSubscriptionEventConsumer struct {
	mock.Mock
}

type MockPlcSubscriptionEventConsumer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPlcSubscriptionEventConsumer) EXPECT() *MockPlcSubscriptionEventConsumer_Expecter {
	return &MockPlcSubscriptionEventConsumer_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: event
func (_m *MockPlcSubscriptionEventConsumer) Execute(event PlcSubscriptionEvent) {
	_m.Called(event)
}

// MockPlcSubscriptionEventConsumer_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPlcSubscriptionEventConsumer_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - event PlcSubscriptionEvent
func (_e *MockPlcSubscriptionEventConsumer_Expecter) Execute(event interface{}) *MockPlcSubscriptionEventConsumer_Execute_Call {
	return &MockPlcSubscriptionEventConsumer_Execute_Call{Call: _e.mock.On("Execute", event)}
}

func (_c *MockPlcSubscriptionEventConsumer_Execute_Call) Run(run func(event PlcSubscriptionEvent)) *MockPlcSubscriptionEventConsumer_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(PlcSubscriptionEvent))
	})
	return _c
}

func (_c *MockPlcSubscriptionEventConsumer_Execute_Call) Return() *MockPlcSubscriptionEventConsumer_Execute_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPlcSubscriptionEventConsumer_Execute_Call) RunAndReturn(run func(PlcSubscriptionEvent)) *MockPlcSubscriptionEventConsumer_Execute_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewMockPlcSubscriptionEventConsumer interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockPlcSubscriptionEventConsumer creates a new instance of MockPlcSubscriptionEventConsumer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockPlcSubscriptionEventConsumer(t mockConstructorTestingTNewMockPlcSubscriptionEventConsumer) *MockPlcSubscriptionEventConsumer {
	mock := &MockPlcSubscriptionEventConsumer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}