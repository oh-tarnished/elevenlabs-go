// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// VoiceService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceService] method instead.
type VoiceService struct {
	Options  []option.RequestOption
	Samples  VoiceSampleService
	Settings VoiceSettingService
	Add      VoiceAddService
	Pvc      VoicePvcService
}

// NewVoiceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVoiceService(opts ...option.RequestOption) (r VoiceService) {
	r = VoiceService{}
	r.Options = opts
	r.Samples = NewVoiceSampleService(opts...)
	r.Settings = NewVoiceSettingService(opts...)
	r.Add = NewVoiceAddService(opts...)
	r.Pvc = NewVoicePvcService(opts...)
	return
}

// Returns metadata about a specific voice.
func (r *VoiceService) Get(ctx context.Context, voiceID string, params VoiceGetParams, opts ...option.RequestOption) (res *Voice, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Returns a list of all available voices for a user.
func (r *VoiceService) List(ctx context.Context, params VoiceListParams, opts ...option.RequestOption) (res *VoiceListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes a voice by its ID.
func (r *VoiceService) Delete(ctx context.Context, voiceID string, body VoiceDeleteParams, opts ...option.RequestOption) (res *VoiceDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Edit a voice created by you.
func (r *VoiceService) Edit(ctx context.Context, voiceID string, params VoiceEditParams, opts ...option.RequestOption) (res *VoiceEditResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s/edit", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets a list of all available voices for a user with search, filtering and
// pagination.
func (r *VoiceService) GetV2(ctx context.Context, params VoiceGetV2Params, opts ...option.RequestOption) (res *VoiceGetV2Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v2/voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type VoiceListResponse struct {
	// A list of available voices.
	Voices []Voice `json:"voices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceListResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceDeleteResponse struct {
	// The status of the voice deletion request. If the request was successful, the
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
func (r VoiceDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceEditResponse struct {
	// The status of the voice edit request. If the request was successful, the status
	// will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceEditResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceGetV2Response struct {
	// Indicates whether there are more voices available in subsequent pages. Use this
	// flag (and next_page_token) for reliable pagination instead of relying on
	// total_count.
	HasMore bool `json:"has_more,required"`
	// The total count of voices matching the query. This value is a live snapshot that
	// reflects the current state of the database and may change between requests as
	// users create, modify, or delete voices. For reliable pagination, use the
	// has_more flag instead of relying on this value. Only request this field when you
	// actually need the total count (e.g., for display purposes), as calculating it
	// incurs a performance cost.
	TotalCount int64 `json:"total_count,required"`
	// The list of voices matching the query.
	Voices []Voice `json:"voices,required"`
	// Token to retrieve the next page of results. Pass this value to the next request
	// to continue pagination. Null if there are no more results.
	NextPageToken string `json:"next_page_token,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore       respjson.Field
		TotalCount    respjson.Field
		Voices        respjson.Field
		NextPageToken respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceGetV2Response) RawJSON() string { return r.JSON.raw }
func (r *VoiceGetV2Response) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceGetParams struct {
	// This parameter is now deprecated. It is ignored and will be removed in a future
	// version.
	WithSettings param.Opt[bool] `query:"with_settings,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceGetParams]'s query parameters as `url.Values`.
func (r VoiceGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceListParams struct {
	// If set to true, legacy premade voices will be included in responses from
	// /v1/voices
	ShowLegacy param.Opt[bool] `query:"show_legacy,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceListParams]'s query parameters as `url.Values`.
func (r VoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoiceDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type VoiceEditParams struct {
	// The name that identifies this voice. This will be displayed in the dropdown of
	// the website.
	Name string `json:"name,required"`
	// A description of the voice.
	Description param.Opt[string] `json:"description,omitzero"`
	// If set will remove background noise for voice samples using our audio isolation
	// model. If the samples do not include background noise, it can make the quality
	// worse.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels VoiceEditParamsLabelsUnion `json:"labels,omitzero"`
	// Audio files to add to the voice
	Files []io.Reader `json:"files,omitzero" format:"binary"`
	paramObj
}

func (r VoiceEditParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VoiceEditParamsLabelsUnion struct {
	OfStringMap map[string]string `json:",omitzero,inline"`
	OfString    param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VoiceEditParamsLabelsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringMap, u.OfString)
}
func (u *VoiceEditParamsLabelsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VoiceEditParamsLabelsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringMap) {
		return &u.OfStringMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type VoiceGetV2Params struct {
	// Category of the voice to filter by. One of 'premade', 'cloned', 'generated',
	// 'professional'
	Category param.Opt[string] `query:"category,omitzero" json:"-"`
	// Collection ID to filter voices by.
	CollectionID param.Opt[string] `query:"collection_id,omitzero" json:"-"`
	// State of the voice's fine tuning to filter by. Applicable only to professional
	// voices clones. One of 'draft', 'not_verified', 'not_started', 'queued',
	// 'fine_tuning', 'fine_tuned', 'failed', 'delayed'
	FineTuningState param.Opt[string] `query:"fine_tuning_state,omitzero" json:"-"`
	// The next page token to use for pagination. Returned from the previous request.
	// Use this in combination with the has_more flag for reliable pagination.
	NextPageToken param.Opt[string] `query:"next_page_token,omitzero" json:"-"`
	// Search term to filter voices by. Searches in name, description, labels,
	// category.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Which field to sort by, one of 'created_at_unix' or 'name'. 'created_at_unix'
	// may not be available for older voices.
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Which direction to sort the voices in. 'asc' or 'desc'.
	SortDirection param.Opt[string] `query:"sort_direction,omitzero" json:"-"`
	// Type of the voice to filter by. One of 'personal', 'community', 'default',
	// 'workspace', 'non-default', 'saved'. 'non-default' is equal to all but
	// 'default'. 'saved' is equal to non-default, but includes default voices if they
	// have been added to a collection.
	VoiceType param.Opt[string] `query:"voice_type,omitzero" json:"-"`
	// Whether to include the total count of voices found in the response. NOTE: The
	// total_count value is a live snapshot and may change between requests as users
	// create, modify, or delete voices. For pagination, rely on the has_more flag
	// instead. Only enable this when you actually need the total count (e.g., for
	// display purposes), as it incurs a performance cost.
	IncludeTotalCount param.Opt[bool] `query:"include_total_count,omitzero" json:"-"`
	// How many voices to return at maximum. Can not exceed 100, defaults to 10. Page 0
	// may include more voices due to default voices being included.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Voice IDs to lookup by. Maximum 100 voice IDs.
	VoiceIDs []string `query:"voice_ids,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoiceGetV2Params]'s query parameters as `url.Values`.
func (r VoiceGetV2Params) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
