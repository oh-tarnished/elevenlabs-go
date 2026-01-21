// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiAgentLlmUsageService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiAgentLlmUsageService] method instead.
type ConvaiAgentLlmUsageService struct {
	Options []option.RequestOption
}

// NewConvaiAgentLlmUsageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiAgentLlmUsageService(opts ...option.RequestOption) (r ConvaiAgentLlmUsageService) {
	r = ConvaiAgentLlmUsageService{}
	r.Options = opts
	return
}

// Calculates expected number of LLM tokens needed for the specified agent.
func (r *ConvaiAgentLlmUsageService) Calculate(ctx context.Context, agentID string, params ConvaiAgentLlmUsageCalculateParams, opts ...option.RequestOption) (res *LlmUsageCalculatorResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agent/%s/llm-usage/calculate", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type Llm string

const (
	LlmGpt4oMini                        Llm = "gpt-4o-mini"
	LlmGpt4o                            Llm = "gpt-4o"
	LlmGpt4                             Llm = "gpt-4"
	LlmGpt4Turbo                        Llm = "gpt-4-turbo"
	LlmGpt4_1                           Llm = "gpt-4.1"
	LlmGpt4_1Mini                       Llm = "gpt-4.1-mini"
	LlmGpt4_1Nano                       Llm = "gpt-4.1-nano"
	LlmGpt5                             Llm = "gpt-5"
	LlmGpt5_1                           Llm = "gpt-5.1"
	LlmGpt5_2                           Llm = "gpt-5.2"
	LlmGpt5_2ChatLatest                 Llm = "gpt-5.2-chat-latest"
	LlmGpt5Mini                         Llm = "gpt-5-mini"
	LlmGpt5Nano                         Llm = "gpt-5-nano"
	LlmGpt3_5Turbo                      Llm = "gpt-3.5-turbo"
	LlmGemini1_5Pro                     Llm = "gemini-1.5-pro"
	LlmGemini1_5Flash                   Llm = "gemini-1.5-flash"
	LlmGemini2_0Flash                   Llm = "gemini-2.0-flash"
	LlmGemini2_0FlashLite               Llm = "gemini-2.0-flash-lite"
	LlmGemini2_5FlashLite               Llm = "gemini-2.5-flash-lite"
	LlmGemini2_5Flash                   Llm = "gemini-2.5-flash"
	LlmGemini3ProPreview                Llm = "gemini-3-pro-preview"
	LlmGemini3FlashPreview              Llm = "gemini-3-flash-preview"
	LlmClaudeSonnet4_5                  Llm = "claude-sonnet-4-5"
	LlmClaudeSonnet4                    Llm = "claude-sonnet-4"
	LlmClaudeHaiku4_5                   Llm = "claude-haiku-4-5"
	LlmClaude3_7Sonnet                  Llm = "claude-3-7-sonnet"
	LlmClaude3_5Sonnet                  Llm = "claude-3-5-sonnet"
	LlmClaude3_5SonnetV1                Llm = "claude-3-5-sonnet-v1"
	LlmClaude3Haiku                     Llm = "claude-3-haiku"
	LlmGrokBeta                         Llm = "grok-beta"
	LlmCustomLlm                        Llm = "custom-llm"
	LlmQwen3_4b                         Llm = "qwen3-4b"
	LlmQwen3_30bA3b                     Llm = "qwen3-30b-a3b"
	LlmGptOss20b                        Llm = "gpt-oss-20b"
	LlmGptOss120b                       Llm = "gpt-oss-120b"
	LlmGlm45AirFp8                      Llm = "glm-45-air-fp8"
	LlmGemini2_5FlashPreview09_2025     Llm = "gemini-2.5-flash-preview-09-2025"
	LlmGemini2_5FlashLitePreview09_2025 Llm = "gemini-2.5-flash-lite-preview-09-2025"
	LlmGemini2_5FlashPreview05_20       Llm = "gemini-2.5-flash-preview-05-20"
	LlmGemini2_5FlashPreview04_17       Llm = "gemini-2.5-flash-preview-04-17"
	LlmGemini2_5FlashLitePreview06_17   Llm = "gemini-2.5-flash-lite-preview-06-17"
	LlmGemini2_0FlashLite001            Llm = "gemini-2.0-flash-lite-001"
	LlmGemini2_0Flash001                Llm = "gemini-2.0-flash-001"
	LlmGemini1_5Flash002                Llm = "gemini-1.5-flash-002"
	LlmGemini1_5Flash001                Llm = "gemini-1.5-flash-001"
	LlmGemini1_5Pro002                  Llm = "gemini-1.5-pro-002"
	LlmGemini1_5Pro001                  Llm = "gemini-1.5-pro-001"
	LlmClaudeSonnet4_20250514           Llm = "claude-sonnet-4@20250514"
	LlmClaudeSonnet4_5_20250929         Llm = "claude-sonnet-4-5@20250929"
	LlmClaudeHaiku4_5_20251001          Llm = "claude-haiku-4-5@20251001"
	LlmClaude3_7Sonnet20250219          Llm = "claude-3-7-sonnet@20250219"
	LlmClaude3_5Sonnet20240620          Llm = "claude-3-5-sonnet@20240620"
	LlmClaude3_5SonnetV2_20241022       Llm = "claude-3-5-sonnet-v2@20241022"
	LlmClaude3Haiku20240307             Llm = "claude-3-haiku@20240307"
	LlmGpt5_2025_08_07                  Llm = "gpt-5-2025-08-07"
	LlmGpt5_1_2025_11_13                Llm = "gpt-5.1-2025-11-13"
	LlmGpt5_2_2025_12_11                Llm = "gpt-5.2-2025-12-11"
	LlmGpt5Mini2025_08_07               Llm = "gpt-5-mini-2025-08-07"
	LlmGpt5Nano2025_08_07               Llm = "gpt-5-nano-2025-08-07"
	LlmGpt4_1_2025_04_14                Llm = "gpt-4.1-2025-04-14"
	LlmGpt4_1Mini2025_04_14             Llm = "gpt-4.1-mini-2025-04-14"
	LlmGpt4_1Nano2025_04_14             Llm = "gpt-4.1-nano-2025-04-14"
	LlmGpt4oMini2024_07_18              Llm = "gpt-4o-mini-2024-07-18"
	LlmGpt4o2024_11_20                  Llm = "gpt-4o-2024-11-20"
	LlmGpt4o2024_08_06                  Llm = "gpt-4o-2024-08-06"
	LlmGpt4o2024_05_13                  Llm = "gpt-4o-2024-05-13"
	LlmGpt4_0613                        Llm = "gpt-4-0613"
	LlmGpt4_0314                        Llm = "gpt-4-0314"
	LlmGpt4Turbo2024_04_09              Llm = "gpt-4-turbo-2024-04-09"
	LlmGpt3_5Turbo0125                  Llm = "gpt-3.5-turbo-0125"
	LlmGpt3_5Turbo1106                  Llm = "gpt-3.5-turbo-1106"
	LlmWattTool8b                       Llm = "watt-tool-8b"
	LlmWattTool70b                      Llm = "watt-tool-70b"
)

type LlmUsageCalculatorResponseModel struct {
	LlmPrices []LlmUsageCalculatorResponseModelLlmPrice `json:"llm_prices,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LlmPrices   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageCalculatorResponseModel) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageCalculatorResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LlmUsageCalculatorResponseModelLlmPrice struct {
	// Any of "gpt-4o-mini", "gpt-4o", "gpt-4", "gpt-4-turbo", "gpt-4.1",
	// "gpt-4.1-mini", "gpt-4.1-nano", "gpt-5", "gpt-5.1", "gpt-5.2",
	// "gpt-5.2-chat-latest", "gpt-5-mini", "gpt-5-nano", "gpt-3.5-turbo",
	// "gemini-1.5-pro", "gemini-1.5-flash", "gemini-2.0-flash",
	// "gemini-2.0-flash-lite", "gemini-2.5-flash-lite", "gemini-2.5-flash",
	// "gemini-3-pro-preview", "gemini-3-flash-preview", "claude-sonnet-4-5",
	// "claude-sonnet-4", "claude-haiku-4-5", "claude-3-7-sonnet", "claude-3-5-sonnet",
	// "claude-3-5-sonnet-v1", "claude-3-haiku", "grok-beta", "custom-llm", "qwen3-4b",
	// "qwen3-30b-a3b", "gpt-oss-20b", "gpt-oss-120b", "glm-45-air-fp8",
	// "gemini-2.5-flash-preview-09-2025", "gemini-2.5-flash-lite-preview-09-2025",
	// "gemini-2.5-flash-preview-05-20", "gemini-2.5-flash-preview-04-17",
	// "gemini-2.5-flash-lite-preview-06-17", "gemini-2.0-flash-lite-001",
	// "gemini-2.0-flash-001", "gemini-1.5-flash-002", "gemini-1.5-flash-001",
	// "gemini-1.5-pro-002", "gemini-1.5-pro-001", "claude-sonnet-4@20250514",
	// "claude-sonnet-4-5@20250929", "claude-haiku-4-5@20251001",
	// "claude-3-7-sonnet@20250219", "claude-3-5-sonnet@20240620",
	// "claude-3-5-sonnet-v2@20241022", "claude-3-haiku@20240307", "gpt-5-2025-08-07",
	// "gpt-5.1-2025-11-13", "gpt-5.2-2025-12-11", "gpt-5-mini-2025-08-07",
	// "gpt-5-nano-2025-08-07", "gpt-4.1-2025-04-14", "gpt-4.1-mini-2025-04-14",
	// "gpt-4.1-nano-2025-04-14", "gpt-4o-mini-2024-07-18", "gpt-4o-2024-11-20",
	// "gpt-4o-2024-08-06", "gpt-4o-2024-05-13", "gpt-4-0613", "gpt-4-0314",
	// "gpt-4-turbo-2024-04-09", "gpt-3.5-turbo-0125", "gpt-3.5-turbo-1106",
	// "watt-tool-8b", "watt-tool-70b".
	Llm            Llm     `json:"llm,required"`
	PricePerMinute float64 `json:"price_per_minute,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Llm            respjson.Field
		PricePerMinute respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LlmUsageCalculatorResponseModelLlmPrice) RawJSON() string { return r.JSON.raw }
func (r *LlmUsageCalculatorResponseModelLlmPrice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentLlmUsageCalculateParams struct {
	// Pages of content in pdf documents OR urls in agent's Knowledge Base.
	NumberOfPages param.Opt[int64] `json:"number_of_pages,omitzero"`
	// Length of the prompt in characters.
	PromptLength param.Opt[int64] `json:"prompt_length,omitzero"`
	// Whether RAG is enabled.
	RagEnabled param.Opt[bool] `json:"rag_enabled,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiAgentLlmUsageCalculateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiAgentLlmUsageCalculateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiAgentLlmUsageCalculateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
