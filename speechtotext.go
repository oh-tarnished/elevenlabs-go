// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// SpeechToTextService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSpeechToTextService] method instead.
type SpeechToTextService struct {
	Options     []option.RequestOption
	Transcripts SpeechToTextTranscriptService
}

// NewSpeechToTextService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSpeechToTextService(opts ...option.RequestOption) (r SpeechToTextService) {
	r = SpeechToTextService{}
	r.Options = opts
	r.Transcripts = NewSpeechToTextTranscriptService(opts...)
	return
}

// Transcribe an audio or video file. If webhook is set to true, the request will
// be processed asynchronously and results sent to configured webhooks. When
// use_multi_channel is true and the provided audio has multiple channels, a
// 'transcripts' object with separate transcripts for each channel is returned.
// Otherwise, returns a single transcript. The optional webhook_metadata parameter
// allows you to attach custom data that will be included in webhook responses for
// request correlation and tracking.
func (r *SpeechToTextService) Transcribe(ctx context.Context, params SpeechToTextTranscribeParams, opts ...option.RequestOption) (res *SpeechToTextTranscribeResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/speech-to-text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Response model for multichannel speech-to-text transcription.
type MultichannelSpeechToText struct {
	// List of transcripts, one for each audio channel. Each transcript contains the
	// text and word-level details for its respective channel.
	Transcripts []SpeechToTextChunk `json:"transcripts,required"`
	// The transcription ID of the response.
	TranscriptionID string `json:"transcription_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Transcripts     respjson.Field
		TranscriptionID respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MultichannelSpeechToText) RawJSON() string { return r.JSON.raw }
func (r *MultichannelSpeechToText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chunk-level detail of the transcription with timing information.
type SpeechToTextChunk struct {
	// The detected language code (e.g. 'eng' for English).
	LanguageCode string `json:"language_code,required"`
	// The confidence score of the language detection (0 to 1).
	LanguageProbability float64 `json:"language_probability,required"`
	// The raw text of the transcription.
	Text string `json:"text,required"`
	// List of words with their timing information.
	Words []SpeechToTextChunkWord `json:"words,required"`
	// Requested additional formats of the transcript.
	AdditionalFormats []SpeechToTextChunkAdditionalFormat `json:"additional_formats,nullable"`
	// The channel index this transcript belongs to (for multichannel audio).
	ChannelIndex int64 `json:"channel_index,nullable"`
	// List of detected entities with their text, type, and character positions in the
	// transcript.
	Entities []SpeechToTextChunkEntity `json:"entities,nullable"`
	// The transcription ID of the response.
	TranscriptionID string `json:"transcription_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LanguageCode        respjson.Field
		LanguageProbability respjson.Field
		Text                respjson.Field
		Words               respjson.Field
		AdditionalFormats   respjson.Field
		ChannelIndex        respjson.Field
		Entities            respjson.Field
		TranscriptionID     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextChunk) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Word-level detail of the transcription with timing information.
type SpeechToTextChunkWord struct {
	// The log of the probability with which this word was predicted. Logprobs are in
	// range [-infinity, 0], higher logprobs indicate a higher confidence the model has
	// in its predictions.
	Logprob float64 `json:"logprob,required"`
	// The word or sound that was transcribed.
	Text string `json:"text,required"`
	// The type of the word or sound. 'audio_event' is used for non-word sounds like
	// laughter or footsteps.
	//
	// Any of "word", "spacing", "audio_event".
	Type string `json:"type,required"`
	// The characters that make up the word and their timing information.
	Characters []SpeechToTextChunkWordCharacter `json:"characters,nullable"`
	// The end time of the word or sound in seconds.
	End float64 `json:"end,nullable"`
	// Unique identifier for the speaker of this word.
	SpeakerID string `json:"speaker_id,nullable"`
	// The start time of the word or sound in seconds.
	Start float64 `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logprob     respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		Characters  respjson.Field
		End         respjson.Field
		SpeakerID   respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextChunkWord) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextChunkWord) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextChunkWordCharacter struct {
	// The character that was transcribed.
	Text string `json:"text,required"`
	// The end time of the character in seconds.
	End float64 `json:"end,nullable"`
	// The start time of the character in seconds.
	Start float64 `json:"start,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		End         respjson.Field
		Start       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextChunkWordCharacter) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextChunkWordCharacter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextChunkAdditionalFormat struct {
	// The content of the additional format.
	Content string `json:"content,required"`
	// The content type of the additional format.
	ContentType string `json:"content_type,required"`
	// The file extension of the additional format.
	FileExtension string `json:"file_extension,required"`
	// Whether the content is base64 encoded.
	IsBase64Encoded bool `json:"is_base64_encoded,required"`
	// The requested format.
	RequestedFormat string `json:"requested_format,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content         respjson.Field
		ContentType     respjson.Field
		FileExtension   respjson.Field
		IsBase64Encoded respjson.Field
		RequestedFormat respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextChunkAdditionalFormat) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextChunkAdditionalFormat) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextChunkEntity struct {
	// End character position in the transcript text.
	EndChar int64 `json:"end_char,required"`
	// The type of entity detected (e.g., 'credit_card', 'email_address',
	// 'person_name').
	EntityType string `json:"entity_type,required"`
	// Start character position in the transcript text.
	StartChar int64 `json:"start_char,required"`
	// The text that was identified as an entity.
	Text string `json:"text,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndChar     respjson.Field
		EntityType  respjson.Field
		StartChar   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextChunkEntity) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextChunkEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Chunk-level detail of the transcription with timing information.
type SpeechToTextTranscribeResponse struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechToTextTranscribeResponse) RawJSON() string { return r.JSON.raw }
func (r *SpeechToTextTranscribeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechToTextTranscribeParams struct {
	// The ID of the model to use for transcription.
	//
	// Any of "scribe_v1", "scribe_v2".
	ModelID SpeechToTextTranscribeParamsModelID `json:"model_id,omitzero,required"`
	// The HTTPS URL of the file to transcribe. Exactly one of the file or
	// cloud_storage_url parameters must be provided. The file must be accessible via
	// HTTPS and the file size must be less than 2GB. Any valid HTTPS URL is accepted,
	// including URLs from cloud storage providers (AWS S3, Google Cloud Storage,
	// Cloudflare R2, etc.), CDNs, or any other HTTPS source. URLs can be pre-signed or
	// include authentication tokens in query parameters.
	CloudStorageURL param.Opt[string] `json:"cloud_storage_url,omitzero"`
	// Diarization threshold to apply during speaker diarization. A higher value means
	// there will be a lower chance of one speaker being diarized as two different
	// speakers but also a higher chance of two different speakers being diarized as
	// one speaker (less total speakers predicted). A low value means there will be a
	// higher chance of one speaker being diarized as two different speakers but also a
	// lower chance of two different speakers being diarized as one speaker (more total
	// speakers predicted). Can only be set when diarize=True and num_speakers=None.
	// Defaults to None, in which case we will choose a threshold based on the model_id
	// (0.22 usually).
	DiarizationThreshold param.Opt[float64] `json:"diarization_threshold,omitzero"`
	// An ISO-639-1 or ISO-639-3 language_code corresponding to the language of the
	// audio file. Can sometimes improve transcription performance if known beforehand.
	// Defaults to null, in this case the language is predicted automatically.
	LanguageCode param.Opt[string] `json:"language_code,omitzero"`
	// The maximum amount of speakers talking in the uploaded file. Can help with
	// predicting who speaks when. The maximum amount of speakers that can be predicted
	// is 32. Defaults to null, in this case the amount of speakers is set to the
	// maximum value the model supports.
	NumSpeakers param.Opt[int64] `json:"num_speakers,omitzero"`
	// If specified, our system will make a best effort to sample deterministically,
	// such that repeated requests with the same seed and parameters should return the
	// same result. Determinism is not guaranteed. Must be an integer between 0
	// and 2147483647.
	Seed param.Opt[int64] `json:"seed,omitzero"`
	// Controls the randomness of the transcription output. Accepts values between 0.0
	// and 2.0, where higher values result in more diverse and less deterministic
	// results. If omitted, we will use a temperature based on the model you selected
	// which is usually 0.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// Optional specific webhook ID to send the transcription result to. Only valid
	// when webhook is set to true. If not provided, transcription will be sent to all
	// configured speech-to-text webhooks.
	WebhookID param.Opt[string] `json:"webhook_id,omitzero"`
	// When enable_logging is set to false zero retention mode will be used for the
	// request. This will mean log and transcript storage features are unavailable for
	// this request. Zero retention mode may only be used by enterprise customers.
	EnableLogging param.Opt[bool] `query:"enable_logging,omitzero" json:"-"`
	// Whether to annotate which speaker is currently talking in the uploaded file.
	Diarize param.Opt[bool] `json:"diarize,omitzero"`
	// Whether to tag audio events like (laughter), (footsteps), etc. in the
	// transcription.
	TagAudioEvents param.Opt[bool] `json:"tag_audio_events,omitzero"`
	// Whether the audio file contains multiple channels where each channel contains a
	// single speaker. When enabled, each channel will be transcribed independently and
	// the results will be combined. Each word in the response will include a
	// 'channel_index' field indicating which channel it was spoken on. A maximum of 5
	// channels is supported.
	UseMultiChannel param.Opt[bool] `json:"use_multi_channel,omitzero"`
	// Whether to send the transcription result to configured speech-to-text webhooks.
	// If set the request will return early without the transcription, which will be
	// delivered later via webhook.
	Webhook param.Opt[bool] `json:"webhook,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Detect entities in the transcript. Can be 'all' to detect all entities, a single
	// entity type or category string, or a list of entity types/categories. Categories
	// include 'pii', 'phi', 'pci', 'other', 'offensive_language'. When enabled,
	// detected entities will be returned in the 'entities' field with their text,
	// type, and character positions.
	EntityDetection SpeechToTextTranscribeParamsEntityDetectionUnion `json:"entity_detection,omitzero"`
	// The file to transcribe. All major audio and video formats are supported. Exactly
	// one of the file or cloud_storage_url parameters must be provided. The file size
	// must be less than 3.0GB.
	File io.Reader `json:"file,omitzero" format:"binary"`
	// Optional metadata to be included in the webhook response. This should be a JSON
	// string representing an object with a maximum depth of 2 levels and maximum size
	// of 16KB. Useful for tracking internal IDs, job references, or other contextual
	// information.
	WebhookMetadata any `json:"webhook_metadata,omitzero"`
	// A list of additional formats to export the transcript to.
	AdditionalFormats []SpeechToTextTranscribeParamsAdditionalFormatUnion `json:"additional_formats,omitzero"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
	// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
	// single channel (mono), and little-endian byte order. Latency will be lower than
	// with passing an encoded waveform.
	//
	// Any of "pcm_s16le_16", "other".
	FileFormat SpeechToTextTranscribeParamsFileFormat `json:"file_format,omitzero"`
	// A list of keyterms to bias the transcription towards. The keyterms are words or
	// phrases you want the model to recognise more accurately. The number of keyterms
	// cannot exceed 100. The length of each keyterm must be less than 50 characters.
	// Keyterms can contain at most 5 words (after normalisation). For example
	// ["hello", "world", "technical term"]
	Keyterms []string `json:"keyterms,omitzero"`
	// The granularity of the timestamps in the transcription. 'word' provides
	// word-level timestamps and 'character' provides character-level timestamps per
	// word.
	//
	// Any of "none", "word", "character".
	TimestampsGranularity SpeechToTextTranscribeParamsTimestampsGranularity `json:"timestamps_granularity,omitzero"`
	paramObj
}

func (r SpeechToTextTranscribeParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// URLQuery serializes [SpeechToTextTranscribeParams]'s query parameters as
// `url.Values`.
func (r SpeechToTextTranscribeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The ID of the model to use for transcription.
type SpeechToTextTranscribeParamsModelID string

const (
	SpeechToTextTranscribeParamsModelIDScribeV1 SpeechToTextTranscribeParamsModelID = "scribe_v1"
	SpeechToTextTranscribeParamsModelIDScribeV2 SpeechToTextTranscribeParamsModelID = "scribe_v2"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SpeechToTextTranscribeParamsAdditionalFormatUnion struct {
	OfSegmentedJson *SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson `json:",omitzero,inline"`
	OfDocx          *SpeechToTextTranscribeParamsAdditionalFormatDocx          `json:",omitzero,inline"`
	OfPdf           *SpeechToTextTranscribeParamsAdditionalFormatPdf           `json:",omitzero,inline"`
	OfTxt           *SpeechToTextTranscribeParamsAdditionalFormatTxt           `json:",omitzero,inline"`
	OfHTML          *SpeechToTextTranscribeParamsAdditionalFormatHTML          `json:",omitzero,inline"`
	OfSrt           *SpeechToTextTranscribeParamsAdditionalFormatSrt           `json:",omitzero,inline"`
	paramUnion
}

func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfSegmentedJson,
		u.OfDocx,
		u.OfPdf,
		u.OfTxt,
		u.OfHTML,
		u.OfSrt)
}
func (u *SpeechToTextTranscribeParamsAdditionalFormatUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SpeechToTextTranscribeParamsAdditionalFormatUnion) asAny() any {
	if !param.IsOmitted(u.OfSegmentedJson) {
		return u.OfSegmentedJson
	} else if !param.IsOmitted(u.OfDocx) {
		return u.OfDocx
	} else if !param.IsOmitted(u.OfPdf) {
		return u.OfPdf
	} else if !param.IsOmitted(u.OfTxt) {
		return u.OfTxt
	} else if !param.IsOmitted(u.OfHTML) {
		return u.OfHTML
	} else if !param.IsOmitted(u.OfSrt) {
		return u.OfSrt
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetFormat() *string {
	if vt := u.OfSegmentedJson; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfDocx; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfPdf; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfTxt; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfHTML; vt != nil {
		return (*string)(&vt.Format)
	} else if vt := u.OfSrt; vt != nil {
		return (*string)(&vt.Format)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetIncludeSpeakers() *bool {
	if vt := u.OfSegmentedJson; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	} else if vt := u.OfDocx; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	} else if vt := u.OfPdf; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	} else if vt := u.OfTxt; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	} else if vt := u.OfHTML; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	} else if vt := u.OfSrt; vt != nil && vt.IncludeSpeakers.Valid() {
		return &vt.IncludeSpeakers.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetIncludeTimestamps() *bool {
	if vt := u.OfSegmentedJson; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	} else if vt := u.OfDocx; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	} else if vt := u.OfPdf; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	} else if vt := u.OfTxt; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	} else if vt := u.OfHTML; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	} else if vt := u.OfSrt; vt != nil && vt.IncludeTimestamps.Valid() {
		return &vt.IncludeTimestamps.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetMaxSegmentChars() *int64 {
	if vt := u.OfSegmentedJson; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	} else if vt := u.OfDocx; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	} else if vt := u.OfPdf; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	} else if vt := u.OfTxt; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	} else if vt := u.OfHTML; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	} else if vt := u.OfSrt; vt != nil && vt.MaxSegmentChars.Valid() {
		return &vt.MaxSegmentChars.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetMaxSegmentDurationS() *float64 {
	if vt := u.OfSegmentedJson; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	} else if vt := u.OfDocx; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	} else if vt := u.OfPdf; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	} else if vt := u.OfTxt; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	} else if vt := u.OfHTML; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	} else if vt := u.OfSrt; vt != nil && vt.MaxSegmentDurationS.Valid() {
		return &vt.MaxSegmentDurationS.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetSegmentOnSilenceLongerThanS() *float64 {
	if vt := u.OfSegmentedJson; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	} else if vt := u.OfDocx; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	} else if vt := u.OfPdf; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	} else if vt := u.OfTxt; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	} else if vt := u.OfHTML; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	} else if vt := u.OfSrt; vt != nil && vt.SegmentOnSilenceLongerThanS.Valid() {
		return &vt.SegmentOnSilenceLongerThanS.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u SpeechToTextTranscribeParamsAdditionalFormatUnion) GetMaxCharactersPerLine() *int64 {
	if vt := u.OfTxt; vt != nil && vt.MaxCharactersPerLine.Valid() {
		return &vt.MaxCharactersPerLine.Value
	} else if vt := u.OfSrt; vt != nil && vt.MaxCharactersPerLine.Valid() {
		return &vt.MaxCharactersPerLine.Value
	}
	return nil
}

func init() {
	apijson.RegisterUnion[SpeechToTextTranscribeParamsAdditionalFormatUnion](
		"format",
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson]("segmented_json"),
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatDocx]("docx"),
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatPdf]("pdf"),
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatTxt]("txt"),
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatHTML]("html"),
		apijson.Discriminator[SpeechToTextTranscribeParamsAdditionalFormatSrt]("srt"),
	)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson struct {
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "segmented_json".
	Format constant.SegmentedJson `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatSegmentedJson) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatDocx struct {
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "docx".
	Format constant.Docx `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatDocx) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatDocx
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatDocx) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatPdf struct {
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "pdf".
	Format constant.Pdf `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatPdf) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatPdf
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatPdf) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatTxt struct {
	MaxCharactersPerLine        param.Opt[int64]   `json:"max_characters_per_line,omitzero"`
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "txt".
	Format constant.Txt `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatTxt) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatTxt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatTxt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatHTML struct {
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "html".
	Format constant.HTML `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatHTML) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatHTML
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatHTML) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Format is required.
type SpeechToTextTranscribeParamsAdditionalFormatSrt struct {
	MaxCharactersPerLine        param.Opt[int64]   `json:"max_characters_per_line,omitzero"`
	MaxSegmentChars             param.Opt[int64]   `json:"max_segment_chars,omitzero"`
	MaxSegmentDurationS         param.Opt[float64] `json:"max_segment_duration_s,omitzero"`
	SegmentOnSilenceLongerThanS param.Opt[float64] `json:"segment_on_silence_longer_than_s,omitzero"`
	IncludeSpeakers             param.Opt[bool]    `json:"include_speakers,omitzero"`
	IncludeTimestamps           param.Opt[bool]    `json:"include_timestamps,omitzero"`
	// This field can be elided, and will marshal its zero value as "srt".
	Format constant.Srt `json:"format,required"`
	paramObj
}

func (r SpeechToTextTranscribeParamsAdditionalFormatSrt) MarshalJSON() (data []byte, err error) {
	type shadow SpeechToTextTranscribeParamsAdditionalFormatSrt
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SpeechToTextTranscribeParamsAdditionalFormatSrt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SpeechToTextTranscribeParamsEntityDetectionUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u SpeechToTextTranscribeParamsEntityDetectionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *SpeechToTextTranscribeParamsEntityDetectionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SpeechToTextTranscribeParamsEntityDetectionUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// The format of input audio. Options are 'pcm_s16le_16' or 'other' For
// `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate,
// single channel (mono), and little-endian byte order. Latency will be lower than
// with passing an encoded waveform.
type SpeechToTextTranscribeParamsFileFormat string

const (
	SpeechToTextTranscribeParamsFileFormatPcmS16le16 SpeechToTextTranscribeParamsFileFormat = "pcm_s16le_16"
	SpeechToTextTranscribeParamsFileFormatOther      SpeechToTextTranscribeParamsFileFormat = "other"
)

// The granularity of the timestamps in the transcription. 'word' provides
// word-level timestamps and 'character' provides character-level timestamps per
// word.
type SpeechToTextTranscribeParamsTimestampsGranularity string

const (
	SpeechToTextTranscribeParamsTimestampsGranularityNone      SpeechToTextTranscribeParamsTimestampsGranularity = "none"
	SpeechToTextTranscribeParamsTimestampsGranularityWord      SpeechToTextTranscribeParamsTimestampsGranularity = "word"
	SpeechToTextTranscribeParamsTimestampsGranularityCharacter SpeechToTextTranscribeParamsTimestampsGranularity = "character"
)
