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

func TestConvaiPhoneNumberGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.PhoneNumbers.Get(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiPhoneNumberGetParams{
			XiAPIKey: elevenlabs.String("xi-api-key"),
		},
	)
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConvaiPhoneNumberUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.PhoneNumbers.Update(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiPhoneNumberUpdateParams{
			AgentID: elevenlabs.String("agent_id"),
			InboundTrunkConfig: elevenlabs.InboundSipTrunkConfigRequestModelParam{
				AllowedAddresses: []string{"string"},
				AllowedNumbers:   []string{"string"},
				Credentials: elevenlabs.SipTrunkCredentialsRequestModelParam{
					Username: "username",
					Password: elevenlabs.String("password"),
				},
				MediaEncryption: elevenlabs.SipMediaEncryptionEnumDisabled,
				RemoteDomains:   []string{"string"},
			},
			Label:        elevenlabs.String("label"),
			LivekitStack: elevenlabs.LivekitStackTypeStandard,
			OutboundTrunkConfig: elevenlabs.OutboundSipTrunkConfigRequestModelParam{
				Address: "address",
				Credentials: elevenlabs.SipTrunkCredentialsRequestModelParam{
					Username: "username",
					Password: elevenlabs.String("password"),
				},
				Headers: map[string]string{
					"foo": "string",
				},
				MediaEncryption: elevenlabs.SipMediaEncryptionEnumDisabled,
				Transport:       elevenlabs.SipTrunkTransportEnumAuto,
			},
			XiAPIKey: elevenlabs.String("xi-api-key"),
		},
	)
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConvaiPhoneNumberListWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.PhoneNumbers.List(context.TODO(), elevenlabs.ConvaiPhoneNumberListParams{
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

func TestConvaiPhoneNumberDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.PhoneNumbers.Delete(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiPhoneNumberDeleteParams{
			XiAPIKey: elevenlabs.String("xi-api-key"),
		},
	)
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConvaiPhoneNumberImportWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.PhoneNumbers.Import(context.TODO(), elevenlabs.ConvaiPhoneNumberImportParams{
		OfCreateTwilioPhoneNumberRequest: &elevenlabs.ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequest{
			Token:       "token",
			Label:       "label",
			PhoneNumber: "phone_number",
			Sid:         "sid",
			Provider:    "twilio",
			RegionConfig: elevenlabs.ConvaiPhoneNumberImportParamsBodyCreateTwilioPhoneNumberRequestRegionConfig{
				Token:        "token",
				EdgeLocation: "ashburn",
				RegionID:     "us1",
			},
			SupportsInbound:  elevenlabs.Bool(true),
			SupportsOutbound: elevenlabs.Bool(true),
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
