// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiPhoneNumberService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiPhoneNumberService] method instead.
type ConvaiPhoneNumberService struct {
	Options []option.RequestOption
}

// NewConvaiPhoneNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiPhoneNumberService(opts ...option.RequestOption) (r ConvaiPhoneNumberService) {
	r = ConvaiPhoneNumberService{}
	r.Options = opts
	return
}

// Retrieve Phone Number details by ID
func (r *ConvaiPhoneNumberService) Get(ctx context.Context, phoneNumberID string, query ConvaiPhoneNumberGetParams, opts ...option.RequestOption) (res *ConvaiPhoneNumberGetResponseUnion, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/phone-numbers/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update assigned agent of a phone number
func (r *ConvaiPhoneNumberService) Update(ctx context.Context, phoneNumberID string, params ConvaiPhoneNumberUpdateParams, opts ...option.RequestOption) (res *ConvaiPhoneNumberUpdateResponseUnion, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/phone-numbers/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Retrieve all Phone Numbers
func (r *ConvaiPhoneNumberService) List(ctx context.Context, query ConvaiPhoneNumberListParams, opts ...option.RequestOption) (res *[]ConvaiPhoneNumberListResponseUnion, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/phone-numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete Phone Number by ID
func (r *ConvaiPhoneNumberService) Delete(ctx context.Context, phoneNumberID string, body ConvaiPhoneNumberDeleteParams, opts ...option.RequestOption) (res *ConvaiPhoneNumberDeleteResponse, err error) {
	if !param.IsOmitted(body.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", body.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if phoneNumberID == "" {
		err = errors.New("missing required phone_number_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/phone-numbers/%s", phoneNumberID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Import Phone Number from provider configuration (Twilio or SIP trunk)
func (r *ConvaiPhoneNumberService) Import(ctx context.Context, params ConvaiPhoneNumberImportParams, opts ...option.RequestOption) (res *ConvaiPhoneNumberImportResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/phone-numbers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// SIP Trunk configuration details for a phone number
type GetPhoneNumberOutboundSipTrunkConfigResponseModel struct {
	// Hostname or IP the SIP INVITE is sent to
	Address string `json:"address,required"`
	// Whether authentication credentials are configured
	HasAuthCredentials bool `json:"has_auth_credentials,required"`
	// Whether or not to encrypt media (data layer).
	//
	// Any of "disabled", "allowed", "required".
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption,required"`
	// Protocol to use for SIP transport
	//
	// Any of "auto", "udp", "tcp", "tls".
	Transport SipTrunkTransportEnum `json:"transport,required"`
	// Whether a LiveKit SIP outbound trunk is configured
	HasOutboundTrunk bool `json:"has_outbound_trunk"`
	// SIP headers for INVITE request
	Headers map[string]string `json:"headers"`
	// SIP trunk username (if available)
	Username string `json:"username,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address            respjson.Field
		HasAuthCredentials respjson.Field
		MediaEncryption    respjson.Field
		Transport          respjson.Field
		HasOutboundTrunk   respjson.Field
		Headers            respjson.Field
		Username           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPhoneNumberOutboundSipTrunkConfigResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GetPhoneNumberOutboundSipTrunkConfigResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetPhoneNumberSipTrunkResponseModel struct {
	// Label for the phone number
	Label string `json:"label,required"`
	// Type of Livekit stack used for this number.
	//
	// Any of "standard", "static".
	LivekitStack LivekitStackType `json:"livekit_stack,required"`
	// Phone number
	PhoneNumber string `json:"phone_number,required"`
	// The ID of the phone number
	PhoneNumberID string `json:"phone_number_id,required"`
	// The agent that is assigned to the phone number
	AssignedAgent PhoneNumberAgentInfo `json:"assigned_agent,nullable"`
	// Configuration of the Inbound SIP trunk - if configured.
	InboundTrunk GetPhoneNumberSipTrunkResponseModelInboundTrunk `json:"inbound_trunk,nullable"`
	// SIP Trunk configuration details for a phone number
	OutboundTrunk GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"outbound_trunk,nullable"`
	// Phone provider
	//
	// Any of "sip_trunk".
	Provider GetPhoneNumberSipTrunkResponseModelProvider `json:"provider"`
	// SIP Trunk configuration details for a phone number
	//
	// Deprecated: deprecated
	ProviderConfig GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"provider_config,nullable"`
	// Whether this phone number supports inbound calls
	SupportsInbound bool `json:"supports_inbound"`
	// Whether this phone number supports outbound calls
	SupportsOutbound bool `json:"supports_outbound"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label            respjson.Field
		LivekitStack     respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		AssignedAgent    respjson.Field
		InboundTrunk     respjson.Field
		OutboundTrunk    respjson.Field
		Provider         respjson.Field
		ProviderConfig   respjson.Field
		SupportsInbound  respjson.Field
		SupportsOutbound respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPhoneNumberSipTrunkResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GetPhoneNumberSipTrunkResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration of the Inbound SIP trunk - if configured.
type GetPhoneNumberSipTrunkResponseModelInboundTrunk struct {
	// List of IP addresses that are allowed to use the trunk. Each item in the list
	// can be an individual IP address or a Classless Inter-Domain Routing notation
	// representing a CIDR block.
	AllowedAddresses []string `json:"allowed_addresses,required"`
	// List of phone numbers that are allowed to use the trunk.
	AllowedNumbers []string `json:"allowed_numbers,required"`
	// Whether authentication credentials are configured
	HasAuthCredentials bool `json:"has_auth_credentials,required"`
	// Any of "disabled", "allowed", "required".
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption,required"`
	// Domains of remote SIP servers used to validate TLS certificates.
	RemoteDomains []string `json:"remote_domains,nullable"`
	// SIP trunk username (if available)
	Username string `json:"username,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedAddresses   respjson.Field
		AllowedNumbers     respjson.Field
		HasAuthCredentials respjson.Field
		MediaEncryption    respjson.Field
		RemoteDomains      respjson.Field
		Username           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPhoneNumberSipTrunkResponseModelInboundTrunk) RawJSON() string { return r.JSON.raw }
func (r *GetPhoneNumberSipTrunkResponseModelInboundTrunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Phone provider
type GetPhoneNumberSipTrunkResponseModelProvider string

const (
	GetPhoneNumberSipTrunkResponseModelProviderSipTrunk GetPhoneNumberSipTrunkResponseModelProvider = "sip_trunk"
)

type GetPhoneNumberTwilioResponseModel struct {
	// Label for the phone number
	Label string `json:"label,required"`
	// Phone number
	PhoneNumber string `json:"phone_number,required"`
	// The ID of the phone number
	PhoneNumberID string `json:"phone_number_id,required"`
	// The agent that is assigned to the phone number
	AssignedAgent PhoneNumberAgentInfo `json:"assigned_agent,nullable"`
	// Phone provider
	//
	// Any of "twilio".
	Provider GetPhoneNumberTwilioResponseModelProvider `json:"provider"`
	// Whether this phone number supports inbound calls
	SupportsInbound bool `json:"supports_inbound"`
	// Whether this phone number supports outbound calls
	SupportsOutbound bool `json:"supports_outbound"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Label            respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		AssignedAgent    respjson.Field
		Provider         respjson.Field
		SupportsInbound  respjson.Field
		SupportsOutbound respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPhoneNumberTwilioResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GetPhoneNumberTwilioResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Phone provider
type GetPhoneNumberTwilioResponseModelProvider string

const (
	GetPhoneNumberTwilioResponseModelProviderTwilio GetPhoneNumberTwilioResponseModelProvider = "twilio"
)

type InboundSipTrunkConfigRequestModelParam struct {
	// List of IP addresses that are allowed to use the trunk. Each item in the list
	// can be an individual IP address or a Classless Inter-Domain Routing notation
	// representing a CIDR block.
	AllowedAddresses []string `json:"allowed_addresses,omitzero"`
	// List of phone numbers that are allowed to use the trunk.
	AllowedNumbers []string `json:"allowed_numbers,omitzero"`
	// Domains of remote SIP servers used to validate TLS certificates.
	RemoteDomains []string `json:"remote_domains,omitzero"`
	// Optional digest authentication credentials (username/password).
	Credentials SipTrunkCredentialsRequestModelParam `json:"credentials,omitzero"`
	// Whether or not to encrypt media (data layer).
	//
	// Any of "disabled", "allowed", "required".
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption,omitzero"`
	paramObj
}

func (r InboundSipTrunkConfigRequestModelParam) MarshalJSON() (data []byte, err error) {
	type shadow InboundSipTrunkConfigRequestModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboundSipTrunkConfigRequestModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Address is required.
type OutboundSipTrunkConfigRequestModelParam struct {
	// Hostname or IP the SIP INVITE is sent to.
	Address string `json:"address,required"`
	// Optional digest authentication credentials (username/password). If not provided,
	// ACL authentication is assumed.
	Credentials SipTrunkCredentialsRequestModelParam `json:"credentials,omitzero"`
	// SIP X-\* headers for INVITE request. These headers are sent as-is and may help
	// identify this call.
	Headers map[string]string `json:"headers,omitzero"`
	// Whether or not to encrypt media (data layer).
	//
	// Any of "disabled", "allowed", "required".
	MediaEncryption SipMediaEncryptionEnum `json:"media_encryption,omitzero"`
	// Protocol to use for SIP transport (signalling layer).
	//
	// Any of "auto", "udp", "tcp", "tls".
	Transport SipTrunkTransportEnum `json:"transport,omitzero"`
	paramObj
}

func (r OutboundSipTrunkConfigRequestModelParam) MarshalJSON() (data []byte, err error) {
	type shadow OutboundSipTrunkConfigRequestModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OutboundSipTrunkConfigRequestModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberAgentInfo struct {
	// The ID of the agent
	AgentID string `json:"agent_id,required"`
	// The name of the agent
	AgentName string `json:"agent_name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentID     respjson.Field
		AgentName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberAgentInfo) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberAgentInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipMediaEncryptionEnum string

const (
	SipMediaEncryptionEnumDisabled SipMediaEncryptionEnum = "disabled"
	SipMediaEncryptionEnumAllowed  SipMediaEncryptionEnum = "allowed"
	SipMediaEncryptionEnumRequired SipMediaEncryptionEnum = "required"
)

// The property Username is required.
type SipTrunkCredentialsRequestModelParam struct {
	// SIP trunk username
	Username string `json:"username,required"`
	// SIP trunk password - if not specified, then remain unchanged
	Password param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r SipTrunkCredentialsRequestModelParam) MarshalJSON() (data []byte, err error) {
	type shadow SipTrunkCredentialsRequestModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SipTrunkCredentialsRequestModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SipTrunkTransportEnum string

const (
	SipTrunkTransportEnumAuto SipTrunkTransportEnum = "auto"
	SipTrunkTransportEnumUdp  SipTrunkTransportEnum = "udp"
	SipTrunkTransportEnumTcp  SipTrunkTransportEnum = "tcp"
	SipTrunkTransportEnumTls  SipTrunkTransportEnum = "tls"
)

// ConvaiPhoneNumberGetResponseUnion contains all possible properties and values
// from [GetPhoneNumberTwilioResponseModel], [GetPhoneNumberSipTrunkResponseModel].
//
// Use the [ConvaiPhoneNumberGetResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiPhoneNumberGetResponseUnion struct {
	Label         string `json:"label"`
	PhoneNumber   string `json:"phone_number"`
	PhoneNumberID string `json:"phone_number_id"`
	// This field is from variant [GetPhoneNumberTwilioResponseModel].
	AssignedAgent PhoneNumberAgentInfo `json:"assigned_agent"`
	// Any of "twilio", "sip_trunk".
	Provider         string `json:"provider"`
	SupportsInbound  bool   `json:"supports_inbound"`
	SupportsOutbound bool   `json:"supports_outbound"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	LivekitStack LivekitStackType `json:"livekit_stack"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	InboundTrunk GetPhoneNumberSipTrunkResponseModelInboundTrunk `json:"inbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	OutboundTrunk GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"outbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	ProviderConfig GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"provider_config"`
	JSON           struct {
		Label            respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		AssignedAgent    respjson.Field
		Provider         respjson.Field
		SupportsInbound  respjson.Field
		SupportsOutbound respjson.Field
		LivekitStack     respjson.Field
		InboundTrunk     respjson.Field
		OutboundTrunk    respjson.Field
		ProviderConfig   respjson.Field
		raw              string
	} `json:"-"`
}

// anyConvaiPhoneNumberGetResponse is implemented by each variant of
// [ConvaiPhoneNumberGetResponseUnion] to add type safety for the return type of
// [ConvaiPhoneNumberGetResponseUnion.AsAny]
type anyConvaiPhoneNumberGetResponse interface {
	implConvaiPhoneNumberGetResponseUnion()
}

func (GetPhoneNumberTwilioResponseModel) implConvaiPhoneNumberGetResponseUnion()   {}
func (GetPhoneNumberSipTrunkResponseModel) implConvaiPhoneNumberGetResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiPhoneNumberGetResponseUnion.AsAny().(type) {
//	case elevenlabs.GetPhoneNumberTwilioResponseModel:
//	case elevenlabs.GetPhoneNumberSipTrunkResponseModel:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiPhoneNumberGetResponseUnion) AsAny() anyConvaiPhoneNumberGetResponse {
	switch u.Provider {
	case "twilio":
		return u.AsTwilio()
	case "sip_trunk":
		return u.AsSipTrunk()
	}
	return nil
}

func (u ConvaiPhoneNumberGetResponseUnion) AsTwilio() (v GetPhoneNumberTwilioResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiPhoneNumberGetResponseUnion) AsSipTrunk() (v GetPhoneNumberSipTrunkResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiPhoneNumberGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiPhoneNumberGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiPhoneNumberUpdateResponseUnion contains all possible properties and values
// from [GetPhoneNumberTwilioResponseModel], [GetPhoneNumberSipTrunkResponseModel].
//
// Use the [ConvaiPhoneNumberUpdateResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiPhoneNumberUpdateResponseUnion struct {
	Label         string `json:"label"`
	PhoneNumber   string `json:"phone_number"`
	PhoneNumberID string `json:"phone_number_id"`
	// This field is from variant [GetPhoneNumberTwilioResponseModel].
	AssignedAgent PhoneNumberAgentInfo `json:"assigned_agent"`
	// Any of "twilio", "sip_trunk".
	Provider         string `json:"provider"`
	SupportsInbound  bool   `json:"supports_inbound"`
	SupportsOutbound bool   `json:"supports_outbound"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	LivekitStack LivekitStackType `json:"livekit_stack"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	InboundTrunk GetPhoneNumberSipTrunkResponseModelInboundTrunk `json:"inbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	OutboundTrunk GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"outbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	ProviderConfig GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"provider_config"`
	JSON           struct {
		Label            respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		AssignedAgent    respjson.Field
		Provider         respjson.Field
		SupportsInbound  respjson.Field
		SupportsOutbound respjson.Field
		LivekitStack     respjson.Field
		InboundTrunk     respjson.Field
		OutboundTrunk    respjson.Field
		ProviderConfig   respjson.Field
		raw              string
	} `json:"-"`
}

// anyConvaiPhoneNumberUpdateResponse is implemented by each variant of
// [ConvaiPhoneNumberUpdateResponseUnion] to add type safety for the return type of
// [ConvaiPhoneNumberUpdateResponseUnion.AsAny]
type anyConvaiPhoneNumberUpdateResponse interface {
	implConvaiPhoneNumberUpdateResponseUnion()
}

func (GetPhoneNumberTwilioResponseModel) implConvaiPhoneNumberUpdateResponseUnion()   {}
func (GetPhoneNumberSipTrunkResponseModel) implConvaiPhoneNumberUpdateResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiPhoneNumberUpdateResponseUnion.AsAny().(type) {
//	case elevenlabs.GetPhoneNumberTwilioResponseModel:
//	case elevenlabs.GetPhoneNumberSipTrunkResponseModel:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiPhoneNumberUpdateResponseUnion) AsAny() anyConvaiPhoneNumberUpdateResponse {
	switch u.Provider {
	case "twilio":
		return u.AsTwilio()
	case "sip_trunk":
		return u.AsSipTrunk()
	}
	return nil
}

func (u ConvaiPhoneNumberUpdateResponseUnion) AsTwilio() (v GetPhoneNumberTwilioResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiPhoneNumberUpdateResponseUnion) AsSipTrunk() (v GetPhoneNumberSipTrunkResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiPhoneNumberUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiPhoneNumberUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiPhoneNumberListResponseUnion contains all possible properties and values
// from [GetPhoneNumberTwilioResponseModel], [GetPhoneNumberSipTrunkResponseModel].
//
// Use the [ConvaiPhoneNumberListResponseUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiPhoneNumberListResponseUnion struct {
	Label         string `json:"label"`
	PhoneNumber   string `json:"phone_number"`
	PhoneNumberID string `json:"phone_number_id"`
	// This field is from variant [GetPhoneNumberTwilioResponseModel].
	AssignedAgent PhoneNumberAgentInfo `json:"assigned_agent"`
	// Any of "twilio", "sip_trunk".
	Provider         string `json:"provider"`
	SupportsInbound  bool   `json:"supports_inbound"`
	SupportsOutbound bool   `json:"supports_outbound"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	LivekitStack LivekitStackType `json:"livekit_stack"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	InboundTrunk GetPhoneNumberSipTrunkResponseModelInboundTrunk `json:"inbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	OutboundTrunk GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"outbound_trunk"`
	// This field is from variant [GetPhoneNumberSipTrunkResponseModel].
	ProviderConfig GetPhoneNumberOutboundSipTrunkConfigResponseModel `json:"provider_config"`
	JSON           struct {
		Label            respjson.Field
		PhoneNumber      respjson.Field
		PhoneNumberID    respjson.Field
		AssignedAgent    respjson.Field
		Provider         respjson.Field
		SupportsInbound  respjson.Field
		SupportsOutbound respjson.Field
		LivekitStack     respjson.Field
		InboundTrunk     respjson.Field
		OutboundTrunk    respjson.Field
		ProviderConfig   respjson.Field
		raw              string
	} `json:"-"`
}

// anyConvaiPhoneNumberListResponse is implemented by each variant of
// [ConvaiPhoneNumberListResponseUnion] to add type safety for the return type of
// [ConvaiPhoneNumberListResponseUnion.AsAny]
type anyConvaiPhoneNumberListResponse interface {
	implConvaiPhoneNumberListResponseUnion()
}

func (GetPhoneNumberTwilioResponseModel) implConvaiPhoneNumberListResponseUnion()   {}
func (GetPhoneNumberSipTrunkResponseModel) implConvaiPhoneNumberListResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiPhoneNumberListResponseUnion.AsAny().(type) {
//	case elevenlabs.GetPhoneNumberTwilioResponseModel:
//	case elevenlabs.GetPhoneNumberSipTrunkResponseModel:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiPhoneNumberListResponseUnion) AsAny() anyConvaiPhoneNumberListResponse {
	switch u.Provider {
	case "twilio":
		return u.AsTwilio()
	case "sip_trunk":
		return u.AsSipTrunk()
	}
	return nil
}

func (u ConvaiPhoneNumberListResponseUnion) AsTwilio() (v GetPhoneNumberTwilioResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiPhoneNumberListResponseUnion) AsSipTrunk() (v GetPhoneNumberSipTrunkResponseModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiPhoneNumberListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiPhoneNumberListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiPhoneNumberDeleteResponse = any

type ConvaiPhoneNumberImportResponse struct {
	// Phone entity ID
	PhoneNumberID string `json:"phone_number_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumberID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiPhoneNumberImportResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiPhoneNumberImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiPhoneNumberGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiPhoneNumberUpdateParams struct {
	AgentID param.Opt[string] `json:"agent_id,omitzero"`
	Label   param.Opt[string] `json:"label,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey           param.Opt[string]                      `header:"xi-api-key,omitzero" json:"-"`
	InboundTrunkConfig InboundSipTrunkConfigRequestModelParam `json:"inbound_trunk_config,omitzero"`
	// Any of "standard", "static".
	LivekitStack        LivekitStackType                        `json:"livekit_stack,omitzero"`
	OutboundTrunkConfig OutboundSipTrunkConfigRequestModelParam `json:"outbound_trunk_config,omitzero"`
	paramObj
}

func (r ConvaiPhoneNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiPhoneNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiPhoneNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiPhoneNumberListParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiPhoneNumberDeleteParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiPhoneNumberImportParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfCreateTwilioPhoneNumberRequest *ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCreateSipTrunkPhoneNumberRequestV2 *ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2 `json:",inline"`

	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (u ConvaiPhoneNumberImportParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateTwilioPhoneNumberRequest, u.OfCreateSipTrunkPhoneNumberRequestV2)
}
func (r *ConvaiPhoneNumberImportParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Token, Label, PhoneNumber, Sid are required.
type ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest struct {
	// Twilio Auth Token
	Token string `json:"token,required"`
	// Label for the phone number
	Label string `json:"label,required"`
	// Phone number
	PhoneNumber string `json:"phone_number,required"`
	// Twilio Account SID
	Sid string `json:"sid,required"`
	// Whether this phone number supports inbound calls
	SupportsInbound param.Opt[bool] `json:"supports_inbound,omitzero"`
	// Whether this phone number supports outbound calls
	SupportsOutbound param.Opt[bool] `json:"supports_outbound,omitzero"`
	// Twilio Additional Region Configuration
	RegionConfig ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig `json:"region_config,omitzero"`
	// Any of "twilio".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest](
		"provider", "twilio",
	)
}

// Twilio Additional Region Configuration
//
// The properties Token, EdgeLocation, RegionID are required.
type ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig struct {
	// Auth Token for this region
	Token string `json:"token,required"`
	// Edge location for this region
	//
	// Any of "ashburn", "dublin", "frankfurt", "sao-paulo", "singapore", "sydney",
	// "tokyo", "umatilla", "roaming".
	EdgeLocation string `json:"edge_location,omitzero,required"`
	// Region ID
	//
	// Any of "us1", "ie1", "au1".
	RegionID string `json:"region_id,omitzero,required"`
	paramObj
}

func (r ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig](
		"edge_location", "ashburn", "dublin", "frankfurt", "sao-paulo", "singapore", "sydney", "tokyo", "umatilla", "roaming",
	)
	apijson.RegisterFieldValidator[ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig](
		"region_id", "us1", "ie1", "au1",
	)
}

// The properties Label, PhoneNumber are required.
type ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2 struct {
	// Label for the phone number
	Label string `json:"label,required"`
	// Phone number
	PhoneNumber string `json:"phone_number,required"`
	// Whether this phone number supports inbound calls
	SupportsInbound param.Opt[bool] `json:"supports_inbound,omitzero"`
	// Whether this phone number supports outbound calls
	SupportsOutbound    param.Opt[bool]                         `json:"supports_outbound,omitzero"`
	InboundTrunkConfig  InboundSipTrunkConfigRequestModelParam  `json:"inbound_trunk_config,omitzero"`
	OutboundTrunkConfig OutboundSipTrunkConfigRequestModelParam `json:"outbound_trunk_config,omitzero"`
	// Any of "sip_trunk".
	Provider string `json:"provider,omitzero"`
	paramObj
}

func (r ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConvaiPhoneNumberImportParamsBodyCreateSipTrunkPhoneNumberRequestV2](
		"provider", "sip_trunk",
	)
}
