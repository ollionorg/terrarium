// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	db "github.com/cldcvr/terrarium/src/pkg/db"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// CreateDependencyAttribute provides a mock function with given fields: e
func (_m *DB) CreateDependencyAttribute(e *db.DependencyAttribute) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.DependencyAttribute) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.DependencyAttribute) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.DependencyAttribute) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDependencyInterface provides a mock function with given fields: e
func (_m *DB) CreateDependencyInterface(e *db.Dependency) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.Dependency) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.Dependency) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.Dependency) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlatform provides a mock function with given fields: p
func (_m *DB) CreatePlatform(p *db.Platform) (uuid.UUID, error) {
	ret := _m.Called(p)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.Platform) (uuid.UUID, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(*db.Platform) uuid.UUID); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.Platform) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreatePlatformComponents provides a mock function with given fields: p
func (_m *DB) CreatePlatformComponents(p *db.PlatformComponent) (uuid.UUID, error) {
	ret := _m.Called(p)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.PlatformComponent) (uuid.UUID, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func(*db.PlatformComponent) uuid.UUID); ok {
		r0 = rf(p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.PlatformComponent) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateRelease provides a mock function with given fields: e
func (_m *DB) CreateRelease(e *db.FarmRelease) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.FarmRelease) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.FarmRelease) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.FarmRelease) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFModule provides a mock function with given fields: e
func (_m *DB) CreateTFModule(e *db.TFModule) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFModule) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFModule) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFModule) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFModuleAttribute provides a mock function with given fields: e
func (_m *DB) CreateTFModuleAttribute(e *db.TFModuleAttribute) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFModuleAttribute) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFModuleAttribute) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFModuleAttribute) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFProvider provides a mock function with given fields: e
func (_m *DB) CreateTFProvider(e *db.TFProvider) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFProvider) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFProvider) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFProvider) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFResourceAttribute provides a mock function with given fields: e
func (_m *DB) CreateTFResourceAttribute(e *db.TFResourceAttribute) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFResourceAttribute) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFResourceAttribute) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFResourceAttribute) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFResourceAttributesMapping provides a mock function with given fields: e
func (_m *DB) CreateTFResourceAttributesMapping(e *db.TFResourceAttributesMapping) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFResourceAttributesMapping) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFResourceAttributesMapping) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFResourceAttributesMapping) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTFResourceType provides a mock function with given fields: e
func (_m *DB) CreateTFResourceType(e *db.TFResourceType) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.TFResourceType) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFResourceType) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFResourceType) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTaxonomy provides a mock function with given fields: e
func (_m *DB) CreateTaxonomy(e *db.Taxonomy) (uuid.UUID, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 error
	if rf, ok := ret.Get(0).(func(*db.Taxonomy) (uuid.UUID, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.Taxonomy) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.Taxonomy) error); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteSQLStatement provides a mock function with given fields: statement
func (_m *DB) ExecuteSQLStatement(statement string) error {
	ret := _m.Called(statement)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(statement)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindReleaseByRepo provides a mock function with given fields: e, repo
func (_m *DB) FindReleaseByRepo(e *db.FarmRelease, repo string) error {
	ret := _m.Called(e, repo)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.FarmRelease, string) error); ok {
		r0 = rf(e, repo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetchdeps provides a mock function with given fields:
func (_m *DB) Fetchdeps() []db.DependencyResult {
	ret := _m.Called()

	var r0 []db.DependencyResult
	if rf, ok := ret.Get(0).(func() []db.DependencyResult); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]db.DependencyResult)
		}
	}

	return r0
}

// GetOrCreateTFProvider provides a mock function with given fields: e
func (_m *DB) GetOrCreateTFProvider(e *db.TFProvider) (uuid.UUID, bool, error) {
	ret := _m.Called(e)

	var r0 uuid.UUID
	var r1 bool
	var r2 error
	if rf, ok := ret.Get(0).(func(*db.TFProvider) (uuid.UUID, bool, error)); ok {
		return rf(e)
	}
	if rf, ok := ret.Get(0).(func(*db.TFProvider) uuid.UUID); ok {
		r0 = rf(e)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uuid.UUID)
		}
	}

	if rf, ok := ret.Get(1).(func(*db.TFProvider) bool); ok {
		r1 = rf(e)
	} else {
		r1 = ret.Get(1).(bool)
	}

	if rf, ok := ret.Get(2).(func(*db.TFProvider) error); ok {
		r2 = rf(e)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetTFProvider provides a mock function with given fields: e, where
func (_m *DB) GetTFProvider(e *db.TFProvider, where *db.TFProvider) error {
	ret := _m.Called(e, where)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.TFProvider, *db.TFProvider) error); ok {
		r0 = rf(e, where)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTFResourceAttribute provides a mock function with given fields: e, where
func (_m *DB) GetTFResourceAttribute(e *db.TFResourceAttribute, where *db.TFResourceAttribute) error {
	ret := _m.Called(e, where)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.TFResourceAttribute, *db.TFResourceAttribute) error); ok {
		r0 = rf(e, where)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTFResourceType provides a mock function with given fields: e, where
func (_m *DB) GetTFResourceType(e *db.TFResourceType, where *db.TFResourceType) error {
	ret := _m.Called(e, where)

	var r0 error
	if rf, ok := ret.Get(0).(func(*db.TFResourceType, *db.TFResourceType) error); ok {
		r0 = rf(e, where)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// QueryDependencies provides a mock function with given fields: filterOps
func (_m *DB) QueryDependencies(filterOps ...db.FilterOption) (db.Dependencies, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Dependencies
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.Dependencies, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.Dependencies); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Dependencies)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryPlatformComponents provides a mock function with given fields: filterOps
func (_m *DB) QueryPlatformComponents(filterOps ...db.FilterOption) (db.PlatformComponents, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.PlatformComponents
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.PlatformComponents, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.PlatformComponents); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.PlatformComponents)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryPlatforms provides a mock function with given fields: filterOps
func (_m *DB) QueryPlatforms(filterOps ...db.FilterOption) (db.Platforms, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Platforms
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.Platforms, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.Platforms); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Platforms)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTFModuleAttributes provides a mock function with given fields: filterOps
func (_m *DB) QueryTFModuleAttributes(filterOps ...db.FilterOption) (db.TFModuleAttributes, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.TFModuleAttributes
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.TFModuleAttributes, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.TFModuleAttributes); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.TFModuleAttributes)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTFModules provides a mock function with given fields: filterOps
func (_m *DB) QueryTFModules(filterOps ...db.FilterOption) (db.TFModules, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.TFModules
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.TFModules, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.TFModules); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.TFModules)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryTaxonomies provides a mock function with given fields: filterOps
func (_m *DB) QueryTaxonomies(filterOps ...db.FilterOption) (db.Taxonomies, error) {
	_va := make([]interface{}, len(filterOps))
	for _i := range filterOps {
		_va[_i] = filterOps[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 db.Taxonomies
	var r1 error
	if rf, ok := ret.Get(0).(func(...db.FilterOption) (db.Taxonomies, error)); ok {
		return rf(filterOps...)
	}
	if rf, ok := ret.Get(0).(func(...db.FilterOption) db.Taxonomies); ok {
		r0 = rf(filterOps...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Taxonomies)
		}
	}

	if rf, ok := ret.Get(1).(func(...db.FilterOption) error); ok {
		r1 = rf(filterOps...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDB interface {
	mock.TestingT
	Cleanup(func())
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDB(t mockConstructorTestingTNewDB) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
