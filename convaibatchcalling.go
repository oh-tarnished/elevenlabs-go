// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
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
)

// ConvaiBatchCallingService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiBatchCallingService] method instead.
type ConvaiBatchCallingService struct {
	Options []option.RequestOption
}

// NewConvaiBatchCallingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiBatchCallingService(opts ...option.RequestOption) (r ConvaiBatchCallingService) {
	r = ConvaiBatchCallingService{}
	r.Options = opts
	return
}

// Get detailed information about a batch call including all recipients.
func (r *ConvaiBatchCallingService) Get(ctx context.Context, batchID string, query ConvaiBatchCallingGetParams, opts ...option.RequestOption) (res *ConvaiBatchCallingGetResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/batch-calling/%s", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all batch calls for the current workspace.
func (r *ConvaiBatchCallingService) List(ctx context.Context, params ConvaiBatchCallingListParams, opts ...option.RequestOption) (res *ConvaiBatchCallingListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/batch-calling/workspace"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Permanently delete a batch call and all recipient records. Conversations remain
// in history.
func (r *ConvaiBatchCallingService) Delete(ctx context.Context, batchID string, body ConvaiBatchCallingDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/batch-calling/%s", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Cancel a running batch call and set all recipients to cancelled status.
func (r *ConvaiBatchCallingService) Cancel(ctx context.Context, batchID string, body ConvaiBatchCallingCancelParams, opts ...option.RequestOption) (res *BatchCallResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/batch-calling/%s/cancel", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Retry a batch call, calling failed and no-response recipients again.
func (r *ConvaiBatchCallingService) Retry(ctx context.Context, batchID string, body ConvaiBatchCallingRetryParams, opts ...option.RequestOption) (res *BatchCallResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if batchID == "" {
		err = errors.New("missing required batch_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/batch-calling/%s/retry", batchID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Submit a batch call request to schedule calls for multiple recipients.
func (r *ConvaiBatchCallingService) Submit(ctx context.Context, params ConvaiBatchCallingSubmitParams, opts ...option.RequestOption) (res *BatchCallResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/batch-calling/submit"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type BatchCallResponse struct {
	ID                string `json:"id,required"`
	AgentID           string `json:"agent_id,required"`
	AgentName         string `json:"agent_name,required"`
	CreatedAtUnix     int64  `json:"created_at_unix,required"`
	LastUpdatedAtUnix int64  `json:"last_updated_at_unix,required"`
	Name              string `json:"name,required"`
	PhoneNumberID     string `json:"phone_number_id,required"`
	// Any of "twilio", "sip_trunk".
	PhoneProvider     TelephonyProvider `json:"phone_provider,required"`
	RetryCount        int64             `json:"retry_count,required"`
	ScheduledTimeUnix int64             `json:"scheduled_time_unix,required"`
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	Status               BatchCallStatus             `json:"status,required"`
	Timezone             string                      `json:"timezone,required"`
	TotalCallsDispatched int64                       `json:"total_calls_dispatched,required"`
	TotalCallsFinished   int64                       `json:"total_calls_finished,required"`
	TotalCallsScheduled  int64                       `json:"total_calls_scheduled,required"`
	WhatsappParams       BatchCallWhatsappParamsResp `json:"whatsapp_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AgentID              respjson.Field
		AgentName            respjson.Field
		CreatedAtUnix        respjson.Field
		LastUpdatedAtUnix    respjson.Field
		Name                 respjson.Field
		PhoneNumberID        respjson.Field
		PhoneProvider        respjson.Field
		RetryCount           respjson.Field
		ScheduledTimeUnix    respjson.Field
		Status               respjson.Field
		Timezone             respjson.Field
		TotalCallsDispatched respjson.Field
		TotalCallsFinished   respjson.Field
		TotalCallsScheduled  respjson.Field
		WhatsappParams       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCallResponse) RawJSON() string { return r.JSON.raw }
func (r *BatchCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BatchCallStatus string

const (
	BatchCallStatusPending    BatchCallStatus = "pending"
	BatchCallStatusInProgress BatchCallStatus = "in_progress"
	BatchCallStatusCompleted  BatchCallStatus = "completed"
	BatchCallStatusFailed     BatchCallStatus = "failed"
	BatchCallStatusCancelled  BatchCallStatus = "cancelled"
)

type BatchCallWhatsappParamsResp struct {
	WhatsappCallPermissionRequestTemplateLanguageCode string `json:"whatsapp_call_permission_request_template_language_code,required"`
	WhatsappCallPermissionRequestTemplateName         string `json:"whatsapp_call_permission_request_template_name,required"`
	WhatsappPhoneNumberID                             string `json:"whatsapp_phone_number_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WhatsappCallPermissionRequestTemplateLanguageCode respjson.Field
		WhatsappCallPermissionRequestTemplateName         respjson.Field
		WhatsappPhoneNumberID                             respjson.Field
		ExtraFields                                       map[string]respjson.Field
		raw                                               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BatchCallWhatsappParamsResp) RawJSON() string { return r.JSON.raw }
func (r *BatchCallWhatsappParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this BatchCallWhatsappParamsResp to a BatchCallWhatsappParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// BatchCallWhatsappParams.Overrides()
func (r BatchCallWhatsappParamsResp) ToParam() BatchCallWhatsappParams {
	return param.Override[BatchCallWhatsappParams](json.RawMessage(r.RawJSON()))
}

// The properties WhatsappCallPermissionRequestTemplateLanguageCode,
// WhatsappCallPermissionRequestTemplateName are required.
type BatchCallWhatsappParams struct {
	WhatsappCallPermissionRequestTemplateLanguageCode string            `json:"whatsapp_call_permission_request_template_language_code,required"`
	WhatsappCallPermissionRequestTemplateName         string            `json:"whatsapp_call_permission_request_template_name,required"`
	WhatsappPhoneNumberID                             param.Opt[string] `json:"whatsapp_phone_number_id,omitzero"`
	paramObj
}

func (r BatchCallWhatsappParams) MarshalJSON() (data []byte, err error) {
	type shadow BatchCallWhatsappParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BatchCallWhatsappParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TelephonyProvider string

const (
	TelephonyProviderTwilio   TelephonyProvider = "twilio"
	TelephonyProviderSipTrunk TelephonyProvider = "sip_trunk"
)

// Detailed response model for a batch call including all recipients.
type ConvaiBatchCallingGetResponse struct {
	ID                string `json:"id,required"`
	AgentID           string `json:"agent_id,required"`
	AgentName         string `json:"agent_name,required"`
	CreatedAtUnix     int64  `json:"created_at_unix,required"`
	LastUpdatedAtUnix int64  `json:"last_updated_at_unix,required"`
	Name              string `json:"name,required"`
	PhoneNumberID     string `json:"phone_number_id,required"`
	// Any of "twilio", "sip_trunk".
	PhoneProvider     TelephonyProvider                        `json:"phone_provider,required"`
	Recipients        []ConvaiBatchCallingGetResponseRecipient `json:"recipients,required"`
	RetryCount        int64                                    `json:"retry_count,required"`
	ScheduledTimeUnix int64                                    `json:"scheduled_time_unix,required"`
	// Any of "pending", "in_progress", "completed", "failed", "cancelled".
	Status               BatchCallStatus             `json:"status,required"`
	Timezone             string                      `json:"timezone,required"`
	TotalCallsDispatched int64                       `json:"total_calls_dispatched,required"`
	TotalCallsFinished   int64                       `json:"total_calls_finished,required"`
	TotalCallsScheduled  int64                       `json:"total_calls_scheduled,required"`
	WhatsappParams       BatchCallWhatsappParamsResp `json:"whatsapp_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		AgentID              respjson.Field
		AgentName            respjson.Field
		CreatedAtUnix        respjson.Field
		LastUpdatedAtUnix    respjson.Field
		Name                 respjson.Field
		PhoneNumberID        respjson.Field
		PhoneProvider        respjson.Field
		Recipients           respjson.Field
		RetryCount           respjson.Field
		ScheduledTimeUnix    respjson.Field
		Status               respjson.Field
		Timezone             respjson.Field
		TotalCallsDispatched respjson.Field
		TotalCallsFinished   respjson.Field
		TotalCallsScheduled  respjson.Field
		WhatsappParams       respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiBatchCallingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiBatchCallingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiBatchCallingGetResponseRecipient struct {
	ID             string `json:"id,required"`
	ConversationID string `json:"conversation_id,required"`
	CreatedAtUnix  int64  `json:"created_at_unix,required"`
	// Any of "pending", "dispatched", "initiated", "in_progress", "completed",
	// "failed", "cancelled", "voicemail".
	Status                           string                                                                 `json:"status,required"`
	UpdatedAtUnix                    int64                                                                  `json:"updated_at_unix,required"`
	ConversationInitiationClientData ConvaiBatchCallingGetResponseRecipientConversationInitiationClientData `json:"conversation_initiation_client_data,nullable"`
	PhoneNumber                      string                                                                 `json:"phone_number,nullable"`
	WhatsappUserID                   string                                                                 `json:"whatsapp_user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                               respjson.Field
		ConversationID                   respjson.Field
		CreatedAtUnix                    respjson.Field
		Status                           respjson.Field
		UpdatedAtUnix                    respjson.Field
		ConversationInitiationClientData respjson.Field
		PhoneNumber                      respjson.Field
		WhatsappUserID                   respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiBatchCallingGetResponseRecipient) RawJSON() string { return r.JSON.raw }
func (r *ConvaiBatchCallingGetResponseRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiBatchCallingGetResponseRecipientConversationInitiationClientData struct {
	ConversationConfigOverride ConversationConfigClientOverride                                                                      `json:"conversation_config_override"`
	CustomLlmExtraBody         any                                                                                                   `json:"custom_llm_extra_body"`
	DynamicVariables           map[string]ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion `json:"dynamic_variables"`
	// Information about the source of conversation initiation
	SourceInfo ConversationInitiationSourceInfo `json:"source_info"`
	// ID of the end user participating in this conversation (for agent owner's user
	// identification)
	UserID string `json:"user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationConfigOverride respjson.Field
		CustomLlmExtraBody         respjson.Field
		DynamicVariables           respjson.Field
		SourceInfo                 respjson.Field
		UserID                     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiBatchCallingGetResponseRecipientConversationInitiationClientData) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiBatchCallingGetResponseRecipientConversationInitiationClientData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfFloat  respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiBatchCallingGetResponseRecipientConversationInitiationClientDataDynamicVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiBatchCallingListResponse struct {
	BatchCalls []BatchCallResponse `json:"batch_calls,required"`
	// Whether there are more batch calls to paginate through
	HasMore bool `json:"has_more"`
	// The next document, used to paginate through the batch calls
	NextDoc string `json:"next_doc,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchCalls  respjson.Field
		HasMore     respjson.Field
		NextDoc     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiBatchCallingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiBatchCallingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiBatchCallingGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiBatchCallingListParams struct {
	LastDoc param.Opt[string] `query:"last_doc,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiBatchCallingListParams]'s query parameters as
// `url.Values`.
func (r ConvaiBatchCallingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiBatchCallingDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiBatchCallingCancelParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiBatchCallingRetryParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiBatchCallingSubmitParams struct {
	AgentID            string                                    `json:"agent_id,required"`
	CallName           string                                    `json:"call_name,required"`
	Recipients         []ConvaiBatchCallingSubmitParamsRecipient `json:"recipients,omitzero,required"`
	AgentPhoneNumberID param.Opt[string]                         `json:"agent_phone_number_id,omitzero"`
	ScheduledTimeUnix  param.Opt[int64]                          `json:"scheduled_time_unix,omitzero"`
	Timezone           param.Opt[string]                         `json:"timezone,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey       param.Opt[string]       `header:"xi-api-key,omitzero" json:"-"`
	WhatsappParams BatchCallWhatsappParams `json:"whatsapp_params,omitzero"`
	paramObj
}

func (r ConvaiBatchCallingSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiBatchCallingSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiBatchCallingSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiBatchCallingSubmitParamsRecipient struct {
	ID                               param.Opt[string]                                 `json:"id,omitzero"`
	PhoneNumber                      param.Opt[string]                                 `json:"phone_number,omitzero"`
	WhatsappUserID                   param.Opt[string]                                 `json:"whatsapp_user_id,omitzero"`
	ConversationInitiationClientData ConversationInitiationClientDataRequestInputParam `json:"conversation_initiation_client_data,omitzero"`
	paramObj
}

func (r ConvaiBatchCallingSubmitParamsRecipient) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiBatchCallingSubmitParamsRecipient
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiBatchCallingSubmitParamsRecipient) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
