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

// WorkspaceInviteService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceInviteService] method instead.
type WorkspaceInviteService struct {
	Options []option.RequestOption
}

// NewWorkspaceInviteService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceInviteService(opts ...option.RequestOption) (r WorkspaceInviteService) {
	r = WorkspaceInviteService{}
	r.Options = opts
	return
}

// Invalidates an existing email invitation. The invitation will still show up in
// the inbox it has been delivered to, but activating it to join the workspace
// won't work. This endpoint may only be called by workspace administrators.
func (r *WorkspaceInviteService) Delete(ctx context.Context, params WorkspaceInviteDeleteParams, opts ...option.RequestOption) (res *WorkspaceInviteDeleteResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/invites"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Sends an email invitation to join your workspace to the provided email. If the
// user doesn't have an account they will be prompted to create one. If the user
// accepts this invite they will be added as a user to your workspace and your
// subscription using one of your seats. This endpoint may only be called by
// workspace administrators. If the user is already in the workspace a 400 error
// will be returned.
func (r *WorkspaceInviteService) Add(ctx context.Context, params WorkspaceInviteAddParams, opts ...option.RequestOption) (res *AddWorkspaceInvite, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/invites/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Sends email invitations to join your workspace to the provided emails. Requires
// all email addresses to be part of a verified domain. If the users don't have an
// account they will be prompted to create one. If the users accept these invites
// they will be added as users to your workspace and your subscription using one of
// your seats. This endpoint may only be called by workspace administrators.
func (r *WorkspaceInviteService) AddBulk(ctx context.Context, params WorkspaceInviteAddBulkParams, opts ...option.RequestOption) (res *AddWorkspaceInvite, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/invites/add-bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AddWorkspaceInvite struct {
	// The status of the workspace invite request. If the request was successful, the
	// status will be 'ok'. Otherwise an error message with status 500 will be
	// returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddWorkspaceInvite) RawJSON() string { return r.JSON.raw }
func (r *AddWorkspaceInvite) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceInviteDeleteResponse struct {
	// The status of the workspace invite deletion request. If the request was
	// successful, the status will be 'ok'. Otherwise an error message with status 500
	// will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceInviteDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceInviteDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceInviteDeleteParams struct {
	// The email of the customer
	Email string `json:"email,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceInviteDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceInviteDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceInviteDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceInviteAddParams struct {
	// The email of the customer
	Email string `json:"email,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The group ids of the user
	GroupIDs []string `json:"group_ids,omitzero"`
	// The workspace permission of the user
	//
	// Any of "external", "admin", "workspace_admin", "workspace_member", "support_l1",
	// "support_l2", "moderator", "sales", "voice_mixer", "voice_admin",
	// "convai_admin", "enterprise_viewer", "quality_check_admin",
	// "workspace_migration_admin", "human_reviewer", "productions_admin", "support",
	// "internal".
	WorkspacePermission WorkspaceInviteAddParamsWorkspacePermission `json:"workspace_permission,omitzero"`
	paramObj
}

func (r WorkspaceInviteAddParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceInviteAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceInviteAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The workspace permission of the user
type WorkspaceInviteAddParamsWorkspacePermission string

const (
	WorkspaceInviteAddParamsWorkspacePermissionExternal                WorkspaceInviteAddParamsWorkspacePermission = "external"
	WorkspaceInviteAddParamsWorkspacePermissionAdmin                   WorkspaceInviteAddParamsWorkspacePermission = "admin"
	WorkspaceInviteAddParamsWorkspacePermissionWorkspaceAdmin          WorkspaceInviteAddParamsWorkspacePermission = "workspace_admin"
	WorkspaceInviteAddParamsWorkspacePermissionWorkspaceMember         WorkspaceInviteAddParamsWorkspacePermission = "workspace_member"
	WorkspaceInviteAddParamsWorkspacePermissionSupportL1               WorkspaceInviteAddParamsWorkspacePermission = "support_l1"
	WorkspaceInviteAddParamsWorkspacePermissionSupportL2               WorkspaceInviteAddParamsWorkspacePermission = "support_l2"
	WorkspaceInviteAddParamsWorkspacePermissionModerator               WorkspaceInviteAddParamsWorkspacePermission = "moderator"
	WorkspaceInviteAddParamsWorkspacePermissionSales                   WorkspaceInviteAddParamsWorkspacePermission = "sales"
	WorkspaceInviteAddParamsWorkspacePermissionVoiceMixer              WorkspaceInviteAddParamsWorkspacePermission = "voice_mixer"
	WorkspaceInviteAddParamsWorkspacePermissionVoiceAdmin              WorkspaceInviteAddParamsWorkspacePermission = "voice_admin"
	WorkspaceInviteAddParamsWorkspacePermissionConvaiAdmin             WorkspaceInviteAddParamsWorkspacePermission = "convai_admin"
	WorkspaceInviteAddParamsWorkspacePermissionEnterpriseViewer        WorkspaceInviteAddParamsWorkspacePermission = "enterprise_viewer"
	WorkspaceInviteAddParamsWorkspacePermissionQualityCheckAdmin       WorkspaceInviteAddParamsWorkspacePermission = "quality_check_admin"
	WorkspaceInviteAddParamsWorkspacePermissionWorkspaceMigrationAdmin WorkspaceInviteAddParamsWorkspacePermission = "workspace_migration_admin"
	WorkspaceInviteAddParamsWorkspacePermissionHumanReviewer           WorkspaceInviteAddParamsWorkspacePermission = "human_reviewer"
	WorkspaceInviteAddParamsWorkspacePermissionProductionsAdmin        WorkspaceInviteAddParamsWorkspacePermission = "productions_admin"
	WorkspaceInviteAddParamsWorkspacePermissionSupport                 WorkspaceInviteAddParamsWorkspacePermission = "support"
	WorkspaceInviteAddParamsWorkspacePermissionInternal                WorkspaceInviteAddParamsWorkspacePermission = "internal"
)

type WorkspaceInviteAddBulkParams struct {
	// The email of the customer
	Emails []string `json:"emails,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The group ids of the user
	GroupIDs []string `json:"group_ids,omitzero"`
	paramObj
}

func (r WorkspaceInviteAddBulkParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceInviteAddBulkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceInviteAddBulkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
