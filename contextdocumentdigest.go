// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
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

// ContextDocumentDigestService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextDocumentDigestService] method instead.
type ContextDocumentDigestService struct {
	options []option.RequestOption
}

// NewContextDocumentDigestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContextDocumentDigestService(opts ...option.RequestOption) (r ContextDocumentDigestService) {
	r = ContextDocumentDigestService{}
	r.options = opts
	return
}

// List recent digest summaries, newest first.
//
// Filter by cadence with `period=daily` or `period=weekly`. Fetch full content
// with `GET /context-documents/tree/by-id/{tree_id}`.
func (r *ContextDocumentDigestService) List(ctx context.Context, query ContextDocumentDigestListParams, opts ...option.RequestOption) (res *ContextDocumentDigestListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/digest/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Generate a date-windowed "what the company did today" digest.
//
// Returns immediately with a tree ID. Poll
// `GET /context-documents/tree/{tree_id}/progress` for completion or fetch the
// result with `GET /context-documents/tree/by-id/{tree_id}`.
func (r *ContextDocumentDigestService) Generate(ctx context.Context, body ContextDocumentDigestGenerateParams, opts ...option.RequestOption) (res *ContextDocumentDigestGenerateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/digest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ContextDocumentDigestListResponse struct {
	Digests []ContextDocumentDigestListResponseDigest `json:"digests" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Digests     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentDigestListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentDigestListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A digest summary. Fetch the full content through the tree-by-ID endpoint.
type ContextDocumentDigestListResponseDigest struct {
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Period      string    `json:"period" api:"required"`
	Status      string    `json:"status" api:"required"`
	TreeID      string    `json:"tree_id" api:"required"`
	WindowEnd   time.Time `json:"window_end" api:"required" format:"date-time"`
	WindowStart time.Time `json:"window_start" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Period      respjson.Field
		Status      respjson.Field
		TreeID      respjson.Field
		WindowEnd   respjson.Field
		WindowStart respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentDigestListResponseDigest) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentDigestListResponseDigest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentDigestGenerateResponse struct {
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
func (r ContextDocumentDigestGenerateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentDigestGenerateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentDigestListParams struct {
	Period param.Opt[string] `query:"period,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContextDocumentDigestListParams]'s query parameters as
// `url.Values`.
func (r ContextDocumentDigestListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextDocumentDigestGenerateParams struct {
	// Exclusive upper bound of the digest window. Defaults to now.
	WindowEnd param.Opt[time.Time] `json:"window_end,omitzero" format:"date-time"`
	// Inclusive lower bound of the digest window. Defaults to midnight UTC today
	// (paired with window_end=now) when omitted. Both bounds must be supplied
	// together.
	WindowStart param.Opt[time.Time] `json:"window_start,omitzero" format:"date-time"`
	// Digest cadence: 'daily' or 'weekly'. Sets the default window when none is given.
	Period param.Opt[string] `json:"period,omitzero"`
	// Integration sources to include (e.g., ['slack', 'github']). Defaults to all.
	Sources []string `json:"sources,omitzero"`
	paramObj
}

func (r ContextDocumentDigestGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentDigestGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentDigestGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
