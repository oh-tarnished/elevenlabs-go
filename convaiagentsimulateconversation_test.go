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

func TestConvaiAgentSimulateConversationSimulateConversationWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.SimulateConversation.SimulateConversation(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentSimulateConversationSimulateConversationParams{
			SimulationSpecification: elevenlabs.ConversationSimulationSpecificationParam{
				SimulatedUserConfig: elevenlabs.AgentConfigAPIModelInputParam{
					DisableFirstMessageInterruptions: elevenlabs.Bool(false),
					DynamicVariables: elevenlabs.DynamicVariablesConfigParam{
						DynamicVariablePlaceholders: map[string]elevenlabs.DynamicVariablesConfigDynamicVariablePlaceholderUnionParam{
							"user_name": {
								OfString: elevenlabs.String("John Doe"),
							},
						},
					},
					FirstMessage: elevenlabs.String("Hello, how can I help you today?"),
					HinglishMode: elevenlabs.Bool(true),
					Language:     elevenlabs.String("en"),
					Prompt: elevenlabs.AgentConfigAPIModelInputPromptParam{
						BackupLlmConfig: elevenlabs.AgentConfigAPIModelInputPromptBackupLlmConfigUnionParam{
							OfDefault: &elevenlabs.BackupLlmDefaultParam{
								Preference: elevenlabs.BackupLlmDefaultPreferenceDefault,
							},
						},
						BuiltInTools: elevenlabs.AgentConfigAPIModelInputPromptBuiltInToolsParam{
							EndCall: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							LanguageDetection: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							PlayKeypadTouchTone: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							SkipTurn: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							TransferToAgent: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							TransferToNumber: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							VoicemailDetection: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
						},
						CascadeTimeoutSeconds: elevenlabs.Float(2),
						CustomLlm: elevenlabs.CustomLlmParam{
							URL: "url",
							APIKey: elevenlabs.ConvaiSecretLocatorParam{
								SecretID: "secret_id",
							},
							APIType:    elevenlabs.CustomLlmAPITypeChatCompletions,
							APIVersion: elevenlabs.String("api_version"),
							ModelID:    elevenlabs.String("model_id"),
							RequestHeaders: map[string]elevenlabs.CustomLlmRequestHeaderUnionParam{
								"foo": {
									OfString: elevenlabs.String("string"),
								},
							},
						},
						IgnoreDefaultPersonality: elevenlabs.Bool(true),
						KnowledgeBase: []elevenlabs.KnowledgeBaseLocatorParam{{
							ID:        "123",
							Name:      "My Knowledge Base",
							Type:      elevenlabs.KnowledgeBaseDocumentTypeFile,
							UsageMode: elevenlabs.DocumentUsageModeAuto,
						}},
						Llm:                elevenlabs.LlmGemini2_0Flash001,
						MaxTokens:          elevenlabs.Int(-1),
						McpServerIDs:       []string{"string"},
						NativeMcpServerIDs: []string{"string"},
						Prompt:             elevenlabs.String("You are a helpful assistant that can answer questions about the topic of the conversation."),
						Rag: elevenlabs.RagConfigParam{
							EmbeddingModel:             elevenlabs.EmbeddingModelE5Mistral7bInstruct,
							Enabled:                    elevenlabs.Bool(true),
							MaxDocumentsLength:         elevenlabs.Int(1),
							MaxRetrievedRagChunksCount: elevenlabs.Int(5),
							MaxVectorDistance:          elevenlabs.Float(0.5),
							QueryRewritePromptOverride: elevenlabs.String("You are given a conversation between a user and an assistant. Rewrite the last question to be self-contained and clear."),
						},
						ReasoningEffort: elevenlabs.LlmReasoningEffortNone,
						Temperature:     elevenlabs.Float(0),
						ThinkingBudget:  elevenlabs.Int(0),
						Timezone:        elevenlabs.String("timezone"),
						ToolIDs:         []string{"string"},
						Tools: []elevenlabs.AgentConfigAPIModelInputPromptToolUnionParam{{
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
						}},
					},
				},
				DynamicVariables: map[string]elevenlabs.ConversationSimulationSpecificationDynamicVariableUnionParam{
					"foo": {
						OfString: elevenlabs.String("string"),
					},
				},
				PartialConversationHistory: []elevenlabs.ConversationHistoryTranscriptInputParam{{
					Role:           elevenlabs.ConversationHistoryTranscriptInputRoleUser,
					TimeInCallSecs: 0,
					AgentMetadata: elevenlabs.AgentMetadataParam{
						AgentID:        "agent_id",
						BranchID:       elevenlabs.String("branch_id"),
						WorkflowNodeID: elevenlabs.String("workflow_node_id"),
					},
					ConversationTurnMetrics: elevenlabs.ConversationTurnMetricsParam{
						Metrics: map[string]elevenlabs.ConversationTurnMetricsMetricParam{
							"foo": {
								ElapsedTime: 0,
							},
						},
					},
					Feedback: elevenlabs.UserFeedbackParam{
						Score:          elevenlabs.UserFeedbackScoreLike,
						TimeInCallSecs: 0,
					},
					Interrupted: elevenlabs.Bool(true),
					LlmOverride: elevenlabs.String("llm_override"),
					LlmUsage: elevenlabs.ConversationHistoryTranscriptInputLlmUsageParam{
						ModelUsage: map[string]elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageParam{
							"foo": {
								Input: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								InputCacheRead: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								InputCacheWrite: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								OutputTotal: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
							},
						},
					},
					Message: elevenlabs.String("message"),
					MultivoiceMessage: elevenlabs.ConversationHistoryMultivoiceMessageModelParam{
						Parts: []elevenlabs.ConversationHistoryMultivoiceMessageModelPartParam{{
							Text:           "text",
							TimeInCallSecs: elevenlabs.Int(0),
							VoiceLabel:     elevenlabs.String("voice_label"),
						}},
					},
					OriginalMessage: elevenlabs.String("original_message"),
					RagRetrievalInfo: elevenlabs.RagRetrievalInfoParam{
						Chunks: []elevenlabs.RagRetrievalInfoChunkParam{{
							ChunkID:        "chunk_id",
							DocumentID:     "document_id",
							VectorDistance: 0,
						}},
						EmbeddingModel: elevenlabs.EmbeddingModelE5Mistral7bInstruct,
						RagLatencySecs: 0,
						RetrievalQuery: "retrieval_query",
					},
					SourceMedium: elevenlabs.ConversationHistoryTranscriptInputSourceMediumAudio,
					ToolCalls: []elevenlabs.ConversationHistoryTranscriptInputToolCallParam{{
						ParamsAsJson:      "params_as_json",
						RequestID:         "request_id",
						ToolHasBeenCalled: true,
						ToolName:          "tool_name",
						ToolDetails: elevenlabs.ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam{
							OfWebhook: &elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetailsParam{
								Method: "method",
								URL:    "url",
								Body:   elevenlabs.String("body"),
								Headers: map[string]string{
									"foo": "string",
								},
								PathParams: map[string]string{
									"foo": "string",
								},
								QueryParams: map[string]string{
									"foo": "string",
								},
								Type: elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetailsTypeWebhook,
							},
						},
						Type: elevenlabs.ToolTypeSystem,
					}},
					ToolResults: []elevenlabs.ConversationHistoryTranscriptInputToolResultUnionParam{{
						OfConversationHistoryTranscriptOtherToolsResultCommonModel: &elevenlabs.ConversationHistoryTranscriptOtherToolsResultCommonModelParam{
							IsError:           true,
							RequestID:         "request_id",
							ResultValue:       "result_value",
							ToolHasBeenCalled: true,
							ToolName:          "tool_name",
							DynamicVariableUpdates: []elevenlabs.DynamicVariableUpdateCommonModelParam{{
								NewValue:      "new_value",
								OldValue:      elevenlabs.String("old_value"),
								ToolName:      "tool_name",
								ToolRequestID: "tool_request_id",
								UpdatedAt:     0,
								VariableName:  "variable_name",
							}},
							ToolLatencySecs: elevenlabs.Float(0),
							Type:            elevenlabs.ConversationHistoryTranscriptOtherToolsResultCommonModelTypeClient,
						},
					}},
				}},
				ToolMockConfig: map[string]elevenlabs.ConversationSimulationSpecificationToolMockConfigParam{
					"foo": {
						DefaultIsError:     elevenlabs.Bool(true),
						DefaultReturnValue: elevenlabs.String("default_return_value"),
					},
				},
			},
			ExtraEvaluationCriteria: []elevenlabs.PromptEvaluationCriteriaParam{{
				ID:                     "1234567890",
				ConversationGoalPrompt: "You are a helpful assistant that can answer questions about the topic of the conversation.",
				Name:                   "Customer satisfaction check",
				Type:                   elevenlabs.PromptEvaluationCriteriaTypePrompt,
				UseKnowledgeBase:       elevenlabs.Bool(false),
			}},
			NewTurnsLimit: elevenlabs.Int(0),
			XiAPIKey:      elevenlabs.String("xi-api-key"),
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

func TestConvaiAgentSimulateConversationStreamWithOptionalParams(t *testing.T) {
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
	err := client.Convai.Agents.SimulateConversation.Stream(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentSimulateConversationStreamParams{
			SimulationSpecification: elevenlabs.ConversationSimulationSpecificationParam{
				SimulatedUserConfig: elevenlabs.AgentConfigAPIModelInputParam{
					DisableFirstMessageInterruptions: elevenlabs.Bool(false),
					DynamicVariables: elevenlabs.DynamicVariablesConfigParam{
						DynamicVariablePlaceholders: map[string]elevenlabs.DynamicVariablesConfigDynamicVariablePlaceholderUnionParam{
							"user_name": {
								OfString: elevenlabs.String("John Doe"),
							},
						},
					},
					FirstMessage: elevenlabs.String("Hello, how can I help you today?"),
					HinglishMode: elevenlabs.Bool(true),
					Language:     elevenlabs.String("en"),
					Prompt: elevenlabs.AgentConfigAPIModelInputPromptParam{
						BackupLlmConfig: elevenlabs.AgentConfigAPIModelInputPromptBackupLlmConfigUnionParam{
							OfDefault: &elevenlabs.BackupLlmDefaultParam{
								Preference: elevenlabs.BackupLlmDefaultPreferenceDefault,
							},
						},
						BuiltInTools: elevenlabs.AgentConfigAPIModelInputPromptBuiltInToolsParam{
							EndCall: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							LanguageDetection: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							PlayKeypadTouchTone: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							SkipTurn: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							TransferToAgent: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							TransferToNumber: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
							VoicemailDetection: elevenlabs.SystemToolConfigInputParam{
								Name: "end_call",
								Params: elevenlabs.SystemToolConfigInputParamsUnionParam{
									OfEndCall: &elevenlabs.EndCallToolConfigParam{
										SystemToolType: elevenlabs.EndCallToolConfigSystemToolTypeEndCall,
									},
								},
								Assignments: []elevenlabs.DynamicVariableAssignmentParam{{
									DynamicVariable: "user_name",
									ValuePath:       "user.name",
									Source:          elevenlabs.DynamicVariableAssignmentSourceResponse,
								}},
								Description:           elevenlabs.String(""),
								DisableInterruptions:  elevenlabs.Bool(true),
								ForcePreToolSpeech:    elevenlabs.Bool(true),
								ResponseTimeoutSecs:   elevenlabs.Int(0),
								ToolCallSound:         elevenlabs.ToolCallSoundTypeTyping,
								ToolCallSoundBehavior: elevenlabs.ToolCallSoundBehaviorAuto,
								Type:                  elevenlabs.SystemToolConfigInputTypeSystem,
							},
						},
						CascadeTimeoutSeconds: elevenlabs.Float(2),
						CustomLlm: elevenlabs.CustomLlmParam{
							URL: "url",
							APIKey: elevenlabs.ConvaiSecretLocatorParam{
								SecretID: "secret_id",
							},
							APIType:    elevenlabs.CustomLlmAPITypeChatCompletions,
							APIVersion: elevenlabs.String("api_version"),
							ModelID:    elevenlabs.String("model_id"),
							RequestHeaders: map[string]elevenlabs.CustomLlmRequestHeaderUnionParam{
								"foo": {
									OfString: elevenlabs.String("string"),
								},
							},
						},
						IgnoreDefaultPersonality: elevenlabs.Bool(true),
						KnowledgeBase: []elevenlabs.KnowledgeBaseLocatorParam{{
							ID:        "123",
							Name:      "My Knowledge Base",
							Type:      elevenlabs.KnowledgeBaseDocumentTypeFile,
							UsageMode: elevenlabs.DocumentUsageModeAuto,
						}},
						Llm:                elevenlabs.LlmGemini2_0Flash001,
						MaxTokens:          elevenlabs.Int(-1),
						McpServerIDs:       []string{"string"},
						NativeMcpServerIDs: []string{"string"},
						Prompt:             elevenlabs.String("You are a helpful assistant that can answer questions about the topic of the conversation."),
						Rag: elevenlabs.RagConfigParam{
							EmbeddingModel:             elevenlabs.EmbeddingModelE5Mistral7bInstruct,
							Enabled:                    elevenlabs.Bool(true),
							MaxDocumentsLength:         elevenlabs.Int(1),
							MaxRetrievedRagChunksCount: elevenlabs.Int(5),
							MaxVectorDistance:          elevenlabs.Float(0.5),
							QueryRewritePromptOverride: elevenlabs.String("You are given a conversation between a user and an assistant. Rewrite the last question to be self-contained and clear."),
						},
						ReasoningEffort: elevenlabs.LlmReasoningEffortNone,
						Temperature:     elevenlabs.Float(0),
						ThinkingBudget:  elevenlabs.Int(0),
						Timezone:        elevenlabs.String("timezone"),
						ToolIDs:         []string{"string"},
						Tools: []elevenlabs.AgentConfigAPIModelInputPromptToolUnionParam{{
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
						}},
					},
				},
				DynamicVariables: map[string]elevenlabs.ConversationSimulationSpecificationDynamicVariableUnionParam{
					"foo": {
						OfString: elevenlabs.String("string"),
					},
				},
				PartialConversationHistory: []elevenlabs.ConversationHistoryTranscriptInputParam{{
					Role:           elevenlabs.ConversationHistoryTranscriptInputRoleUser,
					TimeInCallSecs: 0,
					AgentMetadata: elevenlabs.AgentMetadataParam{
						AgentID:        "agent_id",
						BranchID:       elevenlabs.String("branch_id"),
						WorkflowNodeID: elevenlabs.String("workflow_node_id"),
					},
					ConversationTurnMetrics: elevenlabs.ConversationTurnMetricsParam{
						Metrics: map[string]elevenlabs.ConversationTurnMetricsMetricParam{
							"foo": {
								ElapsedTime: 0,
							},
						},
					},
					Feedback: elevenlabs.UserFeedbackParam{
						Score:          elevenlabs.UserFeedbackScoreLike,
						TimeInCallSecs: 0,
					},
					Interrupted: elevenlabs.Bool(true),
					LlmOverride: elevenlabs.String("llm_override"),
					LlmUsage: elevenlabs.ConversationHistoryTranscriptInputLlmUsageParam{
						ModelUsage: map[string]elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageParam{
							"foo": {
								Input: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								InputCacheRead: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheReadParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								InputCacheWrite: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageInputCacheWriteParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
								OutputTotal: elevenlabs.ConversationHistoryTranscriptInputLlmUsageModelUsageOutputTotalParam{
									Price:  elevenlabs.Float(0),
									Tokens: elevenlabs.Int(0),
								},
							},
						},
					},
					Message: elevenlabs.String("message"),
					MultivoiceMessage: elevenlabs.ConversationHistoryMultivoiceMessageModelParam{
						Parts: []elevenlabs.ConversationHistoryMultivoiceMessageModelPartParam{{
							Text:           "text",
							TimeInCallSecs: elevenlabs.Int(0),
							VoiceLabel:     elevenlabs.String("voice_label"),
						}},
					},
					OriginalMessage: elevenlabs.String("original_message"),
					RagRetrievalInfo: elevenlabs.RagRetrievalInfoParam{
						Chunks: []elevenlabs.RagRetrievalInfoChunkParam{{
							ChunkID:        "chunk_id",
							DocumentID:     "document_id",
							VectorDistance: 0,
						}},
						EmbeddingModel: elevenlabs.EmbeddingModelE5Mistral7bInstruct,
						RagLatencySecs: 0,
						RetrievalQuery: "retrieval_query",
					},
					SourceMedium: elevenlabs.ConversationHistoryTranscriptInputSourceMediumAudio,
					ToolCalls: []elevenlabs.ConversationHistoryTranscriptInputToolCallParam{{
						ParamsAsJson:      "params_as_json",
						RequestID:         "request_id",
						ToolHasBeenCalled: true,
						ToolName:          "tool_name",
						ToolDetails: elevenlabs.ConversationHistoryTranscriptInputToolCallToolDetailsUnionParam{
							OfWebhook: &elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetailsParam{
								Method: "method",
								URL:    "url",
								Body:   elevenlabs.String("body"),
								Headers: map[string]string{
									"foo": "string",
								},
								PathParams: map[string]string{
									"foo": "string",
								},
								QueryParams: map[string]string{
									"foo": "string",
								},
								Type: elevenlabs.ConversationHistoryTranscriptToolCallWebhookDetailsTypeWebhook,
							},
						},
						Type: elevenlabs.ToolTypeSystem,
					}},
					ToolResults: []elevenlabs.ConversationHistoryTranscriptInputToolResultUnionParam{{
						OfConversationHistoryTranscriptOtherToolsResultCommonModel: &elevenlabs.ConversationHistoryTranscriptOtherToolsResultCommonModelParam{
							IsError:           true,
							RequestID:         "request_id",
							ResultValue:       "result_value",
							ToolHasBeenCalled: true,
							ToolName:          "tool_name",
							DynamicVariableUpdates: []elevenlabs.DynamicVariableUpdateCommonModelParam{{
								NewValue:      "new_value",
								OldValue:      elevenlabs.String("old_value"),
								ToolName:      "tool_name",
								ToolRequestID: "tool_request_id",
								UpdatedAt:     0,
								VariableName:  "variable_name",
							}},
							ToolLatencySecs: elevenlabs.Float(0),
							Type:            elevenlabs.ConversationHistoryTranscriptOtherToolsResultCommonModelTypeClient,
						},
					}},
				}},
				ToolMockConfig: map[string]elevenlabs.ConversationSimulationSpecificationToolMockConfigParam{
					"foo": {
						DefaultIsError:     elevenlabs.Bool(true),
						DefaultReturnValue: elevenlabs.String("default_return_value"),
					},
				},
			},
			ExtraEvaluationCriteria: []elevenlabs.PromptEvaluationCriteriaParam{{
				ID:                     "1234567890",
				ConversationGoalPrompt: "You are a helpful assistant that can answer questions about the topic of the conversation.",
				Name:                   "Customer satisfaction check",
				Type:                   elevenlabs.PromptEvaluationCriteriaTypePrompt,
				UseKnowledgeBase:       elevenlabs.Bool(false),
			}},
			NewTurnsLimit: elevenlabs.Int(0),
			XiAPIKey:      elevenlabs.String("xi-api-key"),
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
