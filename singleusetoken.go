// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// SingleUseTokenService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSingleUseTokenService] method instead.
type SingleUseTokenService struct {
	Options []option.RequestOption
}

// NewSingleUseTokenService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSingleUseTokenService(opts ...option.RequestOption) (r SingleUseTokenService) {
	r = SingleUseTokenService{}
	r.Options = opts
	return
}

// Generate a time limited single-use token with embedded authentication for
// frontend clients.
func (r *SingleUseTokenService) New(ctx context.Context, tokenType SingleUseTokenNewParamsTokenType, body SingleUseTokenNewParams, opts ...option.RequestOption) (res *SingleUseTokenNewResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("v1/single-use-token/%v", tokenType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SingleUseTokenNewResponse struct {
	// A time bound single use token that expires after 15 minutes. Will be consumed on
	// use.
	Token string `json:"token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SingleUseTokenNewResponse) RawJSON() string { return r.JSON.raw }
func (r *SingleUseTokenNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SingleUseTokenNewParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type SingleUseTokenNewParamsTokenType string

const (
	SingleUseTokenNewParamsTokenTypeRealtimeScribe SingleUseTokenNewParamsTokenType = "realtime_scribe"
	SingleUseTokenNewParamsTokenTypeTtsWebsocket   SingleUseTokenNewParamsTokenType = "tts_websocket"
)
