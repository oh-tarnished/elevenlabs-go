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

// StudioProjectChapterSnapshotService contains methods and other services that
// help with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStudioProjectChapterSnapshotService] method instead.
type StudioProjectChapterSnapshotService struct {
	Options []option.RequestOption
}

// NewStudioProjectChapterSnapshotService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewStudioProjectChapterSnapshotService(opts ...option.RequestOption) (r StudioProjectChapterSnapshotService) {
	r = StudioProjectChapterSnapshotService{}
	r.Options = opts
	return
}

// Returns the chapter snapshot.
func (r *StudioProjectChapterSnapshotService) Get(ctx context.Context, chapterSnapshotID string, params StudioProjectChapterSnapshotGetParams, opts ...option.RequestOption) (res *StudioProjectChapterSnapshotGetResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if params.ChapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	if chapterSnapshotID == "" {
		err = errors.New("missing required chapter_snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s/snapshots/%s", params.ProjectID, params.ChapterID, chapterSnapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets information about all the snapshots of a chapter. Each snapshot can be
// downloaded as audio. Whenever a chapter is converted a snapshot will
// automatically be created.
func (r *StudioProjectChapterSnapshotService) List(ctx context.Context, chapterID string, params StudioProjectChapterSnapshotListParams, opts ...option.RequestOption) (res *StudioProjectChapterSnapshotListResponse, err error) {
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
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s/snapshots", params.ProjectID, chapterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Stream the audio from a chapter snapshot. Use
// `GET /v1/studio/projects/{project_id}/chapters/{chapter_id}/snapshots` to return
// the snapshots of a chapter.
func (r *StudioProjectChapterSnapshotService) StreamAudio(ctx context.Context, chapterSnapshotID string, params StudioProjectChapterSnapshotStreamAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if params.ChapterID == "" {
		err = errors.New("missing required chapter_id parameter")
		return
	}
	if chapterSnapshotID == "" {
		err = errors.New("missing required chapter_snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/chapters/%s/snapshots/%s/stream", params.ProjectID, params.ChapterID, chapterSnapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type StudioProjectChapterSnapshotGetResponse struct {
	// The ID of the chapter.
	ChapterID string `json:"chapter_id,required"`
	// The ID of the chapter snapshot.
	ChapterSnapshotID   string                       `json:"chapter_snapshot_id,required"`
	CharacterAlignments []CharacterAlignmentSnapshot `json:"character_alignments,required"`
	// The creation date of the chapter snapshot.
	CreatedAtUnix int64 `json:"created_at_unix,required"`
	// The name of the chapter snapshot.
	Name string `json:"name,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChapterID           respjson.Field
		ChapterSnapshotID   respjson.Field
		CharacterAlignments respjson.Field
		CreatedAtUnix       respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterSnapshotGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterSnapshotGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterSnapshotListResponse struct {
	// List of chapter snapshots.
	Snapshots []StudioProjectChapterSnapshotListResponseSnapshot `json:"snapshots,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Snapshots   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterSnapshotListResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterSnapshotListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterSnapshotListResponseSnapshot struct {
	// The ID of the chapter.
	ChapterID string `json:"chapter_id,required"`
	// The ID of the chapter snapshot.
	ChapterSnapshotID string `json:"chapter_snapshot_id,required"`
	// The creation date of the chapter snapshot.
	CreatedAtUnix int64 `json:"created_at_unix,required"`
	// The name of the chapter snapshot.
	Name string `json:"name,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChapterID         respjson.Field
		ChapterSnapshotID respjson.Field
		CreatedAtUnix     respjson.Field
		Name              respjson.Field
		ProjectID         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectChapterSnapshotListResponseSnapshot) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectChapterSnapshotListResponseSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectChapterSnapshotGetParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// The ID of the chapter.
	ChapterID string `path:"chapter_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectChapterSnapshotListParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectChapterSnapshotStreamAudioParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// The ID of the chapter.
	ChapterID string `path:"chapter_id,required" json:"-"`
	// Whether to convert the audio to mpeg format.
	ConvertToMpeg param.Opt[bool] `json:"convert_to_mpeg,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r StudioProjectChapterSnapshotStreamAudioParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectChapterSnapshotStreamAudioParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectChapterSnapshotStreamAudioParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
