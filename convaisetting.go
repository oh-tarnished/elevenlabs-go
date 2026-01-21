// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiSettingService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiSettingService] method instead.
type ConvaiSettingService struct {
	Options   []option.RequestOption
	Dashboard ConvaiSettingDashboardService
}

// NewConvaiSettingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiSettingService(opts ...option.RequestOption) (r ConvaiSettingService) {
	r = ConvaiSettingService{}
	r.Options = opts
	r.Dashboard = NewConvaiSettingDashboardService(opts...)
	return
}

// Retrieve Convai settings for the workspace
func (r *ConvaiSettingService) Get(ctx context.Context, query ConvaiSettingGetParams, opts ...option.RequestOption) (res *GetConvAISettingsResponseModel, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Convai settings for the workspace
func (r *ConvaiSettingService) Update(ctx context.Context, params ConvaiSettingUpdateParams, opts ...option.RequestOption) (res *GetConvAISettingsResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/settings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type ConvAIWebhooks struct {
	// List of event types to send via webhook. Options: transcript, audio,
	// call_initiation_failure.
	//
	// Any of "transcript", "audio", "call_initiation_failure".
	Events            []string `json:"events"`
	PostCallWebhookID string   `json:"post_call_webhook_id,nullable"`
	// DEPRECATED: Use 'events' field instead. Whether to send audio data with
	// post-call webhooks for ConvAI conversations
	//
	// Deprecated: deprecated
	SendAudio bool `json:"send_audio,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Events            respjson.Field
		PostCallWebhookID respjson.Field
		SendAudio         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvAIWebhooks) RawJSON() string { return r.JSON.raw }
func (r *ConvAIWebhooks) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConvAIWebhooks to a ConvAIWebhooksParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConvAIWebhooksParam.Overrides()
func (r ConvAIWebhooks) ToParam() ConvAIWebhooksParam {
	return param.Override[ConvAIWebhooksParam](json.RawMessage(r.RawJSON()))
}

type ConvAIWebhooksParam struct {
	PostCallWebhookID param.Opt[string] `json:"post_call_webhook_id,omitzero"`
	// DEPRECATED: Use 'events' field instead. Whether to send audio data with
	// post-call webhooks for ConvAI conversations
	//
	// Deprecated: deprecated
	SendAudio param.Opt[bool] `json:"send_audio,omitzero"`
	// List of event types to send via webhook. Options: transcript, audio,
	// call_initiation_failure.
	//
	// Any of "transcript", "audio", "call_initiation_failure".
	Events []string `json:"events,omitzero"`
	paramObj
}

func (r ConvAIWebhooksParam) MarshalJSON() (data []byte, err error) {
	type shadow ConvAIWebhooksParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvAIWebhooksParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationInitiationClientDataWebhook struct {
	// The headers to send with the webhook request
	RequestHeaders map[string]ConversationInitiationClientDataWebhookRequestHeaderUnion `json:"request_headers,required"`
	// The URL to send the webhook to
	URL string `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestHeaders respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationInitiationClientDataWebhook) RawJSON() string { return r.JSON.raw }
func (r *ConversationInitiationClientDataWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationInitiationClientDataWebhook to a
// ConversationInitiationClientDataWebhookParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationInitiationClientDataWebhookParam.Overrides()
func (r ConversationInitiationClientDataWebhook) ToParam() ConversationInitiationClientDataWebhookParam {
	return param.Override[ConversationInitiationClientDataWebhookParam](json.RawMessage(r.RawJSON()))
}

// ConversationInitiationClientDataWebhookRequestHeaderUnion contains all possible
// properties and values from [string], [ConvaiSecretLocator].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type ConversationInitiationClientDataWebhookRequestHeaderUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [ConvaiSecretLocator].
	SecretID string `json:"secret_id"`
	JSON     struct {
		OfString respjson.Field
		SecretID respjson.Field
		raw      string
	} `json:"-"`
}

func (u ConversationInitiationClientDataWebhookRequestHeaderUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationInitiationClientDataWebhookRequestHeaderUnion) AsConvAISecretLocator() (v ConvaiSecretLocator) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationInitiationClientDataWebhookRequestHeaderUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationInitiationClientDataWebhookRequestHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties RequestHeaders, URL are required.
type ConversationInitiationClientDataWebhookParam struct {
	// The headers to send with the webhook request
	RequestHeaders map[string]ConversationInitiationClientDataWebhookRequestHeaderUnionParam `json:"request_headers,omitzero,required"`
	// The URL to send the webhook to
	URL string `json:"url,required"`
	paramObj
}

func (r ConversationInitiationClientDataWebhookParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationInitiationClientDataWebhookParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationInitiationClientDataWebhookParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationInitiationClientDataWebhookRequestHeaderUnionParam struct {
	OfString              param.Opt[string]         `json:",omitzero,inline"`
	OfConvAISecretLocator *ConvaiSecretLocatorParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationInitiationClientDataWebhookRequestHeaderUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConvAISecretLocator)
}
func (u *ConversationInitiationClientDataWebhookRequestHeaderUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationInitiationClientDataWebhookRequestHeaderUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfConvAISecretLocator) {
		return u.OfConvAISecretLocator
	}
	return nil
}

type GetConvAISettingsResponseModel struct {
	// Whether the workspace can use MCP servers
	CanUseMcpServers                        bool                                    `json:"can_use_mcp_servers"`
	ConversationInitiationClientDataWebhook ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,nullable"`
	// Any of "standard", "static".
	DefaultLivekitStack    LivekitStackType `json:"default_livekit_stack"`
	RagRetentionPeriodDays int64            `json:"rag_retention_period_days"`
	Webhooks               ConvAIWebhooks   `json:"webhooks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanUseMcpServers                        respjson.Field
		ConversationInitiationClientDataWebhook respjson.Field
		DefaultLivekitStack                     respjson.Field
		RagRetentionPeriodDays                  respjson.Field
		Webhooks                                respjson.Field
		ExtraFields                             map[string]respjson.Field
		raw                                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetConvAISettingsResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GetConvAISettingsResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LivekitStackType string

const (
	LivekitStackTypeStandard LivekitStackType = "standard"
	LivekitStackTypeStatic   LivekitStackType = "static"
)

type ConvaiSettingGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiSettingUpdateParams struct {
	// Whether the workspace can use MCP servers
	CanUseMcpServers       param.Opt[bool]  `json:"can_use_mcp_servers,omitzero"`
	RagRetentionPeriodDays param.Opt[int64] `json:"rag_retention_period_days,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey                                param.Opt[string]                            `header:"xi-api-key,omitzero" json:"-"`
	ConversationInitiationClientDataWebhook ConversationInitiationClientDataWebhookParam `json:"conversation_initiation_client_data_webhook,omitzero"`
	// Any of "standard", "static".
	DefaultLivekitStack LivekitStackType    `json:"default_livekit_stack,omitzero"`
	Webhooks            ConvAIWebhooksParam `json:"webhooks,omitzero"`
	paramObj
}

func (r ConvaiSettingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSettingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSettingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
