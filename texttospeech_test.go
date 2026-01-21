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

func TestTextToSpeechConvertWithOptionalParams(t *testing.T) {
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
	resp, err := client.TextToSpeech.Convert(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.TextToSpeechConvertParams{
			Text:                           "This is a test for the API of ElevenLabs.",
			EnableLogging:                  elevenlabs.Bool(true),
			OptimizeStreamingLatency:       elevenlabs.Int(0),
			OutputFormat:                   elevenlabs.TextToSpeechConvertParamsOutputFormatAlaw8000,
			ApplyLanguageTextNormalization: elevenlabs.Bool(true),
			ApplyTextNormalization:         elevenlabs.TextToSpeechConvertParamsApplyTextNormalizationAuto,
			LanguageCode:                   elevenlabs.String("language_code"),
			ModelID:                        elevenlabs.String("model_id"),
			NextRequestIDs:                 []string{"3tPgBrD1UdW3snUkGw1K", "4D1jAxiRFkolBNUGzXkU", "4c8Z4aWliVR2oipYRXhj"},
			NextText:                       elevenlabs.String("The Willowbrook Festival, held every spring, celebrates the blossoming of the wild bluebells that carpet the nearby forest floors, creating a breathtaking sea of blue under the canopy of fresh green leaves."),
			PreviousRequestIDs:             []string{"09bOJkdYVjKy2oOiiVtR", "0p2uKqOnZyce22SPZ9d5", "1KYvY8WZAKmcjCJ1mvVB"},
			PreviousText:                   elevenlabs.String("In the heart of a lush valley surrounded by towering mountains lies the quaint village of Willowbrook."),
			PronunciationDictionaryLocators: []elevenlabs.PronunciationDictionaryVersionLocatorParam{{
				PronunciationDictionaryID: "test",
				VersionID:                 elevenlabs.String("id2"),
			}},
			Seed:        elevenlabs.Int(12345),
			UsePvcAsIvc: elevenlabs.Bool(true),
			VoiceSettings: elevenlabs.VoiceSettingsParam{
				SimilarityBoost: elevenlabs.Float(1),
				Speed:           elevenlabs.Float(1),
				Stability:       elevenlabs.Float(1),
				Style:           elevenlabs.Float(0),
				UseSpeakerBoost: elevenlabs.Bool(true),
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

func TestTextToSpeechConvertWithTimestampsWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToSpeech.ConvertWithTimestamps(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.TextToSpeechConvertWithTimestampsParams{
			Text:                           "This is a test for the API of ElevenLabs.",
			EnableLogging:                  elevenlabs.Bool(true),
			OptimizeStreamingLatency:       elevenlabs.Int(0),
			OutputFormat:                   elevenlabs.TextToSpeechConvertWithTimestampsParamsOutputFormatAlaw8000,
			ApplyLanguageTextNormalization: elevenlabs.Bool(true),
			ApplyTextNormalization:         elevenlabs.TextToSpeechConvertWithTimestampsParamsApplyTextNormalizationAuto,
			LanguageCode:                   elevenlabs.String("language_code"),
			ModelID:                        elevenlabs.String("model_id"),
			NextRequestIDs:                 []string{"3tPgBrD1UdW3snUkGw1K", "4D1jAxiRFkolBNUGzXkU", "4c8Z4aWliVR2oipYRXhj"},
			NextText:                       elevenlabs.String("The Willowbrook Festival, held every spring, celebrates the blossoming of the wild bluebells that carpet the nearby forest floors, creating a breathtaking sea of blue under the canopy of fresh green leaves."),
			PreviousRequestIDs:             []string{"09bOJkdYVjKy2oOiiVtR", "0p2uKqOnZyce22SPZ9d5", "1KYvY8WZAKmcjCJ1mvVB"},
			PreviousText:                   elevenlabs.String("In the heart of a lush valley surrounded by towering mountains lies the quaint village of Willowbrook."),
			PronunciationDictionaryLocators: []elevenlabs.PronunciationDictionaryVersionLocatorParam{{
				PronunciationDictionaryID: "test",
				VersionID:                 elevenlabs.String("id2"),
			}},
			Seed:        elevenlabs.Int(12345),
			UsePvcAsIvc: elevenlabs.Bool(true),
			VoiceSettings: elevenlabs.VoiceSettingsParam{
				SimilarityBoost: elevenlabs.Float(1),
				Speed:           elevenlabs.Float(1),
				Stability:       elevenlabs.Float(1),
				Style:           elevenlabs.Float(0),
				UseSpeakerBoost: elevenlabs.Bool(true),
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

func TestTextToSpeechStreamConvertWithOptionalParams(t *testing.T) {
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
	resp, err := client.TextToSpeech.StreamConvert(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.TextToSpeechStreamConvertParams{
			Text:                           "This is a test for the API of ElevenLabs.",
			EnableLogging:                  elevenlabs.Bool(true),
			OptimizeStreamingLatency:       elevenlabs.Int(0),
			OutputFormat:                   elevenlabs.TextToSpeechStreamConvertParamsOutputFormatMP3_22050_32,
			ApplyLanguageTextNormalization: elevenlabs.Bool(true),
			ApplyTextNormalization:         elevenlabs.TextToSpeechStreamConvertParamsApplyTextNormalizationAuto,
			LanguageCode:                   elevenlabs.String("language_code"),
			ModelID:                        elevenlabs.String("model_id"),
			NextRequestIDs:                 []string{"3tPgBrD1UdW3snUkGw1K", "4D1jAxiRFkolBNUGzXkU", "4c8Z4aWliVR2oipYRXhj"},
			NextText:                       elevenlabs.String("The Willowbrook Festival, held every spring, celebrates the blossoming of the wild bluebells that carpet the nearby forest floors, creating a breathtaking sea of blue under the canopy of fresh green leaves."),
			PreviousRequestIDs:             []string{"09bOJkdYVjKy2oOiiVtR", "0p2uKqOnZyce22SPZ9d5", "1KYvY8WZAKmcjCJ1mvVB"},
			PreviousText:                   elevenlabs.String("In the heart of a lush valley surrounded by towering mountains lies the quaint village of Willowbrook."),
			PronunciationDictionaryLocators: []elevenlabs.PronunciationDictionaryVersionLocatorParam{{
				PronunciationDictionaryID: "test",
				VersionID:                 elevenlabs.String("id2"),
			}},
			Seed:        elevenlabs.Int(12345),
			UsePvcAsIvc: elevenlabs.Bool(true),
			VoiceSettings: elevenlabs.VoiceSettingsParam{
				SimilarityBoost: elevenlabs.Float(1),
				Speed:           elevenlabs.Float(1),
				Stability:       elevenlabs.Float(1),
				Style:           elevenlabs.Float(0),
				UseSpeakerBoost: elevenlabs.Bool(true),
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

func TestTextToSpeechStreamConvertWithTimestampsWithOptionalParams(t *testing.T) {
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
	_, err := client.TextToSpeech.StreamConvertWithTimestamps(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.TextToSpeechStreamConvertWithTimestampsParams{
			Text:                           "This is a test for the API of ElevenLabs.",
			EnableLogging:                  elevenlabs.Bool(true),
			OptimizeStreamingLatency:       elevenlabs.Int(0),
			OutputFormat:                   elevenlabs.TextToSpeechStreamConvertWithTimestampsParamsOutputFormatMP3_22050_32,
			ApplyLanguageTextNormalization: elevenlabs.Bool(true),
			ApplyTextNormalization:         elevenlabs.TextToSpeechStreamConvertWithTimestampsParamsApplyTextNormalizationAuto,
			LanguageCode:                   elevenlabs.String("language_code"),
			ModelID:                        elevenlabs.String("model_id"),
			NextRequestIDs:                 []string{"3tPgBrD1UdW3snUkGw1K", "4D1jAxiRFkolBNUGzXkU", "4c8Z4aWliVR2oipYRXhj"},
			NextText:                       elevenlabs.String("The Willowbrook Festival, held every spring, celebrates the blossoming of the wild bluebells that carpet the nearby forest floors, creating a breathtaking sea of blue under the canopy of fresh green leaves."),
			PreviousRequestIDs:             []string{"09bOJkdYVjKy2oOiiVtR", "0p2uKqOnZyce22SPZ9d5", "1KYvY8WZAKmcjCJ1mvVB"},
			PreviousText:                   elevenlabs.String("In the heart of a lush valley surrounded by towering mountains lies the quaint village of Willowbrook."),
			PronunciationDictionaryLocators: []elevenlabs.PronunciationDictionaryVersionLocatorParam{{
				PronunciationDictionaryID: "test",
				VersionID:                 elevenlabs.String("id2"),
			}},
			Seed:        elevenlabs.Int(12345),
			UsePvcAsIvc: elevenlabs.Bool(true),
			VoiceSettings: elevenlabs.VoiceSettingsParam{
				SimilarityBoost: elevenlabs.Float(1),
				Speed:           elevenlabs.Float(1),
				Stability:       elevenlabs.Float(1),
				Style:           elevenlabs.Float(0),
				UseSpeakerBoost: elevenlabs.Bool(true),
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
