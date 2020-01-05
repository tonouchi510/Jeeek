// Code generated by goa v3.0.7, DO NOT EDIT.
//
// ExternalActivity client HTTP transport
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the ExternalActivity service endpoint HTTP clients.
type Client struct {
	// RefreshActivitiesOfExternalServices Doer is the HTTP client used to make
	// requests to the Refresh activities of external services endpoint.
	RefreshActivitiesOfExternalServicesDoer goahttp.Doer

	// RefreshQiitaActivity Doer is the HTTP client used to make requests to the
	// Refresh qiita activity endpoint.
	RefreshQiitaActivityDoer goahttp.Doer

	// PickOutPastActivityOfQiita Doer is the HTTP client used to make requests to
	// the Pick out past activity of qiita endpoint.
	PickOutPastActivityOfQiitaDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the ExternalActivity service
// servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		RefreshActivitiesOfExternalServicesDoer: doer,
		RefreshQiitaActivityDoer:                doer,
		PickOutPastActivityOfQiitaDoer:          doer,
		RestoreResponseBody:                     restoreBody,
		scheme:                                  scheme,
		host:                                    host,
		decoder:                                 dec,
		encoder:                                 enc,
	}
}

// RefreshActivitiesOfExternalServices returns an endpoint that makes HTTP
// requests to the ExternalActivity service Refresh activities of external
// services server.
func (c *Client) RefreshActivitiesOfExternalServices() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshActivitiesOfExternalServicesRequest(c.encoder)
		decodeResponse = DecodeRefreshActivitiesOfExternalServicesResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRefreshActivitiesOfExternalServicesRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshActivitiesOfExternalServicesDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("ExternalActivity", "Refresh activities of external services", err)
		}
		return decodeResponse(resp)
	}
}

// RefreshQiitaActivity returns an endpoint that makes HTTP requests to the
// ExternalActivity service Refresh qiita activity server.
func (c *Client) RefreshQiitaActivity() goa.Endpoint {
	var (
		encodeRequest  = EncodeRefreshQiitaActivityRequest(c.encoder)
		decodeResponse = DecodeRefreshQiitaActivityResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildRefreshQiitaActivityRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.RefreshQiitaActivityDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("ExternalActivity", "Refresh qiita activity", err)
		}
		return decodeResponse(resp)
	}
}

// PickOutPastActivityOfQiita returns an endpoint that makes HTTP requests to
// the ExternalActivity service Pick out past activity of qiita server.
func (c *Client) PickOutPastActivityOfQiita() goa.Endpoint {
	var (
		encodeRequest  = EncodePickOutPastActivityOfQiitaRequest(c.encoder)
		decodeResponse = DecodePickOutPastActivityOfQiitaResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildPickOutPastActivityOfQiitaRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.PickOutPastActivityOfQiitaDoer.Do(req)

		if err != nil {
			return nil, goahttp.ErrRequestError("ExternalActivity", "Pick out past activity of qiita", err)
		}
		return decodeResponse(resp)
	}
}
