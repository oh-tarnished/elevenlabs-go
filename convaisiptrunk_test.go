// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs_test

import (
	"context"
	"errors"
	"os"
	"testing"

	elevenlabs "github.com/oh-tarnished/elevenlabs-go"
	"github.com/oh-tarnished/elevenlabs-go/internal/testutil"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

func TestConvaiSipTrunkHandleOutboundCallWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := elevenlabs.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Convai.SipTrunk.HandleOutboundCall(context.TODO(), elevenlabs.ConvaiSipTrunkHandleOutboundCallParams{
		AgentID:            "agent_id",
		AgentPhoneNumberID: "agent_phone_number_id",
		ToNumber:           "to_number",
		ConversationInitiationClientData: elevenlabs.ConversationInitiationClientDataRequestInputParam{
			ConversationConfigOverride: elevenlabs.ConversationInitiationClientDataRequestInputConversationConfigOverrideParam{
				Agent: elevenlabs.ConversationInitiationClientDataRequestInputConversationConfigOverrideAgentParam{
					FirstMessage: elevenlabs.String("Hello, how can I help you today?"),
					Language:     elevenlabs.String("en"),
					Prompt: elevenlabs.PromptAgentAPIModelOverrideParam{
						Llm:                elevenlabs.LlmGemini2_0Flash001,
						NativeMcpServerIDs: []string{"string"},
						Prompt:             elevenlabs.String("You are a helpful assistant that can answer questions about the topic of the conversation."),
					},
				},
				Conversation: elevenlabs.ConversationConfigOverrideParam{
					TextOnly: elevenlabs.Bool(true),
				},
				Tts: elevenlabs.TtsConversationalConfigOverrideParam{
					SimilarityBoost: elevenlabs.Float(0.8),
					Speed:           elevenlabs.Float(1),
					Stability:       elevenlabs.Float(0.5),
					VoiceID:         elevenlabs.String("cjVigY5qzO86Huf0OWal"),
				},
				Turn: elevenlabs.TurnConfigOverrideParam{
					SoftTimeoutConfig: elevenlabs.TurnConfigOverrideSoftTimeoutConfigParam{
						Message: elevenlabs.String("Hhmmmm...yeah."),
					},
				},
			},
			CustomLlmExtraBody: map[string]any{},
			DynamicVariables: map[string]elevenlabs.ConversationInitiationClientDataRequestInputDynamicVariableUnionParam{
				"foo": {
					OfString: elevenlabs.String("string"),
				},
			},
			SourceInfo: elevenlabs.ConversationInitiationSourceInfoParam{
				Source:  elevenlabs.ConversationInitiationSourceUnknown,
				Version: elevenlabs.String("version"),
			},
			UserID: elevenlabs.String("user_id"),
		},
		XiAPIKey: elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
