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

// VoiceGenerationGenerateVoiceService contains methods and other services that
// help with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceGenerationGenerateVoiceService] method instead.
type VoiceGenerationGenerateVoiceService struct {
	Options []option.RequestOption
}

// NewVoiceGenerationGenerateVoiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVoiceGenerationGenerateVoiceService(opts ...option.RequestOption) (r VoiceGenerationGenerateVoiceService) {
	r = VoiceGenerationGenerateVoiceService{}
	r.Options = opts
	return
}

// Generate a random voice based on parameters. This method returns a
// generated_voice_id in the response header, and a sample of the voice in the
// body. If you like the generated voice call /v1/voice-generation/create-voice
// with the generated_voice_id to create the voice.
func (r *VoiceGenerationGenerateVoiceService) NewRandomVoice(ctx context.Context, params VoiceGenerationGenerateVoiceNewRandomVoiceParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/voice-generation/generate-voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get possible parameters for the /v1/voice-generation/generate-voice endpoint.
func (r *VoiceGenerationGenerateVoiceService) GetParameters(ctx context.Context, opts ...option.RequestOption) (res *VoiceGenerationGenerateVoiceGetParametersResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/voice-generation/generate-voice/parameters"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ParameterOptionResp struct {
	Code string `json:"code,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ParameterOptionResp) RawJSON() string { return r.JSON.raw }
func (r *ParameterOptionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceGenerationGenerateVoiceGetParametersResponse struct {
	Accents               []ParameterOptionResp `json:"accents,required"`
	Ages                  []ParameterOptionResp `json:"ages,required"`
	Genders               []ParameterOptionResp `json:"genders,required"`
	MaximumAccentStrength float64               `json:"maximum_accent_strength,required"`
	MaximumCharacters     int64                 `json:"maximum_characters,required"`
	MinimumAccentStrength float64               `json:"minimum_accent_strength,required"`
	MinimumCharacters     int64                 `json:"minimum_characters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accents               respjson.Field
		Ages                  respjson.Field
		Genders               respjson.Field
		MaximumAccentStrength respjson.Field
		MaximumCharacters     respjson.Field
		MinimumAccentStrength respjson.Field
		MinimumCharacters     respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceGenerationGenerateVoiceGetParametersResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceGenerationGenerateVoiceGetParametersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceGenerationGenerateVoiceNewRandomVoiceParams struct {
	// Category code corresponding to the accent of the generated voice. Possible
	// values: british, american, african, australian, indian.
	Accent string `json:"accent,required"`
	// The strength of the accent of the generated voice. Has to be between 0.3 and
	// 2.0.
	AccentStrength float64 `json:"accent_strength,required"`
	// Category code corresponding to the age of the generated voice. Possible values:
	// young, middle_aged, old.
	//
	// Any of "young", "middle_aged", "old".
	Age VoiceGenerationGenerateVoiceNewRandomVoiceParamsAge `json:"age,omitzero,required"`
	// Category code corresponding to the gender of the generated voice. Possible
	// values: female, male.
	//
	// Any of "female", "male".
	Gender VoiceGenerationGenerateVoiceNewRandomVoiceParamsGender `json:"gender,omitzero,required"`
	// Text to generate, text length has to be between 100 and 1000.
	Text string `json:"text,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoiceGenerationGenerateVoiceNewRandomVoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceGenerationGenerateVoiceNewRandomVoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceGenerationGenerateVoiceNewRandomVoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category code corresponding to the age of the generated voice. Possible values:
// young, middle_aged, old.
type VoiceGenerationGenerateVoiceNewRandomVoiceParamsAge string

const (
	VoiceGenerationGenerateVoiceNewRandomVoiceParamsAgeYoung      VoiceGenerationGenerateVoiceNewRandomVoiceParamsAge = "young"
	VoiceGenerationGenerateVoiceNewRandomVoiceParamsAgeMiddleAged VoiceGenerationGenerateVoiceNewRandomVoiceParamsAge = "middle_aged"
	VoiceGenerationGenerateVoiceNewRandomVoiceParamsAgeOld        VoiceGenerationGenerateVoiceNewRandomVoiceParamsAge = "old"
)

// Category code corresponding to the gender of the generated voice. Possible
// values: female, male.
type VoiceGenerationGenerateVoiceNewRandomVoiceParamsGender string

const (
	VoiceGenerationGenerateVoiceNewRandomVoiceParamsGenderFemale VoiceGenerationGenerateVoiceNewRandomVoiceParamsGender = "female"
	VoiceGenerationGenerateVoiceNewRandomVoiceParamsGenderMale   VoiceGenerationGenerateVoiceNewRandomVoiceParamsGender = "male"
)
