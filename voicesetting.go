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
	shimjson "github.com/oh-tarnished/elevenlabs-go/internal/encoding/json"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// VoiceSettingService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceSettingService] method instead.
type VoiceSettingService struct {
	Options []option.RequestOption
}

// NewVoiceSettingService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceSettingService(opts ...option.RequestOption) (r VoiceSettingService) {
	r = VoiceSettingService{}
	r.Options = opts
	return
}

// Edit your settings for a specific voice. "similarity_boost" corresponds to
// "Clarity + Similarity Enhancement" in the web app and "stability" corresponds to
// "Stability" slider in the web app.
func (r *VoiceSettingService) Edit(ctx context.Context, voiceID string, params VoiceSettingEditParams, opts ...option.RequestOption) (res *VoiceSettingEditResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s/settings/edit", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns the settings for a specific voice. "similarity_boost" corresponds
// to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds
// to "Stability" slider in the web app.
func (r *VoiceSettingService) Get(ctx context.Context, voiceID string, query VoiceSettingGetParams, opts ...option.RequestOption) (res *VoiceSettings, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s/settings", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets the default settings for voices. "similarity_boost" corresponds
// to"Clarity + Similarity Enhancement" in the web app and "stability" corresponds
// to "Stability" slider in the web app.
func (r *VoiceSettingService) GetDefault(ctx context.Context, opts ...option.RequestOption) (res *VoiceSettings, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/voices/settings/default"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VoiceSettings struct {
	// Determines how closely the AI should adhere to the original voice when
	// attempting to replicate it.
	SimilarityBoost float64 `json:"similarity_boost,nullable"`
	// Adjusts the speed of the voice. A value of 1.0 is the default speed, while
	// values less than 1.0 slow down the speech, and values greater than 1.0 speed it
	// up.
	Speed float64 `json:"speed,nullable"`
	// Determines how stable the voice is and the randomness between each generation.
	// Lower values introduce broader emotional range for the voice. Higher values can
	// result in a monotonous voice with limited emotion.
	Stability float64 `json:"stability,nullable"`
	// Determines the style exaggeration of the voice. This setting attempts to amplify
	// the style of the original speaker. It does consume additional computational
	// resources and might increase latency if set to anything other than 0.
	Style float64 `json:"style,nullable"`
	// This setting boosts the similarity to the original speaker. Using this setting
	// requires a slightly higher computational load, which in turn increases latency.
	UseSpeakerBoost bool `json:"use_speaker_boost,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SimilarityBoost respjson.Field
		Speed           respjson.Field
		Stability       respjson.Field
		Style           respjson.Field
		UseSpeakerBoost respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSettings) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VoiceSettings to a VoiceSettingsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VoiceSettingsParam.Overrides()
func (r VoiceSettings) ToParam() VoiceSettingsParam {
	return param.Override[VoiceSettingsParam](json.RawMessage(r.RawJSON()))
}

type VoiceSettingsParam struct {
	// Determines how closely the AI should adhere to the original voice when
	// attempting to replicate it.
	SimilarityBoost param.Opt[float64] `json:"similarity_boost,omitzero"`
	// Adjusts the speed of the voice. A value of 1.0 is the default speed, while
	// values less than 1.0 slow down the speech, and values greater than 1.0 speed it
	// up.
	Speed param.Opt[float64] `json:"speed,omitzero"`
	// Determines how stable the voice is and the randomness between each generation.
	// Lower values introduce broader emotional range for the voice. Higher values can
	// result in a monotonous voice with limited emotion.
	Stability param.Opt[float64] `json:"stability,omitzero"`
	// Determines the style exaggeration of the voice. This setting attempts to amplify
	// the style of the original speaker. It does consume additional computational
	// resources and might increase latency if set to anything other than 0.
	Style param.Opt[float64] `json:"style,omitzero"`
	// This setting boosts the similarity to the original speaker. Using this setting
	// requires a slightly higher computational load, which in turn increases latency.
	UseSpeakerBoost param.Opt[bool] `json:"use_speaker_boost,omitzero"`
	paramObj
}

func (r VoiceSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow VoiceSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingEditResponse struct {
	// The status of the voice settings edit request. If the request was successful,
	// the status will be 'ok'. Otherwise an error message with status 500 will be
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
func (r VoiceSettingEditResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceSettingEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSettingEditParams struct {
	// The settings for a specific voice.
	VoiceSettings VoiceSettingsParam
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoiceSettingEditParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VoiceSettings)
}
func (r *VoiceSettingEditParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.VoiceSettings)
}

type VoiceSettingGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
