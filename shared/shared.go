// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"time"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Metadata struct {
	CreatedAt    time.Time      `json:"created_at" api:"nullable" format:"date-time"`
	Events       []Notification `json:"events"`
	IndexedAt    time.Time      `json:"indexed_at" api:"nullable" format:"date-time"`
	LastModified time.Time      `json:"last_modified" api:"nullable" format:"date-time"`
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status      MetadataStatus `json:"status"`
	URL         string         `json:"url" api:"nullable"`
	ExtraFields map[string]any `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt    respjson.Field
		Events       respjson.Field
		IndexedAt    respjson.Field
		LastModified respjson.Field
		Status       respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MetadataStatus string

const (
	MetadataStatusPending       MetadataStatus = "pending"
	MetadataStatusProcessing    MetadataStatus = "processing"
	MetadataStatusCompleted     MetadataStatus = "completed"
	MetadataStatusFailed        MetadataStatus = "failed"
	MetadataStatusPendingReview MetadataStatus = "pending_review"
	MetadataStatusSkipped       MetadataStatus = "skipped"
)

type Notification struct {
	Message string `json:"message" api:"required"`
	// Any of "error", "warning", "info", "success".
	Type NotificationType `json:"type" api:"required"`
	Time time.Time        `json:"time" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Type        respjson.Field
		Time        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Notification) RawJSON() string { return r.JSON.raw }
func (r *Notification) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NotificationType string

const (
	NotificationTypeError   NotificationType = "error"
	NotificationTypeWarning NotificationType = "warning"
	NotificationTypeInfo    NotificationType = "info"
	NotificationTypeSuccess NotificationType = "success"
)

type QueryResult struct {
	Documents []Resource `json:"documents" api:"required"`
	// The answer to the query, if the request was set to answer.
	Answer string `json:"answer" api:"nullable"`
	// Errors that occurred during the query. These are meant to help the developer
	// debug the query, and are not meant to be shown to the user.
	Errors []map[string]string `json:"errors" api:"nullable"`
	// Auditability record attached to an agentic answer.
	//
	// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
	// failed_sources) are derived from in-memory loop state, but `entities` costs one
	// indexed DB lookup, so the whole record is only built on request.
	Provenance QueryResultProvenance `json:"provenance" api:"nullable"`
	// The ID of the query. This can be used to retrieve the query later, or add
	// feedback to it. If the query failed, this will be None.
	QueryID string `json:"query_id" api:"nullable"`
	// The average score of the query feedback, if any.
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Documents   respjson.Field
		Answer      respjson.Field
		Errors      respjson.Field
		Provenance  respjson.Field
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

// Auditability record attached to an agentic answer.
//
// Gated behind `provenance=true` on the request: the cheap parts (sources, steps,
// failed_sources) are derived from in-memory loop state, but `entities` costs one
// indexed DB lookup, so the whole record is only built on request.
type QueryResultProvenance struct {
	Entities      []QueryResultProvenanceEntity `json:"entities"`
	FailedSources []string                      `json:"failed_sources"`
	Sources       []QueryResultProvenanceSource `json:"sources"`
	Steps         []QueryResultProvenanceStep   `json:"steps"`
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
func (r QueryResultProvenance) RawJSON() string { return r.JSON.raw }
func (r *QueryResultProvenance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A canonical entity referenced by the answer's source documents.
type QueryResultProvenanceEntity struct {
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
func (r QueryResultProvenanceEntity) RawJSON() string { return r.JSON.raw }
func (r *QueryResultProvenanceEntity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source document that informed the final answer (the post-rank result set).
type QueryResultProvenanceSource struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield".
	Source string  `json:"source" api:"required"`
	Score  float64 `json:"score" api:"nullable"`
	Title  string  `json:"title" api:"nullable"`
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
func (r QueryResultProvenanceSource) RawJSON() string { return r.JSON.raw }
func (r *QueryResultProvenanceSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One tool invocation in the agent's search trajectory (audit trail).
type QueryResultProvenanceStep struct {
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
func (r QueryResultProvenanceStep) RawJSON() string { return r.JSON.raw }
func (r *QueryResultProvenanceStep) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Resource struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield".
	Source ResourceSource `json:"source" api:"required"`
	// Ordered list of provider folder IDs from immediate parent up to (but not
	// including) provider root. Used by resolve_sync_mode to walk the actual folder
	// tree without depending on intermediate policy records. Empty = resource lives at
	// provider root.
	FolderAncestors []string `json:"folder_ancestors"`
	// Provider folder ID this resource belongs to
	FolderID string   `json:"folder_id" api:"nullable"`
	Metadata Metadata `json:"metadata"`
	// Parent folder ID for policy inheritance
	ParentFolderID string `json:"parent_folder_id" api:"nullable"`
	// The relevance of the resource to the query
	Score float64 `json:"score" api:"nullable"`
	Title string  `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID      respjson.Field
		Source          respjson.Field
		FolderAncestors respjson.Field
		FolderID        respjson.Field
		Metadata        respjson.Field
		ParentFolderID  respjson.Field
		Score           respjson.Field
		Title           respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Resource) RawJSON() string { return r.JSON.raw }
func (r *Resource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ResourceSource string

const (
	ResourceSourceReddit         ResourceSource = "reddit"
	ResourceSourceNotion         ResourceSource = "notion"
	ResourceSourceSlack          ResourceSource = "slack"
	ResourceSourceGoogleCalendar ResourceSource = "google_calendar"
	ResourceSourceGoogleMail     ResourceSource = "google_mail"
	ResourceSourceBox            ResourceSource = "box"
	ResourceSourceDropbox        ResourceSource = "dropbox"
	ResourceSourceGitHub         ResourceSource = "github"
	ResourceSourceGoogleDrive    ResourceSource = "google_drive"
	ResourceSourceVault          ResourceSource = "vault"
	ResourceSourceWebCrawler     ResourceSource = "web_crawler"
	ResourceSourceTrace          ResourceSource = "trace"
	ResourceSourceMicrosoftTeams ResourceSource = "microsoft_teams"
	ResourceSourceGmailActions   ResourceSource = "gmail_actions"
	ResourceSourceGranola        ResourceSource = "granola"
	ResourceSourceFathom         ResourceSource = "fathom"
	ResourceSourceFireflies      ResourceSource = "fireflies"
	ResourceSourceLinear         ResourceSource = "linear"
	ResourceSourceHubspot        ResourceSource = "hubspot"
	ResourceSourceSalesforce     ResourceSource = "salesforce"
	ResourceSourceCoda           ResourceSource = "coda"
	ResourceSourceLightfield     ResourceSource = "lightfield"
)
