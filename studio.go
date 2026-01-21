// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
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

// StudioService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStudioService] method instead.
type StudioService struct {
	Options  []option.RequestOption
	Projects StudioProjectService
}

// NewStudioService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStudioService(opts ...option.RequestOption) (r StudioService) {
	r = StudioService{}
	r.Options = opts
	r.Projects = NewStudioProjectService(opts...)
	return
}

// Create and auto-convert a podcast project. Currently, the LLM cost is covered by
// us but you will still be charged for the audio generation. In the future, you
// will be charged for both the LLM and audio generation costs.
func (r *StudioService) NewPodcast(ctx context.Context, params StudioNewPodcastParams, opts ...option.RequestOption) (res *StudioNewPodcastResponse, err error) {
	if !param.IsOmitted(params.SafetyIdentifier) {
		opts = append(opts, option.WithHeader("safety-identifier", fmt.Sprintf("%s", params.SafetyIdentifier.Value)))
	}
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/studio/podcasts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// The properties Text, Type are required.
type PodcastTextSourceParam struct {
	// The text to create the podcast from.
	Text string `json:"text,required"`
	// The type of source to create.
	//
	// This field can be elided, and will marshal its zero value as "text".
	Type constant.Text `json:"type,required"`
	paramObj
}

func (r PodcastTextSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow PodcastTextSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodcastTextSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Type, URL are required.
type PodcastURLSourceParam struct {
	// The URL to create the podcast from.
	URL string `json:"url,required"`
	// The type of source to create.
	//
	// This field can be elided, and will marshal its zero value as "url".
	Type constant.URL `json:"type,required"`
	paramObj
}

func (r PodcastURLSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow PodcastURLSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PodcastURLSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Project struct {
	// The access level of the project.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	AccessLevel ProjectAccessLevel `json:"access_level,required"`
	// Whether the project can be downloaded.
	CanBeDownloaded bool `json:"can_be_downloaded,required"`
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
	// Whether quality check is enabled for this project.
	//
	// Deprecated: deprecated
	QualityCheckOn bool `json:"quality_check_on,required"`
	// Whether quality check is enabled on the project when bulk converting.
	//
	// Deprecated: deprecated
	QualityCheckOnWhenBulkConvert bool `json:"quality_check_on_when_bulk_convert,required"`
	// The state of the project.
	//
	// Any of "creating", "default", "converting", "in_queue".
	State ProjectState `json:"state,required"`
	// Whether the project uses volume normalization.
	VolumeNormalization bool `json:"volume_normalization,required"`
	// The aspect ratio of the project.
	//
	// Any of "16:9", "9:16", "4:5", "1:1".
	AspectRatio ProjectAspectRatio `json:"aspect_ratio,nullable"`
	// The author of the project.
	Author string `json:"author,nullable"`
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
	// Whether the project is fiction.
	//
	// Any of "fiction", "non-fiction".
	Fiction ProjectFiction `json:"fiction,nullable"`
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
	// The source type of the project.
	//
	// Any of "blank", "book", "article", "genfm", "video", "screenplay".
	SourceType ProjectSourceType `json:"source_type,nullable"`
	// The target audience of the project.
	//
	// Any of "children", "young adult", "adult", "all ages".
	TargetAudience ProjectTargetAudience `json:"target_audience,nullable"`
	// The title of the project.
	Title string `json:"title,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessLevel                   respjson.Field
		CanBeDownloaded               respjson.Field
		CreateDateUnix                respjson.Field
		CreatedByUserID               respjson.Field
		DefaultModelID                respjson.Field
		DefaultParagraphVoiceID       respjson.Field
		DefaultTitleVoiceID           respjson.Field
		Name                          respjson.Field
		ProjectID                     respjson.Field
		QualityCheckOn                respjson.Field
		QualityCheckOnWhenBulkConvert respjson.Field
		State                         respjson.Field
		VolumeNormalization           respjson.Field
		AspectRatio                   respjson.Field
		Author                        respjson.Field
		CaptionStyle                  respjson.Field
		CaptionsEnabled               respjson.Field
		ChaptersEnabled               respjson.Field
		ContentType                   respjson.Field
		CoverImageURL                 respjson.Field
		CreationMeta                  respjson.Field
		Description                   respjson.Field
		Fiction                       respjson.Field
		Genres                        respjson.Field
		IsbnNumber                    respjson.Field
		Language                      respjson.Field
		LastConversionDateUnix        respjson.Field
		MatureContent                 respjson.Field
		OriginalPublicationDate       respjson.Field
		PublicShareID                 respjson.Field
		SourceType                    respjson.Field
		TargetAudience                respjson.Field
		Title                         respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The access level of the project.
type ProjectAccessLevel string

const (
	ProjectAccessLevelAdmin     ProjectAccessLevel = "admin"
	ProjectAccessLevelEditor    ProjectAccessLevel = "editor"
	ProjectAccessLevelCommenter ProjectAccessLevel = "commenter"
	ProjectAccessLevelViewer    ProjectAccessLevel = "viewer"
)

// The state of the project.
type ProjectState string

const (
	ProjectStateCreating   ProjectState = "creating"
	ProjectStateDefault    ProjectState = "default"
	ProjectStateConverting ProjectState = "converting"
	ProjectStateInQueue    ProjectState = "in_queue"
)

// The aspect ratio of the project.
type ProjectAspectRatio string

const (
	ProjectAspectRatio16_9 ProjectAspectRatio = "16:9"
	ProjectAspectRatio9_16 ProjectAspectRatio = "9:16"
	ProjectAspectRatio4_5  ProjectAspectRatio = "4:5"
	ProjectAspectRatio1_1  ProjectAspectRatio = "1:1"
)

// Whether the project is fiction.
type ProjectFiction string

const (
	ProjectFictionFiction    ProjectFiction = "fiction"
	ProjectFictionNonFiction ProjectFiction = "non-fiction"
)

// The source type of the project.
type ProjectSourceType string

const (
	ProjectSourceTypeBlank      ProjectSourceType = "blank"
	ProjectSourceTypeBook       ProjectSourceType = "book"
	ProjectSourceTypeArticle    ProjectSourceType = "article"
	ProjectSourceTypeGenfm      ProjectSourceType = "genfm"
	ProjectSourceTypeVideo      ProjectSourceType = "video"
	ProjectSourceTypeScreenplay ProjectSourceType = "screenplay"
)

// The target audience of the project.
type ProjectTargetAudience string

const (
	ProjectTargetAudienceChildren   ProjectTargetAudience = "children"
	ProjectTargetAudienceYoungAdult ProjectTargetAudience = "young adult"
	ProjectTargetAudienceAdult      ProjectTargetAudience = "adult"
	ProjectTargetAudienceAllAges    ProjectTargetAudience = "all ages"
)

type StudioNewPodcastResponse struct {
	// The project associated with the created podcast.
	Project Project `json:"project,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Project     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioNewPodcastResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioNewPodcastResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioNewPodcastParams struct {
	// The type of podcast to generate. Can be 'conversation', an interaction between
	// two voices, or 'bulletin', a monologue.
	Mode StudioNewPodcastParamsModeUnion `json:"mode,omitzero,required"`
	// The ID of the model to be used for this Studio project, you can query GET
	// /v1/models to list all available models.
	ModelID string `json:"model_id,required"`
	// The source content for the Podcast.
	Source StudioNewPodcastParamsSourceUnion `json:"source,omitzero,required"`
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
	// Additional instructions prompt for the podcast generation used to adjust the
	// podcast's style and tone.
	InstructionsPrompt param.Opt[string] `json:"instructions_prompt,omitzero"`
	// The intro text that will always be added to the beginning of the podcast.
	Intro param.Opt[string] `json:"intro,omitzero"`
	// An optional language of the Studio project. Two-letter language code (ISO
	// 639-1).
	Language param.Opt[string] `json:"language,omitzero"`
	// The outro text that will always be added to the end of the podcast.
	Outro param.Opt[string] `json:"outro,omitzero"`
	// Used for moderation. Your workspace must be allowlisted to use this feature.
	SafetyIdentifier param.Opt[string] `header:"safety-identifier,omitzero" json:"-"`
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
	ApplyTextNormalization StudioNewPodcastParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// A brief summary or highlights of the Studio project's content, providing key
	// points or themes. This should be between 10 and 70 characters.
	Highlights []string `json:"highlights,omitzero"`
	// Duration of the generated podcast. Must be one of: short - produces podcasts
	// shorter than 3 minutes. default - produces podcasts roughly between 3-7 minutes.
	// long - produces podcasts longer than 7 minutes.
	//
	// Any of "short", "default", "long".
	DurationScale StudioNewPodcastParamsDurationScale `json:"duration_scale,omitzero"`
	// Output quality of the generated audio. Must be one of: 'standard' - standard
	// output format, 128kbps with 44.1kHz sample rate. 'high' - high quality output
	// format, 192kbps with 44.1kHz sample rate and major improvements on our side.
	// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and
	// highest improvements on our side. 'ultra_lossless' - ultra quality output
	// format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side
	// in a fully lossless format.
	//
	// Any of "standard", "high", "highest", "ultra", "ultra_lossless".
	QualityPreset StudioNewPodcastParamsQualityPreset `json:"quality_preset,omitzero"`
	paramObj
}

func (r StudioNewPodcastParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioNewPodcastParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioNewPodcastParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StudioNewPodcastParamsModeUnion struct {
	OfPodcastConversationMode *StudioNewPodcastParamsModePodcastConversationMode `json:",omitzero,inline"`
	OfPodcastBulletinMode     *StudioNewPodcastParamsModePodcastBulletinMode     `json:",omitzero,inline"`
	paramUnion
}

func (u StudioNewPodcastParamsModeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPodcastConversationMode, u.OfPodcastBulletinMode)
}
func (u *StudioNewPodcastParamsModeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *StudioNewPodcastParamsModeUnion) asAny() any {
	if !param.IsOmitted(u.OfPodcastConversationMode) {
		return u.OfPodcastConversationMode
	} else if !param.IsOmitted(u.OfPodcastBulletinMode) {
		return u.OfPodcastBulletinMode
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsModeUnion) GetConversation() *StudioNewPodcastParamsModePodcastConversationModeConversation {
	if vt := u.OfPodcastConversationMode; vt != nil {
		return &vt.Conversation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsModeUnion) GetBulletin() *StudioNewPodcastParamsModePodcastBulletinModeBulletin {
	if vt := u.OfPodcastBulletinMode; vt != nil {
		return &vt.Bulletin
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsModeUnion) GetType() *string {
	if vt := u.OfPodcastConversationMode; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPodcastBulletinMode; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// The properties Conversation, Type are required.
type StudioNewPodcastParamsModePodcastConversationMode struct {
	// The voice settings for the conversation.
	Conversation StudioNewPodcastParamsModePodcastConversationModeConversation `json:"conversation,omitzero,required"`
	// The type of podcast to create.
	//
	// This field can be elided, and will marshal its zero value as "conversation".
	Type constant.Conversation `json:"type,required"`
	paramObj
}

func (r StudioNewPodcastParamsModePodcastConversationMode) MarshalJSON() (data []byte, err error) {
	type shadow StudioNewPodcastParamsModePodcastConversationMode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioNewPodcastParamsModePodcastConversationMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The voice settings for the conversation.
//
// The properties GuestVoiceID, HostVoiceID are required.
type StudioNewPodcastParamsModePodcastConversationModeConversation struct {
	// The ID of the guest voice.
	GuestVoiceID string `json:"guest_voice_id,required"`
	// The ID of the host voice.
	HostVoiceID string `json:"host_voice_id,required"`
	paramObj
}

func (r StudioNewPodcastParamsModePodcastConversationModeConversation) MarshalJSON() (data []byte, err error) {
	type shadow StudioNewPodcastParamsModePodcastConversationModeConversation
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioNewPodcastParamsModePodcastConversationModeConversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Bulletin, Type are required.
type StudioNewPodcastParamsModePodcastBulletinMode struct {
	// The voice settings for the bulletin.
	Bulletin StudioNewPodcastParamsModePodcastBulletinModeBulletin `json:"bulletin,omitzero,required"`
	// The type of podcast to create.
	//
	// This field can be elided, and will marshal its zero value as "bulletin".
	Type constant.Bulletin `json:"type,required"`
	paramObj
}

func (r StudioNewPodcastParamsModePodcastBulletinMode) MarshalJSON() (data []byte, err error) {
	type shadow StudioNewPodcastParamsModePodcastBulletinMode
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioNewPodcastParamsModePodcastBulletinMode) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The voice settings for the bulletin.
//
// The property HostVoiceID is required.
type StudioNewPodcastParamsModePodcastBulletinModeBulletin struct {
	// The ID of the host voice.
	HostVoiceID string `json:"host_voice_id,required"`
	paramObj
}

func (r StudioNewPodcastParamsModePodcastBulletinModeBulletin) MarshalJSON() (data []byte, err error) {
	type shadow StudioNewPodcastParamsModePodcastBulletinModeBulletin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioNewPodcastParamsModePodcastBulletinModeBulletin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StudioNewPodcastParamsSourceUnion struct {
	OfPodcastTextSource            *PodcastTextSourceParam                      `json:",omitzero,inline"`
	OfPodcastURLSource             *PodcastURLSourceParam                       `json:",omitzero,inline"`
	OfStudioNewPodcastsSourceArray []StudioNewPodcastParamsSourceArrayItemUnion `json:",omitzero,inline"`
	paramUnion
}

func (u StudioNewPodcastParamsSourceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPodcastTextSource, u.OfPodcastURLSource, u.OfStudioNewPodcastsSourceArray)
}
func (u *StudioNewPodcastParamsSourceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *StudioNewPodcastParamsSourceUnion) asAny() any {
	if !param.IsOmitted(u.OfPodcastTextSource) {
		return u.OfPodcastTextSource
	} else if !param.IsOmitted(u.OfPodcastURLSource) {
		return u.OfPodcastURLSource
	} else if !param.IsOmitted(u.OfStudioNewPodcastsSourceArray) {
		return &u.OfStudioNewPodcastsSourceArray
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceUnion) GetText() *string {
	if vt := u.OfPodcastTextSource; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceUnion) GetURL() *string {
	if vt := u.OfPodcastURLSource; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceUnion) GetType() *string {
	if vt := u.OfPodcastTextSource; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPodcastURLSource; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type StudioNewPodcastParamsSourceArrayItemUnion struct {
	OfPodcastTextSource *PodcastTextSourceParam `json:",omitzero,inline"`
	OfPodcastURLSource  *PodcastURLSourceParam  `json:",omitzero,inline"`
	paramUnion
}

func (u StudioNewPodcastParamsSourceArrayItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPodcastTextSource, u.OfPodcastURLSource)
}
func (u *StudioNewPodcastParamsSourceArrayItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *StudioNewPodcastParamsSourceArrayItemUnion) asAny() any {
	if !param.IsOmitted(u.OfPodcastTextSource) {
		return u.OfPodcastTextSource
	} else if !param.IsOmitted(u.OfPodcastURLSource) {
		return u.OfPodcastURLSource
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceArrayItemUnion) GetText() *string {
	if vt := u.OfPodcastTextSource; vt != nil {
		return &vt.Text
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceArrayItemUnion) GetURL() *string {
	if vt := u.OfPodcastURLSource; vt != nil {
		return &vt.URL
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u StudioNewPodcastParamsSourceArrayItemUnion) GetType() *string {
	if vt := u.OfPodcastTextSource; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPodcastURLSource; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
//
//	When set to 'auto', the system will automatically decide whether to apply text normalization
//	(e.g., spelling out numbers). With 'on', text normalization will always be applied, while
//	with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
type StudioNewPodcastParamsApplyTextNormalization string

const (
	StudioNewPodcastParamsApplyTextNormalizationAuto         StudioNewPodcastParamsApplyTextNormalization = "auto"
	StudioNewPodcastParamsApplyTextNormalizationOn           StudioNewPodcastParamsApplyTextNormalization = "on"
	StudioNewPodcastParamsApplyTextNormalizationOff          StudioNewPodcastParamsApplyTextNormalization = "off"
	StudioNewPodcastParamsApplyTextNormalizationApplyEnglish StudioNewPodcastParamsApplyTextNormalization = "apply_english"
)

// Duration of the generated podcast. Must be one of: short - produces podcasts
// shorter than 3 minutes. default - produces podcasts roughly between 3-7 minutes.
// long - produces podcasts longer than 7 minutes.
type StudioNewPodcastParamsDurationScale string

const (
	StudioNewPodcastParamsDurationScaleShort   StudioNewPodcastParamsDurationScale = "short"
	StudioNewPodcastParamsDurationScaleDefault StudioNewPodcastParamsDurationScale = "default"
	StudioNewPodcastParamsDurationScaleLong    StudioNewPodcastParamsDurationScale = "long"
)

// Output quality of the generated audio. Must be one of: 'standard' - standard
// output format, 128kbps with 44.1kHz sample rate. 'high' - high quality output
// format, 192kbps with 44.1kHz sample rate and major improvements on our side.
// 'ultra' - ultra quality output format, 192kbps with 44.1kHz sample rate and
// highest improvements on our side. 'ultra_lossless' - ultra quality output
// format, 705.6kbps with 44.1kHz sample rate and highest improvements on our side
// in a fully lossless format.
type StudioNewPodcastParamsQualityPreset string

const (
	StudioNewPodcastParamsQualityPresetStandard      StudioNewPodcastParamsQualityPreset = "standard"
	StudioNewPodcastParamsQualityPresetHigh          StudioNewPodcastParamsQualityPreset = "high"
	StudioNewPodcastParamsQualityPresetHighest       StudioNewPodcastParamsQualityPreset = "highest"
	StudioNewPodcastParamsQualityPresetUltra         StudioNewPodcastParamsQualityPreset = "ultra"
	StudioNewPodcastParamsQualityPresetUltraLossless StudioNewPodcastParamsQualityPreset = "ultra_lossless"
)
