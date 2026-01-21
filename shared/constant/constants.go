// File generated from our OpenAPI spec. See CONTRIBUTING.md for details.

package constant

import (
	shimjson "github.com/oh-tarnished/elevenlabs-go/internal/encoding/json"
)

type Constant[T any] interface {
	Default() T
}

// ValueOf gives the default value of a constant from its type. It's helpful when
// constructing constants as variants in a one-of. Note that empty structs are
// marshalled by default. Usage: constant.ValueOf[constant.Foo]()
func ValueOf[T Constant[T]]() T {
	var t T
	return t.Default()
}

type Other string                 // Always "_other"
type Alias string                 // Always "alias"
type All string                   // Always "all"
type AndOperator string           // Always "and_operator"
type Anything string              // Always "anything"
type APIIntegrationWebhook string // Always "api_integration_webhook"
type BooleanLiteral string        // Always "boolean_literal"
type Bulletin string              // Always "bulletin"
type Conversation string          // Always "conversation"
type Docx string                  // Always "docx"
type Dynamic string               // Always "dynamic"
type DynamicVariable string       // Always "dynamic_variable"
type End string                   // Always "end"
type EqOperator string            // Always "eq_operator"
type Exact string                 // Always "exact"
type Expression string            // Always "expression"
type Failure string               // Always "failure"
type File string                  // Always "file"
type Folder string                // Always "folder"
type GtOperator string            // Always "gt_operator"
type GteOperator string           // Always "gte_operator"
type Hmac string                  // Always "hmac"
type HTML string                  // Always "html"
type Llm string                   // Always "llm"
type LtOperator string            // Always "lt_operator"
type LteOperator string           // Always "lte_operator"
type NeqOperator string           // Always "neq_operator"
type New string                   // Always "new"
type NumberLiteral string         // Always "number_literal"
type OrOperator string            // Always "or_operator"
type OverrideAgent string         // Always "override_agent"
type Pdf string                   // Always "pdf"
type PhoneNumber string           // Always "phone_number"
type Phoneme string               // Always "phoneme"
type Regex string                 // Always "regex"
type Result string                // Always "result"
type SegmentedJson string         // Always "segmented_json"
type SipTrunking string           // Always "sip_trunking"
type Srt string                   // Always "srt"
type StandaloneAgent string       // Always "standalone_agent"
type Start string                 // Always "start"
type Stored string                // Always "stored"
type StringLiteral string         // Always "string_literal"
type Success string               // Always "success"
type System string                // Always "system"
type Text string                  // Always "text"
type Tool string                  // Always "tool"
type TtsNode string               // Always "tts_node"
type Twilio string                // Always "twilio"
type Txt string                   // Always "txt"
type Unconditional string         // Always "unconditional"
type Update string                // Always "update"
type URL string                   // Always "url"
type Workflow string              // Always "workflow"

func (c Other) Default() Other                                 { return "_other" }
func (c Alias) Default() Alias                                 { return "alias" }
func (c All) Default() All                                     { return "all" }
func (c AndOperator) Default() AndOperator                     { return "and_operator" }
func (c Anything) Default() Anything                           { return "anything" }
func (c APIIntegrationWebhook) Default() APIIntegrationWebhook { return "api_integration_webhook" }
func (c BooleanLiteral) Default() BooleanLiteral               { return "boolean_literal" }
func (c Bulletin) Default() Bulletin                           { return "bulletin" }
func (c Conversation) Default() Conversation                   { return "conversation" }
func (c Docx) Default() Docx                                   { return "docx" }
func (c Dynamic) Default() Dynamic                             { return "dynamic" }
func (c DynamicVariable) Default() DynamicVariable             { return "dynamic_variable" }
func (c End) Default() End                                     { return "end" }
func (c EqOperator) Default() EqOperator                       { return "eq_operator" }
func (c Exact) Default() Exact                                 { return "exact" }
func (c Expression) Default() Expression                       { return "expression" }
func (c Failure) Default() Failure                             { return "failure" }
func (c File) Default() File                                   { return "file" }
func (c Folder) Default() Folder                               { return "folder" }
func (c GtOperator) Default() GtOperator                       { return "gt_operator" }
func (c GteOperator) Default() GteOperator                     { return "gte_operator" }
func (c Hmac) Default() Hmac                                   { return "hmac" }
func (c HTML) Default() HTML                                   { return "html" }
func (c Llm) Default() Llm                                     { return "llm" }
func (c LtOperator) Default() LtOperator                       { return "lt_operator" }
func (c LteOperator) Default() LteOperator                     { return "lte_operator" }
func (c NeqOperator) Default() NeqOperator                     { return "neq_operator" }
func (c New) Default() New                                     { return "new" }
func (c NumberLiteral) Default() NumberLiteral                 { return "number_literal" }
func (c OrOperator) Default() OrOperator                       { return "or_operator" }
func (c OverrideAgent) Default() OverrideAgent                 { return "override_agent" }
func (c Pdf) Default() Pdf                                     { return "pdf" }
func (c PhoneNumber) Default() PhoneNumber                     { return "phone_number" }
func (c Phoneme) Default() Phoneme                             { return "phoneme" }
func (c Regex) Default() Regex                                 { return "regex" }
func (c Result) Default() Result                               { return "result" }
func (c SegmentedJson) Default() SegmentedJson                 { return "segmented_json" }
func (c SipTrunking) Default() SipTrunking                     { return "sip_trunking" }
func (c Srt) Default() Srt                                     { return "srt" }
func (c StandaloneAgent) Default() StandaloneAgent             { return "standalone_agent" }
func (c Start) Default() Start                                 { return "start" }
func (c Stored) Default() Stored                               { return "stored" }
func (c StringLiteral) Default() StringLiteral                 { return "string_literal" }
func (c Success) Default() Success                             { return "success" }
func (c System) Default() System                               { return "system" }
func (c Text) Default() Text                                   { return "text" }
func (c Tool) Default() Tool                                   { return "tool" }
func (c TtsNode) Default() TtsNode                             { return "tts_node" }
func (c Twilio) Default() Twilio                               { return "twilio" }
func (c Txt) Default() Txt                                     { return "txt" }
func (c Unconditional) Default() Unconditional                 { return "unconditional" }
func (c Update) Default() Update                               { return "update" }
func (c URL) Default() URL                                     { return "url" }
func (c Workflow) Default() Workflow                           { return "workflow" }

func (c Other) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Alias) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c All) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c AndOperator) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Anything) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c APIIntegrationWebhook) MarshalJSON() ([]byte, error) { return marshalString(c) }
func (c BooleanLiteral) MarshalJSON() ([]byte, error)        { return marshalString(c) }
func (c Bulletin) MarshalJSON() ([]byte, error)              { return marshalString(c) }
func (c Conversation) MarshalJSON() ([]byte, error)          { return marshalString(c) }
func (c Docx) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Dynamic) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c DynamicVariable) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c End) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c EqOperator) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Exact) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Expression) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c Failure) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c File) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Folder) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c GtOperator) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c GteOperator) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Hmac) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c HTML) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Llm) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c LtOperator) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c LteOperator) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c NeqOperator) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c New) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c NumberLiteral) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c OrOperator) MarshalJSON() ([]byte, error)            { return marshalString(c) }
func (c OverrideAgent) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Pdf) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c PhoneNumber) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Phoneme) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Regex) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Result) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c SegmentedJson) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c SipTrunking) MarshalJSON() ([]byte, error)           { return marshalString(c) }
func (c Srt) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c StandaloneAgent) MarshalJSON() ([]byte, error)       { return marshalString(c) }
func (c Start) MarshalJSON() ([]byte, error)                 { return marshalString(c) }
func (c Stored) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c StringLiteral) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Success) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c System) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Text) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c Tool) MarshalJSON() ([]byte, error)                  { return marshalString(c) }
func (c TtsNode) MarshalJSON() ([]byte, error)               { return marshalString(c) }
func (c Twilio) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c Txt) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Unconditional) MarshalJSON() ([]byte, error)         { return marshalString(c) }
func (c Update) MarshalJSON() ([]byte, error)                { return marshalString(c) }
func (c URL) MarshalJSON() ([]byte, error)                   { return marshalString(c) }
func (c Workflow) MarshalJSON() ([]byte, error)              { return marshalString(c) }

type constant[T any] interface {
	Constant[T]
	*T
}

func marshalString[T ~string, PT constant[T]](v T) ([]byte, error) {
	var zero T
	if v == zero {
		v = PT(&v).Default()
	}
	return shimjson.Marshal(string(v))
}
