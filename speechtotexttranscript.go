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

// SpeechToTextTranscriptService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeechToTextTranscriptService] method instead.
type SpeechToTextTranscriptService struct {
	Options []option.RequestOption
}

// NewSpeechToTextTranscriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSpeechToTextTranscriptService(opts ...option.RequestOption) (r SpeechToTextTranscriptService) {
	r = SpeechToTextTranscriptService{}
	r.Options = opts
	return
}

// Retrieve a previously generated transcript by its ID.
func (r *SpeechToTextTranscriptService) Get(ctx context.Context, transcriptionID string, query SpeechToTextTranscriptGetParams, opts ...option.RequestOption) (res *SpeechToTextTranscriptGetResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if transcriptionID == "" {
		err = errors.New("missing required transcription_id parameter")
		return
	}
	path := fmt.Sprintf("v1/speech-to-text/transcripts/%s", transcriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a previously generated transcript by its ID.
func (r *SpeechToTextTranscriptService) Delete(ctx context.Context, transcriptionID string, body SpeechToTextTranscriptDeleteParams, opts ...option.RequestOption) (res *SpeechToTextTranscriptDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if transcriptionID == "" {
		err = errors.New("missing required transcription_id parameter")
		return
	}
	path := fmt.Sprintf("v1/speech-to-text/transcripts/%s", transcriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Chunk-level detail of the transcription with timing information.
type SpeechToTextTranscriptGetResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextTranscriptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextTranscriptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextTranscriptDeleteResponse = any

type SpeechToTextTranscriptGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type SpeechToTextTranscriptDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
