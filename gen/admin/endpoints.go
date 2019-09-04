// Code generated by goa v3.0.4, DO NOT EDIT.
//
// Admin endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package admin

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "Admin" service endpoints.
type Endpoints struct {
	AdminHealthCheck   goa.Endpoint
	AdminSignin        goa.Endpoint
	AdminCreateNewUser goa.Endpoint
	AdminUpdateUser    goa.Endpoint
	AdminListUser      goa.Endpoint
	AdminGetUser       goa.Endpoint
	AdminDeleteUser    goa.Endpoint
}

// NewEndpoints wraps the methods of the "Admin" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		AdminHealthCheck:   NewAdminHealthCheckEndpoint(s, a.JWTAuth),
		AdminSignin:        NewAdminSigninEndpoint(s),
		AdminCreateNewUser: NewAdminCreateNewUserEndpoint(s, a.JWTAuth),
		AdminUpdateUser:    NewAdminUpdateUserEndpoint(s, a.JWTAuth),
		AdminListUser:      NewAdminListUserEndpoint(s, a.JWTAuth),
		AdminGetUser:       NewAdminGetUserEndpoint(s, a.JWTAuth),
		AdminDeleteUser:    NewAdminDeleteUserEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "Admin" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.AdminHealthCheck = m(e.AdminHealthCheck)
	e.AdminSignin = m(e.AdminSignin)
	e.AdminCreateNewUser = m(e.AdminCreateNewUser)
	e.AdminUpdateUser = m(e.AdminUpdateUser)
	e.AdminListUser = m(e.AdminListUser)
	e.AdminGetUser = m(e.AdminGetUser)
	e.AdminDeleteUser = m(e.AdminDeleteUser)
}

// NewAdminHealthCheckEndpoint returns an endpoint function that calls the
// method "admin health-check" of service "Admin".
func NewAdminHealthCheckEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SessionTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, err := s.AdminHealthCheck(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekHealthcheck(res, "default")
		return vres, nil
	}
}

// NewAdminSigninEndpoint returns an endpoint function that calls the method
// "admin signin" of service "Admin".
func NewAdminSigninEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminSignInPayload)
		res, err := s.AdminSignin(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekAdminSignin(res, "default")
		return vres, nil
	}
}

// NewAdminCreateNewUserEndpoint returns an endpoint function that calls the
// method "admin create new user" of service "Admin".
func NewAdminCreateNewUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminCreateUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.AdminCreateNewUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewAdminUpdateUserEndpoint returns an endpoint function that calls the
// method "admin update user" of service "Admin".
func NewAdminUpdateUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminUpdateUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.AdminUpdateUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewAdminListUserEndpoint returns an endpoint function that calls the method
// "admin list user" of service "Admin".
func NewAdminListUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*SessionTokenPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.AdminListUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUserCollection(res, view)
		return vres, nil
	}
}

// NewAdminGetUserEndpoint returns an endpoint function that calls the method
// "admin get user" of service "Admin".
func NewAdminGetUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		res, view, err := s.AdminGetUser(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedJeeekUser(res, view)
		return vres, nil
	}
}

// NewAdminDeleteUserEndpoint returns an endpoint function that calls the
// method "admin delete user" of service "Admin".
func NewAdminDeleteUserEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AdminDeleteUserPayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{},
			RequiredScopes: []string{},
		}
		var token string
		if p.Token != nil {
			token = *p.Token
		}
		ctx, err = authJWTFn(ctx, token, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.AdminDeleteUser(ctx, p)
	}
}