// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/apiquery"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiTestInvocationService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiTestInvocationService] method instead.
type ConvaiTestInvocationService struct {
	Options []option.RequestOption
}

// NewConvaiTestInvocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiTestInvocationService(opts ...option.RequestOption) (r ConvaiTestInvocationService) {
	r = ConvaiTestInvocationService{}
	r.Options = opts
	return
}

// Gets a test invocation by ID.
func (r *ConvaiTestInvocationService) Get(ctx context.Context, testInvocationID string, query ConvaiTestInvocationGetParams, opts ...option.RequestOption) (res *GetTestSuiteInvocationResponseModel, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if testInvocationID == "" {
		err = errors.New("missing required test_invocation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/test-invocations/%s", testInvocationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all test invocations with pagination support and optional search
// filtering.
func (r *ConvaiTestInvocationService) List(ctx context.Context, params ConvaiTestInvocationListParams, opts ...option.RequestOption) (res *ConvaiTestInvocationListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/test-invocations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Resubmits specific test runs from a test invocation.
func (r *ConvaiTestInvocationService) Resubmit(ctx context.Context, testInvocationID string, params ConvaiTestInvocationResubmitParams, opts ...option.RequestOption) (res *ConvaiTestInvocationResubmitResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if testInvocationID == "" {
		err = errors.New("missing required test_invocation_id parameter")
		return
	}
	path := fmt.Sprintf("v1/convai/test-invocations/%s/resubmit", testInvocationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type ConvaiTestInvocationListResponse struct {
	// Whether there are more results available
	HasMore bool                                     `json:"has_more,required"`
	Results []ConvaiTestInvocationListResponseResult `json:"results,required"`
	Meta    ConvaiTestInvocationListResponseMeta     `json:"meta"`
	// Cursor for the next page of results
	NextCursor string `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Results     respjson.Field
		Meta        respjson.Field
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiTestInvocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConvaiTestInvocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTestInvocationListResponseResult struct {
	// The ID of the test invocation
	ID string `json:"id,required"`
	// Creation time of the test invocation in unix seconds
	CreatedAtUnixSecs int64 `json:"created_at_unix_secs,required"`
	// Number of test runs that failed
	FailedCount int64 `json:"failed_count,required"`
	// Number of test runs that passed
	PassedCount int64 `json:"passed_count,required"`
	// Number of test runs that are pending
	PendingCount int64 `json:"pending_count,required"`
	// Number of test runs in this invocation
	TestRunCount int64 `json:"test_run_count,required"`
	// Title of the test invocation - either the single test name or count of tests
	Title string `json:"title,required"`
	// The access information of the test invocation
	AccessInfo ResourceAccessInfo `json:"access_info,nullable"`
	// The ID of the agent this test invocation belongs to
	AgentID string `json:"agent_id,nullable"`
	// The ID of the branch this test invocation was run on
	BranchID string `json:"branch_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		CreatedAtUnixSecs respjson.Field
		FailedCount       respjson.Field
		PassedCount       respjson.Field
		PendingCount      respjson.Field
		TestRunCount      respjson.Field
		Title             respjson.Field
		AccessInfo        respjson.Field
		AgentID           respjson.Field
		BranchID          respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiTestInvocationListResponseResult) RawJSON() string { return r.JSON.raw }
func (r *ConvaiTestInvocationListResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTestInvocationListResponseMeta struct {
	Page     int64 `json:"page,nullable"`
	PageSize int64 `json:"page_size,nullable"`
	Total    int64 `json:"total,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page        respjson.Field
		PageSize    respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConvaiTestInvocationListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *ConvaiTestInvocationListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiTestInvocationResubmitResponse = any

type ConvaiTestInvocationGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiTestInvocationListParams struct {
	// Filter by agent ID
	AgentID string `query:"agent_id,required" json:"-"`
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// How many Tests to return at maximum. Can not exceed 100, defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ConvaiTestInvocationListParams]'s query parameters as
// `url.Values`.
func (r ConvaiTestInvocationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ConvaiTestInvocationResubmitParams struct {
	// Agent ID to resubmit tests for
	AgentID string `json:"agent_id,required"`
	// List of test run IDs to resubmit
	TestRunIDs []string `json:"test_run_ids,omitzero,required"`
	// ID of the branch to run the tests on. If not provided, the tests will be run on
	// the agent default configuration.
	BranchID param.Opt[string] `json:"branch_id,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Configuration overrides to use for testing. If not provided, the agent's default
	// configuration will be used.
	AgentConfigOverride AdhocAgentConfigOverrideForTestRequestModelParam `json:"agent_config_override,omitzero"`
	paramObj
}

func (r ConvaiTestInvocationResubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiTestInvocationResubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiTestInvocationResubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
