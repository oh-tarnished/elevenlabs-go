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
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// WorkspaceGroupMemberService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceGroupMemberService] method instead.
type WorkspaceGroupMemberService struct {
	Options []option.RequestOption
}

// NewWorkspaceGroupMemberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkspaceGroupMemberService(opts ...option.RequestOption) (r WorkspaceGroupMemberService) {
	r = WorkspaceGroupMemberService{}
	r.Options = opts
	return
}

// Adds a member of your workspace to the specified group. This endpoint may only
// be called by workspace administrators.
func (r *WorkspaceGroupMemberService) Add(ctx context.Context, groupID string, params WorkspaceGroupMemberAddParams, opts ...option.RequestOption) (res *WorkspaceGroupMemberAddResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/groups/%s/members", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Removes a member from the specified group. This endpoint may only be called by
// workspace administrators.
func (r *WorkspaceGroupMemberService) Remove(ctx context.Context, groupID string, params WorkspaceGroupMemberRemoveParams, opts ...option.RequestOption) (res *WorkspaceGroupMemberRemoveResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if groupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	path := fmt.Sprintf("v1/workspace/groups/%s/members/remove", groupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type WorkspaceGroupMemberAddResponse struct {
	// The status of the workspace group member addition request. If the request was
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
func (r WorkspaceGroupMemberAddResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceGroupMemberAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceGroupMemberRemoveResponse struct {
	// The status of the workspace group member deletion request. If the request was
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
func (r WorkspaceGroupMemberRemoveResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceGroupMemberRemoveResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceGroupMemberAddParams struct {
	// The email of the target workspace member.
	Email string `json:"email,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceGroupMemberAddParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceGroupMemberAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceGroupMemberAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceGroupMemberRemoveParams struct {
	// The email of the target workspace member.
	Email string `json:"email,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r WorkspaceGroupMemberRemoveParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceGroupMemberRemoveParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceGroupMemberRemoveParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
