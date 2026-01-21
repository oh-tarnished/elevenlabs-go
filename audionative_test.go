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

func TestAudioNativeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.AudioNative.New(context.TODO(), elevenlabs.AudioNativeNewParams{
		Name:                            "name",
		ApplyTextNormalization:          elevenlabs.AudioNativeNewParamsApplyTextNormalizationAuto,
		Author:                          elevenlabs.String("author"),
		AutoConvert:                     elevenlabs.Bool(true),
		BackgroundColor:                 elevenlabs.String("background_color"),
		File:                            io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		Image:                           elevenlabs.String("image"),
		ModelID:                         elevenlabs.String("model_id"),
		PronunciationDictionaryLocators: []string{`{"pronunciation_dictionary_id": "21m00Tcm4TlvDq8ikWAM", "version_id": "BdF0s0aZ3oFoKnDYdTox"}`},
		Sessionization:                  elevenlabs.Int(0),
		Small:                           elevenlabs.Bool(true),
		TextColor:                       elevenlabs.String("text_color"),
		Title:                           elevenlabs.String("title"),
		VoiceID:                         elevenlabs.String("voice_id"),
		XiAPIKey:                        elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAudioNativeGetSettingsWithOptionalParams(t *testing.T) {
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
	_, err := client.AudioNative.GetSettings(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.AudioNativeGetSettingsParams{
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

func TestAudioNativeUpdateContentWithOptionalParams(t *testing.T) {
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
	_, err := client.AudioNative.UpdateContent(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.AudioNativeUpdateContentParams{
			AutoConvert: elevenlabs.Bool(true),
			AutoPublish: elevenlabs.Bool(true),
			File:        io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			XiAPIKey:    elevenlabs.String("xi-api-key"),
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
