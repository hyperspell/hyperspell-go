// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"
	"time"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
type Blob struct {
	Data     string `json:"data" api:"required"`
	Mimetype string `json:"mimetype" api:"required"`
	ID       string `json:"id"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "blob".
	Type BlobType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Blob) RawJSON() string { return r.JSON.raw }
func (r *Blob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Blob) ImplCalloutChildUnion()   {}
func (Blob) ImplChunkChildUnion()     {}
func (Blob) ImplCompanyChildUnion()   {}
func (Blob) ImplDealChildUnion()      {}
func (Blob) ImplDocumentChildUnion()  {}
func (Blob) ImplEquationChildUnion()  {}
func (Blob) ImplEventChildUnion()     {}
func (Blob) ImplFileChildUnion()      {}
func (Blob) ImplFootnoteChildUnion()  {}
func (Blob) ImplHeadingChildUnion()   {}
func (Blob) ImplListItemChildUnion()  {}
func (Blob) ImplMessageChildUnion()   {}
func (Blob) ImplParagraphChildUnion() {}
func (Blob) ImplPersonChildUnion()    {}
func (Blob) ImplQuoteChildUnion()     {}
func (Blob) ImplTableCellChildUnion() {}
func (Blob) ImplTaskChildUnion()      {}
func (Blob) ImplToDoChildUnion()      {}
func (Blob) ImplWebsiteChildUnion()   {}

type BlobType string

const (
	BlobTypeBlob BlobType = "blob"
)

type Callout struct {
	ID       string              `json:"id"`
	Children []CalloutChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "callout".
	Type CalloutType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Callout) RawJSON() string { return r.JSON.raw }
func (r *Callout) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Callout) ImplCalloutChildUnion()   {}
func (Callout) ImplChunkChildUnion()     {}
func (Callout) ImplCompanyChildUnion()   {}
func (Callout) ImplDealChildUnion()      {}
func (Callout) ImplDocumentChildUnion()  {}
func (Callout) ImplEquationChildUnion()  {}
func (Callout) ImplEventChildUnion()     {}
func (Callout) ImplFileChildUnion()      {}
func (Callout) ImplFootnoteChildUnion()  {}
func (Callout) ImplHeadingChildUnion()   {}
func (Callout) ImplListItemChildUnion()  {}
func (Callout) ImplMessageChildUnion()   {}
func (Callout) ImplParagraphChildUnion() {}
func (Callout) ImplPersonChildUnion()    {}
func (Callout) ImplQuoteChildUnion()     {}
func (Callout) ImplTableCellChildUnion() {}
func (Callout) ImplTaskChildUnion()      {}
func (Callout) ImplToDoChildUnion()      {}
func (Callout) ImplWebsiteChildUnion()   {}

// CalloutChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [CalloutChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CalloutChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children CalloutChildUnionChildren `json:"children"`
	Text     string                    `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyCalloutChild is implemented by each variant of [CalloutChildUnion] to add
// type safety for the return type of [CalloutChildUnion.AsAny]
type anyCalloutChild interface {
	ImplCalloutChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := CalloutChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CalloutChildUnion) AsAny() anyCalloutChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u CalloutChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CalloutChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CalloutChildUnion) RawJSON() string { return u.JSON.raw }

func (r *CalloutChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CalloutChildUnionChildren is an implicit subunion of [CalloutChildUnion].
// CalloutChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [CalloutChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type CalloutChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *CalloutChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalloutType string

const (
	CalloutTypeCallout CalloutType = "callout"
)

type Chunk struct {
	ID       string            `json:"id"`
	Children []ChunkChildUnion `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "chunk".
	Type ChunkType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Chunk) RawJSON() string { return r.JSON.raw }
func (r *Chunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Chunk) ImplCalloutChildUnion()   {}
func (Chunk) ImplChunkChildUnion()     {}
func (Chunk) ImplCompanyChildUnion()   {}
func (Chunk) ImplDealChildUnion()      {}
func (Chunk) ImplDocumentChildUnion()  {}
func (Chunk) ImplEquationChildUnion()  {}
func (Chunk) ImplEventChildUnion()     {}
func (Chunk) ImplFileChildUnion()      {}
func (Chunk) ImplFootnoteChildUnion()  {}
func (Chunk) ImplHeadingChildUnion()   {}
func (Chunk) ImplListItemChildUnion()  {}
func (Chunk) ImplMessageChildUnion()   {}
func (Chunk) ImplParagraphChildUnion() {}
func (Chunk) ImplPersonChildUnion()    {}
func (Chunk) ImplQuoteChildUnion()     {}
func (Chunk) ImplTableCellChildUnion() {}
func (Chunk) ImplTaskChildUnion()      {}
func (Chunk) ImplToDoChildUnion()      {}
func (Chunk) ImplWebsiteChildUnion()   {}

// ChunkChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [ChunkChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ChunkChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children ChunkChildUnionChildren `json:"children"`
	Text     string                  `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyChunkChild is implemented by each variant of [ChunkChildUnion] to add type
// safety for the return type of [ChunkChildUnion.AsAny]
type anyChunkChild interface {
	ImplChunkChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ChunkChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ChunkChildUnion) AsAny() anyChunkChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u ChunkChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ChunkChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ChunkChildUnion) RawJSON() string { return u.JSON.raw }

func (r *ChunkChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ChunkChildUnionChildren is an implicit subunion of [ChunkChildUnion].
// ChunkChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ChunkChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ChunkChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ChunkChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChunkType string

const (
	ChunkTypeChunk ChunkType = "chunk"
)

type Code struct {
	Text     string `json:"text" api:"required"`
	ID       string `json:"id"`
	Language string `json:"language" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "code".
	Type CodeType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ID          respjson.Field
		Language    respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Code) RawJSON() string { return r.JSON.raw }
func (r *Code) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Code) ImplCalloutChildUnion()   {}
func (Code) ImplChunkChildUnion()     {}
func (Code) ImplCompanyChildUnion()   {}
func (Code) ImplDealChildUnion()      {}
func (Code) ImplDocumentChildUnion()  {}
func (Code) ImplEquationChildUnion()  {}
func (Code) ImplEventChildUnion()     {}
func (Code) ImplFileChildUnion()      {}
func (Code) ImplFootnoteChildUnion()  {}
func (Code) ImplHeadingChildUnion()   {}
func (Code) ImplListItemChildUnion()  {}
func (Code) ImplMessageChildUnion()   {}
func (Code) ImplParagraphChildUnion() {}
func (Code) ImplPersonChildUnion()    {}
func (Code) ImplQuoteChildUnion()     {}
func (Code) ImplTableCellChildUnion() {}
func (Code) ImplTaskChildUnion()      {}
func (Code) ImplToDoChildUnion()      {}
func (Code) ImplWebsiteChildUnion()   {}

type CodeType string

const (
	CodeTypeCode CodeType = "code"
)

type Comment struct {
	Text      string    `json:"text" api:"required"`
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "comment".
	Type CommentType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Comment) RawJSON() string { return r.JSON.raw }
func (r *Comment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Comment) ImplCalloutChildUnion()   {}
func (Comment) ImplChunkChildUnion()     {}
func (Comment) ImplCompanyChildUnion()   {}
func (Comment) ImplDealChildUnion()      {}
func (Comment) ImplDocumentChildUnion()  {}
func (Comment) ImplEquationChildUnion()  {}
func (Comment) ImplEventChildUnion()     {}
func (Comment) ImplFileChildUnion()      {}
func (Comment) ImplFootnoteChildUnion()  {}
func (Comment) ImplHeadingChildUnion()   {}
func (Comment) ImplListItemChildUnion()  {}
func (Comment) ImplMessageChildUnion()   {}
func (Comment) ImplParagraphChildUnion() {}
func (Comment) ImplPersonChildUnion()    {}
func (Comment) ImplQuoteChildUnion()     {}
func (Comment) ImplTableCellChildUnion() {}
func (Comment) ImplTaskChildUnion()      {}
func (Comment) ImplToDoChildUnion()      {}
func (Comment) ImplWebsiteChildUnion()   {}

type CommentType string

const (
	CommentTypeComment CommentType = "comment"
)

// A CRM company/account record (ENG-2476/D10).
type Company struct {
	ID          string              `json:"id"`
	Address     string              `json:"address" api:"nullable"`
	Children    []CompanyChildUnion `json:"children"`
	ContactIDs  []string            `json:"contact_ids" api:"nullable"`
	DealIDs     []string            `json:"deal_ids" api:"nullable"`
	Description string              `json:"description" api:"nullable"`
	Emails      []string            `json:"emails" api:"nullable"`
	Employees   int64               `json:"employees" api:"nullable"`
	ImageURL    string              `json:"image_url" api:"nullable"`
	Industry    string              `json:"industry" api:"nullable"`
	IsActive    bool                `json:"is_active" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata     Metadata `json:"metadata" api:"nullable"`
	Name         string   `json:"name" api:"nullable"`
	PhoneNumbers []string `json:"phone_numbers" api:"nullable"`
	Tags         []string `json:"tags" api:"nullable"`
	Text         string   `json:"text" api:"nullable"`
	Timezone     string   `json:"timezone" api:"nullable"`
	// Any of "company".
	Type     CompanyType `json:"type"`
	Websites []string    `json:"websites" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Address      respjson.Field
		Children     respjson.Field
		ContactIDs   respjson.Field
		DealIDs      respjson.Field
		Description  respjson.Field
		Emails       respjson.Field
		Employees    respjson.Field
		ImageURL     respjson.Field
		Industry     respjson.Field
		IsActive     respjson.Field
		Metadata     respjson.Field
		Name         respjson.Field
		PhoneNumbers respjson.Field
		Tags         respjson.Field
		Text         respjson.Field
		Timezone     respjson.Field
		Type         respjson.Field
		Websites     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Company) RawJSON() string { return r.JSON.raw }
func (r *Company) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Company) ImplScoredDocumentResponseDocumentUnion() {}
func (Company) ImplMemoryListResponseDocumentUnion()     {}
func (Company) ImplMemoryGetResponseDocumentUnion()      {}

// CompanyChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [CompanyChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CompanyChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children CompanyChildUnionChildren `json:"children"`
	Text     string                    `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyCompanyChild is implemented by each variant of [CompanyChildUnion] to add
// type safety for the return type of [CompanyChildUnion.AsAny]
type anyCompanyChild interface {
	ImplCompanyChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := CompanyChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u CompanyChildUnion) AsAny() anyCompanyChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u CompanyChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CompanyChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CompanyChildUnion) RawJSON() string { return u.JSON.raw }

func (r *CompanyChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CompanyChildUnionChildren is an implicit subunion of [CompanyChildUnion].
// CompanyChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [CompanyChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type CompanyChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *CompanyChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CompanyType string

const (
	CompanyTypeCompany CompanyType = "company"
)

type Conversation struct {
	ID       string    `json:"id"`
	Channel  string    `json:"channel" api:"nullable"`
	Children []Message `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "conversation".
	Type ConversationType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Channel     respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Conversation) RawJSON() string { return r.JSON.raw }
func (r *Conversation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Conversation) ImplScoredDocumentResponseDocumentUnion() {}
func (Conversation) ImplMemoryListResponseDocumentUnion()     {}
func (Conversation) ImplMemoryGetResponseDocumentUnion()      {}

type ConversationType string

const (
	ConversationTypeConversation ConversationType = "conversation"
)

// A CRM deal/opportunity record (ENG-2476/D10).
type Deal struct {
	ID         string           `json:"id"`
	Amount     float64          `json:"amount" api:"nullable"`
	Children   []DealChildUnion `json:"children"`
	ClosedAt   time.Time        `json:"closed_at" api:"nullable" format:"date-time"`
	CompanyIDs []string         `json:"company_ids" api:"nullable"`
	ContactIDs []string         `json:"contact_ids" api:"nullable"`
	Currency   string           `json:"currency" api:"nullable"`
	DealSource string           `json:"deal_source" api:"nullable"`
	LostReason string           `json:"lost_reason" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata    Metadata `json:"metadata" api:"nullable"`
	Name        string   `json:"name" api:"nullable"`
	Pipeline    string   `json:"pipeline" api:"nullable"`
	Probability float64  `json:"probability" api:"nullable"`
	Stage       string   `json:"stage" api:"nullable"`
	Tags        []string `json:"tags" api:"nullable"`
	Text        string   `json:"text" api:"nullable"`
	// Any of "deal".
	Type      DealType `json:"type"`
	WonReason string   `json:"won_reason" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Amount      respjson.Field
		Children    respjson.Field
		ClosedAt    respjson.Field
		CompanyIDs  respjson.Field
		ContactIDs  respjson.Field
		Currency    respjson.Field
		DealSource  respjson.Field
		LostReason  respjson.Field
		Metadata    respjson.Field
		Name        respjson.Field
		Pipeline    respjson.Field
		Probability respjson.Field
		Stage       respjson.Field
		Tags        respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		WonReason   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Deal) RawJSON() string { return r.JSON.raw }
func (r *Deal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Deal) ImplScoredDocumentResponseDocumentUnion() {}
func (Deal) ImplMemoryListResponseDocumentUnion()     {}
func (Deal) ImplMemoryGetResponseDocumentUnion()      {}

// DealChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [DealChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DealChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children DealChildUnionChildren `json:"children"`
	Text     string                 `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyDealChild is implemented by each variant of [DealChildUnion] to add type
// safety for the return type of [DealChildUnion.AsAny]
type anyDealChild interface {
	ImplDealChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := DealChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u DealChildUnion) AsAny() anyDealChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u DealChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DealChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DealChildUnion) RawJSON() string { return u.JSON.raw }

func (r *DealChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DealChildUnionChildren is an implicit subunion of [DealChildUnion].
// DealChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [DealChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type DealChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *DealChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DealType string

const (
	DealTypeDeal DealType = "deal"
)

type Divider struct {
	ID string `json:"id"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "divider".
	Type DividerType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Divider) RawJSON() string { return r.JSON.raw }
func (r *Divider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Divider) ImplCalloutChildUnion()   {}
func (Divider) ImplChunkChildUnion()     {}
func (Divider) ImplCompanyChildUnion()   {}
func (Divider) ImplDealChildUnion()      {}
func (Divider) ImplDocumentChildUnion()  {}
func (Divider) ImplEquationChildUnion()  {}
func (Divider) ImplEventChildUnion()     {}
func (Divider) ImplFileChildUnion()      {}
func (Divider) ImplFootnoteChildUnion()  {}
func (Divider) ImplHeadingChildUnion()   {}
func (Divider) ImplListItemChildUnion()  {}
func (Divider) ImplMessageChildUnion()   {}
func (Divider) ImplParagraphChildUnion() {}
func (Divider) ImplPersonChildUnion()    {}
func (Divider) ImplQuoteChildUnion()     {}
func (Divider) ImplTableCellChildUnion() {}
func (Divider) ImplTaskChildUnion()      {}
func (Divider) ImplToDoChildUnion()      {}
func (Divider) ImplWebsiteChildUnion()   {}

type DividerType string

const (
	DividerTypeDivider DividerType = "divider"
)

type Document struct {
	ID       string               `json:"id"`
	Children []DocumentChildUnion `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "document".
	Type DocumentType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Document) RawJSON() string { return r.JSON.raw }
func (r *Document) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Document) ImplScoredDocumentResponseDocumentUnion() {}
func (Document) ImplMemoryListResponseDocumentUnion()     {}
func (Document) ImplMemoryGetResponseDocumentUnion()      {}

// DocumentChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [DocumentChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type DocumentChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children DocumentChildUnionChildren `json:"children"`
	Text     string                     `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyDocumentChild is implemented by each variant of [DocumentChildUnion] to add
// type safety for the return type of [DocumentChildUnion.AsAny]
type anyDocumentChild interface {
	ImplDocumentChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := DocumentChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u DocumentChildUnion) AsAny() anyDocumentChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u DocumentChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u DocumentChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u DocumentChildUnion) RawJSON() string { return u.JSON.raw }

func (r *DocumentChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// DocumentChildUnionChildren is an implicit subunion of [DocumentChildUnion].
// DocumentChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [DocumentChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type DocumentChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *DocumentChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DocumentType string

const (
	DocumentTypeDocument DocumentType = "document"
)

type Equation struct {
	ID       string               `json:"id"`
	Children []EquationChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "equation".
	Type EquationType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Equation) RawJSON() string { return r.JSON.raw }
func (r *Equation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Equation) ImplCalloutChildUnion()   {}
func (Equation) ImplChunkChildUnion()     {}
func (Equation) ImplCompanyChildUnion()   {}
func (Equation) ImplDealChildUnion()      {}
func (Equation) ImplDocumentChildUnion()  {}
func (Equation) ImplEquationChildUnion()  {}
func (Equation) ImplEventChildUnion()     {}
func (Equation) ImplFileChildUnion()      {}
func (Equation) ImplFootnoteChildUnion()  {}
func (Equation) ImplHeadingChildUnion()   {}
func (Equation) ImplListItemChildUnion()  {}
func (Equation) ImplMessageChildUnion()   {}
func (Equation) ImplParagraphChildUnion() {}
func (Equation) ImplPersonChildUnion()    {}
func (Equation) ImplQuoteChildUnion()     {}
func (Equation) ImplTableCellChildUnion() {}
func (Equation) ImplTaskChildUnion()      {}
func (Equation) ImplToDoChildUnion()      {}
func (Equation) ImplWebsiteChildUnion()   {}

// EquationChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [EquationChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EquationChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children EquationChildUnionChildren `json:"children"`
	Text     string                     `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyEquationChild is implemented by each variant of [EquationChildUnion] to add
// type safety for the return type of [EquationChildUnion.AsAny]
type anyEquationChild interface {
	ImplEquationChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := EquationChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EquationChildUnion) AsAny() anyEquationChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u EquationChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EquationChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EquationChildUnion) RawJSON() string { return u.JSON.raw }

func (r *EquationChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EquationChildUnionChildren is an implicit subunion of [EquationChildUnion].
// EquationChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [EquationChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type EquationChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *EquationChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EquationType string

const (
	EquationTypeEquation EquationType = "equation"
)

type Event struct {
	ID         string            `json:"id"`
	Attendees  []Person          `json:"attendees"`
	Children   []EventChildUnion `json:"children"`
	EndAt      time.Time         `json:"end_at" api:"nullable" format:"date-time"`
	Location   string            `json:"location" api:"nullable"`
	MeetingURL string            `json:"meeting_url" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata  `json:"metadata" api:"nullable"`
	StartAt  time.Time `json:"start_at" api:"nullable" format:"date-time"`
	Text     string    `json:"text" api:"nullable"`
	Title    string    `json:"title" api:"nullable"`
	// Any of "event".
	Type EventType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Attendees   respjson.Field
		Children    respjson.Field
		EndAt       respjson.Field
		Location    respjson.Field
		MeetingURL  respjson.Field
		Metadata    respjson.Field
		StartAt     respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Event) RawJSON() string { return r.JSON.raw }
func (r *Event) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Event) ImplScoredDocumentResponseDocumentUnion() {}
func (Event) ImplMemoryListResponseDocumentUnion()     {}
func (Event) ImplMemoryGetResponseDocumentUnion()      {}

// EventChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [EventChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type EventChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children EventChildUnionChildren `json:"children"`
	Text     string                  `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyEventChild is implemented by each variant of [EventChildUnion] to add type
// safety for the return type of [EventChildUnion.AsAny]
type anyEventChild interface {
	ImplEventChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := EventChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u EventChildUnion) AsAny() anyEventChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u EventChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u EventChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u EventChildUnion) RawJSON() string { return u.JSON.raw }

func (r *EventChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// EventChildUnionChildren is an implicit subunion of [EventChildUnion].
// EventChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [EventChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type EventChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *EventChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EventType string

const (
	EventTypeEvent EventType = "event"
)

type File struct {
	ContentType string           `json:"content_type" api:"required"`
	Filename    string           `json:"filename" api:"required"`
	ID          string           `json:"id"`
	Children    []FileChildUnion `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Path     []string `json:"path" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "file".
	Type FileType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContentType respjson.Field
		Filename    respjson.Field
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Path        respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (File) ImplScoredDocumentResponseDocumentUnion() {}
func (File) ImplMemoryListResponseDocumentUnion()     {}
func (File) ImplMemoryGetResponseDocumentUnion()      {}

// FileChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [FileChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FileChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children FileChildUnionChildren `json:"children"`
	Text     string                 `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyFileChild is implemented by each variant of [FileChildUnion] to add type
// safety for the return type of [FileChildUnion.AsAny]
type anyFileChild interface {
	ImplFileChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := FileChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FileChildUnion) AsAny() anyFileChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u FileChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FileChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FileChildUnion) RawJSON() string { return u.JSON.raw }

func (r *FileChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FileChildUnionChildren is an implicit subunion of [FileChildUnion].
// FileChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [FileChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type FileChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *FileChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileType string

const (
	FileTypeFile FileType = "file"
)

type Footnote struct {
	ID       string               `json:"id"`
	Children []FootnoteChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "footnote".
	Type FootnoteType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Footnote) RawJSON() string { return r.JSON.raw }
func (r *Footnote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Footnote) ImplCalloutChildUnion()   {}
func (Footnote) ImplChunkChildUnion()     {}
func (Footnote) ImplCompanyChildUnion()   {}
func (Footnote) ImplDealChildUnion()      {}
func (Footnote) ImplDocumentChildUnion()  {}
func (Footnote) ImplEquationChildUnion()  {}
func (Footnote) ImplEventChildUnion()     {}
func (Footnote) ImplFileChildUnion()      {}
func (Footnote) ImplFootnoteChildUnion()  {}
func (Footnote) ImplHeadingChildUnion()   {}
func (Footnote) ImplListItemChildUnion()  {}
func (Footnote) ImplMessageChildUnion()   {}
func (Footnote) ImplParagraphChildUnion() {}
func (Footnote) ImplPersonChildUnion()    {}
func (Footnote) ImplQuoteChildUnion()     {}
func (Footnote) ImplTableCellChildUnion() {}
func (Footnote) ImplTaskChildUnion()      {}
func (Footnote) ImplToDoChildUnion()      {}
func (Footnote) ImplWebsiteChildUnion()   {}

// FootnoteChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [FootnoteChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type FootnoteChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children FootnoteChildUnionChildren `json:"children"`
	Text     string                     `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyFootnoteChild is implemented by each variant of [FootnoteChildUnion] to add
// type safety for the return type of [FootnoteChildUnion.AsAny]
type anyFootnoteChild interface {
	ImplFootnoteChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := FootnoteChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u FootnoteChildUnion) AsAny() anyFootnoteChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u FootnoteChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FootnoteChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FootnoteChildUnion) RawJSON() string { return u.JSON.raw }

func (r *FootnoteChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FootnoteChildUnionChildren is an implicit subunion of [FootnoteChildUnion].
// FootnoteChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [FootnoteChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type FootnoteChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *FootnoteChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FootnoteType string

const (
	FootnoteTypeFootnote FootnoteType = "footnote"
)

type Heading struct {
	Level    int64               `json:"level" api:"required"`
	ID       string              `json:"id"`
	Children []HeadingChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "heading".
	Type HeadingType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Level       respjson.Field
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Heading) RawJSON() string { return r.JSON.raw }
func (r *Heading) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Heading) ImplCalloutChildUnion()   {}
func (Heading) ImplChunkChildUnion()     {}
func (Heading) ImplCompanyChildUnion()   {}
func (Heading) ImplDealChildUnion()      {}
func (Heading) ImplDocumentChildUnion()  {}
func (Heading) ImplEquationChildUnion()  {}
func (Heading) ImplEventChildUnion()     {}
func (Heading) ImplFileChildUnion()      {}
func (Heading) ImplFootnoteChildUnion()  {}
func (Heading) ImplHeadingChildUnion()   {}
func (Heading) ImplListItemChildUnion()  {}
func (Heading) ImplMessageChildUnion()   {}
func (Heading) ImplParagraphChildUnion() {}
func (Heading) ImplPersonChildUnion()    {}
func (Heading) ImplQuoteChildUnion()     {}
func (Heading) ImplTableCellChildUnion() {}
func (Heading) ImplTaskChildUnion()      {}
func (Heading) ImplToDoChildUnion()      {}
func (Heading) ImplWebsiteChildUnion()   {}

// HeadingChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [HeadingChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type HeadingChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children HeadingChildUnionChildren `json:"children"`
	Text     string                    `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyHeadingChild is implemented by each variant of [HeadingChildUnion] to add
// type safety for the return type of [HeadingChildUnion.AsAny]
type anyHeadingChild interface {
	ImplHeadingChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := HeadingChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u HeadingChildUnion) AsAny() anyHeadingChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u HeadingChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u HeadingChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u HeadingChildUnion) RawJSON() string { return u.JSON.raw }

func (r *HeadingChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HeadingChildUnionChildren is an implicit subunion of [HeadingChildUnion].
// HeadingChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [HeadingChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type HeadingChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *HeadingChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HeadingType string

const (
	HeadingTypeHeading HeadingType = "heading"
)

type Image struct {
	Src  string `json:"src" api:"required"`
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "image".
	Type ImageType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Src         respjson.Field
		Text        respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Image) ImplCalloutChildUnion()   {}
func (Image) ImplChunkChildUnion()     {}
func (Image) ImplCompanyChildUnion()   {}
func (Image) ImplDealChildUnion()      {}
func (Image) ImplDocumentChildUnion()  {}
func (Image) ImplEquationChildUnion()  {}
func (Image) ImplEventChildUnion()     {}
func (Image) ImplFileChildUnion()      {}
func (Image) ImplFootnoteChildUnion()  {}
func (Image) ImplHeadingChildUnion()   {}
func (Image) ImplListItemChildUnion()  {}
func (Image) ImplMessageChildUnion()   {}
func (Image) ImplParagraphChildUnion() {}
func (Image) ImplPersonChildUnion()    {}
func (Image) ImplQuoteChildUnion()     {}
func (Image) ImplTableCellChildUnion() {}
func (Image) ImplTaskChildUnion()      {}
func (Image) ImplToDoChildUnion()      {}
func (Image) ImplWebsiteChildUnion()   {}

type ImageType string

const (
	ImageTypeImage ImageType = "image"
)

type LineBreak struct {
	ID string `json:"id"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "line_break".
	Type LineBreakType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LineBreak) RawJSON() string { return r.JSON.raw }
func (r *LineBreak) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (LineBreak) ImplCalloutChildUnion()   {}
func (LineBreak) ImplChunkChildUnion()     {}
func (LineBreak) ImplCompanyChildUnion()   {}
func (LineBreak) ImplDealChildUnion()      {}
func (LineBreak) ImplDocumentChildUnion()  {}
func (LineBreak) ImplEquationChildUnion()  {}
func (LineBreak) ImplEventChildUnion()     {}
func (LineBreak) ImplFileChildUnion()      {}
func (LineBreak) ImplFootnoteChildUnion()  {}
func (LineBreak) ImplHeadingChildUnion()   {}
func (LineBreak) ImplListItemChildUnion()  {}
func (LineBreak) ImplMessageChildUnion()   {}
func (LineBreak) ImplParagraphChildUnion() {}
func (LineBreak) ImplPersonChildUnion()    {}
func (LineBreak) ImplQuoteChildUnion()     {}
func (LineBreak) ImplTableCellChildUnion() {}
func (LineBreak) ImplTaskChildUnion()      {}
func (LineBreak) ImplToDoChildUnion()      {}
func (LineBreak) ImplWebsiteChildUnion()   {}

type LineBreakType string

const (
	LineBreakTypeLineBreak LineBreakType = "line_break"
)

type Link struct {
	Text string `json:"text" api:"required"`
	URL  string `json:"url" api:"required"`
	ID   string `json:"id"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "link".
	Type LinkType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		URL         respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Link) RawJSON() string { return r.JSON.raw }
func (r *Link) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Link) ImplCalloutChildUnion()   {}
func (Link) ImplChunkChildUnion()     {}
func (Link) ImplCompanyChildUnion()   {}
func (Link) ImplDealChildUnion()      {}
func (Link) ImplDocumentChildUnion()  {}
func (Link) ImplEquationChildUnion()  {}
func (Link) ImplEventChildUnion()     {}
func (Link) ImplFileChildUnion()      {}
func (Link) ImplFootnoteChildUnion()  {}
func (Link) ImplHeadingChildUnion()   {}
func (Link) ImplListItemChildUnion()  {}
func (Link) ImplMessageChildUnion()   {}
func (Link) ImplParagraphChildUnion() {}
func (Link) ImplPersonChildUnion()    {}
func (Link) ImplQuoteChildUnion()     {}
func (Link) ImplTableCellChildUnion() {}
func (Link) ImplTaskChildUnion()      {}
func (Link) ImplToDoChildUnion()      {}
func (Link) ImplWebsiteChildUnion()   {}

type LinkType string

const (
	LinkTypeLink LinkType = "link"
)

type List struct {
	ID       string           `json:"id"`
	Children []ListChildUnion `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Ordered  bool     `json:"ordered"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "list".
	Type ListType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Ordered     respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r List) RawJSON() string { return r.JSON.raw }
func (r *List) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (List) ImplCalloutChildUnion()   {}
func (List) ImplChunkChildUnion()     {}
func (List) ImplCompanyChildUnion()   {}
func (List) ImplDealChildUnion()      {}
func (List) ImplDocumentChildUnion()  {}
func (List) ImplEquationChildUnion()  {}
func (List) ImplEventChildUnion()     {}
func (List) ImplFileChildUnion()      {}
func (List) ImplFootnoteChildUnion()  {}
func (List) ImplHeadingChildUnion()   {}
func (List) ImplListItemChildUnion()  {}
func (List) ImplMessageChildUnion()   {}
func (List) ImplParagraphChildUnion() {}
func (List) ImplPersonChildUnion()    {}
func (List) ImplQuoteChildUnion()     {}
func (List) ImplTableCellChildUnion() {}
func (List) ImplTaskChildUnion()      {}
func (List) ImplToDoChildUnion()      {}
func (List) ImplWebsiteChildUnion()   {}

// ListChildUnion contains all possible properties and values from [ListItem],
// [ToDo].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ListChildUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]ListItemChildUnion], [[]ToDoChildUnion]
	Children ListChildUnionChildren `json:"children"`
	// This field is from variant [ListItem].
	Metadata Metadata `json:"metadata"`
	Text     string   `json:"text"`
	Type     string   `json:"type"`
	// This field is from variant [ToDo].
	Checked bool `json:"checked"`
	JSON    struct {
		ID       respjson.Field
		Children respjson.Field
		Metadata respjson.Field
		Text     respjson.Field
		Type     respjson.Field
		Checked  respjson.Field
		raw      string
	} `json:"-"`
}

func (u ListChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListChildUnion) AsToDo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListChildUnion) RawJSON() string { return u.JSON.raw }

func (r *ListChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListChildUnionChildren is an implicit subunion of [ListChildUnion].
// ListChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ListChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ListChildUnionChildren struct {
	// This field will be present if the value is a [[]ListItemChildUnion] instead of
	// an object.
	OfChildren []ListItemChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ListChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListType string

const (
	ListTypeList ListType = "list"
)

type ListItem struct {
	ID       string               `json:"id"`
	Children []ListItemChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "list_item".
	Type ListItemType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListItem) RawJSON() string { return r.JSON.raw }
func (r *ListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ListItem) ImplCalloutChildUnion()   {}
func (ListItem) ImplChunkChildUnion()     {}
func (ListItem) ImplCompanyChildUnion()   {}
func (ListItem) ImplDealChildUnion()      {}
func (ListItem) ImplDocumentChildUnion()  {}
func (ListItem) ImplEquationChildUnion()  {}
func (ListItem) ImplEventChildUnion()     {}
func (ListItem) ImplFileChildUnion()      {}
func (ListItem) ImplFootnoteChildUnion()  {}
func (ListItem) ImplHeadingChildUnion()   {}
func (ListItem) ImplListItemChildUnion()  {}
func (ListItem) ImplMessageChildUnion()   {}
func (ListItem) ImplParagraphChildUnion() {}
func (ListItem) ImplPersonChildUnion()    {}
func (ListItem) ImplQuoteChildUnion()     {}
func (ListItem) ImplTableCellChildUnion() {}
func (ListItem) ImplTaskChildUnion()      {}
func (ListItem) ImplToDoChildUnion()      {}
func (ListItem) ImplWebsiteChildUnion()   {}

// ListItemChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [ListItemChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ListItemChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children ListItemChildUnionChildren `json:"children"`
	Text     string                     `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyListItemChild is implemented by each variant of [ListItemChildUnion] to add
// type safety for the return type of [ListItemChildUnion.AsAny]
type anyListItemChild interface {
	ImplListItemChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ListItemChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ListItemChildUnion) AsAny() anyListItemChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u ListItemChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ListItemChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ListItemChildUnion) RawJSON() string { return u.JSON.raw }

func (r *ListItemChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ListItemChildUnionChildren is an implicit subunion of [ListItemChildUnion].
// ListItemChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ListItemChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ListItemChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ListItemChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListItemType string

const (
	ListItemTypeListItem ListItemType = "list_item"
)

type Message struct {
	Date   time.Time `json:"date" api:"required" format:"date-time"`
	Sender Person    `json:"sender" api:"required"`
	ID     string    `json:"id"`
	// The channel or platform where the message was posted, if this Message is not
	// explicitly part of a conversation
	Channel  string              `json:"channel" api:"nullable"`
	Children []MessageChildUnion `json:"children"`
	// Provider message id (e.g. Slack ts, Gmail message id) — merge-dedup key
	ExternalID     string   `json:"external_id" api:"nullable"`
	IsSelf         bool     `json:"is_self" api:"nullable"`
	MentionedUsers []Person `json:"mentioned_users" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata   Metadata `json:"metadata" api:"nullable"`
	NumReplies int64    `json:"num_replies" api:"nullable"`
	// The replies or comments to the message
	Replies  []Message `json:"replies" api:"nullable"`
	Text     string    `json:"text" api:"nullable"`
	ThreadID string    `json:"thread_id" api:"nullable"`
	// The subject or title of the message
	Title string `json:"title" api:"nullable"`
	// Any of "message".
	Type      MessageType `json:"type"`
	UpdatedAt time.Time   `json:"updated_at" api:"nullable" format:"date-time"`
	// The number of upvotes, likes, or reactions on the message
	Upvotes int64 `json:"upvotes" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Date           respjson.Field
		Sender         respjson.Field
		ID             respjson.Field
		Channel        respjson.Field
		Children       respjson.Field
		ExternalID     respjson.Field
		IsSelf         respjson.Field
		MentionedUsers respjson.Field
		Metadata       respjson.Field
		NumReplies     respjson.Field
		Replies        respjson.Field
		Text           respjson.Field
		ThreadID       respjson.Field
		Title          respjson.Field
		Type           respjson.Field
		UpdatedAt      respjson.Field
		Upvotes        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Message) RawJSON() string { return r.JSON.raw }
func (r *Message) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Message) ImplScoredDocumentResponseDocumentUnion() {}
func (Message) ImplMemoryListResponseDocumentUnion()     {}
func (Message) ImplMemoryGetResponseDocumentUnion()      {}

// MessageChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [MessageChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MessageChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children MessageChildUnionChildren `json:"children"`
	Text     string                    `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyMessageChild is implemented by each variant of [MessageChildUnion] to add
// type safety for the return type of [MessageChildUnion.AsAny]
type anyMessageChild interface {
	ImplMessageChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MessageChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u MessageChildUnion) AsAny() anyMessageChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u MessageChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MessageChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MessageChildUnion) RawJSON() string { return u.JSON.raw }

func (r *MessageChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MessageChildUnionChildren is an implicit subunion of [MessageChildUnion].
// MessageChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [MessageChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type MessageChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MessageChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessageType string

const (
	MessageTypeMessage MessageType = "message"
)

// Per-block annotations carried by any Hyperdoc node (ENG-1390).
//
// Out-of-band annotations that travel with a block but aren't part of its content:
// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
// types get added here as typed fields as the need arises.
//
// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
// `metadata` (None) is dropped from serialization entirely, and within a populated
// `Metadata` only the set keys survive.
type Metadata struct {
	EditedBy string           `json:"edited_by" api:"nullable"`
	Sources  []MetadataSource `json:"sources" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EditedBy    respjson.Field
		Sources     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A reference to a memory/chunk that a block's content is grounded in (ENG-1390).
//
// Chunks are the unit persisted to the DB — extracted memories become chunks when
// indexed — so `chunk_id` is the stable pointer back to the source. `resource_id`
// and `source` locate the originating document; `score` carries optional retrieval
// relevance. Kept deliberately self-contained (plain `str` for `source` rather
// than the `DocumentProviders` enum) so the hyperdoc format stays free of
// app-layer imports.
type MetadataSource struct {
	ChunkID    string  `json:"chunk_id" api:"required"`
	ResourceID string  `json:"resource_id" api:"nullable"`
	Score      float64 `json:"score" api:"nullable"`
	Source     string  `json:"source" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID     respjson.Field
		ResourceID  respjson.Field
		Score       respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MetadataSource) RawJSON() string { return r.JSON.raw }
func (r *MetadataSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Paragraph struct {
	ID       string                `json:"id"`
	Children []ParagraphChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "paragraph".
	Type ParagraphType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Paragraph) RawJSON() string { return r.JSON.raw }
func (r *Paragraph) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Paragraph) ImplCalloutChildUnion()   {}
func (Paragraph) ImplChunkChildUnion()     {}
func (Paragraph) ImplCompanyChildUnion()   {}
func (Paragraph) ImplDealChildUnion()      {}
func (Paragraph) ImplDocumentChildUnion()  {}
func (Paragraph) ImplEquationChildUnion()  {}
func (Paragraph) ImplEventChildUnion()     {}
func (Paragraph) ImplFileChildUnion()      {}
func (Paragraph) ImplFootnoteChildUnion()  {}
func (Paragraph) ImplHeadingChildUnion()   {}
func (Paragraph) ImplListItemChildUnion()  {}
func (Paragraph) ImplMessageChildUnion()   {}
func (Paragraph) ImplParagraphChildUnion() {}
func (Paragraph) ImplPersonChildUnion()    {}
func (Paragraph) ImplQuoteChildUnion()     {}
func (Paragraph) ImplTableCellChildUnion() {}
func (Paragraph) ImplTaskChildUnion()      {}
func (Paragraph) ImplToDoChildUnion()      {}
func (Paragraph) ImplWebsiteChildUnion()   {}

// ParagraphChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [ParagraphChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ParagraphChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children ParagraphChildUnionChildren `json:"children"`
	Text     string                      `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyParagraphChild is implemented by each variant of [ParagraphChildUnion] to add
// type safety for the return type of [ParagraphChildUnion.AsAny]
type anyParagraphChild interface {
	ImplParagraphChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ParagraphChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ParagraphChildUnion) AsAny() anyParagraphChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u ParagraphChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ParagraphChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ParagraphChildUnion) RawJSON() string { return u.JSON.raw }

func (r *ParagraphChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ParagraphChildUnionChildren is an implicit subunion of [ParagraphChildUnion].
// ParagraphChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [ParagraphChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ParagraphChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ParagraphChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ParagraphType string

const (
	ParagraphTypeParagraph ParagraphType = "paragraph"
)

type Person struct {
	ID          string             `json:"id"`
	Address     string             `json:"address" api:"nullable"`
	AltNames    []string           `json:"alt_names" api:"nullable"`
	Children    []PersonChildUnion `json:"children"`
	Company     string             `json:"company" api:"nullable"`
	CompanyIDs  []string           `json:"company_ids" api:"nullable"`
	DateOfBirth time.Time          `json:"date_of_birth" api:"nullable" format:"date"`
	DealIDs     []string           `json:"deal_ids" api:"nullable"`
	Email       string             `json:"email" api:"nullable"`
	// All known email addresses; `email` holds the primary one
	Emails   []string `json:"emails" api:"nullable"`
	ImageURL string   `json:"image_url" api:"nullable"`
	JobTitle string   `json:"job_title" api:"nullable"`
	LinkURLs []string `json:"link_urls" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata     Metadata `json:"metadata" api:"nullable"`
	Name         string   `json:"name" api:"nullable"`
	PhoneNumbers []string `json:"phone_numbers" api:"nullable"`
	Tags         []string `json:"tags" api:"nullable"`
	Text         string   `json:"text" api:"nullable"`
	// Any of "person".
	Type     PersonType `json:"type"`
	Username string     `json:"username" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Address      respjson.Field
		AltNames     respjson.Field
		Children     respjson.Field
		Company      respjson.Field
		CompanyIDs   respjson.Field
		DateOfBirth  respjson.Field
		DealIDs      respjson.Field
		Email        respjson.Field
		Emails       respjson.Field
		ImageURL     respjson.Field
		JobTitle     respjson.Field
		LinkURLs     respjson.Field
		Metadata     respjson.Field
		Name         respjson.Field
		PhoneNumbers respjson.Field
		Tags         respjson.Field
		Text         respjson.Field
		Type         respjson.Field
		Username     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Person) RawJSON() string { return r.JSON.raw }
func (r *Person) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Person) ImplScoredDocumentResponseDocumentUnion() {}
func (Person) ImplMemoryListResponseDocumentUnion()     {}
func (Person) ImplMemoryGetResponseDocumentUnion()      {}

// PersonChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [PersonChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PersonChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children PersonChildUnionChildren `json:"children"`
	Text     string                   `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyPersonChild is implemented by each variant of [PersonChildUnion] to add type
// safety for the return type of [PersonChildUnion.AsAny]
type anyPersonChild interface {
	ImplPersonChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := PersonChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PersonChildUnion) AsAny() anyPersonChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u PersonChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PersonChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PersonChildUnion) RawJSON() string { return u.JSON.raw }

func (r *PersonChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PersonChildUnionChildren is an implicit subunion of [PersonChildUnion].
// PersonChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [PersonChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type PersonChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *PersonChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PersonType string

const (
	PersonTypePerson PersonType = "person"
)

// Auditability record attached to an agentic answer.
//
// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
// failed_sources) are derived from in-memory loop state, but `entities` costs one
// indexed DB lookup, so the whole record is only built on request.
type Provenance struct {
	Entities      []ProvenanceEntity `json:"entities"`
	FailedSources []string           `json:"failed_sources"`
	Sources       []ProvenanceSource `json:"sources"`
	Steps         []ProvenanceStep   `json:"steps"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities      respjson.Field
		FailedSources respjson.Field
		Sources       respjson.Field
		Steps         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Provenance) RawJSON() string { return r.JSON.raw }
func (r *Provenance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A canonical entity referenced by the answer's source documents.
type ProvenanceEntity struct {
	ID   string `json:"id" api:"required" format:"uuid"`
	Name string `json:"name" api:"required"`
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvenanceEntity) RawJSON() string { return r.JSON.raw }
func (r *ProvenanceEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source document that informed the final answer (the post-rank result set).
type ProvenanceSource struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source ProvenanceSourceSource `json:"source" api:"required"`
	Score  float64                `json:"score" api:"nullable"`
	Title  string                 `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID  respjson.Field
		Source      respjson.Field
		Score       respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvenanceSource) RawJSON() string { return r.JSON.raw }
func (r *ProvenanceSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProvenanceSourceSource string

const (
	ProvenanceSourceSourceReddit         ProvenanceSourceSource = "reddit"
	ProvenanceSourceSourceNotion         ProvenanceSourceSource = "notion"
	ProvenanceSourceSourceSlack          ProvenanceSourceSource = "slack"
	ProvenanceSourceSourceGoogleCalendar ProvenanceSourceSource = "google_calendar"
	ProvenanceSourceSourceGoogleMail     ProvenanceSourceSource = "google_mail"
	ProvenanceSourceSourceBox            ProvenanceSourceSource = "box"
	ProvenanceSourceSourceDropbox        ProvenanceSourceSource = "dropbox"
	ProvenanceSourceSourceGitHub         ProvenanceSourceSource = "github"
	ProvenanceSourceSourceGoogleDrive    ProvenanceSourceSource = "google_drive"
	ProvenanceSourceSourceVault          ProvenanceSourceSource = "vault"
	ProvenanceSourceSourceWebCrawler     ProvenanceSourceSource = "web_crawler"
	ProvenanceSourceSourceTrace          ProvenanceSourceSource = "trace"
	ProvenanceSourceSourceMicrosoftTeams ProvenanceSourceSource = "microsoft_teams"
	ProvenanceSourceSourceGmailActions   ProvenanceSourceSource = "gmail_actions"
	ProvenanceSourceSourceGranola        ProvenanceSourceSource = "granola"
	ProvenanceSourceSourceFathom         ProvenanceSourceSource = "fathom"
	ProvenanceSourceSourceFireflies      ProvenanceSourceSource = "fireflies"
	ProvenanceSourceSourceLinear         ProvenanceSourceSource = "linear"
	ProvenanceSourceSourceHubspot        ProvenanceSourceSource = "hubspot"
	ProvenanceSourceSourceSalesforce     ProvenanceSourceSource = "salesforce"
	ProvenanceSourceSourceCoda           ProvenanceSourceSource = "coda"
	ProvenanceSourceSourceLightfield     ProvenanceSourceSource = "lightfield"
	ProvenanceSourceSourceGong           ProvenanceSourceSource = "gong"
)

// One tool invocation in the agent's search trajectory (audit trail).
type ProvenanceStep struct {
	Iteration   int64  `json:"iteration" api:"required"`
	Status      string `json:"status" api:"required"`
	Tool        string `json:"tool" api:"required"`
	Query       string `json:"query" api:"nullable"`
	ResultCount int64  `json:"result_count"`
	Source      string `json:"source" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Iteration   respjson.Field
		Status      respjson.Field
		Tool        respjson.Field
		Query       respjson.Field
		ResultCount respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProvenanceStep) RawJSON() string { return r.JSON.raw }
func (r *ProvenanceStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueryResult struct {
	// The answer to the query, if the request was set to answer.
	Answer string `json:"answer" api:"nullable"`
	// The matching documents, each carrying its hyperdoc tree plus query-path
	// score/highlights/summary (ENG-2479 Phase 4).
	Documents []ScoredDocumentResponse `json:"documents"`
	// Errors that occurred during the query. These are meant to help the developer
	// debug the query, and are not meant to be shown to the user.
	Errors []map[string]string `json:"errors" api:"nullable"`
	// Auditability record attached to an agentic answer.
	//
	// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
	// failed_sources) are derived from in-memory loop state, but `entities` costs one
	// indexed DB lookup, so the whole record is only built on request.
	Provenance Provenance `json:"provenance" api:"nullable"`
	// The query string that was issued.
	Query string `json:"query" api:"nullable"`
	// The ID of the query. This can be used to retrieve the query later, or add
	// feedback to it. If the query failed, this will be None.
	QueryID string `json:"query_id" api:"nullable"`
	// The average score of the query feedback, if any.
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer      respjson.Field
		Documents   respjson.Field
		Errors      respjson.Field
		Provenance  respjson.Field
		Query       respjson.Field
		QueryID     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueryResult) RawJSON() string { return r.JSON.raw }
func (r *QueryResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Quote struct {
	ID       string            `json:"id"`
	Children []QuoteChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "quote".
	Type QuoteType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Quote) RawJSON() string { return r.JSON.raw }
func (r *Quote) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Quote) ImplCalloutChildUnion()   {}
func (Quote) ImplChunkChildUnion()     {}
func (Quote) ImplCompanyChildUnion()   {}
func (Quote) ImplDealChildUnion()      {}
func (Quote) ImplDocumentChildUnion()  {}
func (Quote) ImplEquationChildUnion()  {}
func (Quote) ImplEventChildUnion()     {}
func (Quote) ImplFileChildUnion()      {}
func (Quote) ImplFootnoteChildUnion()  {}
func (Quote) ImplHeadingChildUnion()   {}
func (Quote) ImplListItemChildUnion()  {}
func (Quote) ImplMessageChildUnion()   {}
func (Quote) ImplParagraphChildUnion() {}
func (Quote) ImplPersonChildUnion()    {}
func (Quote) ImplQuoteChildUnion()     {}
func (Quote) ImplTableCellChildUnion() {}
func (Quote) ImplTaskChildUnion()      {}
func (Quote) ImplToDoChildUnion()      {}
func (Quote) ImplWebsiteChildUnion()   {}

// QuoteChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [QuoteChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type QuoteChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children QuoteChildUnionChildren `json:"children"`
	Text     string                  `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyQuoteChild is implemented by each variant of [QuoteChildUnion] to add type
// safety for the return type of [QuoteChildUnion.AsAny]
type anyQuoteChild interface {
	ImplQuoteChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := QuoteChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u QuoteChildUnion) AsAny() anyQuoteChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u QuoteChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u QuoteChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u QuoteChildUnion) RawJSON() string { return u.JSON.raw }

func (r *QuoteChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// QuoteChildUnionChildren is an implicit subunion of [QuoteChildUnion].
// QuoteChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [QuoteChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type QuoteChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *QuoteChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QuoteType string

const (
	QuoteTypeQuote QuoteType = "quote"
)

// A `DocumentResponse` plus the query-path fields a `ScoredDocument` carries
// (ENG-2479): relevance score, matched highlights, and the concatenated summary of
// those highlights.
type ScoredDocumentResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// `children` for the body — see the `<Hyperdoc />` renderer.
	Document   ScoredDocumentResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                              `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source ScoredDocumentResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// The document's collection, if any.
	Collection string `json:"collection" api:"nullable"`
	// The document's own date (e.g. email sent date, event date).
	DocumentDate time.Time `json:"document_date" api:"nullable" format:"date-time"`
	// The matched chunks that made this document a hit, with per-chunk scores.
	Highlights []any `json:"highlights"`
	// When Hyperspell first indexed the document.
	IngestedAt time.Time `json:"ingested_at" api:"nullable" format:"date-time"`
	// When the source document was last modified.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Filterable custom metadata attached to the document.
	Metadata map[string]any `json:"metadata"`
	// Relevance of the document to the query.
	Score float64 `json:"score" api:"nullable"`
	// Indexing status of the document.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status ScoredDocumentResponseStatus `json:"status" api:"nullable"`
	// Concatenated text of the matched highlights.
	Summary string `json:"summary" api:"nullable"`
	// Human-readable document title.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document       respjson.Field
		ResourceID     respjson.Field
		Source         respjson.Field
		Type           respjson.Field
		Collection     respjson.Field
		DocumentDate   respjson.Field
		Highlights     respjson.Field
		IngestedAt     respjson.Field
		LastModifiedAt respjson.Field
		Metadata       respjson.Field
		Score          respjson.Field
		Status         respjson.Field
		Summary        respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScoredDocumentResponse) RawJSON() string { return r.JSON.raw }
func (r *ScoredDocumentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoredDocumentResponseDocumentUnion contains all possible properties and values
// from [Document], [Website], [Task], [Person], [Message], [Event], [File],
// [Conversation], [Trace], [Transcript], [Company], [Deal].
//
// Use the [ScoredDocumentResponseDocumentUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ScoredDocumentResponseDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]DocumentChildUnion], [[]WebsiteChildUnion],
	// [[]TaskChildUnion], [[]PersonChildUnion], [[]MessageChildUnion],
	// [[]EventChildUnion], [[]FileChildUnion], [[]Message], [[]TraceChildUnion],
	// [[]Utterance], [[]CompanyChildUnion], [[]DealChildUnion]
	Children ScoredDocumentResponseDocumentUnionChildren `json:"children"`
	// This field is from variant [Document].
	Metadata Metadata `json:"metadata"`
	Text     string   `json:"text"`
	Title    string   `json:"title"`
	// Any of "document", "website", "task", "person", "message", "event", "file",
	// "conversation", "trace", "transcript", "company", "deal".
	Type string `json:"type"`
	// This field is from variant [Website].
	URL         string `json:"url"`
	Description string `json:"description"`
	// This field is from variant [Website].
	Favicon  string `json:"favicon"`
	ImageURL string `json:"image_url"`
	// This field is from variant [Website].
	Language string `json:"language"`
	// This field is from variant [Task].
	Comments []Message `json:"comments"`
	// This field is from variant [Task].
	DueAt time.Time `json:"due_at"`
	// This field is from variant [Task].
	Priority TaskPriority `json:"priority"`
	// This field is from variant [Task].
	Status  TaskStatus `json:"status"`
	Address string     `json:"address"`
	// This field is from variant [Person].
	AltNames []string `json:"alt_names"`
	// This field is from variant [Person].
	Company    string   `json:"company"`
	CompanyIDs []string `json:"company_ids"`
	// This field is from variant [Person].
	DateOfBirth time.Time `json:"date_of_birth"`
	DealIDs     []string  `json:"deal_ids"`
	// This field is from variant [Person].
	Email  string   `json:"email"`
	Emails []string `json:"emails"`
	// This field is from variant [Person].
	JobTitle string `json:"job_title"`
	// This field is from variant [Person].
	LinkURLs     []string `json:"link_urls"`
	Name         string   `json:"name"`
	PhoneNumbers []string `json:"phone_numbers"`
	Tags         []string `json:"tags"`
	// This field is from variant [Person].
	Username string `json:"username"`
	// This field is from variant [Message].
	Date time.Time `json:"date"`
	// This field is from variant [Message].
	Sender  Person `json:"sender"`
	Channel string `json:"channel"`
	// This field is from variant [Message].
	ExternalID string `json:"external_id"`
	// This field is from variant [Message].
	IsSelf bool `json:"is_self"`
	// This field is from variant [Message].
	MentionedUsers []Person `json:"mentioned_users"`
	// This field is from variant [Message].
	NumReplies int64 `json:"num_replies"`
	// This field is from variant [Message].
	Replies []Message `json:"replies"`
	// This field is from variant [Message].
	ThreadID string `json:"thread_id"`
	// This field is from variant [Message].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [Message].
	Upvotes int64 `json:"upvotes"`
	// This field is from variant [Event].
	Attendees []Person `json:"attendees"`
	// This field is from variant [Event].
	EndAt time.Time `json:"end_at"`
	// This field is from variant [Event].
	Location string `json:"location"`
	// This field is from variant [Event].
	MeetingURL string `json:"meeting_url"`
	// This field is from variant [Event].
	StartAt time.Time `json:"start_at"`
	// This field is from variant [File].
	ContentType string `json:"content_type"`
	// This field is from variant [File].
	Filename string `json:"filename"`
	// This field is from variant [File].
	Path []string `json:"path"`
	// This field is from variant [Transcript].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [Transcript].
	Participants []Person `json:"participants"`
	// This field is from variant [Transcript].
	StartedAt  time.Time `json:"started_at"`
	ContactIDs []string  `json:"contact_ids"`
	// This field is from variant [Company].
	Employees int64 `json:"employees"`
	// This field is from variant [Company].
	Industry string `json:"industry"`
	// This field is from variant [Company].
	IsActive bool `json:"is_active"`
	// This field is from variant [Company].
	Timezone string `json:"timezone"`
	// This field is from variant [Company].
	Websites []string `json:"websites"`
	// This field is from variant [Deal].
	Amount float64 `json:"amount"`
	// This field is from variant [Deal].
	ClosedAt time.Time `json:"closed_at"`
	// This field is from variant [Deal].
	Currency string `json:"currency"`
	// This field is from variant [Deal].
	DealSource string `json:"deal_source"`
	// This field is from variant [Deal].
	LostReason string `json:"lost_reason"`
	// This field is from variant [Deal].
	Pipeline string `json:"pipeline"`
	// This field is from variant [Deal].
	Probability float64 `json:"probability"`
	// This field is from variant [Deal].
	Stage string `json:"stage"`
	// This field is from variant [Deal].
	WonReason string `json:"won_reason"`
	JSON      struct {
		ID             respjson.Field
		Children       respjson.Field
		Metadata       respjson.Field
		Text           respjson.Field
		Title          respjson.Field
		Type           respjson.Field
		URL            respjson.Field
		Description    respjson.Field
		Favicon        respjson.Field
		ImageURL       respjson.Field
		Language       respjson.Field
		Comments       respjson.Field
		DueAt          respjson.Field
		Priority       respjson.Field
		Status         respjson.Field
		Address        respjson.Field
		AltNames       respjson.Field
		Company        respjson.Field
		CompanyIDs     respjson.Field
		DateOfBirth    respjson.Field
		DealIDs        respjson.Field
		Email          respjson.Field
		Emails         respjson.Field
		JobTitle       respjson.Field
		LinkURLs       respjson.Field
		Name           respjson.Field
		PhoneNumbers   respjson.Field
		Tags           respjson.Field
		Username       respjson.Field
		Date           respjson.Field
		Sender         respjson.Field
		Channel        respjson.Field
		ExternalID     respjson.Field
		IsSelf         respjson.Field
		MentionedUsers respjson.Field
		NumReplies     respjson.Field
		Replies        respjson.Field
		ThreadID       respjson.Field
		UpdatedAt      respjson.Field
		Upvotes        respjson.Field
		Attendees      respjson.Field
		EndAt          respjson.Field
		Location       respjson.Field
		MeetingURL     respjson.Field
		StartAt        respjson.Field
		ContentType    respjson.Field
		Filename       respjson.Field
		Path           respjson.Field
		EndedAt        respjson.Field
		Participants   respjson.Field
		StartedAt      respjson.Field
		ContactIDs     respjson.Field
		Employees      respjson.Field
		Industry       respjson.Field
		IsActive       respjson.Field
		Timezone       respjson.Field
		Websites       respjson.Field
		Amount         respjson.Field
		ClosedAt       respjson.Field
		Currency       respjson.Field
		DealSource     respjson.Field
		LostReason     respjson.Field
		Pipeline       respjson.Field
		Probability    respjson.Field
		Stage          respjson.Field
		WonReason      respjson.Field
		raw            string
	} `json:"-"`
}

// anyScoredDocumentResponseDocument is implemented by each variant of
// [ScoredDocumentResponseDocumentUnion] to add type safety for the return type of
// [ScoredDocumentResponseDocumentUnion.AsAny]
type anyScoredDocumentResponseDocument interface {
	ImplScoredDocumentResponseDocumentUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ScoredDocumentResponseDocumentUnion.AsAny().(type) {
//	case shared.Document:
//	case shared.Website:
//	case shared.Task:
//	case shared.Person:
//	case shared.Message:
//	case shared.Event:
//	case shared.File:
//	case shared.Conversation:
//	case shared.Trace:
//	case shared.Transcript:
//	case shared.Company:
//	case shared.Deal:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ScoredDocumentResponseDocumentUnion) AsAny() anyScoredDocumentResponseDocument {
	switch u.Type {
	case "document":
		return u.AsDocument()
	case "website":
		return u.AsWebsite()
	case "task":
		return u.AsTask()
	case "person":
		return u.AsPerson()
	case "message":
		return u.AsMessage()
	case "event":
		return u.AsEvent()
	case "file":
		return u.AsFile()
	case "conversation":
		return u.AsConversation()
	case "trace":
		return u.AsTrace()
	case "transcript":
		return u.AsTranscript()
	case "company":
		return u.AsCompany()
	case "deal":
		return u.AsDeal()
	}
	return nil
}

func (u ScoredDocumentResponseDocumentUnion) AsDocument() (v Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsWebsite() (v Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsTask() (v Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsPerson() (v Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsMessage() (v Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsEvent() (v Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsFile() (v File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsConversation() (v Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsTrace() (v Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsTranscript() (v Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsCompany() (v Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ScoredDocumentResponseDocumentUnion) AsDeal() (v Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ScoredDocumentResponseDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *ScoredDocumentResponseDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ScoredDocumentResponseDocumentUnionChildren is an implicit subunion of
// [ScoredDocumentResponseDocumentUnion].
// ScoredDocumentResponseDocumentUnionChildren provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [ScoredDocumentResponseDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ScoredDocumentResponseDocumentUnionChildren struct {
	// This field will be present if the value is a [[]DocumentChildUnion] instead of
	// an object.
	OfChildren []DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ScoredDocumentResponseDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScoredDocumentResponseSource string

const (
	ScoredDocumentResponseSourceReddit         ScoredDocumentResponseSource = "reddit"
	ScoredDocumentResponseSourceNotion         ScoredDocumentResponseSource = "notion"
	ScoredDocumentResponseSourceSlack          ScoredDocumentResponseSource = "slack"
	ScoredDocumentResponseSourceGoogleCalendar ScoredDocumentResponseSource = "google_calendar"
	ScoredDocumentResponseSourceGoogleMail     ScoredDocumentResponseSource = "google_mail"
	ScoredDocumentResponseSourceBox            ScoredDocumentResponseSource = "box"
	ScoredDocumentResponseSourceDropbox        ScoredDocumentResponseSource = "dropbox"
	ScoredDocumentResponseSourceGitHub         ScoredDocumentResponseSource = "github"
	ScoredDocumentResponseSourceGoogleDrive    ScoredDocumentResponseSource = "google_drive"
	ScoredDocumentResponseSourceVault          ScoredDocumentResponseSource = "vault"
	ScoredDocumentResponseSourceWebCrawler     ScoredDocumentResponseSource = "web_crawler"
	ScoredDocumentResponseSourceTrace          ScoredDocumentResponseSource = "trace"
	ScoredDocumentResponseSourceMicrosoftTeams ScoredDocumentResponseSource = "microsoft_teams"
	ScoredDocumentResponseSourceGmailActions   ScoredDocumentResponseSource = "gmail_actions"
	ScoredDocumentResponseSourceGranola        ScoredDocumentResponseSource = "granola"
	ScoredDocumentResponseSourceFathom         ScoredDocumentResponseSource = "fathom"
	ScoredDocumentResponseSourceFireflies      ScoredDocumentResponseSource = "fireflies"
	ScoredDocumentResponseSourceLinear         ScoredDocumentResponseSource = "linear"
	ScoredDocumentResponseSourceHubspot        ScoredDocumentResponseSource = "hubspot"
	ScoredDocumentResponseSourceSalesforce     ScoredDocumentResponseSource = "salesforce"
	ScoredDocumentResponseSourceCoda           ScoredDocumentResponseSource = "coda"
	ScoredDocumentResponseSourceLightfield     ScoredDocumentResponseSource = "lightfield"
	ScoredDocumentResponseSourceGong           ScoredDocumentResponseSource = "gong"
)

// Indexing status of the document.
type ScoredDocumentResponseStatus string

const (
	ScoredDocumentResponseStatusPending       ScoredDocumentResponseStatus = "pending"
	ScoredDocumentResponseStatusProcessing    ScoredDocumentResponseStatus = "processing"
	ScoredDocumentResponseStatusCompleted     ScoredDocumentResponseStatus = "completed"
	ScoredDocumentResponseStatusFailed        ScoredDocumentResponseStatus = "failed"
	ScoredDocumentResponseStatusPendingReview ScoredDocumentResponseStatus = "pending_review"
	ScoredDocumentResponseStatusSkipped       ScoredDocumentResponseStatus = "skipped"
)

type Table struct {
	ID       string     `json:"id"`
	Children []TableRow `json:"children"`
	// Whether the first row should be treated as a header
	HasHeader bool `json:"has_header"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "table".
	Type TableType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		HasHeader   respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Table) RawJSON() string { return r.JSON.raw }
func (r *Table) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Table) ImplCalloutChildUnion()   {}
func (Table) ImplChunkChildUnion()     {}
func (Table) ImplCompanyChildUnion()   {}
func (Table) ImplDealChildUnion()      {}
func (Table) ImplDocumentChildUnion()  {}
func (Table) ImplEquationChildUnion()  {}
func (Table) ImplEventChildUnion()     {}
func (Table) ImplFileChildUnion()      {}
func (Table) ImplFootnoteChildUnion()  {}
func (Table) ImplHeadingChildUnion()   {}
func (Table) ImplListItemChildUnion()  {}
func (Table) ImplMessageChildUnion()   {}
func (Table) ImplParagraphChildUnion() {}
func (Table) ImplPersonChildUnion()    {}
func (Table) ImplQuoteChildUnion()     {}
func (Table) ImplTableCellChildUnion() {}
func (Table) ImplTaskChildUnion()      {}
func (Table) ImplToDoChildUnion()      {}
func (Table) ImplWebsiteChildUnion()   {}

type TableType string

const (
	TableTypeTable TableType = "table"
)

type TableCell struct {
	ID string `json:"id"`
	// Any of "left", "center", "right".
	Align    TableCellAlign        `json:"align"`
	Children []TableCellChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "table_cell".
	Type TableCellType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Align       respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TableCell) RawJSON() string { return r.JSON.raw }
func (r *TableCell) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (TableCell) ImplCalloutChildUnion()   {}
func (TableCell) ImplChunkChildUnion()     {}
func (TableCell) ImplCompanyChildUnion()   {}
func (TableCell) ImplDealChildUnion()      {}
func (TableCell) ImplDocumentChildUnion()  {}
func (TableCell) ImplEquationChildUnion()  {}
func (TableCell) ImplEventChildUnion()     {}
func (TableCell) ImplFileChildUnion()      {}
func (TableCell) ImplFootnoteChildUnion()  {}
func (TableCell) ImplHeadingChildUnion()   {}
func (TableCell) ImplListItemChildUnion()  {}
func (TableCell) ImplMessageChildUnion()   {}
func (TableCell) ImplParagraphChildUnion() {}
func (TableCell) ImplPersonChildUnion()    {}
func (TableCell) ImplQuoteChildUnion()     {}
func (TableCell) ImplTableCellChildUnion() {}
func (TableCell) ImplTaskChildUnion()      {}
func (TableCell) ImplToDoChildUnion()      {}
func (TableCell) ImplWebsiteChildUnion()   {}

type TableCellAlign string

const (
	TableCellAlignLeft   TableCellAlign = "left"
	TableCellAlignCenter TableCellAlign = "center"
	TableCellAlignRight  TableCellAlign = "right"
)

// TableCellChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [TableCellChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TableCellChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children TableCellChildUnionChildren `json:"children"`
	Text     string                      `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyTableCellChild is implemented by each variant of [TableCellChildUnion] to add
// type safety for the return type of [TableCellChildUnion.AsAny]
type anyTableCellChild interface {
	ImplTableCellChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := TableCellChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TableCellChildUnion) AsAny() anyTableCellChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u TableCellChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TableCellChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TableCellChildUnion) RawJSON() string { return u.JSON.raw }

func (r *TableCellChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TableCellChildUnionChildren is an implicit subunion of [TableCellChildUnion].
// TableCellChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [TableCellChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type TableCellChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *TableCellChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TableCellType string

const (
	TableCellTypeTableCell TableCellType = "table_cell"
)

type TableRow struct {
	ID       string      `json:"id"`
	Children []TableCell `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "table_row".
	Type TableRowType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TableRow) RawJSON() string { return r.JSON.raw }
func (r *TableRow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (TableRow) ImplCalloutChildUnion()   {}
func (TableRow) ImplChunkChildUnion()     {}
func (TableRow) ImplCompanyChildUnion()   {}
func (TableRow) ImplDealChildUnion()      {}
func (TableRow) ImplDocumentChildUnion()  {}
func (TableRow) ImplEquationChildUnion()  {}
func (TableRow) ImplEventChildUnion()     {}
func (TableRow) ImplFileChildUnion()      {}
func (TableRow) ImplFootnoteChildUnion()  {}
func (TableRow) ImplHeadingChildUnion()   {}
func (TableRow) ImplListItemChildUnion()  {}
func (TableRow) ImplMessageChildUnion()   {}
func (TableRow) ImplParagraphChildUnion() {}
func (TableRow) ImplPersonChildUnion()    {}
func (TableRow) ImplQuoteChildUnion()     {}
func (TableRow) ImplTableCellChildUnion() {}
func (TableRow) ImplTaskChildUnion()      {}
func (TableRow) ImplToDoChildUnion()      {}
func (TableRow) ImplWebsiteChildUnion()   {}

type TableRowType string

const (
	TableRowTypeTableRow TableRowType = "table_row"
)

type Task struct {
	ID       string           `json:"id"`
	Children []TaskChildUnion `json:"children"`
	Comments []Message        `json:"comments" api:"nullable"`
	DueAt    time.Time        `json:"due_at" api:"nullable" format:"date-time"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "urgent", "high", "medium", "low".
	Priority TaskPriority `json:"priority" api:"nullable"`
	// Any of "completed", "not_started", "in_progress", "cancelled".
	Status TaskStatus `json:"status" api:"nullable"`
	Text   string     `json:"text" api:"nullable"`
	// Any of "task".
	Type TaskType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Comments    respjson.Field
		DueAt       respjson.Field
		Metadata    respjson.Field
		Priority    respjson.Field
		Status      respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Task) RawJSON() string { return r.JSON.raw }
func (r *Task) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Task) ImplScoredDocumentResponseDocumentUnion() {}
func (Task) ImplMemoryListResponseDocumentUnion()     {}
func (Task) ImplMemoryGetResponseDocumentUnion()      {}

// TaskChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [TaskChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children TaskChildUnionChildren `json:"children"`
	Text     string                 `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyTaskChild is implemented by each variant of [TaskChildUnion] to add type
// safety for the return type of [TaskChildUnion.AsAny]
type anyTaskChild interface {
	ImplTaskChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := TaskChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TaskChildUnion) AsAny() anyTaskChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u TaskChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskChildUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskChildUnionChildren is an implicit subunion of [TaskChildUnion].
// TaskChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type TaskChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *TaskChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskPriority string

const (
	TaskPriorityUrgent TaskPriority = "urgent"
	TaskPriorityHigh   TaskPriority = "high"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityLow    TaskPriority = "low"
)

type TaskStatus string

const (
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusNotStarted TaskStatus = "not_started"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

type TaskType string

const (
	TaskTypeTask TaskType = "task"
)

type Text struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Any of "bold", "italic", "underline", "strikethrough", "code", "math".
	Marks []string `json:"marks" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "text".
	Type TextType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ID          respjson.Field
		Marks       respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Text) RawJSON() string { return r.JSON.raw }
func (r *Text) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Text) ImplCalloutChildUnion()   {}
func (Text) ImplChunkChildUnion()     {}
func (Text) ImplCompanyChildUnion()   {}
func (Text) ImplDealChildUnion()      {}
func (Text) ImplDocumentChildUnion()  {}
func (Text) ImplEquationChildUnion()  {}
func (Text) ImplEventChildUnion()     {}
func (Text) ImplFileChildUnion()      {}
func (Text) ImplFootnoteChildUnion()  {}
func (Text) ImplHeadingChildUnion()   {}
func (Text) ImplListItemChildUnion()  {}
func (Text) ImplMessageChildUnion()   {}
func (Text) ImplParagraphChildUnion() {}
func (Text) ImplPersonChildUnion()    {}
func (Text) ImplQuoteChildUnion()     {}
func (Text) ImplTableCellChildUnion() {}
func (Text) ImplTaskChildUnion()      {}
func (Text) ImplToDoChildUnion()      {}
func (Text) ImplWebsiteChildUnion()   {}

type TextType string

const (
	TextTypeText TextType = "text"
)

type ToDo struct {
	ID       string           `json:"id"`
	Checked  bool             `json:"checked"`
	Children []ToDoChildUnion `json:"children" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	// Any of "todo".
	Type ToDoType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Checked     respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToDo) RawJSON() string { return r.JSON.raw }
func (r *ToDo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ToDo) ImplCalloutChildUnion()   {}
func (ToDo) ImplChunkChildUnion()     {}
func (ToDo) ImplCompanyChildUnion()   {}
func (ToDo) ImplDealChildUnion()      {}
func (ToDo) ImplDocumentChildUnion()  {}
func (ToDo) ImplEquationChildUnion()  {}
func (ToDo) ImplEventChildUnion()     {}
func (ToDo) ImplFileChildUnion()      {}
func (ToDo) ImplFootnoteChildUnion()  {}
func (ToDo) ImplHeadingChildUnion()   {}
func (ToDo) ImplListItemChildUnion()  {}
func (ToDo) ImplMessageChildUnion()   {}
func (ToDo) ImplParagraphChildUnion() {}
func (ToDo) ImplPersonChildUnion()    {}
func (ToDo) ImplQuoteChildUnion()     {}
func (ToDo) ImplTableCellChildUnion() {}
func (ToDo) ImplTaskChildUnion()      {}
func (ToDo) ImplToDoChildUnion()      {}
func (ToDo) ImplWebsiteChildUnion()   {}

// ToDoChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [ToDoChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ToDoChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children ToDoChildUnionChildren `json:"children"`
	Text     string                 `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyToDoChild is implemented by each variant of [ToDoChildUnion] to add type
// safety for the return type of [ToDoChildUnion.AsAny]
type anyToDoChild interface {
	ImplToDoChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := ToDoChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u ToDoChildUnion) AsAny() anyToDoChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u ToDoChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToDoChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToDoChildUnion) RawJSON() string { return u.JSON.raw }

func (r *ToDoChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToDoChildUnionChildren is an implicit subunion of [ToDoChildUnion].
// ToDoChildUnionChildren provides convenient access to the sub-properties of the
// union.
//
// For type safety it is recommended to directly use a variant of the
// [ToDoChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type ToDoChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *ToDoChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToDoType string

const (
	ToDoTypeTodo ToDoType = "todo"
)

// A tool/function call made by the assistant.
type ToolCall struct {
	ToolCallID string         `json:"tool_call_id" api:"required"`
	ToolName   string         `json:"tool_name" api:"required"`
	ID         string         `json:"id"`
	Args       map[string]any `json:"args"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "tool_call".
	Type ToolCallType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		ID          respjson.Field
		Args        respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolCall) RawJSON() string { return r.JSON.raw }
func (r *ToolCall) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ToolCall) ImplCalloutChildUnion()   {}
func (ToolCall) ImplChunkChildUnion()     {}
func (ToolCall) ImplCompanyChildUnion()   {}
func (ToolCall) ImplDealChildUnion()      {}
func (ToolCall) ImplDocumentChildUnion()  {}
func (ToolCall) ImplEquationChildUnion()  {}
func (ToolCall) ImplEventChildUnion()     {}
func (ToolCall) ImplFileChildUnion()      {}
func (ToolCall) ImplFootnoteChildUnion()  {}
func (ToolCall) ImplHeadingChildUnion()   {}
func (ToolCall) ImplListItemChildUnion()  {}
func (ToolCall) ImplMessageChildUnion()   {}
func (ToolCall) ImplParagraphChildUnion() {}
func (ToolCall) ImplPersonChildUnion()    {}
func (ToolCall) ImplQuoteChildUnion()     {}
func (ToolCall) ImplTableCellChildUnion() {}
func (ToolCall) ImplTaskChildUnion()      {}
func (ToolCall) ImplToDoChildUnion()      {}
func (ToolCall) ImplTraceChildUnion()     {}
func (ToolCall) ImplWebsiteChildUnion()   {}

type ToolCallType string

const (
	ToolCallTypeToolCall ToolCallType = "tool_call"
)

// The result of a tool call.
type ToolResult struct {
	Output     ToolResultOutputUnion `json:"output" api:"required"`
	ToolCallID string                `json:"tool_call_id" api:"required"`
	ToolName   string                `json:"tool_name" api:"required"`
	ID         string                `json:"id"`
	IsError    bool                  `json:"is_error"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "tool_result".
	Type ToolResultType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		ID          respjson.Field
		IsError     respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ToolResult) RawJSON() string { return r.JSON.raw }
func (r *ToolResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (ToolResult) ImplCalloutChildUnion()   {}
func (ToolResult) ImplChunkChildUnion()     {}
func (ToolResult) ImplCompanyChildUnion()   {}
func (ToolResult) ImplDealChildUnion()      {}
func (ToolResult) ImplDocumentChildUnion()  {}
func (ToolResult) ImplEquationChildUnion()  {}
func (ToolResult) ImplEventChildUnion()     {}
func (ToolResult) ImplFileChildUnion()      {}
func (ToolResult) ImplFootnoteChildUnion()  {}
func (ToolResult) ImplHeadingChildUnion()   {}
func (ToolResult) ImplListItemChildUnion()  {}
func (ToolResult) ImplMessageChildUnion()   {}
func (ToolResult) ImplParagraphChildUnion() {}
func (ToolResult) ImplPersonChildUnion()    {}
func (ToolResult) ImplQuoteChildUnion()     {}
func (ToolResult) ImplTableCellChildUnion() {}
func (ToolResult) ImplTaskChildUnion()      {}
func (ToolResult) ImplToDoChildUnion()      {}
func (ToolResult) ImplTraceChildUnion()     {}
func (ToolResult) ImplWebsiteChildUnion()   {}

// ToolResultOutputUnion contains all possible properties and values from [string],
// [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfToolResultOutputMapItem OfAnyArray]
type ToolResultOutputUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfToolResultOutputMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                  respjson.Field
		OfToolResultOutputMapItem respjson.Field
		OfAnyArray                respjson.Field
		raw                       string
	} `json:"-"`
}

func (u ToolResultOutputUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResultOutputUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ToolResultOutputUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ToolResultOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *ToolResultOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ToolResultType string

const (
	ToolResultTypeToolResult ToolResultType = "tool_result"
)

// An agent trace/transcript containing a sequence of steps.
//
// Steps can be TraceMessage (user/assistant messages or thinking), ToolCall
// (function calls), or ToolResult (tool responses).
type Trace struct {
	ID       string            `json:"id"`
	Children []TraceChildUnion `json:"children"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "trace".
	Type TraceType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Children    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Trace) RawJSON() string { return r.JSON.raw }
func (r *Trace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Trace) ImplScoredDocumentResponseDocumentUnion() {}
func (Trace) ImplMemoryListResponseDocumentUnion()     {}
func (Trace) ImplMemoryGetResponseDocumentUnion()      {}

// TraceChildUnion contains all possible properties and values from [TraceMessage],
// [ToolCall], [ToolResult].
//
// Use the [TraceChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TraceChildUnion struct {
	// This field is from variant [TraceMessage].
	Text string `json:"text"`
	ID   string `json:"id"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Metadata Metadata `json:"metadata"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// Any of "trace_message", "tool_call", "tool_result".
	Type       string `json:"type"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	JSON    struct {
		Text        respjson.Field
		ID          respjson.Field
		MessageType respjson.Field
		Metadata    respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		raw         string
	} `json:"-"`
}

// anyTraceChild is implemented by each variant of [TraceChildUnion] to add type
// safety for the return type of [TraceChildUnion.AsAny]
type anyTraceChild interface {
	ImplTraceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := TraceChildUnion.AsAny().(type) {
//	case shared.TraceMessage:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TraceChildUnion) AsAny() anyTraceChild {
	switch u.Type {
	case "trace_message":
		return u.AsTraceMessage()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	}
	return nil
}

func (u TraceChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TraceChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TraceChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TraceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *TraceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TraceType string

const (
	TraceTypeTrace TraceType = "trace"
)

// A message in an agent trace (user message, assistant message, or thinking).
type TraceMessage struct {
	Text string `json:"text" api:"required"`
	ID   string `json:"id"`
	// Any of "message", "thinking".
	MessageType TraceMessageMessageType `json:"message_type"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	// Any of "user", "assistant".
	Role      TraceMessageRole `json:"role"`
	Timestamp time.Time        `json:"timestamp" api:"nullable" format:"date-time"`
	// Any of "trace_message".
	Type TraceMessageType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ID          respjson.Field
		MessageType respjson.Field
		Metadata    respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TraceMessage) RawJSON() string { return r.JSON.raw }
func (r *TraceMessage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (TraceMessage) ImplCalloutChildUnion()   {}
func (TraceMessage) ImplChunkChildUnion()     {}
func (TraceMessage) ImplCompanyChildUnion()   {}
func (TraceMessage) ImplDealChildUnion()      {}
func (TraceMessage) ImplDocumentChildUnion()  {}
func (TraceMessage) ImplEquationChildUnion()  {}
func (TraceMessage) ImplEventChildUnion()     {}
func (TraceMessage) ImplFileChildUnion()      {}
func (TraceMessage) ImplFootnoteChildUnion()  {}
func (TraceMessage) ImplHeadingChildUnion()   {}
func (TraceMessage) ImplListItemChildUnion()  {}
func (TraceMessage) ImplMessageChildUnion()   {}
func (TraceMessage) ImplParagraphChildUnion() {}
func (TraceMessage) ImplPersonChildUnion()    {}
func (TraceMessage) ImplQuoteChildUnion()     {}
func (TraceMessage) ImplTableCellChildUnion() {}
func (TraceMessage) ImplTaskChildUnion()      {}
func (TraceMessage) ImplToDoChildUnion()      {}
func (TraceMessage) ImplTraceChildUnion()     {}
func (TraceMessage) ImplWebsiteChildUnion()   {}

type TraceMessageMessageType string

const (
	TraceMessageMessageTypeMessage  TraceMessageMessageType = "message"
	TraceMessageMessageTypeThinking TraceMessageMessageType = "thinking"
)

type TraceMessageRole string

const (
	TraceMessageRoleUser      TraceMessageRole = "user"
	TraceMessageRoleAssistant TraceMessageRole = "assistant"
)

type TraceMessageType string

const (
	TraceMessageTypeTraceMessage TraceMessageType = "trace_message"
)

// A time-anchored, speaker-attributed transcript — meetings, calls (ENG-2476/D10;
// mirrors the Trace+TraceStep precedent).
//
// Utterance timestamps are relative offsets from `started_at`, which is the
// absolute wall-clock anchor.
type Transcript struct {
	ID       string      `json:"id"`
	Children []Utterance `json:"children"`
	EndedAt  time.Time   `json:"ended_at" api:"nullable" format:"date-time"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata     Metadata  `json:"metadata" api:"nullable"`
	Participants []Person  `json:"participants"`
	StartedAt    time.Time `json:"started_at" api:"nullable" format:"date-time"`
	Text         string    `json:"text" api:"nullable"`
	Title        string    `json:"title" api:"nullable"`
	// Any of "transcript".
	Type TranscriptType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		Children     respjson.Field
		EndedAt      respjson.Field
		Metadata     respjson.Field
		Participants respjson.Field
		StartedAt    respjson.Field
		Text         respjson.Field
		Title        respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transcript) RawJSON() string { return r.JSON.raw }
func (r *Transcript) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Transcript) ImplScoredDocumentResponseDocumentUnion() {}
func (Transcript) ImplMemoryListResponseDocumentUnion()     {}
func (Transcript) ImplMemoryGetResponseDocumentUnion()      {}

type TranscriptType string

const (
	TranscriptTypeTranscript TranscriptType = "transcript"
)

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
type Utterance struct {
	Text string  `json:"text" api:"required"`
	ID   string  `json:"id"`
	End  float64 `json:"end" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Speaker  *Person  `json:"speaker" api:"nullable"`
	Start    float64  `json:"start" api:"nullable"`
	// Any of "utterance".
	Type UtteranceType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ID          respjson.Field
		End         respjson.Field
		Metadata    respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Utterance) RawJSON() string { return r.JSON.raw }
func (r *Utterance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Utterance) ImplCalloutChildUnion()   {}
func (Utterance) ImplChunkChildUnion()     {}
func (Utterance) ImplCompanyChildUnion()   {}
func (Utterance) ImplDealChildUnion()      {}
func (Utterance) ImplDocumentChildUnion()  {}
func (Utterance) ImplEquationChildUnion()  {}
func (Utterance) ImplEventChildUnion()     {}
func (Utterance) ImplFileChildUnion()      {}
func (Utterance) ImplFootnoteChildUnion()  {}
func (Utterance) ImplHeadingChildUnion()   {}
func (Utterance) ImplListItemChildUnion()  {}
func (Utterance) ImplMessageChildUnion()   {}
func (Utterance) ImplParagraphChildUnion() {}
func (Utterance) ImplPersonChildUnion()    {}
func (Utterance) ImplQuoteChildUnion()     {}
func (Utterance) ImplTableCellChildUnion() {}
func (Utterance) ImplTaskChildUnion()      {}
func (Utterance) ImplToDoChildUnion()      {}
func (Utterance) ImplWebsiteChildUnion()   {}

type UtteranceType string

const (
	UtteranceTypeUtterance UtteranceType = "utterance"
)

type Website struct {
	URL         string              `json:"url" api:"required"`
	ID          string              `json:"id"`
	Children    []WebsiteChildUnion `json:"children"`
	Description string              `json:"description" api:"nullable"`
	Favicon     string              `json:"favicon" api:"nullable"`
	ImageURL    string              `json:"image_url" api:"nullable"`
	Language    string              `json:"language" api:"nullable"`
	// Per-block annotations carried by any Hyperdoc node (ENG-1390).
	//
	// Out-of-band annotations that travel with a block but aren't part of its content:
	// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
	// types get added here as typed fields as the need arises.
	//
	// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
	// `metadata` (None) is dropped from serialization entirely, and within a populated
	// `Metadata` only the set keys survive.
	Metadata Metadata `json:"metadata" api:"nullable"`
	Text     string   `json:"text" api:"nullable"`
	Title    string   `json:"title" api:"nullable"`
	// Any of "website".
	Type WebsiteType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ID          respjson.Field
		Children    respjson.Field
		Description respjson.Field
		Favicon     respjson.Field
		ImageURL    respjson.Field
		Language    respjson.Field
		Metadata    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Website) RawJSON() string { return r.JSON.raw }
func (r *Website) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (Website) ImplScoredDocumentResponseDocumentUnion() {}
func (Website) ImplMemoryListResponseDocumentUnion()     {}
func (Website) ImplMemoryGetResponseDocumentUnion()      {}

// WebsiteChildUnion contains all possible properties and values from [Blob],
// [Callout], [Chunk], [Code], [Comment], [Divider], [Equation], [Footnote],
// [Heading], [Image], [Link], [LineBreak], [List], [ListItem], [Paragraph],
// [Quote], [Table], [TableCell], [TableRow], [Text], [ToDo], [ToolCall],
// [ToolResult], [TraceMessage], [Utterance].
//
// Use the [WebsiteChildUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebsiteChildUnion struct {
	// This field is from variant [Blob].
	Data string `json:"data"`
	// This field is from variant [Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [Blob].
	Metadata Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "paragraph", "quote", "table", "table_cell", "table_row", "text", "todo",
	// "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]CalloutChildUnion], [[]ChunkChildUnion],
	// [[]EquationChildUnion], [[]FootnoteChildUnion], [[]HeadingChildUnion],
	// [[]ListChildUnion], [[]ListItemChildUnion], [[]ParagraphChildUnion],
	// [[]QuoteChildUnion], [[]TableRow], [[]TableCellChildUnion], [[]TableCell],
	// [[]ToDoChildUnion]
	Children WebsiteChildUnionChildren `json:"children"`
	Text     string                    `json:"text"`
	// This field is from variant [Callout].
	Title string `json:"title"`
	// This field is from variant [Code].
	Language string `json:"language"`
	// This field is from variant [Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [Heading].
	Level int64 `json:"level"`
	// This field is from variant [Image].
	Src string `json:"src"`
	// This field is from variant [Link].
	URL string `json:"url"`
	// This field is from variant [List].
	Ordered bool `json:"ordered"`
	// This field is from variant [Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [TableCell].
	Align TableCellAlign `json:"align"`
	// This field is from variant [Text].
	Marks []string `json:"marks"`
	// This field is from variant [ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [ToolResult].
	Output ToolResultOutputUnion `json:"output"`
	// This field is from variant [ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [TraceMessage].
	MessageType TraceMessageMessageType `json:"message_type"`
	// This field is from variant [TraceMessage].
	Role TraceMessageRole `json:"role"`
	// This field is from variant [TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [Utterance].
	End float64 `json:"end"`
	// This field is from variant [Utterance].
	Speaker Person `json:"speaker"`
	// This field is from variant [Utterance].
	Start float64 `json:"start"`
	JSON  struct {
		Data        respjson.Field
		Mimetype    respjson.Field
		ID          respjson.Field
		Metadata    respjson.Field
		Type        respjson.Field
		Children    respjson.Field
		Text        respjson.Field
		Title       respjson.Field
		Language    respjson.Field
		CreatedAt   respjson.Field
		Level       respjson.Field
		Src         respjson.Field
		URL         respjson.Field
		Ordered     respjson.Field
		HasHeader   respjson.Field
		Align       respjson.Field
		Marks       respjson.Field
		Checked     respjson.Field
		ToolCallID  respjson.Field
		ToolName    respjson.Field
		Args        respjson.Field
		Output      respjson.Field
		IsError     respjson.Field
		MessageType respjson.Field
		Role        respjson.Field
		Timestamp   respjson.Field
		End         respjson.Field
		Speaker     respjson.Field
		Start       respjson.Field
		raw         string
	} `json:"-"`
}

// anyWebsiteChild is implemented by each variant of [WebsiteChildUnion] to add
// type safety for the return type of [WebsiteChildUnion.AsAny]
type anyWebsiteChild interface {
	ImplWebsiteChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := WebsiteChildUnion.AsAny().(type) {
//	case shared.Blob:
//	case shared.Callout:
//	case shared.Chunk:
//	case shared.Code:
//	case shared.Comment:
//	case shared.Divider:
//	case shared.Equation:
//	case shared.Footnote:
//	case shared.Heading:
//	case shared.Image:
//	case shared.Link:
//	case shared.LineBreak:
//	case shared.List:
//	case shared.ListItem:
//	case shared.Paragraph:
//	case shared.Quote:
//	case shared.Table:
//	case shared.TableCell:
//	case shared.TableRow:
//	case shared.Text:
//	case shared.ToDo:
//	case shared.ToolCall:
//	case shared.ToolResult:
//	case shared.TraceMessage:
//	case shared.Utterance:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u WebsiteChildUnion) AsAny() anyWebsiteChild {
	switch u.Type {
	case "blob":
		return u.AsBlob()
	case "callout":
		return u.AsCallout()
	case "chunk":
		return u.AsChunk()
	case "code":
		return u.AsCode()
	case "comment":
		return u.AsComment()
	case "divider":
		return u.AsDivider()
	case "equation":
		return u.AsEquation()
	case "footnote":
		return u.AsFootnote()
	case "heading":
		return u.AsHeading()
	case "image":
		return u.AsImage()
	case "link":
		return u.AsLink()
	case "line_break":
		return u.AsLineBreak()
	case "list":
		return u.AsList()
	case "list_item":
		return u.AsListItem()
	case "paragraph":
		return u.AsParagraph()
	case "quote":
		return u.AsQuote()
	case "table":
		return u.AsTable()
	case "table_cell":
		return u.AsTableCell()
	case "table_row":
		return u.AsTableRow()
	case "text":
		return u.AsText()
	case "todo":
		return u.AsTodo()
	case "tool_call":
		return u.AsToolCall()
	case "tool_result":
		return u.AsToolResult()
	case "trace_message":
		return u.AsTraceMessage()
	case "utterance":
		return u.AsUtterance()
	}
	return nil
}

func (u WebsiteChildUnion) AsBlob() (v Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsCallout() (v Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsChunk() (v Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsCode() (v Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsComment() (v Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsDivider() (v Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsEquation() (v Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsFootnote() (v Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsHeading() (v Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsImage() (v Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsLink() (v Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsLineBreak() (v LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsList() (v List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsListItem() (v ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsParagraph() (v Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsQuote() (v Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsTable() (v Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsTableCell() (v TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsTableRow() (v TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsText() (v Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsTodo() (v ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsToolCall() (v ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsToolResult() (v ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsTraceMessage() (v TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebsiteChildUnion) AsUtterance() (v Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebsiteChildUnion) RawJSON() string { return u.JSON.raw }

func (r *WebsiteChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebsiteChildUnionChildren is an implicit subunion of [WebsiteChildUnion].
// WebsiteChildUnionChildren provides convenient access to the sub-properties of
// the union.
//
// For type safety it is recommended to directly use a variant of the
// [WebsiteChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type WebsiteChildUnionChildren struct {
	// This field will be present if the value is a [[]CalloutChildUnion] instead of an
	// object.
	OfChildren []CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *WebsiteChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebsiteType string

const (
	WebsiteTypeWebsite WebsiteType = "website"
)
