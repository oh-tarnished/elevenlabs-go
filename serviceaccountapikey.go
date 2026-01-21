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
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ServiceAccountAPIKeyService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceAccountAPIKeyService] method instead.
type ServiceAccountAPIKeyService struct {
	Options []option.RequestOption
}

// NewServiceAccountAPIKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewServiceAccountAPIKeyService(opts ...option.RequestOption) (r ServiceAccountAPIKeyService) {
	r = ServiceAccountAPIKeyService{}
	r.Options = opts
	return
}

// Create a new API key for a service account
func (r *ServiceAccountAPIKeyService) New(ctx context.Context, serviceAccountUserID string, params ServiceAccountAPIKeyNewParams, opts ...option.RequestOption) (res *ServiceAccountAPIKeyNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if serviceAccountUserID == "" {
		err = errors.New("missing required service_account_user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/service-accounts/%s/api-keys", serviceAccountUserID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Update an existing API key for a service account
func (r *ServiceAccountAPIKeyService) Update(ctx context.Context, apiKeyID string, params ServiceAccountAPIKeyUpdateParams, opts ...option.RequestOption) (res *ServiceAccountAPIKeyUpdateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ServiceAccountUserID == "" {
		err = errors.New("missing required service_account_user_id parameter")
		return
	}
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/service-accounts/%s/api-keys/%s", params.ServiceAccountUserID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get all API keys for a service account
func (r *ServiceAccountAPIKeyService) List(ctx context.Context, serviceAccountUserID string, query ServiceAccountAPIKeyListParams, opts ...option.RequestOption) (res *ServiceAccountAPIKeyListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if serviceAccountUserID == "" {
		err = errors.New("missing required service_account_user_id parameter")
		return
	}
	path := fmt.Sprintf("v1/service-accounts/%s/api-keys", serviceAccountUserID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an existing API key for a service account
func (r *ServiceAccountAPIKeyService) Delete(ctx context.Context, apiKeyID string, params ServiceAccountAPIKeyDeleteParams, opts ...option.RequestOption) (res *ServiceAccountAPIKeyDeleteResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ServiceAccountUserID == "" {
		err = errors.New("missing required service_account_user_id parameter")
		return
	}
	if apiKeyID == "" {
		err = errors.New("missing required api_key_id parameter")
		return
	}
	path := fmt.Sprintf("v1/service-accounts/%s/api-keys/%s", params.ServiceAccountUserID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type PermissionType string

const (
	PermissionTypeTextToSpeech                   PermissionType = "text_to_speech"
	PermissionTypeSpeechToSpeech                 PermissionType = "speech_to_speech"
	PermissionTypeSpeechToText                   PermissionType = "speech_to_text"
	PermissionTypeModelsRead                     PermissionType = "models_read"
	PermissionTypeModelsWrite                    PermissionType = "models_write"
	PermissionTypeVoicesRead                     PermissionType = "voices_read"
	PermissionTypeVoicesWrite                    PermissionType = "voices_write"
	PermissionTypeSpeechHistoryRead              PermissionType = "speech_history_read"
	PermissionTypeSpeechHistoryWrite             PermissionType = "speech_history_write"
	PermissionTypeSoundGeneration                PermissionType = "sound_generation"
	PermissionTypeAudioIsolation                 PermissionType = "audio_isolation"
	PermissionTypeVoiceGeneration                PermissionType = "voice_generation"
	PermissionTypeDubbingRead                    PermissionType = "dubbing_read"
	PermissionTypeDubbingWrite                   PermissionType = "dubbing_write"
	PermissionTypePronunciationDictionariesRead  PermissionType = "pronunciation_dictionaries_read"
	PermissionTypePronunciationDictionariesWrite PermissionType = "pronunciation_dictionaries_write"
	PermissionTypeUserRead                       PermissionType = "user_read"
	PermissionTypeUserWrite                      PermissionType = "user_write"
	PermissionTypeProjectsRead                   PermissionType = "projects_read"
	PermissionTypeProjectsWrite                  PermissionType = "projects_write"
	PermissionTypeAudioNativeRead                PermissionType = "audio_native_read"
	PermissionTypeAudioNativeWrite               PermissionType = "audio_native_write"
	PermissionTypeWorkspaceRead                  PermissionType = "workspace_read"
	PermissionTypeWorkspaceWrite                 PermissionType = "workspace_write"
	PermissionTypeForcedAlignment                PermissionType = "forced_alignment"
	PermissionTypeConvaiRead                     PermissionType = "convai_read"
	PermissionTypeConvaiWrite                    PermissionType = "convai_write"
	PermissionTypeMusicGeneration                PermissionType = "music_generation"
	PermissionTypeAddVoiceFromVoiceLibrary       PermissionType = "add_voice_from_voice_library"
	PermissionTypeCreateInstantVoiceClone        PermissionType = "create_instant_voice_clone"
	PermissionTypeCreateProfessionalVoiceClone   PermissionType = "create_professional_voice_clone"
	PermissionTypePublishVoiceToVoiceLibrary     PermissionType = "publish_voice_to_voice_library"
	PermissionTypeShareVoiceExternally           PermissionType = "share_voice_externally"
	PermissionTypeCreateUserAPIKey               PermissionType = "create_user_api_key"
	PermissionTypeWorkspaceAnalyticsFullRead     PermissionType = "workspace_analytics_full_read"
)

type WorkspaceAPIKey struct {
	HashedXiAPIKey       string           `json:"hashed_xi_api_key,required"`
	Hint                 string           `json:"hint,required"`
	KeyID                string           `json:"key_id,required"`
	Name                 string           `json:"name,required"`
	ServiceAccountUserID string           `json:"service_account_user_id,required"`
	CharacterCount       int64            `json:"character_count,nullable"`
	CharacterLimit       int64            `json:"character_limit,nullable"`
	CreatedAtUnix        int64            `json:"created_at_unix,nullable"`
	IsDisabled           bool             `json:"is_disabled"`
	Permissions          []PermissionType `json:"permissions,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HashedXiAPIKey       respjson.Field
		Hint                 respjson.Field
		KeyID                respjson.Field
		Name                 respjson.Field
		ServiceAccountUserID respjson.Field
		CharacterCount       respjson.Field
		CharacterLimit       respjson.Field
		CreatedAtUnix        respjson.Field
		IsDisabled           respjson.Field
		Permissions          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceAPIKey) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceAPIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyNewResponse struct {
	KeyID    string `json:"key_id,required"`
	XiAPIKey string `json:"xi-api-key,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		KeyID       respjson.Field
		XiAPIKey    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountAPIKeyNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountAPIKeyNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyUpdateResponse = any

type ServiceAccountAPIKeyListResponse struct {
	APIKeys []WorkspaceAPIKey `json:"api-keys,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		APIKeys     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountAPIKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountAPIKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyDeleteResponse = any

type ServiceAccountAPIKeyNewParams struct {
	Name string `json:"name,required"`
	// The permissions of the XI API.
	Permissions ServiceAccountAPIKeyNewParamsPermissionsUnion `json:"permissions,omitzero,required"`
	// The character limit of the XI API key. If provided this will limit the usage of
	// this api key to n characters per month where n is the chosen value. Requests
	// that incur charges will fail after reaching this monthly limit.
	CharacterLimit param.Opt[int64] `json:"character_limit,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ServiceAccountAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ServiceAccountAPIKeyNewParamsPermissionsUnion struct {
	OfPermissionTypeArray []PermissionType `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ServiceAccountAPIKeyNewParamsPermissionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPermissionTypeArray, u.OfAll)
}
func (u *ServiceAccountAPIKeyNewParamsPermissionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ServiceAccountAPIKeyNewParamsPermissionsUnion) asAny() any {
	if !param.IsOmitted(u.OfPermissionTypeArray) {
		return &u.OfPermissionTypeArray
	} else if !param.IsOmitted(u.OfAll) {
		return &u.OfAll
	}
	return nil
}

type ServiceAccountAPIKeyUpdateParams struct {
	ServiceAccountUserID string `path:"service_account_user_id,required" json:"-"`
	// Whether to enable or disable the API key.
	IsEnabled bool `json:"is_enabled,required"`
	// The name of the XI API key to use (used for identification purposes only).
	Name string `json:"name,required"`
	// The permissions of the XI API.
	Permissions ServiceAccountAPIKeyUpdateParamsPermissionsUnion `json:"permissions,omitzero,required"`
	// The character limit of the XI API key. If provided this will limit the usage of
	// this api key to n characters per month where n is the chosen value. Requests
	// that incur charges will fail after reaching this monthly limit.
	CharacterLimit param.Opt[int64] `json:"character_limit,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ServiceAccountAPIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountAPIKeyUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountAPIKeyUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ServiceAccountAPIKeyUpdateParamsPermissionsUnion struct {
	OfPermissionTypeArray []PermissionType `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.All]()
	OfAll constant.All `json:",omitzero,inline"`
	paramUnion
}

func (u ServiceAccountAPIKeyUpdateParamsPermissionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPermissionTypeArray, u.OfAll)
}
func (u *ServiceAccountAPIKeyUpdateParamsPermissionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ServiceAccountAPIKeyUpdateParamsPermissionsUnion) asAny() any {
	if !param.IsOmitted(u.OfPermissionTypeArray) {
		return &u.OfPermissionTypeArray
	} else if !param.IsOmitted(u.OfAll) {
		return &u.OfAll
	}
	return nil
}

type ServiceAccountAPIKeyListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ServiceAccountAPIKeyDeleteParams struct {
	ServiceAccountUserID string `path:"service_account_user_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
