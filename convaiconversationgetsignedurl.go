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

// ConvaiConversationGetSignedURLService contains methods and other services that
// help with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiConversationGetSignedURLService] method instead.
type ConvaiConversationGetSignedURLService struct {
	Options []option.RequestOption
}

// NewConvaiConversationGetSignedURLService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiConversationGetSignedURLService(opts ...option.RequestOption) (r ConvaiConversationGetSignedURLService) {
	r = ConvaiConversationGetSignedURLService{}
	r.Options = opts
	return
}

// Get a signed url to start a conversation with an agent with an agent that
// requires authorization
//
// Deprecated: deprecated
func (r *ConvaiConversationGetSignedURLService) List(ctx context.Context, params ConvaiConversationGetSignedURLListParams, opts ...option.RequestOption) (res *ConversationSignedURLResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/conversation/get_signed_url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get a signed url to start a conversation with an agent with an agent that
// requires authorization
func (r *ConvaiConversationGetSignedURLService) GetGetSignedURL(ctx context.Context, params ConvaiConversationGetSignedURLGetGetSignedURLParams, opts ...option.RequestOption) (res *ConversationSignedURLResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/conversation/get-signed-url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type ConversationSignedURLResponseModel struct {
	SignedURL string `json:"signed_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignedURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationSignedURLResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ConversationSignedURLResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetSignedURLListParams struct {
	// The id of the agent you're taking the action on.
	AgentID string `query:"agent_id,required" json:"-"`
	// The ID of the branch to use
	BranchID param.Opt[string] `query:"branch_id,omitzero" json:"-"`
	// Whether to include a conversation_id with the response. If included, the
	// conversation_signature cannot be used again.
	IncludeConversationID param.Opt[bool] `query:"include_conversation_id,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiConversationGetSignedURLListParams]'s query
// parameters as `url.Values`.
func (r ConvaiConversationGetSignedURLListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiConversationGetSignedURLGetGetSignedURLParams struct {
	// The id of the agent you're taking the action on.
	AgentID string `query:"agent_id,required" json:"-"`
	// The ID of the branch to use
	BranchID param.Opt[string] `query:"branch_id,omitzero" json:"-"`
	// Whether to include a conversation_id with the response. If included, the
	// conversation_signature cannot be used again.
	IncludeConversationID param.Opt[bool] `query:"include_conversation_id,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiConversationGetSignedURLGetGetSignedURLParams]'s
// query parameters as `url.Values`.
func (r ConvaiConversationGetSignedURLGetGetSignedURLParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
