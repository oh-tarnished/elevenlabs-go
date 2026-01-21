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

func TestConvaiAgentTestingNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.New(context.TODO(), elevenlabs.ConvaiAgentTestingNewParams{
		ChatHistory: []elevenlabs.ConversationHistoryTranscriptInputParam{{
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
		FailureExamples: []elevenlabs.AgentFailureResponseExampleParam{{
			Response: "response",
		}},
		Name:             "name",
		SuccessCondition: "success_condition",
		SuccessExamples: []elevenlabs.AgentSuccessfulResponseExampleParam{{
			Response: "response",
		}},
		DynamicVariables: map[string]elevenlabs.ConvaiAgentTestingNewParamsDynamicVariableUnion{
			"foo": {
				OfString: elevenlabs.String("string"),
			},
		},
		FromConversationMetadata: elevenlabs.TestFromConversationMetadataInputParam{
			AgentID:        "agent_id",
			ConversationID: "conversation_id",
			OriginalAgentReply: []elevenlabs.ConversationHistoryTranscriptInputParam{{
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
			WorkflowNodeID: elevenlabs.String("workflow_node_id"),
		},
		ToolCallParameters: elevenlabs.UnitTestToolCallEvaluationInputParam{
			Parameters: []elevenlabs.UnitTestToolCallParameter{{
				Eval: elevenlabs.UnitTestToolCallParameterEvalUnion{
					OfLlm: &elevenlabs.UnitTestToolCallParameterEvalLlm{
						Description: "description",
					},
				},
				Path: "path",
			}},
			ReferencedTool: elevenlabs.ReferencedToolCommonParam{
				ID:   "id",
				Type: elevenlabs.ReferencedToolCommonTypeSystem,
			},
			VerifyAbsence: elevenlabs.Bool(true),
		},
		Type:     elevenlabs.UnitTestCommonModelLlm,
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

func TestConvaiAgentTestingGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.Get(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiAgentTestingGetParams{
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

func TestConvaiAgentTestingUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.Update(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiAgentTestingUpdateParams{
			ChatHistory: []elevenlabs.ConversationHistoryTranscriptInputParam{{
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
			FailureExamples: []elevenlabs.AgentFailureResponseExampleParam{{
				Response: "response",
			}},
			Name:             "name",
			SuccessCondition: "success_condition",
			SuccessExamples: []elevenlabs.AgentSuccessfulResponseExampleParam{{
				Response: "response",
			}},
			DynamicVariables: map[string]elevenlabs.ConvaiAgentTestingUpdateParamsDynamicVariableUnion{
				"foo": {
					OfString: elevenlabs.String("string"),
				},
			},
			FromConversationMetadata: elevenlabs.TestFromConversationMetadataInputParam{
				AgentID:        "agent_id",
				ConversationID: "conversation_id",
				OriginalAgentReply: []elevenlabs.ConversationHistoryTranscriptInputParam{{
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
				WorkflowNodeID: elevenlabs.String("workflow_node_id"),
			},
			ToolCallParameters: elevenlabs.UnitTestToolCallEvaluationInputParam{
				Parameters: []elevenlabs.UnitTestToolCallParameter{{
					Eval: elevenlabs.UnitTestToolCallParameterEvalUnion{
						OfLlm: &elevenlabs.UnitTestToolCallParameterEvalLlm{
							Description: "description",
						},
					},
					Path: "path",
				}},
				ReferencedTool: elevenlabs.ReferencedToolCommonParam{
					ID:   "id",
					Type: elevenlabs.ReferencedToolCommonTypeSystem,
				},
				VerifyAbsence: elevenlabs.Bool(true),
			},
			Type:     elevenlabs.UnitTestCommonModelLlm,
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

func TestConvaiAgentTestingListWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.List(context.TODO(), elevenlabs.ConvaiAgentTestingListParams{
		Cursor:   elevenlabs.String("cursor"),
		PageSize: elevenlabs.Int(1),
		Search:   elevenlabs.String("search"),
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

func TestConvaiAgentTestingDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.Delete(
		context.TODO(),
		"TeaqRRdTcIfIu2i7BYfT",
		elevenlabs.ConvaiAgentTestingDeleteParams{
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

func TestConvaiAgentTestingGetSummariesWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.AgentTesting.GetSummaries(context.TODO(), elevenlabs.ConvaiAgentTestingGetSummariesParams{
		TestIDs:  []string{"test_id_1", "test_id_2"},
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
