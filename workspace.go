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

// WorkspaceService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	Options   []option.RequestOption
	Groups    WorkspaceGroupService
	Invites   WorkspaceInviteService
	Resources WorkspaceResourceService
	Webhooks  WorkspaceWebhookService
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r WorkspaceService) {
	r = WorkspaceService{}
	r.Options = opts
	r.Groups = NewWorkspaceGroupService(opts...)
	r.Invites = NewWorkspaceInviteService(opts...)
	r.Resources = NewWorkspaceResourceService(opts...)
	r.Webhooks = NewWorkspaceWebhookService(opts...)
	return
}

// Updates attributes of a workspace member. Apart from the email identifier, all
// parameters will remain unchanged unless specified. This endpoint may only be
// called by workspace administrators.
func (r *WorkspaceService) UpdateMember(ctx context.Context, params WorkspaceUpdateMemberParams, opts ...option.RequestOption) (res *WorkspaceUpdateMemberResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type WorkspaceUpdateMemberResponse struct {
	// The status of the workspace member update request. If the request was
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
func (r WorkspaceUpdateMemberResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceUpdateMemberResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceUpdateMemberParams struct {
	// Email of the target user.
	Email string `json:"email,required"`
	// Whether to lock or unlock the user account.
	IsLocked param.Opt[bool] `json:"is_locked,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Role dictating permissions in the workspace.
	//
	// Any of "workspace_admin", "workspace_member".
	WorkspaceRole WorkspaceUpdateMemberParamsWorkspaceRole `json:"workspace_role,omitzero"`
	paramObj
}

func (r WorkspaceUpdateMemberParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceUpdateMemberParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceUpdateMemberParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Role dictating permissions in the workspace.
type WorkspaceUpdateMemberParamsWorkspaceRole string

const (
	WorkspaceUpdateMemberParamsWorkspaceRoleWorkspaceAdmin  WorkspaceUpdateMemberParamsWorkspaceRole = "workspace_admin"
	WorkspaceUpdateMemberParamsWorkspaceRoleWorkspaceMember WorkspaceUpdateMemberParamsWorkspaceRole = "workspace_member"
)
