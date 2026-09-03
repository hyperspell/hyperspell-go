// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/apiquery"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/pagination"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
	"github.com/hyperspell/hyperspell-go/shared"
)

// LiveService contains methods and other services that help with interacting with
// the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLiveService] method instead.
type LiveService struct {
	options []option.RequestOption
}

// NewLiveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewLiveService(opts ...option.RequestOption) (r LiveService) {
	r = LiveService{}
	r.options = opts
	return
}

// Fetch one resource live by id. A single fetch may fan out into several resources
// (e.g. a thread → its messages); all are returned.
func (r *LiveService) GetResource(ctx context.Context, resourceID string, params LiveGetResourceParams, opts ...option.RequestOption) (res *LiveGetResourceResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if resourceID == "" {
		err = errors.New("missing required resource_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("live/%v/resources/%s", params.Source, url.PathEscape(resourceID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Page through a source's resources live (no indexing side effect).
//
// The cursor is opaque and integration-defined — pass back the `next_cursor` from
// the previous page verbatim.
func (r *LiveService) ListResources(ctx context.Context, source LiveListResourcesParamsSource, query LiveListResourcesParams, opts ...option.RequestOption) (res *pagination.CursorPage[LiveListResourcesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("live/%v/resources", source)
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

// Page through a source's resources live (no indexing side effect).
//
// The cursor is opaque and integration-defined — pass back the `next_cursor` from
// the previous page verbatim.
func (r *LiveService) ListResourcesAutoPaging(ctx context.Context, source LiveListResourcesParamsSource, query LiveListResourcesParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[LiveListResourcesResponse] {
	return pagination.NewCursorPageAutoPager(r.ListResources(ctx, source, query, opts...))
}

// List the user's connected sources and the live capabilities each supports.
func (r *LiveService) ListSources(ctx context.Context, opts ...option.RequestOption) (res *LiveListSourcesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "live/sources"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search a source live for content that may not be indexed yet. With `index=true`,
// each hit is queued for indexing (no-op for live-only sources like Google
// Calendar — see `notes` in the response).
func (r *LiveService) Search(ctx context.Context, source LiveSearchParamsSource, body LiveSearchParams, opts ...option.RequestOption) (res *LiveSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := fmt.Sprintf("live/%v/search", source)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A fetch/search result: the live documents plus what happened to them.
type LiveGetResourceResponse struct {
	Documents []LiveGetResourceResponseDocument `json:"documents" api:"required"`
	Indexed   bool                              `json:"indexed"`
	Notes     []string                          `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		Indexed     respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiveGetResourceResponse) RawJSON() string { return r.JSON.raw }
func (r *LiveGetResourceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document-shaped API response containing the hyperdoc tree.
type LiveGetResourceResponseDocument struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// through `children` for the body.
	Document   LiveGetResourceResponseDocumentDocumentUnion `json:"document" api:"required"`
	ResourceID string                                       `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source string `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// Extracted memories (chunks with summaries) for this document, in document order.
	// Present only when explicitly requested via `include_chunks`; omitted otherwise.
	Chunks []LiveGetResourceResponseDocumentChunk `json:"chunks" api:"nullable"`
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
	Status string `json:"status" api:"nullable"`
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
func (r LiveGetResourceResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *LiveGetResourceResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveGetResourceResponseDocumentDocumentUnion contains all possible properties
// and values from [shared.Document], [shared.Website], [shared.Task],
// [shared.Person], [shared.Message], [shared.Event], [shared.File],
// [shared.Conversation], [shared.Trace], [shared.Transcript], [shared.Company],
// [shared.Deal], [LiveGetResourceResponseDocumentDocumentInvoice].
//
// Use the [LiveGetResourceResponseDocumentDocumentUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveGetResourceResponseDocumentDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]shared.DocumentChildUnion],
	// [[]shared.WebsiteChildUnion], [[]shared.TaskChildUnion],
	// [[]shared.PersonChildUnion], [[]shared.MessageChildUnion],
	// [[]shared.EventChildUnion], [[]shared.FileChildUnion], [[]shared.Message],
	// [[]shared.TraceChildUnion], [[]shared.Utterance], [[]shared.CompanyChildUnion],
	// [[]shared.DealChildUnion],
	// [[]LiveGetResourceResponseDocumentDocumentInvoiceChildUnion]
	Children LiveGetResourceResponseDocumentDocumentUnionChildren `json:"children"`
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
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	AttachmentNames []string `json:"attachment_names"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	BalanceAmount float64 `json:"balance_amount"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	ContactID string `json:"contact_id"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	ContactName string `json:"contact_name"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	InvoiceType string `json:"invoice_type"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	Notes string `json:"notes"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	Number string `json:"number"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	PaidAmount float64 `json:"paid_amount"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	PaidAt time.Time `json:"paid_at"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	PostedAt time.Time `json:"posted_at"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	Reference string `json:"reference"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	RefundAmount float64 `json:"refund_amount"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	RefundReason string `json:"refund_reason"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	RefundedAt time.Time `json:"refunded_at"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
	TaxAmount float64 `json:"tax_amount"`
	// This field is from variant [LiveGetResourceResponseDocumentDocumentInvoice].
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

// anyLiveGetResourceResponseDocumentDocument is implemented by each variant of
// [LiveGetResourceResponseDocumentDocumentUnion] to add type safety for the return
// type of [LiveGetResourceResponseDocumentDocumentUnion.AsAny]
type anyLiveGetResourceResponseDocumentDocument interface {
	ImplLiveGetResourceResponseDocumentDocumentUnion()
}

func (LiveGetResourceResponseDocumentDocumentInvoice) ImplLiveGetResourceResponseDocumentDocumentUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveGetResourceResponseDocumentDocumentUnion.AsAny().(type) {
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
//	case hyperspell.LiveGetResourceResponseDocumentDocumentInvoice:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LiveGetResourceResponseDocumentDocumentUnion) AsAny() anyLiveGetResourceResponseDocumentDocument {
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

func (u LiveGetResourceResponseDocumentDocumentUnion) AsDocument() (v shared.Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsWebsite() (v shared.Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsPerson() (v shared.Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsEvent() (v shared.Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsFile() (v shared.File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsConversation() (v shared.Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsTrace() (v shared.Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsTranscript() (v shared.Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsCompany() (v shared.Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsDeal() (v shared.Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentUnion) AsInvoice() (v LiveGetResourceResponseDocumentDocumentInvoice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveGetResourceResponseDocumentDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveGetResourceResponseDocumentDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveGetResourceResponseDocumentDocumentUnionChildren is an implicit subunion of
// [LiveGetResourceResponseDocumentDocumentUnion].
// LiveGetResourceResponseDocumentDocumentUnionChildren provides convenient access
// to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveGetResourceResponseDocumentDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveGetResourceResponseDocumentDocumentUnionChildren struct {
	// This field will be present if the value is a [[]shared.DocumentChildUnion]
	// instead of an object.
	OfChildren []shared.DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveGetResourceResponseDocumentDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A customer invoice, vendor bill, or credit memo.
//
// Line items are included in `children`.
type LiveGetResourceResponseDocumentDocumentInvoice struct {
	ID              string                                                     `json:"id"`
	AttachmentNames []string                                                   `json:"attachment_names" api:"nullable"`
	BalanceAmount   float64                                                    `json:"balance_amount" api:"nullable"`
	CancelledAt     time.Time                                                  `json:"cancelled_at" api:"nullable" format:"date-time"`
	Children        []LiveGetResourceResponseDocumentDocumentInvoiceChildUnion `json:"children"`
	ContactID       string                                                     `json:"contact_id" api:"nullable"`
	ContactName     string                                                     `json:"contact_name" api:"nullable"`
	Currency        string                                                     `json:"currency" api:"nullable"`
	DueAt           time.Time                                                  `json:"due_at" api:"nullable" format:"date-time"`
	InvoiceType     string                                                     `json:"invoice_type" api:"nullable"`
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
func (r LiveGetResourceResponseDocumentDocumentInvoice) RawJSON() string { return r.JSON.raw }
func (r *LiveGetResourceResponseDocumentDocumentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveGetResourceResponseDocumentDocumentInvoiceChildUnion contains all possible
// properties and values from [shared.Blob], [shared.Callout], [shared.Chunk],
// [shared.Code], [shared.Comment], [shared.Divider], [shared.Equation],
// [shared.Footnote], [shared.Heading], [shared.Image], [shared.Link],
// [shared.LineBreak], [shared.List], [shared.ListItem], [shared.Page],
// [shared.Paragraph], [shared.Quote], [shared.Table], [shared.TableCell],
// [shared.TableRow], [shared.Text], [shared.ToDo], [shared.ToolCall],
// [shared.ToolResult], [shared.TraceMessage], [shared.Utterance].
//
// Use the [LiveGetResourceResponseDocumentDocumentInvoiceChildUnion.AsAny] method
// to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveGetResourceResponseDocumentDocumentInvoiceChildUnion struct {
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
	Children LiveGetResourceResponseDocumentDocumentInvoiceChildUnionChildren `json:"children"`
	Text     string                                                           `json:"text"`
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

// anyLiveGetResourceResponseDocumentDocumentInvoiceChild is implemented by each
// variant of [LiveGetResourceResponseDocumentDocumentInvoiceChildUnion] to add
// type safety for the return type of
// [LiveGetResourceResponseDocumentDocumentInvoiceChildUnion.AsAny]
type anyLiveGetResourceResponseDocumentDocumentInvoiceChild interface {
	ImplLiveGetResourceResponseDocumentDocumentInvoiceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveGetResourceResponseDocumentDocumentInvoiceChildUnion.AsAny().(type) {
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
func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsAny() anyLiveGetResourceResponseDocumentDocumentInvoiceChild {
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

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsBlob() (v shared.Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsCallout() (v shared.Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsChunk() (v shared.Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsCode() (v shared.Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsComment() (v shared.Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsDivider() (v shared.Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsEquation() (v shared.Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsFootnote() (v shared.Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsHeading() (v shared.Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsImage() (v shared.Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsLink() (v shared.Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsLineBreak() (v shared.LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsList() (v shared.List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsListItem() (v shared.ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsPage() (v shared.Page) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsParagraph() (v shared.Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsQuote() (v shared.Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsTable() (v shared.Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsTableCell() (v shared.TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsTableRow() (v shared.TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsText() (v shared.Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsTodo() (v shared.ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsToolCall() (v shared.ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsToolResult() (v shared.ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsTraceMessage() (v shared.TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) AsUtterance() (v shared.Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveGetResourceResponseDocumentDocumentInvoiceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveGetResourceResponseDocumentDocumentInvoiceChildUnionChildren is an implicit
// subunion of [LiveGetResourceResponseDocumentDocumentInvoiceChildUnion].
// LiveGetResourceResponseDocumentDocumentInvoiceChildUnionChildren provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveGetResourceResponseDocumentDocumentInvoiceChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveGetResourceResponseDocumentDocumentInvoiceChildUnionChildren struct {
	// This field will be present if the value is a [[]shared.CalloutChildUnion]
	// instead of an object.
	OfChildren []shared.CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveGetResourceResponseDocumentDocumentInvoiceChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A searchable chunk extracted from a document during ingestion.
//
// `summary` is null when no summary was generated for the chunk.
type LiveGetResourceResponseDocumentChunk struct {
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
func (r LiveGetResourceResponseDocumentChunk) RawJSON() string { return r.JSON.raw }
func (r *LiveGetResourceResponseDocumentChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document-shaped API response containing the hyperdoc tree.
type LiveListResourcesResponse struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// through `children` for the body.
	Document   LiveListResourcesResponseDocumentUnion `json:"document" api:"required"`
	ResourceID string                                 `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source LiveListResourcesResponseSource `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// Extracted memories (chunks with summaries) for this document, in document order.
	// Present only when explicitly requested via `include_chunks`; omitted otherwise.
	Chunks []LiveListResourcesResponseChunk `json:"chunks" api:"nullable"`
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
	Status LiveListResourcesResponseStatus `json:"status" api:"nullable"`
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
func (r LiveListResourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *LiveListResourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveListResourcesResponseDocumentUnion contains all possible properties and
// values from [shared.Document], [shared.Website], [shared.Task], [shared.Person],
// [shared.Message], [shared.Event], [shared.File], [shared.Conversation],
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal],
// [LiveListResourcesResponseDocumentInvoice].
//
// Use the [LiveListResourcesResponseDocumentUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveListResourcesResponseDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]shared.DocumentChildUnion],
	// [[]shared.WebsiteChildUnion], [[]shared.TaskChildUnion],
	// [[]shared.PersonChildUnion], [[]shared.MessageChildUnion],
	// [[]shared.EventChildUnion], [[]shared.FileChildUnion], [[]shared.Message],
	// [[]shared.TraceChildUnion], [[]shared.Utterance], [[]shared.CompanyChildUnion],
	// [[]shared.DealChildUnion],
	// [[]LiveListResourcesResponseDocumentInvoiceChildUnion]
	Children LiveListResourcesResponseDocumentUnionChildren `json:"children"`
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
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	AttachmentNames []string `json:"attachment_names"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	BalanceAmount float64 `json:"balance_amount"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	ContactID string `json:"contact_id"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	ContactName string `json:"contact_name"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	InvoiceType string `json:"invoice_type"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	Notes string `json:"notes"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	Number string `json:"number"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	PaidAmount float64 `json:"paid_amount"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	PaidAt time.Time `json:"paid_at"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	PostedAt time.Time `json:"posted_at"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	Reference string `json:"reference"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	RefundAmount float64 `json:"refund_amount"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	RefundReason string `json:"refund_reason"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	RefundedAt time.Time `json:"refunded_at"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
	TaxAmount float64 `json:"tax_amount"`
	// This field is from variant [LiveListResourcesResponseDocumentInvoice].
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

// anyLiveListResourcesResponseDocument is implemented by each variant of
// [LiveListResourcesResponseDocumentUnion] to add type safety for the return type
// of [LiveListResourcesResponseDocumentUnion.AsAny]
type anyLiveListResourcesResponseDocument interface {
	ImplLiveListResourcesResponseDocumentUnion()
}

func (LiveListResourcesResponseDocumentInvoice) ImplLiveListResourcesResponseDocumentUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveListResourcesResponseDocumentUnion.AsAny().(type) {
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
//	case hyperspell.LiveListResourcesResponseDocumentInvoice:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LiveListResourcesResponseDocumentUnion) AsAny() anyLiveListResourcesResponseDocument {
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

func (u LiveListResourcesResponseDocumentUnion) AsDocument() (v shared.Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsWebsite() (v shared.Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsPerson() (v shared.Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsEvent() (v shared.Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsFile() (v shared.File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsConversation() (v shared.Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsTrace() (v shared.Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsTranscript() (v shared.Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsCompany() (v shared.Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsDeal() (v shared.Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentUnion) AsInvoice() (v LiveListResourcesResponseDocumentInvoice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveListResourcesResponseDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveListResourcesResponseDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveListResourcesResponseDocumentUnionChildren is an implicit subunion of
// [LiveListResourcesResponseDocumentUnion].
// LiveListResourcesResponseDocumentUnionChildren provides convenient access to the
// sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveListResourcesResponseDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveListResourcesResponseDocumentUnionChildren struct {
	// This field will be present if the value is a [[]shared.DocumentChildUnion]
	// instead of an object.
	OfChildren []shared.DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveListResourcesResponseDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A customer invoice, vendor bill, or credit memo.
//
// Line items are included in `children`.
type LiveListResourcesResponseDocumentInvoice struct {
	ID              string                                               `json:"id"`
	AttachmentNames []string                                             `json:"attachment_names" api:"nullable"`
	BalanceAmount   float64                                              `json:"balance_amount" api:"nullable"`
	CancelledAt     time.Time                                            `json:"cancelled_at" api:"nullable" format:"date-time"`
	Children        []LiveListResourcesResponseDocumentInvoiceChildUnion `json:"children"`
	ContactID       string                                               `json:"contact_id" api:"nullable"`
	ContactName     string                                               `json:"contact_name" api:"nullable"`
	Currency        string                                               `json:"currency" api:"nullable"`
	DueAt           time.Time                                            `json:"due_at" api:"nullable" format:"date-time"`
	InvoiceType     string                                               `json:"invoice_type" api:"nullable"`
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
func (r LiveListResourcesResponseDocumentInvoice) RawJSON() string { return r.JSON.raw }
func (r *LiveListResourcesResponseDocumentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveListResourcesResponseDocumentInvoiceChildUnion contains all possible
// properties and values from [shared.Blob], [shared.Callout], [shared.Chunk],
// [shared.Code], [shared.Comment], [shared.Divider], [shared.Equation],
// [shared.Footnote], [shared.Heading], [shared.Image], [shared.Link],
// [shared.LineBreak], [shared.List], [shared.ListItem], [shared.Page],
// [shared.Paragraph], [shared.Quote], [shared.Table], [shared.TableCell],
// [shared.TableRow], [shared.Text], [shared.ToDo], [shared.ToolCall],
// [shared.ToolResult], [shared.TraceMessage], [shared.Utterance].
//
// Use the [LiveListResourcesResponseDocumentInvoiceChildUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveListResourcesResponseDocumentInvoiceChildUnion struct {
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
	Children LiveListResourcesResponseDocumentInvoiceChildUnionChildren `json:"children"`
	Text     string                                                     `json:"text"`
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

// anyLiveListResourcesResponseDocumentInvoiceChild is implemented by each variant
// of [LiveListResourcesResponseDocumentInvoiceChildUnion] to add type safety for
// the return type of [LiveListResourcesResponseDocumentInvoiceChildUnion.AsAny]
type anyLiveListResourcesResponseDocumentInvoiceChild interface {
	ImplLiveListResourcesResponseDocumentInvoiceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveListResourcesResponseDocumentInvoiceChildUnion.AsAny().(type) {
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
func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsAny() anyLiveListResourcesResponseDocumentInvoiceChild {
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

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsBlob() (v shared.Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsCallout() (v shared.Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsChunk() (v shared.Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsCode() (v shared.Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsComment() (v shared.Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsDivider() (v shared.Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsEquation() (v shared.Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsFootnote() (v shared.Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsHeading() (v shared.Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsImage() (v shared.Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsLink() (v shared.Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsLineBreak() (v shared.LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsList() (v shared.List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsListItem() (v shared.ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsPage() (v shared.Page) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsParagraph() (v shared.Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsQuote() (v shared.Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsTable() (v shared.Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsTableCell() (v shared.TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsTableRow() (v shared.TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsText() (v shared.Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsTodo() (v shared.ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsToolCall() (v shared.ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsToolResult() (v shared.ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsTraceMessage() (v shared.TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveListResourcesResponseDocumentInvoiceChildUnion) AsUtterance() (v shared.Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveListResourcesResponseDocumentInvoiceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveListResourcesResponseDocumentInvoiceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveListResourcesResponseDocumentInvoiceChildUnionChildren is an implicit
// subunion of [LiveListResourcesResponseDocumentInvoiceChildUnion].
// LiveListResourcesResponseDocumentInvoiceChildUnionChildren provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveListResourcesResponseDocumentInvoiceChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveListResourcesResponseDocumentInvoiceChildUnionChildren struct {
	// This field will be present if the value is a [[]shared.CalloutChildUnion]
	// instead of an object.
	OfChildren []shared.CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveListResourcesResponseDocumentInvoiceChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LiveListResourcesResponseSource string

const (
	LiveListResourcesResponseSourceReddit           LiveListResourcesResponseSource = "reddit"
	LiveListResourcesResponseSourceNotion           LiveListResourcesResponseSource = "notion"
	LiveListResourcesResponseSourceSlack            LiveListResourcesResponseSource = "slack"
	LiveListResourcesResponseSourceGoogleCalendar   LiveListResourcesResponseSource = "google_calendar"
	LiveListResourcesResponseSourceGoogleMail       LiveListResourcesResponseSource = "google_mail"
	LiveListResourcesResponseSourceImap             LiveListResourcesResponseSource = "imap"
	LiveListResourcesResponseSourceGoogleMeet       LiveListResourcesResponseSource = "google_meet"
	LiveListResourcesResponseSourceBox              LiveListResourcesResponseSource = "box"
	LiveListResourcesResponseSourceDropbox          LiveListResourcesResponseSource = "dropbox"
	LiveListResourcesResponseSourceGitHub           LiveListResourcesResponseSource = "github"
	LiveListResourcesResponseSourceGitlab           LiveListResourcesResponseSource = "gitlab"
	LiveListResourcesResponseSourceGoogleDrive      LiveListResourcesResponseSource = "google_drive"
	LiveListResourcesResponseSourceVault            LiveListResourcesResponseSource = "vault"
	LiveListResourcesResponseSourceWebCrawler       LiveListResourcesResponseSource = "web_crawler"
	LiveListResourcesResponseSourceTrace            LiveListResourcesResponseSource = "trace"
	LiveListResourcesResponseSourceMicrosoftOutlook LiveListResourcesResponseSource = "microsoft_outlook"
	LiveListResourcesResponseSourceMicrosoftTeams   LiveListResourcesResponseSource = "microsoft_teams"
	LiveListResourcesResponseSourceGranola          LiveListResourcesResponseSource = "granola"
	LiveListResourcesResponseSourceFathom           LiveListResourcesResponseSource = "fathom"
	LiveListResourcesResponseSourceFireflies        LiveListResourcesResponseSource = "fireflies"
	LiveListResourcesResponseSourceFigma            LiveListResourcesResponseSource = "figma"
	LiveListResourcesResponseSourceLinear           LiveListResourcesResponseSource = "linear"
	LiveListResourcesResponseSourceHubspot          LiveListResourcesResponseSource = "hubspot"
	LiveListResourcesResponseSourceSalesforce       LiveListResourcesResponseSource = "salesforce"
	LiveListResourcesResponseSourceCoda             LiveListResourcesResponseSource = "coda"
	LiveListResourcesResponseSourceConfluence       LiveListResourcesResponseSource = "confluence"
	LiveListResourcesResponseSourceJira             LiveListResourcesResponseSource = "jira"
	LiveListResourcesResponseSourceMetabase         LiveListResourcesResponseSource = "metabase"
	LiveListResourcesResponseSourceGong             LiveListResourcesResponseSource = "gong"
	LiveListResourcesResponseSourceClickup          LiveListResourcesResponseSource = "clickup"
	LiveListResourcesResponseSourceLightfield       LiveListResourcesResponseSource = "lightfield"
	LiveListResourcesResponseSourcePylon            LiveListResourcesResponseSource = "pylon"
	LiveListResourcesResponseSourceFellow           LiveListResourcesResponseSource = "fellow"
	LiveListResourcesResponseSourceOdoo             LiveListResourcesResponseSource = "odoo"
	LiveListResourcesResponseSourceExternalMcp      LiveListResourcesResponseSource = "external_mcp"
)

// A searchable chunk extracted from a document during ingestion.
//
// `summary` is null when no summary was generated for the chunk.
type LiveListResourcesResponseChunk struct {
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
func (r LiveListResourcesResponseChunk) RawJSON() string { return r.JSON.raw }
func (r *LiveListResourcesResponseChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indexing status of the document.
type LiveListResourcesResponseStatus string

const (
	LiveListResourcesResponseStatusPending       LiveListResourcesResponseStatus = "pending"
	LiveListResourcesResponseStatusProcessing    LiveListResourcesResponseStatus = "processing"
	LiveListResourcesResponseStatusCompleted     LiveListResourcesResponseStatus = "completed"
	LiveListResourcesResponseStatusFailed        LiveListResourcesResponseStatus = "failed"
	LiveListResourcesResponseStatusPendingReview LiveListResourcesResponseStatus = "pending_review"
	LiveListResourcesResponseStatusSkipped       LiveListResourcesResponseStatus = "skipped"
	LiveListResourcesResponseStatusFiltered      LiveListResourcesResponseStatus = "filtered"
	LiveListResourcesResponseStatusCancelled     LiveListResourcesResponseStatus = "cancelled"
)

type LiveListSourcesResponse struct {
	Sources []LiveListSourcesResponseSource `json:"sources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sources     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiveListSourcesResponse) RawJSON() string { return r.JSON.raw }
func (r *LiveListSourcesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LiveListSourcesResponseSource struct {
	// Any of "nango", "unified", "whitelabel".
	AuthProvider string `json:"auth_provider" api:"required"`
	// Any of "list_resources", "fetch_resource", "search_live", "passthrough",
	// "resolve", "query_structured".
	Capabilities []string `json:"capabilities" api:"required"`
	Source       string   `json:"source" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthProvider respjson.Field
		Capabilities respjson.Field
		Source       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiveListSourcesResponseSource) RawJSON() string { return r.JSON.raw }
func (r *LiveListSourcesResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A fetch/search result: the live documents plus what happened to them.
type LiveSearchResponse struct {
	Documents []LiveSearchResponseDocument `json:"documents" api:"required"`
	Indexed   bool                         `json:"indexed"`
	Notes     []string                     `json:"notes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		Indexed     respjson.Field
		Notes       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LiveSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *LiveSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A document-shaped API response containing the hyperdoc tree.
type LiveSearchResponseDocument struct {
	// The full hyperdoc tree. Switch on `type` for the document frame and recurse
	// through `children` for the body.
	Document   LiveSearchResponseDocumentDocumentUnion `json:"document" api:"required"`
	ResourceID string                                  `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source string `json:"source" api:"required"`
	// Hyperdoc document type discriminator (document, message, file, event, ...).
	Type string `json:"type" api:"required"`
	// Extracted memories (chunks with summaries) for this document, in document order.
	// Present only when explicitly requested via `include_chunks`; omitted otherwise.
	Chunks []LiveSearchResponseDocumentChunk `json:"chunks" api:"nullable"`
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
	Status string `json:"status" api:"nullable"`
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
func (r LiveSearchResponseDocument) RawJSON() string { return r.JSON.raw }
func (r *LiveSearchResponseDocument) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveSearchResponseDocumentDocumentUnion contains all possible properties and
// values from [shared.Document], [shared.Website], [shared.Task], [shared.Person],
// [shared.Message], [shared.Event], [shared.File], [shared.Conversation],
// [shared.Trace], [shared.Transcript], [shared.Company], [shared.Deal],
// [LiveSearchResponseDocumentDocumentInvoice].
//
// Use the [LiveSearchResponseDocumentDocumentUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveSearchResponseDocumentDocumentUnion struct {
	ID string `json:"id"`
	// This field is a union of [[]shared.DocumentChildUnion],
	// [[]shared.WebsiteChildUnion], [[]shared.TaskChildUnion],
	// [[]shared.PersonChildUnion], [[]shared.MessageChildUnion],
	// [[]shared.EventChildUnion], [[]shared.FileChildUnion], [[]shared.Message],
	// [[]shared.TraceChildUnion], [[]shared.Utterance], [[]shared.CompanyChildUnion],
	// [[]shared.DealChildUnion],
	// [[]LiveSearchResponseDocumentDocumentInvoiceChildUnion]
	Children LiveSearchResponseDocumentDocumentUnionChildren `json:"children"`
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
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	AttachmentNames []string `json:"attachment_names"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	BalanceAmount float64 `json:"balance_amount"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	CancelledAt time.Time `json:"cancelled_at"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	ContactID string `json:"contact_id"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	ContactName string `json:"contact_name"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	InvoiceType string `json:"invoice_type"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	Notes string `json:"notes"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	Number string `json:"number"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	OrganizationID string `json:"organization_id"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	PaidAmount float64 `json:"paid_amount"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	PaidAt time.Time `json:"paid_at"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	PostedAt time.Time `json:"posted_at"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	Reference string `json:"reference"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	RefundAmount float64 `json:"refund_amount"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	RefundReason string `json:"refund_reason"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	RefundedAt time.Time `json:"refunded_at"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
	TaxAmount float64 `json:"tax_amount"`
	// This field is from variant [LiveSearchResponseDocumentDocumentInvoice].
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

// anyLiveSearchResponseDocumentDocument is implemented by each variant of
// [LiveSearchResponseDocumentDocumentUnion] to add type safety for the return type
// of [LiveSearchResponseDocumentDocumentUnion.AsAny]
type anyLiveSearchResponseDocumentDocument interface {
	ImplLiveSearchResponseDocumentDocumentUnion()
}

func (LiveSearchResponseDocumentDocumentInvoice) ImplLiveSearchResponseDocumentDocumentUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveSearchResponseDocumentDocumentUnion.AsAny().(type) {
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
//	case hyperspell.LiveSearchResponseDocumentDocumentInvoice:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u LiveSearchResponseDocumentDocumentUnion) AsAny() anyLiveSearchResponseDocumentDocument {
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

func (u LiveSearchResponseDocumentDocumentUnion) AsDocument() (v shared.Document) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsWebsite() (v shared.Website) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsTask() (v shared.Task) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsPerson() (v shared.Person) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsMessage() (v shared.Message) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsEvent() (v shared.Event) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsFile() (v shared.File) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsConversation() (v shared.Conversation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsTrace() (v shared.Trace) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsTranscript() (v shared.Transcript) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsCompany() (v shared.Company) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsDeal() (v shared.Deal) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentUnion) AsInvoice() (v LiveSearchResponseDocumentDocumentInvoice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveSearchResponseDocumentDocumentUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveSearchResponseDocumentDocumentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveSearchResponseDocumentDocumentUnionChildren is an implicit subunion of
// [LiveSearchResponseDocumentDocumentUnion].
// LiveSearchResponseDocumentDocumentUnionChildren provides convenient access to
// the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveSearchResponseDocumentDocumentUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveSearchResponseDocumentDocumentUnionChildren struct {
	// This field will be present if the value is a [[]shared.DocumentChildUnion]
	// instead of an object.
	OfChildren []shared.DocumentChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveSearchResponseDocumentDocumentUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A customer invoice, vendor bill, or credit memo.
//
// Line items are included in `children`.
type LiveSearchResponseDocumentDocumentInvoice struct {
	ID              string                                                `json:"id"`
	AttachmentNames []string                                              `json:"attachment_names" api:"nullable"`
	BalanceAmount   float64                                               `json:"balance_amount" api:"nullable"`
	CancelledAt     time.Time                                             `json:"cancelled_at" api:"nullable" format:"date-time"`
	Children        []LiveSearchResponseDocumentDocumentInvoiceChildUnion `json:"children"`
	ContactID       string                                                `json:"contact_id" api:"nullable"`
	ContactName     string                                                `json:"contact_name" api:"nullable"`
	Currency        string                                                `json:"currency" api:"nullable"`
	DueAt           time.Time                                             `json:"due_at" api:"nullable" format:"date-time"`
	InvoiceType     string                                                `json:"invoice_type" api:"nullable"`
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
func (r LiveSearchResponseDocumentDocumentInvoice) RawJSON() string { return r.JSON.raw }
func (r *LiveSearchResponseDocumentDocumentInvoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveSearchResponseDocumentDocumentInvoiceChildUnion contains all possible
// properties and values from [shared.Blob], [shared.Callout], [shared.Chunk],
// [shared.Code], [shared.Comment], [shared.Divider], [shared.Equation],
// [shared.Footnote], [shared.Heading], [shared.Image], [shared.Link],
// [shared.LineBreak], [shared.List], [shared.ListItem], [shared.Page],
// [shared.Paragraph], [shared.Quote], [shared.Table], [shared.TableCell],
// [shared.TableRow], [shared.Text], [shared.ToDo], [shared.ToolCall],
// [shared.ToolResult], [shared.TraceMessage], [shared.Utterance].
//
// Use the [LiveSearchResponseDocumentDocumentInvoiceChildUnion.AsAny] method to
// switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type LiveSearchResponseDocumentDocumentInvoiceChildUnion struct {
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
	Children LiveSearchResponseDocumentDocumentInvoiceChildUnionChildren `json:"children"`
	Text     string                                                      `json:"text"`
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

// anyLiveSearchResponseDocumentDocumentInvoiceChild is implemented by each variant
// of [LiveSearchResponseDocumentDocumentInvoiceChildUnion] to add type safety for
// the return type of [LiveSearchResponseDocumentDocumentInvoiceChildUnion.AsAny]
type anyLiveSearchResponseDocumentDocumentInvoiceChild interface {
	ImplLiveSearchResponseDocumentDocumentInvoiceChildUnion()
}

// Use the following switch statement to find the correct variant
//
//	switch variant := LiveSearchResponseDocumentDocumentInvoiceChildUnion.AsAny().(type) {
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
func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsAny() anyLiveSearchResponseDocumentDocumentInvoiceChild {
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

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsBlob() (v shared.Blob) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsCallout() (v shared.Callout) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsChunk() (v shared.Chunk) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsCode() (v shared.Code) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsComment() (v shared.Comment) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsDivider() (v shared.Divider) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsEquation() (v shared.Equation) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsFootnote() (v shared.Footnote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsHeading() (v shared.Heading) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsImage() (v shared.Image) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsLink() (v shared.Link) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsLineBreak() (v shared.LineBreak) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsList() (v shared.List) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsListItem() (v shared.ListItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsPage() (v shared.Page) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsParagraph() (v shared.Paragraph) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsQuote() (v shared.Quote) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsTable() (v shared.Table) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsTableCell() (v shared.TableCell) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsTableRow() (v shared.TableRow) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsText() (v shared.Text) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsTodo() (v shared.ToDo) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsToolCall() (v shared.ToolCall) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsToolResult() (v shared.ToolResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsTraceMessage() (v shared.TraceMessage) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) AsUtterance() (v shared.Utterance) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u LiveSearchResponseDocumentDocumentInvoiceChildUnion) RawJSON() string { return u.JSON.raw }

func (r *LiveSearchResponseDocumentDocumentInvoiceChildUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiveSearchResponseDocumentDocumentInvoiceChildUnionChildren is an implicit
// subunion of [LiveSearchResponseDocumentDocumentInvoiceChildUnion].
// LiveSearchResponseDocumentDocumentInvoiceChildUnionChildren provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [LiveSearchResponseDocumentDocumentInvoiceChildUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfChildren]
type LiveSearchResponseDocumentDocumentInvoiceChildUnionChildren struct {
	// This field will be present if the value is a [[]shared.CalloutChildUnion]
	// instead of an object.
	OfChildren []shared.CalloutChildUnion `json:",inline"`
	JSON       struct {
		OfChildren respjson.Field
		raw        string
	} `json:"-"`
}

func (r *LiveSearchResponseDocumentDocumentInvoiceChildUnionChildren) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A searchable chunk extracted from a document during ingestion.
//
// `summary` is null when no summary was generated for the chunk.
type LiveSearchResponseDocumentChunk struct {
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
func (r LiveSearchResponseDocumentChunk) RawJSON() string { return r.JSON.raw }
func (r *LiveSearchResponseDocumentChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LiveGetResourceParams struct {
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "imap",
	// "google_meet", "box", "dropbox", "github", "gitlab", "google_drive", "vault",
	// "web_crawler", "trace", "microsoft_outlook", "microsoft_teams", "granola",
	// "fathom", "fireflies", "figma", "linear", "hubspot", "salesforce", "coda",
	// "confluence", "jira", "metabase", "gong", "clickup", "lightfield", "pylon",
	// "fellow", "odoo", "external_mcp".
	Source LiveGetResourceParamsSource `path:"source,omitzero" api:"required" json:"-"`
	// Specific connection id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	// Also queue this resource for indexing.
	Index param.Opt[bool] `query:"index,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LiveGetResourceParams]'s query parameters as `url.Values`.
func (r LiveGetResourceParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LiveGetResourceParamsSource string

const (
	LiveGetResourceParamsSourceReddit           LiveGetResourceParamsSource = "reddit"
	LiveGetResourceParamsSourceNotion           LiveGetResourceParamsSource = "notion"
	LiveGetResourceParamsSourceSlack            LiveGetResourceParamsSource = "slack"
	LiveGetResourceParamsSourceGoogleCalendar   LiveGetResourceParamsSource = "google_calendar"
	LiveGetResourceParamsSourceGoogleMail       LiveGetResourceParamsSource = "google_mail"
	LiveGetResourceParamsSourceImap             LiveGetResourceParamsSource = "imap"
	LiveGetResourceParamsSourceGoogleMeet       LiveGetResourceParamsSource = "google_meet"
	LiveGetResourceParamsSourceBox              LiveGetResourceParamsSource = "box"
	LiveGetResourceParamsSourceDropbox          LiveGetResourceParamsSource = "dropbox"
	LiveGetResourceParamsSourceGitHub           LiveGetResourceParamsSource = "github"
	LiveGetResourceParamsSourceGitlab           LiveGetResourceParamsSource = "gitlab"
	LiveGetResourceParamsSourceGoogleDrive      LiveGetResourceParamsSource = "google_drive"
	LiveGetResourceParamsSourceVault            LiveGetResourceParamsSource = "vault"
	LiveGetResourceParamsSourceWebCrawler       LiveGetResourceParamsSource = "web_crawler"
	LiveGetResourceParamsSourceTrace            LiveGetResourceParamsSource = "trace"
	LiveGetResourceParamsSourceMicrosoftOutlook LiveGetResourceParamsSource = "microsoft_outlook"
	LiveGetResourceParamsSourceMicrosoftTeams   LiveGetResourceParamsSource = "microsoft_teams"
	LiveGetResourceParamsSourceGranola          LiveGetResourceParamsSource = "granola"
	LiveGetResourceParamsSourceFathom           LiveGetResourceParamsSource = "fathom"
	LiveGetResourceParamsSourceFireflies        LiveGetResourceParamsSource = "fireflies"
	LiveGetResourceParamsSourceFigma            LiveGetResourceParamsSource = "figma"
	LiveGetResourceParamsSourceLinear           LiveGetResourceParamsSource = "linear"
	LiveGetResourceParamsSourceHubspot          LiveGetResourceParamsSource = "hubspot"
	LiveGetResourceParamsSourceSalesforce       LiveGetResourceParamsSource = "salesforce"
	LiveGetResourceParamsSourceCoda             LiveGetResourceParamsSource = "coda"
	LiveGetResourceParamsSourceConfluence       LiveGetResourceParamsSource = "confluence"
	LiveGetResourceParamsSourceJira             LiveGetResourceParamsSource = "jira"
	LiveGetResourceParamsSourceMetabase         LiveGetResourceParamsSource = "metabase"
	LiveGetResourceParamsSourceGong             LiveGetResourceParamsSource = "gong"
	LiveGetResourceParamsSourceClickup          LiveGetResourceParamsSource = "clickup"
	LiveGetResourceParamsSourceLightfield       LiveGetResourceParamsSource = "lightfield"
	LiveGetResourceParamsSourcePylon            LiveGetResourceParamsSource = "pylon"
	LiveGetResourceParamsSourceFellow           LiveGetResourceParamsSource = "fellow"
	LiveGetResourceParamsSourceOdoo             LiveGetResourceParamsSource = "odoo"
	LiveGetResourceParamsSourceExternalMcp      LiveGetResourceParamsSource = "external_mcp"
)

type LiveListResourcesParams struct {
	// Specific connection id.
	ConnectionID param.Opt[string] `query:"connection_id,omitzero" json:"-"`
	Cursor       param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Size         param.Opt[int64]  `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [LiveListResourcesParams]'s query parameters as
// `url.Values`.
func (r LiveListResourcesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LiveListResourcesParamsSource string

const (
	LiveListResourcesParamsSourceReddit           LiveListResourcesParamsSource = "reddit"
	LiveListResourcesParamsSourceNotion           LiveListResourcesParamsSource = "notion"
	LiveListResourcesParamsSourceSlack            LiveListResourcesParamsSource = "slack"
	LiveListResourcesParamsSourceGoogleCalendar   LiveListResourcesParamsSource = "google_calendar"
	LiveListResourcesParamsSourceGoogleMail       LiveListResourcesParamsSource = "google_mail"
	LiveListResourcesParamsSourceImap             LiveListResourcesParamsSource = "imap"
	LiveListResourcesParamsSourceGoogleMeet       LiveListResourcesParamsSource = "google_meet"
	LiveListResourcesParamsSourceBox              LiveListResourcesParamsSource = "box"
	LiveListResourcesParamsSourceDropbox          LiveListResourcesParamsSource = "dropbox"
	LiveListResourcesParamsSourceGitHub           LiveListResourcesParamsSource = "github"
	LiveListResourcesParamsSourceGitlab           LiveListResourcesParamsSource = "gitlab"
	LiveListResourcesParamsSourceGoogleDrive      LiveListResourcesParamsSource = "google_drive"
	LiveListResourcesParamsSourceVault            LiveListResourcesParamsSource = "vault"
	LiveListResourcesParamsSourceWebCrawler       LiveListResourcesParamsSource = "web_crawler"
	LiveListResourcesParamsSourceTrace            LiveListResourcesParamsSource = "trace"
	LiveListResourcesParamsSourceMicrosoftOutlook LiveListResourcesParamsSource = "microsoft_outlook"
	LiveListResourcesParamsSourceMicrosoftTeams   LiveListResourcesParamsSource = "microsoft_teams"
	LiveListResourcesParamsSourceGranola          LiveListResourcesParamsSource = "granola"
	LiveListResourcesParamsSourceFathom           LiveListResourcesParamsSource = "fathom"
	LiveListResourcesParamsSourceFireflies        LiveListResourcesParamsSource = "fireflies"
	LiveListResourcesParamsSourceFigma            LiveListResourcesParamsSource = "figma"
	LiveListResourcesParamsSourceLinear           LiveListResourcesParamsSource = "linear"
	LiveListResourcesParamsSourceHubspot          LiveListResourcesParamsSource = "hubspot"
	LiveListResourcesParamsSourceSalesforce       LiveListResourcesParamsSource = "salesforce"
	LiveListResourcesParamsSourceCoda             LiveListResourcesParamsSource = "coda"
	LiveListResourcesParamsSourceConfluence       LiveListResourcesParamsSource = "confluence"
	LiveListResourcesParamsSourceJira             LiveListResourcesParamsSource = "jira"
	LiveListResourcesParamsSourceMetabase         LiveListResourcesParamsSource = "metabase"
	LiveListResourcesParamsSourceGong             LiveListResourcesParamsSource = "gong"
	LiveListResourcesParamsSourceClickup          LiveListResourcesParamsSource = "clickup"
	LiveListResourcesParamsSourceLightfield       LiveListResourcesParamsSource = "lightfield"
	LiveListResourcesParamsSourcePylon            LiveListResourcesParamsSource = "pylon"
	LiveListResourcesParamsSourceFellow           LiveListResourcesParamsSource = "fellow"
	LiveListResourcesParamsSourceOdoo             LiveListResourcesParamsSource = "odoo"
	LiveListResourcesParamsSourceExternalMcp      LiveListResourcesParamsSource = "external_mcp"
)

type LiveSearchParams struct {
	// Live search query.
	Query string `json:"query" api:"required"`
	// Specific connection id when the user has multiple for this source.
	ConnectionID param.Opt[string] `json:"connection_id,omitzero"`
	// If true, queue each hit for indexing so it's on-hand next time.
	Index param.Opt[bool] `json:"index,omitzero"`
	paramObj
}

func (r LiveSearchParams) MarshalJSON() (data []byte, err error) {
	type shadow LiveSearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LiveSearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type LiveSearchParamsSource string

const (
	LiveSearchParamsSourceReddit           LiveSearchParamsSource = "reddit"
	LiveSearchParamsSourceNotion           LiveSearchParamsSource = "notion"
	LiveSearchParamsSourceSlack            LiveSearchParamsSource = "slack"
	LiveSearchParamsSourceGoogleCalendar   LiveSearchParamsSource = "google_calendar"
	LiveSearchParamsSourceGoogleMail       LiveSearchParamsSource = "google_mail"
	LiveSearchParamsSourceImap             LiveSearchParamsSource = "imap"
	LiveSearchParamsSourceGoogleMeet       LiveSearchParamsSource = "google_meet"
	LiveSearchParamsSourceBox              LiveSearchParamsSource = "box"
	LiveSearchParamsSourceDropbox          LiveSearchParamsSource = "dropbox"
	LiveSearchParamsSourceGitHub           LiveSearchParamsSource = "github"
	LiveSearchParamsSourceGitlab           LiveSearchParamsSource = "gitlab"
	LiveSearchParamsSourceGoogleDrive      LiveSearchParamsSource = "google_drive"
	LiveSearchParamsSourceVault            LiveSearchParamsSource = "vault"
	LiveSearchParamsSourceWebCrawler       LiveSearchParamsSource = "web_crawler"
	LiveSearchParamsSourceTrace            LiveSearchParamsSource = "trace"
	LiveSearchParamsSourceMicrosoftOutlook LiveSearchParamsSource = "microsoft_outlook"
	LiveSearchParamsSourceMicrosoftTeams   LiveSearchParamsSource = "microsoft_teams"
	LiveSearchParamsSourceGranola          LiveSearchParamsSource = "granola"
	LiveSearchParamsSourceFathom           LiveSearchParamsSource = "fathom"
	LiveSearchParamsSourceFireflies        LiveSearchParamsSource = "fireflies"
	LiveSearchParamsSourceFigma            LiveSearchParamsSource = "figma"
	LiveSearchParamsSourceLinear           LiveSearchParamsSource = "linear"
	LiveSearchParamsSourceHubspot          LiveSearchParamsSource = "hubspot"
	LiveSearchParamsSourceSalesforce       LiveSearchParamsSource = "salesforce"
	LiveSearchParamsSourceCoda             LiveSearchParamsSource = "coda"
	LiveSearchParamsSourceConfluence       LiveSearchParamsSource = "confluence"
	LiveSearchParamsSourceJira             LiveSearchParamsSource = "jira"
	LiveSearchParamsSourceMetabase         LiveSearchParamsSource = "metabase"
	LiveSearchParamsSourceGong             LiveSearchParamsSource = "gong"
	LiveSearchParamsSourceClickup          LiveSearchParamsSource = "clickup"
	LiveSearchParamsSourceLightfield       LiveSearchParamsSource = "lightfield"
	LiveSearchParamsSourcePylon            LiveSearchParamsSource = "pylon"
	LiveSearchParamsSourceFellow           LiveSearchParamsSource = "fellow"
	LiveSearchParamsSourceOdoo             LiveSearchParamsSource = "odoo"
	LiveSearchParamsSourceExternalMcp      LiveSearchParamsSource = "external_mcp"
)
