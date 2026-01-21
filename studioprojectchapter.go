// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// StudioProjectChapterService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStudioProjectChapterService] method instead.
type StudioProjectChapterService struct {
	Options   []option.RequestOption
	Snapshots StudioProjectChapterSnapshotService
}

// NewStudioProjectChapterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStudioProjectChapterService(opts ...option.RequestOption) (r StudioProjectChapterService) {
	r = StudioProjectChapterService{}
	r.Options = opts
	r.Snapshots = NewStudioProjectChapterSnapshotService(opts...)
	return
}

// Creates a new chapter either as blank or from a URL.
func (r *StudioProjectChapterService) New(ctx context.Context, projectID string, params StudioProjectChapterNewParams, opts ...option.RequestOption) (res *StudioProjectChapterNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns information about a specific chapter.
func (r *StudioProjectChapterService) Get(ctx context.Context, chapterID string, params StudioProjectChapterGetParams, opts ...option.RequestOption) (res *ChapterWithContent, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if chapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s", params.ProjectID, chapterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a chapter.
func (r *StudioProjectChapterService) Update(ctx context.Context, chapterID string, params StudioProjectChapterUpdateParams, opts ...option.RequestOption) (res *StudioProjectChapterUpdateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if chapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s", params.ProjectID, chapterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns a list of a Studio project's chapters.
func (r *StudioProjectChapterService) List(ctx context.Context, projectID string, query StudioProjectChapterListParams, opts ...option.RequestOption) (res *StudioProjectChapterListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a chapter.
func (r *StudioProjectChapterService) Delete(ctx context.Context, chapterID string, params StudioProjectChapterDeleteParams, opts ...option.RequestOption) (res *StudioProjectChapterDeleteResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if chapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s", params.ProjectID, chapterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Starts conversion of a specific chapter.
func (r *StudioProjectChapterService) Convert(ctx context.Context, chapterID string, params StudioProjectChapterConvertParams, opts ...option.RequestOption) (res *StudioProjectChapterConvertResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if chapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s/convert", params.ProjectID, chapterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type ChapterStatistics struct {
	// The number of converted characters.
	CharactersConverted int64 `json:"characters_converted,required"`
	// The number of unconverted characters.
	CharactersUnconverted int64 `json:"characters_unconverted,required"`
	// The number of converted paragraphs.
	ParagraphsConverted int64 `json:"paragraphs_converted,required"`
	// The number of unconverted paragraphs.
	ParagraphsUnconverted int64 `json:"paragraphs_unconverted,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharactersConverted   respjson.Field
		CharactersUnconverted respjson.Field
		ParagraphsConverted   respjson.Field
		ParagraphsUnconverted respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterStatistics) RawJSON() string { return r.JSON.raw }
func (r *ChapterStatistics) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterWithContent struct {
	// Whether the chapter can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded,required"`
	// The ID of the chapter.
	ChapterID string                    `json:"chapter_id,required"`
	Content   ChapterWithContentContent `json:"content,required"`
	// The name of the chapter.
	Name string `json:"name,required"`
	// The state of the chapter.
	//
	// Any of "default", "converting".
	State ChapterWithContentState `json:"state,required"`
	// The conversion progress of the chapter.
	ConversionProgress float64 `json:"conversion_progress,nullable"`
	// Whether the chapter has a video.
	HasVideo bool `json:"has_video,nullable"`
	// The last conversion date of the chapter.
	LastConversionDateUnix int64 `json:"last_conversion_date_unix,nullable"`
	// The last conversion error of the chapter.
	LastConversionError string `json:"last_conversion_error,nullable"`
	// The statistics of the chapter.
	Statistics ChapterStatistics `json:"statistics,nullable"`
	// List of voice ids used by the chapter
	VoiceIDs []string `json:"voice_ids,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanBeDownloaded        respjson.Field
		ChapterID              respjson.Field
		Content                respjson.Field
		Name                   respjson.Field
		State                  respjson.Field
		ConversionProgress     respjson.Field
		HasVideo               respjson.Field
		LastConversionDateUnix respjson.Field
		LastConversionError    respjson.Field
		Statistics             respjson.Field
		VoiceIDs               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterWithContent) RawJSON() string { return r.JSON.raw }
func (r *ChapterWithContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterWithContentContent struct {
	Blocks []ChapterWithContentContentBlock `json:"blocks,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Blocks      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterWithContentContent) RawJSON() string { return r.JSON.raw }
func (r *ChapterWithContentContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterWithContentContentBlock struct {
	BlockID string                                    `json:"block_id,required"`
	Nodes   []ChapterWithContentContentBlockNodeUnion `json:"nodes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BlockID     respjson.Field
		Nodes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterWithContentContentBlock) RawJSON() string { return r.JSON.raw }
func (r *ChapterWithContentContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChapterWithContentContentBlockNodeUnion contains all possible properties and
// values from
// [ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel],
// [ChapterWithContentContentBlockNodeChapterContentBlockExtendableNodeResponseModel].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChapterWithContentContentBlockNodeUnion struct {
	// This field is from variant
	// [ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel].
	Text string `json:"text"`
	Type string `json:"type"`
	// This field is from variant
	// [ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel].
	VoiceID string `json:"voice_id"`
	JSON    struct {
		Text    respjson.Field
		Type    respjson.Field
		VoiceID respjson.Field
		raw     string
	} `json:"-"`
}

func (u ChapterWithContentContentBlockNodeUnion) AsChapterContentBlockTtsNodeResponseModel() (v ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChapterWithContentContentBlockNodeUnion) AsChapterContentBlockExtendableNodeResponseModel() (v ChapterWithContentContentBlockNodeChapterContentBlockExtendableNodeResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChapterWithContentContentBlockNodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ChapterWithContentContentBlockNodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel struct {
	Text    string           `json:"text,required"`
	Type    constant.TtsNode `json:"type,required"`
	VoiceID string           `json:"voice_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Type        respjson.Field
		VoiceID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *ChapterWithContentContentBlockNodeChapterContentBlockTtsNodeResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Not used. Make sure you anticipate new types in the future.
type ChapterWithContentContentBlockNodeChapterContentBlockExtendableNodeResponseModel struct {
	Type constant.Other `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChapterWithContentContentBlockNodeChapterContentBlockExtendableNodeResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *ChapterWithContentContentBlockNodeChapterContentBlockExtendableNodeResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the chapter.
type ChapterWithContentState string

const (
	ChapterWithContentStateDefault    ChapterWithContentState = "default"
	ChapterWithContentStateConverting ChapterWithContentState = "converting"
)

type StudioProjectChapterNewResponse struct {
	Chapter ChapterWithContent `json:"chapter,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapter     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterUpdateResponse struct {
	Chapter ChapterWithContent `json:"chapter,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapter     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterListResponse struct {
	Chapters []Chapter `json:"chapters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapters    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterListResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterDeleteResponse struct {
	// The status of the studio chapter deletion request. If the request was
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
func (r StudioProjectChapterDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterConvertResponse struct {
	// The status of the studio chapter conversion request. If the request was
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
func (r StudioProjectChapterConvertResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterConvertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterNewParams struct {
	// The name of the chapter, used for identification only.
	Name string `json:"name,required"`
	// An optional URL from which we will extract content to initialize the Studio
	// project. If this is set, 'from_url' and 'from_content' must be null. If neither
	// 'from_url', 'from_document', 'from_content' are provided we will initialize the
	// Studio project as blank.
	FromURL param.Opt[string] `json:"from_url,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r StudioProjectChapterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterGetParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectChapterUpdateParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// The name of the chapter, used for identification only.
	Name param.Opt[string] `json:"name,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The chapter content to use.
	Content StudioProjectChapterUpdateParamsContent `json:"content,omitzero"`
	paramObj
}

func (r StudioProjectChapterUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The chapter content to use.
//
// The property Blocks is required.
type StudioProjectChapterUpdateParamsContent struct {
	Blocks []StudioProjectChapterUpdateParamsContentBlock `json:"blocks,omitzero,required"`
	paramObj
}

func (r StudioProjectChapterUpdateParamsContent) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterUpdateParamsContent
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterUpdateParamsContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Nodes is required.
type StudioProjectChapterUpdateParamsContentBlock struct {
	Nodes   []StudioProjectChapterUpdateParamsContentBlockNode `json:"nodes,omitzero,required"`
	BlockID param.Opt[string]                                  `json:"block_id,omitzero"`
	// Any of "p", "h1", "h2", "h3".
	SubType string `json:"sub_type,omitzero"`
	paramObj
}

func (r StudioProjectChapterUpdateParamsContentBlock) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterUpdateParamsContentBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterUpdateParamsContentBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[StudioProjectChapterUpdateParamsContentBlock](
		"sub_type", "p", "h1", "h2", "h3",
	)
}

// The properties Text, Type, VoiceID are required.
type StudioProjectChapterUpdateParamsContentBlockNode struct {
	Text    string `json:"text,required"`
	VoiceID string `json:"voice_id,required"`
	// This field can be elided, and will marshal its zero value as "tts_node".
	Type constant.TtsNode `json:"type,required"`
	paramObj
}

func (r StudioProjectChapterUpdateParamsContentBlockNode) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterUpdateParamsContentBlockNode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterUpdateParamsContentBlockNode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectChapterDeleteParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectChapterConvertParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
