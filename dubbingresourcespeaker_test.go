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

func TestDubbingResourceSpeakerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Resource.Speaker.New(
		context.TODO(),
		"dubbing_id",
		elevenlabs.DubbingResourceSpeakerNewParams{
			SpeakerName:     elevenlabs.String("speaker_name"),
			VoiceID:         elevenlabs.String("voice_id"),
			VoiceSimilarity: elevenlabs.Float(0),
			VoiceStability:  elevenlabs.Float(0),
			VoiceStyle:      elevenlabs.Float(0),
			XiAPIKey:        elevenlabs.String("xi-api-key"),
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

func TestDubbingResourceSpeakerNewSegmentWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Resource.Speaker.NewSegment(
		context.TODO(),
		"speaker_id",
		elevenlabs.DubbingResourceSpeakerNewSegmentParams{
			DubbingID: "dubbing_id",
			EndTime:   0,
			StartTime: 0,
			Text:      elevenlabs.String("text"),
			Translations: map[string]string{
				"foo": "string",
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

func TestDubbingResourceSpeakerFindSimilarVoicesWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Resource.Speaker.FindSimilarVoices(
		context.TODO(),
		"speaker_id",
		elevenlabs.DubbingResourceSpeakerFindSimilarVoicesParams{
			DubbingID: "dubbing_id",
			XiAPIKey:  elevenlabs.String("xi-api-key"),
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

func TestDubbingResourceSpeakerUpdateMetadataWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Resource.Speaker.UpdateMetadata(
		context.TODO(),
		"speaker_id",
		elevenlabs.DubbingResourceSpeakerUpdateMetadataParams{
			DubbingID:       "dubbing_id",
			Languages:       []string{"string"},
			SpeakerName:     elevenlabs.String("speaker_name"),
			VoiceID:         elevenlabs.String("voice_id"),
			VoiceSimilarity: elevenlabs.Float(0),
			VoiceStability:  elevenlabs.Float(0),
			VoiceStyle:      elevenlabs.Float(0),
			XiAPIKey:        elevenlabs.String("xi-api-key"),
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
