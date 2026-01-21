// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// DubbingService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDubbingService] method instead.
type DubbingService struct {
	Options     []option.RequestOption
	Resource    DubbingResourceService
	Transcripts DubbingTranscriptService
}

// NewDubbingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDubbingService(opts ...option.RequestOption) (r DubbingService) {
	r = DubbingService{}
	r.Options = opts
	r.Resource = NewDubbingResourceService(opts...)
	r.Transcripts = NewDubbingTranscriptService(opts...)
	return
}

// Dubs a provided audio or video file into given language.
func (r *DubbingService) New(ctx context.Context, params DubbingNewParams, opts ...option.RequestOption) (res *DubbingNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/dubbing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns metadata about a dubbing project, including whether it's still in
// progress or not
func (r *DubbingService) Get(ctx context.Context, dubbingID string, query DubbingGetParams, opts ...option.RequestOption) (res *DubbingMetadata, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/%s", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List the dubs you have access to.
func (r *DubbingService) List(ctx context.Context, params DubbingListParams, opts ...option.RequestOption) (res *DubbingListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/dubbing"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Deletes a dubbing project.
func (r *DubbingService) Delete(ctx context.Context, dubbingID string, body DubbingDeleteParams, opts ...option.RequestOption) (res *DubbingDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if dubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/%s", dubbingID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns dub as a streamed MP3 or MP4 file. If this dub has been edited using
// Dubbing Studio you need to use the resource render endpoint as this endpoint
// only returns the original automatic dub result.
func (r *DubbingService) GetAudio(ctx context.Context, languageCode string, params DubbingGetAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if languageCode == "" {
		err = errors.New("missing required language_code parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/%s/audio/%s", params.DubbingID, languageCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns transcript for the dub as an SRT or WEBVTT file.
//
// Deprecated: deprecated
func (r *DubbingService) GetTranscript(ctx context.Context, languageCode string, params DubbingGetTranscriptParams, opts ...option.RequestOption) (res *DubbingGetTranscriptResponseUnion, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if languageCode == "" {
		err = errors.New("missing required language_code parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/%s/transcript/%s", params.DubbingID, languageCode)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type DubbingMetadata struct {
	// Timestamp this dub was created.
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The ID of the dubbing project.
	DubbingID string `json:"dubbing_id,required"`
	// The name of the dubbing project.
	Name string `json:"name,required"`
	// Once dubbing has completed, the ISO-639-1 code of the original media's source
	// language.
	SourceLanguage string `json:"source_language,required"`
	// The state this dub is in.
	Status string `json:"status,required"`
	// The ISO-639-1 code of the languages this media has been dubbed into.
	TargetLanguages []string `json:"target_languages,required"`
	// Whether this dubbing project is editable in Dubbing Studio.
	Editable bool `json:"editable"`
	// Error message indicate, if this dub has failed, what happened.
	Error string `json:"error,nullable"`
	// Metadata, such as the length in seconds and content type, of the dubbed content.
	MediaMetadata DubbingMetadataMediaMetadata `json:"media_metadata,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DubbingID       respjson.Field
		Name            respjson.Field
		SourceLanguage  respjson.Field
		Status          respjson.Field
		TargetLanguages respjson.Field
		Editable        respjson.Field
		Error           respjson.Field
		MediaMetadata   respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingMetadata) RawJSON() string { return r.JSON.raw }
func (r *DubbingMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata, such as the length in seconds and content type, of the dubbed content.
type DubbingMetadataMediaMetadata struct {
	// The content type of the media.
	ContentType string `json:"content_type,required"`
	// The duration of the media in seconds.
	Duration float64 `json:"duration,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Duration    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingMetadataMediaMetadata) RawJSON() string { return r.JSON.raw }
func (r *DubbingMetadataMediaMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingTranscriptUtterance struct {
	EndS      float64                          `json:"end_s"`
	SpeakerID string                           `json:"speaker_id"`
	StartS    float64                          `json:"start_s"`
	Text      string                           `json:"text"`
	Words     []DubbingTranscriptUtteranceWord `json:"words"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndS        respjson.Field
		SpeakerID   respjson.Field
		StartS      respjson.Field
		Text        respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingTranscriptUtterance) RawJSON() string { return r.JSON.raw }
func (r *DubbingTranscriptUtterance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingTranscriptUtteranceWord struct {
	Characters []DubbingTranscriptUtteranceWordCharacter `json:"characters"`
	EndS       float64                                   `json:"end_s"`
	StartS     float64                                   `json:"start_s"`
	Text       string                                    `json:"text"`
	WordType   string                                    `json:"word_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Characters  respjson.Field
		EndS        respjson.Field
		StartS      respjson.Field
		Text        respjson.Field
		WordType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingTranscriptUtteranceWord) RawJSON() string { return r.JSON.raw }
func (r *DubbingTranscriptUtteranceWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingTranscriptUtteranceWordCharacter struct {
	EndS   float64 `json:"end_s"`
	StartS float64 `json:"start_s"`
	Text   string  `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndS        respjson.Field
		StartS      respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingTranscriptUtteranceWordCharacter) RawJSON() string { return r.JSON.raw }
func (r *DubbingTranscriptUtteranceWordCharacter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingNewResponse struct {
	// The ID of the dubbing project.
	DubbingID string `json:"dubbing_id,required"`
	// The expected duration of the dubbing project in seconds.
	ExpectedDurationSec float64 `json:"expected_duration_sec,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DubbingID           respjson.Field
		ExpectedDurationSec respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingListResponse struct {
	Dubs       []DubbingMetadata `json:"dubs,required"`
	HasMore    bool              `json:"has_more,required"`
	NextCursor string            `json:"next_cursor,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Dubs        respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingListResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingDeleteResponse struct {
	// The status of the dubbing project. If the request was successful, the status
	// will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DubbingGetTranscriptResponseUnion contains all possible properties and values
// from [DubbingGetTranscriptResponseDubbingTranscriptResponseModel], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type DubbingGetTranscriptResponseUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant
	// [DubbingGetTranscriptResponseDubbingTranscriptResponseModel].
	Language string `json:"language"`
	// This field is from variant
	// [DubbingGetTranscriptResponseDubbingTranscriptResponseModel].
	Utterances []DubbingTranscriptUtterance `json:"utterances"`
	JSON       struct {
		OfString   respjson.Field
		Language   respjson.Field
		Utterances respjson.Field
		raw        string
	} `json:"-"`
}

func (u DubbingGetTranscriptResponseUnion) AsDubbingTranscriptResponseModel() (v DubbingGetTranscriptResponseDubbingTranscriptResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DubbingGetTranscriptResponseUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DubbingGetTranscriptResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *DubbingGetTranscriptResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingGetTranscriptResponseDubbingTranscriptResponseModel struct {
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
func (r DubbingGetTranscriptResponseDubbingTranscriptResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *DubbingGetTranscriptResponseDubbingTranscriptResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingNewParams struct {
	// Frames per second to use when parsing a CSV file for dubbing. If not provided,
	// FPS will be inferred from timecodes.
	CsvFps param.Opt[float64] `json:"csv_fps,omitzero"`
	// End time of the source video/audio file.
	EndTime param.Opt[int64] `json:"end_time,omitzero"`
	// Name of the dubbing project.
	Name param.Opt[string] `json:"name,omitzero"`
	// URL of the source video/audio file.
	SourceURL param.Opt[string] `json:"source_url,omitzero"`
	// Start time of the source video/audio file.
	StartTime param.Opt[int64] `json:"start_time,omitzero"`
	// [Experimental] An accent to apply when selecting voices from the library and to
	// use to inform translation of the dialect to prefer.
	TargetAccent param.Opt[string] `json:"target_accent,omitzero"`
	// The Target language to dub the content into. Expects a valid iso639-1 or
	// iso639-3 language code.
	TargetLang param.Opt[string] `json:"target_lang,omitzero"`
	// [BETA] Whether transcripts should have profanities censored with the words
	// '[censored]'
	UseProfanityFilter param.Opt[bool] `json:"use_profanity_filter,omitzero"`
	// Instead of using a voice clone in dubbing, use a similar voice from the
	// ElevenLabs Voice Library. Voices used from the library will contribute towards a
	// workspace's custom voices limit, and if there aren't enough available slots the
	// dub will fail. Using this feature requires the caller to have the
	// 'add_voice_from_voice_library' permission on their workspace to access new
	// voices.
	DisableVoiceCloning param.Opt[bool] `json:"disable_voice_cloning,omitzero"`
	// An advanced setting. Whether to drop background audio from the final dub. This
	// can improve dub quality where it's known that audio shouldn't have a background
	// track such as for speeches or monologues.
	DropBackgroundAudio param.Opt[bool] `json:"drop_background_audio,omitzero"`
	// Whether to prepare dub for edits in dubbing studio or edits as a dubbing
	// resource.
	DubbingStudio param.Opt[bool] `json:"dubbing_studio,omitzero"`
	// Whether to use the highest resolution available.
	HighestResolution param.Opt[bool] `json:"highest_resolution,omitzero"`
	// Number of speakers to use for the dubbing. Set to 0 to automatically detect the
	// number of speakers
	NumSpeakers param.Opt[int64] `json:"num_speakers,omitzero"`
	// Source language. Expects a valid iso639-1 or iso639-3 language code.
	SourceLang param.Opt[string] `json:"source_lang,omitzero"`
	// Whether to apply watermark to the output video.
	Watermark param.Opt[bool] `json:"watermark,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// For use only with csv input
	BackgroundAudioFile io.Reader `json:"background_audio_file,omitzero" format:"binary"`
	// CSV file containing transcription/translation metadata
	CsvFile io.Reader `json:"csv_file,omitzero" format:"binary"`
	// A list of file paths to audio recordings intended for voice cloning
	File io.Reader `json:"file,omitzero" format:"binary"`
	// For use only with csv input
	ForegroundAudioFile io.Reader `json:"foreground_audio_file,omitzero" format:"binary"`
	// The mode in which to run this Dubbing job. Defaults to automatic, use manual if
	// specifically providing a CSV transcript to use. Note that manual mode is
	// experimental and production use is strongly discouraged.
	//
	// Any of "automatic", "manual".
	Mode DubbingNewParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r DubbingNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// The mode in which to run this Dubbing job. Defaults to automatic, use manual if
// specifically providing a CSV transcript to use. Note that manual mode is
// experimental and production use is strongly discouraged.
type DubbingNewParamsMode string

const (
	DubbingNewParamsModeAutomatic DubbingNewParamsMode = "automatic"
	DubbingNewParamsModeManual    DubbingNewParamsMode = "manual"
)

type DubbingGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type DubbingListParams struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// How many dubs to return at maximum. Can not exceed 200, defaults to 100.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// What state the dub is currently in.
	//
	// Any of "dubbing", "dubbed", "failed".
	DubbingStatus DubbingListParamsDubbingStatus `query:"dubbing_status,omitzero" json:"-"`
	// Filters who created the resources being listed, whether it was the user running
	// the request or someone else that shared the resource with them.
	//
	// Any of "personal", "others", "all".
	FilterByCreator DubbingListParamsFilterByCreator `query:"filter_by_creator,omitzero" json:"-"`
	// The field to use for ordering results from this query.
	//
	// Any of "created_at".
	OrderBy DubbingListParamsOrderBy `query:"order_by,omitzero" json:"-"`
	// The order direction to use for results from this query.
	//
	// Any of "DESCENDING", "ASCENDING".
	OrderDirection DubbingListParamsOrderDirection `query:"order_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DubbingListParams]'s query parameters as `url.Values`.
func (r DubbingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// What state the dub is currently in.
type DubbingListParamsDubbingStatus string

const (
	DubbingListParamsDubbingStatusDubbing DubbingListParamsDubbingStatus = "dubbing"
	DubbingListParamsDubbingStatusDubbed  DubbingListParamsDubbingStatus = "dubbed"
	DubbingListParamsDubbingStatusFailed  DubbingListParamsDubbingStatus = "failed"
)

// Filters who created the resources being listed, whether it was the user running
// the request or someone else that shared the resource with them.
type DubbingListParamsFilterByCreator string

const (
	DubbingListParamsFilterByCreatorPersonal DubbingListParamsFilterByCreator = "personal"
	DubbingListParamsFilterByCreatorOthers   DubbingListParamsFilterByCreator = "others"
	DubbingListParamsFilterByCreatorAll      DubbingListParamsFilterByCreator = "all"
)

// The field to use for ordering results from this query.
type DubbingListParamsOrderBy string

const (
	DubbingListParamsOrderByCreatedAt DubbingListParamsOrderBy = "created_at"
)

// The order direction to use for results from this query.
type DubbingListParamsOrderDirection string

const (
	DubbingListParamsOrderDirectionDescending DubbingListParamsOrderDirection = "DESCENDING"
	DubbingListParamsOrderDirectionAscending  DubbingListParamsOrderDirection = "ASCENDING"
)

type DubbingDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type DubbingGetAudioParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type DubbingGetTranscriptParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Format to return transcript in. For subtitles use either 'srt' or 'webvtt', and
	// for a full transcript use 'json'. The 'json' format is not yet supported for
	// Dubbing Studio.
	//
	// Any of "srt", "webvtt", "json".
	FormatType DubbingGetTranscriptParamsFormatType `query:"format_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DubbingGetTranscriptParams]'s query parameters as
// `url.Values`.
func (r DubbingGetTranscriptParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Format to return transcript in. For subtitles use either 'srt' or 'webvtt', and
// for a full transcript use 'json'. The 'json' format is not yet supported for
// Dubbing Studio.
type DubbingGetTranscriptParamsFormatType string

const (
	DubbingGetTranscriptParamsFormatTypeSrt    DubbingGetTranscriptParamsFormatType = "srt"
	DubbingGetTranscriptParamsFormatTypeWebvtt DubbingGetTranscriptParamsFormatType = "webvtt"
	DubbingGetTranscriptParamsFormatTypeJson   DubbingGetTranscriptParamsFormatType = "json"
)
