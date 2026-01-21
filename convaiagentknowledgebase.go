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

// ConvaiAgentKnowledgeBaseService contains methods and other services that help
// with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiAgentKnowledgeBaseService] method instead.
type ConvaiAgentKnowledgeBaseService struct {
	Options []option.RequestOption
}

// NewConvaiAgentKnowledgeBaseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiAgentKnowledgeBaseService(opts ...option.RequestOption) (r ConvaiAgentKnowledgeBaseService) {
	r = ConvaiAgentKnowledgeBaseService{}
	r.Options = opts
	return
}

// Returns the number of pages in the agent's knowledge base.
func (r *ConvaiAgentKnowledgeBaseService) GetSize(ctx context.Context, agentID string, query ConvaiAgentKnowledgeBaseGetSizeParams, opts ...option.RequestOption) (res *ConvaiAgentKnowledgeBaseGetSizeResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/agent/%s/knowledge-base/size", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ConvaiAgentKnowledgeBaseGetSizeResponse struct {
	NumberOfPages float64 `json:"number_of_pages,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NumberOfPages respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiAgentKnowledgeBaseGetSizeResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiAgentKnowledgeBaseGetSizeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiAgentKnowledgeBaseGetSizeParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
