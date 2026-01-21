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

func TestConvaiSettingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Settings.Get(context.TODO(), elevenlabs.ConvaiSettingGetParams{
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

func TestConvaiSettingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Settings.Update(context.TODO(), elevenlabs.ConvaiSettingUpdateParams{
		CanUseMcpServers: elevenlabs.Bool(true),
		ConversationInitiationClientDataWebhook: elevenlabs.ConversationInitiationClientDataWebhookParam{
			RequestHeaders: map[string]elevenlabs.ConversationInitiationClientDataWebhookRequestHeaderUnionParam{
				"Content-Type": {
					OfString: elevenlabs.String("application/json"),
				},
			},
			URL: "https://example.com/webhook",
		},
		DefaultLivekitStack:    elevenlabs.LivekitStackTypeStandard,
		RagRetentionPeriodDays: elevenlabs.Int(1),
		Webhooks: elevenlabs.ConvAIWebhooksParam{
			Events:            []string{"transcript"},
			PostCallWebhookID: elevenlabs.String("post_call_webhook_id"),
			SendAudio:         elevenlabs.Bool(true),
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
