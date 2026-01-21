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

// DubbingResourceSegmentService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDubbingResourceSegmentService] method instead.
type DubbingResourceSegmentService struct {
	Options []option.RequestOption
}

// NewDubbingResourceSegmentService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDubbingResourceSegmentService(opts ...option.RequestOption) (r DubbingResourceSegmentService) {
	r = DubbingResourceSegmentService{}
	r.Options = opts
	return
}

// Modifies a single segment with new text and/or start/end times. Will update the
// values for only a specific language of a segment. Does not automatically
// regenerate the dub.
func (r *DubbingResourceSegmentService) Update(ctx context.Context, language string, params DubbingResourceSegmentUpdateParams, opts ...option.RequestOption) (res *DubbingResourceSegmentUpdateResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if params.SegmentID == "" {
		err = errors.New("missing required segment_id parameter")
		return
	}
	if language == "" {
		err = errors.New("missing required language parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/segment/%s/%s", params.DubbingID, params.SegmentID, language)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Deletes a single segment from the dubbing.
func (r *DubbingResourceSegmentService) Delete(ctx context.Context, segmentID string, params DubbingResourceSegmentDeleteParams, opts ...option.RequestOption) (res *DubbingResourceSegmentDeleteResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DubbingID == "" {
		err = errors.New("missing required dubbing_id parameter")
		return
	}
	if segmentID == "" {
		err = errors.New("missing required segment_id parameter")
		return
	}
	path := fmt.Sprintf("v1/dubbing/resource/%s/segment/%s", params.DubbingID, segmentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type DubbingResourceSegmentUpdateResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSegmentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSegmentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSegmentDeleteResponse struct {
	Version int64 `json:"version,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DubbingResourceSegmentDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DubbingResourceSegmentDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSegmentUpdateParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// ID of the segment
	SegmentID string             `path:"segment_id,required" json:"-"`
	EndTime   param.Opt[float64] `json:"end_time,omitzero"`
	StartTime param.Opt[float64] `json:"start_time,omitzero"`
	Text      param.Opt[string]  `json:"text,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r DubbingResourceSegmentUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DubbingResourceSegmentUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DubbingResourceSegmentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DubbingResourceSegmentDeleteParams struct {
	// ID of the dubbing project.
	DubbingID string `path:"dubbing_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
