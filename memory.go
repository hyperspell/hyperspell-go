// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

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
	"time"

	"github.com/hyperspell/hyperspell-go/internal/apiform"
	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/apiquery"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/pagination"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
	"github.com/hyperspell/hyperspell-go/shared"
)

// MemoryService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMemoryService] method instead.
type MemoryService struct {
	options []option.RequestOption
}

// NewMemoryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMemoryService(opts ...option.RequestOption) (r MemoryService) {
	r = MemoryService{}
	r.options = opts
	return
}

// Updates an existing document in the index. You can update the text, collection,
// title, and metadata. The document must already exist or a 404 will be returned.
// This works for documents from any source (vault, slack, gmail, etc.).
//
// To remove a collection, set it to null explicitly.
func (r *MemoryService) Update(ctx context.Context, resourceID string, params MemoryUpdateParams, opts ...option.RequestOption) (res *MemoryStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("memories/update/%v/%s", params.Source, url.PathEscape(resourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// This endpoint allows you to paginate through all documents in the index. You can
// filter the documents by title, date, metadata, etc.
func (r *MemoryService) List(ctx context.Context, query MemoryListParams, opts ...option.RequestOption) (res *pagination.CursorPage[MemoryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "memories/list"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// This endpoint allows you to paginate through all documents in the index. You can
// filter the documents by title, date, metadata, etc.
func (r *MemoryService) ListAutoPaging(ctx context.Context, query MemoryListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[MemoryListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a memory and its associated chunks from the index.
//
// This removes the memory completely from the vector index and database. The
// operation deletes:
//
//  1. All chunks associated with the resource (including embeddings)
//  2. The documents row AND any legacy resources rows sharing the identity —
//     leaving either one behind would resurrect the memory through the double-read
//     path (ENG-2477).
//
// Args: source: The document provider (e.g., gmail, notion, vault) resource_id:
// The unique identifier of the resource to delete api_token: Authentication token
//
// Returns: MemoryDeletionResponse with deletion details
//
// Raises: DocumentNotFound: If the resource doesn't exist or user doesn't have
// access
func (r *MemoryService) Delete(ctx context.Context, resourceID string, body MemoryDeleteParams, opts ...option.RequestOption) (res *MemoryDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("memories/delete/%v/%s", body.Source, url.PathEscape(resourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Adds an arbitrary document to the index. This can be any text, email, call
// transcript, etc. The document will be processed and made available for querying
// once the processing is complete.
func (r *MemoryService) Add(ctx context.Context, body MemoryAddParams, opts ...option.RequestOption) (res *MemoryStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Adds multiple documents to the index in a single request.
//
// All items are validated before any database operations occur. If any item fails
// validation, the entire batch is rejected with a 422 error detailing which items
// failed and why.
//
// Maximum 100 items per request. Each item follows the same schema as the
// single-item /memories/add endpoint.
func (r *MemoryService) AddBulk(ctx context.Context, body MemoryAddBulkParams, opts ...option.RequestOption) (res *MemoryAddBulkResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/add/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a document by provider and resource_id, as a document-shaped response
// carrying the full hyperdoc tree (ENG-2479 Phase 4).
func (r *MemoryService) Get(ctx context.Context, resourceID string, query MemoryGetParams, opts ...option.RequestOption) (res *MemoryGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("memories/get/%v/%s", query.Source, url.PathEscape(resourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieves documents matching the query.
func (r *MemoryService) Search(ctx context.Context, body MemorySearchParams, opts ...option.RequestOption) (res *shared.QueryResult, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// This endpoint shows the indexing progress of documents, both by provider and
// total.
func (r *MemoryService) Status(ctx context.Context, opts ...option.RequestOption) (res *MemoryStatusResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/status"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// This endpoint will upload a file to the index and return a resource_id. The file
// will be processed in the background and the memory will be available for
// querying once the processing is complete. You can use the `resource_id` to query
// the memory later, and check the status of the memory.
func (r *MemoryService) Upload(ctx context.Context, body MemoryUploadParams, opts ...option.RequestOption) (res *MemoryStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/upload"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MemoryStatus struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryStatusSource `json:"source" api:"required"`
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status MemoryStatusStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID  respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatus) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryStatusSource string

const (
	MemoryStatusSourceReddit         MemoryStatusSource = "reddit"
	MemoryStatusSourceNotion         MemoryStatusSource = "notion"
	MemoryStatusSourceSlack          MemoryStatusSource = "slack"
	MemoryStatusSourceGoogleCalendar MemoryStatusSource = "google_calendar"
	MemoryStatusSourceGoogleMail     MemoryStatusSource = "google_mail"
	MemoryStatusSourceBox            MemoryStatusSource = "box"
	MemoryStatusSourceDropbox        MemoryStatusSource = "dropbox"
	MemoryStatusSourceGitHub         MemoryStatusSource = "github"
	MemoryStatusSourceGoogleDrive    MemoryStatusSource = "google_drive"
	MemoryStatusSourceVault          MemoryStatusSource = "vault"
	MemoryStatusSourceWebCrawler     MemoryStatusSource = "web_crawler"
	MemoryStatusSourceTrace          MemoryStatusSource = "trace"
	MemoryStatusSourceMicrosoftTeams MemoryStatusSource = "microsoft_teams"
	MemoryStatusSourceGmailActions   MemoryStatusSource = "gmail_actions"
	MemoryStatusSourceGranola        MemoryStatusSource = "granola"
	MemoryStatusSourceFathom         MemoryStatusSource = "fathom"
	MemoryStatusSourceFireflies      MemoryStatusSource = "fireflies"
	MemoryStatusSourceLinear         MemoryStatusSource = "linear"
	MemoryStatusSourceHubspot        MemoryStatusSource = "hubspot"
	MemoryStatusSourceSalesforce     MemoryStatusSource = "salesforce"
	MemoryStatusSourceCoda           MemoryStatusSource = "coda"
	MemoryStatusSourceLightfield     MemoryStatusSource = "lightfield"
	MemoryStatusSourceGong           MemoryStatusSource = "gong"
)

type MemoryStatusStatus string

const (
	MemoryStatusStatusPending       MemoryStatusStatus = "pending"
	MemoryStatusStatusProcessing    MemoryStatusStatus = "processing"
	MemoryStatusStatusCompleted     MemoryStatusStatus = "completed"
	MemoryStatusStatusFailed        MemoryStatusStatus = "failed"
	MemoryStatusStatusPendingReview MemoryStatusStatus = "pending_review"
	MemoryStatusStatusSkipped       MemoryStatusStatus = "skipped"
)

// A document-shaped API response carrying the hyperdoc tree (ENG-2479/D12).
type MemoryListResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// `children` for the body — see the `<Hyperdoc />` renderer.
	Document   MemoryListResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                          `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryListResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// The document's collection, if any.
	Collection string `json:"collection" api:"nullable"`
	// The document's own date (e.g. email sent date, event date).
	DocumentDate time.Time `json:"document_date" api:"nullable" format:"date-time"`
	// When Hyperspell first indexed the document.
	IngestedAt time.Time `json:"ingested_at" api:"nullable" format:"date-time"`
	// When the source document was last modified.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Filterable custom metadata attached to the document.
	Metadata map[string]any `json:"metadata"`
	// Indexing status of the document.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status MemoryListResponseStatus `json:"status" api:"nullable"`
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
		IngestedAt     respjson.Field
		LastModifiedAt respjson.Field
		Metadata       respjson.Field
		Status         respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryListResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryListResponseDocumentUnion contains all possible properties and values from
// [shared.Document], [shared.Website], [shared.Task], [shared.Person],
// [shared.Message], [shared.Event], [shared.File], [shared.Conversation],
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal].
//
// Use the [MemoryListResponseDocumentUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MemoryListResponseDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]shared.DocumentChildUnion],
	// [[]shared.WebsiteChildUnion], [[]shared.TaskChildUnion],
	// [[]shared.PersonChildUnion], [[]shared.MessageChildUnion],
	// [[]shared.EventChildUnion], [[]shared.FileChildUnion], [[]shared.Message],
	// [[]shared.TraceChildUnion], [[]shared.Utterance], [[]shared.CompanyChildUnion],
	// [[]shared.DealChildUnion]
	Children MemoryListResponseDocumentUnionChildren `json:"children"`
	// This field is from variant [shared.Document].
	Metadata shared.Metadata `json:"metadata"`
	Text     string          `json:"text"`
	Title    string          `json:"title"`
	// Any of "document", "website", "task", "person", "message", "event", "file",
	// "conversation", "trace", "transcript", "company", "deal".
	Type string `json:"type"`
	// This field is from variant [shared.Website].
	URL         string `json:"url"`
	Description string `json:"description"`
	// This field is from variant [shared.Website].
	Favicon  string `json:"favicon"`
	ImageURL string `json:"image_url"`
	// This field is from variant [shared.Website].
	Language string `json:"language"`
	// This field is from variant [shared.Task].
	Comments []shared.Message `json:"comments"`
	// This field is from variant [shared.Task].
	DueAt time.Time `json:"due_at"`
	// This field is from variant [shared.Task].
	Priority shared.TaskPriority `json:"priority"`
	// This field is from variant [shared.Task].
	Status  shared.TaskStatus `json:"status"`
	Address string            `json:"address"`
	// This field is from variant [shared.Person].
	AltNames []string `json:"alt_names"`
	// This field is from variant [shared.Person].
	Company    string   `json:"company"`
	CompanyIDs []string `json:"company_ids"`
	// This field is from variant [shared.Person].
	DateOfBirth time.Time `json:"date_of_birth"`
	DealIDs     []string  `json:"deal_ids"`
	// This field is from variant [shared.Person].
	Email  string   `json:"email"`
	Emails []string `json:"emails"`
	// This field is from variant [shared.Person].
	JobTitle string `json:"job_title"`
	// This field is from variant [shared.Person].
	LinkURLs     []string `json:"link_urls"`
	Name         string   `json:"name"`
	PhoneNumbers []string `json:"phone_numbers"`
	Tags         []string `json:"tags"`
	// This field is from variant [shared.Person].
	Username string `json:"username"`
	// This field is from variant [shared.Message].
	Date time.Time `json:"date"`
	// This field is from variant [shared.Message].
	Sender  shared.Person `json:"sender"`
	Channel string        `json:"channel"`
	// This field is from variant [shared.Message].
	ExternalID string `json:"external_id"`
	// This field is from variant [shared.Message].
	IsSelf bool `json:"is_self"`
	// This field is from variant [shared.Message].
	MentionedUsers []shared.Person `json:"mentioned_users"`
	// This field is from variant [shared.Message].
	NumReplies int64 `json:"num_replies"`
	// This field is from variant [shared.Message].
	Replies []shared.Message `json:"replies"`
	// This field is from variant [shared.Message].
	ThreadID string `json:"thread_id"`
	// This field is from variant [shared.Message].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [shared.Message].
	Upvotes int64 `json:"upvotes"`
	// This field is from variant [shared.Event].
	Attendees []shared.Person `json:"attendees"`
	// This field is from variant [shared.Event].
	EndAt time.Time `json:"end_at"`
	// This field is from variant [shared.Event].
	Location string `json:"location"`
	// This field is from variant [shared.Event].
	MeetingURL string `json:"meeting_url"`
	// This field is from variant [shared.Event].
	StartAt time.Time `json:"start_at"`
	// This field is from variant [shared.File].
	ContentType string `json:"content_type"`
	// This field is from variant [shared.File].
	Filename string `json:"filename"`
	// This field is from variant [shared.File].
	Path []string `json:"path"`
	// This field is from variant [shared.Transcript].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [shared.Transcript].
	Participants []shared.Person `json:"participants"`
	// This field is from variant [shared.Transcript].
	StartedAt  time.Time `json:"started_at"`
	ContactIDs []string  `json:"contact_ids"`
	// This field is from variant [shared.Company].
	Employees int64 `json:"employees"`
	// This field is from variant [shared.Company].
	Industry string `json:"industry"`
	// This field is from variant [shared.Company].
	IsActive bool `json:"is_active"`
	// This field is from variant [shared.Company].
	Timezone string `json:"timezone"`
	// This field is from variant [shared.Company].
	Websites []string `json:"websites"`
	// This field is from variant [shared.Deal].
	Amount float64 `json:"amount"`
	// This field is from variant [shared.Deal].
	ClosedAt time.Time `json:"closed_at"`
	// This field is from variant [shared.Deal].
	Currency string `json:"currency"`
	// This field is from variant [shared.Deal].
	DealSource string `json:"deal_source"`
	// This field is from variant [shared.Deal].
	LostReason string `json:"lost_reason"`
	// This field is from variant [shared.Deal].
	Pipeline string `json:"pipeline"`
	// This field is from variant [shared.Deal].
	Probability float64 `json:"probability"`
	// This field is from variant [shared.Deal].
	Stage string `json:"stage"`
	// This field is from variant [shared.Deal].
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

// anyMemoryListResponseDocument is implemented by each variant of
// [MemoryListResponseDocumentUnion] to add type safety for the return type of
// [MemoryListResponseDocumentUnion.AsAny]
type anyMemoryListResponseDocument interface {
	ImplMemoryListResponseDocumentUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MemoryListResponseDocumentUnion.AsAny().(type) {
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
func (u MemoryListResponseDocumentUnion) AsAny() anyMemoryListResponseDocument {
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

func (u MemoryListResponseDocumentUnion) AsDocument() (v shared.Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsWebsite() (v shared.Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsPerson() (v shared.Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsEvent() (v shared.Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsFile() (v shared.File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsConversation() (v shared.Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsTrace() (v shared.Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsTranscript() (v shared.Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsCompany() (v shared.Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentUnion) AsDeal() (v shared.Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemoryListResponseDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *MemoryListResponseDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryListResponseDocumentUnionChildren is an implicit subunion of
// [MemoryListResponseDocumentUnion]. MemoryListResponseDocumentUnionChildren
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MemoryListResponseDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type MemoryListResponseDocumentUnionChildren struct {
	// This field will be present if the value is a [[]shared.DocumentChildUnion]
	// instead of an object.
	OfChildren []shared.DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MemoryListResponseDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryListResponseSource string

const (
	MemoryListResponseSourceReddit         MemoryListResponseSource = "reddit"
	MemoryListResponseSourceNotion         MemoryListResponseSource = "notion"
	MemoryListResponseSourceSlack          MemoryListResponseSource = "slack"
	MemoryListResponseSourceGoogleCalendar MemoryListResponseSource = "google_calendar"
	MemoryListResponseSourceGoogleMail     MemoryListResponseSource = "google_mail"
	MemoryListResponseSourceBox            MemoryListResponseSource = "box"
	MemoryListResponseSourceDropbox        MemoryListResponseSource = "dropbox"
	MemoryListResponseSourceGitHub         MemoryListResponseSource = "github"
	MemoryListResponseSourceGoogleDrive    MemoryListResponseSource = "google_drive"
	MemoryListResponseSourceVault          MemoryListResponseSource = "vault"
	MemoryListResponseSourceWebCrawler     MemoryListResponseSource = "web_crawler"
	MemoryListResponseSourceTrace          MemoryListResponseSource = "trace"
	MemoryListResponseSourceMicrosoftTeams MemoryListResponseSource = "microsoft_teams"
	MemoryListResponseSourceGmailActions   MemoryListResponseSource = "gmail_actions"
	MemoryListResponseSourceGranola        MemoryListResponseSource = "granola"
	MemoryListResponseSourceFathom         MemoryListResponseSource = "fathom"
	MemoryListResponseSourceFireflies      MemoryListResponseSource = "fireflies"
	MemoryListResponseSourceLinear         MemoryListResponseSource = "linear"
	MemoryListResponseSourceHubspot        MemoryListResponseSource = "hubspot"
	MemoryListResponseSourceSalesforce     MemoryListResponseSource = "salesforce"
	MemoryListResponseSourceCoda           MemoryListResponseSource = "coda"
	MemoryListResponseSourceLightfield     MemoryListResponseSource = "lightfield"
	MemoryListResponseSourceGong           MemoryListResponseSource = "gong"
)

// Indexing status of the document.
type MemoryListResponseStatus string

const (
	MemoryListResponseStatusPending       MemoryListResponseStatus = "pending"
	MemoryListResponseStatusProcessing    MemoryListResponseStatus = "processing"
	MemoryListResponseStatusCompleted     MemoryListResponseStatus = "completed"
	MemoryListResponseStatusFailed        MemoryListResponseStatus = "failed"
	MemoryListResponseStatusPendingReview MemoryListResponseStatus = "pending_review"
	MemoryListResponseStatusSkipped       MemoryListResponseStatus = "skipped"
)

type MemoryDeleteResponse struct {
	ChunksDeleted int64  `json:"chunks_deleted" api:"required"`
	Message       string `json:"message" api:"required"`
	ResourceID    string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source  MemoryDeleteResponseSource `json:"source" api:"required"`
	Success bool                       `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunksDeleted respjson.Field
		Message       respjson.Field
		ResourceID    respjson.Field
		Source        respjson.Field
		Success       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryDeleteResponseSource string

const (
	MemoryDeleteResponseSourceReddit         MemoryDeleteResponseSource = "reddit"
	MemoryDeleteResponseSourceNotion         MemoryDeleteResponseSource = "notion"
	MemoryDeleteResponseSourceSlack          MemoryDeleteResponseSource = "slack"
	MemoryDeleteResponseSourceGoogleCalendar MemoryDeleteResponseSource = "google_calendar"
	MemoryDeleteResponseSourceGoogleMail     MemoryDeleteResponseSource = "google_mail"
	MemoryDeleteResponseSourceBox            MemoryDeleteResponseSource = "box"
	MemoryDeleteResponseSourceDropbox        MemoryDeleteResponseSource = "dropbox"
	MemoryDeleteResponseSourceGitHub         MemoryDeleteResponseSource = "github"
	MemoryDeleteResponseSourceGoogleDrive    MemoryDeleteResponseSource = "google_drive"
	MemoryDeleteResponseSourceVault          MemoryDeleteResponseSource = "vault"
	MemoryDeleteResponseSourceWebCrawler     MemoryDeleteResponseSource = "web_crawler"
	MemoryDeleteResponseSourceTrace          MemoryDeleteResponseSource = "trace"
	MemoryDeleteResponseSourceMicrosoftTeams MemoryDeleteResponseSource = "microsoft_teams"
	MemoryDeleteResponseSourceGmailActions   MemoryDeleteResponseSource = "gmail_actions"
	MemoryDeleteResponseSourceGranola        MemoryDeleteResponseSource = "granola"
	MemoryDeleteResponseSourceFathom         MemoryDeleteResponseSource = "fathom"
	MemoryDeleteResponseSourceFireflies      MemoryDeleteResponseSource = "fireflies"
	MemoryDeleteResponseSourceLinear         MemoryDeleteResponseSource = "linear"
	MemoryDeleteResponseSourceHubspot        MemoryDeleteResponseSource = "hubspot"
	MemoryDeleteResponseSourceSalesforce     MemoryDeleteResponseSource = "salesforce"
	MemoryDeleteResponseSourceCoda           MemoryDeleteResponseSource = "coda"
	MemoryDeleteResponseSourceLightfield     MemoryDeleteResponseSource = "lightfield"
	MemoryDeleteResponseSourceGong           MemoryDeleteResponseSource = "gong"
)

// Response schema for successful bulk ingestion.
type MemoryAddBulkResponse struct {
	// Number of items successfully processed
	Count int64 `json:"count" api:"required"`
	// Status of each ingested item
	Items   []MemoryStatus `json:"items" api:"required"`
	Success bool           `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Items       respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryAddBulkResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryAddBulkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document-shaped API response carrying the hyperdoc tree (ENG-2479/D12).
type MemoryGetResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// `children` for the body — see the `<Hyperdoc />` renderer.
	Document   MemoryGetResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                         `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryGetResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// The document's collection, if any.
	Collection string `json:"collection" api:"nullable"`
	// The document's own date (e.g. email sent date, event date).
	DocumentDate time.Time `json:"document_date" api:"nullable" format:"date-time"`
	// When Hyperspell first indexed the document.
	IngestedAt time.Time `json:"ingested_at" api:"nullable" format:"date-time"`
	// When the source document was last modified.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Filterable custom metadata attached to the document.
	Metadata map[string]any `json:"metadata"`
	// Indexing status of the document.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status MemoryGetResponseStatus `json:"status" api:"nullable"`
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
		IngestedAt     respjson.Field
		LastModifiedAt respjson.Field
		Metadata       respjson.Field
		Status         respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryGetResponseDocumentUnion contains all possible properties and values from
// [shared.Document], [shared.Website], [shared.Task], [shared.Person],
// [shared.Message], [shared.Event], [shared.File], [shared.Conversation],
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal].
//
// Use the [MemoryGetResponseDocumentUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MemoryGetResponseDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]shared.DocumentChildUnion],
	// [[]shared.WebsiteChildUnion], [[]shared.TaskChildUnion],
	// [[]shared.PersonChildUnion], [[]shared.MessageChildUnion],
	// [[]shared.EventChildUnion], [[]shared.FileChildUnion], [[]shared.Message],
	// [[]shared.TraceChildUnion], [[]shared.Utterance], [[]shared.CompanyChildUnion],
	// [[]shared.DealChildUnion]
	Children MemoryGetResponseDocumentUnionChildren `json:"children"`
	// This field is from variant [shared.Document].
	Metadata shared.Metadata `json:"metadata"`
	Text     string          `json:"text"`
	Title    string          `json:"title"`
	// Any of "document", "website", "task", "person", "message", "event", "file",
	// "conversation", "trace", "transcript", "company", "deal".
	Type string `json:"type"`
	// This field is from variant [shared.Website].
	URL         string `json:"url"`
	Description string `json:"description"`
	// This field is from variant [shared.Website].
	Favicon  string `json:"favicon"`
	ImageURL string `json:"image_url"`
	// This field is from variant [shared.Website].
	Language string `json:"language"`
	// This field is from variant [shared.Task].
	Comments []shared.Message `json:"comments"`
	// This field is from variant [shared.Task].
	DueAt time.Time `json:"due_at"`
	// This field is from variant [shared.Task].
	Priority shared.TaskPriority `json:"priority"`
	// This field is from variant [shared.Task].
	Status  shared.TaskStatus `json:"status"`
	Address string            `json:"address"`
	// This field is from variant [shared.Person].
	AltNames []string `json:"alt_names"`
	// This field is from variant [shared.Person].
	Company    string   `json:"company"`
	CompanyIDs []string `json:"company_ids"`
	// This field is from variant [shared.Person].
	DateOfBirth time.Time `json:"date_of_birth"`
	DealIDs     []string  `json:"deal_ids"`
	// This field is from variant [shared.Person].
	Email  string   `json:"email"`
	Emails []string `json:"emails"`
	// This field is from variant [shared.Person].
	JobTitle string `json:"job_title"`
	// This field is from variant [shared.Person].
	LinkURLs     []string `json:"link_urls"`
	Name         string   `json:"name"`
	PhoneNumbers []string `json:"phone_numbers"`
	Tags         []string `json:"tags"`
	// This field is from variant [shared.Person].
	Username string `json:"username"`
	// This field is from variant [shared.Message].
	Date time.Time `json:"date"`
	// This field is from variant [shared.Message].
	Sender  shared.Person `json:"sender"`
	Channel string        `json:"channel"`
	// This field is from variant [shared.Message].
	ExternalID string `json:"external_id"`
	// This field is from variant [shared.Message].
	IsSelf bool `json:"is_self"`
	// This field is from variant [shared.Message].
	MentionedUsers []shared.Person `json:"mentioned_users"`
	// This field is from variant [shared.Message].
	NumReplies int64 `json:"num_replies"`
	// This field is from variant [shared.Message].
	Replies []shared.Message `json:"replies"`
	// This field is from variant [shared.Message].
	ThreadID string `json:"thread_id"`
	// This field is from variant [shared.Message].
	UpdatedAt time.Time `json:"updated_at"`
	// This field is from variant [shared.Message].
	Upvotes int64 `json:"upvotes"`
	// This field is from variant [shared.Event].
	Attendees []shared.Person `json:"attendees"`
	// This field is from variant [shared.Event].
	EndAt time.Time `json:"end_at"`
	// This field is from variant [shared.Event].
	Location string `json:"location"`
	// This field is from variant [shared.Event].
	MeetingURL string `json:"meeting_url"`
	// This field is from variant [shared.Event].
	StartAt time.Time `json:"start_at"`
	// This field is from variant [shared.File].
	ContentType string `json:"content_type"`
	// This field is from variant [shared.File].
	Filename string `json:"filename"`
	// This field is from variant [shared.File].
	Path []string `json:"path"`
	// This field is from variant [shared.Transcript].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [shared.Transcript].
	Participants []shared.Person `json:"participants"`
	// This field is from variant [shared.Transcript].
	StartedAt  time.Time `json:"started_at"`
	ContactIDs []string  `json:"contact_ids"`
	// This field is from variant [shared.Company].
	Employees int64 `json:"employees"`
	// This field is from variant [shared.Company].
	Industry string `json:"industry"`
	// This field is from variant [shared.Company].
	IsActive bool `json:"is_active"`
	// This field is from variant [shared.Company].
	Timezone string `json:"timezone"`
	// This field is from variant [shared.Company].
	Websites []string `json:"websites"`
	// This field is from variant [shared.Deal].
	Amount float64 `json:"amount"`
	// This field is from variant [shared.Deal].
	ClosedAt time.Time `json:"closed_at"`
	// This field is from variant [shared.Deal].
	Currency string `json:"currency"`
	// This field is from variant [shared.Deal].
	DealSource string `json:"deal_source"`
	// This field is from variant [shared.Deal].
	LostReason string `json:"lost_reason"`
	// This field is from variant [shared.Deal].
	Pipeline string `json:"pipeline"`
	// This field is from variant [shared.Deal].
	Probability float64 `json:"probability"`
	// This field is from variant [shared.Deal].
	Stage string `json:"stage"`
	// This field is from variant [shared.Deal].
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

// anyMemoryGetResponseDocument is implemented by each variant of
// [MemoryGetResponseDocumentUnion] to add type safety for the return type of
// [MemoryGetResponseDocumentUnion.AsAny]
type anyMemoryGetResponseDocument interface {
	ImplMemoryGetResponseDocumentUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MemoryGetResponseDocumentUnion.AsAny().(type) {
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
func (u MemoryGetResponseDocumentUnion) AsAny() anyMemoryGetResponseDocument {
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

func (u MemoryGetResponseDocumentUnion) AsDocument() (v shared.Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsWebsite() (v shared.Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsPerson() (v shared.Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsEvent() (v shared.Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsFile() (v shared.File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsConversation() (v shared.Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsTrace() (v shared.Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsTranscript() (v shared.Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsCompany() (v shared.Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentUnion) AsDeal() (v shared.Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemoryGetResponseDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *MemoryGetResponseDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryGetResponseDocumentUnionChildren is an implicit subunion of
// [MemoryGetResponseDocumentUnion]. MemoryGetResponseDocumentUnionChildren
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MemoryGetResponseDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type MemoryGetResponseDocumentUnionChildren struct {
	// This field will be present if the value is a [[]shared.DocumentChildUnion]
	// instead of an object.
	OfChildren []shared.DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MemoryGetResponseDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryGetResponseSource string

const (
	MemoryGetResponseSourceReddit         MemoryGetResponseSource = "reddit"
	MemoryGetResponseSourceNotion         MemoryGetResponseSource = "notion"
	MemoryGetResponseSourceSlack          MemoryGetResponseSource = "slack"
	MemoryGetResponseSourceGoogleCalendar MemoryGetResponseSource = "google_calendar"
	MemoryGetResponseSourceGoogleMail     MemoryGetResponseSource = "google_mail"
	MemoryGetResponseSourceBox            MemoryGetResponseSource = "box"
	MemoryGetResponseSourceDropbox        MemoryGetResponseSource = "dropbox"
	MemoryGetResponseSourceGitHub         MemoryGetResponseSource = "github"
	MemoryGetResponseSourceGoogleDrive    MemoryGetResponseSource = "google_drive"
	MemoryGetResponseSourceVault          MemoryGetResponseSource = "vault"
	MemoryGetResponseSourceWebCrawler     MemoryGetResponseSource = "web_crawler"
	MemoryGetResponseSourceTrace          MemoryGetResponseSource = "trace"
	MemoryGetResponseSourceMicrosoftTeams MemoryGetResponseSource = "microsoft_teams"
	MemoryGetResponseSourceGmailActions   MemoryGetResponseSource = "gmail_actions"
	MemoryGetResponseSourceGranola        MemoryGetResponseSource = "granola"
	MemoryGetResponseSourceFathom         MemoryGetResponseSource = "fathom"
	MemoryGetResponseSourceFireflies      MemoryGetResponseSource = "fireflies"
	MemoryGetResponseSourceLinear         MemoryGetResponseSource = "linear"
	MemoryGetResponseSourceHubspot        MemoryGetResponseSource = "hubspot"
	MemoryGetResponseSourceSalesforce     MemoryGetResponseSource = "salesforce"
	MemoryGetResponseSourceCoda           MemoryGetResponseSource = "coda"
	MemoryGetResponseSourceLightfield     MemoryGetResponseSource = "lightfield"
	MemoryGetResponseSourceGong           MemoryGetResponseSource = "gong"
)

// Indexing status of the document.
type MemoryGetResponseStatus string

const (
	MemoryGetResponseStatusPending       MemoryGetResponseStatus = "pending"
	MemoryGetResponseStatusProcessing    MemoryGetResponseStatus = "processing"
	MemoryGetResponseStatusCompleted     MemoryGetResponseStatus = "completed"
	MemoryGetResponseStatusFailed        MemoryGetResponseStatus = "failed"
	MemoryGetResponseStatusPendingReview MemoryGetResponseStatus = "pending_review"
	MemoryGetResponseStatusSkipped       MemoryGetResponseStatus = "skipped"
)

type MemoryStatusResponse struct {
	Providers map[string]map[string]int64 `json:"providers" api:"required"`
	Total     map[string]int64            `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Providers   respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryUpdateParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryUpdateParamsSource `path:"source,omitzero" api:"required" json:"-"`
	// The collection to move the document to — deprecated, set the collection using
	// metadata instead.
	Collection any `json:"collection,omitzero"`
	// Date of the document for ranking and filtering.
	Date any `json:"date,omitzero"`
	// Custom metadata for filtering. Keys must be alphanumeric with underscores, max
	// 64 chars. Values must be string, number, boolean, or null. Will be merged with
	// existing metadata.
	Metadata map[string]MemoryUpdateParamsMetadataUnion `json:"metadata,omitzero"`
	// Full text of the document. If provided, the document will be re-indexed.
	Text any `json:"text,omitzero"`
	// Title of the document.
	Title any `json:"title,omitzero"`
	paramObj
}

func (r MemoryUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryUpdateParamsSource string

const (
	MemoryUpdateParamsSourceReddit         MemoryUpdateParamsSource = "reddit"
	MemoryUpdateParamsSourceNotion         MemoryUpdateParamsSource = "notion"
	MemoryUpdateParamsSourceSlack          MemoryUpdateParamsSource = "slack"
	MemoryUpdateParamsSourceGoogleCalendar MemoryUpdateParamsSource = "google_calendar"
	MemoryUpdateParamsSourceGoogleMail     MemoryUpdateParamsSource = "google_mail"
	MemoryUpdateParamsSourceBox            MemoryUpdateParamsSource = "box"
	MemoryUpdateParamsSourceDropbox        MemoryUpdateParamsSource = "dropbox"
	MemoryUpdateParamsSourceGitHub         MemoryUpdateParamsSource = "github"
	MemoryUpdateParamsSourceGoogleDrive    MemoryUpdateParamsSource = "google_drive"
	MemoryUpdateParamsSourceVault          MemoryUpdateParamsSource = "vault"
	MemoryUpdateParamsSourceWebCrawler     MemoryUpdateParamsSource = "web_crawler"
	MemoryUpdateParamsSourceTrace          MemoryUpdateParamsSource = "trace"
	MemoryUpdateParamsSourceMicrosoftTeams MemoryUpdateParamsSource = "microsoft_teams"
	MemoryUpdateParamsSourceGmailActions   MemoryUpdateParamsSource = "gmail_actions"
	MemoryUpdateParamsSourceGranola        MemoryUpdateParamsSource = "granola"
	MemoryUpdateParamsSourceFathom         MemoryUpdateParamsSource = "fathom"
	MemoryUpdateParamsSourceFireflies      MemoryUpdateParamsSource = "fireflies"
	MemoryUpdateParamsSourceLinear         MemoryUpdateParamsSource = "linear"
	MemoryUpdateParamsSourceHubspot        MemoryUpdateParamsSource = "hubspot"
	MemoryUpdateParamsSourceSalesforce     MemoryUpdateParamsSource = "salesforce"
	MemoryUpdateParamsSourceCoda           MemoryUpdateParamsSource = "coda"
	MemoryUpdateParamsSourceLightfield     MemoryUpdateParamsSource = "lightfield"
	MemoryUpdateParamsSourceGong           MemoryUpdateParamsSource = "gong"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MemoryUpdateParamsMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MemoryUpdateParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MemoryUpdateParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MemoryListParams struct {
	// Filter documents by collection.
	Collection param.Opt[string] `query:"collection,omitzero" json:"-"`
	Cursor     param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter documents by metadata using MongoDB-style operators. Example:
	// {"department": "engineering", "priority": {"$gt": 3}}
	Filter param.Opt[string] `query:"filter,omitzero" json:"-"`
	Size   param.Opt[int64]  `query:"size,omitzero" json:"-"`
	// Filter documents by source.
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryListParamsSource `query:"source,omitzero" json:"-"`
	// Filter documents by status.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status MemoryListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemoryListParams]'s query parameters as `url.Values`.
func (r MemoryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter documents by source.
type MemoryListParamsSource string

const (
	MemoryListParamsSourceReddit         MemoryListParamsSource = "reddit"
	MemoryListParamsSourceNotion         MemoryListParamsSource = "notion"
	MemoryListParamsSourceSlack          MemoryListParamsSource = "slack"
	MemoryListParamsSourceGoogleCalendar MemoryListParamsSource = "google_calendar"
	MemoryListParamsSourceGoogleMail     MemoryListParamsSource = "google_mail"
	MemoryListParamsSourceBox            MemoryListParamsSource = "box"
	MemoryListParamsSourceDropbox        MemoryListParamsSource = "dropbox"
	MemoryListParamsSourceGitHub         MemoryListParamsSource = "github"
	MemoryListParamsSourceGoogleDrive    MemoryListParamsSource = "google_drive"
	MemoryListParamsSourceVault          MemoryListParamsSource = "vault"
	MemoryListParamsSourceWebCrawler     MemoryListParamsSource = "web_crawler"
	MemoryListParamsSourceTrace          MemoryListParamsSource = "trace"
	MemoryListParamsSourceMicrosoftTeams MemoryListParamsSource = "microsoft_teams"
	MemoryListParamsSourceGmailActions   MemoryListParamsSource = "gmail_actions"
	MemoryListParamsSourceGranola        MemoryListParamsSource = "granola"
	MemoryListParamsSourceFathom         MemoryListParamsSource = "fathom"
	MemoryListParamsSourceFireflies      MemoryListParamsSource = "fireflies"
	MemoryListParamsSourceLinear         MemoryListParamsSource = "linear"
	MemoryListParamsSourceHubspot        MemoryListParamsSource = "hubspot"
	MemoryListParamsSourceSalesforce     MemoryListParamsSource = "salesforce"
	MemoryListParamsSourceCoda           MemoryListParamsSource = "coda"
	MemoryListParamsSourceLightfield     MemoryListParamsSource = "lightfield"
	MemoryListParamsSourceGong           MemoryListParamsSource = "gong"
)

// Filter documents by status.
type MemoryListParamsStatus string

const (
	MemoryListParamsStatusPending       MemoryListParamsStatus = "pending"
	MemoryListParamsStatusProcessing    MemoryListParamsStatus = "processing"
	MemoryListParamsStatusCompleted     MemoryListParamsStatus = "completed"
	MemoryListParamsStatusFailed        MemoryListParamsStatus = "failed"
	MemoryListParamsStatusPendingReview MemoryListParamsStatus = "pending_review"
	MemoryListParamsStatusSkipped       MemoryListParamsStatus = "skipped"
)

type MemoryDeleteParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryDeleteParamsSource `path:"source,omitzero" api:"required" json:"-"`
	paramObj
}

type MemoryDeleteParamsSource string

const (
	MemoryDeleteParamsSourceReddit         MemoryDeleteParamsSource = "reddit"
	MemoryDeleteParamsSourceNotion         MemoryDeleteParamsSource = "notion"
	MemoryDeleteParamsSourceSlack          MemoryDeleteParamsSource = "slack"
	MemoryDeleteParamsSourceGoogleCalendar MemoryDeleteParamsSource = "google_calendar"
	MemoryDeleteParamsSourceGoogleMail     MemoryDeleteParamsSource = "google_mail"
	MemoryDeleteParamsSourceBox            MemoryDeleteParamsSource = "box"
	MemoryDeleteParamsSourceDropbox        MemoryDeleteParamsSource = "dropbox"
	MemoryDeleteParamsSourceGitHub         MemoryDeleteParamsSource = "github"
	MemoryDeleteParamsSourceGoogleDrive    MemoryDeleteParamsSource = "google_drive"
	MemoryDeleteParamsSourceVault          MemoryDeleteParamsSource = "vault"
	MemoryDeleteParamsSourceWebCrawler     MemoryDeleteParamsSource = "web_crawler"
	MemoryDeleteParamsSourceTrace          MemoryDeleteParamsSource = "trace"
	MemoryDeleteParamsSourceMicrosoftTeams MemoryDeleteParamsSource = "microsoft_teams"
	MemoryDeleteParamsSourceGmailActions   MemoryDeleteParamsSource = "gmail_actions"
	MemoryDeleteParamsSourceGranola        MemoryDeleteParamsSource = "granola"
	MemoryDeleteParamsSourceFathom         MemoryDeleteParamsSource = "fathom"
	MemoryDeleteParamsSourceFireflies      MemoryDeleteParamsSource = "fireflies"
	MemoryDeleteParamsSourceLinear         MemoryDeleteParamsSource = "linear"
	MemoryDeleteParamsSourceHubspot        MemoryDeleteParamsSource = "hubspot"
	MemoryDeleteParamsSourceSalesforce     MemoryDeleteParamsSource = "salesforce"
	MemoryDeleteParamsSourceCoda           MemoryDeleteParamsSource = "coda"
	MemoryDeleteParamsSourceLightfield     MemoryDeleteParamsSource = "lightfield"
	MemoryDeleteParamsSourceGong           MemoryDeleteParamsSource = "gong"
)

type MemoryAddParams struct {
	// Full text of the document.
	Text string `json:"text" api:"required"`
	// The collection to add the document to — deprecated, set the collection using
	// metadata instead.
	Collection param.Opt[string] `json:"collection,omitzero"`
	// Title of the document.
	Title param.Opt[string] `json:"title,omitzero"`
	// Date of the document. Depending on the document, this could be the creation date
	// or date the document was last updated (eg. for a chat transcript, this would be
	// the date of the last message). This helps the ranking algorithm and allows you
	// to filter by date range.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// The resource ID to add the document to. If not provided, a new resource ID will
	// be generated. If provided, the document will be updated if it already exists.
	ResourceID param.Opt[string] `json:"resource_id,omitzero"`
	// Custom metadata for filtering. Keys must be alphanumeric with underscores, max
	// 64 chars. Values must be string, number, boolean, or null.
	Metadata map[string]MemoryAddParamsMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r MemoryAddParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MemoryAddParamsMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MemoryAddParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MemoryAddParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MemoryAddBulkParams struct {
	// List of memories to ingest. Maximum 100 items.
	Items []MemoryAddBulkParamsItem `json:"items,omitzero" api:"required"`
	paramObj
}

func (r MemoryAddBulkParams) MarshalJSON() (data []byte, err error) {
	type shadow MemoryAddBulkParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryAddBulkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Text is required.
type MemoryAddBulkParamsItem struct {
	// Full text of the document.
	Text string `json:"text" api:"required"`
	// The collection to add the document to — deprecated, set the collection using
	// metadata instead.
	//
	// Deprecated: deprecated
	Collection param.Opt[string] `json:"collection,omitzero"`
	// Title of the document.
	Title param.Opt[string] `json:"title,omitzero"`
	// Date of the document. Depending on the document, this could be the creation date
	// or date the document was last updated (eg. for a chat transcript, this would be
	// the date of the last message). This helps the ranking algorithm and allows you
	// to filter by date range.
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// The resource ID to add the document to. If not provided, a new resource ID will
	// be generated. If provided, the document will be updated if it already exists.
	ResourceID param.Opt[string] `json:"resource_id,omitzero"`
	// Custom metadata for filtering. Keys must be alphanumeric with underscores, max
	// 64 chars. Values must be string, number, boolean, or null.
	Metadata map[string]MemoryAddBulkParamsItemMetadataUnion `json:"metadata,omitzero"`
	paramObj
}

func (r MemoryAddBulkParamsItem) MarshalJSON() (data []byte, err error) {
	type shadow MemoryAddBulkParamsItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemoryAddBulkParamsItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type MemoryAddBulkParamsItemMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u MemoryAddBulkParamsItemMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *MemoryAddBulkParamsItemMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

type MemoryGetParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Source MemoryGetParamsSource `path:"source,omitzero" api:"required" json:"-"`
	paramObj
}

type MemoryGetParamsSource string

const (
	MemoryGetParamsSourceReddit         MemoryGetParamsSource = "reddit"
	MemoryGetParamsSourceNotion         MemoryGetParamsSource = "notion"
	MemoryGetParamsSourceSlack          MemoryGetParamsSource = "slack"
	MemoryGetParamsSourceGoogleCalendar MemoryGetParamsSource = "google_calendar"
	MemoryGetParamsSourceGoogleMail     MemoryGetParamsSource = "google_mail"
	MemoryGetParamsSourceBox            MemoryGetParamsSource = "box"
	MemoryGetParamsSourceDropbox        MemoryGetParamsSource = "dropbox"
	MemoryGetParamsSourceGitHub         MemoryGetParamsSource = "github"
	MemoryGetParamsSourceGoogleDrive    MemoryGetParamsSource = "google_drive"
	MemoryGetParamsSourceVault          MemoryGetParamsSource = "vault"
	MemoryGetParamsSourceWebCrawler     MemoryGetParamsSource = "web_crawler"
	MemoryGetParamsSourceTrace          MemoryGetParamsSource = "trace"
	MemoryGetParamsSourceMicrosoftTeams MemoryGetParamsSource = "microsoft_teams"
	MemoryGetParamsSourceGmailActions   MemoryGetParamsSource = "gmail_actions"
	MemoryGetParamsSourceGranola        MemoryGetParamsSource = "granola"
	MemoryGetParamsSourceFathom         MemoryGetParamsSource = "fathom"
	MemoryGetParamsSourceFireflies      MemoryGetParamsSource = "fireflies"
	MemoryGetParamsSourceLinear         MemoryGetParamsSource = "linear"
	MemoryGetParamsSourceHubspot        MemoryGetParamsSource = "hubspot"
	MemoryGetParamsSourceSalesforce     MemoryGetParamsSource = "salesforce"
	MemoryGetParamsSourceCoda           MemoryGetParamsSource = "coda"
	MemoryGetParamsSourceLightfield     MemoryGetParamsSource = "lightfield"
	MemoryGetParamsSourceGong           MemoryGetParamsSource = "gong"
)

type MemorySearchParams struct {
	// Query to run.
	Query string `json:"query" api:"required"`
	// If true, the query will be answered along with matching source documents.
	Answer param.Opt[bool] `json:"answer,omitzero"`
	// Maximum number of results to return.
	MaxResults param.Opt[int64] `json:"max_results,omitzero"`
	// If true (effort='very_high' only), attach a provenance record to the response:
	// the source documents and entities the answer was grounded in, the agent's search
	// trajectory, and any sources that failed. Adds one indexed lookup; intended for
	// auditability / compliance use cases.
	Provenance param.Opt[bool] `json:"provenance,omitzero"`
	// How much compute to spend on retrieval. Mirrors the dial popularized by
	// frontier-model APIs (OpenAI reasoning_effort, etc.). 'minimal' = verbatim
	// single-shot retrieval (fastest). 'low' = LLM rewrites the query for better
	// retrieval and extracts date filters. 'medium' = rewrite + agentic refinement
	// loop (the answer LLM may request additional retrieval rounds, up to 3). 'high' =
	// rewrite + extended refinement (up to 6 rounds). Higher = better recall, more
	// latency, more cost.
	//
	// Any of "minimal", "low", "medium", "high", "very_high".
	Effort MemorySearchParamsEffort `json:"effort,omitzero"`
	// Search options for the query.
	Options MemorySearchParamsOptions `json:"options,omitzero"`
	// Only query documents from these sources.
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield", "gong".
	Sources []string `json:"sources,omitzero"`
	paramObj
}

func (r MemorySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How much compute to spend on retrieval. Mirrors the dial popularized by
// frontier-model APIs (OpenAI reasoning_effort, etc.). 'minimal' = verbatim
// single-shot retrieval (fastest). 'low' = LLM rewrites the query for better
// retrieval and extracts date filters. 'medium' = rewrite + agentic refinement
// loop (the answer LLM may request additional retrieval rounds, up to 3). 'high' =
// rewrite + extended refinement (up to 6 rounds). Higher = better recall, more
// latency, more cost.
type MemorySearchParamsEffort string

const (
	MemorySearchParamsEffortMinimal  MemorySearchParamsEffort = "minimal"
	MemorySearchParamsEffortLow      MemorySearchParamsEffort = "low"
	MemorySearchParamsEffortMedium   MemorySearchParamsEffort = "medium"
	MemorySearchParamsEffortHigh     MemorySearchParamsEffort = "high"
	MemorySearchParamsEffortVeryHigh MemorySearchParamsEffort = "very_high"
)

// Search options for the query.
type MemorySearchParamsOptions struct {
	// Only query documents created on or after this date.
	After param.Opt[time.Time] `json:"after,omitzero" format:"date-time"`
	// Only query documents created before this date.
	Before param.Opt[time.Time] `json:"before,omitzero" format:"date-time"`
	// When set, multiplies each result's score by an exponential-decay factor based on
	// the document's most recent activity timestamp (source-reported last_modified,
	// falling back to document_date). A document one half-life old gets its score
	// halved. Resources with no recency timestamp are passed through unchanged. Leave
	// unset to disable.
	RecencyHalfLifeDays param.Opt[float64] `json:"recency_half_life_days,omitzero"`
	// Maximum number of results to return.
	MaxResults param.Opt[int64] `json:"max_results,omitzero"`
	// Metadata filters using MongoDB-style operators. Example: {'status': 'published',
	// 'priority': {'$gt': 3}}
	Filter map[string]any `json:"filter,omitzero"`
	// Only return results from these specific resource IDs. Useful for scoping
	// searches to specific documents (e.g., a specific email thread or uploaded file).
	ResourceIDs []string `json:"resource_ids,omitzero"`
	// Model to use for answer generation when answer=True
	//
	// Any of "llama-3.1", "gemma2", "qwen-qwq", "mistral-saba", "llama-4-scout",
	// "deepseek-r1", "gpt-oss-20b", "gpt-oss-120b".
	AnswerModel string `json:"answer_model,omitzero"`
	// Search options for Box
	Box MemorySearchParamsOptionsBox `json:"box,omitzero"`
	// Search options for Google Calendar
	GoogleCalendar MemorySearchParamsOptionsGoogleCalendar `json:"google_calendar,omitzero"`
	// Search options for Google Drive
	GoogleDrive MemorySearchParamsOptionsGoogleDrive `json:"google_drive,omitzero"`
	// Search options for Gmail
	GoogleMail MemorySearchParamsOptionsGoogleMail `json:"google_mail,omitzero"`
	// Filter by memory type. Defaults to generic memories only. Pass multiple types to
	// include procedures, etc.
	//
	// Any of "procedure", "memory", "mood".
	MemoryTypes []string `json:"memory_types,omitzero"`
	// Search options for Notion
	Notion MemorySearchParamsOptionsNotion `json:"notion,omitzero"`
	// Search options for Slack
	Slack MemorySearchParamsOptionsSlack `json:"slack,omitzero"`
	// Search options for vault
	Vault MemorySearchParamsOptionsVault `json:"vault,omitzero"`
	// Search options for Web Crawler
	WebCrawler MemorySearchParamsOptionsWebCrawler `json:"web_crawler,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptions) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MemorySearchParamsOptions](
		"answer_model", "llama-3.1", "gemma2", "qwen-qwq", "mistral-saba", "llama-4-scout", "deepseek-r1", "gpt-oss-20b", "gpt-oss-120b",
	)
}

// Search options for Box
type MemorySearchParamsOptionsBox struct {
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsBox) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsBox
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsBox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Google Calendar
type MemorySearchParamsOptionsGoogleCalendar struct {
	// The ID of the calendar to search. If not provided, it will use the ID of the
	// default calendar. You can get the list of calendars with the
	// `/integrations/google_calendar/list` endpoint.
	CalendarID param.Opt[string] `json:"calendar_id,omitzero"`
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsGoogleCalendar) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsGoogleCalendar
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsGoogleCalendar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Google Drive
type MemorySearchParamsOptionsGoogleDrive struct {
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsGoogleDrive) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsGoogleDrive
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsGoogleDrive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Gmail
type MemorySearchParamsOptionsGoogleMail struct {
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// List of label IDs to filter messages (e.g., ['INBOX', 'SENT', 'DRAFT']).
	// Multiple labels are combined with OR logic - messages matching ANY specified
	// label will be returned. If empty, no label filtering is applied (searches all
	// accessible messages).
	LabelIDs []string `json:"label_ids,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsGoogleMail) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsGoogleMail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsGoogleMail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Notion
type MemorySearchParamsOptionsNotion struct {
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// List of Notion page IDs to search. If not provided, all pages in the workspace
	// will be searched.
	NotionPageIDs []string `json:"notion_page_ids,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsNotion) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsNotion
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsNotion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Slack
type MemorySearchParamsOptionsSlack struct {
	// If set, pass 'exclude_archived' to Slack. If None, omit the param.
	ExcludeArchived param.Opt[bool] `json:"exclude_archived,omitzero"`
	// Include direct messages (im) when listing conversations.
	IncludeDms param.Opt[bool] `json:"include_dms,omitzero"`
	// Include group DMs (mpim) when listing conversations.
	IncludeGroupDms param.Opt[bool] `json:"include_group_dms,omitzero"`
	// Include private channels when constructing Slack 'types'. Defaults to False to
	// preserve existing cassette query params.
	IncludePrivate param.Opt[bool] `json:"include_private,omitzero"`
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// List of Slack channels to include (by id, name, or #name).
	Channels []string `json:"channels,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsSlack) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsSlack
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsSlack) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for vault
type MemorySearchParamsOptionsVault struct {
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsVault) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsVault
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsVault) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Search options for Web Crawler
type MemorySearchParamsOptionsWebCrawler struct {
	// The URL to crawl
	URL param.Opt[string] `json:"url,omitzero"`
	// Maximum depth to crawl from the starting URL
	MaxDepth param.Opt[int64] `json:"max_depth,omitzero"`
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsWebCrawler) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsWebCrawler
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsWebCrawler) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryUploadParams struct {
	// The file to ingest.
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// The collection to add the document to — deprecated, set the collection using
	// metadata instead.
	Collection param.Opt[string] `json:"collection,omitzero"`
	// Custom metadata as JSON string for filtering. Keys must be alphanumeric with
	// underscores, max 64 chars. Values must be string, number, or boolean.
	Metadata param.Opt[string] `json:"metadata,omitzero"`
	paramObj
}

func (r MemoryUploadParams) MarshalMultipart() (data []byte, contentType string, err error) {
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
