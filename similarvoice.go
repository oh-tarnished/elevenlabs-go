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

// SimilarVoiceService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimilarVoiceService] method instead.
type SimilarVoiceService struct {
	Options []option.RequestOption
}

// NewSimilarVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSimilarVoiceService(opts ...option.RequestOption) (r SimilarVoiceService) {
	r = SimilarVoiceService{}
	r.Options = opts
	return
}

// Returns a list of shared voices similar to the provided audio sample. If neither
// similarity_threshold nor top_k is provided, we will apply default values.
func (r *SimilarVoiceService) New(ctx context.Context, params SimilarVoiceNewParams, opts ...option.RequestOption) (res *LibraryVoices, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/similar-voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SimilarVoiceNewParams struct {
	// Threshold for voice similarity between provided sample and library voices.
	// Values range from 0 to 2. The smaller the value the more similar voices will be
	// returned.
	SimilarityThreshold param.Opt[float64] `json:"similarity_threshold,omitzero"`
	// Number of most similar voices to return. If similarity_threshold is provided,
	// less than this number of voices may be returned. Values range from 1 to 100.
	TopK param.Opt[int64] `json:"top_k,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey  param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	AudioFile io.Reader         `json:"audio_file,omitzero" format:"binary"`
	paramObj
}

func (r SimilarVoiceNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
