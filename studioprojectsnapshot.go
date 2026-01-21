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

// StudioProjectSnapshotService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStudioProjectSnapshotService] method instead.
type StudioProjectSnapshotService struct {
	Options []option.RequestOption
}

// NewStudioProjectSnapshotService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStudioProjectSnapshotService(opts ...option.RequestOption) (r StudioProjectSnapshotService) {
	r = StudioProjectSnapshotService{}
	r.Options = opts
	return
}

// Returns the project snapshot.
func (r *StudioProjectSnapshotService) Get(ctx context.Context, projectSnapshotID string, params StudioProjectSnapshotGetParams, opts ...option.RequestOption) (res *StudioProjectSnapshotGetResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if projectSnapshotID == "" {
		err = errors.New("missing required project_snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/snapshots/%s", params.ProjectID, projectSnapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieves a list of snapshots for a Studio project.
func (r *StudioProjectSnapshotService) List(ctx context.Context, projectID string, query StudioProjectSnapshotListParams, opts ...option.RequestOption) (res *StudioProjectSnapshotListResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/snapshots", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a compressed archive of the Studio project's audio.
func (r *StudioProjectSnapshotService) ArchiveAudio(ctx context.Context, projectSnapshotID string, params StudioProjectSnapshotArchiveAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/x-zip")}, opts...)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if projectSnapshotID == "" {
		err = errors.New("missing required project_snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/snapshots/%s/archive", params.ProjectID, projectSnapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Stream the audio from a Studio project snapshot.
func (r *StudioProjectSnapshotService) StreamAudio(ctx context.Context, projectSnapshotID string, params StudioProjectSnapshotStreamAudioParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ProjectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	if projectSnapshotID == "" {
		err = errors.New("missing required project_snapshot_id parameter")
		return
	}
	path := fmt.Sprintf("v1/studio/projects/%s/snapshots/%s/stream", params.ProjectID, projectSnapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

type CharacterAlignmentSnapshot struct {
	CharacterEndTimesSeconds   []float64 `json:"character_end_times_seconds,required"`
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds,required"`
	Characters                 []string  `json:"characters,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterEndTimesSeconds   respjson.Field
		CharacterStartTimesSeconds respjson.Field
		Characters                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CharacterAlignmentSnapshot) RawJSON() string { return r.JSON.raw }
func (r *CharacterAlignmentSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectSnapshotGetResponse struct {
	// The total duration of the audio in seconds.
	AudioDurationSecs   float64                      `json:"audio_duration_secs,required"`
	CharacterAlignments []CharacterAlignmentSnapshot `json:"character_alignments,required"`
	// The creation date of the project snapshot.
	CreatedAtUnix int64 `json:"created_at_unix,required"`
	// The name of the project snapshot.
	Name string `json:"name,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// The ID of the project snapshot.
	ProjectSnapshotID string `json:"project_snapshot_id,required"`
	// (Deprecated)
	AudioUpload any `json:"audio_upload,nullable"`
	// (Deprecated)
	ZipUpload any `json:"zip_upload,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioDurationSecs   respjson.Field
		CharacterAlignments respjson.Field
		CreatedAtUnix       respjson.Field
		Name                respjson.Field
		ProjectID           respjson.Field
		ProjectSnapshotID   respjson.Field
		AudioUpload         respjson.Field
		ZipUpload           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectSnapshotGetResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectSnapshotGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectSnapshotListResponse struct {
	// List of project snapshots.
	Snapshots []StudioProjectSnapshotListResponseSnapshot `json:"snapshots,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Snapshots   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectSnapshotListResponse) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectSnapshotListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectSnapshotListResponseSnapshot struct {
	// The creation date of the project snapshot.
	CreatedAtUnix int64 `json:"created_at_unix,required"`
	// The name of the project snapshot.
	Name string `json:"name,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// The ID of the project snapshot.
	ProjectSnapshotID string `json:"project_snapshot_id,required"`
	// (Deprecated)
	AudioUpload any `json:"audio_upload,nullable"`
	// (Deprecated)
	ZipUpload any `json:"zip_upload,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAtUnix     respjson.Field
		Name              respjson.Field
		ProjectID         respjson.Field
		ProjectSnapshotID respjson.Field
		AudioUpload       respjson.Field
		ZipUpload         respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r StudioProjectSnapshotListResponseSnapshot) RawJSON() string { return r.JSON.raw }
func (r *StudioProjectSnapshotListResponseSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type StudioProjectSnapshotGetParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectSnapshotListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectSnapshotArchiveAudioParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type StudioProjectSnapshotStreamAudioParams struct {
	// The ID of the Studio project.
	ProjectID string `path:"project_id,required" json:"-"`
	// Whether to convert the audio to mpeg format.
	ConvertToMpeg param.Opt[bool] `json:"convert_to_mpeg,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r StudioProjectSnapshotStreamAudioParams) MarshalJSON() (data []byte, err error) {
	type shadow StudioProjectSnapshotStreamAudioParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *StudioProjectSnapshotStreamAudioParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
