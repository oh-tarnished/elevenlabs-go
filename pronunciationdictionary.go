// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package elevenlabs

import (
	"bytes"
	"context"
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

// PronunciationDictionaryService contains methods and other services that help
// with interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPronunciationDictionaryService] method instead.
type PronunciationDictionaryService struct {
	Options []option.RequestOption
}

// NewPronunciationDictionaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPronunciationDictionaryService(opts ...option.RequestOption) (r PronunciationDictionaryService) {
	r = PronunciationDictionaryService{}
	r.Options = opts
	return
}

// Get metadata for a pronunciation dictionary
func (r *PronunciationDictionaryService) Get(ctx context.Context, pronunciationDictionaryID string, query PronunciationDictionaryGetParams, opts ...option.RequestOption) (res *GetPronunciationDictionaryMetadata, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if pronunciationDictionaryID == "" {
		err = errors.New("missing required pronunciation_dictionary_id parameter")
		return
	}
	path := fmt.Sprintf("v1/pronunciation-dictionaries/%s", pronunciationDictionaryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Partially update the pronunciation dictionary without changing the version
func (r *PronunciationDictionaryService) Update(ctx context.Context, pronunciationDictionaryID string, params PronunciationDictionaryUpdateParams, opts ...option.RequestOption) (res *GetPronunciationDictionaryMetadata, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if pronunciationDictionaryID == "" {
		err = errors.New("missing required pronunciation_dictionary_id parameter")
		return
	}
	path := fmt.Sprintf("v1/pronunciation-dictionaries/%s", pronunciationDictionaryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Get a list of the pronunciation dictionaries you have access to and their
// metadata
func (r *PronunciationDictionaryService) List(ctx context.Context, params PronunciationDictionaryListParams, opts ...option.RequestOption) (res *PronunciationDictionaryListResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/pronunciation-dictionaries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return
}

// Add rules to the pronunciation dictionary. If a rule with the same
// string_to_replace already exists, it will be replaced.
func (r *PronunciationDictionaryService) AddRules(ctx context.Context, pronunciationDictionaryID string, params PronunciationDictionaryAddRulesParams, opts ...option.RequestOption) (res *PronunciationDictionaryRules, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if pronunciationDictionaryID == "" {
		err = errors.New("missing required pronunciation_dictionary_id parameter")
		return
	}
	path := fmt.Sprintf("v1/pronunciation-dictionaries/%s/add-rules", pronunciationDictionaryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a new pronunciation dictionary from a lexicon .PLS file
func (r *PronunciationDictionaryService) NewFromFile(ctx context.Context, params PronunciationDictionaryNewFromFileParams, opts ...option.RequestOption) (res *AddPronunciationDictionary, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/pronunciation-dictionaries/add-from-file"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Creates a new pronunciation dictionary from provided rules.
func (r *PronunciationDictionaryService) NewFromRules(ctx context.Context, params PronunciationDictionaryNewFromRulesParams, opts ...option.RequestOption) (res *AddPronunciationDictionary, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/pronunciation-dictionaries/add-from-rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a PLS file with a pronunciation dictionary version rules
func (r *PronunciationDictionaryService) DownloadPls(ctx context.Context, versionID string, params PronunciationDictionaryDownloadPlsParams, opts ...option.RequestOption) (res *io.Reader, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if params.DictionaryID == "" {
		err = errors.New("missing required dictionary_id parameter")
		return
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return
	}
	path := fmt.Sprintf("v1/pronunciation-dictionaries/%s/%s/download", params.DictionaryID, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Remove rules from the pronunciation dictionary
func (r *PronunciationDictionaryService) RemoveRules(ctx context.Context, pronunciationDictionaryID string, params PronunciationDictionaryRemoveRulesParams, opts ...option.RequestOption) (res *PronunciationDictionaryRules, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if pronunciationDictionaryID == "" {
		err = errors.New("missing required pronunciation_dictionary_id parameter")
		return
	}
	path := fmt.Sprintf("v1/pronunciation-dictionaries/%s/remove-rules", pronunciationDictionaryID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AddPronunciationDictionary struct {
	// The ID of the created pronunciation dictionary.
	ID string `json:"id,required"`
	// The user ID of the creator of the pronunciation dictionary.
	CreatedBy string `json:"created_by,required"`
	// The creation time of the pronunciation dictionary in Unix timestamp.
	CreationTimeUnix int64 `json:"creation_time_unix,required"`
	// The name of the created pronunciation dictionary.
	Name string `json:"name,required"`
	// The permission on the resource of the pronunciation dictionary.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	PermissionOnResource AddPronunciationDictionaryPermissionOnResource `json:"permission_on_resource,required"`
	// The ID of the created pronunciation dictionary version.
	VersionID string `json:"version_id,required"`
	// The number of rules in the version of the pronunciation dictionary.
	VersionRulesNum int64 `json:"version_rules_num,required"`
	// The description of the pronunciation dictionary.
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedBy            respjson.Field
		CreationTimeUnix     respjson.Field
		Name                 respjson.Field
		PermissionOnResource respjson.Field
		VersionID            respjson.Field
		VersionRulesNum      respjson.Field
		Description          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AddPronunciationDictionary) RawJSON() string { return r.JSON.raw }
func (r *AddPronunciationDictionary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The permission on the resource of the pronunciation dictionary.
type AddPronunciationDictionaryPermissionOnResource string

const (
	AddPronunciationDictionaryPermissionOnResourceAdmin     AddPronunciationDictionaryPermissionOnResource = "admin"
	AddPronunciationDictionaryPermissionOnResourceEditor    AddPronunciationDictionaryPermissionOnResource = "editor"
	AddPronunciationDictionaryPermissionOnResourceCommenter AddPronunciationDictionaryPermissionOnResource = "commenter"
	AddPronunciationDictionaryPermissionOnResourceViewer    AddPronunciationDictionaryPermissionOnResource = "viewer"
)

type GetPronunciationDictionaryMetadata struct {
	// The ID of the pronunciation dictionary.
	ID string `json:"id,required"`
	// The user ID of the creator of the pronunciation dictionary.
	CreatedBy string `json:"created_by,required"`
	// The creation time of the pronunciation dictionary in Unix timestamp.
	CreationTimeUnix int64 `json:"creation_time_unix,required"`
	// The ID of the latest version of the pronunciation dictionary.
	LatestVersionID string `json:"latest_version_id,required"`
	// The number of rules in the latest version of the pronunciation dictionary.
	LatestVersionRulesNum int64 `json:"latest_version_rules_num,required"`
	// The name of the pronunciation dictionary.
	Name string `json:"name,required"`
	// The permission on the resource of the pronunciation dictionary.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	PermissionOnResource GetPronunciationDictionaryMetadataPermissionOnResource `json:"permission_on_resource,required"`
	// The archive time of the pronunciation dictionary in Unix timestamp.
	ArchivedTimeUnix int64 `json:"archived_time_unix,nullable"`
	// The description of the pronunciation dictionary.
	Description string `json:"description,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		CreatedBy             respjson.Field
		CreationTimeUnix      respjson.Field
		LatestVersionID       respjson.Field
		LatestVersionRulesNum respjson.Field
		Name                  respjson.Field
		PermissionOnResource  respjson.Field
		ArchivedTimeUnix      respjson.Field
		Description           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GetPronunciationDictionaryMetadata) RawJSON() string { return r.JSON.raw }
func (r *GetPronunciationDictionaryMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The permission on the resource of the pronunciation dictionary.
type GetPronunciationDictionaryMetadataPermissionOnResource string

const (
	GetPronunciationDictionaryMetadataPermissionOnResourceAdmin     GetPronunciationDictionaryMetadataPermissionOnResource = "admin"
	GetPronunciationDictionaryMetadataPermissionOnResourceEditor    GetPronunciationDictionaryMetadataPermissionOnResource = "editor"
	GetPronunciationDictionaryMetadataPermissionOnResourceCommenter GetPronunciationDictionaryMetadataPermissionOnResource = "commenter"
	GetPronunciationDictionaryMetadataPermissionOnResourceViewer    GetPronunciationDictionaryMetadataPermissionOnResource = "viewer"
)

// The properties Alias, StringToReplace, Type are required.
type PronunciationDictionaryAliasRuleParam struct {
	// The alias for the string to be replaced.
	Alias string `json:"alias,required"`
	// The string to replace. Must be a non-empty string.
	StringToReplace string `json:"string_to_replace,required"`
	// The type of the rule.
	//
	// This field can be elided, and will marshal its zero value as "alias".
	Type constant.Alias `json:"type,required"`
	paramObj
}

func (r PronunciationDictionaryAliasRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryAliasRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryAliasRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Alphabet, Phoneme, StringToReplace, Type are required.
type PronunciationDictionaryPhonemeRuleParam struct {
	// The alphabet to use with the phoneme rule.
	Alphabet string `json:"alphabet,required"`
	// The phoneme rule.
	Phoneme string `json:"phoneme,required"`
	// The string to replace. Must be a non-empty string.
	StringToReplace string `json:"string_to_replace,required"`
	// The type of the rule.
	//
	// This field can be elided, and will marshal its zero value as "phoneme".
	Type constant.Phoneme `json:"type,required"`
	paramObj
}

func (r PronunciationDictionaryPhonemeRuleParam) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryPhonemeRuleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryPhonemeRuleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PronunciationDictionaryRules struct {
	// The ID of the pronunciation dictionary.
	ID string `json:"id,required"`
	// The version ID of the pronunciation dictionary.
	VersionID string `json:"version_id,required"`
	// The number of rules in the version of the pronunciation dictionary.
	VersionRulesNum int64 `json:"version_rules_num,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		VersionID       respjson.Field
		VersionRulesNum respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictionaryRules) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictionaryRules) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PronunciationDictionaryListResponse struct {
	// Whether there are more pronunciation dictionaries to fetch.
	HasMore bool `json:"has_more,required"`
	// A list of pronunciation dictionaries and their metadata.
	PronunciationDictionaries []GetPronunciationDictionaryMetadata `json:"pronunciation_dictionaries,required"`
	// The next cursor to use for pagination.
	NextCursor string `json:"next_cursor,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore                   respjson.Field
		PronunciationDictionaries respjson.Field
		NextCursor                respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictionaryListResponse) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictionaryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PronunciationDictionaryGetParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type PronunciationDictionaryUpdateParams struct {
	// The name of the pronunciation dictionary, used for identification only.
	Archived param.Opt[bool] `json:"archived,omitzero"`
	// The name of the pronunciation dictionary, used for identification only.
	Name param.Opt[string] `json:"name,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r PronunciationDictionaryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PronunciationDictionaryListParams struct {
	// Used for fetching next page. Cursor is returned in the response.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Which direction to sort the voices in. 'ascending' or 'descending'.
	SortDirection param.Opt[string] `query:"sort_direction,omitzero" json:"-"`
	// How many pronunciation dictionaries to return at maximum. Can not exceed 100,
	// defaults to 30.
	PageSize param.Opt[int64] `query:"page_size,omitzero" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Which field to sort by, one of 'created_at_unix' or 'name'.
	//
	// Any of "creation_time_unix", "name".
	Sort PronunciationDictionaryListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PronunciationDictionaryListParams]'s query parameters as
// `url.Values`.
func (r PronunciationDictionaryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which field to sort by, one of 'created_at_unix' or 'name'.
type PronunciationDictionaryListParamsSort string

const (
	PronunciationDictionaryListParamsSortCreationTimeUnix PronunciationDictionaryListParamsSort = "creation_time_unix"
	PronunciationDictionaryListParamsSortName             PronunciationDictionaryListParamsSort = "name"
)

type PronunciationDictionaryAddRulesParams struct {
	// List of pronunciation rules. Rule can be either: an alias rule:
	// {'string_to_replace': 'a', 'type': 'alias', 'alias': 'b', } or a phoneme rule:
	// {'string_to_replace': 'a', 'type': 'phoneme', 'phoneme': 'b', 'alphabet': 'ipa'
	// }
	Rules []PronunciationDictionaryAddRulesParamsRuleUnion `json:"rules,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r PronunciationDictionaryAddRulesParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryAddRulesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryAddRulesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PronunciationDictionaryAddRulesParamsRuleUnion struct {
	OfPronunciationDictionaryAliasRuleRequestModel   *PronunciationDictionaryAliasRuleParam   `json:",omitzero,inline"`
	OfPronunciationDictionaryPhonemeRuleRequestModel *PronunciationDictionaryPhonemeRuleParam `json:",omitzero,inline"`
	paramUnion
}

func (u PronunciationDictionaryAddRulesParamsRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPronunciationDictionaryAliasRuleRequestModel, u.OfPronunciationDictionaryPhonemeRuleRequestModel)
}
func (u *PronunciationDictionaryAddRulesParamsRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PronunciationDictionaryAddRulesParamsRuleUnion) asAny() any {
	if !param.IsOmitted(u.OfPronunciationDictionaryAliasRuleRequestModel) {
		return u.OfPronunciationDictionaryAliasRuleRequestModel
	} else if !param.IsOmitted(u.OfPronunciationDictionaryPhonemeRuleRequestModel) {
		return u.OfPronunciationDictionaryPhonemeRuleRequestModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryAddRulesParamsRuleUnion) GetAlias() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return &vt.Alias
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryAddRulesParamsRuleUnion) GetAlphabet() *string {
	if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return &vt.Alphabet
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryAddRulesParamsRuleUnion) GetPhoneme() *string {
	if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return &vt.Phoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryAddRulesParamsRuleUnion) GetStringToReplace() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return (*string)(&vt.StringToReplace)
	} else if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return (*string)(&vt.StringToReplace)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryAddRulesParamsRuleUnion) GetType() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

type PronunciationDictionaryNewFromFileParams struct {
	// The name of the pronunciation dictionary, used for identification only.
	Name string `json:"name,required"`
	// A description of the pronunciation dictionary, used for identification only.
	Description param.Opt[string] `json:"description,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// A lexicon .pls file which we will use to initialize the project with.
	File io.Reader `json:"file,omitzero" format:"binary"`
	// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no
	// access.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	WorkspaceAccess PronunciationDictionaryNewFromFileParamsWorkspaceAccess `json:"workspace_access,omitzero"`
	paramObj
}

func (r PronunciationDictionaryNewFromFileParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no
// access.
type PronunciationDictionaryNewFromFileParamsWorkspaceAccess string

const (
	PronunciationDictionaryNewFromFileParamsWorkspaceAccessAdmin     PronunciationDictionaryNewFromFileParamsWorkspaceAccess = "admin"
	PronunciationDictionaryNewFromFileParamsWorkspaceAccessEditor    PronunciationDictionaryNewFromFileParamsWorkspaceAccess = "editor"
	PronunciationDictionaryNewFromFileParamsWorkspaceAccessCommenter PronunciationDictionaryNewFromFileParamsWorkspaceAccess = "commenter"
	PronunciationDictionaryNewFromFileParamsWorkspaceAccessViewer    PronunciationDictionaryNewFromFileParamsWorkspaceAccess = "viewer"
)

type PronunciationDictionaryNewFromRulesParams struct {
	// The name of the pronunciation dictionary, used for identification only.
	Name string `json:"name,required"`
	// List of pronunciation rules. Rule can be either: an alias rule:
	// {'string_to_replace': 'a', 'type': 'alias', 'alias': 'b', } or a phoneme rule:
	// {'string_to_replace': 'a', 'type': 'phoneme', 'phoneme': 'b', 'alphabet': 'ipa'
	// }
	Rules []PronunciationDictionaryNewFromRulesParamsRuleUnion `json:"rules,omitzero,required"`
	// A description of the pronunciation dictionary, used for identification only.
	Description param.Opt[string] `json:"description,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no
	// access.
	//
	// Any of "admin", "editor", "commenter", "viewer".
	WorkspaceAccess PronunciationDictionaryNewFromRulesParamsWorkspaceAccess `json:"workspace_access,omitzero"`
	paramObj
}

func (r PronunciationDictionaryNewFromRulesParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryNewFromRulesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryNewFromRulesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PronunciationDictionaryNewFromRulesParamsRuleUnion struct {
	OfPronunciationDictionaryAliasRuleRequestModel   *PronunciationDictionaryAliasRuleParam   `json:",omitzero,inline"`
	OfPronunciationDictionaryPhonemeRuleRequestModel *PronunciationDictionaryPhonemeRuleParam `json:",omitzero,inline"`
	paramUnion
}

func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPronunciationDictionaryAliasRuleRequestModel, u.OfPronunciationDictionaryPhonemeRuleRequestModel)
}
func (u *PronunciationDictionaryNewFromRulesParamsRuleUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PronunciationDictionaryNewFromRulesParamsRuleUnion) asAny() any {
	if !param.IsOmitted(u.OfPronunciationDictionaryAliasRuleRequestModel) {
		return u.OfPronunciationDictionaryAliasRuleRequestModel
	} else if !param.IsOmitted(u.OfPronunciationDictionaryPhonemeRuleRequestModel) {
		return u.OfPronunciationDictionaryPhonemeRuleRequestModel
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) GetAlias() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return &vt.Alias
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) GetAlphabet() *string {
	if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return &vt.Alphabet
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) GetPhoneme() *string {
	if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return &vt.Phoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) GetStringToReplace() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return (*string)(&vt.StringToReplace)
	} else if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return (*string)(&vt.StringToReplace)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictionaryNewFromRulesParamsRuleUnion) GetType() *string {
	if vt := u.OfPronunciationDictionaryAliasRuleRequestModel; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPronunciationDictionaryPhonemeRuleRequestModel; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Should be one of 'admin', 'editor' or 'viewer'. If not provided, defaults to no
// access.
type PronunciationDictionaryNewFromRulesParamsWorkspaceAccess string

const (
	PronunciationDictionaryNewFromRulesParamsWorkspaceAccessAdmin     PronunciationDictionaryNewFromRulesParamsWorkspaceAccess = "admin"
	PronunciationDictionaryNewFromRulesParamsWorkspaceAccessEditor    PronunciationDictionaryNewFromRulesParamsWorkspaceAccess = "editor"
	PronunciationDictionaryNewFromRulesParamsWorkspaceAccessCommenter PronunciationDictionaryNewFromRulesParamsWorkspaceAccess = "commenter"
	PronunciationDictionaryNewFromRulesParamsWorkspaceAccessViewer    PronunciationDictionaryNewFromRulesParamsWorkspaceAccess = "viewer"
)

type PronunciationDictionaryDownloadPlsParams struct {
	// The id of the pronunciation dictionary
	DictionaryID string `path:"dictionary_id,required" json:"-"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type PronunciationDictionaryRemoveRulesParams struct {
	// List of strings to remove from the pronunciation dictionary.
	RuleStrings []string `json:"rule_strings,omitzero,required"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

func (r PronunciationDictionaryRemoveRulesParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictionaryRemoveRulesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictionaryRemoveRulesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
