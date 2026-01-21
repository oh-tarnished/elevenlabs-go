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
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
)

// SpeechToSpeechService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeechToSpeechService] method instead.
type SpeechToSpeechService struct {
	Options []option.RequestOption
}

// NewSpeechToSpeechService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeechToSpeechService(opts ...option.RequestOption) (r SpeechToSpeechService) {
	r = SpeechToSpeechService{}
	r.Options = opts
	return
}

// Stream audio from one voice to another. Maintain full control over emotion,
// timing and delivery.
func (r *SpeechToSpeechService) StreamTransform(ctx context.Context, voiceID string, params SpeechToSpeechStreamTransformParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/speech-to-speech/%s/stream", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Transform audio from one voice to another. Maintain full control over emotion,
// timing and delivery.
func (r *SpeechToSpeechService) Transform(ctx context.Context, voiceID string, params SpeechToSpeechTransformParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if voiceID == "" {
		err = errors.New("missing required voice_id parameter")
		return
	}
	path := fmt.Sprintf("v1/speech-to-speech/%s", voiceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SpeechToSpeechStreamTransformParams struct {
	// The audio file which holds the content and emotion that will control the
	// generated speech.
	Audio io.Reader `json:"audio,omitzero,required" format:"binary"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request. Needs to be send as a JSON encoded string.
	VoiceSettings param.Opt[string] `json:"voice_settings,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for speech to speech, you can check
	// this using the can_do_voice_conversion property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If set, will remove the background noise from your audio input using our audio
	// isolation model. Only applies to Voice Changer.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
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
	FileFormat SpeechToSpeechStreamTransformParamsFileFormat `json:"file_format,omitzero"`
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
	OutputFormat SpeechToSpeechStreamTransformParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	paramObj
}

func (r SpeechToSpeechStreamTransformParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [SpeechToSpeechStreamTransformParams]'s query parameters as
// `url.Values`.
func (r SpeechToSpeechStreamTransformParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
// Note that the μ-law format (sometimes written mu-law, often approximated as
// u-law) is commonly used for Twilio audio inputs.
type SpeechToSpeechStreamTransformParamsOutputFormat string

const (
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_22050_32  SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_22050_32"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_24000_48  SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_24000_48"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_44100_32  SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_44100_32"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_44100_64  SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_44100_64"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_44100_96  SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_44100_96"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_44100_128 SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_44100_128"
	SpeechToSpeechStreamTransformParamsOutputFormatMP3_44100_192 SpeechToSpeechStreamTransformParamsOutputFormat = "mp3_44100_192"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm8000       SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_8000"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm16000      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_16000"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm22050      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_22050"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm24000      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_24000"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm32000      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_32000"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm44100      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_44100"
	SpeechToSpeechStreamTransformParamsOutputFormatPcm48000      SpeechToSpeechStreamTransformParamsOutputFormat = "pcm_48000"
	SpeechToSpeechStreamTransformParamsOutputFormatUlaw8000      SpeechToSpeechStreamTransformParamsOutputFormat = "ulaw_8000"
	SpeechToSpeechStreamTransformParamsOutputFormatAlaw8000      SpeechToSpeechStreamTransformParamsOutputFormat = "alaw_8000"
	SpeechToSpeechStreamTransformParamsOutputFormatOpus48000_32  SpeechToSpeechStreamTransformParamsOutputFormat = "opus_48000_32"
	SpeechToSpeechStreamTransformParamsOutputFormatOpus48000_64  SpeechToSpeechStreamTransformParamsOutputFormat = "opus_48000_64"
	SpeechToSpeechStreamTransformParamsOutputFormatOpus48000_96  SpeechToSpeechStreamTransformParamsOutputFormat = "opus_48000_96"
	SpeechToSpeechStreamTransformParamsOutputFormatOpus48000_128 SpeechToSpeechStreamTransformParamsOutputFormat = "opus_48000_128"
	SpeechToSpeechStreamTransformParamsOutputFormatOpus48000_192 SpeechToSpeechStreamTransformParamsOutputFormat = "opus_48000_192"
)

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
// single channel (mono), and little-endian byte order. Latency will be lower than
// with passing an encoded waveform.
type SpeechToSpeechStreamTransformParamsFileFormat string

const (
	SpeechToSpeechStreamTransformParamsFileFormatPcmS16le16 SpeechToSpeechStreamTransformParamsFileFormat = "pcm_s16le_16"
	SpeechToSpeechStreamTransformParamsFileFormatOther      SpeechToSpeechStreamTransformParamsFileFormat = "other"
)

type SpeechToSpeechTransformParams struct {
	// The audio file which holds the content and emotion that will control the
	// generated speech.
	Audio io.Reader `json:"audio,omitzero,required" format:"binary"`
	// You can turn on latency optimizations at some cost of quality. The best possible
	// final latency varies by model. Possible values: 0 - default mode (no latency
	// optimizations) 1 - normal latency optimizations (about 50% of possible latency
	// improvement of option 3) 2 - strong latency optimizations (about 75% of possible
	// latency improvement of option 3) 3 - max latency optimizations 4 - max latency
	// optimizations, but also with text normalizer turned off for even more latency
	// savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	OptimizeStreamingLatency param.Opt[int64] `query:"optimize_streaming_latency,omitzero" json:"-"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be integer between 0
	// and 4294967295.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Voice settings overriding stored settings for the given voice. They are applied
	// only on the given request. Needs to be send as a JSON encoded string.
	VoiceSettings param.Opt[string] `json:"voice_settings,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean history features are unavailable for this request,
	// including request stitching. Zero retention mode may only be used by enterprise
	// customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// Identifier of the model that will be used, you can query them using GET
	// /v1/models. The model needs to have support for speech to speech, you can check
	// this using the can_do_voice_conversion property.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// If set, will remove the background noise from your audio input using our audio
	// isolation model. Only applies to Voice Changer.
	RemoveBackgroundNoise param.Opt[bool] `json:"remove_background_noise,omitzero"`
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
	FileFormat SpeechToSpeechTransformParamsFileFormat `json:"file_format,omitzero"`
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
	OutputFormat SpeechToSpeechTransformParamsOutputFormat `query:"output_format,omitzero" json:"-"`
	paramObj
}

func (r SpeechToSpeechTransformParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [SpeechToSpeechTransformParams]'s query parameters as
// `url.Values`.
func (r SpeechToSpeechTransformParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So
// an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3
// with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM
// with 44.1kHz sample rate requires you to be subscribed to Pro tier or above.
// Note that the μ-law format (sometimes written mu-law, often approximated as
// u-law) is commonly used for Twilio audio inputs.
type SpeechToSpeechTransformParamsOutputFormat string

const (
	SpeechToSpeechTransformParamsOutputFormatMP3_22050_32  SpeechToSpeechTransformParamsOutputFormat = "mp3_22050_32"
	SpeechToSpeechTransformParamsOutputFormatMP3_24000_48  SpeechToSpeechTransformParamsOutputFormat = "mp3_24000_48"
	SpeechToSpeechTransformParamsOutputFormatMP3_44100_32  SpeechToSpeechTransformParamsOutputFormat = "mp3_44100_32"
	SpeechToSpeechTransformParamsOutputFormatMP3_44100_64  SpeechToSpeechTransformParamsOutputFormat = "mp3_44100_64"
	SpeechToSpeechTransformParamsOutputFormatMP3_44100_96  SpeechToSpeechTransformParamsOutputFormat = "mp3_44100_96"
	SpeechToSpeechTransformParamsOutputFormatMP3_44100_128 SpeechToSpeechTransformParamsOutputFormat = "mp3_44100_128"
	SpeechToSpeechTransformParamsOutputFormatMP3_44100_192 SpeechToSpeechTransformParamsOutputFormat = "mp3_44100_192"
	SpeechToSpeechTransformParamsOutputFormatPcm8000       SpeechToSpeechTransformParamsOutputFormat = "pcm_8000"
	SpeechToSpeechTransformParamsOutputFormatPcm16000      SpeechToSpeechTransformParamsOutputFormat = "pcm_16000"
	SpeechToSpeechTransformParamsOutputFormatPcm22050      SpeechToSpeechTransformParamsOutputFormat = "pcm_22050"
	SpeechToSpeechTransformParamsOutputFormatPcm24000      SpeechToSpeechTransformParamsOutputFormat = "pcm_24000"
	SpeechToSpeechTransformParamsOutputFormatPcm32000      SpeechToSpeechTransformParamsOutputFormat = "pcm_32000"
	SpeechToSpeechTransformParamsOutputFormatPcm44100      SpeechToSpeechTransformParamsOutputFormat = "pcm_44100"
	SpeechToSpeechTransformParamsOutputFormatPcm48000      SpeechToSpeechTransformParamsOutputFormat = "pcm_48000"
	SpeechToSpeechTransformParamsOutputFormatUlaw8000      SpeechToSpeechTransformParamsOutputFormat = "ulaw_8000"
	SpeechToSpeechTransformParamsOutputFormatAlaw8000      SpeechToSpeechTransformParamsOutputFormat = "alaw_8000"
	SpeechToSpeechTransformParamsOutputFormatOpus48000_32  SpeechToSpeechTransformParamsOutputFormat = "opus_48000_32"
	SpeechToSpeechTransformParamsOutputFormatOpus48000_64  SpeechToSpeechTransformParamsOutputFormat = "opus_48000_64"
	SpeechToSpeechTransformParamsOutputFormatOpus48000_96  SpeechToSpeechTransformParamsOutputFormat = "opus_48000_96"
	SpeechToSpeechTransformParamsOutputFormatOpus48000_128 SpeechToSpeechTransformParamsOutputFormat = "opus_48000_128"
	SpeechToSpeechTransformParamsOutputFormatOpus48000_192 SpeechToSpeechTransformParamsOutputFormat = "opus_48000_192"
)

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
// single channel (mono), and little-endian byte order. Latency will be lower than
// with passing an encoded waveform.
type SpeechToSpeechTransformParamsFileFormat string

const (
	SpeechToSpeechTransformParamsFileFormatPcmS16le16 SpeechToSpeechTransformParamsFileFormat = "pcm_s16le_16"
	SpeechToSpeechTransformParamsFileFormatOther      SpeechToSpeechTransformParamsFileFormat = "other"
)
