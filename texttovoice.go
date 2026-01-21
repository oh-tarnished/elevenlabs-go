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

// TextToVoiceService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTextToVoiceService] method instead.
type TextToVoiceService struct {
	Options []option.RequestOption
}

// NewTextToVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTextToVoiceService(opts ...option.RequestOption) (r TextToVoiceService) {
	r = TextToVoiceService{}
	r.Options = opts
	return
}

// Generate a custom voice based on voice description. This method returns a list
// of voice previews. Each preview has a generated_voice_id and a sample of the
// voice as base64 encoded mp3 audio. If you like the a voice previewand want to
// create the voice call /v1/text-to-voice/create-voice-from-preview with the
// generated_voice_id to create the voice.
func (r *TextToVoiceService) NewPreviews(ctx context.Context, params TextToVoiceNewPreviewsParams, opts ...option.RequestOption) (res *VoicePreviews, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/text-to-voice/create-previews"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a voice from previously generated voice preview. This endpoint should be
// called after you fetched a generated_voice_id using POST
// /v1/text-to-voice/design or POST /v1/text-to-voice/:voice_id/remix.
func (r *TextToVoiceService) NewVoice(ctx context.Context, params TextToVoiceNewVoiceParams, opts ...option.RequestOption) (res *Voice, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/text-to-voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Design a voice via a prompt. This method returns a list of voice previews. Each
// preview has a generated_voice_id and a sample of the voice as base64 encoded mp3
// audio. To create a voice use the generated_voice_id of the preferred preview
// with the /v1/text-to-voice endpoint.
func (r *TextToVoiceService) DesignVoice(ctx context.Context, params TextToVoiceDesignVoiceParams, opts ...option.RequestOption) (res *VoicePreviews, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/text-to-voice/design"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Remix an existing voice via a prompt. This method returns a list of voice
// previews. Each preview has a generated_voice_id and a sample of the voice as
// base64 encoded mp3 audio. To create a voice use the generated_voice_id of the
// preferred preview with the /v1/text-to-voice endpoint.
func (r *TextToVoiceService) RemixVoice(ctx context.Context, voiceID string, params TextToVoiceRemixVoiceParams, opts ...option.RequestOption) (res *VoicePreviews, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-voice/%s/remix", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Stream a voice preview that was created via the /v1/text-to-voice/design
// endpoint.
func (r *TextToVoiceService) StreamPreview(ctx context.Context, generatedVoiceID string, query TextToVoiceStreamPreviewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if generatedVoiceID == "" {
		err = errors.New("missing required generated_voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/text-to-voice/%s/stream", generatedVoiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VoicePreviews struct {
	// The previews of the generated voices.
	Previews []VoicePreviewsPreview `json:"previews,required"`
	// The text used to preview the voices.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Previews    respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePreviews) RawJSON() string { return r.JSON.raw }
func (r *VoicePreviews) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePreviewsPreview struct {
	// The base64 encoded audio of the preview.
	AudioBase64 string `json:"audio_base_64,required"`
	// The duration of the preview in seconds.
	DurationSecs float64 `json:"duration_secs,required"`
	// The ID of the generated voice. Use it to create a voice from the preview.
	GeneratedVoiceID string `json:"generated_voice_id,required"`
	// The language of the preview.
	Language string `json:"language,required"`
	// The media type of the preview.
	MediaType string `json:"media_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64      respjson.Field
		DurationSecs     respjson.Field
		GeneratedVoiceID respjson.Field
		Language         respjson.Field
		MediaType        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePreviewsPreview) RawJSON() string { return r.JSON.raw }
func (r *VoicePreviewsPreview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToVoiceNewPreviewsParams struct {
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description,required"`
	// Random number that controls the voice generation. Same seed with same inputs
	// produces same voice.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Text to generate, text length has to be between 100 and 1000.
	Text param.Opt[string] `json:"text,omitzero"`
	// Whether to automatically generate a text suitable for the voice description.
	AutoGenerateText param.Opt[bool] `json:"auto_generate_text,omitzero"`
	// Controls how closely the AI follows the prompt. Lower numbers give the AI more
	// freedom to be creative, while higher numbers force it to stick more to the
	// prompt. High numbers can cause voice to sound artificial or robotic. We
	// recommend to use longer, more detailed prompts at lower Guidance Scale.
	GuidanceScale param.Opt[float64] `json:"guidance_scale,omitzero"`
	// Controls the volume level of the generated voice. -1 is quietest, 1 is loudest,
	// 0 corresponds to roughly -24 LUFS.
	Loudness param.Opt[float64] `json:"loudness,omitzero"`
	// Higher quality results in better voice output but less variety.
	Quality param.Opt[float64] `json:"quality,omitzero"`
	// Whether to enhance the voice description using AI to add more detail and improve
	// voice generation quality. When enabled, the system will automatically expand
	// simple prompts into more detailed voice descriptions. Defaults to False
	ShouldEnhance param.Opt[bool] `json:"should_enhance,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
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
	paramObj
}

func (r TextToVoiceNewPreviewsParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToVoiceNewPreviewsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToVoiceNewPreviewsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToVoiceNewPreviewsParams]'s query parameters as
// `url.Values`.
func (r TextToVoiceNewPreviewsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TextToVoiceNewVoiceParams struct {
	// The generated_voice_id to create, call POST /v1/text-to-voice/create-previews
	// and fetch the generated_voice_id from the response header if don't have one yet.
	GeneratedVoiceID string `json:"generated_voice_id,required"`
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description,required"`
	// Name to use for the created voice.
	VoiceName string `json:"voice_name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Optional, metadata to add to the created voice. Defaults to None.
	Labels map[string]string `json:"labels,omitzero"`
	// List of voice ids that the user has played but not selected. Used for RLHF.
	PlayedNotSelectedVoiceIDs []string `json:"played_not_selected_voice_ids,omitzero"`
	paramObj
}

func (r TextToVoiceNewVoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToVoiceNewVoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToVoiceNewVoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TextToVoiceDesignVoiceParams struct {
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description,required"`
	// Controls the balance of prompt versus reference audio when generating voice
	// samples. 0 means almost no prompt influence, 1 means almost no reference audio
	// influence. Only supported when using the eleven_ttv_v3 model.
	PromptStrength param.Opt[float64] `json:"prompt_strength,omitzero"`
	// Higher quality results in better voice output but less variety.
	Quality param.Opt[float64] `json:"quality,omitzero"`
	// Reference audio to use for the voice generation. The audio should be base64
	// encoded. Only supported when using the eleven_ttv_v3 model.
	ReferenceAudioBase64 param.Opt[string] `json:"reference_audio_base64,omitzero"`
	// The remixing session id.
	RemixingSessionID param.Opt[string] `json:"remixing_session_id,omitzero"`
	// The id of the remixing session iteration where these generations should be
	// attached to. If not provided, a new iteration will be created.
	RemixingSessionIterationID param.Opt[string] `json:"remixing_session_iteration_id,omitzero"`
	// Random number that controls the voice generation. Same seed with same inputs
	// produces same voice.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Text to generate, text length has to be between 100 and 1000.
	Text param.Opt[string] `json:"text,omitzero"`
	// Whether to automatically generate a text suitable for the voice description.
	AutoGenerateText param.Opt[bool] `json:"auto_generate_text,omitzero"`
	// Controls how closely the AI follows the prompt. Lower numbers give the AI more
	// freedom to be creative, while higher numbers force it to stick more to the
	// prompt. High numbers can cause voice to sound artificial or robotic. We
	// recommend to use longer, more detailed prompts at lower Guidance Scale.
	GuidanceScale param.Opt[float64] `json:"guidance_scale,omitzero"`
	// Controls the volume level of the generated voice. -1 is quietest, 1 is loudest,
	// 0 corresponds to roughly -24 LUFS.
	Loudness param.Opt[float64] `json:"loudness,omitzero"`
	// Whether to enhance the voice description using AI to add more detail and improve
	// voice generation quality. When enabled, the system will automatically expand
	// simple prompts into more detailed voice descriptions. Defaults to False
	ShouldEnhance param.Opt[bool] `json:"should_enhance,omitzero"`
	// Determines whether the Text to Voice previews should be included in the
	// response. If true, only the generated IDs will be returned which can then be
	// streamed via the /v1/text-to-voice/:generated_voice_id/stream endpoint.
	StreamPreviews param.Opt[bool] `json:"stream_previews,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
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
	// Model to use for the voice generation. Possible values:
	// eleven_multilingual_ttv_v2, eleven_ttv_v3.
	//
	// Any of "eleven_multilingual_ttv_v2", "eleven_ttv_v3".
	ModelID TextToVoiceDesignVoiceParamsModelID `json:"model_id,omitzero"`
	paramObj
}

func (r TextToVoiceDesignVoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToVoiceDesignVoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToVoiceDesignVoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToVoiceDesignVoiceParams]'s query parameters as
// `url.Values`.
func (r TextToVoiceDesignVoiceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Model to use for the voice generation. Possible values:
// eleven_multilingual_ttv_v2, eleven_ttv_v3.
type TextToVoiceDesignVoiceParamsModelID string

const (
	TextToVoiceDesignVoiceParamsModelIDElevenMultilingualTtvV2 TextToVoiceDesignVoiceParamsModelID = "eleven_multilingual_ttv_v2"
	TextToVoiceDesignVoiceParamsModelIDElevenTtvV3             TextToVoiceDesignVoiceParamsModelID = "eleven_ttv_v3"
)

type TextToVoiceRemixVoiceParams struct {
	// Description of the changes to make to the voice.
	VoiceDescription string `json:"voice_description,required"`
	// Controls the balance of prompt versus reference audio when generating voice
	// samples. 0 means almost no prompt influence, 1 means almost no reference audio
	// influence. Only supported when using the eleven_ttv_v3 model.
	PromptStrength param.Opt[float64] `json:"prompt_strength,omitzero"`
	// The remixing session id.
	RemixingSessionID param.Opt[string] `json:"remixing_session_id,omitzero"`
	// The id of the remixing session iteration where these generations should be
	// attached to. If not provided, a new iteration will be created.
	RemixingSessionIterationID param.Opt[string] `json:"remixing_session_iteration_id,omitzero"`
	// Random number that controls the voice generation. Same seed with same inputs
	// produces same voice.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Text to generate, text length has to be between 100 and 1000.
	Text param.Opt[string] `json:"text,omitzero"`
	// Whether to automatically generate a text suitable for the voice description.
	AutoGenerateText param.Opt[bool] `json:"auto_generate_text,omitzero"`
	// Controls how closely the AI follows the prompt. Lower numbers give the AI more
	// freedom to be creative, while higher numbers force it to stick more to the
	// prompt. High numbers can cause voice to sound artificial or robotic. We
	// recommend to use longer, more detailed prompts at lower Guidance Scale.
	GuidanceScale param.Opt[float64] `json:"guidance_scale,omitzero"`
	// Controls the volume level of the generated voice. -1 is quietest, 1 is loudest,
	// 0 corresponds to roughly -24 LUFS.
	Loudness param.Opt[float64] `json:"loudness,omitzero"`
	// Determines whether the Text to Voice previews should be included in the
	// response. If true, only the generated IDs will be returned which can then be
	// streamed via the /v1/text-to-voice/:generated_voice_id/stream endpoint.
	StreamPreviews param.Opt[bool] `json:"stream_previews,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
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
	paramObj
}

func (r TextToVoiceRemixVoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow TextToVoiceRemixVoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TextToVoiceRemixVoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [TextToVoiceRemixVoiceParams]'s query parameters as
// `url.Values`.
func (r TextToVoiceRemixVoiceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TextToVoiceStreamPreviewParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
