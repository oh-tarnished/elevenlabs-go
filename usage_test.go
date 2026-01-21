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

func TestUsageGetCharacterStatsWithOptionalParams(t *testing.T) {
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
	_, err := client.Usage.GetCharacterStats(context.TODO(), elevenlabs.UsageGetCharacterStatsParams{
		EndUnix:                 0,
		StartUnix:               0,
		AggregationBucketSize:   elevenlabs.Int(0),
		AggregationInterval:     elevenlabs.UsageGetCharacterStatsParamsAggregationIntervalHour,
		BreakdownType:           elevenlabs.UsageGetCharacterStatsParamsBreakdownTypeNone,
		IncludeWorkspaceMetrics: elevenlabs.Bool(true),
		Metric:                  elevenlabs.UsageGetCharacterStatsParamsMetricCredits,
		XiAPIKey:                elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
