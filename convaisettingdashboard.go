// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// ConvaiSettingDashboardService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConvaiSettingDashboardService] method instead.
type ConvaiSettingDashboardService struct {
	Options []option.RequestOption
}

// NewConvaiSettingDashboardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConvaiSettingDashboardService(opts ...option.RequestOption) (r ConvaiSettingDashboardService) {
	r = ConvaiSettingDashboardService{}
	r.Options = opts
	return
}

// Retrieve Convai dashboard settings for the workspace
func (r *ConvaiSettingDashboardService) Get(ctx context.Context, query ConvaiSettingDashboardGetParams, opts ...option.RequestOption) (res *GetConvAIDashboardSettingsResponseModel, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/settings/dashboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update Convai dashboard settings for the workspace
func (r *ConvaiSettingDashboardService) Update(ctx context.Context, params ConvaiSettingDashboardUpdateParams, opts ...option.RequestOption) (res *GetConvAIDashboardSettingsResponseModel, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/convai/settings/dashboard"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

type DashboardCallSuccessChartModel struct {
	Name string `json:"name,required"`
	// Any of "call_success".
	Type DashboardCallSuccessChartModelType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DashboardCallSuccessChartModel) RawJSON() string { return r.JSON.raw }
func (r *DashboardCallSuccessChartModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DashboardCallSuccessChartModel to a
// DashboardCallSuccessChartModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DashboardCallSuccessChartModelParam.Overrides()
func (r DashboardCallSuccessChartModel) ToParam() DashboardCallSuccessChartModelParam {
	return param.Override[DashboardCallSuccessChartModelParam](json.RawMessage(r.RawJSON()))
}

type DashboardCallSuccessChartModelType string

const (
	DashboardCallSuccessChartModelTypeCallSuccess DashboardCallSuccessChartModelType = "call_success"
)

// The property Name is required.
type DashboardCallSuccessChartModelParam struct {
	Name string `json:"name,required"`
	// Any of "call_success".
	Type DashboardCallSuccessChartModelType `json:"type,omitzero"`
	paramObj
}

func (r DashboardCallSuccessChartModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DashboardCallSuccessChartModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DashboardCallSuccessChartModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardCriteriaChartModel struct {
	CriteriaID string `json:"criteria_id,required"`
	Name       string `json:"name,required"`
	// Any of "criteria".
	Type DashboardCriteriaChartModelType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CriteriaID  respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DashboardCriteriaChartModel) RawJSON() string { return r.JSON.raw }
func (r *DashboardCriteriaChartModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DashboardCriteriaChartModel to a
// DashboardCriteriaChartModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DashboardCriteriaChartModelParam.Overrides()
func (r DashboardCriteriaChartModel) ToParam() DashboardCriteriaChartModelParam {
	return param.Override[DashboardCriteriaChartModelParam](json.RawMessage(r.RawJSON()))
}

type DashboardCriteriaChartModelType string

const (
	DashboardCriteriaChartModelTypeCriteria DashboardCriteriaChartModelType = "criteria"
)

// The properties CriteriaID, Name are required.
type DashboardCriteriaChartModelParam struct {
	CriteriaID string `json:"criteria_id,required"`
	Name       string `json:"name,required"`
	// Any of "criteria".
	Type DashboardCriteriaChartModelType `json:"type,omitzero"`
	paramObj
}

func (r DashboardCriteriaChartModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DashboardCriteriaChartModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DashboardCriteriaChartModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardDataCollectionChartModel struct {
	DataCollectionID string `json:"data_collection_id,required"`
	Name             string `json:"name,required"`
	// Any of "data_collection".
	Type DashboardDataCollectionChartModelType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataCollectionID respjson.Field
		Name             respjson.Field
		Type             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DashboardDataCollectionChartModel) RawJSON() string { return r.JSON.raw }
func (r *DashboardDataCollectionChartModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DashboardDataCollectionChartModel to a
// DashboardDataCollectionChartModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DashboardDataCollectionChartModelParam.Overrides()
func (r DashboardDataCollectionChartModel) ToParam() DashboardDataCollectionChartModelParam {
	return param.Override[DashboardDataCollectionChartModelParam](json.RawMessage(r.RawJSON()))
}

type DashboardDataCollectionChartModelType string

const (
	DashboardDataCollectionChartModelTypeDataCollection DashboardDataCollectionChartModelType = "data_collection"
)

// The properties DataCollectionID, Name are required.
type DashboardDataCollectionChartModelParam struct {
	DataCollectionID string `json:"data_collection_id,required"`
	Name             string `json:"name,required"`
	// Any of "data_collection".
	Type DashboardDataCollectionChartModelType `json:"type,omitzero"`
	paramObj
}

func (r DashboardDataCollectionChartModelParam) MarshalJSON() (data []byte, err error) {
	type shadow DashboardDataCollectionChartModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DashboardDataCollectionChartModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetConvAIDashboardSettingsResponseModel struct {
	Charts []GetConvAIDashboardSettingsResponseModelChartUnion `json:"charts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Charts      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetConvAIDashboardSettingsResponseModel) RawJSON() string { return r.JSON.raw }
func (r *GetConvAIDashboardSettingsResponseModel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetConvAIDashboardSettingsResponseModelChartUnion contains all possible
// properties and values from [DashboardCallSuccessChartModel],
// [DashboardCriteriaChartModel], [DashboardDataCollectionChartModel].
//
// Use the [GetConvAIDashboardSettingsResponseModelChartUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type GetConvAIDashboardSettingsResponseModelChartUnion struct {
	Name string `json:"name"`
	// Any of "call_success", "criteria", "data_collection".
	Type string `json:"type"`
	// This field is from variant [DashboardCriteriaChartModel].
	CriteriaID string `json:"criteria_id"`
	// This field is from variant [DashboardDataCollectionChartModel].
	DataCollectionID string `json:"data_collection_id"`
	JSON             struct {
		Name             respjson.Field
		Type             respjson.Field
		CriteriaID       respjson.Field
		DataCollectionID respjson.Field
		raw              string
	} `json:"-"`
}

// anyGetConvAIDashboardSettingsResponseModelChart is implemented by each variant
// of [GetConvAIDashboardSettingsResponseModelChartUnion] to add type safety for
// the return type of [GetConvAIDashboardSettingsResponseModelChartUnion.AsAny]
type anyGetConvAIDashboardSettingsResponseModelChart interface {
	implGetConvAIDashboardSettingsResponseModelChartUnion()
}

func (DashboardCallSuccessChartModel) implGetConvAIDashboardSettingsResponseModelChartUnion()    {}
func (DashboardCriteriaChartModel) implGetConvAIDashboardSettingsResponseModelChartUnion()       {}
func (DashboardDataCollectionChartModel) implGetConvAIDashboardSettingsResponseModelChartUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := GetConvAIDashboardSettingsResponseModelChartUnion.AsAny().(type) {
//	case elevenlabs.DashboardCallSuccessChartModel:
//	case elevenlabs.DashboardCriteriaChartModel:
//	case elevenlabs.DashboardDataCollectionChartModel:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u GetConvAIDashboardSettingsResponseModelChartUnion) AsAny() anyGetConvAIDashboardSettingsResponseModelChart {
	switch u.Type {
	case "call_success":
		return u.AsCallSuccess()
	case "criteria":
		return u.AsCriteria()
	case "data_collection":
		return u.AsDataCollection()
	}
	return nil
}

func (u GetConvAIDashboardSettingsResponseModelChartUnion) AsCallSuccess() (v DashboardCallSuccessChartModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetConvAIDashboardSettingsResponseModelChartUnion) AsCriteria() (v DashboardCriteriaChartModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GetConvAIDashboardSettingsResponseModelChartUnion) AsDataCollection() (v DashboardDataCollectionChartModel) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GetConvAIDashboardSettingsResponseModelChartUnion) RawJSON() string { return u.JSON.raw }

func (r *GetConvAIDashboardSettingsResponseModelChartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConvaiSettingDashboardGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type ConvaiSettingDashboardUpdateParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string]                              `header:"xi-api-key,omitzero" json:"-"`
	Charts   []ConvaiSettingDashboardUpdateParamsChartUnion `json:"charts,omitzero"`
	paramObj
}

func (r ConvaiSettingDashboardUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ConvaiSettingDashboardUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConvaiSettingDashboardUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ConvaiSettingDashboardUpdateParamsChartUnion struct {
	OfCallSuccess    *DashboardCallSuccessChartModelParam    `json:",omitzero,inline"`
	OfCriteria       *DashboardCriteriaChartModelParam       `json:",omitzero,inline"`
	OfDataCollection *DashboardDataCollectionChartModelParam `json:",omitzero,inline"`
	paramUnion
}

func (u ConvaiSettingDashboardUpdateParamsChartUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCallSuccess, u.OfCriteria, u.OfDataCollection)
}
func (u *ConvaiSettingDashboardUpdateParamsChartUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ConvaiSettingDashboardUpdateParamsChartUnion) asAny() any {
	if !param.IsOmitted(u.OfCallSuccess) {
		return u.OfCallSuccess
	} else if !param.IsOmitted(u.OfCriteria) {
		return u.OfCriteria
	} else if !param.IsOmitted(u.OfDataCollection) {
		return u.OfDataCollection
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConvaiSettingDashboardUpdateParamsChartUnion) GetCriteriaID() *string {
	if vt := u.OfCriteria; vt != nil {
		return &vt.CriteriaID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConvaiSettingDashboardUpdateParamsChartUnion) GetDataCollectionID() *string {
	if vt := u.OfDataCollection; vt != nil {
		return &vt.DataCollectionID
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConvaiSettingDashboardUpdateParamsChartUnion) GetName() *string {
	if vt := u.OfCallSuccess; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfCriteria; vt != nil {
		return (*string)(&vt.Name)
	} else if vt := u.OfDataCollection; vt != nil {
		return (*string)(&vt.Name)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ConvaiSettingDashboardUpdateParamsChartUnion) GetType() *string {
	if vt := u.OfCallSuccess; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfCriteria; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfDataCollection; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[ConvaiSettingDashboardUpdateParamsChartUnion](
		"type",
		apijson.Discriminator[DashboardCallSuccessChartModelParam]("call_success"),
		apijson.Discriminator[DashboardCriteriaChartModelParam]("criteria"),
		apijson.Discriminator[DashboardDataCollectionChartModelParam]("data_collection"),
	)
}
