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

// VoiceSampleService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceSampleService] method instead.
type VoiceSampleService struct {
	Options []option.RequestOption
}

// NewVoiceSampleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceSampleService(opts ...option.RequestOption) (r VoiceSampleService) {
	r = VoiceSampleService{}
	r.Options = opts
	return
}

// Removes a sample by its ID.
func (r *VoiceSampleService) Delete(ctx context.Context, sampleID string, params VoiceSampleDeleteParams, opts ...option.RequestOption) (res *VoiceSampleDeleteResponse, err error) {
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
	path := fmt.Sprintf("v1/voices/%s/samples/%s", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns the audio corresponding to a sample attached to a voice.
func (r *VoiceSampleService) GetAudio(ctx context.Context, sampleID string, params VoiceSampleGetAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/*")}, opts...)
	if params.VoiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	if sampleID == "" {
		err = errors.New("missing required sample_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/%s/samples/%s/audio", params.VoiceID, sampleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VoiceSampleDeleteResponse struct {
	// The status of the sample deletion request. If the request was successful, the
	// status will be 'ok'. Otherwise an error message with status 500 will be
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
func (r VoiceSampleDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceSampleDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSampleDeleteParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type VoiceSampleGetAudioParams struct {
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all
	// the available voices.
	VoiceID string `path:"voice_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
