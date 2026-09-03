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

// EntityService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEntityService] method instead.
type EntityService struct {
	options []option.RequestOption
}

// NewEntityService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEntityService(opts ...option.RequestOption) (r EntityService) {
	r = EntityService{}
	r.options = opts
	return
}

// List entities available to the current app.
//
// Results can be filtered by type, status, name, and supporting-document count.
// Use the returned cursor to retrieve the next page.
func (r *EntityService) List(ctx context.Context, query EntityListParams, opts ...option.RequestOption) (res *pagination.EntityCursorPage[EntityListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "entities"
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

// List entities available to the current app.
//
// Results can be filtered by type, status, name, and supporting-document count.
// Use the returned cursor to retrieve the next page.
func (r *EntityService) ListAutoPaging(ctx context.Context, query EntityListParams, opts ...option.RequestOption) *pagination.EntityCursorPageAutoPager[EntityListResponse] {
	return pagination.NewEntityCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Fetch a single entity belonging to the current app.
//
// Returns 404 when the entity does not exist or is not visible to the app.
func (r *EntityService) Get(ctx context.Context, entityID string, opts ...option.RequestOption) (res *EntityGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if entityID == "" {
		err = errors.New("missing required entity_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("entities/%s", entityID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Search the current app's entities by meaning.
func (r *EntityService) Search(ctx context.Context, body EntitySearchParams, opts ...option.RequestOption) (res *EntitySearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "entities/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type EntityListResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// How strongly the entity's current identity has been established.
	//
	// Any of "provisional", "confirmed".
	Status                  EntityListResponseStatus `json:"status" api:"required"`
	Type                    string                   `json:"type" api:"required"`
	UpdatedAt               time.Time                `json:"updated_at" api:"required" format:"date-time"`
	Attributes              map[string]any           `json:"attributes"`
	Description             string                   `json:"description" api:"nullable"`
	HardLinkedMentionCount  int64                    `json:"hard_linked_mention_count" api:"nullable"`
	ProminenceCalculatedAt  time.Time                `json:"prominence_calculated_at" api:"nullable" format:"date-time"`
	ProminenceVersion       string                   `json:"prominence_version" api:"nullable"`
	RecordCount             int64                    `json:"record_count" api:"nullable"`
	SupportingDocumentCount int64                    `json:"supporting_document_count" api:"nullable"`
	SupportingScopeCount    int64                    `json:"supporting_scope_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Attributes              respjson.Field
		Description             respjson.Field
		HardLinkedMentionCount  respjson.Field
		ProminenceCalculatedAt  respjson.Field
		ProminenceVersion       respjson.Field
		RecordCount             respjson.Field
		SupportingDocumentCount respjson.Field
		SupportingScopeCount    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityListResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How strongly the entity's current identity has been established.
type EntityListResponseStatus string

const (
	EntityListResponseStatusProvisional EntityListResponseStatus = "provisional"
	EntityListResponseStatusConfirmed   EntityListResponseStatus = "confirmed"
)

type EntityGetResponse struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// How strongly the entity's current identity has been established.
	//
	// Any of "provisional", "confirmed".
	Status                  EntityGetResponseStatus `json:"status" api:"required"`
	Type                    string                  `json:"type" api:"required"`
	UpdatedAt               time.Time               `json:"updated_at" api:"required" format:"date-time"`
	Attributes              map[string]any          `json:"attributes"`
	Description             string                  `json:"description" api:"nullable"`
	HardLinkedMentionCount  int64                   `json:"hard_linked_mention_count" api:"nullable"`
	ProminenceCalculatedAt  time.Time               `json:"prominence_calculated_at" api:"nullable" format:"date-time"`
	ProminenceVersion       string                  `json:"prominence_version" api:"nullable"`
	RecordCount             int64                   `json:"record_count" api:"nullable"`
	SupportingDocumentCount int64                   `json:"supporting_document_count" api:"nullable"`
	SupportingScopeCount    int64                   `json:"supporting_scope_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Attributes              respjson.Field
		Description             respjson.Field
		HardLinkedMentionCount  respjson.Field
		ProminenceCalculatedAt  respjson.Field
		ProminenceVersion       respjson.Field
		RecordCount             respjson.Field
		SupportingDocumentCount respjson.Field
		SupportingScopeCount    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EntityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// How strongly the entity's current identity has been established.
type EntityGetResponseStatus string

const (
	EntityGetResponseStatusProvisional EntityGetResponseStatus = "provisional"
	EntityGetResponseStatusConfirmed   EntityGetResponseStatus = "confirmed"
)

type EntitySearchResponse struct {
	Items []EntitySearchResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *EntitySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntitySearchResponseItem struct {
	ID         string    `json:"id" api:"required" format:"uuid"`
	CreatedAt  time.Time `json:"created_at" api:"required" format:"date-time"`
	Name       string    `json:"name" api:"required"`
	Similarity float64   `json:"similarity" api:"required"`
	// How strongly the entity's current identity has been established.
	//
	// Any of "provisional", "confirmed".
	Status                  string         `json:"status" api:"required"`
	Type                    string         `json:"type" api:"required"`
	UpdatedAt               time.Time      `json:"updated_at" api:"required" format:"date-time"`
	Attributes              map[string]any `json:"attributes"`
	Description             string         `json:"description" api:"nullable"`
	HardLinkedMentionCount  int64          `json:"hard_linked_mention_count" api:"nullable"`
	ProminenceCalculatedAt  time.Time      `json:"prominence_calculated_at" api:"nullable" format:"date-time"`
	ProminenceVersion       string         `json:"prominence_version" api:"nullable"`
	RecordCount             int64          `json:"record_count" api:"nullable"`
	SupportingDocumentCount int64          `json:"supporting_document_count" api:"nullable"`
	SupportingScopeCount    int64          `json:"supporting_scope_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                      respjson.Field
		CreatedAt               respjson.Field
		Name                    respjson.Field
		Similarity              respjson.Field
		Status                  respjson.Field
		Type                    respjson.Field
		UpdatedAt               respjson.Field
		Attributes              respjson.Field
		Description             respjson.Field
		HardLinkedMentionCount  respjson.Field
		ProminenceCalculatedAt  respjson.Field
		ProminenceVersion       respjson.Field
		RecordCount             respjson.Field
		SupportingDocumentCount respjson.Field
		SupportingScopeCount    respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EntitySearchResponseItem) RawJSON() string { return r.JSON.raw }
func (r *EntitySearchResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EntityListParams struct {
	Cursor                 param.Opt[string] `query:"cursor,omitzero" json:"-"`
	MinSupportingDocuments param.Opt[int64]  `query:"min_supporting_documents,omitzero" json:"-"`
	NamePrefix             param.Opt[string] `query:"name_prefix,omitzero" json:"-"`
	Search                 param.Opt[string] `query:"search,omitzero" json:"-"`
	Type                   param.Opt[string] `query:"type,omitzero" json:"-"`
	Limit                  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// How strongly the entity's current identity has been established.
	//
	// Any of "provisional", "confirmed".
	Status EntityListParamsStatus `query:"status,omitzero" json:"-"`
	// Any of "id", "name", "type", "prominence".
	SortBy EntityListParamsSortBy `query:"sort_by,omitzero" json:"-"`
	// Any of "asc", "desc".
	SortDir EntityListParamsSortDir `query:"sort_dir,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EntityListParams]'s query parameters as `url.Values`.
func (r EntityListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EntityListParamsSortBy string

const (
	EntityListParamsSortByID         EntityListParamsSortBy = "id"
	EntityListParamsSortByName       EntityListParamsSortBy = "name"
	EntityListParamsSortByType       EntityListParamsSortBy = "type"
	EntityListParamsSortByProminence EntityListParamsSortBy = "prominence"
)

type EntityListParamsSortDir string

const (
	EntityListParamsSortDirAsc  EntityListParamsSortDir = "asc"
	EntityListParamsSortDirDesc EntityListParamsSortDir = "desc"
)

// How strongly the entity's current identity has been established.
type EntityListParamsStatus string

const (
	EntityListParamsStatusProvisional EntityListParamsStatus = "provisional"
	EntityListParamsStatusConfirmed   EntityListParamsStatus = "confirmed"
)

type EntitySearchParams struct {
	Query string            `json:"query" api:"required"`
	Type  param.Opt[string] `json:"type,omitzero"`
	Limit param.Opt[int64]  `json:"limit,omitzero"`
	paramObj
}

func (r EntitySearchParams) MarshalJSON() (data []byte, err error) {
	type shadow EntitySearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EntitySearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
