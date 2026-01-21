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
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiConversationService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiConversationService] method instead.
type ConvaiConversationService struct {
	Options      []option.RequestOption
	GetSignedURL ConvaiConversationGetSignedURLService
}

// NewConvaiConversationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiConversationService(opts ...option.RequestOption) (r ConvaiConversationService) {
	r = ConvaiConversationService{}
	r.Options = opts
	r.GetSignedURL = NewConvaiConversationGetSignedURLService(opts...)
	return
}

// Get the details of a particular conversation
func (r *ConvaiConversationService) Get(ctx context.Context, conversationID string, query ConvaiConversationGetParams, opts ...option.RequestOption) (res *ConvaiConversationGetResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all conversations of agents that user owns. With option to restrict to a
// specific agent.
func (r *ConvaiConversationService) List(ctx context.Context, params ConvaiConversationListParams, opts ...option.RequestOption) (res *ConvaiConversationListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/conversations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a particular conversation
func (r *ConvaiConversationService) Delete(ctx context.Context, conversationID string, body ConvaiConversationDeleteParams, opts ...option.RequestOption) (res *ConvaiConversationDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/conversations/%s", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the audio recording of a particular conversation
func (r *ConvaiConversationService) GetAudio(ctx context.Context, conversationID string, query ConvaiConversationGetAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/conversations/%s/audio", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a WebRTC session token for real-time communication.
func (r *ConvaiConversationService) GetToken(ctx context.Context, params ConvaiConversationGetTokenParams, opts ...option.RequestOption) (res *ConvaiConversationGetTokenResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/conversation/token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Send the feedback for the given conversation
func (r *ConvaiConversationService) SendFeedback(ctx context.Context, conversationID string, body ConvaiConversationSendFeedbackParams, opts ...option.RequestOption) (res *ConvaiConversationSendFeedbackResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if conversationID == "" {
		err = errors.New("missing required conversation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/conversations/%s/feedback", conversationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentDefinitionSource string

const (
	AgentDefinitionSourceCli      AgentDefinitionSource = "cli"
	AgentDefinitionSourceUi       AgentDefinitionSource = "ui"
	AgentDefinitionSourceAPI      AgentDefinitionSource = "api"
	AgentDefinitionSourceTemplate AgentDefinitionSource = "template"
	AgentDefinitionSourceUnknown  AgentDefinitionSource = "unknown"
)

type ConversationConfigClientOverride struct {
	// Agent specific configuration
	Agent ConversationConfigClientOverrideAgent `json:"agent,nullable"`
	// Configuration for conversational events
	Conversation ConversationConfigOverride `json:"conversation,nullable"`
	// Configuration for conversational text to speech
	Tts TtsConversationalConfigOverride `json:"tts,nullable"`
	// Configuration for turn detection
	Turn TurnConfigOverride `json:"turn,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent        respjson.Field
		Conversation respjson.Field
		Tts          respjson.Field
		Turn         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationConfigClientOverride) RawJSON() string { return r.JSON.raw }
func (r *ConversationConfigClientOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Agent specific configuration
type ConversationConfigClientOverrideAgent struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits
	// for the user to start the discussion.
	FirstMessage string `json:"first_message,nullable"`
	// Language of the agent - used for ASR and TTS
	Language string `json:"language,nullable"`
	// The prompt for the agent
	Prompt PromptAgentAPIModelOverride `json:"prompt,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FirstMessage respjson.Field
		Language     respjson.Field
		Prompt       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationConfigClientOverrideAgent) RawJSON() string { return r.JSON.raw }
func (r *ConversationConfigClientOverrideAgent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationConfigOverride struct {
	// If enabled audio will not be processed and only text will be used, use to avoid
	// audio pricing.
	TextOnly bool `json:"text_only,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TextOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationConfigOverride) RawJSON() string { return r.JSON.raw }
func (r *ConversationConfigOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationConfigOverride to a
// ConversationConfigOverrideParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationConfigOverrideParam.Overrides()
func (r ConversationConfigOverride) ToParam() ConversationConfigOverrideParam {
	return param.Override[ConversationConfigOverrideParam](json.RawMessage(r.RawJSON()))
}

type ConversationConfigOverrideParam struct {
	// If enabled audio will not be processed and only text will be used, use to avoid
	// audio pricing.
	TextOnly param.Opt[bool] `json:"text_only,omitzero"`
	paramObj
}

func (r ConversationConfigOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationConfigOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationConfigOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Enum representing the possible sources for conversation initiation.
type ConversationInitiationSource string

const (
	ConversationInitiationSourceUnknown            ConversationInitiationSource = "unknown"
	ConversationInitiationSourceAndroidSDK         ConversationInitiationSource = "android_sdk"
	ConversationInitiationSourceNodeJsSDK          ConversationInitiationSource = "node_js_sdk"
	ConversationInitiationSourceReactNativeSDK     ConversationInitiationSource = "react_native_sdk"
	ConversationInitiationSourceReactSDK           ConversationInitiationSource = "react_sdk"
	ConversationInitiationSourceJsSDK              ConversationInitiationSource = "js_sdk"
	ConversationInitiationSourcePythonSDK          ConversationInitiationSource = "python_sdk"
	ConversationInitiationSourceWidget             ConversationInitiationSource = "widget"
	ConversationInitiationSourceSipTrunk           ConversationInitiationSource = "sip_trunk"
	ConversationInitiationSourceTwilio             ConversationInitiationSource = "twilio"
	ConversationInitiationSourceGenesys            ConversationInitiationSource = "genesys"
	ConversationInitiationSourceSwiftSDK           ConversationInitiationSource = "swift_sdk"
	ConversationInitiationSourceWhatsapp           ConversationInitiationSource = "whatsapp"
	ConversationInitiationSourceFlutterSDK         ConversationInitiationSource = "flutter_sdk"
	ConversationInitiationSourceZendeskIntegration ConversationInitiationSource = "zendesk_integration"
)

// Information about the source of conversation initiation
type ConversationInitiationSourceInfo struct {
	// Enum representing the possible sources for conversation initiation.
	//
	// Any of "unknown", "android_sdk", "node_js_sdk", "react_native_sdk", "react_sdk",
	// "js_sdk", "python_sdk", "widget", "sip_trunk", "twilio", "genesys", "swift_sdk",
	// "whatsapp", "flutter_sdk", "zendesk_integration".
	Source ConversationInitiationSource `json:"source,nullable"`
	// The SDK version number
	Version string `json:"version,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Source      respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationInitiationSourceInfo) RawJSON() string { return r.JSON.raw }
func (r *ConversationInitiationSourceInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationInitiationSourceInfo to a
// ConversationInitiationSourceInfoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationInitiationSourceInfoParam.Overrides()
func (r ConversationInitiationSourceInfo) ToParam() ConversationInitiationSourceInfoParam {
	return param.Override[ConversationInitiationSourceInfoParam](json.RawMessage(r.RawJSON()))
}

// Information about the source of conversation initiation
type ConversationInitiationSourceInfoParam struct {
	// The SDK version number
	Version param.Opt[string] `json:"version,omitzero"`
	// Enum representing the possible sources for conversation initiation.
	//
	// Any of "unknown", "android_sdk", "node_js_sdk", "react_native_sdk", "react_sdk",
	// "js_sdk", "python_sdk", "widget", "sip_trunk", "twilio", "genesys", "swift_sdk",
	// "whatsapp", "flutter_sdk", "zendesk_integration".
	Source ConversationInitiationSource `json:"source,omitzero"`
	paramObj
}

func (r ConversationInitiationSourceInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationInitiationSourceInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationInitiationSourceInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluationSuccessResult string

const (
	EvaluationSuccessResultSuccess EvaluationSuccessResult = "success"
	EvaluationSuccessResultFailure EvaluationSuccessResult = "failure"
	EvaluationSuccessResultUnknown EvaluationSuccessResult = "unknown"
)

type FeatureStatusCommon struct {
	Enabled bool `json:"enabled"`
	Used    bool `json:"used"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Used        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureStatusCommon) RawJSON() string { return r.JSON.raw }
func (r *FeatureStatusCommon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PromptAgentAPIModelOverride struct {
	// The LLM to query with the prompt and the chat history. If using data residency,
	// the LLM must be supported in the data residency environment
	//
	// Any of "gpt-4o-mini", "gpt-4o", "gpt-4", "gpt-4-turbo", "gpt-4.1",
	// "gpt-4.1-mini", "gpt-4.1-nano", "gpt-5", "gpt-5.1", "gpt-5.2",
	// "gpt-5.2-chat-latest", "gpt-5-mini", "gpt-5-nano", "gpt-3.5-turbo",
	// "gemini-1.5-pro", "gemini-1.5-flash", "gemini-2.0-flash",
	// "gemini-2.0-flash-lite", "gemini-2.5-flash-lite", "gemini-2.5-flash",
	// "gemini-3-pro-preview", "gemini-3-flash-preview", "claude-sonnet-4-5",
	// "claude-sonnet-4", "claude-haiku-4-5", "claude-3-7-sonnet", "claude-3-5-sonnet",
	// "claude-3-5-sonnet-v1", "claude-3-haiku", "grok-beta", "custom-llm", "qwen3-4b",
	// "qwen3-30b-a3b", "gpt-oss-20b", "gpt-oss-120b", "glm-45-air-fp8",
	// "gemini-2.5-flash-preview-09-2025", "gemini-2.5-flash-lite-preview-09-2025",
	// "gemini-2.5-flash-preview-05-20", "gemini-2.5-flash-preview-04-17",
	// "gemini-2.5-flash-lite-preview-06-17", "gemini-2.0-flash-lite-001",
	// "gemini-2.0-flash-001", "gemini-1.5-flash-002", "gemini-1.5-flash-001",
	// "gemini-1.5-pro-002", "gemini-1.5-pro-001", "claude-sonnet-4@20250514",
	// "claude-sonnet-4-5@20250929", "claude-haiku-4-5@20251001",
	// "claude-3-7-sonnet@20250219", "claude-3-5-sonnet@20240620",
	// "claude-3-5-sonnet-v2@20241022", "claude-3-haiku@20240307", "gpt-5-2025-08-07",
	// "gpt-5.1-2025-11-13", "gpt-5.2-2025-12-11", "gpt-5-mini-2025-08-07",
	// "gpt-5-nano-2025-08-07", "gpt-4.1-2025-04-14", "gpt-4.1-mini-2025-04-14",
	// "gpt-4.1-nano-2025-04-14", "gpt-4o-mini-2024-07-18", "gpt-4o-2024-11-20",
	// "gpt-4o-2024-08-06", "gpt-4o-2024-05-13", "gpt-4-0613", "gpt-4-0314",
	// "gpt-4-turbo-2024-04-09", "gpt-3.5-turbo-0125", "gpt-3.5-turbo-1106",
	// "watt-tool-8b", "watt-tool-70b".
	Llm Llm `json:"llm,nullable"`
	// A list of Native MCP server ids to be used by the agent
	NativeMcpServerIDs []string `json:"native_mcp_server_ids,nullable"`
	// The prompt for the agent
	Prompt string `json:"prompt,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Llm                respjson.Field
		NativeMcpServerIDs respjson.Field
		Prompt             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptAgentAPIModelOverride) RawJSON() string { return r.JSON.raw }
func (r *PromptAgentAPIModelOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PromptAgentAPIModelOverride to a
// PromptAgentAPIModelOverrideParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PromptAgentAPIModelOverrideParam.Overrides()
func (r PromptAgentAPIModelOverride) ToParam() PromptAgentAPIModelOverrideParam {
	return param.Override[PromptAgentAPIModelOverrideParam](json.RawMessage(r.RawJSON()))
}

type PromptAgentAPIModelOverrideParam struct {
	// The prompt for the agent
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// A list of Native MCP server ids to be used by the agent
	NativeMcpServerIDs []string `json:"native_mcp_server_ids,omitzero"`
	// The LLM to query with the prompt and the chat history. If using data residency,
	// the LLM must be supported in the data residency environment
	//
	// Any of "gpt-4o-mini", "gpt-4o", "gpt-4", "gpt-4-turbo", "gpt-4.1",
	// "gpt-4.1-mini", "gpt-4.1-nano", "gpt-5", "gpt-5.1", "gpt-5.2",
	// "gpt-5.2-chat-latest", "gpt-5-mini", "gpt-5-nano", "gpt-3.5-turbo",
	// "gemini-1.5-pro", "gemini-1.5-flash", "gemini-2.0-flash",
	// "gemini-2.0-flash-lite", "gemini-2.5-flash-lite", "gemini-2.5-flash",
	// "gemini-3-pro-preview", "gemini-3-flash-preview", "claude-sonnet-4-5",
	// "claude-sonnet-4", "claude-haiku-4-5", "claude-3-7-sonnet", "claude-3-5-sonnet",
	// "claude-3-5-sonnet-v1", "claude-3-haiku", "grok-beta", "custom-llm", "qwen3-4b",
	// "qwen3-30b-a3b", "gpt-oss-20b", "gpt-oss-120b", "glm-45-air-fp8",
	// "gemini-2.5-flash-preview-09-2025", "gemini-2.5-flash-lite-preview-09-2025",
	// "gemini-2.5-flash-preview-05-20", "gemini-2.5-flash-preview-04-17",
	// "gemini-2.5-flash-lite-preview-06-17", "gemini-2.0-flash-lite-001",
	// "gemini-2.0-flash-001", "gemini-1.5-flash-002", "gemini-1.5-flash-001",
	// "gemini-1.5-pro-002", "gemini-1.5-pro-001", "claude-sonnet-4@20250514",
	// "claude-sonnet-4-5@20250929", "claude-haiku-4-5@20251001",
	// "claude-3-7-sonnet@20250219", "claude-3-5-sonnet@20240620",
	// "claude-3-5-sonnet-v2@20241022", "claude-3-haiku@20240307", "gpt-5-2025-08-07",
	// "gpt-5.1-2025-11-13", "gpt-5.2-2025-12-11", "gpt-5-mini-2025-08-07",
	// "gpt-5-nano-2025-08-07", "gpt-4.1-2025-04-14", "gpt-4.1-mini-2025-04-14",
	// "gpt-4.1-nano-2025-04-14", "gpt-4o-mini-2024-07-18", "gpt-4o-2024-11-20",
	// "gpt-4o-2024-08-06", "gpt-4o-2024-05-13", "gpt-4-0613", "gpt-4-0314",
	// "gpt-4-turbo-2024-04-09", "gpt-3.5-turbo-0125", "gpt-3.5-turbo-1106",
	// "watt-tool-8b", "watt-tool-70b".
	Llm Llm `json:"llm,omitzero"`
	paramObj
}

func (r PromptAgentAPIModelOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptAgentAPIModelOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptAgentAPIModelOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TtsConversationalConfigOverride struct {
	// The similarity boost for generated speech
	SimilarityBoost float64 `json:"similarity_boost,nullable"`
	// The speed of generated speech
	Speed float64 `json:"speed,nullable"`
	// The stability of generated speech
	Stability float64 `json:"stability,nullable"`
	// The voice ID to use for TTS
	VoiceID string `json:"voice_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SimilarityBoost respjson.Field
		Speed           respjson.Field
		Stability       respjson.Field
		VoiceID         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TtsConversationalConfigOverride) RawJSON() string { return r.JSON.raw }
func (r *TtsConversationalConfigOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TtsConversationalConfigOverride to a
// TtsConversationalConfigOverrideParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TtsConversationalConfigOverrideParam.Overrides()
func (r TtsConversationalConfigOverride) ToParam() TtsConversationalConfigOverrideParam {
	return param.Override[TtsConversationalConfigOverrideParam](json.RawMessage(r.RawJSON()))
}

type TtsConversationalConfigOverrideParam struct {
	// The similarity boost for generated speech
	SimilarityBoost param.Opt[float64] `json:"similarity_boost,omitzero"`
	// The speed of generated speech
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// The stability of generated speech
	Stability param.Opt[float64] `json:"stability,omitzero"`
	// The voice ID to use for TTS
	VoiceID param.Opt[string] `json:"voice_id,omitzero"`
	paramObj
}

func (r TtsConversationalConfigOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow TtsConversationalConfigOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TtsConversationalConfigOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TurnConfigOverride struct {
	// Configuration for soft timeout functionality. Provides immediate feedback during
	// longer LLM responses.
	SoftTimeoutConfig TurnConfigOverrideSoftTimeoutConfig `json:"soft_timeout_config,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SoftTimeoutConfig respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnConfigOverride) RawJSON() string { return r.JSON.raw }
func (r *TurnConfigOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TurnConfigOverride to a TurnConfigOverrideParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TurnConfigOverrideParam.Overrides()
func (r TurnConfigOverride) ToParam() TurnConfigOverrideParam {
	return param.Override[TurnConfigOverrideParam](json.RawMessage(r.RawJSON()))
}

// Configuration for soft timeout functionality. Provides immediate feedback during
// longer LLM responses.
type TurnConfigOverrideSoftTimeoutConfig struct {
	// Message to show when soft timeout is reached while waiting for LLM response
	Message string `json:"message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TurnConfigOverrideSoftTimeoutConfig) RawJSON() string { return r.JSON.raw }
func (r *TurnConfigOverrideSoftTimeoutConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TurnConfigOverrideParam struct {
	// Configuration for soft timeout functionality. Provides immediate feedback during
	// longer LLM responses.
	SoftTimeoutConfig TurnConfigOverrideSoftTimeoutConfigParam `json:"soft_timeout_config,omitzero"`
	paramObj
}

func (r TurnConfigOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow TurnConfigOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TurnConfigOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for soft timeout functionality. Provides immediate feedback during
// longer LLM responses.
type TurnConfigOverrideSoftTimeoutConfigParam struct {
	// Message to show when soft timeout is reached while waiting for LLM response
	Message param.Opt[string] `json:"message,omitzero"`
	paramObj
}

func (r TurnConfigOverrideSoftTimeoutConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TurnConfigOverrideSoftTimeoutConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TurnConfigOverrideSoftTimeoutConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserFeedbackScore string

const (
	UserFeedbackScoreLike    UserFeedbackScore = "like"
	UserFeedbackScoreDislike UserFeedbackScore = "dislike"
)

type ConvaiConversationGetResponse struct {
	AgentID          string                                `json:"agent_id,required"`
	ConversationID   string                                `json:"conversation_id,required"`
	HasAudio         bool                                  `json:"has_audio,required"`
	HasResponseAudio bool                                  `json:"has_response_audio,required"`
	HasUserAudio     bool                                  `json:"has_user_audio,required"`
	Metadata         ConvaiConversationGetResponseMetadata `json:"metadata,required"`
	// Any of "initiated", "in-progress", "processing", "done", "failed".
	Status                           ConvaiConversationGetResponseStatus                           `json:"status,required"`
	Transcript                       []ConversationHistoryTranscriptCommonModelOutput              `json:"transcript,required"`
	Analysis                         ConversationHistoryAnalysisCommonModel                        `json:"analysis,nullable"`
	BranchID                         string                                                        `json:"branch_id,nullable"`
	ConversationInitiationClientData ConvaiConversationGetResponseConversationInitiationClientData `json:"conversation_initiation_client_data"`
	UserID                           string                                                        `json:"user_id,nullable"`
	// The ID of the agent version used for this conversation
	VersionID string `json:"version_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID                          respjson.Field
		ConversationID                   respjson.Field
		HasAudio                         respjson.Field
		HasResponseAudio                 respjson.Field
		HasUserAudio                     respjson.Field
		Metadata                         respjson.Field
		Status                           respjson.Field
		Transcript                       respjson.Field
		Analysis                         respjson.Field
		BranchID                         respjson.Field
		ConversationInitiationClientData respjson.Field
		UserID                           respjson.Field
		VersionID                        respjson.Field
		ExtraFields                      map[string]respjson.Field
		raw                              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadata struct {
	CallDurationSecs     int64 `json:"call_duration_secs,required"`
	StartTimeUnixSecs    int64 `json:"start_time_unix_secs,required"`
	AcceptedTimeUnixSecs int64 `json:"accepted_time_unix_secs,nullable"`
	// Any of "cli", "ui", "api", "template", "unknown".
	AgentCreatedFrom AgentDefinitionSource `json:"agent_created_from"`
	// Any of "cli", "ui", "api", "template", "unknown".
	AgentLastUpdatedFrom AgentDefinitionSource `json:"agent_last_updated_from"`
	// Metadata for async conversation delivery (Zendesk, Slack, etc.).
	AsyncMetadata ConvaiConversationGetResponseMetadataAsyncMetadata `json:"async_metadata,nullable"`
	// Any of "invalid", "public", "authorization_header", "signed_url",
	// "shareable_link", "livekit_token", "livekit_token_website", "genesys_api_key",
	// "whatsapp".
	AuthorizationMethod string                                         `json:"authorization_method"`
	BatchCall           ConvaiConversationGetResponseMetadataBatchCall `json:"batch_call,nullable"`
	Charging            ConvaiConversationGetResponseMetadataCharging  `json:"charging"`
	// Enum representing the possible sources for conversation initiation.
	//
	// Any of "unknown", "android_sdk", "node_js_sdk", "react_native_sdk", "react_sdk",
	// "js_sdk", "python_sdk", "widget", "sip_trunk", "twilio", "genesys", "swift_sdk",
	// "whatsapp", "flutter_sdk", "zendesk_integration".
	ConversationInitiationSource        ConversationInitiationSource                          `json:"conversation_initiation_source"`
	ConversationInitiationSourceVersion string                                                `json:"conversation_initiation_source_version,nullable"`
	Cost                                int64                                                 `json:"cost,nullable"`
	DeletionSettings                    ConvaiConversationGetResponseMetadataDeletionSettings `json:"deletion_settings"`
	ElevenAssistant                     ConvaiConversationGetResponseMetadataElevenAssistant  `json:"eleven_assistant"`
	Error                               ConvaiConversationGetResponseMetadataError            `json:"error,nullable"`
	FeaturesUsage                       ConvaiConversationGetResponseMetadataFeaturesUsage    `json:"features_usage"`
	Feedback                            ConvaiConversationGetResponseMetadataFeedback         `json:"feedback"`
	InitiatorID                         string                                                `json:"initiator_id,nullable"`
	MainLanguage                        string                                                `json:"main_language,nullable"`
	PhoneCall                           ConvaiConversationGetResponseMetadataPhoneCallUnion   `json:"phone_call,nullable"`
	RagUsage                            ConvaiConversationGetResponseMetadataRagUsage         `json:"rag_usage,nullable"`
	TerminationReason                   string                                                `json:"termination_reason"`
	TextOnly                            bool                                                  `json:"text_only"`
	Timezone                            string                                                `json:"timezone,nullable"`
	Warnings                            []string                                              `json:"warnings"`
	Whatsapp                            ConvaiConversationGetResponseMetadataWhatsapp         `json:"whatsapp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallDurationSecs                    respjson.Field
		StartTimeUnixSecs                   respjson.Field
		AcceptedTimeUnixSecs                respjson.Field
		AgentCreatedFrom                    respjson.Field
		AgentLastUpdatedFrom                respjson.Field
		AsyncMetadata                       respjson.Field
		AuthorizationMethod                 respjson.Field
		BatchCall                           respjson.Field
		Charging                            respjson.Field
		ConversationInitiationSource        respjson.Field
		ConversationInitiationSourceVersion respjson.Field
		Cost                                respjson.Field
		DeletionSettings                    respjson.Field
		ElevenAssistant                     respjson.Field
		Error                               respjson.Field
		FeaturesUsage                       respjson.Field
		Feedback                            respjson.Field
		InitiatorID                         respjson.Field
		MainLanguage                        respjson.Field
		PhoneCall                           respjson.Field
		RagUsage                            respjson.Field
		TerminationReason                   respjson.Field
		TextOnly                            respjson.Field
		Timezone                            respjson.Field
		Warnings                            respjson.Field
		Whatsapp                            respjson.Field
		ExtraFields                         map[string]respjson.Field
		raw                                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for async conversation delivery (Zendesk, Slack, etc.).
type ConvaiConversationGetResponseMetadataAsyncMetadata struct {
	// Any of "pending", "success", "failed".
	DeliveryStatus     string `json:"delivery_status,required"`
	DeliveryTimestamp  int64  `json:"delivery_timestamp,required"`
	ExternalID         string `json:"external_id,required"`
	ExternalSystem     string `json:"external_system,required"`
	DeliveryError      string `json:"delivery_error,nullable"`
	LastRetryTimestamp int64  `json:"last_retry_timestamp,nullable"`
	RetryCount         int64  `json:"retry_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeliveryStatus     respjson.Field
		DeliveryTimestamp  respjson.Field
		ExternalID         respjson.Field
		ExternalSystem     respjson.Field
		DeliveryError      respjson.Field
		LastRetryTimestamp respjson.Field
		RetryCount         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataAsyncMetadata) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataAsyncMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataBatchCall struct {
	BatchCallID          string `json:"batch_call_id,required"`
	BatchCallRecipientID string `json:"batch_call_recipient_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchCallID          respjson.Field
		BatchCallRecipientID respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataBatchCall) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataBatchCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataCharging struct {
	CallCharge             int64                                                 `json:"call_charge,nullable"`
	DevDiscount            bool                                                  `json:"dev_discount"`
	FreeLlmDollarsConsumed float64                                               `json:"free_llm_dollars_consumed"`
	FreeMinutesConsumed    float64                                               `json:"free_minutes_consumed"`
	IsBurst                bool                                                  `json:"is_burst"`
	LlmCharge              int64                                                 `json:"llm_charge,nullable"`
	LlmPrice               float64                                               `json:"llm_price,nullable"`
	LlmUsage               ConvaiConversationGetResponseMetadataChargingLlmUsage `json:"llm_usage"`
	Tier                   string                                                `json:"tier,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallCharge             respjson.Field
		DevDiscount            respjson.Field
		FreeLlmDollarsConsumed respjson.Field
		FreeMinutesConsumed    respjson.Field
		IsBurst                respjson.Field
		LlmCharge              respjson.Field
		LlmPrice               respjson.Field
		LlmUsage               respjson.Field
		Tier                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataCharging) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataCharging) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataChargingLlmUsage struct {
	InitiatedGeneration    LlmUsageOutput `json:"initiated_generation"`
	IrreversibleGeneration LlmUsageOutput `json:"irreversible_generation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InitiatedGeneration    respjson.Field
		IrreversibleGeneration respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataChargingLlmUsage) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataChargingLlmUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataDeletionSettings struct {
	DeleteAudio                     bool  `json:"delete_audio"`
	DeleteTranscriptAndPii          bool  `json:"delete_transcript_and_pii"`
	DeletedAudioAtTimeUnixSecs      int64 `json:"deleted_audio_at_time_unix_secs,nullable"`
	DeletedLogsAtTimeUnixSecs       int64 `json:"deleted_logs_at_time_unix_secs,nullable"`
	DeletedTranscriptAtTimeUnixSecs int64 `json:"deleted_transcript_at_time_unix_secs,nullable"`
	DeletionTimeUnixSecs            int64 `json:"deletion_time_unix_secs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DeleteAudio                     respjson.Field
		DeleteTranscriptAndPii          respjson.Field
		DeletedAudioAtTimeUnixSecs      respjson.Field
		DeletedLogsAtTimeUnixSecs       respjson.Field
		DeletedTranscriptAtTimeUnixSecs respjson.Field
		DeletionTimeUnixSecs            respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataDeletionSettings) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataDeletionSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataElevenAssistant struct {
	IsElevenAssistant bool `json:"is_eleven_assistant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsElevenAssistant respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataElevenAssistant) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataElevenAssistant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataError struct {
	Code   int64  `json:"code,required"`
	Reason string `json:"reason,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Reason      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataError) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataFeaturesUsage struct {
	AgentTesting               ConvaiConversationGetResponseMetadataFeaturesUsageAgentTesting `json:"agent_testing"`
	DtmfTones                  FeatureStatusCommon                                            `json:"dtmf_tones"`
	ExternalMcpServers         FeatureStatusCommon                                            `json:"external_mcp_servers"`
	IsLivekit                  bool                                                           `json:"is_livekit"`
	LanguageDetection          FeatureStatusCommon                                            `json:"language_detection"`
	Multivoice                 FeatureStatusCommon                                            `json:"multivoice"`
	PiiZrmAgent                bool                                                           `json:"pii_zrm_agent"`
	PiiZrmWorkspace            bool                                                           `json:"pii_zrm_workspace"`
	ToolDynamicVariableUpdates FeatureStatusCommon                                            `json:"tool_dynamic_variable_updates"`
	TransferToAgent            FeatureStatusCommon                                            `json:"transfer_to_agent"`
	TransferToNumber           FeatureStatusCommon                                            `json:"transfer_to_number"`
	Versioning                 FeatureStatusCommon                                            `json:"versioning"`
	VoicemailDetection         FeatureStatusCommon                                            `json:"voicemail_detection"`
	Workflow                   ConvaiConversationGetResponseMetadataFeaturesUsageWorkflow     `json:"workflow"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentTesting               respjson.Field
		DtmfTones                  respjson.Field
		ExternalMcpServers         respjson.Field
		IsLivekit                  respjson.Field
		LanguageDetection          respjson.Field
		Multivoice                 respjson.Field
		PiiZrmAgent                respjson.Field
		PiiZrmWorkspace            respjson.Field
		ToolDynamicVariableUpdates respjson.Field
		TransferToAgent            respjson.Field
		TransferToNumber           respjson.Field
		Versioning                 respjson.Field
		VoicemailDetection         respjson.Field
		Workflow                   respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataFeaturesUsage) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataFeaturesUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataFeaturesUsageAgentTesting struct {
	Enabled                       bool `json:"enabled"`
	TestsRanAfterLastModification bool `json:"tests_ran_after_last_modification"`
	TestsRanInLast7Days           bool `json:"tests_ran_in_last_7_days"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                       respjson.Field
		TestsRanAfterLastModification respjson.Field
		TestsRanInLast7Days           respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataFeaturesUsageAgentTesting) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiConversationGetResponseMetadataFeaturesUsageAgentTesting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataFeaturesUsageWorkflow struct {
	Enabled             bool                `json:"enabled"`
	EndNode             FeatureStatusCommon `json:"end_node"`
	PhoneNumberNode     FeatureStatusCommon `json:"phone_number_node"`
	StandaloneAgentNode FeatureStatusCommon `json:"standalone_agent_node"`
	ToolNode            FeatureStatusCommon `json:"tool_node"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled             respjson.Field
		EndNode             respjson.Field
		PhoneNumberNode     respjson.Field
		StandaloneAgentNode respjson.Field
		ToolNode            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataFeaturesUsageWorkflow) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiConversationGetResponseMetadataFeaturesUsageWorkflow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataFeedback struct {
	Comment  string `json:"comment,nullable"`
	Dislikes int64  `json:"dislikes"`
	Likes    int64  `json:"likes"`
	// Any of "like", "dislike".
	OverallScore UserFeedbackScore `json:"overall_score,nullable"`
	Rating       int64             `json:"rating,nullable"`
	// Any of "thumbs", "rating".
	Type string `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comment      respjson.Field
		Dislikes     respjson.Field
		Likes        respjson.Field
		OverallScore respjson.Field
		Rating       respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataFeedback) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiConversationGetResponseMetadataPhoneCallUnion contains all possible
// properties and values from
// [ConvaiConversationGetResponseMetadataPhoneCallTwilio],
// [ConvaiConversationGetResponseMetadataPhoneCallSipTrunking].
//
// Use the [ConvaiConversationGetResponseMetadataPhoneCallUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiConversationGetResponseMetadataPhoneCallUnion struct {
	AgentNumber    string `json:"agent_number"`
	CallSid        string `json:"call_sid"`
	Direction      string `json:"direction"`
	ExternalNumber string `json:"external_number"`
	PhoneNumberID  string `json:"phone_number_id"`
	// This field is from variant
	// [ConvaiConversationGetResponseMetadataPhoneCallTwilio].
	StreamSid string `json:"stream_sid"`
	// Any of "twilio", "sip_trunking".
	Type string `json:"type"`
	JSON struct {
		AgentNumber    respjson.Field
		CallSid        respjson.Field
		Direction      respjson.Field
		ExternalNumber respjson.Field
		PhoneNumberID  respjson.Field
		StreamSid      respjson.Field
		Type           respjson.Field
		raw            string
	} `json:"-"`
}

// anyConvaiConversationGetResponseMetadataPhoneCall is implemented by each variant
// of [ConvaiConversationGetResponseMetadataPhoneCallUnion] to add type safety for
// the return type of [ConvaiConversationGetResponseMetadataPhoneCallUnion.AsAny]
type anyConvaiConversationGetResponseMetadataPhoneCall interface {
	implConvaiConversationGetResponseMetadataPhoneCallUnion()
}

func (ConvaiConversationGetResponseMetadataPhoneCallTwilio) implConvaiConversationGetResponseMetadataPhoneCallUnion() {
}
func (ConvaiConversationGetResponseMetadataPhoneCallSipTrunking) implConvaiConversationGetResponseMetadataPhoneCallUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiConversationGetResponseMetadataPhoneCallUnion.AsAny().(type) {
//	case elevenlabs.ConvaiConversationGetResponseMetadataPhoneCallTwilio:
//	case elevenlabs.ConvaiConversationGetResponseMetadataPhoneCallSipTrunking:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiConversationGetResponseMetadataPhoneCallUnion) AsAny() anyConvaiConversationGetResponseMetadataPhoneCall {
	switch u.Type {
	case "twilio":
		return u.AsTwilio()
	case "sip_trunking":
		return u.AsSipTrunking()
	}
	return nil
}

func (u ConvaiConversationGetResponseMetadataPhoneCallUnion) AsTwilio() (v ConvaiConversationGetResponseMetadataPhoneCallTwilio) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiConversationGetResponseMetadataPhoneCallUnion) AsSipTrunking() (v ConvaiConversationGetResponseMetadataPhoneCallSipTrunking) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiConversationGetResponseMetadataPhoneCallUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiConversationGetResponseMetadataPhoneCallUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataPhoneCallTwilio struct {
	AgentNumber string `json:"agent_number,required"`
	CallSid     string `json:"call_sid,required"`
	// Any of "inbound", "outbound".
	Direction      string          `json:"direction,required"`
	ExternalNumber string          `json:"external_number,required"`
	PhoneNumberID  string          `json:"phone_number_id,required"`
	StreamSid      string          `json:"stream_sid,required"`
	Type           constant.Twilio `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentNumber    respjson.Field
		CallSid        respjson.Field
		Direction      respjson.Field
		ExternalNumber respjson.Field
		PhoneNumberID  respjson.Field
		StreamSid      respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataPhoneCallTwilio) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataPhoneCallTwilio) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataPhoneCallSipTrunking struct {
	AgentNumber string `json:"agent_number,required"`
	CallSid     string `json:"call_sid,required"`
	// Any of "inbound", "outbound".
	Direction      string               `json:"direction,required"`
	ExternalNumber string               `json:"external_number,required"`
	PhoneNumberID  string               `json:"phone_number_id,required"`
	Type           constant.SipTrunking `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentNumber    respjson.Field
		CallSid        respjson.Field
		Direction      respjson.Field
		ExternalNumber respjson.Field
		PhoneNumberID  respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataPhoneCallSipTrunking) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiConversationGetResponseMetadataPhoneCallSipTrunking) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataRagUsage struct {
	EmbeddingModel string `json:"embedding_model,required"`
	UsageCount     int64  `json:"usage_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EmbeddingModel respjson.Field
		UsageCount     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataRagUsage) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataRagUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseMetadataWhatsapp struct {
	WhatsappUserID string `json:"whatsapp_user_id,required"`
	// Any of "inbound", "outbound", "unknown".
	Direction             string `json:"direction"`
	WhatsappPhoneNumberID string `json:"whatsapp_phone_number_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		WhatsappUserID        respjson.Field
		Direction             respjson.Field
		WhatsappPhoneNumberID respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetResponseMetadataWhatsapp) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetResponseMetadataWhatsapp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationGetResponseStatus string

const (
	ConvaiConversationGetResponseStatusInitiated  ConvaiConversationGetResponseStatus = "initiated"
	ConvaiConversationGetResponseStatusInProgress ConvaiConversationGetResponseStatus = "in-progress"
	ConvaiConversationGetResponseStatusProcessing ConvaiConversationGetResponseStatus = "processing"
	ConvaiConversationGetResponseStatusDone       ConvaiConversationGetResponseStatus = "done"
	ConvaiConversationGetResponseStatusFailed     ConvaiConversationGetResponseStatus = "failed"
)

type ConvaiConversationGetResponseConversationInitiationClientData struct {
	ConversationConfigOverride ConversationConfigClientOverride                                                             `json:"conversation_config_override"`
	CustomLlmExtraBody         any                                                                                          `json:"custom_llm_extra_body"`
	DynamicVariables           map[string]ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion `json:"dynamic_variables"`
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
func (r ConvaiConversationGetResponseConversationInitiationClientData) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiConversationGetResponseConversationInitiationClientData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion struct {
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

func (u ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiConversationGetResponseConversationInitiationClientDataDynamicVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationListResponse struct {
	Conversations []ConvaiConversationListResponseConversation `json:"conversations,required"`
	HasMore       bool                                         `json:"has_more,required"`
	NextCursor    string                                       `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conversations respjson.Field
		HasMore       respjson.Field
		NextCursor    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationListResponseConversation struct {
	AgentID          string `json:"agent_id,required"`
	CallDurationSecs int64  `json:"call_duration_secs,required"`
	// Any of "success", "failure", "unknown".
	CallSuccessful    EvaluationSuccessResult `json:"call_successful,required"`
	ConversationID    string                  `json:"conversation_id,required"`
	MessageCount      int64                   `json:"message_count,required"`
	StartTimeUnixSecs int64                   `json:"start_time_unix_secs,required"`
	// Any of "initiated", "in-progress", "processing", "done", "failed".
	Status           string `json:"status,required"`
	AgentName        string `json:"agent_name,nullable"`
	BranchID         string `json:"branch_id,nullable"`
	CallSummaryTitle string `json:"call_summary_title,nullable"`
	// Any of "inbound", "outbound".
	Direction         string  `json:"direction,nullable"`
	Rating            float64 `json:"rating,nullable"`
	TranscriptSummary string  `json:"transcript_summary,nullable"`
	VersionID         string  `json:"version_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID           respjson.Field
		CallDurationSecs  respjson.Field
		CallSuccessful    respjson.Field
		ConversationID    respjson.Field
		MessageCount      respjson.Field
		StartTimeUnixSecs respjson.Field
		Status            respjson.Field
		AgentName         respjson.Field
		BranchID          respjson.Field
		CallSummaryTitle  respjson.Field
		Direction         respjson.Field
		Rating            respjson.Field
		TranscriptSummary respjson.Field
		VersionID         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationListResponseConversation) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationListResponseConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationDeleteResponse = any

type ConvaiConversationGetTokenResponse struct {
	Token string `json:"token,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiConversationGetTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiConversationGetTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiConversationSendFeedbackResponse = any

type ConvaiConversationGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiConversationListParams struct {
	// The id of the agent you're taking the action on.
	AgentID param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// Filter conversations by branch ID.
	BranchID param.Opt[string] `query:"branch_id,omitzero" json:"-"`
	// Maximum call duration in seconds.
	CallDurationMaxSecs param.Opt[int64] `query:"call_duration_max_secs,omitzero" json:"-"`
	// Minimum call duration in seconds.
	CallDurationMinSecs param.Opt[int64] `query:"call_duration_min_secs,omitzero" json:"-"`
	// Unix timestamp (in seconds) to filter conversations after to this start date.
	CallStartAfterUnix param.Opt[int64] `query:"call_start_after_unix,omitzero" json:"-"`
	// Unix timestamp (in seconds) to filter conversations up to this start date.
	CallStartBeforeUnix param.Opt[int64] `query:"call_start_before_unix,omitzero" json:"-"`
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter conversations with user feedback comments.
	HasFeedbackComment param.Opt[bool] `query:"has_feedback_comment,omitzero" json:"-"`
	// Maximum overall rating (1-5).
	RatingMax param.Opt[int64] `query:"rating_max,omitzero" json:"-"`
	// Minimum overall rating (1-5).
	RatingMin param.Opt[int64] `query:"rating_min,omitzero" json:"-"`
	// Full-text or fuzzy search over transcript messages
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Filter conversations by the user ID who initiated them.
	UserID param.Opt[string] `query:"user_id,omitzero" json:"-"`
	// How many conversations to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Data collection filters. Repeat param. Format: id:op:value where op is one of
	// eq|neq|gt|gte|lt|lte|in|exists|missing. For in, pipe-delimit values.
	DataCollectionParams []string `query:"data_collection_params,omitzero" json:"-"`
	// Evaluation filters. Repeat param. Format: criteria_id:result. Example:
	// eval=value_framing:success
	EvaluationParams []string `query:"evaluation_params,omitzero" json:"-"`
	// Filter conversations by detected main language (language code).
	MainLanguages []string `query:"main_languages,omitzero" json:"-"`
	// Filter conversations by tool names used during the call.
	ToolNames []string `query:"tool_names,omitzero" json:"-"`
	// The result of the success evaluation
	//
	// Any of "success", "failure", "unknown".
	CallSuccessful EvaluationSuccessResult `query:"call_successful,omitzero" json:"-"`
	// Enum representing the possible sources for conversation initiation.
	//
	// Any of "unknown", "android_sdk", "node_js_sdk", "react_native_sdk", "react_sdk",
	// "js_sdk", "python_sdk", "widget", "sip_trunk", "twilio", "genesys", "swift_sdk",
	// "whatsapp", "flutter_sdk", "zendesk_integration".
	ConversationInitiationSource ConversationInitiationSource `query:"conversation_initiation_source,omitzero" json:"-"`
	// Whether to include transcript summaries in the response.
	//
	// Any of "exclude", "include".
	SummaryMode ConvaiConversationListParamsSummaryMode `query:"summary_mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiConversationListParams]'s query parameters as
// `url.Values`.
func (r ConvaiConversationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Whether to include transcript summaries in the response.
type ConvaiConversationListParamsSummaryMode string

const (
	ConvaiConversationListParamsSummaryModeExclude ConvaiConversationListParamsSummaryMode = "exclude"
	ConvaiConversationListParamsSummaryModeInclude ConvaiConversationListParamsSummaryMode = "include"
)

type ConvaiConversationDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiConversationGetAudioParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiConversationGetTokenParams struct {
	// The id of the agent you're taking the action on.
	AgentID string `query:"agent_id,required" json:"-"`
	// The ID of the branch to use
	BranchID param.Opt[string] `query:"branch_id,omitzero" json:"-"`
	// Optional custom participant name. If not provided, user ID will be used
	ParticipantName param.Opt[string] `query:"participant_name,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiConversationGetTokenParams]'s query parameters as
// `url.Values`.
func (r ConvaiConversationGetTokenParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiConversationSendFeedbackParams struct {
	// Either 'like' or 'dislike' to indicate the feedback for the conversation.
	//
	// Any of "like", "dislike".
	Feedback UserFeedbackScore `json:"feedback,omitzero"`
	paramObj
}

func (r ConvaiConversationSendFeedbackParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiConversationSendFeedbackParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiConversationSendFeedbackParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
