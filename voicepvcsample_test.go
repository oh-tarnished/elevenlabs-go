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

func TestVoicePvcSampleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.Update(
		context.TODO(),
		"VW7YKqPnjY4h39yTbx2L",
		elevenlabs.VoicePvcSampleUpdateParams{
			VoiceID:               "21m00Tcm4TlvDq8ikWAM",
			FileName:              elevenlabs.String("sample.mp3"),
			RemoveBackgroundNoise: elevenlabs.Bool(true),
			SelectedSpeakerIDs:    []string{"string"},
			TrimEndTime:           elevenlabs.Int(10),
			TrimStartTime:         elevenlabs.Int(0),
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

func TestVoicePvcSampleDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.Delete(
		context.TODO(),
		"VW7YKqPnjY4h39yTbx2L",
		elevenlabs.VoicePvcSampleDeleteParams{
			VoiceID:  "21m00Tcm4TlvDq8ikWAM",
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

func TestVoicePvcSampleAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.Add(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoicePvcSampleAddParams{
			Files:                 []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
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

func TestVoicePvcSampleGetAudioWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.GetAudio(
		context.TODO(),
		"VW7YKqPnjY4h39yTbx2L",
		elevenlabs.VoicePvcSampleGetAudioParams{
			VoiceID:               "21m00Tcm4TlvDq8ikWAM",
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

func TestVoicePvcSampleGetWaveformWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.GetWaveform(
		context.TODO(),
		"VW7YKqPnjY4h39yTbx2L",
		elevenlabs.VoicePvcSampleGetWaveformParams{
			VoiceID:  "21m00Tcm4TlvDq8ikWAM",
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

func TestVoicePvcSampleSeparateSpeakersWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Samples.SeparateSpeakers(
		context.TODO(),
		"VW7YKqPnjY4h39yTbx2L",
		elevenlabs.VoicePvcSampleSeparateSpeakersParams{
			VoiceID:  "21m00Tcm4TlvDq8ikWAM",
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
