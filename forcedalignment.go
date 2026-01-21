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
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ForcedAlignmentService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewForcedAlignmentService] method instead.
type ForcedAlignmentService struct {
	Options []option.RequestOption
}

// NewForcedAlignmentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewForcedAlignmentService(opts ...option.RequestOption) (r ForcedAlignmentService) {
	r = ForcedAlignmentService{}
	r.Options = opts
	return
}

// Force align an audio file to text. Use this endpoint to get the timing
// information for each character and word in an audio file based on a provided
// text transcript.
func (r *ForcedAlignmentService) New(ctx context.Context, params ForcedAlignmentNewParams, opts ...option.RequestOption) (res *ForcedAlignmentNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/forced-alignment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Model representing the response from the aligner service.
type ForcedAlignmentNewResponse struct {
	// List of characters with their timing information.
	Characters []ForcedAlignmentNewResponseCharacter `json:"characters,required"`
	// The average alignment loss/confidence score for the entire transcript,
	// calculated from all characters.
	Loss float64 `json:"loss,required"`
	// List of words with their timing information.
	Words []ForcedAlignmentNewResponseWord `json:"words,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Characters  respjson.Field
		Loss        respjson.Field
		Words       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForcedAlignmentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ForcedAlignmentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representing a single character with its timing information from the
// aligner.
type ForcedAlignmentNewResponseCharacter struct {
	// The end time of the character in seconds.
	End float64 `json:"end,required"`
	// The start time of the character in seconds.
	Start float64 `json:"start,required"`
	// The character that was transcribed.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForcedAlignmentNewResponseCharacter) RawJSON() string { return r.JSON.raw }
func (r *ForcedAlignmentNewResponseCharacter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model representing a single word with its timing information from the aligner.
type ForcedAlignmentNewResponseWord struct {
	// The end time of the word in seconds.
	End float64 `json:"end,required"`
	// The average alignment loss/confidence score for this word, calculated from its
	// constituent characters.
	Loss float64 `json:"loss,required"`
	// The start time of the word in seconds.
	Start float64 `json:"start,required"`
	// The word that was transcribed.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		End         respjson.Field
		Loss        respjson.Field
		Start       respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ForcedAlignmentNewResponseWord) RawJSON() string { return r.JSON.raw }
func (r *ForcedAlignmentNewResponseWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ForcedAlignmentNewParams struct {
	// The file to align. All major audio formats are supported. The file size must be
	// less than 1GB.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// The text to align with the audio. The input text can be in any format, however
	// diarization is not supported at this time.
	Text string `json:"text,required"`
	// If true, the file will be streamed to the server and processed in chunks. This
	// is useful for large files that cannot be loaded into memory. The default is
	// false.
	EnabledSpooledFile param.Opt[bool] `json:"enabled_spooled_file,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ForcedAlignmentNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
