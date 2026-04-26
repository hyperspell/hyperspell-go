// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

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
func (r *MemoryService) List(ctx context.Context, query MemoryListParams, opts ...option.RequestOption) (res *pagination.CursorPage[shared.Resource], err error) {
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
func (r *MemoryService) ListAutoPaging(ctx context.Context, query MemoryListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[shared.Resource] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Delete a memory and its associated chunks from the index.
//
// This removes the memory completely from the vector index and database. The
// operation deletes:
//
// 1. All chunks associated with the resource (including embeddings)
// 2. The resource record itself
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

// Retrieves a document by provider and resource_id.
func (r *MemoryService) Get(ctx context.Context, resourceID string, query MemoryGetParams, opts ...option.RequestOption) (res *Memory, err error) {
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

// Response model for the GET /memories/get endpoint.
type Memory struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
	Source MemorySource `json:"source" api:"required"`
	// The type of document (e.g. Document, Website, Email)
	Type string `json:"type" api:"required"`
	// The structured content of the document
	Data []any `json:"data" api:"nullable"`
	// Summaries of all memories extracted from this document
	Memories    []string        `json:"memories"`
	Metadata    shared.Metadata `json:"metadata"`
	Title       string          `json:"title" api:"nullable"`
	ExtraFields map[string]any  `json:"" api:"extrafields"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID  respjson.Field
		Source      respjson.Field
		Type        respjson.Field
		Data        respjson.Field
		Memories    respjson.Field
		Metadata    respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Memory) RawJSON() string { return r.JSON.raw }
func (r *Memory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MemorySource string

const (
	MemorySourceReddit         MemorySource = "reddit"
	MemorySourceNotion         MemorySource = "notion"
	MemorySourceSlack          MemorySource = "slack"
	MemorySourceGoogleCalendar MemorySource = "google_calendar"
	MemorySourceGoogleMail     MemorySource = "google_mail"
	MemorySourceBox            MemorySource = "box"
	MemorySourceDropbox        MemorySource = "dropbox"
	MemorySourceGitHub         MemorySource = "github"
	MemorySourceGoogleDrive    MemorySource = "google_drive"
	MemorySourceVault          MemorySource = "vault"
	MemorySourceWebCrawler     MemorySource = "web_crawler"
	MemorySourceTrace          MemorySource = "trace"
	MemorySourceMicrosoftTeams MemorySource = "microsoft_teams"
	MemorySourceGmailActions   MemorySource = "gmail_actions"
)

type MemoryStatus struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
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

type MemoryDeleteResponse struct {
	ChunksDeleted int64  `json:"chunks_deleted" api:"required"`
	Message       string `json:"message" api:"required"`
	ResourceID    string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
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
	// "microsoft_teams", "gmail_actions".
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
	// "microsoft_teams", "gmail_actions".
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
	// "microsoft_teams", "gmail_actions".
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
	// "microsoft_teams", "gmail_actions".
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
)

type MemorySearchParams struct {
	// Query to run.
	Query string `json:"query" api:"required"`
	// If true, the query will be answered along with matching source documents.
	Answer param.Opt[bool] `json:"answer,omitzero"`
	// Effort level. 0 = pass query through verbatim. 1 = LLM rewrites the query for
	// better retrieval and extracts date filters.
	Effort param.Opt[int64] `json:"effort,omitzero"`
	// Maximum number of results to return.
	MaxResults param.Opt[int64] `json:"max_results,omitzero"`
	// Search options for the query.
	Options MemorySearchParamsOptions `json:"options,omitzero"`
	// Only query documents from these sources.
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
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
	// Search options for Reddit
	Reddit MemorySearchParamsOptionsReddit `json:"reddit,omitzero"`
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

// Search options for Reddit
type MemorySearchParamsOptionsReddit struct {
	// The subreddit to search. If not provided, the query will be searched for in all
	// subreddits.
	Subreddit param.Opt[string] `json:"subreddit,omitzero"`
	// Weight of results from this source. A weight greater than 1.0 means more results
	// from this source will be returned, a weight less than 1.0 means fewer results
	// will be returned. This will only affect results if multiple sources are queried
	// at the same time.
	Weight param.Opt[float64] `json:"weight,omitzero"`
	// The time period to search. Defaults to 'month'.
	//
	// Any of "hour", "day", "week", "month", "year", "all".
	Period string `json:"period,omitzero"`
	// The sort order of the posts. Defaults to 'relevance'.
	//
	// Any of "relevance", "new", "hot", "top", "comments".
	Sort string `json:"sort,omitzero"`
	paramObj
}

func (r MemorySearchParamsOptionsReddit) MarshalJSON() (data []byte, err error) {
	type shadow MemorySearchParamsOptionsReddit
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MemorySearchParamsOptionsReddit) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MemorySearchParamsOptionsReddit](
		"period", "hour", "day", "week", "month", "year", "all",
	)
	apijson.RegisterFieldValidator[MemorySearchParamsOptionsReddit](
		"sort", "relevance", "new", "hot", "top", "comments",
	)
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
