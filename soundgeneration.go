// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
)

// SoundGenerationService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSoundGenerationService] method instead.
type SoundGenerationService struct {
	Options []option.RequestOption
}

// NewSoundGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSoundGenerationService(opts ...option.RequestOption) (r SoundGenerationService) {
	r = SoundGenerationService{}
	r.Options = opts
	return
}

// Turn text into sound effects for your videos, voice-overs or video games using
// the most advanced sound effects models in the world.
func (r *SoundGenerationService) New(ctx context.Context, params SoundGenerationNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	path := "v1/sound-generation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AllowedOutputFormats string

const (
	AllowedOutputFormatsMP3_22050_32  AllowedOutputFormats = "mp3_22050_32"
	AllowedOutputFormatsMP3_24000_48  AllowedOutputFormats = "mp3_24000_48"
	AllowedOutputFormatsMP3_44100_32  AllowedOutputFormats = "mp3_44100_32"
	AllowedOutputFormatsMP3_44100_64  AllowedOutputFormats = "mp3_44100_64"
	AllowedOutputFormatsMP3_44100_96  AllowedOutputFormats = "mp3_44100_96"
	AllowedOutputFormatsMP3_44100_128 AllowedOutputFormats = "mp3_44100_128"
	AllowedOutputFormatsMP3_44100_192 AllowedOutputFormats = "mp3_44100_192"
	AllowedOutputFormatsPcm8000       AllowedOutputFormats = "pcm_8000"
	AllowedOutputFormatsPcm16000      AllowedOutputFormats = "pcm_16000"
	AllowedOutputFormatsPcm22050      AllowedOutputFormats = "pcm_22050"
	AllowedOutputFormatsPcm24000      AllowedOutputFormats = "pcm_24000"
	AllowedOutputFormatsPcm32000      AllowedOutputFormats = "pcm_32000"
	AllowedOutputFormatsPcm44100      AllowedOutputFormats = "pcm_44100"
	AllowedOutputFormatsPcm48000      AllowedOutputFormats = "pcm_48000"
	AllowedOutputFormatsUlaw8000      AllowedOutputFormats = "ulaw_8000"
	AllowedOutputFormatsAlaw8000      AllowedOutputFormats = "alaw_8000"
	AllowedOutputFormatsOpus48000_32  AllowedOutputFormats = "opus_48000_32"
	AllowedOutputFormatsOpus48000_64  AllowedOutputFormats = "opus_48000_64"
	AllowedOutputFormatsOpus48000_96  AllowedOutputFormats = "opus_48000_96"
	AllowedOutputFormatsOpus48000_128 AllowedOutputFormats = "opus_48000_128"
	AllowedOutputFormatsOpus48000_192 AllowedOutputFormats = "opus_48000_192"
)

type SoundGenerationNewParams struct {
	// The text that will get converted into a sound effect.
	Text string `json:"text,required"`
	// The duration of the sound which will be generated in seconds. Must be at least
	// 0.5 and at most 30. If set to None we will guess the optimal duration using the
	// prompt. Defaults to None.
	DurationSeconds param.Opt[float64] `json:"duration_seconds,omitzero"`
	// A higher prompt influence makes your generation follow the prompt more closely
	// while also making generations less variable. Must be a value between 0 and 1.
	// Defaults to 0.3.
	PromptInfluence param.Opt[float64] `json:"prompt_influence,omitzero"`
	// Whether to create a sound effect that loops smoothly. Only available for the
	// 'eleven_text_to_sound_v2 model'.
	Loop param.Opt[bool] `json:"loop,omitzero"`
	// The model ID to use for the sound generation.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
	// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
	// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
	// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
	// Note that the μ-law format (sometimes written mu-law, often approximated as
	// u-law) is commonly used for Twilio audio inputs.
	//
	// Any of "mp3_22050_32", "mp3_24000_48", "mp3_44100_32", "mp3_44100_64",
	// "mp3_44100_96", "mp3_44100_128", "mp3_44100_192", "pcm_8000", "pcm_16000",
	// "pcm_22050", "pcm_24000", "pcm_32000", "pcm_44100", "pcm_48000", "ulaw_8000",
	// "alaw_8000", "opus_48000_32", "opus_48000_64", "opus_48000_96",
	// "opus_48000_128", "opus_48000_192".
	OutputFormat AllowedOutputFormats `query:"output_format,omitzero" json:"-"`
	paramObj
}

func (r SoundGenerationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SoundGenerationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SoundGenerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SoundGenerationNewParams]'s query parameters as
// `url.Values`.
func (r SoundGenerationNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
