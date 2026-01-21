// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	elevenlabs "github.com/oh-tarnished/elevenlabs-go"
	"github.com/oh-tarnished/elevenlabs-go/internal/testutil"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

func TestTextToVoiceNewPreviewsWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToVoice.NewPreviews(context.TODO(), elevenlabs.TextToVoiceNewPreviewsParams{
		VoiceDescription: "A sassy squeaky mouse",
		OutputFormat:     elevenlabs.AllowedOutputFormatsMP3_22050_32,
		AutoGenerateText: elevenlabs.Bool(true),
		GuidanceScale:    elevenlabs.Float(5),
		Loudness:         elevenlabs.Float(0.5),
		Quality:          elevenlabs.Float(0.9),
		Seed:             elevenlabs.Int(11),
		ShouldEnhance:    elevenlabs.Bool(true),
		Text:             elevenlabs.String("Every act of kindness, no matter how small, carries value and can make a difference, as no gesture of goodwill is ever wasted."),
		XiAPIKey:         elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTextToVoiceNewVoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToVoice.NewVoice(context.TODO(), elevenlabs.TextToVoiceNewVoiceParams{
		GeneratedVoiceID: "37HceQefKmEi3bGovXjL",
		VoiceDescription: "A sassy squeaky mouse",
		VoiceName:        "Sassy squeaky mouse",
		Labels: map[string]string{
			"language": "en",
		},
		PlayedNotSelectedVoiceIDs: []string{"string"},
		XiAPIKey:                  elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTextToVoiceDesignVoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToVoice.DesignVoice(context.TODO(), elevenlabs.TextToVoiceDesignVoiceParams{
		VoiceDescription:           "A sassy squeaky mouse",
		OutputFormat:               elevenlabs.AllowedOutputFormatsMP3_22050_32,
		AutoGenerateText:           elevenlabs.Bool(true),
		GuidanceScale:              elevenlabs.Float(5),
		Loudness:                   elevenlabs.Float(0.5),
		ModelID:                    elevenlabs.TextToVoiceDesignVoiceParamsModelIDElevenMultilingualTtvV2,
		PromptStrength:             elevenlabs.Float(0.25),
		Quality:                    elevenlabs.Float(0.9),
		ReferenceAudioBase64:       elevenlabs.String("reference_audio_base64"),
		RemixingSessionID:          elevenlabs.String("123"),
		RemixingSessionIterationID: elevenlabs.String("123"),
		Seed:                       elevenlabs.Int(11),
		ShouldEnhance:              elevenlabs.Bool(true),
		StreamPreviews:             elevenlabs.Bool(true),
		Text:                       elevenlabs.String("Every act of kindness, no matter how small, carries value and can make a difference, as no gesture of goodwill is ever wasted."),
		XiAPIKey:                   elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTextToVoiceRemixVoiceWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToVoice.RemixVoice(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.TextToVoiceRemixVoiceParams{
			VoiceDescription:           "Make the voice have a higher pitch.",
			OutputFormat:               elevenlabs.AllowedOutputFormatsMP3_22050_32,
			AutoGenerateText:           elevenlabs.Bool(true),
			GuidanceScale:              elevenlabs.Float(5),
			Loudness:                   elevenlabs.Float(0.5),
			PromptStrength:             elevenlabs.Float(0.25),
			RemixingSessionID:          elevenlabs.String("123"),
			RemixingSessionIterationID: elevenlabs.String("123"),
			Seed:                       elevenlabs.Int(11),
			StreamPreviews:             elevenlabs.Bool(true),
			Text:                       elevenlabs.String("Every act of kindness, no matter how small, carries value and can make a difference, as no gesture of goodwill is ever wasted."),
			XiAPIKey:                   elevenlabs.String("xi-api-key"),
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

func TestTextToVoiceStreamPreviewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := elevenlabs.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	resp, err := client.TextToVoice.StreamPreview(
		context.TODO(),
		"37HceQefKmEi3bGovXjL",
		elevenlabs.TextToVoiceStreamPreviewParams{
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}
