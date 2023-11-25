// Code generated by mockery v2.38.0. DO NOT EDIT.

package utils

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// MockFileService is an autogenerated mock type for the FileService type
type MockFileService struct {
	mock.Mock
}

// CreateFile provides a mock function with given fields: dirPath
func (_m *MockFileService) CreateFile(dirPath string) *os.File {
	ret := _m.Called(dirPath)

	if len(ret) == 0 {
		panic("no return value specified for CreateFile")
	}

	var r0 *os.File
	if rf, ok := ret.Get(0).(func(string) *os.File); ok {
		r0 = rf(dirPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*os.File)
		}
	}

	return r0
}

// FileExists provides a mock function with given fields: dirPath
func (_m *MockFileService) FileExists(dirPath string) bool {
	ret := _m.Called(dirPath)

	if len(ret) == 0 {
		panic("no return value specified for FileExists")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(dirPath)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ReadFileToByteArray provides a mock function with given fields: filePath
func (_m *MockFileService) ReadFileToByteArray(filePath string) []byte {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for ReadFileToByteArray")
	}

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	return r0
}

// ReadLine provides a mock function with given fields: filePath
func (_m *MockFileService) ReadLine(filePath string) string {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for ReadLine")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// RemoveFile provides a mock function with given fields: dirPath
func (_m *MockFileService) RemoveFile(dirPath string) {
	_m.Called(dirPath)
}

// WriteByteArrayToFile provides a mock function with given fields: b, filePath
func (_m *MockFileService) WriteByteArrayToFile(b []byte, filePath string) {
	_m.Called(b, filePath)
}

// WriteToFile provides a mock function with given fields: filePath, content
func (_m *MockFileService) WriteToFile(filePath string, content string) {
	_m.Called(filePath, content)
}

// NewMockFileService creates a new instance of MockFileService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockFileService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockFileService {
	mock := &MockFileService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
