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

func TestPronunciationDictionaryGetWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.Get(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.PronunciationDictionaryGetParams{
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

func TestPronunciationDictionaryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.Update(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.PronunciationDictionaryUpdateParams{
			Archived: elevenlabs.Bool(true),
			Name:     elevenlabs.String("My Dictionary"),
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

func TestPronunciationDictionaryListWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.List(context.TODO(), elevenlabs.PronunciationDictionaryListParams{
		Cursor:        elevenlabs.String("cursor"),
		PageSize:      elevenlabs.Int(1),
		Sort:          elevenlabs.PronunciationDictionaryListParamsSortCreationTimeUnix,
		SortDirection: elevenlabs.String("descending"),
		XiAPIKey:      elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPronunciationDictionaryAddRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.AddRules(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.PronunciationDictionaryAddRulesParams{
			Rules: []elevenlabs.PronunciationDictionaryAddRulesParamsRuleUnion{{
				OfPronunciationDictionaryAliasRuleRequestModel: &elevenlabs.PronunciationDictionaryAliasRuleParam{
					Alias:           "tie-land",
					StringToReplace: "Thailand",
				},
			}},
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

func TestPronunciationDictionaryNewFromFileWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.NewFromFile(context.TODO(), elevenlabs.PronunciationDictionaryNewFromFileParams{
		Name:            "My Dictionary",
		Description:     elevenlabs.String("Contains pronunciation's of our character names"),
		File:            io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		WorkspaceAccess: elevenlabs.PronunciationDictionaryNewFromFileParamsWorkspaceAccessViewer,
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

func TestPronunciationDictionaryNewFromRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.NewFromRules(context.TODO(), elevenlabs.PronunciationDictionaryNewFromRulesParams{
		Name: "My Dictionary",
		Rules: []elevenlabs.PronunciationDictionaryNewFromRulesParamsRuleUnion{{
			OfPronunciationDictionaryAliasRuleRequestModel: &elevenlabs.PronunciationDictionaryAliasRuleParam{
				Alias:           "tie-land",
				StringToReplace: "Thailand",
			},
		}},
		Description:     elevenlabs.String("Contains pronunciation's of our character names"),
		WorkspaceAccess: elevenlabs.PronunciationDictionaryNewFromRulesParamsWorkspaceAccessViewer,
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

func TestPronunciationDictionaryDownloadPlsWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.DownloadPls(
		context.TODO(),
		"BdF0s0aZ3oFoKnDYdTox",
		elevenlabs.PronunciationDictionaryDownloadPlsParams{
			DictionaryID: "21m00Tcm4TlvDq8ikWAM",
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

func TestPronunciationDictionaryRemoveRulesWithOptionalParams(t *testing.T) {
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
	_, err := client.PronunciationDictionaries.RemoveRules(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.PronunciationDictionaryRemoveRulesParams{
			RuleStrings: []string{"string"},
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
