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
	"slices"

	"github.com/oh-tarnished/elevenlabs-go/internal/apiform"
	"github.com/oh-tarnished/elevenlabs-go/internal/apijson"
	"github.com/oh-tarnished/elevenlabs-go/internal/requestconfig"
	"github.com/oh-tarnished/elevenlabs-go/option"
	"github.com/oh-tarnished/elevenlabs-go/packages/param"
	"github.com/oh-tarnished/elevenlabs-go/packages/respjson"
)

// AudioNativeService contains methods and other services that help with
// interacting with the elevenlabs-personal API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAudioNativeService] method instead.
type AudioNativeService struct {
	Options []option.RequestOption
}

// NewAudioNativeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAudioNativeService(opts ...option.RequestOption) (r AudioNativeService) {
	r = AudioNativeService{}
	r.Options = opts
	return
}

// Creates Audio Native enabled project, optionally starts conversion and returns
// project ID and embeddable HTML snippet.
func (r *AudioNativeService) New(ctx context.Context, params AudioNativeNewParams, opts ...option.RequestOption) (res *AudioNativeNewResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "v1/audio-native"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get player settings for the specific project.
func (r *AudioNativeService) GetSettings(ctx context.Context, projectID string, query AudioNativeGetSettingsParams, opts ...option.RequestOption) (res *AudioNativeGetSettingsResponse, err error) {
	if !param.IsOmitted(query.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", query.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/audio-native/%s/settings", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates content for the specific AudioNative Project.
func (r *AudioNativeService) UpdateContent(ctx context.Context, projectID string, params AudioNativeUpdateContentParams, opts ...option.RequestOption) (res *AudioNativeUpdateContentResponse, err error) {
	if !param.IsOmitted(params.XiAPIKey) {
		opts = append(opts, option.WithHeader("xi-api-key", fmt.Sprintf("%s", params.XiAPIKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if projectID == "" {
		err = errors.New("missing required project_id parameter")
		return
	}
	path := fmt.Sprintf("v1/audio-native/%s/content", projectID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type AudioNativeNewResponse struct {
	// Whether the project is currently being converted.
	Converting bool `json:"converting,required"`
	// The HTML snippet to embed the Audio Native player.
	HTMLSnippet string `json:"html_snippet,required"`
	// The ID of the created Audio Native project.
	ProjectID string `json:"project_id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Converting  respjson.Field
		HTMLSnippet respjson.Field
		ProjectID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioNativeNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioNativeNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioNativeGetSettingsResponse struct {
	// Whether the project is enabled.
	Enabled bool `json:"enabled,required"`
	// The settings of the project.
	Settings AudioNativeGetSettingsResponseSettings `json:"settings,nullable"`
	// The ID of the latest snapshot of the project.
	SnapshotID string `json:"snapshot_id,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Settings    respjson.Field
		SnapshotID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioNativeGetSettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioNativeGetSettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The settings of the project.
type AudioNativeGetSettingsResponseSettings struct {
	// The author of the project.
	Author string `json:"author,required"`
	// The background color of the project.
	BackgroundColor string `json:"background_color,required"`
	// The image of the project.
	Image string `json:"image,required"`
	// The sessionization of the project. Specifies for how many minutes to persist the
	// session across page reloads.
	Sessionization int64 `json:"sessionization,required"`
	// Whether the project is small.
	Small bool `json:"small,required"`
	// The text color of the project.
	TextColor string `json:"text_color,required"`
	// The title of the project.
	Title string `json:"title,required"`
	// The path of the audio file.
	AudioPath string `json:"audio_path,nullable"`
	// The URL of the audio file.
	AudioURL string `json:"audio_url,nullable"`
	// Current state of the project
	//
	// Any of "processing", "ready".
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Author          respjson.Field
		BackgroundColor respjson.Field
		Image           respjson.Field
		Sessionization  respjson.Field
		Small           respjson.Field
		TextColor       respjson.Field
		Title           respjson.Field
		AudioPath       respjson.Field
		AudioURL        respjson.Field
		Status          respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioNativeGetSettingsResponseSettings) RawJSON() string { return r.JSON.raw }
func (r *AudioNativeGetSettingsResponseSettings) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioNativeUpdateContentResponse struct {
	// Whether the project is currently being converted.
	Converting bool `json:"converting,required"`
	// The HTML snippet to embed the Audio Native player.
	HTMLSnippet string `json:"html_snippet,required"`
	// The ID of the project.
	ProjectID string `json:"project_id,required"`
	// Whether the project is currently being published.
	Publishing bool `json:"publishing,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Converting  respjson.Field
		HTMLSnippet respjson.Field
		ProjectID   respjson.Field
		Publishing  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AudioNativeUpdateContentResponse) RawJSON() string { return r.JSON.raw }
func (r *AudioNativeUpdateContentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AudioNativeNewParams struct {
	// Project name.
	Name string `json:"name,required"`
	// Author used in the player and inserted at the start of the uploaded article. If
	// not provided, the default author set in the Player settings is used.
	Author param.Opt[string] `json:"author,omitzero"`
	// Background color used in the player. If not provided, default background color
	// set in the Player settings is used.
	BackgroundColor param.Opt[string] `json:"background_color,omitzero"`
	// (Deprecated) Image URL used in the player. If not provided, default image set in
	// the Player settings is used.
	Image param.Opt[string] `json:"image,omitzero"`
	// TTS Model ID used in the player. If not provided, default model ID set in the
	// Player settings is used.
	ModelID param.Opt[string] `json:"model_id,omitzero"`
	// Text color used in the player. If not provided, default text color set in the
	// Player settings is used.
	TextColor param.Opt[string] `json:"text_color,omitzero"`
	// Title used in the player and inserted at the top of the uploaded article. If not
	// provided, the default title set in the Player settings is used.
	Title param.Opt[string] `json:"title,omitzero"`
	// Voice ID used to voice the content. If not provided, default voice ID set in the
	// Player settings is used.
	VoiceID param.Opt[string] `json:"voice_id,omitzero"`
	// Whether to auto convert the project to audio or not.
	AutoConvert param.Opt[bool] `json:"auto_convert,omitzero"`
	// (Deprecated) Specifies for how many minutes to persist the session across page
	// reloads. If not provided, default sessionization set in the Player settings is
	// used.
	Sessionization param.Opt[int64] `json:"sessionization,omitzero"`
	// (Deprecated) Whether to use small player or not. If not provided, default value
	// set in the Player settings is used.
	Small param.Opt[bool] `json:"small,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
	//
	//	When set to 'auto', the system will automatically decide whether to apply text normalization
	//	(e.g., spelling out numbers). With 'on', text normalization will always be applied, while
	//	with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
	//
	// Any of "auto", "on", "off", "apply_english".
	ApplyTextNormalization AudioNativeNewParamsApplyTextNormalization `json:"apply_text_normalization,omitzero"`
	// Either txt or HTML input file containing the article content. HTML should be
	// formatted as follows '&lt;html&gt;&lt;body&gt;&lt;div&gt;&lt;p&gt;Your
	// content&lt;/p&gt;&lt;h3&gt;More of your content&lt;/h3&gt;&lt;p&gt;Some more of
	// your content&lt;/p&gt;&lt;/div&gt;&lt;/body&gt;&lt;/html&gt;'
	File io.Reader `json:"file,omitzero" format:"binary"`
	// A list of pronunciation dictionary locators (pronunciation_dictionary_id,
	// version_id) encoded as a list of JSON strings for pronunciation dictionaries to
	// be applied to the text. A list of json encoded strings is required as adding
	// projects may occur through formData as opposed to jsonBody. To specify multiple
	// dictionaries use multiple --form lines in your curl, such as --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"Vmd4Zor6fplcA7WrINey\",\"version_id\":\"hRPaxjlTdR7wFMhV4w0b\"}"'
	// --form
	// 'pronunciation_dictionary_locators="{\"pronunciation_dictionary_id\":\"JzWtcGQMJ6bnlWwyMo7e\",\"version_id\":\"lbmwxiLu4q6txYxgdZqn\"}"'.
	PronunciationDictionaryLocators []string `json:"pronunciation_dictionary_locators,omitzero"`
	paramObj
}

func (r AudioNativeNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
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

// This parameter controls text normalization with four modes: 'auto', 'on', 'apply_english' and 'off'.
//
//	When set to 'auto', the system will automatically decide whether to apply text normalization
//	(e.g., spelling out numbers). With 'on', text normalization will always be applied, while
//	with 'off', it will be skipped. 'apply_english' is the same as 'on' but will assume that text is in English.
type AudioNativeNewParamsApplyTextNormalization string

const (
	AudioNativeNewParamsApplyTextNormalizationAuto         AudioNativeNewParamsApplyTextNormalization = "auto"
	AudioNativeNewParamsApplyTextNormalizationOn           AudioNativeNewParamsApplyTextNormalization = "on"
	AudioNativeNewParamsApplyTextNormalizationOff          AudioNativeNewParamsApplyTextNormalization = "off"
	AudioNativeNewParamsApplyTextNormalizationApplyEnglish AudioNativeNewParamsApplyTextNormalization = "apply_english"
)

type AudioNativeGetSettingsParams struct {
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	paramObj
}

type AudioNativeUpdateContentParams struct {
	// Whether to auto convert the project to audio or not.
	AutoConvert param.Opt[bool] `json:"auto_convert,omitzero"`
	// Whether to auto publish the new project snapshot after it's converted.
	AutoPublish param.Opt[bool] `json:"auto_publish,omitzero"`
	// Your API key. This is required by most endpoints to access our API
	// programmatically. You can view your xi-api-key using the 'Profile' tab on the
	// website.
	XiAPIKey param.Opt[string] `header:"xi-api-key,omitzero" json:"-"`
	// Either txt or HTML input file containing the article content. HTML should be
	// formatted as follows '&lt;html&gt;&lt;body&gt;&lt;div&gt;&lt;p&gt;Your
	// content&lt;/p&gt;&lt;h5&gt;More of your content&lt;/h5&gt;&lt;p&gt;Some more of
	// your content&lt;/p&gt;&lt;/div&gt;&lt;/body&gt;&lt;/html&gt;'
	File io.Reader `json:"file,omitzero" format:"binary"`
	paramObj
}

func (r AudioNativeUpdateContentParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
