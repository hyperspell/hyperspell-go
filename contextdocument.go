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
	"github.com/hyperspell/hyperspell-go/packages/pagination"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// ContextDocumentService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextDocumentService] method instead.
type ContextDocumentService struct {
	options []option.RequestOption
	Trees   ContextDocumentTreeService
	Digests ContextDocumentDigestService
	Config  ContextDocumentConfigService
}

// NewContextDocumentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewContextDocumentService(opts ...option.RequestOption) (r ContextDocumentService) {
	r = ContextDocumentService{}
	r.options = opts
	r.Trees = NewContextDocumentTreeService(opts...)
	r.Digests = NewContextDocumentDigestService(opts...)
	r.Config = NewContextDocumentConfigService(opts...)
	return
}

// List context documents for the authenticated app, most recent first.
func (r *ContextDocumentService) List(ctx context.Context, query ContextDocumentListParams, opts ...option.RequestOption) (res *pagination.ContextDocumentsCursorPage[ContextDocumentListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "context-documents"
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

// List context documents for the authenticated app, most recent first.
func (r *ContextDocumentService) ListAutoPaging(ctx context.Context, query ContextDocumentListParams, opts ...option.RequestOption) *pagination.ContextDocumentsCursorPageAutoPager[ContextDocumentListResponse] {
	return pagination.NewContextDocumentsCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Generate an LLM-synthesized context document from the app's synced data.
//
// Generation runs asynchronously. The endpoint returns immediately with status
// `PROCESSING`; synthesis time depends on the amount of source data.
func (r *ContextDocumentService) Generate(ctx context.Context, body ContextDocumentGenerateParams, opts ...option.RequestOption) (res *ContextDocumentGenerateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/generate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a specific context document by ID.
func (r *ContextDocumentService) Get(ctx context.Context, documentID string, opts ...option.RequestOption) (res *ContextDocumentGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if documentID == "" {
		err = errors.New("missing required document_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("context-documents/%s", url.PathEscape(documentID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ContextDocumentListResponse struct {
	CompletedAt time.Time `json:"completed_at" api:"required" format:"date-time"`
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	DocumentID  string    `json:"document_id" api:"required"`
	Model       string    `json:"model" api:"required"`
	Sources     []string  `json:"sources" api:"required"`
	Status      string    `json:"status" api:"required"`
	TokenCount  int64     `json:"token_count" api:"required"`
	Error       string    `json:"error" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		Model       respjson.Field
		Sources     respjson.Field
		Status      respjson.Field
		TokenCount  respjson.Field
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentListResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentGenerateResponse struct {
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	DocumentID string    `json:"document_id" api:"required"`
	Status     string    `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentGenerateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentGenerateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentGetResponse struct {
	CompletedAt time.Time      `json:"completed_at" api:"required" format:"date-time"`
	Content     string         `json:"content" api:"required"`
	CreatedAt   time.Time      `json:"created_at" api:"required" format:"date-time"`
	DocumentID  string         `json:"document_id" api:"required"`
	Error       string         `json:"error" api:"required"`
	Metadata    map[string]any `json:"metadata" api:"required"`
	Model       string         `json:"model" api:"required"`
	Prompt      string         `json:"prompt" api:"required"`
	Sources     []string       `json:"sources" api:"required"`
	Status      string         `json:"status" api:"required"`
	TokenUsage  map[string]any `json:"token_usage" api:"required"`
	UserID      string         `json:"user_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		DocumentID  respjson.Field
		Error       respjson.Field
		Metadata    respjson.Field
		Model       respjson.Field
		Prompt      respjson.Field
		Sources     respjson.Field
		Status      respjson.Field
		TokenUsage  respjson.Field
		UserID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Any of "processing", "completed", "failed".
	Status ContextDocumentListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ContextDocumentListParams]'s query parameters as
// `url.Values`.
func (r ContextDocumentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ContextDocumentListParamsStatus string

const (
	ContextDocumentListParamsStatusProcessing ContextDocumentListParamsStatus = "processing"
	ContextDocumentListParamsStatusCompleted  ContextDocumentListParamsStatus = "completed"
	ContextDocumentListParamsStatusFailed     ContextDocumentListParamsStatus = "failed"
)

type ContextDocumentGenerateParams struct {
	// Custom prompt template. Replaces the standard summary prompt.
	Prompt param.Opt[string] `json:"prompt,omitzero"`
	// Scope generation to a specific user's data.
	UserID param.Opt[string] `json:"user_id,omitzero"`
	// Model used for final synthesis.
	Model param.Opt[string] `json:"model,omitzero"`
	// Integration sources to include (e.g., ['gmail', 'slack']). Defaults to all
	// connected integrations.
	Sources []string `json:"sources,omitzero"`
	paramObj
}

func (r ContextDocumentGenerateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentGenerateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentGenerateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
