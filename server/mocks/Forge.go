// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	http "net/http"

	mock "github.com/stretchr/testify/mock"

	model "github.com/crowci/crow/v3/server/model"

	types "github.com/crowci/crow/v3/server/forge/types"
)

// Forge is an autogenerated mock type for the Forge type
type Forge struct {
	mock.Mock
}

// Activate provides a mock function with given fields: ctx, u, r, link
func (_m *Forge) Activate(ctx context.Context, u *model.User, r *model.Repo, link string) error {
	ret := _m.Called(ctx, u, r, link)

	if len(ret) == 0 {
		panic("no return value specified for Activate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, string) error); ok {
		r0 = rf(ctx, u, r, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Auth provides a mock function with given fields: ctx, token, secret
func (_m *Forge) Auth(ctx context.Context, token string, secret string) (string, error) {
	ret := _m.Called(ctx, token, secret)

	if len(ret) == 0 {
		panic("no return value specified for Auth")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (string, error)); ok {
		return rf(ctx, token, secret)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) string); ok {
		r0 = rf(ctx, token, secret)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, token, secret)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BranchHead provides a mock function with given fields: ctx, u, r, branch
func (_m *Forge) BranchHead(ctx context.Context, u *model.User, r *model.Repo, branch string) (*model.Commit, error) {
	ret := _m.Called(ctx, u, r, branch)

	if len(ret) == 0 {
		panic("no return value specified for BranchHead")
	}

	var r0 *model.Commit
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, string) (*model.Commit, error)); ok {
		return rf(ctx, u, r, branch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, string) *model.Commit); ok {
		r0 = rf(ctx, u, r, branch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Commit)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Repo, string) error); ok {
		r1 = rf(ctx, u, r, branch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Branches provides a mock function with given fields: ctx, u, r, p
func (_m *Forge) Branches(ctx context.Context, u *model.User, r *model.Repo, p *model.ListOptions) ([]string, error) {
	ret := _m.Called(ctx, u, r, p)

	if len(ret) == 0 {
		panic("no return value specified for Branches")
	}

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) ([]string, error)); ok {
		return rf(ctx, u, r, p)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) []string); ok {
		r0 = rf(ctx, u, r, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) error); ok {
		r1 = rf(ctx, u, r, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deactivate provides a mock function with given fields: ctx, u, r, link
func (_m *Forge) Deactivate(ctx context.Context, u *model.User, r *model.Repo, link string) error {
	ret := _m.Called(ctx, u, r, link)

	if len(ret) == 0 {
		panic("no return value specified for Deactivate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, string) error); ok {
		r0 = rf(ctx, u, r, link)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Dir provides a mock function with given fields: ctx, u, r, b, f
func (_m *Forge) Dir(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, f string) ([]*types.FileMeta, error) {
	ret := _m.Called(ctx, u, r, b, f)

	if len(ret) == 0 {
		panic("no return value specified for Dir")
	}

	var r0 []*types.FileMeta
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) ([]*types.FileMeta, error)); ok {
		return rf(ctx, u, r, b, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) []*types.FileMeta); ok {
		r0 = rf(ctx, u, r, b, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.FileMeta)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) error); ok {
		r1 = rf(ctx, u, r, b, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// File provides a mock function with given fields: ctx, u, r, b, f
func (_m *Forge) File(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, f string) ([]byte, error) {
	ret := _m.Called(ctx, u, r, b, f)

	if len(ret) == 0 {
		panic("no return value specified for File")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) ([]byte, error)); ok {
		return rf(ctx, u, r, b, f)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) []byte); ok {
		r0 = rf(ctx, u, r, b, f)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, string) error); ok {
		r1 = rf(ctx, u, r, b, f)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Hook provides a mock function with given fields: ctx, r
func (_m *Forge) Hook(ctx context.Context, r *http.Request) (*model.Repo, *model.Pipeline, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for Hook")
	}

	var r0 *model.Repo
	var r1 *model.Pipeline
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) (*model.Repo, *model.Pipeline, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) *model.Repo); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *http.Request) *model.Pipeline); ok {
		r1 = rf(ctx, r)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*model.Pipeline)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *http.Request) error); ok {
		r2 = rf(ctx, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: ctx, r
func (_m *Forge) Login(ctx context.Context, r *types.OAuthRequest) (*model.User, string, error) {
	ret := _m.Called(ctx, r)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *model.User
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.OAuthRequest) (*model.User, string, error)); ok {
		return rf(ctx, r)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *types.OAuthRequest) *model.User); ok {
		r0 = rf(ctx, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *types.OAuthRequest) string); ok {
		r1 = rf(ctx, r)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(context.Context, *types.OAuthRequest) error); ok {
		r2 = rf(ctx, r)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Name provides a mock function with given fields:
func (_m *Forge) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Netrc provides a mock function with given fields: u, r
func (_m *Forge) Netrc(u *model.User, r *model.Repo) (*model.Netrc, error) {
	ret := _m.Called(u, r)

	if len(ret) == 0 {
		panic("no return value specified for Netrc")
	}

	var r0 *model.Netrc
	var r1 error
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo) (*model.Netrc, error)); ok {
		return rf(u, r)
	}
	if rf, ok := ret.Get(0).(func(*model.User, *model.Repo) *model.Netrc); ok {
		r0 = rf(u, r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Netrc)
		}
	}

	if rf, ok := ret.Get(1).(func(*model.User, *model.Repo) error); ok {
		r1 = rf(u, r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Org provides a mock function with given fields: ctx, u, org
func (_m *Forge) Org(ctx context.Context, u *model.User, org string) (*model.Org, error) {
	ret := _m.Called(ctx, u, org)

	if len(ret) == 0 {
		panic("no return value specified for Org")
	}

	var r0 *model.Org
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, string) (*model.Org, error)); ok {
		return rf(ctx, u, org)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, string) *model.Org); ok {
		r0 = rf(ctx, u, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Org)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, string) error); ok {
		r1 = rf(ctx, u, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OrgMembership provides a mock function with given fields: ctx, u, org
func (_m *Forge) OrgMembership(ctx context.Context, u *model.User, org string) (*model.OrgPerm, error) {
	ret := _m.Called(ctx, u, org)

	if len(ret) == 0 {
		panic("no return value specified for OrgMembership")
	}

	var r0 *model.OrgPerm
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, string) (*model.OrgPerm, error)); ok {
		return rf(ctx, u, org)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, string) *model.OrgPerm); ok {
		r0 = rf(ctx, u, org)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.OrgPerm)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, string) error); ok {
		r1 = rf(ctx, u, org)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PullRequests provides a mock function with given fields: ctx, u, r, p
func (_m *Forge) PullRequests(ctx context.Context, u *model.User, r *model.Repo, p *model.ListOptions) ([]*model.PullRequest, error) {
	ret := _m.Called(ctx, u, r, p)

	if len(ret) == 0 {
		panic("no return value specified for PullRequests")
	}

	var r0 []*model.PullRequest
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) ([]*model.PullRequest, error)); ok {
		return rf(ctx, u, r, p)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) []*model.PullRequest); ok {
		r0 = rf(ctx, u, r, p)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.PullRequest)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, *model.Repo, *model.ListOptions) error); ok {
		r1 = rf(ctx, u, r, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repo provides a mock function with given fields: ctx, u, remoteID, owner, name
func (_m *Forge) Repo(ctx context.Context, u *model.User, remoteID model.ForgeRemoteID, owner string, name string) (*model.Repo, error) {
	ret := _m.Called(ctx, u, remoteID, owner, name)

	if len(ret) == 0 {
		panic("no return value specified for Repo")
	}

	var r0 *model.Repo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, model.ForgeRemoteID, string, string) (*model.Repo, error)); ok {
		return rf(ctx, u, remoteID, owner, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, model.ForgeRemoteID, string, string) *model.Repo); ok {
		r0 = rf(ctx, u, remoteID, owner, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Repo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User, model.ForgeRemoteID, string, string) error); ok {
		r1 = rf(ctx, u, remoteID, owner, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Repos provides a mock function with given fields: ctx, u
func (_m *Forge) Repos(ctx context.Context, u *model.User) ([]*model.Repo, error) {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for Repos")
	}

	var r0 []*model.Repo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) ([]*model.Repo, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) []*model.Repo); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Repo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Status provides a mock function with given fields: ctx, u, r, b, p
func (_m *Forge) Status(ctx context.Context, u *model.User, r *model.Repo, b *model.Pipeline, p *model.Workflow) error {
	ret := _m.Called(ctx, u, r, b, p)

	if len(ret) == 0 {
		panic("no return value specified for Status")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User, *model.Repo, *model.Pipeline, *model.Workflow) error); ok {
		r0 = rf(ctx, u, r, b, p)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Teams provides a mock function with given fields: ctx, u
func (_m *Forge) Teams(ctx context.Context, u *model.User) ([]*model.Team, error) {
	ret := _m.Called(ctx, u)

	if len(ret) == 0 {
		panic("no return value specified for Teams")
	}

	var r0 []*model.Team
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) ([]*model.Team, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) []*model.Team); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Team)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URL provides a mock function with given fields:
func (_m *Forge) URL() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for URL")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewForge creates a new instance of Forge. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewForge(t interface {
	mock.TestingT
	Cleanup(func())
}) *Forge {
	mock := &Forge{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
