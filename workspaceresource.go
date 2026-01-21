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
)

// WorkspaceResourceService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceResourceService] method instead.
type WorkspaceResourceService struct {
	Options []option.RequestOption
}

// NewWorkspaceResourceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceResourceService(opts ...option.RequestOption) (r WorkspaceResourceService) {
	r = WorkspaceResourceService{}
	r.Options = opts
	return
}

// Gets the metadata of a resource by ID.
func (r *WorkspaceResourceService) Get(ctx context.Context, resourceID string, params WorkspaceResourceGetParams, opts ...option.RequestOption) (res *WorkspaceResourceGetResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/resources/%s", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Copies a workspace resource to another workspace.
func (r *WorkspaceResourceService) CopyToWorkspace(ctx context.Context, resourceID string, body WorkspaceResourceCopyToWorkspaceParams, opts ...option.RequestOption) (res *WorkspaceResourceCopyToWorkspaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/resources/%s/copy-to-workspace", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Grants a role on a workspace resource to a user or a group. It overrides any
// existing role this user/service account/group/workspace api key has on the
// resource. To target a user or service account, pass only the user email. The
// user must be in your workspace. To target a group, pass only the group id. To
// target a workspace api key, pass the api key id. The resource will be shared
// with the service account associated with the api key. You must have admin access
// to the resource to share it.
func (r *WorkspaceResourceService) Share(ctx context.Context, resourceID string, params WorkspaceResourceShareParams, opts ...option.RequestOption) (res *WorkspaceResourceShareResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/resources/%s/share", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Removes any existing role on a workspace resource from a user, service account,
// group or workspace api key. To target a user or service account, pass only the
// user email. The user must be in your workspace. To target a group, pass only the
// group id. To target a workspace api key, pass the api key id. The resource will
// be unshared from the service account associated with the api key. You must have
// admin access to the resource to unshare it. You cannot remove permissions from
// the user who created the resource.
func (r *WorkspaceResourceService) Unshare(ctx context.Context, resourceID string, params WorkspaceResourceUnshareParams, opts ...option.RequestOption) (res *WorkspaceResourceUnshareResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/resources/%s/unshare", resourceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Resource types that can be shared in the workspace. The name always need to
// match the collection names
type WorkspaceResourceType string

const (
	WorkspaceResourceTypeVoice                                  WorkspaceResourceType = "voice"
	WorkspaceResourceTypeVoiceCollection                        WorkspaceResourceType = "voice_collection"
	WorkspaceResourceTypePronunciationDictionary                WorkspaceResourceType = "pronunciation_dictionary"
	WorkspaceResourceTypeDubbing                                WorkspaceResourceType = "dubbing"
	WorkspaceResourceTypeProject                                WorkspaceResourceType = "project"
	WorkspaceResourceTypeConvaiAgents                           WorkspaceResourceType = "convai_agents"
	WorkspaceResourceTypeConvaiKnowledgeBaseDocuments           WorkspaceResourceType = "convai_knowledge_base_documents"
	WorkspaceResourceTypeConvaiTools                            WorkspaceResourceType = "convai_tools"
	WorkspaceResourceTypeConvaiSettings                         WorkspaceResourceType = "convai_settings"
	WorkspaceResourceTypeConvaiSecrets                          WorkspaceResourceType = "convai_secrets"
	WorkspaceResourceTypeWorkspaceAuthConnections               WorkspaceResourceType = "workspace_auth_connections"
	WorkspaceResourceTypeConvaiPhoneNumbers                     WorkspaceResourceType = "convai_phone_numbers"
	WorkspaceResourceTypeConvaiMcpServers                       WorkspaceResourceType = "convai_mcp_servers"
	WorkspaceResourceTypeConvaiAPIIntegrationConnections        WorkspaceResourceType = "convai_api_integration_connections"
	WorkspaceResourceTypeConvaiAPIIntegrationTriggerConnections WorkspaceResourceType = "convai_api_integration_trigger_connections"
	WorkspaceResourceTypeConvaiBatchCalls                       WorkspaceResourceType = "convai_batch_calls"
	WorkspaceResourceTypeConvaiAgentResponseTests               WorkspaceResourceType = "convai_agent_response_tests"
	WorkspaceResourceTypeConvaiTestSuiteInvocations             WorkspaceResourceType = "convai_test_suite_invocations"
	WorkspaceResourceTypeConvaiCrawlJobs                        WorkspaceResourceType = "convai_crawl_jobs"
	WorkspaceResourceTypeConvaiCrawlTasks                       WorkspaceResourceType = "convai_crawl_tasks"
	WorkspaceResourceTypeConvaiWhatsappAccounts                 WorkspaceResourceType = "convai_whatsapp_accounts"
	WorkspaceResourceTypeConvaiAgentVersions                    WorkspaceResourceType = "convai_agent_versions"
	WorkspaceResourceTypeConvaiAgentBranches                    WorkspaceResourceType = "convai_agent_branches"
	WorkspaceResourceTypeConvaiAgentVersionsDeployments         WorkspaceResourceType = "convai_agent_versions_deployments"
	WorkspaceResourceTypeDashboard                              WorkspaceResourceType = "dashboard"
	WorkspaceResourceTypeDashboardConfiguration                 WorkspaceResourceType = "dashboard_configuration"
	WorkspaceResourceTypeConvaiAgentDrafts                      WorkspaceResourceType = "convai_agent_drafts"
	WorkspaceResourceTypeResourceLocators                       WorkspaceResourceType = "resource_locators"
	WorkspaceResourceTypeAssets                                 WorkspaceResourceType = "assets"
	WorkspaceResourceTypeContentGenerations                     WorkspaceResourceType = "content_generations"
)

type WorkspaceResourceGetResponse struct {
	// The access level for anonymous users. If None, the resource is not shared
	// publicly.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	AnonymousAccessLevelOverride WorkspaceResourceGetResponseAnonymousAccessLevelOverride `json:"anonymous_access_level_override,required"`
	// The ID of the user who created the resource.
	CreatorUserID string `json:"creator_user_id,required"`
	// The ID of the resource.
	ResourceID string `json:"resource_id,required"`
	// The type of the resource.
	//
	// Any of "voice", "voice_collection", "pronunciation_dictionary", "dubbing",
	// "project", "convai_agents", "convai_knowledge_base_documents", "convai_tools",
	// "convai_settings", "convai_secrets", "workspace_auth_connections",
	// "convai_phone_numbers", "convai_mcp_servers",
	// "convai_api_integration_connections",
	// "convai_api_integration_trigger_connections", "convai_batch_calls",
	// "convai_agent_response_tests", "convai_test_suite_invocations",
	// "convai_crawl_jobs", "convai_crawl_tasks", "convai_whatsapp_accounts",
	// "convai_agent_versions", "convai_agent_branches",
	// "convai_agent_versions_deployments", "dashboard", "dashboard_configuration",
	// "convai_agent_drafts", "resource_locators", "assets", "content_generations".
	ResourceType WorkspaceResourceType `json:"resource_type,required"`
	// A mapping of roles to group IDs. When the resource is shared with a user, the
	// group id is the user's id.
	RoleToGroupIDs map[string][]string `json:"role_to_group_ids,required"`
	// List of options for sharing the resource further in the workspace. These are
	// users who don't have access to the resource yet.
	ShareOptions []WorkspaceResourceGetResponseShareOption `json:"share_options,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AnonymousAccessLevelOverride respjson.Field
		CreatorUserID                respjson.Field
		ResourceID                   respjson.Field
		ResourceType                 respjson.Field
		RoleToGroupIDs               respjson.Field
		ShareOptions                 respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceResourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceResourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The access level for anonymous users. If None, the resource is not shared
// publicly.
type WorkspaceResourceGetResponseAnonymousAccessLevelOverride string

const (
	WorkspaceResourceGetResponseAnonymousAccessLevelOverrideAdmin     WorkspaceResourceGetResponseAnonymousAccessLevelOverride = "admin"
	WorkspaceResourceGetResponseAnonymousAccessLevelOverrideEditor    WorkspaceResourceGetResponseAnonymousAccessLevelOverride = "editor"
	WorkspaceResourceGetResponseAnonymousAccessLevelOverrideCommenter WorkspaceResourceGetResponseAnonymousAccessLevelOverride = "commenter"
	WorkspaceResourceGetResponseAnonymousAccessLevelOverrideViewer    WorkspaceResourceGetResponseAnonymousAccessLevelOverride = "viewer"
)

type WorkspaceResourceGetResponseShareOption struct {
	// The ID of the principal.
	ID string `json:"id,required"`
	// The name of the principal.
	Name string `json:"name,required"`
	// The type of the principal: user, group, or service account (under 'key').
	//
	// Any of "user", "group", "key".
	Type string `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceResourceGetResponseShareOption) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceResourceGetResponseShareOption) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceResourceCopyToWorkspaceResponse = any

type WorkspaceResourceShareResponse = any

type WorkspaceResourceUnshareResponse = any

type WorkspaceResourceGetParams struct {
	// Resource type of the target resource.
	//
	// Any of "voice", "voice_collection", "pronunciation_dictionary", "dubbing",
	// "project", "convai_agents", "convai_knowledge_base_documents", "convai_tools",
	// "convai_settings", "convai_secrets", "workspace_auth_connections",
	// "convai_phone_numbers", "convai_mcp_servers",
	// "convai_api_integration_connections",
	// "convai_api_integration_trigger_connections", "convai_batch_calls",
	// "convai_agent_response_tests", "convai_test_suite_invocations",
	// "convai_crawl_jobs", "convai_crawl_tasks", "convai_whatsapp_accounts",
	// "convai_agent_versions", "convai_agent_branches",
	// "convai_agent_versions_deployments", "dashboard", "dashboard_configuration",
	// "convai_agent_drafts", "resource_locators", "assets", "content_generations".
	ResourceType WorkspaceResourceType `query:"resource_type,omitzero,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceResourceGetParams]'s query parameters as
// `url.Values`.
func (r WorkspaceResourceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WorkspaceResourceCopyToWorkspaceParams struct {
	// Resource type of the target resource.
	//
	// Any of "voice", "voice_collection", "pronunciation_dictionary", "dubbing",
	// "project", "convai_agents", "convai_knowledge_base_documents", "convai_tools",
	// "convai_settings", "convai_secrets", "workspace_auth_connections",
	// "convai_phone_numbers", "convai_mcp_servers",
	// "convai_api_integration_connections",
	// "convai_api_integration_trigger_connections", "convai_batch_calls",
	// "convai_agent_response_tests", "convai_test_suite_invocations",
	// "convai_crawl_jobs", "convai_crawl_tasks", "convai_whatsapp_accounts",
	// "convai_agent_versions", "convai_agent_branches",
	// "convai_agent_versions_deployments", "dashboard", "dashboard_configuration",
	// "convai_agent_drafts", "resource_locators", "assets", "content_generations".
	ResourceType WorkspaceResourceType `json:"resource_type,omitzero,required"`
	// The ID of the target user.
	TargetUserID string `json:"target_user_id,required"`
	paramObj
}

func (r WorkspaceResourceCopyToWorkspaceParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceResourceCopyToWorkspaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceResourceCopyToWorkspaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceResourceShareParams struct {
	// Resource type of the target resource.
	//
	// Any of "voice", "voice_collection", "pronunciation_dictionary", "dubbing",
	// "project", "convai_agents", "convai_knowledge_base_documents", "convai_tools",
	// "convai_settings", "convai_secrets", "workspace_auth_connections",
	// "convai_phone_numbers", "convai_mcp_servers",
	// "convai_api_integration_connections",
	// "convai_api_integration_trigger_connections", "convai_batch_calls",
	// "convai_agent_response_tests", "convai_test_suite_invocations",
	// "convai_crawl_jobs", "convai_crawl_tasks", "convai_whatsapp_accounts",
	// "convai_agent_versions", "convai_agent_branches",
	// "convai_agent_versions_deployments", "dashboard", "dashboard_configuration",
	// "convai_agent_drafts", "resource_locators", "assets", "content_generations".
	ResourceType WorkspaceResourceType `json:"resource_type,omitzero,required"`
	// Role to update the target principal with.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	Role WorkspaceResourceShareParamsRole `json:"role,omitzero,required"`
	// The ID of the target group. To target the permissions principals have by default
	// on this resource, use the value 'default'.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// The email of the user or service account.
	UserEmail param.Opt[string] `json:"user_email,omitzero"`
	// The ID of the target workspace API key. This isn't the same as the key itself
	// that would you pass in the header for authentication. Workspace admins can find
	// this in the workspace settings UI.
	WorkspaceAPIKeyID param.Opt[string] `json:"workspace_api_key_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceResourceShareParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceResourceShareParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceResourceShareParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role to update the target principal with.
type WorkspaceResourceShareParamsRole string

const (
	WorkspaceResourceShareParamsRoleAdmin     WorkspaceResourceShareParamsRole = "admin"
	WorkspaceResourceShareParamsRoleEditor    WorkspaceResourceShareParamsRole = "editor"
	WorkspaceResourceShareParamsRoleCommenter WorkspaceResourceShareParamsRole = "commenter"
	WorkspaceResourceShareParamsRoleViewer    WorkspaceResourceShareParamsRole = "viewer"
)

type WorkspaceResourceUnshareParams struct {
	// Resource type of the target resource.
	//
	// Any of "voice", "voice_collection", "pronunciation_dictionary", "dubbing",
	// "project", "convai_agents", "convai_knowledge_base_documents", "convai_tools",
	// "convai_settings", "convai_secrets", "workspace_auth_connections",
	// "convai_phone_numbers", "convai_mcp_servers",
	// "convai_api_integration_connections",
	// "convai_api_integration_trigger_connections", "convai_batch_calls",
	// "convai_agent_response_tests", "convai_test_suite_invocations",
	// "convai_crawl_jobs", "convai_crawl_tasks", "convai_whatsapp_accounts",
	// "convai_agent_versions", "convai_agent_branches",
	// "convai_agent_versions_deployments", "dashboard", "dashboard_configuration",
	// "convai_agent_drafts", "resource_locators", "assets", "content_generations".
	ResourceType WorkspaceResourceType `json:"resource_type,omitzero,required"`
	// The ID of the target group. To target the permissions principals have by default
	// on this resource, use the value 'default'.
	GroupID param.Opt[string] `json:"group_id,omitzero"`
	// The email of the user or service account.
	UserEmail param.Opt[string] `json:"user_email,omitzero"`
	// The ID of the target workspace API key. This isn't the same as the key itself
	// that would you pass in the header for authentication. Workspace admins can find
	// this in the workspace settings UI.
	WorkspaceAPIKeyID param.Opt[string] `json:"workspace_api_key_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceResourceUnshareParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceResourceUnshareParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceResourceUnshareParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
