// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"

	elevenlabs "github.com/oh-tarnished/elevenlabs-go"
	"github.com/oh-tarnished/elevenlabs-go/internal/testutil"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

func TestVoiceGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Get(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoiceGetParams{
			WithSettings: elevenlabs.Bool(true),
			XiAPIKey:     elevenlabs.String("xi-api-key"),
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

func TestVoiceListWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.List(context.TODO(), elevenlabs.VoiceListParams{
		ShowLegacy: elevenlabs.Bool(true),
		XiAPIKey:   elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Delete(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoiceDeleteParams{
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

func TestVoiceEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Edit(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoiceEditParams{
			Name:        "John Smith",
			Description: elevenlabs.String("An old American male voice with a slight hoarseness in his throat. Perfect for news."),
			Files:       []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
			Labels: elevenlabs.VoiceEditParamsLabelsUnion{
				OfString: elevenlabs.String(`{"language": "en", "accent": "en-US", "gender": "male", "age": "middle-aged"}`),
			},
			RemoveBackgroundNoise: elevenlabs.Bool(true),
			XiAPIKey:              elevenlabs.String("xi-api-key"),
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

func TestVoiceGetV2WithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.GetV2(context.TODO(), elevenlabs.VoiceGetV2Params{
		Category:          elevenlabs.String("category"),
		CollectionID:      elevenlabs.String("collection_id"),
		FineTuningState:   elevenlabs.String("fine_tuning_state"),
		IncludeTotalCount: elevenlabs.Bool(true),
		NextPageToken:     elevenlabs.String("0"),
		PageSize:          elevenlabs.Int(0),
		Search:            elevenlabs.String("search"),
		Sort:              elevenlabs.String("created_at_unix"),
		SortDirection:     elevenlabs.String("desc"),
		VoiceIDs:          []string{"string"},
		VoiceType:         elevenlabs.String("voice_type"),
		XiAPIKey:          elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
