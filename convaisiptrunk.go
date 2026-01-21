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

// ConvaiSipTrunkService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiSipTrunkService] method instead.
type ConvaiSipTrunkService struct {
	Options []option.RequestOption
}

// NewConvaiSipTrunkService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiSipTrunkService(opts ...option.RequestOption) (r ConvaiSipTrunkService) {
	r = ConvaiSipTrunkService{}
	r.Options = opts
	return
}

// Handle an outbound call via SIP trunk
func (r *ConvaiSipTrunkService) HandleOutboundCall(ctx context.Context, params ConvaiSipTrunkHandleOutboundCallParams, opts ...option.RequestOption) (res *ConvaiSipTrunkHandleOutboundCallResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/sip-trunk/outbound-call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ConvaiSipTrunkHandleOutboundCallResponse struct {
	ConversationID string `json:"conversation_id,required"`
	Message        string `json:"message,required"`
	SipCallID      string `json:"sip_call_id,required"`
	Success        bool   `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		Message        respjson.Field
		SipCallID      respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiSipTrunkHandleOutboundCallResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiSipTrunkHandleOutboundCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSipTrunkHandleOutboundCallParams struct {
	AgentID            string `json:"agent_id,required"`
	AgentPhoneNumberID string `json:"agent_phone_number_id,required"`
	ToNumber           string `json:"to_number,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey                         param.Opt[string]                                 `header:"xi-api-key,omitzero" json:"-"`
	ConversationInitiationClientData ConversationInitiationClientDataRequestInputParam `json:"conversation_initiation_client_data,omitzero"`
	paramObj
}

func (r ConvaiSipTrunkHandleOutboundCallParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSipTrunkHandleOutboundCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSipTrunkHandleOutboundCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
