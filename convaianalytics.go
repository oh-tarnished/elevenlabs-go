// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiAnalyticsService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiAnalyticsService] method instead.
type ConvaiAnalyticsService struct {
	Options []option.RequestOption
}

// NewConvaiAnalyticsService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiAnalyticsService(opts ...option.RequestOption) (r ConvaiAnalyticsService) {
	r = ConvaiAnalyticsService{}
	r.Options = opts
	return
}

// Get the live count of the ongoing conversations.
func (r *ConvaiAnalyticsService) GetLiveCount(ctx context.Context, params ConvaiAnalyticsGetLiveCountParams, opts ...option.RequestOption) (res *ConvaiAnalyticsGetLiveCountResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/analytics/live-count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ConvaiAnalyticsGetLiveCountResponse struct {
	// The number of active ongoing conversations.
	Count int64 `json:"count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAnalyticsGetLiveCountResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAnalyticsGetLiveCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAnalyticsGetLiveCountParams struct {
	// The id of an agent to restrict the analytics to.
	AgentID param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiAnalyticsGetLiveCountParams]'s query parameters as
// `url.Values`.
func (r ConvaiAnalyticsGetLiveCountParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
