// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
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

// WorkspaceGroupService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceGroupService] method instead.
type WorkspaceGroupService struct {
	Options []option.RequestOption
	Members WorkspaceGroupMemberService
}

// NewWorkspaceGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceGroupService(opts ...option.RequestOption) (r WorkspaceGroupService) {
	r = WorkspaceGroupService{}
	r.Options = opts
	r.Members = NewWorkspaceGroupMemberService(opts...)
	return
}

// Searches for user groups in the workspace. Multiple or no groups may be
// returned.
func (r *WorkspaceGroupService) Search(ctx context.Context, params WorkspaceGroupSearchParams, opts ...option.RequestOption) (res *[]WorkspaceGroupSearchResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/workspace/groups/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type WorkspaceGroupSearchResponse struct {
	// The ID of the workspace group.
	ID string `json:"id,required"`
	// The emails of the members of the workspace group.
	MembersEmails []string `json:"members_emails,required"`
	// The name of the workspace group.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		MembersEmails respjson.Field
		Name          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceGroupSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceGroupSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceGroupSearchParams struct {
	// Name of the target group.
	Name string `query:"name,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WorkspaceGroupSearchParams]'s query parameters as
// `url.Values`.
func (r WorkspaceGroupSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
