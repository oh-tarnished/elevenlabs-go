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

// VoicePvcService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoicePvcService] method instead.
type VoicePvcService struct {
	Options []option.RequestOption
	Samples VoicePvcSampleService
	Captcha VoicePvcCaptchaService
}

// NewVoicePvcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoicePvcService(opts ...option.RequestOption) (r VoicePvcService) {
	r = VoicePvcService{}
	r.Options = opts
	r.Samples = NewVoicePvcSampleService(opts...)
	r.Captcha = NewVoicePvcCaptchaService(opts...)
	return
}

// Creates a new PVC voice with metadata but no samples
func (r *VoicePvcService) New(ctx context.Context, params VoicePvcNewParams, opts ...option.RequestOption) (res *AddVoiceResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/voices/pvc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Edit PVC voice metadata
func (r *VoicePvcService) Edit(ctx context.Context, voiceID string, params VoicePvcEditParams, opts ...option.RequestOption) (res *AddVoiceResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Request manual verification for a PVC voice.
func (r *VoicePvcService) RequestVerification(ctx context.Context, voiceID string, params VoicePvcRequestVerificationParams, opts ...option.RequestOption) (res *VoicePvcRequestVerificationResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/verification", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Start PVC training process for a voice.
func (r *VoicePvcService) Train(ctx context.Context, voiceID string, params VoicePvcTrainParams, opts ...option.RequestOption) (res *VoicePvcTrainResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/voices/pvc/%s/train", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type VoicePvcRequestVerificationResponse struct {
	// The status of the request PVC manual verification request. If the request was
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
func (r VoicePvcRequestVerificationResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcRequestVerificationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcTrainResponse struct {
	// The status of the start PVC voice training request. If the request was
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
func (r VoicePvcTrainResponse) RawJSON() string { return r.JSON.raw }
func (r *VoicePvcTrainResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcNewParams struct {
	// Language used in the samples.
	Language string `json:"language,required"`
	// The name that identifies this voice. This will be displayed in the dropdown of
	// the website.
	Name string `json:"name,required"`
	// Description to use for the created voice.
	Description param.Opt[string] `json:"description,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r VoicePvcNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VoicePvcNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicePvcNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcEditParams struct {
	// Description to use for the created voice.
	Description param.Opt[string] `json:"description,omitzero"`
	// Language used in the samples.
	Language param.Opt[string] `json:"language,omitzero"`
	// The name that identifies this voice. This will be displayed in the dropdown of
	// the website.
	Name param.Opt[string] `json:"name,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Labels for the voice. Keys can be language, accent, gender, or age.
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r VoicePvcEditParams) MarshalJSON() (data []byte, err error) {
	type shadow VoicePvcEditParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicePvcEditParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoicePvcRequestVerificationParams struct {
	// Verification documents
	Files []io.Reader `json:"files,omitzero,required" format:"binary"`
	// Extra text to be used in the manual verification process.
	ExtraText param.Opt[string] `json:"extra_text,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoicePvcRequestVerificationParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

type VoicePvcTrainParams struct {
	// The model ID to use for the conversion.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r VoicePvcTrainParams) MarshalJSON() (data []byte, err error) {
	type shadow VoicePvcTrainParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoicePvcTrainParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
