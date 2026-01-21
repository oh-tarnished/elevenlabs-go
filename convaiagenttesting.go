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

// ConvaiAgentTestingService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiAgentTestingService] method instead.
type ConvaiAgentTestingService struct {
	Options []option.RequestOption
}

// NewConvaiAgentTestingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiAgentTestingService(opts ...option.RequestOption) (r ConvaiAgentTestingService) {
	r = ConvaiAgentTestingService{}
	r.Options = opts
	return
}

// Creates a new agent response test.
func (r *ConvaiAgentTestingService) New(ctx context.Context, params ConvaiAgentTestingNewParams, opts ...option.RequestOption) (res *ConvaiAgentTestingNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/agent-testing/create"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets an agent response test by ID.
func (r *ConvaiAgentTestingService) Get(ctx context.Context, testID string, query ConvaiAgentTestingGetParams, opts ...option.RequestOption) (res *GetUnitTestResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agent-testing/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an agent response test by ID.
func (r *ConvaiAgentTestingService) Update(ctx context.Context, testID string, params ConvaiAgentTestingUpdateParams, opts ...option.RequestOption) (res *GetUnitTestResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agent-testing/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Lists all agent response tests with pagination support and optional search
// filtering.
func (r *ConvaiAgentTestingService) List(ctx context.Context, params ConvaiAgentTestingListParams, opts ...option.RequestOption) (res *ConvaiAgentTestingListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/agent-testing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes an agent response test by ID.
func (r *ConvaiAgentTestingService) Delete(ctx context.Context, testID string, body ConvaiAgentTestingDeleteParams, opts ...option.RequestOption) (res *ConvaiAgentTestingDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if testID == "" {
		err = errors.New("missing required test_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agent-testing/%s", testID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets multiple agent response tests by their IDs. Returns a dictionary mapping
// test IDs to test summaries.
func (r *ConvaiAgentTestingService) GetSummaries(ctx context.Context, params ConvaiAgentTestingGetSummariesParams, opts ...option.RequestOption) (res *ConvaiAgentTestingGetSummariesResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/agent-testing/summaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AgentFailureResponseExample struct {
	Response string           `json:"response,required"`
	Type     constant.Failure `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentFailureResponseExample) RawJSON() string { return r.JSON.raw }
func (r *AgentFailureResponseExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentFailureResponseExample to a
// AgentFailureResponseExampleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentFailureResponseExampleParam.Overrides()
func (r AgentFailureResponseExample) ToParam() AgentFailureResponseExampleParam {
	return param.Override[AgentFailureResponseExampleParam](json.RawMessage(r.RawJSON()))
}

// The properties Response, Type are required.
type AgentFailureResponseExampleParam struct {
	Response string `json:"response,required"`
	// This field can be elided, and will marshal its zero value as "failure".
	Type constant.Failure `json:"type,required"`
	paramObj
}

func (r AgentFailureResponseExampleParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentFailureResponseExampleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentFailureResponseExampleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentSuccessfulResponseExample struct {
	Response string           `json:"response,required"`
	Type     constant.Success `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Response    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSuccessfulResponseExample) RawJSON() string { return r.JSON.raw }
func (r *AgentSuccessfulResponseExample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentSuccessfulResponseExample to a
// AgentSuccessfulResponseExampleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentSuccessfulResponseExampleParam.Overrides()
func (r AgentSuccessfulResponseExample) ToParam() AgentSuccessfulResponseExampleParam {
	return param.Override[AgentSuccessfulResponseExampleParam](json.RawMessage(r.RawJSON()))
}

// The properties Response, Type are required.
type AgentSuccessfulResponseExampleParam struct {
	Response string `json:"response,required"`
	// This field can be elided, and will marshal its zero value as "success".
	Type constant.Success `json:"type,required"`
	paramObj
}

func (r AgentSuccessfulResponseExampleParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSuccessfulResponseExampleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSuccessfulResponseExampleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Role, TimeInCallSecs are required.
type ConversationHistoryTranscriptInputParam struct {
	// Any of "user", "agent".
	Role            ConversationHistoryTranscriptInputRole          `json:"role,omitzero,required"`
	TimeInCallSecs  int64                                           `json:"time_in_call_secs,required"`
	LlmOverride     param.Opt[string]                               `json:"llm_override,omitzero"`
	Message         param.Opt[string]                               `json:"message,omitzero"`
	OriginalMessage param.Opt[string]                               `json:"original_message,omitzero"`
	Interrupted     param.Opt[bool]                                 `json:"interrupted,omitzero"`
	LlmUsage        ConversationHistoryTranscriptInputLlmUsageParam `json:"llm_usage,omitzero"`
	// Any of "audio", "text".
	SourceMedium            ConversationHistoryTranscriptInputSourceMedium `json:"source_medium,omitzero"`
	AgentMetadata           AgentMetadataParam                             `json:"agent_metadata,omitzero"`
	ConversationTurnMetrics ConversationTurnMetricsParam                   `json:"conversation_turn_metrics,omitzero"`
	Feedback                UserFeedbackParam                              `json:"feedback,omitzero"`
	// Represents a message from a multi-voice agent.
	MultivoiceMessage ConversationHistoryMultivoiceMessageModelParam           `json:"multivoice_message,omitzero"`
	RagRetrievalInfo  RagRetrievalInfoParam                                    `json:"rag_retrieval_info,omitzero"`
	ToolCalls         []ConversationHistoryTranscriptInputToolCallParam        `json:"tool_calls,omitzero"`
	ToolResults       []ConversationHistoryTranscriptInputToolResultUnionParam `json:"tool_results,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputRole string

const (
	ConversationHistoryTranscriptInputRoleUser  ConversationHistoryTranscriptInputRole = "user"
	ConversationHistoryTranscriptInputRoleAgent ConversationHistoryTranscriptInputRole = "agent"
)

type ConversationHistoryTranscriptInputLlmUsageParam struct {
	ModelUsage map[string]ConversationHistoryTranscriptInputLlmUsageModelUsageParam `json:"model_usage,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputLlmUsageModelUsageParam struct {
	Input           ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam           `json:"input,omitzero"`
	InputCacheRead  ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam  `json:"input_cache_read,omitzero"`
	InputCacheWrite ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam `json:"input_cache_write,omitzero"`
	OutputTotal     ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam     `json:"output_total,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageModelUsageParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageModelUsageParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageModelUsageParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam struct {
	Price  param.Opt[float64] `json:"price,omitzero"`
	Tokens param.Opt[int64]   `json:"tokens,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam struct {
	Price  param.Opt[float64] `json:"price,omitzero"`
	Tokens param.Opt[int64]   `json:"tokens,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam struct {
	Price  param.Opt[float64] `json:"price,omitzero"`
	Tokens param.Opt[int64]   `json:"tokens,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam struct {
	Price  param.Opt[float64] `json:"price,omitzero"`
	Tokens param.Opt[int64]   `json:"tokens,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConversationHistoryTranscriptInputSourceMedium string

const (
	ConversationHistoryTranscriptInputSourceMediumAudio ConversationHistoryTranscriptInputSourceMedium = "audio"
	ConversationHistoryTranscriptInputSourceMediumText  ConversationHistoryTranscriptInputSourceMedium = "text"
)

// The properties ParamsAsJson, RequestID, ToolHasBeenCalled, ToolName are
// required.
type ConversationHistoryTranscriptInputToolCallParam struct {
	ParamsAsJson      string                                                          `json:"params_as_json,required"`
	RequestID         string                                                          `json:"request_id,required"`
	ToolHasBeenCalled bool                                                            `json:"tool_has_been_called,required"`
	ToolName          string                                                          `json:"tool_name,required"`
	ToolDetails       ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam `json:"tool_details,omitzero"`
	// Any of "system", "webhook", "client", "mcp", "workflow",
	// "api_integration_webhook", "api_integration_mcp".
	Type ToolType `json:"type,omitzero"`
	paramObj
}

func (r ConversationHistoryTranscriptInputToolCallParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptInputToolCallParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptInputToolCallParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam struct {
	OfWebhook               *ConversationHistoryTranscriptToolCallWebhookDetailsParam               `json:",omitzero,inline"`
	OfClient                *ConversationHistoryTranscriptToolCallClientDetailsParam                `json:",omitzero,inline"`
	OfMcp                   *ConversationHistoryTranscriptToolCallMcpDetailsParam                   `json:",omitzero,inline"`
	OfAPIIntegrationWebhook *ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook, u.OfClient, u.OfMcp, u.OfAPIIntegrationWebhook)
}
func (u *ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfClient) {
		return u.OfClient
	} else if !param.IsOmitted(u.OfMcp) {
		return u.OfMcp
	} else if !param.IsOmitted(u.OfAPIIntegrationWebhook) {
		return u.OfAPIIntegrationWebhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Method
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetBody() *string {
	if vt := u.OfWebhook; vt != nil && vt.Body.Valid() {
		return &vt.Body.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetHeaders() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetPathParams() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.PathParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetQueryParams() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetApprovalPolicy() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ApprovalPolicy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetIntegrationType() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.IntegrationType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetMcpServerID() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.McpServerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetMcpServerName() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.McpServerName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetMcpToolDescription() *string {
	if vt := u.OfMcp; vt != nil && vt.McpToolDescription.Valid() {
		return &vt.McpToolDescription.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetMcpToolName() *string {
	if vt := u.OfMcp; vt != nil && vt.McpToolName.Valid() {
		return &vt.McpToolName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetRequiresApproval() *bool {
	if vt := u.OfMcp; vt != nil && vt.RequiresApproval.Valid() {
		return &vt.RequiresApproval.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetCredentialID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.CredentialID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetIntegrationConnectionID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.IntegrationConnectionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetIntegrationID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.IntegrationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetWebhookDetails() *ConversationHistoryTranscriptToolCallWebhookDetailsParam {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.WebhookDetails
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetType() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam) GetParameters() (res conversationHistoryTranscriptInputToolCallToolDetailsUnionParamParameters) {
	if vt := u.OfClient; vt != nil {
		res.any = &vt.Parameters
	} else if vt := u.OfMcp; vt != nil {
		res.any = &vt.Parameters
	}
	return
}

// Can have the runtime types [*string], [\*map[string]string]
type conversationHistoryTranscriptInputToolCallToolDetailsUnionParamParameters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *map[string]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u conversationHistoryTranscriptInputToolCallToolDetailsUnionParamParameters) AsAny() any {
	return u.any
}

func init() {
	apijson.RegisterUnion[ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam](
		"type",
		apijson.Discriminator[ConversationHistoryTranscriptToolCallWebhookDetailsParam]("webhook"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallClientDetailsParam]("client"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallMcpDetailsParam]("mcp"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam]("api_integration_webhook"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConversationHistoryTranscriptInputToolResultUnionParam struct {
	OfConversationHistoryTranscriptOtherToolsResultCommonModel                 *ConversationHistoryTranscriptOtherToolsResultCommonModelParam                 `json:",omitzero,inline"`
	OfConversationHistoryTranscriptSystemToolResultCommonModel                 *ConversationHistoryTranscriptSystemToolResultCommonModelParam                 `json:",omitzero,inline"`
	OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel *ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam `json:",omitzero,inline"`
	OfConversationHistoryTranscriptWorkflowToolsResultCommonModel              *ConversationHistoryTranscriptWorkflowToolsResultInputParam                    `json:",omitzero,inline"`
	paramUnion
}

func (u ConversationHistoryTranscriptInputToolResultUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfConversationHistoryTranscriptOtherToolsResultCommonModel, u.OfConversationHistoryTranscriptSystemToolResultCommonModel, u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel, u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel)
}
func (u *ConversationHistoryTranscriptInputToolResultUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConversationHistoryTranscriptInputToolResultUnionParam) asAny() any {
	if !param.IsOmitted(u.OfConversationHistoryTranscriptOtherToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptOtherToolsResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptSystemToolResultCommonModel) {
		return u.OfConversationHistoryTranscriptSystemToolResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetCredentialID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.CredentialID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetIntegrationConnectionID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.IntegrationConnectionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetIntegrationID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.IntegrationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetIsError() *bool {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetRequestID() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetResultValue() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetToolHasBeenCalled() *bool {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetToolName() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetToolLatencySecs() *float64 {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetType() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's DynamicVariableUpdates property,
// if present.
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetDynamicVariableUpdates() []DynamicVariableUpdateCommonModelParam {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ConversationHistoryTranscriptInputToolResultUnionParam) GetResult() (res conversationHistoryTranscriptInputToolResultUnionParamResult) {
	if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		res.any = vt.Result.asAny()
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		res.any = &vt.Result
	}
	return
}

// Can have the runtime types
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam],
// [*WorkflowToolResponseInputParam]
type conversationHistoryTranscriptInputToolResultUnionParamResult struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam:
//	case *elevenlabs.WorkflowToolResponseInputParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetLanguage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetLanguage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetCondition() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetCondition()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetToAgent() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetToAgent()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetDelayMs() *int64 {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDelayMs()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetEnableTransferredAgentFirstMessage() *bool {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetEnableTransferredAgentFirstMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetTransferMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetTransferMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetAgentMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetAgentMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetConferenceName() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetConferenceName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetClientMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetClientMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetDtmfTones() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDtmfTones()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetVoicemailMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetVoicemailMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetSteps() []WorkflowToolResponseInputStepUnionParam {
	switch vt := u.any.(type) {
	case *WorkflowToolResponseInputParam:
		return vt.Steps
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetReason() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetReason()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetResultType() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetResultType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetStatus() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetStatus()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetFromAgent() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetFromAgent()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetError() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetError()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetTransferNumber() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetTransferNumber()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetNote() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetNote()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u conversationHistoryTranscriptInputToolResultUnionParamResult) GetDetails() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDetails()
	}
	return nil
}

// The properties IsError, RequestID, ResultValue, ToolHasBeenCalled, ToolName,
// Type are required.
type ConversationHistoryTranscriptWorkflowToolsResultInputParam struct {
	IsError                bool                                    `json:"is_error,required"`
	RequestID              string                                  `json:"request_id,required"`
	ResultValue            string                                  `json:"result_value,required"`
	ToolHasBeenCalled      bool                                    `json:"tool_has_been_called,required"`
	ToolName               string                                  `json:"tool_name,required"`
	ToolLatencySecs        param.Opt[float64]                      `json:"tool_latency_secs,omitzero"`
	DynamicVariableUpdates []DynamicVariableUpdateCommonModelParam `json:"dynamic_variable_updates,omitzero"`
	// A common model for workflow tool responses.
	Result WorkflowToolResponseInputParam `json:"result,omitzero"`
	// This field can be elided, and will marshal its zero value as "workflow".
	Type constant.Workflow `json:"type,required"`
	paramObj
}

func (r ConversationHistoryTranscriptWorkflowToolsResultInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ConversationHistoryTranscriptWorkflowToolsResultInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConversationHistoryTranscriptWorkflowToolsResultInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetUnitTestResponse struct {
	ID          string                                           `json:"id,required"`
	ChatHistory []ConversationHistoryTranscriptCommonModelOutput `json:"chat_history,required"`
	// Non-empty list of example responses that should be considered failures
	FailureExamples []AgentFailureResponseExample `json:"failure_examples,required"`
	Name            string                        `json:"name,required"`
	// A prompt that evaluates whether the agent's response is successful. Should
	// return True or False.
	SuccessCondition string `json:"success_condition,required"`
	// Non-empty list of example responses that should be considered successful
	SuccessExamples []AgentSuccessfulResponseExample `json:"success_examples,required"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]GetUnitTestResponseDynamicVariableUnion `json:"dynamic_variables"`
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata TestFromConversationMetadataOutput `json:"from_conversation_metadata,nullable"`
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not
	// evaluated.
	ToolCallParameters UnitTestToolCallEvaluationOutput `json:"tool_call_parameters,nullable"`
	// Any of "llm", "tool".
	Type UnitTestCommonModel `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		ChatHistory              respjson.Field
		FailureExamples          respjson.Field
		Name                     respjson.Field
		SuccessCondition         respjson.Field
		SuccessExamples          respjson.Field
		DynamicVariables         respjson.Field
		FromConversationMetadata respjson.Field
		ToolCallParameters       respjson.Field
		Type                     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetUnitTestResponse) RawJSON() string { return r.JSON.raw }
func (r *GetUnitTestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetUnitTestResponseDynamicVariableUnion contains all possible properties and
// values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type GetUnitTestResponseDynamicVariableUnion struct {
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

func (u GetUnitTestResponseDynamicVariableUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetUnitTestResponseDynamicVariableUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetUnitTestResponseDynamicVariableUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetUnitTestResponseDynamicVariableUnion) RawJSON() string { return u.JSON.raw }

func (r *GetUnitTestResponseDynamicVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reference to a tool for unit test evaluation.
type ReferencedToolCommon struct {
	// The ID of the tool
	ID string `json:"id,required"`
	// The type of the tool
	//
	// Any of "system", "webhook", "client", "workflow", "api_integration_webhook".
	Type ReferencedToolCommonType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedToolCommon) RawJSON() string { return r.JSON.raw }
func (r *ReferencedToolCommon) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ReferencedToolCommon to a ReferencedToolCommonParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ReferencedToolCommonParam.Overrides()
func (r ReferencedToolCommon) ToParam() ReferencedToolCommonParam {
	return param.Override[ReferencedToolCommonParam](json.RawMessage(r.RawJSON()))
}

// The type of the tool
type ReferencedToolCommonType string

const (
	ReferencedToolCommonTypeSystem                ReferencedToolCommonType = "system"
	ReferencedToolCommonTypeWebhook               ReferencedToolCommonType = "webhook"
	ReferencedToolCommonTypeClient                ReferencedToolCommonType = "client"
	ReferencedToolCommonTypeWorkflow              ReferencedToolCommonType = "workflow"
	ReferencedToolCommonTypeAPIIntegrationWebhook ReferencedToolCommonType = "api_integration_webhook"
)

// Reference to a tool for unit test evaluation.
//
// The properties ID, Type are required.
type ReferencedToolCommonParam struct {
	// The ID of the tool
	ID string `json:"id,required"`
	// The type of the tool
	//
	// Any of "system", "webhook", "client", "workflow", "api_integration_webhook".
	Type ReferencedToolCommonType `json:"type,omitzero,required"`
	paramObj
}

func (r ReferencedToolCommonParam) MarshalJSON() (data []byte, err error) {
	type shadow ReferencedToolCommonParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReferencedToolCommonParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AgentID, ConversationID are required.
type TestFromConversationMetadataInputParam struct {
	AgentID            string                                    `json:"agent_id,required"`
	ConversationID     string                                    `json:"conversation_id,required"`
	WorkflowNodeID     param.Opt[string]                         `json:"workflow_node_id,omitzero"`
	OriginalAgentReply []ConversationHistoryTranscriptInputParam `json:"original_agent_reply,omitzero"`
	paramObj
}

func (r TestFromConversationMetadataInputParam) MarshalJSON() (data []byte, err error) {
	type shadow TestFromConversationMetadataInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TestFromConversationMetadataInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestFromConversationMetadataOutput struct {
	AgentID            string                                           `json:"agent_id,required"`
	ConversationID     string                                           `json:"conversation_id,required"`
	OriginalAgentReply []ConversationHistoryTranscriptCommonModelOutput `json:"original_agent_reply"`
	WorkflowNodeID     string                                           `json:"workflow_node_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID            respjson.Field
		ConversationID     respjson.Field
		OriginalAgentReply respjson.Field
		WorkflowNodeID     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestFromConversationMetadataOutput) RawJSON() string { return r.JSON.raw }
func (r *TestFromConversationMetadataOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestCommonModel string

const (
	UnitTestCommonModelLlm  UnitTestCommonModel = "llm"
	UnitTestCommonModelTool UnitTestCommonModel = "tool"
)

type UnitTestToolCallEvaluationInputParam struct {
	// Whether to verify that the tool was NOT called.
	VerifyAbsence param.Opt[bool] `json:"verify_absence,omitzero"`
	// Parameters to evaluate for the agent's tool call. If empty, the tool call
	// parameters are not evaluated.
	Parameters []UnitTestToolCallParameter `json:"parameters,omitzero"`
	// Reference to a tool for unit test evaluation.
	ReferencedTool ReferencedToolCommonParam `json:"referenced_tool,omitzero"`
	paramObj
}

func (r UnitTestToolCallEvaluationInputParam) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallEvaluationInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallEvaluationInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallEvaluationOutput struct {
	// Parameters to evaluate for the agent's tool call. If empty, the tool call
	// parameters are not evaluated.
	Parameters []UnitTestToolCallParameterResp `json:"parameters"`
	// Reference to a tool for unit test evaluation.
	ReferencedTool ReferencedToolCommon `json:"referenced_tool,nullable"`
	// Whether to verify that the tool was NOT called.
	VerifyAbsence bool `json:"verify_absence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameters     respjson.Field
		ReferencedTool respjson.Field
		VerifyAbsence  respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallEvaluationOutput) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallEvaluationOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallParameterResp struct {
	Eval UnitTestToolCallParameterEvalUnionResp `json:"eval,required"`
	Path string                                 `json:"path,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Eval        respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallParameterResp) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallParameterResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UnitTestToolCallParameterResp to a
// UnitTestToolCallParameter.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UnitTestToolCallParameter.Overrides()
func (r UnitTestToolCallParameterResp) ToParam() UnitTestToolCallParameter {
	return param.Override[UnitTestToolCallParameter](json.RawMessage(r.RawJSON()))
}

// UnitTestToolCallParameterEvalUnionResp contains all possible properties and
// values from [UnitTestToolCallParameterEvalLlmResp],
// [UnitTestToolCallParameterEvalRegexResp],
// [UnitTestToolCallParameterEvalExactResp],
// [UnitTestToolCallParameterEvalAnythingResp].
//
// Use the [UnitTestToolCallParameterEvalUnionResp.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UnitTestToolCallParameterEvalUnionResp struct {
	// This field is from variant [UnitTestToolCallParameterEvalLlmResp].
	Description string `json:"description"`
	// Any of "llm", "regex", "exact", "anything".
	Type string `json:"type"`
	// This field is from variant [UnitTestToolCallParameterEvalRegexResp].
	Pattern string `json:"pattern"`
	// This field is from variant [UnitTestToolCallParameterEvalExactResp].
	ExpectedValue string `json:"expected_value"`
	JSON          struct {
		Description   respjson.Field
		Type          respjson.Field
		Pattern       respjson.Field
		ExpectedValue respjson.Field
		raw           string
	} `json:"-"`
}

// anyUnitTestToolCallParameterEvalResp is implemented by each variant of
// [UnitTestToolCallParameterEvalUnionResp] to add type safety for the return type
// of [UnitTestToolCallParameterEvalUnionResp.AsAny]
type anyUnitTestToolCallParameterEvalResp interface {
	implUnitTestToolCallParameterEvalUnionResp()
}

func (UnitTestToolCallParameterEvalLlmResp) implUnitTestToolCallParameterEvalUnionResp()      {}
func (UnitTestToolCallParameterEvalRegexResp) implUnitTestToolCallParameterEvalUnionResp()    {}
func (UnitTestToolCallParameterEvalExactResp) implUnitTestToolCallParameterEvalUnionResp()    {}
func (UnitTestToolCallParameterEvalAnythingResp) implUnitTestToolCallParameterEvalUnionResp() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := UnitTestToolCallParameterEvalUnionResp.AsAny().(type) {
//	case elevenlabs.UnitTestToolCallParameterEvalLlmResp:
//	case elevenlabs.UnitTestToolCallParameterEvalRegexResp:
//	case elevenlabs.UnitTestToolCallParameterEvalExactResp:
//	case elevenlabs.UnitTestToolCallParameterEvalAnythingResp:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u UnitTestToolCallParameterEvalUnionResp) AsAny() anyUnitTestToolCallParameterEvalResp {
	switch u.Type {
	case "llm":
		return u.AsLlm()
	case "regex":
		return u.AsRegex()
	case "exact":
		return u.AsExact()
	case "anything":
		return u.AsAnything()
	}
	return nil
}

func (u UnitTestToolCallParameterEvalUnionResp) AsLlm() (v UnitTestToolCallParameterEvalLlmResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnitTestToolCallParameterEvalUnionResp) AsRegex() (v UnitTestToolCallParameterEvalRegexResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnitTestToolCallParameterEvalUnionResp) AsExact() (v UnitTestToolCallParameterEvalExactResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UnitTestToolCallParameterEvalUnionResp) AsAnything() (v UnitTestToolCallParameterEvalAnythingResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UnitTestToolCallParameterEvalUnionResp) RawJSON() string { return u.JSON.raw }

func (r *UnitTestToolCallParameterEvalUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallParameterEvalLlmResp struct {
	// A description of the evaluation strategy to use for the test.
	Description string       `json:"description,required"`
	Type        constant.Llm `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallParameterEvalLlmResp) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallParameterEvalLlmResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallParameterEvalRegexResp struct {
	// A regex pattern to match the agent's response against.
	Pattern string         `json:"pattern,required"`
	Type    constant.Regex `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pattern     respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallParameterEvalRegexResp) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallParameterEvalRegexResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallParameterEvalExactResp struct {
	// The exact string value that the parameter must match.
	ExpectedValue string         `json:"expected_value,required"`
	Type          constant.Exact `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpectedValue respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallParameterEvalExactResp) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallParameterEvalExactResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UnitTestToolCallParameterEvalAnythingResp struct {
	Type constant.Anything `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UnitTestToolCallParameterEvalAnythingResp) RawJSON() string { return r.JSON.raw }
func (r *UnitTestToolCallParameterEvalAnythingResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Eval, Path are required.
type UnitTestToolCallParameter struct {
	Eval UnitTestToolCallParameterEvalUnion `json:"eval,omitzero,required"`
	Path string                             `json:"path,required"`
	paramObj
}

func (r UnitTestToolCallParameter) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallParameter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallParameter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type UnitTestToolCallParameterEvalUnion struct {
	OfLlm      *UnitTestToolCallParameterEvalLlm      `json:",omitzero,inline"`
	OfRegex    *UnitTestToolCallParameterEvalRegex    `json:",omitzero,inline"`
	OfExact    *UnitTestToolCallParameterEvalExact    `json:",omitzero,inline"`
	OfAnything *UnitTestToolCallParameterEvalAnything `json:",omitzero,inline"`
	paramUnion
}

func (u UnitTestToolCallParameterEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLlm, u.OfRegex, u.OfExact, u.OfAnything)
}
func (u *UnitTestToolCallParameterEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *UnitTestToolCallParameterEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfLlm) {
		return u.OfLlm
	} else if !param.IsOmitted(u.OfRegex) {
		return u.OfRegex
	} else if !param.IsOmitted(u.OfExact) {
		return u.OfExact
	} else if !param.IsOmitted(u.OfAnything) {
		return u.OfAnything
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UnitTestToolCallParameterEvalUnion) GetDescription() *string {
	if vt := u.OfLlm; vt != nil {
		return &vt.Description
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UnitTestToolCallParameterEvalUnion) GetPattern() *string {
	if vt := u.OfRegex; vt != nil {
		return &vt.Pattern
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UnitTestToolCallParameterEvalUnion) GetExpectedValue() *string {
	if vt := u.OfExact; vt != nil {
		return &vt.ExpectedValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u UnitTestToolCallParameterEvalUnion) GetType() *string {
	if vt := u.OfLlm; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfRegex; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfExact; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAnything; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[UnitTestToolCallParameterEvalUnion](
		"type",
		apijson.Discriminator[UnitTestToolCallParameterEvalLlm]("llm"),
		apijson.Discriminator[UnitTestToolCallParameterEvalRegex]("regex"),
		apijson.Discriminator[UnitTestToolCallParameterEvalExact]("exact"),
		apijson.Discriminator[UnitTestToolCallParameterEvalAnything]("anything"),
	)
}

// The properties Description, Type are required.
type UnitTestToolCallParameterEvalLlm struct {
	// A description of the evaluation strategy to use for the test.
	Description string `json:"description,required"`
	// This field can be elided, and will marshal its zero value as "llm".
	Type constant.Llm `json:"type,required"`
	paramObj
}

func (r UnitTestToolCallParameterEvalLlm) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallParameterEvalLlm
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallParameterEvalLlm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Pattern, Type are required.
type UnitTestToolCallParameterEvalRegex struct {
	// A regex pattern to match the agent's response against.
	Pattern string `json:"pattern,required"`
	// This field can be elided, and will marshal its zero value as "regex".
	Type constant.Regex `json:"type,required"`
	paramObj
}

func (r UnitTestToolCallParameterEvalRegex) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallParameterEvalRegex
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallParameterEvalRegex) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ExpectedValue, Type are required.
type UnitTestToolCallParameterEvalExact struct {
	// The exact string value that the parameter must match.
	ExpectedValue string `json:"expected_value,required"`
	// This field can be elided, and will marshal its zero value as "exact".
	Type constant.Exact `json:"type,required"`
	paramObj
}

func (r UnitTestToolCallParameterEvalExact) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallParameterEvalExact
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallParameterEvalExact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func NewUnitTestToolCallParameterEvalAnything() UnitTestToolCallParameterEvalAnything {
	return UnitTestToolCallParameterEvalAnything{
		Type: "anything",
	}
}

// This struct has a constant value, construct it with
// [NewUnitTestToolCallParameterEvalAnything].
type UnitTestToolCallParameterEvalAnything struct {
	Type constant.Anything `json:"type,required"`
	paramObj
}

func (r UnitTestToolCallParameterEvalAnything) MarshalJSON() (data []byte, err error) {
	type shadow UnitTestToolCallParameterEvalAnything
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UnitTestToolCallParameterEvalAnything) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties IsSuccessful, NodeID, Requests, Results, StepLatencySecs are
// required.
type WorkflowToolNestedToolsStepInputParam struct {
	IsSuccessful    bool                                               `json:"is_successful,required"`
	NodeID          string                                             `json:"node_id,required"`
	Requests        []WorkflowToolNestedToolsStepInputRequestParam     `json:"requests,omitzero,required"`
	Results         []WorkflowToolNestedToolsStepInputResultUnionParam `json:"results,omitzero,required"`
	StepLatencySecs float64                                            `json:"step_latency_secs,required"`
	// Any of "nested_tools".
	Type WorkflowToolNestedToolsStepInputType `json:"type,omitzero"`
	paramObj
}

func (r WorkflowToolNestedToolsStepInputParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowToolNestedToolsStepInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowToolNestedToolsStepInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ParamsAsJson, RequestID, ToolHasBeenCalled, ToolName are
// required.
type WorkflowToolNestedToolsStepInputRequestParam struct {
	ParamsAsJson      string                                                       `json:"params_as_json,required"`
	RequestID         string                                                       `json:"request_id,required"`
	ToolHasBeenCalled bool                                                         `json:"tool_has_been_called,required"`
	ToolName          string                                                       `json:"tool_name,required"`
	ToolDetails       WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam `json:"tool_details,omitzero"`
	// Any of "system", "webhook", "client", "mcp", "workflow",
	// "api_integration_webhook", "api_integration_mcp".
	Type ToolType `json:"type,omitzero"`
	paramObj
}

func (r WorkflowToolNestedToolsStepInputRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowToolNestedToolsStepInputRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowToolNestedToolsStepInputRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam struct {
	OfWebhook               *ConversationHistoryTranscriptToolCallWebhookDetailsParam               `json:",omitzero,inline"`
	OfClient                *ConversationHistoryTranscriptToolCallClientDetailsParam                `json:",omitzero,inline"`
	OfMcp                   *ConversationHistoryTranscriptToolCallMcpDetailsParam                   `json:",omitzero,inline"`
	OfAPIIntegrationWebhook *ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook, u.OfClient, u.OfMcp, u.OfAPIIntegrationWebhook)
}
func (u *WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfClient) {
		return u.OfClient
	} else if !param.IsOmitted(u.OfMcp) {
		return u.OfMcp
	} else if !param.IsOmitted(u.OfAPIIntegrationWebhook) {
		return u.OfAPIIntegrationWebhook
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetMethod() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.Method
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetURL() *string {
	if vt := u.OfWebhook; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetBody() *string {
	if vt := u.OfWebhook; vt != nil && vt.Body.Valid() {
		return &vt.Body.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetHeaders() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.Headers
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetPathParams() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.PathParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetQueryParams() map[string]string {
	if vt := u.OfWebhook; vt != nil {
		return vt.QueryParams
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetApprovalPolicy() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.ApprovalPolicy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetIntegrationType() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.IntegrationType
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetMcpServerID() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.McpServerID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetMcpServerName() *string {
	if vt := u.OfMcp; vt != nil {
		return &vt.McpServerName
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetMcpToolDescription() *string {
	if vt := u.OfMcp; vt != nil && vt.McpToolDescription.Valid() {
		return &vt.McpToolDescription.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetMcpToolName() *string {
	if vt := u.OfMcp; vt != nil && vt.McpToolName.Valid() {
		return &vt.McpToolName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetRequiresApproval() *bool {
	if vt := u.OfMcp; vt != nil && vt.RequiresApproval.Valid() {
		return &vt.RequiresApproval.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetCredentialID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.CredentialID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetIntegrationConnectionID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.IntegrationConnectionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetIntegrationID() *string {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.IntegrationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetWebhookDetails() *ConversationHistoryTranscriptToolCallWebhookDetailsParam {
	if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return &vt.WebhookDetails
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetType() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMcp; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfAPIIntegrationWebhook; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam) GetParameters() (res workflowToolNestedToolsStepInputRequestToolDetailsUnionParamParameters) {
	if vt := u.OfClient; vt != nil {
		res.any = &vt.Parameters
	} else if vt := u.OfMcp; vt != nil {
		res.any = &vt.Parameters
	}
	return
}

// Can have the runtime types [*string], [\*map[string]string]
type workflowToolNestedToolsStepInputRequestToolDetailsUnionParamParameters struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *map[string]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowToolNestedToolsStepInputRequestToolDetailsUnionParamParameters) AsAny() any {
	return u.any
}

func init() {
	apijson.RegisterUnion[WorkflowToolNestedToolsStepInputRequestToolDetailsUnionParam](
		"type",
		apijson.Discriminator[ConversationHistoryTranscriptToolCallWebhookDetailsParam]("webhook"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallClientDetailsParam]("client"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallMcpDetailsParam]("mcp"),
		apijson.Discriminator[ConversationHistoryTranscriptToolCallAPIIntegrationWebhookDetailsParam]("api_integration_webhook"),
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowToolNestedToolsStepInputResultUnionParam struct {
	OfConversationHistoryTranscriptOtherToolsResultCommonModel                 *ConversationHistoryTranscriptOtherToolsResultCommonModelParam                 `json:",omitzero,inline"`
	OfConversationHistoryTranscriptSystemToolResultCommonModel                 *ConversationHistoryTranscriptSystemToolResultCommonModelParam                 `json:",omitzero,inline"`
	OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel *ConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModelParam `json:",omitzero,inline"`
	OfConversationHistoryTranscriptWorkflowToolsResultCommonModel              *ConversationHistoryTranscriptWorkflowToolsResultInputParam                    `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowToolNestedToolsStepInputResultUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfConversationHistoryTranscriptOtherToolsResultCommonModel, u.OfConversationHistoryTranscriptSystemToolResultCommonModel, u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel, u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel)
}
func (u *WorkflowToolNestedToolsStepInputResultUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowToolNestedToolsStepInputResultUnionParam) asAny() any {
	if !param.IsOmitted(u.OfConversationHistoryTranscriptOtherToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptOtherToolsResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptSystemToolResultCommonModel) {
		return u.OfConversationHistoryTranscriptSystemToolResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel
	} else if !param.IsOmitted(u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel) {
		return u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetCredentialID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.CredentialID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetIntegrationConnectionID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.IntegrationConnectionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetIntegrationID() *string {
	if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return &vt.IntegrationID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetIsError() *bool {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.IsError)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetRequestID() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.RequestID)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetResultValue() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ResultValue)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetToolHasBeenCalled() *bool {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*bool)(&vt.ToolHasBeenCalled)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetToolName() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.ToolName)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetToolLatencySecs() *float64 {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil && vt.ToolLatencySecs.Valid() {
		return &vt.ToolLatencySecs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetType() *string {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's DynamicVariableUpdates property,
// if present.
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetDynamicVariableUpdates() []DynamicVariableUpdateCommonModelParam {
	if vt := u.OfConversationHistoryTranscriptOtherToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptAPIIntegrationWebhookToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		return vt.DynamicVariableUpdates
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u WorkflowToolNestedToolsStepInputResultUnionParam) GetResult() (res workflowToolNestedToolsStepInputResultUnionParamResult) {
	if vt := u.OfConversationHistoryTranscriptSystemToolResultCommonModel; vt != nil {
		res.any = vt.Result.asAny()
	} else if vt := u.OfConversationHistoryTranscriptWorkflowToolsResultCommonModel; vt != nil {
		res.any = &vt.Result
	}
	return
}

// Can have the runtime types
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam],
// [*ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam],
// [*WorkflowToolResponseInputParam]
type workflowToolNestedToolsStepInputResultUnionParamResult struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultEndCallSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultLanguageDetectionSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToAgentErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberTwilioSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberSipSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTransferToNumberErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultSkipTurnSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultPlayDtmfErrorParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultVoicemailDetectionSuccessParam:
//	case *elevenlabs.ConversationHistoryTranscriptSystemToolResultCommonModelResultTestingToolResultParam:
//	case *elevenlabs.WorkflowToolResponseInputParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u workflowToolNestedToolsStepInputResultUnionParamResult) AsAny() any { return u.any }

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetLanguage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetLanguage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetCondition() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetCondition()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetToAgent() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetToAgent()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetDelayMs() *int64 {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDelayMs()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetEnableTransferredAgentFirstMessage() *bool {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetEnableTransferredAgentFirstMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetTransferMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetTransferMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetAgentMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetAgentMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetConferenceName() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetConferenceName()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetClientMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetClientMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetDtmfTones() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDtmfTones()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetVoicemailMessage() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetVoicemailMessage()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetSteps() []WorkflowToolResponseInputStepUnionParam {
	switch vt := u.any.(type) {
	case *WorkflowToolResponseInputParam:
		return vt.Steps
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetReason() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetReason()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetResultType() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetResultType()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetStatus() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetStatus()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetFromAgent() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetFromAgent()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetError() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetError()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetTransferNumber() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetTransferNumber()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetNote() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetNote()
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u workflowToolNestedToolsStepInputResultUnionParamResult) GetDetails() *string {
	switch vt := u.any.(type) {
	case *ConversationHistoryTranscriptSystemToolResultCommonModelResultUnionParam:
		return vt.GetDetails()
	}
	return nil
}

type WorkflowToolNestedToolsStepInputType string

const (
	WorkflowToolNestedToolsStepInputTypeNestedTools WorkflowToolNestedToolsStepInputType = "nested_tools"
)

// A common model for workflow tool responses.
type WorkflowToolResponseInputParam struct {
	Steps []WorkflowToolResponseInputStepUnionParam `json:"steps,omitzero"`
	paramObj
}

func (r WorkflowToolResponseInputParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowToolResponseInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowToolResponseInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WorkflowToolResponseInputStepUnionParam struct {
	OfEdge                  *WorkflowToolResponseInputStepEdgeParam                  `json:",omitzero,inline"`
	OfNestedTools           *WorkflowToolNestedToolsStepInputParam                   `json:",omitzero,inline"`
	OfMaxIterationsExceeded *WorkflowToolResponseInputStepMaxIterationsExceededParam `json:",omitzero,inline"`
	paramUnion
}

func (u WorkflowToolResponseInputStepUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEdge, u.OfNestedTools, u.OfMaxIterationsExceeded)
}
func (u *WorkflowToolResponseInputStepUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WorkflowToolResponseInputStepUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEdge) {
		return u.OfEdge
	} else if !param.IsOmitted(u.OfNestedTools) {
		return u.OfNestedTools
	} else if !param.IsOmitted(u.OfMaxIterationsExceeded) {
		return u.OfMaxIterationsExceeded
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetEdgeID() *string {
	if vt := u.OfEdge; vt != nil {
		return &vt.EdgeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetTargetNodeID() *string {
	if vt := u.OfEdge; vt != nil {
		return &vt.TargetNodeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetIsSuccessful() *bool {
	if vt := u.OfNestedTools; vt != nil {
		return &vt.IsSuccessful
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetNodeID() *string {
	if vt := u.OfNestedTools; vt != nil {
		return &vt.NodeID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetRequests() []WorkflowToolNestedToolsStepInputRequestParam {
	if vt := u.OfNestedTools; vt != nil {
		return vt.Requests
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetResults() []WorkflowToolNestedToolsStepInputResultUnionParam {
	if vt := u.OfNestedTools; vt != nil {
		return vt.Results
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetMaxIterations() *int64 {
	if vt := u.OfMaxIterationsExceeded; vt != nil {
		return &vt.MaxIterations
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetStepLatencySecs() *float64 {
	if vt := u.OfEdge; vt != nil {
		return (*float64)(&vt.StepLatencySecs)
	} else if vt := u.OfNestedTools; vt != nil {
		return (*float64)(&vt.StepLatencySecs)
	} else if vt := u.OfMaxIterationsExceeded; vt != nil {
		return (*float64)(&vt.StepLatencySecs)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WorkflowToolResponseInputStepUnionParam) GetType() *string {
	if vt := u.OfEdge; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNestedTools; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfMaxIterationsExceeded; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[WorkflowToolResponseInputStepUnionParam](
		"type",
		apijson.Discriminator[WorkflowToolResponseInputStepEdgeParam]("edge"),
		apijson.Discriminator[WorkflowToolNestedToolsStepInputParam]("nested_tools"),
		apijson.Discriminator[WorkflowToolResponseInputStepMaxIterationsExceededParam]("max_iterations_exceeded"),
	)
}

// The properties EdgeID, StepLatencySecs, TargetNodeID are required.
type WorkflowToolResponseInputStepEdgeParam struct {
	EdgeID          string  `json:"edge_id,required"`
	StepLatencySecs float64 `json:"step_latency_secs,required"`
	TargetNodeID    string  `json:"target_node_id,required"`
	// Any of "edge".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WorkflowToolResponseInputStepEdgeParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowToolResponseInputStepEdgeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowToolResponseInputStepEdgeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowToolResponseInputStepEdgeParam](
		"type", "edge",
	)
}

// The properties MaxIterations, StepLatencySecs are required.
type WorkflowToolResponseInputStepMaxIterationsExceededParam struct {
	MaxIterations   int64   `json:"max_iterations,required"`
	StepLatencySecs float64 `json:"step_latency_secs,required"`
	// Any of "max_iterations_exceeded".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r WorkflowToolResponseInputStepMaxIterationsExceededParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkflowToolResponseInputStepMaxIterationsExceededParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkflowToolResponseInputStepMaxIterationsExceededParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WorkflowToolResponseInputStepMaxIterationsExceededParam](
		"type", "max_iterations_exceeded",
	)
}

type ConvaiAgentTestingNewResponse struct {
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentTestingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentTestingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentTestingListResponse struct {
	HasMore    bool                                 `json:"has_more,required"`
	Tests      []ConvaiAgentTestingListResponseTest `json:"tests,required"`
	NextCursor string                               `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Tests       respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentTestingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentTestingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentTestingListResponseTest struct {
	// The ID of the test
	ID string `json:"id,required"`
	// Creation time of the test in unix seconds
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs,required"`
	// Last update time of the test in unix seconds
	LastUpdatedAtUnixSecs int64 `json:"last_updated_at_unix_secs,required"`
	// Name of the test
	Name string `json:"name,required"`
	// Type of the test
	//
	// Any of "llm", "tool".
	Type UnitTestCommonModel `json:"type,required"`
	// The access information of the test
	AccessInfo ResourceAccessInfo `json:"access_info,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAtUnixSecs     respjson.Field
		LastUpdatedAtUnixSecs respjson.Field
		Name                  respjson.Field
		Type                  respjson.Field
		AccessInfo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentTestingListResponseTest) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentTestingListResponseTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentTestingDeleteResponse = any

type ConvaiAgentTestingGetSummariesResponse struct {
	// Dictionary mapping test IDs to their summary information
	Tests map[string]ConvaiAgentTestingGetSummariesResponseTest `json:"tests,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tests       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentTestingGetSummariesResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentTestingGetSummariesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentTestingGetSummariesResponseTest struct {
	// The ID of the test
	ID string `json:"id,required"`
	// Creation time of the test in unix seconds
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs,required"`
	// Last update time of the test in unix seconds
	LastUpdatedAtUnixSecs int64 `json:"last_updated_at_unix_secs,required"`
	// Name of the test
	Name string `json:"name,required"`
	// Type of the test
	//
	// Any of "llm", "tool".
	Type UnitTestCommonModel `json:"type,required"`
	// The access information of the test
	AccessInfo ResourceAccessInfo `json:"access_info,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedAtUnixSecs     respjson.Field
		LastUpdatedAtUnixSecs respjson.Field
		Name                  respjson.Field
		Type                  respjson.Field
		AccessInfo            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentTestingGetSummariesResponseTest) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentTestingGetSummariesResponseTest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentTestingNewParams struct {
	ChatHistory []ConversationHistoryTranscriptInputParam `json:"chat_history,omitzero,required"`
	// Non-empty list of example responses that should be considered failures
	FailureExamples []AgentFailureResponseExampleParam `json:"failure_examples,omitzero,required"`
	Name            string                             `json:"name,required"`
	// A prompt that evaluates whether the agent's response is successful. Should
	// return True or False.
	SuccessCondition string `json:"success_condition,required"`
	// Non-empty list of example responses that should be considered successful
	SuccessExamples []AgentSuccessfulResponseExampleParam `json:"success_examples,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]ConvaiAgentTestingNewParamsDynamicVariableUnion `json:"dynamic_variables,omitzero"`
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata TestFromConversationMetadataInputParam `json:"from_conversation_metadata,omitzero"`
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not
	// evaluated.
	ToolCallParameters UnitTestToolCallEvaluationInputParam `json:"tool_call_parameters,omitzero"`
	// Any of "llm", "tool".
	Type UnitTestCommonModel `json:"type,omitzero"`
	paramObj
}

func (r ConvaiAgentTestingNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentTestingNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentTestingNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiAgentTestingNewParamsDynamicVariableUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiAgentTestingNewParamsDynamicVariableUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ConvaiAgentTestingNewParamsDynamicVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiAgentTestingNewParamsDynamicVariableUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ConvaiAgentTestingGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiAgentTestingUpdateParams struct {
	ChatHistory []ConversationHistoryTranscriptInputParam `json:"chat_history,omitzero,required"`
	// Non-empty list of example responses that should be considered failures
	FailureExamples []AgentFailureResponseExampleParam `json:"failure_examples,omitzero,required"`
	Name            string                             `json:"name,required"`
	// A prompt that evaluates whether the agent's response is successful. Should
	// return True or False.
	SuccessCondition string `json:"success_condition,required"`
	// Non-empty list of example responses that should be considered successful
	SuccessExamples []AgentSuccessfulResponseExampleParam `json:"success_examples,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Dynamic variables to replace in the agent config during testing
	DynamicVariables map[string]ConvaiAgentTestingUpdateParamsDynamicVariableUnion `json:"dynamic_variables,omitzero"`
	// Metadata of a conversation this test was created from (if applicable).
	FromConversationMetadata TestFromConversationMetadataInputParam `json:"from_conversation_metadata,omitzero"`
	// How to evaluate the agent's tool call (if any). If empty, the tool call is not
	// evaluated.
	ToolCallParameters UnitTestToolCallEvaluationInputParam `json:"tool_call_parameters,omitzero"`
	// Any of "llm", "tool".
	Type UnitTestCommonModel `json:"type,omitzero"`
	paramObj
}

func (r ConvaiAgentTestingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentTestingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentTestingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiAgentTestingUpdateParamsDynamicVariableUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiAgentTestingUpdateParamsDynamicVariableUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ConvaiAgentTestingUpdateParamsDynamicVariableUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiAgentTestingUpdateParamsDynamicVariableUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ConvaiAgentTestingListParams struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Search query to filter tests by name.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// How many Tests to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiAgentTestingListParams]'s query parameters as
// `url.Values`.
func (r ConvaiAgentTestingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiAgentTestingDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiAgentTestingGetSummariesParams struct {
	// List of test IDs to fetch. No duplicates allowed.
	TestIDs []string `json:"test_ids,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiAgentTestingGetSummariesParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentTestingGetSummariesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentTestingGetSummariesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
