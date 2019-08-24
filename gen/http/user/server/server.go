// Code generated by goa v3.0.4, DO NOT EDIT.
//
// User HTTP server
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"net/http"

	user "github.com/tonouchi510/Jeeek/gen/user"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the User service endpoint HTTP handlers.
type Server struct {
	Mounts         []*MountPoint
	GetCurrentUser http.Handler
	UpdateUser     http.Handler
	ListUser       http.Handler
	GetUser        http.Handler
	DeleteUser     http.Handler
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

// New instantiates HTTP handlers for all the User service endpoints.
func New(
	e *user.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"GetCurrentUser", "GET", "/v1/users/me"},
			{"UpdateUser", "PUT", "/v1/users"},
			{"ListUser", "GET", "/v1/users"},
			{"GetUser", "GET", "/v1/users/{user_id}"},
			{"DeleteUser", "DELETE", "/v1/users"},
		},
		GetCurrentUser: NewGetCurrentUserHandler(e.GetCurrentUser, mux, dec, enc, eh),
		UpdateUser:     NewUpdateUserHandler(e.UpdateUser, mux, dec, enc, eh),
		ListUser:       NewListUserHandler(e.ListUser, mux, dec, enc, eh),
		GetUser:        NewGetUserHandler(e.GetUser, mux, dec, enc, eh),
		DeleteUser:     NewDeleteUserHandler(e.DeleteUser, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "User" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.GetCurrentUser = m(s.GetCurrentUser)
	s.UpdateUser = m(s.UpdateUser)
	s.ListUser = m(s.ListUser)
	s.GetUser = m(s.GetUser)
	s.DeleteUser = m(s.DeleteUser)
}

// Mount configures the mux to serve the User endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountGetCurrentUserHandler(mux, h.GetCurrentUser)
	MountUpdateUserHandler(mux, h.UpdateUser)
	MountListUserHandler(mux, h.ListUser)
	MountGetUserHandler(mux, h.GetUser)
	MountDeleteUserHandler(mux, h.DeleteUser)
}

// MountGetCurrentUserHandler configures the mux to serve the "User" service
// "Get current user" endpoint.
func MountGetCurrentUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users/me", f)
}

// NewGetCurrentUserHandler creates a HTTP handler which loads the HTTP request
// and calls the "User" service "Get current user" endpoint.
func NewGetCurrentUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeGetCurrentUserRequest(mux, dec)
		encodeResponse = EncodeGetCurrentUserResponse(enc)
		encodeError    = EncodeGetCurrentUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Get current user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
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

// MountUpdateUserHandler configures the mux to serve the "User" service
// "Update user" endpoint.
func MountUpdateUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("PUT", "/v1/users", f)
}

// NewUpdateUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "Update user" endpoint.
func NewUpdateUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeUpdateUserRequest(mux, dec)
		encodeResponse = EncodeUpdateUserResponse(enc)
		encodeError    = EncodeUpdateUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Update user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
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

// MountListUserHandler configures the mux to serve the "User" service "List
// user" endpoint.
func MountListUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users", f)
}

// NewListUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "List user" endpoint.
func NewListUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeListUserRequest(mux, dec)
		encodeResponse = EncodeListUserResponse(enc)
		encodeError    = EncodeListUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "List user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
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

// MountGetUserHandler configures the mux to serve the "User" service "Get
// user" endpoint.
func MountGetUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/users/{user_id}", f)
}

// NewGetUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "Get user" endpoint.
func NewGetUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeGetUserRequest(mux, dec)
		encodeResponse = EncodeGetUserResponse(enc)
		encodeError    = EncodeGetUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Get user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
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

// MountDeleteUserHandler configures the mux to serve the "User" service
// "Delete user" endpoint.
func MountDeleteUserHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/v1/users", f)
}

// NewDeleteUserHandler creates a HTTP handler which loads the HTTP request and
// calls the "User" service "Delete user" endpoint.
func NewDeleteUserHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUserRequest(mux, dec)
		encodeResponse = EncodeDeleteUserResponse(enc)
		encodeError    = EncodeDeleteUserError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Delete user")
		ctx = context.WithValue(ctx, goa.ServiceKey, "User")
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
