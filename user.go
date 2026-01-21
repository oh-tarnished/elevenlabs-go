// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Gets information about the user
func (r *UserService) GetInfo(ctx context.Context, query UserGetInfoParams, opts ...option.RequestOption) (res *UserGetInfoResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/user"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets extended information about the users subscription
func (r *UserService) GetSubscription(ctx context.Context, query UserGetSubscriptionParams, opts ...option.RequestOption) (res *UserGetSubscriptionResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/user/subscription"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type BillingPeriod string

const (
	BillingPeriodMonthlyPeriod BillingPeriod = "monthly_period"
	BillingPeriod3MonthPeriod  BillingPeriod = "3_month_period"
	BillingPeriod6MonthPeriod  BillingPeriod = "6_month_period"
	BillingPeriodAnnualPeriod  BillingPeriod = "annual_period"
)

type CharacterRefreshPeriod string

const (
	CharacterRefreshPeriodMonthlyPeriod CharacterRefreshPeriod = "monthly_period"
	CharacterRefreshPeriod3MonthPeriod  CharacterRefreshPeriod = "3_month_period"
	CharacterRefreshPeriod6MonthPeriod  CharacterRefreshPeriod = "6_month_period"
	CharacterRefreshPeriodAnnualPeriod  CharacterRefreshPeriod = "annual_period"
)

type Invoice struct {
	// The amount due in cents.
	AmountDueCents int64 `json:"amount_due_cents,required"`
	// The discounts applied to the invoice.
	Discounts []InvoiceDiscount `json:"discounts,required"`
	// The Unix timestamp of the next payment attempt.
	NextPaymentAttemptUnix int64 `json:"next_payment_attempt_unix,required"`
	// Deprecated. Use [payment_intent_statusses] instead. The status of this invoice's
	// first payment intent. None when there is no payment intent.
	//
	// Any of "canceled", "processing", "requires_action", "requires_capture",
	// "requires_confirmation", "requires_payment_method", "succeeded".
	PaymentIntentStatus InvoicePaymentIntentStatus `json:"payment_intent_status,required"`
	// The statuses of this invoice's payment intents. Empty list when there are no
	// payment intents.
	//
	// Any of "canceled", "processing", "requires_action", "requires_capture",
	// "requires_confirmation", "requires_payment_method", "succeeded".
	PaymentIntentStatusses []string `json:"payment_intent_statusses,required"`
	// Deprecated. Use [discounts] instead. The discount applied to the invoice. E.g.
	// [20.0f] for 20 cents off.
	DiscountAmountOff float64 `json:"discount_amount_off,nullable"`
	// Deprecated. Use [discounts] instead. The discount applied to the invoice. E.g.
	// [20.0f] for 20% off.
	DiscountPercentOff float64 `json:"discount_percent_off,nullable"`
	// The subtotal amount in cents before tax (exclusive of tax and discounts).
	SubtotalCents int64 `json:"subtotal_cents,nullable"`
	// The tax amount in cents.
	TaxCents int64 `json:"tax_cents,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AmountDueCents         respjson.Field
		Discounts              respjson.Field
		NextPaymentAttemptUnix respjson.Field
		PaymentIntentStatus    respjson.Field
		PaymentIntentStatusses respjson.Field
		DiscountAmountOff      respjson.Field
		DiscountPercentOff     respjson.Field
		SubtotalCents          respjson.Field
		TaxCents               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Invoice) RawJSON() string { return r.JSON.raw }
func (r *Invoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InvoiceDiscount struct {
	// The discount applied to the invoice. E.g. [20.0f] for 20 cents off.
	DiscountAmountOff float64 `json:"discount_amount_off,nullable"`
	// The discount applied to the invoice. E.g. [20.0f] for 20% off.
	DiscountPercentOff float64 `json:"discount_percent_off,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiscountAmountOff  respjson.Field
		DiscountPercentOff respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InvoiceDiscount) RawJSON() string { return r.JSON.raw }
func (r *InvoiceDiscount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Deprecated. Use [payment_intent_statusses] instead. The status of this invoice's
// first payment intent. None when there is no payment intent.
type InvoicePaymentIntentStatus string

const (
	InvoicePaymentIntentStatusCanceled              InvoicePaymentIntentStatus = "canceled"
	InvoicePaymentIntentStatusProcessing            InvoicePaymentIntentStatus = "processing"
	InvoicePaymentIntentStatusRequiresAction        InvoicePaymentIntentStatus = "requires_action"
	InvoicePaymentIntentStatusRequiresCapture       InvoicePaymentIntentStatus = "requires_capture"
	InvoicePaymentIntentStatusRequiresConfirmation  InvoicePaymentIntentStatus = "requires_confirmation"
	InvoicePaymentIntentStatusRequiresPaymentMethod InvoicePaymentIntentStatus = "requires_payment_method"
	InvoicePaymentIntentStatusSucceeded             InvoicePaymentIntentStatus = "succeeded"
)

type SubscriptionStatus string

const (
	SubscriptionStatusTrialing     SubscriptionStatus = "trialing"
	SubscriptionStatusActive       SubscriptionStatus = "active"
	SubscriptionStatusIncomplete   SubscriptionStatus = "incomplete"
	SubscriptionStatusPastDue      SubscriptionStatus = "past_due"
	SubscriptionStatusFree         SubscriptionStatus = "free"
	SubscriptionStatusFreeDisabled SubscriptionStatus = "free_disabled"
)

type UserGetInfoResponse struct {
	// This field is deprecated and will be removed in a future major version. Instead
	// use subscription.trust_on_invoice_creation.
	//
	// Deprecated: deprecated
	CanUseDelayedPaymentMethods bool `json:"can_use_delayed_payment_methods,required"`
	// The unix timestamp of the user's creation. 0 if the user was created before the
	// unix timestamp was added.
	CreatedAt int64 `json:"created_at,required"`
	// Whether the user is new. This field is deprecated and will be removed in the
	// future. Use 'created_at' instead.
	//
	// Deprecated: deprecated
	IsNewUser bool `json:"is_new_user,required"`
	// Whether the user's onboarding checklist is completed.
	IsOnboardingChecklistCompleted bool `json:"is_onboarding_checklist_completed,required"`
	// Whether the user's onboarding is completed.
	IsOnboardingCompleted bool `json:"is_onboarding_completed,required"`
	// Details of the user's subscription.
	Subscription UserGetInfoResponseSubscription `json:"subscription,required"`
	// The unique identifier of the user.
	UserID string `json:"user_id,required"`
	// First name of the user.
	FirstName string `json:"first_name,nullable"`
	// Whether the user's API key is hashed.
	IsAPIKeyHashed bool `json:"is_api_key_hashed"`
	// The Partnerstack partner default link of the user.
	PartnerstackPartnerDefaultLink string `json:"partnerstack_partner_default_link,nullable"`
	// The referral link code of the user.
	ReferralLinkCode string `json:"referral_link_code,nullable"`
	// The API key of the user.
	XiAPIKey string `json:"xi_api_key,nullable"`
	// The preview of the user's API key.
	XiAPIKeyPreview string `json:"xi_api_key_preview,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanUseDelayedPaymentMethods    respjson.Field
		CreatedAt                      respjson.Field
		IsNewUser                      respjson.Field
		IsOnboardingChecklistCompleted respjson.Field
		IsOnboardingCompleted          respjson.Field
		Subscription                   respjson.Field
		UserID                         respjson.Field
		FirstName                      respjson.Field
		IsAPIKeyHashed                 respjson.Field
		PartnerstackPartnerDefaultLink respjson.Field
		ReferralLinkCode               respjson.Field
		XiAPIKey                       respjson.Field
		XiAPIKeyPreview                respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetInfoResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetInfoResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the user's subscription.
type UserGetInfoResponseSubscription struct {
	// Whether the user is allowed to extend their character limit.
	AllowedToExtendCharacterLimit bool `json:"allowed_to_extend_character_limit,required"`
	// Whether the user can extend their character limit.
	CanExtendCharacterLimit bool `json:"can_extend_character_limit,required"`
	// Whether the user can extend their voice limit.
	CanExtendVoiceLimit bool `json:"can_extend_voice_limit,required"`
	// Whether the user can use instant voice cloning.
	CanUseInstantVoiceCloning bool `json:"can_use_instant_voice_cloning,required"`
	// Whether the user can use professional voice cloning.
	CanUseProfessionalVoiceCloning bool `json:"can_use_professional_voice_cloning,required"`
	// The number of characters used by the user.
	CharacterCount int64 `json:"character_count,required"`
	// The maximum number of characters allowed in the current billing period.
	CharacterLimit int64 `json:"character_limit,required"`
	// Maximum number of characters that the character limit can be exceeded by.
	// Managed by the workspace admin.
	MaxCharacterLimitExtension int64 `json:"max_character_limit_extension,required"`
	// The maximum number of professional voices allowed for the user.
	ProfessionalVoiceLimit int64 `json:"professional_voice_limit,required"`
	// The number of professional voice slots used by the workspace/user if single
	// seat.
	ProfessionalVoiceSlotsUsed int64 `json:"professional_voice_slots_used,required"`
	// The status of the user's subscription.
	//
	// Any of "trialing", "active", "incomplete", "past_due", "free", "free_disabled".
	Status SubscriptionStatus `json:"status,required"`
	// The tier of the user's subscription.
	Tier string `json:"tier,required"`
	// The number of voice add/edits used by the user.
	VoiceAddEditCounter int64 `json:"voice_add_edit_counter,required"`
	// The maximum number of voice slots allowed for the user.
	VoiceLimit int64 `json:"voice_limit,required"`
	// The number of voice slots used by the user.
	VoiceSlotsUsed int64 `json:"voice_slots_used,required"`
	// The billing period of the user's subscription.
	//
	// Any of "monthly_period", "3_month_period", "6_month_period", "annual_period".
	BillingPeriod BillingPeriod `json:"billing_period,nullable"`
	// The character refresh period of the user's subscription.
	//
	// Any of "monthly_period", "3_month_period", "6_month_period", "annual_period".
	CharacterRefreshPeriod CharacterRefreshPeriod `json:"character_refresh_period,nullable"`
	// The currency of the user's subscription.
	//
	// Any of "usd", "eur", "inr".
	Currency string `json:"currency,nullable"`
	// The maximum number of voice add/edits allowed for the user.
	MaxVoiceAddEdits int64 `json:"max_voice_add_edits,nullable"`
	// The Unix timestamp of the next character count reset.
	NextCharacterCountResetUnix int64 `json:"next_character_count_reset_unix,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedToExtendCharacterLimit  respjson.Field
		CanExtendCharacterLimit        respjson.Field
		CanExtendVoiceLimit            respjson.Field
		CanUseInstantVoiceCloning      respjson.Field
		CanUseProfessionalVoiceCloning respjson.Field
		CharacterCount                 respjson.Field
		CharacterLimit                 respjson.Field
		MaxCharacterLimitExtension     respjson.Field
		ProfessionalVoiceLimit         respjson.Field
		ProfessionalVoiceSlotsUsed     respjson.Field
		Status                         respjson.Field
		Tier                           respjson.Field
		VoiceAddEditCounter            respjson.Field
		VoiceLimit                     respjson.Field
		VoiceSlotsUsed                 respjson.Field
		BillingPeriod                  respjson.Field
		CharacterRefreshPeriod         respjson.Field
		Currency                       respjson.Field
		MaxVoiceAddEdits               respjson.Field
		NextCharacterCountResetUnix    respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetInfoResponseSubscription) RawJSON() string { return r.JSON.raw }
func (r *UserGetInfoResponseSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetSubscriptionResponse struct {
	// Whether the user is allowed to extend their character limit.
	AllowedToExtendCharacterLimit bool `json:"allowed_to_extend_character_limit,required"`
	// Whether the user can extend their character limit.
	CanExtendCharacterLimit bool `json:"can_extend_character_limit,required"`
	// Whether the user can extend their voice limit.
	CanExtendVoiceLimit bool `json:"can_extend_voice_limit,required"`
	// Whether the user can use instant voice cloning.
	CanUseInstantVoiceCloning bool `json:"can_use_instant_voice_cloning,required"`
	// Whether the user can use professional voice cloning.
	CanUseProfessionalVoiceCloning bool `json:"can_use_professional_voice_cloning,required"`
	// The number of characters used by the user.
	CharacterCount int64 `json:"character_count,required"`
	// The maximum number of characters allowed in the current billing period.
	CharacterLimit int64 `json:"character_limit,required"`
	// Whether the user has open invoices.
	HasOpenInvoices bool `json:"has_open_invoices,required"`
	// Maximum number of characters that the character limit can be exceeded by.
	// Managed by the workspace admin.
	MaxCharacterLimitExtension int64 `json:"max_character_limit_extension,required"`
	// The open invoices for the user.
	OpenInvoices []Invoice `json:"open_invoices,required"`
	// The maximum number of professional voices allowed for the user.
	ProfessionalVoiceLimit int64 `json:"professional_voice_limit,required"`
	// The number of professional voice slots used by the workspace/user if single
	// seat.
	ProfessionalVoiceSlotsUsed int64 `json:"professional_voice_slots_used,required"`
	// The status of the user's subscription.
	//
	// Any of "trialing", "active", "incomplete", "past_due", "free", "free_disabled".
	Status SubscriptionStatus `json:"status,required"`
	// The tier of the user's subscription.
	Tier string `json:"tier,required"`
	// The number of voice add/edits used by the user.
	VoiceAddEditCounter int64 `json:"voice_add_edit_counter,required"`
	// The maximum number of voice slots allowed for the user.
	VoiceLimit int64 `json:"voice_limit,required"`
	// The number of voice slots used by the user.
	VoiceSlotsUsed int64 `json:"voice_slots_used,required"`
	// The billing period of the user's subscription.
	//
	// Any of "monthly_period", "3_month_period", "6_month_period", "annual_period".
	BillingPeriod BillingPeriod `json:"billing_period,nullable"`
	// The character refresh period of the user's subscription.
	//
	// Any of "monthly_period", "3_month_period", "6_month_period", "annual_period".
	CharacterRefreshPeriod CharacterRefreshPeriod `json:"character_refresh_period,nullable"`
	// The currency of the user's subscription.
	//
	// Any of "usd", "eur", "inr".
	Currency UserGetSubscriptionResponseCurrency `json:"currency,nullable"`
	// The maximum number of voice add/edits allowed for the user.
	MaxVoiceAddEdits int64 `json:"max_voice_add_edits,nullable"`
	// The Unix timestamp of the next character count reset.
	NextCharacterCountResetUnix int64 `json:"next_character_count_reset_unix,nullable"`
	// The next invoice for the user.
	NextInvoice Invoice `json:"next_invoice,nullable"`
	// The pending change for the user.
	PendingChange UserGetSubscriptionResponsePendingChangeUnion `json:"pending_change,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedToExtendCharacterLimit  respjson.Field
		CanExtendCharacterLimit        respjson.Field
		CanExtendVoiceLimit            respjson.Field
		CanUseInstantVoiceCloning      respjson.Field
		CanUseProfessionalVoiceCloning respjson.Field
		CharacterCount                 respjson.Field
		CharacterLimit                 respjson.Field
		HasOpenInvoices                respjson.Field
		MaxCharacterLimitExtension     respjson.Field
		OpenInvoices                   respjson.Field
		ProfessionalVoiceLimit         respjson.Field
		ProfessionalVoiceSlotsUsed     respjson.Field
		Status                         respjson.Field
		Tier                           respjson.Field
		VoiceAddEditCounter            respjson.Field
		VoiceLimit                     respjson.Field
		VoiceSlotsUsed                 respjson.Field
		BillingPeriod                  respjson.Field
		CharacterRefreshPeriod         respjson.Field
		Currency                       respjson.Field
		MaxVoiceAddEdits               respjson.Field
		NextCharacterCountResetUnix    respjson.Field
		NextInvoice                    respjson.Field
		PendingChange                  respjson.Field
		ExtraFields                    map[string]respjson.Field
		raw                            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetSubscriptionResponse) RawJSON() string { return r.JSON.raw }
func (r *UserGetSubscriptionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The currency of the user's subscription.
type UserGetSubscriptionResponseCurrency string

const (
	UserGetSubscriptionResponseCurrencyUsd UserGetSubscriptionResponseCurrency = "usd"
	UserGetSubscriptionResponseCurrencyEur UserGetSubscriptionResponseCurrency = "eur"
	UserGetSubscriptionResponseCurrencyInr UserGetSubscriptionResponseCurrency = "inr"
)

// UserGetSubscriptionResponsePendingChangeUnion contains all possible properties
// and values from
// [UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel],
// [UserGetSubscriptionResponsePendingChangePendingCancellationResponseModel].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type UserGetSubscriptionResponsePendingChangeUnion struct {
	// This field is from variant
	// [UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel].
	NextBillingPeriod BillingPeriod `json:"next_billing_period"`
	// This field is from variant
	// [UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel].
	NextTier         string `json:"next_tier"`
	TimestampSeconds int64  `json:"timestamp_seconds"`
	Kind             string `json:"kind"`
	JSON             struct {
		NextBillingPeriod respjson.Field
		NextTier          respjson.Field
		TimestampSeconds  respjson.Field
		Kind              respjson.Field
		raw               string
	} `json:"-"`
}

func (u UserGetSubscriptionResponsePendingChangeUnion) AsPendingSubscriptionSwitchResponseModel() (v UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u UserGetSubscriptionResponsePendingChangeUnion) AsPendingCancellationResponseModel() (v UserGetSubscriptionResponsePendingChangePendingCancellationResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u UserGetSubscriptionResponsePendingChangeUnion) RawJSON() string { return u.JSON.raw }

func (r *UserGetSubscriptionResponsePendingChangeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel struct {
	// The billing period to change to.
	//
	// Any of "monthly_period", "3_month_period", "6_month_period", "annual_period".
	NextBillingPeriod BillingPeriod `json:"next_billing_period,required"`
	// The tier to change to.
	//
	// Any of "free", "starter", "creator", "pro", "growing_business",
	// "scale_2024_08_10", "grant_tier_1_2025_07_23", "grant_tier_2_2025_07_23",
	// "trial", "enterprise".
	NextTier string `json:"next_tier,required"`
	// The timestamp of the change.
	TimestampSeconds int64 `json:"timestamp_seconds,required"`
	// Any of "change".
	Kind string `json:"kind"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextBillingPeriod respjson.Field
		NextTier          respjson.Field
		TimestampSeconds  respjson.Field
		Kind              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *UserGetSubscriptionResponsePendingChangePendingSubscriptionSwitchResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetSubscriptionResponsePendingChangePendingCancellationResponseModel struct {
	// The timestamp of the cancellation.
	TimestampSeconds int64 `json:"timestamp_seconds,required"`
	// Any of "cancellation".
	Kind string `json:"kind"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TimestampSeconds respjson.Field
		Kind             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UserGetSubscriptionResponsePendingChangePendingCancellationResponseModel) RawJSON() string {
	return r.JSON.raw
}
func (r *UserGetSubscriptionResponsePendingChangePendingCancellationResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserGetInfoParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type UserGetSubscriptionParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
