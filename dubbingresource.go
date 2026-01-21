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

// DubbingResourceService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDubbingResourceService] method instead.
type DubbingResourceService struct {
	Options []option.RequestOption
	Speaker DubbingResourceSpeakerService
	Segment DubbingResourceSegmentService
}

// NewDubbingResourceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDubbingResourceService(opts ...option.RequestOption) (r DubbingResourceService) {
	r = DubbingResourceService{}
	r.Options = opts
	r.Speaker = NewDubbingResourceSpeakerService(opts...)
	r.Segment = NewDubbingResourceSegmentService(opts...)
	return
}

// Given a dubbing ID generated from the '/v1/dubbing' endpoint with studio
// enabled, returns the dubbing resource.
func (r *DubbingResourceService) Get(ctx context.Context, dubbingID string, query DubbingResourceGetParams, opts ...option.RequestOption) (res *DubbingResourceGetResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Adds the given ElevenLab Turbo V2/V2.5 language code to the resource. Does not
// automatically generate transcripts/translations/audio.
func (r *DubbingResourceService) AddLanguage(ctx context.Context, dubbingID string, params DubbingResourceAddLanguageParams, opts ...option.RequestOption) (res *DubbingResourceAddLanguageResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/language", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Regenerate the dubs for either the entire resource or the specified
// segments/languages. Will automatically transcribe and translate any missing
// transcriptions and translations.
func (r *DubbingResourceService) Dub(ctx context.Context, dubbingID string, params DubbingResourceDubParams, opts ...option.RequestOption) (res *DubbingResourceDubResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/dub", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Change the attribution of one or more segments to a different speaker.
func (r *DubbingResourceService) MigrateSegments(ctx context.Context, dubbingID string, params DubbingResourceMigrateSegmentsParams, opts ...option.RequestOption) (res *DubbingResourceMigrateSegmentsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/migrate-segments", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Regenerate the output media for a language using the latest Studio state. Please
// ensure all segments have been dubbed before rendering, otherwise they will be
// omitted. Renders are generated asynchronously, and to check the status of all
// renders please use the 'Get Dubbing Resource' endpoint.
func (r *DubbingResourceService) Render(ctx context.Context, language string, params DubbingResourceRenderParams, opts ...option.RequestOption) (res *DubbingResourceRenderResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/render/%s", params.DubbingID, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Regenerate the transcriptions for the specified segments. Does not automatically
// regenerate translations or dubs.
func (r *DubbingResourceService) Transcribe(ctx context.Context, dubbingID string, params DubbingResourceTranscribeParams, opts ...option.RequestOption) (res *DubbingResourceTranscribeResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/transcribe", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Regenerate the translations for either the entire resource or the specified
// segments/languages. Will automatically transcribe missing transcriptions. Will
// not automatically regenerate the dubs.
func (r *DubbingResourceService) Translate(ctx context.Context, dubbingID string, params DubbingResourceTranslateParams, opts ...option.RequestOption) (res *DubbingResourceTranslateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/translate", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type DubbingMediaReference struct {
	BucketName     string  `json:"bucket_name,required"`
	ContentType    string  `json:"content_type,required"`
	DurationSecs   float64 `json:"duration_secs,required"`
	IsAudio        bool    `json:"is_audio,required"`
	RandomPathSlug string  `json:"random_path_slug,required"`
	Src            string  `json:"src,required"`
	URL            string  `json:"url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BucketName     respjson.Field
		ContentType    respjson.Field
		DurationSecs   respjson.Field
		IsAudio        respjson.Field
		RandomPathSlug respjson.Field
		Src            respjson.Field
		URL            respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingMediaReference) RawJSON() string { return r.JSON.raw }
func (r *DubbingMediaReference) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponse struct {
	ID              string                                              `json:"id,required"`
	Background      DubbingMediaReference                               `json:"background,required"`
	Foreground      DubbingMediaReference                               `json:"foreground,required"`
	Input           DubbingMediaReference                               `json:"input,required"`
	Renders         map[string]DubbingResourceGetResponseRender         `json:"renders,required"`
	SourceLanguage  string                                              `json:"source_language,required"`
	SpeakerSegments map[string]DubbingResourceGetResponseSpeakerSegment `json:"speaker_segments,required"`
	SpeakerTracks   map[string]DubbingResourceGetResponseSpeakerTrack   `json:"speaker_tracks,required"`
	TargetLanguages []string                                            `json:"target_languages,required"`
	Version         int64                                               `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Background      respjson.Field
		Foreground      respjson.Field
		Input           respjson.Field
		Renders         respjson.Field
		SourceLanguage  respjson.Field
		SpeakerSegments respjson.Field
		SpeakerTracks   respjson.Field
		TargetLanguages respjson.Field
		Version         respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseRender struct {
	ID       string                `json:"id,required"`
	Language string                `json:"language,required"`
	MediaRef DubbingMediaReference `json:"media_ref,required"`
	// Any of "complete", "processing", "failed".
	Status string `json:"status,required"`
	// Any of "mp4", "aac", "mp3", "wav", "aaf", "tracks_zip", "clips_zip".
	Type    string `json:"type,required"`
	Version int64  `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Language    respjson.Field
		MediaRef    respjson.Field
		Status      respjson.Field
		Type        respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseRender) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseRender) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseSpeakerSegment struct {
	ID        string                                                 `json:"id,required"`
	Dubs      map[string]DubbingResourceGetResponseSpeakerSegmentDub `json:"dubs,required"`
	EndTime   float64                                                `json:"end_time,required"`
	StartTime float64                                                `json:"start_time,required"`
	Subtitles []DubbingResourceGetResponseSpeakerSegmentSubtitle     `json:"subtitles,required"`
	Text      string                                                 `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Dubs        respjson.Field
		EndTime     respjson.Field
		StartTime   respjson.Field
		Subtitles   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseSpeakerSegment) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseSpeakerSegment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseSpeakerSegmentDub struct {
	AudioStale bool                                                  `json:"audio_stale,required"`
	EndTime    float64                                               `json:"end_time,required"`
	MediaRef   DubbingMediaReference                                 `json:"media_ref,required"`
	StartTime  float64                                               `json:"start_time,required"`
	Subtitles  []DubbingResourceGetResponseSpeakerSegmentDubSubtitle `json:"subtitles,required"`
	Text       string                                                `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioStale  respjson.Field
		EndTime     respjson.Field
		MediaRef    respjson.Field
		StartTime   respjson.Field
		Subtitles   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseSpeakerSegmentDub) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseSpeakerSegmentDub) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseSpeakerSegmentDubSubtitle struct {
	EndTime   float64  `json:"end_time,required"`
	Lines     []string `json:"lines,required"`
	StartTime float64  `json:"start_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		Lines       respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseSpeakerSegmentDubSubtitle) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseSpeakerSegmentDubSubtitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseSpeakerSegmentSubtitle struct {
	EndTime   float64  `json:"end_time,required"`
	Lines     []string `json:"lines,required"`
	StartTime float64  `json:"start_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		Lines       respjson.Field
		StartTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseSpeakerSegmentSubtitle) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseSpeakerSegmentSubtitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetResponseSpeakerTrack struct {
	ID          string                `json:"id,required"`
	MediaRef    DubbingMediaReference `json:"media_ref,required"`
	Segments    []string              `json:"segments,required"`
	SpeakerName string                `json:"speaker_name,required"`
	Voices      map[string]string     `json:"voices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		MediaRef    respjson.Field
		Segments    respjson.Field
		SpeakerName respjson.Field
		Voices      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceGetResponseSpeakerTrack) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceGetResponseSpeakerTrack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceAddLanguageResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceAddLanguageResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceAddLanguageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceDubResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceDubResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceDubResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceMigrateSegmentsResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceMigrateSegmentsResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceMigrateSegmentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceRenderResponse struct {
	RenderID string `json:"render_id,required"`
	Version  int64  `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RenderID    respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceRenderResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceRenderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceTranscribeResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceTranscribeResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceTranscribeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceTranslateResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceTranslateResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceTranslateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type DubbingResourceAddLanguageParams struct {
	// The Target language.
	Language param.Opt[string] `json:"language,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceAddLanguageParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceAddLanguageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceAddLanguageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceDubParams struct {
	// Dub only these languages for each segment.
	Languages []string `json:"languages,omitzero,required"`
	// Dub only this list of segments.
	Segments []string `json:"segments,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceDubParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceDubParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceDubParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceMigrateSegmentsParams struct {
	SegmentIDs []string `json:"segment_ids,omitzero,required"`
	SpeakerID  string   `json:"speaker_id,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceMigrateSegmentsParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceMigrateSegmentsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceMigrateSegmentsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceRenderParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// The type of the render. One of ['mp4', 'aac', 'mp3', 'wav', 'aaf', 'tracks_zip',
	// 'clips_zip']
	//
	// Any of "mp4", "aac", "mp3", "wav", "aaf", "tracks_zip", "clips_zip".
	RenderType DubbingResourceRenderParamsRenderType `json:"render_type,omitzero,required"`
	// Whether to normalize the volume of the rendered audio.
	NormalizeVolume param.Opt[bool] `json:"normalize_volume,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceRenderParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceRenderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceRenderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the render. One of ['mp4', 'aac', 'mp3', 'wav', 'aaf', 'tracks_zip',
// 'clips_zip']
type DubbingResourceRenderParamsRenderType string

const (
	DubbingResourceRenderParamsRenderTypeMP4       DubbingResourceRenderParamsRenderType = "mp4"
	DubbingResourceRenderParamsRenderTypeAac       DubbingResourceRenderParamsRenderType = "aac"
	DubbingResourceRenderParamsRenderTypeMP3       DubbingResourceRenderParamsRenderType = "mp3"
	DubbingResourceRenderParamsRenderTypeWav       DubbingResourceRenderParamsRenderType = "wav"
	DubbingResourceRenderParamsRenderTypeAaf       DubbingResourceRenderParamsRenderType = "aaf"
	DubbingResourceRenderParamsRenderTypeTracksZip DubbingResourceRenderParamsRenderType = "tracks_zip"
	DubbingResourceRenderParamsRenderTypeClipsZip  DubbingResourceRenderParamsRenderType = "clips_zip"
)

type DubbingResourceTranscribeParams struct {
	// Transcribe this specific list of segments.
	Segments []string `json:"segments,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceTranscribeParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceTranscribeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceTranscribeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceTranslateParams struct {
	// Translate only these languages for each segment.
	Languages []string `json:"languages,omitzero,required"`
	// Translate only this list of segments.
	Segments []string `json:"segments,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceTranslateParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceTranslateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceTranslateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
