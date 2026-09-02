// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
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
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// ContextDocumentTreeService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextDocumentTreeService] method instead.
type ContextDocumentTreeService struct {
	options []option.RequestOption
}

// NewContextDocumentTreeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContextDocumentTreeService(opts ...option.RequestOption) (r ContextDocumentTreeService) {
	r = ContextDocumentTreeService{}
	r.options = opts
	return
}

// Generate a three-tier context document tree for local push delivery.
//
// Creates company, workstream, and personal context documents from the app's
// synced data. Returns immediately with a tree ID; use
// `GET /context-documents/tree/latest` to retrieve the result.
func (r *ContextDocumentTreeService) Generate(ctx context.Context, body ContextDocumentTreeGenerateParams, opts ...option.RequestOption) (res *ContextDocumentTreeGenerateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/tree"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch a specific tree by its tree ID instead of selecting the latest one.
func (r *ContextDocumentTreeService) Get(ctx context.Context, treeID string, opts ...option.RequestOption) (res *ContextDocumentTreeGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if treeID == "" {
		err = errors.New("missing required tree_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("context-documents/tree/by-id/%s", url.PathEscape(treeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get the most recent context document tree for the authenticated app.
//
// By default, the endpoint returns the latest ready tree. Readiness depends on
// whether the app has `require_review` enabled:
//
// - `require_review=False` (default): return the latest completed tree.
// - `require_review=True`: return the latest published tree.
//
// `status` filters to a specific status (case-insensitive). When no ready tree
// exists yet, the endpoint returns the newest available generation state.
func (r *ContextDocumentTreeService) GetLatest(ctx context.Context, query ContextDocumentTreeGetLatestParams, opts ...option.RequestOption) (res *ContextDocumentTreeGetLatestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/tree/latest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Return the generation progress for a single tree.
//
// Active generations include phase and counter data. Completed generations, and
// generations without detailed progress data, return status only.
func (r *ContextDocumentTreeService) Progress(ctx context.Context, treeID string, opts ...option.RequestOption) (res *ContextDocumentTreeProgressResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if treeID == "" {
		err = errors.New("missing required tree_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("context-documents/tree/%s/progress", url.PathEscape(treeID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ContextDocumentTreeGenerateResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Status    string    `json:"status" api:"required"`
	TreeID    string    `json:"tree_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGenerateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGenerateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGetResponse struct {
	CompletedAt time.Time                            `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time                            `json:"created_at" api:"required" format:"date-time"`
	Error       string                               `json:"error" api:"required"`
	Files       []ContextDocumentTreeGetResponseFile `json:"files" api:"required"`
	Meta        map[string]any                       `json:"meta" api:"required"`
	Status      string                               `json:"status" api:"required"`
	TreeID      string                               `json:"tree_id" api:"required"`
	Version     int64                                `json:"version" api:"required"`
	// Status of a newer generation that is processing or recently failed.
	//
	// This can accompany the last ready tree so clients can report progress while
	// continuing to use ready content.
	Generating ContextDocumentTreeGetResponseGenerating `json:"generating" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Error       respjson.Field
		Files       respjson.Field
		Meta        respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		Version     respjson.Field
		Generating  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGetResponseFile struct {
	Content    string         `json:"content" api:"required"`
	Path       string         `json:"path" api:"required"`
	Team       string         `json:"team" api:"required"`
	Tier       string         `json:"tier" api:"required"`
	UpdatedAt  string         `json:"updated_at" api:"required"`
	Error      string         `json:"error" api:"nullable"`
	Provenance map[string]any `json:"provenance" api:"nullable"`
	Status     string         `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Path        respjson.Field
		Team        respjson.Field
		Tier        respjson.Field
		UpdatedAt   respjson.Field
		Error       respjson.Field
		Provenance  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetResponseFile) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a newer generation that is processing or recently failed.
//
// This can accompany the last ready tree so clients can report progress while
// continuing to use ready content.
type ContextDocumentTreeGetResponseGenerating struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "processing", "failed".
	Status   string         `json:"status" api:"required"`
	TreeID   string         `json:"tree_id" api:"required"`
	Error    string         `json:"error" api:"nullable"`
	Progress map[string]any `json:"progress" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		Error       respjson.Field
		Progress    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetResponseGenerating) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetResponseGenerating) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGetLatestResponse struct {
	CompletedAt time.Time                                  `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	Error       string                                     `json:"error" api:"required"`
	Files       []ContextDocumentTreeGetLatestResponseFile `json:"files" api:"required"`
	Meta        map[string]any                             `json:"meta" api:"required"`
	Status      string                                     `json:"status" api:"required"`
	TreeID      string                                     `json:"tree_id" api:"required"`
	Version     int64                                      `json:"version" api:"required"`
	// Status of a newer generation that is processing or recently failed.
	//
	// This can accompany the last ready tree so clients can report progress while
	// continuing to use ready content.
	Generating ContextDocumentTreeGetLatestResponseGenerating `json:"generating" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Error       respjson.Field
		Files       respjson.Field
		Meta        respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		Version     respjson.Field
		Generating  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetLatestResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetLatestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGetLatestResponseFile struct {
	Content    string         `json:"content" api:"required"`
	Path       string         `json:"path" api:"required"`
	Team       string         `json:"team" api:"required"`
	Tier       string         `json:"tier" api:"required"`
	UpdatedAt  string         `json:"updated_at" api:"required"`
	Error      string         `json:"error" api:"nullable"`
	Provenance map[string]any `json:"provenance" api:"nullable"`
	Status     string         `json:"status" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Path        respjson.Field
		Team        respjson.Field
		Tier        respjson.Field
		UpdatedAt   respjson.Field
		Error       respjson.Field
		Provenance  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetLatestResponseFile) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetLatestResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a newer generation that is processing or recently failed.
//
// This can accompany the last ready tree so clients can report progress while
// continuing to use ready content.
type ContextDocumentTreeGetLatestResponseGenerating struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "processing", "failed".
	Status   string         `json:"status" api:"required"`
	TreeID   string         `json:"tree_id" api:"required"`
	Error    string         `json:"error" api:"nullable"`
	Progress map[string]any `json:"progress" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		Error       respjson.Field
		Progress    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeGetLatestResponseGenerating) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeGetLatestResponseGenerating) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response shape for GET /context-documents/tree/{tree_id}/progress.
type ContextDocumentTreeProgressResponse struct {
	Status        string   `json:"status" api:"required"`
	TreeID        string   `json:"tree_id" api:"required"`
	CompletedDocs int64    `json:"completed_docs" api:"nullable"`
	FailedDocs    int64    `json:"failed_docs" api:"nullable"`
	FailedKeys    []string `json:"failed_keys" api:"nullable"`
	// Generation phase. Values: discover, search, select, synthesize, finalize,
	// personal, done. Null when detailed progress is unavailable.
	Phase     string `json:"phase" api:"nullable"`
	TotalDocs int64  `json:"total_docs" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status        respjson.Field
		TreeID        respjson.Field
		CompletedDocs respjson.Field
		FailedDocs    respjson.Field
		FailedKeys    respjson.Field
		Phase         respjson.Field
		TotalDocs     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentTreeProgressResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentTreeProgressResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGenerateParams struct {
	// User ID for personal tier scoping. When set, personal/context.md is generated
	// from this user's data only. Company and workstream tiers still use all data.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Generate docs for this workstream only (skip auto-detection).
	WorkstreamName param.Opt[string] `json:"workstream_name,omitzero"`
	// Integration sources to include (e.g., ['gmail', 'slack']). Defaults to all.
	Sources []string `json:"sources,omitzero"`
	paramObj
}

func (r ContextDocumentTreeGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentTreeGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentTreeGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentTreeGetLatestParams struct {
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContextDocumentTreeGetLatestParams]'s query parameters as
// `url.Values`.
func (r ContextDocumentTreeGetLatestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
