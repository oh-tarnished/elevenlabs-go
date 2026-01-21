// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// VoiceGenerationService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVoiceGenerationService] method instead.
type VoiceGenerationService struct {
	Options       []option.RequestOption
	GenerateVoice VoiceGenerationGenerateVoiceService
}

// NewVoiceGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVoiceGenerationService(opts ...option.RequestOption) (r VoiceGenerationService) {
	r = VoiceGenerationService{}
	r.Options = opts
	r.GenerateVoice = NewVoiceGenerationGenerateVoiceService(opts...)
	return
}

// Create a previously generated voice. This endpoint should be called after you
// fetched a generated_voice_id using /v1/voice-generation/generate-voice.
func (r *VoiceGenerationService) NewVoice(ctx context.Context, params VoiceGenerationNewVoiceParams, opts ...option.RequestOption) (res *Voice, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/voice-generation/create-voice"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type VerificationAttempt struct {
	// Whether the verification attempt was accepted.
	Accepted bool `json:"accepted,required"`
	// The date of the verification attempt in Unix time.
	DateUnix int64 `json:"date_unix,required"`
	// The Levenshtein distance of the verification attempt.
	LevenshteinDistance float64 `json:"levenshtein_distance,required"`
	// The similarity of the verification attempt.
	Similarity float64 `json:"similarity,required"`
	// The text of the verification attempt.
	Text string `json:"text,required"`
	// The recording of the verification attempt.
	Recording VerificationAttemptRecording `json:"recording,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted            respjson.Field
		DateUnix            respjson.Field
		LevenshteinDistance respjson.Field
		Similarity          respjson.Field
		Text                respjson.Field
		Recording           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationAttempt) RawJSON() string { return r.JSON.raw }
func (r *VerificationAttempt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The recording of the verification attempt.
type VerificationAttemptRecording struct {
	// The MIME type of the recording.
	MimeType string `json:"mime_type,required"`
	// The ID of the recording.
	RecordingID string `json:"recording_id,required"`
	// The size of the recording in bytes.
	SizeBytes int64 `json:"size_bytes,required"`
	// The transcription of the recording.
	Transcription string `json:"transcription,required"`
	// The date of the recording in Unix time.
	UploadDateUnix int64 `json:"upload_date_unix,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MimeType       respjson.Field
		RecordingID    respjson.Field
		SizeBytes      respjson.Field
		Transcription  respjson.Field
		UploadDateUnix respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationAttemptRecording) RawJSON() string { return r.JSON.raw }
func (r *VerificationAttemptRecording) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerifiedVoiceLanguage struct {
	// The language of the voice.
	Language string `json:"language,required"`
	// The voice's model ID.
	ModelID string `json:"model_id,required"`
	// The voice's accent, if applicable.
	Accent string `json:"accent,nullable"`
	// The voice's locale, if applicable.
	Locale string `json:"locale,nullable"`
	// The voice's preview URL, if applicable.
	PreviewURL string `json:"preview_url,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Language    respjson.Field
		ModelID     respjson.Field
		Accent      respjson.Field
		Locale      respjson.Field
		PreviewURL  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerifiedVoiceLanguage) RawJSON() string { return r.JSON.raw }
func (r *VerifiedVoiceLanguage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Voice struct {
	// The tiers the voice is available for.
	AvailableForTiers []string `json:"available_for_tiers,required"`
	// The category of the voice.
	//
	// Any of "generated", "cloned", "premade", "professional", "famous",
	// "high_quality".
	Category VoiceCategory `json:"category,required"`
	// The base model IDs for high-quality voices.
	HighQualityBaseModelIDs []string `json:"high_quality_base_model_ids,required"`
	// Labels associated with the voice.
	Labels map[string]string `json:"labels,required"`
	// The name of the voice.
	Name string `json:"name,required"`
	// The ID of the voice.
	VoiceID string `json:"voice_id,required"`
	// The IDs of collections this voice belongs to.
	CollectionIDs []string `json:"collection_ids,nullable"`
	// The creation time of the voice in Unix time.
	CreatedAtUnix int64 `json:"created_at_unix,nullable"`
	// The description of the voice.
	Description string `json:"description,nullable"`
	// Timestamp when the voice was marked as favorite in Unix time.
	FavoritedAtUnix int64 `json:"favorited_at_unix,nullable"`
	// Fine-tuning information for the voice.
	FineTuning VoiceFineTuning `json:"fine_tuning,nullable"`
	// Whether the voice is legacy.
	IsLegacy bool `json:"is_legacy"`
	// Whether the voice is mixed.
	IsMixed bool `json:"is_mixed"`
	// Whether the voice is owned by the user.
	IsOwner bool `json:"is_owner,nullable"`
	// The permission on the resource of the voice.
	PermissionOnResource string `json:"permission_on_resource,nullable"`
	// The preview URL of the voice.
	PreviewURL string `json:"preview_url,nullable"`
	// The safety controls of the voice.
	//
	// Any of "NONE", "BAN", "CAPTCHA", "ENTERPRISE_BAN", "ENTERPRISE_CAPTCHA".
	SafetyControl VoiceSafetyControl `json:"safety_control,nullable"`
	// List of samples associated with the voice.
	Samples []SampleResponse `json:"samples,nullable"`
	// The settings of the voice.
	Settings VoiceSettings `json:"settings,nullable"`
	// The sharing information of the voice.
	Sharing VoiceSharing `json:"sharing,nullable"`
	// The verified languages of the voice.
	VerifiedLanguages []VerifiedVoiceLanguage `json:"verified_languages,nullable"`
	// The voice verification of the voice.
	VoiceVerification VoiceVoiceVerification `json:"voice_verification,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AvailableForTiers       respjson.Field
		Category                respjson.Field
		HighQualityBaseModelIDs respjson.Field
		Labels                  respjson.Field
		Name                    respjson.Field
		VoiceID                 respjson.Field
		CollectionIDs           respjson.Field
		CreatedAtUnix           respjson.Field
		Description             respjson.Field
		FavoritedAtUnix         respjson.Field
		FineTuning              respjson.Field
		IsLegacy                respjson.Field
		IsMixed                 respjson.Field
		IsOwner                 respjson.Field
		PermissionOnResource    respjson.Field
		PreviewURL              respjson.Field
		SafetyControl           respjson.Field
		Samples                 respjson.Field
		Settings                respjson.Field
		Sharing                 respjson.Field
		VerifiedLanguages       respjson.Field
		VoiceVerification       respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Voice) RawJSON() string { return r.JSON.raw }
func (r *Voice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The category of the voice.
type VoiceCategory string

const (
	VoiceCategoryGenerated    VoiceCategory = "generated"
	VoiceCategoryCloned       VoiceCategory = "cloned"
	VoiceCategoryPremade      VoiceCategory = "premade"
	VoiceCategoryProfessional VoiceCategory = "professional"
	VoiceCategoryFamous       VoiceCategory = "famous"
	VoiceCategoryHighQuality  VoiceCategory = "high_quality"
)

// Fine-tuning information for the voice.
type VoiceFineTuning struct {
	// Whether the user is allowed to fine-tune the voice.
	IsAllowedToFineTune bool `json:"is_allowed_to_fine_tune,required"`
	// Whether a manual verification was requested for the fine-tuning process.
	ManualVerificationRequested bool `json:"manual_verification_requested,required"`
	// The state of the fine-tuning process for each model.
	//
	// Any of "not_started", "queued", "fine_tuning", "fine_tuned", "failed",
	// "delayed".
	State map[string]string `json:"state,required"`
	// The number of verification attempts in the fine-tuning process.
	VerificationAttemptsCount int64 `json:"verification_attempts_count,required"`
	// List of verification failures in the fine-tuning process.
	VerificationFailures []string `json:"verification_failures,required"`
	// The duration of the dataset in seconds.
	DatasetDurationSeconds float64 `json:"dataset_duration_seconds,nullable"`
	// The language of the fine-tuning process.
	Language string `json:"language,nullable"`
	// The manual verification of the fine-tuning process.
	ManualVerification VoiceFineTuningManualVerification `json:"manual_verification,nullable"`
	// The maximum number of verification attempts.
	MaxVerificationAttempts int64 `json:"max_verification_attempts,nullable"`
	// The message of the fine-tuning process.
	Message map[string]string `json:"message,nullable"`
	// The next maximum verification attempts reset time in Unix milliseconds.
	NextMaxVerificationAttemptsResetUnixMs int64 `json:"next_max_verification_attempts_reset_unix_ms,nullable"`
	// The progress of the fine-tuning process.
	Progress map[string]float64 `json:"progress,nullable"`
	// List of slice IDs.
	SliceIDs []string `json:"slice_ids,nullable"`
	// The number of verification attempts.
	VerificationAttempts []VerificationAttempt `json:"verification_attempts,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAllowedToFineTune                    respjson.Field
		ManualVerificationRequested            respjson.Field
		State                                  respjson.Field
		VerificationAttemptsCount              respjson.Field
		VerificationFailures                   respjson.Field
		DatasetDurationSeconds                 respjson.Field
		Language                               respjson.Field
		ManualVerification                     respjson.Field
		MaxVerificationAttempts                respjson.Field
		Message                                respjson.Field
		NextMaxVerificationAttemptsResetUnixMs respjson.Field
		Progress                               respjson.Field
		SliceIDs                               respjson.Field
		VerificationAttempts                   respjson.Field
		ExtraFields                            map[string]respjson.Field
		raw                                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceFineTuning) RawJSON() string { return r.JSON.raw }
func (r *VoiceFineTuning) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The manual verification of the fine-tuning process.
type VoiceFineTuningManualVerification struct {
	// The extra text of the manual verification.
	ExtraText string `json:"extra_text,required"`
	// The files of the manual verification.
	Files []VoiceFineTuningManualVerificationFile `json:"files,required"`
	// The date of the manual verification in Unix time.
	RequestTimeUnix int64 `json:"request_time_unix,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraText       respjson.Field
		Files           respjson.Field
		RequestTimeUnix respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceFineTuningManualVerification) RawJSON() string { return r.JSON.raw }
func (r *VoiceFineTuningManualVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceFineTuningManualVerificationFile struct {
	// The ID of the file.
	FileID string `json:"file_id,required"`
	// The name of the file.
	FileName string `json:"file_name,required"`
	// The MIME type of the file.
	MimeType string `json:"mime_type,required"`
	// The size of the file in bytes.
	SizeBytes int64 `json:"size_bytes,required"`
	// The date of the file in Unix time.
	UploadDateUnix int64 `json:"upload_date_unix,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FileID         respjson.Field
		FileName       respjson.Field
		MimeType       respjson.Field
		SizeBytes      respjson.Field
		UploadDateUnix respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceFineTuningManualVerificationFile) RawJSON() string { return r.JSON.raw }
func (r *VoiceFineTuningManualVerificationFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The safety controls of the voice.
type VoiceSafetyControl string

const (
	VoiceSafetyControlNone              VoiceSafetyControl = "NONE"
	VoiceSafetyControlBan               VoiceSafetyControl = "BAN"
	VoiceSafetyControlCaptcha           VoiceSafetyControl = "CAPTCHA"
	VoiceSafetyControlEnterpriseBan     VoiceSafetyControl = "ENTERPRISE_BAN"
	VoiceSafetyControlEnterpriseCaptcha VoiceSafetyControl = "ENTERPRISE_CAPTCHA"
)

// The sharing information of the voice.
type VoiceSharing struct {
	// The category of the voice.
	//
	// Any of "generated", "cloned", "premade", "professional", "famous",
	// "high_quality".
	Category string `json:"category,required"`
	// The number of clones on the voice.
	ClonedByCount int64 `json:"cloned_by_count,required"`
	// The date of the voice sharing in Unix time.
	DateUnix int64 `json:"date_unix,required"`
	// Whether the voice is enabled in the library.
	EnabledInLibrary bool `json:"enabled_in_library,required"`
	// Whether the voice is featured.
	Featured bool `json:"featured,required"`
	// Whether financial rewards are enabled.
	FinancialRewardsEnabled bool `json:"financial_rewards_enabled,required"`
	// Whether free users are allowed.
	FreeUsersAllowed bool `json:"free_users_allowed,required"`
	// The labels of the voice.
	Labels map[string]string `json:"labels,required"`
	// The number of likes on the voice.
	LikedByCount int64 `json:"liked_by_count,required"`
	// Whether live moderation is enabled.
	LiveModerationEnabled bool `json:"live_moderation_enabled,required"`
	// The name of the voice.
	Name string `json:"name,required"`
	// The notice period of the voice sharing.
	NoticePeriod int64 `json:"notice_period,required"`
	// The ID of the original voice.
	OriginalVoiceID string `json:"original_voice_id,required"`
	// The ID of the public owner.
	PublicOwnerID string `json:"public_owner_id,required"`
	// The review status of the voice.
	//
	// Any of "not_requested", "pending", "declined", "allowed",
	// "allowed_with_changes".
	ReviewStatus string `json:"review_status,required"`
	// The status of the voice sharing.
	//
	// Any of "enabled", "disabled", "copied", "copied_disabled".
	Status string `json:"status,required"`
	// Whether voice mixing is allowed.
	VoiceMixingAllowed bool `json:"voice_mixing_allowed,required"`
	// A list of whitelisted emails.
	WhitelistedEmails []string `json:"whitelisted_emails,required"`
	// The ban reason of the voice.
	BanReason string `json:"ban_reason,nullable"`
	// The description of the voice.
	Description string `json:"description,nullable"`
	// The date of the voice sharing in Unix time.
	DisableAtUnix int64 `json:"disable_at_unix,nullable"`
	// The rate of the voice sharing in USD per 1000 credits.
	FiatRate float64 `json:"fiat_rate,nullable"`
	// The sample ID of the history item.
	HistoryItemSampleID string `json:"history_item_sample_id,nullable"`
	// The image URL of the voice.
	ImageURL string `json:"image_url,nullable"`
	// The Instagram username of the voice.
	InstagramUsername string `json:"instagram_username,nullable"`
	// The moderation check of the voice.
	ModerationCheck VoiceSharingModerationCheck `json:"moderation_check,nullable"`
	// The rate of the voice sharing.
	Rate float64 `json:"rate,nullable"`
	// Whether the reader app is enabled.
	ReaderAppEnabled bool `json:"reader_app_enabled,nullable"`
	// The reader restricted on of the voice.
	ReaderRestrictedOn []VoiceSharingReaderRestrictedOn `json:"reader_restricted_on,nullable"`
	// The review message of the voice.
	ReviewMessage string `json:"review_message,nullable"`
	// The TikTok username of the voice.
	TiktokUsername string `json:"tiktok_username,nullable"`
	// The Twitter/X username of the voice.
	TwitterUsername string `json:"twitter_username,nullable"`
	// The YouTube username of the voice.
	YoutubeUsername string `json:"youtube_username,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category                respjson.Field
		ClonedByCount           respjson.Field
		DateUnix                respjson.Field
		EnabledInLibrary        respjson.Field
		Featured                respjson.Field
		FinancialRewardsEnabled respjson.Field
		FreeUsersAllowed        respjson.Field
		Labels                  respjson.Field
		LikedByCount            respjson.Field
		LiveModerationEnabled   respjson.Field
		Name                    respjson.Field
		NoticePeriod            respjson.Field
		OriginalVoiceID         respjson.Field
		PublicOwnerID           respjson.Field
		ReviewStatus            respjson.Field
		Status                  respjson.Field
		VoiceMixingAllowed      respjson.Field
		WhitelistedEmails       respjson.Field
		BanReason               respjson.Field
		Description             respjson.Field
		DisableAtUnix           respjson.Field
		FiatRate                respjson.Field
		HistoryItemSampleID     respjson.Field
		ImageURL                respjson.Field
		InstagramUsername       respjson.Field
		ModerationCheck         respjson.Field
		Rate                    respjson.Field
		ReaderAppEnabled        respjson.Field
		ReaderRestrictedOn      respjson.Field
		ReviewMessage           respjson.Field
		TiktokUsername          respjson.Field
		TwitterUsername         respjson.Field
		YoutubeUsername         respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSharing) RawJSON() string { return r.JSON.raw }
func (r *VoiceSharing) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The moderation check of the voice.
type VoiceSharingModerationCheck struct {
	// A list of CAPTCHA check values.
	CaptchaChecks []float64 `json:"captcha_checks,nullable"`
	// A list of captcha IDs.
	CaptchaIDs []string `json:"captcha_ids,nullable"`
	// The date the moderation check was made in Unix time.
	DateCheckedUnix int64 `json:"date_checked_unix,nullable"`
	// Whether the description check was successful.
	DescriptionCheck bool `json:"description_check,nullable"`
	// The description value of the voice.
	DescriptionValue string `json:"description_value,nullable"`
	// Whether the name check was successful.
	NameCheck bool `json:"name_check,nullable"`
	// The name value of the voice.
	NameValue string `json:"name_value,nullable"`
	// A list of sample checks.
	SampleChecks []float64 `json:"sample_checks,nullable"`
	// A list of sample IDs.
	SampleIDs []string `json:"sample_ids,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CaptchaChecks    respjson.Field
		CaptchaIDs       respjson.Field
		DateCheckedUnix  respjson.Field
		DescriptionCheck respjson.Field
		DescriptionValue respjson.Field
		NameCheck        respjson.Field
		NameValue        respjson.Field
		SampleChecks     respjson.Field
		SampleIDs        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSharingModerationCheck) RawJSON() string { return r.JSON.raw }
func (r *VoiceSharingModerationCheck) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceSharingReaderRestrictedOn struct {
	// The ID of the resource.
	ResourceID string `json:"resource_id,required"`
	// The type of resource.
	//
	// Any of "read", "collection".
	ResourceType string `json:"resource_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID   respjson.Field
		ResourceType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceSharingReaderRestrictedOn) RawJSON() string { return r.JSON.raw }
func (r *VoiceSharingReaderRestrictedOn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The voice verification of the voice.
type VoiceVoiceVerification struct {
	// Whether the voice has been verified.
	IsVerified bool `json:"is_verified,required"`
	// Whether the voice requires verification.
	RequiresVerification bool `json:"requires_verification,required"`
	// The number of verification attempts.
	VerificationAttemptsCount int64 `json:"verification_attempts_count,required"`
	// List of verification failures.
	VerificationFailures []string `json:"verification_failures,required"`
	// The language of the voice.
	Language string `json:"language,nullable"`
	// Number of times a verification was attempted.
	VerificationAttempts []VerificationAttempt `json:"verification_attempts,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsVerified                respjson.Field
		RequiresVerification      respjson.Field
		VerificationAttemptsCount respjson.Field
		VerificationFailures      respjson.Field
		Language                  respjson.Field
		VerificationAttempts      respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VoiceVoiceVerification) RawJSON() string { return r.JSON.raw }
func (r *VoiceVoiceVerification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VoiceGenerationNewVoiceParams struct {
	// The generated_voice_id to create, call POST /v1/text-to-voice/create-previews
	// and fetch the generated_voice_id from the response header if don't have one yet.
	GeneratedVoiceID string `json:"generated_voice_id,required"`
	// Description to use for the created voice.
	VoiceDescription string `json:"voice_description,required"`
	// Name to use for the created voice.
	VoiceName string `json:"voice_name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Optional, metadata to add to the created voice. Defaults to None.
	Labels map[string]string `json:"labels,omitzero"`
	// List of voice ids that the user has played but not selected. Used for RLHF.
	PlayedNotSelectedVoiceIDs []string `json:"played_not_selected_voice_ids,omitzero"`
	paramObj
}

func (r VoiceGenerationNewVoiceParams) MarshalJSON() (data []byte, err error) {
	type shadow VoiceGenerationNewVoiceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VoiceGenerationNewVoiceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
