// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"github.com/hyperspell/hyperspell-go/internal/apierror"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// This is an alias to an internal type.
type QueryResult = shared.QueryResult

// A `DocumentResponse` plus the query-path fields a `ScoredDocument` carries
// (ENG-2479): relevance score, matched highlights, and the concatenated summary of
// those highlights.
//
// This is an alias to an internal type.
type QueryResultDocument = shared.QueryResultDocument

// The full hyperdoc tree. Switch on `type` for the document frame and recurse
// `children` for the body — see the `<Hyperdoc />` renderer.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentUnion = shared.QueryResultDocumentDocumentUnion

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocument = shared.QueryResultDocumentDocumentDocument

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildUnion = shared.QueryResultDocumentDocumentDocumentChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildBlob = shared.QueryResultDocumentDocumentDocumentChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCallout = shared.QueryResultDocumentDocumentDocumentChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildUnion = shared.QueryResultDocumentDocumentDocumentChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildBlob = shared.QueryResultDocumentDocumentDocumentChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildCode = shared.QueryResultDocumentDocumentDocumentChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildComment = shared.QueryResultDocumentDocumentDocumentChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildDivider = shared.QueryResultDocumentDocumentDocumentChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildImage = shared.QueryResultDocumentDocumentDocumentChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildLink = shared.QueryResultDocumentDocumentDocumentChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildText = shared.QueryResultDocumentDocumentDocumentChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildToolCall = shared.QueryResultDocumentDocumentDocumentChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildToolResult = shared.QueryResultDocumentDocumentDocumentChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunk = shared.QueryResultDocumentDocumentDocumentChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildUnion = shared.QueryResultDocumentDocumentDocumentChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildBlob = shared.QueryResultDocumentDocumentDocumentChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildCode = shared.QueryResultDocumentDocumentDocumentChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildComment = shared.QueryResultDocumentDocumentDocumentChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildDivider = shared.QueryResultDocumentDocumentDocumentChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildImage = shared.QueryResultDocumentDocumentDocumentChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildLink = shared.QueryResultDocumentDocumentDocumentChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildText = shared.QueryResultDocumentDocumentDocumentChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildToolCall = shared.QueryResultDocumentDocumentDocumentChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildToolResult = shared.QueryResultDocumentDocumentDocumentChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildCode = shared.QueryResultDocumentDocumentDocumentChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildComment = shared.QueryResultDocumentDocumentDocumentChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildDivider = shared.QueryResultDocumentDocumentDocumentChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquation = shared.QueryResultDocumentDocumentDocumentChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildUnion = shared.QueryResultDocumentDocumentDocumentChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildBlob = shared.QueryResultDocumentDocumentDocumentChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildCode = shared.QueryResultDocumentDocumentDocumentChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildComment = shared.QueryResultDocumentDocumentDocumentChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildDivider = shared.QueryResultDocumentDocumentDocumentChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildImage = shared.QueryResultDocumentDocumentDocumentChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildLink = shared.QueryResultDocumentDocumentDocumentChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildText = shared.QueryResultDocumentDocumentDocumentChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildToolCall = shared.QueryResultDocumentDocumentDocumentChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildToolResult = shared.QueryResultDocumentDocumentDocumentChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnote = shared.QueryResultDocumentDocumentDocumentChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildUnion = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildBlob = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildCode = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildComment = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildDivider = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildImage = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildLink = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildText = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeading = shared.QueryResultDocumentDocumentDocumentChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildUnion = shared.QueryResultDocumentDocumentDocumentChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildBlob = shared.QueryResultDocumentDocumentDocumentChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildCode = shared.QueryResultDocumentDocumentDocumentChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildComment = shared.QueryResultDocumentDocumentDocumentChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildDivider = shared.QueryResultDocumentDocumentDocumentChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildImage = shared.QueryResultDocumentDocumentDocumentChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildLink = shared.QueryResultDocumentDocumentDocumentChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildText = shared.QueryResultDocumentDocumentDocumentChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildToolCall = shared.QueryResultDocumentDocumentDocumentChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildToolResult = shared.QueryResultDocumentDocumentDocumentChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildImage = shared.QueryResultDocumentDocumentDocumentChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildLink = shared.QueryResultDocumentDocumentDocumentChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildList = shared.QueryResultDocumentDocumentDocumentChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItem = shared.QueryResultDocumentDocumentDocumentChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildUnion = shared.QueryResultDocumentDocumentDocumentChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildBlob = shared.QueryResultDocumentDocumentDocumentChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildCode = shared.QueryResultDocumentDocumentDocumentChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildComment = shared.QueryResultDocumentDocumentDocumentChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildDivider = shared.QueryResultDocumentDocumentDocumentChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildImage = shared.QueryResultDocumentDocumentDocumentChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildLink = shared.QueryResultDocumentDocumentDocumentChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildText = shared.QueryResultDocumentDocumentDocumentChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildToolCall = shared.QueryResultDocumentDocumentDocumentChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildToolResult = shared.QueryResultDocumentDocumentDocumentChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraph = shared.QueryResultDocumentDocumentDocumentChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildUnion = shared.QueryResultDocumentDocumentDocumentChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildBlob = shared.QueryResultDocumentDocumentDocumentChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildCode = shared.QueryResultDocumentDocumentDocumentChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildComment = shared.QueryResultDocumentDocumentDocumentChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildDivider = shared.QueryResultDocumentDocumentDocumentChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildImage = shared.QueryResultDocumentDocumentDocumentChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildLink = shared.QueryResultDocumentDocumentDocumentChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildText = shared.QueryResultDocumentDocumentDocumentChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildToolCall = shared.QueryResultDocumentDocumentDocumentChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildToolResult = shared.QueryResultDocumentDocumentDocumentChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuote = shared.QueryResultDocumentDocumentDocumentChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildUnion = shared.QueryResultDocumentDocumentDocumentChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildBlob = shared.QueryResultDocumentDocumentDocumentChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildCode = shared.QueryResultDocumentDocumentDocumentChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildComment = shared.QueryResultDocumentDocumentDocumentChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildDivider = shared.QueryResultDocumentDocumentDocumentChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildImage = shared.QueryResultDocumentDocumentDocumentChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildLink = shared.QueryResultDocumentDocumentDocumentChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildText = shared.QueryResultDocumentDocumentDocumentChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildToolCall = shared.QueryResultDocumentDocumentDocumentChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildToolResult = shared.QueryResultDocumentDocumentDocumentChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTable = shared.QueryResultDocumentDocumentDocumentChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCell = shared.QueryResultDocumentDocumentDocumentChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildUnion = shared.QueryResultDocumentDocumentDocumentChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildBlob = shared.QueryResultDocumentDocumentDocumentChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildCode = shared.QueryResultDocumentDocumentDocumentChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildComment = shared.QueryResultDocumentDocumentDocumentChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildDivider = shared.QueryResultDocumentDocumentDocumentChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildImage = shared.QueryResultDocumentDocumentDocumentChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildLink = shared.QueryResultDocumentDocumentDocumentChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildText = shared.QueryResultDocumentDocumentDocumentChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildToolCall = shared.QueryResultDocumentDocumentDocumentChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildToolResult = shared.QueryResultDocumentDocumentDocumentChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTableRow = shared.QueryResultDocumentDocumentDocumentChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildText = shared.QueryResultDocumentDocumentDocumentChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodo = shared.QueryResultDocumentDocumentDocumentChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildUnion = shared.QueryResultDocumentDocumentDocumentChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildBlob = shared.QueryResultDocumentDocumentDocumentChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildCode = shared.QueryResultDocumentDocumentDocumentChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildComment = shared.QueryResultDocumentDocumentDocumentChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildDivider = shared.QueryResultDocumentDocumentDocumentChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildImage = shared.QueryResultDocumentDocumentDocumentChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildLink = shared.QueryResultDocumentDocumentDocumentChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildLineBreak = shared.QueryResultDocumentDocumentDocumentChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildText = shared.QueryResultDocumentDocumentDocumentChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildToolCall = shared.QueryResultDocumentDocumentDocumentChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildToolResult = shared.QueryResultDocumentDocumentDocumentChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildToolCall = shared.QueryResultDocumentDocumentDocumentChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildToolResult = shared.QueryResultDocumentDocumentDocumentChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDocumentChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildTraceMessage = shared.QueryResultDocumentDocumentDocumentChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDocumentChildUtterance = shared.QueryResultDocumentDocumentDocumentChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsite = shared.QueryResultDocumentDocumentWebsite

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildUnion = shared.QueryResultDocumentDocumentWebsiteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildBlob = shared.QueryResultDocumentDocumentWebsiteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCallout = shared.QueryResultDocumentDocumentWebsiteChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildUnion = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildBlob = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildCode = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildComment = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildDivider = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildImage = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildLink = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildText = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunk = shared.QueryResultDocumentDocumentWebsiteChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildUnion = shared.QueryResultDocumentDocumentWebsiteChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildBlob = shared.QueryResultDocumentDocumentWebsiteChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildCode = shared.QueryResultDocumentDocumentWebsiteChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildComment = shared.QueryResultDocumentDocumentWebsiteChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildDivider = shared.QueryResultDocumentDocumentWebsiteChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildImage = shared.QueryResultDocumentDocumentWebsiteChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildLink = shared.QueryResultDocumentDocumentWebsiteChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildText = shared.QueryResultDocumentDocumentWebsiteChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildCode = shared.QueryResultDocumentDocumentWebsiteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildComment = shared.QueryResultDocumentDocumentWebsiteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildDivider = shared.QueryResultDocumentDocumentWebsiteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquation = shared.QueryResultDocumentDocumentWebsiteChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildUnion = shared.QueryResultDocumentDocumentWebsiteChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildBlob = shared.QueryResultDocumentDocumentWebsiteChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildCode = shared.QueryResultDocumentDocumentWebsiteChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildComment = shared.QueryResultDocumentDocumentWebsiteChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildDivider = shared.QueryResultDocumentDocumentWebsiteChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildImage = shared.QueryResultDocumentDocumentWebsiteChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildLink = shared.QueryResultDocumentDocumentWebsiteChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildText = shared.QueryResultDocumentDocumentWebsiteChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnote = shared.QueryResultDocumentDocumentWebsiteChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildUnion = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildBlob = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildCode = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildComment = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildDivider = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildImage = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildLink = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildText = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeading = shared.QueryResultDocumentDocumentWebsiteChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildUnion = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildBlob = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildCode = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildComment = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildDivider = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildImage = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildLink = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildText = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildImage = shared.QueryResultDocumentDocumentWebsiteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildLink = shared.QueryResultDocumentDocumentWebsiteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildList = shared.QueryResultDocumentDocumentWebsiteChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItem = shared.QueryResultDocumentDocumentWebsiteChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildUnion = shared.QueryResultDocumentDocumentWebsiteChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildBlob = shared.QueryResultDocumentDocumentWebsiteChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildCode = shared.QueryResultDocumentDocumentWebsiteChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildComment = shared.QueryResultDocumentDocumentWebsiteChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildDivider = shared.QueryResultDocumentDocumentWebsiteChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildImage = shared.QueryResultDocumentDocumentWebsiteChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildLink = shared.QueryResultDocumentDocumentWebsiteChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildText = shared.QueryResultDocumentDocumentWebsiteChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraph = shared.QueryResultDocumentDocumentWebsiteChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildUnion = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildBlob = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildCode = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildComment = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildDivider = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildImage = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildLink = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildText = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuote = shared.QueryResultDocumentDocumentWebsiteChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildUnion = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildBlob = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildCode = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildComment = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildDivider = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildImage = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildLink = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildText = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTable = shared.QueryResultDocumentDocumentWebsiteChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCell = shared.QueryResultDocumentDocumentWebsiteChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildUnion = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildBlob = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildCode = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildComment = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildDivider = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildImage = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildLink = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildText = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTableRow = shared.QueryResultDocumentDocumentWebsiteChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildText = shared.QueryResultDocumentDocumentWebsiteChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodo = shared.QueryResultDocumentDocumentWebsiteChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildUnion = shared.QueryResultDocumentDocumentWebsiteChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildBlob = shared.QueryResultDocumentDocumentWebsiteChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildCode = shared.QueryResultDocumentDocumentWebsiteChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildComment = shared.QueryResultDocumentDocumentWebsiteChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildDivider = shared.QueryResultDocumentDocumentWebsiteChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildImage = shared.QueryResultDocumentDocumentWebsiteChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildLink = shared.QueryResultDocumentDocumentWebsiteChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildLineBreak = shared.QueryResultDocumentDocumentWebsiteChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildText = shared.QueryResultDocumentDocumentWebsiteChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildToolCall = shared.QueryResultDocumentDocumentWebsiteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildToolResult = shared.QueryResultDocumentDocumentWebsiteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentWebsiteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildTraceMessage = shared.QueryResultDocumentDocumentWebsiteChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentWebsiteChildUtterance = shared.QueryResultDocumentDocumentWebsiteChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentTask = shared.QueryResultDocumentDocumentTask

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildUnion = shared.QueryResultDocumentDocumentTaskChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildBlob = shared.QueryResultDocumentDocumentTaskChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCallout = shared.QueryResultDocumentDocumentTaskChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildUnion = shared.QueryResultDocumentDocumentTaskChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildBlob = shared.QueryResultDocumentDocumentTaskChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildCode = shared.QueryResultDocumentDocumentTaskChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildComment = shared.QueryResultDocumentDocumentTaskChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildDivider = shared.QueryResultDocumentDocumentTaskChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildImage = shared.QueryResultDocumentDocumentTaskChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildLink = shared.QueryResultDocumentDocumentTaskChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentTaskChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildText = shared.QueryResultDocumentDocumentTaskChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildToolCall = shared.QueryResultDocumentDocumentTaskChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildToolResult = shared.QueryResultDocumentDocumentTaskChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunk = shared.QueryResultDocumentDocumentTaskChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildUnion = shared.QueryResultDocumentDocumentTaskChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildBlob = shared.QueryResultDocumentDocumentTaskChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildCode = shared.QueryResultDocumentDocumentTaskChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildComment = shared.QueryResultDocumentDocumentTaskChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildDivider = shared.QueryResultDocumentDocumentTaskChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildImage = shared.QueryResultDocumentDocumentTaskChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildLink = shared.QueryResultDocumentDocumentTaskChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildLineBreak = shared.QueryResultDocumentDocumentTaskChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildText = shared.QueryResultDocumentDocumentTaskChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildToolCall = shared.QueryResultDocumentDocumentTaskChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildToolResult = shared.QueryResultDocumentDocumentTaskChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildCode = shared.QueryResultDocumentDocumentTaskChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildComment = shared.QueryResultDocumentDocumentTaskChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildDivider = shared.QueryResultDocumentDocumentTaskChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquation = shared.QueryResultDocumentDocumentTaskChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildUnion = shared.QueryResultDocumentDocumentTaskChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildBlob = shared.QueryResultDocumentDocumentTaskChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildCode = shared.QueryResultDocumentDocumentTaskChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildComment = shared.QueryResultDocumentDocumentTaskChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildDivider = shared.QueryResultDocumentDocumentTaskChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildImage = shared.QueryResultDocumentDocumentTaskChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildLink = shared.QueryResultDocumentDocumentTaskChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildLineBreak = shared.QueryResultDocumentDocumentTaskChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildText = shared.QueryResultDocumentDocumentTaskChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildToolCall = shared.QueryResultDocumentDocumentTaskChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildToolResult = shared.QueryResultDocumentDocumentTaskChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnote = shared.QueryResultDocumentDocumentTaskChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildUnion = shared.QueryResultDocumentDocumentTaskChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildBlob = shared.QueryResultDocumentDocumentTaskChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildCode = shared.QueryResultDocumentDocumentTaskChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildComment = shared.QueryResultDocumentDocumentTaskChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildDivider = shared.QueryResultDocumentDocumentTaskChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildImage = shared.QueryResultDocumentDocumentTaskChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildLink = shared.QueryResultDocumentDocumentTaskChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentTaskChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildText = shared.QueryResultDocumentDocumentTaskChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentTaskChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentTaskChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeading = shared.QueryResultDocumentDocumentTaskChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildUnion = shared.QueryResultDocumentDocumentTaskChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildBlob = shared.QueryResultDocumentDocumentTaskChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildCode = shared.QueryResultDocumentDocumentTaskChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildComment = shared.QueryResultDocumentDocumentTaskChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildDivider = shared.QueryResultDocumentDocumentTaskChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildImage = shared.QueryResultDocumentDocumentTaskChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildLink = shared.QueryResultDocumentDocumentTaskChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentTaskChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildText = shared.QueryResultDocumentDocumentTaskChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildToolCall = shared.QueryResultDocumentDocumentTaskChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildToolResult = shared.QueryResultDocumentDocumentTaskChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildImage = shared.QueryResultDocumentDocumentTaskChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildLink = shared.QueryResultDocumentDocumentTaskChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildLineBreak = shared.QueryResultDocumentDocumentTaskChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildList = shared.QueryResultDocumentDocumentTaskChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItem = shared.QueryResultDocumentDocumentTaskChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildUnion = shared.QueryResultDocumentDocumentTaskChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildBlob = shared.QueryResultDocumentDocumentTaskChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildCode = shared.QueryResultDocumentDocumentTaskChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildComment = shared.QueryResultDocumentDocumentTaskChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildDivider = shared.QueryResultDocumentDocumentTaskChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildImage = shared.QueryResultDocumentDocumentTaskChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildLink = shared.QueryResultDocumentDocumentTaskChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildLineBreak = shared.QueryResultDocumentDocumentTaskChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildText = shared.QueryResultDocumentDocumentTaskChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildToolCall = shared.QueryResultDocumentDocumentTaskChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildToolResult = shared.QueryResultDocumentDocumentTaskChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraph = shared.QueryResultDocumentDocumentTaskChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildUnion = shared.QueryResultDocumentDocumentTaskChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildBlob = shared.QueryResultDocumentDocumentTaskChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildCode = shared.QueryResultDocumentDocumentTaskChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildComment = shared.QueryResultDocumentDocumentTaskChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildDivider = shared.QueryResultDocumentDocumentTaskChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildImage = shared.QueryResultDocumentDocumentTaskChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildLink = shared.QueryResultDocumentDocumentTaskChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentTaskChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildText = shared.QueryResultDocumentDocumentTaskChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildToolCall = shared.QueryResultDocumentDocumentTaskChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildToolResult = shared.QueryResultDocumentDocumentTaskChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuote = shared.QueryResultDocumentDocumentTaskChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildUnion = shared.QueryResultDocumentDocumentTaskChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildBlob = shared.QueryResultDocumentDocumentTaskChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildCode = shared.QueryResultDocumentDocumentTaskChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildComment = shared.QueryResultDocumentDocumentTaskChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildDivider = shared.QueryResultDocumentDocumentTaskChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildImage = shared.QueryResultDocumentDocumentTaskChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildLink = shared.QueryResultDocumentDocumentTaskChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentTaskChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildText = shared.QueryResultDocumentDocumentTaskChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildToolCall = shared.QueryResultDocumentDocumentTaskChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildToolResult = shared.QueryResultDocumentDocumentTaskChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTable = shared.QueryResultDocumentDocumentTaskChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCell = shared.QueryResultDocumentDocumentTaskChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildUnion = shared.QueryResultDocumentDocumentTaskChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildBlob = shared.QueryResultDocumentDocumentTaskChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildCode = shared.QueryResultDocumentDocumentTaskChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildComment = shared.QueryResultDocumentDocumentTaskChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildDivider = shared.QueryResultDocumentDocumentTaskChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildImage = shared.QueryResultDocumentDocumentTaskChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildLink = shared.QueryResultDocumentDocumentTaskChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentTaskChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildText = shared.QueryResultDocumentDocumentTaskChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildToolCall = shared.QueryResultDocumentDocumentTaskChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildToolResult = shared.QueryResultDocumentDocumentTaskChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTableRow = shared.QueryResultDocumentDocumentTaskChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildText = shared.QueryResultDocumentDocumentTaskChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodo = shared.QueryResultDocumentDocumentTaskChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildUnion = shared.QueryResultDocumentDocumentTaskChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildBlob = shared.QueryResultDocumentDocumentTaskChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildCode = shared.QueryResultDocumentDocumentTaskChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildComment = shared.QueryResultDocumentDocumentTaskChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildDivider = shared.QueryResultDocumentDocumentTaskChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildImage = shared.QueryResultDocumentDocumentTaskChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildLink = shared.QueryResultDocumentDocumentTaskChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildLineBreak = shared.QueryResultDocumentDocumentTaskChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildText = shared.QueryResultDocumentDocumentTaskChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildToolCall = shared.QueryResultDocumentDocumentTaskChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildToolResult = shared.QueryResultDocumentDocumentTaskChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildToolCall = shared.QueryResultDocumentDocumentTaskChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildToolResult = shared.QueryResultDocumentDocumentTaskChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildTraceMessage = shared.QueryResultDocumentDocumentTaskChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskChildUtterance = shared.QueryResultDocumentDocumentTaskChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskComment = shared.QueryResultDocumentDocumentTaskComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSender = shared.QueryResultDocumentDocumentTaskCommentSender

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildUnion = shared.QueryResultDocumentDocumentTaskCommentSenderChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildBlob = shared.QueryResultDocumentDocumentTaskCommentSenderChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildCode = shared.QueryResultDocumentDocumentTaskCommentSenderChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildComment = shared.QueryResultDocumentDocumentTaskCommentSenderChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildDivider = shared.QueryResultDocumentDocumentTaskCommentSenderChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildImage = shared.QueryResultDocumentDocumentTaskCommentSenderChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildLink = shared.QueryResultDocumentDocumentTaskCommentSenderChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentSenderChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildText = shared.QueryResultDocumentDocumentTaskCommentSenderChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildToolCall = shared.QueryResultDocumentDocumentTaskCommentSenderChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildToolResult = shared.QueryResultDocumentDocumentTaskCommentSenderChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentSenderChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentSenderChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentSenderChildTraceMessage

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCallout = shared.QueryResultDocumentDocumentTaskCommentChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildCode = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildComment = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildImage = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildLink = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildText = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunk = shared.QueryResultDocumentDocumentTaskCommentChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildCode = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildComment = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildImage = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildLink = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildText = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildCode = shared.QueryResultDocumentDocumentTaskCommentChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildComment = shared.QueryResultDocumentDocumentTaskCommentChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquation = shared.QueryResultDocumentDocumentTaskCommentChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildCode = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildComment = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildImage = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildLink = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildText = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnote = shared.QueryResultDocumentDocumentTaskCommentChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildCode = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildComment = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildImage = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildLink = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildText = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeading = shared.QueryResultDocumentDocumentTaskCommentChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildCode = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildComment = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildImage = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildLink = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildText = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildImage = shared.QueryResultDocumentDocumentTaskCommentChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildLink = shared.QueryResultDocumentDocumentTaskCommentChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildList = shared.QueryResultDocumentDocumentTaskCommentChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItem = shared.QueryResultDocumentDocumentTaskCommentChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildCode = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildComment = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildImage = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildLink = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildText = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraph = shared.QueryResultDocumentDocumentTaskCommentChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildCode = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildComment = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildImage = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildLink = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildText = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuote = shared.QueryResultDocumentDocumentTaskCommentChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildCode = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildComment = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildImage = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildLink = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildText = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTable = shared.QueryResultDocumentDocumentTaskCommentChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCell = shared.QueryResultDocumentDocumentTaskCommentChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildCode = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildComment = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildImage = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildLink = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildText = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTableRow = shared.QueryResultDocumentDocumentTaskCommentChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildText = shared.QueryResultDocumentDocumentTaskCommentChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodo = shared.QueryResultDocumentDocumentTaskCommentChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildUnion = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildBlob = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildCode = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildComment = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildDivider = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildImage = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildLink = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildText = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildToolCall = shared.QueryResultDocumentDocumentTaskCommentChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildToolResult = shared.QueryResultDocumentDocumentTaskCommentChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentChildUtterance = shared.QueryResultDocumentDocumentTaskCommentChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUser = shared.QueryResultDocumentDocumentTaskCommentMentionedUser

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildUnion = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildBlob = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildCode = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildComment = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildDivider = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildImage = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildLink = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildLineBreak = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildText = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildToolCall = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildToolResult = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTaskCommentMentionedUserChildTraceMessage = shared.QueryResultDocumentDocumentTaskCommentMentionedUserChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentPerson = shared.QueryResultDocumentDocumentPerson

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildUnion = shared.QueryResultDocumentDocumentPersonChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildBlob = shared.QueryResultDocumentDocumentPersonChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildCode = shared.QueryResultDocumentDocumentPersonChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildComment = shared.QueryResultDocumentDocumentPersonChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildDivider = shared.QueryResultDocumentDocumentPersonChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildImage = shared.QueryResultDocumentDocumentPersonChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildLink = shared.QueryResultDocumentDocumentPersonChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildLineBreak = shared.QueryResultDocumentDocumentPersonChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildText = shared.QueryResultDocumentDocumentPersonChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildToolCall = shared.QueryResultDocumentDocumentPersonChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildToolResult = shared.QueryResultDocumentDocumentPersonChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildToolResultOutputUnion = shared.QueryResultDocumentDocumentPersonChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentPersonChildTraceMessage = shared.QueryResultDocumentDocumentPersonChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessage = shared.QueryResultDocumentDocumentMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSender = shared.QueryResultDocumentDocumentMessageSender

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildUnion = shared.QueryResultDocumentDocumentMessageSenderChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildBlob = shared.QueryResultDocumentDocumentMessageSenderChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildCode = shared.QueryResultDocumentDocumentMessageSenderChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildComment = shared.QueryResultDocumentDocumentMessageSenderChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildDivider = shared.QueryResultDocumentDocumentMessageSenderChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildImage = shared.QueryResultDocumentDocumentMessageSenderChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildLink = shared.QueryResultDocumentDocumentMessageSenderChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildLineBreak = shared.QueryResultDocumentDocumentMessageSenderChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildText = shared.QueryResultDocumentDocumentMessageSenderChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildToolCall = shared.QueryResultDocumentDocumentMessageSenderChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildToolResult = shared.QueryResultDocumentDocumentMessageSenderChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageSenderChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageSenderChildTraceMessage = shared.QueryResultDocumentDocumentMessageSenderChildTraceMessage

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildUnion = shared.QueryResultDocumentDocumentMessageChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildBlob = shared.QueryResultDocumentDocumentMessageChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCallout = shared.QueryResultDocumentDocumentMessageChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildUnion = shared.QueryResultDocumentDocumentMessageChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildBlob = shared.QueryResultDocumentDocumentMessageChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildCode = shared.QueryResultDocumentDocumentMessageChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildComment = shared.QueryResultDocumentDocumentMessageChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildDivider = shared.QueryResultDocumentDocumentMessageChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildImage = shared.QueryResultDocumentDocumentMessageChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildLink = shared.QueryResultDocumentDocumentMessageChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentMessageChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildText = shared.QueryResultDocumentDocumentMessageChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildToolCall = shared.QueryResultDocumentDocumentMessageChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildToolResult = shared.QueryResultDocumentDocumentMessageChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunk = shared.QueryResultDocumentDocumentMessageChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildUnion = shared.QueryResultDocumentDocumentMessageChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildBlob = shared.QueryResultDocumentDocumentMessageChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildCode = shared.QueryResultDocumentDocumentMessageChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildComment = shared.QueryResultDocumentDocumentMessageChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildDivider = shared.QueryResultDocumentDocumentMessageChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildImage = shared.QueryResultDocumentDocumentMessageChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildLink = shared.QueryResultDocumentDocumentMessageChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildLineBreak = shared.QueryResultDocumentDocumentMessageChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildText = shared.QueryResultDocumentDocumentMessageChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildToolCall = shared.QueryResultDocumentDocumentMessageChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildToolResult = shared.QueryResultDocumentDocumentMessageChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildCode = shared.QueryResultDocumentDocumentMessageChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildComment = shared.QueryResultDocumentDocumentMessageChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildDivider = shared.QueryResultDocumentDocumentMessageChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquation = shared.QueryResultDocumentDocumentMessageChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildUnion = shared.QueryResultDocumentDocumentMessageChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildBlob = shared.QueryResultDocumentDocumentMessageChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildCode = shared.QueryResultDocumentDocumentMessageChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildComment = shared.QueryResultDocumentDocumentMessageChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildDivider = shared.QueryResultDocumentDocumentMessageChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildImage = shared.QueryResultDocumentDocumentMessageChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildLink = shared.QueryResultDocumentDocumentMessageChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildLineBreak = shared.QueryResultDocumentDocumentMessageChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildText = shared.QueryResultDocumentDocumentMessageChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildToolCall = shared.QueryResultDocumentDocumentMessageChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildToolResult = shared.QueryResultDocumentDocumentMessageChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnote = shared.QueryResultDocumentDocumentMessageChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildUnion = shared.QueryResultDocumentDocumentMessageChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildBlob = shared.QueryResultDocumentDocumentMessageChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildCode = shared.QueryResultDocumentDocumentMessageChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildComment = shared.QueryResultDocumentDocumentMessageChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildDivider = shared.QueryResultDocumentDocumentMessageChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildImage = shared.QueryResultDocumentDocumentMessageChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildLink = shared.QueryResultDocumentDocumentMessageChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentMessageChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildText = shared.QueryResultDocumentDocumentMessageChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentMessageChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentMessageChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeading = shared.QueryResultDocumentDocumentMessageChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildUnion = shared.QueryResultDocumentDocumentMessageChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildBlob = shared.QueryResultDocumentDocumentMessageChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildCode = shared.QueryResultDocumentDocumentMessageChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildComment = shared.QueryResultDocumentDocumentMessageChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildDivider = shared.QueryResultDocumentDocumentMessageChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildImage = shared.QueryResultDocumentDocumentMessageChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildLink = shared.QueryResultDocumentDocumentMessageChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentMessageChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildText = shared.QueryResultDocumentDocumentMessageChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildToolCall = shared.QueryResultDocumentDocumentMessageChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildToolResult = shared.QueryResultDocumentDocumentMessageChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildImage = shared.QueryResultDocumentDocumentMessageChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildLink = shared.QueryResultDocumentDocumentMessageChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildLineBreak = shared.QueryResultDocumentDocumentMessageChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildList = shared.QueryResultDocumentDocumentMessageChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItem = shared.QueryResultDocumentDocumentMessageChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildUnion = shared.QueryResultDocumentDocumentMessageChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildBlob = shared.QueryResultDocumentDocumentMessageChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildCode = shared.QueryResultDocumentDocumentMessageChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildComment = shared.QueryResultDocumentDocumentMessageChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildDivider = shared.QueryResultDocumentDocumentMessageChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildImage = shared.QueryResultDocumentDocumentMessageChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildLink = shared.QueryResultDocumentDocumentMessageChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildLineBreak = shared.QueryResultDocumentDocumentMessageChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildText = shared.QueryResultDocumentDocumentMessageChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildToolCall = shared.QueryResultDocumentDocumentMessageChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildToolResult = shared.QueryResultDocumentDocumentMessageChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraph = shared.QueryResultDocumentDocumentMessageChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildUnion = shared.QueryResultDocumentDocumentMessageChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildBlob = shared.QueryResultDocumentDocumentMessageChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildCode = shared.QueryResultDocumentDocumentMessageChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildComment = shared.QueryResultDocumentDocumentMessageChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildDivider = shared.QueryResultDocumentDocumentMessageChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildImage = shared.QueryResultDocumentDocumentMessageChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildLink = shared.QueryResultDocumentDocumentMessageChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentMessageChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildText = shared.QueryResultDocumentDocumentMessageChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildToolCall = shared.QueryResultDocumentDocumentMessageChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildToolResult = shared.QueryResultDocumentDocumentMessageChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuote = shared.QueryResultDocumentDocumentMessageChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildUnion = shared.QueryResultDocumentDocumentMessageChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildBlob = shared.QueryResultDocumentDocumentMessageChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildCode = shared.QueryResultDocumentDocumentMessageChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildComment = shared.QueryResultDocumentDocumentMessageChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildDivider = shared.QueryResultDocumentDocumentMessageChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildImage = shared.QueryResultDocumentDocumentMessageChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildLink = shared.QueryResultDocumentDocumentMessageChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentMessageChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildText = shared.QueryResultDocumentDocumentMessageChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildToolCall = shared.QueryResultDocumentDocumentMessageChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildToolResult = shared.QueryResultDocumentDocumentMessageChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTable = shared.QueryResultDocumentDocumentMessageChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCell = shared.QueryResultDocumentDocumentMessageChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildUnion = shared.QueryResultDocumentDocumentMessageChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildBlob = shared.QueryResultDocumentDocumentMessageChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildCode = shared.QueryResultDocumentDocumentMessageChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildComment = shared.QueryResultDocumentDocumentMessageChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildDivider = shared.QueryResultDocumentDocumentMessageChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildImage = shared.QueryResultDocumentDocumentMessageChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildLink = shared.QueryResultDocumentDocumentMessageChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentMessageChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildText = shared.QueryResultDocumentDocumentMessageChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildToolCall = shared.QueryResultDocumentDocumentMessageChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildToolResult = shared.QueryResultDocumentDocumentMessageChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTableRow = shared.QueryResultDocumentDocumentMessageChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildText = shared.QueryResultDocumentDocumentMessageChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodo = shared.QueryResultDocumentDocumentMessageChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildUnion = shared.QueryResultDocumentDocumentMessageChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildBlob = shared.QueryResultDocumentDocumentMessageChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildCode = shared.QueryResultDocumentDocumentMessageChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildComment = shared.QueryResultDocumentDocumentMessageChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildDivider = shared.QueryResultDocumentDocumentMessageChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildImage = shared.QueryResultDocumentDocumentMessageChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildLink = shared.QueryResultDocumentDocumentMessageChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildLineBreak = shared.QueryResultDocumentDocumentMessageChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildText = shared.QueryResultDocumentDocumentMessageChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildToolCall = shared.QueryResultDocumentDocumentMessageChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildToolResult = shared.QueryResultDocumentDocumentMessageChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildToolCall = shared.QueryResultDocumentDocumentMessageChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildToolResult = shared.QueryResultDocumentDocumentMessageChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildTraceMessage = shared.QueryResultDocumentDocumentMessageChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageChildUtterance = shared.QueryResultDocumentDocumentMessageChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUser = shared.QueryResultDocumentDocumentMessageMentionedUser

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildUnion = shared.QueryResultDocumentDocumentMessageMentionedUserChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildBlob = shared.QueryResultDocumentDocumentMessageMentionedUserChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildCode = shared.QueryResultDocumentDocumentMessageMentionedUserChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildComment = shared.QueryResultDocumentDocumentMessageMentionedUserChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildDivider = shared.QueryResultDocumentDocumentMessageMentionedUserChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildImage = shared.QueryResultDocumentDocumentMessageMentionedUserChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildLink = shared.QueryResultDocumentDocumentMessageMentionedUserChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildLineBreak = shared.QueryResultDocumentDocumentMessageMentionedUserChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildText = shared.QueryResultDocumentDocumentMessageMentionedUserChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildToolCall = shared.QueryResultDocumentDocumentMessageMentionedUserChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildToolResult = shared.QueryResultDocumentDocumentMessageMentionedUserChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildToolResultOutputUnion = shared.QueryResultDocumentDocumentMessageMentionedUserChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentMessageMentionedUserChildTraceMessage = shared.QueryResultDocumentDocumentMessageMentionedUserChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEvent = shared.QueryResultDocumentDocumentEvent

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendee = shared.QueryResultDocumentDocumentEventAttendee

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildUnion = shared.QueryResultDocumentDocumentEventAttendeeChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildBlob = shared.QueryResultDocumentDocumentEventAttendeeChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildCode = shared.QueryResultDocumentDocumentEventAttendeeChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildComment = shared.QueryResultDocumentDocumentEventAttendeeChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildDivider = shared.QueryResultDocumentDocumentEventAttendeeChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildImage = shared.QueryResultDocumentDocumentEventAttendeeChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildLink = shared.QueryResultDocumentDocumentEventAttendeeChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildLineBreak = shared.QueryResultDocumentDocumentEventAttendeeChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildText = shared.QueryResultDocumentDocumentEventAttendeeChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildToolCall = shared.QueryResultDocumentDocumentEventAttendeeChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildToolResult = shared.QueryResultDocumentDocumentEventAttendeeChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventAttendeeChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventAttendeeChildTraceMessage = shared.QueryResultDocumentDocumentEventAttendeeChildTraceMessage

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildUnion = shared.QueryResultDocumentDocumentEventChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildBlob = shared.QueryResultDocumentDocumentEventChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCallout = shared.QueryResultDocumentDocumentEventChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildUnion = shared.QueryResultDocumentDocumentEventChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildBlob = shared.QueryResultDocumentDocumentEventChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildCode = shared.QueryResultDocumentDocumentEventChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildComment = shared.QueryResultDocumentDocumentEventChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildDivider = shared.QueryResultDocumentDocumentEventChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildImage = shared.QueryResultDocumentDocumentEventChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildLink = shared.QueryResultDocumentDocumentEventChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentEventChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildText = shared.QueryResultDocumentDocumentEventChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildToolCall = shared.QueryResultDocumentDocumentEventChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildToolResult = shared.QueryResultDocumentDocumentEventChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentEventChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunk = shared.QueryResultDocumentDocumentEventChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildUnion = shared.QueryResultDocumentDocumentEventChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildBlob = shared.QueryResultDocumentDocumentEventChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildCode = shared.QueryResultDocumentDocumentEventChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildComment = shared.QueryResultDocumentDocumentEventChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildDivider = shared.QueryResultDocumentDocumentEventChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildImage = shared.QueryResultDocumentDocumentEventChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildLink = shared.QueryResultDocumentDocumentEventChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildLineBreak = shared.QueryResultDocumentDocumentEventChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildText = shared.QueryResultDocumentDocumentEventChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildToolCall = shared.QueryResultDocumentDocumentEventChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildToolResult = shared.QueryResultDocumentDocumentEventChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentEventChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildCode = shared.QueryResultDocumentDocumentEventChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildComment = shared.QueryResultDocumentDocumentEventChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildDivider = shared.QueryResultDocumentDocumentEventChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquation = shared.QueryResultDocumentDocumentEventChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildUnion = shared.QueryResultDocumentDocumentEventChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildBlob = shared.QueryResultDocumentDocumentEventChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildCode = shared.QueryResultDocumentDocumentEventChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildComment = shared.QueryResultDocumentDocumentEventChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildDivider = shared.QueryResultDocumentDocumentEventChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildImage = shared.QueryResultDocumentDocumentEventChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildLink = shared.QueryResultDocumentDocumentEventChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildLineBreak = shared.QueryResultDocumentDocumentEventChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildText = shared.QueryResultDocumentDocumentEventChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildToolCall = shared.QueryResultDocumentDocumentEventChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildToolResult = shared.QueryResultDocumentDocumentEventChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentEventChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnote = shared.QueryResultDocumentDocumentEventChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildUnion = shared.QueryResultDocumentDocumentEventChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildBlob = shared.QueryResultDocumentDocumentEventChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildCode = shared.QueryResultDocumentDocumentEventChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildComment = shared.QueryResultDocumentDocumentEventChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildDivider = shared.QueryResultDocumentDocumentEventChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildImage = shared.QueryResultDocumentDocumentEventChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildLink = shared.QueryResultDocumentDocumentEventChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentEventChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildText = shared.QueryResultDocumentDocumentEventChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentEventChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentEventChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentEventChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeading = shared.QueryResultDocumentDocumentEventChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildUnion = shared.QueryResultDocumentDocumentEventChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildBlob = shared.QueryResultDocumentDocumentEventChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildCode = shared.QueryResultDocumentDocumentEventChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildComment = shared.QueryResultDocumentDocumentEventChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildDivider = shared.QueryResultDocumentDocumentEventChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildImage = shared.QueryResultDocumentDocumentEventChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildLink = shared.QueryResultDocumentDocumentEventChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentEventChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildText = shared.QueryResultDocumentDocumentEventChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildToolCall = shared.QueryResultDocumentDocumentEventChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildToolResult = shared.QueryResultDocumentDocumentEventChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentEventChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildImage = shared.QueryResultDocumentDocumentEventChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildLink = shared.QueryResultDocumentDocumentEventChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildLineBreak = shared.QueryResultDocumentDocumentEventChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildList = shared.QueryResultDocumentDocumentEventChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItem = shared.QueryResultDocumentDocumentEventChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildUnion = shared.QueryResultDocumentDocumentEventChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildBlob = shared.QueryResultDocumentDocumentEventChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildCode = shared.QueryResultDocumentDocumentEventChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildComment = shared.QueryResultDocumentDocumentEventChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildDivider = shared.QueryResultDocumentDocumentEventChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildImage = shared.QueryResultDocumentDocumentEventChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildLink = shared.QueryResultDocumentDocumentEventChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildLineBreak = shared.QueryResultDocumentDocumentEventChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildText = shared.QueryResultDocumentDocumentEventChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildToolCall = shared.QueryResultDocumentDocumentEventChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildToolResult = shared.QueryResultDocumentDocumentEventChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentEventChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraph = shared.QueryResultDocumentDocumentEventChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildUnion = shared.QueryResultDocumentDocumentEventChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildBlob = shared.QueryResultDocumentDocumentEventChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildCode = shared.QueryResultDocumentDocumentEventChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildComment = shared.QueryResultDocumentDocumentEventChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildDivider = shared.QueryResultDocumentDocumentEventChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildImage = shared.QueryResultDocumentDocumentEventChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildLink = shared.QueryResultDocumentDocumentEventChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentEventChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildText = shared.QueryResultDocumentDocumentEventChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildToolCall = shared.QueryResultDocumentDocumentEventChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildToolResult = shared.QueryResultDocumentDocumentEventChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentEventChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuote = shared.QueryResultDocumentDocumentEventChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildUnion = shared.QueryResultDocumentDocumentEventChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildBlob = shared.QueryResultDocumentDocumentEventChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildCode = shared.QueryResultDocumentDocumentEventChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildComment = shared.QueryResultDocumentDocumentEventChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildDivider = shared.QueryResultDocumentDocumentEventChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildImage = shared.QueryResultDocumentDocumentEventChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildLink = shared.QueryResultDocumentDocumentEventChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentEventChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildText = shared.QueryResultDocumentDocumentEventChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildToolCall = shared.QueryResultDocumentDocumentEventChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildToolResult = shared.QueryResultDocumentDocumentEventChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentEventChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTable = shared.QueryResultDocumentDocumentEventChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCell = shared.QueryResultDocumentDocumentEventChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildUnion = shared.QueryResultDocumentDocumentEventChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildBlob = shared.QueryResultDocumentDocumentEventChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildCode = shared.QueryResultDocumentDocumentEventChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildComment = shared.QueryResultDocumentDocumentEventChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildDivider = shared.QueryResultDocumentDocumentEventChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildImage = shared.QueryResultDocumentDocumentEventChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildLink = shared.QueryResultDocumentDocumentEventChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentEventChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildText = shared.QueryResultDocumentDocumentEventChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildToolCall = shared.QueryResultDocumentDocumentEventChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildToolResult = shared.QueryResultDocumentDocumentEventChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentEventChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTableRow = shared.QueryResultDocumentDocumentEventChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildText = shared.QueryResultDocumentDocumentEventChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodo = shared.QueryResultDocumentDocumentEventChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildUnion = shared.QueryResultDocumentDocumentEventChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildBlob = shared.QueryResultDocumentDocumentEventChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildCode = shared.QueryResultDocumentDocumentEventChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildComment = shared.QueryResultDocumentDocumentEventChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildDivider = shared.QueryResultDocumentDocumentEventChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildImage = shared.QueryResultDocumentDocumentEventChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildLink = shared.QueryResultDocumentDocumentEventChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildLineBreak = shared.QueryResultDocumentDocumentEventChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildText = shared.QueryResultDocumentDocumentEventChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildToolCall = shared.QueryResultDocumentDocumentEventChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildToolResult = shared.QueryResultDocumentDocumentEventChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentEventChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildToolCall = shared.QueryResultDocumentDocumentEventChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildToolResult = shared.QueryResultDocumentDocumentEventChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildToolResultOutputUnion = shared.QueryResultDocumentDocumentEventChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildTraceMessage = shared.QueryResultDocumentDocumentEventChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentEventChildUtterance = shared.QueryResultDocumentDocumentEventChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentFile = shared.QueryResultDocumentDocumentFile

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildUnion = shared.QueryResultDocumentDocumentFileChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildBlob = shared.QueryResultDocumentDocumentFileChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCallout = shared.QueryResultDocumentDocumentFileChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildUnion = shared.QueryResultDocumentDocumentFileChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildBlob = shared.QueryResultDocumentDocumentFileChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildCode = shared.QueryResultDocumentDocumentFileChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildComment = shared.QueryResultDocumentDocumentFileChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildDivider = shared.QueryResultDocumentDocumentFileChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildImage = shared.QueryResultDocumentDocumentFileChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildLink = shared.QueryResultDocumentDocumentFileChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentFileChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildText = shared.QueryResultDocumentDocumentFileChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildToolCall = shared.QueryResultDocumentDocumentFileChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildToolResult = shared.QueryResultDocumentDocumentFileChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentFileChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunk = shared.QueryResultDocumentDocumentFileChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildUnion = shared.QueryResultDocumentDocumentFileChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildBlob = shared.QueryResultDocumentDocumentFileChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildCode = shared.QueryResultDocumentDocumentFileChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildComment = shared.QueryResultDocumentDocumentFileChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildDivider = shared.QueryResultDocumentDocumentFileChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildImage = shared.QueryResultDocumentDocumentFileChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildLink = shared.QueryResultDocumentDocumentFileChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildLineBreak = shared.QueryResultDocumentDocumentFileChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildText = shared.QueryResultDocumentDocumentFileChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildToolCall = shared.QueryResultDocumentDocumentFileChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildToolResult = shared.QueryResultDocumentDocumentFileChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentFileChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildCode = shared.QueryResultDocumentDocumentFileChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildComment = shared.QueryResultDocumentDocumentFileChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildDivider = shared.QueryResultDocumentDocumentFileChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquation = shared.QueryResultDocumentDocumentFileChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildUnion = shared.QueryResultDocumentDocumentFileChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildBlob = shared.QueryResultDocumentDocumentFileChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildCode = shared.QueryResultDocumentDocumentFileChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildComment = shared.QueryResultDocumentDocumentFileChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildDivider = shared.QueryResultDocumentDocumentFileChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildImage = shared.QueryResultDocumentDocumentFileChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildLink = shared.QueryResultDocumentDocumentFileChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildLineBreak = shared.QueryResultDocumentDocumentFileChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildText = shared.QueryResultDocumentDocumentFileChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildToolCall = shared.QueryResultDocumentDocumentFileChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildToolResult = shared.QueryResultDocumentDocumentFileChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentFileChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnote = shared.QueryResultDocumentDocumentFileChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildUnion = shared.QueryResultDocumentDocumentFileChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildBlob = shared.QueryResultDocumentDocumentFileChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildCode = shared.QueryResultDocumentDocumentFileChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildComment = shared.QueryResultDocumentDocumentFileChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildDivider = shared.QueryResultDocumentDocumentFileChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildImage = shared.QueryResultDocumentDocumentFileChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildLink = shared.QueryResultDocumentDocumentFileChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentFileChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildText = shared.QueryResultDocumentDocumentFileChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentFileChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentFileChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentFileChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeading = shared.QueryResultDocumentDocumentFileChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildUnion = shared.QueryResultDocumentDocumentFileChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildBlob = shared.QueryResultDocumentDocumentFileChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildCode = shared.QueryResultDocumentDocumentFileChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildComment = shared.QueryResultDocumentDocumentFileChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildDivider = shared.QueryResultDocumentDocumentFileChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildImage = shared.QueryResultDocumentDocumentFileChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildLink = shared.QueryResultDocumentDocumentFileChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentFileChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildText = shared.QueryResultDocumentDocumentFileChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildToolCall = shared.QueryResultDocumentDocumentFileChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildToolResult = shared.QueryResultDocumentDocumentFileChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentFileChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildImage = shared.QueryResultDocumentDocumentFileChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildLink = shared.QueryResultDocumentDocumentFileChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildLineBreak = shared.QueryResultDocumentDocumentFileChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildList = shared.QueryResultDocumentDocumentFileChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItem = shared.QueryResultDocumentDocumentFileChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildUnion = shared.QueryResultDocumentDocumentFileChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildBlob = shared.QueryResultDocumentDocumentFileChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildCode = shared.QueryResultDocumentDocumentFileChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildComment = shared.QueryResultDocumentDocumentFileChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildDivider = shared.QueryResultDocumentDocumentFileChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildImage = shared.QueryResultDocumentDocumentFileChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildLink = shared.QueryResultDocumentDocumentFileChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildLineBreak = shared.QueryResultDocumentDocumentFileChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildText = shared.QueryResultDocumentDocumentFileChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildToolCall = shared.QueryResultDocumentDocumentFileChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildToolResult = shared.QueryResultDocumentDocumentFileChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentFileChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraph = shared.QueryResultDocumentDocumentFileChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildUnion = shared.QueryResultDocumentDocumentFileChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildBlob = shared.QueryResultDocumentDocumentFileChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildCode = shared.QueryResultDocumentDocumentFileChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildComment = shared.QueryResultDocumentDocumentFileChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildDivider = shared.QueryResultDocumentDocumentFileChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildImage = shared.QueryResultDocumentDocumentFileChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildLink = shared.QueryResultDocumentDocumentFileChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentFileChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildText = shared.QueryResultDocumentDocumentFileChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildToolCall = shared.QueryResultDocumentDocumentFileChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildToolResult = shared.QueryResultDocumentDocumentFileChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentFileChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuote = shared.QueryResultDocumentDocumentFileChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildUnion = shared.QueryResultDocumentDocumentFileChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildBlob = shared.QueryResultDocumentDocumentFileChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildCode = shared.QueryResultDocumentDocumentFileChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildComment = shared.QueryResultDocumentDocumentFileChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildDivider = shared.QueryResultDocumentDocumentFileChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildImage = shared.QueryResultDocumentDocumentFileChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildLink = shared.QueryResultDocumentDocumentFileChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentFileChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildText = shared.QueryResultDocumentDocumentFileChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildToolCall = shared.QueryResultDocumentDocumentFileChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildToolResult = shared.QueryResultDocumentDocumentFileChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentFileChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTable = shared.QueryResultDocumentDocumentFileChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCell = shared.QueryResultDocumentDocumentFileChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildUnion = shared.QueryResultDocumentDocumentFileChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildBlob = shared.QueryResultDocumentDocumentFileChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildCode = shared.QueryResultDocumentDocumentFileChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildComment = shared.QueryResultDocumentDocumentFileChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildDivider = shared.QueryResultDocumentDocumentFileChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildImage = shared.QueryResultDocumentDocumentFileChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildLink = shared.QueryResultDocumentDocumentFileChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentFileChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildText = shared.QueryResultDocumentDocumentFileChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildToolCall = shared.QueryResultDocumentDocumentFileChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildToolResult = shared.QueryResultDocumentDocumentFileChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentFileChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTableRow = shared.QueryResultDocumentDocumentFileChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildText = shared.QueryResultDocumentDocumentFileChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodo = shared.QueryResultDocumentDocumentFileChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildUnion = shared.QueryResultDocumentDocumentFileChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildBlob = shared.QueryResultDocumentDocumentFileChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildCode = shared.QueryResultDocumentDocumentFileChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildComment = shared.QueryResultDocumentDocumentFileChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildDivider = shared.QueryResultDocumentDocumentFileChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildImage = shared.QueryResultDocumentDocumentFileChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildLink = shared.QueryResultDocumentDocumentFileChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildLineBreak = shared.QueryResultDocumentDocumentFileChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildText = shared.QueryResultDocumentDocumentFileChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildToolCall = shared.QueryResultDocumentDocumentFileChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildToolResult = shared.QueryResultDocumentDocumentFileChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentFileChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildToolCall = shared.QueryResultDocumentDocumentFileChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildToolResult = shared.QueryResultDocumentDocumentFileChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildToolResultOutputUnion = shared.QueryResultDocumentDocumentFileChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildTraceMessage = shared.QueryResultDocumentDocumentFileChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentFileChildUtterance = shared.QueryResultDocumentDocumentFileChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversation = shared.QueryResultDocumentDocumentConversation

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChild = shared.QueryResultDocumentDocumentConversationChild

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSender = shared.QueryResultDocumentDocumentConversationChildSender

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildUnion = shared.QueryResultDocumentDocumentConversationChildSenderChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildBlob = shared.QueryResultDocumentDocumentConversationChildSenderChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildCode = shared.QueryResultDocumentDocumentConversationChildSenderChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildComment = shared.QueryResultDocumentDocumentConversationChildSenderChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildDivider = shared.QueryResultDocumentDocumentConversationChildSenderChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildImage = shared.QueryResultDocumentDocumentConversationChildSenderChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildLink = shared.QueryResultDocumentDocumentConversationChildSenderChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildLineBreak = shared.QueryResultDocumentDocumentConversationChildSenderChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildText = shared.QueryResultDocumentDocumentConversationChildSenderChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildToolCall = shared.QueryResultDocumentDocumentConversationChildSenderChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildToolResult = shared.QueryResultDocumentDocumentConversationChildSenderChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildSenderChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildSenderChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildSenderChildTraceMessage

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildUnion = shared.QueryResultDocumentDocumentConversationChildChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildBlob = shared.QueryResultDocumentDocumentConversationChildChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCallout = shared.QueryResultDocumentDocumentConversationChildChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildUnion = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildBlob = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildCode = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildComment = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildDivider = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildImage = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildLink = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildText = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunk = shared.QueryResultDocumentDocumentConversationChildChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildUnion = shared.QueryResultDocumentDocumentConversationChildChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildBlob = shared.QueryResultDocumentDocumentConversationChildChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildCode = shared.QueryResultDocumentDocumentConversationChildChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildComment = shared.QueryResultDocumentDocumentConversationChildChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildDivider = shared.QueryResultDocumentDocumentConversationChildChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildImage = shared.QueryResultDocumentDocumentConversationChildChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildLink = shared.QueryResultDocumentDocumentConversationChildChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildText = shared.QueryResultDocumentDocumentConversationChildChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildCode = shared.QueryResultDocumentDocumentConversationChildChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildComment = shared.QueryResultDocumentDocumentConversationChildChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildDivider = shared.QueryResultDocumentDocumentConversationChildChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquation = shared.QueryResultDocumentDocumentConversationChildChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildUnion = shared.QueryResultDocumentDocumentConversationChildChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildBlob = shared.QueryResultDocumentDocumentConversationChildChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildCode = shared.QueryResultDocumentDocumentConversationChildChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildComment = shared.QueryResultDocumentDocumentConversationChildChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildDivider = shared.QueryResultDocumentDocumentConversationChildChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildImage = shared.QueryResultDocumentDocumentConversationChildChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildLink = shared.QueryResultDocumentDocumentConversationChildChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildText = shared.QueryResultDocumentDocumentConversationChildChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnote = shared.QueryResultDocumentDocumentConversationChildChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildUnion = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildBlob = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildCode = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildComment = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildDivider = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildImage = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildLink = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildText = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeading = shared.QueryResultDocumentDocumentConversationChildChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildUnion = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildBlob = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildCode = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildComment = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildDivider = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildImage = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildLink = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildText = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildImage = shared.QueryResultDocumentDocumentConversationChildChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildLink = shared.QueryResultDocumentDocumentConversationChildChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildList = shared.QueryResultDocumentDocumentConversationChildChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItem = shared.QueryResultDocumentDocumentConversationChildChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildUnion = shared.QueryResultDocumentDocumentConversationChildChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildBlob = shared.QueryResultDocumentDocumentConversationChildChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildCode = shared.QueryResultDocumentDocumentConversationChildChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildComment = shared.QueryResultDocumentDocumentConversationChildChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildDivider = shared.QueryResultDocumentDocumentConversationChildChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildImage = shared.QueryResultDocumentDocumentConversationChildChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildLink = shared.QueryResultDocumentDocumentConversationChildChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildText = shared.QueryResultDocumentDocumentConversationChildChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraph = shared.QueryResultDocumentDocumentConversationChildChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildUnion = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildBlob = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildCode = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildComment = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildDivider = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildImage = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildLink = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildText = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuote = shared.QueryResultDocumentDocumentConversationChildChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildUnion = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildBlob = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildCode = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildComment = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildDivider = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildImage = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildLink = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildText = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTable = shared.QueryResultDocumentDocumentConversationChildChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCell = shared.QueryResultDocumentDocumentConversationChildChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildUnion = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildBlob = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildCode = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildComment = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildDivider = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildImage = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildLink = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildText = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTableRow = shared.QueryResultDocumentDocumentConversationChildChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildText = shared.QueryResultDocumentDocumentConversationChildChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodo = shared.QueryResultDocumentDocumentConversationChildChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildUnion = shared.QueryResultDocumentDocumentConversationChildChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildBlob = shared.QueryResultDocumentDocumentConversationChildChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildCode = shared.QueryResultDocumentDocumentConversationChildChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildComment = shared.QueryResultDocumentDocumentConversationChildChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildDivider = shared.QueryResultDocumentDocumentConversationChildChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildImage = shared.QueryResultDocumentDocumentConversationChildChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildLink = shared.QueryResultDocumentDocumentConversationChildChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildLineBreak = shared.QueryResultDocumentDocumentConversationChildChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildText = shared.QueryResultDocumentDocumentConversationChildChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildToolCall = shared.QueryResultDocumentDocumentConversationChildChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildToolResult = shared.QueryResultDocumentDocumentConversationChildChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildChildUtterance = shared.QueryResultDocumentDocumentConversationChildChildUtterance

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUser = shared.QueryResultDocumentDocumentConversationChildMentionedUser

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildUnion = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildBlob = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildCode = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildComment = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildDivider = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildImage = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildLink = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildLineBreak = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildText = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildToolCall = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildToolResult = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildToolResultOutputUnion = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentConversationChildMentionedUserChildTraceMessage = shared.QueryResultDocumentDocumentConversationChildMentionedUserChildTraceMessage

// An agent trace/transcript containing a sequence of steps.
//
// Steps can be TraceMessage (user/assistant messages or thinking), ToolCall
// (function calls), or ToolResult (tool responses).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTrace = shared.QueryResultDocumentDocumentTrace

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTraceChildUnion = shared.QueryResultDocumentDocumentTraceChildUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTraceChildTraceMessage = shared.QueryResultDocumentDocumentTraceChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTraceChildToolCall = shared.QueryResultDocumentDocumentTraceChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTraceChildToolResult = shared.QueryResultDocumentDocumentTraceChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTraceChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTraceChildToolResultOutputUnion

// A time-anchored, speaker-attributed transcript — meetings, calls (ENG-2476/D10;
// mirrors the Trace+TraceStep precedent).
//
// Utterance timestamps are relative offsets from `started_at`, which is the
// absolute wall-clock anchor.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscript = shared.QueryResultDocumentDocumentTranscript

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptChild = shared.QueryResultDocumentDocumentTranscriptChild

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipant = shared.QueryResultDocumentDocumentTranscriptParticipant

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildUnion = shared.QueryResultDocumentDocumentTranscriptParticipantChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildBlob = shared.QueryResultDocumentDocumentTranscriptParticipantChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildCode = shared.QueryResultDocumentDocumentTranscriptParticipantChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildComment = shared.QueryResultDocumentDocumentTranscriptParticipantChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildDivider = shared.QueryResultDocumentDocumentTranscriptParticipantChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildImage = shared.QueryResultDocumentDocumentTranscriptParticipantChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildLink = shared.QueryResultDocumentDocumentTranscriptParticipantChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildLineBreak = shared.QueryResultDocumentDocumentTranscriptParticipantChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildText = shared.QueryResultDocumentDocumentTranscriptParticipantChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildToolCall = shared.QueryResultDocumentDocumentTranscriptParticipantChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildToolResult = shared.QueryResultDocumentDocumentTranscriptParticipantChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildToolResultOutputUnion = shared.QueryResultDocumentDocumentTranscriptParticipantChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentTranscriptParticipantChildTraceMessage = shared.QueryResultDocumentDocumentTranscriptParticipantChildTraceMessage

// A CRM company/account record (ENG-2476/D10).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompany = shared.QueryResultDocumentDocumentCompany

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildUnion = shared.QueryResultDocumentDocumentCompanyChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildBlob = shared.QueryResultDocumentDocumentCompanyChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCallout = shared.QueryResultDocumentDocumentCompanyChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildUnion = shared.QueryResultDocumentDocumentCompanyChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildBlob = shared.QueryResultDocumentDocumentCompanyChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildCode = shared.QueryResultDocumentDocumentCompanyChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildComment = shared.QueryResultDocumentDocumentCompanyChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildDivider = shared.QueryResultDocumentDocumentCompanyChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildImage = shared.QueryResultDocumentDocumentCompanyChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildLink = shared.QueryResultDocumentDocumentCompanyChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildText = shared.QueryResultDocumentDocumentCompanyChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildToolCall = shared.QueryResultDocumentDocumentCompanyChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildToolResult = shared.QueryResultDocumentDocumentCompanyChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunk = shared.QueryResultDocumentDocumentCompanyChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildUnion = shared.QueryResultDocumentDocumentCompanyChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildBlob = shared.QueryResultDocumentDocumentCompanyChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildCode = shared.QueryResultDocumentDocumentCompanyChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildComment = shared.QueryResultDocumentDocumentCompanyChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildDivider = shared.QueryResultDocumentDocumentCompanyChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildImage = shared.QueryResultDocumentDocumentCompanyChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildLink = shared.QueryResultDocumentDocumentCompanyChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildText = shared.QueryResultDocumentDocumentCompanyChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildToolCall = shared.QueryResultDocumentDocumentCompanyChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildToolResult = shared.QueryResultDocumentDocumentCompanyChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildCode = shared.QueryResultDocumentDocumentCompanyChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildComment = shared.QueryResultDocumentDocumentCompanyChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildDivider = shared.QueryResultDocumentDocumentCompanyChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquation = shared.QueryResultDocumentDocumentCompanyChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildUnion = shared.QueryResultDocumentDocumentCompanyChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildBlob = shared.QueryResultDocumentDocumentCompanyChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildCode = shared.QueryResultDocumentDocumentCompanyChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildComment = shared.QueryResultDocumentDocumentCompanyChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildDivider = shared.QueryResultDocumentDocumentCompanyChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildImage = shared.QueryResultDocumentDocumentCompanyChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildLink = shared.QueryResultDocumentDocumentCompanyChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildText = shared.QueryResultDocumentDocumentCompanyChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildToolCall = shared.QueryResultDocumentDocumentCompanyChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildToolResult = shared.QueryResultDocumentDocumentCompanyChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnote = shared.QueryResultDocumentDocumentCompanyChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildUnion = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildBlob = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildCode = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildComment = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildDivider = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildImage = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildLink = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildText = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeading = shared.QueryResultDocumentDocumentCompanyChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildUnion = shared.QueryResultDocumentDocumentCompanyChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildBlob = shared.QueryResultDocumentDocumentCompanyChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildCode = shared.QueryResultDocumentDocumentCompanyChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildComment = shared.QueryResultDocumentDocumentCompanyChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildDivider = shared.QueryResultDocumentDocumentCompanyChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildImage = shared.QueryResultDocumentDocumentCompanyChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildLink = shared.QueryResultDocumentDocumentCompanyChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildText = shared.QueryResultDocumentDocumentCompanyChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildToolCall = shared.QueryResultDocumentDocumentCompanyChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildToolResult = shared.QueryResultDocumentDocumentCompanyChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildImage = shared.QueryResultDocumentDocumentCompanyChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildLink = shared.QueryResultDocumentDocumentCompanyChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildList = shared.QueryResultDocumentDocumentCompanyChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItem = shared.QueryResultDocumentDocumentCompanyChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildUnion = shared.QueryResultDocumentDocumentCompanyChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildBlob = shared.QueryResultDocumentDocumentCompanyChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildCode = shared.QueryResultDocumentDocumentCompanyChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildComment = shared.QueryResultDocumentDocumentCompanyChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildDivider = shared.QueryResultDocumentDocumentCompanyChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildImage = shared.QueryResultDocumentDocumentCompanyChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildLink = shared.QueryResultDocumentDocumentCompanyChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildText = shared.QueryResultDocumentDocumentCompanyChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildToolCall = shared.QueryResultDocumentDocumentCompanyChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildToolResult = shared.QueryResultDocumentDocumentCompanyChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraph = shared.QueryResultDocumentDocumentCompanyChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildUnion = shared.QueryResultDocumentDocumentCompanyChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildBlob = shared.QueryResultDocumentDocumentCompanyChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildCode = shared.QueryResultDocumentDocumentCompanyChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildComment = shared.QueryResultDocumentDocumentCompanyChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildDivider = shared.QueryResultDocumentDocumentCompanyChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildImage = shared.QueryResultDocumentDocumentCompanyChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildLink = shared.QueryResultDocumentDocumentCompanyChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildText = shared.QueryResultDocumentDocumentCompanyChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildToolCall = shared.QueryResultDocumentDocumentCompanyChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildToolResult = shared.QueryResultDocumentDocumentCompanyChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuote = shared.QueryResultDocumentDocumentCompanyChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildUnion = shared.QueryResultDocumentDocumentCompanyChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildBlob = shared.QueryResultDocumentDocumentCompanyChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildCode = shared.QueryResultDocumentDocumentCompanyChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildComment = shared.QueryResultDocumentDocumentCompanyChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildDivider = shared.QueryResultDocumentDocumentCompanyChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildImage = shared.QueryResultDocumentDocumentCompanyChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildLink = shared.QueryResultDocumentDocumentCompanyChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildText = shared.QueryResultDocumentDocumentCompanyChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildToolCall = shared.QueryResultDocumentDocumentCompanyChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildToolResult = shared.QueryResultDocumentDocumentCompanyChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTable = shared.QueryResultDocumentDocumentCompanyChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCell = shared.QueryResultDocumentDocumentCompanyChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildUnion = shared.QueryResultDocumentDocumentCompanyChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildBlob = shared.QueryResultDocumentDocumentCompanyChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildCode = shared.QueryResultDocumentDocumentCompanyChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildComment = shared.QueryResultDocumentDocumentCompanyChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildDivider = shared.QueryResultDocumentDocumentCompanyChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildImage = shared.QueryResultDocumentDocumentCompanyChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildLink = shared.QueryResultDocumentDocumentCompanyChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildText = shared.QueryResultDocumentDocumentCompanyChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildToolCall = shared.QueryResultDocumentDocumentCompanyChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildToolResult = shared.QueryResultDocumentDocumentCompanyChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTableRow = shared.QueryResultDocumentDocumentCompanyChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildText = shared.QueryResultDocumentDocumentCompanyChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodo = shared.QueryResultDocumentDocumentCompanyChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildUnion = shared.QueryResultDocumentDocumentCompanyChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildBlob = shared.QueryResultDocumentDocumentCompanyChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildCode = shared.QueryResultDocumentDocumentCompanyChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildComment = shared.QueryResultDocumentDocumentCompanyChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildDivider = shared.QueryResultDocumentDocumentCompanyChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildImage = shared.QueryResultDocumentDocumentCompanyChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildLink = shared.QueryResultDocumentDocumentCompanyChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildLineBreak = shared.QueryResultDocumentDocumentCompanyChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildText = shared.QueryResultDocumentDocumentCompanyChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildToolCall = shared.QueryResultDocumentDocumentCompanyChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildToolResult = shared.QueryResultDocumentDocumentCompanyChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildToolCall = shared.QueryResultDocumentDocumentCompanyChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildToolResult = shared.QueryResultDocumentDocumentCompanyChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildToolResultOutputUnion = shared.QueryResultDocumentDocumentCompanyChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildTraceMessage = shared.QueryResultDocumentDocumentCompanyChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentCompanyChildUtterance = shared.QueryResultDocumentDocumentCompanyChildUtterance

// A CRM deal/opportunity record (ENG-2476/D10).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDeal = shared.QueryResultDocumentDocumentDeal

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildUnion = shared.QueryResultDocumentDocumentDealChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildBlob = shared.QueryResultDocumentDocumentDealChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCallout = shared.QueryResultDocumentDocumentDealChildCallout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildUnion = shared.QueryResultDocumentDocumentDealChildCalloutChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildBlob = shared.QueryResultDocumentDocumentDealChildCalloutChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildCode = shared.QueryResultDocumentDocumentDealChildCalloutChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildComment = shared.QueryResultDocumentDocumentDealChildCalloutChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildDivider = shared.QueryResultDocumentDocumentDealChildCalloutChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildImage = shared.QueryResultDocumentDocumentDealChildCalloutChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildLink = shared.QueryResultDocumentDocumentDealChildCalloutChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildLineBreak = shared.QueryResultDocumentDocumentDealChildCalloutChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildText = shared.QueryResultDocumentDocumentDealChildCalloutChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildToolCall = shared.QueryResultDocumentDocumentDealChildCalloutChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildToolResult = shared.QueryResultDocumentDocumentDealChildCalloutChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildCalloutChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCalloutChildTraceMessage = shared.QueryResultDocumentDocumentDealChildCalloutChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunk = shared.QueryResultDocumentDocumentDealChildChunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildUnion = shared.QueryResultDocumentDocumentDealChildChunkChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildBlob = shared.QueryResultDocumentDocumentDealChildChunkChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildCode = shared.QueryResultDocumentDocumentDealChildChunkChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildComment = shared.QueryResultDocumentDocumentDealChildChunkChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildDivider = shared.QueryResultDocumentDocumentDealChildChunkChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildImage = shared.QueryResultDocumentDocumentDealChildChunkChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildLink = shared.QueryResultDocumentDocumentDealChildChunkChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildLineBreak = shared.QueryResultDocumentDocumentDealChildChunkChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildText = shared.QueryResultDocumentDocumentDealChildChunkChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildToolCall = shared.QueryResultDocumentDocumentDealChildChunkChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildToolResult = shared.QueryResultDocumentDocumentDealChildChunkChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildChunkChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildChunkChildTraceMessage = shared.QueryResultDocumentDocumentDealChildChunkChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildCode = shared.QueryResultDocumentDocumentDealChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildComment = shared.QueryResultDocumentDocumentDealChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildDivider = shared.QueryResultDocumentDocumentDealChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquation = shared.QueryResultDocumentDocumentDealChildEquation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildUnion = shared.QueryResultDocumentDocumentDealChildEquationChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildBlob = shared.QueryResultDocumentDocumentDealChildEquationChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildCode = shared.QueryResultDocumentDocumentDealChildEquationChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildComment = shared.QueryResultDocumentDocumentDealChildEquationChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildDivider = shared.QueryResultDocumentDocumentDealChildEquationChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildImage = shared.QueryResultDocumentDocumentDealChildEquationChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildLink = shared.QueryResultDocumentDocumentDealChildEquationChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildLineBreak = shared.QueryResultDocumentDocumentDealChildEquationChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildText = shared.QueryResultDocumentDocumentDealChildEquationChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildToolCall = shared.QueryResultDocumentDocumentDealChildEquationChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildToolResult = shared.QueryResultDocumentDocumentDealChildEquationChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildEquationChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildEquationChildTraceMessage = shared.QueryResultDocumentDocumentDealChildEquationChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnote = shared.QueryResultDocumentDocumentDealChildFootnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildUnion = shared.QueryResultDocumentDocumentDealChildFootnoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildBlob = shared.QueryResultDocumentDocumentDealChildFootnoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildCode = shared.QueryResultDocumentDocumentDealChildFootnoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildComment = shared.QueryResultDocumentDocumentDealChildFootnoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildDivider = shared.QueryResultDocumentDocumentDealChildFootnoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildImage = shared.QueryResultDocumentDocumentDealChildFootnoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildLink = shared.QueryResultDocumentDocumentDealChildFootnoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildLineBreak = shared.QueryResultDocumentDocumentDealChildFootnoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildText = shared.QueryResultDocumentDocumentDealChildFootnoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildToolCall = shared.QueryResultDocumentDocumentDealChildFootnoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildToolResult = shared.QueryResultDocumentDocumentDealChildFootnoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildFootnoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildFootnoteChildTraceMessage = shared.QueryResultDocumentDocumentDealChildFootnoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeading = shared.QueryResultDocumentDocumentDealChildHeading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildUnion = shared.QueryResultDocumentDocumentDealChildHeadingChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildBlob = shared.QueryResultDocumentDocumentDealChildHeadingChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildCode = shared.QueryResultDocumentDocumentDealChildHeadingChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildComment = shared.QueryResultDocumentDocumentDealChildHeadingChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildDivider = shared.QueryResultDocumentDocumentDealChildHeadingChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildImage = shared.QueryResultDocumentDocumentDealChildHeadingChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildLink = shared.QueryResultDocumentDocumentDealChildHeadingChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildLineBreak = shared.QueryResultDocumentDocumentDealChildHeadingChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildText = shared.QueryResultDocumentDocumentDealChildHeadingChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildToolCall = shared.QueryResultDocumentDocumentDealChildHeadingChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildToolResult = shared.QueryResultDocumentDocumentDealChildHeadingChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildHeadingChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildHeadingChildTraceMessage = shared.QueryResultDocumentDocumentDealChildHeadingChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildImage = shared.QueryResultDocumentDocumentDealChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildLink = shared.QueryResultDocumentDocumentDealChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildLineBreak = shared.QueryResultDocumentDocumentDealChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildList = shared.QueryResultDocumentDocumentDealChildList

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItem = shared.QueryResultDocumentDocumentDealChildListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildUnion = shared.QueryResultDocumentDocumentDealChildListItemChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildBlob = shared.QueryResultDocumentDocumentDealChildListItemChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildCode = shared.QueryResultDocumentDocumentDealChildListItemChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildComment = shared.QueryResultDocumentDocumentDealChildListItemChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildDivider = shared.QueryResultDocumentDocumentDealChildListItemChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildImage = shared.QueryResultDocumentDocumentDealChildListItemChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildLink = shared.QueryResultDocumentDocumentDealChildListItemChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildLineBreak = shared.QueryResultDocumentDocumentDealChildListItemChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildText = shared.QueryResultDocumentDocumentDealChildListItemChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildToolCall = shared.QueryResultDocumentDocumentDealChildListItemChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildToolResult = shared.QueryResultDocumentDocumentDealChildListItemChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildListItemChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildListItemChildTraceMessage = shared.QueryResultDocumentDocumentDealChildListItemChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraph = shared.QueryResultDocumentDocumentDealChildParagraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildUnion = shared.QueryResultDocumentDocumentDealChildParagraphChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildBlob = shared.QueryResultDocumentDocumentDealChildParagraphChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildCode = shared.QueryResultDocumentDocumentDealChildParagraphChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildComment = shared.QueryResultDocumentDocumentDealChildParagraphChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildDivider = shared.QueryResultDocumentDocumentDealChildParagraphChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildImage = shared.QueryResultDocumentDocumentDealChildParagraphChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildLink = shared.QueryResultDocumentDocumentDealChildParagraphChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildLineBreak = shared.QueryResultDocumentDocumentDealChildParagraphChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildText = shared.QueryResultDocumentDocumentDealChildParagraphChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildToolCall = shared.QueryResultDocumentDocumentDealChildParagraphChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildToolResult = shared.QueryResultDocumentDocumentDealChildParagraphChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildParagraphChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildParagraphChildTraceMessage = shared.QueryResultDocumentDocumentDealChildParagraphChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuote = shared.QueryResultDocumentDocumentDealChildQuote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildUnion = shared.QueryResultDocumentDocumentDealChildQuoteChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildBlob = shared.QueryResultDocumentDocumentDealChildQuoteChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildCode = shared.QueryResultDocumentDocumentDealChildQuoteChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildComment = shared.QueryResultDocumentDocumentDealChildQuoteChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildDivider = shared.QueryResultDocumentDocumentDealChildQuoteChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildImage = shared.QueryResultDocumentDocumentDealChildQuoteChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildLink = shared.QueryResultDocumentDocumentDealChildQuoteChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildLineBreak = shared.QueryResultDocumentDocumentDealChildQuoteChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildText = shared.QueryResultDocumentDocumentDealChildQuoteChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildToolCall = shared.QueryResultDocumentDocumentDealChildQuoteChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildToolResult = shared.QueryResultDocumentDocumentDealChildQuoteChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildQuoteChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildQuoteChildTraceMessage = shared.QueryResultDocumentDocumentDealChildQuoteChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTable = shared.QueryResultDocumentDocumentDealChildTable

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCell = shared.QueryResultDocumentDocumentDealChildTableCell

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildUnion = shared.QueryResultDocumentDocumentDealChildTableCellChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildBlob = shared.QueryResultDocumentDocumentDealChildTableCellChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildCode = shared.QueryResultDocumentDocumentDealChildTableCellChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildComment = shared.QueryResultDocumentDocumentDealChildTableCellChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildDivider = shared.QueryResultDocumentDocumentDealChildTableCellChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildImage = shared.QueryResultDocumentDocumentDealChildTableCellChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildLink = shared.QueryResultDocumentDocumentDealChildTableCellChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildLineBreak = shared.QueryResultDocumentDocumentDealChildTableCellChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildText = shared.QueryResultDocumentDocumentDealChildTableCellChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildToolCall = shared.QueryResultDocumentDocumentDealChildTableCellChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildToolResult = shared.QueryResultDocumentDocumentDealChildTableCellChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildTableCellChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableCellChildTraceMessage = shared.QueryResultDocumentDocumentDealChildTableCellChildTraceMessage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTableRow = shared.QueryResultDocumentDocumentDealChildTableRow

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildText = shared.QueryResultDocumentDocumentDealChildText

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodo = shared.QueryResultDocumentDocumentDealChildTodo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildUnion = shared.QueryResultDocumentDocumentDealChildTodoChildUnion

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildBlob = shared.QueryResultDocumentDocumentDealChildTodoChildBlob

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildCode = shared.QueryResultDocumentDocumentDealChildTodoChildCode

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildComment = shared.QueryResultDocumentDocumentDealChildTodoChildComment

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildDivider = shared.QueryResultDocumentDocumentDealChildTodoChildDivider

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildImage = shared.QueryResultDocumentDocumentDealChildTodoChildImage

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildLink = shared.QueryResultDocumentDocumentDealChildTodoChildLink

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildLineBreak = shared.QueryResultDocumentDocumentDealChildTodoChildLineBreak

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildText = shared.QueryResultDocumentDocumentDealChildTodoChildText

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildToolCall = shared.QueryResultDocumentDocumentDealChildTodoChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildToolResult = shared.QueryResultDocumentDocumentDealChildTodoChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildTodoChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTodoChildTraceMessage = shared.QueryResultDocumentDocumentDealChildTodoChildTraceMessage

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildToolCall = shared.QueryResultDocumentDocumentDealChildToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildToolResult = shared.QueryResultDocumentDocumentDealChildToolResult

// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildToolResultOutputUnion = shared.QueryResultDocumentDocumentDealChildToolResultOutputUnion

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildTraceMessage = shared.QueryResultDocumentDocumentDealChildTraceMessage

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type QueryResultDocumentDocumentDealChildUtterance = shared.QueryResultDocumentDocumentDealChildUtterance

// Auditability record attached to an agentic answer.
//
// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
// failed_sources) are derived from in-memory loop state, but `entities` costs one
// indexed DB lookup, so the whole record is only built on request.
//
// This is an alias to an internal type.
type QueryResultProvenance = shared.QueryResultProvenance

// A canonical entity referenced by the answer's source documents.
//
// This is an alias to an internal type.
type QueryResultProvenanceEntity = shared.QueryResultProvenanceEntity

// A source document that informed the final answer (the post-rank result set).
//
// This is an alias to an internal type.
type QueryResultProvenanceSource = shared.QueryResultProvenanceSource

// One tool invocation in the agent's search trajectory (audit trail).
//
// This is an alias to an internal type.
type QueryResultProvenanceStep = shared.QueryResultProvenanceStep
