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

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type Blob = shared.Blob

// This is an alias to an internal type.
type BlobType = shared.BlobType

// Equals "blob"
const BlobTypeBlob = shared.BlobTypeBlob

// This is an alias to an internal type.
type Callout = shared.Callout

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type CalloutChildUnion = shared.CalloutChildUnion

// This is an alias to an internal type.
type CalloutType = shared.CalloutType

// Equals "callout"
const CalloutTypeCallout = shared.CalloutTypeCallout

// This is an alias to an internal type.
type Chunk = shared.Chunk

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type ChunkChildUnion = shared.ChunkChildUnion

// This is an alias to an internal type.
type ChunkType = shared.ChunkType

// Equals "chunk"
const ChunkTypeChunk = shared.ChunkTypeChunk

// This is an alias to an internal type.
type Code = shared.Code

// This is an alias to an internal type.
type CodeType = shared.CodeType

// Equals "code"
const CodeTypeCode = shared.CodeTypeCode

// This is an alias to an internal type.
type Comment = shared.Comment

// This is an alias to an internal type.
type CommentType = shared.CommentType

// Equals "comment"
const CommentTypeComment = shared.CommentTypeComment

// A CRM company/account record (ENG-2476/D10).
//
// This is an alias to an internal type.
type Company = shared.Company

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type CompanyChildUnion = shared.CompanyChildUnion

// This is an alias to an internal type.
type CompanyType = shared.CompanyType

// Equals "company"
const CompanyTypeCompany = shared.CompanyTypeCompany

// This is an alias to an internal type.
type Conversation = shared.Conversation

// This is an alias to an internal type.
type ConversationType = shared.ConversationType

// Equals "conversation"
const ConversationTypeConversation = shared.ConversationTypeConversation

// A CRM deal/opportunity record (ENG-2476/D10).
//
// This is an alias to an internal type.
type Deal = shared.Deal

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type DealChildUnion = shared.DealChildUnion

// This is an alias to an internal type.
type DealType = shared.DealType

// Equals "deal"
const DealTypeDeal = shared.DealTypeDeal

// This is an alias to an internal type.
type Divider = shared.Divider

// This is an alias to an internal type.
type DividerType = shared.DividerType

// Equals "divider"
const DividerTypeDivider = shared.DividerTypeDivider

// This is an alias to an internal type.
type Document = shared.Document

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type DocumentChildUnion = shared.DocumentChildUnion

// This is an alias to an internal type.
type DocumentType = shared.DocumentType

// Equals "document"
const DocumentTypeDocument = shared.DocumentTypeDocument

// This is an alias to an internal type.
type Equation = shared.Equation

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type EquationChildUnion = shared.EquationChildUnion

// This is an alias to an internal type.
type EquationType = shared.EquationType

// Equals "equation"
const EquationTypeEquation = shared.EquationTypeEquation

// This is an alias to an internal type.
type Event = shared.Event

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type EventChildUnion = shared.EventChildUnion

// This is an alias to an internal type.
type EventType = shared.EventType

// Equals "event"
const EventTypeEvent = shared.EventTypeEvent

// This is an alias to an internal type.
type File = shared.File

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type FileChildUnion = shared.FileChildUnion

// This is an alias to an internal type.
type FileType = shared.FileType

// Equals "file"
const FileTypeFile = shared.FileTypeFile

// This is an alias to an internal type.
type Footnote = shared.Footnote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type FootnoteChildUnion = shared.FootnoteChildUnion

// This is an alias to an internal type.
type FootnoteType = shared.FootnoteType

// Equals "footnote"
const FootnoteTypeFootnote = shared.FootnoteTypeFootnote

// This is an alias to an internal type.
type Heading = shared.Heading

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type HeadingChildUnion = shared.HeadingChildUnion

// This is an alias to an internal type.
type HeadingType = shared.HeadingType

// Equals "heading"
const HeadingTypeHeading = shared.HeadingTypeHeading

// This is an alias to an internal type.
type Image = shared.Image

// This is an alias to an internal type.
type ImageType = shared.ImageType

// Equals "image"
const ImageTypeImage = shared.ImageTypeImage

// This is an alias to an internal type.
type LineBreak = shared.LineBreak

// This is an alias to an internal type.
type LineBreakType = shared.LineBreakType

// Equals "line_break"
const LineBreakTypeLineBreak = shared.LineBreakTypeLineBreak

// This is an alias to an internal type.
type Link = shared.Link

// This is an alias to an internal type.
type LinkType = shared.LinkType

// Equals "link"
const LinkTypeLink = shared.LinkTypeLink

// This is an alias to an internal type.
type List = shared.List

// This is an alias to an internal type.
type ListChildUnion = shared.ListChildUnion

// This is an alias to an internal type.
type ListType = shared.ListType

// Equals "list"
const ListTypeList = shared.ListTypeList

// This is an alias to an internal type.
type ListItem = shared.ListItem

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type ListItemChildUnion = shared.ListItemChildUnion

// This is an alias to an internal type.
type ListItemType = shared.ListItemType

// Equals "list_item"
const ListItemTypeListItem = shared.ListItemTypeListItem

// This is an alias to an internal type.
type Message = shared.Message

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type MessageChildUnion = shared.MessageChildUnion

// This is an alias to an internal type.
type MessageType = shared.MessageType

// Equals "message"
const MessageTypeMessage = shared.MessageTypeMessage

// Per-block annotations carried by any Hyperdoc node (ENG-1390).
//
// Out-of-band annotations that travel with a block but aren't part of its content:
// provenance (`sources`) and human edit attribution (`edited_by`). New annotation
// types get added here as typed fields as the need arises.
//
// Empty by default. Because `Node.model_dump` forces `exclude_none=True`, an unset
// `metadata` (None) is dropped from serialization entirely, and within a populated
// `Metadata` only the set keys survive.
//
// This is an alias to an internal type.
type Metadata = shared.Metadata

// A reference to a memory/chunk that a block's content is grounded in (ENG-1390).
//
// Chunks are the unit persisted to the DB — extracted memories become chunks when
// indexed — so `chunk_id` is the stable pointer back to the source. `resource_id`
// and `source` locate the originating document; `score` carries optional retrieval
// relevance. Kept deliberately self-contained (plain `str` for `source` rather
// than the `DocumentProviders` enum) so the hyperdoc format stays free of
// app-layer imports.
//
// This is an alias to an internal type.
type MetadataSource = shared.MetadataSource

// This is an alias to an internal type.
type Paragraph = shared.Paragraph

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type ParagraphChildUnion = shared.ParagraphChildUnion

// This is an alias to an internal type.
type ParagraphType = shared.ParagraphType

// Equals "paragraph"
const ParagraphTypeParagraph = shared.ParagraphTypeParagraph

// This is an alias to an internal type.
type Person = shared.Person

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type PersonChildUnion = shared.PersonChildUnion

// This is an alias to an internal type.
type PersonType = shared.PersonType

// Equals "person"
const PersonTypePerson = shared.PersonTypePerson

// Auditability record attached to an agentic answer.
//
// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
// failed_sources) are derived from in-memory loop state, but `entities` costs one
// indexed DB lookup, so the whole record is only built on request.
//
// This is an alias to an internal type.
type Provenance = shared.Provenance

// A canonical entity referenced by the answer's source documents.
//
// This is an alias to an internal type.
type ProvenanceEntity = shared.ProvenanceEntity

// A source document that informed the final answer (the post-rank result set).
//
// This is an alias to an internal type.
type ProvenanceSource = shared.ProvenanceSource

// This is an alias to an internal type.
type ProvenanceSourceSource = shared.ProvenanceSourceSource

// Equals "reddit"
const ProvenanceSourceSourceReddit = shared.ProvenanceSourceSourceReddit

// Equals "notion"
const ProvenanceSourceSourceNotion = shared.ProvenanceSourceSourceNotion

// Equals "slack"
const ProvenanceSourceSourceSlack = shared.ProvenanceSourceSourceSlack

// Equals "google_calendar"
const ProvenanceSourceSourceGoogleCalendar = shared.ProvenanceSourceSourceGoogleCalendar

// Equals "google_mail"
const ProvenanceSourceSourceGoogleMail = shared.ProvenanceSourceSourceGoogleMail

// Equals "box"
const ProvenanceSourceSourceBox = shared.ProvenanceSourceSourceBox

// Equals "dropbox"
const ProvenanceSourceSourceDropbox = shared.ProvenanceSourceSourceDropbox

// Equals "github"
const ProvenanceSourceSourceGitHub = shared.ProvenanceSourceSourceGitHub

// Equals "google_drive"
const ProvenanceSourceSourceGoogleDrive = shared.ProvenanceSourceSourceGoogleDrive

// Equals "vault"
const ProvenanceSourceSourceVault = shared.ProvenanceSourceSourceVault

// Equals "web_crawler"
const ProvenanceSourceSourceWebCrawler = shared.ProvenanceSourceSourceWebCrawler

// Equals "trace"
const ProvenanceSourceSourceTrace = shared.ProvenanceSourceSourceTrace

// Equals "microsoft_teams"
const ProvenanceSourceSourceMicrosoftTeams = shared.ProvenanceSourceSourceMicrosoftTeams

// Equals "gmail_actions"
const ProvenanceSourceSourceGmailActions = shared.ProvenanceSourceSourceGmailActions

// Equals "granola"
const ProvenanceSourceSourceGranola = shared.ProvenanceSourceSourceGranola

// Equals "fathom"
const ProvenanceSourceSourceFathom = shared.ProvenanceSourceSourceFathom

// Equals "fireflies"
const ProvenanceSourceSourceFireflies = shared.ProvenanceSourceSourceFireflies

// Equals "linear"
const ProvenanceSourceSourceLinear = shared.ProvenanceSourceSourceLinear

// Equals "hubspot"
const ProvenanceSourceSourceHubspot = shared.ProvenanceSourceSourceHubspot

// Equals "salesforce"
const ProvenanceSourceSourceSalesforce = shared.ProvenanceSourceSourceSalesforce

// Equals "coda"
const ProvenanceSourceSourceCoda = shared.ProvenanceSourceSourceCoda

// Equals "lightfield"
const ProvenanceSourceSourceLightfield = shared.ProvenanceSourceSourceLightfield

// Equals "gong"
const ProvenanceSourceSourceGong = shared.ProvenanceSourceSourceGong

// One tool invocation in the agent's search trajectory (audit trail).
//
// This is an alias to an internal type.
type ProvenanceStep = shared.ProvenanceStep

// This is an alias to an internal type.
type QueryResult = shared.QueryResult

// This is an alias to an internal type.
type Quote = shared.Quote

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type QuoteChildUnion = shared.QuoteChildUnion

// This is an alias to an internal type.
type QuoteType = shared.QuoteType

// Equals "quote"
const QuoteTypeQuote = shared.QuoteTypeQuote

// A `DocumentResponse` plus the query-path fields a `ScoredDocument` carries
// (ENG-2479): relevance score, matched highlights, and the concatenated summary of
// those highlights.
//
// This is an alias to an internal type.
type ScoredDocumentResponse = shared.ScoredDocumentResponse

// The full hyperdoc tree. Switch on `type` for the document frame and recurse
// `children` for the body — see the `<Hyperdoc />` renderer.
//
// This is an alias to an internal type.
type ScoredDocumentResponseDocumentUnion = shared.ScoredDocumentResponseDocumentUnion

// This is an alias to an internal type.
type ScoredDocumentResponseSource = shared.ScoredDocumentResponseSource

// Equals "reddit"
const ScoredDocumentResponseSourceReddit = shared.ScoredDocumentResponseSourceReddit

// Equals "notion"
const ScoredDocumentResponseSourceNotion = shared.ScoredDocumentResponseSourceNotion

// Equals "slack"
const ScoredDocumentResponseSourceSlack = shared.ScoredDocumentResponseSourceSlack

// Equals "google_calendar"
const ScoredDocumentResponseSourceGoogleCalendar = shared.ScoredDocumentResponseSourceGoogleCalendar

// Equals "google_mail"
const ScoredDocumentResponseSourceGoogleMail = shared.ScoredDocumentResponseSourceGoogleMail

// Equals "box"
const ScoredDocumentResponseSourceBox = shared.ScoredDocumentResponseSourceBox

// Equals "dropbox"
const ScoredDocumentResponseSourceDropbox = shared.ScoredDocumentResponseSourceDropbox

// Equals "github"
const ScoredDocumentResponseSourceGitHub = shared.ScoredDocumentResponseSourceGitHub

// Equals "google_drive"
const ScoredDocumentResponseSourceGoogleDrive = shared.ScoredDocumentResponseSourceGoogleDrive

// Equals "vault"
const ScoredDocumentResponseSourceVault = shared.ScoredDocumentResponseSourceVault

// Equals "web_crawler"
const ScoredDocumentResponseSourceWebCrawler = shared.ScoredDocumentResponseSourceWebCrawler

// Equals "trace"
const ScoredDocumentResponseSourceTrace = shared.ScoredDocumentResponseSourceTrace

// Equals "microsoft_teams"
const ScoredDocumentResponseSourceMicrosoftTeams = shared.ScoredDocumentResponseSourceMicrosoftTeams

// Equals "gmail_actions"
const ScoredDocumentResponseSourceGmailActions = shared.ScoredDocumentResponseSourceGmailActions

// Equals "granola"
const ScoredDocumentResponseSourceGranola = shared.ScoredDocumentResponseSourceGranola

// Equals "fathom"
const ScoredDocumentResponseSourceFathom = shared.ScoredDocumentResponseSourceFathom

// Equals "fireflies"
const ScoredDocumentResponseSourceFireflies = shared.ScoredDocumentResponseSourceFireflies

// Equals "linear"
const ScoredDocumentResponseSourceLinear = shared.ScoredDocumentResponseSourceLinear

// Equals "hubspot"
const ScoredDocumentResponseSourceHubspot = shared.ScoredDocumentResponseSourceHubspot

// Equals "salesforce"
const ScoredDocumentResponseSourceSalesforce = shared.ScoredDocumentResponseSourceSalesforce

// Equals "coda"
const ScoredDocumentResponseSourceCoda = shared.ScoredDocumentResponseSourceCoda

// Equals "lightfield"
const ScoredDocumentResponseSourceLightfield = shared.ScoredDocumentResponseSourceLightfield

// Equals "gong"
const ScoredDocumentResponseSourceGong = shared.ScoredDocumentResponseSourceGong

// Indexing status of the document.
//
// This is an alias to an internal type.
type ScoredDocumentResponseStatus = shared.ScoredDocumentResponseStatus

// Equals "pending"
const ScoredDocumentResponseStatusPending = shared.ScoredDocumentResponseStatusPending

// Equals "processing"
const ScoredDocumentResponseStatusProcessing = shared.ScoredDocumentResponseStatusProcessing

// Equals "completed"
const ScoredDocumentResponseStatusCompleted = shared.ScoredDocumentResponseStatusCompleted

// Equals "failed"
const ScoredDocumentResponseStatusFailed = shared.ScoredDocumentResponseStatusFailed

// Equals "pending_review"
const ScoredDocumentResponseStatusPendingReview = shared.ScoredDocumentResponseStatusPendingReview

// Equals "skipped"
const ScoredDocumentResponseStatusSkipped = shared.ScoredDocumentResponseStatusSkipped

// This is an alias to an internal type.
type Table = shared.Table

// This is an alias to an internal type.
type TableType = shared.TableType

// Equals "table"
const TableTypeTable = shared.TableTypeTable

// This is an alias to an internal type.
type TableCell = shared.TableCell

// This is an alias to an internal type.
type TableCellAlign = shared.TableCellAlign

// Equals "left"
const TableCellAlignLeft = shared.TableCellAlignLeft

// Equals "center"
const TableCellAlignCenter = shared.TableCellAlignCenter

// Equals "right"
const TableCellAlignRight = shared.TableCellAlignRight

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type TableCellChildUnion = shared.TableCellChildUnion

// This is an alias to an internal type.
type TableCellType = shared.TableCellType

// Equals "table_cell"
const TableCellTypeTableCell = shared.TableCellTypeTableCell

// This is an alias to an internal type.
type TableRow = shared.TableRow

// This is an alias to an internal type.
type TableRowType = shared.TableRowType

// Equals "table_row"
const TableRowTypeTableRow = shared.TableRowTypeTableRow

// This is an alias to an internal type.
type Task = shared.Task

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type TaskChildUnion = shared.TaskChildUnion

// This is an alias to an internal type.
type TaskPriority = shared.TaskPriority

// Equals "urgent"
const TaskPriorityUrgent = shared.TaskPriorityUrgent

// Equals "high"
const TaskPriorityHigh = shared.TaskPriorityHigh

// Equals "medium"
const TaskPriorityMedium = shared.TaskPriorityMedium

// Equals "low"
const TaskPriorityLow = shared.TaskPriorityLow

// This is an alias to an internal type.
type TaskStatus = shared.TaskStatus

// Equals "completed"
const TaskStatusCompleted = shared.TaskStatusCompleted

// Equals "not_started"
const TaskStatusNotStarted = shared.TaskStatusNotStarted

// Equals "in_progress"
const TaskStatusInProgress = shared.TaskStatusInProgress

// Equals "cancelled"
const TaskStatusCancelled = shared.TaskStatusCancelled

// This is an alias to an internal type.
type TaskType = shared.TaskType

// Equals "task"
const TaskTypeTask = shared.TaskTypeTask

// This is an alias to an internal type.
type Text = shared.Text

// This is an alias to an internal type.
type TextType = shared.TextType

// Equals "text"
const TextTypeText = shared.TextTypeText

// This is an alias to an internal type.
type ToDo = shared.ToDo

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type ToDoChildUnion = shared.ToDoChildUnion

// This is an alias to an internal type.
type ToDoType = shared.ToDoType

// Equals "todo"
const ToDoTypeTodo = shared.ToDoTypeTodo

// A tool/function call made by the assistant.
//
// This is an alias to an internal type.
type ToolCall = shared.ToolCall

// This is an alias to an internal type.
type ToolCallType = shared.ToolCallType

// Equals "tool_call"
const ToolCallTypeToolCall = shared.ToolCallTypeToolCall

// The result of a tool call.
//
// This is an alias to an internal type.
type ToolResult = shared.ToolResult

// This is an alias to an internal type.
type ToolResultOutputUnion = shared.ToolResultOutputUnion

// This is an alias to an internal type.
type ToolResultType = shared.ToolResultType

// Equals "tool_result"
const ToolResultTypeToolResult = shared.ToolResultTypeToolResult

// An agent trace/transcript containing a sequence of steps.
//
// Steps can be TraceMessage (user/assistant messages or thinking), ToolCall
// (function calls), or ToolResult (tool responses).
//
// This is an alias to an internal type.
type Trace = shared.Trace

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type TraceChildUnion = shared.TraceChildUnion

// This is an alias to an internal type.
type TraceType = shared.TraceType

// Equals "trace"
const TraceTypeTrace = shared.TraceTypeTrace

// A message in an agent trace (user message, assistant message, or thinking).
//
// This is an alias to an internal type.
type TraceMessage = shared.TraceMessage

// This is an alias to an internal type.
type TraceMessageMessageType = shared.TraceMessageMessageType

// Equals "message"
const TraceMessageMessageTypeMessage = shared.TraceMessageMessageTypeMessage

// Equals "thinking"
const TraceMessageMessageTypeThinking = shared.TraceMessageMessageTypeThinking

// This is an alias to an internal type.
type TraceMessageRole = shared.TraceMessageRole

// Equals "user"
const TraceMessageRoleUser = shared.TraceMessageRoleUser

// Equals "assistant"
const TraceMessageRoleAssistant = shared.TraceMessageRoleAssistant

// This is an alias to an internal type.
type TraceMessageType = shared.TraceMessageType

// Equals "trace_message"
const TraceMessageTypeTraceMessage = shared.TraceMessageTypeTraceMessage

// A time-anchored, speaker-attributed transcript — meetings, calls (ENG-2476/D10;
// mirrors the Trace+TraceStep precedent).
//
// Utterance timestamps are relative offsets from `started_at`, which is the
// absolute wall-clock anchor.
//
// This is an alias to an internal type.
type Transcript = shared.Transcript

// This is an alias to an internal type.
type TranscriptType = shared.TranscriptType

// Equals "transcript"
const TranscriptTypeTranscript = shared.TranscriptTypeTranscript

// A speaker-attributed segment of a transcript (ENG-2476/D10).
//
// "Utterance" is the standard name for this across transcription providers
// (AssemblyAI, Deepgram, Rev). Timestamps are relative offsets in seconds —
// provider-native; absolute times derive from `Transcript.started_at`.
//
// This is an alias to an internal type.
type Utterance = shared.Utterance

// This is an alias to an internal type.
type UtteranceType = shared.UtteranceType

// Equals "utterance"
const UtteranceTypeUtterance = shared.UtteranceTypeUtterance

// This is an alias to an internal type.
type Website = shared.Website

// Represents embedded binary data using data URI scheme.
//
// Format: data:[<media type>][;base64],<data> Example:
// data:text/html;base64,PGh0bWw+...
//
// This is an alias to an internal type.
type WebsiteChildUnion = shared.WebsiteChildUnion

// This is an alias to an internal type.
type WebsiteType = shared.WebsiteType

// Equals "website"
const WebsiteTypeWebsite = shared.WebsiteTypeWebsite
