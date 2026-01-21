// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
)

// ConvaiLlmUsageService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiLlmUsageService] method instead.
type ConvaiLlmUsageService struct {
	Options []option.RequestOption
}

// NewConvaiLlmUsageService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConvaiLlmUsageService(opts ...option.RequestOption) (r ConvaiLlmUsageService) {
	r = ConvaiLlmUsageService{}
	r.Options = opts
	return
}

// Returns a list of LLM models and the expected cost for using them based on the
// provided values.
func (r *ConvaiLlmUsageService) Calculate(ctx context.Context, body ConvaiLlmUsageCalculateParams, opts ...option.RequestOption) (res *LlmUsageCalculatorResponseModel, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/llm-usage/calculate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ConvaiLlmUsageCalculateParams struct {
	// Pages of content in PDF documents or URLs in the agent's knowledge base.
	NumberOfPages int64 `json:"number_of_pages,required"`
	// Length of the prompt in characters.
	PromptLength int64 `json:"prompt_length,required"`
	// Whether RAG is enabled.
	RagEnabled bool `json:"rag_enabled,required"`
	paramObj
}

func (r ConvaiLlmUsageCalculateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiLlmUsageCalculateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiLlmUsageCalculateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
