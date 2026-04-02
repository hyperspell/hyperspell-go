// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hyperspell-go/internal/apijson"
	"github.com/stainless-sdks/hyperspell-go/internal/apiquery"
	"github.com/stainless-sdks/hyperspell-go/internal/requestconfig"
	"github.com/stainless-sdks/hyperspell-go/option"
	"github.com/stainless-sdks/hyperspell-go/packages/pagination"
	"github.com/stainless-sdks/hyperspell-go/packages/param"
	"github.com/stainless-sdks/hyperspell-go/packages/respjson"
)

// VaultService contains methods and other services that help with interacting with
// the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVaultService] method instead.
type VaultService struct {
	options []option.RequestOption
}

// NewVaultService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVaultService(opts ...option.RequestOption) (r VaultService) {
	r = VaultService{}
	r.options = opts
	return
}

// This endpoint lists all collections, and how many documents are in each
// collection. All documents that do not have a collection assigned are in the
// `null` collection.
//
// Deprecated: This method will be removed in the future
func (r *VaultService) List(ctx context.Context, query VaultListParams, opts ...option.RequestOption) (res *pagination.CursorPage[VaultListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "vault/list"
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

// This endpoint lists all collections, and how many documents are in each
// collection. All documents that do not have a collection assigned are in the
// `null` collection.
//
// Deprecated: This method will be removed in the future
func (r *VaultService) ListAutoPaging(ctx context.Context, query VaultListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[VaultListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type VaultListResponse struct {
	Collection    string `json:"collection" api:"required"`
	DocumentCount int64  `json:"document_count" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Collection    respjson.Field
		DocumentCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VaultListResponse) RawJSON() string { return r.JSON.raw }
func (r *VaultListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VaultListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Size   param.Opt[int64]  `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VaultListParams]'s query parameters as `url.Values`.
func (r VaultListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
