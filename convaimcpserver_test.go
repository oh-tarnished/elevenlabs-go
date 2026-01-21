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

func TestConvaiMcpServerNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.New(context.TODO(), elevenlabs.ConvaiMcpServerNewParams{
		Config: elevenlabs.ConvaiMcpServerNewParamsConfig{
			Name: "name",
			URL: elevenlabs.ConvaiMcpServerNewParamsConfigURLUnion{
				OfString: elevenlabs.String("string"),
			},
			ApprovalPolicy:       elevenlabs.McpApprovalPolicyAutoApproveAll,
			Description:          elevenlabs.String("description"),
			DisableCompression:   elevenlabs.Bool(true),
			DisableInterruptions: elevenlabs.Bool(true),
			ExecutionMode:        elevenlabs.ToolExecutionModeImmediate,
			ForcePreToolSpeech:   elevenlabs.Bool(true),
			RequestHeaders: map[string]elevenlabs.ConvaiMcpServerNewParamsConfigRequestHeaderUnion{
				"foo": {
					OfString: elevenlabs.String("string"),
				},
			},
			SecretToken: elevenlabs.ConvaiMcpServerNewParamsConfigSecretTokenUnion{
				OfConvAISecretLocator: &elevenlabs.ConvaiSecretLocatorParam{
					SecretID: "secret_id",
				},
			},
			ToolApprovalHashes: []elevenlabs.McpToolApprovalHashParam{{
				ToolHash:       "tool_hash",
				ToolName:       "tool_name",
				ApprovalPolicy: elevenlabs.McpToolApprovalPolicyAutoApproved,
			}},
			ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
			ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
			ToolConfigOverrides: []elevenlabs.McpToolConfigOverrideParam{{
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
			}},
			Transport: elevenlabs.McpServerTransportSse,
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

func TestConvaiMcpServerGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.Get(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerGetParams{
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

func TestConvaiMcpServerUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.Update(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerUpdateParams{
			ApprovalPolicy:       elevenlabs.McpApprovalPolicyAutoApproveAll,
			DisableCompression:   elevenlabs.Bool(true),
			DisableInterruptions: elevenlabs.Bool(true),
			ExecutionMode:        elevenlabs.ToolExecutionModeImmediate,
			ForcePreToolSpeech:   elevenlabs.Bool(true),
			RequestHeaders: map[string]elevenlabs.ConvaiMcpServerUpdateParamsRequestHeaderUnion{
				"foo": {
					OfString: elevenlabs.String("string"),
				},
			},
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

func TestConvaiMcpServerListWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.List(context.TODO(), elevenlabs.ConvaiMcpServerListParams{
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

func TestConvaiMcpServerDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.Delete(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerDeleteParams{
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

func TestConvaiMcpServerListToolsWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.ListTools(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerListToolsParams{
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

func TestConvaiMcpServerUpdateApprovalPolicyWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.McpServers.UpdateApprovalPolicy(
		context.TODO(),
		"mcp_server_id",
		elevenlabs.ConvaiMcpServerUpdateApprovalPolicyParams{
			ApprovalPolicy: elevenlabs.McpApprovalPolicyAutoApproveAll,
			XiAPIKey:       elevenlabs.String("xi-api-key"),
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
