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

func TestConvaiAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.New(context.TODO(), elevenlabs.ConvaiAgentNewParams{
		ConversationConfig: elevenlabs.ConversationalConfigAPIModelInputParam{
			Agent: elevenlabs.AgentConfigAPIModelInputParam{
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
			Asr: elevenlabs.AsrConversationalConfigParam{
				Keywords:             []string{"hello", "world"},
				Provider:             elevenlabs.AsrConversationalConfigProviderElevenlabs,
				Quality:              elevenlabs.AsrConversationalConfigQualityHigh,
				UserInputAudioFormat: elevenlabs.AsrConversationalConfigUserInputAudioFormatPcm16000,
			},
			Conversation: elevenlabs.ConversationConfigParam{
				ClientEvents:       []elevenlabs.ClientEvent{elevenlabs.ClientEventAudio, elevenlabs.ClientEventInterruption},
				MaxDurationSeconds: elevenlabs.Int(600),
				MonitoringEnabled:  elevenlabs.Bool(true),
				MonitoringEvents:   []elevenlabs.ClientEvent{elevenlabs.ClientEventConversationInitiationMetadata},
				TextOnly:           elevenlabs.Bool(true),
			},
			LanguagePresets: map[string]elevenlabs.ConversationalConfigAPIModelInputLanguagePresetParam{
				"foo": {
					Overrides: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetOverridesParam{
						Agent: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetOverridesAgentParam{
							FirstMessage: elevenlabs.String("Hello, how can I help you today?"),
							Language:     elevenlabs.String("en"),
							Prompt: elevenlabs.PromptAgentAPIModelOverrideParam{
								Llm:                elevenlabs.LlmGemini2_0Flash001,
								NativeMcpServerIDs: []string{"string"},
								Prompt:             elevenlabs.String("You are a helpful assistant that can answer questions about the topic of the conversation."),
							},
						},
						Conversation: elevenlabs.ConversationConfigOverrideParam{
							TextOnly: elevenlabs.Bool(true),
						},
						Tts: elevenlabs.TtsConversationalConfigOverrideParam{
							SimilarityBoost: elevenlabs.Float(0.8),
							Speed:           elevenlabs.Float(1),
							Stability:       elevenlabs.Float(0.5),
							VoiceID:         elevenlabs.String("cjVigY5qzO86Huf0OWal"),
						},
						Turn: elevenlabs.TurnConfigOverrideParam{
							SoftTimeoutConfig: elevenlabs.TurnConfigOverrideSoftTimeoutConfigParam{
								Message: elevenlabs.String("Hhmmmm...yeah."),
							},
						},
					},
					FirstMessageTranslation: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetFirstMessageTranslationParam{
						SourceHash: "source_hash",
						Text:       "text",
					},
					SoftTimeoutTranslation: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetSoftTimeoutTranslationParam{
						SourceHash: "source_hash",
						Text:       "text",
					},
				},
			},
			Tts: elevenlabs.ConversationalConfigAPIModelInputTtsParam{
				AgentOutputAudioFormat:   elevenlabs.TtsOutputFormatPcm16000,
				ModelID:                  elevenlabs.TtsConversationalModelElevenTurboV2,
				OptimizeStreamingLatency: 3,
				PronunciationDictionaryLocators: []elevenlabs.PydanticPronunciationDictionaryVersionLocatorParam{{
					PronunciationDictionaryID: "pronunciation_dictionary_id",
					VersionID:                 elevenlabs.String("version_id"),
				}},
				SimilarityBoost: elevenlabs.Float(0.8),
				Speed:           elevenlabs.Float(1),
				Stability:       elevenlabs.Float(0.5),
				SupportedVoices: []elevenlabs.SupportedVoiceParam{{
					Label:                    "x",
					VoiceID:                  "x",
					Description:              elevenlabs.String("description"),
					Language:                 elevenlabs.String("language"),
					ModelFamily:              elevenlabs.SupportedVoiceModelFamilyTurbo,
					OptimizeStreamingLatency: 0,
					SimilarityBoost:          elevenlabs.Float(0),
					Speed:                    elevenlabs.Float(0.7),
					Stability:                elevenlabs.Float(0),
				}},
				TextNormalisationType: elevenlabs.TextNormalisationTypeSystemPrompt,
				VoiceID:               elevenlabs.String("cjVigY5qzO86Huf0OWal"),
			},
			Turn: elevenlabs.TurnConfigParam{
				InitialWaitTime:       elevenlabs.Float(0),
				Mode:                  elevenlabs.TurnConfigModeTurn,
				SilenceEndCallTimeout: elevenlabs.Float(-1),
				SoftTimeoutConfig: elevenlabs.TurnConfigSoftTimeoutConfigParam{
					Message:                elevenlabs.String("Hhmmmm...yeah."),
					TimeoutSeconds:         elevenlabs.Float(-1),
					UseLlmGeneratedMessage: elevenlabs.Bool(false),
				},
				SpellingPatience: elevenlabs.TurnConfigSpellingPatienceAuto,
				TurnEagerness:    elevenlabs.TurnConfigTurnEagernessNormal,
				TurnTimeout:      elevenlabs.Float(7),
			},
			Vad: elevenlabs.VadConfigParam{
				BackgroundVoiceDetection: elevenlabs.Bool(false),
			},
		},
		Name: elevenlabs.String("My agent"),
		PlatformSettings: elevenlabs.AgentPlatformSettingsRequestModelParam{
			Archived: elevenlabs.Bool(true),
			Auth: elevenlabs.AuthSettingsParam{
				Allowlist: []elevenlabs.AllowlistItemParam{{
					Hostname: "https://example.com",
				}},
				EnableAuth:     elevenlabs.Bool(true),
				ShareableToken: elevenlabs.String("1234567890"),
			},
			CallLimits: elevenlabs.AgentCallLimitsParam{
				AgentConcurrencyLimit: elevenlabs.Int(-1),
				BurstingEnabled:       elevenlabs.Bool(true),
				DailyLimit:            elevenlabs.Int(100000),
			},
			DataCollection: map[string]elevenlabs.AgentPlatformSettingsRequestModelDataCollectionParam{
				"foo": {
					Type: "string",
					ConstantValue: elevenlabs.AgentPlatformSettingsRequestModelDataCollectionConstantValueUnionParam{
						OfString: elevenlabs.String("string"),
					},
					Description:      elevenlabs.String("A user-provided message"),
					DynamicVariable:  elevenlabs.String("dynamic_variable"),
					Enum:             []string{"string"},
					IsSystemProvided: elevenlabs.Bool(true),
				},
			},
			Evaluation: elevenlabs.EvaluationSettingsParam{
				Criteria: []elevenlabs.PromptEvaluationCriteriaParam{{
					ID:                     "1234567890",
					ConversationGoalPrompt: "You are a helpful assistant that can answer questions about the topic of the conversation.",
					Name:                   "Customer satisfaction check",
					Type:                   elevenlabs.PromptEvaluationCriteriaTypePrompt,
					UseKnowledgeBase:       elevenlabs.Bool(false),
				}},
			},
			Guardrails: elevenlabs.AgentPlatformSettingsRequestModelGuardrailsParam{
				Moderation: elevenlabs.AgentPlatformSettingsRequestModelGuardrailsModerationParam{
					Config: elevenlabs.ModerationConfigParam{
						Harassment: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						HarassmentThreatening: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						Hate: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						HateThreatening: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						SelfHarm: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						SelfHarmInstructions: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						SelfHarmIntent: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						Sexual: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						SexualMinors: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						Violence: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
						ViolenceGraphic: elevenlabs.ThresholdGuardrailParam{
							IsEnabled: elevenlabs.Bool(true),
							Threshold: elevenlabs.Float(0),
						},
					},
				},
				Version: "1",
			},
			Overrides: elevenlabs.AgentPlatformSettingsRequestModelOverridesParam{
				ConversationConfigOverride: elevenlabs.AgentPlatformSettingsRequestModelOverridesConversationConfigOverrideParam{
					Agent: elevenlabs.AgentConfigOverrideConfigParam{
						FirstMessage: elevenlabs.Bool(true),
						Language:     elevenlabs.Bool(true),
						Prompt: elevenlabs.AgentConfigOverrideConfigPromptParam{
							Llm:                elevenlabs.Bool(true),
							NativeMcpServerIDs: elevenlabs.Bool(true),
							Prompt:             elevenlabs.Bool(true),
						},
					},
					Conversation: elevenlabs.ConversationConfigOverrideConfigParam{
						TextOnly: elevenlabs.Bool(true),
					},
					Tts: elevenlabs.TtsConversationalConfigOverrideConfigParam{
						SimilarityBoost: elevenlabs.Bool(true),
						Speed:           elevenlabs.Bool(true),
						Stability:       elevenlabs.Bool(true),
						VoiceID:         elevenlabs.Bool(true),
					},
					Turn: elevenlabs.TurnConfigOverrideConfigParam{
						SoftTimeoutConfig: elevenlabs.TurnConfigOverrideConfigSoftTimeoutConfigParam{
							Message: elevenlabs.Bool(true),
						},
					},
				},
				CustomLlmExtraBody: elevenlabs.Bool(true),
				EnableConversationInitiationClientDataFromWebhook: elevenlabs.Bool(true),
			},
			Privacy: elevenlabs.PrivacyConfigParam{
				ApplyToExistingConversations: elevenlabs.Bool(false),
				DeleteAudio:                  elevenlabs.Bool(false),
				DeleteTranscriptAndPii:       elevenlabs.Bool(false),
				RecordVoice:                  elevenlabs.Bool(true),
				RetentionDays:                elevenlabs.Int(-1),
				ZeroRetentionMode:            elevenlabs.Bool(false),
			},
			Testing: elevenlabs.AgentTestingSettingsParam{
				AttachedTests: []elevenlabs.AgentTestingSettingsAttachedTestParam{{
					TestID:         "test_123",
					WorkflowNodeID: elevenlabs.String("node_abc"),
				}, {
					TestID:         "test_456",
					WorkflowNodeID: elevenlabs.String("workflow_node_id"),
				}},
			},
			Widget: elevenlabs.AgentPlatformSettingsRequestModelWidgetParam{
				ActionText:     elevenlabs.String("action_text"),
				AlwaysExpanded: elevenlabs.Bool(true),
				Avatar: elevenlabs.AgentPlatformSettingsRequestModelWidgetAvatarUnionParam{
					OfOrbAvatar: &elevenlabs.OrbAvatarParam{
						Color1: elevenlabs.String("#2792dc"),
						Color2: elevenlabs.String("#9ce6e6"),
						Type:   elevenlabs.OrbAvatarTypeOrb,
					},
				},
				BgColor:                       elevenlabs.String("bg_color"),
				BorderColor:                   elevenlabs.String("border_color"),
				BorderRadius:                  elevenlabs.Int(0),
				BtnColor:                      elevenlabs.String("btn_color"),
				BtnRadius:                     elevenlabs.Int(0),
				BtnTextColor:                  elevenlabs.String("btn_text_color"),
				ConversationModeToggleEnabled: elevenlabs.Bool(true),
				CustomAvatarPath:              elevenlabs.String("https://example.com/avatar.png"),
				DefaultExpanded:               elevenlabs.Bool(true),
				DisableBanner:                 elevenlabs.Bool(true),
				EndCallText:                   elevenlabs.String("end_call_text"),
				EndFeedback: elevenlabs.WidgetEndFeedbackConfigParam{
					Type: elevenlabs.WidgetEndFeedbackConfigTypeRating,
				},
				ExpandText:   elevenlabs.String("expand_text"),
				Expandable:   elevenlabs.WidgetExpandableNever,
				FeedbackMode: elevenlabs.WidgetFeedbackModeNone,
				FocusColor:   elevenlabs.String("focus_color"),
				LanguagePresets: map[string]elevenlabs.AgentPlatformSettingsRequestModelWidgetLanguagePresetParam{
					"foo": {
						TermsHTML: elevenlabs.String("terms_html"),
						TermsKey:  elevenlabs.String("terms_key"),
						TermsText: elevenlabs.String("terms_text"),
						TermsTranslation: elevenlabs.AgentPlatformSettingsRequestModelWidgetLanguagePresetTermsTranslationParam{
							SourceHash: "source_hash",
							Text:       "text",
						},
						TextContents: elevenlabs.WidgetTextContentsParam{
							AcceptTerms:                     elevenlabs.String("accept_terms"),
							AgentEndedConversation:          elevenlabs.String("agent_ended_conversation"),
							ChangeLanguage:                  elevenlabs.String("change_language"),
							ChattingStatus:                  elevenlabs.String("chatting_status"),
							Collapse:                        elevenlabs.String("collapse"),
							ConnectingStatus:                elevenlabs.String("connecting_status"),
							ConversationID:                  elevenlabs.String("conversation_id"),
							Copied:                          elevenlabs.String("copied"),
							CopyID:                          elevenlabs.String("copy_id"),
							DismissTerms:                    elevenlabs.String("dismiss_terms"),
							EndCall:                         elevenlabs.String("end_call"),
							ErrorOccurred:                   elevenlabs.String("error_occurred"),
							Expand:                          elevenlabs.String("expand"),
							FollowUpFeedbackPlaceholder:     elevenlabs.String("follow_up_feedback_placeholder"),
							GoBack:                          elevenlabs.String("go_back"),
							InitiateFeedback:                elevenlabs.String("initiate_feedback"),
							InputLabel:                      elevenlabs.String("input_label"),
							InputPlaceholder:                elevenlabs.String("input_placeholder"),
							InputPlaceholderNewConversation: elevenlabs.String("input_placeholder_new_conversation"),
							InputPlaceholderTextOnly:        elevenlabs.String("input_placeholder_text_only"),
							ListeningStatus:                 elevenlabs.String("listening_status"),
							MainLabel:                       elevenlabs.String("main_label"),
							MuteMicrophone:                  elevenlabs.String("mute_microphone"),
							NewCall:                         elevenlabs.String("new_call"),
							RequestFollowUpFeedback:         elevenlabs.String("request_follow_up_feedback"),
							SpeakingStatus:                  elevenlabs.String("speaking_status"),
							StartCall:                       elevenlabs.String("start_call"),
							StartChat:                       elevenlabs.String("start_chat"),
							Submit:                          elevenlabs.String("submit"),
							ThanksForFeedback:               elevenlabs.String("thanks_for_feedback"),
							ThanksForFeedbackDetails:        elevenlabs.String("thanks_for_feedback_details"),
							UserEndedConversation:           elevenlabs.String("user_ended_conversation"),
						},
					},
				},
				LanguageSelector:      elevenlabs.Bool(false),
				ListeningText:         elevenlabs.String("listening_text"),
				MarkdownLinkAllowHTTP: elevenlabs.Bool(true),
				MarkdownLinkAllowedHosts: []elevenlabs.AllowlistItemParam{{
					Hostname: "hostname",
				}},
				MarkdownLinkIncludeWww:  elevenlabs.Bool(true),
				MicMutingEnabled:        elevenlabs.Bool(true),
				OverrideLink:            elevenlabs.String("override_link"),
				Placement:               elevenlabs.WidgetPlacementTopLeft,
				ShareablePageShowTerms:  elevenlabs.Bool(true),
				ShareablePageText:       elevenlabs.String("shareable_page_text"),
				ShowAvatarWhenCollapsed: elevenlabs.Bool(true),
				SpeakingText:            elevenlabs.String("speaking_text"),
				StartCallText:           elevenlabs.String("start_call_text"),
				Styles: elevenlabs.WidgetStylesParam{
					Accent:              elevenlabs.String("accent"),
					AccentActive:        elevenlabs.String("accent_active"),
					AccentBorder:        elevenlabs.String("accent_border"),
					AccentHover:         elevenlabs.String("accent_hover"),
					AccentPrimary:       elevenlabs.String("accent_primary"),
					AccentSubtle:        elevenlabs.String("accent_subtle"),
					Base:                elevenlabs.String("base"),
					BaseActive:          elevenlabs.String("base_active"),
					BaseBorder:          elevenlabs.String("base_border"),
					BaseError:           elevenlabs.String("base_error"),
					BaseHover:           elevenlabs.String("base_hover"),
					BasePrimary:         elevenlabs.String("base_primary"),
					BaseSubtle:          elevenlabs.String("base_subtle"),
					BubbleRadius:        elevenlabs.Float(0),
					ButtonRadius:        elevenlabs.Float(0),
					CompactSheetRadius:  elevenlabs.Float(0),
					DropdownSheetRadius: elevenlabs.Float(0),
					InputRadius:         elevenlabs.Float(0),
					OverlayPadding:      elevenlabs.Float(0),
					SheetRadius:         elevenlabs.Float(0),
				},
				SupportsTextOnly: elevenlabs.Bool(true),
				TermsHTML:        elevenlabs.String("terms_html"),
				TermsKey:         elevenlabs.String("terms_key"),
				TermsText:        elevenlabs.String("terms_text"),
				TextColor:        elevenlabs.String("text_color"),
				TextContents: elevenlabs.WidgetTextContentsParam{
					AcceptTerms:                     elevenlabs.String("accept_terms"),
					AgentEndedConversation:          elevenlabs.String("agent_ended_conversation"),
					ChangeLanguage:                  elevenlabs.String("change_language"),
					ChattingStatus:                  elevenlabs.String("chatting_status"),
					Collapse:                        elevenlabs.String("collapse"),
					ConnectingStatus:                elevenlabs.String("connecting_status"),
					ConversationID:                  elevenlabs.String("conversation_id"),
					Copied:                          elevenlabs.String("copied"),
					CopyID:                          elevenlabs.String("copy_id"),
					DismissTerms:                    elevenlabs.String("dismiss_terms"),
					EndCall:                         elevenlabs.String("end_call"),
					ErrorOccurred:                   elevenlabs.String("error_occurred"),
					Expand:                          elevenlabs.String("expand"),
					FollowUpFeedbackPlaceholder:     elevenlabs.String("follow_up_feedback_placeholder"),
					GoBack:                          elevenlabs.String("go_back"),
					InitiateFeedback:                elevenlabs.String("initiate_feedback"),
					InputLabel:                      elevenlabs.String("input_label"),
					InputPlaceholder:                elevenlabs.String("input_placeholder"),
					InputPlaceholderNewConversation: elevenlabs.String("input_placeholder_new_conversation"),
					InputPlaceholderTextOnly:        elevenlabs.String("input_placeholder_text_only"),
					ListeningStatus:                 elevenlabs.String("listening_status"),
					MainLabel:                       elevenlabs.String("main_label"),
					MuteMicrophone:                  elevenlabs.String("mute_microphone"),
					NewCall:                         elevenlabs.String("new_call"),
					RequestFollowUpFeedback:         elevenlabs.String("request_follow_up_feedback"),
					SpeakingStatus:                  elevenlabs.String("speaking_status"),
					StartCall:                       elevenlabs.String("start_call"),
					StartChat:                       elevenlabs.String("start_chat"),
					Submit:                          elevenlabs.String("submit"),
					ThanksForFeedback:               elevenlabs.String("thanks_for_feedback"),
					ThanksForFeedbackDetails:        elevenlabs.String("thanks_for_feedback_details"),
					UserEndedConversation:           elevenlabs.String("user_ended_conversation"),
				},
				TextInputEnabled:  elevenlabs.Bool(true),
				TranscriptEnabled: elevenlabs.Bool(true),
				Variant:           elevenlabs.EmbedVariantTiny,
			},
			WorkspaceOverrides: elevenlabs.AgentPlatformSettingsRequestModelWorkspaceOverridesParam{
				ConversationInitiationClientDataWebhook: elevenlabs.ConversationInitiationClientDataWebhookParam{
					RequestHeaders: map[string]elevenlabs.ConversationInitiationClientDataWebhookRequestHeaderUnionParam{
						"Content-Type": {
							OfString: elevenlabs.String("application/json"),
						},
					},
					URL: "https://example.com/webhook",
				},
				Webhooks: elevenlabs.ConvAIWebhooksParam{
					Events:            []string{"transcript"},
					PostCallWebhookID: elevenlabs.String("post_call_webhook_id"),
					SendAudio:         elevenlabs.Bool(true),
				},
			},
		},
		Tags: []string{"Customer Support", "Technical Help", "Eleven"},
		Workflow: elevenlabs.AgentWorkflowRequestModelParam{
			Edges: map[string]elevenlabs.AgentWorkflowRequestModelEdgeParam{
				"entry_to_tool_a": {
					Source: "entry_node",
					Target: "tool_node_a",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"start_to_entry": {
					Source: "start_node",
					Target: "entry_node",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_a_to_failure": {
					Source: "tool_node_a",
					Target: "failure_node",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_a_to_tool_b": {
					Source: "tool_node_a",
					Target: "tool_node_b",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_b_to_agent_transfer": {
					Source: "tool_node_b",
					Target: "success_transfer",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_b_to_conversation": {
					Source: "tool_node_b",
					Target: "success_conversation",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_b_to_end": {
					Source: "tool_node_b",
					Target: "success_end",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
				"tool_b_to_phone": {
					Source: "tool_node_b",
					Target: "success_phone",
					BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
					ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
						OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
							Label: elevenlabs.String("label"),
							Type:  "unconditional",
						},
					},
				},
			},
			Nodes: map[string]elevenlabs.AgentWorkflowRequestModelNodeUnionParam{
				"entry_node": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"entry_to_tool_a"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"failure_node": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"string"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"start_node": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"start_to_entry"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"success_conversation": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"string"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"success_end": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"string"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"success_phone": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"string"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"success_transfer": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"string"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"tool_node_a": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"tool_a_to_failure", "tool_a_to_tool_b"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
				"tool_node_b": {
					OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
						EdgeOrder: []string{"tool_b_to_conversation", "tool_b_to_end", "tool_b_to_phone", "tool_b_to_agent_transfer"},
						Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
							X: elevenlabs.Float(0),
							Y: elevenlabs.Float(0),
						},
						Type: "start",
					},
				},
			},
			PreventSubagentLoops: elevenlabs.Bool(true),
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

func TestConvaiAgentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.Get(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentGetParams{
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

func TestConvaiAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.Update(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentUpdateParams{
			ConversationConfig: map[string]any{},
			Name:               elevenlabs.String("My agent"),
			PlatformSettings:   map[string]any{},
			Tags:               []string{"Customer Support", "Technical Help", "Eleven"},
			VersionDescription: elevenlabs.String("version_description"),
			Workflow: elevenlabs.AgentWorkflowRequestModelParam{
				Edges: map[string]elevenlabs.AgentWorkflowRequestModelEdgeParam{
					"entry_to_tool_a": {
						Source: "entry_node",
						Target: "tool_node_a",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"start_to_entry": {
						Source: "start_node",
						Target: "entry_node",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_a_to_failure": {
						Source: "tool_node_a",
						Target: "failure_node",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_a_to_tool_b": {
						Source: "tool_node_a",
						Target: "tool_node_b",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_b_to_agent_transfer": {
						Source: "tool_node_b",
						Target: "success_transfer",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_b_to_conversation": {
						Source: "tool_node_b",
						Target: "success_conversation",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_b_to_end": {
						Source: "tool_node_b",
						Target: "success_end",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
					"tool_b_to_phone": {
						Source: "tool_node_b",
						Target: "success_phone",
						BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
						ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
							OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
								Label: elevenlabs.String("label"),
								Type:  "unconditional",
							},
						},
					},
				},
				Nodes: map[string]elevenlabs.AgentWorkflowRequestModelNodeUnionParam{
					"entry_node": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"entry_to_tool_a"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"failure_node": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"string"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"start_node": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"start_to_entry"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"success_conversation": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"string"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"success_end": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"string"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"success_phone": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"string"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"success_transfer": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"string"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"tool_node_a": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"tool_a_to_failure", "tool_a_to_tool_b"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
					"tool_node_b": {
						OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
							EdgeOrder: []string{"tool_b_to_conversation", "tool_b_to_end", "tool_b_to_phone", "tool_b_to_agent_transfer"},
							Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
								X: elevenlabs.Float(0),
								Y: elevenlabs.Float(0),
							},
							Type: "start",
						},
					},
				},
				PreventSubagentLoops: elevenlabs.Bool(true),
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

func TestConvaiAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.List(context.TODO(), elevenlabs.ConvaiAgentListParams{
		Archived:            elevenlabs.Bool(false),
		Cursor:              elevenlabs.String("cursor"),
		PageSize:            elevenlabs.Int(1),
		Search:              elevenlabs.String("search"),
		ShowOnlyOwnedAgents: elevenlabs.Bool(true),
		SortBy:              elevenlabs.ConvaiAgentListParamsSortByName,
		SortDirection:       elevenlabs.SortDirectionAsc,
		XiAPIKey:            elevenlabs.String("xi-api-key"),
	})
	if err != nil {
		var apierr *elevenlabs.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestConvaiAgentDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Convai.Agents.Delete(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentDeleteParams{
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

func TestConvaiAgentAvatarWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.Avatar(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentAvatarParams{
			AvatarFile: io.Reader(bytes.NewBuffer([]byte("some file contents"))),
			XiAPIKey:   elevenlabs.String("xi-api-key"),
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

func TestConvaiAgentDuplicateWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.Duplicate(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentDuplicateParams{
			Name:     elevenlabs.String("My agent"),
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

func TestConvaiAgentGetLinkWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.GetLink(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentGetLinkParams{
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

func TestConvaiAgentGetSummariesWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.GetSummaries(context.TODO(), elevenlabs.ConvaiAgentGetSummariesParams{
		AgentIDs: []string{"string"},
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

func TestConvaiAgentGetWidgetWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.GetWidget(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentGetWidgetParams{
			ConversationSignature: elevenlabs.String("conversation_signature"),
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

func TestConvaiAgentRunTestsWithOptionalParams(t *testing.T) {
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
	_, err := client.Convai.Agents.RunTests(
		context.TODO(),
		"agent_3701k3ttaq12ewp8b7qv5rfyszkz",
		elevenlabs.ConvaiAgentRunTestsParams{
			Tests: []elevenlabs.ConvaiAgentRunTestsParamsTest{{
				TestID:         "test_id",
				WorkflowNodeID: elevenlabs.String("workflow_node_id"),
			}},
			AgentConfigOverride: elevenlabs.AdhocAgentConfigOverrideForTestRequestModelParam{
				ConversationConfig: elevenlabs.ConversationalConfigAPIModelInputParam{
					Agent: elevenlabs.AgentConfigAPIModelInputParam{
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
					Asr: elevenlabs.AsrConversationalConfigParam{
						Keywords:             []string{"hello", "world"},
						Provider:             elevenlabs.AsrConversationalConfigProviderElevenlabs,
						Quality:              elevenlabs.AsrConversationalConfigQualityHigh,
						UserInputAudioFormat: elevenlabs.AsrConversationalConfigUserInputAudioFormatPcm16000,
					},
					Conversation: elevenlabs.ConversationConfigParam{
						ClientEvents:       []elevenlabs.ClientEvent{elevenlabs.ClientEventAudio, elevenlabs.ClientEventInterruption},
						MaxDurationSeconds: elevenlabs.Int(600),
						MonitoringEnabled:  elevenlabs.Bool(true),
						MonitoringEvents:   []elevenlabs.ClientEvent{elevenlabs.ClientEventConversationInitiationMetadata},
						TextOnly:           elevenlabs.Bool(true),
					},
					LanguagePresets: map[string]elevenlabs.ConversationalConfigAPIModelInputLanguagePresetParam{
						"foo": {
							Overrides: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetOverridesParam{
								Agent: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetOverridesAgentParam{
									FirstMessage: elevenlabs.String("Hello, how can I help you today?"),
									Language:     elevenlabs.String("en"),
									Prompt: elevenlabs.PromptAgentAPIModelOverrideParam{
										Llm:                elevenlabs.LlmGemini2_0Flash001,
										NativeMcpServerIDs: []string{"string"},
										Prompt:             elevenlabs.String("You are a helpful assistant that can answer questions about the topic of the conversation."),
									},
								},
								Conversation: elevenlabs.ConversationConfigOverrideParam{
									TextOnly: elevenlabs.Bool(true),
								},
								Tts: elevenlabs.TtsConversationalConfigOverrideParam{
									SimilarityBoost: elevenlabs.Float(0.8),
									Speed:           elevenlabs.Float(1),
									Stability:       elevenlabs.Float(0.5),
									VoiceID:         elevenlabs.String("cjVigY5qzO86Huf0OWal"),
								},
								Turn: elevenlabs.TurnConfigOverrideParam{
									SoftTimeoutConfig: elevenlabs.TurnConfigOverrideSoftTimeoutConfigParam{
										Message: elevenlabs.String("Hhmmmm...yeah."),
									},
								},
							},
							FirstMessageTranslation: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetFirstMessageTranslationParam{
								SourceHash: "source_hash",
								Text:       "text",
							},
							SoftTimeoutTranslation: elevenlabs.ConversationalConfigAPIModelInputLanguagePresetSoftTimeoutTranslationParam{
								SourceHash: "source_hash",
								Text:       "text",
							},
						},
					},
					Tts: elevenlabs.ConversationalConfigAPIModelInputTtsParam{
						AgentOutputAudioFormat:   elevenlabs.TtsOutputFormatPcm16000,
						ModelID:                  elevenlabs.TtsConversationalModelElevenTurboV2,
						OptimizeStreamingLatency: 3,
						PronunciationDictionaryLocators: []elevenlabs.PydanticPronunciationDictionaryVersionLocatorParam{{
							PronunciationDictionaryID: "pronunciation_dictionary_id",
							VersionID:                 elevenlabs.String("version_id"),
						}},
						SimilarityBoost: elevenlabs.Float(0.8),
						Speed:           elevenlabs.Float(1),
						Stability:       elevenlabs.Float(0.5),
						SupportedVoices: []elevenlabs.SupportedVoiceParam{{
							Label:                    "x",
							VoiceID:                  "x",
							Description:              elevenlabs.String("description"),
							Language:                 elevenlabs.String("language"),
							ModelFamily:              elevenlabs.SupportedVoiceModelFamilyTurbo,
							OptimizeStreamingLatency: 0,
							SimilarityBoost:          elevenlabs.Float(0),
							Speed:                    elevenlabs.Float(0.7),
							Stability:                elevenlabs.Float(0),
						}},
						TextNormalisationType: elevenlabs.TextNormalisationTypeSystemPrompt,
						VoiceID:               elevenlabs.String("cjVigY5qzO86Huf0OWal"),
					},
					Turn: elevenlabs.TurnConfigParam{
						InitialWaitTime:       elevenlabs.Float(0),
						Mode:                  elevenlabs.TurnConfigModeTurn,
						SilenceEndCallTimeout: elevenlabs.Float(-1),
						SoftTimeoutConfig: elevenlabs.TurnConfigSoftTimeoutConfigParam{
							Message:                elevenlabs.String("Hhmmmm...yeah."),
							TimeoutSeconds:         elevenlabs.Float(-1),
							UseLlmGeneratedMessage: elevenlabs.Bool(false),
						},
						SpellingPatience: elevenlabs.TurnConfigSpellingPatienceAuto,
						TurnEagerness:    elevenlabs.TurnConfigTurnEagernessNormal,
						TurnTimeout:      elevenlabs.Float(7),
					},
					Vad: elevenlabs.VadConfigParam{
						BackgroundVoiceDetection: elevenlabs.Bool(false),
					},
				},
				PlatformSettings: elevenlabs.AgentPlatformSettingsRequestModelParam{
					Archived: elevenlabs.Bool(true),
					Auth: elevenlabs.AuthSettingsParam{
						Allowlist: []elevenlabs.AllowlistItemParam{{
							Hostname: "https://example.com",
						}},
						EnableAuth:     elevenlabs.Bool(true),
						ShareableToken: elevenlabs.String("1234567890"),
					},
					CallLimits: elevenlabs.AgentCallLimitsParam{
						AgentConcurrencyLimit: elevenlabs.Int(-1),
						BurstingEnabled:       elevenlabs.Bool(true),
						DailyLimit:            elevenlabs.Int(100000),
					},
					DataCollection: map[string]elevenlabs.AgentPlatformSettingsRequestModelDataCollectionParam{
						"foo": {
							Type: "string",
							ConstantValue: elevenlabs.AgentPlatformSettingsRequestModelDataCollectionConstantValueUnionParam{
								OfString: elevenlabs.String("string"),
							},
							Description:      elevenlabs.String("A user-provided message"),
							DynamicVariable:  elevenlabs.String("dynamic_variable"),
							Enum:             []string{"string"},
							IsSystemProvided: elevenlabs.Bool(true),
						},
					},
					Evaluation: elevenlabs.EvaluationSettingsParam{
						Criteria: []elevenlabs.PromptEvaluationCriteriaParam{{
							ID:                     "1234567890",
							ConversationGoalPrompt: "You are a helpful assistant that can answer questions about the topic of the conversation.",
							Name:                   "Customer satisfaction check",
							Type:                   elevenlabs.PromptEvaluationCriteriaTypePrompt,
							UseKnowledgeBase:       elevenlabs.Bool(false),
						}},
					},
					Guardrails: elevenlabs.AgentPlatformSettingsRequestModelGuardrailsParam{
						Moderation: elevenlabs.AgentPlatformSettingsRequestModelGuardrailsModerationParam{
							Config: elevenlabs.ModerationConfigParam{
								Harassment: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								HarassmentThreatening: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								Hate: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								HateThreatening: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								SelfHarm: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								SelfHarmInstructions: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								SelfHarmIntent: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								Sexual: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								SexualMinors: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								Violence: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
								ViolenceGraphic: elevenlabs.ThresholdGuardrailParam{
									IsEnabled: elevenlabs.Bool(true),
									Threshold: elevenlabs.Float(0),
								},
							},
						},
						Version: "1",
					},
					Overrides: elevenlabs.AgentPlatformSettingsRequestModelOverridesParam{
						ConversationConfigOverride: elevenlabs.AgentPlatformSettingsRequestModelOverridesConversationConfigOverrideParam{
							Agent: elevenlabs.AgentConfigOverrideConfigParam{
								FirstMessage: elevenlabs.Bool(true),
								Language:     elevenlabs.Bool(true),
								Prompt: elevenlabs.AgentConfigOverrideConfigPromptParam{
									Llm:                elevenlabs.Bool(true),
									NativeMcpServerIDs: elevenlabs.Bool(true),
									Prompt:             elevenlabs.Bool(true),
								},
							},
							Conversation: elevenlabs.ConversationConfigOverrideConfigParam{
								TextOnly: elevenlabs.Bool(true),
							},
							Tts: elevenlabs.TtsConversationalConfigOverrideConfigParam{
								SimilarityBoost: elevenlabs.Bool(true),
								Speed:           elevenlabs.Bool(true),
								Stability:       elevenlabs.Bool(true),
								VoiceID:         elevenlabs.Bool(true),
							},
							Turn: elevenlabs.TurnConfigOverrideConfigParam{
								SoftTimeoutConfig: elevenlabs.TurnConfigOverrideConfigSoftTimeoutConfigParam{
									Message: elevenlabs.Bool(true),
								},
							},
						},
						CustomLlmExtraBody: elevenlabs.Bool(true),
						EnableConversationInitiationClientDataFromWebhook: elevenlabs.Bool(true),
					},
					Privacy: elevenlabs.PrivacyConfigParam{
						ApplyToExistingConversations: elevenlabs.Bool(false),
						DeleteAudio:                  elevenlabs.Bool(false),
						DeleteTranscriptAndPii:       elevenlabs.Bool(false),
						RecordVoice:                  elevenlabs.Bool(true),
						RetentionDays:                elevenlabs.Int(-1),
						ZeroRetentionMode:            elevenlabs.Bool(false),
					},
					Testing: elevenlabs.AgentTestingSettingsParam{
						AttachedTests: []elevenlabs.AgentTestingSettingsAttachedTestParam{{
							TestID:         "test_123",
							WorkflowNodeID: elevenlabs.String("node_abc"),
						}, {
							TestID:         "test_456",
							WorkflowNodeID: elevenlabs.String("workflow_node_id"),
						}},
					},
					Widget: elevenlabs.AgentPlatformSettingsRequestModelWidgetParam{
						ActionText:     elevenlabs.String("action_text"),
						AlwaysExpanded: elevenlabs.Bool(true),
						Avatar: elevenlabs.AgentPlatformSettingsRequestModelWidgetAvatarUnionParam{
							OfOrbAvatar: &elevenlabs.OrbAvatarParam{
								Color1: elevenlabs.String("#2792dc"),
								Color2: elevenlabs.String("#9ce6e6"),
								Type:   elevenlabs.OrbAvatarTypeOrb,
							},
						},
						BgColor:                       elevenlabs.String("bg_color"),
						BorderColor:                   elevenlabs.String("border_color"),
						BorderRadius:                  elevenlabs.Int(0),
						BtnColor:                      elevenlabs.String("btn_color"),
						BtnRadius:                     elevenlabs.Int(0),
						BtnTextColor:                  elevenlabs.String("btn_text_color"),
						ConversationModeToggleEnabled: elevenlabs.Bool(true),
						CustomAvatarPath:              elevenlabs.String("https://example.com/avatar.png"),
						DefaultExpanded:               elevenlabs.Bool(true),
						DisableBanner:                 elevenlabs.Bool(true),
						EndCallText:                   elevenlabs.String("end_call_text"),
						EndFeedback: elevenlabs.WidgetEndFeedbackConfigParam{
							Type: elevenlabs.WidgetEndFeedbackConfigTypeRating,
						},
						ExpandText:   elevenlabs.String("expand_text"),
						Expandable:   elevenlabs.WidgetExpandableNever,
						FeedbackMode: elevenlabs.WidgetFeedbackModeNone,
						FocusColor:   elevenlabs.String("focus_color"),
						LanguagePresets: map[string]elevenlabs.AgentPlatformSettingsRequestModelWidgetLanguagePresetParam{
							"foo": {
								TermsHTML: elevenlabs.String("terms_html"),
								TermsKey:  elevenlabs.String("terms_key"),
								TermsText: elevenlabs.String("terms_text"),
								TermsTranslation: elevenlabs.AgentPlatformSettingsRequestModelWidgetLanguagePresetTermsTranslationParam{
									SourceHash: "source_hash",
									Text:       "text",
								},
								TextContents: elevenlabs.WidgetTextContentsParam{
									AcceptTerms:                     elevenlabs.String("accept_terms"),
									AgentEndedConversation:          elevenlabs.String("agent_ended_conversation"),
									ChangeLanguage:                  elevenlabs.String("change_language"),
									ChattingStatus:                  elevenlabs.String("chatting_status"),
									Collapse:                        elevenlabs.String("collapse"),
									ConnectingStatus:                elevenlabs.String("connecting_status"),
									ConversationID:                  elevenlabs.String("conversation_id"),
									Copied:                          elevenlabs.String("copied"),
									CopyID:                          elevenlabs.String("copy_id"),
									DismissTerms:                    elevenlabs.String("dismiss_terms"),
									EndCall:                         elevenlabs.String("end_call"),
									ErrorOccurred:                   elevenlabs.String("error_occurred"),
									Expand:                          elevenlabs.String("expand"),
									FollowUpFeedbackPlaceholder:     elevenlabs.String("follow_up_feedback_placeholder"),
									GoBack:                          elevenlabs.String("go_back"),
									InitiateFeedback:                elevenlabs.String("initiate_feedback"),
									InputLabel:                      elevenlabs.String("input_label"),
									InputPlaceholder:                elevenlabs.String("input_placeholder"),
									InputPlaceholderNewConversation: elevenlabs.String("input_placeholder_new_conversation"),
									InputPlaceholderTextOnly:        elevenlabs.String("input_placeholder_text_only"),
									ListeningStatus:                 elevenlabs.String("listening_status"),
									MainLabel:                       elevenlabs.String("main_label"),
									MuteMicrophone:                  elevenlabs.String("mute_microphone"),
									NewCall:                         elevenlabs.String("new_call"),
									RequestFollowUpFeedback:         elevenlabs.String("request_follow_up_feedback"),
									SpeakingStatus:                  elevenlabs.String("speaking_status"),
									StartCall:                       elevenlabs.String("start_call"),
									StartChat:                       elevenlabs.String("start_chat"),
									Submit:                          elevenlabs.String("submit"),
									ThanksForFeedback:               elevenlabs.String("thanks_for_feedback"),
									ThanksForFeedbackDetails:        elevenlabs.String("thanks_for_feedback_details"),
									UserEndedConversation:           elevenlabs.String("user_ended_conversation"),
								},
							},
						},
						LanguageSelector:      elevenlabs.Bool(false),
						ListeningText:         elevenlabs.String("listening_text"),
						MarkdownLinkAllowHTTP: elevenlabs.Bool(true),
						MarkdownLinkAllowedHosts: []elevenlabs.AllowlistItemParam{{
							Hostname: "hostname",
						}},
						MarkdownLinkIncludeWww:  elevenlabs.Bool(true),
						MicMutingEnabled:        elevenlabs.Bool(true),
						OverrideLink:            elevenlabs.String("override_link"),
						Placement:               elevenlabs.WidgetPlacementTopLeft,
						ShareablePageShowTerms:  elevenlabs.Bool(true),
						ShareablePageText:       elevenlabs.String("shareable_page_text"),
						ShowAvatarWhenCollapsed: elevenlabs.Bool(true),
						SpeakingText:            elevenlabs.String("speaking_text"),
						StartCallText:           elevenlabs.String("start_call_text"),
						Styles: elevenlabs.WidgetStylesParam{
							Accent:              elevenlabs.String("accent"),
							AccentActive:        elevenlabs.String("accent_active"),
							AccentBorder:        elevenlabs.String("accent_border"),
							AccentHover:         elevenlabs.String("accent_hover"),
							AccentPrimary:       elevenlabs.String("accent_primary"),
							AccentSubtle:        elevenlabs.String("accent_subtle"),
							Base:                elevenlabs.String("base"),
							BaseActive:          elevenlabs.String("base_active"),
							BaseBorder:          elevenlabs.String("base_border"),
							BaseError:           elevenlabs.String("base_error"),
							BaseHover:           elevenlabs.String("base_hover"),
							BasePrimary:         elevenlabs.String("base_primary"),
							BaseSubtle:          elevenlabs.String("base_subtle"),
							BubbleRadius:        elevenlabs.Float(0),
							ButtonRadius:        elevenlabs.Float(0),
							CompactSheetRadius:  elevenlabs.Float(0),
							DropdownSheetRadius: elevenlabs.Float(0),
							InputRadius:         elevenlabs.Float(0),
							OverlayPadding:      elevenlabs.Float(0),
							SheetRadius:         elevenlabs.Float(0),
						},
						SupportsTextOnly: elevenlabs.Bool(true),
						TermsHTML:        elevenlabs.String("terms_html"),
						TermsKey:         elevenlabs.String("terms_key"),
						TermsText:        elevenlabs.String("terms_text"),
						TextColor:        elevenlabs.String("text_color"),
						TextContents: elevenlabs.WidgetTextContentsParam{
							AcceptTerms:                     elevenlabs.String("accept_terms"),
							AgentEndedConversation:          elevenlabs.String("agent_ended_conversation"),
							ChangeLanguage:                  elevenlabs.String("change_language"),
							ChattingStatus:                  elevenlabs.String("chatting_status"),
							Collapse:                        elevenlabs.String("collapse"),
							ConnectingStatus:                elevenlabs.String("connecting_status"),
							ConversationID:                  elevenlabs.String("conversation_id"),
							Copied:                          elevenlabs.String("copied"),
							CopyID:                          elevenlabs.String("copy_id"),
							DismissTerms:                    elevenlabs.String("dismiss_terms"),
							EndCall:                         elevenlabs.String("end_call"),
							ErrorOccurred:                   elevenlabs.String("error_occurred"),
							Expand:                          elevenlabs.String("expand"),
							FollowUpFeedbackPlaceholder:     elevenlabs.String("follow_up_feedback_placeholder"),
							GoBack:                          elevenlabs.String("go_back"),
							InitiateFeedback:                elevenlabs.String("initiate_feedback"),
							InputLabel:                      elevenlabs.String("input_label"),
							InputPlaceholder:                elevenlabs.String("input_placeholder"),
							InputPlaceholderNewConversation: elevenlabs.String("input_placeholder_new_conversation"),
							InputPlaceholderTextOnly:        elevenlabs.String("input_placeholder_text_only"),
							ListeningStatus:                 elevenlabs.String("listening_status"),
							MainLabel:                       elevenlabs.String("main_label"),
							MuteMicrophone:                  elevenlabs.String("mute_microphone"),
							NewCall:                         elevenlabs.String("new_call"),
							RequestFollowUpFeedback:         elevenlabs.String("request_follow_up_feedback"),
							SpeakingStatus:                  elevenlabs.String("speaking_status"),
							StartCall:                       elevenlabs.String("start_call"),
							StartChat:                       elevenlabs.String("start_chat"),
							Submit:                          elevenlabs.String("submit"),
							ThanksForFeedback:               elevenlabs.String("thanks_for_feedback"),
							ThanksForFeedbackDetails:        elevenlabs.String("thanks_for_feedback_details"),
							UserEndedConversation:           elevenlabs.String("user_ended_conversation"),
						},
						TextInputEnabled:  elevenlabs.Bool(true),
						TranscriptEnabled: elevenlabs.Bool(true),
						Variant:           elevenlabs.EmbedVariantTiny,
					},
					WorkspaceOverrides: elevenlabs.AgentPlatformSettingsRequestModelWorkspaceOverridesParam{
						ConversationInitiationClientDataWebhook: elevenlabs.ConversationInitiationClientDataWebhookParam{
							RequestHeaders: map[string]elevenlabs.ConversationInitiationClientDataWebhookRequestHeaderUnionParam{
								"Content-Type": {
									OfString: elevenlabs.String("application/json"),
								},
							},
							URL: "https://example.com/webhook",
						},
						Webhooks: elevenlabs.ConvAIWebhooksParam{
							Events:            []string{"transcript"},
							PostCallWebhookID: elevenlabs.String("post_call_webhook_id"),
							SendAudio:         elevenlabs.Bool(true),
						},
					},
				},
				Workflow: elevenlabs.AgentWorkflowRequestModelParam{
					Edges: map[string]elevenlabs.AgentWorkflowRequestModelEdgeParam{
						"entry_to_tool_a": {
							Source: "entry_node",
							Target: "tool_node_a",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"start_to_entry": {
							Source: "start_node",
							Target: "entry_node",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_a_to_failure": {
							Source: "tool_node_a",
							Target: "failure_node",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_a_to_tool_b": {
							Source: "tool_node_a",
							Target: "tool_node_b",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_b_to_agent_transfer": {
							Source: "tool_node_b",
							Target: "success_transfer",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_b_to_conversation": {
							Source: "tool_node_b",
							Target: "success_conversation",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_b_to_end": {
							Source: "tool_node_b",
							Target: "success_end",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
						"tool_b_to_phone": {
							Source: "tool_node_b",
							Target: "success_phone",
							BackwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeBackwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
							ForwardCondition: elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnionParam{
								OfUnconditional: &elevenlabs.AgentWorkflowRequestModelEdgeForwardConditionUnconditionalParam{
									Label: elevenlabs.String("label"),
									Type:  "unconditional",
								},
							},
						},
					},
					Nodes: map[string]elevenlabs.AgentWorkflowRequestModelNodeUnionParam{
						"entry_node": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"entry_to_tool_a"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"failure_node": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"string"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"start_node": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"start_to_entry"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"success_conversation": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"string"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"success_end": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"string"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"success_phone": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"string"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"success_transfer": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"string"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"tool_node_a": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"tool_a_to_failure", "tool_a_to_tool_b"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
						"tool_node_b": {
							OfStart: &elevenlabs.AgentWorkflowRequestModelNodeStartParam{
								EdgeOrder: []string{"tool_b_to_conversation", "tool_b_to_end", "tool_b_to_phone", "tool_b_to_agent_transfer"},
								Position: elevenlabs.AgentWorkflowRequestModelNodeStartPositionParam{
									X: elevenlabs.Float(0),
									Y: elevenlabs.Float(0),
								},
								Type: "start",
							},
						},
					},
					PreventSubagentLoops: elevenlabs.Bool(true),
				},
			},
			BranchID: elevenlabs.String("branch_id"),
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
