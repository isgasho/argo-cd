// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	repo "github.com/argoproj/argo-cd/util/repo"
)

// Repo is an autogenerated mock type for the Repo type
type Repo struct {
	mock.Mock
}

// GetApp provides a mock function with given fields: app, resolvedRevision
func (_m *Repo) GetApp(app string, resolvedRevision string) (string, error) {
	ret := _m.Called(app, resolvedRevision)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(app, resolvedRevision)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, resolvedRevision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Init provides a mock function with given fields:
func (_m *Repo) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListApps provides a mock function with given fields: resolvedRevision
func (_m *Repo) ListApps(resolvedRevision string) (map[string]string, error) {
	ret := _m.Called(resolvedRevision)

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func(string) map[string]string); ok {
		r0 = rf(resolvedRevision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(resolvedRevision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LockKey provides a mock function with given fields:
func (_m *Repo) LockKey() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ResolveAppRevision provides a mock function with given fields: app, revision
func (_m *Repo) ResolveAppRevision(app string, revision string) (string, error) {
	ret := _m.Called(app, revision)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(app, revision)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResolveRevision provides a mock function with given fields: revision
func (_m *Repo) ResolveRevision(revision string) (string, error) {
	ret := _m.Called(revision)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(revision)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(revision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RevisionMetadata provides a mock function with given fields: app, resolvedRevision
func (_m *Repo) RevisionMetadata(app string, resolvedRevision string) (*repo.RevisionMetadata, error) {
	ret := _m.Called(app, resolvedRevision)

	var r0 *repo.RevisionMetadata
	if rf, ok := ret.Get(0).(func(string, string) *repo.RevisionMetadata); ok {
		r0 = rf(app, resolvedRevision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*repo.RevisionMetadata)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(app, resolvedRevision)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
