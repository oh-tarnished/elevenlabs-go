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

func TestVoiceAddNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Add.New(context.TODO(), elevenlabs.VoiceAddNewParams{
		Files:       []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
		Name:        "John Smith",
		Description: elevenlabs.String("An old American male voice with a slight hoarseness in his throat. Perfect for news."),
		Labels: elevenlabs.VoiceAddNewParamsLabelsUnion{
			OfString: elevenlabs.String(`{"language": "en", "accent": "en-US", "gender": "male", "age": "middle-aged"}`),
		},
		RemoveBackgroundNoise: elevenlabs.Bool(true),
		XiAPIKey:              elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestVoiceAddAddSharedWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Add.AddShared(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoiceAddAddSharedParams{
			PublicUserID: "63e06b7e7cafdc46be4d2e0b3f045940231ae058d508589653d74d1265a574ca",
			NewName:      "John Smith",
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
