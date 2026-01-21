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

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// StudioProjectService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStudioProjectService] method instead.
type StudioProjectService struct {
	Options   []option.RequestOption
	Snapshots StudioProjectSnapshotService
	Chapters  StudioProjectChapterService
}

// NewStudioProjectService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewStudioProjectService(opts ...option.RequestOption) (r StudioProjectService) {
	r = StudioProjectService{}
	r.Options = opts
	r.Snapshots = NewStudioProjectSnapshotService(opts...)
	r.Chapters = NewStudioProjectChapterService(opts...)
	return
}

// Creates a new Studio project, it can be either initialized as blank, from a
// document or from a URL.
func (r *StudioProjectService) New(ctx context.Context, params StudioProjectNewParams, opts ...option.RequestOption) (res *StudioProjectNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/studio/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns information about a specific Studio project. This endpoint returns more
// detailed information about a project than `GET /v1/studio`.
func (r *StudioProjectService) Get(ctx context.Context, projectID string, params StudioProjectGetParams, opts ...option.RequestOption) (res *StudioProjectGetResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Updates the specified Studio project by setting the values of the parameters
// passed.
func (r *StudioProjectService) Update(ctx context.Context, projectID string, params StudioProjectUpdateParams, opts ...option.RequestOption) (res *EditProject, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns a list of your Studio projects with metadata.
func (r *StudioProjectService) List(ctx context.Context, query StudioProjectListParams, opts ...option.RequestOption) (res *StudioProjectListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/studio/projects"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a Studio project.
func (r *StudioProjectService) Delete(ctx context.Context, projectID string, body StudioProjectDeleteParams, opts ...option.RequestOption) (res *StudioProjectDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Starts conversion of a Studio project and all of its chapters.
func (r *StudioProjectService) Convert(ctx context.Context, projectID string, body StudioProjectConvertParams, opts ...option.RequestOption) (res *StudioProjectConvertResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/convert", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create a set of pronunciation dictionaries acting on a project. This will
// automatically mark text within this project as requiring reconverting where the
// new dictionary would apply or the old one no longer does.
func (r *StudioProjectService) NewPronunciationDictionary(ctx context.Context, projectID string, params StudioProjectNewPronunciationDictionaryParams, opts ...option.RequestOption) (res *StudioProjectNewPronunciationDictionaryResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/pronunciation-dictionaries", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates Studio project content.
func (r *StudioProjectService) UpdateContent(ctx context.Context, projectID string, params StudioProjectUpdateContentParams, opts ...option.RequestOption) (res *EditProject, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/content", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Defines asset positioning and transformation on canvas.
type CanvasPlacement struct {
	PivotX float64 `json:"pivot_x"`
	PivotY float64 `json:"pivot_y"`
	ScaleX float64 `json:"scale_x"`
	ScaleY float64 `json:"scale_y"`
	SkewX  float64 `json:"skew_x"`
	SkewY  float64 `json:"skew_y"`
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PivotX      respjson.Field
		PivotY      respjson.Field
		ScaleX      respjson.Field
		ScaleY      respjson.Field
		SkewX       respjson.Field
		SkewY       respjson.Field
		X           respjson.Field
		Y           respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CanvasPlacement) RawJSON() string { return r.JSON.raw }
func (r *CanvasPlacement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyle struct {
	AutoBreakEnabled    bool                            `json:"auto_break_enabled,nullable"`
	BackgroundColor     string                          `json:"background_color,nullable"`
	BackgroundEnabled   bool                            `json:"background_enabled,nullable"`
	BackgroundOpacity   float64                         `json:"background_opacity,nullable"`
	CharacterAnimation  CaptionStyleCharacterAnimation  `json:"character_animation,nullable"`
	HorizontalPlacement CaptionStyleHorizontalPlacement `json:"horizontal_placement,nullable"`
	MaxLinesPerSection  int64                           `json:"max_lines_per_section,nullable"`
	MaxWordsPerLine     int64                           `json:"max_words_per_line,nullable"`
	SectionAnimation    CaptionStyleSectionAnimation    `json:"section_animation,nullable"`
	Template            CaptionStyleTemplate            `json:"template,nullable"`
	// Any of "start", "center", "end".
	TextAlign CaptionStyleTextAlign `json:"text_align,nullable"`
	TextColor string                `json:"text_color,nullable"`
	TextFont  string                `json:"text_font,nullable"`
	TextScale float64               `json:"text_scale,nullable"`
	// Any of "normal", "italic".
	TextStyle CaptionStyleTextStyle `json:"text_style,nullable"`
	// Any of "normal", "bold".
	TextWeight                    CaptionStyleTextWeight        `json:"text_weight,nullable"`
	VerticalPlacement             CaptionStyleVerticalPlacement `json:"vertical_placement,nullable"`
	WidthPct                      float64                       `json:"width_pct,nullable"`
	WordAnimation                 CaptionStyleWordAnimation     `json:"word_animation,nullable"`
	WordHighlightsBackgroundColor string                        `json:"word_highlights_background_color,nullable"`
	WordHighlightsColor           string                        `json:"word_highlights_color,nullable"`
	WordHighlightsEnabled         bool                          `json:"word_highlights_enabled,nullable"`
	WordHighlightsOpacity         float64                       `json:"word_highlights_opacity,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AutoBreakEnabled              respjson.Field
		BackgroundColor               respjson.Field
		BackgroundEnabled             respjson.Field
		BackgroundOpacity             respjson.Field
		CharacterAnimation            respjson.Field
		HorizontalPlacement           respjson.Field
		MaxLinesPerSection            respjson.Field
		MaxWordsPerLine               respjson.Field
		SectionAnimation              respjson.Field
		Template                      respjson.Field
		TextAlign                     respjson.Field
		TextColor                     respjson.Field
		TextFont                      respjson.Field
		TextScale                     respjson.Field
		TextStyle                     respjson.Field
		TextWeight                    respjson.Field
		VerticalPlacement             respjson.Field
		WidthPct                      respjson.Field
		WordAnimation                 respjson.Field
		WordHighlightsBackgroundColor respjson.Field
		WordHighlightsColor           respjson.Field
		WordHighlightsEnabled         respjson.Field
		WordHighlightsOpacity         respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyle) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleCharacterAnimation struct {
	// Any of "none", "fade".
	EnterType string `json:"enter_type,required"`
	// Any of "none", "fade".
	ExitType string `json:"exit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnterType   respjson.Field
		ExitType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleCharacterAnimation) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleCharacterAnimation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleHorizontalPlacement struct {
	// Any of "left", "center", "right".
	Align        string  `json:"align,required"`
	TranslatePct float64 `json:"translate_pct,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Align        respjson.Field
		TranslatePct respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleHorizontalPlacement) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleHorizontalPlacement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleSectionAnimation struct {
	// Any of "none", "fade", "scale".
	EnterType string `json:"enter_type,required"`
	// Any of "none", "fade", "scale".
	ExitType string `json:"exit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnterType   respjson.Field
		ExitType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleSectionAnimation) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleSectionAnimation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleTemplate struct {
	Key             string `json:"key,required"`
	Label           string `json:"label,required"`
	RequiresHighFps bool   `json:"requires_high_fps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key             respjson.Field
		Label           respjson.Field
		RequiresHighFps respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleTemplate) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleTextAlign string

const (
	CaptionStyleTextAlignStart  CaptionStyleTextAlign = "start"
	CaptionStyleTextAlignCenter CaptionStyleTextAlign = "center"
	CaptionStyleTextAlignEnd    CaptionStyleTextAlign = "end"
)

type CaptionStyleTextStyle string

const (
	CaptionStyleTextStyleNormal CaptionStyleTextStyle = "normal"
	CaptionStyleTextStyleItalic CaptionStyleTextStyle = "italic"
)

type CaptionStyleTextWeight string

const (
	CaptionStyleTextWeightNormal CaptionStyleTextWeight = "normal"
	CaptionStyleTextWeightBold   CaptionStyleTextWeight = "bold"
)

type CaptionStyleVerticalPlacement struct {
	// Any of "top", "center", "bottom".
	Align        string  `json:"align,required"`
	TranslatePct float64 `json:"translate_pct,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Align        respjson.Field
		TranslatePct respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleVerticalPlacement) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleVerticalPlacement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CaptionStyleWordAnimation struct {
	// Any of "none", "fade", "scale".
	EnterType string `json:"enter_type,required"`
	// Any of "none", "fade", "scale".
	ExitType string `json:"exit_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnterType   respjson.Field
		ExitType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CaptionStyleWordAnimation) RawJSON() string { return r.JSON.raw }
func (r *CaptionStyleWordAnimation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Chapter struct {
	// Whether the chapter can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded,required"`
	// The ID of the chapter.
	ChapterID string `json:"chapter_id,required"`
	// The name of the chapter.
	Name string `json:"name,required"`
	// The state of the chapter.
	//
	// Any of "default", "converting".
	State ChapterState `json:"state,required"`
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
func (r Chapter) RawJSON() string { return r.JSON.raw }
func (r *Chapter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the chapter.
type ChapterState string

const (
	ChapterStateDefault    ChapterState = "default"
	ChapterStateConverting ChapterState = "converting"
)

type EditProject struct {
	Project Project `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EditProject) RawJSON() string { return r.JSON.raw }
func (r *EditProject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectCreationMeta struct {
	// The progress of the project creation.
	CreationProgress float64 `json:"creation_progress,required"`
	// The status of the project creation action.
	//
	// Any of "pending", "creating", "finished", "failed".
	Status ProjectCreationMetaStatus `json:"status,required"`
	// The type of the project creation action.
	//
	// Any of "blank", "generate_podcast", "auto_assign_voices", "dub_video".
	Type ProjectCreationMetaType `json:"type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreationProgress respjson.Field
		Status           respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectCreationMeta) RawJSON() string { return r.JSON.raw }
func (r *ProjectCreationMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the project creation action.
type ProjectCreationMetaStatus string

const (
	ProjectCreationMetaStatusPending  ProjectCreationMetaStatus = "pending"
	ProjectCreationMetaStatusCreating ProjectCreationMetaStatus = "creating"
	ProjectCreationMetaStatusFinished ProjectCreationMetaStatus = "finished"
	ProjectCreationMetaStatusFailed   ProjectCreationMetaStatus = "failed"
)

// The type of the project creation action.
type ProjectCreationMetaType string

const (
	ProjectCreationMetaTypeBlank            ProjectCreationMetaType = "blank"
	ProjectCreationMetaTypeGeneratePodcast  ProjectCreationMetaType = "generate_podcast"
	ProjectCreationMetaTypeAutoAssignVoices ProjectCreationMetaType = "auto_assign_voices"
	ProjectCreationMetaTypeDubVideo         ProjectCreationMetaType = "dub_video"
)

type ReadLegalTerms struct {
	EndDate   string `json:"end_date,nullable"`
	StartDate string `json:"start_date,nullable"`
	Terms     string `json:"terms,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndDate     respjson.Field
		StartDate   respjson.Field
		Terms       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReadLegalTerms) RawJSON() string { return r.JSON.raw }
func (r *ReadLegalTerms) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectNewResponse struct {
	Project Project `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectNewResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponse struct {
	// The access level of the project.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	AccessLevel StudioProjectGetResponseAccessLevel `json:"access_level,required"`
	// Whether text normalization is applied to the project.
	//
	// Any of "auto", "on", "off", "apply_english".
	ApplyTextNormalization StudioProjectGetResponseApplyTextNormalization `json:"apply_text_normalization,required"`
	// List of uploaded assets e.g. videos, audios.
	Assets []StudioProjectGetResponseAssetUnion `json:"assets,required"`
	// Whether the project can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded,required"`
	// List of chapters of the project and their metadata.
	Chapters []Chapter `json:"chapters,required"`
	// The creation date of the project.
	CreateDateUnix int64 `json:"create_date_unix,required"`
	// The user ID who created the project.
	CreatedByUserID string `json:"created_by_user_id,required"`
	// The default model ID.
	DefaultModelID string `json:"default_model_id,required"`
	// The default paragraph voice ID.
	DefaultParagraphVoiceID string `json:"default_paragraph_voice_id,required"`
	// The default title voice ID.
	DefaultTitleVoiceID string `json:"default_title_voice_id,required"`
	// The name of the project.
	Name string `json:"name,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// List of pronunciation dictionary locators.
	PronunciationDictionaryLocators []StudioProjectGetResponsePronunciationDictionaryLocator `json:"pronunciation_dictionary_locators,required"`
	// List of pronunciation dictionary versions of the project and their metadata.
	PronunciationDictionaryVersions []StudioProjectGetResponsePronunciationDictionaryVersion `json:"pronunciation_dictionary_versions,required"`
	// Whether quality check is enabled for this project.
	//
	// Deprecated: deprecated
	QualityCheckOn bool `json:"quality_check_on,required"`
	// Whether quality check is enabled on the project when bulk converting.
	//
	// Deprecated: deprecated
	QualityCheckOnWhenBulkConvert bool `json:"quality_check_on_when_bulk_convert,required"`
	// The quality preset level of the project.
	//
	// Any of "standard", "high", "highest", "ultra", "ultra_lossless".
	QualityPreset StudioProjectGetResponseQualityPreset `json:"quality_preset,required"`
	// The state of the project.
	//
	// Any of "creating", "default", "converting", "in_queue".
	State StudioProjectGetResponseState `json:"state,required"`
	// List of configured project voices.
	Voices []StudioProjectGetResponseVoice `json:"voices,required"`
	// Whether the project uses volume normalization.
	VolumeNormalization bool `json:"volume_normalization,required"`
	// The aspect ratio of the project.
	//
	// Any of "16:9", "9:16", "4:5", "1:1".
	AspectRatio StudioProjectGetResponseAspectRatio `json:"aspect_ratio,nullable"`
	// The author of the project.
	Author string `json:"author,nullable"`
	// List of voices used by the project.
	BaseVoices []Voice `json:"base_voices,nullable"`
	// Global styling to be applied to all captions
	CaptionStyle CaptionStyle `json:"caption_style,nullable"`
	// Whether captions are enabled for the project.
	CaptionsEnabled bool `json:"captions_enabled,nullable"`
	// Whether chapters are enabled for the project.
	ChaptersEnabled bool `json:"chapters_enabled,nullable"`
	// The content type of the project, e.g. 'Novel' or 'Short Story'
	ContentType string `json:"content_type,nullable"`
	// The cover image URL of the project.
	CoverImageURL string `json:"cover_image_url,nullable"`
	// The creation meta of the project.
	CreationMeta ProjectCreationMeta `json:"creation_meta,nullable"`
	// The description of the project.
	Description string `json:"description,nullable"`
	// Experimental features for the project.
	Experimental any `json:"experimental"`
	// Whether the project is fiction.
	//
	// Any of "fiction", "non-fiction".
	Fiction StudioProjectGetResponseFiction `json:"fiction,nullable"`
	// List of genres of the project.
	Genres []string `json:"genres,nullable"`
	// The ISBN number of the project.
	IsbnNumber string `json:"isbn_number,nullable"`
	// Two-letter language code (ISO 639-1) of the language of the project.
	Language string `json:"language,nullable"`
	// The last conversion date of the project.
	LastConversionDateUnix int64 `json:"last_conversion_date_unix,nullable"`
	// Whether the project contains mature content.
	MatureContent bool `json:"mature_content,nullable"`
	// The original publication date of the project.
	OriginalPublicationDate string `json:"original_publication_date,nullable"`
	// The public share ID of the project.
	PublicShareID string `json:"public_share_id,nullable"`
	// The ElevenReader data if the book was published.
	PublishingRead StudioProjectGetResponsePublishingRead `json:"publishing_read,nullable"`
	// The source type of the project.
	//
	// Any of "blank", "book", "article", "genfm", "video", "screenplay".
	SourceType StudioProjectGetResponseSourceType `json:"source_type,nullable"`
	// The target audience of the project.
	//
	// Any of "children", "young adult", "adult", "all ages".
	TargetAudience StudioProjectGetResponseTargetAudience `json:"target_audience,nullable"`
	// The title of the project.
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel                     respjson.Field
		ApplyTextNormalization          respjson.Field
		Assets                          respjson.Field
		CanBeDownloaded                 respjson.Field
		Chapters                        respjson.Field
		CreateDateUnix                  respjson.Field
		CreatedByUserID                 respjson.Field
		DefaultModelID                  respjson.Field
		DefaultParagraphVoiceID         respjson.Field
		DefaultTitleVoiceID             respjson.Field
		Name                            respjson.Field
		ProjectID                       respjson.Field
		PronunciationDictionaryLocators respjson.Field
		PronunciationDictionaryVersions respjson.Field
		QualityCheckOn                  respjson.Field
		QualityCheckOnWhenBulkConvert   respjson.Field
		QualityPreset                   respjson.Field
		State                           respjson.Field
		Voices                          respjson.Field
		VolumeNormalization             respjson.Field
		AspectRatio                     respjson.Field
		Author                          respjson.Field
		BaseVoices                      respjson.Field
		CaptionStyle                    respjson.Field
		CaptionsEnabled                 respjson.Field
		ChaptersEnabled                 respjson.Field
		ContentType                     respjson.Field
		CoverImageURL                   respjson.Field
		CreationMeta                    respjson.Field
		Description                     respjson.Field
		Experimental                    respjson.Field
		Fiction                         respjson.Field
		Genres                          respjson.Field
		IsbnNumber                      respjson.Field
		Language                        respjson.Field
		LastConversionDateUnix          respjson.Field
		MatureContent                   respjson.Field
		OriginalPublicationDate         respjson.Field
		PublicShareID                   respjson.Field
		PublishingRead                  respjson.Field
		SourceType                      respjson.Field
		TargetAudience                  respjson.Field
		Title                           respjson.Field
		ExtraFields                     map[string]respjson.Field
		raw                             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The access level of the project.
type StudioProjectGetResponseAccessLevel string

const (
	StudioProjectGetResponseAccessLevelAdmin     StudioProjectGetResponseAccessLevel = "admin"
	StudioProjectGetResponseAccessLevelEditor    StudioProjectGetResponseAccessLevel = "editor"
	StudioProjectGetResponseAccessLevelCommenter StudioProjectGetResponseAccessLevel = "commenter"
	StudioProjectGetResponseAccessLevelViewer    StudioProjectGetResponseAccessLevel = "viewer"
)

// Whether text normalization is applied to the project.
type StudioProjectGetResponseApplyTextNormalization string

const (
	StudioProjectGetResponseApplyTextNormalizationAuto         StudioProjectGetResponseApplyTextNormalization = "auto"
	StudioProjectGetResponseApplyTextNormalizationOn           StudioProjectGetResponseApplyTextNormalization = "on"
	StudioProjectGetResponseApplyTextNormalizationOff          StudioProjectGetResponseApplyTextNormalization = "off"
	StudioProjectGetResponseApplyTextNormalizationApplyEnglish StudioProjectGetResponseApplyTextNormalization = "apply_english"
)

// StudioProjectGetResponseAssetUnion contains all possible properties and values
// from [StudioProjectGetResponseAssetProjectVideoResponseModel],
// [StudioProjectGetResponseAssetProjectExternalAudioResponseModel],
// [StudioProjectGetResponseAssetProjectImageResponseModel].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type StudioProjectGetResponseAssetUnion struct {
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	Codec           string   `json:"codec"`
	CreatedAtMs     int64    `json:"created_at_ms"`
	DurationMs      int64    `json:"duration_ms"`
	EndTimeMs       int64    `json:"end_time_ms"`
	Filename        string   `json:"filename"`
	Height          int64    `json:"height"`
	Muted           bool     `json:"muted"`
	OffsetMs        int64    `json:"offset_ms"`
	Order           string   `json:"order"`
	PendingBlockIDs []string `json:"pending_block_ids"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	PreviewJobProgress float64 `json:"preview_job_progress"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	SignedPreviewURL string `json:"signed_preview_url"`
	SignedURL        string `json:"signed_url"`
	StartTimeMs      int64  `json:"start_time_ms"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	ThumbnailIntervalSeconds float64 `json:"thumbnail_interval_seconds"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	ThumbnailSheets []StudioProjectGetResponseAssetProjectVideoResponseModelThumbnailSheet `json:"thumbnail_sheets"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	ThumbnailSize []int64 `json:"thumbnail_size"`
	UpdatedAtMs   int64   `json:"updated_at_ms"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	VideoID      string  `json:"video_id"`
	VolumeGainDB float64 `json:"volume_gain_db"`
	Width        int64   `json:"width"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	AssetPreviewSignedURL string `json:"asset_preview_signed_url"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	AudioTrackReady bool `json:"audio_track_ready"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	CanvasPlacement   CanvasPlacement `json:"canvas_placement"`
	CurrentSnapshotID string          `json:"current_snapshot_id"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	Error                string  `json:"error"`
	ImportSpeechProgress float64 `json:"import_speech_progress"`
	SourceAssetID        string  `json:"source_asset_id"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectVideoResponseModel].
	SourceVideoID  string `json:"source_video_id"`
	SpeechImported bool   `json:"speech_imported"`
	TrackID        string `json:"track_id"`
	Type           string `json:"type"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectExternalAudioResponseModel].
	ExternalAudioID string `json:"external_audio_id"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectExternalAudioResponseModel].
	FadeInMs int64 `json:"fade_in_ms"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectExternalAudioResponseModel].
	FadeOutMs int64 `json:"fade_out_ms"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectExternalAudioResponseModel].
	SourceExternalAudioID string `json:"source_external_audio_id"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectImageResponseModel].
	FileSizeBytes int64 `json:"file_size_bytes"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectImageResponseModel].
	ImageID string `json:"image_id"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectImageResponseModel].
	ThumbnailSignedURL string `json:"thumbnail_signed_url"`
	// This field is from variant
	// [StudioProjectGetResponseAssetProjectImageResponseModel].
	Source string `json:"source"`
	JSON   struct {
		Codec                    respjson.Field
		CreatedAtMs              respjson.Field
		DurationMs               respjson.Field
		EndTimeMs                respjson.Field
		Filename                 respjson.Field
		Height                   respjson.Field
		Muted                    respjson.Field
		OffsetMs                 respjson.Field
		Order                    respjson.Field
		PendingBlockIDs          respjson.Field
		PreviewJobProgress       respjson.Field
		SignedPreviewURL         respjson.Field
		SignedURL                respjson.Field
		StartTimeMs              respjson.Field
		ThumbnailIntervalSeconds respjson.Field
		ThumbnailSheets          respjson.Field
		ThumbnailSize            respjson.Field
		UpdatedAtMs              respjson.Field
		VideoID                  respjson.Field
		VolumeGainDB             respjson.Field
		Width                    respjson.Field
		AssetPreviewSignedURL    respjson.Field
		AudioTrackReady          respjson.Field
		CanvasPlacement          respjson.Field
		CurrentSnapshotID        respjson.Field
		Error                    respjson.Field
		ImportSpeechProgress     respjson.Field
		SourceAssetID            respjson.Field
		SourceVideoID            respjson.Field
		SpeechImported           respjson.Field
		TrackID                  respjson.Field
		Type                     respjson.Field
		ExternalAudioID          respjson.Field
		FadeInMs                 respjson.Field
		FadeOutMs                respjson.Field
		SourceExternalAudioID    respjson.Field
		FileSizeBytes            respjson.Field
		ImageID                  respjson.Field
		ThumbnailSignedURL       respjson.Field
		Source                   respjson.Field
		raw                      string
	} `json:"-"`
}

func (u StudioProjectGetResponseAssetUnion) AsProjectVideoResponseModel() (v StudioProjectGetResponseAssetProjectVideoResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StudioProjectGetResponseAssetUnion) AsProjectExternalAudioResponseModel() (v StudioProjectGetResponseAssetProjectExternalAudioResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u StudioProjectGetResponseAssetUnion) AsProjectImageResponseModel() (v StudioProjectGetResponseAssetProjectImageResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u StudioProjectGetResponseAssetUnion) RawJSON() string { return u.JSON.raw }

func (r *StudioProjectGetResponseAssetUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponseAssetProjectVideoResponseModel struct {
	Codec                    string                                                                 `json:"codec,required"`
	CreatedAtMs              int64                                                                  `json:"created_at_ms,required"`
	DurationMs               int64                                                                  `json:"duration_ms,required"`
	EndTimeMs                int64                                                                  `json:"end_time_ms,required"`
	Filename                 string                                                                 `json:"filename,required"`
	Height                   int64                                                                  `json:"height,required"`
	Muted                    bool                                                                   `json:"muted,required"`
	OffsetMs                 int64                                                                  `json:"offset_ms,required"`
	Order                    string                                                                 `json:"order,required"`
	PendingBlockIDs          []string                                                               `json:"pending_block_ids,required"`
	PreviewJobProgress       float64                                                                `json:"preview_job_progress,required"`
	SignedPreviewURL         string                                                                 `json:"signed_preview_url,required"`
	SignedURL                string                                                                 `json:"signed_url,required"`
	StartTimeMs              int64                                                                  `json:"start_time_ms,required"`
	ThumbnailIntervalSeconds float64                                                                `json:"thumbnail_interval_seconds,required"`
	ThumbnailSheets          []StudioProjectGetResponseAssetProjectVideoResponseModelThumbnailSheet `json:"thumbnail_sheets,required"`
	ThumbnailSize            []int64                                                                `json:"thumbnail_size,required"`
	UpdatedAtMs              int64                                                                  `json:"updated_at_ms,required"`
	VideoID                  string                                                                 `json:"video_id,required"`
	VolumeGainDB             float64                                                                `json:"volume_gain_db,required"`
	Width                    int64                                                                  `json:"width,required"`
	AssetPreviewSignedURL    string                                                                 `json:"asset_preview_signed_url,nullable"`
	AudioTrackReady          bool                                                                   `json:"audio_track_ready"`
	// Defines asset positioning and transformation on canvas.
	CanvasPlacement      CanvasPlacement `json:"canvas_placement"`
	CurrentSnapshotID    string          `json:"current_snapshot_id,nullable"`
	Error                string          `json:"error,nullable"`
	ImportSpeechProgress float64         `json:"import_speech_progress,nullable"`
	SourceAssetID        string          `json:"source_asset_id,nullable"`
	SourceVideoID        string          `json:"source_video_id,nullable"`
	SpeechImported       bool            `json:"speech_imported"`
	TrackID              string          `json:"track_id"`
	// Any of "video".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Codec                    respjson.Field
		CreatedAtMs              respjson.Field
		DurationMs               respjson.Field
		EndTimeMs                respjson.Field
		Filename                 respjson.Field
		Height                   respjson.Field
		Muted                    respjson.Field
		OffsetMs                 respjson.Field
		Order                    respjson.Field
		PendingBlockIDs          respjson.Field
		PreviewJobProgress       respjson.Field
		SignedPreviewURL         respjson.Field
		SignedURL                respjson.Field
		StartTimeMs              respjson.Field
		ThumbnailIntervalSeconds respjson.Field
		ThumbnailSheets          respjson.Field
		ThumbnailSize            respjson.Field
		UpdatedAtMs              respjson.Field
		VideoID                  respjson.Field
		VolumeGainDB             respjson.Field
		Width                    respjson.Field
		AssetPreviewSignedURL    respjson.Field
		AudioTrackReady          respjson.Field
		CanvasPlacement          respjson.Field
		CurrentSnapshotID        respjson.Field
		Error                    respjson.Field
		ImportSpeechProgress     respjson.Field
		SourceAssetID            respjson.Field
		SourceVideoID            respjson.Field
		SpeechImported           respjson.Field
		TrackID                  respjson.Field
		Type                     respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponseAssetProjectVideoResponseModel) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponseAssetProjectVideoResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponseAssetProjectVideoResponseModelThumbnailSheet struct {
	SignedCloudURL      string `json:"signed_cloud_url,required"`
	StartThumbnailIndex int64  `json:"start_thumbnail_index,required"`
	ThumbnailCount      int64  `json:"thumbnail_count,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignedCloudURL      respjson.Field
		StartThumbnailIndex respjson.Field
		ThumbnailCount      respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponseAssetProjectVideoResponseModelThumbnailSheet) RawJSON() string {
	return r.JSON.raw
}
func (r *StudioProjectGetResponseAssetProjectVideoResponseModelThumbnailSheet) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponseAssetProjectExternalAudioResponseModel struct {
	CreatedAtMs           int64    `json:"created_at_ms,required"`
	DurationMs            int64    `json:"duration_ms,required"`
	EndTimeMs             int64    `json:"end_time_ms,required"`
	ExternalAudioID       string   `json:"external_audio_id,required"`
	Filename              string   `json:"filename,required"`
	OffsetMs              int64    `json:"offset_ms,required"`
	Order                 string   `json:"order,required"`
	PendingBlockIDs       []string `json:"pending_block_ids,required"`
	SignedURL             string   `json:"signed_url,required"`
	StartTimeMs           int64    `json:"start_time_ms,required"`
	TrackID               string   `json:"track_id,required"`
	UpdatedAtMs           int64    `json:"updated_at_ms,required"`
	CurrentSnapshotID     string   `json:"current_snapshot_id,nullable"`
	FadeInMs              int64    `json:"fade_in_ms"`
	FadeOutMs             int64    `json:"fade_out_ms"`
	ImportSpeechProgress  float64  `json:"import_speech_progress,nullable"`
	Muted                 bool     `json:"muted"`
	SourceAssetID         string   `json:"source_asset_id,nullable"`
	SourceExternalAudioID string   `json:"source_external_audio_id,nullable"`
	SpeechImported        bool     `json:"speech_imported"`
	// Any of "audio".
	Type         string  `json:"type"`
	VolumeGainDB float64 `json:"volume_gain_db"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAtMs           respjson.Field
		DurationMs            respjson.Field
		EndTimeMs             respjson.Field
		ExternalAudioID       respjson.Field
		Filename              respjson.Field
		OffsetMs              respjson.Field
		Order                 respjson.Field
		PendingBlockIDs       respjson.Field
		SignedURL             respjson.Field
		StartTimeMs           respjson.Field
		TrackID               respjson.Field
		UpdatedAtMs           respjson.Field
		CurrentSnapshotID     respjson.Field
		FadeInMs              respjson.Field
		FadeOutMs             respjson.Field
		ImportSpeechProgress  respjson.Field
		Muted                 respjson.Field
		SourceAssetID         respjson.Field
		SourceExternalAudioID respjson.Field
		SpeechImported        respjson.Field
		Type                  respjson.Field
		VolumeGainDB          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponseAssetProjectExternalAudioResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *StudioProjectGetResponseAssetProjectExternalAudioResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponseAssetProjectImageResponseModel struct {
	// Defines asset positioning and transformation on canvas.
	CanvasPlacement    CanvasPlacement `json:"canvas_placement,required"`
	CreatedAtMs        int64           `json:"created_at_ms,required"`
	DurationMs         int64           `json:"duration_ms,required"`
	FileSizeBytes      int64           `json:"file_size_bytes,required"`
	Filename           string          `json:"filename,required"`
	Height             int64           `json:"height,required"`
	ImageID            string          `json:"image_id,required"`
	OffsetMs           int64           `json:"offset_ms,required"`
	Order              string          `json:"order,required"`
	SignedURL          string          `json:"signed_url,required"`
	ThumbnailSignedURL string          `json:"thumbnail_signed_url,required"`
	UpdatedAtMs        int64           `json:"updated_at_ms,required"`
	Width              int64           `json:"width,required"`
	CurrentSnapshotID  string          `json:"current_snapshot_id,nullable"`
	// Any of "upload".
	Source        string `json:"source"`
	SourceAssetID string `json:"source_asset_id,nullable"`
	TrackID       string `json:"track_id"`
	// Any of "image".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanvasPlacement    respjson.Field
		CreatedAtMs        respjson.Field
		DurationMs         respjson.Field
		FileSizeBytes      respjson.Field
		Filename           respjson.Field
		Height             respjson.Field
		ImageID            respjson.Field
		OffsetMs           respjson.Field
		Order              respjson.Field
		SignedURL          respjson.Field
		ThumbnailSignedURL respjson.Field
		UpdatedAtMs        respjson.Field
		Width              respjson.Field
		CurrentSnapshotID  respjson.Field
		Source             respjson.Field
		SourceAssetID      respjson.Field
		TrackID            respjson.Field
		Type               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponseAssetProjectImageResponseModel) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponseAssetProjectImageResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePronunciationDictionaryLocator struct {
	PronunciationDictionaryID string `json:"pronunciation_dictionary_id,required"`
	VersionID                 string `json:"version_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PronunciationDictionaryID respjson.Field
		VersionID                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePronunciationDictionaryLocator) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePronunciationDictionaryLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePronunciationDictionaryVersion struct {
	CreatedBy        string `json:"created_by,required"`
	CreationTimeUnix int64  `json:"creation_time_unix,required"`
	DictionaryName   string `json:"dictionary_name,required"`
	// Any of "admin", "editor", "commenter", "viewer".
	PermissionOnResource      string `json:"permission_on_resource,required"`
	PronunciationDictionaryID string `json:"pronunciation_dictionary_id,required"`
	VersionID                 string `json:"version_id,required"`
	VersionName               string `json:"version_name,required"`
	VersionRulesNum           int64  `json:"version_rules_num,required"`
	ArchivedTimeUnix          int64  `json:"archived_time_unix,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy                 respjson.Field
		CreationTimeUnix          respjson.Field
		DictionaryName            respjson.Field
		PermissionOnResource      respjson.Field
		PronunciationDictionaryID respjson.Field
		VersionID                 respjson.Field
		VersionName               respjson.Field
		VersionRulesNum           respjson.Field
		ArchivedTimeUnix          respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePronunciationDictionaryVersion) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePronunciationDictionaryVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The quality preset level of the project.
type StudioProjectGetResponseQualityPreset string

const (
	StudioProjectGetResponseQualityPresetStandard      StudioProjectGetResponseQualityPreset = "standard"
	StudioProjectGetResponseQualityPresetHigh          StudioProjectGetResponseQualityPreset = "high"
	StudioProjectGetResponseQualityPresetHighest       StudioProjectGetResponseQualityPreset = "highest"
	StudioProjectGetResponseQualityPresetUltra         StudioProjectGetResponseQualityPreset = "ultra"
	StudioProjectGetResponseQualityPresetUltraLossless StudioProjectGetResponseQualityPreset = "ultra_lossless"
)

// The state of the project.
type StudioProjectGetResponseState string

const (
	StudioProjectGetResponseStateCreating   StudioProjectGetResponseState = "creating"
	StudioProjectGetResponseStateDefault    StudioProjectGetResponseState = "default"
	StudioProjectGetResponseStateConverting StudioProjectGetResponseState = "converting"
	StudioProjectGetResponseStateInQueue    StudioProjectGetResponseState = "in_queue"
)

type StudioProjectGetResponseVoice struct {
	Alias           string  `json:"alias,required"`
	IsPinned        bool    `json:"is_pinned,required"`
	SimilarityBoost float64 `json:"similarity_boost,required"`
	Speed           float64 `json:"speed,required"`
	Stability       float64 `json:"stability,required"`
	Style           float64 `json:"style,required"`
	UseSpeakerBoost bool    `json:"use_speaker_boost,required"`
	VoiceID         string  `json:"voice_id,required"`
	VolumeGain      float64 `json:"volume_gain,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias           respjson.Field
		IsPinned        respjson.Field
		SimilarityBoost respjson.Field
		Speed           respjson.Field
		Stability       respjson.Field
		Style           respjson.Field
		UseSpeakerBoost respjson.Field
		VoiceID         respjson.Field
		VolumeGain      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponseVoice) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponseVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The aspect ratio of the project.
type StudioProjectGetResponseAspectRatio string

const (
	StudioProjectGetResponseAspectRatio16_9 StudioProjectGetResponseAspectRatio = "16:9"
	StudioProjectGetResponseAspectRatio9_16 StudioProjectGetResponseAspectRatio = "9:16"
	StudioProjectGetResponseAspectRatio4_5  StudioProjectGetResponseAspectRatio = "4:5"
	StudioProjectGetResponseAspectRatio1_1  StudioProjectGetResponseAspectRatio = "1:1"
)

// Whether the project is fiction.
type StudioProjectGetResponseFiction string

const (
	StudioProjectGetResponseFictionFiction    StudioProjectGetResponseFiction = "fiction"
	StudioProjectGetResponseFictionNonFiction StudioProjectGetResponseFiction = "non-fiction"
)

// The ElevenReader data if the book was published.
type StudioProjectGetResponsePublishingRead struct {
	Chapters               []StudioProjectGetResponsePublishingReadChapter     `json:"chapters,required"`
	CharCount              int64                                               `json:"char_count,required"`
	CreatedAtUnix          int64                                               `json:"created_at_unix,required"`
	ReadID                 string                                              `json:"read_id,required"`
	UpdatedAtUnix          int64                                               `json:"updated_at_unix,required"`
	WordCount              int64                                               `json:"word_count,required"`
	ArticleImageURL        string                                              `json:"article_image_url,nullable"`
	Author                 string                                              `json:"author,nullable"`
	CanUseAssistant        bool                                                `json:"can_use_assistant"`
	ContentGuidelinesTerms ReadLegalTerms                                      `json:"content_guidelines_terms,nullable"`
	ContentType            string                                              `json:"content_type,nullable"`
	Contributors           []StudioProjectGetResponsePublishingReadContributor `json:"contributors,nullable"`
	Copyright              string                                              `json:"copyright,nullable"`
	// Any of "usd".
	Currency    string `json:"currency,nullable"`
	Description string `json:"description,nullable"`
	// Any of "text", "audio-only", "text-with-audio".
	DisplayMode string `json:"display_mode,nullable"`
	// Any of "worldwide".
	DistributionTerritories []string `json:"distribution_territories,nullable"`
	Ean                     string   `json:"ean,nullable"`
	Edition                 string   `json:"edition,nullable"`
	Fiction                 string   `json:"fiction,nullable"`
	// Any of "Fantasy", "Romance", "Science Fiction", "Mystery and Thriller", "Action
	// and Adventure", "Dystopia", "Business and Economics", "Technology", "Detective
	// and Crime", "Horror", "Biography and Memoir", "Education and Learning",
	// "History", "Children's Literature", "Fairy Tales and Folklore", "Fan Fiction",
	// "General Fiction", "Health and Wellness", "Historical Fiction", "Humor",
	// "Literary Classics", "Philosophy", "Poetry", "Politics and Government",
	// "Psychology", "Science and Nature", "Self-Help", "Spirituality and Religion",
	// "Travel", "True Crime", "Other", "Adult Romance".
	Genre                      []string       `json:"genre,nullable"`
	Isbn                       string         `json:"isbn,nullable"`
	Language                   string         `json:"language,nullable"`
	LastUpdatedFromProjectUnix int64          `json:"last_updated_from_project_unix,nullable"`
	LegalTerms                 ReadLegalTerms `json:"legal_terms,nullable"`
	ListPrice                  float64        `json:"list_price,nullable"`
	MatureContent              bool           `json:"mature_content,nullable"`
	Origin                     string         `json:"origin,nullable"`
	OriginalAudioDocumentID    string         `json:"original_audio_document_id,nullable"`
	OriginalFileType           string         `json:"original_file_type,nullable"`
	// Any of "none", "engagement_based", "fixed_payout".
	PayoutType          string                                                   `json:"payout_type,nullable"`
	PreviewAudioObject  StudioProjectGetResponsePublishingReadPreviewAudioObject `json:"preview_audio_object,nullable"`
	PublicationDate     string                                                   `json:"publication_date,nullable"`
	PublishedAtUnix     int64                                                    `json:"published_at_unix,nullable"`
	Publisher           string                                                   `json:"publisher,nullable"`
	PublisherProfileID  string                                                   `json:"publisher_profile_id,nullable"`
	PublishingProjectID string                                                   `json:"publishing_project_id,nullable"`
	PublishingState     string                                                   `json:"publishing_state"`
	QualityScore        int64                                                    `json:"quality_score,nullable"`
	ReadSlug            string                                                   `json:"read_slug,nullable"`
	Review              StudioProjectGetResponsePublishingReadReview             `json:"review,nullable"`
	SampleConfig        StudioProjectGetResponsePublishingReadSampleConfig       `json:"sample_config,nullable"`
	SeriesID            string                                                   `json:"series_id,nullable"`
	Subtitle            string                                                   `json:"subtitle,nullable"`
	// Any of "children", "young adult", "adult", "all ages".
	TargetAudience string `json:"target_audience,nullable"`
	Title          string `json:"title,nullable"`
	VoiceID        string `json:"voice_id,nullable"`
	Volume         int64  `json:"volume,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Chapters                   respjson.Field
		CharCount                  respjson.Field
		CreatedAtUnix              respjson.Field
		ReadID                     respjson.Field
		UpdatedAtUnix              respjson.Field
		WordCount                  respjson.Field
		ArticleImageURL            respjson.Field
		Author                     respjson.Field
		CanUseAssistant            respjson.Field
		ContentGuidelinesTerms     respjson.Field
		ContentType                respjson.Field
		Contributors               respjson.Field
		Copyright                  respjson.Field
		Currency                   respjson.Field
		Description                respjson.Field
		DisplayMode                respjson.Field
		DistributionTerritories    respjson.Field
		Ean                        respjson.Field
		Edition                    respjson.Field
		Fiction                    respjson.Field
		Genre                      respjson.Field
		Isbn                       respjson.Field
		Language                   respjson.Field
		LastUpdatedFromProjectUnix respjson.Field
		LegalTerms                 respjson.Field
		ListPrice                  respjson.Field
		MatureContent              respjson.Field
		Origin                     respjson.Field
		OriginalAudioDocumentID    respjson.Field
		OriginalFileType           respjson.Field
		PayoutType                 respjson.Field
		PreviewAudioObject         respjson.Field
		PublicationDate            respjson.Field
		PublishedAtUnix            respjson.Field
		Publisher                  respjson.Field
		PublisherProfileID         respjson.Field
		PublishingProjectID        respjson.Field
		PublishingState            respjson.Field
		QualityScore               respjson.Field
		ReadSlug                   respjson.Field
		Review                     respjson.Field
		SampleConfig               respjson.Field
		SeriesID                   respjson.Field
		Subtitle                   respjson.Field
		TargetAudience             respjson.Field
		Title                      respjson.Field
		VoiceID                    respjson.Field
		Volume                     respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingRead) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePublishingReadChapter struct {
	ChapterName        string  `json:"chapter_name,required"`
	CharCount          int64   `json:"char_count,required"`
	StartingCharOffset int64   `json:"starting_char_offset,required"`
	WordCount          int64   `json:"word_count,required"`
	DurationSeconds    float64 `json:"duration_seconds,nullable"`
	FileNumber         string  `json:"file_number,nullable"`
	HasParsedHTML      bool    `json:"has_parsed_html"`
	HasSummary         bool    `json:"has_summary"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChapterName        respjson.Field
		CharCount          respjson.Field
		StartingCharOffset respjson.Field
		WordCount          respjson.Field
		DurationSeconds    respjson.Field
		FileNumber         respjson.Field
		HasParsedHTML      respjson.Field
		HasSummary         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingReadChapter) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingReadChapter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePublishingReadContributor struct {
	Name string `json:"name,required"`
	Role string `json:"role,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingReadContributor) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingReadContributor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePublishingReadPreviewAudioObject struct {
	AudioURL        string `json:"audio_url,required"`
	Text            string `json:"text,required"`
	VoiceID         string `json:"voice_id,required"`
	GeneratedAtUnix int64  `json:"generated_at_unix,nullable"`
	IsAutoGenerated bool   `json:"is_auto_generated,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioURL        respjson.Field
		Text            respjson.Field
		VoiceID         respjson.Field
		GeneratedAtUnix respjson.Field
		IsAutoGenerated respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingReadPreviewAudioObject) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingReadPreviewAudioObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePublishingReadReview struct {
	// Any of "approved", "edits_required", "rejected".
	ReviewStatus   string `json:"review_status,required"`
	ReviewedAtUnix int64  `json:"reviewed_at_unix,required"`
	// Any of "lacks_structure", "doesnt_open", "not_literary_work",
	// "language_not_supported", "too_short", "duplicate", "promotional",
	// "formatting_issues", "low_quality", "metadata_incomplete",
	// "metadata_inaccurate", "typos", "review_error", "spam", "legal_violation",
	// "content_policy", "public_domain", "other".
	RejectReasons   []string         `json:"reject_reasons,nullable"`
	RejectedDetails string           `json:"rejected_details,nullable"`
	ReviewedBy      string           `json:"reviewed_by,nullable"`
	ScoresBreakdown map[string]int64 `json:"scores_breakdown,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReviewStatus    respjson.Field
		ReviewedAtUnix  respjson.Field
		RejectReasons   respjson.Field
		RejectedDetails respjson.Field
		ReviewedBy      respjson.Field
		ScoresBreakdown respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingReadReview) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingReadReview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectGetResponsePublishingReadSampleConfig struct {
	ChapterIDs []string `json:"chapter_ids,nullable"`
	IsSample   bool     `json:"is_sample"`
	ParentID   string   `json:"parent_id,nullable"`
	// Any of "read", "collection".
	ParentType string `json:"parent_type,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChapterIDs  respjson.Field
		IsSample    respjson.Field
		ParentID    respjson.Field
		ParentType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectGetResponsePublishingReadSampleConfig) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectGetResponsePublishingReadSampleConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source type of the project.
type StudioProjectGetResponseSourceType string

const (
	StudioProjectGetResponseSourceTypeBlank      StudioProjectGetResponseSourceType = "blank"
	StudioProjectGetResponseSourceTypeBook       StudioProjectGetResponseSourceType = "book"
	StudioProjectGetResponseSourceTypeArticle    StudioProjectGetResponseSourceType = "article"
	StudioProjectGetResponseSourceTypeGenfm      StudioProjectGetResponseSourceType = "genfm"
	StudioProjectGetResponseSourceTypeVideo      StudioProjectGetResponseSourceType = "video"
	StudioProjectGetResponseSourceTypeScreenplay StudioProjectGetResponseSourceType = "screenplay"
)

// The target audience of the project.
type StudioProjectGetResponseTargetAudience string

const (
	StudioProjectGetResponseTargetAudienceChildren   StudioProjectGetResponseTargetAudience = "children"
	StudioProjectGetResponseTargetAudienceYoungAdult StudioProjectGetResponseTargetAudience = "young adult"
	StudioProjectGetResponseTargetAudienceAdult      StudioProjectGetResponseTargetAudience = "adult"
	StudioProjectGetResponseTargetAudienceAllAges    StudioProjectGetResponseTargetAudience = "all ages"
)

type StudioProjectListResponse struct {
	// A list of projects with their metadata.
	Projects []Project `json:"projects,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Projects    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectDeleteResponse struct {
	// The status of the studio project deletion request. If the request was
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
func (r StudioProjectDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectConvertResponse struct {
	// The status of the studio project conversion request. If the request was
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
func (r StudioProjectConvertResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectConvertResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectNewPronunciationDictionaryResponse struct {
	// The status of the create pronunciation dictionary request. If the request was
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
func (r StudioProjectNewPronunciationDictionaryResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectNewPronunciationDictionaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectNewParams struct {
	// The name of the Studio project, used for identification only.
	Name string `json:"name,required"`
	// An optional name of the author of the Studio project, this will be added as
	// metadata to the mp3 file on Studio project or chapter download.
	Author param.Opt[string] `json:"author,omitzero"`
	// [Alpha Feature] Whether automatically assign voices to phrases in the create
	// Project.
	AutoAssignVoices param.Opt[bool] `json:"auto_assign_voices,omitzero"`
	// A url that will be called by our service when the Studio project is converted. Request will contain a json blob containing the status of the conversion
	//
	//	Messages:
	//	1. When project was converted successfully:
	//	{
	//	  type: "project_conversion_status",
	//	  event_timestamp: 1234567890,
	//	  data: {
	//	    request_id: "1234567890",
	//	    project_id: "21m00Tcm4TlvDq8ikWAM",
	//	    conversion_status: "success",
	//	    project_snapshot_id: "22m00Tcm4TlvDq8ikMAT",
	//	    error_details: None,
	//	  }
	//	}
	//	2. When project conversion failed:
	//	{
	//	  type: "project_conversion_status",
	//	  event_timestamp: 1234567890,
	//	  data: {
	//	    request_id: "1234567890",
	//	    project_id: "21m00Tcm4TlvDq8ikWAM",
	//	    conversion_status: "error",
	//	    project_snapshot_id: None,
	//	    error_details: "Error details if conversion failed"
	//	  }
	//	}
	//
	//	3. When chapter was converted successfully:
	//	{
	//	  type: "chapter_conversion_status",
	//	  event_timestamp: 1234567890,
	//	  data: {
	//	    request_id: "1234567890",
	//	    project_id: "21m00Tcm4TlvDq8ikWAM",
	//	    chapter_id: "22m00Tcm4TlvDq8ikMAT",
	//	    conversion_status: "success",
	//	    chapter_snapshot_id: "23m00Tcm4TlvDq8ikMAV",
	//	    error_details: None,
	//	  }
	//	}
	//	4. When chapter conversion failed:
	//	{
	//	  type: "chapter_conversion_status",
	//	  event_timestamp: 1234567890,
	//	  data: {
	//	    request_id: "1234567890",
	//	    project_id: "21m00Tcm4TlvDq8ikWAM",
	//	    chapter_id: "22m00Tcm4TlvDq8ikMAT",
	//	    conversion_status: "error",
	//	    chapter_snapshot_id: None,
	//	    error_details: "Error details if conversion failed"
	//	  }
	//	}
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// An optional content type of the Studio project.
	ContentType param.Opt[string] `json:"content_type,omitzero"`
	// The ID of the model to be used for this Studio project, you can query GET
	// /v1/models to list all available models.
	DefaultModelID param.Opt[string] `json:"default_model_id,omitzero"`
	// The voice_id that corresponds to the default voice used for new paragraphs.
	DefaultParagraphVoiceID param.Opt[string] `json:"default_paragraph_voice_id,omitzero"`
	// The voice_id that corresponds to the default voice used for new titles.
	DefaultTitleVoiceID param.Opt[string] `json:"default_title_voice_id,omitzero"`
	// An optional description of the Studio project.
	Description param.Opt[string] `json:"description,omitzero"`
	// An optional URL from which we will extract content to initialize the Studio
	// project. If this is set, 'from_url' and 'from_content' must be null. If neither
	// 'from_url', 'from_document', 'from_content' are provided we will initialize the
	// Studio project as blank.
	FromURL param.Opt[string] `json:"from_url,omitzero"`
	// An optional ISBN number of the Studio project you want to create, this will be
	// added as metadata to the mp3 file on Studio project or chapter download.
	IsbnNumber param.Opt[string] `json:"isbn_number,omitzero"`
	// An optional language of the Studio project. Two-letter language code (ISO
	// 639-1).
	Language param.Opt[string] `json:"language,omitzero"`
	// An optional specification of whether this Studio project contains mature
	// content.
	MatureContent param.Opt[bool] `json:"mature_content,omitzero"`
	// An optional original publication date of the Studio project, in the format
	// YYYY-MM-DD or YYYY.
	OriginalPublicationDate param.Opt[string] `json:"original_publication_date,omitzero"`
	// An optional name of the author of the Studio project, this will be added as
	// metadata to the mp3 file on Studio project or chapter download.
	Title param.Opt[string] `json:"title,omitzero"`
	// [Deprecated] When the Studio project is downloaded, should the returned audio
	// have postprocessing in order to make it compliant with audiobook normalized
	// volume requirements
	AcxVolumeNormalization param.Opt[bool] `json:"acx_volume_normalization,omitzero"`
	// Whether to auto convert the Studio project to audio or not.
	AutoConvert param.Opt[bool] `json:"auto_convert,omitzero"`
	// An optional content to initialize the Studio project with. If this is set, 'from_url' and 'from_document' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	//
	//	Example:
	//	[{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]
	FromContentJson param.Opt[string] `json:"from_content_json,omitzero"`
	// When the Studio project is downloaded, should the returned audio have
	// postprocessing in order to make it compliant with audiobook normalized volume
	// requirements
	VolumeNormalization param.Opt[bool] `json:"volume_normalization,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
	//
	//	When set to 'auto', the system will automatically decide whether to apply text normalization
	//	(e.g., spelling out numbers). With 'on', text normalization will always be applied, while
	//	with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
	//
	// Any of "auto", "on", "off", "apply_english".
	ApplyTextNormalization StudioProjectNewParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// An optional specification of whether the content of this Studio project is
	// fiction.
	//
	// Any of "fiction", "non-fiction".
	Fiction StudioProjectNewParamsFiction `json:"fiction,omitzero"`
	// An optional .epub, .pdf, .txt or similar file can be provided. If provided, we
	// will initialize the Studio project with its content. If this is set, 'from_url'
	// and 'from_content' must be null. If neither 'from_url', 'from_document',
	// 'from_content' are provided we will initialize the Studio project as blank.
	FromDocument io.Reader `json:"from_document,omitzero" format:"binary"`
	// The type of Studio project to create.
	//
	// Any of "blank", "book", "article", "genfm", "video", "screenplay".
	SourceType StudioProjectNewParamsSourceType `json:"source_type,omitzero"`
	// An optional target audience of the Studio project.
	//
	// Any of "children", "young adult", "adult", "all ages".
	TargetAudience StudioProjectNewParamsTargetAudience `json:"target_audience,omitzero"`
	// An optional list of genres associated with the Studio project.
	Genres []string `json:"genres,omitzero"`
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id,
	// version_id) encoded as a list of JSON strings for pronunciation dictionaries to
	// be applied to the text. A list of json encoded strings is required as adding
	// projects may occur through formData as opposed to jsonBody. To specify multiple
	// dictionaries use multiple --form lines in your curl, such as --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"'
	// --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []string `json:"pronunciation_dictionary_locators,omitzero"`
	// Output quality of the generated audio. Must be one of: 'standard' - standard
	// output format, 128kbps with 44.1kHz sample rate. 'high' - high quality output
	// format, 192kbps with 44.1kHz sample rate and major improvements on our side.
	// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and
	// highest improvements on our side. 'ultra_lossless' - ultra quality output
	// format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side
	// in a fully lossless format.
	//
	// Any of "standard", "high", "ultra", "ultra_lossless".
	QualityPreset StudioProjectNewParamsQualityPreset `json:"quality_preset,omitzero"`
	// Optional voice settings overrides for the project, encoded as a list of JSON strings.
	//
	//	Example:
	//	["{\"voice_id\": \"21m00Tcm4TlvDq8ikWAM\", \"stability\": 0.7, \"similarity_boost\": 0.8, \"style\": 0.5, \"speed\": 1.0, \"use_speaker_boost\": true}"]
	VoiceSettings []string `json:"voice_settings,omitzero"`
	paramObj
}

func (r StudioProjectNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
//
//	When set to 'auto', the system will automatically decide whether to apply text normalization
//	(e.g., spelling out numbers). With 'on', text normalization will always be applied, while
//	with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
type StudioProjectNewParamsApplyTextNormalization string

const (
	StudioProjectNewParamsApplyTextNormalizationAuto         StudioProjectNewParamsApplyTextNormalization = "auto"
	StudioProjectNewParamsApplyTextNormalizationOn           StudioProjectNewParamsApplyTextNormalization = "on"
	StudioProjectNewParamsApplyTextNormalizationOff          StudioProjectNewParamsApplyTextNormalization = "off"
	StudioProjectNewParamsApplyTextNormalizationApplyEnglish StudioProjectNewParamsApplyTextNormalization = "apply_english"
)

// An optional specification of whether the content of this Studio project is
// fiction.
type StudioProjectNewParamsFiction string

const (
	StudioProjectNewParamsFictionFiction    StudioProjectNewParamsFiction = "fiction"
	StudioProjectNewParamsFictionNonFiction StudioProjectNewParamsFiction = "non-fiction"
)

// Output quality of the generated audio. Must be one of: 'standard' - standard
// output format, 128kbps with 44.1kHz sample rate. 'high' - high quality output
// format, 192kbps with 44.1kHz sample rate and major improvements on our side.
// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and
// highest improvements on our side. 'ultra_lossless' - ultra quality output
// format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side
// in a fully lossless format.
type StudioProjectNewParamsQualityPreset string

const (
	StudioProjectNewParamsQualityPresetStandard      StudioProjectNewParamsQualityPreset = "standard"
	StudioProjectNewParamsQualityPresetHigh          StudioProjectNewParamsQualityPreset = "high"
	StudioProjectNewParamsQualityPresetUltra         StudioProjectNewParamsQualityPreset = "ultra"
	StudioProjectNewParamsQualityPresetUltraLossless StudioProjectNewParamsQualityPreset = "ultra_lossless"
)

// The type of Studio project to create.
type StudioProjectNewParamsSourceType string

const (
	StudioProjectNewParamsSourceTypeBlank      StudioProjectNewParamsSourceType = "blank"
	StudioProjectNewParamsSourceTypeBook       StudioProjectNewParamsSourceType = "book"
	StudioProjectNewParamsSourceTypeArticle    StudioProjectNewParamsSourceType = "article"
	StudioProjectNewParamsSourceTypeGenfm      StudioProjectNewParamsSourceType = "genfm"
	StudioProjectNewParamsSourceTypeVideo      StudioProjectNewParamsSourceType = "video"
	StudioProjectNewParamsSourceTypeScreenplay StudioProjectNewParamsSourceType = "screenplay"
)

// An optional target audience of the Studio project.
type StudioProjectNewParamsTargetAudience string

const (
	StudioProjectNewParamsTargetAudienceChildren   StudioProjectNewParamsTargetAudience = "children"
	StudioProjectNewParamsTargetAudienceYoungAdult StudioProjectNewParamsTargetAudience = "young adult"
	StudioProjectNewParamsTargetAudienceAdult      StudioProjectNewParamsTargetAudience = "adult"
	StudioProjectNewParamsTargetAudienceAllAges    StudioProjectNewParamsTargetAudience = "all ages"
)

type StudioProjectGetParams struct {
	// The share ID of the project
	ShareID param.Opt[string] `query:"share_id,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [StudioProjectGetParams]'s query parameters as `url.Values`.
func (r StudioProjectGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type StudioProjectUpdateParams struct {
	// The voice_id that corresponds to the default voice used for new paragraphs.
	DefaultParagraphVoiceID string `json:"default_paragraph_voice_id,required"`
	// The voice_id that corresponds to the default voice used for new titles.
	DefaultTitleVoiceID string `json:"default_title_voice_id,required"`
	// The name of the Studio project, used for identification only.
	Name string `json:"name,required"`
	// An optional name of the author of the Studio project, this will be added as
	// metadata to the mp3 file on Studio project or chapter download.
	Author param.Opt[string] `json:"author,omitzero"`
	// An optional ISBN number of the Studio project you want to create, this will be
	// added as metadata to the mp3 file on Studio project or chapter download.
	IsbnNumber param.Opt[string] `json:"isbn_number,omitzero"`
	// An optional name of the author of the Studio project, this will be added as
	// metadata to the mp3 file on Studio project or chapter download.
	Title param.Opt[string] `json:"title,omitzero"`
	// When the Studio project is downloaded, should the returned audio have
	// postprocessing in order to make it compliant with audiobook normalized volume
	// requirements
	VolumeNormalization param.Opt[bool] `json:"volume_normalization,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r StudioProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectConvertParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectNewPronunciationDictionaryParams struct {
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id,
	// version_id) encoded as a list of JSON strings for pronunciation dictionaries to
	// be applied to the text. A list of json encoded strings is required as adding
	// projects may occur through formData as opposed to jsonBody. To specify multiple
	// dictionaries use multiple --form lines in your curl, such as --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"'
	// --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator `json:"pronunciation_dictionary_locators,omitzero,required"`
	// This will automatically mark text in this project for reconversion when the new
	// dictionary applies or the old one no longer does.
	InvalidateAffectedText param.Opt[bool] `json:"invalidate_affected_text,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r StudioProjectNewPronunciationDictionaryParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectNewPronunciationDictionaryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectNewPronunciationDictionaryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties PronunciationDictionaryID, VersionID are required.
type StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator struct {
	VersionID                 param.Opt[string] `json:"version_id,omitzero,required"`
	PronunciationDictionaryID string            `json:"pronunciation_dictionary_id,required"`
	paramObj
}

func (r StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectUpdateContentParams struct {
	// An optional URL from which we will extract content to initialize the Studio
	// project. If this is set, 'from_url' and 'from_content' must be null. If neither
	// 'from_url', 'from_document', 'from_content' are provided we will initialize the
	// Studio project as blank.
	FromURL param.Opt[string] `json:"from_url,omitzero"`
	// Whether to auto convert the Studio project to audio or not.
	AutoConvert param.Opt[bool] `json:"auto_convert,omitzero"`
	// An optional content to initialize the Studio project with. If this is set, 'from_url' and 'from_document' must be null. If neither 'from_url', 'from_document', 'from_content' are provided we will initialize the Studio project as blank.
	//
	//	Example:
	//	[{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]
	FromContentJson param.Opt[string] `json:"from_content_json,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// An optional .epub, .pdf, .txt or similar file can be provided. If provided, we
	// will initialize the Studio project with its content. If this is set, 'from_url'
	// and 'from_content' must be null. If neither 'from_url', 'from_document',
	// 'from_content' are provided we will initialize the Studio project as blank.
	FromDocument io.Reader `json:"from_document,omitzero" format:"binary"`
	paramObj
}

func (r StudioProjectUpdateContentParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
