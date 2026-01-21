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

func TestDubbingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.New(context.TODO(), elevenlabs.DubbingNewParams{
		BackgroundAudioFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		CsvFile:             io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		CsvFps:              elevenlabs.Float(0),
		DisableVoiceCloning: elevenlabs.Bool(true),
		DropBackgroundAudio: elevenlabs.Bool(true),
		DubbingStudio:       elevenlabs.Bool(true),
		EndTime:             elevenlabs.Int(0),
		File:                io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		ForegroundAudioFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		HighestResolution:   elevenlabs.Bool(true),
		Mode:                elevenlabs.DubbingNewParamsModeAutomatic,
		Name:                elevenlabs.String("name"),
		NumSpeakers:         elevenlabs.Int(0),
		SourceLang:          elevenlabs.String("source_lang"),
		SourceURL:           elevenlabs.String("source_url"),
		StartTime:           elevenlabs.Int(0),
		TargetAccent:        elevenlabs.String("target_accent"),
		TargetLang:          elevenlabs.String("target_lang"),
		UseProfanityFilter:  elevenlabs.Bool(true),
		Watermark:           elevenlabs.Bool(true),
		XiAPIKey:            elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDubbingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Get(
		context.TODO(),
		"dubbing_id",
		elevenlabs.DubbingGetParams{
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

func TestDubbingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.List(context.TODO(), elevenlabs.DubbingListParams{
		Cursor:          elevenlabs.String("cursor"),
		DubbingStatus:   elevenlabs.DubbingListParamsDubbingStatusDubbing,
		FilterByCreator: elevenlabs.DubbingListParamsFilterByCreatorAll,
		OrderBy:         elevenlabs.DubbingListParamsOrderByCreatedAt,
		OrderDirection:  elevenlabs.DubbingListParamsOrderDirectionDescending,
		PageSize:        elevenlabs.Int(1),
		XiAPIKey:        elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDubbingDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.Delete(
		context.TODO(),
		"dubbing_id",
		elevenlabs.DubbingDeleteParams{
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

func TestDubbingGetAudioWithOptionalParams(t *testing.T) {
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
	resp, err := client.Dubbing.GetAudio(
		context.TODO(),
		"language_code",
		elevenlabs.DubbingGetAudioParams{
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

func TestDubbingGetTranscriptWithOptionalParams(t *testing.T) {
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
	_, err := client.Dubbing.GetTranscript(
		context.TODO(),
		"source",
		elevenlabs.DubbingGetTranscriptParams{
			DubbingID:  "dubbing_id",
			FormatType: elevenlabs.DubbingGetTranscriptParamsFormatTypeSrt,
			XiAPIKey:   elevenlabs.String("xi-api-key"),
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
