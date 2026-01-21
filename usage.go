// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
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

// UsageService contains methods and other services that help with interacting with
// the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUsageService] method instead.
type UsageService struct {
	Options []option.RequestOption
}

// NewUsageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUsageService(opts ...option.RequestOption) (r UsageService) {
	r = UsageService{}
	r.Options = opts
	return
}

// Returns the usage metrics for the current user or the entire workspace they are
// part of. The response provides a time axis based on the specified aggregation
// interval (default: day), with usage values for each interval along that axis.
// Usage is broken down by the selected breakdown type. For example, breakdown type
// "voice" will return the usage of each voice for each interval along the time
// axis.
func (r *UsageService) GetCharacterStats(ctx context.Context, params UsageGetCharacterStatsParams, opts ...option.RequestOption) (res *UsageGetCharacterStatsResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/usage/character-stats"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

type UsageGetCharacterStatsResponse struct {
	// The time axis with unix timestamps for each day.
	Time []int64 `json:"time,required"`
	// The usage of each breakdown type along the time axis.
	Usage map[string][]float64 `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Time        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UsageGetCharacterStatsResponse) RawJSON() string { return r.JSON.raw }
func (r *UsageGetCharacterStatsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UsageGetCharacterStatsParams struct {
	// UTC Unix timestamp for the end of the usage window, in milliseconds. To include
	// the last day of the window, the timestamp should be at 23:59:59 of that day.
	EndUnix int64 `query:"end_unix,required" json:"-"`
	// UTC Unix timestamp for the start of the usage window, in milliseconds. To
	// include the first day of the window, the timestamp should be at 00:00:00 of that
	// day.
	StartUnix int64 `query:"start_unix,required" json:"-"`
	// Aggregation bucket size in seconds. Overrides the aggregation interval.
	AggregationBucketSize param.Opt[int64] `query:"aggregation_bucket_size,omitzero" json:"-"`
	// Whether or not to include the statistics of the entire workspace.
	IncludeWorkspaceMetrics param.Opt[bool] `query:"include_workspace_metrics,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// How to aggregate usage data over time. Can be "hour", "day", "week", "month", or
	// "cumulative".
	//
	// Any of "hour", "day", "week", "month", "cumulative".
	AggregationInterval UsageGetCharacterStatsParamsAggregationInterval `query:"aggregation_interval,omitzero" json:"-"`
	// How to break down the information. Cannot be "user" if include_workspace_metrics
	// is False.
	//
	// Any of "none", "voice", "voice_multiplier", "user", "groups", "api_keys",
	// "all_api_keys", "product_type", "model", "resource", "request_queue", "region",
	// "subresource_id", "reporting_workspace_id", "has_api_key", "request_source".
	BreakdownType UsageGetCharacterStatsParamsBreakdownType `query:"breakdown_type,omitzero" json:"-"`
	// Which metric to aggregate.
	//
	// Any of "credits", "tts_characters", "minutes_used", "request_count", "ttfb_avg",
	// "ttfb_p95", "fiat_units_spent", "concurrency", "concurrency_average".
	Metric UsageGetCharacterStatsParamsMetric `query:"metric,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UsageGetCharacterStatsParams]'s query parameters as
// `url.Values`.
func (r UsageGetCharacterStatsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// How to aggregate usage data over time. Can be "hour", "day", "week", "month", or
// "cumulative".
type UsageGetCharacterStatsParamsAggregationInterval string

const (
	UsageGetCharacterStatsParamsAggregationIntervalHour       UsageGetCharacterStatsParamsAggregationInterval = "hour"
	UsageGetCharacterStatsParamsAggregationIntervalDay        UsageGetCharacterStatsParamsAggregationInterval = "day"
	UsageGetCharacterStatsParamsAggregationIntervalWeek       UsageGetCharacterStatsParamsAggregationInterval = "week"
	UsageGetCharacterStatsParamsAggregationIntervalMonth      UsageGetCharacterStatsParamsAggregationInterval = "month"
	UsageGetCharacterStatsParamsAggregationIntervalCumulative UsageGetCharacterStatsParamsAggregationInterval = "cumulative"
)

// How to break down the information. Cannot be "user" if include_workspace_metrics
// is False.
type UsageGetCharacterStatsParamsBreakdownType string

const (
	UsageGetCharacterStatsParamsBreakdownTypeNone                 UsageGetCharacterStatsParamsBreakdownType = "none"
	UsageGetCharacterStatsParamsBreakdownTypeVoice                UsageGetCharacterStatsParamsBreakdownType = "voice"
	UsageGetCharacterStatsParamsBreakdownTypeVoiceMultiplier      UsageGetCharacterStatsParamsBreakdownType = "voice_multiplier"
	UsageGetCharacterStatsParamsBreakdownTypeUser                 UsageGetCharacterStatsParamsBreakdownType = "user"
	UsageGetCharacterStatsParamsBreakdownTypeGroups               UsageGetCharacterStatsParamsBreakdownType = "groups"
	UsageGetCharacterStatsParamsBreakdownTypeAPIKeys              UsageGetCharacterStatsParamsBreakdownType = "api_keys"
	UsageGetCharacterStatsParamsBreakdownTypeAllAPIKeys           UsageGetCharacterStatsParamsBreakdownType = "all_api_keys"
	UsageGetCharacterStatsParamsBreakdownTypeProductType          UsageGetCharacterStatsParamsBreakdownType = "product_type"
	UsageGetCharacterStatsParamsBreakdownTypeModel                UsageGetCharacterStatsParamsBreakdownType = "model"
	UsageGetCharacterStatsParamsBreakdownTypeResource             UsageGetCharacterStatsParamsBreakdownType = "resource"
	UsageGetCharacterStatsParamsBreakdownTypeRequestQueue         UsageGetCharacterStatsParamsBreakdownType = "request_queue"
	UsageGetCharacterStatsParamsBreakdownTypeRegion               UsageGetCharacterStatsParamsBreakdownType = "region"
	UsageGetCharacterStatsParamsBreakdownTypeSubresourceID        UsageGetCharacterStatsParamsBreakdownType = "subresource_id"
	UsageGetCharacterStatsParamsBreakdownTypeReportingWorkspaceID UsageGetCharacterStatsParamsBreakdownType = "reporting_workspace_id"
	UsageGetCharacterStatsParamsBreakdownTypeHasAPIKey            UsageGetCharacterStatsParamsBreakdownType = "has_api_key"
	UsageGetCharacterStatsParamsBreakdownTypeRequestSource        UsageGetCharacterStatsParamsBreakdownType = "request_source"
)

// Which metric to aggregate.
type UsageGetCharacterStatsParamsMetric string

const (
	UsageGetCharacterStatsParamsMetricCredits            UsageGetCharacterStatsParamsMetric = "credits"
	UsageGetCharacterStatsParamsMetricTtsCharacters      UsageGetCharacterStatsParamsMetric = "tts_characters"
	UsageGetCharacterStatsParamsMetricMinutesUsed        UsageGetCharacterStatsParamsMetric = "minutes_used"
	UsageGetCharacterStatsParamsMetricRequestCount       UsageGetCharacterStatsParamsMetric = "request_count"
	UsageGetCharacterStatsParamsMetricTtfbAvg            UsageGetCharacterStatsParamsMetric = "ttfb_avg"
	UsageGetCharacterStatsParamsMetricTtfbP95            UsageGetCharacterStatsParamsMetric = "ttfb_p95"
	UsageGetCharacterStatsParamsMetricFiatUnitsSpent     UsageGetCharacterStatsParamsMetric = "fiat_units_spent"
	UsageGetCharacterStatsParamsMetricConcurrency        UsageGetCharacterStatsParamsMetric = "concurrency"
	UsageGetCharacterStatsParamsMetricConcurrencyAverage UsageGetCharacterStatsParamsMetric = "concurrency_average"
)
