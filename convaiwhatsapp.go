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

// ConvaiWhatsappService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiWhatsappService] method instead.
type ConvaiWhatsappService struct {
	Options []option.RequestOption
}

// NewConvaiWhatsappService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiWhatsappService(opts ...option.RequestOption) (r ConvaiWhatsappService) {
	r = ConvaiWhatsappService{}
	r.Options = opts
	return
}

// Make an outbound call via WhatsApp
func (r *ConvaiWhatsappService) OutboundCall(ctx context.Context, params ConvaiWhatsappOutboundCallParams, opts ...option.RequestOption) (res *ConvaiWhatsappOutboundCallResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/whatsapp/outbound-call"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ConvaiWhatsappOutboundCallResponse struct {
	ConversationID string `json:"conversation_id,required"`
	Message        string `json:"message,required"`
	Success        bool   `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConversationID respjson.Field
		Message        respjson.Field
		Success        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiWhatsappOutboundCallResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiWhatsappOutboundCallResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiWhatsappOutboundCallParams struct {
	AgentID                                           string `json:"agent_id,required"`
	WhatsappCallPermissionRequestTemplateLanguageCode string `json:"whatsapp_call_permission_request_template_language_code,required"`
	WhatsappCallPermissionRequestTemplateName         string `json:"whatsapp_call_permission_request_template_name,required"`
	WhatsappPhoneNumberID                             string `json:"whatsapp_phone_number_id,required"`
	WhatsappUserID                                    string `json:"whatsapp_user_id,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey                         param.Opt[string]                                 `header:"xi-api-key,omitzero" json:"-"`
	ConversationInitiationClientData ConversationInitiationClientDataRequestInputParam `json:"conversation_initiation_client_data,omitzero"`
	paramObj
}

func (r ConvaiWhatsappOutboundCallParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiWhatsappOutboundCallParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiWhatsappOutboundCallParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
