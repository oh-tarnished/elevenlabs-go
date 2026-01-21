// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
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
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// WorkspaceWebhookService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceWebhookService] method instead.
type WorkspaceWebhookService struct {
	Options []option.RequestOption
}

// NewWorkspaceWebhookService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceWebhookService(opts ...option.RequestOption) (r WorkspaceWebhookService) {
	r = WorkspaceWebhookService{}
	r.Options = opts
	return
}

// Create a new webhook for the workspace with the specified authentication type.
func (r *WorkspaceWebhookService) New(ctx context.Context, params WorkspaceWebhookNewParams, opts ...option.RequestOption) (res *WorkspaceWebhookNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update the specified workspace webhook
func (r *WorkspaceWebhookService) Update(ctx context.Context, webhookID string, params WorkspaceWebhookUpdateParams, opts ...option.RequestOption) (res *WorkspaceWebhookUpdateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all webhooks for a workspace
func (r *WorkspaceWebhookService) List(ctx context.Context, params WorkspaceWebhookListParams, opts ...option.RequestOption) (res *WorkspaceWebhookListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/webhooks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete the specified workspace webhook
func (r *WorkspaceWebhookService) Delete(ctx context.Context, webhookID string, body WorkspaceWebhookDeleteParams, opts ...option.RequestOption) (res *WorkspaceWebhookDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if webhookID == "" {
		err = errors.New("missing required webhook_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/webhooks/%s", webhookID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type WorkspaceWebhookNewResponse struct {
	WebhookID     string `json:"webhook_id,required"`
	WebhookSecret string `json:"webhook_secret,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WebhookID     respjson.Field
		WebhookSecret respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookUpdateResponse struct {
	// The status of the workspace webhook patch request. If the request was
	// successful, the status will be 'ok'. Otherwise an error message with status 500
	// will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookListResponse struct {
	// List of webhooks currently configured for the workspace
	Webhooks []WorkspaceWebhookListResponseWebhook `json:"webhooks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Webhooks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookListResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookListResponseWebhook struct {
	// The authentication mode used to secure the webhook.
	//
	// Any of "hmac", "oauth2", "mtls".
	AuthType string `json:"auth_type,required"`
	// Original creation time of the webhook.
	CreatedAtUnix int64 `json:"created_at_unix,required"`
	// Whether the webhook has been automatically disabled due to repeated consecutive
	// failures over a long period of time.
	IsAutoDisabled bool `json:"is_auto_disabled,required"`
	// Whether the webhook has been manually disabled by a user.
	IsDisabled bool `json:"is_disabled,required"`
	// The display name for this webhook.
	Name string `json:"name,required"`
	// The unique ID for this webhook.
	WebhookID string `json:"webhook_id,required"`
	// The HTTPS callback URL that is called when this webhook is triggered in the
	// platform.
	WebhookURL string `json:"webhook_url,required"`
	// The most recent error code returned from the callback URL.
	MostRecentFailureErrorCode int64 `json:"most_recent_failure_error_code,nullable"`
	// The most recent time the webhook failed, failures are any non-200 codes returned
	// by the callback URL.
	MostRecentFailureTimestamp int64 `json:"most_recent_failure_timestamp,nullable"`
	// The list of products that are currently configured to trigger this webhook.
	Usage []WorkspaceWebhookListResponseWebhookUsage `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthType                   respjson.Field
		CreatedAtUnix              respjson.Field
		IsAutoDisabled             respjson.Field
		IsDisabled                 respjson.Field
		Name                       respjson.Field
		WebhookID                  respjson.Field
		WebhookURL                 respjson.Field
		MostRecentFailureErrorCode respjson.Field
		MostRecentFailureTimestamp respjson.Field
		Usage                      respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookListResponseWebhook) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookListResponseWebhook) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookListResponseWebhookUsage struct {
	// Any of "ConvAI Agent Settings", "ConvAI Settings", "Voice Library Removal
	// Notices", "Speech to Text".
	UsageType string `json:"usage_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UsageType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookListResponseWebhookUsage) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookListResponseWebhookUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookDeleteResponse struct {
	// The status of the workspace webhook deletion request. If the request was
	// successful, the status will be 'ok'. Otherwise an error message with status 500
	// will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceWebhookDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceWebhookDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookNewParams struct {
	// Webhook settings object containing auth_type and corresponding configuration
	Settings WorkspaceWebhookNewParamsSettings `json:"settings,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceWebhookNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceWebhookNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceWebhookNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Webhook settings object containing auth_type and corresponding configuration
//
// The properties AuthType, Name, WebhookURL are required.
type WorkspaceWebhookNewParamsSettings struct {
	// The display name for this webhook
	Name string `json:"name,required"`
	// The HTTPS callback URL that will be called when this webhook is triggered
	WebhookURL string `json:"webhook_url,required"`
	// The authentication type for this webhook
	//
	// This field can be elided, and will marshal its zero value as "hmac".
	AuthType constant.Hmac `json:"auth_type,required"`
	paramObj
}

func (r WorkspaceWebhookNewParamsSettings) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceWebhookNewParamsSettings
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceWebhookNewParamsSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookUpdateParams struct {
	// Whether to disable or enable the webhook
	IsDisabled bool `json:"is_disabled,required"`
	// The display name of the webhook (used for display purposes only).
	Name string `json:"name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceWebhookUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceWebhookUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceWebhookUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceWebhookListParams struct {
	// Whether to include active usages of the webhook, only usable by admins
	IncludeUsages param.Opt[bool] `query:"include_usages,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceWebhookListParams]'s query parameters as
// `url.Values`.
func (r WorkspaceWebhookListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceWebhookDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
