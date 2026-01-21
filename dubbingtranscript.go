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

// DubbingTranscriptService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDubbingTranscriptService] method instead.
type DubbingTranscriptService struct {
	Options []option.RequestOption
}

// NewDubbingTranscriptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDubbingTranscriptService(opts ...option.RequestOption) (r DubbingTranscriptService) {
	r = DubbingTranscriptService{}
	r.Options = opts
	return
}

// Fetch the transcript for one of the languages in a dub.
func (r *DubbingTranscriptService) Get(ctx context.Context, formatType DubbingTranscriptGetParamsFormatType, params DubbingTranscriptGetParams, opts ...option.RequestOption) (res *DubbingTranscriptGetResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if params.LanguageCode == "" {
		err = errors.New("missing required language_code parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/%s/transcripts/%s/format/%v", params.DubbingID, params.LanguageCode, formatType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DubbingTranscriptGetResponse struct {
	// Any of "srt", "webvtt", "json".
	TranscriptFormat DubbingTranscriptGetResponseTranscriptFormat `json:"transcript_format,required"`
	Json             DubbingTranscriptGetResponseJson             `json:"json,nullable"`
	Srt              string                                       `json:"srt,nullable"`
	Webvtt           string                                       `json:"webvtt,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TranscriptFormat respjson.Field
		Json             respjson.Field
		Srt              respjson.Field
		Webvtt           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingTranscriptGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingTranscriptGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingTranscriptGetResponseTranscriptFormat string

const (
	DubbingTranscriptGetResponseTranscriptFormatSrt    DubbingTranscriptGetResponseTranscriptFormat = "srt"
	DubbingTranscriptGetResponseTranscriptFormatWebvtt DubbingTranscriptGetResponseTranscriptFormat = "webvtt"
	DubbingTranscriptGetResponseTranscriptFormatJson   DubbingTranscriptGetResponseTranscriptFormat = "json"
)

type DubbingTranscriptGetResponseJson struct {
	Language   string                       `json:"language,required"`
	Utterances []DubbingTranscriptUtterance `json:"utterances,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		Utterances  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingTranscriptGetResponseJson) RawJSON() string { return r.JSON.raw }
func (r *DubbingTranscriptGetResponseJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingTranscriptGetParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// ISO-693 language code to retrieve the transcript for. Use 'source' to fetch the
	// transcript of the original media.
	LanguageCode string `path:"language_code,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// Format to return transcript in. For subtitles use either 'srt' or 'webvtt', and
// for a full transcript use 'json'. The 'json' format is not yet supported for
// Dubbing Studio.
type DubbingTranscriptGetParamsFormatType string

const (
	DubbingTranscriptGetParamsFormatTypeSrt    DubbingTranscriptGetParamsFormatType = "srt"
	DubbingTranscriptGetParamsFormatTypeWebvtt DubbingTranscriptGetParamsFormatType = "webvtt"
	DubbingTranscriptGetParamsFormatTypeJson   DubbingTranscriptGetParamsFormatType = "json"
)
