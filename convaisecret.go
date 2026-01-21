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
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiSecretService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiSecretService] method instead.
type ConvaiSecretService struct {
	Options []option.RequestOption
}

// NewConvaiSecretService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiSecretService(opts ...option.RequestOption) (r ConvaiSecretService) {
	r = ConvaiSecretService{}
	r.Options = opts
	return
}

// Create a new secret for the workspace
func (r *ConvaiSecretService) New(ctx context.Context, params ConvaiSecretNewParams, opts ...option.RequestOption) (res *WorkspaceSecret, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an existing secret for the workspace
func (r *ConvaiSecretService) Update(ctx context.Context, secretID string, params ConvaiSecretUpdateParams, opts ...option.RequestOption) (res *WorkspaceSecret, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/secrets/%s", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get all workspace secrets for the user
func (r *ConvaiSecretService) List(ctx context.Context, query ConvaiSecretListParams, opts ...option.RequestOption) (res *ConvaiSecretListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a workspace secret if it's not in use
func (r *ConvaiSecretService) Delete(ctx context.Context, secretID string, body ConvaiSecretDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if secretID == "" {
		err = errors.New("missing required secret_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/secrets/%s", secretID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type WorkspaceSecret struct {
	Name     string          `json:"name,required"`
	SecretID string          `json:"secret_id,required"`
	Type     constant.Stored `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		SecretID    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceSecret) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListResponse struct {
	Secrets []ConvaiSecretListResponseSecret `json:"secrets,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Secrets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListResponseSecret struct {
	Name     string                               `json:"name,required"`
	SecretID string                               `json:"secret_id,required"`
	Type     constant.Stored                      `json:"type,required"`
	UsedBy   ConvaiSecretListResponseSecretUsedBy `json:"used_by,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		SecretID    respjson.Field
		Type        respjson.Field
		UsedBy      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponseSecret) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponseSecret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListResponseSecretUsedBy struct {
	Agents []ConvaiSecretListResponseSecretUsedByAgentUnion `json:"agents,required"`
	// Any of "conversation_initiation_webhook".
	Others       []string                                          `json:"others,required"`
	Tools        []ConvaiSecretListResponseSecretUsedByToolUnion   `json:"tools,required"`
	PhoneNumbers []ConvaiSecretListResponseSecretUsedByPhoneNumber `json:"phone_numbers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agents       respjson.Field
		Others       respjson.Field
		Tools        respjson.Field
		PhoneNumbers respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponseSecretUsedBy) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponseSecretUsedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiSecretListResponseSecretUsedByAgentUnion contains all possible properties
// and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiSecretListResponseSecretUsedByAgentUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiSecretListResponseSecretUsedByAgentUnion struct {
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

// anyConvaiSecretListResponseSecretUsedByAgent is implemented by each variant of
// [ConvaiSecretListResponseSecretUsedByAgentUnion] to add type safety for the
// return type of [ConvaiSecretListResponseSecretUsedByAgentUnion.AsAny]
type anyConvaiSecretListResponseSecretUsedByAgent interface {
	implConvaiSecretListResponseSecretUsedByAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiSecretListResponseSecretUsedByAgentUnion() {}
func (DependentUnknownAgentIdentifier) implConvaiSecretListResponseSecretUsedByAgentUnion()   {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiSecretListResponseSecretUsedByAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiSecretListResponseSecretUsedByAgentUnion) AsAny() anyConvaiSecretListResponseSecretUsedByAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiSecretListResponseSecretUsedByAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiSecretListResponseSecretUsedByAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiSecretListResponseSecretUsedByAgentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiSecretListResponseSecretUsedByAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiSecretListResponseSecretUsedByToolUnion contains all possible properties
// and values from [ConvaiSecretListResponseSecretUsedByToolAvailable],
// [ConvaiSecretListResponseSecretUsedByToolUnknown].
//
// Use the [ConvaiSecretListResponseSecretUsedByToolUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiSecretListResponseSecretUsedByToolUnion struct {
	// This field is from variant [ConvaiSecretListResponseSecretUsedByToolAvailable].
	ID string `json:"id"`
	// This field is from variant [ConvaiSecretListResponseSecretUsedByToolAvailable].
	AccessLevel string `json:"access_level"`
	// This field is from variant [ConvaiSecretListResponseSecretUsedByToolAvailable].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [ConvaiSecretListResponseSecretUsedByToolAvailable].
	Name string `json:"name"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                respjson.Field
		AccessLevel       respjson.Field
		CreatedAtUnixSecs respjson.Field
		Name              respjson.Field
		Type              respjson.Field
		raw               string
	} `json:"-"`
}

// anyConvaiSecretListResponseSecretUsedByTool is implemented by each variant of
// [ConvaiSecretListResponseSecretUsedByToolUnion] to add type safety for the
// return type of [ConvaiSecretListResponseSecretUsedByToolUnion.AsAny]
type anyConvaiSecretListResponseSecretUsedByTool interface {
	implConvaiSecretListResponseSecretUsedByToolUnion()
}

func (ConvaiSecretListResponseSecretUsedByToolAvailable) implConvaiSecretListResponseSecretUsedByToolUnion() {
}
func (ConvaiSecretListResponseSecretUsedByToolUnknown) implConvaiSecretListResponseSecretUsedByToolUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiSecretListResponseSecretUsedByToolUnion.AsAny().(type) {
//	case elevenlabs.ConvaiSecretListResponseSecretUsedByToolAvailable:
//	case elevenlabs.ConvaiSecretListResponseSecretUsedByToolUnknown:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiSecretListResponseSecretUsedByToolUnion) AsAny() anyConvaiSecretListResponseSecretUsedByTool {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiSecretListResponseSecretUsedByToolUnion) AsAvailable() (v ConvaiSecretListResponseSecretUsedByToolAvailable) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiSecretListResponseSecretUsedByToolUnion) AsUnknown() (v ConvaiSecretListResponseSecretUsedByToolUnknown) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiSecretListResponseSecretUsedByToolUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiSecretListResponseSecretUsedByToolUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListResponseSecretUsedByToolAvailable struct {
	ID string `json:"id,required"`
	// Any of "admin", "editor", "commenter", "viewer".
	AccessLevel       string `json:"access_level,required"`
	CreatedAtUnixSecs int64  `json:"created_at_unix_secs,required"`
	Name              string `json:"name,required"`
	// Any of "available".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		AccessLevel       respjson.Field
		CreatedAtUnixSecs respjson.Field
		Name              respjson.Field
		Type              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponseSecretUsedByToolAvailable) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponseSecretUsedByToolAvailable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A model that represents an tool dependent on a knowledge base/tools to which the
// user has no direct access.
type ConvaiSecretListResponseSecretUsedByToolUnknown struct {
	// Any of "unknown".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponseSecretUsedByToolUnknown) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponseSecretUsedByToolUnknown) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListResponseSecretUsedByPhoneNumber struct {
	Label         string `json:"label,required"`
	PhoneNumber   string `json:"phone_number,required"`
	PhoneNumberID string `json:"phone_number_id,required"`
	// Any of "twilio", "sip_trunk".
	Provider TelephonyProvider `json:"provider,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label         respjson.Field
		PhoneNumber   respjson.Field
		PhoneNumberID respjson.Field
		Provider      respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSecretListResponseSecretUsedByPhoneNumber) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSecretListResponseSecretUsedByPhoneNumber) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretNewParams struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// This field can be elided, and will marshal its zero value as "new".
	Type constant.New `json:"type,required"`
	paramObj
}

func (r ConvaiSecretNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSecretNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretUpdateParams struct {
	Name  string `json:"name,required"`
	Value string `json:"value,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// This field can be elided, and will marshal its zero value as "update".
	Type constant.Update `json:"type,required"`
	paramObj
}

func (r ConvaiSecretUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSecretUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSecretUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSecretListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiSecretDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
