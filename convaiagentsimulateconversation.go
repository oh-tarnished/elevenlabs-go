// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiAgentSimulateConversationService contains methods and other services that
// help with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiAgentSimulateConversationService] method instead.
type ConvaiAgentSimulateConversationService struct {
	Options []option.RequestOption
}

// NewConvaiAgentSimulateConversationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiAgentSimulateConversationService(opts ...option.RequestOption) (r ConvaiAgentSimulateConversationService) {
	r = ConvaiAgentSimulateConversationService{}
	r.Options = opts
	return
}

// Run a conversation between the agent and a simulated user.
func (r *ConvaiAgentSimulateConversationService) SimulateConversation(ctx context.Context, agentID string, params ConvaiAgentSimulateConversationSimulateConversationParams, opts ...option.RequestOption) (res *ConvaiAgentSimulateConversationSimulateConversationResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agents/%s/simulate-conversation", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Run a conversation between the agent and a simulated user and stream back the
// response. Response is streamed back as partial lists of messages that should be
// concatenated and once the conversation has complete a single final message with
// the conversation analysis will be sent.
func (r *ConvaiAgentSimulateConversationService) Stream(ctx context.Context, agentID string, params ConvaiAgentSimulateConversationStreamParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agents/%s/simulate-conversation/stream", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type AgentMetadata struct {
	AgentID        string `json:"agent_id,required"`
	BranchID       string `json:"branch_id,nullable"`
	WorkflowNodeID string `json:"workflow_node_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID        respjson.Field
		BranchID       respjson.Field
		WorkflowNodeID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentMetadata) RawJSON() string { return r.JSON.raw }
func (r *AgentMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentMetadata to a AgentMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentMetadataParam.Overrides()
func (r AgentMetadata) ToParam() AgentMetadataParam {
	return param.Override[AgentMetadataParam](json.RawMessage(r.RawJSON()))
}

// The property AgentID is required.
type AgentMetadataParam struct {
	AgentID        string            `json:"agent_id,required"`
	BranchID       param.Opt[string] `json:"branch_id,omitzero"`
	WorkflowNodeID param.Opt[string] `json:"workflow_node_id,omitzero"`
	paramObj
}

func (r AgentMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryAnalysisCommonModel struct {
	// Any of "success", "failure", "unknown".
	CallSuccessful                EvaluationSuccessResult                                                   `json:"call_successful,required"`
	TranscriptSummary             string                                                                    `json:"transcript_summary,required"`
	CallSummaryTitle              string                                                                    `json:"call_summary_title,nullable"`
	DataCollectionResults         map[string]ConversationHistoryAnalysisCommonModelDataCollectionResult     `json:"data_collection_results"`
	DataCollectionResultsList     []ConversationHistoryAnalysisCommonModelDataCollectionResultsList         `json:"data_collection_results_list"`
	EvaluationCriteriaResults     map[string]ConversationHistoryAnalysisCommonModelEvaluationCriteriaResult `json:"evaluation_criteria_results"`
	EvaluationCriteriaResultsList []ConversationHistoryAnalysisCommonModelEvaluationCriteriaResultsList     `json:"evaluation_criteria_results_list"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSuccessful                respjson.Field
		TranscriptSummary             respjson.Field
		CallSummaryTitle              respjson.Field
		DataCollectionResults         respjson.Field
		DataCollectionResultsList     respjson.Field
		EvaluationCriteriaResults     respjson.Field
		EvaluationCriteriaResultsList respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModel) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryAnalysisCommonModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryAnalysisCommonModelDataCollectionResult struct {
	DataCollectionID string `json:"data_collection_id,required"`
	Rationale        string `json:"rationale,required"`
	// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
	// fields can be set: description (LLM provides value), dynamic_variable (value
	// from variable), is_system_provided (system provides value), or constant_value
	// (fixed value). These are mutually exclusive.
	JsonSchema ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchema `json:"json_schema,nullable"`
	Value      any                                                                  `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataCollectionID respjson.Field
		Rationale        respjson.Field
		JsonSchema       respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelDataCollectionResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelDataCollectionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
type ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchema struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,required"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion `json:"constant_value"`
	// The description of the property. When set, the LLM will provide the value based
	// on this description. Mutually exclusive with dynamic_variable,
	// is_system_provided, and constant_value.
	Description string `json:"description"`
	// The name of the dynamic variable to use for this property's value. Mutually
	// exclusive with description, is_system_provided, and constant_value.
	DynamicVariable string `json:"dynamic_variable"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,nullable"`
	// If true, the value will be populated by the system at runtime. Used by API
	// Integration Webhook tools for templating. Mutually exclusive with description,
	// dynamic_variable, and constant_value.
	IsSystemProvided bool `json:"is_system_provided"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		ConstantValue    respjson.Field
		Description      respjson.Field
		DynamicVariable  respjson.Field
		Enum             respjson.Field
		IsSystemProvided respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchema) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion struct {
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

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationHistoryAnalysisCommonModelDataCollectionResultJsonSchemaConstantValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryAnalysisCommonModelDataCollectionResultsList struct {
	DataCollectionID string `json:"data_collection_id,required"`
	Rationale        string `json:"rationale,required"`
	// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
	// fields can be set: description (LLM provides value), dynamic_variable (value
	// from variable), is_system_provided (system provides value), or constant_value
	// (fixed value). These are mutually exclusive.
	JsonSchema ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchema `json:"json_schema,nullable"`
	Value      any                                                                       `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataCollectionID respjson.Field
		Rationale        respjson.Field
		JsonSchema       respjson.Field
		Value            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelDataCollectionResultsList) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelDataCollectionResultsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
type ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchema struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,required"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion `json:"constant_value"`
	// The description of the property. When set, the LLM will provide the value based
	// on this description. Mutually exclusive with dynamic_variable,
	// is_system_provided, and constant_value.
	Description string `json:"description"`
	// The name of the dynamic variable to use for this property's value. Mutually
	// exclusive with description, is_system_provided, and constant_value.
	DynamicVariable string `json:"dynamic_variable"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,nullable"`
	// If true, the value will be populated by the system at runtime. Used by API
	// Integration Webhook tools for templating. Mutually exclusive with description,
	// dynamic_variable, and constant_value.
	IsSystemProvided bool `json:"is_system_provided"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type             respjson.Field
		ConstantValue    respjson.Field
		Description      respjson.Field
		DynamicVariable  respjson.Field
		Enum             respjson.Field
		IsSystemProvided respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchema) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion struct {
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

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationHistoryAnalysisCommonModelDataCollectionResultsListJsonSchemaConstantValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryAnalysisCommonModelEvaluationCriteriaResult struct {
	CriteriaID string `json:"criteria_id,required"`
	Rationale  string `json:"rationale,required"`
	// Any of "success", "failure", "unknown".
	Result EvaluationSuccessResult `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CriteriaID  respjson.Field
		Rationale   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelEvaluationCriteriaResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelEvaluationCriteriaResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryAnalysisCommonModelEvaluationCriteriaResultsList struct {
	CriteriaID string `json:"criteria_id,required"`
	Rationale  string `json:"rationale,required"`
	// Any of "success", "failure", "unknown".
	Result EvaluationSuccessResult `json:"result,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CriteriaID  respjson.Field
		Rationale   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryAnalysisCommonModelEvaluationCriteriaResultsList) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryAnalysisCommonModelEvaluationCriteriaResultsList) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a message from a multi-voice agent.
type ConversationHistoryMultivoiceMessageModel struct {
	Parts []ConversationHistoryMultivoiceMessageModelPart `json:"parts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryMultivoiceMessageModel) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryMultivoiceMessageModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryMultivoiceMessageModel to a
// ConversationHistoryMultivoiceMessageModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryMultivoiceMessageModelParam.Overrides()
func (r ConversationHistoryMultivoiceMessageModel) ToParam() ConversationHistoryMultivoiceMessageModelParam {
	return param.Override[ConversationHistoryMultivoiceMessageModelParam](json.RawMessage(r.RawJSON()))
}

// Represents a single voice part of a multi-voice message.
type ConversationHistoryMultivoiceMessageModelPart struct {
	Text           string `json:"text,required"`
	TimeInCallSecs int64  `json:"time_in_call_secs,required"`
	VoiceLabel     string `json:"voice_label,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text           respjson.Field
		TimeInCallSecs respjson.Field
		VoiceLabel     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryMultivoiceMessageModelPart) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryMultivoiceMessageModelPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a message from a multi-voice agent.
//
// The property Parts is required.
type ConversationHistoryMultivoiceMessageModelParam struct {
	Parts []ConversationHistoryMultivoiceMessageModelPartParam `json:"parts,omitzero,required"`
	paramObj
}

func (r ConversationHistoryMultivoiceMessageModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryMultivoiceMessageModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryMultivoiceMessageModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents a single voice part of a multi-voice message.
//
// The properties Text, TimeInCallSecs, VoiceLabel are required.
type ConversationHistoryMultivoiceMessageModelPartParam struct {
	TimeInCallSecs param.Opt[int64]  `json:"time_in_call_secs,omitzero,required"`
	VoiceLabel     param.Opt[string] `json:"voice_label,omitzero,required"`
	Text           string            `json:"text,required"`
	paramObj
}

func (r ConversationHistoryMultivoiceMessageModelPartParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryMultivoiceMessageModelPartParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryMultivoiceMessageModelPartParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel struct {
	CredentialID            string                             `json:"credential_id,required"`
	IntegrationConnectionID string                             `json:"integration_connection_id,required"`
	IntegrationID           string                             `json:"integration_id,required"`
	IsError                 bool                               `json:"is_error,required"`
	RequestID               string                             `json:"request_id,required"`
	ResultValue             string                             `json:"result_value,required"`
	ToolHasBeenCalled       bool                               `json:"tool_has_been_called,required"`
	ToolName                string                             `json:"tool_name,required"`
	Type                    constant.APIIntegrationWebhook     `json:"type,required"`
	DynamicVariableUpdates  []DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates"`
	ToolLatencySecs         float64                            `json:"tool_latency_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		IsError                 respjson.Field
		RequestID               respjson.Field
		ResultValue             respjson.Field
		ToolHasBeenCalled       respjson.Field
		ToolName                respjson.Field
		Type                    respjson.Field
		DynamicVariableUpdates  respjson.Field
		ToolLatencySecs         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this
// ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel to a
// ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam.Overrides()
func (r ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) ToParam() ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam {
	return param.Override[ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam](json.RawMessage(r.RawJSON()))
}

// The properties CredentialID, IntegrationConnectionID, IntegrationID, IsError,
// RequestID, ResultValue, ToolHasBeenCalled, ToolName, Type are required.
type ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam struct {
	CredentialID            string                                  `json:"credential_id,required"`
	IntegrationConnectionID string                                  `json:"integration_connection_id,required"`
	IntegrationID           string                                  `json:"integration_id,required"`
	IsError                 bool                                    `json:"is_error,required"`
	RequestID               string                                  `json:"request_id,required"`
	ResultValue             string                                  `json:"result_value,required"`
	ToolHasBeenCalled       bool                                    `json:"tool_has_been_called,required"`
	ToolName                string                                  `json:"tool_name,required"`
	ToolLatencySecs         param.Opt[float64]                      `json:"tool_latency_secs,omitzero"`
	DynamicVariableUpdates  []DynamicVariableUpdateCommonModelParam `json:"dynamic_variable_updates,omitzero"`
	// This field can be elided, and will marshal its zero value as
	// "api_integration_webhook".
	Type constant.APIIntegrationWebhook `json:"type,required"`
	paramObj
}

func (r ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptCommonModelOutput struct {
	// Any of "user", "agent".
	Role                    ConversationHistoryTranscriptCommonModelOutputRole `json:"role,required"`
	TimeInCallSecs          int64                                              `json:"time_in_call_secs,required"`
	AgentMetadata           AgentMetadata                                      `json:"agent_metadata,nullable"`
	ConversationTurnMetrics ConversationTurnMetrics                            `json:"conversation_turn_metrics,nullable"`
	Feedback                UserFeedback                                       `json:"feedback,nullable"`
	Interrupted             bool                                               `json:"interrupted"`
	LlmOverride             string                                             `json:"llm_override,nullable"`
	LlmUsage                LlmUsageOutput                                     `json:"llm_usage,nullable"`
	Message                 string                                             `json:"message,nullable"`
	// Represents a message from a multi-voice agent.
	MultivoiceMessage ConversationHistoryMultivoiceMessageModel `json:"multivoice_message,nullable"`
	OriginalMessage   string                                    `json:"original_message,nullable"`
	RagRetrievalInfo  RagRetrievalInfo                          `json:"rag_retrieval_info,nullable"`
	// Any of "audio", "text".
	SourceMedium ConversationHistoryTranscriptCommonModelOutputSourceMedium      `json:"source_medium,nullable"`
	ToolCalls    []ConversationHistoryTranscriptCommonModelOutputToolCall        `json:"tool_calls"`
	ToolResults  []ConversationHistoryTranscriptCommonModelOutputToolResultUnion `json:"tool_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Role                    respjson.Field
		TimeInCallSecs          respjson.Field
		AgentMetadata           respjson.Field
		ConversationTurnMetrics respjson.Field
		Feedback                respjson.Field
		Interrupted             respjson.Field
		LlmOverride             respjson.Field
		LlmUsage                respjson.Field
		Message                 respjson.Field
		MultivoiceMessage       respjson.Field
		OriginalMessage         respjson.Field
		RagRetrievalInfo        respjson.Field
		SourceMedium            respjson.Field
		ToolCalls               respjson.Field
		ToolResults             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptCommonModelOutput) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptCommonModelOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptCommonModelOutputRole string

const (
	ConversationHistoryTranscriptCommonModelOutputRoleUser  ConversationHistoryTranscriptCommonModelOutputRole = "user"
	ConversationHistoryTranscriptCommonModelOutputRoleAgent ConversationHistoryTranscriptCommonModelOutputRole = "agent"
)

type ConversationHistoryTranscriptCommonModelOutputSourceMedium string

const (
	ConversationHistoryTranscriptCommonModelOutputSourceMediumAudio ConversationHistoryTranscriptCommonModelOutputSourceMedium = "audio"
	ConversationHistoryTranscriptCommonModelOutputSourceMediumText  ConversationHistoryTranscriptCommonModelOutputSourceMedium = "text"
)

type ConversationHistoryTranscriptCommonModelOutputToolCall struct {
	ParamsAsJson      string                                                                 `json:"params_as_json,required"`
	RequestID         string                                                                 `json:"request_id,required"`
	ToolHasBeenCalled bool                                                                   `json:"tool_has_been_called,required"`
	ToolName          string                                                                 `json:"tool_name,required"`
	ToolDetails       ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion `json:"tool_details,nullable"`
	// Any of "system", "webhook", "client", "mcp", "workflow",
	// "api_integration_webhook", "api_integration_mcp".
	Type ToolType `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParamsAsJson      respjson.Field
		RequestID         respjson.Field
		ToolHasBeenCalled respjson.Field
		ToolName          respjson.Field
		ToolDetails       respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptCommonModelOutputToolCall) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptCommonModelOutputToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion contains
// all possible properties and values from
// [ConversationHistoryTranscriptToolCallWebhookDetails],
// [ConversationHistoryTranscriptToolCallClientDetails],
// [ConversationHistoryTranscriptToolCallMcpDetails],
// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
//
// Use the
// [ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion struct {
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Method string `json:"method"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Body string `json:"body"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	PathParams map[string]string `json:"path_params"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	QueryParams map[string]string `json:"query_params"`
	// Any of "webhook", "client", "mcp", "api_integration_webhook".
	Type       string `json:"type"`
	Parameters string `json:"parameters"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	ApprovalPolicy string `json:"approval_policy"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	IntegrationType string `json:"integration_type"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpServerID string `json:"mcp_server_id"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpServerName string `json:"mcp_server_name"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpToolDescription string `json:"mcp_tool_description"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpToolName string `json:"mcp_tool_name"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	RequiresApproval bool `json:"requires_approval"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	CredentialID string `json:"credential_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	IntegrationConnectionID string `json:"integration_connection_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	IntegrationID string `json:"integration_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	WebhookDetails ConversationHistoryTranscriptToolCallWebhookDetails `json:"webhook_details"`
	JSON           struct {
		Method                  respjson.Field
		URL                     respjson.Field
		Body                    respjson.Field
		Headers                 respjson.Field
		PathParams              respjson.Field
		QueryParams             respjson.Field
		Type                    respjson.Field
		Parameters              respjson.Field
		ApprovalPolicy          respjson.Field
		IntegrationType         respjson.Field
		McpServerID             respjson.Field
		McpServerName           respjson.Field
		McpToolDescription      respjson.Field
		McpToolName             respjson.Field
		RequiresApproval        respjson.Field
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		WebhookDetails          respjson.Field
		raw                     string
	} `json:"-"`
}

// anyConversationHistoryTranscriptCommonModelOutputToolCallToolDetails is
// implemented by each variant of
// [ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion] to add
// type safety for the return type of
// [ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion.AsAny]
type anyConversationHistoryTranscriptCommonModelOutputToolCallToolDetails interface {
	implConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion()
}

func (ConversationHistoryTranscriptToolCallWebhookDetails) implConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallClientDetails) implConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallMcpDetails) implConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) implConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion.AsAny().(type) {
//	case elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallClientDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallMcpDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) AsAny() anyConversationHistoryTranscriptCommonModelOutputToolCallToolDetails {
	switch u.Type {
	case "webhook":
		return u.AsWebhook()
	case "client":
		return u.AsClient()
	case "mcp":
		return u.AsMcp()
	case "api_integration_webhook":
		return u.AsAPIIntegrationWebhook()
	}
	return nil
}

func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) AsWebhook() (v ConversationHistoryTranscriptToolCallWebhookDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) AsClient() (v ConversationHistoryTranscriptToolCallClientDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) AsMcp() (v ConversationHistoryTranscriptToolCallMcpDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) AsAPIIntegrationWebhook() (v ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationHistoryTranscriptCommonModelOutputToolCallToolDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationHistoryTranscriptCommonModelOutputToolResultUnion contains all
// possible properties and values from
// [ConversationHistoryTranscriptOtherToolsResultCommonModel],
// [ConversationHistoryTranscriptSystemToolResultCommonModel],
// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel],
// [ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationHistoryTranscriptCommonModelOutputToolResultUnion struct {
	IsError                bool                               `json:"is_error"`
	RequestID              string                             `json:"request_id"`
	ResultValue            string                             `json:"result_value"`
	ToolHasBeenCalled      bool                               `json:"tool_has_been_called"`
	ToolName               string                             `json:"tool_name"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates"`
	ToolLatencySecs        float64                            `json:"tool_latency_secs"`
	Type                   string                             `json:"type"`
	// This field is a union of
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion],
	// [WorkflowToolResponseModelOutput]
	Result ConversationHistoryTranscriptCommonModelOutputToolResultUnionResult `json:"result"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	CredentialID string `json:"credential_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	IntegrationConnectionID string `json:"integration_connection_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	IntegrationID string `json:"integration_id"`
	JSON          struct {
		IsError                 respjson.Field
		RequestID               respjson.Field
		ResultValue             respjson.Field
		ToolHasBeenCalled       respjson.Field
		ToolName                respjson.Field
		DynamicVariableUpdates  respjson.Field
		ToolLatencySecs         respjson.Field
		Type                    respjson.Field
		Result                  respjson.Field
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		raw                     string
	} `json:"-"`
}

func (u ConversationHistoryTranscriptCommonModelOutputToolResultUnion) AsConversationHistoryTranscriptOtherToolsResultCommonModel() (v ConversationHistoryTranscriptOtherToolsResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolResultUnion) AsConversationHistoryTranscriptSystemToolResultCommonModel() (v ConversationHistoryTranscriptSystemToolResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolResultUnion) AsConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel() (v ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptCommonModelOutputToolResultUnion) AsConversationHistoryTranscriptWorkflowToolsResultCommonModel() (v ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationHistoryTranscriptCommonModelOutputToolResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationHistoryTranscriptCommonModelOutputToolResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConversationHistoryTranscriptCommonModelOutputToolResultUnionResult is an
// implicit subunion of
// [ConversationHistoryTranscriptCommonModelOutputToolResultUnion].
// ConversationHistoryTranscriptCommonModelOutputToolResultUnionResult provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConversationHistoryTranscriptCommonModelOutputToolResultUnion].
type ConversationHistoryTranscriptCommonModelOutputToolResultUnionResult struct {
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Message    string `json:"message"`
	Reason     string `json:"reason"`
	ResultType string `json:"result_type"`
	Status     string `json:"status"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Language string `json:"language"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Condition string `json:"condition"`
	FromAgent string `json:"from_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ToAgent string `json:"to_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	DelayMs int64 `json:"delay_ms"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	EnableTransferredAgentFirstMessage bool `json:"enable_transferred_agent_first_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	TransferMessage string `json:"transfer_message"`
	Error           string `json:"error"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	AgentMessage string `json:"agent_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ConferenceName string `json:"conference_name"`
	TransferNumber string `json:"transfer_number"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ClientMessage string `json:"client_message"`
	Note          string `json:"note"`
	Details       string `json:"details"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	DtmfTones string `json:"dtmf_tones"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	VoicemailMessage string `json:"voicemail_message"`
	// This field is from variant [WorkflowToolResponseModelOutput].
	Steps []WorkflowToolResponseModelOutputStepUnion `json:"steps"`
	JSON  struct {
		Message                            respjson.Field
		Reason                             respjson.Field
		ResultType                         respjson.Field
		Status                             respjson.Field
		Language                           respjson.Field
		Condition                          respjson.Field
		FromAgent                          respjson.Field
		ToAgent                            respjson.Field
		DelayMs                            respjson.Field
		EnableTransferredAgentFirstMessage respjson.Field
		TransferMessage                    respjson.Field
		Error                              respjson.Field
		AgentMessage                       respjson.Field
		ConferenceName                     respjson.Field
		TransferNumber                     respjson.Field
		ClientMessage                      respjson.Field
		Note                               respjson.Field
		Details                            respjson.Field
		DtmfTones                          respjson.Field
		VoicemailMessage                   respjson.Field
		Steps                              respjson.Field
		raw                                string
	} `json:"-"`
}

func (r *ConversationHistoryTranscriptCommonModelOutputToolResultUnionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptOtherToolsResultCommonModel struct {
	IsError                bool                               `json:"is_error,required"`
	RequestID              string                             `json:"request_id,required"`
	ResultValue            string                             `json:"result_value,required"`
	ToolHasBeenCalled      bool                               `json:"tool_has_been_called,required"`
	ToolName               string                             `json:"tool_name,required"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates"`
	ToolLatencySecs        float64                            `json:"tool_latency_secs"`
	// Any of "client", "webhook", "mcp".
	Type ConversationHistoryTranscriptOtherToolsResultCommonModelType `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsError                respjson.Field
		RequestID              respjson.Field
		ResultValue            respjson.Field
		ToolHasBeenCalled      respjson.Field
		ToolName               respjson.Field
		DynamicVariableUpdates respjson.Field
		ToolLatencySecs        respjson.Field
		Type                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptOtherToolsResultCommonModel) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptOtherToolsResultCommonModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryTranscriptOtherToolsResultCommonModel
// to a ConversationHistoryTranscriptOtherToolsResultCommonModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptOtherToolsResultCommonModelParam.Overrides()
func (r ConversationHistoryTranscriptOtherToolsResultCommonModel) ToParam() ConversationHistoryTranscriptOtherToolsResultCommonModelParam {
	return param.Override[ConversationHistoryTranscriptOtherToolsResultCommonModelParam](json.RawMessage(r.RawJSON()))
}

type ConversationHistoryTranscriptOtherToolsResultCommonModelType string

const (
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeClient  ConversationHistoryTranscriptOtherToolsResultCommonModelType = "client"
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeWebhook ConversationHistoryTranscriptOtherToolsResultCommonModelType = "webhook"
	ConversationHistoryTranscriptOtherToolsResultCommonModelTypeMcp     ConversationHistoryTranscriptOtherToolsResultCommonModelType = "mcp"
)

// The properties IsError, RequestID, ResultValue, ToolHasBeenCalled, ToolName are
// required.
type ConversationHistoryTranscriptOtherToolsResultCommonModelParam struct {
	IsError           bool               `json:"is_error,required"`
	RequestID         string             `json:"request_id,required"`
	ResultValue       string             `json:"result_value,required"`
	ToolHasBeenCalled bool               `json:"tool_has_been_called,required"`
	ToolName          string             `json:"tool_name,required"`
	ToolLatencySecs   param.Opt[float64] `json:"tool_latency_secs,omitzero"`
	// Any of "client", "webhook", "mcp".
	Type                   ConversationHistoryTranscriptOtherToolsResultCommonModelType `json:"type,omitzero"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModelParam                      `json:"dynamic_variable_updates,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptOtherToolsResultCommonModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptOtherToolsResultCommonModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptOtherToolsResultCommonModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModel struct {
	IsError                bool                                                                `json:"is_error,required"`
	RequestID              string                                                              `json:"request_id,required"`
	ResultValue            string                                                              `json:"result_value,required"`
	ToolHasBeenCalled      bool                                                                `json:"tool_has_been_called,required"`
	ToolName               string                                                              `json:"tool_name,required"`
	Type                   constant.System                                                     `json:"type,required"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModel                                  `json:"dynamic_variable_updates"`
	Result                 ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion `json:"result,nullable"`
	ToolLatencySecs        float64                                                             `json:"tool_latency_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsError                respjson.Field
		RequestID              respjson.Field
		ResultValue            respjson.Field
		ToolHasBeenCalled      respjson.Field
		ToolName               respjson.Field
		Type                   respjson.Field
		DynamicVariableUpdates respjson.Field
		Result                 respjson.Field
		ToolLatencySecs        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModel) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptSystemToolResultCommonModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryTranscriptSystemToolResultCommonModel
// to a ConversationHistoryTranscriptSystemToolResultCommonModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptSystemToolResultCommonModelParam.Overrides()
func (r ConversationHistoryTranscriptSystemToolResultCommonModel) ToParam() ConversationHistoryTranscriptSystemToolResultCommonModelParam {
	return param.Override[ConversationHistoryTranscriptSystemToolResultCommonModelParam](json.RawMessage(r.RawJSON()))
}

// ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion contains all
// possible properties and values from
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess],
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult].
//
// Use the
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion struct {
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess].
	Message string `json:"message"`
	Reason  string `json:"reason"`
	// Any of "end_call_success", "language_detection_success",
	// "transfer_to_agent_success", "transfer_to_agent_error",
	// "transfer_to_number_twilio_success", "transfer_to_number_sip_success",
	// "transfer_to_number_error", "skip_turn_success", "play_dtmf_success",
	// "play_dtmf_error", "voicemail_detection_success", "testing_tool_result".
	ResultType string `json:"result_type"`
	Status     string `json:"status"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess].
	Language string `json:"language"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess].
	Condition string `json:"condition"`
	FromAgent string `json:"from_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess].
	ToAgent string `json:"to_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess].
	DelayMs int64 `json:"delay_ms"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess].
	EnableTransferredAgentFirstMessage bool `json:"enable_transferred_agent_first_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess].
	TransferMessage string `json:"transfer_message"`
	Error           string `json:"error"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess].
	AgentMessage string `json:"agent_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess].
	ConferenceName string `json:"conference_name"`
	TransferNumber string `json:"transfer_number"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess].
	ClientMessage string `json:"client_message"`
	Note          string `json:"note"`
	Details       string `json:"details"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess].
	DtmfTones string `json:"dtmf_tones"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess].
	VoicemailMessage string `json:"voicemail_message"`
	JSON             struct {
		Message                            respjson.Field
		Reason                             respjson.Field
		ResultType                         respjson.Field
		Status                             respjson.Field
		Language                           respjson.Field
		Condition                          respjson.Field
		FromAgent                          respjson.Field
		ToAgent                            respjson.Field
		DelayMs                            respjson.Field
		EnableTransferredAgentFirstMessage respjson.Field
		TransferMessage                    respjson.Field
		Error                              respjson.Field
		AgentMessage                       respjson.Field
		ConferenceName                     respjson.Field
		TransferNumber                     respjson.Field
		ClientMessage                      respjson.Field
		Note                               respjson.Field
		Details                            respjson.Field
		DtmfTones                          respjson.Field
		VoicemailMessage                   respjson.Field
		raw                                string
	} `json:"-"`
}

// anyConversationHistoryTranscriptSystemToolResultCommonModelResult is implemented
// by each variant of
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion] to add
// type safety for the return type of
// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion.AsAny]
type anyConversationHistoryTranscriptSystemToolResultCommonModelResult interface {
	implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion()
}

func (ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}
func (ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult) implConversationHistoryTranscriptSystemToolResultCommonModelResultUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion.AsAny().(type) {
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess:
//	case elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsAny() anyConversationHistoryTranscriptSystemToolResultCommonModelResult {
	switch u.ResultType {
	case "end_call_success":
		return u.AsEndCallSuccess()
	case "language_detection_success":
		return u.AsLanguageDetectionSuccess()
	case "transfer_to_agent_success":
		return u.AsTransferToAgentSuccess()
	case "transfer_to_agent_error":
		return u.AsTransferToAgentError()
	case "transfer_to_number_twilio_success":
		return u.AsTransferToNumberTwilioSuccess()
	case "transfer_to_number_sip_success":
		return u.AsTransferToNumberSipSuccess()
	case "transfer_to_number_error":
		return u.AsTransferToNumberError()
	case "skip_turn_success":
		return u.AsSkipTurnSuccess()
	case "play_dtmf_success":
		return u.AsPlayDtmfSuccess()
	case "play_dtmf_error":
		return u.AsPlayDtmfError()
	case "voicemail_detection_success":
		return u.AsVoicemailDetectionSuccess()
	case "testing_tool_result":
		return u.AsTestingToolResult()
	}
	return nil
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsEndCallSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsLanguageDetectionSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTransferToAgentSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTransferToAgentError() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTransferToNumberTwilioSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTransferToNumberSipSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTransferToNumberError() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsSkipTurnSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsPlayDtmfSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsPlayDtmfError() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsVoicemailDetectionSuccess() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) AsTestingToolResult() (v ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess struct {
	Message string `json:"message,nullable"`
	Reason  string `json:"reason,nullable"`
	// Any of "end_call_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Reason      respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess struct {
	Language string `json:"language,nullable"`
	Reason   string `json:"reason,nullable"`
	// Any of "language_detection_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Reason      respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess struct {
	Condition                          string `json:"condition,required"`
	FromAgent                          string `json:"from_agent,required"`
	ToAgent                            string `json:"to_agent,required"`
	DelayMs                            int64  `json:"delay_ms"`
	EnableTransferredAgentFirstMessage bool   `json:"enable_transferred_agent_first_message"`
	// Any of "transfer_to_agent_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status          string `json:"status"`
	TransferMessage string `json:"transfer_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Condition                          respjson.Field
		FromAgent                          respjson.Field
		ToAgent                            respjson.Field
		DelayMs                            respjson.Field
		EnableTransferredAgentFirstMessage respjson.Field
		ResultType                         respjson.Field
		Status                             respjson.Field
		TransferMessage                    respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError struct {
	Error     string `json:"error,required"`
	FromAgent string `json:"from_agent,required"`
	// Any of "transfer_to_agent_error".
	ResultType string `json:"result_type"`
	// Any of "error".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		FromAgent   respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess struct {
	AgentMessage   string `json:"agent_message,required"`
	ConferenceName string `json:"conference_name,required"`
	TransferNumber string `json:"transfer_number,required"`
	ClientMessage  string `json:"client_message,nullable"`
	Note           string `json:"note,nullable"`
	Reason         string `json:"reason,nullable"`
	// Any of "transfer_to_number_twilio_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentMessage   respjson.Field
		ConferenceName respjson.Field
		TransferNumber respjson.Field
		ClientMessage  respjson.Field
		Note           respjson.Field
		Reason         respjson.Field
		ResultType     respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess struct {
	TransferNumber string `json:"transfer_number,required"`
	Note           string `json:"note,nullable"`
	Reason         string `json:"reason,nullable"`
	// Any of "transfer_to_number_sip_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TransferNumber respjson.Field
		Note           respjson.Field
		Reason         respjson.Field
		ResultType     respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError struct {
	Error   string `json:"error,required"`
	Details string `json:"details,nullable"`
	// Any of "transfer_to_number_error".
	ResultType string `json:"result_type"`
	// Any of "error".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Details     respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess struct {
	Reason string `json:"reason,nullable"`
	// Any of "skip_turn_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess struct {
	DtmfTones string `json:"dtmf_tones,required"`
	Reason    string `json:"reason,nullable"`
	// Any of "play_dtmf_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DtmfTones   respjson.Field
		Reason      respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError struct {
	Error   string `json:"error,required"`
	Details string `json:"details,nullable"`
	// Any of "play_dtmf_error".
	ResultType string `json:"result_type"`
	// Any of "error".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Details     respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess struct {
	Reason string `json:"reason,nullable"`
	// Any of "voicemail_detection_success".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status           string `json:"status"`
	VoicemailMessage string `json:"voicemail_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason           respjson.Field
		ResultType       respjson.Field
		Status           respjson.Field
		VoicemailMessage respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult struct {
	Reason string `json:"reason"`
	// Any of "testing_tool_result".
	ResultType string `json:"result_type"`
	// Any of "success".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ResultType  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IsError, RequestID, ResultValue, ToolHasBeenCalled, ToolName,
// Type are required.
type ConversationHistoryTranscriptSystemToolResultCommonModelParam struct {
	IsError                bool                                                                     `json:"is_error,required"`
	RequestID              string                                                                   `json:"request_id,required"`
	ResultValue            string                                                                   `json:"result_value,required"`
	ToolHasBeenCalled      bool                                                                     `json:"tool_has_been_called,required"`
	ToolName               string                                                                   `json:"tool_name,required"`
	ToolLatencySecs        param.Opt[float64]                                                       `json:"tool_latency_secs,omitzero"`
	Result                 ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam `json:"result,omitzero"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModelParam                                  `json:"dynamic_variable_updates,omitzero"`
	// This field can be elided, and will marshal its zero value as "system".
	Type constant.System `json:"type,required"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam struct {
	OfEndCallSuccess                *ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam                `json:",omitzero,inline"`
	OfLanguageDetectionSuccess      *ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam      `json:",omitzero,inline"`
	OfTransferToAgentSuccess        *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam        `json:",omitzero,inline"`
	OfTransferToAgentError          *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam          `json:",omitzero,inline"`
	OfTransferToNumberTwilioSuccess *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam `json:",omitzero,inline"`
	OfTransferToNumberSipSuccess    *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam    `json:",omitzero,inline"`
	OfTransferToNumberError         *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam         `json:",omitzero,inline"`
	OfSkipTurnSuccess               *ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam               `json:",omitzero,inline"`
	OfPlayDtmfSuccess               *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam               `json:",omitzero,inline"`
	OfPlayDtmfError                 *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam                 `json:",omitzero,inline"`
	OfVoicemailDetectionSuccess     *ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam     `json:",omitzero,inline"`
	OfTestingToolResult             *ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam             `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEndCallSuccess,
		u.OfLanguageDetectionSuccess,
		u.OfTransferToAgentSuccess,
		u.OfTransferToAgentError,
		u.OfTransferToNumberTwilioSuccess,
		u.OfTransferToNumberSipSuccess,
		u.OfTransferToNumberError,
		u.OfSkipTurnSuccess,
		u.OfPlayDtmfSuccess,
		u.OfPlayDtmfError,
		u.OfVoicemailDetectionSuccess,
		u.OfTestingToolResult)
}
func (u *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEndCallSuccess) {
		return u.OfEndCallSuccess
	} else if !param.IsOmitted(u.OfLanguageDetectionSuccess) {
		return u.OfLanguageDetectionSuccess
	} else if !param.IsOmitted(u.OfTransferToAgentSuccess) {
		return u.OfTransferToAgentSuccess
	} else if !param.IsOmitted(u.OfTransferToAgentError) {
		return u.OfTransferToAgentError
	} else if !param.IsOmitted(u.OfTransferToNumberTwilioSuccess) {
		return u.OfTransferToNumberTwilioSuccess
	} else if !param.IsOmitted(u.OfTransferToNumberSipSuccess) {
		return u.OfTransferToNumberSipSuccess
	} else if !param.IsOmitted(u.OfTransferToNumberError) {
		return u.OfTransferToNumberError
	} else if !param.IsOmitted(u.OfSkipTurnSuccess) {
		return u.OfSkipTurnSuccess
	} else if !param.IsOmitted(u.OfPlayDtmfSuccess) {
		return u.OfPlayDtmfSuccess
	} else if !param.IsOmitted(u.OfPlayDtmfError) {
		return u.OfPlayDtmfError
	} else if !param.IsOmitted(u.OfVoicemailDetectionSuccess) {
		return u.OfVoicemailDetectionSuccess
	} else if !param.IsOmitted(u.OfTestingToolResult) {
		return u.OfTestingToolResult
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetMessage() *string {
	if vt := u.OfEndCallSuccess; vt != nil && vt.Message.Valid() {
		return &vt.Message.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetLanguage() *string {
	if vt := u.OfLanguageDetectionSuccess; vt != nil && vt.Language.Valid() {
		return &vt.Language.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetCondition() *string {
	if vt := u.OfTransferToAgentSuccess; vt != nil {
		return &vt.Condition
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetToAgent() *string {
	if vt := u.OfTransferToAgentSuccess; vt != nil {
		return &vt.ToAgent
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetDelayMs() *int64 {
	if vt := u.OfTransferToAgentSuccess; vt != nil && vt.DelayMs.Valid() {
		return &vt.DelayMs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetEnableTransferredAgentFirstMessage() *bool {
	if vt := u.OfTransferToAgentSuccess; vt != nil && vt.EnableTransferredAgentFirstMessage.Valid() {
		return &vt.EnableTransferredAgentFirstMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetTransferMessage() *string {
	if vt := u.OfTransferToAgentSuccess; vt != nil && vt.TransferMessage.Valid() {
		return &vt.TransferMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetAgentMessage() *string {
	if vt := u.OfTransferToNumberTwilioSuccess; vt != nil {
		return &vt.AgentMessage
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetConferenceName() *string {
	if vt := u.OfTransferToNumberTwilioSuccess; vt != nil {
		return &vt.ConferenceName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetClientMessage() *string {
	if vt := u.OfTransferToNumberTwilioSuccess; vt != nil && vt.ClientMessage.Valid() {
		return &vt.ClientMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetDtmfTones() *string {
	if vt := u.OfPlayDtmfSuccess; vt != nil {
		return &vt.DtmfTones
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetVoicemailMessage() *string {
	if vt := u.OfVoicemailDetectionSuccess; vt != nil && vt.VoicemailMessage.Valid() {
		return &vt.VoicemailMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetReason() *string {
	if vt := u.OfEndCallSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfLanguageDetectionSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfTransferToNumberTwilioSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfTransferToNumberSipSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfSkipTurnSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfPlayDtmfSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfVoicemailDetectionSuccess; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	} else if vt := u.OfTestingToolResult; vt != nil && vt.Reason.Valid() {
		return &vt.Reason.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetResultType() *string {
	if vt := u.OfEndCallSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfLanguageDetectionSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTransferToAgentSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTransferToAgentError; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTransferToNumberTwilioSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTransferToNumberSipSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTransferToNumberError; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfSkipTurnSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfPlayDtmfSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfPlayDtmfError; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfVoicemailDetectionSuccess; vt != nil {
		return (*string)(&vt.ResultType)
	} else if vt := u.OfTestingToolResult; vt != nil {
		return (*string)(&vt.ResultType)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetStatus() *string {
	if vt := u.OfEndCallSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfLanguageDetectionSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTransferToAgentSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTransferToAgentError; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTransferToNumberTwilioSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTransferToNumberSipSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTransferToNumberError; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfSkipTurnSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfPlayDtmfSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfPlayDtmfError; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfVoicemailDetectionSuccess; vt != nil {
		return (*string)(&vt.Status)
	} else if vt := u.OfTestingToolResult; vt != nil {
		return (*string)(&vt.Status)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetFromAgent() *string {
	if vt := u.OfTransferToAgentSuccess; vt != nil {
		return (*string)(&vt.FromAgent)
	} else if vt := u.OfTransferToAgentError; vt != nil {
		return (*string)(&vt.FromAgent)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetError() *string {
	if vt := u.OfTransferToAgentError; vt != nil {
		return (*string)(&vt.Error)
	} else if vt := u.OfTransferToNumberError; vt != nil {
		return (*string)(&vt.Error)
	} else if vt := u.OfPlayDtmfError; vt != nil {
		return (*string)(&vt.Error)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetTransferNumber() *string {
	if vt := u.OfTransferToNumberTwilioSuccess; vt != nil {
		return (*string)(&vt.TransferNumber)
	} else if vt := u.OfTransferToNumberSipSuccess; vt != nil {
		return (*string)(&vt.TransferNumber)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetNote() *string {
	if vt := u.OfTransferToNumberTwilioSuccess; vt != nil && vt.Note.Valid() {
		return &vt.Note.Value
	} else if vt := u.OfTransferToNumberSipSuccess; vt != nil && vt.Note.Valid() {
		return &vt.Note.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam) GetDetails() *string {
	if vt := u.OfTransferToNumberError; vt != nil && vt.Details.Valid() {
		return &vt.Details.Value
	} else if vt := u.OfPlayDtmfError; vt != nil && vt.Details.Valid() {
		return &vt.Details.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam](
		"result_type",
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam]("end_call_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam]("language_detection_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam]("transfer_to_agent_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam]("transfer_to_agent_error"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam]("transfer_to_number_twilio_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam]("transfer_to_number_sip_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam]("transfer_to_number_error"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam]("skip_turn_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam]("play_dtmf_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam]("play_dtmf_error"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam]("voicemail_detection_success"),
		apijson.Discriminator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam]("testing_tool_result"),
	)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam struct {
	Message param.Opt[string] `json:"message,omitzero"`
	Reason  param.Opt[string] `json:"reason,omitzero"`
	// Any of "end_call_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam](
		"result_type", "end_call_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam](
		"status", "success",
	)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam struct {
	Language param.Opt[string] `json:"language,omitzero"`
	Reason   param.Opt[string] `json:"reason,omitzero"`
	// Any of "language_detection_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam](
		"result_type", "language_detection_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam](
		"status", "success",
	)
}

// The properties Condition, FromAgent, ToAgent are required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam struct {
	Condition                          string            `json:"condition,required"`
	FromAgent                          string            `json:"from_agent,required"`
	ToAgent                            string            `json:"to_agent,required"`
	TransferMessage                    param.Opt[string] `json:"transfer_message,omitzero"`
	DelayMs                            param.Opt[int64]  `json:"delay_ms,omitzero"`
	EnableTransferredAgentFirstMessage param.Opt[bool]   `json:"enable_transferred_agent_first_message,omitzero"`
	// Any of "transfer_to_agent_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam](
		"result_type", "transfer_to_agent_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam](
		"status", "success",
	)
}

// The properties Error, FromAgent are required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam struct {
	Error     string `json:"error,required"`
	FromAgent string `json:"from_agent,required"`
	// Any of "transfer_to_agent_error".
	ResultType string `json:"result_type,omitzero"`
	// Any of "error".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam](
		"result_type", "transfer_to_agent_error",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam](
		"status", "error",
	)
}

// The properties AgentMessage, ConferenceName, TransferNumber are required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam struct {
	AgentMessage   string            `json:"agent_message,required"`
	ConferenceName string            `json:"conference_name,required"`
	TransferNumber string            `json:"transfer_number,required"`
	ClientMessage  param.Opt[string] `json:"client_message,omitzero"`
	Note           param.Opt[string] `json:"note,omitzero"`
	Reason         param.Opt[string] `json:"reason,omitzero"`
	// Any of "transfer_to_number_twilio_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam](
		"result_type", "transfer_to_number_twilio_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam](
		"status", "success",
	)
}

// The property TransferNumber is required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam struct {
	TransferNumber string            `json:"transfer_number,required"`
	Note           param.Opt[string] `json:"note,omitzero"`
	Reason         param.Opt[string] `json:"reason,omitzero"`
	// Any of "transfer_to_number_sip_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam](
		"result_type", "transfer_to_number_sip_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam](
		"status", "success",
	)
}

// The property Error is required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam struct {
	Error   string            `json:"error,required"`
	Details param.Opt[string] `json:"details,omitzero"`
	// Any of "transfer_to_number_error".
	ResultType string `json:"result_type,omitzero"`
	// Any of "error".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam](
		"result_type", "transfer_to_number_error",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam](
		"status", "error",
	)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam struct {
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Any of "skip_turn_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam](
		"result_type", "skip_turn_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam](
		"status", "success",
	)
}

// The property DtmfTones is required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam struct {
	DtmfTones string            `json:"dtmf_tones,required"`
	Reason    param.Opt[string] `json:"reason,omitzero"`
	// Any of "play_dtmf_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam](
		"result_type", "play_dtmf_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam](
		"status", "success",
	)
}

// The property Error is required.
type ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam struct {
	Error   string            `json:"error,required"`
	Details param.Opt[string] `json:"details,omitzero"`
	// Any of "play_dtmf_error".
	ResultType string `json:"result_type,omitzero"`
	// Any of "error".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam](
		"result_type", "play_dtmf_error",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam](
		"status", "error",
	)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam struct {
	Reason           param.Opt[string] `json:"reason,omitzero"`
	VoicemailMessage param.Opt[string] `json:"voicemail_message,omitzero"`
	// Any of "voicemail_detection_success".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam](
		"result_type", "voicemail_detection_success",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam](
		"status", "success",
	)
}

type ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam struct {
	Reason param.Opt[string] `json:"reason,omitzero"`
	// Any of "testing_tool_result".
	ResultType string `json:"result_type,omitzero"`
	// Any of "success".
	Status string `json:"status,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam](
		"result_type", "testing_tool_result",
	)
	apijson.RegisterFieldValidator[ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam](
		"status", "success",
	)
}

type ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails struct {
	CredentialID            string                                              `json:"credential_id,required"`
	IntegrationConnectionID string                                              `json:"integration_connection_id,required"`
	IntegrationID           string                                              `json:"integration_id,required"`
	WebhookDetails          ConversationHistoryTranscriptToolCallWebhookDetails `json:"webhook_details,required"`
	// Any of "api_integration_webhook".
	Type ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		WebhookDetails          respjson.Field
		Type                    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this
// ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails to a
// ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam.Overrides()
func (r ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) ToParam() ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam {
	return param.Override[ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam](json.RawMessage(r.RawJSON()))
}

type ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsType string

const (
	ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsTypeAPIIntegrationWebhook ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsType = "api_integration_webhook"
)

// The properties CredentialID, IntegrationConnectionID, IntegrationID,
// WebhookDetails are required.
type ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam struct {
	CredentialID            string                                                   `json:"credential_id,required"`
	IntegrationConnectionID string                                                   `json:"integration_connection_id,required"`
	IntegrationID           string                                                   `json:"integration_id,required"`
	WebhookDetails          ConversationHistoryTranscriptToolCallWebhookDetailsParam `json:"webhook_details,omitzero,required"`
	// Any of "api_integration_webhook".
	Type ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsType `json:"type,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptToolCallClientDetails struct {
	Parameters string `json:"parameters,required"`
	// Any of "client".
	Type ConversationHistoryTranscriptToolCallClientDetailsType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptToolCallClientDetails) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptToolCallClientDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryTranscriptToolCallClientDetails to a
// ConversationHistoryTranscriptToolCallClientDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptToolCallClientDetailsParam.Overrides()
func (r ConversationHistoryTranscriptToolCallClientDetails) ToParam() ConversationHistoryTranscriptToolCallClientDetailsParam {
	return param.Override[ConversationHistoryTranscriptToolCallClientDetailsParam](json.RawMessage(r.RawJSON()))
}

type ConversationHistoryTranscriptToolCallClientDetailsType string

const (
	ConversationHistoryTranscriptToolCallClientDetailsTypeClient ConversationHistoryTranscriptToolCallClientDetailsType = "client"
)

// The property Parameters is required.
type ConversationHistoryTranscriptToolCallClientDetailsParam struct {
	Parameters string `json:"parameters,required"`
	// Any of "client".
	Type ConversationHistoryTranscriptToolCallClientDetailsType `json:"type,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptToolCallClientDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptToolCallClientDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptToolCallClientDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptToolCallMcpDetails struct {
	ApprovalPolicy     string            `json:"approval_policy,required"`
	IntegrationType    string            `json:"integration_type,required"`
	McpServerID        string            `json:"mcp_server_id,required"`
	McpServerName      string            `json:"mcp_server_name,required"`
	McpToolDescription string            `json:"mcp_tool_description"`
	McpToolName        string            `json:"mcp_tool_name"`
	Parameters         map[string]string `json:"parameters"`
	RequiresApproval   bool              `json:"requires_approval"`
	// Any of "mcp".
	Type ConversationHistoryTranscriptToolCallMcpDetailsType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApprovalPolicy     respjson.Field
		IntegrationType    respjson.Field
		McpServerID        respjson.Field
		McpServerName      respjson.Field
		McpToolDescription respjson.Field
		McpToolName        respjson.Field
		Parameters         respjson.Field
		RequiresApproval   respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptToolCallMcpDetails) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptToolCallMcpDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryTranscriptToolCallMcpDetails to a
// ConversationHistoryTranscriptToolCallMcpDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptToolCallMcpDetailsParam.Overrides()
func (r ConversationHistoryTranscriptToolCallMcpDetails) ToParam() ConversationHistoryTranscriptToolCallMcpDetailsParam {
	return param.Override[ConversationHistoryTranscriptToolCallMcpDetailsParam](json.RawMessage(r.RawJSON()))
}

type ConversationHistoryTranscriptToolCallMcpDetailsType string

const (
	ConversationHistoryTranscriptToolCallMcpDetailsTypeMcp ConversationHistoryTranscriptToolCallMcpDetailsType = "mcp"
)

// The properties ApprovalPolicy, IntegrationType, McpServerID, McpServerName are
// required.
type ConversationHistoryTranscriptToolCallMcpDetailsParam struct {
	ApprovalPolicy     string            `json:"approval_policy,required"`
	IntegrationType    string            `json:"integration_type,required"`
	McpServerID        string            `json:"mcp_server_id,required"`
	McpServerName      string            `json:"mcp_server_name,required"`
	McpToolDescription param.Opt[string] `json:"mcp_tool_description,omitzero"`
	McpToolName        param.Opt[string] `json:"mcp_tool_name,omitzero"`
	RequiresApproval   param.Opt[bool]   `json:"requires_approval,omitzero"`
	Parameters         map[string]string `json:"parameters,omitzero"`
	// Any of "mcp".
	Type ConversationHistoryTranscriptToolCallMcpDetailsType `json:"type,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptToolCallMcpDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptToolCallMcpDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptToolCallMcpDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptToolCallWebhookDetails struct {
	Method      string            `json:"method,required"`
	URL         string            `json:"url,required"`
	Body        string            `json:"body,nullable"`
	Headers     map[string]string `json:"headers"`
	PathParams  map[string]string `json:"path_params"`
	QueryParams map[string]string `json:"query_params"`
	// Any of "webhook".
	Type ConversationHistoryTranscriptToolCallWebhookDetailsType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Method      respjson.Field
		URL         respjson.Field
		Body        respjson.Field
		Headers     respjson.Field
		PathParams  respjson.Field
		QueryParams respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptToolCallWebhookDetails) RawJSON() string { return r.JSON.raw }
func (r *ConversationHistoryTranscriptToolCallWebhookDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationHistoryTranscriptToolCallWebhookDetails to a
// ConversationHistoryTranscriptToolCallWebhookDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationHistoryTranscriptToolCallWebhookDetailsParam.Overrides()
func (r ConversationHistoryTranscriptToolCallWebhookDetails) ToParam() ConversationHistoryTranscriptToolCallWebhookDetailsParam {
	return param.Override[ConversationHistoryTranscriptToolCallWebhookDetailsParam](json.RawMessage(r.RawJSON()))
}

type ConversationHistoryTranscriptToolCallWebhookDetailsType string

const (
	ConversationHistoryTranscriptToolCallWebhookDetailsTypeWebhook ConversationHistoryTranscriptToolCallWebhookDetailsType = "webhook"
)

// The properties Method, URL are required.
type ConversationHistoryTranscriptToolCallWebhookDetailsParam struct {
	Method      string            `json:"method,required"`
	URL         string            `json:"url,required"`
	Body        param.Opt[string] `json:"body,omitzero"`
	Headers     map[string]string `json:"headers,omitzero"`
	PathParams  map[string]string `json:"path_params,omitzero"`
	QueryParams map[string]string `json:"query_params,omitzero"`
	// Any of "webhook".
	Type ConversationHistoryTranscriptToolCallWebhookDetailsType `json:"type,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptToolCallWebhookDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptToolCallWebhookDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptToolCallWebhookDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput struct {
	IsError                bool                               `json:"is_error,required"`
	RequestID              string                             `json:"request_id,required"`
	ResultValue            string                             `json:"result_value,required"`
	ToolHasBeenCalled      bool                               `json:"tool_has_been_called,required"`
	ToolName               string                             `json:"tool_name,required"`
	Type                   constant.Workflow                  `json:"type,required"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates"`
	// A common model for workflow tool responses.
	Result          *WorkflowToolResponseModelOutput `json:"result,nullable"`
	ToolLatencySecs float64                          `json:"tool_latency_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsError                respjson.Field
		RequestID              respjson.Field
		ResultValue            respjson.Field
		ToolHasBeenCalled      respjson.Field
		ToolName               respjson.Field
		Type                   respjson.Field
		DynamicVariableUpdates respjson.Field
		Result                 respjson.Field
		ToolLatencySecs        respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A specification that will be used to simulate a conversation between an agent
// and an AI user.
//
// The property SimulatedUserConfig is required.
type ConversationSimulationSpecificationParam struct {
	SimulatedUserConfig AgentConfigAPIModelInputParam                                           `json:"simulated_user_config,omitzero,required"`
	DynamicVariables    map[string]ConversationSimulationSpecificationDynamicVariableUnionParam `json:"dynamic_variables,omitzero"`
	// A partial conversation history to start the simulation from. If empty,
	// simulation starts fresh.
	PartialConversationHistory []ConversationHistoryTranscriptInputParam                         `json:"partial_conversation_history,omitzero"`
	ToolMockConfig             map[string]ConversationSimulationSpecificationToolMockConfigParam `json:"tool_mock_config,omitzero"`
	paramObj
}

func (r ConversationSimulationSpecificationParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationSimulationSpecificationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationSimulationSpecificationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationSimulationSpecificationDynamicVariableUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationSimulationSpecificationDynamicVariableUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ConversationSimulationSpecificationDynamicVariableUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationSimulationSpecificationDynamicVariableUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ConversationSimulationSpecificationToolMockConfigParam struct {
	DefaultIsError     param.Opt[bool]   `json:"default_is_error,omitzero"`
	DefaultReturnValue param.Opt[string] `json:"default_return_value,omitzero"`
	paramObj
}

func (r ConversationSimulationSpecificationToolMockConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationSimulationSpecificationToolMockConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationSimulationSpecificationToolMockConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationTurnMetrics struct {
	Metrics map[string]ConversationTurnMetricsMetric `json:"metrics"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metrics     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationTurnMetrics) RawJSON() string { return r.JSON.raw }
func (r *ConversationTurnMetrics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConversationTurnMetrics to a ConversationTurnMetricsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConversationTurnMetricsParam.Overrides()
func (r ConversationTurnMetrics) ToParam() ConversationTurnMetricsParam {
	return param.Override[ConversationTurnMetricsParam](json.RawMessage(r.RawJSON()))
}

type ConversationTurnMetricsMetric struct {
	ElapsedTime float64 `json:"elapsed_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ElapsedTime respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConversationTurnMetricsMetric) RawJSON() string { return r.JSON.raw }
func (r *ConversationTurnMetricsMetric) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationTurnMetricsParam struct {
	Metrics map[string]ConversationTurnMetricsMetricParam `json:"metrics,omitzero"`
	paramObj
}

func (r ConversationTurnMetricsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationTurnMetricsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationTurnMetricsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ElapsedTime is required.
type ConversationTurnMetricsMetricParam struct {
	ElapsedTime float64 `json:"elapsed_time,required"`
	paramObj
}

func (r ConversationTurnMetricsMetricParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationTurnMetricsMetricParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationTurnMetricsMetricParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Tracks a dynamic variable update that occurred during tool execution.
type DynamicVariableUpdateCommonModel struct {
	NewValue      string  `json:"new_value,required"`
	OldValue      string  `json:"old_value,required"`
	ToolName      string  `json:"tool_name,required"`
	ToolRequestID string  `json:"tool_request_id,required"`
	UpdatedAt     float64 `json:"updated_at,required"`
	VariableName  string  `json:"variable_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewValue      respjson.Field
		OldValue      respjson.Field
		ToolName      respjson.Field
		ToolRequestID respjson.Field
		UpdatedAt     respjson.Field
		VariableName  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicVariableUpdateCommonModel) RawJSON() string { return r.JSON.raw }
func (r *DynamicVariableUpdateCommonModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DynamicVariableUpdateCommonModel to a
// DynamicVariableUpdateCommonModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DynamicVariableUpdateCommonModelParam.Overrides()
func (r DynamicVariableUpdateCommonModel) ToParam() DynamicVariableUpdateCommonModelParam {
	return param.Override[DynamicVariableUpdateCommonModelParam](json.RawMessage(r.RawJSON()))
}

// Tracks a dynamic variable update that occurred during tool execution.
//
// The properties NewValue, OldValue, ToolName, ToolRequestID, UpdatedAt,
// VariableName are required.
type DynamicVariableUpdateCommonModelParam struct {
	OldValue      param.Opt[string] `json:"old_value,omitzero,required"`
	NewValue      string            `json:"new_value,required"`
	ToolName      string            `json:"tool_name,required"`
	ToolRequestID string            `json:"tool_request_id,required"`
	UpdatedAt     float64           `json:"updated_at,required"`
	VariableName  string            `json:"variable_name,required"`
	paramObj
}

func (r DynamicVariableUpdateCommonModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DynamicVariableUpdateCommonModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DynamicVariableUpdateCommonModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutput struct {
	ModelUsage map[string]LlmUsageOutputModelUsage `json:"model_usage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ModelUsage  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutput) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutputModelUsage struct {
	Input           LlmUsageOutputModelUsageInput           `json:"input"`
	InputCacheRead  LlmUsageOutputModelUsageInputCacheRead  `json:"input_cache_read"`
	InputCacheWrite LlmUsageOutputModelUsageInputCacheWrite `json:"input_cache_write"`
	OutputTotal     LlmUsageOutputModelUsageOutputTotal     `json:"output_total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input           respjson.Field
		InputCacheRead  respjson.Field
		InputCacheWrite respjson.Field
		OutputTotal     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutputModelUsage) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutputModelUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutputModelUsageInput struct {
	Price  float64 `json:"price"`
	Tokens int64   `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutputModelUsageInput) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutputModelUsageInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutputModelUsageInputCacheRead struct {
	Price  float64 `json:"price"`
	Tokens int64   `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutputModelUsageInputCacheRead) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutputModelUsageInputCacheRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutputModelUsageInputCacheWrite struct {
	Price  float64 `json:"price"`
	Tokens int64   `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutputModelUsageInputCacheWrite) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutputModelUsageInputCacheWrite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageOutputModelUsageOutputTotal struct {
	Price  float64 `json:"price"`
	Tokens int64   `json:"tokens"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Price       respjson.Field
		Tokens      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageOutputModelUsageOutputTotal) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageOutputModelUsageOutputTotal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An evaluation using the transcript and a prompt for a yes/no achieved answer
type PromptEvaluationCriteria struct {
	// The unique identifier for the evaluation criteria
	ID string `json:"id,required"`
	// The prompt that the agent should use to evaluate the conversation
	ConversationGoalPrompt string `json:"conversation_goal_prompt,required"`
	Name                   string `json:"name,required"`
	// The type of evaluation criteria
	//
	// Any of "prompt".
	Type PromptEvaluationCriteriaType `json:"type"`
	// When evaluating the prompt, should the agent's knowledge base be used.
	UseKnowledgeBase bool `json:"use_knowledge_base"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		ConversationGoalPrompt respjson.Field
		Name                   respjson.Field
		Type                   respjson.Field
		UseKnowledgeBase       respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PromptEvaluationCriteria) RawJSON() string { return r.JSON.raw }
func (r *PromptEvaluationCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PromptEvaluationCriteria to a
// PromptEvaluationCriteriaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PromptEvaluationCriteriaParam.Overrides()
func (r PromptEvaluationCriteria) ToParam() PromptEvaluationCriteriaParam {
	return param.Override[PromptEvaluationCriteriaParam](json.RawMessage(r.RawJSON()))
}

// The type of evaluation criteria
type PromptEvaluationCriteriaType string

const (
	PromptEvaluationCriteriaTypePrompt PromptEvaluationCriteriaType = "prompt"
)

// An evaluation using the transcript and a prompt for a yes/no achieved answer
//
// The properties ID, ConversationGoalPrompt, Name are required.
type PromptEvaluationCriteriaParam struct {
	// The unique identifier for the evaluation criteria
	ID string `json:"id,required"`
	// The prompt that the agent should use to evaluate the conversation
	ConversationGoalPrompt string `json:"conversation_goal_prompt,required"`
	Name                   string `json:"name,required"`
	// When evaluating the prompt, should the agent's knowledge base be used.
	UseKnowledgeBase param.Opt[bool] `json:"use_knowledge_base,omitzero"`
	// The type of evaluation criteria
	//
	// Any of "prompt".
	Type PromptEvaluationCriteriaType `json:"type,omitzero"`
	paramObj
}

func (r PromptEvaluationCriteriaParam) MarshalJSON() (data []byte, err error) {
	type shadow PromptEvaluationCriteriaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PromptEvaluationCriteriaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RagRetrievalInfo struct {
	Chunks []RagRetrievalInfoChunk `json:"chunks,required"`
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	EmbeddingModel EmbeddingModel `json:"embedding_model,required"`
	RagLatencySecs float64        `json:"rag_latency_secs,required"`
	RetrievalQuery string         `json:"retrieval_query,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chunks         respjson.Field
		EmbeddingModel respjson.Field
		RagLatencySecs respjson.Field
		RetrievalQuery respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RagRetrievalInfo) RawJSON() string { return r.JSON.raw }
func (r *RagRetrievalInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RagRetrievalInfo to a RagRetrievalInfoParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RagRetrievalInfoParam.Overrides()
func (r RagRetrievalInfo) ToParam() RagRetrievalInfoParam {
	return param.Override[RagRetrievalInfoParam](json.RawMessage(r.RawJSON()))
}

type RagRetrievalInfoChunk struct {
	ChunkID        string  `json:"chunk_id,required"`
	DocumentID     string  `json:"document_id,required"`
	VectorDistance float64 `json:"vector_distance,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID        respjson.Field
		DocumentID     respjson.Field
		VectorDistance respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RagRetrievalInfoChunk) RawJSON() string { return r.JSON.raw }
func (r *RagRetrievalInfoChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Chunks, EmbeddingModel, RagLatencySecs, RetrievalQuery are
// required.
type RagRetrievalInfoParam struct {
	Chunks []RagRetrievalInfoChunkParam `json:"chunks,omitzero,required"`
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	EmbeddingModel EmbeddingModel `json:"embedding_model,omitzero,required"`
	RagLatencySecs float64        `json:"rag_latency_secs,required"`
	RetrievalQuery string         `json:"retrieval_query,required"`
	paramObj
}

func (r RagRetrievalInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow RagRetrievalInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RagRetrievalInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ChunkID, DocumentID, VectorDistance are required.
type RagRetrievalInfoChunkParam struct {
	ChunkID        string  `json:"chunk_id,required"`
	DocumentID     string  `json:"document_id,required"`
	VectorDistance float64 `json:"vector_distance,required"`
	paramObj
}

func (r RagRetrievalInfoChunkParam) MarshalJSON() (data []byte, err error) {
	type shadow RagRetrievalInfoChunkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RagRetrievalInfoChunkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolType string

const (
	ToolTypeSystem                ToolType = "system"
	ToolTypeWebhook               ToolType = "webhook"
	ToolTypeClient                ToolType = "client"
	ToolTypeMcp                   ToolType = "mcp"
	ToolTypeWorkflow              ToolType = "workflow"
	ToolTypeAPIIntegrationWebhook ToolType = "api_integration_webhook"
	ToolTypeAPIIntegrationMcp     ToolType = "api_integration_mcp"
)

type UserFeedback struct {
	// Any of "like", "dislike".
	Score          UserFeedbackScore `json:"score,required"`
	TimeInCallSecs int64             `json:"time_in_call_secs,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Score          respjson.Field
		TimeInCallSecs respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserFeedback) RawJSON() string { return r.JSON.raw }
func (r *UserFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UserFeedback to a UserFeedbackParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UserFeedbackParam.Overrides()
func (r UserFeedback) ToParam() UserFeedbackParam {
	return param.Override[UserFeedbackParam](json.RawMessage(r.RawJSON()))
}

// The properties Score, TimeInCallSecs are required.
type UserFeedbackParam struct {
	// Any of "like", "dislike".
	Score          UserFeedbackScore `json:"score,omitzero,required"`
	TimeInCallSecs int64             `json:"time_in_call_secs,required"`
	paramObj
}

func (r UserFeedbackParam) MarshalJSON() (data []byte, err error) {
	type shadow UserFeedbackParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserFeedbackParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowToolNestedToolsStepModelOutput struct {
	IsSuccessful    bool                                                `json:"is_successful,required"`
	NodeID          string                                              `json:"node_id,required"`
	Requests        []WorkflowToolNestedToolsStepModelOutputRequest     `json:"requests,required"`
	Results         []WorkflowToolNestedToolsStepModelOutputResultUnion `json:"results,required"`
	StepLatencySecs float64                                             `json:"step_latency_secs,required"`
	// Any of "nested_tools".
	Type WorkflowToolNestedToolsStepModelOutputType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccessful    respjson.Field
		NodeID          respjson.Field
		Requests        respjson.Field
		Results         respjson.Field
		StepLatencySecs respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowToolNestedToolsStepModelOutput) RawJSON() string { return r.JSON.raw }
func (r *WorkflowToolNestedToolsStepModelOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowToolNestedToolsStepModelOutputRequest struct {
	ParamsAsJson      string                                                        `json:"params_as_json,required"`
	RequestID         string                                                        `json:"request_id,required"`
	ToolHasBeenCalled bool                                                          `json:"tool_has_been_called,required"`
	ToolName          string                                                        `json:"tool_name,required"`
	ToolDetails       WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion `json:"tool_details,nullable"`
	// Any of "system", "webhook", "client", "mcp", "workflow",
	// "api_integration_webhook", "api_integration_mcp".
	Type ToolType `json:"type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ParamsAsJson      respjson.Field
		RequestID         respjson.Field
		ToolHasBeenCalled respjson.Field
		ToolName          respjson.Field
		ToolDetails       respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowToolNestedToolsStepModelOutputRequest) RawJSON() string { return r.JSON.raw }
func (r *WorkflowToolNestedToolsStepModelOutputRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion contains all
// possible properties and values from
// [ConversationHistoryTranscriptToolCallWebhookDetails],
// [ConversationHistoryTranscriptToolCallClientDetails],
// [ConversationHistoryTranscriptToolCallMcpDetails],
// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
//
// Use the [WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion struct {
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Method string `json:"method"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	URL string `json:"url"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Body string `json:"body"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	Headers map[string]string `json:"headers"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	PathParams map[string]string `json:"path_params"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallWebhookDetails].
	QueryParams map[string]string `json:"query_params"`
	// Any of "webhook", "client", "mcp", "api_integration_webhook".
	Type       string `json:"type"`
	Parameters string `json:"parameters"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	ApprovalPolicy string `json:"approval_policy"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	IntegrationType string `json:"integration_type"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpServerID string `json:"mcp_server_id"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpServerName string `json:"mcp_server_name"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpToolDescription string `json:"mcp_tool_description"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	McpToolName string `json:"mcp_tool_name"`
	// This field is from variant [ConversationHistoryTranscriptToolCallMcpDetails].
	RequiresApproval bool `json:"requires_approval"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	CredentialID string `json:"credential_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	IntegrationConnectionID string `json:"integration_connection_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	IntegrationID string `json:"integration_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails].
	WebhookDetails ConversationHistoryTranscriptToolCallWebhookDetails `json:"webhook_details"`
	JSON           struct {
		Method                  respjson.Field
		URL                     respjson.Field
		Body                    respjson.Field
		Headers                 respjson.Field
		PathParams              respjson.Field
		QueryParams             respjson.Field
		Type                    respjson.Field
		Parameters              respjson.Field
		ApprovalPolicy          respjson.Field
		IntegrationType         respjson.Field
		McpServerID             respjson.Field
		McpServerName           respjson.Field
		McpToolDescription      respjson.Field
		McpToolName             respjson.Field
		RequiresApproval        respjson.Field
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		WebhookDetails          respjson.Field
		raw                     string
	} `json:"-"`
}

// anyWorkflowToolNestedToolsStepModelOutputRequestToolDetails is implemented by
// each variant of [WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion]
// to add type safety for the return type of
// [WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion.AsAny]
type anyWorkflowToolNestedToolsStepModelOutputRequestToolDetails interface {
	implWorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion()
}

func (ConversationHistoryTranscriptToolCallWebhookDetails) implWorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallClientDetails) implWorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallMcpDetails) implWorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion() {
}
func (ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) implWorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion.AsAny().(type) {
//	case elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallClientDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallMcpDetails:
//	case elevenlabs.ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) AsAny() anyWorkflowToolNestedToolsStepModelOutputRequestToolDetails {
	switch u.Type {
	case "webhook":
		return u.AsWebhook()
	case "client":
		return u.AsClient()
	case "mcp":
		return u.AsMcp()
	case "api_integration_webhook":
		return u.AsAPIIntegrationWebhook()
	}
	return nil
}

func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) AsWebhook() (v ConversationHistoryTranscriptToolCallWebhookDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) AsClient() (v ConversationHistoryTranscriptToolCallClientDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) AsMcp() (v ConversationHistoryTranscriptToolCallMcpDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) AsAPIIntegrationWebhook() (v ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetails) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WorkflowToolNestedToolsStepModelOutputRequestToolDetailsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowToolNestedToolsStepModelOutputResultUnion contains all possible
// properties and values from
// [ConversationHistoryTranscriptOtherToolsResultCommonModel],
// [ConversationHistoryTranscriptSystemToolResultCommonModel],
// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel],
// [ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowToolNestedToolsStepModelOutputResultUnion struct {
	IsError                bool                               `json:"is_error"`
	RequestID              string                             `json:"request_id"`
	ResultValue            string                             `json:"result_value"`
	ToolHasBeenCalled      bool                               `json:"tool_has_been_called"`
	ToolName               string                             `json:"tool_name"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModel `json:"dynamic_variable_updates"`
	ToolLatencySecs        float64                            `json:"tool_latency_secs"`
	Type                   string                             `json:"type"`
	// This field is a union of
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion],
	// [WorkflowToolResponseModelOutput]
	Result WorkflowToolNestedToolsStepModelOutputResultUnionResult `json:"result"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	CredentialID string `json:"credential_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	IntegrationConnectionID string `json:"integration_connection_id"`
	// This field is from variant
	// [ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel].
	IntegrationID string `json:"integration_id"`
	JSON          struct {
		IsError                 respjson.Field
		RequestID               respjson.Field
		ResultValue             respjson.Field
		ToolHasBeenCalled       respjson.Field
		ToolName                respjson.Field
		DynamicVariableUpdates  respjson.Field
		ToolLatencySecs         respjson.Field
		Type                    respjson.Field
		Result                  respjson.Field
		CredentialID            respjson.Field
		IntegrationConnectionID respjson.Field
		IntegrationID           respjson.Field
		raw                     string
	} `json:"-"`
}

func (u WorkflowToolNestedToolsStepModelOutputResultUnion) AsConversationHistoryTranscriptOtherToolsResultCommonModel() (v ConversationHistoryTranscriptOtherToolsResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputResultUnion) AsConversationHistoryTranscriptSystemToolResultCommonModel() (v ConversationHistoryTranscriptSystemToolResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputResultUnion) AsConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel() (v ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolNestedToolsStepModelOutputResultUnion) AsConversationHistoryTranscriptWorkflowToolsResultCommonModel() (v ConversationHistoryTranscriptWorkflowToolsResultCommonModelOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowToolNestedToolsStepModelOutputResultUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowToolNestedToolsStepModelOutputResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowToolNestedToolsStepModelOutputResultUnionResult is an implicit subunion
// of [WorkflowToolNestedToolsStepModelOutputResultUnion].
// WorkflowToolNestedToolsStepModelOutputResultUnionResult provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [WorkflowToolNestedToolsStepModelOutputResultUnion].
type WorkflowToolNestedToolsStepModelOutputResultUnionResult struct {
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Message    string `json:"message"`
	Reason     string `json:"reason"`
	ResultType string `json:"result_type"`
	Status     string `json:"status"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Language string `json:"language"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	Condition string `json:"condition"`
	FromAgent string `json:"from_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ToAgent string `json:"to_agent"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	DelayMs int64 `json:"delay_ms"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	EnableTransferredAgentFirstMessage bool `json:"enable_transferred_agent_first_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	TransferMessage string `json:"transfer_message"`
	Error           string `json:"error"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	AgentMessage string `json:"agent_message"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ConferenceName string `json:"conference_name"`
	TransferNumber string `json:"transfer_number"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	ClientMessage string `json:"client_message"`
	Note          string `json:"note"`
	Details       string `json:"details"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	DtmfTones string `json:"dtmf_tones"`
	// This field is from variant
	// [ConversationHistoryTranscriptSystemToolResultCommonModelResultUnion].
	VoicemailMessage string `json:"voicemail_message"`
	// This field is from variant [WorkflowToolResponseModelOutput].
	Steps []WorkflowToolResponseModelOutputStepUnion `json:"steps"`
	JSON  struct {
		Message                            respjson.Field
		Reason                             respjson.Field
		ResultType                         respjson.Field
		Status                             respjson.Field
		Language                           respjson.Field
		Condition                          respjson.Field
		FromAgent                          respjson.Field
		ToAgent                            respjson.Field
		DelayMs                            respjson.Field
		EnableTransferredAgentFirstMessage respjson.Field
		TransferMessage                    respjson.Field
		Error                              respjson.Field
		AgentMessage                       respjson.Field
		ConferenceName                     respjson.Field
		TransferNumber                     respjson.Field
		ClientMessage                      respjson.Field
		Note                               respjson.Field
		Details                            respjson.Field
		DtmfTones                          respjson.Field
		VoicemailMessage                   respjson.Field
		Steps                              respjson.Field
		raw                                string
	} `json:"-"`
}

func (r *WorkflowToolNestedToolsStepModelOutputResultUnionResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowToolNestedToolsStepModelOutputType string

const (
	WorkflowToolNestedToolsStepModelOutputTypeNestedTools WorkflowToolNestedToolsStepModelOutputType = "nested_tools"
)

// A common model for workflow tool responses.
type WorkflowToolResponseModelOutput struct {
	Steps []WorkflowToolResponseModelOutputStepUnion `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Steps       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowToolResponseModelOutput) RawJSON() string { return r.JSON.raw }
func (r *WorkflowToolResponseModelOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WorkflowToolResponseModelOutputStepUnion contains all possible properties and
// values from [WorkflowToolResponseModelOutputStepEdge],
// [WorkflowToolNestedToolsStepModelOutput],
// [WorkflowToolResponseModelOutputStepMaxIterationsExceeded].
//
// Use the [WorkflowToolResponseModelOutputStepUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WorkflowToolResponseModelOutputStepUnion struct {
	// This field is from variant [WorkflowToolResponseModelOutputStepEdge].
	EdgeID          string  `json:"edge_id"`
	StepLatencySecs float64 `json:"step_latency_secs"`
	// This field is from variant [WorkflowToolResponseModelOutputStepEdge].
	TargetNodeID string `json:"target_node_id"`
	// Any of "edge", "nested_tools", "max_iterations_exceeded".
	Type string `json:"type"`
	// This field is from variant [WorkflowToolNestedToolsStepModelOutput].
	IsSuccessful bool `json:"is_successful"`
	// This field is from variant [WorkflowToolNestedToolsStepModelOutput].
	NodeID string `json:"node_id"`
	// This field is from variant [WorkflowToolNestedToolsStepModelOutput].
	Requests []WorkflowToolNestedToolsStepModelOutputRequest `json:"requests"`
	// This field is from variant [WorkflowToolNestedToolsStepModelOutput].
	Results []WorkflowToolNestedToolsStepModelOutputResultUnion `json:"results"`
	// This field is from variant
	// [WorkflowToolResponseModelOutputStepMaxIterationsExceeded].
	MaxIterations int64 `json:"max_iterations"`
	JSON          struct {
		EdgeID          respjson.Field
		StepLatencySecs respjson.Field
		TargetNodeID    respjson.Field
		Type            respjson.Field
		IsSuccessful    respjson.Field
		NodeID          respjson.Field
		Requests        respjson.Field
		Results         respjson.Field
		MaxIterations   respjson.Field
		raw             string
	} `json:"-"`
}

// anyWorkflowToolResponseModelOutputStep is implemented by each variant of
// [WorkflowToolResponseModelOutputStepUnion] to add type safety for the return
// type of [WorkflowToolResponseModelOutputStepUnion.AsAny]
type anyWorkflowToolResponseModelOutputStep interface {
	implWorkflowToolResponseModelOutputStepUnion()
}

func (WorkflowToolResponseModelOutputStepEdge) implWorkflowToolResponseModelOutputStepUnion() {}
func (WorkflowToolNestedToolsStepModelOutput) implWorkflowToolResponseModelOutputStepUnion()  {}
func (WorkflowToolResponseModelOutputStepMaxIterationsExceeded) implWorkflowToolResponseModelOutputStepUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := WorkflowToolResponseModelOutputStepUnion.AsAny().(type) {
//	case elevenlabs.WorkflowToolResponseModelOutputStepEdge:
//	case elevenlabs.WorkflowToolNestedToolsStepModelOutput:
//	case elevenlabs.WorkflowToolResponseModelOutputStepMaxIterationsExceeded:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WorkflowToolResponseModelOutputStepUnion) AsAny() anyWorkflowToolResponseModelOutputStep {
	switch u.Type {
	case "edge":
		return u.AsEdge()
	case "nested_tools":
		return u.AsNestedTools()
	case "max_iterations_exceeded":
		return u.AsMaxIterationsExceeded()
	}
	return nil
}

func (u WorkflowToolResponseModelOutputStepUnion) AsEdge() (v WorkflowToolResponseModelOutputStepEdge) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolResponseModelOutputStepUnion) AsNestedTools() (v WorkflowToolNestedToolsStepModelOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WorkflowToolResponseModelOutputStepUnion) AsMaxIterationsExceeded() (v WorkflowToolResponseModelOutputStepMaxIterationsExceeded) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WorkflowToolResponseModelOutputStepUnion) RawJSON() string { return u.JSON.raw }

func (r *WorkflowToolResponseModelOutputStepUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowToolResponseModelOutputStepEdge struct {
	EdgeID          string  `json:"edge_id,required"`
	StepLatencySecs float64 `json:"step_latency_secs,required"`
	TargetNodeID    string  `json:"target_node_id,required"`
	// Any of "edge".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EdgeID          respjson.Field
		StepLatencySecs respjson.Field
		TargetNodeID    respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowToolResponseModelOutputStepEdge) RawJSON() string { return r.JSON.raw }
func (r *WorkflowToolResponseModelOutputStepEdge) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkflowToolResponseModelOutputStepMaxIterationsExceeded struct {
	MaxIterations   int64   `json:"max_iterations,required"`
	StepLatencySecs float64 `json:"step_latency_secs,required"`
	// Any of "max_iterations_exceeded".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxIterations   respjson.Field
		StepLatencySecs respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkflowToolResponseModelOutputStepMaxIterationsExceeded) RawJSON() string { return r.JSON.raw }
func (r *WorkflowToolResponseModelOutputStepMaxIterationsExceeded) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentSimulateConversationSimulateConversationResponse struct {
	Analysis              ConversationHistoryAnalysisCommonModel           `json:"analysis,required"`
	SimulatedConversation []ConversationHistoryTranscriptCommonModelOutput `json:"simulated_conversation,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Analysis              respjson.Field
		SimulatedConversation respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentSimulateConversationSimulateConversationResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiAgentSimulateConversationSimulateConversationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentSimulateConversationSimulateConversationParams struct {
	// A specification detailing how the conversation should be simulated
	SimulationSpecification ConversationSimulationSpecificationParam `json:"simulation_specification,omitzero,required"`
	// Maximum number of new turns to generate in the conversation simulation
	NewTurnsLimit param.Opt[int64] `json:"new_turns_limit,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of evaluation criteria to test
	ExtraEvaluationCriteria []PromptEvaluationCriteriaParam `json:"extra_evaluation_criteria,omitzero"`
	paramObj
}

func (r ConvaiAgentSimulateConversationSimulateConversationParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentSimulateConversationSimulateConversationParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentSimulateConversationSimulateConversationParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentSimulateConversationStreamParams struct {
	// A specification detailing how the conversation should be simulated
	SimulationSpecification ConversationSimulationSpecificationParam `json:"simulation_specification,omitzero,required"`
	// Maximum number of new turns to generate in the conversation simulation
	NewTurnsLimit param.Opt[int64] `json:"new_turns_limit,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of evaluation criteria to test
	ExtraEvaluationCriteria []PromptEvaluationCriteriaParam `json:"extra_evaluation_criteria,omitzero"`
	paramObj
}

func (r ConvaiAgentSimulateConversationStreamParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentSimulateConversationStreamParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentSimulateConversationStreamParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
