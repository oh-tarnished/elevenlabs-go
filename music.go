// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"encoding/json"
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

// MusicService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMusicService] method instead.
type MusicService struct {
	Options []option.RequestOption
}

// NewMusicService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMusicService(opts ...option.RequestOption) (r MusicService) {
	r = MusicService{}
	r.Options = opts
	return
}

// Compose a song from a prompt or a composition plan.
func (r *MusicService) Compose(ctx context.Context, params MusicComposeParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/*")}, opts...)
	path := "v1/music"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Compose a song from a prompt or a composition plan.
func (r *MusicService) ComposeWithDetailedResponse(ctx context.Context, params MusicComposeWithDetailedResponseParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "multipart/mixed")}, opts...)
	path := "v1/music/detailed"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Generate a composition plan from a prompt.
func (r *MusicService) GenerateCompositionPlan(ctx context.Context, params MusicGenerateCompositionPlanParams, opts ...option.RequestOption) (res *MusicPrompt, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/music/plan"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Separate an audio file into individual stems. This endpoint might have high
// latency, depending on the length of the audio file.
func (r *MusicService) SeparateStems(ctx context.Context, params MusicSeparateStemsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	path := "v1/music/stem-separation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Stream a composed song from a prompt or a composition plan.
func (r *MusicService) StreamComposedMusic(ctx context.Context, params MusicStreamComposedMusicParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/*")}, opts...)
	path := "v1/music/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type MusicPrompt struct {
	// The styles and musical directions that should not be present in the entire song.
	// Use English language for best result.
	NegativeGlobalStyles []string `json:"negative_global_styles,required"`
	// The styles and musical directions that should be present in the entire song. Use
	// English language for best result.
	PositiveGlobalStyles []string `json:"positive_global_styles,required"`
	// The sections of the song.
	Sections []MusicPromptSection `json:"sections,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NegativeGlobalStyles respjson.Field
		PositiveGlobalStyles respjson.Field
		Sections             respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MusicPrompt) RawJSON() string { return r.JSON.raw }
func (r *MusicPrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MusicPrompt to a MusicPromptParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MusicPromptParam.Overrides()
func (r MusicPrompt) ToParam() MusicPromptParam {
	return param.Override[MusicPromptParam](json.RawMessage(r.RawJSON()))
}

type MusicPromptSection struct {
	// The duration of the section in milliseconds. Must be between 3000ms and
	// 120000ms.
	DurationMs int64 `json:"duration_ms,required"`
	// The lyrics of the section. Max 200 characters per line.
	Lines []string `json:"lines,required"`
	// The styles and musical directions that should not be present in this section.
	// Use English language for best result.
	NegativeLocalStyles []string `json:"negative_local_styles,required"`
	// The styles and musical directions that should be present in this section. Use
	// English language for best result.
	PositiveLocalStyles []string `json:"positive_local_styles,required"`
	// The name of the section. Must be between 1 and 100 characters.
	SectionName string `json:"section_name,required"`
	// Optional source to extract the section from. Used for inpainting. Only available
	// to enterprise clients with access to the inpainting API.
	SourceFrom MusicPromptSectionSourceFrom `json:"source_from,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationMs          respjson.Field
		Lines               respjson.Field
		NegativeLocalStyles respjson.Field
		PositiveLocalStyles respjson.Field
		SectionName         respjson.Field
		SourceFrom          respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MusicPromptSection) RawJSON() string { return r.JSON.raw }
func (r *MusicPromptSection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional source to extract the section from. Used for inpainting. Only available
// to enterprise clients with access to the inpainting API.
type MusicPromptSectionSourceFrom struct {
	// The range to extract from the source song.
	Range TimeRange `json:"range,required"`
	// The ID of the song to source the section from. You can find the song ID in the
	// response headers when you generate a song.
	SongID string `json:"song_id,required"`
	// The ranges to exclude from the 'range'.
	NegativeRanges []TimeRange `json:"negative_ranges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Range          respjson.Field
		SongID         respjson.Field
		NegativeRanges respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MusicPromptSectionSourceFrom) RawJSON() string { return r.JSON.raw }
func (r *MusicPromptSectionSourceFrom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties NegativeGlobalStyles, PositiveGlobalStyles, Sections are
// required.
type MusicPromptParam struct {
	// The styles and musical directions that should not be present in the entire song.
	// Use English language for best result.
	NegativeGlobalStyles []string `json:"negative_global_styles,omitzero,required"`
	// The styles and musical directions that should be present in the entire song. Use
	// English language for best result.
	PositiveGlobalStyles []string `json:"positive_global_styles,omitzero,required"`
	// The sections of the song.
	Sections []MusicPromptSectionParam `json:"sections,omitzero,required"`
	paramObj
}

func (r MusicPromptParam) MarshalJSON() (data []byte, err error) {
	type shadow MusicPromptParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicPromptParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DurationMs, Lines, NegativeLocalStyles, PositiveLocalStyles,
// SectionName are required.
type MusicPromptSectionParam struct {
	// The duration of the section in milliseconds. Must be between 3000ms and
	// 120000ms.
	DurationMs int64 `json:"duration_ms,required"`
	// The lyrics of the section. Max 200 characters per line.
	Lines []string `json:"lines,omitzero,required"`
	// The styles and musical directions that should not be present in this section.
	// Use English language for best result.
	NegativeLocalStyles []string `json:"negative_local_styles,omitzero,required"`
	// The styles and musical directions that should be present in this section. Use
	// English language for best result.
	PositiveLocalStyles []string `json:"positive_local_styles,omitzero,required"`
	// The name of the section. Must be between 1 and 100 characters.
	SectionName string `json:"section_name,required"`
	// Optional source to extract the section from. Used for inpainting. Only available
	// to enterprise clients with access to the inpainting API.
	SourceFrom MusicPromptSectionSourceFromParam `json:"source_from,omitzero"`
	paramObj
}

func (r MusicPromptSectionParam) MarshalJSON() (data []byte, err error) {
	type shadow MusicPromptSectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicPromptSectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional source to extract the section from. Used for inpainting. Only available
// to enterprise clients with access to the inpainting API.
//
// The properties Range, SongID are required.
type MusicPromptSectionSourceFromParam struct {
	// The range to extract from the source song.
	Range TimeRangeParam `json:"range,omitzero,required"`
	// The ID of the song to source the section from. You can find the song ID in the
	// response headers when you generate a song.
	SongID string `json:"song_id,required"`
	// The ranges to exclude from the 'range'.
	NegativeRanges []TimeRangeParam `json:"negative_ranges,omitzero"`
	paramObj
}

func (r MusicPromptSectionSourceFromParam) MarshalJSON() (data []byte, err error) {
	type shadow MusicPromptSectionSourceFromParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicPromptSectionSourceFromParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TimeRange struct {
	EndMs   int64 `json:"end_ms,required"`
	StartMs int64 `json:"start_ms,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndMs       respjson.Field
		StartMs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeRange) RawJSON() string { return r.JSON.raw }
func (r *TimeRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimeRange to a TimeRangeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimeRangeParam.Overrides()
func (r TimeRange) ToParam() TimeRangeParam {
	return param.Override[TimeRangeParam](json.RawMessage(r.RawJSON()))
}

// The properties EndMs, StartMs are required.
type TimeRangeParam struct {
	EndMs   int64 `json:"end_ms,required"`
	StartMs int64 `json:"start_ms,required"`
	paramObj
}

func (r TimeRangeParam) MarshalJSON() (data []byte, err error) {
	type shadow TimeRangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimeRangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MusicComposeParams struct {
	// The ID of the finetune to use for the generation
	FinetuneID param.Opt[string] `json:"finetune_id,omitzero"`
	// The length of the song to generate in milliseconds. Used only in conjunction
	// with `prompt`. Must be between 3000ms and 600000ms. Optional - if not provided,
	// the model will choose a length based on the prompt.
	MusicLengthMs param.Opt[int64] `json:"music_length_ms,omitzero"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with
	// `composition_plan`.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// The seed to use for the generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// If true, guarantees that the generated song will be instrumental. If false, the
	// song may or may not be instrumental depending on the `prompt`. Can only be used
	// with `prompt`.
	ForceInstrumental param.Opt[bool] `json:"force_instrumental,omitzero"`
	// Controls how strictly section durations in the `composition_plan` are enforced.
	// Only used with `composition_plan`. When set to true, the model will precisely
	// respect each section's `duration_ms` from the plan. When set to false, the model
	// may adjust individual section durations which will generally lead to better
	// generation quality and improved latency, while always preserving the total song
	// duration from the plan.
	RespectSectionsDurations param.Opt[bool] `json:"respect_sections_durations,omitzero"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2pa param.Opt[bool] `json:"sign_with_c2pa,omitzero"`
	// Whether to store the generated song for inpainting. Only available to enterprise
	// clients with access to the inpainting API.
	StoreForInpainting param.Opt[bool] `json:"store_for_inpainting,omitzero"`
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
	// A detailed composition plan to guide music generation. Cannot be used in
	// conjunction with `prompt`.
	CompositionPlan MusicPromptParam `json:"composition_plan,omitzero"`
	// The model to use for the generation.
	//
	// Any of "music_v1".
	ModelID MusicComposeParamsModelID `json:"model_id,omitzero"`
	// A music prompt. Deprecated. Use `composition_plan` instead.
	MusicPrompt MusicPromptParam `json:"music_prompt,omitzero"`
	paramObj
}

func (r MusicComposeParams) MarshalJSON() (data []byte, err error) {
	type shadow MusicComposeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicComposeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MusicComposeParams]'s query parameters as `url.Values`.
func (r MusicComposeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The model to use for the generation.
type MusicComposeParamsModelID string

const (
	MusicComposeParamsModelIDMusicV1 MusicComposeParamsModelID = "music_v1"
)

type MusicComposeWithDetailedResponseParams struct {
	// The ID of the finetune to use for the generation
	FinetuneID param.Opt[string] `json:"finetune_id,omitzero"`
	// The length of the song to generate in milliseconds. Used only in conjunction
	// with `prompt`. Must be between 3000ms and 600000ms. Optional - if not provided,
	// the model will choose a length based on the prompt.
	MusicLengthMs param.Opt[int64] `json:"music_length_ms,omitzero"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with
	// `composition_plan`.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// The seed to use for the generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// If true, guarantees that the generated song will be instrumental. If false, the
	// song may or may not be instrumental depending on the `prompt`. Can only be used
	// with `prompt`.
	ForceInstrumental param.Opt[bool] `json:"force_instrumental,omitzero"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2pa param.Opt[bool] `json:"sign_with_c2pa,omitzero"`
	// Whether to store the generated song for inpainting. Only available to enterprise
	// clients with access to the inpainting API.
	StoreForInpainting param.Opt[bool] `json:"store_for_inpainting,omitzero"`
	// Whether to return the timestamps of the words in the generated song.
	WithTimestamps param.Opt[bool] `json:"with_timestamps,omitzero"`
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
	// A detailed composition plan to guide music generation. Cannot be used in
	// conjunction with `prompt`.
	CompositionPlan MusicPromptParam `json:"composition_plan,omitzero"`
	// The model to use for the generation.
	//
	// Any of "music_v1".
	ModelID MusicComposeWithDetailedResponseParamsModelID `json:"model_id,omitzero"`
	// A music prompt. Deprecated. Use `composition_plan` instead.
	MusicPrompt MusicPromptParam `json:"music_prompt,omitzero"`
	paramObj
}

func (r MusicComposeWithDetailedResponseParams) MarshalJSON() (data []byte, err error) {
	type shadow MusicComposeWithDetailedResponseParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicComposeWithDetailedResponseParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MusicComposeWithDetailedResponseParams]'s query parameters
// as `url.Values`.
func (r MusicComposeWithDetailedResponseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The model to use for the generation.
type MusicComposeWithDetailedResponseParamsModelID string

const (
	MusicComposeWithDetailedResponseParamsModelIDMusicV1 MusicComposeWithDetailedResponseParamsModelID = "music_v1"
)

type MusicGenerateCompositionPlanParams struct {
	// A simple text prompt to compose a plan from.
	Prompt string `json:"prompt,required"`
	// The length of the composition plan to generate in milliseconds. Must be between
	// 3000ms and 600000ms. Optional - if not provided, the model will choose a length
	// based on the prompt.
	MusicLengthMs param.Opt[int64] `json:"music_length_ms,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The model to use for the generation.
	//
	// Any of "music_v1".
	ModelID MusicGenerateCompositionPlanParamsModelID `json:"model_id,omitzero"`
	// An optional composition plan to use as a source for the new composition plan.
	SourceCompositionPlan MusicPromptParam `json:"source_composition_plan,omitzero"`
	paramObj
}

func (r MusicGenerateCompositionPlanParams) MarshalJSON() (data []byte, err error) {
	type shadow MusicGenerateCompositionPlanParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicGenerateCompositionPlanParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The model to use for the generation.
type MusicGenerateCompositionPlanParamsModelID string

const (
	MusicGenerateCompositionPlanParamsModelIDMusicV1 MusicGenerateCompositionPlanParamsModelID = "music_v1"
)

type MusicSeparateStemsParams struct {
	// The audio file to separate into stems.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2pa param.Opt[bool] `json:"sign_with_c2pa,omitzero"`
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
	// The id of the stem variation to use.
	//
	// Any of "two_stems_v1", "six_stems_v1".
	StemVariationID MusicSeparateStemsParamsStemVariationID `json:"stem_variation_id,omitzero"`
	paramObj
}

func (r MusicSeparateStemsParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [MusicSeparateStemsParams]'s query parameters as
// `url.Values`.
func (r MusicSeparateStemsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The id of the stem variation to use.
type MusicSeparateStemsParamsStemVariationID string

const (
	MusicSeparateStemsParamsStemVariationIDTwoStemsV1 MusicSeparateStemsParamsStemVariationID = "two_stems_v1"
	MusicSeparateStemsParamsStemVariationIDSixStemsV1 MusicSeparateStemsParamsStemVariationID = "six_stems_v1"
)

type MusicStreamComposedMusicParams struct {
	// The ID of the finetune to use for the generation
	FinetuneID param.Opt[string] `json:"finetune_id,omitzero"`
	// The length of the song to generate in milliseconds. Used only in conjunction
	// with `prompt`. Must be between 3000ms and 600000ms. Optional - if not provided,
	// the model will choose a length based on the prompt.
	MusicLengthMs param.Opt[int64] `json:"music_length_ms,omitzero"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with
	// `composition_plan`.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// The seed to use for the generation.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// If true, guarantees that the generated song will be instrumental. If false, the
	// song may or may not be instrumental depending on the `prompt`. Can only be used
	// with `prompt`.
	ForceInstrumental param.Opt[bool] `json:"force_instrumental,omitzero"`
	// Whether to store the generated song for inpainting. Only available to enterprise
	// clients with access to the inpainting API.
	StoreForInpainting param.Opt[bool] `json:"store_for_inpainting,omitzero"`
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
	// A detailed composition plan to guide music generation. Cannot be used in
	// conjunction with `prompt`.
	CompositionPlan MusicPromptParam `json:"composition_plan,omitzero"`
	// The model to use for the generation.
	//
	// Any of "music_v1".
	ModelID MusicStreamComposedMusicParamsModelID `json:"model_id,omitzero"`
	// A music prompt. Deprecated. Use `composition_plan` instead.
	MusicPrompt MusicPromptParam `json:"music_prompt,omitzero"`
	paramObj
}

func (r MusicStreamComposedMusicParams) MarshalJSON() (data []byte, err error) {
	type shadow MusicStreamComposedMusicParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MusicStreamComposedMusicParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [MusicStreamComposedMusicParams]'s query parameters as
// `url.Values`.
func (r MusicStreamComposedMusicParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The model to use for the generation.
type MusicStreamComposedMusicParamsModelID string

const (
	MusicStreamComposedMusicParamsModelIDMusicV1 MusicStreamComposedMusicParamsModelID = "music_v1"
)
