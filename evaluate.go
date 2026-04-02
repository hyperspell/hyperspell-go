// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
	"github.com/hyperspell/hyperspell-go/shared"
)

// EvaluateService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEvaluateService] method instead.
type EvaluateService struct {
	options []option.RequestOption
}

// NewEvaluateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEvaluateService(opts ...option.RequestOption) (r EvaluateService) {
	r = EvaluateService{}
	r.options = opts
	return
}

// Retrieve the result of a previous query.
func (r *EvaluateService) GetQuery(ctx context.Context, queryID string, opts ...option.RequestOption) (res *shared.QueryResult, err error) {
	opts = slices.Concat(r.options, opts)
	if queryID == "" {
		err = errors.New("missing required query_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("evaluate/query/%s", url.PathEscape(queryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Score an individual highlight.
func (r *EvaluateService) ScoreHighlight(ctx context.Context, highlightID string, body EvaluateScoreHighlightParams, opts ...option.RequestOption) (res *EvaluateScoreHighlightResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if highlightID == "" {
		err = errors.New("missing required highlight_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("evaluate/highlight/%s", url.PathEscape(highlightID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Score the result of a query.
func (r *EvaluateService) ScoreQuery(ctx context.Context, queryID string, body EvaluateScoreQueryParams, opts ...option.RequestOption) (res *EvaluateScoreQueryResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if queryID == "" {
		err = errors.New("missing required query_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("evaluate/query/%s", url.PathEscape(queryID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EvaluateScoreHighlightResponse struct {
	// A message describing the result.
	Message string `json:"message" api:"required"`
	// Whether the feedback was successfully saved.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluateScoreHighlightResponse) RawJSON() string { return r.JSON.raw }
func (r *EvaluateScoreHighlightResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluateScoreQueryResponse struct {
	// A message describing the result.
	Message string `json:"message" api:"required"`
	// Whether the feedback was successfully saved.
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EvaluateScoreQueryResponse) RawJSON() string { return r.JSON.raw }
func (r *EvaluateScoreQueryResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluateScoreHighlightParams struct {
	// Comment on the chunk
	Comment param.Opt[string] `json:"comment,omitzero"`
	// Rating of the chunk from -1 (bad) to +1 (good).
	Score param.Opt[float64] `json:"score,omitzero"`
	paramObj
}

func (r EvaluateScoreHighlightParams) MarshalJSON() (data []byte, err error) {
	type shadow EvaluateScoreHighlightParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvaluateScoreHighlightParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EvaluateScoreQueryParams struct {
	// Rating of the query result from -1 (bad) to +1 (good).
	Score param.Opt[float64] `json:"score,omitzero"`
	paramObj
}

func (r EvaluateScoreQueryParams) MarshalJSON() (data []byte, err error) {
	type shadow EvaluateScoreQueryParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EvaluateScoreQueryParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
