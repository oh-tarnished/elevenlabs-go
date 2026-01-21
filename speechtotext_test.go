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

func TestSpeechToTextTranscribeWithOptionalParams(t *testing.T) {
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
	_, err := client.SpeechToText.Transcribe(context.TODO(), elevenlabs.SpeechToTextTranscribeParams{
		ModelID:       elevenlabs.SpeechToTextTranscribeParamsModelIDScribeV1,
		EnableLogging: elevenlabs.Bool(true),
		AdditionalFormats: []elevenlabs.SpeechToTextTranscribeParamsAdditionalFormatUnion{{
			OfSrt: &elevenlabs.SpeechToTextTranscribeParamsAdditionalFormatSrt{
				IncludeSpeakers:             elevenlabs.Bool(true),
				IncludeTimestamps:           elevenlabs.Bool(true),
				MaxCharactersPerLine:        elevenlabs.Int(0),
				MaxSegmentChars:             elevenlabs.Int(0),
				MaxSegmentDurationS:         elevenlabs.Float(0),
				SegmentOnSilenceLongerThanS: elevenlabs.Float(0),
			},
		}, {
			OfTxt: &elevenlabs.SpeechToTextTranscribeParamsAdditionalFormatTxt{
				IncludeSpeakers:             elevenlabs.Bool(false),
				IncludeTimestamps:           elevenlabs.Bool(true),
				MaxCharactersPerLine:        elevenlabs.Int(0),
				MaxSegmentChars:             elevenlabs.Int(0),
				MaxSegmentDurationS:         elevenlabs.Float(0),
				SegmentOnSilenceLongerThanS: elevenlabs.Float(0),
			},
		}},
		CloudStorageURL:      elevenlabs.String("https://storage.googleapis.com/my-bucket/folder/audio.mp3"),
		DiarizationThreshold: elevenlabs.Float(0.1),
		Diarize:              elevenlabs.Bool(true),
		EntityDetection: elevenlabs.SpeechToTextTranscribeParamsEntityDetectionUnion{
			OfString: elevenlabs.String("string"),
		},
		File:                  io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		FileFormat:            elevenlabs.SpeechToTextTranscribeParamsFileFormatPcmS16le16,
		Keyterms:              []string{"string"},
		LanguageCode:          elevenlabs.String("language_code"),
		NumSpeakers:           elevenlabs.Int(1),
		Seed:                  elevenlabs.Int(12345),
		TagAudioEvents:        elevenlabs.Bool(true),
		Temperature:           elevenlabs.Float(0),
		TimestampsGranularity: elevenlabs.SpeechToTextTranscribeParamsTimestampsGranularityNone,
		UseMultiChannel:       elevenlabs.Bool(true),
		Webhook:               elevenlabs.Bool(true),
		WebhookID:             elevenlabs.String("webhook_id"),
		WebhookMetadata:       `{"user_id": "123", "session_id": "abc-def-ghi"}`,
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
