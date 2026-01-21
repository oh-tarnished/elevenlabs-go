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

func TestConvaiToolNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.New(context.TODO(), elevenlabs.ConvaiToolNewParams{
		ToolRequestModel: elevenlabs.ToolRequestModelParam{
			ToolConfig: elevenlabs.ToolRequestModelToolConfigUnionParam{
				OfWebhook: &elevenlabs.WebhookToolConfigInputParam{
					APISchema: elevenlabs.WebhookToolConfigInputAPISchemaParam{
						URL: "https://example.com/agents/{agent_id}",
						AuthConnection: elevenlabs.AuthConnectionLocatorParam{
							AuthConnectionID: "auth_connection_id",
						},
						ContentType: "application/json",
						Method:      "GET",
						PathParamsSchema: map[string]elevenlabs.WebhookToolConfigInputAPISchemaPathParamsSchemaParam{
							"agent_id": {
								Type: "string",
								ConstantValue: elevenlabs.WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam{
									OfString: elevenlabs.String("string"),
								},
								Description:      elevenlabs.String("A user-provided message"),
								DynamicVariable:  elevenlabs.String("dynamic_variable"),
								Enum:             []string{"string"},
								IsSystemProvided: elevenlabs.Bool(true),
							},
						},
						QueryParamsSchema: elevenlabs.QueryParamsJsonSchema{
							Properties: map[string]elevenlabs.QueryParamsJsonSchemaProperty{
								"foo": {
									Type: "string",
									ConstantValue: elevenlabs.QueryParamsJsonSchemaPropertyConstantValueUnion{
										OfString: elevenlabs.String("string"),
									},
									Description:      elevenlabs.String("A user-provided message"),
									DynamicVariable:  elevenlabs.String("dynamic_variable"),
									Enum:             []string{"string"},
									IsSystemProvided: elevenlabs.Bool(true),
								},
							},
							Required: []string{"string"},
						},
						RequestBodySchema: elevenlabs.ObjectJsonSchemaPropertyInputParam{
							Description: elevenlabs.String("description"),
							Properties: map[string]elevenlabs.ObjectJsonSchemaPropertyInputPropertyUnionParam{
								"foo": {
									OfLiteralJsonSchemaProperty: &elevenlabs.ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam{
										Type: "string",
										ConstantValue: elevenlabs.ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam{
											OfString: elevenlabs.String("string"),
										},
										Description:      elevenlabs.String("A user-provided message"),
										DynamicVariable:  elevenlabs.String("dynamic_variable"),
										Enum:             []string{"string"},
										IsSystemProvided: elevenlabs.Bool(true),
									},
								},
							},
							Required: []string{"string"},
							Type:     elevenlabs.ObjectJsonSchemaPropertyInputTypeObject,
						},
						RequestHeaders: map[string]elevenlabs.WebhookToolConfigInputAPISchemaRequestHeaderUnionParam{
							"Authorization": {
								OfString: elevenlabs.String("Bearer {api_key}"),
							},
						},
					},
					Description: "description",
					Name:        "name",
					Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
						DynamicVariable: "user_name",
						ValuePath:       "user.name",
						Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
					}},
					DisableInterruptions: elevenlabs.Bool(true),
					DynamicVariables: elevenlabs.DynamicVariablesConfigParam{
						DynamicVariablePlaceholders: map[string]elevenlabs.DynamicVariablesConfigDynamicVariablePlaceholderUnionParam{
							"user_name": {
								OfString: elevenlabs.String("John Doe"),
							},
						},
					},
					ExecutionMode:         elevenlabs.ToolExecutionModeImmediate,
					ForcePreToolSpeech:    elevenlabs.Bool(true),
					ResponseTimeoutSecs:   elevenlabs.Int(20),
					ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
					ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
					Type:                  elevenlabs.WebhookToolConfigInputTypeWebhook,
				},
			},
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

func TestConvaiToolGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.Get(
		context.TODO(),
		"tool_id",
		elevenlabs.ConvaiToolGetParams{
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

func TestConvaiToolUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.Update(
		context.TODO(),
		"tool_id",
		elevenlabs.ConvaiToolUpdateParams{
			ToolRequestModel: elevenlabs.ToolRequestModelParam{
				ToolConfig: elevenlabs.ToolRequestModelToolConfigUnionParam{
					OfWebhook: &elevenlabs.WebhookToolConfigInputParam{
						APISchema: elevenlabs.WebhookToolConfigInputAPISchemaParam{
							URL: "https://example.com/agents/{agent_id}",
							AuthConnection: elevenlabs.AuthConnectionLocatorParam{
								AuthConnectionID: "auth_connection_id",
							},
							ContentType: "application/json",
							Method:      "GET",
							PathParamsSchema: map[string]elevenlabs.WebhookToolConfigInputAPISchemaPathParamsSchemaParam{
								"agent_id": {
									Type: "string",
									ConstantValue: elevenlabs.WebhookToolConfigInputAPISchemaPathParamsSchemaConstantValueUnionParam{
										OfString: elevenlabs.String("string"),
									},
									Description:      elevenlabs.String("A user-provided message"),
									DynamicVariable:  elevenlabs.String("dynamic_variable"),
									Enum:             []string{"string"},
									IsSystemProvided: elevenlabs.Bool(true),
								},
							},
							QueryParamsSchema: elevenlabs.QueryParamsJsonSchema{
								Properties: map[string]elevenlabs.QueryParamsJsonSchemaProperty{
									"foo": {
										Type: "string",
										ConstantValue: elevenlabs.QueryParamsJsonSchemaPropertyConstantValueUnion{
											OfString: elevenlabs.String("string"),
										},
										Description:      elevenlabs.String("A user-provided message"),
										DynamicVariable:  elevenlabs.String("dynamic_variable"),
										Enum:             []string{"string"},
										IsSystemProvided: elevenlabs.Bool(true),
									},
								},
								Required: []string{"string"},
							},
							RequestBodySchema: elevenlabs.ObjectJsonSchemaPropertyInputParam{
								Description: elevenlabs.String("description"),
								Properties: map[string]elevenlabs.ObjectJsonSchemaPropertyInputPropertyUnionParam{
									"foo": {
										OfLiteralJsonSchemaProperty: &elevenlabs.ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyParam{
											Type: "string",
											ConstantValue: elevenlabs.ObjectJsonSchemaPropertyInputPropertyLiteralJsonSchemaPropertyConstantValueUnionParam{
												OfString: elevenlabs.String("string"),
											},
											Description:      elevenlabs.String("A user-provided message"),
											DynamicVariable:  elevenlabs.String("dynamic_variable"),
											Enum:             []string{"string"},
											IsSystemProvided: elevenlabs.Bool(true),
										},
									},
								},
								Required: []string{"string"},
								Type:     elevenlabs.ObjectJsonSchemaPropertyInputTypeObject,
							},
							RequestHeaders: map[string]elevenlabs.WebhookToolConfigInputAPISchemaRequestHeaderUnionParam{
								"Authorization": {
									OfString: elevenlabs.String("Bearer {api_key}"),
								},
							},
						},
						Description: "description",
						Name:        "name",
						Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
							DynamicVariable: "user_name",
							ValuePath:       "user.name",
							Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
						}},
						DisableInterruptions: elevenlabs.Bool(true),
						DynamicVariables: elevenlabs.DynamicVariablesConfigParam{
							DynamicVariablePlaceholders: map[string]elevenlabs.DynamicVariablesConfigDynamicVariablePlaceholderUnionParam{
								"user_name": {
									OfString: elevenlabs.String("John Doe"),
								},
							},
						},
						ExecutionMode:         elevenlabs.ToolExecutionModeImmediate,
						ForcePreToolSpeech:    elevenlabs.Bool(true),
						ResponseTimeoutSecs:   elevenlabs.Int(20),
						ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
						ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
						Type:                  elevenlabs.WebhookToolConfigInputTypeWebhook,
					},
				},
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

func TestConvaiToolListWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.List(context.TODO(), elevenlabs.ConvaiToolListParams{
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

func TestConvaiToolDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.Delete(
		context.TODO(),
		"tool_id",
		elevenlabs.ConvaiToolDeleteParams{
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

func TestConvaiToolListDependentAgentsWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Tools.ListDependentAgents(
		context.TODO(),
		"tool_id",
		elevenlabs.ConvaiToolListDependentAgentsParams{
			Cursor:   elevenlabs.String("cursor"),
			PageSize: elevenlabs.Int(1),
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
