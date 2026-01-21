// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	elevenlabs "github.com/oh-tarnished/elevenlabs-go"
	"github.com/oh-tarnished/elevenlabs-go/option"
)

func TestSpeechToSpeechStreamTransformWithOptionalParams(t *testing.T) {
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
	resp, err := client.SpeechToSpeech.StreamTransform(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.SpeechToSpeechStreamTransformParams{
			Audio:                    io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			EnableLogging:            elevenlabs.Bool(true),
			OptimizeStreamingLatency: elevenlabs.Int(0),
			OutputFormat:             elevenlabs.SpeechToSpeechStreamTransformParamsOutputFormatMP3_22050_32,
			FileFormat:               elevenlabs.SpeechToSpeechStreamTransformParamsFileFormatPcmS16le16,
			ModelID:                  elevenlabs.String("model_id"),
			RemoveBackgroundNoise:    elevenlabs.Bool(true),
			Seed:                     elevenlabs.Int(12345),
			VoiceSettings:            elevenlabs.String("voice_settings"),
			XiAPIKey:                 elevenlabs.String("xi-api-key"),
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

func TestSpeechToSpeechTransformWithOptionalParams(t *testing.T) {
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
	resp, err := client.SpeechToSpeech.Transform(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.SpeechToSpeechTransformParams{
			Audio:                    io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			EnableLogging:            elevenlabs.Bool(true),
			OptimizeStreamingLatency: elevenlabs.Int(0),
			OutputFormat:             elevenlabs.SpeechToSpeechTransformParamsOutputFormatMP3_22050_32,
			FileFormat:               elevenlabs.SpeechToSpeechTransformParamsFileFormatPcmS16le16,
			ModelID:                  elevenlabs.String("model_id"),
			RemoveBackgroundNoise:    elevenlabs.Bool(true),
			Seed:                     elevenlabs.Int(12345),
			VoiceSettings:            elevenlabs.String("voice_settings"),
			XiAPIKey:                 elevenlabs.String("xi-api-key"),
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
