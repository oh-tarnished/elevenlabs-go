// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
)

// AudioIsolationService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioIsolationService] method instead.
type AudioIsolationService struct {
	Options []option.RequestOption
}

// NewAudioIsolationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioIsolationService(opts ...option.RequestOption) (r AudioIsolationService) {
	r = AudioIsolationService{}
	r.Options = opts
	return
}

// Removes background noise from audio
func (r *AudioIsolationService) New(ctx context.Context, params AudioIsolationNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/audio-isolation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Removes background noise from audio and streams the result
func (r *AudioIsolationService) Stream(ctx context.Context, params AudioIsolationStreamParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/audio-isolation/stream"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AudioIsolationNewParams struct {
	// The audio file from which vocals/speech will be isolated from.
	Audio io.Reader `json:"audio,omitzero,required" format:"binary"`
	// Optional preview image base64 for tracking this generation.
	PreviewB64 param.Opt[string] `json:"preview_b64,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
	// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
	// single channel (mono), and little-endian byte order. Latency will be lower than
	// with passing an encoded waveform.
	//
	// Any of "pcm_s16le_16", "other".
	FileFormat AudioIsolationNewParamsFileFormat `json:"file_format,omitzero"`
	paramObj
}

func (r AudioIsolationNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
// single channel (mono), and little-endian byte order. Latency will be lower than
// with passing an encoded waveform.
type AudioIsolationNewParamsFileFormat string

const (
	AudioIsolationNewParamsFileFormatPcmS16le16 AudioIsolationNewParamsFileFormat = "pcm_s16le_16"
	AudioIsolationNewParamsFileFormatOther      AudioIsolationNewParamsFileFormat = "other"
)

type AudioIsolationStreamParams struct {
	// The audio file from which vocals/speech will be isolated from.
	Audio io.Reader `json:"audio,omitzero,required" format:"binary"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
	// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
	// single channel (mono), and little-endian byte order. Latency will be lower than
	// with passing an encoded waveform.
	//
	// Any of "pcm_s16le_16", "other".
	FileFormat AudioIsolationStreamParamsFileFormat `json:"file_format,omitzero"`
	paramObj
}

func (r AudioIsolationStreamParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
// single channel (mono), and little-endian byte order. Latency will be lower than
// with passing an encoded waveform.
type AudioIsolationStreamParamsFileFormat string

const (
	AudioIsolationStreamParamsFileFormatPcmS16le16 AudioIsolationStreamParamsFileFormat = "pcm_s16le_16"
	AudioIsolationStreamParamsFileFormatOther      AudioIsolationStreamParamsFileFormat = "other"
)
