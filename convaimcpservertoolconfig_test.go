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

func TestConvaiMcpServerToolConfigNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.ToolConfigs.New(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerToolConfigNewParams{
			ToolName: "tool_name",
			Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
				DynamicVariable: "user_name",
				ValuePath:       "user.name",
				Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
			}},
			DisableInterruptions:  elevenlabs.Bool(true),
			ExecutionMode:         elevenlabs.ToolExecutionModeImmediate,
			ForcePreToolSpeech:    elevenlabs.Bool(true),
			ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
			ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
			XiAPIKey:              elevenlabs.String("xi-api-key"),
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

func TestConvaiMcpServerToolConfigGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.ToolConfigs.Get(
		context.TODO(),
		"tool_name",
		elevenlabs.ConvaiMcpServerToolConfigGetParams{
			McpServerID: "mcp_server_id",
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

func TestConvaiMcpServerToolConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.ToolConfigs.Update(
		context.TODO(),
		"tool_name",
		elevenlabs.ConvaiMcpServerToolConfigUpdateParams{
			McpServerID: "mcp_server_id",
			Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
				DynamicVariable: "user_name",
				ValuePath:       "user.name",
				Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
			}},
			DisableInterruptions:  elevenlabs.Bool(true),
			ExecutionMode:         elevenlabs.ToolExecutionModeImmediate,
			ForcePreToolSpeech:    elevenlabs.Bool(true),
			ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
			ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
			XiAPIKey:              elevenlabs.String("xi-api-key"),
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

func TestConvaiMcpServerToolConfigDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.ToolConfigs.Delete(
		context.TODO(),
		"tool_name",
		elevenlabs.ConvaiMcpServerToolConfigDeleteParams{
			McpServerID: "mcp_server_id",
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
