// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
)

// ConvaiMcpServerToolApprovalService contains methods and other services that help
// with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiMcpServerToolApprovalService] method instead.
type ConvaiMcpServerToolApprovalService struct {
	Options []option.RequestOption
}

// NewConvaiMcpServerToolApprovalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiMcpServerToolApprovalService(opts ...option.RequestOption) (r ConvaiMcpServerToolApprovalService) {
	r = ConvaiMcpServerToolApprovalService{}
	r.Options = opts
	return
}

// Add approval for a specific MCP tool when using per-tool approval mode.
func (r *ConvaiMcpServerToolApprovalService) New(ctx context.Context, mcpServerID string, params ConvaiMcpServerToolApprovalNewParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mcpServerID == "" {
		err = errors.New("missing required mcp_server_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-approvals", mcpServerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remove approval for a specific MCP tool when using per-tool approval mode.
func (r *ConvaiMcpServerToolApprovalService) Delete(ctx context.Context, toolName string, params ConvaiMcpServerToolApprovalDeleteParams, opts ...option.RequestOption) (res *McpServerResponseModel, err error) {
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
	path := fmt.Sprintf("v1/convai/mcp-servers/%s/tool-approvals/%s", params.McpServerID, toolName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Defines the tool-level approval policy.
type McpToolApprovalPolicy string

const (
	McpToolApprovalPolicyAutoApproved     McpToolApprovalPolicy = "auto_approved"
	McpToolApprovalPolicyRequiresApproval McpToolApprovalPolicy = "requires_approval"
)

type ConvaiMcpServerToolApprovalNewParams struct {
	// The description of the MCP tool
	ToolDescription string `json:"tool_description,required"`
	// The name of the MCP tool
	ToolName string `json:"tool_name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The tool-level approval policy
	//
	// Any of "auto_approved", "requires_approval".
	ApprovalPolicy McpToolApprovalPolicy `json:"approval_policy,omitzero"`
	// The input schema of the MCP tool (the schema defined on the MCP server before
	// ElevenLabs does any extra processing)
	InputSchema any `json:"input_schema,omitzero"`
	paramObj
}

func (r ConvaiMcpServerToolApprovalNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiMcpServerToolApprovalNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiMcpServerToolApprovalNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiMcpServerToolApprovalDeleteParams struct {
	// ID of the MCP Server.
	McpServerID string `path:"mcp_server_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
