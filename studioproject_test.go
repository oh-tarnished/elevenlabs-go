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

func TestStudioProjectNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.New(context.TODO(), elevenlabs.StudioProjectNewParams{
		Name:                            "Project 1",
		AcxVolumeNormalization:          elevenlabs.Bool(false),
		ApplyTextNormalization:          elevenlabs.StudioProjectNewParamsApplyTextNormalizationAuto,
		Author:                          elevenlabs.String("William Shakespeare"),
		AutoAssignVoices:                elevenlabs.Bool(true),
		AutoConvert:                     elevenlabs.Bool(true),
		CallbackURL:                     elevenlabs.String("callback_url"),
		ContentType:                     elevenlabs.String("Book"),
		DefaultModelID:                  elevenlabs.String("eleven_multilingual_v2"),
		DefaultParagraphVoiceID:         elevenlabs.String("21m00Tcm4TlvDq8ikWAM"),
		DefaultTitleVoiceID:             elevenlabs.String("21m00Tcm4TlvDq8ikWAM"),
		Description:                     elevenlabs.String("A tragic love story between two young lovers."),
		Fiction:                         elevenlabs.StudioProjectNewParamsFictionFiction,
		FromContentJson:                 elevenlabs.String(`[{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]`),
		FromDocument:                    io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		FromURL:                         elevenlabs.String("https://blog.elevenlabs.io/the_first_ai_that_can_laugh/"),
		Genres:                          []string{"Romance", "Drama"},
		IsbnNumber:                      elevenlabs.String("0-306-40615-2"),
		Language:                        elevenlabs.String("en"),
		MatureContent:                   elevenlabs.Bool(false),
		OriginalPublicationDate:         elevenlabs.String("1597-01-01"),
		PronunciationDictionaryLocators: []string{`{"pronunciation_dictionary_id": "21m00Tcm4TlvDq8ikWAM", "version_id": "BdF0s0aZ3oFoKnDYdTox"}`},
		QualityPreset:                   elevenlabs.StudioProjectNewParamsQualityPresetStandard,
		SourceType:                      elevenlabs.StudioProjectNewParamsSourceTypeBook,
		TargetAudience:                  elevenlabs.StudioProjectNewParamsTargetAudienceAdult,
		Title:                           elevenlabs.String("Romeo and Juliet"),
		VoiceSettings:                   []string{`{"voice_id": "21m00Tcm4TlvDq8ikWAM", "stability": 0.7, "similarity_boost": 0.8, "style": 0.5, "speed": 1.0, "use_speaker_boost": true}`},
		VolumeNormalization:             elevenlabs.Bool(false),
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

func TestStudioProjectGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.Get(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectGetParams{
			ShareID:  elevenlabs.String("share_id"),
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

func TestStudioProjectUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.Update(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectUpdateParams{
			DefaultParagraphVoiceID: "21m00Tcm4TlvDq8ikWAM",
			DefaultTitleVoiceID:     "21m00Tcm4TlvDq8ikWAM",
			Name:                    "Project 1",
			Author:                  elevenlabs.String("William Shakespeare"),
			IsbnNumber:              elevenlabs.String("0-306-40615-2"),
			Title:                   elevenlabs.String("Romeo and Juliet"),
			VolumeNormalization:     elevenlabs.Bool(false),
			XiAPIKey:                elevenlabs.String("xi-api-key"),
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

func TestStudioProjectListWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.List(context.TODO(), elevenlabs.StudioProjectListParams{
		XiAPIKey: elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestStudioProjectDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.Delete(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectDeleteParams{
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

func TestStudioProjectConvertWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.Convert(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectConvertParams{
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

func TestStudioProjectNewPronunciationDictionaryWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.NewPronunciationDictionary(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectNewPronunciationDictionaryParams{
			PronunciationDictionaryLocators: []elevenlabs.StudioProjectNewPronunciationDictionaryParamsPronunciationDictionaryLocator{{
				PronunciationDictionaryID: "pronunciation_dictionary_id",
				VersionID:                 elevenlabs.String("version_id"),
			}},
			InvalidateAffectedText: elevenlabs.Bool(false),
			XiAPIKey:               elevenlabs.String("xi-api-key"),
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

func TestStudioProjectUpdateContentWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.Projects.UpdateContent(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.StudioProjectUpdateContentParams{
			AutoConvert:     elevenlabs.Bool(true),
			FromContentJson: elevenlabs.String(`[{"name": "Chapter A", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "A", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "B", "type": "tts_node"}]}, {"sub_type": "h1", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "C", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "D", "type": "tts_node"}]}]}, {"name": "Chapter B", "blocks": [{"sub_type": "p", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "E", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "F", "type": "tts_node"}]}, {"sub_type": "h2", "nodes": [{"voice_id": "6lCwbsX1yVjD49QmpkT0", "text": "G", "type": "tts_node"}, {"voice_id": "6lCwbsX1yVjD49QmpkT1", "text": "H", "type": "tts_node"}]}]}]`),
			FromDocument:    io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			FromURL:         elevenlabs.String("https://blog.elevenlabs.io/the_first_ai_that_can_laugh/"),
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
