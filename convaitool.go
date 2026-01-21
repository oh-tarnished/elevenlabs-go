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
	shimjson "github.com/oh-tarnished/elevenlabs-go/internal/encoding/json"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiToolService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiToolService] method instead.
type ConvaiToolService struct {
	Options []option.RequestOption
}

// NewConvaiToolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiToolService(opts ...option.RequestOption) (r ConvaiToolService) {
	r = ConvaiToolService{}
	r.Options = opts
	return
}

// Add a new tool to the available tools in the workspace.
func (r *ConvaiToolService) New(ctx context.Context, params ConvaiToolNewParams, opts ...option.RequestOption) (res *ToolResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get tool that is available in the workspace.
func (r *ConvaiToolService) Get(ctx context.Context, toolID string, query ConvaiToolGetParams, opts ...option.RequestOption) (res *ToolResponseModel, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update tool that is available in the workspace.
func (r *ConvaiToolService) Update(ctx context.Context, toolID string, params ConvaiToolUpdateParams, opts ...option.RequestOption) (res *ToolResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get all available tools in the workspace.
func (r *ConvaiToolService) List(ctx context.Context, query ConvaiToolListParams, opts ...option.RequestOption) (res *ConvaiToolListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete tool from the workspace.
func (r *ConvaiToolService) Delete(ctx context.Context, toolID string, body ConvaiToolDeleteParams, opts ...option.RequestOption) (res *ConvaiToolDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/tools/%s", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a list of agents depending on this tool
func (r *ConvaiToolService) ListDependentAgents(ctx context.Context, toolID string, params ConvaiToolListDependentAgentsParams, opts ...option.RequestOption) (res *ConvaiToolListDependentAgentsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/tools/%s/dependent-agents", toolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Used to reference an auth connection from the workspace's auth connection store.
type AuthConnectionLocator struct {
	AuthConnectionID string `json:"auth_connection_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthConnectionID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthConnectionLocator) RawJSON() string { return r.JSON.raw }
func (r *AuthConnectionLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AuthConnectionLocator to a AuthConnectionLocatorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AuthConnectionLocatorParam.Overrides()
func (r AuthConnectionLocator) ToParam() AuthConnectionLocatorParam {
	return param.Override[AuthConnectionLocatorParam](json.RawMessage(r.RawJSON()))
}

// Used to reference an auth connection from the workspace's auth connection store.
//
// The property AuthConnectionID is required.
type AuthConnectionLocatorParam struct {
	AuthConnectionID string `json:"auth_connection_id,required"`
	paramObj
}

func (r AuthConnectionLocatorParam) MarshalJSON() (data []byte, err error) {
	type shadow AuthConnectionLocatorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthConnectionLocatorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A client tool is one that sends an event to the user's client to trigger
// something client side
//
// The properties Description, Name are required.
type ClientToolConfigInputParam struct {
	// Description of when the tool should be used and what it does.
	Description string `json:"description,required"`
	Name        string `json:"name,required"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If true, calling this tool should block the conversation until the client
	// responds with some response which is passed to the llm. If false then we will
	// continue the conversation without waiting for the client to respond, this is
	// useful to show content to a user but not block the conversation
	ExpectsResponse param.Opt[bool] `json:"expects_response,omitzero"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// The maximum time in seconds to wait for the tool call to complete. Must be
	// between 1 and 120 seconds (inclusive).
	ResponseTimeoutSecs param.Opt[int64] `json:"response_timeout_secs,omitzero"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// Configuration for dynamic variables
	DynamicVariables DynamicVariablesConfigParam `json:"dynamic_variables,omitzero"`
	// Determines when and how the tool executes: 'immediate' executes the tool right
	// away when requested by the LLM, 'post_tool_speech' waits for the agent to finish
	// speaking before executing, 'async' runs the tool in the background without
	// blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// Schema for any parameters to pass to the client
	Parameters ObjectJsonSchemaPropertyInputParam `json:"parameters,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	// The type of tool
	//
	// Any of "client".
	Type ClientToolConfigInputType `json:"type,omitzero"`
	paramObj
}

func (r ClientToolConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ClientToolConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientToolConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tool
type ClientToolConfigInputType string

const (
	ClientToolConfigInputTypeClient ClientToolConfigInputType = "client"
)

// A client tool is one that sends an event to the user's client to trigger
// something client side
type ClientToolConfigOutput struct {
	// Description of when the tool should be used and what it does.
	Description string `json:"description,required"`
	Name        string `json:"name,required"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignment `json:"assignments"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions bool `json:"disable_interruptions"`
	// Configuration for dynamic variables
	DynamicVariables DynamicVariablesConfig `json:"dynamic_variables"`
	// Determines when and how the tool executes: 'immediate' executes the tool right
	// away when requested by the LLM, 'post_tool_speech' waits for the agent to finish
	// speaking before executing, 'async' runs the tool in the background without
	// blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode"`
	// If true, calling this tool should block the conversation until the client
	// responds with some response which is passed to the llm. If false then we will
	// continue the conversation without waiting for the client to respond, this is
	// useful to show content to a user but not block the conversation
	ExpectsResponse bool `json:"expects_response"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech bool `json:"force_pre_tool_speech"`
	// Schema for any parameters to pass to the client
	Parameters ObjectJsonSchemaPropertyOutput `json:"parameters,nullable"`
	// The maximum time in seconds to wait for the tool call to complete. Must be
	// between 1 and 120 seconds (inclusive).
	ResponseTimeoutSecs int64 `json:"response_timeout_secs"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,nullable"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior"`
	// The type of tool
	//
	// Any of "client".
	Type ClientToolConfigOutputType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description           respjson.Field
		Name                  respjson.Field
		Assignments           respjson.Field
		DisableInterruptions  respjson.Field
		DynamicVariables      respjson.Field
		ExecutionMode         respjson.Field
		ExpectsResponse       respjson.Field
		ForcePreToolSpeech    respjson.Field
		Parameters            respjson.Field
		ResponseTimeoutSecs   respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientToolConfigOutput) RawJSON() string { return r.JSON.raw }
func (r *ClientToolConfigOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tool
type ClientToolConfigOutputType string

const (
	ClientToolConfigOutputTypeClient ClientToolConfigOutputType = "client"
)

type DynamicVariablesConfig struct {
	// A dictionary of dynamic variable placeholders and their values
	DynamicVariablePlaceholders map[string]DynamicVariablesConfigDynamicVariablePlaceholderUnion `json:"dynamic_variable_placeholders"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DynamicVariablePlaceholders respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicVariablesConfig) RawJSON() string { return r.JSON.raw }
func (r *DynamicVariablesConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DynamicVariablesConfig to a DynamicVariablesConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DynamicVariablesConfigParam.Overrides()
func (r DynamicVariablesConfig) ToParam() DynamicVariablesConfigParam {
	return param.Override[DynamicVariablesConfigParam](json.RawMessage(r.RawJSON()))
}

// DynamicVariablesConfigDynamicVariablePlaceholderUnion contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type DynamicVariablesConfigDynamicVariablePlaceholderUnion struct {
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

func (u DynamicVariablesConfigDynamicVariablePlaceholderUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DynamicVariablesConfigDynamicVariablePlaceholderUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DynamicVariablesConfigDynamicVariablePlaceholderUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DynamicVariablesConfigDynamicVariablePlaceholderUnion) RawJSON() string { return u.JSON.raw }

func (r *DynamicVariablesConfigDynamicVariablePlaceholderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DynamicVariablesConfigParam struct {
	// A dictionary of dynamic variable placeholders and their values
	DynamicVariablePlaceholders map[string]DynamicVariablesConfigDynamicVariablePlaceholderUnionParam `json:"dynamic_variable_placeholders,omitzero"`
	paramObj
}

func (r DynamicVariablesConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DynamicVariablesConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DynamicVariablesConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type DynamicVariablesConfigDynamicVariablePlaceholderUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u DynamicVariablesConfigDynamicVariablePlaceholderUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *DynamicVariablesConfigDynamicVariablePlaceholderUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *DynamicVariablesConfigDynamicVariablePlaceholderUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type EndCallToolConfig struct {
	// Any of "end_call".
	SystemToolType EndCallToolConfigSystemToolType `json:"system_tool_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemToolType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EndCallToolConfig) RawJSON() string { return r.JSON.raw }
func (r *EndCallToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EndCallToolConfig to a EndCallToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EndCallToolConfigParam.Overrides()
func (r EndCallToolConfig) ToParam() EndCallToolConfigParam {
	return param.Override[EndCallToolConfigParam](json.RawMessage(r.RawJSON()))
}

type EndCallToolConfigSystemToolType string

const (
	EndCallToolConfigSystemToolTypeEndCall EndCallToolConfigSystemToolType = "end_call"
)

type EndCallToolConfigParam struct {
	// Any of "end_call".
	SystemToolType EndCallToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r EndCallToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow EndCallToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EndCallToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LanguageDetectionToolConfig struct {
	// Any of "language_detection".
	SystemToolType LanguageDetectionToolConfigSystemToolType `json:"system_tool_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemToolType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LanguageDetectionToolConfig) RawJSON() string { return r.JSON.raw }
func (r *LanguageDetectionToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this LanguageDetectionToolConfig to a
// LanguageDetectionToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// LanguageDetectionToolConfigParam.Overrides()
func (r LanguageDetectionToolConfig) ToParam() LanguageDetectionToolConfigParam {
	return param.Override[LanguageDetectionToolConfigParam](json.RawMessage(r.RawJSON()))
}

type LanguageDetectionToolConfigSystemToolType string

const (
	LanguageDetectionToolConfigSystemToolTypeLanguageDetection LanguageDetectionToolConfigSystemToolType = "language_detection"
)

type LanguageDetectionToolConfigParam struct {
	// Any of "language_detection".
	SystemToolType LanguageDetectionToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r LanguageDetectionToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow LanguageDetectionToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LanguageDetectionToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectJsonSchemaPropertyInputParam struct {
	Description param.Opt[string]                                          `json:"description,omitzero"`
	Properties  map[string]ObjectJsonSchemaPropertyInputPropertyUnionParam `json:"properties,omitzero"`
	Required    []string                                                   `json:"required,omitzero"`
	// Any of "object".
	Type ObjectJsonSchemaPropertyInputType `json:"type,omitzero"`
	paramObj
}

func (r ObjectJsonSchemaPropertyInputParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectJsonSchemaPropertyInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectJsonSchemaPropertyInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectJsonSchemaPropertyInputPropertyUnionParam struct {
	OfLiteralJsonSchemaProperty *ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam `json:",omitzero,inline"`
	OfObjectJsonSchemaProperty  *ObjectJsonSchemaPropertyInputParam                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfLiteralJsonSchemaProperty, u.OfObjectJsonSchemaProperty)
}
func (u *ObjectJsonSchemaPropertyInputPropertyUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ObjectJsonSchemaPropertyInputPropertyUnionParam) asAny() any {
	if !param.IsOmitted(u.OfLiteralJsonSchemaProperty) {
		return u.OfLiteralJsonSchemaProperty
	} else if !param.IsOmitted(u.OfObjectJsonSchemaProperty) {
		return u.OfObjectJsonSchemaProperty
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetConstantValue() *ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil {
		return &vt.ConstantValue
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetDynamicVariable() *string {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil && vt.DynamicVariable.Valid() {
		return &vt.DynamicVariable.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetEnum() []string {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil {
		return vt.Enum
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetIsSystemProvided() *bool {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil && vt.IsSystemProvided.Valid() {
		return &vt.IsSystemProvided.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetProperties() map[string]ObjectJsonSchemaPropertyInputPropertyUnionParam {
	if vt := u.OfObjectJsonSchemaProperty; vt != nil {
		return vt.Properties
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetRequired() []string {
	if vt := u.OfObjectJsonSchemaProperty; vt != nil {
		return vt.Required
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetType() *string {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfObjectJsonSchemaProperty; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ObjectJsonSchemaPropertyInputPropertyUnionParam) GetDescription() *string {
	if vt := u.OfLiteralJsonSchemaProperty; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	} else if vt := u.OfObjectJsonSchemaProperty; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
//
// The property Type is required.
type ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,omitzero,required"`
	// The description of the property. When set, the LLM will provide the value based
	// on this description. Mutually exclusive with dynamic_variable,
	// is_system_provided, and constant_value.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the dynamic variable to use for this property's value. Mutually
	// exclusive with description, is_system_provided, and constant_value.
	DynamicVariable param.Opt[string] `json:"dynamic_variable,omitzero"`
	// If true, the value will be populated by the system at runtime. Used by API
	// Integration Webhook tools for templating. Mutually exclusive with description,
	// dynamic_variable, and constant_value.
	IsSystemProvided param.Opt[bool] `json:"is_system_provided,omitzero"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,omitzero"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam `json:"constant_value,omitzero"`
	paramObj
}

func (r ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam) MarshalJSON() (data []byte, err error) {
	type shadow ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam](
		"type", "boolean", "string", "integer", "number",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ObjectJsonSchemaPropertyInputType string

const (
	ObjectJsonSchemaPropertyInputTypeObject ObjectJsonSchemaPropertyInputType = "object"
)

type ObjectJsonSchemaPropertyOutput struct {
	Description string                                                 `json:"description"`
	Properties  map[string]ObjectJsonSchemaPropertyOutputPropertyUnion `json:"properties"`
	Required    []string                                               `json:"required"`
	// Any of "object".
	Type ObjectJsonSchemaPropertyOutputType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Properties  respjson.Field
		Required    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ObjectJsonSchemaPropertyOutput) RawJSON() string { return r.JSON.raw }
func (r *ObjectJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectJsonSchemaPropertyOutputPropertyUnion contains all possible properties and
// values from [ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty],
// [ObjectJsonSchemaPropertyOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ObjectJsonSchemaPropertyOutputPropertyUnion struct {
	Type string `json:"type"`
	// This field is from variant
	// [ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty].
	ConstantValue ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion `json:"constant_value"`
	Description   string                                                                            `json:"description"`
	// This field is from variant
	// [ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty].
	DynamicVariable string `json:"dynamic_variable"`
	// This field is from variant
	// [ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty].
	Enum []string `json:"enum"`
	// This field is from variant
	// [ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty].
	IsSystemProvided bool `json:"is_system_provided"`
	// This field is from variant [ObjectJsonSchemaPropertyOutput].
	Properties map[string]ObjectJsonSchemaPropertyOutputPropertyUnion `json:"properties"`
	// This field is from variant [ObjectJsonSchemaPropertyOutput].
	Required []string `json:"required"`
	JSON     struct {
		Type             respjson.Field
		ConstantValue    respjson.Field
		Description      respjson.Field
		DynamicVariable  respjson.Field
		Enum             respjson.Field
		IsSystemProvided respjson.Field
		Properties       respjson.Field
		Required         respjson.Field
		raw              string
	} `json:"-"`
}

func (u ObjectJsonSchemaPropertyOutputPropertyUnion) AsLiteralJsonSchemaProperty() (v ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectJsonSchemaPropertyOutputPropertyUnion) AsObjectJsonSchemaProperty() (v ObjectJsonSchemaPropertyOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectJsonSchemaPropertyOutputPropertyUnion) RawJSON() string { return u.JSON.raw }

func (r *ObjectJsonSchemaPropertyOutputPropertyUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
type ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,required"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion `json:"constant_value"`
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
func (r ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty) RawJSON() string {
	return r.JSON.raw
}
func (r *ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion
// contains all possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion struct {
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

func (u ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ObjectJsonSchemaPropertyOutputPropertyLiteralJsonSchemaPropertyConstantValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ObjectJsonSchemaPropertyOutputType string

const (
	ObjectJsonSchemaPropertyOutputTypeObject ObjectJsonSchemaPropertyOutputType = "object"
)

type PhoneNumberTransfer struct {
	Condition string `json:"condition,required"`
	// Custom SIP headers to include when transferring the call. Each header can be
	// either a static value or a dynamic variable reference.
	CustomSipHeaders []PhoneNumberTransferCustomSipHeaderUnion `json:"custom_sip_headers"`
	// Deprecated: deprecated
	PhoneNumber         string                                      `json:"phone_number,nullable"`
	TransferDestination PhoneNumberTransferTransferDestinationUnion `json:"transfer_destination,nullable"`
	// Any of "blind", "conference", "sip_refer".
	TransferType PhoneNumberTransferTransferType `json:"transfer_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Condition           respjson.Field
		CustomSipHeaders    respjson.Field
		PhoneNumber         respjson.Field
		TransferDestination respjson.Field
		TransferType        respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransfer) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PhoneNumberTransfer to a PhoneNumberTransferParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PhoneNumberTransferParam.Overrides()
func (r PhoneNumberTransfer) ToParam() PhoneNumberTransferParam {
	return param.Override[PhoneNumberTransferParam](json.RawMessage(r.RawJSON()))
}

// PhoneNumberTransferCustomSipHeaderUnion contains all possible properties and
// values from [PhoneNumberTransferCustomSipHeaderStatic],
// [PhoneNumberTransferCustomSipHeaderDynamic].
//
// Use the [PhoneNumberTransferCustomSipHeaderUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PhoneNumberTransferCustomSipHeaderUnion struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// Any of "static", "dynamic".
	Type string `json:"type"`
	JSON struct {
		Key   respjson.Field
		Value respjson.Field
		Type  respjson.Field
		raw   string
	} `json:"-"`
}

// anyPhoneNumberTransferCustomSipHeader is implemented by each variant of
// [PhoneNumberTransferCustomSipHeaderUnion] to add type safety for the return type
// of [PhoneNumberTransferCustomSipHeaderUnion.AsAny]
type anyPhoneNumberTransferCustomSipHeader interface {
	implPhoneNumberTransferCustomSipHeaderUnion()
}

func (PhoneNumberTransferCustomSipHeaderStatic) implPhoneNumberTransferCustomSipHeaderUnion()  {}
func (PhoneNumberTransferCustomSipHeaderDynamic) implPhoneNumberTransferCustomSipHeaderUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PhoneNumberTransferCustomSipHeaderUnion.AsAny().(type) {
//	case elevenlabs.PhoneNumberTransferCustomSipHeaderStatic:
//	case elevenlabs.PhoneNumberTransferCustomSipHeaderDynamic:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PhoneNumberTransferCustomSipHeaderUnion) AsAny() anyPhoneNumberTransferCustomSipHeader {
	switch u.Type {
	case "static":
		return u.AsStatic()
	case "dynamic":
		return u.AsDynamic()
	}
	return nil
}

func (u PhoneNumberTransferCustomSipHeaderUnion) AsStatic() (v PhoneNumberTransferCustomSipHeaderStatic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PhoneNumberTransferCustomSipHeaderUnion) AsDynamic() (v PhoneNumberTransferCustomSipHeaderDynamic) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PhoneNumberTransferCustomSipHeaderUnion) RawJSON() string { return u.JSON.raw }

func (r *PhoneNumberTransferCustomSipHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom SIP header for phone transfers with a static (validated) value.
type PhoneNumberTransferCustomSipHeaderStatic struct {
	// The SIP header name (e.g., 'X-Customer-ID')
	Key string `json:"key,required"`
	// The header value
	Value string `json:"value,required"`
	// Any of "static".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Value       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferCustomSipHeaderStatic) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberTransferCustomSipHeaderStatic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom SIP header for phone transfers with a dynamic variable reference. The
// value is a variable name that will be resolved at runtime. Value is not
// validated here since it will be substituted with actual value later.
type PhoneNumberTransferCustomSipHeaderDynamic struct {
	// The SIP header name (e.g., 'X-Customer-ID')
	Key  string           `json:"key,required"`
	Type constant.Dynamic `json:"type,required"`
	// The dynamic variable name to resolve
	Value string `json:"value,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key         respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferCustomSipHeaderDynamic) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberTransferCustomSipHeaderDynamic) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PhoneNumberTransferTransferDestinationUnion contains all possible properties and
// values from [PhoneNumberTransferTransferDestinationPhone],
// [PhoneNumberTransferTransferDestinationSipUri],
// [PhoneNumberTransferTransferDestinationPhoneDynamicVariable],
// [PhoneNumberTransferTransferDestinationSipUriDynamicVariable].
//
// Use the [PhoneNumberTransferTransferDestinationUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PhoneNumberTransferTransferDestinationUnion struct {
	PhoneNumber string `json:"phone_number"`
	// Any of "phone", "sip_uri", "phone_dynamic_variable", "sip_uri_dynamic_variable".
	Type   string `json:"type"`
	SipUri string `json:"sip_uri"`
	JSON   struct {
		PhoneNumber respjson.Field
		Type        respjson.Field
		SipUri      respjson.Field
		raw         string
	} `json:"-"`
}

// anyPhoneNumberTransferTransferDestination is implemented by each variant of
// [PhoneNumberTransferTransferDestinationUnion] to add type safety for the return
// type of [PhoneNumberTransferTransferDestinationUnion.AsAny]
type anyPhoneNumberTransferTransferDestination interface {
	implPhoneNumberTransferTransferDestinationUnion()
}

func (PhoneNumberTransferTransferDestinationPhone) implPhoneNumberTransferTransferDestinationUnion() {
}
func (PhoneNumberTransferTransferDestinationSipUri) implPhoneNumberTransferTransferDestinationUnion() {
}
func (PhoneNumberTransferTransferDestinationPhoneDynamicVariable) implPhoneNumberTransferTransferDestinationUnion() {
}
func (PhoneNumberTransferTransferDestinationSipUriDynamicVariable) implPhoneNumberTransferTransferDestinationUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := PhoneNumberTransferTransferDestinationUnion.AsAny().(type) {
//	case elevenlabs.PhoneNumberTransferTransferDestinationPhone:
//	case elevenlabs.PhoneNumberTransferTransferDestinationSipUri:
//	case elevenlabs.PhoneNumberTransferTransferDestinationPhoneDynamicVariable:
//	case elevenlabs.PhoneNumberTransferTransferDestinationSipUriDynamicVariable:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PhoneNumberTransferTransferDestinationUnion) AsAny() anyPhoneNumberTransferTransferDestination {
	switch u.Type {
	case "phone":
		return u.AsPhone()
	case "sip_uri":
		return u.AsSipUri()
	case "phone_dynamic_variable":
		return u.AsPhoneDynamicVariable()
	case "sip_uri_dynamic_variable":
		return u.AsSipUriDynamicVariable()
	}
	return nil
}

func (u PhoneNumberTransferTransferDestinationUnion) AsPhone() (v PhoneNumberTransferTransferDestinationPhone) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PhoneNumberTransferTransferDestinationUnion) AsSipUri() (v PhoneNumberTransferTransferDestinationSipUri) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PhoneNumberTransferTransferDestinationUnion) AsPhoneDynamicVariable() (v PhoneNumberTransferTransferDestinationPhoneDynamicVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PhoneNumberTransferTransferDestinationUnion) AsSipUriDynamicVariable() (v PhoneNumberTransferTransferDestinationSipUriDynamicVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PhoneNumberTransferTransferDestinationUnion) RawJSON() string { return u.JSON.raw }

func (r *PhoneNumberTransferTransferDestinationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberTransferTransferDestinationPhone struct {
	PhoneNumber string `json:"phone_number,required"`
	// Any of "phone".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferTransferDestinationPhone) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberTransferTransferDestinationPhone) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberTransferTransferDestinationSipUri struct {
	SipUri string `json:"sip_uri,required"`
	// Any of "sip_uri".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SipUri      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferTransferDestinationSipUri) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberTransferTransferDestinationSipUri) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberTransferTransferDestinationPhoneDynamicVariable struct {
	PhoneNumber string `json:"phone_number,required"`
	// Any of "phone_dynamic_variable".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferTransferDestinationPhoneDynamicVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberTransferTransferDestinationPhoneDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberTransferTransferDestinationSipUriDynamicVariable struct {
	SipUri string `json:"sip_uri,required"`
	// Any of "sip_uri_dynamic_variable".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SipUri      respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberTransferTransferDestinationSipUriDynamicVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *PhoneNumberTransferTransferDestinationSipUriDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberTransferTransferType string

const (
	PhoneNumberTransferTransferTypeBlind      PhoneNumberTransferTransferType = "blind"
	PhoneNumberTransferTransferTypeConference PhoneNumberTransferTransferType = "conference"
	PhoneNumberTransferTransferTypeSipRefer   PhoneNumberTransferTransferType = "sip_refer"
)

// The property Condition is required.
type PhoneNumberTransferParam struct {
	Condition string `json:"condition,required"`
	// Deprecated: deprecated
	PhoneNumber         param.Opt[string]                                `json:"phone_number,omitzero"`
	TransferDestination PhoneNumberTransferTransferDestinationUnionParam `json:"transfer_destination,omitzero"`
	// Custom SIP headers to include when transferring the call. Each header can be
	// either a static value or a dynamic variable reference.
	CustomSipHeaders []PhoneNumberTransferCustomSipHeaderUnionParam `json:"custom_sip_headers,omitzero"`
	// Any of "blind", "conference", "sip_refer".
	TransferType PhoneNumberTransferTransferType `json:"transfer_type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PhoneNumberTransferCustomSipHeaderUnionParam struct {
	OfStatic  *PhoneNumberTransferCustomSipHeaderStaticParam  `json:",omitzero,inline"`
	OfDynamic *PhoneNumberTransferCustomSipHeaderDynamicParam `json:",omitzero,inline"`
	paramUnion
}

func (u PhoneNumberTransferCustomSipHeaderUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStatic, u.OfDynamic)
}
func (u *PhoneNumberTransferCustomSipHeaderUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PhoneNumberTransferCustomSipHeaderUnionParam) asAny() any {
	if !param.IsOmitted(u.OfStatic) {
		return u.OfStatic
	} else if !param.IsOmitted(u.OfDynamic) {
		return u.OfDynamic
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferCustomSipHeaderUnionParam) GetKey() *string {
	if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Key)
	} else if vt := u.OfDynamic; vt != nil {
		return (*string)(&vt.Key)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferCustomSipHeaderUnionParam) GetValue() *string {
	if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfDynamic; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferCustomSipHeaderUnionParam) GetType() *string {
	if vt := u.OfStatic; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDynamic; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[PhoneNumberTransferCustomSipHeaderUnionParam](
		"type",
		apijson.Discriminator[PhoneNumberTransferCustomSipHeaderStaticParam]("static"),
		apijson.Discriminator[PhoneNumberTransferCustomSipHeaderDynamicParam]("dynamic"),
	)
}

// Custom SIP header for phone transfers with a static (validated) value.
//
// The properties Key, Value are required.
type PhoneNumberTransferCustomSipHeaderStaticParam struct {
	// The SIP header name (e.g., 'X-Customer-ID')
	Key string `json:"key,required"`
	// The header value
	Value string `json:"value,required"`
	// Any of "static".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferCustomSipHeaderStaticParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferCustomSipHeaderStaticParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferCustomSipHeaderStaticParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberTransferCustomSipHeaderStaticParam](
		"type", "static",
	)
}

// Custom SIP header for phone transfers with a dynamic variable reference. The
// value is a variable name that will be resolved at runtime. Value is not
// validated here since it will be substituted with actual value later.
//
// The properties Key, Type, Value are required.
type PhoneNumberTransferCustomSipHeaderDynamicParam struct {
	// The SIP header name (e.g., 'X-Customer-ID')
	Key string `json:"key,required"`
	// The dynamic variable name to resolve
	Value string `json:"value,required"`
	// This field can be elided, and will marshal its zero value as "dynamic".
	Type constant.Dynamic `json:"type,required"`
	paramObj
}

func (r PhoneNumberTransferCustomSipHeaderDynamicParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferCustomSipHeaderDynamicParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferCustomSipHeaderDynamicParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PhoneNumberTransferTransferDestinationUnionParam struct {
	OfPhone                 *PhoneNumberTransferTransferDestinationPhoneParam                 `json:",omitzero,inline"`
	OfSipUri                *PhoneNumberTransferTransferDestinationSipUriParam                `json:",omitzero,inline"`
	OfPhoneDynamicVariable  *PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam  `json:",omitzero,inline"`
	OfSipUriDynamicVariable *PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam `json:",omitzero,inline"`
	paramUnion
}

func (u PhoneNumberTransferTransferDestinationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPhone, u.OfSipUri, u.OfPhoneDynamicVariable, u.OfSipUriDynamicVariable)
}
func (u *PhoneNumberTransferTransferDestinationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PhoneNumberTransferTransferDestinationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfPhone) {
		return u.OfPhone
	} else if !param.IsOmitted(u.OfSipUri) {
		return u.OfSipUri
	} else if !param.IsOmitted(u.OfPhoneDynamicVariable) {
		return u.OfPhoneDynamicVariable
	} else if !param.IsOmitted(u.OfSipUriDynamicVariable) {
		return u.OfSipUriDynamicVariable
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferTransferDestinationUnionParam) GetPhoneNumber() *string {
	if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.PhoneNumber)
	} else if vt := u.OfPhoneDynamicVariable; vt != nil {
		return (*string)(&vt.PhoneNumber)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferTransferDestinationUnionParam) GetType() *string {
	if vt := u.OfPhone; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSipUri; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPhoneDynamicVariable; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSipUriDynamicVariable; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PhoneNumberTransferTransferDestinationUnionParam) GetSipUri() *string {
	if vt := u.OfSipUri; vt != nil {
		return (*string)(&vt.SipUri)
	} else if vt := u.OfSipUriDynamicVariable; vt != nil {
		return (*string)(&vt.SipUri)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[PhoneNumberTransferTransferDestinationUnionParam](
		"type",
		apijson.Discriminator[PhoneNumberTransferTransferDestinationPhoneParam]("phone"),
		apijson.Discriminator[PhoneNumberTransferTransferDestinationSipUriParam]("sip_uri"),
		apijson.Discriminator[PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam]("phone_dynamic_variable"),
		apijson.Discriminator[PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam]("sip_uri_dynamic_variable"),
	)
}

// The property PhoneNumber is required.
type PhoneNumberTransferTransferDestinationPhoneParam struct {
	PhoneNumber string `json:"phone_number,required"`
	// Any of "phone".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferTransferDestinationPhoneParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferTransferDestinationPhoneParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferTransferDestinationPhoneParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberTransferTransferDestinationPhoneParam](
		"type", "phone",
	)
}

// The property SipUri is required.
type PhoneNumberTransferTransferDestinationSipUriParam struct {
	SipUri string `json:"sip_uri,required"`
	// Any of "sip_uri".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferTransferDestinationSipUriParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferTransferDestinationSipUriParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferTransferDestinationSipUriParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberTransferTransferDestinationSipUriParam](
		"type", "sip_uri",
	)
}

// The property PhoneNumber is required.
type PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam struct {
	PhoneNumber string `json:"phone_number,required"`
	// Any of "phone_dynamic_variable".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberTransferTransferDestinationPhoneDynamicVariableParam](
		"type", "phone_dynamic_variable",
	)
}

// The property SipUri is required.
type PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam struct {
	SipUri string `json:"sip_uri,required"`
	// Any of "sip_uri_dynamic_variable".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PhoneNumberTransferTransferDestinationSipUriDynamicVariableParam](
		"type", "sip_uri_dynamic_variable",
	)
}

// Allows the agent to play DTMF tones during a phone call.
//
// This tool can be used to interact with automated phone systems, such as
// navigating phone menus, entering extensions, or inputting numeric codes.
type PlayDtmfToolConfig struct {
	// Any of "play_keypad_touch_tone".
	SystemToolType PlayDtmfToolConfigSystemToolType `json:"system_tool_type"`
	// If true, send DTMF tones out-of-band using RFC 4733 (useful for SIP calls only).
	// If false, send DTMF as in-band audio tones (default, works for all call types).
	UseOutOfBandDtmf bool `json:"use_out_of_band_dtmf"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemToolType   respjson.Field
		UseOutOfBandDtmf respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PlayDtmfToolConfig) RawJSON() string { return r.JSON.raw }
func (r *PlayDtmfToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PlayDtmfToolConfig to a PlayDtmfToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PlayDtmfToolConfigParam.Overrides()
func (r PlayDtmfToolConfig) ToParam() PlayDtmfToolConfigParam {
	return param.Override[PlayDtmfToolConfigParam](json.RawMessage(r.RawJSON()))
}

type PlayDtmfToolConfigSystemToolType string

const (
	PlayDtmfToolConfigSystemToolTypePlayKeypadTouchTone PlayDtmfToolConfigSystemToolType = "play_keypad_touch_tone"
)

// Allows the agent to play DTMF tones during a phone call.
//
// This tool can be used to interact with automated phone systems, such as
// navigating phone menus, entering extensions, or inputting numeric codes.
type PlayDtmfToolConfigParam struct {
	// If true, send DTMF tones out-of-band using RFC 4733 (useful for SIP calls only).
	// If false, send DTMF as in-band audio tones (default, works for all call types).
	UseOutOfBandDtmf param.Opt[bool] `json:"use_out_of_band_dtmf,omitzero"`
	// Any of "play_keypad_touch_tone".
	SystemToolType PlayDtmfToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r PlayDtmfToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow PlayDtmfToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PlayDtmfToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryParamsJsonSchemaResp struct {
	Properties map[string]QueryParamsJsonSchemaPropertyResp `json:"properties,required"`
	Required   []string                                     `json:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Properties  respjson.Field
		Required    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryParamsJsonSchemaResp) RawJSON() string { return r.JSON.raw }
func (r *QueryParamsJsonSchemaResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this QueryParamsJsonSchemaResp to a QueryParamsJsonSchema.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// QueryParamsJsonSchema.Overrides()
func (r QueryParamsJsonSchemaResp) ToParam() QueryParamsJsonSchema {
	return param.Override[QueryParamsJsonSchema](json.RawMessage(r.RawJSON()))
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
type QueryParamsJsonSchemaPropertyResp struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,required"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue QueryParamsJsonSchemaPropertyConstantValueUnionResp `json:"constant_value"`
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
func (r QueryParamsJsonSchemaPropertyResp) RawJSON() string { return r.JSON.raw }
func (r *QueryParamsJsonSchemaPropertyResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QueryParamsJsonSchemaPropertyConstantValueUnionResp contains all possible
// properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type QueryParamsJsonSchemaPropertyConstantValueUnionResp struct {
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

func (u QueryParamsJsonSchemaPropertyConstantValueUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryParamsJsonSchemaPropertyConstantValueUnionResp) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QueryParamsJsonSchemaPropertyConstantValueUnionResp) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QueryParamsJsonSchemaPropertyConstantValueUnionResp) RawJSON() string { return u.JSON.raw }

func (r *QueryParamsJsonSchemaPropertyConstantValueUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Properties is required.
type QueryParamsJsonSchema struct {
	Properties map[string]QueryParamsJsonSchemaProperty `json:"properties,omitzero,required"`
	Required   []string                                 `json:"required,omitzero"`
	paramObj
}

func (r QueryParamsJsonSchema) MarshalJSON() (data []byte, err error) {
	type shadow QueryParamsJsonSchema
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryParamsJsonSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
//
// The property Type is required.
type QueryParamsJsonSchemaProperty struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,omitzero,required"`
	// The description of the property. When set, the LLM will provide the value based
	// on this description. Mutually exclusive with dynamic_variable,
	// is_system_provided, and constant_value.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the dynamic variable to use for this property's value. Mutually
	// exclusive with description, is_system_provided, and constant_value.
	DynamicVariable param.Opt[string] `json:"dynamic_variable,omitzero"`
	// If true, the value will be populated by the system at runtime. Used by API
	// Integration Webhook tools for templating. Mutually exclusive with description,
	// dynamic_variable, and constant_value.
	IsSystemProvided param.Opt[bool] `json:"is_system_provided,omitzero"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,omitzero"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue QueryParamsJsonSchemaPropertyConstantValueUnion `json:"constant_value,omitzero"`
	paramObj
}

func (r QueryParamsJsonSchemaProperty) MarshalJSON() (data []byte, err error) {
	type shadow QueryParamsJsonSchemaProperty
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueryParamsJsonSchemaProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[QueryParamsJsonSchemaProperty](
		"type", "boolean", "string", "integer", "number",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type QueryParamsJsonSchemaPropertyConstantValueUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u QueryParamsJsonSchemaPropertyConstantValueUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *QueryParamsJsonSchemaPropertyConstantValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *QueryParamsJsonSchemaPropertyConstantValueUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ResourceAccessInfo struct {
	// Email of the agent's creator
	CreatorEmail string `json:"creator_email,required"`
	// Name of the agent's creator
	CreatorName string `json:"creator_name,required"`
	// Whether the user making the request is the creator of the agent
	IsCreator bool `json:"is_creator,required"`
	// The role of the user making the request
	//
	// Any of "admin", "editor", "commenter", "viewer".
	Role ResourceAccessInfoRole `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatorEmail respjson.Field
		CreatorName  respjson.Field
		IsCreator    respjson.Field
		Role         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ResourceAccessInfo) RawJSON() string { return r.JSON.raw }
func (r *ResourceAccessInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The role of the user making the request
type ResourceAccessInfoRole string

const (
	ResourceAccessInfoRoleAdmin     ResourceAccessInfoRole = "admin"
	ResourceAccessInfoRoleEditor    ResourceAccessInfoRole = "editor"
	ResourceAccessInfoRoleCommenter ResourceAccessInfoRole = "commenter"
	ResourceAccessInfoRoleViewer    ResourceAccessInfoRole = "viewer"
)

// Allows the agent to explicitly skip its turn.
//
// This tool should be invoked by the LLM when the user indicates they would like
// to think or take a short pause before continuing the conversation—e.g. when they
// say: "Give me a second", "Let me think", or "One moment please". After calling
// this tool, the assistant should not speak until the user speaks again, or
// another normal turn-taking condition is met. The tool itself has no parameters
// and performs no side-effects other than informing the backend that the current
// turn generation is complete.
type SkipTurnToolConfig struct {
	// Any of "skip_turn".
	SystemToolType SkipTurnToolConfigSystemToolType `json:"system_tool_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemToolType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkipTurnToolConfig) RawJSON() string { return r.JSON.raw }
func (r *SkipTurnToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SkipTurnToolConfig to a SkipTurnToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SkipTurnToolConfigParam.Overrides()
func (r SkipTurnToolConfig) ToParam() SkipTurnToolConfigParam {
	return param.Override[SkipTurnToolConfigParam](json.RawMessage(r.RawJSON()))
}

type SkipTurnToolConfigSystemToolType string

const (
	SkipTurnToolConfigSystemToolTypeSkipTurn SkipTurnToolConfigSystemToolType = "skip_turn"
)

// Allows the agent to explicitly skip its turn.
//
// This tool should be invoked by the LLM when the user indicates they would like
// to think or take a short pause before continuing the conversation—e.g. when they
// say: "Give me a second", "Let me think", or "One moment please". After calling
// this tool, the assistant should not speak until the user speaks again, or
// another normal turn-taking condition is met. The tool itself has no parameters
// and performs no side-effects other than informing the backend that the current
// turn generation is complete.
type SkipTurnToolConfigParam struct {
	// Any of "skip_turn".
	SystemToolType SkipTurnToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r SkipTurnToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SkipTurnToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SkipTurnToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A system tool is a tool that is used to call a system method in the server
//
// The properties Name, Params are required.
type SystemToolConfigInputParam struct {
	Name string `json:"name,required"`
	// Allows the agent to explicitly skip its turn.
	//
	// This tool should be invoked by the LLM when the user indicates they would like
	// to think or take a short pause before continuing the conversation—e.g. when they
	// say: "Give me a second", "Let me think", or "One moment please". After calling
	// this tool, the assistant should not speak until the user speaks again, or
	// another normal turn-taking condition is met. The tool itself has no parameters
	// and performs no side-effects other than informing the backend that the current
	// turn generation is complete.
	Params SystemToolConfigInputParamsUnionParam `json:"params,omitzero,required"`
	// Description of when the tool should be used and what it does. Leave empty to use
	// the default description that's optimized for the specific tool type.
	Description param.Opt[string] `json:"description,omitzero"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// The maximum time in seconds to wait for the tool call to complete.
	ResponseTimeoutSecs param.Opt[int64] `json:"response_timeout_secs,omitzero"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	// The type of tool
	//
	// Any of "system".
	Type SystemToolConfigInputType `json:"type,omitzero"`
	paramObj
}

func (r SystemToolConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemToolConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemToolConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SystemToolConfigInputParamsUnionParam struct {
	OfEndCall             *EndCallToolConfigParam                           `json:",omitzero,inline"`
	OfLanguageDetection   *LanguageDetectionToolConfigParam                 `json:",omitzero,inline"`
	OfTransferToAgent     *TransferToAgentToolConfigParam                   `json:",omitzero,inline"`
	OfTransferToNumber    *SystemToolConfigInputParamsTransferToNumberParam `json:",omitzero,inline"`
	OfSkipTurn            *SkipTurnToolConfigParam                          `json:",omitzero,inline"`
	OfPlayKeypadTouchTone *PlayDtmfToolConfigParam                          `json:",omitzero,inline"`
	OfVoicemailDetection  *VoicemailDetectionToolConfigParam                `json:",omitzero,inline"`
	paramUnion
}

func (u SystemToolConfigInputParamsUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEndCall,
		u.OfLanguageDetection,
		u.OfTransferToAgent,
		u.OfTransferToNumber,
		u.OfSkipTurn,
		u.OfPlayKeypadTouchTone,
		u.OfVoicemailDetection)
}
func (u *SystemToolConfigInputParamsUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SystemToolConfigInputParamsUnionParam) asAny() any {
	if !param.IsOmitted(u.OfEndCall) {
		return u.OfEndCall
	} else if !param.IsOmitted(u.OfLanguageDetection) {
		return u.OfLanguageDetection
	} else if !param.IsOmitted(u.OfTransferToAgent) {
		return u.OfTransferToAgent
	} else if !param.IsOmitted(u.OfTransferToNumber) {
		return u.OfTransferToNumber
	} else if !param.IsOmitted(u.OfSkipTurn) {
		return u.OfSkipTurn
	} else if !param.IsOmitted(u.OfPlayKeypadTouchTone) {
		return u.OfPlayKeypadTouchTone
	} else if !param.IsOmitted(u.OfVoicemailDetection) {
		return u.OfVoicemailDetection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemToolConfigInputParamsUnionParam) GetEnableClientMessage() *bool {
	if vt := u.OfTransferToNumber; vt != nil && vt.EnableClientMessage.Valid() {
		return &vt.EnableClientMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemToolConfigInputParamsUnionParam) GetUseOutOfBandDtmf() *bool {
	if vt := u.OfPlayKeypadTouchTone; vt != nil && vt.UseOutOfBandDtmf.Valid() {
		return &vt.UseOutOfBandDtmf.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemToolConfigInputParamsUnionParam) GetVoicemailMessage() *string {
	if vt := u.OfVoicemailDetection; vt != nil && vt.VoicemailMessage.Valid() {
		return &vt.VoicemailMessage.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SystemToolConfigInputParamsUnionParam) GetSystemToolType() *string {
	if vt := u.OfEndCall; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfLanguageDetection; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfTransferToAgent; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfTransferToNumber; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfSkipTurn; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfPlayKeypadTouchTone; vt != nil {
		return (*string)(&vt.SystemToolType)
	} else if vt := u.OfVoicemailDetection; vt != nil {
		return (*string)(&vt.SystemToolType)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u SystemToolConfigInputParamsUnionParam) GetTransfers() (res systemToolConfigInputParamsUnionParamTransfers) {
	if vt := u.OfTransferToAgent; vt != nil {
		res.any = &vt.Transfers
	} else if vt := u.OfTransferToNumber; vt != nil {
		res.any = &vt.Transfers
	}
	return
}

// Can have the runtime types [_[]TransferToAgentToolConfigTransferParam],
// [_[]PhoneNumberTransferParam]
type systemToolConfigInputParamsUnionParamTransfers struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *[]elevenlabs.TransferToAgentToolConfigTransferParam:
//	case *[]elevenlabs.PhoneNumberTransferParam:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u systemToolConfigInputParamsUnionParamTransfers) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[SystemToolConfigInputParamsUnionParam](
		"system_tool_type",
		apijson.Discriminator[EndCallToolConfigParam]("end_call"),
		apijson.Discriminator[LanguageDetectionToolConfigParam]("language_detection"),
		apijson.Discriminator[TransferToAgentToolConfigParam]("transfer_to_agent"),
		apijson.Discriminator[SystemToolConfigInputParamsTransferToNumberParam]("transfer_to_number"),
		apijson.Discriminator[SkipTurnToolConfigParam]("skip_turn"),
		apijson.Discriminator[PlayDtmfToolConfigParam]("play_keypad_touch_tone"),
		apijson.Discriminator[VoicemailDetectionToolConfigParam]("voicemail_detection"),
	)
}

// The property Transfers is required.
type SystemToolConfigInputParamsTransferToNumberParam struct {
	Transfers []PhoneNumberTransferParam `json:"transfers,omitzero,required"`
	// Whether to play a message to the client while they wait for transfer. Defaults
	// to true for backward compatibility.
	EnableClientMessage param.Opt[bool] `json:"enable_client_message,omitzero"`
	// Any of "transfer_to_number".
	SystemToolType string `json:"system_tool_type,omitzero"`
	paramObj
}

func (r SystemToolConfigInputParamsTransferToNumberParam) MarshalJSON() (data []byte, err error) {
	type shadow SystemToolConfigInputParamsTransferToNumberParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SystemToolConfigInputParamsTransferToNumberParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SystemToolConfigInputParamsTransferToNumberParam](
		"system_tool_type", "transfer_to_number",
	)
}

// The type of tool
type SystemToolConfigInputType string

const (
	SystemToolConfigInputTypeSystem SystemToolConfigInputType = "system"
)

// A system tool is a tool that is used to call a system method in the server
type SystemToolConfigOutput struct {
	Name string `json:"name,required"`
	// Allows the agent to explicitly skip its turn.
	//
	// This tool should be invoked by the LLM when the user indicates they would like
	// to think or take a short pause before continuing the conversation—e.g. when they
	// say: "Give me a second", "Let me think", or "One moment please". After calling
	// this tool, the assistant should not speak until the user speaks again, or
	// another normal turn-taking condition is met. The tool itself has no parameters
	// and performs no side-effects other than informing the backend that the current
	// turn generation is complete.
	Params SystemToolConfigOutputParamsUnion `json:"params,required"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignment `json:"assignments"`
	// Description of when the tool should be used and what it does. Leave empty to use
	// the default description that's optimized for the specific tool type.
	Description string `json:"description"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions bool `json:"disable_interruptions"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech bool `json:"force_pre_tool_speech"`
	// The maximum time in seconds to wait for the tool call to complete.
	ResponseTimeoutSecs int64 `json:"response_timeout_secs"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,nullable"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior"`
	// The type of tool
	//
	// Any of "system".
	Type SystemToolConfigOutputType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                  respjson.Field
		Params                respjson.Field
		Assignments           respjson.Field
		Description           respjson.Field
		DisableInterruptions  respjson.Field
		ForcePreToolSpeech    respjson.Field
		ResponseTimeoutSecs   respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemToolConfigOutput) RawJSON() string { return r.JSON.raw }
func (r *SystemToolConfigOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SystemToolConfigOutputParamsUnion contains all possible properties and values
// from [EndCallToolConfig], [LanguageDetectionToolConfig],
// [TransferToAgentToolConfig], [SystemToolConfigOutputParamsTransferToNumber],
// [SkipTurnToolConfig], [PlayDtmfToolConfig], [VoicemailDetectionToolConfig].
//
// Use the [SystemToolConfigOutputParamsUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SystemToolConfigOutputParamsUnion struct {
	// Any of "end_call", "language_detection", "transfer_to_agent",
	// "transfer_to_number", "skip_turn", "play_keypad_touch_tone",
	// "voicemail_detection".
	SystemToolType string `json:"system_tool_type"`
	// This field is a union of [[]TransferToAgentToolConfigTransfer],
	// [[]PhoneNumberTransfer]
	Transfers SystemToolConfigOutputParamsUnionTransfers `json:"transfers"`
	// This field is from variant [SystemToolConfigOutputParamsTransferToNumber].
	EnableClientMessage bool `json:"enable_client_message"`
	// This field is from variant [PlayDtmfToolConfig].
	UseOutOfBandDtmf bool `json:"use_out_of_band_dtmf"`
	// This field is from variant [VoicemailDetectionToolConfig].
	VoicemailMessage string `json:"voicemail_message"`
	JSON             struct {
		SystemToolType      respjson.Field
		Transfers           respjson.Field
		EnableClientMessage respjson.Field
		UseOutOfBandDtmf    respjson.Field
		VoicemailMessage    respjson.Field
		raw                 string
	} `json:"-"`
}

// anySystemToolConfigOutputParams is implemented by each variant of
// [SystemToolConfigOutputParamsUnion] to add type safety for the return type of
// [SystemToolConfigOutputParamsUnion.AsAny]
type anySystemToolConfigOutputParams interface {
	implSystemToolConfigOutputParamsUnion()
}

func (EndCallToolConfig) implSystemToolConfigOutputParamsUnion()                            {}
func (LanguageDetectionToolConfig) implSystemToolConfigOutputParamsUnion()                  {}
func (TransferToAgentToolConfig) implSystemToolConfigOutputParamsUnion()                    {}
func (SystemToolConfigOutputParamsTransferToNumber) implSystemToolConfigOutputParamsUnion() {}
func (SkipTurnToolConfig) implSystemToolConfigOutputParamsUnion()                           {}
func (PlayDtmfToolConfig) implSystemToolConfigOutputParamsUnion()                           {}
func (VoicemailDetectionToolConfig) implSystemToolConfigOutputParamsUnion()                 {}

// Use the following switch statement to find the correct variant
//
//	switch variant := SystemToolConfigOutputParamsUnion.AsAny().(type) {
//	case elevenlabs.EndCallToolConfig:
//	case elevenlabs.LanguageDetectionToolConfig:
//	case elevenlabs.TransferToAgentToolConfig:
//	case elevenlabs.SystemToolConfigOutputParamsTransferToNumber:
//	case elevenlabs.SkipTurnToolConfig:
//	case elevenlabs.PlayDtmfToolConfig:
//	case elevenlabs.VoicemailDetectionToolConfig:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u SystemToolConfigOutputParamsUnion) AsAny() anySystemToolConfigOutputParams {
	switch u.SystemToolType {
	case "end_call":
		return u.AsEndCall()
	case "language_detection":
		return u.AsLanguageDetection()
	case "transfer_to_agent":
		return u.AsTransferToAgent()
	case "transfer_to_number":
		return u.AsTransferToNumber()
	case "skip_turn":
		return u.AsSkipTurn()
	case "play_keypad_touch_tone":
		return u.AsPlayKeypadTouchTone()
	case "voicemail_detection":
		return u.AsVoicemailDetection()
	}
	return nil
}

func (u SystemToolConfigOutputParamsUnion) AsEndCall() (v EndCallToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsLanguageDetection() (v LanguageDetectionToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsTransferToAgent() (v TransferToAgentToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsTransferToNumber() (v SystemToolConfigOutputParamsTransferToNumber) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsSkipTurn() (v SkipTurnToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsPlayKeypadTouchTone() (v PlayDtmfToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SystemToolConfigOutputParamsUnion) AsVoicemailDetection() (v VoicemailDetectionToolConfig) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SystemToolConfigOutputParamsUnion) RawJSON() string { return u.JSON.raw }

func (r *SystemToolConfigOutputParamsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SystemToolConfigOutputParamsUnionTransfers is an implicit subunion of
// [SystemToolConfigOutputParamsUnion]. SystemToolConfigOutputParamsUnionTransfers
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [SystemToolConfigOutputParamsUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfTransfers]
type SystemToolConfigOutputParamsUnionTransfers struct {
	// This field will be present if the value is a
	// [[]TransferToAgentToolConfigTransfer] instead of an object.
	OfTransfers []TransferToAgentToolConfigTransfer `json:",inline"`
	JSON        struct {
		OfTransfers respjson.Field
		raw         string
	} `json:"-"`
}

func (r *SystemToolConfigOutputParamsUnionTransfers) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SystemToolConfigOutputParamsTransferToNumber struct {
	Transfers []PhoneNumberTransfer `json:"transfers,required"`
	// Whether to play a message to the client while they wait for transfer. Defaults
	// to true for backward compatibility.
	EnableClientMessage bool `json:"enable_client_message"`
	// Any of "transfer_to_number".
	SystemToolType string `json:"system_tool_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transfers           respjson.Field
		EnableClientMessage respjson.Field
		SystemToolType      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SystemToolConfigOutputParamsTransferToNumber) RawJSON() string { return r.JSON.raw }
func (r *SystemToolConfigOutputParamsTransferToNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tool
type SystemToolConfigOutputType string

const (
	SystemToolConfigOutputTypeSystem SystemToolConfigOutputType = "system"
)

// The property ToolConfig is required.
type ToolRequestModelParam struct {
	// Configuration for the tool
	ToolConfig ToolRequestModelToolConfigUnionParam `json:"tool_config,omitzero,required"`
	paramObj
}

func (r ToolRequestModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ToolRequestModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ToolRequestModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ToolRequestModelToolConfigUnionParam struct {
	OfWebhook *WebhookToolConfigInputParam `json:",omitzero,inline"`
	OfClient  *ClientToolConfigInputParam  `json:",omitzero,inline"`
	OfSystem  *SystemToolConfigInputParam  `json:",omitzero,inline"`
	paramUnion
}

func (u ToolRequestModelToolConfigUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebhook, u.OfClient, u.OfSystem)
}
func (u *ToolRequestModelToolConfigUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ToolRequestModelToolConfigUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebhook) {
		return u.OfWebhook
	} else if !param.IsOmitted(u.OfClient) {
		return u.OfClient
	} else if !param.IsOmitted(u.OfSystem) {
		return u.OfSystem
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetAPISchema() *WebhookToolConfigInputAPISchemaParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.APISchema
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetExpectsResponse() *bool {
	if vt := u.OfClient; vt != nil && vt.ExpectsResponse.Valid() {
		return &vt.ExpectsResponse.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetParameters() *ObjectJsonSchemaPropertyInputParam {
	if vt := u.OfClient; vt != nil {
		return &vt.Parameters
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetParams() *SystemToolConfigInputParamsUnionParam {
	if vt := u.OfSystem; vt != nil {
		return &vt.Params
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetDescription() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Description)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.Description)
	} else if vt := u.OfSystem; vt != nil && vt.Description.Valid() {
		return &vt.Description.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetName() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetDisableInterruptions() *bool {
	if vt := u.OfWebhook; vt != nil && vt.DisableInterruptions.Valid() {
		return &vt.DisableInterruptions.Value
	} else if vt := u.OfClient; vt != nil && vt.DisableInterruptions.Valid() {
		return &vt.DisableInterruptions.Value
	} else if vt := u.OfSystem; vt != nil && vt.DisableInterruptions.Valid() {
		return &vt.DisableInterruptions.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetExecutionMode() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ExecutionMode)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.ExecutionMode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetForcePreToolSpeech() *bool {
	if vt := u.OfWebhook; vt != nil && vt.ForcePreToolSpeech.Valid() {
		return &vt.ForcePreToolSpeech.Value
	} else if vt := u.OfClient; vt != nil && vt.ForcePreToolSpeech.Valid() {
		return &vt.ForcePreToolSpeech.Value
	} else if vt := u.OfSystem; vt != nil && vt.ForcePreToolSpeech.Valid() {
		return &vt.ForcePreToolSpeech.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetResponseTimeoutSecs() *int64 {
	if vt := u.OfWebhook; vt != nil && vt.ResponseTimeoutSecs.Valid() {
		return &vt.ResponseTimeoutSecs.Value
	} else if vt := u.OfClient; vt != nil && vt.ResponseTimeoutSecs.Valid() {
		return &vt.ResponseTimeoutSecs.Value
	} else if vt := u.OfSystem; vt != nil && vt.ResponseTimeoutSecs.Valid() {
		return &vt.ResponseTimeoutSecs.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetToolCallSound() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ToolCallSound)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.ToolCallSound)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.ToolCallSound)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetToolCallSoundBehavior() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.ToolCallSoundBehavior)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.ToolCallSoundBehavior)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.ToolCallSoundBehavior)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetType() *string {
	if vt := u.OfWebhook; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfClient; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfSystem; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a pointer to the underlying variant's Assignments property, if present.
func (u ToolRequestModelToolConfigUnionParam) GetAssignments() []DynamicVariableAssignmentParam {
	if vt := u.OfWebhook; vt != nil {
		return vt.Assignments
	} else if vt := u.OfClient; vt != nil {
		return vt.Assignments
	} else if vt := u.OfSystem; vt != nil {
		return vt.Assignments
	}
	return nil
}

// Returns a pointer to the underlying variant's DynamicVariables property, if
// present.
func (u ToolRequestModelToolConfigUnionParam) GetDynamicVariables() *DynamicVariablesConfigParam {
	if vt := u.OfWebhook; vt != nil {
		return &vt.DynamicVariables
	} else if vt := u.OfClient; vt != nil {
		return &vt.DynamicVariables
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ToolRequestModelToolConfigUnionParam](
		"type",
		apijson.Discriminator[WebhookToolConfigInputParam]("webhook"),
		apijson.Discriminator[ClientToolConfigInputParam]("client"),
		apijson.Discriminator[SystemToolConfigInputParam]("system"),
	)
}

type ToolResponseModel struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// The type of tool
	ToolConfig ToolResponseModelToolConfigUnion `json:"tool_config,required"`
	UsageStats ToolResponseModelUsageStats      `json:"usage_stats,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccessInfo  respjson.Field
		ToolConfig  respjson.Field
		UsageStats  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ToolResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToolResponseModelToolConfigUnion contains all possible properties and values
// from [WebhookToolConfigOutput], [ClientToolConfigOutput],
// [SystemToolConfigOutput].
//
// Use the [ToolResponseModelToolConfigUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToolResponseModelToolConfigUnion struct {
	// This field is from variant [WebhookToolConfigOutput].
	APISchema            WebhookToolConfigOutputAPISchema `json:"api_schema"`
	Description          string                           `json:"description"`
	Name                 string                           `json:"name"`
	Assignments          []DynamicVariableAssignment      `json:"assignments"`
	DisableInterruptions bool                             `json:"disable_interruptions"`
	// This field is from variant [WebhookToolConfigOutput].
	DynamicVariables DynamicVariablesConfig `json:"dynamic_variables"`
	// This field is from variant [WebhookToolConfigOutput].
	ExecutionMode       ToolExecutionMode `json:"execution_mode"`
	ForcePreToolSpeech  bool              `json:"force_pre_tool_speech"`
	ResponseTimeoutSecs int64             `json:"response_timeout_secs"`
	// This field is from variant [WebhookToolConfigOutput].
	ToolCallSound ToolCallSoundType `json:"tool_call_sound"`
	// This field is from variant [WebhookToolConfigOutput].
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior"`
	// Any of "webhook", "client", "system".
	Type string `json:"type"`
	// This field is from variant [ClientToolConfigOutput].
	ExpectsResponse bool `json:"expects_response"`
	// This field is from variant [ClientToolConfigOutput].
	Parameters ObjectJsonSchemaPropertyOutput `json:"parameters"`
	// This field is from variant [SystemToolConfigOutput].
	Params SystemToolConfigOutputParamsUnion `json:"params"`
	JSON   struct {
		APISchema             respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Assignments           respjson.Field
		DisableInterruptions  respjson.Field
		DynamicVariables      respjson.Field
		ExecutionMode         respjson.Field
		ForcePreToolSpeech    respjson.Field
		ResponseTimeoutSecs   respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		Type                  respjson.Field
		ExpectsResponse       respjson.Field
		Parameters            respjson.Field
		Params                respjson.Field
		raw                   string
	} `json:"-"`
}

// anyToolResponseModelToolConfig is implemented by each variant of
// [ToolResponseModelToolConfigUnion] to add type safety for the return type of
// [ToolResponseModelToolConfigUnion.AsAny]
type anyToolResponseModelToolConfig interface {
	implToolResponseModelToolConfigUnion()
}

func (WebhookToolConfigOutput) implToolResponseModelToolConfigUnion() {}
func (ClientToolConfigOutput) implToolResponseModelToolConfigUnion()  {}
func (SystemToolConfigOutput) implToolResponseModelToolConfigUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToolResponseModelToolConfigUnion.AsAny().(type) {
//	case elevenlabs.WebhookToolConfigOutput:
//	case elevenlabs.ClientToolConfigOutput:
//	case elevenlabs.SystemToolConfigOutput:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToolResponseModelToolConfigUnion) AsAny() anyToolResponseModelToolConfig {
	switch u.Type {
	case "webhook":
		return u.AsWebhook()
	case "client":
		return u.AsClient()
	case "system":
		return u.AsSystem()
	}
	return nil
}

func (u ToolResponseModelToolConfigUnion) AsWebhook() (v WebhookToolConfigOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseModelToolConfigUnion) AsClient() (v ClientToolConfigOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResponseModelToolConfigUnion) AsSystem() (v SystemToolConfigOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolResponseModelToolConfigUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolResponseModelToolConfigUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolResponseModelUsageStats struct {
	AvgLatencySecs float64 `json:"avg_latency_secs,required"`
	// The total number of calls to the tool
	TotalCalls int64 `json:"total_calls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvgLatencySecs respjson.Field
		TotalCalls     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResponseModelUsageStats) RawJSON() string { return r.JSON.raw }
func (r *ToolResponseModelUsageStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferToAgentToolConfig struct {
	Transfers []TransferToAgentToolConfigTransfer `json:"transfers,required"`
	// Any of "transfer_to_agent".
	SystemToolType TransferToAgentToolConfigSystemToolType `json:"system_tool_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transfers      respjson.Field
		SystemToolType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferToAgentToolConfig) RawJSON() string { return r.JSON.raw }
func (r *TransferToAgentToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TransferToAgentToolConfig to a
// TransferToAgentToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TransferToAgentToolConfigParam.Overrides()
func (r TransferToAgentToolConfig) ToParam() TransferToAgentToolConfigParam {
	return param.Override[TransferToAgentToolConfigParam](json.RawMessage(r.RawJSON()))
}

type TransferToAgentToolConfigTransfer struct {
	AgentID                            string `json:"agent_id,required"`
	Condition                          string `json:"condition,required"`
	DelayMs                            int64  `json:"delay_ms"`
	EnableTransferredAgentFirstMessage bool   `json:"enable_transferred_agent_first_message"`
	TransferMessage                    string `json:"transfer_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID                            respjson.Field
		Condition                          respjson.Field
		DelayMs                            respjson.Field
		EnableTransferredAgentFirstMessage respjson.Field
		TransferMessage                    respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransferToAgentToolConfigTransfer) RawJSON() string { return r.JSON.raw }
func (r *TransferToAgentToolConfigTransfer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransferToAgentToolConfigSystemToolType string

const (
	TransferToAgentToolConfigSystemToolTypeTransferToAgent TransferToAgentToolConfigSystemToolType = "transfer_to_agent"
)

// The property Transfers is required.
type TransferToAgentToolConfigParam struct {
	Transfers []TransferToAgentToolConfigTransferParam `json:"transfers,omitzero,required"`
	// Any of "transfer_to_agent".
	SystemToolType TransferToAgentToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r TransferToAgentToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToAgentToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToAgentToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AgentID, Condition are required.
type TransferToAgentToolConfigTransferParam struct {
	AgentID                            string            `json:"agent_id,required"`
	Condition                          string            `json:"condition,required"`
	TransferMessage                    param.Opt[string] `json:"transfer_message,omitzero"`
	DelayMs                            param.Opt[int64]  `json:"delay_ms,omitzero"`
	EnableTransferredAgentFirstMessage param.Opt[bool]   `json:"enable_transferred_agent_first_message,omitzero"`
	paramObj
}

func (r TransferToAgentToolConfigTransferParam) MarshalJSON() (data []byte, err error) {
	type shadow TransferToAgentToolConfigTransferParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TransferToAgentToolConfigTransferParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Allows the agent to detect when a voicemail system is encountered.
//
// This tool should be invoked by the LLM when it detects that the call has been
// answered by a voicemail system rather than a human. If a voicemail message is
// configured, it will be played; otherwise the call will end immediately.
type VoicemailDetectionToolConfig struct {
	// Any of "voicemail_detection".
	SystemToolType VoicemailDetectionToolConfigSystemToolType `json:"system_tool_type"`
	// Optional message to leave on voicemail when detected. If not provided, the call
	// will end immediately when voicemail is detected. Supports dynamic variables
	// (e.g., {{system__time}}, {{system__call_duration_secs}}, {{custom_variable}}).
	VoicemailMessage string `json:"voicemail_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SystemToolType   respjson.Field
		VoicemailMessage respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicemailDetectionToolConfig) RawJSON() string { return r.JSON.raw }
func (r *VoicemailDetectionToolConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VoicemailDetectionToolConfig to a
// VoicemailDetectionToolConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VoicemailDetectionToolConfigParam.Overrides()
func (r VoicemailDetectionToolConfig) ToParam() VoicemailDetectionToolConfigParam {
	return param.Override[VoicemailDetectionToolConfigParam](json.RawMessage(r.RawJSON()))
}

type VoicemailDetectionToolConfigSystemToolType string

const (
	VoicemailDetectionToolConfigSystemToolTypeVoicemailDetection VoicemailDetectionToolConfigSystemToolType = "voicemail_detection"
)

// Allows the agent to detect when a voicemail system is encountered.
//
// This tool should be invoked by the LLM when it detects that the call has been
// answered by a voicemail system rather than a human. If a voicemail message is
// configured, it will be played; otherwise the call will end immediately.
type VoicemailDetectionToolConfigParam struct {
	// Optional message to leave on voicemail when detected. If not provided, the call
	// will end immediately when voicemail is detected. Supports dynamic variables
	// (e.g., {{system__time}}, {{system__call_duration_secs}}, {{custom_variable}}).
	VoicemailMessage param.Opt[string] `json:"voicemail_message,omitzero"`
	// Any of "voicemail_detection".
	SystemToolType VoicemailDetectionToolConfigSystemToolType `json:"system_tool_type,omitzero"`
	paramObj
}

func (r VoicemailDetectionToolConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow VoicemailDetectionToolConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicemailDetectionToolConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A webhook tool is a tool that calls an external webhook from our server
//
// The properties APISchema, Description, Name are required.
type WebhookToolConfigInputParam struct {
	// The schema for the outgoing webhoook, including parameters and URL specification
	APISchema WebhookToolConfigInputAPISchemaParam `json:"api_schema,omitzero,required"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description,required"`
	Name        string `json:"name,required"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// The maximum time in seconds to wait for the tool call to complete. Must be
	// between 5 and 120 seconds (inclusive).
	ResponseTimeoutSecs param.Opt[int64] `json:"response_timeout_secs,omitzero"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// Configuration for dynamic variables
	DynamicVariables DynamicVariablesConfigParam `json:"dynamic_variables,omitzero"`
	// Determines when and how the tool executes: 'immediate' executes the tool right
	// away when requested by the LLM, 'post_tool_speech' waits for the agent to finish
	// speaking before executing, 'async' runs the tool in the background without
	// blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	// The type of tool
	//
	// Any of "webhook".
	Type WebhookToolConfigInputType `json:"type,omitzero"`
	paramObj
}

func (r WebhookToolConfigInputParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolConfigInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolConfigInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schema for the outgoing webhoook, including parameters and URL specification
//
// The property URL is required.
type WebhookToolConfigInputAPISchemaParam struct {
	// The URL that the webhook will be sent to. May include path parameters, e.g.
	// https://example.com/agents/{agent_id}
	URL string `json:"url,required"`
	// Used to reference an auth connection from the workspace's auth connection store.
	AuthConnection AuthConnectionLocatorParam `json:"auth_connection,omitzero"`
	// Content type for the request body. Only applies to POST/PUT/PATCH requests.
	//
	// Any of "application/json", "application/x-www-form-urlencoded".
	ContentType string `json:"content_type,omitzero"`
	// The HTTP method to use for the webhook
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Schema for path parameters, if any. The keys should match the placeholders in
	// the URL.
	PathParamsSchema map[string]WebhookToolConfigInputAPISchemaPathParamsSchemaParam `json:"path_params_schema,omitzero"`
	// Schema for any query params, if any. These will be added to end of the URL as
	// query params. Note: properties in a query param must all be literal types
	QueryParamsSchema QueryParamsJsonSchema `json:"query_params_schema,omitzero"`
	// Schema for the body parameters, if any. Used for POST/PATCH/PUT requests. The
	// schema should be an object which will be sent as the json body
	RequestBodySchema ObjectJsonSchemaPropertyInputParam `json:"request_body_schema,omitzero"`
	// Headers that should be included in the request
	RequestHeaders map[string]WebhookToolConfigInputAPISchemaRequestHeaderUnionParam `json:"request_headers,omitzero"`
	paramObj
}

func (r WebhookToolConfigInputAPISchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolConfigInputAPISchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolConfigInputAPISchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolConfigInputAPISchemaParam](
		"content_type", "application/json", "application/x-www-form-urlencoded",
	)
	apijson.RegisterFieldValidator[WebhookToolConfigInputAPISchemaParam](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
//
// The property Type is required.
type WebhookToolConfigInputAPISchemaPathParamsSchemaParam struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,omitzero,required"`
	// The description of the property. When set, the LLM will provide the value based
	// on this description. Mutually exclusive with dynamic_variable,
	// is_system_provided, and constant_value.
	Description param.Opt[string] `json:"description,omitzero"`
	// The name of the dynamic variable to use for this property's value. Mutually
	// exclusive with description, is_system_provided, and constant_value.
	DynamicVariable param.Opt[string] `json:"dynamic_variable,omitzero"`
	// If true, the value will be populated by the system at runtime. Used by API
	// Integration Webhook tools for templating. Mutually exclusive with description,
	// dynamic_variable, and constant_value.
	IsSystemProvided param.Opt[bool] `json:"is_system_provided,omitzero"`
	// List of allowed string values for string type parameters
	Enum []string `json:"enum,omitzero"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam `json:"constant_value,omitzero"`
	paramObj
}

func (r WebhookToolConfigInputAPISchemaPathParamsSchemaParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolConfigInputAPISchemaPathParamsSchemaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolConfigInputAPISchemaPathParamsSchemaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebhookToolConfigInputAPISchemaPathParamsSchemaParam](
		"type", "boolean", "string", "integer", "number",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebhookToolConfigInputAPISchemaRequestHeaderUnionParam struct {
	OfString                param.Opt[string]                                                       `json:",omitzero,inline"`
	OfConvAISecretLocator   *ConvaiSecretLocatorParam                                               `json:",omitzero,inline"`
	OfConvAIDynamicVariable *WebhookToolConfigInputAPISchemaRequestHeaderConvAIDynamicVariableParam `json:",omitzero,inline"`
	paramUnion
}

func (u WebhookToolConfigInputAPISchemaRequestHeaderUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConvAISecretLocator, u.OfConvAIDynamicVariable)
}
func (u *WebhookToolConfigInputAPISchemaRequestHeaderUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WebhookToolConfigInputAPISchemaRequestHeaderUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfConvAISecretLocator) {
		return u.OfConvAISecretLocator
	} else if !param.IsOmitted(u.OfConvAIDynamicVariable) {
		return u.OfConvAIDynamicVariable
	}
	return nil
}

// Used to reference a dynamic variable.
//
// The property VariableName is required.
type WebhookToolConfigInputAPISchemaRequestHeaderConvAIDynamicVariableParam struct {
	VariableName string `json:"variable_name,required"`
	paramObj
}

func (r WebhookToolConfigInputAPISchemaRequestHeaderConvAIDynamicVariableParam) MarshalJSON() (data []byte, err error) {
	type shadow WebhookToolConfigInputAPISchemaRequestHeaderConvAIDynamicVariableParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebhookToolConfigInputAPISchemaRequestHeaderConvAIDynamicVariableParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tool
type WebhookToolConfigInputType string

const (
	WebhookToolConfigInputTypeWebhook WebhookToolConfigInputType = "webhook"
)

// A webhook tool is a tool that calls an external webhook from our server
type WebhookToolConfigOutput struct {
	// The schema for the outgoing webhoook, including parameters and URL specification
	APISchema WebhookToolConfigOutputAPISchema `json:"api_schema,required"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description,required"`
	Name        string `json:"name,required"`
	// Configuration for extracting values from tool responses and assigning them to
	// dynamic variables
	Assignments []DynamicVariableAssignment `json:"assignments"`
	// If true, the user will not be able to interrupt the agent while this tool is
	// running.
	DisableInterruptions bool `json:"disable_interruptions"`
	// Configuration for dynamic variables
	DynamicVariables DynamicVariablesConfig `json:"dynamic_variables"`
	// Determines when and how the tool executes: 'immediate' executes the tool right
	// away when requested by the LLM, 'post_tool_speech' waits for the agent to finish
	// speaking before executing, 'async' runs the tool in the background without
	// blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode"`
	// If true, the agent will speak before the tool call.
	ForcePreToolSpeech bool `json:"force_pre_tool_speech"`
	// The maximum time in seconds to wait for the tool call to complete. Must be
	// between 5 and 120 seconds (inclusive).
	ResponseTimeoutSecs int64 `json:"response_timeout_secs"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,nullable"`
	// Determines when the tool call sound should play. 'auto' only plays when there's
	// pre-tool speech, 'always' plays for every tool call.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior"`
	// The type of tool
	//
	// Any of "webhook".
	Type WebhookToolConfigOutputType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APISchema             respjson.Field
		Description           respjson.Field
		Name                  respjson.Field
		Assignments           respjson.Field
		DisableInterruptions  respjson.Field
		DynamicVariables      respjson.Field
		ExecutionMode         respjson.Field
		ForcePreToolSpeech    respjson.Field
		ResponseTimeoutSecs   respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookToolConfigOutput) RawJSON() string { return r.JSON.raw }
func (r *WebhookToolConfigOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The schema for the outgoing webhoook, including parameters and URL specification
type WebhookToolConfigOutputAPISchema struct {
	// The URL that the webhook will be sent to. May include path parameters, e.g.
	// https://example.com/agents/{agent_id}
	URL string `json:"url,required"`
	// Used to reference an auth connection from the workspace's auth connection store.
	AuthConnection AuthConnectionLocator `json:"auth_connection,nullable"`
	// Content type for the request body. Only applies to POST/PUT/PATCH requests.
	//
	// Any of "application/json", "application/x-www-form-urlencoded".
	ContentType string `json:"content_type"`
	// The HTTP method to use for the webhook
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method"`
	// Schema for path parameters, if any. The keys should match the placeholders in
	// the URL.
	PathParamsSchema map[string]WebhookToolConfigOutputAPISchemaPathParamsSchema `json:"path_params_schema"`
	// Schema for any query params, if any. These will be added to end of the URL as
	// query params. Note: properties in a query param must all be literal types
	QueryParamsSchema QueryParamsJsonSchemaResp `json:"query_params_schema,nullable"`
	// Schema for the body parameters, if any. Used for POST/PATCH/PUT requests. The
	// schema should be an object which will be sent as the json body
	RequestBodySchema ObjectJsonSchemaPropertyOutput `json:"request_body_schema,nullable"`
	// Headers that should be included in the request
	RequestHeaders map[string]WebhookToolConfigOutputAPISchemaRequestHeaderUnion `json:"request_headers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL               respjson.Field
		AuthConnection    respjson.Field
		ContentType       respjson.Field
		Method            respjson.Field
		PathParamsSchema  respjson.Field
		QueryParamsSchema respjson.Field
		RequestBodySchema respjson.Field
		RequestHeaders    respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookToolConfigOutputAPISchema) RawJSON() string { return r.JSON.raw }
func (r *WebhookToolConfigOutputAPISchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema property for literal JSON types. IMPORTANT: Only ONE of the following
// fields can be set: description (LLM provides value), dynamic_variable (value
// from variable), is_system_provided (system provides value), or constant_value
// (fixed value). These are mutually exclusive.
type WebhookToolConfigOutputAPISchemaPathParamsSchema struct {
	// Any of "boolean", "string", "integer", "number".
	Type string `json:"type,required"`
	// A constant value to use for this property. Mutually exclusive with description,
	// dynamic_variable, and is_system_provided.
	ConstantValue WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion `json:"constant_value"`
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
func (r WebhookToolConfigOutputAPISchemaPathParamsSchema) RawJSON() string { return r.JSON.raw }
func (r *WebhookToolConfigOutputAPISchemaPathParamsSchema) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion contains all
// possible properties and values from [string], [float64], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool]
type WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion struct {
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

func (u WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *WebhookToolConfigOutputAPISchemaPathParamsSchemaConstantValueUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebhookToolConfigOutputAPISchemaRequestHeaderUnion contains all possible
// properties and values from [string], [ConvaiSecretLocator],
// [WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type WebhookToolConfigOutputAPISchemaRequestHeaderUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [ConvaiSecretLocator].
	SecretID string `json:"secret_id"`
	// This field is from variant
	// [WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable].
	VariableName string `json:"variable_name"`
	JSON         struct {
		OfString     respjson.Field
		SecretID     respjson.Field
		VariableName respjson.Field
		raw          string
	} `json:"-"`
}

func (u WebhookToolConfigOutputAPISchemaRequestHeaderUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookToolConfigOutputAPISchemaRequestHeaderUnion) AsConvAISecretLocator() (v ConvaiSecretLocator) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebhookToolConfigOutputAPISchemaRequestHeaderUnion) AsConvAIDynamicVariable() (v WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebhookToolConfigOutputAPISchemaRequestHeaderUnion) RawJSON() string { return u.JSON.raw }

func (r *WebhookToolConfigOutputAPISchemaRequestHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Used to reference a dynamic variable.
type WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable struct {
	VariableName string `json:"variable_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VariableName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *WebhookToolConfigOutputAPISchemaRequestHeaderConvAIDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of tool
type WebhookToolConfigOutputType string

const (
	WebhookToolConfigOutputTypeWebhook WebhookToolConfigOutputType = "webhook"
)

type ConvaiToolListResponse struct {
	Tools []ToolResponseModel `json:"tools,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Tools       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiToolListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiToolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiToolDeleteResponse = any

type ConvaiToolListDependentAgentsResponse struct {
	Agents     []ConvaiToolListDependentAgentsResponseAgentUnion `json:"agents,required"`
	HasMore    bool                                              `json:"has_more,required"`
	NextCursor string                                            `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agents      respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiToolListDependentAgentsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiToolListDependentAgentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiToolListDependentAgentsResponseAgentUnion contains all possible properties
// and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiToolListDependentAgentsResponseAgentUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiToolListDependentAgentsResponseAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiToolListDependentAgentsResponseAgent is implemented by each variant of
// [ConvaiToolListDependentAgentsResponseAgentUnion] to add type safety for the
// return type of [ConvaiToolListDependentAgentsResponseAgentUnion.AsAny]
type anyConvaiToolListDependentAgentsResponseAgent interface {
	implConvaiToolListDependentAgentsResponseAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiToolListDependentAgentsResponseAgentUnion() {}
func (DependentUnknownAgentIdentifier) implConvaiToolListDependentAgentsResponseAgentUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiToolListDependentAgentsResponseAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiToolListDependentAgentsResponseAgentUnion) AsAny() anyConvaiToolListDependentAgentsResponseAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiToolListDependentAgentsResponseAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiToolListDependentAgentsResponseAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiToolListDependentAgentsResponseAgentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiToolListDependentAgentsResponseAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiToolNewParams struct {
	// A tool that an agent can provide to LLM.
	ToolRequestModel ToolRequestModelParam
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiToolNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ToolRequestModel)
}
func (r *ConvaiToolNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ToolRequestModel)
}

type ConvaiToolGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiToolUpdateParams struct {
	// A tool that an agent can provide to LLM.
	ToolRequestModel ToolRequestModelParam
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiToolUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ToolRequestModel)
}
func (r *ConvaiToolUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ToolRequestModel)
}

type ConvaiToolListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiToolDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiToolListDependentAgentsParams struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// How many documents to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiToolListDependentAgentsParams]'s query parameters as
// `url.Values`.
func (r ConvaiToolListDependentAgentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
