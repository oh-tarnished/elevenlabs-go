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

// ServiceAccountService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceAccountService] method instead.
type ServiceAccountService struct {
	Options []option.RequestOption
	APIKeys ServiceAccountAPIKeyService
}

// NewServiceAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServiceAccountService(opts ...option.RequestOption) (r ServiceAccountService) {
	r = ServiceAccountService{}
	r.Options = opts
	r.APIKeys = NewServiceAccountAPIKeyService(opts...)
	return
}

// List all service accounts in the workspace
func (r *ServiceAccountService) List(ctx context.Context, query ServiceAccountListParams, opts ...option.RequestOption) (res *ServiceAccountListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/service-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ServiceAccountListResponse struct {
	ServiceAccounts []ServiceAccountListResponseServiceAccount `json:"service-accounts,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ServiceAccounts respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountListResponseServiceAccount struct {
	APIKeys              []WorkspaceAPIKey `json:"api-keys,required"`
	Name                 string            `json:"name,required"`
	ServiceAccountUserID string            `json:"service_account_user_id,required"`
	CreatedAtUnix        int64             `json:"created_at_unix,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeys              respjson.Field
		Name                 respjson.Field
		ServiceAccountUserID respjson.Field
		CreatedAtUnix        respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountListResponseServiceAccount) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountListResponseServiceAccount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
