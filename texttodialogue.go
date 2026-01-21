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

// TextToDialogueService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextToDialogueService] method instead.
type TextToDialogueService struct {
	Options []option.RequestOption
	Stream  TextToDialogueStreamService
}

// NewTextToDialogueService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTextToDialogueService(opts ...option.RequestOption) (r TextToDialogueService) {
	r = TextToDialogueService{}
	r.Options = opts
	r.Stream = NewTextToDialogueStreamService(opts...)
	return
}

// Converts a list of text and voice ID pairs into speech (dialogue) and returns
// audio.
func (r *TextToDialogueService) New(ctx context.Context, params TextToDialogueNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/text-to-dialogue"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Generate dialogue from text with precise character-level timing information for
// audio-text synchronization.
func (r *TextToDialogueService) NewWithTimestamps(ctx context.Context, params TextToDialogueNewWithTimestampsParams, opts ...option.RequestOption) (res *TextToDialogueNewWithTimestampsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/text-to-dialogue/with-timestamps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// The properties Text, VoiceID are required.
type DialogueInputParam struct {
	// The text to be converted into speech.
	Text string `json:"text,required"`
	// The ID of the voice to be used for the generation.
	VoiceID string `json:"voice_id,required"`
	paramObj
}

func (r DialogueInputParam) MarshalJSON() (data []byte, err error) {
	type shadow DialogueInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DialogueInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelSettingsParam struct {
	// Determines how stable the voice is and the randomness between each generation.
	// Lower values introduce broader emotional range for the voice. Higher values can
	// result in a monotonous voice with limited emotion.
	Stability param.Opt[float64] `json:"stability,omitzero"`
	paramObj
}

func (r ModelSettingsParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelSettingsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSettingsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToDialogueNewWithTimestampsResponse struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64,required"`
	// Voice segments for the audio
	VoiceSegments []VoiceSegment `json:"voice_segments,required"`
	// Timestamp information for each character in the original text
	Alignment CharacterAlignmentResponse `json:"alignment,nullable"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment CharacterAlignmentResponse `json:"normalized_alignment,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64         respjson.Field
		VoiceSegments       respjson.Field
		Alignment           respjson.Field
		NormalizedAlignment respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToDialogueNewWithTimestampsResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToDialogueNewWithTimestampsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToDialogueNewParams struct {
	// A list of dialogue inputs, each containing text and a voice ID which will be
	// converted into speech.
	Inputs []DialogueInputParam `json:"inputs,omitzero,required"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to
	// the text. They will be applied in order. You may have up to 3 locators per
	// request
	PronunciationDictionaryLocators []PronunciationDictionaryVersionLocatorParam `json:"pronunciation_dictionary_locators,omitzero"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
	// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
	// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
	// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
	// Note that the μ-law format (sometimes written mu-law, often approximated as
	// u-law) is commonly used for Twilio audio inputs.
	//
	// Any of "mp3_22050_32", "mp3_24000_48", "mp3_44100_32", "mp3_44100_64",
	// "mp3_44100_96", "mp3_44100_128", "mp3_44100_192", "pcm_8000", "pcm_16000",
	// "pcm_22050", "pcm_24000", "pcm_32000", "pcm_44100", "pcm_48000", "ulaw_8000",
	// "alaw_8000", "opus_48000_32", "opus_48000_64", "opus_48000_96",
	// "opus_48000_128", "opus_48000_192".
	OutputFormat AllowedOutputFormats `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToDialogueNewParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Settings controlling the dialogue generation.
	Settings ModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r TextToDialogueNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToDialogueNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToDialogueNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToDialogueNewParams]'s query parameters as
// `url.Values`.
func (r TextToDialogueNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToDialogueNewParamsApplyTextNormalization string

const (
	TextToDialogueNewParamsApplyTextNormalizationAuto TextToDialogueNewParamsApplyTextNormalization = "auto"
	TextToDialogueNewParamsApplyTextNormalizationOn   TextToDialogueNewParamsApplyTextNormalization = "on"
	TextToDialogueNewParamsApplyTextNormalizationOff  TextToDialogueNewParamsApplyTextNormalization = "off"
)

type TextToDialogueNewWithTimestampsParams struct {
	// A list of dialogue inputs, each containing text and a voice ID which will be
	// converted into speech.
	Inputs []DialogueInputParam `json:"inputs,omitzero,required"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to
	// the text. They will be applied in order. You may have up to 3 locators per
	// request
	PronunciationDictionaryLocators []PronunciationDictionaryVersionLocatorParam `json:"pronunciation_dictionary_locators,omitzero"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
	// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
	// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
	// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
	// Note that the μ-law format (sometimes written mu-law, often approximated as
	// u-law) is commonly used for Twilio audio inputs.
	//
	// Any of "mp3_22050_32", "mp3_24000_48", "mp3_44100_32", "mp3_44100_64",
	// "mp3_44100_96", "mp3_44100_128", "mp3_44100_192", "pcm_8000", "pcm_16000",
	// "pcm_22050", "pcm_24000", "pcm_32000", "pcm_44100", "pcm_48000", "ulaw_8000",
	// "alaw_8000", "opus_48000_32", "opus_48000_64", "opus_48000_96",
	// "opus_48000_128", "opus_48000_192".
	OutputFormat AllowedOutputFormats `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToDialogueNewWithTimestampsParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Settings controlling the dialogue generation.
	Settings ModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r TextToDialogueNewWithTimestampsParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToDialogueNewWithTimestampsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToDialogueNewWithTimestampsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToDialogueNewWithTimestampsParams]'s query parameters
// as `url.Values`.
func (r TextToDialogueNewWithTimestampsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToDialogueNewWithTimestampsParamsApplyTextNormalization string

const (
	TextToDialogueNewWithTimestampsParamsApplyTextNormalizationAuto TextToDialogueNewWithTimestampsParamsApplyTextNormalization = "auto"
	TextToDialogueNewWithTimestampsParamsApplyTextNormalizationOn   TextToDialogueNewWithTimestampsParamsApplyTextNormalization = "on"
	TextToDialogueNewWithTimestampsParamsApplyTextNormalizationOff  TextToDialogueNewWithTimestampsParamsApplyTextNormalization = "off"
)
