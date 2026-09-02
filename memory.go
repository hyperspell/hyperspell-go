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

// Delete a memory accessible to the authenticated credential.
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
// All items are validated before processing begins. If any item fails validation,
// the entire batch is rejected with a 422 error detailing which items failed and
// why.
//
// Maximum 100 items per request. Each item follows the same schema as the
// single-item /memories/add endpoint.
func (r *MemoryService) AddBulk(ctx context.Context, body MemoryAddBulkParams, opts ...option.RequestOption) (res *MemoryAddBulkResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "memories/add/bulk"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a document by provider and resource ID, including its full hyperdoc
// tree.
func (r *MemoryService) Get(ctx context.Context, resourceID string, params MemoryGetParams, opts ...option.RequestOption) (res *MemoryGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("memories/get/%v/%s", params.Source, url.PathEscape(resourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
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
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryStatusSource `json:"source" api:"required"`
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped", "filtered", "cancelled".
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
	MemoryStatusSourceReddit           MemoryStatusSource = "reddit"
	MemoryStatusSourceNotion           MemoryStatusSource = "notion"
	MemoryStatusSourceSlack            MemoryStatusSource = "slack"
	MemoryStatusSourceGoogleCalendar   MemoryStatusSource = "google_calendar"
	MemoryStatusSourceGoogleMail       MemoryStatusSource = "google_mail"
	MemoryStatusSourceImap             MemoryStatusSource = "imap"
	MemoryStatusSourceGoogleMeet       MemoryStatusSource = "google_meet"
	MemoryStatusSourceBox              MemoryStatusSource = "box"
	MemoryStatusSourceDropbox          MemoryStatusSource = "dropbox"
	MemoryStatusSourceGitHub           MemoryStatusSource = "github"
	MemoryStatusSourceGitlab           MemoryStatusSource = "gitlab"
	MemoryStatusSourceGoogleDrive      MemoryStatusSource = "google_drive"
	MemoryStatusSourceVault            MemoryStatusSource = "vault"
	MemoryStatusSourceWebCrawler       MemoryStatusSource = "web_crawler"
	MemoryStatusSourceTrace            MemoryStatusSource = "trace"
	MemoryStatusSourceMicrosoftOutlook MemoryStatusSource = "microsoft_outlook"
	MemoryStatusSourceMicrosoftTeams   MemoryStatusSource = "microsoft_teams"
	MemoryStatusSourceGranola          MemoryStatusSource = "granola"
	MemoryStatusSourceFathom           MemoryStatusSource = "fathom"
	MemoryStatusSourceFireflies        MemoryStatusSource = "fireflies"
	MemoryStatusSourceFigma            MemoryStatusSource = "figma"
	MemoryStatusSourceLinear           MemoryStatusSource = "linear"
	MemoryStatusSourceHubspot          MemoryStatusSource = "hubspot"
	MemoryStatusSourceSalesforce       MemoryStatusSource = "salesforce"
	MemoryStatusSourceCoda             MemoryStatusSource = "coda"
	MemoryStatusSourceConfluence       MemoryStatusSource = "confluence"
	MemoryStatusSourceJira             MemoryStatusSource = "jira"
	MemoryStatusSourceMetabase         MemoryStatusSource = "metabase"
	MemoryStatusSourceGong             MemoryStatusSource = "gong"
	MemoryStatusSourceClickup          MemoryStatusSource = "clickup"
	MemoryStatusSourceLightfield       MemoryStatusSource = "lightfield"
	MemoryStatusSourcePylon            MemoryStatusSource = "pylon"
	MemoryStatusSourceFellow           MemoryStatusSource = "fellow"
	MemoryStatusSourceOdoo             MemoryStatusSource = "odoo"
	MemoryStatusSourceExternalMcp      MemoryStatusSource = "external_mcp"
)

type MemoryStatusStatus string

const (
	MemoryStatusStatusPending       MemoryStatusStatus = "pending"
	MemoryStatusStatusProcessing    MemoryStatusStatus = "processing"
	MemoryStatusStatusCompleted     MemoryStatusStatus = "completed"
	MemoryStatusStatusFailed        MemoryStatusStatus = "failed"
	MemoryStatusStatusPendingReview MemoryStatusStatus = "pending_review"
	MemoryStatusStatusSkipped       MemoryStatusStatus = "skipped"
	MemoryStatusStatusFiltered      MemoryStatusStatus = "filtered"
	MemoryStatusStatusCancelled     MemoryStatusStatus = "cancelled"
)

// A document-shaped API response containing the hyperdoc tree.
type MemoryListResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// through `children` for the body.
	Document   MemoryListResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                          `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryListResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// Extracted memories (chunks with summaries) for this document, in document order.
	// Present only when explicitly requested via `include_chunks`; omitted otherwise.
	Chunks []MemoryListResponseChunk `json:"chunks" api:"nullable"`
	// The document's collection, if any.
	Collection string `json:"collection" api:"nullable"`
	// The document's own date (e.g. email sent date, event date).
	DocumentDate time.Time `json:"document_date" api:"nullable" format:"date-time"`
	// When Hyperspell first indexed the document.
	IngestedAt time.Time `json:"ingested_at" api:"nullable" format:"date-time"`
	// When the source document was last modified, if supplied by the source.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Filterable custom metadata attached to the document.
	Metadata map[string]any `json:"metadata"`
	// Indexing status of the document.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped", "filtered", "cancelled".
	Status MemoryListResponseStatus `json:"status" api:"nullable"`
	// Human-readable document title.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document       respjson.Field
		ResourceID     respjson.Field
		Source         respjson.Field
		Type           respjson.Field
		Chunks         respjson.Field
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
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal],
// [MemoryListResponseDocumentInvoice].
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
	// [[]shared.DealChildUnion], [[]MemoryListResponseDocumentInvoiceChildUnion]
	Children MemoryListResponseDocumentUnionChildren `json:"children"`
	// This field is from variant [shared.Document].
	Metadata shared.Metadata `json:"metadata"`
	Text     string          `json:"text"`
	Title    string          `json:"title"`
	// Any of "document", "website", "task", "person", "message", "event", "file",
	// "conversation", "trace", "transcript", "company", "deal", "invoice".
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
	DueAt    time.Time        `json:"due_at"`
	// This field is from variant [shared.Task].
	Priority shared.TaskPriority `json:"priority"`
	Status   string              `json:"status"`
	Address  string              `json:"address"`
	// This field is from variant [shared.Person].
	AltNames []string `json:"alt_names"`
	// This field is from variant [shared.Person].
	BuyingRoles []string `json:"buying_roles"`
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
	EmploymentRole string `json:"employment_role"`
	// This field is from variant [shared.Person].
	EmploymentSeniority string `json:"employment_seniority"`
	// This field is from variant [shared.Person].
	EmploymentSubRole string `json:"employment_sub_role"`
	Industry          string `json:"industry"`
	// This field is from variant [shared.Person].
	IsAppUser bool `json:"is_app_user"`
	// This field is from variant [shared.Person].
	IsBot bool `json:"is_bot"`
	// This field is from variant [shared.Person].
	JobTitle string `json:"job_title"`
	// This field is from variant [shared.Person].
	LastSalesActivityAt string `json:"last_sales_activity_at"`
	// This field is from variant [shared.Person].
	LastSalesActivityType string `json:"last_sales_activity_type"`
	// This field is from variant [shared.Person].
	LeadStatus string `json:"lead_status"`
	// This field is from variant [shared.Person].
	LifecycleStage string `json:"lifecycle_stage"`
	// This field is from variant [shared.Person].
	LinkURLs []string `json:"link_urls"`
	// This field is from variant [shared.Person].
	LinkedinURL string `json:"linkedin_url"`
	// This field is from variant [shared.Person].
	MarketingContactStatus string `json:"marketing_contact_status"`
	Name                   string `json:"name"`
	// This field is from variant [shared.Person].
	OriginalSource string `json:"original_source"`
	// This field is from variant [shared.Person].
	Persona      string   `json:"persona"`
	PhoneNumbers []string `json:"phone_numbers"`
	Tags         []string `json:"tags"`
	Timezone     string   `json:"timezone"`
	// This field is from variant [shared.Person].
	Username string `json:"username"`
	// This field is from variant [shared.Person].
	Website string `json:"website"`
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
	ContentTruncated bool `json:"content_truncated"`
	// This field is from variant [shared.File].
	Path         []string        `json:"path"`
	Participants []shared.Person `json:"participants"`
	// This field is from variant [shared.Transcript].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [shared.Transcript].
	StartedAt  time.Time `json:"started_at"`
	ContactIDs []string  `json:"contact_ids"`
	// This field is from variant [shared.Company].
	Employees int64 `json:"employees"`
	// This field is from variant [shared.Company].
	IsActive bool `json:"is_active"`
	// This field is from variant [shared.Company].
	Websites []string `json:"websites"`
	// This field is from variant [shared.Deal].
	Amount float64 `json:"amount"`
	// This field is from variant [shared.Deal].
	ClosedAt time.Time `json:"closed_at"`
	Currency string    `json:"currency"`
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
	// This field is from variant [MemoryListResponseDocumentInvoice].
	AttachmentNames []string `json:"attachment_names"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	BalanceAmount float64 `json:"balance_amount"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	ContactID string `json:"contact_id"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	ContactName string `json:"contact_name"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	InvoiceType string `json:"invoice_type"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	Notes string `json:"notes"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	Number string `json:"number"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	PaidAmount float64 `json:"paid_amount"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	PaidAt time.Time `json:"paid_at"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	PostedAt time.Time `json:"posted_at"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	Reference string `json:"reference"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	RefundAmount float64 `json:"refund_amount"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	RefundReason string `json:"refund_reason"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	RefundedAt time.Time `json:"refunded_at"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	TaxAmount float64 `json:"tax_amount"`
	// This field is from variant [MemoryListResponseDocumentInvoice].
	TotalAmount float64 `json:"total_amount"`
	JSON        struct {
		ID                     respjson.Field
		Children               respjson.Field
		Metadata               respjson.Field
		Text                   respjson.Field
		Title                  respjson.Field
		Type                   respjson.Field
		URL                    respjson.Field
		Description            respjson.Field
		Favicon                respjson.Field
		ImageURL               respjson.Field
		Language               respjson.Field
		Comments               respjson.Field
		DueAt                  respjson.Field
		Priority               respjson.Field
		Status                 respjson.Field
		Address                respjson.Field
		AltNames               respjson.Field
		BuyingRoles            respjson.Field
		Company                respjson.Field
		CompanyIDs             respjson.Field
		DateOfBirth            respjson.Field
		DealIDs                respjson.Field
		Email                  respjson.Field
		Emails                 respjson.Field
		EmploymentRole         respjson.Field
		EmploymentSeniority    respjson.Field
		EmploymentSubRole      respjson.Field
		Industry               respjson.Field
		IsAppUser              respjson.Field
		IsBot                  respjson.Field
		JobTitle               respjson.Field
		LastSalesActivityAt    respjson.Field
		LastSalesActivityType  respjson.Field
		LeadStatus             respjson.Field
		LifecycleStage         respjson.Field
		LinkURLs               respjson.Field
		LinkedinURL            respjson.Field
		MarketingContactStatus respjson.Field
		Name                   respjson.Field
		OriginalSource         respjson.Field
		Persona                respjson.Field
		PhoneNumbers           respjson.Field
		Tags                   respjson.Field
		Timezone               respjson.Field
		Username               respjson.Field
		Website                respjson.Field
		Date                   respjson.Field
		Sender                 respjson.Field
		Channel                respjson.Field
		ExternalID             respjson.Field
		IsSelf                 respjson.Field
		MentionedUsers         respjson.Field
		NumReplies             respjson.Field
		Replies                respjson.Field
		ThreadID               respjson.Field
		UpdatedAt              respjson.Field
		Upvotes                respjson.Field
		Attendees              respjson.Field
		EndAt                  respjson.Field
		Location               respjson.Field
		MeetingURL             respjson.Field
		StartAt                respjson.Field
		ContentType            respjson.Field
		Filename               respjson.Field
		ContentTruncated       respjson.Field
		Path                   respjson.Field
		Participants           respjson.Field
		EndedAt                respjson.Field
		StartedAt              respjson.Field
		ContactIDs             respjson.Field
		Employees              respjson.Field
		IsActive               respjson.Field
		Websites               respjson.Field
		Amount                 respjson.Field
		ClosedAt               respjson.Field
		Currency               respjson.Field
		DealSource             respjson.Field
		LostReason             respjson.Field
		Pipeline               respjson.Field
		Probability            respjson.Field
		Stage                  respjson.Field
		WonReason              respjson.Field
		AttachmentNames        respjson.Field
		BalanceAmount          respjson.Field
		CancelledAt            respjson.Field
		ContactID              respjson.Field
		ContactName            respjson.Field
		InvoiceType            respjson.Field
		Notes                  respjson.Field
		Number                 respjson.Field
		OrganizationID         respjson.Field
		PaidAmount             respjson.Field
		PaidAt                 respjson.Field
		PostedAt               respjson.Field
		Reference              respjson.Field
		RefundAmount           respjson.Field
		RefundReason           respjson.Field
		RefundedAt             respjson.Field
		TaxAmount              respjson.Field
		TotalAmount            respjson.Field
		raw                    string
	} `json:"-"`
}

// anyMemoryListResponseDocument is implemented by each variant of
// [MemoryListResponseDocumentUnion] to add type safety for the return type of
// [MemoryListResponseDocumentUnion.AsAny]
type anyMemoryListResponseDocument interface {
	ImplMemoryListResponseDocumentUnion()
}

func (MemoryListResponseDocumentInvoice) ImplMemoryListResponseDocumentUnion() {}

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
//	case hyperspell.MemoryListResponseDocumentInvoice:
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
	case "invoice":
		return u.AsInvoice()
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

func (u MemoryListResponseDocumentUnion) AsInvoice() (v MemoryListResponseDocumentInvoice) {
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

// A customer invoice, vendor bill, or credit memo.
//
// Line items are included in `children`.
type MemoryListResponseDocumentInvoice struct {
	ID              string                                        `json:"id"`
	AttachmentNames []string                                      `json:"attachment_names" api:"nullable"`
	BalanceAmount   float64                                       `json:"balance_amount" api:"nullable"`
	CancelledAt     time.Time                                     `json:"cancelled_at" api:"nullable" format:"date-time"`
	Children        []MemoryListResponseDocumentInvoiceChildUnion `json:"children"`
	ContactID       string                                        `json:"contact_id" api:"nullable"`
	ContactName     string                                        `json:"contact_name" api:"nullable"`
	Currency        string                                        `json:"currency" api:"nullable"`
	DueAt           time.Time                                     `json:"due_at" api:"nullable" format:"date-time"`
	InvoiceType     string                                        `json:"invoice_type" api:"nullable"`
	// Optional annotations carried by a hyperdoc node.
	//
	// Includes source provenance and human edit attribution. Unset metadata is omitted
	// from serialized responses.
	Metadata       shared.Metadata `json:"metadata" api:"nullable"`
	Notes          string          `json:"notes" api:"nullable"`
	Number         string          `json:"number" api:"nullable"`
	OrganizationID string          `json:"organization_id" api:"nullable"`
	PaidAmount     float64         `json:"paid_amount" api:"nullable"`
	PaidAt         time.Time       `json:"paid_at" api:"nullable" format:"date-time"`
	PostedAt       time.Time       `json:"posted_at" api:"nullable" format:"date-time"`
	Reference      string          `json:"reference" api:"nullable"`
	RefundAmount   float64         `json:"refund_amount" api:"nullable"`
	RefundReason   string          `json:"refund_reason" api:"nullable"`
	RefundedAt     time.Time       `json:"refunded_at" api:"nullable" format:"date-time"`
	Status         string          `json:"status" api:"nullable"`
	TaxAmount      float64         `json:"tax_amount" api:"nullable"`
	Text           string          `json:"text" api:"nullable"`
	TotalAmount    float64         `json:"total_amount" api:"nullable"`
	// Any of "invoice".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttachmentNames respjson.Field
		BalanceAmount   respjson.Field
		CancelledAt     respjson.Field
		Children        respjson.Field
		ContactID       respjson.Field
		ContactName     respjson.Field
		Currency        respjson.Field
		DueAt           respjson.Field
		InvoiceType     respjson.Field
		Metadata        respjson.Field
		Notes           respjson.Field
		Number          respjson.Field
		OrganizationID  respjson.Field
		PaidAmount      respjson.Field
		PaidAt          respjson.Field
		PostedAt        respjson.Field
		Reference       respjson.Field
		RefundAmount    respjson.Field
		RefundReason    respjson.Field
		RefundedAt      respjson.Field
		Status          respjson.Field
		TaxAmount       respjson.Field
		Text            respjson.Field
		TotalAmount     respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryListResponseDocumentInvoice) RawJSON() string { return r.JSON.raw }
func (r *MemoryListResponseDocumentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryListResponseDocumentInvoiceChildUnion contains all possible properties and
// values from [shared.Blob], [shared.Callout], [shared.Chunk], [shared.Code],
// [shared.Comment], [shared.Divider], [shared.Equation], [shared.Footnote],
// [shared.Heading], [shared.Image], [shared.Link], [shared.LineBreak],
// [shared.List], [shared.ListItem], [shared.Page], [shared.Paragraph],
// [shared.Quote], [shared.Table], [shared.TableCell], [shared.TableRow],
// [shared.Text], [shared.ToDo], [shared.ToolCall], [shared.ToolResult],
// [shared.TraceMessage], [shared.Utterance].
//
// Use the [MemoryListResponseDocumentInvoiceChildUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MemoryListResponseDocumentInvoiceChildUnion struct {
	// This field is from variant [shared.Blob].
	Data string `json:"data"`
	// This field is from variant [shared.Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [shared.Blob].
	Metadata shared.Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "page", "paragraph", "quote", "table", "table_cell", "table_row", "text",
	// "todo", "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]shared.CalloutChildUnion],
	// [[]shared.ChunkChildUnion], [[]shared.EquationChildUnion],
	// [[]shared.FootnoteChildUnion], [[]shared.HeadingChildUnion],
	// [[]shared.ListChildUnion], [[]shared.ListItemChildUnion],
	// [[]shared.PageChildUnion], [[]shared.ParagraphChildUnion],
	// [[]shared.QuoteChildUnion], [[]shared.TableRow], [[]shared.TableCellChildUnion],
	// [[]shared.TableCell], [[]shared.ToDoChildUnion]
	Children MemoryListResponseDocumentInvoiceChildUnionChildren `json:"children"`
	Text     string                                              `json:"text"`
	// This field is from variant [shared.Callout].
	Title string `json:"title"`
	// This field is from variant [shared.Code].
	Language string `json:"language"`
	// This field is from variant [shared.Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [shared.Heading].
	Level int64 `json:"level"`
	// This field is from variant [shared.Image].
	Src string `json:"src"`
	// This field is from variant [shared.Link].
	URL string `json:"url"`
	// This field is from variant [shared.List].
	Ordered bool `json:"ordered"`
	// This field is from variant [shared.Page].
	PageNumber int64 `json:"page_number"`
	// This field is from variant [shared.Page].
	PreviewURL string `json:"preview_url"`
	// This field is from variant [shared.Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [shared.TableCell].
	Align shared.TableCellAlign `json:"align"`
	// This field is from variant [shared.Text].
	Marks []string `json:"marks"`
	// This field is from variant [shared.ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [shared.ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [shared.ToolResult].
	Output shared.ToolResultOutputUnion `json:"output"`
	// This field is from variant [shared.ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [shared.TraceMessage].
	MessageType shared.TraceMessageMessageType `json:"message_type"`
	// This field is from variant [shared.TraceMessage].
	Role shared.TraceMessageRole `json:"role"`
	// This field is from variant [shared.TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [shared.Utterance].
	End float64 `json:"end"`
	// This field is from variant [shared.Utterance].
	Speaker shared.Person `json:"speaker"`
	// This field is from variant [shared.Utterance].
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
		PageNumber  respjson.Field
		PreviewURL  respjson.Field
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

// anyMemoryListResponseDocumentInvoiceChild is implemented by each variant of
// [MemoryListResponseDocumentInvoiceChildUnion] to add type safety for the return
// type of [MemoryListResponseDocumentInvoiceChildUnion.AsAny]
type anyMemoryListResponseDocumentInvoiceChild interface {
	ImplMemoryListResponseDocumentInvoiceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MemoryListResponseDocumentInvoiceChildUnion.AsAny().(type) {
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
//	case shared.Page:
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
func (u MemoryListResponseDocumentInvoiceChildUnion) AsAny() anyMemoryListResponseDocumentInvoiceChild {
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
	case "page":
		return u.AsPage()
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

func (u MemoryListResponseDocumentInvoiceChildUnion) AsBlob() (v shared.Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsCallout() (v shared.Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsChunk() (v shared.Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsCode() (v shared.Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsComment() (v shared.Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsDivider() (v shared.Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsEquation() (v shared.Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsFootnote() (v shared.Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsHeading() (v shared.Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsImage() (v shared.Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsLink() (v shared.Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsLineBreak() (v shared.LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsList() (v shared.List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsListItem() (v shared.ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsPage() (v shared.Page) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsParagraph() (v shared.Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsQuote() (v shared.Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsTable() (v shared.Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsTableCell() (v shared.TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsTableRow() (v shared.TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsText() (v shared.Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsTodo() (v shared.ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsToolCall() (v shared.ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsToolResult() (v shared.ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsTraceMessage() (v shared.TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryListResponseDocumentInvoiceChildUnion) AsUtterance() (v shared.Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemoryListResponseDocumentInvoiceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *MemoryListResponseDocumentInvoiceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryListResponseDocumentInvoiceChildUnionChildren is an implicit subunion of
// [MemoryListResponseDocumentInvoiceChildUnion].
// MemoryListResponseDocumentInvoiceChildUnionChildren provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MemoryListResponseDocumentInvoiceChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type MemoryListResponseDocumentInvoiceChildUnionChildren struct {
	// This field will be present if the value is a [[]shared.CalloutChildUnion]
	// instead of an object.
	OfChildren []shared.CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MemoryListResponseDocumentInvoiceChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryListResponseSource string

const (
	MemoryListResponseSourceReddit           MemoryListResponseSource = "reddit"
	MemoryListResponseSourceNotion           MemoryListResponseSource = "notion"
	MemoryListResponseSourceSlack            MemoryListResponseSource = "slack"
	MemoryListResponseSourceGoogleCalendar   MemoryListResponseSource = "google_calendar"
	MemoryListResponseSourceGoogleMail       MemoryListResponseSource = "google_mail"
	MemoryListResponseSourceImap             MemoryListResponseSource = "imap"
	MemoryListResponseSourceGoogleMeet       MemoryListResponseSource = "google_meet"
	MemoryListResponseSourceBox              MemoryListResponseSource = "box"
	MemoryListResponseSourceDropbox          MemoryListResponseSource = "dropbox"
	MemoryListResponseSourceGitHub           MemoryListResponseSource = "github"
	MemoryListResponseSourceGitlab           MemoryListResponseSource = "gitlab"
	MemoryListResponseSourceGoogleDrive      MemoryListResponseSource = "google_drive"
	MemoryListResponseSourceVault            MemoryListResponseSource = "vault"
	MemoryListResponseSourceWebCrawler       MemoryListResponseSource = "web_crawler"
	MemoryListResponseSourceTrace            MemoryListResponseSource = "trace"
	MemoryListResponseSourceMicrosoftOutlook MemoryListResponseSource = "microsoft_outlook"
	MemoryListResponseSourceMicrosoftTeams   MemoryListResponseSource = "microsoft_teams"
	MemoryListResponseSourceGranola          MemoryListResponseSource = "granola"
	MemoryListResponseSourceFathom           MemoryListResponseSource = "fathom"
	MemoryListResponseSourceFireflies        MemoryListResponseSource = "fireflies"
	MemoryListResponseSourceFigma            MemoryListResponseSource = "figma"
	MemoryListResponseSourceLinear           MemoryListResponseSource = "linear"
	MemoryListResponseSourceHubspot          MemoryListResponseSource = "hubspot"
	MemoryListResponseSourceSalesforce       MemoryListResponseSource = "salesforce"
	MemoryListResponseSourceCoda             MemoryListResponseSource = "coda"
	MemoryListResponseSourceConfluence       MemoryListResponseSource = "confluence"
	MemoryListResponseSourceJira             MemoryListResponseSource = "jira"
	MemoryListResponseSourceMetabase         MemoryListResponseSource = "metabase"
	MemoryListResponseSourceGong             MemoryListResponseSource = "gong"
	MemoryListResponseSourceClickup          MemoryListResponseSource = "clickup"
	MemoryListResponseSourceLightfield       MemoryListResponseSource = "lightfield"
	MemoryListResponseSourcePylon            MemoryListResponseSource = "pylon"
	MemoryListResponseSourceFellow           MemoryListResponseSource = "fellow"
	MemoryListResponseSourceOdoo             MemoryListResponseSource = "odoo"
	MemoryListResponseSourceExternalMcp      MemoryListResponseSource = "external_mcp"
)

// A searchable chunk extracted from a document during ingestion.
//
// `summary` is null when no summary was generated for the chunk.
type MemoryListResponseChunk struct {
	// Stable identifier of the chunk.
	ChunkID string `json:"chunk_id" api:"required"`
	// LLM-generated summary of the chunk, if one was produced.
	Summary string `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryListResponseChunk) RawJSON() string { return r.JSON.raw }
func (r *MemoryListResponseChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexing status of the document.
type MemoryListResponseStatus string

const (
	MemoryListResponseStatusPending       MemoryListResponseStatus = "pending"
	MemoryListResponseStatusProcessing    MemoryListResponseStatus = "processing"
	MemoryListResponseStatusCompleted     MemoryListResponseStatus = "completed"
	MemoryListResponseStatusFailed        MemoryListResponseStatus = "failed"
	MemoryListResponseStatusPendingReview MemoryListResponseStatus = "pending_review"
	MemoryListResponseStatusSkipped       MemoryListResponseStatus = "skipped"
	MemoryListResponseStatusFiltered      MemoryListResponseStatus = "filtered"
	MemoryListResponseStatusCancelled     MemoryListResponseStatus = "cancelled"
)

type MemoryDeleteResponse struct {
	ChunksDeleted int64  `json:"chunks_deleted" api:"required"`
	Message       string `json:"message" api:"required"`
	ResourceID    string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
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
	MemoryDeleteResponseSourceReddit           MemoryDeleteResponseSource = "reddit"
	MemoryDeleteResponseSourceNotion           MemoryDeleteResponseSource = "notion"
	MemoryDeleteResponseSourceSlack            MemoryDeleteResponseSource = "slack"
	MemoryDeleteResponseSourceGoogleCalendar   MemoryDeleteResponseSource = "google_calendar"
	MemoryDeleteResponseSourceGoogleMail       MemoryDeleteResponseSource = "google_mail"
	MemoryDeleteResponseSourceImap             MemoryDeleteResponseSource = "imap"
	MemoryDeleteResponseSourceGoogleMeet       MemoryDeleteResponseSource = "google_meet"
	MemoryDeleteResponseSourceBox              MemoryDeleteResponseSource = "box"
	MemoryDeleteResponseSourceDropbox          MemoryDeleteResponseSource = "dropbox"
	MemoryDeleteResponseSourceGitHub           MemoryDeleteResponseSource = "github"
	MemoryDeleteResponseSourceGitlab           MemoryDeleteResponseSource = "gitlab"
	MemoryDeleteResponseSourceGoogleDrive      MemoryDeleteResponseSource = "google_drive"
	MemoryDeleteResponseSourceVault            MemoryDeleteResponseSource = "vault"
	MemoryDeleteResponseSourceWebCrawler       MemoryDeleteResponseSource = "web_crawler"
	MemoryDeleteResponseSourceTrace            MemoryDeleteResponseSource = "trace"
	MemoryDeleteResponseSourceMicrosoftOutlook MemoryDeleteResponseSource = "microsoft_outlook"
	MemoryDeleteResponseSourceMicrosoftTeams   MemoryDeleteResponseSource = "microsoft_teams"
	MemoryDeleteResponseSourceGranola          MemoryDeleteResponseSource = "granola"
	MemoryDeleteResponseSourceFathom           MemoryDeleteResponseSource = "fathom"
	MemoryDeleteResponseSourceFireflies        MemoryDeleteResponseSource = "fireflies"
	MemoryDeleteResponseSourceFigma            MemoryDeleteResponseSource = "figma"
	MemoryDeleteResponseSourceLinear           MemoryDeleteResponseSource = "linear"
	MemoryDeleteResponseSourceHubspot          MemoryDeleteResponseSource = "hubspot"
	MemoryDeleteResponseSourceSalesforce       MemoryDeleteResponseSource = "salesforce"
	MemoryDeleteResponseSourceCoda             MemoryDeleteResponseSource = "coda"
	MemoryDeleteResponseSourceConfluence       MemoryDeleteResponseSource = "confluence"
	MemoryDeleteResponseSourceJira             MemoryDeleteResponseSource = "jira"
	MemoryDeleteResponseSourceMetabase         MemoryDeleteResponseSource = "metabase"
	MemoryDeleteResponseSourceGong             MemoryDeleteResponseSource = "gong"
	MemoryDeleteResponseSourceClickup          MemoryDeleteResponseSource = "clickup"
	MemoryDeleteResponseSourceLightfield       MemoryDeleteResponseSource = "lightfield"
	MemoryDeleteResponseSourcePylon            MemoryDeleteResponseSource = "pylon"
	MemoryDeleteResponseSourceFellow           MemoryDeleteResponseSource = "fellow"
	MemoryDeleteResponseSourceOdoo             MemoryDeleteResponseSource = "odoo"
	MemoryDeleteResponseSourceExternalMcp      MemoryDeleteResponseSource = "external_mcp"
)

// Response schema for successful bulk ingestion.
type MemoryAddBulkResponse struct {
	// Number of items successfully processed
	Count int64 `json:"count" api:"required"`
	// Status of each ingested item
	Items []MemoryStatus `json:"items" api:"required"`
	// Items not ingested because their resource_id is already owned by another user on
	// this app. Empty in the common case; a non-empty list is a partial success, not
	// an error.
	Skipped []MemoryAddBulkResponseSkipped `json:"skipped"`
	Success bool                           `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Items       respjson.Field
		Skipped     respjson.Field
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

// A bulk item that was neither written nor indexed, with the reason.
//
// `owned_by_another_user` means the resource ID already belongs to another user in
// the app. The bulk endpoint skips that item without modifying the existing
// document. Single-item `/memories/add` returns 409 instead.
type MemoryAddBulkResponseSkipped struct {
	// Why the item was skipped (e.g. 'owned_by_another_user')
	Reason string `json:"reason" api:"required"`
	// Resource ID of the skipped item
	ResourceID string `json:"resource_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Reason      respjson.Field
		ResourceID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryAddBulkResponseSkipped) RawJSON() string { return r.JSON.raw }
func (r *MemoryAddBulkResponseSkipped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document-shaped API response containing the hyperdoc tree.
type MemoryGetResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// through `children` for the body.
	Document   MemoryGetResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                         `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryGetResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// Extracted memories (chunks with summaries) for this document, in document order.
	// Present only when explicitly requested via `include_chunks`; omitted otherwise.
	Chunks []MemoryGetResponseChunk `json:"chunks" api:"nullable"`
	// The document's collection, if any.
	Collection string `json:"collection" api:"nullable"`
	// The document's own date (e.g. email sent date, event date).
	DocumentDate time.Time `json:"document_date" api:"nullable" format:"date-time"`
	// When Hyperspell first indexed the document.
	IngestedAt time.Time `json:"ingested_at" api:"nullable" format:"date-time"`
	// When the source document was last modified, if supplied by the source.
	LastModifiedAt time.Time `json:"last_modified_at" api:"nullable" format:"date-time"`
	// Filterable custom metadata attached to the document.
	Metadata map[string]any `json:"metadata"`
	// Indexing status of the document.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped", "filtered", "cancelled".
	Status MemoryGetResponseStatus `json:"status" api:"nullable"`
	// Human-readable document title.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Document       respjson.Field
		ResourceID     respjson.Field
		Source         respjson.Field
		Type           respjson.Field
		Chunks         respjson.Field
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
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal],
// [MemoryGetResponseDocumentInvoice].
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
	// [[]shared.DealChildUnion], [[]MemoryGetResponseDocumentInvoiceChildUnion]
	Children MemoryGetResponseDocumentUnionChildren `json:"children"`
	// This field is from variant [shared.Document].
	Metadata shared.Metadata `json:"metadata"`
	Text     string          `json:"text"`
	Title    string          `json:"title"`
	// Any of "document", "website", "task", "person", "message", "event", "file",
	// "conversation", "trace", "transcript", "company", "deal", "invoice".
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
	DueAt    time.Time        `json:"due_at"`
	// This field is from variant [shared.Task].
	Priority shared.TaskPriority `json:"priority"`
	Status   string              `json:"status"`
	Address  string              `json:"address"`
	// This field is from variant [shared.Person].
	AltNames []string `json:"alt_names"`
	// This field is from variant [shared.Person].
	BuyingRoles []string `json:"buying_roles"`
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
	EmploymentRole string `json:"employment_role"`
	// This field is from variant [shared.Person].
	EmploymentSeniority string `json:"employment_seniority"`
	// This field is from variant [shared.Person].
	EmploymentSubRole string `json:"employment_sub_role"`
	Industry          string `json:"industry"`
	// This field is from variant [shared.Person].
	IsAppUser bool `json:"is_app_user"`
	// This field is from variant [shared.Person].
	IsBot bool `json:"is_bot"`
	// This field is from variant [shared.Person].
	JobTitle string `json:"job_title"`
	// This field is from variant [shared.Person].
	LastSalesActivityAt string `json:"last_sales_activity_at"`
	// This field is from variant [shared.Person].
	LastSalesActivityType string `json:"last_sales_activity_type"`
	// This field is from variant [shared.Person].
	LeadStatus string `json:"lead_status"`
	// This field is from variant [shared.Person].
	LifecycleStage string `json:"lifecycle_stage"`
	// This field is from variant [shared.Person].
	LinkURLs []string `json:"link_urls"`
	// This field is from variant [shared.Person].
	LinkedinURL string `json:"linkedin_url"`
	// This field is from variant [shared.Person].
	MarketingContactStatus string `json:"marketing_contact_status"`
	Name                   string `json:"name"`
	// This field is from variant [shared.Person].
	OriginalSource string `json:"original_source"`
	// This field is from variant [shared.Person].
	Persona      string   `json:"persona"`
	PhoneNumbers []string `json:"phone_numbers"`
	Tags         []string `json:"tags"`
	Timezone     string   `json:"timezone"`
	// This field is from variant [shared.Person].
	Username string `json:"username"`
	// This field is from variant [shared.Person].
	Website string `json:"website"`
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
	ContentTruncated bool `json:"content_truncated"`
	// This field is from variant [shared.File].
	Path         []string        `json:"path"`
	Participants []shared.Person `json:"participants"`
	// This field is from variant [shared.Transcript].
	EndedAt time.Time `json:"ended_at"`
	// This field is from variant [shared.Transcript].
	StartedAt  time.Time `json:"started_at"`
	ContactIDs []string  `json:"contact_ids"`
	// This field is from variant [shared.Company].
	Employees int64 `json:"employees"`
	// This field is from variant [shared.Company].
	IsActive bool `json:"is_active"`
	// This field is from variant [shared.Company].
	Websites []string `json:"websites"`
	// This field is from variant [shared.Deal].
	Amount float64 `json:"amount"`
	// This field is from variant [shared.Deal].
	ClosedAt time.Time `json:"closed_at"`
	Currency string    `json:"currency"`
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
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	AttachmentNames []string `json:"attachment_names"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	BalanceAmount float64 `json:"balance_amount"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	ContactID string `json:"contact_id"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	ContactName string `json:"contact_name"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	InvoiceType string `json:"invoice_type"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	Notes string `json:"notes"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	Number string `json:"number"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	PaidAmount float64 `json:"paid_amount"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	PaidAt time.Time `json:"paid_at"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	PostedAt time.Time `json:"posted_at"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	Reference string `json:"reference"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	RefundAmount float64 `json:"refund_amount"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	RefundReason string `json:"refund_reason"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	RefundedAt time.Time `json:"refunded_at"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	TaxAmount float64 `json:"tax_amount"`
	// This field is from variant [MemoryGetResponseDocumentInvoice].
	TotalAmount float64 `json:"total_amount"`
	JSON        struct {
		ID                     respjson.Field
		Children               respjson.Field
		Metadata               respjson.Field
		Text                   respjson.Field
		Title                  respjson.Field
		Type                   respjson.Field
		URL                    respjson.Field
		Description            respjson.Field
		Favicon                respjson.Field
		ImageURL               respjson.Field
		Language               respjson.Field
		Comments               respjson.Field
		DueAt                  respjson.Field
		Priority               respjson.Field
		Status                 respjson.Field
		Address                respjson.Field
		AltNames               respjson.Field
		BuyingRoles            respjson.Field
		Company                respjson.Field
		CompanyIDs             respjson.Field
		DateOfBirth            respjson.Field
		DealIDs                respjson.Field
		Email                  respjson.Field
		Emails                 respjson.Field
		EmploymentRole         respjson.Field
		EmploymentSeniority    respjson.Field
		EmploymentSubRole      respjson.Field
		Industry               respjson.Field
		IsAppUser              respjson.Field
		IsBot                  respjson.Field
		JobTitle               respjson.Field
		LastSalesActivityAt    respjson.Field
		LastSalesActivityType  respjson.Field
		LeadStatus             respjson.Field
		LifecycleStage         respjson.Field
		LinkURLs               respjson.Field
		LinkedinURL            respjson.Field
		MarketingContactStatus respjson.Field
		Name                   respjson.Field
		OriginalSource         respjson.Field
		Persona                respjson.Field
		PhoneNumbers           respjson.Field
		Tags                   respjson.Field
		Timezone               respjson.Field
		Username               respjson.Field
		Website                respjson.Field
		Date                   respjson.Field
		Sender                 respjson.Field
		Channel                respjson.Field
		ExternalID             respjson.Field
		IsSelf                 respjson.Field
		MentionedUsers         respjson.Field
		NumReplies             respjson.Field
		Replies                respjson.Field
		ThreadID               respjson.Field
		UpdatedAt              respjson.Field
		Upvotes                respjson.Field
		Attendees              respjson.Field
		EndAt                  respjson.Field
		Location               respjson.Field
		MeetingURL             respjson.Field
		StartAt                respjson.Field
		ContentType            respjson.Field
		Filename               respjson.Field
		ContentTruncated       respjson.Field
		Path                   respjson.Field
		Participants           respjson.Field
		EndedAt                respjson.Field
		StartedAt              respjson.Field
		ContactIDs             respjson.Field
		Employees              respjson.Field
		IsActive               respjson.Field
		Websites               respjson.Field
		Amount                 respjson.Field
		ClosedAt               respjson.Field
		Currency               respjson.Field
		DealSource             respjson.Field
		LostReason             respjson.Field
		Pipeline               respjson.Field
		Probability            respjson.Field
		Stage                  respjson.Field
		WonReason              respjson.Field
		AttachmentNames        respjson.Field
		BalanceAmount          respjson.Field
		CancelledAt            respjson.Field
		ContactID              respjson.Field
		ContactName            respjson.Field
		InvoiceType            respjson.Field
		Notes                  respjson.Field
		Number                 respjson.Field
		OrganizationID         respjson.Field
		PaidAmount             respjson.Field
		PaidAt                 respjson.Field
		PostedAt               respjson.Field
		Reference              respjson.Field
		RefundAmount           respjson.Field
		RefundReason           respjson.Field
		RefundedAt             respjson.Field
		TaxAmount              respjson.Field
		TotalAmount            respjson.Field
		raw                    string
	} `json:"-"`
}

// anyMemoryGetResponseDocument is implemented by each variant of
// [MemoryGetResponseDocumentUnion] to add type safety for the return type of
// [MemoryGetResponseDocumentUnion.AsAny]
type anyMemoryGetResponseDocument interface {
	ImplMemoryGetResponseDocumentUnion()
}

func (MemoryGetResponseDocumentInvoice) ImplMemoryGetResponseDocumentUnion() {}

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
//	case hyperspell.MemoryGetResponseDocumentInvoice:
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
	case "invoice":
		return u.AsInvoice()
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

func (u MemoryGetResponseDocumentUnion) AsInvoice() (v MemoryGetResponseDocumentInvoice) {
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

// A customer invoice, vendor bill, or credit memo.
//
// Line items are included in `children`.
type MemoryGetResponseDocumentInvoice struct {
	ID              string                                       `json:"id"`
	AttachmentNames []string                                     `json:"attachment_names" api:"nullable"`
	BalanceAmount   float64                                      `json:"balance_amount" api:"nullable"`
	CancelledAt     time.Time                                    `json:"cancelled_at" api:"nullable" format:"date-time"`
	Children        []MemoryGetResponseDocumentInvoiceChildUnion `json:"children"`
	ContactID       string                                       `json:"contact_id" api:"nullable"`
	ContactName     string                                       `json:"contact_name" api:"nullable"`
	Currency        string                                       `json:"currency" api:"nullable"`
	DueAt           time.Time                                    `json:"due_at" api:"nullable" format:"date-time"`
	InvoiceType     string                                       `json:"invoice_type" api:"nullable"`
	// Optional annotations carried by a hyperdoc node.
	//
	// Includes source provenance and human edit attribution. Unset metadata is omitted
	// from serialized responses.
	Metadata       shared.Metadata `json:"metadata" api:"nullable"`
	Notes          string          `json:"notes" api:"nullable"`
	Number         string          `json:"number" api:"nullable"`
	OrganizationID string          `json:"organization_id" api:"nullable"`
	PaidAmount     float64         `json:"paid_amount" api:"nullable"`
	PaidAt         time.Time       `json:"paid_at" api:"nullable" format:"date-time"`
	PostedAt       time.Time       `json:"posted_at" api:"nullable" format:"date-time"`
	Reference      string          `json:"reference" api:"nullable"`
	RefundAmount   float64         `json:"refund_amount" api:"nullable"`
	RefundReason   string          `json:"refund_reason" api:"nullable"`
	RefundedAt     time.Time       `json:"refunded_at" api:"nullable" format:"date-time"`
	Status         string          `json:"status" api:"nullable"`
	TaxAmount      float64         `json:"tax_amount" api:"nullable"`
	Text           string          `json:"text" api:"nullable"`
	TotalAmount    float64         `json:"total_amount" api:"nullable"`
	// Any of "invoice".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		AttachmentNames respjson.Field
		BalanceAmount   respjson.Field
		CancelledAt     respjson.Field
		Children        respjson.Field
		ContactID       respjson.Field
		ContactName     respjson.Field
		Currency        respjson.Field
		DueAt           respjson.Field
		InvoiceType     respjson.Field
		Metadata        respjson.Field
		Notes           respjson.Field
		Number          respjson.Field
		OrganizationID  respjson.Field
		PaidAmount      respjson.Field
		PaidAt          respjson.Field
		PostedAt        respjson.Field
		Reference       respjson.Field
		RefundAmount    respjson.Field
		RefundReason    respjson.Field
		RefundedAt      respjson.Field
		Status          respjson.Field
		TaxAmount       respjson.Field
		Text            respjson.Field
		TotalAmount     respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryGetResponseDocumentInvoice) RawJSON() string { return r.JSON.raw }
func (r *MemoryGetResponseDocumentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryGetResponseDocumentInvoiceChildUnion contains all possible properties and
// values from [shared.Blob], [shared.Callout], [shared.Chunk], [shared.Code],
// [shared.Comment], [shared.Divider], [shared.Equation], [shared.Footnote],
// [shared.Heading], [shared.Image], [shared.Link], [shared.LineBreak],
// [shared.List], [shared.ListItem], [shared.Page], [shared.Paragraph],
// [shared.Quote], [shared.Table], [shared.TableCell], [shared.TableRow],
// [shared.Text], [shared.ToDo], [shared.ToolCall], [shared.ToolResult],
// [shared.TraceMessage], [shared.Utterance].
//
// Use the [MemoryGetResponseDocumentInvoiceChildUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type MemoryGetResponseDocumentInvoiceChildUnion struct {
	// This field is from variant [shared.Blob].
	Data string `json:"data"`
	// This field is from variant [shared.Blob].
	Mimetype string `json:"mimetype"`
	ID       string `json:"id"`
	// This field is from variant [shared.Blob].
	Metadata shared.Metadata `json:"metadata"`
	// Any of "blob", "callout", "chunk", "code", "comment", "divider", "equation",
	// "footnote", "heading", "image", "link", "line_break", "list", "list_item",
	// "page", "paragraph", "quote", "table", "table_cell", "table_row", "text",
	// "todo", "tool_call", "tool_result", "trace_message", "utterance".
	Type string `json:"type"`
	// This field is a union of [[]shared.CalloutChildUnion],
	// [[]shared.ChunkChildUnion], [[]shared.EquationChildUnion],
	// [[]shared.FootnoteChildUnion], [[]shared.HeadingChildUnion],
	// [[]shared.ListChildUnion], [[]shared.ListItemChildUnion],
	// [[]shared.PageChildUnion], [[]shared.ParagraphChildUnion],
	// [[]shared.QuoteChildUnion], [[]shared.TableRow], [[]shared.TableCellChildUnion],
	// [[]shared.TableCell], [[]shared.ToDoChildUnion]
	Children MemoryGetResponseDocumentInvoiceChildUnionChildren `json:"children"`
	Text     string                                             `json:"text"`
	// This field is from variant [shared.Callout].
	Title string `json:"title"`
	// This field is from variant [shared.Code].
	Language string `json:"language"`
	// This field is from variant [shared.Comment].
	CreatedAt time.Time `json:"created_at"`
	// This field is from variant [shared.Heading].
	Level int64 `json:"level"`
	// This field is from variant [shared.Image].
	Src string `json:"src"`
	// This field is from variant [shared.Link].
	URL string `json:"url"`
	// This field is from variant [shared.List].
	Ordered bool `json:"ordered"`
	// This field is from variant [shared.Page].
	PageNumber int64 `json:"page_number"`
	// This field is from variant [shared.Page].
	PreviewURL string `json:"preview_url"`
	// This field is from variant [shared.Table].
	HasHeader bool `json:"has_header"`
	// This field is from variant [shared.TableCell].
	Align shared.TableCellAlign `json:"align"`
	// This field is from variant [shared.Text].
	Marks []string `json:"marks"`
	// This field is from variant [shared.ToDo].
	Checked    bool   `json:"checked"`
	ToolCallID string `json:"tool_call_id"`
	ToolName   string `json:"tool_name"`
	// This field is from variant [shared.ToolCall].
	Args map[string]any `json:"args"`
	// This field is from variant [shared.ToolResult].
	Output shared.ToolResultOutputUnion `json:"output"`
	// This field is from variant [shared.ToolResult].
	IsError bool `json:"is_error"`
	// This field is from variant [shared.TraceMessage].
	MessageType shared.TraceMessageMessageType `json:"message_type"`
	// This field is from variant [shared.TraceMessage].
	Role shared.TraceMessageRole `json:"role"`
	// This field is from variant [shared.TraceMessage].
	Timestamp time.Time `json:"timestamp"`
	// This field is from variant [shared.Utterance].
	End float64 `json:"end"`
	// This field is from variant [shared.Utterance].
	Speaker shared.Person `json:"speaker"`
	// This field is from variant [shared.Utterance].
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
		PageNumber  respjson.Field
		PreviewURL  respjson.Field
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

// anyMemoryGetResponseDocumentInvoiceChild is implemented by each variant of
// [MemoryGetResponseDocumentInvoiceChildUnion] to add type safety for the return
// type of [MemoryGetResponseDocumentInvoiceChildUnion.AsAny]
type anyMemoryGetResponseDocumentInvoiceChild interface {
	ImplMemoryGetResponseDocumentInvoiceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := MemoryGetResponseDocumentInvoiceChildUnion.AsAny().(type) {
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
//	case shared.Page:
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
func (u MemoryGetResponseDocumentInvoiceChildUnion) AsAny() anyMemoryGetResponseDocumentInvoiceChild {
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
	case "page":
		return u.AsPage()
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

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsBlob() (v shared.Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsCallout() (v shared.Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsChunk() (v shared.Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsCode() (v shared.Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsComment() (v shared.Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsDivider() (v shared.Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsEquation() (v shared.Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsFootnote() (v shared.Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsHeading() (v shared.Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsImage() (v shared.Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsLink() (v shared.Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsLineBreak() (v shared.LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsList() (v shared.List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsListItem() (v shared.ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsPage() (v shared.Page) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsParagraph() (v shared.Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsQuote() (v shared.Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsTable() (v shared.Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsTableCell() (v shared.TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsTableRow() (v shared.TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsText() (v shared.Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsTodo() (v shared.ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsToolCall() (v shared.ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsToolResult() (v shared.ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsTraceMessage() (v shared.TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MemoryGetResponseDocumentInvoiceChildUnion) AsUtterance() (v shared.Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MemoryGetResponseDocumentInvoiceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *MemoryGetResponseDocumentInvoiceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MemoryGetResponseDocumentInvoiceChildUnionChildren is an implicit subunion of
// [MemoryGetResponseDocumentInvoiceChildUnion].
// MemoryGetResponseDocumentInvoiceChildUnionChildren provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [MemoryGetResponseDocumentInvoiceChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type MemoryGetResponseDocumentInvoiceChildUnionChildren struct {
	// This field will be present if the value is a [[]shared.CalloutChildUnion]
	// instead of an object.
	OfChildren []shared.CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *MemoryGetResponseDocumentInvoiceChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryGetResponseSource string

const (
	MemoryGetResponseSourceReddit           MemoryGetResponseSource = "reddit"
	MemoryGetResponseSourceNotion           MemoryGetResponseSource = "notion"
	MemoryGetResponseSourceSlack            MemoryGetResponseSource = "slack"
	MemoryGetResponseSourceGoogleCalendar   MemoryGetResponseSource = "google_calendar"
	MemoryGetResponseSourceGoogleMail       MemoryGetResponseSource = "google_mail"
	MemoryGetResponseSourceImap             MemoryGetResponseSource = "imap"
	MemoryGetResponseSourceGoogleMeet       MemoryGetResponseSource = "google_meet"
	MemoryGetResponseSourceBox              MemoryGetResponseSource = "box"
	MemoryGetResponseSourceDropbox          MemoryGetResponseSource = "dropbox"
	MemoryGetResponseSourceGitHub           MemoryGetResponseSource = "github"
	MemoryGetResponseSourceGitlab           MemoryGetResponseSource = "gitlab"
	MemoryGetResponseSourceGoogleDrive      MemoryGetResponseSource = "google_drive"
	MemoryGetResponseSourceVault            MemoryGetResponseSource = "vault"
	MemoryGetResponseSourceWebCrawler       MemoryGetResponseSource = "web_crawler"
	MemoryGetResponseSourceTrace            MemoryGetResponseSource = "trace"
	MemoryGetResponseSourceMicrosoftOutlook MemoryGetResponseSource = "microsoft_outlook"
	MemoryGetResponseSourceMicrosoftTeams   MemoryGetResponseSource = "microsoft_teams"
	MemoryGetResponseSourceGranola          MemoryGetResponseSource = "granola"
	MemoryGetResponseSourceFathom           MemoryGetResponseSource = "fathom"
	MemoryGetResponseSourceFireflies        MemoryGetResponseSource = "fireflies"
	MemoryGetResponseSourceFigma            MemoryGetResponseSource = "figma"
	MemoryGetResponseSourceLinear           MemoryGetResponseSource = "linear"
	MemoryGetResponseSourceHubspot          MemoryGetResponseSource = "hubspot"
	MemoryGetResponseSourceSalesforce       MemoryGetResponseSource = "salesforce"
	MemoryGetResponseSourceCoda             MemoryGetResponseSource = "coda"
	MemoryGetResponseSourceConfluence       MemoryGetResponseSource = "confluence"
	MemoryGetResponseSourceJira             MemoryGetResponseSource = "jira"
	MemoryGetResponseSourceMetabase         MemoryGetResponseSource = "metabase"
	MemoryGetResponseSourceGong             MemoryGetResponseSource = "gong"
	MemoryGetResponseSourceClickup          MemoryGetResponseSource = "clickup"
	MemoryGetResponseSourceLightfield       MemoryGetResponseSource = "lightfield"
	MemoryGetResponseSourcePylon            MemoryGetResponseSource = "pylon"
	MemoryGetResponseSourceFellow           MemoryGetResponseSource = "fellow"
	MemoryGetResponseSourceOdoo             MemoryGetResponseSource = "odoo"
	MemoryGetResponseSourceExternalMcp      MemoryGetResponseSource = "external_mcp"
)

// A searchable chunk extracted from a document during ingestion.
//
// `summary` is null when no summary was generated for the chunk.
type MemoryGetResponseChunk struct {
	// Stable identifier of the chunk.
	ChunkID string `json:"chunk_id" api:"required"`
	// LLM-generated summary of the chunk, if one was produced.
	Summary string `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChunkID     respjson.Field
		Summary     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryGetResponseChunk) RawJSON() string { return r.JSON.raw }
func (r *MemoryGetResponseChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexing status of the document.
type MemoryGetResponseStatus string

const (
	MemoryGetResponseStatusPending       MemoryGetResponseStatus = "pending"
	MemoryGetResponseStatusProcessing    MemoryGetResponseStatus = "processing"
	MemoryGetResponseStatusCompleted     MemoryGetResponseStatus = "completed"
	MemoryGetResponseStatusFailed        MemoryGetResponseStatus = "failed"
	MemoryGetResponseStatusPendingReview MemoryGetResponseStatus = "pending_review"
	MemoryGetResponseStatusSkipped       MemoryGetResponseStatus = "skipped"
	MemoryGetResponseStatusFiltered      MemoryGetResponseStatus = "filtered"
	MemoryGetResponseStatusCancelled     MemoryGetResponseStatus = "cancelled"
)

type MemoryStatusResponse struct {
	Providers    map[string]map[string]int64       `json:"providers" api:"required"`
	Total        map[string]int64                  `json:"total" api:"required"`
	Integrations []MemoryStatusResponseIntegration `json:"integrations"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Providers    respjson.Field
		Total        respjson.Field
		Integrations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Health summary for a configured integration.
//
// `provider` uses lowercase snake_case naming (e.g. `google_drive`).
type MemoryStatusResponseIntegration struct {
	Connections []MemoryStatusResponseIntegrationConnection `json:"connections" api:"required"`
	// The current error for a connection.
	//
	// `detail` contains a sanitized summary suitable for display.
	Error         MemoryStatusResponseIntegrationError `json:"error" api:"required"`
	IntegrationID string                               `json:"integration_id" api:"required" format:"uuid"`
	LastSyncedAt  time.Time                            `json:"last_synced_at" api:"required" format:"date-time"`
	Provider      string                               `json:"provider" api:"required"`
	// Current health status of a connection or integration.
	//
	// Any of "broken", "stalled", "error", "rate_limited", "syncing", "connected",
	// "live", "never_synced", "not_connected".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connections   respjson.Field
		Error         respjson.Field
		IntegrationID respjson.Field
		LastSyncedAt  respjson.Field
		Provider      respjson.Field
		Status        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current health of one connection.
type MemoryStatusResponseIntegrationConnection struct {
	ID string `json:"id" api:"required"`
	// The current error for a connection.
	//
	// `detail` contains a sanitized summary suitable for display.
	Error          MemoryStatusResponseIntegrationConnectionError `json:"error" api:"required"`
	Label          string                                         `json:"label" api:"required"`
	LastActivityAt time.Time                                      `json:"last_activity_at" api:"required" format:"date-time"`
	LastSyncedAt   time.Time                                      `json:"last_synced_at" api:"required" format:"date-time"`
	// Current health status of a connection or integration.
	//
	// Any of "broken", "stalled", "error", "rate_limited", "syncing", "connected",
	// "live", "never_synced", "not_connected".
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Error          respjson.Field
		Label          respjson.Field
		LastActivityAt respjson.Field
		LastSyncedAt   respjson.Field
		Status         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponseIntegrationConnection) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponseIntegrationConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current error for a connection.
//
// `detail` contains a sanitized summary suitable for display.
type MemoryStatusResponseIntegrationConnectionError struct {
	At     time.Time `json:"at" api:"required" format:"date-time"`
	Detail string    `json:"detail" api:"required"`
	// Classification of the most recent synchronization or indexing failure.
	//
	// Any of "auth", "rate_limited", "provider", "internal".
	Kind    string    `json:"kind" api:"required"`
	Origin  string    `json:"origin" api:"nullable"`
	RetryAt time.Time `json:"retry_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At          respjson.Field
		Detail      respjson.Field
		Kind        respjson.Field
		Origin      respjson.Field
		RetryAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponseIntegrationConnectionError) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponseIntegrationConnectionError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current error for a connection.
//
// `detail` contains a sanitized summary suitable for display.
type MemoryStatusResponseIntegrationError struct {
	At     time.Time `json:"at" api:"required" format:"date-time"`
	Detail string    `json:"detail" api:"required"`
	// Classification of the most recent synchronization or indexing failure.
	//
	// Any of "auth", "rate_limited", "provider", "internal".
	Kind    string    `json:"kind" api:"required"`
	Origin  string    `json:"origin" api:"nullable"`
	RetryAt time.Time `json:"retry_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		At          respjson.Field
		Detail      respjson.Field
		Kind        respjson.Field
		Origin      respjson.Field
		RetryAt     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MemoryStatusResponseIntegrationError) RawJSON() string { return r.JSON.raw }
func (r *MemoryStatusResponseIntegrationError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemoryUpdateParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
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
	MemoryUpdateParamsSourceReddit           MemoryUpdateParamsSource = "reddit"
	MemoryUpdateParamsSourceNotion           MemoryUpdateParamsSource = "notion"
	MemoryUpdateParamsSourceSlack            MemoryUpdateParamsSource = "slack"
	MemoryUpdateParamsSourceGoogleCalendar   MemoryUpdateParamsSource = "google_calendar"
	MemoryUpdateParamsSourceGoogleMail       MemoryUpdateParamsSource = "google_mail"
	MemoryUpdateParamsSourceImap             MemoryUpdateParamsSource = "imap"
	MemoryUpdateParamsSourceGoogleMeet       MemoryUpdateParamsSource = "google_meet"
	MemoryUpdateParamsSourceBox              MemoryUpdateParamsSource = "box"
	MemoryUpdateParamsSourceDropbox          MemoryUpdateParamsSource = "dropbox"
	MemoryUpdateParamsSourceGitHub           MemoryUpdateParamsSource = "github"
	MemoryUpdateParamsSourceGitlab           MemoryUpdateParamsSource = "gitlab"
	MemoryUpdateParamsSourceGoogleDrive      MemoryUpdateParamsSource = "google_drive"
	MemoryUpdateParamsSourceVault            MemoryUpdateParamsSource = "vault"
	MemoryUpdateParamsSourceWebCrawler       MemoryUpdateParamsSource = "web_crawler"
	MemoryUpdateParamsSourceTrace            MemoryUpdateParamsSource = "trace"
	MemoryUpdateParamsSourceMicrosoftOutlook MemoryUpdateParamsSource = "microsoft_outlook"
	MemoryUpdateParamsSourceMicrosoftTeams   MemoryUpdateParamsSource = "microsoft_teams"
	MemoryUpdateParamsSourceGranola          MemoryUpdateParamsSource = "granola"
	MemoryUpdateParamsSourceFathom           MemoryUpdateParamsSource = "fathom"
	MemoryUpdateParamsSourceFireflies        MemoryUpdateParamsSource = "fireflies"
	MemoryUpdateParamsSourceFigma            MemoryUpdateParamsSource = "figma"
	MemoryUpdateParamsSourceLinear           MemoryUpdateParamsSource = "linear"
	MemoryUpdateParamsSourceHubspot          MemoryUpdateParamsSource = "hubspot"
	MemoryUpdateParamsSourceSalesforce       MemoryUpdateParamsSource = "salesforce"
	MemoryUpdateParamsSourceCoda             MemoryUpdateParamsSource = "coda"
	MemoryUpdateParamsSourceConfluence       MemoryUpdateParamsSource = "confluence"
	MemoryUpdateParamsSourceJira             MemoryUpdateParamsSource = "jira"
	MemoryUpdateParamsSourceMetabase         MemoryUpdateParamsSource = "metabase"
	MemoryUpdateParamsSourceGong             MemoryUpdateParamsSource = "gong"
	MemoryUpdateParamsSourceClickup          MemoryUpdateParamsSource = "clickup"
	MemoryUpdateParamsSourceLightfield       MemoryUpdateParamsSource = "lightfield"
	MemoryUpdateParamsSourcePylon            MemoryUpdateParamsSource = "pylon"
	MemoryUpdateParamsSourceFellow           MemoryUpdateParamsSource = "fellow"
	MemoryUpdateParamsSourceOdoo             MemoryUpdateParamsSource = "odoo"
	MemoryUpdateParamsSourceExternalMcp      MemoryUpdateParamsSource = "external_mcp"
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
	// When > 0, include up to this many extracted memories (chunks with summaries) per
	// document in each item's `chunks` field, in document order. 0 (default) omits
	// them.
	IncludeChunks param.Opt[int64] `query:"include_chunks,omitzero" json:"-"`
	Size          param.Opt[int64] `query:"size,omitzero" json:"-"`
	// Filter documents by source.
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryListParamsSource `query:"source,omitzero" json:"-"`
	// Filter documents by status.
	//
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped", "filtered", "cancelled".
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
	MemoryListParamsSourceReddit           MemoryListParamsSource = "reddit"
	MemoryListParamsSourceNotion           MemoryListParamsSource = "notion"
	MemoryListParamsSourceSlack            MemoryListParamsSource = "slack"
	MemoryListParamsSourceGoogleCalendar   MemoryListParamsSource = "google_calendar"
	MemoryListParamsSourceGoogleMail       MemoryListParamsSource = "google_mail"
	MemoryListParamsSourceImap             MemoryListParamsSource = "imap"
	MemoryListParamsSourceGoogleMeet       MemoryListParamsSource = "google_meet"
	MemoryListParamsSourceBox              MemoryListParamsSource = "box"
	MemoryListParamsSourceDropbox          MemoryListParamsSource = "dropbox"
	MemoryListParamsSourceGitHub           MemoryListParamsSource = "github"
	MemoryListParamsSourceGitlab           MemoryListParamsSource = "gitlab"
	MemoryListParamsSourceGoogleDrive      MemoryListParamsSource = "google_drive"
	MemoryListParamsSourceVault            MemoryListParamsSource = "vault"
	MemoryListParamsSourceWebCrawler       MemoryListParamsSource = "web_crawler"
	MemoryListParamsSourceTrace            MemoryListParamsSource = "trace"
	MemoryListParamsSourceMicrosoftOutlook MemoryListParamsSource = "microsoft_outlook"
	MemoryListParamsSourceMicrosoftTeams   MemoryListParamsSource = "microsoft_teams"
	MemoryListParamsSourceGranola          MemoryListParamsSource = "granola"
	MemoryListParamsSourceFathom           MemoryListParamsSource = "fathom"
	MemoryListParamsSourceFireflies        MemoryListParamsSource = "fireflies"
	MemoryListParamsSourceFigma            MemoryListParamsSource = "figma"
	MemoryListParamsSourceLinear           MemoryListParamsSource = "linear"
	MemoryListParamsSourceHubspot          MemoryListParamsSource = "hubspot"
	MemoryListParamsSourceSalesforce       MemoryListParamsSource = "salesforce"
	MemoryListParamsSourceCoda             MemoryListParamsSource = "coda"
	MemoryListParamsSourceConfluence       MemoryListParamsSource = "confluence"
	MemoryListParamsSourceJira             MemoryListParamsSource = "jira"
	MemoryListParamsSourceMetabase         MemoryListParamsSource = "metabase"
	MemoryListParamsSourceGong             MemoryListParamsSource = "gong"
	MemoryListParamsSourceClickup          MemoryListParamsSource = "clickup"
	MemoryListParamsSourceLightfield       MemoryListParamsSource = "lightfield"
	MemoryListParamsSourcePylon            MemoryListParamsSource = "pylon"
	MemoryListParamsSourceFellow           MemoryListParamsSource = "fellow"
	MemoryListParamsSourceOdoo             MemoryListParamsSource = "odoo"
	MemoryListParamsSourceExternalMcp      MemoryListParamsSource = "external_mcp"
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
	MemoryListParamsStatusFiltered      MemoryListParamsStatus = "filtered"
	MemoryListParamsStatusCancelled     MemoryListParamsStatus = "cancelled"
)

type MemoryDeleteParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryDeleteParamsSource `path:"source,omitzero" api:"required" json:"-"`
	paramObj
}

type MemoryDeleteParamsSource string

const (
	MemoryDeleteParamsSourceReddit           MemoryDeleteParamsSource = "reddit"
	MemoryDeleteParamsSourceNotion           MemoryDeleteParamsSource = "notion"
	MemoryDeleteParamsSourceSlack            MemoryDeleteParamsSource = "slack"
	MemoryDeleteParamsSourceGoogleCalendar   MemoryDeleteParamsSource = "google_calendar"
	MemoryDeleteParamsSourceGoogleMail       MemoryDeleteParamsSource = "google_mail"
	MemoryDeleteParamsSourceImap             MemoryDeleteParamsSource = "imap"
	MemoryDeleteParamsSourceGoogleMeet       MemoryDeleteParamsSource = "google_meet"
	MemoryDeleteParamsSourceBox              MemoryDeleteParamsSource = "box"
	MemoryDeleteParamsSourceDropbox          MemoryDeleteParamsSource = "dropbox"
	MemoryDeleteParamsSourceGitHub           MemoryDeleteParamsSource = "github"
	MemoryDeleteParamsSourceGitlab           MemoryDeleteParamsSource = "gitlab"
	MemoryDeleteParamsSourceGoogleDrive      MemoryDeleteParamsSource = "google_drive"
	MemoryDeleteParamsSourceVault            MemoryDeleteParamsSource = "vault"
	MemoryDeleteParamsSourceWebCrawler       MemoryDeleteParamsSource = "web_crawler"
	MemoryDeleteParamsSourceTrace            MemoryDeleteParamsSource = "trace"
	MemoryDeleteParamsSourceMicrosoftOutlook MemoryDeleteParamsSource = "microsoft_outlook"
	MemoryDeleteParamsSourceMicrosoftTeams   MemoryDeleteParamsSource = "microsoft_teams"
	MemoryDeleteParamsSourceGranola          MemoryDeleteParamsSource = "granola"
	MemoryDeleteParamsSourceFathom           MemoryDeleteParamsSource = "fathom"
	MemoryDeleteParamsSourceFireflies        MemoryDeleteParamsSource = "fireflies"
	MemoryDeleteParamsSourceFigma            MemoryDeleteParamsSource = "figma"
	MemoryDeleteParamsSourceLinear           MemoryDeleteParamsSource = "linear"
	MemoryDeleteParamsSourceHubspot          MemoryDeleteParamsSource = "hubspot"
	MemoryDeleteParamsSourceSalesforce       MemoryDeleteParamsSource = "salesforce"
	MemoryDeleteParamsSourceCoda             MemoryDeleteParamsSource = "coda"
	MemoryDeleteParamsSourceConfluence       MemoryDeleteParamsSource = "confluence"
	MemoryDeleteParamsSourceJira             MemoryDeleteParamsSource = "jira"
	MemoryDeleteParamsSourceMetabase         MemoryDeleteParamsSource = "metabase"
	MemoryDeleteParamsSourceGong             MemoryDeleteParamsSource = "gong"
	MemoryDeleteParamsSourceClickup          MemoryDeleteParamsSource = "clickup"
	MemoryDeleteParamsSourceLightfield       MemoryDeleteParamsSource = "lightfield"
	MemoryDeleteParamsSourcePylon            MemoryDeleteParamsSource = "pylon"
	MemoryDeleteParamsSourceFellow           MemoryDeleteParamsSource = "fellow"
	MemoryDeleteParamsSourceOdoo             MemoryDeleteParamsSource = "odoo"
	MemoryDeleteParamsSourceExternalMcp      MemoryDeleteParamsSource = "external_mcp"
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
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source MemoryGetParamsSource `path:"source,omitzero" api:"required" json:"-"`
	// When true, include the document's extracted memories (chunks with summaries) in
	// the `chunks` field, in document order.
	IncludeChunks param.Opt[bool] `query:"include_chunks,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MemoryGetParams]'s query parameters as `url.Values`.
func (r MemoryGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MemoryGetParamsSource string

const (
	MemoryGetParamsSourceReddit           MemoryGetParamsSource = "reddit"
	MemoryGetParamsSourceNotion           MemoryGetParamsSource = "notion"
	MemoryGetParamsSourceSlack            MemoryGetParamsSource = "slack"
	MemoryGetParamsSourceGoogleCalendar   MemoryGetParamsSource = "google_calendar"
	MemoryGetParamsSourceGoogleMail       MemoryGetParamsSource = "google_mail"
	MemoryGetParamsSourceImap             MemoryGetParamsSource = "imap"
	MemoryGetParamsSourceGoogleMeet       MemoryGetParamsSource = "google_meet"
	MemoryGetParamsSourceBox              MemoryGetParamsSource = "box"
	MemoryGetParamsSourceDropbox          MemoryGetParamsSource = "dropbox"
	MemoryGetParamsSourceGitHub           MemoryGetParamsSource = "github"
	MemoryGetParamsSourceGitlab           MemoryGetParamsSource = "gitlab"
	MemoryGetParamsSourceGoogleDrive      MemoryGetParamsSource = "google_drive"
	MemoryGetParamsSourceVault            MemoryGetParamsSource = "vault"
	MemoryGetParamsSourceWebCrawler       MemoryGetParamsSource = "web_crawler"
	MemoryGetParamsSourceTrace            MemoryGetParamsSource = "trace"
	MemoryGetParamsSourceMicrosoftOutlook MemoryGetParamsSource = "microsoft_outlook"
	MemoryGetParamsSourceMicrosoftTeams   MemoryGetParamsSource = "microsoft_teams"
	MemoryGetParamsSourceGranola          MemoryGetParamsSource = "granola"
	MemoryGetParamsSourceFathom           MemoryGetParamsSource = "fathom"
	MemoryGetParamsSourceFireflies        MemoryGetParamsSource = "fireflies"
	MemoryGetParamsSourceFigma            MemoryGetParamsSource = "figma"
	MemoryGetParamsSourceLinear           MemoryGetParamsSource = "linear"
	MemoryGetParamsSourceHubspot          MemoryGetParamsSource = "hubspot"
	MemoryGetParamsSourceSalesforce       MemoryGetParamsSource = "salesforce"
	MemoryGetParamsSourceCoda             MemoryGetParamsSource = "coda"
	MemoryGetParamsSourceConfluence       MemoryGetParamsSource = "confluence"
	MemoryGetParamsSourceJira             MemoryGetParamsSource = "jira"
	MemoryGetParamsSourceMetabase         MemoryGetParamsSource = "metabase"
	MemoryGetParamsSourceGong             MemoryGetParamsSource = "gong"
	MemoryGetParamsSourceClickup          MemoryGetParamsSource = "clickup"
	MemoryGetParamsSourceLightfield       MemoryGetParamsSource = "lightfield"
	MemoryGetParamsSourcePylon            MemoryGetParamsSource = "pylon"
	MemoryGetParamsSourceFellow           MemoryGetParamsSource = "fellow"
	MemoryGetParamsSourceOdoo             MemoryGetParamsSource = "odoo"
	MemoryGetParamsSourceExternalMcp      MemoryGetParamsSource = "external_mcp"
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
	// trajectory, and any sources that failed. Intended for auditability and
	// compliance use cases.
	Provenance param.Opt[bool] `json:"provenance,omitzero"`
	// Controls retrieval thoroughness. 'minimal' performs direct retrieval. 'low'
	// improves the query and extracts date filters. 'medium' adds up to 3 refinement
	// rounds; 'high' allows up to 6. Higher levels can improve recall but add latency
	// and cost.
	//
	// Any of "minimal", "low", "medium", "high", "very_high".
	Effort MemorySearchParamsEffort `json:"effort,omitzero"`
	// Search options for the query.
	Options MemorySearchParamsOptions `json:"options,omitzero"`
	// Only query documents from these sources. Names are case-insensitive and accept
	// either separator, so `Google Drive`'s provider may be given as `google_drive`,
	// `google-drive`, or `GOOGLE_DRIVE`.
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
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

// Controls retrieval thoroughness. 'minimal' performs direct retrieval. 'low'
// improves the query and extracts date filters. 'medium' adds up to 3 refinement
// rounds; 'high' allows up to 6. Higher levels can improve recall but add latency
// and cost.
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
	// IANA timezone used to interpret date-only bounds and relative calendar phrases.
	// Defaults to UTC.
	Timezone param.Opt[string] `json:"timezone,omitzero"`
	// Metadata filters using MongoDB-style operators. Example: {'status': 'published',
	// 'priority': {'$gt': 3}}
	Filter map[string]any `json:"filter,omitzero"`
	// Only return results from these specific resource IDs. Useful for scoping
	// searches to specific documents (e.g., a specific email thread or uploaded file).
	ResourceIDs []string `json:"resource_ids,omitzero"`
	// Model to use for answer generation when answer=True
	//
	// Any of "llama-3.1", "gemma2", "qwen-qwq", "mistral-saba", "llama-4-scout",
	// "deepseek-r1", "gpt-oss-20b", "gpt-oss-120b", "claude-sonnet-4-6",
	// "claude-sonnet-5", "claude-opus-4-7", "claude-opus-4-8".
	AnswerModel string `json:"answer_model,omitzero"`
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
		"answer_model", "llama-3.1", "gemma2", "qwen-qwq", "mistral-saba", "llama-4-scout", "deepseek-r1", "gpt-oss-20b", "gpt-oss-120b", "claude-sonnet-4-6", "claude-sonnet-5", "claude-opus-4-7", "claude-opus-4-8",
	)
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
