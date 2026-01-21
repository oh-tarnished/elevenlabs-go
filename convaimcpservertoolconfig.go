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
)

// ConvaiMcpServerToolConfigService contains methods and other services that help
// with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiMcpServerToolConfigService] method instead.
type ConvaiMcpServerToolConfigService struct {
	Options []option.RequestOption
}

// NewConvaiMcpServerToolConfigService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiMcpServerToolConfigService(opts ...option.RequestOption) (r ConvaiMcpServerToolConfigService) {
	r = ConvaiMcpServerToolConfigService{}
	r.Options = opts
	return
}

// Create configuration overrides for a specific MCP tool.
func (r *ConvaiMcpServerToolConfigService) New(ctx context.Context, mcpServerID string, params ConvaiMcpServerToolConfigNewParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-configs", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve configuration overrides for a specific MCP tool.
func (r *ConvaiMcpServerToolConfigService) Get(ctx context.Context, toolName string, params ConvaiMcpServerToolConfigGetParams, opts ...option.RequestOption) (res *McpToolConfigOverride, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.McpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	if toolName == "" {
		err = errors.New("missing required tool_name parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-configs/%s", params.McpServerID, toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update configuration overrides for a specific MCP tool.
func (r *ConvaiMcpServerToolConfigService) Update(ctx context.Context, toolName string, params ConvaiMcpServerToolConfigUpdateParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.McpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	if toolName == "" {
		err = errors.New("missing required tool_name parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-configs/%s", params.McpServerID, toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Remove configuration overrides for a specific MCP tool.
func (r *ConvaiMcpServerToolConfigService) Delete(ctx context.Context, toolName string, params ConvaiMcpServerToolConfigDeleteParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.McpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	if toolName == "" {
		err = errors.New("missing required tool_name parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-configs/%s", params.McpServerID, toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Configuration for extracting values from tool responses and assigning them to
// dynamic variables.
type DynamicVariableAssignment struct {
	// The name of the dynamic variable to assign the extracted value to
	DynamicVariable string `json:"dynamic_variable,required"`
	// Dot notation path to extract the value from the source (e.g., 'user.name' or
	// 'data.0.id')
	ValuePath string `json:"value_path,required"`
	// The source to extract the value from. Currently only 'response' is supported.
	//
	// Any of "response".
	Source DynamicVariableAssignmentSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DynamicVariable respjson.Field
		ValuePath       respjson.Field
		Source          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DynamicVariableAssignment) RawJSON() string { return r.JSON.raw }
func (r *DynamicVariableAssignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DynamicVariableAssignment to a
// DynamicVariableAssignmentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DynamicVariableAssignmentParam.Overrides()
func (r DynamicVariableAssignment) ToParam() DynamicVariableAssignmentParam {
	return param.Override[DynamicVariableAssignmentParam](json.RawMessage(r.RawJSON()))
}

// The source to extract the value from. Currently only 'response' is supported.
type DynamicVariableAssignmentSource string

const (
	DynamicVariableAssignmentSourceResponse DynamicVariableAssignmentSource = "response"
)

// Configuration for extracting values from tool responses and assigning them to
// dynamic variables.
//
// The properties DynamicVariable, ValuePath are required.
type DynamicVariableAssignmentParam struct {
	// The name of the dynamic variable to assign the extracted value to
	DynamicVariable string `json:"dynamic_variable,required"`
	// Dot notation path to extract the value from the source (e.g., 'user.name' or
	// 'data.0.id')
	ValuePath string `json:"value_path,required"`
	// The source to extract the value from. Currently only 'response' is supported.
	//
	// Any of "response".
	Source DynamicVariableAssignmentSource `json:"source,omitzero"`
	paramObj
}

func (r DynamicVariableAssignmentParam) MarshalJSON() (data []byte, err error) {
	type shadow DynamicVariableAssignmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DynamicVariableAssignmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type McpToolConfigOverride struct {
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// Dynamic variable assignments for this MCP tool
	Assignments []DynamicVariableAssignment `json:"assignments"`
	// If set, overrides the server's disable_interruptions setting for this tool
	DisableInterruptions bool `json:"disable_interruptions,nullable"`
	// If set, overrides the server's execution_mode setting for this tool
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,nullable"`
	// If set, overrides the server's force_pre_tool_speech setting for this tool
	ForcePreToolSpeech bool `json:"force_pre_tool_speech,nullable"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,nullable"`
	// Determines how the tool call sound should be played.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolName              respjson.Field
		Assignments           respjson.Field
		DisableInterruptions  respjson.Field
		ExecutionMode         respjson.Field
		ForcePreToolSpeech    respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpToolConfigOverride) RawJSON() string { return r.JSON.raw }
func (r *McpToolConfigOverride) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this McpToolConfigOverride to a McpToolConfigOverrideParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// McpToolConfigOverrideParam.Overrides()
func (r McpToolConfigOverride) ToParam() McpToolConfigOverrideParam {
	return param.Override[McpToolConfigOverrideParam](json.RawMessage(r.RawJSON()))
}

// The property ToolName is required.
type McpToolConfigOverrideParam struct {
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// If set, overrides the server's disable_interruptions setting for this tool
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If set, overrides the server's force_pre_tool_speech setting for this tool
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// Dynamic variable assignments for this MCP tool
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// If set, overrides the server's execution_mode setting for this tool
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines how the tool call sound should be played.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	paramObj
}

func (r McpToolConfigOverrideParam) MarshalJSON() (data []byte, err error) {
	type shadow McpToolConfigOverrideParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *McpToolConfigOverrideParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerToolConfigNewParams struct {
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// If set, overrides the server's disable_interruptions setting for this tool
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If set, overrides the server's force_pre_tool_speech setting for this tool
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Dynamic variable assignments for this MCP tool
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// If set, overrides the server's execution_mode setting for this tool
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines how the tool call sound should be played.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	paramObj
}

func (r ConvaiMcpServerToolConfigNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerToolConfigNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerToolConfigNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerToolConfigGetParams struct {
	// ID of the MCP Server.
	McpServerID string `path:"mcp_server_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiMcpServerToolConfigUpdateParams struct {
	// ID of the MCP Server.
	McpServerID string `path:"mcp_server_id,required" json:"-"`
	// If set, overrides the server's disable_interruptions setting for this tool
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If set, overrides the server's force_pre_tool_speech setting for this tool
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Dynamic variable assignments for this MCP tool
	Assignments []DynamicVariableAssignmentParam `json:"assignments,omitzero"`
	// If set, overrides the server's execution_mode setting for this tool
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines how the tool call sound should be played.
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	paramObj
}

func (r ConvaiMcpServerToolConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerToolConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerToolConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerToolConfigDeleteParams struct {
	// ID of the MCP Server.
	McpServerID string `path:"mcp_server_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
