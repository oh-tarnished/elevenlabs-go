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

// TextToDialogueStreamService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextToDialogueStreamService] method instead.
type TextToDialogueStreamService struct {
	Options []option.RequestOption
}

// NewTextToDialogueStreamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTextToDialogueStreamService(opts ...option.RequestOption) (r TextToDialogueStreamService) {
	r = TextToDialogueStreamService{}
	r.Options = opts
	return
}

// Converts a list of text and voice ID pairs into speech (dialogue) and returns an
// audio stream.
func (r *TextToDialogueStreamService) New(ctx context.Context, params TextToDialogueStreamNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/text-to-dialogue/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Converts a list of text and voice ID pairs into speech (dialogue) and returns a
// stream of JSON blobs containing audio as a base64 encoded string and timestamps
func (r *TextToDialogueStreamService) NewWithTimestamps(ctx context.Context, params TextToDialogueStreamNewWithTimestampsParams, opts ...option.RequestOption) (res *TextToDialogueStreamNewWithTimestampsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/text-to-dialogue/stream/with-timestamps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type VoiceSegment struct {
	// End index in the characters array (exclusive)
	CharacterEndIndex int64 `json:"character_end_index,required"`
	// Start index in the characters array
	CharacterStartIndex int64 `json:"character_start_index,required"`
	// Line of the dialogue (script) that this segment is a part of.
	DialogueInputIndex int64 `json:"dialogue_input_index,required"`
	// End time of this voice segment
	EndTimeSeconds float64 `json:"end_time_seconds,required"`
	// Start time of this voice segment
	StartTimeSeconds float64 `json:"start_time_seconds,required"`
	// The voice ID used for this segment
	VoiceID string `json:"voice_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterEndIndex   respjson.Field
		CharacterStartIndex respjson.Field
		DialogueInputIndex  respjson.Field
		EndTimeSeconds      respjson.Field
		StartTimeSeconds    respjson.Field
		VoiceID             respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSegment) RawJSON() string { return r.JSON.raw }
func (r *VoiceSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToDialogueStreamNewWithTimestampsResponse struct {
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
func (r TextToDialogueStreamNewWithTimestampsResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToDialogueStreamNewWithTimestampsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToDialogueStreamNewParams struct {
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
	ApplyTextNormalization TextToDialogueStreamNewParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Settings controlling the dialogue generation.
	Settings ModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r TextToDialogueStreamNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToDialogueStreamNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToDialogueStreamNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToDialogueStreamNewParams]'s query parameters as
// `url.Values`.
func (r TextToDialogueStreamNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToDialogueStreamNewParamsApplyTextNormalization string

const (
	TextToDialogueStreamNewParamsApplyTextNormalizationAuto TextToDialogueStreamNewParamsApplyTextNormalization = "auto"
	TextToDialogueStreamNewParamsApplyTextNormalizationOn   TextToDialogueStreamNewParamsApplyTextNormalization = "on"
	TextToDialogueStreamNewParamsApplyTextNormalizationOff  TextToDialogueStreamNewParamsApplyTextNormalization = "off"
)

type TextToDialogueStreamNewWithTimestampsParams struct {
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
	ApplyTextNormalization TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Settings controlling the dialogue generation.
	Settings ModelSettingsParam `json:"settings,omitzero"`
	paramObj
}

func (r TextToDialogueStreamNewWithTimestampsParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToDialogueStreamNewWithTimestampsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToDialogueStreamNewWithTimestampsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToDialogueStreamNewWithTimestampsParams]'s query
// parameters as `url.Values`.
func (r TextToDialogueStreamNewWithTimestampsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalization string

const (
	TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalizationAuto TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalization = "auto"
	TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalizationOn   TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalization = "on"
	TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalizationOff  TextToDialogueStreamNewWithTimestampsParamsApplyTextNormalization = "off"
)
