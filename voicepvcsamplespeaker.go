// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// VoicePvcSampleSpeakerService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoicePvcSampleSpeakerService] method instead.
type VoicePvcSampleSpeakerService struct {
	Options []option.RequestOption
}

// NewVoicePvcSampleSpeakerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVoicePvcSampleSpeakerService(opts ...option.RequestOption) (r VoicePvcSampleSpeakerService) {
	r = VoicePvcSampleSpeakerService{}
	r.Options = opts
	return
}

// Retrieve the separated audio for a specific speaker.
func (r *VoicePvcSampleSpeakerService) GetAudio(ctx context.Context, speakerID string, params VoicePvcSampleSpeakerGetAudioParams, opts ...option.RequestOption) (res *VoicePvcSampleSpeakerGetAudioResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if params.SampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	if speakerID == "" {
		err = errors.New("missing required speaker_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s/speakers/%s/audio", params.VoiceID, params.SampleID, speakerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the status of the speaker separation process and the list of detected
// speakers if complete.
func (r *VoicePvcSampleSpeakerService) GetStatus(ctx context.Context, sampleID string, params VoicePvcSampleSpeakerGetStatusParams, opts ...option.RequestOption) (res *SpeakerSeparationResponse, err error) {
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
	path := fmt.Sprintf("v1/voices/pvc/%s/samples/%s/speakers", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type SpeakerSeparationResponse struct {
	// The ID of the sample.
	SampleID string `json:"sample_id,required"`
	// The status of the speaker separation.
	//
	// Any of "not_started", "pending", "completed", "failed".
	Status SpeakerSeparationResponseStatus `json:"status,required"`
	// The ID of the voice.
	VoiceID string `json:"voice_id,required"`
	// The IDs of the selected speakers.
	SelectedSpeakerIDs []string `json:"selected_speaker_ids,nullable"`
	// The speakers of the sample.
	Speakers map[string]SpeakerSeparationResponseSpeaker `json:"speakers,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SampleID           respjson.Field
		Status             respjson.Field
		VoiceID            respjson.Field
		SelectedSpeakerIDs respjson.Field
		Speakers           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeakerSeparationResponse) RawJSON() string { return r.JSON.raw }
func (r *SpeakerSeparationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the speaker separation.
type SpeakerSeparationResponseStatus string

const (
	SpeakerSeparationResponseStatusNotStarted SpeakerSeparationResponseStatus = "not_started"
	SpeakerSeparationResponseStatusPending    SpeakerSeparationResponseStatus = "pending"
	SpeakerSeparationResponseStatusCompleted  SpeakerSeparationResponseStatus = "completed"
	SpeakerSeparationResponseStatusFailed     SpeakerSeparationResponseStatus = "failed"
)

type SpeakerSeparationResponseSpeaker struct {
	// The duration of the speaker segment in seconds.
	DurationSecs float64 `json:"duration_secs,required"`
	// The ID of the speaker.
	SpeakerID string `json:"speaker_id,required"`
	// The utterances of the speaker.
	Utterances []SpeakerSeparationResponseSpeakerUtterance `json:"utterances,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DurationSecs respjson.Field
		SpeakerID    respjson.Field
		Utterances   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeakerSeparationResponseSpeaker) RawJSON() string { return r.JSON.raw }
func (r *SpeakerSeparationResponseSpeaker) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeakerSeparationResponseSpeakerUtterance struct {
	// The end time of the utterance in seconds.
	End float64 `json:"end,required"`
	// The start time of the utterance in seconds.
	Start float64 `json:"start,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeakerSeparationResponseSpeakerUtterance) RawJSON() string { return r.JSON.raw }
func (r *SpeakerSeparationResponseSpeakerUtterance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleSpeakerGetAudioResponse struct {
	// The base64 encoded audio.
	AudioBase64 string `json:"audio_base_64,required"`
	// The duration of the audio in seconds.
	DurationSecs float64 `json:"duration_secs,required"`
	// The media type of the audio.
	MediaType string `json:"media_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioBase64  respjson.Field
		DurationSecs respjson.Field
		MediaType    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoicePvcSampleSpeakerGetAudioResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcSampleSpeakerGetAudioResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcSampleSpeakerGetAudioParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Sample ID to be used
	SampleID string `path:"sample_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type VoicePvcSampleSpeakerGetStatusParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
