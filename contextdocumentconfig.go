// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"slices"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// ContextDocumentConfigService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewContextDocumentConfigService] method instead.
type ContextDocumentConfigService struct {
	options []option.RequestOption
}

// NewContextDocumentConfigService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewContextDocumentConfigService(opts ...option.RequestOption) (r ContextDocumentConfigService) {
	r = ContextDocumentConfigService{}
	r.options = opts
	return
}

// Update the supplied generation settings.
//
// Changes apply to the next generation. This endpoint does not start a generation
// or modify existing context documents.
func (r *ContextDocumentConfigService) Update(ctx context.Context, body ContextDocumentConfigUpdateParams, opts ...option.RequestOption) (res *ContextDocumentConfigUpdateResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Read the customer-editable generation config for the authenticated app.
func (r *ContextDocumentConfigService) Get(ctx context.Context, opts ...option.RequestOption) (res *ContextDocumentConfigGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Reset customer-editable generation settings to their defaults.
//
// Existing context documents remain unchanged. `detected_domain` is retained and
// used for future generations unless a new domain override is set.
func (r *ContextDocumentConfigService) Reset(ctx context.Context, opts ...option.RequestOption) (res *ContextDocumentConfigResetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "context-documents/config/reset"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Brain-generation settings that customers can view and edit.
type ContextDocumentConfigUpdateResponse struct {
	Prompts        map[string]any    `json:"prompts" api:"required"`
	SourceWeights  map[string]string `json:"source_weights" api:"required"`
	Structure      map[string]any    `json:"structure" api:"required"`
	DetectedDomain string            `json:"detected_domain" api:"nullable"`
	Domain         string            `json:"domain" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompts        respjson.Field
		SourceWeights  respjson.Field
		Structure      respjson.Field
		DetectedDomain respjson.Field
		Domain         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentConfigUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentConfigUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brain-generation settings that customers can view and edit.
type ContextDocumentConfigGetResponse struct {
	Prompts        map[string]any    `json:"prompts" api:"required"`
	SourceWeights  map[string]string `json:"source_weights" api:"required"`
	Structure      map[string]any    `json:"structure" api:"required"`
	DetectedDomain string            `json:"detected_domain" api:"nullable"`
	Domain         string            `json:"domain" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompts        respjson.Field
		SourceWeights  respjson.Field
		Structure      respjson.Field
		DetectedDomain respjson.Field
		Domain         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentConfigGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentConfigGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Brain-generation settings that customers can view and edit.
type ContextDocumentConfigResetResponse struct {
	Prompts        map[string]any    `json:"prompts" api:"required"`
	SourceWeights  map[string]string `json:"source_weights" api:"required"`
	Structure      map[string]any    `json:"structure" api:"required"`
	DetectedDomain string            `json:"detected_domain" api:"nullable"`
	Domain         string            `json:"domain" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompts        respjson.Field
		SourceWeights  respjson.Field
		Structure      respjson.Field
		DetectedDomain respjson.Field
		Domain         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContextDocumentConfigResetResponse) RawJSON() string { return r.JSON.raw }
func (r *ContextDocumentConfigResetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContextDocumentConfigUpdateParams struct {
	DetectionPrompt param.Opt[string] `json:"detection_prompt,omitzero"`
	Domain          param.Opt[string] `json:"domain,omitzero"`
	PersonalPrompt  param.Opt[string] `json:"personal_prompt,omitzero"`
	CompanyPrompts  map[string]string `json:"company_prompts,omitzero"`
	SourceWeights   map[string]string `json:"source_weights,omitzero"`
	// Per-tier document definitions for custom generation.
	Structure         ContextDocumentConfigUpdateParamsStructure `json:"structure,omitzero"`
	WorkstreamPrompts map[string]string                          `json:"workstream_prompts,omitzero"`
	paramObj
}

func (r ContextDocumentConfigUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentConfigUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentConfigUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Per-tier document definitions for custom generation.
type ContextDocumentConfigUpdateParamsStructure struct {
	Company    []ContextDocumentConfigUpdateParamsStructureCompany    `json:"company,omitzero"`
	Workstream []ContextDocumentConfigUpdateParamsStructureWorkstream `json:"workstream,omitzero"`
	paramObj
}

func (r ContextDocumentConfigUpdateParamsStructure) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentConfigUpdateParamsStructure
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentConfigUpdateParamsStructure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One document in a context-tree tier: what to generate and how to retrieve for
// it.
//
// The properties Filename, Key, Prompt, SearchQueries are required.
type ContextDocumentConfigUpdateParamsStructureCompany struct {
	Filename      string   `json:"filename" api:"required"`
	Key           string   `json:"key" api:"required"`
	Prompt        string   `json:"prompt" api:"required"`
	SearchQueries []string `json:"search_queries,omitzero" api:"required"`
	paramObj
}

func (r ContextDocumentConfigUpdateParamsStructureCompany) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentConfigUpdateParamsStructureCompany
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentConfigUpdateParamsStructureCompany) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One document in a context-tree tier: what to generate and how to retrieve for
// it.
//
// The properties Filename, Key, Prompt, SearchQueries are required.
type ContextDocumentConfigUpdateParamsStructureWorkstream struct {
	Filename      string   `json:"filename" api:"required"`
	Key           string   `json:"key" api:"required"`
	Prompt        string   `json:"prompt" api:"required"`
	SearchQueries []string `json:"search_queries,omitzero" api:"required"`
	paramObj
}

func (r ContextDocumentConfigUpdateParamsStructureWorkstream) MarshalJSON() (data []byte, err error) {
	type shadow ContextDocumentConfigUpdateParamsStructureWorkstream
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ContextDocumentConfigUpdateParamsStructureWorkstream) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
