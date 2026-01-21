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

// ConvaiMcpServerService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiMcpServerService] method instead.
type ConvaiMcpServerService struct {
	Options       []option.RequestOption
	ToolApprovals ConvaiMcpServerToolApprovalService
	ToolConfigs   ConvaiMcpServerToolConfigService
}

// NewConvaiMcpServerService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiMcpServerService(opts ...option.RequestOption) (r ConvaiMcpServerService) {
	r = ConvaiMcpServerService{}
	r.Options = opts
	r.ToolApprovals = NewConvaiMcpServerToolApprovalService(opts...)
	r.ToolConfigs = NewConvaiMcpServerToolConfigService(opts...)
	return
}

// Create a new MCP server configuration in the workspace.
func (r *ConvaiMcpServerService) New(ctx context.Context, params ConvaiMcpServerNewParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/mcp-servers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a specific MCP server configuration from the workspace.
func (r *ConvaiMcpServerService) Get(ctx context.Context, mcpServerID string, query ConvaiMcpServerGetParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the configuration settings for an MCP server.
func (r *ConvaiMcpServerService) Update(ctx context.Context, mcpServerID string, params ConvaiMcpServerUpdateParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve all MCP server configurations available in the workspace.
func (r *ConvaiMcpServerService) List(ctx context.Context, query ConvaiMcpServerListParams, opts ...option.RequestOption) (res *ConvaiMcpServerListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/mcp-servers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a specific MCP server configuration from the workspace.
func (r *ConvaiMcpServerService) Delete(ctx context.Context, mcpServerID string, body ConvaiMcpServerDeleteParams, opts ...option.RequestOption) (res *ConvaiMcpServerDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Retrieve all tools available for a specific MCP server configuration.
func (r *ConvaiMcpServerService) ListTools(ctx context.Context, mcpServerID string, query ConvaiMcpServerListToolsParams, opts ...option.RequestOption) (res *ConvaiMcpServerListToolsResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tools", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the approval policy configuration for an MCP server. DEPRECATED: Use
// PATCH /mcp-servers/{id} endpoint instead.
//
// Deprecated: deprecated
func (r *ConvaiMcpServerService) UpdateApprovalPolicy(ctx context.Context, mcpServerID string, params ConvaiMcpServerUpdateApprovalPolicyParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/approval-policy", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Used to reference a secret from the agent's secret store.
type ConvaiSecretLocator struct {
	SecretID string `json:"secret_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SecretID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretLocator) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConvaiSecretLocator to a ConvaiSecretLocatorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConvaiSecretLocatorParam.Overrides()
func (r ConvaiSecretLocator) ToParam() ConvaiSecretLocatorParam {
	return param.Override[ConvaiSecretLocatorParam](json.RawMessage(r.RawJSON()))
}

// Used to reference a secret from the agent's secret store.
//
// The property SecretID is required.
type ConvaiSecretLocatorParam struct {
	SecretID string `json:"secret_id,required"`
	paramObj
}

func (r ConvaiSecretLocatorParam) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSecretLocatorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSecretLocatorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// User-specific secret model that are not shared with other users in a workspace.
type ConvaiUserSecretDBModel struct {
	ID             string `json:"id,required"`
	EncryptedValue string `json:"encrypted_value,required"`
	Name           string `json:"name,required"`
	Nonce          string `json:"nonce,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		EncryptedValue respjson.Field
		Name           respjson.Field
		Nonce          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiUserSecretDBModel) RawJSON() string { return r.JSON.raw }
func (r *ConvaiUserSecretDBModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConvaiUserSecretDBModel to a ConvaiUserSecretDBModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConvaiUserSecretDBModelParam.Overrides()
func (r ConvaiUserSecretDBModel) ToParam() ConvaiUserSecretDBModelParam {
	return param.Override[ConvaiUserSecretDBModelParam](json.RawMessage(r.RawJSON()))
}

// User-specific secret model that are not shared with other users in a workspace.
//
// The properties ID, EncryptedValue, Name, Nonce are required.
type ConvaiUserSecretDBModelParam struct {
	ID             string `json:"id,required"`
	EncryptedValue string `json:"encrypted_value,required"`
	Name           string `json:"name,required"`
	Nonce          string `json:"nonce,required"`
	paramObj
}

func (r ConvaiUserSecretDBModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiUserSecretDBModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiUserSecretDBModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Defines the MCP server-level approval policy for tool execution.
type McpApprovalPolicy string

const (
	McpApprovalPolicyAutoApproveAll         McpApprovalPolicy = "auto_approve_all"
	McpApprovalPolicyRequireApprovalAll     McpApprovalPolicy = "require_approval_all"
	McpApprovalPolicyRequireApprovalPerTool McpApprovalPolicy = "require_approval_per_tool"
)

// Response model representing an MCP Server configuration.
type McpServerResponseModel struct {
	ID     string                       `json:"id,required"`
	Config McpServerResponseModelConfig `json:"config,required"`
	// The metadata of the MCP Server
	Metadata McpServerResponseModelMetadata `json:"metadata,required"`
	// The access information of the MCP Server
	AccessInfo ResourceAccessInfo `json:"access_info,nullable"`
	// List of agents that depend on this MCP Server.
	DependentAgents []McpServerResponseModelDependentAgentUnion `json:"dependent_agents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Config          respjson.Field
		Metadata        respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpServerResponseModel) RawJSON() string { return r.JSON.raw }
func (r *McpServerResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type McpServerResponseModelConfig struct {
	Name string `json:"name,required"`
	// The URL of the MCP server, if this contains a secret please store as a workspace
	// secret, otherwise store as a plain string. Must use https
	URL McpServerResponseModelConfigURLUnion `json:"url,required"`
	// Defines the MCP server-level approval policy for tool execution.
	//
	// Any of "auto_approve_all", "require_approval_all", "require_approval_per_tool".
	ApprovalPolicy McpApprovalPolicy `json:"approval_policy"`
	Description    string            `json:"description"`
	// Whether to disable HTTP compression for this MCP server. Enable this if the
	// server does not support compressed responses.
	DisableCompression bool `json:"disable_compression"`
	// If true, the user will not be able to interrupt the agent while any tool from
	// this MCP server is running.
	DisableInterruptions bool `json:"disable_interruptions"`
	// Determines when and how all tools from this MCP server execute: 'immediate'
	// executes the tool right away when requested by the LLM, 'post_tool_speech' waits
	// for the agent to finish speaking before executing, 'async' runs the tool in the
	// background without blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode"`
	// If true, all tools from this MCP server will require pre-tool execution speech
	ForcePreToolSpeech bool `json:"force_pre_tool_speech"`
	// The headers included in the request
	RequestHeaders map[string]McpServerResponseModelConfigRequestHeaderUnion `json:"request_headers"`
	// The secret token (Authorization header) stored as a workspace secret or in-place
	// secret
	SecretToken McpServerResponseModelConfigSecretTokenUnion `json:"secret_token,nullable"`
	// List of tool approval hashes for per-tool approval when approval_policy is
	// REQUIRE_APPROVAL_PER_TOOL
	ToolApprovalHashes []McpToolApprovalHash `json:"tool_approval_hashes"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,nullable"`
	// Determines when the tool call sound should play for all tools from this MCP
	// server
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior"`
	// List of per-tool configuration overrides that override the server-level defaults
	// for specific tools
	ToolConfigOverrides []McpToolConfigOverride `json:"tool_config_overrides"`
	// The transport type used to connect to the MCP server
	//
	// Any of "SSE", "STREAMABLE_HTTP".
	Transport McpServerTransport `json:"transport"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                  respjson.Field
		URL                   respjson.Field
		ApprovalPolicy        respjson.Field
		Description           respjson.Field
		DisableCompression    respjson.Field
		DisableInterruptions  respjson.Field
		ExecutionMode         respjson.Field
		ForcePreToolSpeech    respjson.Field
		RequestHeaders        respjson.Field
		SecretToken           respjson.Field
		ToolApprovalHashes    respjson.Field
		ToolCallSound         respjson.Field
		ToolCallSoundBehavior respjson.Field
		ToolConfigOverrides   respjson.Field
		Transport             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpServerResponseModelConfig) RawJSON() string { return r.JSON.raw }
func (r *McpServerResponseModelConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// McpServerResponseModelConfigURLUnion contains all possible properties and values
// from [string], [ConvaiSecretLocator].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type McpServerResponseModelConfigURLUnion struct {
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

func (u McpServerResponseModelConfigURLUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u McpServerResponseModelConfigURLUnion) AsConvAISecretLocator() (v ConvaiSecretLocator) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u McpServerResponseModelConfigURLUnion) RawJSON() string { return u.JSON.raw }

func (r *McpServerResponseModelConfigURLUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// McpServerResponseModelConfigRequestHeaderUnion contains all possible properties
// and values from [string], [ConvaiSecretLocator],
// [McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type McpServerResponseModelConfigRequestHeaderUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [ConvaiSecretLocator].
	SecretID string `json:"secret_id"`
	// This field is from variant
	// [McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable].
	VariableName string `json:"variable_name"`
	JSON         struct {
		OfString     respjson.Field
		SecretID     respjson.Field
		VariableName respjson.Field
		raw          string
	} `json:"-"`
}

func (u McpServerResponseModelConfigRequestHeaderUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u McpServerResponseModelConfigRequestHeaderUnion) AsConvAISecretLocator() (v ConvaiSecretLocator) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u McpServerResponseModelConfigRequestHeaderUnion) AsConvAIDynamicVariable() (v McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u McpServerResponseModelConfigRequestHeaderUnion) RawJSON() string { return u.JSON.raw }

func (r *McpServerResponseModelConfigRequestHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Used to reference a dynamic variable.
type McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable struct {
	VariableName string `json:"variable_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VariableName respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable) RawJSON() string {
	return r.JSON.raw
}
func (r *McpServerResponseModelConfigRequestHeaderConvAIDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// McpServerResponseModelConfigSecretTokenUnion contains all possible properties
// and values from [ConvaiSecretLocator], [ConvaiUserSecretDBModel].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type McpServerResponseModelConfigSecretTokenUnion struct {
	// This field is from variant [ConvaiSecretLocator].
	SecretID string `json:"secret_id"`
	// This field is from variant [ConvaiUserSecretDBModel].
	ID string `json:"id"`
	// This field is from variant [ConvaiUserSecretDBModel].
	EncryptedValue string `json:"encrypted_value"`
	// This field is from variant [ConvaiUserSecretDBModel].
	Name string `json:"name"`
	// This field is from variant [ConvaiUserSecretDBModel].
	Nonce string `json:"nonce"`
	JSON  struct {
		SecretID       respjson.Field
		ID             respjson.Field
		EncryptedValue respjson.Field
		Name           respjson.Field
		Nonce          respjson.Field
		raw            string
	} `json:"-"`
}

func (u McpServerResponseModelConfigSecretTokenUnion) AsConvAISecretLocator() (v ConvaiSecretLocator) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u McpServerResponseModelConfigSecretTokenUnion) AsConvAIUserSecretDBModel() (v ConvaiUserSecretDBModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u McpServerResponseModelConfigSecretTokenUnion) RawJSON() string { return u.JSON.raw }

func (r *McpServerResponseModelConfigSecretTokenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The metadata of the MCP Server
type McpServerResponseModelMetadata struct {
	CreatedAt   int64  `json:"created_at,required"`
	OwnerUserID string `json:"owner_user_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		OwnerUserID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpServerResponseModelMetadata) RawJSON() string { return r.JSON.raw }
func (r *McpServerResponseModelMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// McpServerResponseModelDependentAgentUnion contains all possible properties and
// values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [McpServerResponseModelDependentAgentUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type McpServerResponseModelDependentAgentUnion struct {
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

// anyMcpServerResponseModelDependentAgent is implemented by each variant of
// [McpServerResponseModelDependentAgentUnion] to add type safety for the return
// type of [McpServerResponseModelDependentAgentUnion.AsAny]
type anyMcpServerResponseModelDependentAgent interface {
	implMcpServerResponseModelDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implMcpServerResponseModelDependentAgentUnion() {}
func (DependentUnknownAgentIdentifier) implMcpServerResponseModelDependentAgentUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := McpServerResponseModelDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u McpServerResponseModelDependentAgentUnion) AsAny() anyMcpServerResponseModelDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u McpServerResponseModelDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u McpServerResponseModelDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u McpServerResponseModelDependentAgentUnion) RawJSON() string { return u.JSON.raw }

func (r *McpServerResponseModelDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Supported MCP server transport types.
type McpServerTransport string

const (
	McpServerTransportSse            McpServerTransport = "SSE"
	McpServerTransportStreamableHTTP McpServerTransport = "STREAMABLE_HTTP"
)

// Model for storing tool approval hashes for per-tool approval.
type McpToolApprovalHash struct {
	// SHA256 hash of the tool's parameters and description
	ToolHash string `json:"tool_hash,required"`
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// The approval policy for this tool
	//
	// Any of "auto_approved", "requires_approval".
	ApprovalPolicy McpToolApprovalPolicy `json:"approval_policy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolHash       respjson.Field
		ToolName       respjson.Field
		ApprovalPolicy respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpToolApprovalHash) RawJSON() string { return r.JSON.raw }
func (r *McpToolApprovalHash) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this McpToolApprovalHash to a McpToolApprovalHashParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// McpToolApprovalHashParam.Overrides()
func (r McpToolApprovalHash) ToParam() McpToolApprovalHashParam {
	return param.Override[McpToolApprovalHashParam](json.RawMessage(r.RawJSON()))
}

// Model for storing tool approval hashes for per-tool approval.
//
// The properties ToolHash, ToolName are required.
type McpToolApprovalHashParam struct {
	// SHA256 hash of the tool's parameters and description
	ToolHash string `json:"tool_hash,required"`
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// The approval policy for this tool
	//
	// Any of "auto_approved", "requires_approval".
	ApprovalPolicy McpToolApprovalPolicy `json:"approval_policy,omitzero"`
	paramObj
}

func (r McpToolApprovalHashParam) MarshalJSON() (data []byte, err error) {
	type shadow McpToolApprovalHashParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *McpToolApprovalHashParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determines how the tool call sound should be played.
type ToolCallSoundBehavior string

const (
	ToolCallSoundBehaviorAuto   ToolCallSoundBehavior = "auto"
	ToolCallSoundBehaviorAlways ToolCallSoundBehavior = "always"
)

// Predefined tool call sound types.
type ToolCallSoundType string

const (
	ToolCallSoundTypeTyping    ToolCallSoundType = "typing"
	ToolCallSoundTypeElevator1 ToolCallSoundType = "elevator1"
	ToolCallSoundTypeElevator2 ToolCallSoundType = "elevator2"
	ToolCallSoundTypeElevator3 ToolCallSoundType = "elevator3"
	ToolCallSoundTypeElevator4 ToolCallSoundType = "elevator4"
)

type ToolExecutionMode string

const (
	ToolExecutionModeImmediate      ToolExecutionMode = "immediate"
	ToolExecutionModePostToolSpeech ToolExecutionMode = "post_tool_speech"
	ToolExecutionModeAsync          ToolExecutionMode = "async"
)

// Response model for a list of MCP Server configurations.
type ConvaiMcpServerListResponse struct {
	McpServers []McpServerResponseModel `json:"mcp_servers,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		McpServers  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiMcpServerListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiMcpServerListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerDeleteResponse = any

// Response model for testing tools available on an MCP server.
type ConvaiMcpServerListToolsResponse struct {
	// Indicates if the operation was successful.
	Success bool `json:"success,required"`
	// A list of tools available on the MCP server.
	Tools []ConvaiMcpServerListToolsResponseTool `json:"tools,required"`
	// Error message if the operation was not successful.
	ErrorMessage string `json:"error_message,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success      respjson.Field
		Tools        respjson.Field
		ErrorMessage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiMcpServerListToolsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiMcpServerListToolsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Definition for a tool the client can call.
type ConvaiMcpServerListToolsResponseTool struct {
	InputSchema any    `json:"inputSchema,required"`
	Name        string `json:"name,required"`
	Meta        any    `json:"_meta,nullable"`
	// Additional properties describing a Tool to clients.
	//
	// NOTE: all properties in ToolAnnotations are **hints**. They are not guaranteed
	// to provide a faithful description of tool behavior (including descriptive
	// properties like `title`).
	//
	// Clients should never make tool use decisions based on ToolAnnotations received
	// from untrusted servers.
	Annotations  ConvaiMcpServerListToolsResponseToolAnnotations `json:"annotations,nullable"`
	Description  string                                          `json:"description,nullable"`
	OutputSchema any                                             `json:"outputSchema,nullable"`
	Title        string                                          `json:"title,nullable"`
	ExtraFields  map[string]any                                  `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputSchema  respjson.Field
		Name         respjson.Field
		Meta         respjson.Field
		Annotations  respjson.Field
		Description  respjson.Field
		OutputSchema respjson.Field
		Title        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiMcpServerListToolsResponseTool) RawJSON() string { return r.JSON.raw }
func (r *ConvaiMcpServerListToolsResponseTool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Additional properties describing a Tool to clients.
//
// NOTE: all properties in ToolAnnotations are **hints**. They are not guaranteed
// to provide a faithful description of tool behavior (including descriptive
// properties like `title`).
//
// Clients should never make tool use decisions based on ToolAnnotations received
// from untrusted servers.
type ConvaiMcpServerListToolsResponseToolAnnotations struct {
	DestructiveHint bool           `json:"destructiveHint,nullable"`
	IdempotentHint  bool           `json:"idempotentHint,nullable"`
	OpenWorldHint   bool           `json:"openWorldHint,nullable"`
	ReadOnlyHint    bool           `json:"readOnlyHint,nullable"`
	Title           string         `json:"title,nullable"`
	ExtraFields     map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DestructiveHint respjson.Field
		IdempotentHint  respjson.Field
		OpenWorldHint   respjson.Field
		ReadOnlyHint    respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiMcpServerListToolsResponseToolAnnotations) RawJSON() string { return r.JSON.raw }
func (r *ConvaiMcpServerListToolsResponseToolAnnotations) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerNewParams struct {
	// Configuration details for the MCP Server.
	Config ConvaiMcpServerNewParamsConfig `json:"config,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiMcpServerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration details for the MCP Server.
//
// The properties Name, URL are required.
type ConvaiMcpServerNewParamsConfig struct {
	Name string `json:"name,required"`
	// The URL of the MCP server, if this contains a secret please store as a workspace
	// secret, otherwise store as a plain string. Must use https
	URL         ConvaiMcpServerNewParamsConfigURLUnion `json:"url,omitzero,required"`
	Description param.Opt[string]                      `json:"description,omitzero"`
	// Whether to disable HTTP compression for this MCP server. Enable this if the
	// server does not support compressed responses.
	DisableCompression param.Opt[bool] `json:"disable_compression,omitzero"`
	// If true, the user will not be able to interrupt the agent while any tool from
	// this MCP server is running.
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If true, all tools from this MCP server will require pre-tool execution speech
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// The secret token (Authorization header) stored as a workspace secret or in-place
	// secret
	SecretToken ConvaiMcpServerNewParamsConfigSecretTokenUnion `json:"secret_token,omitzero"`
	// Defines the MCP server-level approval policy for tool execution.
	//
	// Any of "auto_approve_all", "require_approval_all", "require_approval_per_tool".
	ApprovalPolicy McpApprovalPolicy `json:"approval_policy,omitzero"`
	// Determines when and how all tools from this MCP server execute: 'immediate'
	// executes the tool right away when requested by the LLM, 'post_tool_speech' waits
	// for the agent to finish speaking before executing, 'async' runs the tool in the
	// background without blocking - best for long-running operations.
	//
	// Any of "immediate", "post_tool_speech", "async".
	ExecutionMode ToolExecutionMode `json:"execution_mode,omitzero"`
	// The headers included in the request
	RequestHeaders map[string]ConvaiMcpServerNewParamsConfigRequestHeaderUnion `json:"request_headers,omitzero"`
	// List of tool approval hashes for per-tool approval when approval_policy is
	// REQUIRE_APPROVAL_PER_TOOL
	ToolApprovalHashes []McpToolApprovalHashParam `json:"tool_approval_hashes,omitzero"`
	// Predefined tool call sound types.
	//
	// Any of "typing", "elevator1", "elevator2", "elevator3", "elevator4".
	ToolCallSound ToolCallSoundType `json:"tool_call_sound,omitzero"`
	// Determines when the tool call sound should play for all tools from this MCP
	// server
	//
	// Any of "auto", "always".
	ToolCallSoundBehavior ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitzero"`
	// List of per-tool configuration overrides that override the server-level defaults
	// for specific tools
	ToolConfigOverrides []McpToolConfigOverrideParam `json:"tool_config_overrides,omitzero"`
	// The transport type used to connect to the MCP server
	//
	// Any of "SSE", "STREAMABLE_HTTP".
	Transport McpServerTransport `json:"transport,omitzero"`
	paramObj
}

func (r ConvaiMcpServerNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerNewParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerNewParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiMcpServerNewParamsConfigURLUnion struct {
	OfString              param.Opt[string]         `json:",omitzero,inline"`
	OfConvAISecretLocator *ConvaiSecretLocatorParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiMcpServerNewParamsConfigURLUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConvAISecretLocator)
}
func (u *ConvaiMcpServerNewParamsConfigURLUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiMcpServerNewParamsConfigURLUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfConvAISecretLocator) {
		return u.OfConvAISecretLocator
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiMcpServerNewParamsConfigRequestHeaderUnion struct {
	OfString                param.Opt[string]                                                 `json:",omitzero,inline"`
	OfConvAISecretLocator   *ConvaiSecretLocatorParam                                         `json:",omitzero,inline"`
	OfConvAIDynamicVariable *ConvaiMcpServerNewParamsConfigRequestHeaderConvAIDynamicVariable `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiMcpServerNewParamsConfigRequestHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConvAISecretLocator, u.OfConvAIDynamicVariable)
}
func (u *ConvaiMcpServerNewParamsConfigRequestHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiMcpServerNewParamsConfigRequestHeaderUnion) asAny() any {
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
type ConvaiMcpServerNewParamsConfigRequestHeaderConvAIDynamicVariable struct {
	VariableName string `json:"variable_name,required"`
	paramObj
}

func (r ConvaiMcpServerNewParamsConfigRequestHeaderConvAIDynamicVariable) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerNewParamsConfigRequestHeaderConvAIDynamicVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerNewParamsConfigRequestHeaderConvAIDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiMcpServerNewParamsConfigSecretTokenUnion struct {
	OfConvAISecretLocator     *ConvaiSecretLocatorParam     `json:",omitzero,inline"`
	OfConvAIUserSecretDBModel *ConvaiUserSecretDBModelParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiMcpServerNewParamsConfigSecretTokenUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfConvAISecretLocator, u.OfConvAIUserSecretDBModel)
}
func (u *ConvaiMcpServerNewParamsConfigSecretTokenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiMcpServerNewParamsConfigSecretTokenUnion) asAny() any {
	if !param.IsOmitted(u.OfConvAISecretLocator) {
		return u.OfConvAISecretLocator
	} else if !param.IsOmitted(u.OfConvAIUserSecretDBModel) {
		return u.OfConvAIUserSecretDBModel
	}
	return nil
}

type ConvaiMcpServerGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiMcpServerUpdateParams struct {
	// Whether to disable HTTP compression for this MCP server
	DisableCompression param.Opt[bool] `json:"disable_compression,omitzero"`
	// If set, overrides the server's disable_interruptions setting for this tool
	DisableInterruptions param.Opt[bool] `json:"disable_interruptions,omitzero"`
	// If set, overrides the server's force_pre_tool_speech setting for this tool
	ForcePreToolSpeech param.Opt[bool] `json:"force_pre_tool_speech,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The headers to include in requests to the MCP server
	RequestHeaders map[string]ConvaiMcpServerUpdateParamsRequestHeaderUnion `json:"request_headers,omitzero"`
	// Defines the MCP server-level approval policy for tool execution.
	//
	// Any of "auto_approve_all", "require_approval_all", "require_approval_per_tool".
	ApprovalPolicy McpApprovalPolicy `json:"approval_policy,omitzero"`
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

func (r ConvaiMcpServerUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiMcpServerUpdateParamsRequestHeaderUnion struct {
	OfString                param.Opt[string]                                              `json:",omitzero,inline"`
	OfConvAISecretLocator   *ConvaiSecretLocatorParam                                      `json:",omitzero,inline"`
	OfConvAIDynamicVariable *ConvaiMcpServerUpdateParamsRequestHeaderConvAIDynamicVariable `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiMcpServerUpdateParamsRequestHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfConvAISecretLocator, u.OfConvAIDynamicVariable)
}
func (u *ConvaiMcpServerUpdateParamsRequestHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiMcpServerUpdateParamsRequestHeaderUnion) asAny() any {
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
type ConvaiMcpServerUpdateParamsRequestHeaderConvAIDynamicVariable struct {
	VariableName string `json:"variable_name,required"`
	paramObj
}

func (r ConvaiMcpServerUpdateParamsRequestHeaderConvAIDynamicVariable) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerUpdateParamsRequestHeaderConvAIDynamicVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerUpdateParamsRequestHeaderConvAIDynamicVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiMcpServerDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiMcpServerListToolsParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiMcpServerUpdateApprovalPolicyParams struct {
	// The approval mode to set for the MCP server
	//
	// Any of "auto_approve_all", "require_approval_all", "require_approval_per_tool".
	ApprovalPolicy McpApprovalPolicy `json:"approval_policy,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiMcpServerUpdateApprovalPolicyParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerUpdateApprovalPolicyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerUpdateApprovalPolicyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
