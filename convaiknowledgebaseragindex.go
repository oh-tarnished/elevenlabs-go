// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiKnowledgeBaseRagIndexService contains methods and other services that help
// with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiKnowledgeBaseRagIndexService] method instead.
type ConvaiKnowledgeBaseRagIndexService struct {
	Options []option.RequestOption
}

// NewConvaiKnowledgeBaseRagIndexService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConvaiKnowledgeBaseRagIndexService(opts ...option.RequestOption) (r ConvaiKnowledgeBaseRagIndexService) {
	r = ConvaiKnowledgeBaseRagIndexService{}
	r.Options = opts
	return
}

// Delete RAG index for the knowledgebase document.
func (r *ConvaiKnowledgeBaseRagIndexService) Delete(ctx context.Context, ragIndexID string, params ConvaiKnowledgeBaseRagIndexDeleteParams, opts ...option.RequestOption) (res *DocumentIndexResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DocumentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	if ragIndexID == "" {
		err = errors.New("missing required rag_index_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/rag-index/%s", params.DocumentationID, ragIndexID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// In case the document is not RAG indexed, it triggers rag indexing task,
// otherwise it just returns the current status.
func (r *ConvaiKnowledgeBaseRagIndexService) Compute(ctx context.Context, documentationID string, params ConvaiKnowledgeBaseRagIndexComputeParams, opts ...option.RequestOption) (res *DocumentIndexResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/rag-index", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieves and/or creates RAG indexes for multiple knowledge base documents in a
// single request.
func (r *ConvaiKnowledgeBaseRagIndexService) ComputeBatch(ctx context.Context, params ConvaiKnowledgeBaseRagIndexComputeBatchParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseRagIndexComputeBatchResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/rag-index"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Provides information about all RAG indexes of the specified knowledgebase
// document.
func (r *ConvaiKnowledgeBaseRagIndexService) GetIndexes(ctx context.Context, documentationID string, query ConvaiKnowledgeBaseRagIndexGetIndexesParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseRagIndexGetIndexesResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/rag-index", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Provides total size and other information of RAG indexes used by knowledgebase
// documents
func (r *ConvaiKnowledgeBaseRagIndexService) GetOverview(ctx context.Context, query ConvaiKnowledgeBaseRagIndexGetOverviewParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseRagIndexGetOverviewResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/rag-index"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DocumentIndexResponse struct {
	ID                      string                                       `json:"id,required"`
	DocumentModelIndexUsage DocumentIndexResponseDocumentModelIndexUsage `json:"document_model_index_usage,required"`
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	Model              EmbeddingModel `json:"model,required"`
	ProgressPercentage float64        `json:"progress_percentage,required"`
	// Any of "created", "processing", "failed", "succeeded", "rag_limit_exceeded",
	// "document_too_small", "cannot_index_folder".
	Status DocumentIndexResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		DocumentModelIndexUsage respjson.Field
		Model                   respjson.Field
		ProgressPercentage      respjson.Field
		Status                  respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentIndexResponse) RawJSON() string { return r.JSON.raw }
func (r *DocumentIndexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentIndexResponseDocumentModelIndexUsage struct {
	UsedBytes int64 `json:"used_bytes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		UsedBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DocumentIndexResponseDocumentModelIndexUsage) RawJSON() string { return r.JSON.raw }
func (r *DocumentIndexResponseDocumentModelIndexUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentIndexResponseStatus string

const (
	DocumentIndexResponseStatusCreated           DocumentIndexResponseStatus = "created"
	DocumentIndexResponseStatusProcessing        DocumentIndexResponseStatus = "processing"
	DocumentIndexResponseStatusFailed            DocumentIndexResponseStatus = "failed"
	DocumentIndexResponseStatusSucceeded         DocumentIndexResponseStatus = "succeeded"
	DocumentIndexResponseStatusRagLimitExceeded  DocumentIndexResponseStatus = "rag_limit_exceeded"
	DocumentIndexResponseStatusDocumentTooSmall  DocumentIndexResponseStatus = "document_too_small"
	DocumentIndexResponseStatusCannotIndexFolder DocumentIndexResponseStatus = "cannot_index_folder"
)

type EmbeddingModel string

const (
	EmbeddingModelE5Mistral7bInstruct         EmbeddingModel = "e5_mistral_7b_instruct"
	EmbeddingModelMultilingualE5LargeInstruct EmbeddingModel = "multilingual_e5_large_instruct"
)

type ConvaiKnowledgeBaseRagIndexComputeBatchResponse map[string]ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion

// ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion contains all possible
// properties and values from
// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess],
// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure].
//
// Use the [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion struct {
	// This field is from variant
	// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess].
	Data DocumentIndexResponse `json:"data"`
	// Any of "success", "failure".
	Status string `json:"status"`
	// This field is from variant
	// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure].
	ErrorCode int64 `json:"error_code"`
	// This field is from variant
	// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure].
	ErrorMessage string `json:"error_message"`
	// This field is from variant
	// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure].
	ErrorStatus string `json:"error_status"`
	JSON        struct {
		Data         respjson.Field
		Status       respjson.Field
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ErrorStatus  respjson.Field
		raw          string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseRagIndexComputeBatchResponseItem is implemented by each
// variant of [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion] to add
// type safety for the return type of
// [ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion.AsAny]
type anyConvaiKnowledgeBaseRagIndexComputeBatchResponseItem interface {
	implConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion()
}

func (ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess) implConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion() {
}
func (ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure) implConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion.AsAny().(type) {
//	case elevenlabs.ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess:
//	case elevenlabs.ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion) AsAny() anyConvaiKnowledgeBaseRagIndexComputeBatchResponseItem {
	switch u.Status {
	case "success":
		return u.AsSuccess()
	case "failure":
		return u.AsFailure()
	}
	return nil
}

func (u ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion) AsSuccess() (v ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion) AsFailure() (v ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess struct {
	Data   DocumentIndexResponse `json:"data,required"`
	Status constant.Success      `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure struct {
	ErrorCode    int64            `json:"error_code,required"`
	ErrorMessage string           `json:"error_message,required"`
	ErrorStatus  string           `json:"error_status,required"`
	Status       constant.Failure `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorCode    respjson.Field
		ErrorMessage respjson.Field
		ErrorStatus  respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseRagIndexComputeBatchResponseItemFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexGetIndexesResponse struct {
	Indexes []DocumentIndexResponse `json:"indexes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Indexes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseRagIndexGetIndexesResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseRagIndexGetIndexesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexGetOverviewResponse struct {
	Models         []ConvaiKnowledgeBaseRagIndexGetOverviewResponseModel `json:"models,required"`
	TotalMaxBytes  int64                                                 `json:"total_max_bytes,required"`
	TotalUsedBytes int64                                                 `json:"total_used_bytes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Models         respjson.Field
		TotalMaxBytes  respjson.Field
		TotalUsedBytes respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseRagIndexGetOverviewResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseRagIndexGetOverviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexGetOverviewResponseModel struct {
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	Model     EmbeddingModel `json:"model,required"`
	UsedBytes int64          `json:"used_bytes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		UsedBytes   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseRagIndexGetOverviewResponseModel) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseRagIndexGetOverviewResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexDeleteParams struct {
	// The id of a document from the knowledge base. This is returned on document
	// addition.
	DocumentationID string `path:"documentation_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiKnowledgeBaseRagIndexComputeParams struct {
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	Model EmbeddingModel `json:"model,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseRagIndexComputeParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseRagIndexComputeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseRagIndexComputeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexComputeBatchParams struct {
	// List of requested RAG indexes.
	Items []ConvaiKnowledgeBaseRagIndexComputeBatchParamsItem `json:"items,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseRagIndexComputeBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseRagIndexComputeBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseRagIndexComputeBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties CreateIfMissing, DocumentID, Model are required.
type ConvaiKnowledgeBaseRagIndexComputeBatchParamsItem struct {
	// Whether to create the RAG index if it does not exist
	CreateIfMissing bool `json:"create_if_missing,required"`
	// ID of the knowledgebase document for which to retrieve the index
	DocumentID string `json:"document_id,required"`
	// Embedding model to use for the RAG index
	//
	// Any of "e5_mistral_7b_instruct", "multilingual_e5_large_instruct".
	Model EmbeddingModel `json:"model,omitzero,required"`
	paramObj
}

func (r ConvaiKnowledgeBaseRagIndexComputeBatchParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseRagIndexComputeBatchParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseRagIndexComputeBatchParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseRagIndexGetIndexesParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiKnowledgeBaseRagIndexGetOverviewParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}
