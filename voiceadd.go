// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// VoiceAddService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceAddService] method instead.
type VoiceAddService struct {
	Options []option.RequestOption
}

// NewVoiceAddService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceAddService(opts ...option.RequestOption) (r VoiceAddService) {
	r = VoiceAddService{}
	r.Options = opts
	return
}

// Add a new voice to your collection of voices in VoiceLab.
func (r *VoiceAddService) New(ctx context.Context, params VoiceAddNewParams, opts ...option.RequestOption) (res *VoiceAddNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/voices/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Add a shared voice to your collection of voices.
func (r *VoiceAddService) AddShared(ctx context.Context, voiceID string, params VoiceAddAddSharedParams, opts ...option.RequestOption) (res *AddVoiceResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.PublicUserID == "" {
		err = errors.New("missing required public_user_id parameter")
		return
	}
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/add/%s/%s", params.PublicUserID, voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AddVoiceResponse struct {
	// The ID of the voice.
	VoiceID string `json:"voice_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VoiceID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddVoiceResponse) RawJSON() string { return r.JSON.raw }
func (r *AddVoiceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAddNewResponse struct {
	// Whether the voice requires verification
	RequiresVerification bool `json:"requires_verification,required"`
	// The ID of the newly created voice.
	VoiceID string `json:"voice_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequiresVerification respjson.Field
		VoiceID              respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceAddNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VoiceAddNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceAddNewParams struct {
	// A list of file paths to audio recordings intended for voice cloning.
	Files []io.Reader `json:"files,omitzero,required" format:"binary"`
	// The name that identifies this voice. This will be displayed in the dropdown of
	// the website.
	Name string `json:"name,required"`
	// A description of the voice.
	Description param.Opt[string] `json:"description,omitzero"`
	// If set will remove background noise for voice samples using our audio isolation
	// model. If the samples do not include background noise, it can make the quality
	// worse.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels VoiceAddNewParamsLabelsUnion `json:"labels,omitzero"`
	paramObj
}

func (r VoiceAddNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type VoiceAddNewParamsLabelsUnion struct {
	OfStringMap map[string]string `json:",omitzero,inline"`
	OfString    param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u VoiceAddNewParamsLabelsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfStringMap, u.OfString)
}
func (u *VoiceAddNewParamsLabelsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *VoiceAddNewParamsLabelsUnion) asAny() any {
	if !param.IsOmitted(u.OfStringMap) {
		return &u.OfStringMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type VoiceAddAddSharedParams struct {
	// Public user ID used to publicly identify ElevenLabs users.
	PublicUserID string `path:"public_user_id,required" json:"-"`
	// The name that identifies this voice. This will be displayed in the dropdown of
	// the website.
	NewName string `json:"new_name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoiceAddAddSharedParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceAddAddSharedParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceAddAddSharedParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
