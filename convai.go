// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"github.com/oh-tarnished/elevenlabs-go/option"
)

// ConvaiService contains methods and other services that help with interacting
// with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiService] method instead.
type ConvaiService struct {
	Options          []option.RequestOption
	Twilio           ConvaiTwilioService
	Whatsapp         ConvaiWhatsappService
	AgentTesting     ConvaiAgentTestingService
	TestInvocations  ConvaiTestInvocationService
	PhoneNumbers     ConvaiPhoneNumberService
	LlmUsage         ConvaiLlmUsageService
	Analytics        ConvaiAnalyticsService
	KnowledgeBase    ConvaiKnowledgeBaseService
	Tools            ConvaiToolService
	Settings         ConvaiSettingService
	Secrets          ConvaiSecretService
	BatchCalling     ConvaiBatchCallingService
	SipTrunk         ConvaiSipTrunkService
	McpServers       ConvaiMcpServerService
	WhatsappAccounts ConvaiWhatsappAccountService
	Conversations    ConvaiConversationService
	Agents           ConvaiAgentService
}

// NewConvaiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewConvaiService(opts ...option.RequestOption) (r ConvaiService) {
	r = ConvaiService{}
	r.Options = opts
	r.Twilio = NewConvaiTwilioService(opts...)
	r.Whatsapp = NewConvaiWhatsappService(opts...)
	r.AgentTesting = NewConvaiAgentTestingService(opts...)
	r.TestInvocations = NewConvaiTestInvocationService(opts...)
	r.PhoneNumbers = NewConvaiPhoneNumberService(opts...)
	r.LlmUsage = NewConvaiLlmUsageService(opts...)
	r.Analytics = NewConvaiAnalyticsService(opts...)
	r.KnowledgeBase = NewConvaiKnowledgeBaseService(opts...)
	r.Tools = NewConvaiToolService(opts...)
	r.Settings = NewConvaiSettingService(opts...)
	r.Secrets = NewConvaiSecretService(opts...)
	r.BatchCalling = NewConvaiBatchCallingService(opts...)
	r.SipTrunk = NewConvaiSipTrunkService(opts...)
	r.McpServers = NewConvaiMcpServerService(opts...)
	r.WhatsappAccounts = NewConvaiWhatsappAccountService(opts...)
	r.Conversations = NewConvaiConversationService(opts...)
	r.Agents = NewConvaiAgentService(opts...)
	return
}
