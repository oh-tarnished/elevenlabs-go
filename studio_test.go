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

func TestStudioNewPodcastWithOptionalParams(t *testing.T) {
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
	_, err := client.Studio.NewPodcast(context.TODO(), elevenlabs.StudioNewPodcastParams{
		Mode: elevenlabs.StudioNewPodcastParamsModeUnion{
			OfPodcastConversationMode: &elevenlabs.StudioNewPodcastParamsModePodcastConversationMode{
				Conversation: elevenlabs.StudioNewPodcastParamsModePodcastConversationModeConversation{
					GuestVoiceID: "bYTqZQo3Jz7LQtmGTgwi",
					HostVoiceID:  "6lCwbsX1yVjD49QmpkTR",
				},
			},
		},
		ModelID: "eleven_multilingual_v2",
		Source: elevenlabs.StudioNewPodcastParamsSourceUnion{
			OfPodcastURLSource: &elevenlabs.PodcastURLSourceParam{
				URL: "https://en.wikipedia.org/wiki/Cognitive_science",
			},
		},
		ApplyTextNormalization: elevenlabs.StudioNewPodcastParamsApplyTextNormalizationAuto,
		CallbackURL:            elevenlabs.String("callback_url"),
		DurationScale:          elevenlabs.StudioNewPodcastParamsDurationScaleShort,
		Highlights:             []string{"Emphasize the importance of AI on education"},
		InstructionsPrompt:     elevenlabs.String("Ensure the podcast remains factual, accurate and appropriate for all audiences."),
		Intro:                  elevenlabs.String("Welcome to the podcast."),
		Language:               elevenlabs.String("en"),
		Outro:                  elevenlabs.String("Thank you for listening!"),
		QualityPreset:          elevenlabs.StudioNewPodcastParamsQualityPresetStandard,
		SafetyIdentifier:       elevenlabs.String("safety-identifier"),
		XiAPIKey:               elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
