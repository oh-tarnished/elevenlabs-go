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

// ConvaiWhatsappAccountService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiWhatsappAccountService] method instead.
type ConvaiWhatsappAccountService struct {
	Options []option.RequestOption
}

// NewConvaiWhatsappAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiWhatsappAccountService(opts ...option.RequestOption) (r ConvaiWhatsappAccountService) {
	r = ConvaiWhatsappAccountService{}
	r.Options = opts
	return
}

// Get a WhatsApp account
func (r *ConvaiWhatsappAccountService) Get(ctx context.Context, phoneNumberID string, query ConvaiWhatsappAccountGetParams, opts ...option.RequestOption) (res *GetWhatsAppAccountResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/whatsapp-accounts/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a WhatsApp account
func (r *ConvaiWhatsappAccountService) Update(ctx context.Context, phoneNumberID string, params ConvaiWhatsappAccountUpdateParams, opts ...option.RequestOption) (res *ConvaiWhatsappAccountUpdateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/whatsapp-accounts/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// List all WhatsApp accounts
func (r *ConvaiWhatsappAccountService) List(ctx context.Context, query ConvaiWhatsappAccountListParams, opts ...option.RequestOption) (res *ConvaiWhatsappAccountListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/whatsapp-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a WhatsApp account
func (r *ConvaiWhatsappAccountService) Delete(ctx context.Context, phoneNumberID string, body ConvaiWhatsappAccountDeleteParams, opts ...option.RequestOption) (res *ConvaiWhatsappAccountDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/whatsapp-accounts/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Import a WhatsApp account
func (r *ConvaiWhatsappAccountService) Import(ctx context.Context, params ConvaiWhatsappAccountImportParams, opts ...option.RequestOption) (res *ConvaiWhatsappAccountImportResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/whatsapp-accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type GetWhatsAppAccountResponse struct {
	AssignedAgentName   string `json:"assigned_agent_name,required"`
	BusinessAccountID   string `json:"business_account_id,required"`
	BusinessAccountName string `json:"business_account_name,required"`
	PhoneNumber         string `json:"phone_number,required"`
	PhoneNumberID       string `json:"phone_number_id,required"`
	PhoneNumberName     string `json:"phone_number_name,required"`
	AssignedAgentID     string `json:"assigned_agent_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssignedAgentName   respjson.Field
		BusinessAccountID   respjson.Field
		BusinessAccountName respjson.Field
		PhoneNumber         respjson.Field
		PhoneNumberID       respjson.Field
		PhoneNumberName     respjson.Field
		AssignedAgentID     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetWhatsAppAccountResponse) RawJSON() string { return r.JSON.raw }
func (r *GetWhatsAppAccountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiWhatsappAccountUpdateResponse = any

type ConvaiWhatsappAccountListResponse struct {
	Items []GetWhatsAppAccountResponse `json:"items,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiWhatsappAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiWhatsappAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiWhatsappAccountDeleteResponse = any

type ConvaiWhatsappAccountImportResponse struct {
	PhoneNumberID string `json:"phone_number_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumberID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiWhatsappAccountImportResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiWhatsappAccountImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiWhatsappAccountGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiWhatsappAccountUpdateParams struct {
	AssignedAgentID param.Opt[string] `json:"assigned_agent_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiWhatsappAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiWhatsappAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiWhatsappAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiWhatsappAccountListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiWhatsappAccountDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiWhatsappAccountImportParams struct {
	BusinessAccountID string `json:"business_account_id,required"`
	PhoneNumberID     string `json:"phone_number_id,required"`
	TokenCode         string `json:"token_code,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiWhatsappAccountImportParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiWhatsappAccountImportParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiWhatsappAccountImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
