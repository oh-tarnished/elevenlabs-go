// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
	"github.com/oh-tarnished/elevenlabs-go/shared/constant"
)

// ConvaiKnowledgeBaseService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiKnowledgeBaseService] method instead.
type ConvaiKnowledgeBaseService struct {
	Options  []option.RequestOption
	RagIndex ConvaiKnowledgeBaseRagIndexService
}

// NewConvaiKnowledgeBaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiKnowledgeBaseService(opts ...option.RequestOption) (r ConvaiKnowledgeBaseService) {
	r = ConvaiKnowledgeBaseService{}
	r.Options = opts
	r.RagIndex = NewConvaiKnowledgeBaseRagIndexService(opts...)
	return
}

// Get a list of available knowledge base documents
func (r *ConvaiKnowledgeBaseService) List(ctx context.Context, params ConvaiKnowledgeBaseListParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Uploads a file or reference a webpage to use as part of the shared knowledge
// base
//
// Deprecated: deprecated
func (r *ConvaiKnowledgeBaseService) Add(ctx context.Context, params ConvaiKnowledgeBaseAddParams, opts ...option.RequestOption) (res *AddKnowledgeBaseResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Moves multiple entities from one folder to another.
func (r *ConvaiKnowledgeBaseService) BulkMoveToFolder(ctx context.Context, params ConvaiKnowledgeBaseBulkMoveToFolderParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v1/convai/knowledge-base/bulk-move"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Create a knowledge base document generated form the uploaded file.
func (r *ConvaiKnowledgeBaseService) NewFileDocument(ctx context.Context, params ConvaiKnowledgeBaseNewFileDocumentParams, opts ...option.RequestOption) (res *AddKnowledgeBaseResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a folder used for grouping documents together.
func (r *ConvaiKnowledgeBaseService) NewFolder(ctx context.Context, params ConvaiKnowledgeBaseNewFolderParams, opts ...option.RequestOption) (res *AddKnowledgeBaseResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/folder"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a knowledge base document containing the provided text.
func (r *ConvaiKnowledgeBaseService) NewTextDocument(ctx context.Context, params ConvaiKnowledgeBaseNewTextDocumentParams, opts ...option.RequestOption) (res *AddKnowledgeBaseResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/text"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Create a knowledge base document generated by scraping the given webpage.
func (r *ConvaiKnowledgeBaseService) NewURLDocument(ctx context.Context, params ConvaiKnowledgeBaseNewURLDocumentParams, opts ...option.RequestOption) (res *AddKnowledgeBaseResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/url"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a document from the knowledge base
func (r *ConvaiKnowledgeBaseService) DeleteDocument(ctx context.Context, documentationID string, params ConvaiKnowledgeBaseDeleteDocumentParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseDeleteDocumentResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return
}

// Get details about a specific documentation part used by RAG.
func (r *ConvaiKnowledgeBaseService) GetChunk(ctx context.Context, chunkID string, params ConvaiKnowledgeBaseGetChunkParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseGetChunkResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.DocumentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	if chunkID == "" {
		err = errors.New("missing required chunk_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/chunk/%s", params.DocumentationID, chunkID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get the entire content of a document from the knowledge base
func (r *ConvaiKnowledgeBaseService) GetContent(ctx context.Context, documentationID string, query ConvaiKnowledgeBaseGetContentParams, opts ...option.RequestOption) (res *string, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/html")}, opts...)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/content", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of agents depending on this knowledge base document
func (r *ConvaiKnowledgeBaseService) GetDependentAgents(ctx context.Context, documentationID string, params ConvaiKnowledgeBaseGetDependentAgentsParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseGetDependentAgentsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/dependent-agents", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Get a signed URL to download the original source file of a file-type document
// from the knowledge base
func (r *ConvaiKnowledgeBaseService) GetSourceFileURL(ctx context.Context, documentationID string, query ConvaiKnowledgeBaseGetSourceFileURLParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseGetSourceFileURLResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/source-file-url", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gets multiple knowledge base document summaries by their IDs.
func (r *ConvaiKnowledgeBaseService) ListSummaries(ctx context.Context, params ConvaiKnowledgeBaseListSummariesParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseListSummariesResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/knowledge-base/summaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Moves the entity from one folder to another.
func (r *ConvaiKnowledgeBaseService) MoveToFolder(ctx context.Context, documentID string, params ConvaiKnowledgeBaseMoveToFolderParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s/move", documentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Get details about a specific documentation making up the agent's knowledge base
func (r *ConvaiKnowledgeBaseService) GetDocument(ctx context.Context, documentationID string, params ConvaiKnowledgeBaseGetDocumentParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseGetDocumentResponseUnion, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Update the name of a document
func (r *ConvaiKnowledgeBaseService) UpdateDocument(ctx context.Context, documentationID string, params ConvaiKnowledgeBaseUpdateDocumentParams, opts ...option.RequestOption) (res *ConvaiKnowledgeBaseUpdateDocumentResponseUnion, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if documentationID == "" {
		err = errors.New("missing required documentation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/knowledge-base/%s", documentationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type AddKnowledgeBaseResponse struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		FolderPath  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddKnowledgeBaseResponse) RawJSON() string { return r.JSON.raw }
func (r *AddKnowledgeBaseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DependentAvailableAgentIdentifier struct {
	ID string `json:"id,required"`
	// Any of "admin", "editor", "commenter", "viewer".
	AccessLevel       DependentAvailableAgentIdentifierAccessLevel `json:"access_level,required"`
	CreatedAtUnixSecs int64                                        `json:"created_at_unix_secs,required"`
	Name              string                                       `json:"name,required"`
	// If the agent is a transitive dependent, contains IDs of the resources that the
	// agent depends on directly.
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available".
	Type DependentAvailableAgentIdentifierType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DependentAvailableAgentIdentifier) RawJSON() string { return r.JSON.raw }
func (r *DependentAvailableAgentIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DependentAvailableAgentIdentifierAccessLevel string

const (
	DependentAvailableAgentIdentifierAccessLevelAdmin     DependentAvailableAgentIdentifierAccessLevel = "admin"
	DependentAvailableAgentIdentifierAccessLevelEditor    DependentAvailableAgentIdentifierAccessLevel = "editor"
	DependentAvailableAgentIdentifierAccessLevelCommenter DependentAvailableAgentIdentifierAccessLevel = "commenter"
	DependentAvailableAgentIdentifierAccessLevelViewer    DependentAvailableAgentIdentifierAccessLevel = "viewer"
)

type DependentAvailableAgentIdentifierType string

const (
	DependentAvailableAgentIdentifierTypeAvailable DependentAvailableAgentIdentifierType = "available"
)

// A model that represents an agent dependent on a knowledge base/tools to which
// the user has no direct access.
type DependentUnknownAgentIdentifier struct {
	// If the agent is a transitive dependent, contains IDs of the resources that the
	// agent depends on directly.
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "unknown".
	Type DependentUnknownAgentIdentifierType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DependentUnknownAgentIdentifier) RawJSON() string { return r.JSON.raw }
func (r *DependentUnknownAgentIdentifier) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DependentUnknownAgentIdentifierType string

const (
	DependentUnknownAgentIdentifierTypeUnknown DependentUnknownAgentIdentifierType = "unknown"
)

type DocumentUsageMode string

const (
	DocumentUsageModePrompt DocumentUsageMode = "prompt"
	DocumentUsageModeAuto   DocumentUsageMode = "auto"
)

type GetKnowledgeBaseFileResponse struct {
	ID                 string                                `json:"id,required"`
	AccessInfo         ResourceAccessInfo                    `json:"access_info,required"`
	ExtractedInnerHTML string                                `json:"extracted_inner_html,required"`
	Filename           string                                `json:"filename,required"`
	Metadata           KnowledgeBaseDocumentMetadataResponse `json:"metadata,required"`
	Name               string                                `json:"name,required"`
	SupportedUsages    []DocumentUsageMode                   `json:"supported_usages,required"`
	Type               constant.File                         `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessInfo         respjson.Field
		ExtractedInnerHTML respjson.Field
		Filename           respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		SupportedUsages    respjson.Field
		Type               respjson.Field
		FolderParentID     respjson.Field
		FolderPath         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetKnowledgeBaseFileResponse) RawJSON() string { return r.JSON.raw }
func (r *GetKnowledgeBaseFileResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetKnowledgeBaseFolderResponse struct {
	ID              string                                `json:"id,required"`
	AccessInfo      ResourceAccessInfo                    `json:"access_info,required"`
	ChildrenCount   int64                                 `json:"children_count,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse `json:"metadata,required"`
	Name            string                                `json:"name,required"`
	SupportedUsages []DocumentUsageMode                   `json:"supported_usages,required"`
	Type            constant.Folder                       `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		ChildrenCount   respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetKnowledgeBaseFolderResponse) RawJSON() string { return r.JSON.raw }
func (r *GetKnowledgeBaseFolderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetKnowledgeBaseTextResponse struct {
	ID                 string                                `json:"id,required"`
	AccessInfo         ResourceAccessInfo                    `json:"access_info,required"`
	ExtractedInnerHTML string                                `json:"extracted_inner_html,required"`
	Metadata           KnowledgeBaseDocumentMetadataResponse `json:"metadata,required"`
	Name               string                                `json:"name,required"`
	SupportedUsages    []DocumentUsageMode                   `json:"supported_usages,required"`
	Type               constant.Text                         `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessInfo         respjson.Field
		ExtractedInnerHTML respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		SupportedUsages    respjson.Field
		Type               respjson.Field
		FolderParentID     respjson.Field
		FolderPath         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetKnowledgeBaseTextResponse) RawJSON() string { return r.JSON.raw }
func (r *GetKnowledgeBaseTextResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetKnowledgeBaseURLResponse struct {
	ID                 string                                `json:"id,required"`
	AccessInfo         ResourceAccessInfo                    `json:"access_info,required"`
	ExtractedInnerHTML string                                `json:"extracted_inner_html,required"`
	Metadata           KnowledgeBaseDocumentMetadataResponse `json:"metadata,required"`
	Name               string                                `json:"name,required"`
	SupportedUsages    []DocumentUsageMode                   `json:"supported_usages,required"`
	Type               constant.URL                          `json:"type,required"`
	URL                string                                `json:"url,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		AccessInfo         respjson.Field
		ExtractedInnerHTML respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		SupportedUsages    respjson.Field
		Type               respjson.Field
		URL                respjson.Field
		FolderParentID     respjson.Field
		FolderPath         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetKnowledgeBaseURLResponse) RawJSON() string { return r.JSON.raw }
func (r *GetKnowledgeBaseURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseDocumentMetadataResponse struct {
	CreatedAtUnixSecs     int64 `json:"created_at_unix_secs,required"`
	LastUpdatedAtUnixSecs int64 `json:"last_updated_at_unix_secs,required"`
	SizeBytes             int64 `json:"size_bytes,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAtUnixSecs     respjson.Field
		LastUpdatedAtUnixSecs respjson.Field
		SizeBytes             respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseDocumentMetadataResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseDocumentMetadataResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseDocumentType string

const (
	KnowledgeBaseDocumentTypeFile   KnowledgeBaseDocumentType = "file"
	KnowledgeBaseDocumentTypeURL    KnowledgeBaseDocumentType = "url"
	KnowledgeBaseDocumentTypeText   KnowledgeBaseDocumentType = "text"
	KnowledgeBaseDocumentTypeFolder KnowledgeBaseDocumentType = "folder"
)

type KnowledgeBaseFolderPathSegmentResponse struct {
	ID   string `json:"id,required"`
	Name string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseFolderPathSegmentResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseFolderPathSegmentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowledgeBaseFolderPathSegmentSummaryResponse struct {
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowledgeBaseFolderPathSegmentSummaryResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowledgeBaseFolderPathSegmentSummaryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListResponse struct {
	Documents  []ConvaiKnowledgeBaseListResponseDocumentUnion `json:"documents,required"`
	HasMore    bool                                           `json:"has_more,required"`
	NextCursor string                                         `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentUnion contains all possible properties
// and values from [ConvaiKnowledgeBaseListResponseDocumentURL],
// [ConvaiKnowledgeBaseListResponseDocumentFile],
// [ConvaiKnowledgeBaseListResponseDocumentText],
// [ConvaiKnowledgeBaseListResponseDocumentFolder].
//
// Use the [ConvaiKnowledgeBaseListResponseDocumentUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListResponseDocumentUnion struct {
	ID string `json:"id"`
	// This field is from variant [ConvaiKnowledgeBaseListResponseDocumentURL].
	AccessInfo ResourceAccessInfo `json:"access_info"`
	// This field is a union of
	// [[]ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion]
	DependentAgents ConvaiKnowledgeBaseListResponseDocumentUnionDependentAgents `json:"dependent_agents"`
	// This field is from variant [ConvaiKnowledgeBaseListResponseDocumentURL].
	Metadata        KnowledgeBaseDocumentMetadataResponse `json:"metadata"`
	Name            string                                `json:"name"`
	SupportedUsages []DocumentUsageMode                   `json:"supported_usages"`
	// Any of "url", "file", "text", "folder".
	Type string `json:"type"`
	// This field is from variant [ConvaiKnowledgeBaseListResponseDocumentURL].
	URL            string                                          `json:"url"`
	FolderParentID string                                          `json:"folder_parent_id"`
	FolderPath     []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// This field is from variant [ConvaiKnowledgeBaseListResponseDocumentFolder].
	ChildrenCount int64 `json:"children_count"`
	JSON          struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ChildrenCount   respjson.Field
		raw             string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListResponseDocument is implemented by each variant of
// [ConvaiKnowledgeBaseListResponseDocumentUnion] to add type safety for the return
// type of [ConvaiKnowledgeBaseListResponseDocumentUnion.AsAny]
type anyConvaiKnowledgeBaseListResponseDocument interface {
	implConvaiKnowledgeBaseListResponseDocumentUnion()
}

func (ConvaiKnowledgeBaseListResponseDocumentURL) implConvaiKnowledgeBaseListResponseDocumentUnion() {
}
func (ConvaiKnowledgeBaseListResponseDocumentFile) implConvaiKnowledgeBaseListResponseDocumentUnion() {
}
func (ConvaiKnowledgeBaseListResponseDocumentText) implConvaiKnowledgeBaseListResponseDocumentUnion() {
}
func (ConvaiKnowledgeBaseListResponseDocumentFolder) implConvaiKnowledgeBaseListResponseDocumentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListResponseDocumentUnion.AsAny().(type) {
//	case elevenlabs.ConvaiKnowledgeBaseListResponseDocumentURL:
//	case elevenlabs.ConvaiKnowledgeBaseListResponseDocumentFile:
//	case elevenlabs.ConvaiKnowledgeBaseListResponseDocumentText:
//	case elevenlabs.ConvaiKnowledgeBaseListResponseDocumentFolder:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListResponseDocumentUnion) AsAny() anyConvaiKnowledgeBaseListResponseDocument {
	switch u.Type {
	case "url":
		return u.AsURL()
	case "file":
		return u.AsFile()
	case "text":
		return u.AsText()
	case "folder":
		return u.AsFolder()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListResponseDocumentUnion) AsURL() (v ConvaiKnowledgeBaseListResponseDocumentURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentUnion) AsFile() (v ConvaiKnowledgeBaseListResponseDocumentFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentUnion) AsText() (v ConvaiKnowledgeBaseListResponseDocumentText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentUnion) AsFolder() (v ConvaiKnowledgeBaseListResponseDocumentFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListResponseDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseListResponseDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentUnionDependentAgents is an implicit
// subunion of [ConvaiKnowledgeBaseListResponseDocumentUnion].
// ConvaiKnowledgeBaseListResponseDocumentUnionDependentAgents provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConvaiKnowledgeBaseListResponseDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDependentAgents]
type ConvaiKnowledgeBaseListResponseDocumentUnionDependentAgents struct {
	// This field will be present if the value is a
	// [[]ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion] instead of an
	// object.
	OfDependentAgents []ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion `json:",inline"`
	JSON              struct {
		OfDependentAgents respjson.Field
		raw               string
	} `json:"-"`
}

func (r *ConvaiKnowledgeBaseListResponseDocumentUnionDependentAgents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListResponseDocumentURL struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                           `json:"metadata,required"`
	Name            string                                                          `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                             `json:"supported_usages,required"`
	Type            constant.URL                                                    `json:"type,required"`
	URL             string                                                          `json:"url,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListResponseDocumentURL) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListResponseDocumentURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion contains all
// possible properties and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListResponseDocumentURLDependentAgent is implemented by
// each variant of [ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListResponseDocumentURLDependentAgent interface {
	implConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListResponseDocumentURLDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListResponseDocumentURLDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListResponseDocumentFile struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                            `json:"metadata,required"`
	Name            string                                                           `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                              `json:"supported_usages,required"`
	Type            constant.File                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListResponseDocumentFile) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListResponseDocumentFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion contains all
// possible properties and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListResponseDocumentFileDependentAgent is implemented by
// each variant of [ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListResponseDocumentFileDependentAgent interface {
	implConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListResponseDocumentFileDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListResponseDocumentFileDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListResponseDocumentText struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                            `json:"metadata,required"`
	Name            string                                                           `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                              `json:"supported_usages,required"`
	Type            constant.Text                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListResponseDocumentText) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListResponseDocumentText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion contains all
// possible properties and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListResponseDocumentTextDependentAgent is implemented by
// each variant of [ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListResponseDocumentTextDependentAgent interface {
	implConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListResponseDocumentTextDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListResponseDocumentTextDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListResponseDocumentFolder struct {
	ID            string             `json:"id,required"`
	AccessInfo    ResourceAccessInfo `json:"access_info,required"`
	ChildrenCount int64              `json:"children_count,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                              `json:"metadata,required"`
	Name            string                                                             `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                                `json:"supported_usages,required"`
	Type            constant.Folder                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		ChildrenCount   respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListResponseDocumentFolder) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListResponseDocumentFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion contains all
// possible properties and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListResponseDocumentFolderDependentAgent is implemented by
// each variant of
// [ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion] to add type
// safety for the return type of
// [ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListResponseDocumentFolderDependentAgent interface {
	implConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListResponseDocumentFolderDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListResponseDocumentFolderDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseDeleteDocumentResponse = any

type ConvaiKnowledgeBaseGetChunkResponse struct {
	ID      string `json:"id,required"`
	Content string `json:"content,required"`
	Name    string `json:"name,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Content     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseGetChunkResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseGetChunkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseGetDependentAgentsResponse struct {
	Agents     []ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion `json:"agents,required"`
	HasMore    bool                                                      `json:"has_more,required"`
	NextCursor string                                                    `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agents      respjson.Field
		HasMore     respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseGetDependentAgentsResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseGetDependentAgentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion contains all possible
// properties and values from [DependentAvailableAgentIdentifier],
// [DependentUnknownAgentIdentifier].
//
// Use the [ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseGetDependentAgentsResponseAgent is implemented by each
// variant of [ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion] to add type
// safety for the return type of
// [ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion.AsAny]
type anyConvaiKnowledgeBaseGetDependentAgentsResponseAgent interface {
	implConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion) AsAny() anyConvaiKnowledgeBaseGetDependentAgentsResponseAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseGetDependentAgentsResponseAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseGetSourceFileURLResponse struct {
	// Signed URL to download the source file directly
	SignedURL string `json:"signed_url,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SignedURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseGetSourceFileURLResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseGetSourceFileURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponse map[string]ConvaiKnowledgeBaseListSummariesResponseItemUnion

// ConvaiKnowledgeBaseListSummariesResponseItemUnion contains all possible
// properties and values from
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccess],
// [ConvaiKnowledgeBaseListSummariesResponseItemFailure].
//
// Use the [ConvaiKnowledgeBaseListSummariesResponseItemUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemUnion struct {
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemSuccess].
	Data ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion `json:"data"`
	// Any of "success", "failure".
	Status string `json:"status"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemFailure].
	ErrorCode int64 `json:"error_code"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemFailure].
	ErrorMessage string `json:"error_message"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemFailure].
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

// anyConvaiKnowledgeBaseListSummariesResponseItem is implemented by each variant
// of [ConvaiKnowledgeBaseListSummariesResponseItemUnion] to add type safety for
// the return type of [ConvaiKnowledgeBaseListSummariesResponseItemUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItem interface {
	implConvaiKnowledgeBaseListSummariesResponseItemUnion()
}

func (ConvaiKnowledgeBaseListSummariesResponseItemSuccess) implConvaiKnowledgeBaseListSummariesResponseItemUnion() {
}
func (ConvaiKnowledgeBaseListSummariesResponseItemFailure) implConvaiKnowledgeBaseListSummariesResponseItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemUnion.AsAny().(type) {
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemSuccess:
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemFailure:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItem {
	switch u.Status {
	case "success":
		return u.AsSuccess()
	case "failure":
		return u.AsFailure()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemUnion) AsSuccess() (v ConvaiKnowledgeBaseListSummariesResponseItemSuccess) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemUnion) AsFailure() (v ConvaiKnowledgeBaseListSummariesResponseItemFailure) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseListSummariesResponseItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemSuccess struct {
	Data   ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion `json:"data,required"`
	Status constant.Success                                             `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListSummariesResponseItemSuccess) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccess) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion contains all
// possible properties and values from
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL],
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile],
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText],
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder].
//
// Use the [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion struct {
	ID string `json:"id"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL].
	AccessInfo ResourceAccessInfo `json:"access_info"`
	// This field is a union of
	// [[]ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion],
	// [[]ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion]
	DependentAgents ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnionDependentAgents `json:"dependent_agents"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL].
	Metadata        KnowledgeBaseDocumentMetadataResponse `json:"metadata"`
	Name            string                                `json:"name"`
	SupportedUsages []DocumentUsageMode                   `json:"supported_usages"`
	// Any of "url", "file", "text", "folder".
	Type string `json:"type"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL].
	URL            string                                          `json:"url"`
	FolderParentID string                                          `json:"folder_parent_id"`
	FolderPath     []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// This field is from variant
	// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder].
	ChildrenCount int64 `json:"children_count"`
	JSON          struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ChildrenCount   respjson.Field
		raw             string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListSummariesResponseItemSuccessData is implemented by
// each variant of [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItemSuccessData interface {
	implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion()
}

func (ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion() {
}
func (ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion() {
}
func (ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion() {
}
func (ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion.AsAny().(type) {
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL:
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile:
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText:
//	case elevenlabs.ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItemSuccessData {
	switch u.Type {
	case "url":
		return u.AsURL()
	case "file":
		return u.AsFile()
	case "text":
		return u.AsText()
	case "folder":
		return u.AsFolder()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) AsURL() (v ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) AsFile() (v ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) AsText() (v ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) AsFolder() (v ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnionDependentAgents is
// an implicit subunion of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion].
// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnionDependentAgents
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfDependentAgents]
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnionDependentAgents struct {
	// This field will be present if the value is a
	// [[]ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion]
	// instead of an object.
	OfDependentAgents []ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion `json:",inline"`
	JSON              struct {
		OfDependentAgents respjson.Field
		raw               string
	} `json:"-"`
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataUnionDependentAgents) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                                           `json:"metadata,required"`
	Name            string                                                                          `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                                             `json:"supported_usages,required"`
	Type            constant.URL                                                                    `json:"type,required"`
	URL             string                                                                          `json:"url,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		URL             respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion
// contains all possible properties and values from
// [DependentAvailableAgentIdentifier], [DependentUnknownAgentIdentifier].
//
// Use the
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgent is
// implemented by each variant of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgent interface {
	implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataURLDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                                            `json:"metadata,required"`
	Name            string                                                                           `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                                              `json:"supported_usages,required"`
	Type            constant.File                                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion
// contains all possible properties and values from
// [DependentAvailableAgentIdentifier], [DependentUnknownAgentIdentifier].
//
// Use the
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgent is
// implemented by each variant of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgent interface {
	implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFileDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText struct {
	ID         string             `json:"id,required"`
	AccessInfo ResourceAccessInfo `json:"access_info,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                                            `json:"metadata,required"`
	Name            string                                                                           `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                                              `json:"supported_usages,required"`
	Type            constant.Text                                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataText) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion
// contains all possible properties and values from
// [DependentAvailableAgentIdentifier], [DependentUnknownAgentIdentifier].
//
// Use the
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgent is
// implemented by each variant of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgent interface {
	implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataTextDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder struct {
	ID            string             `json:"id,required"`
	AccessInfo    ResourceAccessInfo `json:"access_info,required"`
	ChildrenCount int64              `json:"children_count,required"`
	// This field is deprecated and will be removed in the future, use the separate
	// endpoint to get dependent agents instead.
	//
	// Deprecated: deprecated
	DependentAgents []ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion `json:"dependent_agents,required"`
	Metadata        KnowledgeBaseDocumentMetadataResponse                                              `json:"metadata,required"`
	Name            string                                                                             `json:"name,required"`
	SupportedUsages []DocumentUsageMode                                                                `json:"supported_usages,required"`
	Type            constant.Folder                                                                    `json:"type,required"`
	// The ID of the parent folder, or null if the document is at the root level.
	FolderParentID string `json:"folder_parent_id,nullable"`
	// The folder path segments leading to this entity, from root to parent folder.
	FolderPath []KnowledgeBaseFolderPathSegmentSummaryResponse `json:"folder_path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AccessInfo      respjson.Field
		ChildrenCount   respjson.Field
		DependentAgents respjson.Field
		Metadata        respjson.Field
		Name            respjson.Field
		SupportedUsages respjson.Field
		Type            respjson.Field
		FolderParentID  respjson.Field
		FolderPath      respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder) RawJSON() string {
	return r.JSON.raw
}
func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion
// contains all possible properties and values from
// [DependentAvailableAgentIdentifier], [DependentUnknownAgentIdentifier].
//
// Use the
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion.AsAny]
// method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion struct {
	// This field is from variant [DependentAvailableAgentIdentifier].
	ID string `json:"id"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	AccessLevel DependentAvailableAgentIdentifierAccessLevel `json:"access_level"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs"`
	// This field is from variant [DependentAvailableAgentIdentifier].
	Name                  string   `json:"name"`
	ReferencedResourceIDs []string `json:"referenced_resource_ids"`
	// Any of "available", "unknown".
	Type string `json:"type"`
	JSON struct {
		ID                    respjson.Field
		AccessLevel           respjson.Field
		CreatedAtUnixSecs     respjson.Field
		Name                  respjson.Field
		ReferencedResourceIDs respjson.Field
		Type                  respjson.Field
		raw                   string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgent
// is implemented by each variant of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion]
// to add type safety for the return type of
// [ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion.AsAny]
type anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgent interface {
	implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion()
}

func (DependentAvailableAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion() {
}
func (DependentUnknownAgentIdentifier) implConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion.AsAny().(type) {
//	case elevenlabs.DependentAvailableAgentIdentifier:
//	case elevenlabs.DependentUnknownAgentIdentifier:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion) AsAny() anyConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgent {
	switch u.Type {
	case "available":
		return u.AsAvailable()
	case "unknown":
		return u.AsUnknown()
	}
	return nil
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion) AsAvailable() (v DependentAvailableAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion) AsUnknown() (v DependentUnknownAgentIdentifier) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ConvaiKnowledgeBaseListSummariesResponseItemSuccessDataFolderDependentAgentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListSummariesResponseItemFailure struct {
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
func (r ConvaiKnowledgeBaseListSummariesResponseItemFailure) RawJSON() string { return r.JSON.raw }
func (r *ConvaiKnowledgeBaseListSummariesResponseItemFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseGetDocumentResponseUnion contains all possible properties and
// values from [GetKnowledgeBaseURLResponse], [GetKnowledgeBaseFileResponse],
// [GetKnowledgeBaseTextResponse], [GetKnowledgeBaseFolderResponse].
//
// Use the [ConvaiKnowledgeBaseGetDocumentResponseUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseGetDocumentResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	AccessInfo         ResourceAccessInfo `json:"access_info"`
	ExtractedInnerHTML string             `json:"extracted_inner_html"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	Metadata        KnowledgeBaseDocumentMetadataResponse `json:"metadata"`
	Name            string                                `json:"name"`
	SupportedUsages []DocumentUsageMode                   `json:"supported_usages"`
	// Any of "url", "file", "text", "folder".
	Type string `json:"type"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	URL            string                                   `json:"url"`
	FolderParentID string                                   `json:"folder_parent_id"`
	FolderPath     []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// This field is from variant [GetKnowledgeBaseFileResponse].
	Filename string `json:"filename"`
	// This field is from variant [GetKnowledgeBaseFolderResponse].
	ChildrenCount int64 `json:"children_count"`
	JSON          struct {
		ID                 respjson.Field
		AccessInfo         respjson.Field
		ExtractedInnerHTML respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		SupportedUsages    respjson.Field
		Type               respjson.Field
		URL                respjson.Field
		FolderParentID     respjson.Field
		FolderPath         respjson.Field
		Filename           respjson.Field
		ChildrenCount      respjson.Field
		raw                string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseGetDocumentResponse is implemented by each variant of
// [ConvaiKnowledgeBaseGetDocumentResponseUnion] to add type safety for the return
// type of [ConvaiKnowledgeBaseGetDocumentResponseUnion.AsAny]
type anyConvaiKnowledgeBaseGetDocumentResponse interface {
	implConvaiKnowledgeBaseGetDocumentResponseUnion()
}

func (GetKnowledgeBaseURLResponse) implConvaiKnowledgeBaseGetDocumentResponseUnion()    {}
func (GetKnowledgeBaseFileResponse) implConvaiKnowledgeBaseGetDocumentResponseUnion()   {}
func (GetKnowledgeBaseTextResponse) implConvaiKnowledgeBaseGetDocumentResponseUnion()   {}
func (GetKnowledgeBaseFolderResponse) implConvaiKnowledgeBaseGetDocumentResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseGetDocumentResponseUnion.AsAny().(type) {
//	case elevenlabs.GetKnowledgeBaseURLResponse:
//	case elevenlabs.GetKnowledgeBaseFileResponse:
//	case elevenlabs.GetKnowledgeBaseTextResponse:
//	case elevenlabs.GetKnowledgeBaseFolderResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) AsAny() anyConvaiKnowledgeBaseGetDocumentResponse {
	switch u.Type {
	case "url":
		return u.AsURL()
	case "file":
		return u.AsFile()
	case "text":
		return u.AsText()
	case "folder":
		return u.AsFolder()
	}
	return nil
}

func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) AsURL() (v GetKnowledgeBaseURLResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) AsFile() (v GetKnowledgeBaseFileResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) AsText() (v GetKnowledgeBaseTextResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) AsFolder() (v GetKnowledgeBaseFolderResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseGetDocumentResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseGetDocumentResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ConvaiKnowledgeBaseUpdateDocumentResponseUnion contains all possible properties
// and values from [GetKnowledgeBaseURLResponse], [GetKnowledgeBaseFileResponse],
// [GetKnowledgeBaseTextResponse], [GetKnowledgeBaseFolderResponse].
//
// Use the [ConvaiKnowledgeBaseUpdateDocumentResponseUnion.AsAny] method to switch
// on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ConvaiKnowledgeBaseUpdateDocumentResponseUnion struct {
	ID string `json:"id"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	AccessInfo         ResourceAccessInfo `json:"access_info"`
	ExtractedInnerHTML string             `json:"extracted_inner_html"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	Metadata        KnowledgeBaseDocumentMetadataResponse `json:"metadata"`
	Name            string                                `json:"name"`
	SupportedUsages []DocumentUsageMode                   `json:"supported_usages"`
	// Any of "url", "file", "text", "folder".
	Type string `json:"type"`
	// This field is from variant [GetKnowledgeBaseURLResponse].
	URL            string                                   `json:"url"`
	FolderParentID string                                   `json:"folder_parent_id"`
	FolderPath     []KnowledgeBaseFolderPathSegmentResponse `json:"folder_path"`
	// This field is from variant [GetKnowledgeBaseFileResponse].
	Filename string `json:"filename"`
	// This field is from variant [GetKnowledgeBaseFolderResponse].
	ChildrenCount int64 `json:"children_count"`
	JSON          struct {
		ID                 respjson.Field
		AccessInfo         respjson.Field
		ExtractedInnerHTML respjson.Field
		Metadata           respjson.Field
		Name               respjson.Field
		SupportedUsages    respjson.Field
		Type               respjson.Field
		URL                respjson.Field
		FolderParentID     respjson.Field
		FolderPath         respjson.Field
		Filename           respjson.Field
		ChildrenCount      respjson.Field
		raw                string
	} `json:"-"`
}

// anyConvaiKnowledgeBaseUpdateDocumentResponse is implemented by each variant of
// [ConvaiKnowledgeBaseUpdateDocumentResponseUnion] to add type safety for the
// return type of [ConvaiKnowledgeBaseUpdateDocumentResponseUnion.AsAny]
type anyConvaiKnowledgeBaseUpdateDocumentResponse interface {
	implConvaiKnowledgeBaseUpdateDocumentResponseUnion()
}

func (GetKnowledgeBaseURLResponse) implConvaiKnowledgeBaseUpdateDocumentResponseUnion()    {}
func (GetKnowledgeBaseFileResponse) implConvaiKnowledgeBaseUpdateDocumentResponseUnion()   {}
func (GetKnowledgeBaseTextResponse) implConvaiKnowledgeBaseUpdateDocumentResponseUnion()   {}
func (GetKnowledgeBaseFolderResponse) implConvaiKnowledgeBaseUpdateDocumentResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := ConvaiKnowledgeBaseUpdateDocumentResponseUnion.AsAny().(type) {
//	case elevenlabs.GetKnowledgeBaseURLResponse:
//	case elevenlabs.GetKnowledgeBaseFileResponse:
//	case elevenlabs.GetKnowledgeBaseTextResponse:
//	case elevenlabs.GetKnowledgeBaseFolderResponse:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) AsAny() anyConvaiKnowledgeBaseUpdateDocumentResponse {
	switch u.Type {
	case "url":
		return u.AsURL()
	case "file":
		return u.AsFile()
	case "text":
		return u.AsText()
	case "folder":
		return u.AsFolder()
	}
	return nil
}

func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) AsURL() (v GetKnowledgeBaseURLResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) AsFile() (v GetKnowledgeBaseFileResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) AsText() (v GetKnowledgeBaseTextResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) AsFolder() (v GetKnowledgeBaseFolderResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ConvaiKnowledgeBaseUpdateDocumentResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *ConvaiKnowledgeBaseUpdateDocumentResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseListParams struct {
	// If set, the endpoint will return only documents that are descendants of the
	// given folder.
	AncestorFolderID param.Opt[string] `query:"ancestor_folder_id,omitzero" json:"-"`
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// If set, the endpoint will return only documents that are direct children of the
	// given folder.
	ParentFolderID param.Opt[string] `query:"parent_folder_id,omitzero" json:"-"`
	// If specified, the endpoint returns only such knowledge base documents whose
	// names start with this string.
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// Whether folders should be returned first in the list of documents.
	FoldersFirst param.Opt[bool] `query:"folders_first,omitzero" json:"-"`
	// How many documents to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// If set to true, the endpoint will return only documents owned by you (and not
	// shared from somebody else).
	ShowOnlyOwnedDocuments param.Opt[bool] `query:"show_only_owned_documents,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// The field to sort the results by
	//
	// Any of "name", "created_at", "updated_at", "size".
	SortBy ConvaiKnowledgeBaseListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// If present, the endpoint will return only documents of the given types.
	Types []KnowledgeBaseDocumentType `query:"types,omitzero" json:"-"`
	// The direction to sort the results
	//
	// Any of "asc", "desc".
	SortDirection SortDirection `query:"sort_direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiKnowledgeBaseListParams]'s query parameters as
// `url.Values`.
func (r ConvaiKnowledgeBaseListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The field to sort the results by
type ConvaiKnowledgeBaseListParamsSortBy string

const (
	ConvaiKnowledgeBaseListParamsSortByName      ConvaiKnowledgeBaseListParamsSortBy = "name"
	ConvaiKnowledgeBaseListParamsSortByCreatedAt ConvaiKnowledgeBaseListParamsSortBy = "created_at"
	ConvaiKnowledgeBaseListParamsSortByUpdatedAt ConvaiKnowledgeBaseListParamsSortBy = "updated_at"
	ConvaiKnowledgeBaseListParamsSortBySize      ConvaiKnowledgeBaseListParamsSortBy = "size"
)

type ConvaiKnowledgeBaseAddParams struct {
	// A custom, human-readable name for the document.
	Name    param.Opt[string] `json:"name,omitzero"`
	AgentID param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// URL to a page of documentation that the agent will have access to in order to
	// interact with users.
	URL param.Opt[string] `json:"url,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Documentation that the agent will have access to in order to interact with
	// users.
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r ConvaiKnowledgeBaseAddParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [ConvaiKnowledgeBaseAddParams]'s query parameters as
// `url.Values`.
func (r ConvaiKnowledgeBaseAddParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiKnowledgeBaseBulkMoveToFolderParams struct {
	// The ids of documents or folders from the knowledge base.
	DocumentIDs []string `json:"document_ids,omitzero,required"`
	// The folder to move the entities to. If not set, the entities will be moved to
	// the root folder.
	MoveTo param.Opt[string] `json:"move_to,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseBulkMoveToFolderParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseBulkMoveToFolderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseBulkMoveToFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseNewFileDocumentParams struct {
	// Documentation that the agent will have access to in order to interact with
	// users.
	File io.Reader `json:"file,omitzero,required" format:"binary"`
	// A custom, human-readable name for the document.
	Name param.Opt[string] `json:"name,omitzero"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderID param.Opt[string] `json:"parent_folder_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseNewFileDocumentParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

type ConvaiKnowledgeBaseNewFolderParams struct {
	// A custom, human-readable name for the document.
	Name string `json:"name,required"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderID param.Opt[string] `json:"parent_folder_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseNewFolderParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseNewFolderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseNewFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseNewTextDocumentParams struct {
	// Text content to be added to the knowledge base.
	Text string `json:"text,required"`
	// A custom, human-readable name for the document.
	Name param.Opt[string] `json:"name,omitzero"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderID param.Opt[string] `json:"parent_folder_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseNewTextDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseNewTextDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseNewTextDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseNewURLDocumentParams struct {
	// URL to a page of documentation that the agent will have access to in order to
	// interact with users.
	URL string `json:"url,required"`
	// A custom, human-readable name for the document.
	Name param.Opt[string] `json:"name,omitzero"`
	// If set, the created document or folder will be placed inside the given folder.
	ParentFolderID param.Opt[string] `json:"parent_folder_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseNewURLDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseNewURLDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseNewURLDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseDeleteDocumentParams struct {
	// If set to true, the document will be deleted regardless of whether it is used by
	// any agents and it will be deleted from the dependent agents.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiKnowledgeBaseDeleteDocumentParams]'s query parameters
// as `url.Values`.
func (r ConvaiKnowledgeBaseDeleteDocumentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiKnowledgeBaseGetChunkParams struct {
	// The id of a document from the knowledge base. This is returned on document
	// addition.
	DocumentationID string `path:"documentation_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiKnowledgeBaseGetContentParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiKnowledgeBaseGetDependentAgentsParams struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// How many documents to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Type of dependent agents to return.
	//
	// Any of "direct", "transitive", "all".
	DependentType ConvaiKnowledgeBaseGetDependentAgentsParamsDependentType `query:"dependent_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiKnowledgeBaseGetDependentAgentsParams]'s query
// parameters as `url.Values`.
func (r ConvaiKnowledgeBaseGetDependentAgentsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Type of dependent agents to return.
type ConvaiKnowledgeBaseGetDependentAgentsParamsDependentType string

const (
	ConvaiKnowledgeBaseGetDependentAgentsParamsDependentTypeDirect     ConvaiKnowledgeBaseGetDependentAgentsParamsDependentType = "direct"
	ConvaiKnowledgeBaseGetDependentAgentsParamsDependentTypeTransitive ConvaiKnowledgeBaseGetDependentAgentsParamsDependentType = "transitive"
	ConvaiKnowledgeBaseGetDependentAgentsParamsDependentTypeAll        ConvaiKnowledgeBaseGetDependentAgentsParamsDependentType = "all"
)

type ConvaiKnowledgeBaseGetSourceFileURLParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiKnowledgeBaseListSummariesParams struct {
	// The ids of knowledge base documents.
	DocumentIDs []string `query:"document_ids,omitzero,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiKnowledgeBaseListSummariesParams]'s query parameters
// as `url.Values`.
func (r ConvaiKnowledgeBaseListSummariesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiKnowledgeBaseMoveToFolderParams struct {
	// The folder to move the entities to. If not set, the entities will be moved to
	// the root folder.
	MoveTo param.Opt[string] `json:"move_to,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseMoveToFolderParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseMoveToFolderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseMoveToFolderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiKnowledgeBaseGetDocumentParams struct {
	AgentID param.Opt[string] `query:"agent_id,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiKnowledgeBaseGetDocumentParams]'s query parameters as
// `url.Values`.
func (r ConvaiKnowledgeBaseGetDocumentParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiKnowledgeBaseUpdateDocumentParams struct {
	// A custom, human-readable name for the document.
	Name string `json:"name,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r ConvaiKnowledgeBaseUpdateDocumentParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiKnowledgeBaseUpdateDocumentParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiKnowledgeBaseUpdateDocumentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
