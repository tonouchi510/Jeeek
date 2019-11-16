// Code generated by goa v3.0.7, DO NOT EDIT.
//
// Activity HTTP server
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package server

import (
	"context"
	"net/http"

	activity "github.com/tonouchi510/Jeeek/gen/activity"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Server lists the Activity service endpoint HTTP handlers.
type Server struct {
	Mounts                               []*MountPoint
	FetchQiitaArticleByQiitaUserID       http.Handler
	BatchJobMethodToRefreshQiitaActivity http.Handler
	PickOutPastActivityOfQiita           http.Handler
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

// New instantiates HTTP handlers for all the Activity service endpoints.
func New(
	e *activity.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"FetchQiitaArticleByQiitaUserID", "GET", "/v1/activity/qiita/{user_id}"},
			{"BatchJobMethodToRefreshQiitaActivity", "GET", "/v1/activity/qiita/batch"},
			{"PickOutPastActivityOfQiita", "GET", "/v1/activity/qiita/initialization"},
		},
		FetchQiitaArticleByQiitaUserID:       NewFetchQiitaArticleByQiitaUserIDHandler(e.FetchQiitaArticleByQiitaUserID, mux, dec, enc, eh),
		BatchJobMethodToRefreshQiitaActivity: NewBatchJobMethodToRefreshQiitaActivityHandler(e.BatchJobMethodToRefreshQiitaActivity, mux, dec, enc, eh),
		PickOutPastActivityOfQiita:           NewPickOutPastActivityOfQiitaHandler(e.PickOutPastActivityOfQiita, mux, dec, enc, eh),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "Activity" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.FetchQiitaArticleByQiitaUserID = m(s.FetchQiitaArticleByQiitaUserID)
	s.BatchJobMethodToRefreshQiitaActivity = m(s.BatchJobMethodToRefreshQiitaActivity)
	s.PickOutPastActivityOfQiita = m(s.PickOutPastActivityOfQiita)
}

// Mount configures the mux to serve the Activity endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountFetchQiitaArticleByQiitaUserIDHandler(mux, h.FetchQiitaArticleByQiitaUserID)
	MountBatchJobMethodToRefreshQiitaActivityHandler(mux, h.BatchJobMethodToRefreshQiitaActivity)
	MountPickOutPastActivityOfQiitaHandler(mux, h.PickOutPastActivityOfQiita)
}

// MountFetchQiitaArticleByQiitaUserIDHandler configures the mux to serve the
// "Activity" service "Fetch qiita article by qiita-user-id" endpoint.
func MountFetchQiitaArticleByQiitaUserIDHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/activity/qiita/{user_id}", f)
}

// NewFetchQiitaArticleByQiitaUserIDHandler creates a HTTP handler which loads
// the HTTP request and calls the "Activity" service "Fetch qiita article by
// qiita-user-id" endpoint.
func NewFetchQiitaArticleByQiitaUserIDHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeFetchQiitaArticleByQiitaUserIDRequest(mux, dec)
		encodeResponse = EncodeFetchQiitaArticleByQiitaUserIDResponse(enc)
		encodeError    = EncodeFetchQiitaArticleByQiitaUserIDError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Fetch qiita article by qiita-user-id")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Activity")
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

// MountBatchJobMethodToRefreshQiitaActivityHandler configures the mux to serve
// the "Activity" service "Batch job method to refresh qiita activity" endpoint.
func MountBatchJobMethodToRefreshQiitaActivityHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/activity/qiita/batch", f)
}

// NewBatchJobMethodToRefreshQiitaActivityHandler creates a HTTP handler which
// loads the HTTP request and calls the "Activity" service "Batch job method to
// refresh qiita activity" endpoint.
func NewBatchJobMethodToRefreshQiitaActivityHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		encodeResponse = EncodeBatchJobMethodToRefreshQiitaActivityResponse(enc)
		encodeError    = EncodeBatchJobMethodToRefreshQiitaActivityError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "Batch job method to refresh qiita activity")
		ctx = context.WithValue(ctx, goa.ServiceKey, "Activity")
		var err error

		res, err := endpoint(ctx, nil)

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
// "Activity" service "Pick out past activity of qiita" endpoint.
func MountPickOutPastActivityOfQiitaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/v1/activity/qiita/initialization", f)
}

// NewPickOutPastActivityOfQiitaHandler creates a HTTP handler which loads the
// HTTP request and calls the "Activity" service "Pick out past activity of
// qiita" endpoint.
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
		ctx = context.WithValue(ctx, goa.ServiceKey, "Activity")
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
