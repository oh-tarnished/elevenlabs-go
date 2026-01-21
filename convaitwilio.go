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

// ConvaiTwilioService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiTwilioService] method instead.
type ConvaiTwilioService struct {
	Options []option.RequestOption
}

// NewConvaiTwilioService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiTwilioService(opts ...option.RequestOption) (r ConvaiTwilioService) {
	r = ConvaiTwilioService{}
	r.Options = opts
	return
}

// Handle an outbound call via Twilio
func (r *ConvaiTwilioService) OutboundCall(ctx context.Context, params ConvaiTwilioOutboundCallParams, opts ...option.RequestOption) (res *ConvaiTwilioOutboundCallResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/twilio/outbound-call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Register a Twilio call and return TwiML to connect the call
func (r *ConvaiTwilioService) RegisterCall(ctx context.Context, params ConvaiTwilioRegisterCallParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	path := "v1/convai/twilio/register-call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ConversationInitiationClientDataRequestInputParam struct {
	// ID of the end user participating in this conversation (for agent owner's user
	// identification)
	UserID                     param.Opt[string]                                                                `json:"user_id,omitzero"`
	ConversationConfigOverride ConversationInitiationClientDataRequestInputConversationConfigOverrideParam      `json:"conversation_config_override,omitzero"`
	CustomLlmExtraBody         any                                                                              `json:"custom_llm_extra_body,omitzero"`
	DynamicVariables           map[string]ConversationInitiationClientDataRequestInputDynamicVariableUnionParam `json:"dynamic_variables,omitzero"`
	// Information about the source of conversation initiation
	SourceInfo ConversationInitiationSourceInfoParam `json:"source_info,omitzero"`
	paramObj
}

func (r ConversationInitiationClientDataRequestInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationInitiationClientDataRequestInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationInitiationClientDataRequestInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationInitiationClientDataRequestInputConversationConfigOverrideParam struct {
	// Agent specific configuration
	Agent ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam `json:"agent,omitzero"`
	// Configuration for conversational events
	Conversation ConversationConfigOverrideParam `json:"conversation,omitzero"`
	// Configuration for conversational text to speech
	Tts TtsConversationalConfigOverrideParam `json:"tts,omitzero"`
	// Configuration for turn detection
	Turn TurnConfigOverrideParam `json:"turn,omitzero"`
	paramObj
}

func (r ConversationInitiationClientDataRequestInputConversationConfigOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationInitiationClientDataRequestInputConversationConfigOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationInitiationClientDataRequestInputConversationConfigOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Agent specific configuration
type ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits
	// for the user to start the discussion.
	FirstMessage param.Opt[string] `json:"first_message,omitzero"`
	// Language of the agent - used for ASR and TTS
	Language param.Opt[string] `json:"language,omitzero"`
	// The prompt for the agent
	Prompt PromptAgentAPIModelOverrideParam `json:"prompt,omitzero"`
	paramObj
}

func (r ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationInitiationClientDataRequestInputDynamicVariableUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationInitiationClientDataRequestInputDynamicVariableUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ConversationInitiationClientDataRequestInputDynamicVariableUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationInitiationClientDataRequestInputDynamicVariableUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ConvaiTwilioOutboundCallResponse struct {
	CallSid        string `json:"callSid,required"`
	ConversationID string `json:"conversation_id,required"`
	Message        string `json:"message,required"`
	Success        bool   `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSid        respjson.Field
		ConversationID respjson.Field
		Message        respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiTwilioOutboundCallResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiTwilioOutboundCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTwilioOutboundCallParams struct {
	AgentID            string `json:"agent_id,required"`
	AgentPhoneNumberID string `json:"agent_phone_number_id,required"`
	ToNumber           string `json:"to_number,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey                         param.Opt[string]                                 `header:"xi-api-key,omitzero" json:"-"`
	ConversationInitiationClientData ConversationInitiationClientDataRequestInputParam `json:"conversation_initiation_client_data,omitzero"`
	paramObj
}

func (r ConvaiTwilioOutboundCallParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiTwilioOutboundCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiTwilioOutboundCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTwilioRegisterCallParams struct {
	AgentID    string `json:"agent_id,required"`
	FromNumber string `json:"from_number,required"`
	ToNumber   string `json:"to_number,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey                         param.Opt[string]                                 `header:"xi-api-key,omitzero" json:"-"`
	ConversationInitiationClientData ConversationInitiationClientDataRequestInputParam `json:"conversation_initiation_client_data,omitzero"`
	// Any of "inbound", "outbound".
	Direction ConvaiTwilioRegisterCallParamsDirection `json:"direction,omitzero"`
	paramObj
}

func (r ConvaiTwilioRegisterCallParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiTwilioRegisterCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiTwilioRegisterCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTwilioRegisterCallParamsDirection string

const (
	ConvaiTwilioRegisterCallParamsDirectionInbound  ConvaiTwilioRegisterCallParamsDirection = "inbound"
	ConvaiTwilioRegisterCallParamsDirectionOutbound ConvaiTwilioRegisterCallParamsDirection = "outbound"
)
