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

// ModelService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Gets a list of available models.
func (r *ModelService) List(ctx context.Context, query ModelListParams, opts ...option.RequestOption) (res *[]ModelListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelListResponse struct {
	// Whether the model can be finetuned.
	CanBeFinetuned bool `json:"can_be_finetuned,required"`
	// Whether the model can do text-to-speech.
	CanDoTextToSpeech bool `json:"can_do_text_to_speech,required"`
	// Whether the model can do voice conversion.
	CanDoVoiceConversion bool `json:"can_do_voice_conversion,required"`
	// Whether the model can use speaker boost.
	CanUseSpeakerBoost bool `json:"can_use_speaker_boost,required"`
	// Whether the model can use style.
	CanUseStyle bool `json:"can_use_style,required"`
	// The concurrency group for the model.
	ConcurrencyGroup string `json:"concurrency_group,required"`
	// The description of the model.
	Description string `json:"description,required"`
	// The languages supported by the model.
	Languages []ModelListResponseLanguage `json:"languages,required"`
	// The maximum number of characters that can be requested by a free user.
	MaxCharactersRequestFreeUser int64 `json:"max_characters_request_free_user,required"`
	// The maximum number of characters that can be requested by a subscribed user.
	MaxCharactersRequestSubscribedUser int64 `json:"max_characters_request_subscribed_user,required"`
	// The maximum length of text that can be requested for this model.
	MaximumTextLengthPerRequest int64 `json:"maximum_text_length_per_request,required"`
	// The unique identifier of the model.
	ModelID string `json:"model_id,required"`
	// The rates for the model.
	ModelRates ModelListResponseModelRates `json:"model_rates,required"`
	// The name of the model.
	Name string `json:"name,required"`
	// Whether the model requires alpha access.
	RequiresAlphaAccess bool `json:"requires_alpha_access,required"`
	// Whether the model serves pro voices.
	ServesProVoices bool `json:"serves_pro_voices,required"`
	// The cost factor for the model.
	TokenCostFactor float64 `json:"token_cost_factor,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeFinetuned                     respjson.Field
		CanDoTextToSpeech                  respjson.Field
		CanDoVoiceConversion               respjson.Field
		CanUseSpeakerBoost                 respjson.Field
		CanUseStyle                        respjson.Field
		ConcurrencyGroup                   respjson.Field
		Description                        respjson.Field
		Languages                          respjson.Field
		MaxCharactersRequestFreeUser       respjson.Field
		MaxCharactersRequestSubscribedUser respjson.Field
		MaximumTextLengthPerRequest        respjson.Field
		ModelID                            respjson.Field
		ModelRates                         respjson.Field
		Name                               respjson.Field
		RequiresAlphaAccess                respjson.Field
		ServesProVoices                    respjson.Field
		TokenCostFactor                    respjson.Field
		ExtraFields                        map[string]respjson.Field
		raw                                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponse) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListResponseLanguage struct {
	// The unique identifier of the language.
	LanguageID string `json:"language_id,required"`
	// The name of the language.
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageID  respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseLanguage) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The rates for the model.
type ModelListResponseModelRates struct {
	// The cost multiplier for characters.
	CharacterCostMultiplier float64 `json:"character_cost_multiplier,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterCostMultiplier respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelListResponseModelRates) RawJSON() string { return r.JSON.raw }
func (r *ModelListResponseModelRates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
