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

// DubbingResourceSpeakerService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDubbingResourceSpeakerService] method instead.
type DubbingResourceSpeakerService struct {
	Options []option.RequestOption
}

// NewDubbingResourceSpeakerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDubbingResourceSpeakerService(opts ...option.RequestOption) (r DubbingResourceSpeakerService) {
	r = DubbingResourceSpeakerService{}
	r.Options = opts
	return
}

// Create A New Speaker
func (r *DubbingResourceSpeakerService) New(ctx context.Context, dubbingID string, params DubbingResourceSpeakerNewParams, opts ...option.RequestOption) (res *DubbingResourceSpeakerNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/speaker", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a new segment in dubbing resource with a start and end time for the
// speaker in every available language. Does not automatically generate
// transcripts/translations/audio.
func (r *DubbingResourceSpeakerService) NewSegment(ctx context.Context, speakerID string, params DubbingResourceSpeakerNewSegmentParams, opts ...option.RequestOption) (res *DubbingResourceSpeakerNewSegmentResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if speakerID == "" {
		err = errors.New("missing required speaker_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/speaker/%s/segment", params.DubbingID, speakerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Fetch the top 10 similar voices to a speaker, including the voice IDs, names,
// descriptions, and, where possible, a sample audio recording.
func (r *DubbingResourceSpeakerService) FindSimilarVoices(ctx context.Context, speakerID string, params DubbingResourceSpeakerFindSimilarVoicesParams, opts ...option.RequestOption) (res *DubbingResourceSpeakerFindSimilarVoicesResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if speakerID == "" {
		err = errors.New("missing required speaker_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/speaker/%s/similar-voices", params.DubbingID, speakerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Amend the metadata associated with a speaker, such as their voice. Both voice
// cloning and using voices from the ElevenLabs library are supported.
func (r *DubbingResourceSpeakerService) UpdateMetadata(ctx context.Context, speakerID string, params DubbingResourceSpeakerUpdateMetadataParams, opts ...option.RequestOption) (res *DubbingResourceSpeakerUpdateMetadataResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if speakerID == "" {
		err = errors.New("missing required speaker_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/speaker/%s", params.DubbingID, speakerID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type DubbingResourceSpeakerNewResponse struct {
	SpeakerID string `json:"speaker_id,required"`
	Version   int64  `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SpeakerID   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSpeakerNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSpeakerNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerNewSegmentResponse struct {
	NewSegment string `json:"new_segment,required"`
	Version    int64  `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NewSegment  respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSpeakerNewSegmentResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSpeakerNewSegmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerFindSimilarVoicesResponse struct {
	Voices []DubbingResourceSpeakerFindSimilarVoicesResponseVoice `json:"voices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSpeakerFindSimilarVoicesResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSpeakerFindSimilarVoicesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerFindSimilarVoicesResponseVoice struct {
	// Any of "premade", "cloned", "generated", "professional", "famous".
	Category    string `json:"category,required"`
	Name        string `json:"name,required"`
	VoiceID     string `json:"voice_id,required"`
	Description string `json:"description,nullable"`
	PreviewURL  string `json:"preview_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category    respjson.Field
		Name        respjson.Field
		VoiceID     respjson.Field
		Description respjson.Field
		PreviewURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSpeakerFindSimilarVoicesResponseVoice) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSpeakerFindSimilarVoicesResponseVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerUpdateMetadataResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSpeakerUpdateMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSpeakerUpdateMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerNewParams struct {
	// Name to attribute to this speaker.
	SpeakerName param.Opt[string] `json:"speaker_name,omitzero"`
	// Either the identifier of a voice from the ElevenLabs voice library, or one of
	// ['track-clone', 'clip-clone'].
	VoiceID param.Opt[string] `json:"voice_id,omitzero"`
	// For models that support it, the voice similarity value to use. This will default
	// to 1.0, with a valid range of [0.0, 1.0].
	VoiceSimilarity param.Opt[float64] `json:"voice_similarity,omitzero"`
	// For models that support it, the voice similarity value to use. This will default
	// to 0.65, with a valid range of [0.0, 1.0].
	VoiceStability param.Opt[float64] `json:"voice_stability,omitzero"`
	// For models that support it, the voice style value to use. This will default to
	// 1.0, with a valid range of [0.0, 1.0].
	VoiceStyle param.Opt[float64] `json:"voice_style,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceSpeakerNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceSpeakerNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceSpeakerNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerNewSegmentParams struct {
	// ID of the dubbing project.
	DubbingID string            `path:"dubbing_id,required" json:"-"`
	EndTime   float64           `json:"end_time,required"`
	StartTime float64           `json:"start_time,required"`
	Text      param.Opt[string] `json:"text,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey     param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	Translations map[string]string `json:"translations,omitzero"`
	paramObj
}

func (r DubbingResourceSpeakerNewSegmentParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceSpeakerNewSegmentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceSpeakerNewSegmentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSpeakerFindSimilarVoicesParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type DubbingResourceSpeakerUpdateMetadataParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// Name to attribute to this speaker.
	SpeakerName param.Opt[string] `json:"speaker_name,omitzero"`
	// Either the identifier of a voice from the ElevenLabs voice library, or one of
	// ['track-clone', 'clip-clone'].
	VoiceID param.Opt[string] `json:"voice_id,omitzero"`
	// For models that support it, the voice similarity value to use. This will default
	// to 1.0, with a valid range of [0.0, 1.0].
	VoiceSimilarity param.Opt[float64] `json:"voice_similarity,omitzero"`
	// For models that support it, the voice similarity value to use. This will default
	// to 0.65, with a valid range of [0.0, 1.0].
	VoiceStability param.Opt[float64] `json:"voice_stability,omitzero"`
	// For models that support it, the voice style value to use. This will default to
	// 1.0, with a valid range of [0.0, 1.0].
	VoiceStyle param.Opt[float64] `json:"voice_style,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Languages to apply these changes to. If empty, will apply to all languages.
	Languages []string `json:"languages,omitzero"`
	paramObj
}

func (r DubbingResourceSpeakerUpdateMetadataParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceSpeakerUpdateMetadataParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceSpeakerUpdateMetadataParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
