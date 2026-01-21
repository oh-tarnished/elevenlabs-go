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
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// SharedVoiceService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSharedVoiceService] method instead.
type SharedVoiceService struct {
	Options []option.RequestOption
}

// NewSharedVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSharedVoiceService(opts ...option.RequestOption) (r SharedVoiceService) {
	r = SharedVoiceService{}
	r.Options = opts
	return
}

// Retrieves a list of shared voices.
func (r *SharedVoiceService) List(ctx context.Context, params SharedVoiceListParams, opts ...option.RequestOption) (res *LibraryVoices, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/shared-voices"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type LibraryVoices struct {
	// Whether there are more shared voices in subsequent pages.
	HasMore bool `json:"has_more,required"`
	// The list of shared voices
	Voices     []LibraryVoicesVoice `json:"voices,required"`
	LastSortID string               `json:"last_sort_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Voices      respjson.Field
		LastSortID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LibraryVoices) RawJSON() string { return r.JSON.raw }
func (r *LibraryVoices) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LibraryVoicesVoice struct {
	// The accent of the voice.
	Accent string `json:"accent,required"`
	// The age of the voice.
	Age string `json:"age,required"`
	// The category of the voice.
	//
	// Any of "generated", "cloned", "premade", "professional", "famous",
	// "high_quality".
	Category string `json:"category,required"`
	// The number of times the voice has been cloned.
	ClonedByCount int64 `json:"cloned_by_count,required"`
	// The date the voice was added to the library in Unix time.
	DateUnix int64 `json:"date_unix,required"`
	// The descriptive of the voice.
	Descriptive string `json:"descriptive,required"`
	// Whether the voice is featured.
	Featured bool `json:"featured,required"`
	// Whether free users are allowed to use the voice.
	FreeUsersAllowed bool `json:"free_users_allowed,required"`
	// The gender of the voice.
	Gender string `json:"gender,required"`
	// Whether live moderation is enabled for the voice.
	LiveModerationEnabled bool `json:"live_moderation_enabled,required"`
	// The name of the voice.
	Name string `json:"name,required"`
	// The play API usage character count of the voice in the last year.
	PlayAPIUsageCharacterCount1y int64 `json:"play_api_usage_character_count_1y,required"`
	// The public owner id of the voice.
	PublicOwnerID string `json:"public_owner_id,required"`
	// The usage character count of the voice in the last year.
	UsageCharacterCount1y int64 `json:"usage_character_count_1y,required"`
	// The usage character count of the voice in the last 7 days.
	UsageCharacterCount7d int64 `json:"usage_character_count_7d,required"`
	// The use case of the voice.
	UseCase string `json:"use_case,required"`
	// The id of the voice.
	VoiceID string `json:"voice_id,required"`
	// The description of the voice.
	Description string `json:"description,nullable"`
	// The rate of the voice in USD per 1000 credits. null if default
	FiatRate float64 `json:"fiat_rate,nullable"`
	// The image URL of the voice.
	ImageURL string `json:"image_url,nullable"`
	// The Instagram username of the voice.
	InstagramUsername string `json:"instagram_username,nullable"`
	// Whether the voice was added by the user.
	IsAddedByUser bool `json:"is_added_by_user,nullable"`
	// The language of the voice.
	Language string `json:"language,nullable"`
	// The locale of the voice.
	Locale string `json:"locale,nullable"`
	// The notice period of the voice.
	NoticePeriod int64 `json:"notice_period,nullable"`
	// The preview URL of the voice.
	PreviewURL string `json:"preview_url,nullable"`
	// The rate multiplier of the voice.
	Rate float64 `json:"rate,nullable"`
	// The TikTok username of the voice.
	TiktokUsername string `json:"tiktok_username,nullable"`
	// The Twitter username of the voice.
	TwitterUsername string `json:"twitter_username,nullable"`
	// The verified languages of the voice.
	VerifiedLanguages []VerifiedVoiceLanguage `json:"verified_languages,nullable"`
	// The YouTube username of the voice.
	YoutubeUsername string `json:"youtube_username,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accent                       respjson.Field
		Age                          respjson.Field
		Category                     respjson.Field
		ClonedByCount                respjson.Field
		DateUnix                     respjson.Field
		Descriptive                  respjson.Field
		Featured                     respjson.Field
		FreeUsersAllowed             respjson.Field
		Gender                       respjson.Field
		LiveModerationEnabled        respjson.Field
		Name                         respjson.Field
		PlayAPIUsageCharacterCount1y respjson.Field
		PublicOwnerID                respjson.Field
		UsageCharacterCount1y        respjson.Field
		UsageCharacterCount7d        respjson.Field
		UseCase                      respjson.Field
		VoiceID                      respjson.Field
		Description                  respjson.Field
		FiatRate                     respjson.Field
		ImageURL                     respjson.Field
		InstagramUsername            respjson.Field
		IsAddedByUser                respjson.Field
		Language                     respjson.Field
		Locale                       respjson.Field
		NoticePeriod                 respjson.Field
		PreviewURL                   respjson.Field
		Rate                         respjson.Field
		TiktokUsername               respjson.Field
		TwitterUsername              respjson.Field
		VerifiedLanguages            respjson.Field
		YoutubeUsername              respjson.Field
		ExtraFields                  map[string]respjson.Field
		raw                          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LibraryVoicesVoice) RawJSON() string { return r.JSON.raw }
func (r *LibraryVoicesVoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SharedVoiceListParams struct {
	// Accent used for filtering
	Accent param.Opt[string] `query:"accent,omitzero" json:"-"`
	// Age used for filtering
	Age param.Opt[string] `query:"age,omitzero" json:"-"`
	// Gender used for filtering
	Gender param.Opt[string] `query:"gender,omitzero" json:"-"`
	// Include/exclude voices with custom rates
	IncludeCustomRates param.Opt[bool] `query:"include_custom_rates,omitzero" json:"-"`
	// Include/exclude voices that are live moderated
	IncludeLiveModerated param.Opt[bool] `query:"include_live_moderated,omitzero" json:"-"`
	// Language used for filtering
	Language param.Opt[string] `query:"language,omitzero" json:"-"`
	// Locale used for filtering
	Locale param.Opt[string] `query:"locale,omitzero" json:"-"`
	// Filter voices with a minimum notice period of the given number of days.
	MinNoticePeriodDays param.Opt[int64] `query:"min_notice_period_days,omitzero" json:"-"`
	// Filter voices by public owner ID
	OwnerID param.Opt[string] `query:"owner_id,omitzero" json:"-"`
	// Search term used for filtering
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Sort criteria
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Filter featured voices
	Featured param.Opt[bool]  `query:"featured,omitzero" json:"-"`
	Page     param.Opt[int64] `query:"page,omitzero" json:"-"`
	// How many shared voices to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Filter voices that are enabled for the reader app
	ReaderAppEnabled param.Opt[bool] `query:"reader_app_enabled,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Voice category used for filtering
	//
	// Any of "professional", "famous", "high_quality".
	Category SharedVoiceListParamsCategory `query:"category,omitzero" json:"-"`
	// Search term used for filtering
	Descriptives []string `query:"descriptives,omitzero" json:"-"`
	// Use-case used for filtering
	UseCases []string `query:"use_cases,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SharedVoiceListParams]'s query parameters as `url.Values`.
func (r SharedVoiceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Voice category used for filtering
type SharedVoiceListParamsCategory string

const (
	SharedVoiceListParamsCategoryProfessional SharedVoiceListParamsCategory = "professional"
	SharedVoiceListParamsCategoryFamous       SharedVoiceListParamsCategory = "famous"
	SharedVoiceListParamsCategoryHighQuality  SharedVoiceListParamsCategory = "high_quality"
)
