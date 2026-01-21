// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
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

// TextToSpeechService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextToSpeechService] method instead.
type TextToSpeechService struct {
	Options []option.RequestOption
}

// NewTextToSpeechService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTextToSpeechService(opts ...option.RequestOption) (r TextToSpeechService) {
	r = TextToSpeechService{}
	r.Options = opts
	return
}

// Converts text into speech using a voice of your choice and returns audio.
func (r *TextToSpeechService) Convert(ctx context.Context, voiceID string, params TextToSpeechConvertParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-speech/%s", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Generate speech from text with precise character-level timing information for
// audio-text synchronization.
func (r *TextToSpeechService) ConvertWithTimestamps(ctx context.Context, voiceID string, params TextToSpeechConvertWithTimestampsParams, opts ...option.RequestOption) (res *TextToSpeechConvertWithTimestampsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-speech/%s/with-timestamps", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Converts text into speech using a voice of your choice and returns audio as an
// audio stream.
func (r *TextToSpeechService) StreamConvert(ctx context.Context, voiceID string, params TextToSpeechStreamConvertParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-speech/%s/stream", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Converts text into speech using a voice of your choice and returns a stream of
// JSONs containing audio as a base64 encoded string together with information on
// when which character was spoken.
func (r *TextToSpeechService) StreamConvertWithTimestamps(ctx context.Context, voiceID string, params TextToSpeechStreamConvertWithTimestampsParams, opts ...option.RequestOption) (res *TextToSpeechStreamConvertWithTimestampsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-speech/%s/stream/with-timestamps", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type CharacterAlignmentResponse struct {
	CharacterEndTimesSeconds   []float64 `json:"character_end_times_seconds,required"`
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds,required"`
	Characters                 []string  `json:"characters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterEndTimesSeconds   respjson.Field
		CharacterStartTimesSeconds respjson.Field
		Characters                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CharacterAlignmentResponse) RawJSON() string { return r.JSON.raw }
func (r *CharacterAlignmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property PronunciationDictionaryID is required.
type PronunciationDictionaryVersionLocatorParam struct {
	// The ID of the pronunciation dictionary.
	PronunciationDictionaryID string `json:"pronunciation_dictionary_id,required"`
	// The ID of the version of the pronunciation dictionary. If not provided, the
	// latest version will be used.
	VersionID param.Opt[string] `json:"version_id,omitzero"`
	paramObj
}

func (r PronunciationDictionaryVersionLocatorParam) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryVersionLocatorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryVersionLocatorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechConvertWithTimestampsResponse struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64,required"`
	// Timestamp information for each character in the original text
	Alignment CharacterAlignmentResponse `json:"alignment,nullable"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment CharacterAlignmentResponse `json:"normalized_alignment,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64         respjson.Field
		Alignment           respjson.Field
		NormalizedAlignment respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechConvertWithTimestampsResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechConvertWithTimestampsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechStreamConvertWithTimestampsResponse struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64,required"`
	// Timestamp information for each character in the original text
	Alignment CharacterAlignmentResponse `json:"alignment,nullable"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment CharacterAlignmentResponse `json:"normalized_alignment,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64         respjson.Field
		Alignment           respjson.Field
		NormalizedAlignment respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TextToSpeechStreamConvertWithTimestampsResponse) RawJSON() string { return r.JSON.raw }
func (r *TextToSpeechStreamConvertWithTimestampsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToSpeechConvertParams struct {
	// The text that will get converted into speech.
	Text string `json:"text,required"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// The text that comes after the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	NextText param.Opt[string] `json:"next_text,omitzero"`
	// The text that came before the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	PreviousText param.Opt[string] `json:"previous_text,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// This parameter controls language text normalization. This helps with proper
	// pronunciation of text in some supported languages. WARNING: This parameter can
	// heavily increase the latency of the request. Currently only supported for
	// Japanese.
	ApplyLanguageTextNormalization param.Opt[bool] `json:"apply_language_text_normalization,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If true, we won't use PVC version of the voice for the generation but the IVC
	// version. This is a temporary workaround for higher latency in PVC versions.
	UsePvcAsIvc param.Opt[bool] `json:"use_pvc_as_ivc,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of request_id of the samples that come after this generation.
	// next_request_ids is especially useful for maintaining the speech's continuity
	// when regenerating a sample that has had some audio quality issues. For example,
	// if you have generated 3 speech clips, and you want to improve clip 2, passing
	// the request id of clip 3 as a next_request_id (and that of clip 1 as a
	// previous_request_id) will help maintain natural flow in the combined speech. The
	// results will be best when the same model is used across the generations. In case
	// both next_text and next_request_ids is send, next_text will be ignored. A
	// maximum of 3 request_ids can be send.
	NextRequestIDs []string `json:"next_request_ids,omitzero"`
	// A list of request_id of the samples that were generated before this generation.
	// Can be used to improve the speech's continuity when splitting up a large task
	// into multiple requests. The results will be best when the same model is used
	// across the generations. In case both previous_text and previous_request_ids is
	// send, previous_text will be ignored. A maximum of 3 request_ids can be send.
	PreviousRequestIDs []string `json:"previous_request_ids,omitzero"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to
	// the text. They will be applied in order. You may have up to 3 locators per
	// request
	PronunciationDictionaryLocators []PronunciationDictionaryVersionLocatorParam `json:"pronunciation_dictionary_locators,omitzero"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
	// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
	// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
	// and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro
	// tier or above. Note that the μ-law format (sometimes written mu-law, often
	// approximated as u-law) is commonly used for Twilio audio inputs.
	//
	// Any of "alaw_8000", "mp3_22050_32", "mp3_24000_48", "mp3_44100_128",
	// "mp3_44100_192", "mp3_44100_32", "mp3_44100_64", "mp3_44100_96",
	// "opus_48000_128", "opus_48000_192", "opus_48000_32", "opus_48000_64",
	// "opus_48000_96", "pcm_16000", "pcm_22050", "pcm_24000", "pcm_32000",
	// "pcm_44100", "pcm_48000", "pcm_8000", "ulaw_8000", "wav_16000", "wav_22050",
	// "wav_24000", "wav_32000", "wav_44100", "wav_48000", "wav_8000".
	OutputFormat TextToSpeechConvertParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToSpeechConvertParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request.
	VoiceSettings VoiceSettingsParam `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechConvertParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechConvertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechConvertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToSpeechConvertParams]'s query parameters as
// `url.Values`.
func (r TextToSpeechConvertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro
// tier or above. Note that the μ-law format (sometimes written mu-law, often
// approximated as u-law) is commonly used for Twilio audio inputs.
type TextToSpeechConvertParamsOutputFormat string

const (
	TextToSpeechConvertParamsOutputFormatAlaw8000      TextToSpeechConvertParamsOutputFormat = "alaw_8000"
	TextToSpeechConvertParamsOutputFormatMP3_22050_32  TextToSpeechConvertParamsOutputFormat = "mp3_22050_32"
	TextToSpeechConvertParamsOutputFormatMP3_24000_48  TextToSpeechConvertParamsOutputFormat = "mp3_24000_48"
	TextToSpeechConvertParamsOutputFormatMP3_44100_128 TextToSpeechConvertParamsOutputFormat = "mp3_44100_128"
	TextToSpeechConvertParamsOutputFormatMP3_44100_192 TextToSpeechConvertParamsOutputFormat = "mp3_44100_192"
	TextToSpeechConvertParamsOutputFormatMP3_44100_32  TextToSpeechConvertParamsOutputFormat = "mp3_44100_32"
	TextToSpeechConvertParamsOutputFormatMP3_44100_64  TextToSpeechConvertParamsOutputFormat = "mp3_44100_64"
	TextToSpeechConvertParamsOutputFormatMP3_44100_96  TextToSpeechConvertParamsOutputFormat = "mp3_44100_96"
	TextToSpeechConvertParamsOutputFormatOpus48000_128 TextToSpeechConvertParamsOutputFormat = "opus_48000_128"
	TextToSpeechConvertParamsOutputFormatOpus48000_192 TextToSpeechConvertParamsOutputFormat = "opus_48000_192"
	TextToSpeechConvertParamsOutputFormatOpus48000_32  TextToSpeechConvertParamsOutputFormat = "opus_48000_32"
	TextToSpeechConvertParamsOutputFormatOpus48000_64  TextToSpeechConvertParamsOutputFormat = "opus_48000_64"
	TextToSpeechConvertParamsOutputFormatOpus48000_96  TextToSpeechConvertParamsOutputFormat = "opus_48000_96"
	TextToSpeechConvertParamsOutputFormatPcm16000      TextToSpeechConvertParamsOutputFormat = "pcm_16000"
	TextToSpeechConvertParamsOutputFormatPcm22050      TextToSpeechConvertParamsOutputFormat = "pcm_22050"
	TextToSpeechConvertParamsOutputFormatPcm24000      TextToSpeechConvertParamsOutputFormat = "pcm_24000"
	TextToSpeechConvertParamsOutputFormatPcm32000      TextToSpeechConvertParamsOutputFormat = "pcm_32000"
	TextToSpeechConvertParamsOutputFormatPcm44100      TextToSpeechConvertParamsOutputFormat = "pcm_44100"
	TextToSpeechConvertParamsOutputFormatPcm48000      TextToSpeechConvertParamsOutputFormat = "pcm_48000"
	TextToSpeechConvertParamsOutputFormatPcm8000       TextToSpeechConvertParamsOutputFormat = "pcm_8000"
	TextToSpeechConvertParamsOutputFormatUlaw8000      TextToSpeechConvertParamsOutputFormat = "ulaw_8000"
	TextToSpeechConvertParamsOutputFormatWav16000      TextToSpeechConvertParamsOutputFormat = "wav_16000"
	TextToSpeechConvertParamsOutputFormatWav22050      TextToSpeechConvertParamsOutputFormat = "wav_22050"
	TextToSpeechConvertParamsOutputFormatWav24000      TextToSpeechConvertParamsOutputFormat = "wav_24000"
	TextToSpeechConvertParamsOutputFormatWav32000      TextToSpeechConvertParamsOutputFormat = "wav_32000"
	TextToSpeechConvertParamsOutputFormatWav44100      TextToSpeechConvertParamsOutputFormat = "wav_44100"
	TextToSpeechConvertParamsOutputFormatWav48000      TextToSpeechConvertParamsOutputFormat = "wav_48000"
	TextToSpeechConvertParamsOutputFormatWav8000       TextToSpeechConvertParamsOutputFormat = "wav_8000"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToSpeechConvertParamsApplyTextNormalization string

const (
	TextToSpeechConvertParamsApplyTextNormalizationAuto TextToSpeechConvertParamsApplyTextNormalization = "auto"
	TextToSpeechConvertParamsApplyTextNormalizationOn   TextToSpeechConvertParamsApplyTextNormalization = "on"
	TextToSpeechConvertParamsApplyTextNormalizationOff  TextToSpeechConvertParamsApplyTextNormalization = "off"
)

type TextToSpeechConvertWithTimestampsParams struct {
	// The text that will get converted into speech.
	Text string `json:"text,required"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// The text that comes after the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	NextText param.Opt[string] `json:"next_text,omitzero"`
	// The text that came before the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	PreviousText param.Opt[string] `json:"previous_text,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// This parameter controls language text normalization. This helps with proper
	// pronunciation of text in some supported languages. WARNING: This parameter can
	// heavily increase the latency of the request. Currently only supported for
	// Japanese.
	ApplyLanguageTextNormalization param.Opt[bool] `json:"apply_language_text_normalization,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If true, we won't use PVC version of the voice for the generation but the IVC
	// version. This is a temporary workaround for higher latency in PVC versions.
	UsePvcAsIvc param.Opt[bool] `json:"use_pvc_as_ivc,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
	// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
	// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
	// and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro
	// tier or above. Note that the μ-law format (sometimes written mu-law, often
	// approximated as u-law) is commonly used for Twilio audio inputs.
	//
	// Any of "alaw_8000", "mp3_22050_32", "mp3_24000_48", "mp3_44100_128",
	// "mp3_44100_192", "mp3_44100_32", "mp3_44100_64", "mp3_44100_96",
	// "opus_48000_128", "opus_48000_192", "opus_48000_32", "opus_48000_64",
	// "opus_48000_96", "pcm_16000", "pcm_22050", "pcm_24000", "pcm_32000",
	// "pcm_44100", "pcm_48000", "pcm_8000", "ulaw_8000", "wav_16000", "wav_22050",
	// "wav_24000", "wav_32000", "wav_44100", "wav_48000", "wav_8000".
	OutputFormat TextToSpeechConvertWithTimestampsParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToSpeechConvertWithTimestampsParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// A list of request_id of the samples that come after this generation.
	// next_request_ids is especially useful for maintaining the speech's continuity
	// when regenerating a sample that has had some audio quality issues. For example,
	// if you have generated 3 speech clips, and you want to improve clip 2, passing
	// the request id of clip 3 as a next_request_id (and that of clip 1 as a
	// previous_request_id) will help maintain natural flow in the combined speech. The
	// results will be best when the same model is used across the generations. In case
	// both next_text and next_request_ids is send, next_text will be ignored. A
	// maximum of 3 request_ids can be send.
	NextRequestIDs []string `json:"next_request_ids,omitzero"`
	// A list of request_id of the samples that were generated before this generation.
	// Can be used to improve the speech's continuity when splitting up a large task
	// into multiple requests. The results will be best when the same model is used
	// across the generations. In case both previous_text and previous_request_ids is
	// send, previous_text will be ignored. A maximum of 3 request_ids can be send.
	PreviousRequestIDs []string `json:"previous_request_ids,omitzero"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to
	// the text. They will be applied in order. You may have up to 3 locators per
	// request
	PronunciationDictionaryLocators []PronunciationDictionaryVersionLocatorParam `json:"pronunciation_dictionary_locators,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request.
	VoiceSettings VoiceSettingsParam `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechConvertWithTimestampsParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechConvertWithTimestampsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechConvertWithTimestampsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToSpeechConvertWithTimestampsParams]'s query parameters
// as `url.Values`.
func (r TextToSpeechConvertWithTimestampsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro
// tier or above. Note that the μ-law format (sometimes written mu-law, often
// approximated as u-law) is commonly used for Twilio audio inputs.
type TextToSpeechConvertWithTimestampsParamsOutputFormat string

const (
	TextToSpeechConvertWithTimestampsParamsOutputFormatAlaw8000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "alaw_8000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_22050_32  TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_22050_32"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_24000_48  TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_24000_48"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_44100_128 TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_44100_128"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_44100_192 TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_44100_192"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_44100_32  TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_44100_32"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_44100_64  TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_44100_64"
	TextToSpeechConvertWithTimestampsParamsOutputFormatMP3_44100_96  TextToSpeechConvertWithTimestampsParamsOutputFormat = "mp3_44100_96"
	TextToSpeechConvertWithTimestampsParamsOutputFormatOpus48000_128 TextToSpeechConvertWithTimestampsParamsOutputFormat = "opus_48000_128"
	TextToSpeechConvertWithTimestampsParamsOutputFormatOpus48000_192 TextToSpeechConvertWithTimestampsParamsOutputFormat = "opus_48000_192"
	TextToSpeechConvertWithTimestampsParamsOutputFormatOpus48000_32  TextToSpeechConvertWithTimestampsParamsOutputFormat = "opus_48000_32"
	TextToSpeechConvertWithTimestampsParamsOutputFormatOpus48000_64  TextToSpeechConvertWithTimestampsParamsOutputFormat = "opus_48000_64"
	TextToSpeechConvertWithTimestampsParamsOutputFormatOpus48000_96  TextToSpeechConvertWithTimestampsParamsOutputFormat = "opus_48000_96"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm16000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_16000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm22050      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_22050"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm24000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_24000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm32000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_32000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm44100      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_44100"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm48000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_48000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatPcm8000       TextToSpeechConvertWithTimestampsParamsOutputFormat = "pcm_8000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatUlaw8000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "ulaw_8000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav16000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_16000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav22050      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_22050"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav24000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_24000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav32000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_32000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav44100      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_44100"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav48000      TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_48000"
	TextToSpeechConvertWithTimestampsParamsOutputFormatWav8000       TextToSpeechConvertWithTimestampsParamsOutputFormat = "wav_8000"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToSpeechConvertWithTimestampsParamsApplyTextNormalization string

const (
	TextToSpeechConvertWithTimestampsParamsApplyTextNormalizationAuto TextToSpeechConvertWithTimestampsParamsApplyTextNormalization = "auto"
	TextToSpeechConvertWithTimestampsParamsApplyTextNormalizationOn   TextToSpeechConvertWithTimestampsParamsApplyTextNormalization = "on"
	TextToSpeechConvertWithTimestampsParamsApplyTextNormalizationOff  TextToSpeechConvertWithTimestampsParamsApplyTextNormalization = "off"
)

type TextToSpeechStreamConvertParams struct {
	// The text that will get converted into speech.
	Text string `json:"text,required"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// The text that comes after the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	NextText param.Opt[string] `json:"next_text,omitzero"`
	// The text that came before the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	PreviousText param.Opt[string] `json:"previous_text,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// This parameter controls language text normalization. This helps with proper
	// pronunciation of text in some supported languages. WARNING: This parameter can
	// heavily increase the latency of the request. Currently only supported for
	// Japanese.
	ApplyLanguageTextNormalization param.Opt[bool] `json:"apply_language_text_normalization,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If true, we won't use PVC version of the voice for the generation but the IVC
	// version. This is a temporary workaround for higher latency in PVC versions.
	UsePvcAsIvc param.Opt[bool] `json:"use_pvc_as_ivc,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of request_id of the samples that come after this generation.
	// next_request_ids is especially useful for maintaining the speech's continuity
	// when regenerating a sample that has had some audio quality issues. For example,
	// if you have generated 3 speech clips, and you want to improve clip 2, passing
	// the request id of clip 3 as a next_request_id (and that of clip 1 as a
	// previous_request_id) will help maintain natural flow in the combined speech. The
	// results will be best when the same model is used across the generations. In case
	// both next_text and next_request_ids is send, next_text will be ignored. A
	// maximum of 3 request_ids can be send.
	NextRequestIDs []string `json:"next_request_ids,omitzero"`
	// A list of request_id of the samples that were generated before this generation.
	// Can be used to improve the speech's continuity when splitting up a large task
	// into multiple requests. The results will be best when the same model is used
	// across the generations. In case both previous_text and previous_request_ids is
	// send, previous_text will be ignored. A maximum of 3 request_ids can be send.
	PreviousRequestIDs []string `json:"previous_request_ids,omitzero"`
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
	OutputFormat TextToSpeechStreamConvertParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToSpeechStreamConvertParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request.
	VoiceSettings VoiceSettingsParam `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechStreamConvertParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechStreamConvertParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechStreamConvertParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToSpeechStreamConvertParams]'s query parameters as
// `url.Values`.
func (r TextToSpeechStreamConvertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
// Note that the μ-law format (sometimes written mu-law, often approximated as
// u-law) is commonly used for Twilio audio inputs.
type TextToSpeechStreamConvertParamsOutputFormat string

const (
	TextToSpeechStreamConvertParamsOutputFormatMP3_22050_32  TextToSpeechStreamConvertParamsOutputFormat = "mp3_22050_32"
	TextToSpeechStreamConvertParamsOutputFormatMP3_24000_48  TextToSpeechStreamConvertParamsOutputFormat = "mp3_24000_48"
	TextToSpeechStreamConvertParamsOutputFormatMP3_44100_32  TextToSpeechStreamConvertParamsOutputFormat = "mp3_44100_32"
	TextToSpeechStreamConvertParamsOutputFormatMP3_44100_64  TextToSpeechStreamConvertParamsOutputFormat = "mp3_44100_64"
	TextToSpeechStreamConvertParamsOutputFormatMP3_44100_96  TextToSpeechStreamConvertParamsOutputFormat = "mp3_44100_96"
	TextToSpeechStreamConvertParamsOutputFormatMP3_44100_128 TextToSpeechStreamConvertParamsOutputFormat = "mp3_44100_128"
	TextToSpeechStreamConvertParamsOutputFormatMP3_44100_192 TextToSpeechStreamConvertParamsOutputFormat = "mp3_44100_192"
	TextToSpeechStreamConvertParamsOutputFormatPcm8000       TextToSpeechStreamConvertParamsOutputFormat = "pcm_8000"
	TextToSpeechStreamConvertParamsOutputFormatPcm16000      TextToSpeechStreamConvertParamsOutputFormat = "pcm_16000"
	TextToSpeechStreamConvertParamsOutputFormatPcm22050      TextToSpeechStreamConvertParamsOutputFormat = "pcm_22050"
	TextToSpeechStreamConvertParamsOutputFormatPcm24000      TextToSpeechStreamConvertParamsOutputFormat = "pcm_24000"
	TextToSpeechStreamConvertParamsOutputFormatPcm32000      TextToSpeechStreamConvertParamsOutputFormat = "pcm_32000"
	TextToSpeechStreamConvertParamsOutputFormatPcm44100      TextToSpeechStreamConvertParamsOutputFormat = "pcm_44100"
	TextToSpeechStreamConvertParamsOutputFormatPcm48000      TextToSpeechStreamConvertParamsOutputFormat = "pcm_48000"
	TextToSpeechStreamConvertParamsOutputFormatUlaw8000      TextToSpeechStreamConvertParamsOutputFormat = "ulaw_8000"
	TextToSpeechStreamConvertParamsOutputFormatAlaw8000      TextToSpeechStreamConvertParamsOutputFormat = "alaw_8000"
	TextToSpeechStreamConvertParamsOutputFormatOpus48000_32  TextToSpeechStreamConvertParamsOutputFormat = "opus_48000_32"
	TextToSpeechStreamConvertParamsOutputFormatOpus48000_64  TextToSpeechStreamConvertParamsOutputFormat = "opus_48000_64"
	TextToSpeechStreamConvertParamsOutputFormatOpus48000_96  TextToSpeechStreamConvertParamsOutputFormat = "opus_48000_96"
	TextToSpeechStreamConvertParamsOutputFormatOpus48000_128 TextToSpeechStreamConvertParamsOutputFormat = "opus_48000_128"
	TextToSpeechStreamConvertParamsOutputFormatOpus48000_192 TextToSpeechStreamConvertParamsOutputFormat = "opus_48000_192"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToSpeechStreamConvertParamsApplyTextNormalization string

const (
	TextToSpeechStreamConvertParamsApplyTextNormalizationAuto TextToSpeechStreamConvertParamsApplyTextNormalization = "auto"
	TextToSpeechStreamConvertParamsApplyTextNormalizationOn   TextToSpeechStreamConvertParamsApplyTextNormalization = "on"
	TextToSpeechStreamConvertParamsApplyTextNormalizationOff  TextToSpeechStreamConvertParamsApplyTextNormalization = "off"
)

type TextToSpeechStreamConvertWithTimestampsParams struct {
	// The text that will get converted into speech.
	Text string `json:"text,required"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// Language code (ISO 639-1) used to enforce a language for the model and text
	// normalization. If the model does not support provided language code, an error
	// will be returned.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// The text that comes after the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	NextText param.Opt[string] `json:"next_text,omitzero"`
	// The text that came before the text of the current request. Can be used to
	// improve the speech's continuity when concatenating together multiple generations
	// or to influence the speech's continuity in the current generation.
	PreviousText param.Opt[string] `json:"previous_text,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// This parameter controls language text normalization. This helps with proper
	// pronunciation of text in some supported languages. WARNING: This parameter can
	// heavily increase the latency of the request. Currently only supported for
	// Japanese.
	ApplyLanguageTextNormalization param.Opt[bool] `json:"apply_language_text_normalization,omitzero"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for text to speech, you can check
	// this using the can_do_text_to_speech property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If true, we won't use PVC version of the voice for the generation but the IVC
	// version. This is a temporary workaround for higher latency in PVC versions.
	UsePvcAsIvc param.Opt[bool] `json:"use_pvc_as_ivc,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A list of request_id of the samples that come after this generation.
	// next_request_ids is especially useful for maintaining the speech's continuity
	// when regenerating a sample that has had some audio quality issues. For example,
	// if you have generated 3 speech clips, and you want to improve clip 2, passing
	// the request id of clip 3 as a next_request_id (and that of clip 1 as a
	// previous_request_id) will help maintain natural flow in the combined speech. The
	// results will be best when the same model is used across the generations. In case
	// both next_text and next_request_ids is send, next_text will be ignored. A
	// maximum of 3 request_ids can be send.
	NextRequestIDs []string `json:"next_request_ids,omitzero"`
	// A list of request_id of the samples that were generated before this generation.
	// Can be used to improve the speech's continuity when splitting up a large task
	// into multiple requests. The results will be best when the same model is used
	// across the generations. In case both previous_text and previous_request_ids is
	// send, previous_text will be ignored. A maximum of 3 request_ids can be send.
	PreviousRequestIDs []string `json:"previous_request_ids,omitzero"`
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
	OutputFormat TextToSpeechStreamConvertWithTimestampsParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and
	// 'off'. When set to 'auto', the system will automatically decide whether to apply
	// text normalization (e.g., spelling out numbers). With 'on', text normalization
	// will always be applied, while with 'off', it will be skipped.
	//
	// Any of "auto", "on", "off".
	ApplyTextNormalization TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request.
	VoiceSettings VoiceSettingsParam `json:"voice_settings,omitzero"`
	paramObj
}

func (r TextToSpeechStreamConvertWithTimestampsParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToSpeechStreamConvertWithTimestampsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToSpeechStreamConvertWithTimestampsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToSpeechStreamConvertWithTimestampsParams]'s query
// parameters as `url.Values`.
func (r TextToSpeechStreamConvertWithTimestampsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
// Note that the μ-law format (sometimes written mu-law, often approximated as
// u-law) is commonly used for Twilio audio inputs.
type TextToSpeechStreamConvertWithTimestampsParamsOutputFormat string

const (
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_22050_32  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_22050_32"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_24000_48  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_24000_48"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_44100_32  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_44100_32"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_44100_64  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_44100_64"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_44100_96  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_44100_96"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_44100_128 TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_44100_128"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_44100_192 TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "mp3_44100_192"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm8000       TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_8000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm16000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_16000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm22050      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_22050"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm24000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_24000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm32000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_32000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm44100      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_44100"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatPcm48000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "pcm_48000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatUlaw8000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "ulaw_8000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatAlaw8000      TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "alaw_8000"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatOpus48000_32  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "opus_48000_32"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatOpus48000_64  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "opus_48000_64"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatOpus48000_96  TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "opus_48000_96"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatOpus48000_128 TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "opus_48000_128"
	TextToSpeechStreamConvertWithTimestampsParamsOutputFormatOpus48000_192 TextToSpeechStreamConvertWithTimestampsParamsOutputFormat = "opus_48000_192"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and
// 'off'. When set to 'auto', the system will automatically decide whether to apply
// text normalization (e.g., spelling out numbers). With 'on', text normalization
// will always be applied, while with 'off', it will be skipped.
type TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalization string

const (
	TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalizationAuto TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalization = "auto"
	TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalizationOn   TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalization = "on"
	TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalizationOff  TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalization = "off"
)
