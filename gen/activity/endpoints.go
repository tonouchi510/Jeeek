// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity endpoints
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package activity

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "Activity" service endpoints.
type Endpoints struct {
	ManualPostOfActivity                      goa.Endpoint
	RefreshActivitiesOfAllCooperationServices goa.Endpoint
	RefreshQiitaActivities                    goa.Endpoint
	PickOutAllPastActivitiesOfQiita           goa.Endpoint
}

// NewEndpoints wraps the methods of the "Activity" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		ManualPostOfActivity:                      NewManualPostOfActivityEndpoint(s, a.JWTAuth),
		RefreshActivitiesOfAllCooperationServices: NewRefreshActivitiesOfAllCooperationServicesEndpoint(s, a.JWTAuth),
		RefreshQiitaActivities:                    NewRefreshQiitaActivitiesEndpoint(s, a.JWTAuth),
		PickOutAllPastActivitiesOfQiita:           NewPickOutAllPastActivitiesOfQiitaEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "Activity" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.ManualPostOfActivity = m(e.ManualPostOfActivity)
	e.RefreshActivitiesOfAllCooperationServices = m(e.RefreshActivitiesOfAllCooperationServices)
	e.RefreshQiitaActivities = m(e.RefreshQiitaActivities)
	e.PickOutAllPastActivitiesOfQiita = m(e.PickOutAllPastActivitiesOfQiita)
}

// NewManualPostOfActivityEndpoint returns an endpoint function that calls the
// method "Manual post of activity" of service "Activity".
func NewManualPostOfActivityEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ActivityPostPayload)
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
		return nil, s.ManualPostOfActivity(ctx, p)
	}
}

// NewRefreshActivitiesOfAllCooperationServicesEndpoint returns an endpoint
// function that calls the method "Refresh activities of all cooperation
// services" of service "Activity".
func NewRefreshActivitiesOfAllCooperationServicesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
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
		return nil, s.RefreshActivitiesOfAllCooperationServices(ctx, p)
	}
}

// NewRefreshQiitaActivitiesEndpoint returns an endpoint function that calls
// the method "Refresh qiita activities" of service "Activity".
func NewRefreshQiitaActivitiesEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
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
		return nil, s.RefreshQiitaActivities(ctx, p)
	}
}

// NewPickOutAllPastActivitiesOfQiitaEndpoint returns an endpoint function that
// calls the method "Pick out all past activities of qiita" of service
// "Activity".
func NewPickOutAllPastActivitiesOfQiitaEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
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
		return nil, s.PickOutAllPastActivitiesOfQiita(ctx, p)
	}
}
