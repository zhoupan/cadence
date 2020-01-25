// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	archiver "github.com/uber/cadence/common/archiver"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Client is an autogenerated mock type for the Client type
type Client struct {
	mock.Mock
}

// Exist provides a mock function with given fields: ctx, URI, fileName
func (_m *Client) Exist(ctx context.Context, URI archiver.URI, fileName string) (bool, error) {
	ret := _m.Called(ctx, URI, fileName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, archiver.URI, string) bool); ok {
		r0 = rf(ctx, URI, fileName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, archiver.URI, string) error); ok {
		r1 = rf(ctx, URI, fileName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, URI, file
func (_m *Client) Get(ctx context.Context, URI archiver.URI, file string) ([]byte, error) {
	ret := _m.Called(ctx, URI, file)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, archiver.URI, string) []byte); ok {
		r0 = rf(ctx, URI, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, archiver.URI, string) error); ok {
		r1 = rf(ctx, URI, file)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Query provides a mock function with given fields: ctx, URI, fileNamePrefix
func (_m *Client) Query(ctx context.Context, URI archiver.URI, fileNamePrefix string) ([]string, error) {
	ret := _m.Called(ctx, URI, fileNamePrefix)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, archiver.URI, string) []string); ok {
		r0 = rf(ctx, URI, fileNamePrefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, archiver.URI, string) error); ok {
		r1 = rf(ctx, URI, fileNamePrefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Upload provides a mock function with given fields: ctx, URI, fileName, file
func (_m *Client) Upload(ctx context.Context, URI archiver.URI, fileName string, file []byte) error {
	ret := _m.Called(ctx, URI, fileName, file)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, archiver.URI, string, []byte) error); ok {
		r0 = rf(ctx, URI, fileName, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
