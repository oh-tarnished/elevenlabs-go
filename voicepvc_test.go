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

func TestVoicePvcNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.New(context.TODO(), elevenlabs.VoicePvcNewParams{
		Language:    "en",
		Name:        "John Smith",
		Description: elevenlabs.String("An old American male voice with a slight hoarseness in his throat. Perfect for news."),
		Labels: map[string]string{
			"0":  "{",
			"1":  "\"",
			"2":  "l",
			"3":  "a",
			"4":  "n",
			"5":  "g",
			"6":  "u",
			"7":  "a",
			"8":  "g",
			"9":  "e",
			"10": "\"",
			"11": ":",
			"12": " ",
			"13": "\"",
			"14": "e",
			"15": "n",
			"16": "\"",
			"17": ",",
			"18": " ",
			"19": "\"",
			"20": "a",
			"21": "c",
			"22": "c",
			"23": "e",
			"24": "n",
			"25": "t",
			"26": "\"",
			"27": ":",
			"28": " ",
			"29": "\"",
			"30": "e",
			"31": "n",
			"32": "-",
			"33": "U",
			"34": "S",
			"35": "\"",
			"36": ",",
			"37": " ",
			"38": "\"",
			"39": "g",
			"40": "e",
			"41": "n",
			"42": "d",
			"43": "e",
			"44": "r",
			"45": "\"",
			"46": ":",
			"47": " ",
			"48": "\"",
			"49": "m",
			"50": "a",
			"51": "l",
			"52": "e",
			"53": "\"",
			"54": ",",
			"55": " ",
			"56": "\"",
			"57": "a",
			"58": "g",
			"59": "e",
			"60": "\"",
			"61": ":",
			"62": " ",
			"63": "\"",
			"64": "m",
			"65": "i",
			"66": "d",
			"67": "d",
			"68": "l",
			"69": "e",
			"70": "-",
			"71": "a",
			"72": "g",
			"73": "e",
			"74": "d",
			"75": "\"",
			"76": "}",
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

func TestVoicePvcEditWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Edit(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoicePvcEditParams{
			Description: elevenlabs.String("An old American male voice with a slight hoarseness in his throat. Perfect for news."),
			Labels: map[string]string{
				"0":  "{",
				"1":  "\"",
				"2":  "l",
				"3":  "a",
				"4":  "n",
				"5":  "g",
				"6":  "u",
				"7":  "a",
				"8":  "g",
				"9":  "e",
				"10": "\"",
				"11": ":",
				"12": " ",
				"13": "\"",
				"14": "e",
				"15": "n",
				"16": "\"",
				"17": ",",
				"18": " ",
				"19": "\"",
				"20": "a",
				"21": "c",
				"22": "c",
				"23": "e",
				"24": "n",
				"25": "t",
				"26": "\"",
				"27": ":",
				"28": " ",
				"29": "\"",
				"30": "e",
				"31": "n",
				"32": "-",
				"33": "U",
				"34": "S",
				"35": "\"",
				"36": ",",
				"37": " ",
				"38": "\"",
				"39": "g",
				"40": "e",
				"41": "n",
				"42": "d",
				"43": "e",
				"44": "r",
				"45": "\"",
				"46": ":",
				"47": " ",
				"48": "\"",
				"49": "m",
				"50": "a",
				"51": "l",
				"52": "e",
				"53": "\"",
				"54": ",",
				"55": " ",
				"56": "\"",
				"57": "a",
				"58": "g",
				"59": "e",
				"60": "\"",
				"61": ":",
				"62": " ",
				"63": "\"",
				"64": "m",
				"65": "i",
				"66": "d",
				"67": "d",
				"68": "l",
				"69": "e",
				"70": "-",
				"71": "a",
				"72": "g",
				"73": "e",
				"74": "d",
				"75": "\"",
				"76": "}",
			},
			Language: elevenlabs.String("en"),
			Name:     elevenlabs.String("John Smith"),
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

func TestVoicePvcRequestVerificationWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.RequestVerification(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoicePvcRequestVerificationParams{
			Files:     []io.Reader{io.Reader(bytes.NewBuffer([]byte("some file contents")))},
			ExtraText: elevenlabs.String("extra_text"),
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
}

func TestVoicePvcTrainWithOptionalParams(t *testing.T) {
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
	_, err := client.Voices.Pvc.Train(
		context.TODO(),
		"21m00Tcm4TlvDq8ikWAM",
		elevenlabs.VoicePvcTrainParams{
			ModelID:  elevenlabs.String("eleven_turbo_v2"),
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
