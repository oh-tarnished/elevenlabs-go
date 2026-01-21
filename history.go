// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// HistoryService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHistoryService] method instead.
type HistoryService struct {
	Options []option.RequestOption
}

// NewHistoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHistoryService(opts ...option.RequestOption) (r HistoryService) {
	r = HistoryService{}
	r.Options = opts
	return
}

// Retrieves a history item.
func (r *HistoryService) Get(ctx context.Context, historyItemID string, query HistoryGetParams, opts ...option.RequestOption) (res *SpeechHistoryItem, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if historyItemID == "" {
		err = errors.New("missing required history_item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/history/%s", historyItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of your generated audio.
func (r *HistoryService) List(ctx context.Context, params HistoryListParams, opts ...option.RequestOption) (res *HistoryListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/history"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Delete a history item by its ID
func (r *HistoryService) Delete(ctx context.Context, historyItemID string, body HistoryDeleteParams, opts ...option.RequestOption) (res *HistoryDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if historyItemID == "" {
		err = errors.New("missing required history_item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/history/%s", historyItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Download one or more history items. If one history item ID is provided, we will
// return a single audio file. If more than one history item IDs are provided, we
// will provide the history items packed into a .zip file.
func (r *HistoryService) Download(ctx context.Context, params HistoryDownloadParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/zip")}, opts...)
	path := "v1/history/download"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns the audio of an history item.
func (r *HistoryService) GetAudio(ctx context.Context, historyItemID string, query HistoryGetAudioParams, opts ...option.RequestOption) (res *http.Response, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "audio/mpeg")}, opts...)
	if historyItemID == "" {
		err = errors.New("missing required history_item_id parameter")
		return
	}
	path := fmt.Sprintf("v1/history/%s/audio", historyItemID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HistoryAlignment struct {
	// The end times of the characters in seconds.
	CharacterEndTimesSeconds []float64 `json:"character_end_times_seconds,required"`
	// The start times of the characters in seconds.
	CharacterStartTimesSeconds []float64 `json:"character_start_times_seconds,required"`
	// The characters in the alignment.
	Characters []string `json:"characters,required"`
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
func (r HistoryAlignment) RawJSON() string { return r.JSON.raw }
func (r *HistoryAlignment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechHistoryItem struct {
	// The character count change from.
	CharacterCountChangeFrom int64 `json:"character_count_change_from,required"`
	// The character count change to.
	CharacterCountChangeTo int64 `json:"character_count_change_to,required"`
	// The content type of the generated item.
	ContentType string `json:"content_type,required"`
	// Unix timestamp of when the item was created.
	DateUnix int64 `json:"date_unix,required"`
	// The ID of the history item.
	HistoryItemID string `json:"history_item_id,required"`
	// The state of the history item.
	//
	// Any of "created", "deleted", "processing".
	State SpeechHistoryItemState `json:"state,required"`
	// The alignments of the history item.
	Alignments SpeechHistoryItemAlignments `json:"alignments,nullable"`
	// The dialogue (voice and text pairs) used to generate the audio item. If this is
	// set then the top level `text` and `voice_id` fields will be empty.
	Dialogue []SpeechHistoryItemDialogue `json:"dialogue,nullable"`
	// Feedback associated with the generated item. Returns null if no feedback has
	// been provided.
	Feedback SpeechHistoryItemFeedback `json:"feedback,nullable"`
	// The ID of the model.
	ModelID string `json:"model_id,nullable"`
	// The ID of the request.
	RequestID string `json:"request_id,nullable"`
	// The settings of the history item.
	Settings any `json:"settings,nullable"`
	// The ID of the share link.
	ShareLinkID string `json:"share_link_id,nullable"`
	// The source of the history item. Either TTS (text to speech), STS (speech to
	// text), AN (audio native), Projects, Dubbing, PlayAPI, PD (pronunciation
	// dictionary) or ConvAI (Agents Platform).
	//
	// Any of "TTS", "STS", "Projects", "PD", "AN", "Dubbing", "PlayAPI", "ConvAI",
	// "VoiceGeneration".
	Source SpeechHistoryItemSource `json:"source,nullable"`
	// The text used to generate the audio item.
	Text string `json:"text,nullable"`
	// The category of the voice. Either 'premade', 'cloned', 'generated' or
	// 'professional'.
	//
	// Any of "premade", "cloned", "generated", "professional".
	VoiceCategory SpeechHistoryItemVoiceCategory `json:"voice_category,nullable"`
	// The ID of the voice used.
	VoiceID string `json:"voice_id,nullable"`
	// The name of the voice.
	VoiceName string `json:"voice_name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CharacterCountChangeFrom respjson.Field
		CharacterCountChangeTo   respjson.Field
		ContentType              respjson.Field
		DateUnix                 respjson.Field
		HistoryItemID            respjson.Field
		State                    respjson.Field
		Alignments               respjson.Field
		Dialogue                 respjson.Field
		Feedback                 respjson.Field
		ModelID                  respjson.Field
		RequestID                respjson.Field
		Settings                 respjson.Field
		ShareLinkID              respjson.Field
		Source                   respjson.Field
		Text                     respjson.Field
		VoiceCategory            respjson.Field
		VoiceID                  respjson.Field
		VoiceName                respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechHistoryItem) RawJSON() string { return r.JSON.raw }
func (r *SpeechHistoryItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The state of the history item.
type SpeechHistoryItemState string

const (
	SpeechHistoryItemStateCreated    SpeechHistoryItemState = "created"
	SpeechHistoryItemStateDeleted    SpeechHistoryItemState = "deleted"
	SpeechHistoryItemStateProcessing SpeechHistoryItemState = "processing"
)

// The alignments of the history item.
type SpeechHistoryItemAlignments struct {
	// The alignment of the text.
	Alignment HistoryAlignment `json:"alignment,required"`
	// The normalized alignment of the text.
	NormalizedAlignment HistoryAlignment `json:"normalized_alignment,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alignment           respjson.Field
		NormalizedAlignment respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechHistoryItemAlignments) RawJSON() string { return r.JSON.raw }
func (r *SpeechHistoryItemAlignments) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SpeechHistoryItemDialogue struct {
	// The text of the dialogue input line.
	Text string `json:"text,required"`
	// The ID of the voice used for this dialogue input line.
	VoiceID string `json:"voice_id,required"`
	// The name of the voice used for this dialogue input line.
	VoiceName string `json:"voice_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		VoiceID     respjson.Field
		VoiceName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechHistoryItemDialogue) RawJSON() string { return r.JSON.raw }
func (r *SpeechHistoryItemDialogue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Feedback associated with the generated item. Returns null if no feedback has
// been provided.
type SpeechHistoryItemFeedback struct {
	// Whether the user thinks the audio quality is good.
	AudioQuality bool `json:"audio_quality,required"`
	// Whether the user provided emotions.
	Emotions bool `json:"emotions,required"`
	// The feedback text provided by the user.
	Feedback string `json:"feedback,required"`
	// Whether the user thinks there are glitches in the audio.
	Glitches bool `json:"glitches,required"`
	// Whether the user thinks the clone is inaccurate.
	InaccurateClone bool `json:"inaccurate_clone,required"`
	// Whether the user provided other feedback.
	Other bool `json:"other,required"`
	// Whether the user liked the generated item.
	ThumbsUp bool `json:"thumbs_up,required"`
	// The review status of the item. Defaults to 'not_reviewed'.
	ReviewStatus string `json:"review_status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AudioQuality    respjson.Field
		Emotions        respjson.Field
		Feedback        respjson.Field
		Glitches        respjson.Field
		InaccurateClone respjson.Field
		Other           respjson.Field
		ThumbsUp        respjson.Field
		ReviewStatus    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SpeechHistoryItemFeedback) RawJSON() string { return r.JSON.raw }
func (r *SpeechHistoryItemFeedback) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The source of the history item. Either TTS (text to speech), STS (speech to
// text), AN (audio native), Projects, Dubbing, PlayAPI, PD (pronunciation
// dictionary) or ConvAI (Agents Platform).
type SpeechHistoryItemSource string

const (
	SpeechHistoryItemSourceTts             SpeechHistoryItemSource = "TTS"
	SpeechHistoryItemSourceSts             SpeechHistoryItemSource = "STS"
	SpeechHistoryItemSourceProjects        SpeechHistoryItemSource = "Projects"
	SpeechHistoryItemSourcePd              SpeechHistoryItemSource = "PD"
	SpeechHistoryItemSourceAn              SpeechHistoryItemSource = "AN"
	SpeechHistoryItemSourceDubbing         SpeechHistoryItemSource = "Dubbing"
	SpeechHistoryItemSourcePlayAPI         SpeechHistoryItemSource = "PlayAPI"
	SpeechHistoryItemSourceConvAI          SpeechHistoryItemSource = "ConvAI"
	SpeechHistoryItemSourceVoiceGeneration SpeechHistoryItemSource = "VoiceGeneration"
)

// The category of the voice. Either 'premade', 'cloned', 'generated' or
// 'professional'.
type SpeechHistoryItemVoiceCategory string

const (
	SpeechHistoryItemVoiceCategoryPremade      SpeechHistoryItemVoiceCategory = "premade"
	SpeechHistoryItemVoiceCategoryCloned       SpeechHistoryItemVoiceCategory = "cloned"
	SpeechHistoryItemVoiceCategoryGenerated    SpeechHistoryItemVoiceCategory = "generated"
	SpeechHistoryItemVoiceCategoryProfessional SpeechHistoryItemVoiceCategory = "professional"
)

type HistoryListResponse struct {
	// Whether there are more history items to fetch.
	HasMore bool `json:"has_more,required"`
	// A list of speech history items.
	History []SpeechHistoryItem `json:"history,required"`
	// The ID of the last history item.
	LastHistoryItemID string `json:"last_history_item_id,nullable"`
	// The timestamp of the last history item.
	ScannedUntil int64 `json:"scanned_until,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore           respjson.Field
		History           respjson.Field
		LastHistoryItemID respjson.Field
		ScannedUntil      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HistoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *HistoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryDeleteResponse struct {
	// The status of the deletion request. If the request was successful, the status
	// will be 'ok'. Otherwise an error message with http code 500 will be returned.
	Status string `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HistoryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *HistoryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type HistoryListParams struct {
	// Unix timestamp to filter history items after this date (inclusive).
	DateAfterUnix param.Opt[int64] `query:"date_after_unix,omitzero" json:"-"`
	// Unix timestamp to filter history items before this date (exclusive).
	DateBeforeUnix param.Opt[int64] `query:"date_before_unix,omitzero" json:"-"`
	// Model ID to filter history items by.
	ModelID param.Opt[string] `query:"model_id,omitzero" json:"-"`
	// search term used for filtering
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// After which ID to start fetching, use this parameter to paginate across a large
	// collection of history items. In case this parameter is not provided history
	// items will be fetched starting from the most recently created one ordered
	// descending by their creation date.
	StartAfterHistoryItemID param.Opt[string] `query:"start_after_history_item_id,omitzero" json:"-"`
	// Voice ID to be filtered for, you can use GET https://api.elevenlabs.io/v1/voices
	// to receive a list of voices and their IDs.
	VoiceID param.Opt[string] `query:"voice_id,omitzero" json:"-"`
	// How many history items to return at maximum. Can not exceed 1000, defaults
	// to 100.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Sort direction for the results.
	//
	// Any of "asc", "desc".
	SortDirection HistoryListParamsSortDirection `query:"sort_direction,omitzero" json:"-"`
	// Source of the generated history item
	//
	// Any of "TTS", "STS".
	Source HistoryListParamsSource `query:"source,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [HistoryListParams]'s query parameters as `url.Values`.
func (r HistoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort direction for the results.
type HistoryListParamsSortDirection string

const (
	HistoryListParamsSortDirectionAsc  HistoryListParamsSortDirection = "asc"
	HistoryListParamsSortDirectionDesc HistoryListParamsSortDirection = "desc"
)

// Source of the generated history item
type HistoryListParamsSource string

const (
	HistoryListParamsSourceTts HistoryListParamsSource = "TTS"
	HistoryListParamsSourceSts HistoryListParamsSource = "STS"
)

type HistoryDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type HistoryDownloadParams struct {
	// A list of history items to download, you can get IDs of history items and other
	// metadata using the GET https://api.elevenlabs.io/v1/history endpoint.
	HistoryItemIDs []string `json:"history_item_ids,omitzero,required"`
	// Output format to transcode the audio file, can be wav or default.
	OutputFormat param.Opt[string] `json:"output_format,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r HistoryDownloadParams) MarshalJSON() (data []byte, err error) {
	type shadow HistoryDownloadParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *HistoryDownloadParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HistoryGetAudioParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
