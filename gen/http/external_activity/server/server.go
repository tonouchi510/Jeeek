// Code generated by goa v3.0.7, DO NOT EDIT.
//
// ExternalActivity HTTP server
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"net/http"

	externalactivity "github.com/tonouchi510/Jeeek/gen/external_activity"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the ExternalActivity service endpoint HTTP handlers.
type Server struct {
	Mounts                              []*MountPoint
	RefreshActivitiesOfExternalServices http.Handler
	RefreshQiitaActivity                http.Handler
	PickOutPastActivityOfQiita          http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the ExternalActivity service
// endpoints.
func New(
	e *externalactivity.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"RefreshActivitiesOfExternalServices", "GET", "/v1/external_activity/batch"},
			{"RefreshQiitaActivity", "GET", "/v1/external_activity/qiita"},
			{"PickOutPastActivityOfQiita", "GET", "/v1/external_activity/qiita/initialization"},
		},
		RefreshActivitiesOfExternalServices: NewRefreshActivitiesOfExternalServicesHandler(e.RefreshActivitiesOfExternalServices, mux, dec, enc, eh),
		RefreshQiitaActivity:                NewRefreshQiitaActivityHandler(e.RefreshQiitaActivity, mux, dec, enc, eh),
		PickOutPastActivityOfQiita:          NewPickOutPastActivityOfQiitaHandler(e.PickOutPastActivityOfQiita, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "ExternalActivity" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.RefreshActivitiesOfExternalServices = m(s.RefreshActivitiesOfExternalServices)
	s.RefreshQiitaActivity = m(s.RefreshQiitaActivity)
	s.PickOutPastActivityOfQiita = m(s.PickOutPastActivityOfQiita)
}

// Mount configures the mux to serve the ExternalActivity endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountRefreshActivitiesOfExternalServicesHandler(mux, h.RefreshActivitiesOfExternalServices)
	MountRefreshQiitaActivityHandler(mux, h.RefreshQiitaActivity)
	MountPickOutPastActivityOfQiitaHandler(mux, h.PickOutPastActivityOfQiita)
}

// MountRefreshActivitiesOfExternalServicesHandler configures the mux to serve
// the "ExternalActivity" service "Refresh activities of external services"
// endpoint.
func MountRefreshActivitiesOfExternalServicesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/external_activity/batch", f)
}

// NewRefreshActivitiesOfExternalServicesHandler creates a HTTP handler which
// loads the HTTP request and calls the "ExternalActivity" service "Refresh
// activities of external services" endpoint.
func NewRefreshActivitiesOfExternalServicesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRefreshActivitiesOfExternalServicesRequest(mux, dec)
		encodeResponse = EncodeRefreshActivitiesOfExternalServicesResponse(enc)
		encodeError    = EncodeRefreshActivitiesOfExternalServicesError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Refresh activities of external services")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ExternalActivity")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountRefreshQiitaActivityHandler configures the mux to serve the
// "ExternalActivity" service "Refresh qiita activity" endpoint.
func MountRefreshQiitaActivityHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/external_activity/qiita", f)
}

// NewRefreshQiitaActivityHandler creates a HTTP handler which loads the HTTP
// request and calls the "ExternalActivity" service "Refresh qiita activity"
// endpoint.
func NewRefreshQiitaActivityHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeRefreshQiitaActivityRequest(mux, dec)
		encodeResponse = EncodeRefreshQiitaActivityResponse(enc)
		encodeError    = EncodeRefreshQiitaActivityError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Refresh qiita activity")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ExternalActivity")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountPickOutPastActivityOfQiitaHandler configures the mux to serve the
// "ExternalActivity" service "Pick out past activity of qiita" endpoint.
func MountPickOutPastActivityOfQiitaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/external_activity/qiita/initialization", f)
}

// NewPickOutPastActivityOfQiitaHandler creates a HTTP handler which loads the
// HTTP request and calls the "ExternalActivity" service "Pick out past
// activity of qiita" endpoint.
func NewPickOutPastActivityOfQiitaHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodePickOutPastActivityOfQiitaRequest(mux, dec)
		encodeResponse = EncodePickOutPastActivityOfQiitaResponse(enc)
		encodeError    = EncodePickOutPastActivityOfQiitaError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Pick out past activity of qiita")
		ctx = context.WithValue(ctx, goa.ServiceKey, "ExternalActivity")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}
