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

func TestMusicComposeWithOptionalParams(t *testing.T) {
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
	resp, err := client.Music.Compose(context.TODO(), elevenlabs.MusicComposeParams{
		OutputFormat: elevenlabs.AllowedOutputFormatsMP3_22050_32,
		CompositionPlan: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		FinetuneID:        elevenlabs.String("finetune_id"),
		ForceInstrumental: elevenlabs.Bool(true),
		ModelID:           elevenlabs.MusicComposeParamsModelIDMusicV1,
		MusicLengthMs:     elevenlabs.Int(3000),
		MusicPrompt: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		Prompt:                   elevenlabs.String("prompt"),
		RespectSectionsDurations: elevenlabs.Bool(true),
		Seed:                     elevenlabs.Int(0),
		SignWithC2pa:             elevenlabs.Bool(true),
		StoreForInpainting:       elevenlabs.Bool(true),
		XiAPIKey:                 elevenlabs.String("xi-api-key"),
	})
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

func TestMusicComposeWithDetailedResponseWithOptionalParams(t *testing.T) {
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
	resp, err := client.Music.ComposeWithDetailedResponse(context.TODO(), elevenlabs.MusicComposeWithDetailedResponseParams{
		OutputFormat: elevenlabs.AllowedOutputFormatsMP3_22050_32,
		CompositionPlan: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		FinetuneID:        elevenlabs.String("finetune_id"),
		ForceInstrumental: elevenlabs.Bool(true),
		ModelID:           elevenlabs.MusicComposeWithDetailedResponseParamsModelIDMusicV1,
		MusicLengthMs:     elevenlabs.Int(3000),
		MusicPrompt: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		Prompt:             elevenlabs.String("prompt"),
		Seed:               elevenlabs.Int(0),
		SignWithC2pa:       elevenlabs.Bool(true),
		StoreForInpainting: elevenlabs.Bool(true),
		WithTimestamps:     elevenlabs.Bool(true),
		XiAPIKey:           elevenlabs.String("xi-api-key"),
	})
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

func TestMusicGenerateCompositionPlanWithOptionalParams(t *testing.T) {
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
	_, err := client.Music.GenerateCompositionPlan(context.TODO(), elevenlabs.MusicGenerateCompositionPlanParams{
		Prompt:        "prompt",
		ModelID:       elevenlabs.MusicGenerateCompositionPlanParamsModelIDMusicV1,
		MusicLengthMs: elevenlabs.Int(3000),
		SourceCompositionPlan: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
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

func TestMusicSeparateStemsWithOptionalParams(t *testing.T) {
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
	resp, err := client.Music.SeparateStems(context.TODO(), elevenlabs.MusicSeparateStemsParams{
		File:            io.Reader(bytes.NewBuffer([]byte("some file contents"))),
		OutputFormat:    elevenlabs.AllowedOutputFormatsMP3_22050_32,
		SignWithC2pa:    elevenlabs.Bool(true),
		StemVariationID: elevenlabs.MusicSeparateStemsParamsStemVariationIDTwoStemsV1,
		XiAPIKey:        elevenlabs.String("xi-api-key"),
	})
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

func TestMusicStreamComposedMusicWithOptionalParams(t *testing.T) {
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
	resp, err := client.Music.StreamComposedMusic(context.TODO(), elevenlabs.MusicStreamComposedMusicParams{
		OutputFormat: elevenlabs.AllowedOutputFormatsMP3_22050_32,
		CompositionPlan: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		FinetuneID:        elevenlabs.String("finetune_id"),
		ForceInstrumental: elevenlabs.Bool(true),
		ModelID:           elevenlabs.MusicStreamComposedMusicParamsModelIDMusicV1,
		MusicLengthMs:     elevenlabs.Int(3000),
		MusicPrompt: elevenlabs.MusicPromptParam{
			NegativeGlobalStyles: []string{"metal", "hip-hop", "country"},
			PositiveGlobalStyles: []string{"pop", "rock", "jazz"},
			Sections: []elevenlabs.MusicPromptSectionParam{{
				DurationMs:          10000,
				Lines:               []string{"Verse 1 lyrics"},
				NegativeLocalStyles: []string{"metal", "hip-hop", "country"},
				PositiveLocalStyles: []string{"pop", "rock", "jazz"},
				SectionName:         "Verse 1",
				SourceFrom: elevenlabs.MusicPromptSectionSourceFromParam{
					Range: elevenlabs.TimeRangeParam{
						EndMs:   0,
						StartMs: 0,
					},
					SongID: "x",
					NegativeRanges: []elevenlabs.TimeRangeParam{{
						EndMs:   0,
						StartMs: 0,
					}},
				},
			}},
		},
		Prompt:             elevenlabs.String("prompt"),
		Seed:               elevenlabs.Int(0),
		StoreForInpainting: elevenlabs.Bool(true),
		XiAPIKey:           elevenlabs.String("xi-api-key"),
	})
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
