// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	ads "ads-server/internal/ads"

	context "context"

	mock "github.com/stretchr/testify/mock"

	url "net/url"

	users "ads-server/internal/users"
)

// IApp is an autogenerated mock type for the IApp type
type IApp struct {
	mock.Mock
}

// CreateAd provides a mock function with given fields: ctx, uID, title, text
func (_m *IApp) CreateAd(ctx context.Context, uID int64, title string, text string) (*ads.Ad, error) {
	ret := _m.Called(ctx, uID, title, text)

	var r0 *ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (*ads.Ad, error)); ok {
		return rf(ctx, uID, title, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *ads.Ad); ok {
		r0 = rf(ctx, uID, title, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, uID, title, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: ctx, name, email
func (_m *IApp) CreateUser(ctx context.Context, name string, email string) (*users.User, error) {
	ret := _m.Called(ctx, name, email)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*users.User, error)); ok {
		return rf(ctx, name, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *users.User); ok {
		r0 = rf(ctx, name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAd provides a mock function with given fields: ctx, adID, uID
func (_m *IApp) DeleteAd(ctx context.Context, adID int64, uID int64) error {
	ret := _m.Called(ctx, adID, uID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, adID, uID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *IApp) DeleteUser(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Filter provides a mock function with given fields: ctx, params
func (_m *IApp) Filter(ctx context.Context, params url.Values) ([]*ads.Ad, error) {
	ret := _m.Called(ctx, params)

	var r0 []*ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, url.Values) ([]*ads.Ad, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, url.Values) []*ads.Ad); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, url.Values) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUser provides a mock function with given fields: ctx, id
func (_m *IApp) FindUser(ctx context.Context, id int64) (*users.User, error) {
	ret := _m.Called(ctx, id)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*users.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *users.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdByID provides a mock function with given fields: ctx, id
func (_m *IApp) GetAdByID(ctx context.Context, id int64) (*ads.Ad, error) {
	ret := _m.Called(ctx, id)

	var r0 *ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*ads.Ad, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *ads.Ad); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAdByName provides a mock function with given fields: ctx, title
func (_m *IApp) GetAdByName(ctx context.Context, title string) []*ads.Ad {
	ret := _m.Called(ctx, title)

	var r0 []*ads.Ad
	if rf, ok := ret.Get(0).(func(context.Context, string) []*ads.Ad); ok {
		r0 = rf(ctx, title)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*ads.Ad)
		}
	}

	return r0
}

// PublishAd provides a mock function with given fields: ctx, adID, uID, action
func (_m *IApp) PublishAd(ctx context.Context, adID int64, uID int64, action bool) (*ads.Ad, error) {
	ret := _m.Called(ctx, adID, uID, action)

	var r0 *ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, bool) (*ads.Ad, error)); ok {
		return rf(ctx, adID, uID, action)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, bool) *ads.Ad); ok {
		r0 = rf(ctx, adID, uID, action)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, bool) error); ok {
		r1 = rf(ctx, adID, uID, action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAd provides a mock function with given fields: ctx, adID, uID, title, text
func (_m *IApp) UpdateAd(ctx context.Context, adID int64, uID int64, title string, text string) (*ads.Ad, error) {
	ret := _m.Called(ctx, adID, uID, title, text)

	var r0 *ads.Ad
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, string) (*ads.Ad, error)); ok {
		return rf(ctx, adID, uID, title, text)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64, string, string) *ads.Ad); ok {
		r0 = rf(ctx, adID, uID, title, text)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ads.Ad)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, int64, string, string) error); ok {
		r1 = rf(ctx, adID, uID, title, text)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, name, email
func (_m *IApp) UpdateUser(ctx context.Context, id int64, name string, email string) (*users.User, error) {
	ret := _m.Called(ctx, id, name, email)

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) (*users.User, error)); ok {
		return rf(ctx, id, name, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64, string, string) *users.User); ok {
		r0 = rf(ctx, id, name, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64, string, string) error); ok {
		r1 = rf(ctx, id, name, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIApp interface {
	mock.TestingT
	Cleanup(func())
}

// NewIApp creates a new instance of IApp. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIApp(t mockConstructorTestingTNewIApp) *IApp {
	mock := &IApp{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
