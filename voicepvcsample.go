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

// VoicePvcSampleService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoicePvcSampleService] method instead.
type VoicePvcSampleService struct {
	Options  []option.RequestOption
	Speakers VoicePvcSampleSpeakerService
}

// NewVoicePvcSampleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoicePvcSampleService(opts ...option.RequestOption) (r VoicePvcSampleService) {
	r = VoicePvcSampleService{}
	r.Options = opts
	r.Speakers = NewVoicePvcSampleSpeakerService(opts...)
	return
}

// Update a PVC voice sample - apply noise removal, select speaker, change trim
// times or file name.
func (r *VoicePvcSampleService) Update(ctx context.Context, sampleID string, params VoicePvcSampleUpdateParams, opts ...option.RequestOption) (res *AddVoiceResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a sample from a PVC voice.
func (r *VoicePvcSampleService) Delete(ctx context.Context, sampleID string, params VoicePvcSampleDeleteParams, opts ...option.RequestOption) (res *VoicePvcSampleDeleteResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Add audio samples to a PVC voice
func (r *VoicePvcSampleService) Add(ctx context.Context, voiceID string, params VoicePvcSampleAddParams, opts ...option.RequestOption) (res *[]SampleResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve the first 30 seconds of voice sample audio with or without noise
// removal.
func (r *VoicePvcSampleService) GetAudio(ctx context.Context, sampleID string, params VoicePvcSampleGetAudioParams, opts ...option.RequestOption) (res *VoicePvcSampleGetAudioResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s/audio", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Retrieve the visual waveform of a voice sample.
func (r *VoicePvcSampleService) GetWaveform(ctx context.Context, sampleID string, params VoicePvcSampleGetWaveformParams, opts ...option.RequestOption) (res *VoicePvcSampleGetWaveformResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s/waveform", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Start speaker separation process for a sample
func (r *VoicePvcSampleService) SeparateSpeakers(ctx context.Context, sampleID string, params VoicePvcSampleSeparateSpeakersParams, opts ...option.RequestOption) (res *VoicePvcSampleSeparateSpeakersResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s/separate-speakers", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type SampleResponse struct {
	// The name of the sample file.
	FileName string `json:"file_name,required"`
	// The hash of the sample file.
	Hash string `json:"hash,required"`
	// The MIME type of the sample file.
	MimeType string `json:"mime_type,required"`
	// The ID of the sample.
	SampleID string `json:"sample_id,required"`
	// The size of the sample file in bytes.
	SizeBytes               int64                     `json:"size_bytes,required"`
	DurationSecs            float64                   `json:"duration_secs,nullable"`
	HasIsolatedAudio        bool                      `json:"has_isolated_audio,nullable"`
	HasIsolatedAudioPreview bool                      `json:"has_isolated_audio_preview,nullable"`
	RemoveBackgroundNoise   bool                      `json:"remove_background_noise,nullable"`
	SpeakerSeparation       SpeakerSeparationResponse `json:"speaker_separation,nullable"`
	TrimEnd                 int64                     `json:"trim_end,nullable"`
	TrimStart               int64                     `json:"trim_start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileName                respjson.Field
		Hash                    respjson.Field
		MimeType                respjson.Field
		SampleID                respjson.Field
		SizeBytes               respjson.Field
		DurationSecs            respjson.Field
		HasIsolatedAudio        respjson.Field
		HasIsolatedAudioPreview respjson.Field
		RemoveBackgroundNoise   respjson.Field
		SpeakerSeparation       respjson.Field
		TrimEnd                 respjson.Field
		TrimStart               respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SampleResponse) RawJSON() string { return r.JSON.raw }
func (r *SampleResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleDeleteResponse struct {
	// The status of the voice sample deletion request. If the request was successful,
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
func (r VoicePvcSampleDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcSampleDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleGetAudioResponse struct {
	// The base64 encoded audio.
	AudioBase64 string `json:"audio_base_64,required"`
	// The media type of the audio.
	MediaType string `json:"media_type,required"`
	// The ID of the sample.
	SampleID string `json:"sample_id,required"`
	// The ID of the voice.
	VoiceID string `json:"voice_id,required"`
	// The duration of the audio in seconds.
	DurationSecs float64 `json:"duration_secs,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64  respjson.Field
		MediaType    respjson.Field
		SampleID     respjson.Field
		VoiceID      respjson.Field
		DurationSecs respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePvcSampleGetAudioResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcSampleGetAudioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleGetWaveformResponse struct {
	// The ID of the sample.
	SampleID string `json:"sample_id,required"`
	// The visual waveform of the sample, represented as a list of floats.
	VisualWaveform []float64 `json:"visual_waveform,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SampleID       respjson.Field
		VisualWaveform respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePvcSampleGetWaveformResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcSampleGetWaveformResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleSeparateSpeakersResponse struct {
	// The status of the start speaker seperation request. If the request was
	// successful, the status will be 'ok'. Otherwise an error message with status 500
	// will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePvcSampleSeparateSpeakersResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcSampleSeparateSpeakersResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleUpdateParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// The name of the audio file to be used for PVC training.
	FileName param.Opt[string] `json:"file_name,omitzero"`
	// The end time of the audio to be used for PVC training. Time should be in
	// milliseconds
	TrimEndTime param.Opt[int64] `json:"trim_end_time,omitzero"`
	// The start time of the audio to be used for PVC training. Time should be in
	// milliseconds
	TrimStartTime param.Opt[int64] `json:"trim_start_time,omitzero"`
	// If set will remove background noise for voice samples using our audio isolation
	// model. If the samples do not include background noise, it can make the quality
	// worse.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Speaker IDs to be used for PVC training. Make sure you send all the speaker IDs
	// you want to use for PVC training in one request because the last request will
	// override the previous ones.
	SelectedSpeakerIDs []string `json:"selected_speaker_ids,omitzero"`
	paramObj
}

func (r VoicePvcSampleUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow VoicePvcSampleUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicePvcSampleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleDeleteParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type VoicePvcSampleAddParams struct {
	// Audio files used to create the voice.
	Files []io.Reader `json:"files,omitzero,required" format:"binary"`
	// If set will remove background noise for voice samples using our audio isolation
	// model. If the samples do not include background noise, it can make the quality
	// worse.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoicePvcSampleAddParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type VoicePvcSampleGetAudioParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// If set will remove background noise for voice samples using our audio isolation
	// model. If the samples do not include background noise, it can make the quality
	// worse.
	RemoveBackgroundNoise param.Opt[bool] `query:"remove_background_noise,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VoicePvcSampleGetAudioParams]'s query parameters as
// `url.Values`.
func (r VoicePvcSampleGetAudioParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VoicePvcSampleGetWaveformParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type VoicePvcSampleSeparateSpeakersParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
