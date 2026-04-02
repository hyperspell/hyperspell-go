// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/hyperspell-go/internal/apijson"
	"github.com/stainless-sdks/hyperspell-go/internal/apiquery"
	"github.com/stainless-sdks/hyperspell-go/internal/requestconfig"
	"github.com/stainless-sdks/hyperspell-go/option"
	"github.com/stainless-sdks/hyperspell-go/packages/param"
	"github.com/stainless-sdks/hyperspell-go/packages/respjson"
)

// FolderService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFolderService] method instead.
type FolderService struct {
	options []option.RequestOption
}

// NewFolderService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFolderService(opts ...option.RequestOption) (r FolderService) {
	r = FolderService{}
	r.options = opts
	return
}

// List one level of folders from the user's connected source.
//
// Returns folders decorated with their explicit folder policy (if any). Use
// parent_id to drill into subfolders.
func (r *FolderService) List(ctx context.Context, connectionID string, query FolderListParams, opts ...option.RequestOption) (res *FolderListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/folders", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a folder policy for a specific connection.
func (r *FolderService) DeletePolicy(ctx context.Context, policyID string, body FolderDeletePolicyParams, opts ...option.RequestOption) (res *FolderDeletePolicyResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.ConnectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	if policyID == "" {
		err = errors.New("missing required policy_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/folder-policies/%s", body.ConnectionID, policyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List all folder policies for a specific connection.
func (r *FolderService) ListPolicies(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *FolderListPoliciesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/folder-policies", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create or update a folder policy for a specific connection.
func (r *FolderService) SetPolicies(ctx context.Context, connectionID string, body FolderSetPoliciesParams, opts ...option.RequestOption) (res *FolderSetPoliciesResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/folder-policies", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FolderListResponse struct {
	// Folders at this level
	Folders []FolderListResponseFolder `json:"folders" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Folders     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderListResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderListResponseFolder struct {
	// Whether this folder may have subfolders
	HasChildren bool `json:"has_children" api:"required"`
	// Display name of the folder
	Name string `json:"name" api:"required"`
	// Folder ID from the source provider
	ProviderFolderID string `json:"provider_folder_id" api:"required"`
	// Parent folder's provider ID
	ParentFolderID string `json:"parent_folder_id" api:"nullable"`
	// Explicit policy on this folder, or null if inheriting/default
	Policy FolderListResponseFolderPolicy `json:"policy" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasChildren      respjson.Field
		Name             respjson.Field
		ProviderFolderID respjson.Field
		ParentFolderID   respjson.Field
		Policy           respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderListResponseFolder) RawJSON() string { return r.JSON.raw }
func (r *FolderListResponseFolder) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Explicit policy on this folder, or null if inheriting/default
type FolderListResponseFolderPolicy struct {
	// Policy UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Sync mode set on this folder
	//
	// Any of "sync", "skip", "manual".
	SyncMode string `json:"sync_mode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		SyncMode    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderListResponseFolderPolicy) RawJSON() string { return r.JSON.raw }
func (r *FolderListResponseFolderPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderDeletePolicyResponse struct {
	// Whether the deletion was successful
	Success bool `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderDeletePolicyResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderDeletePolicyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderListPoliciesResponse struct {
	// List of folder policies
	Policies []FolderListPoliciesResponsePolicy `json:"policies" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Policies    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderListPoliciesResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderListPoliciesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderListPoliciesResponsePolicy struct {
	// Unique policy ID
	ID string `json:"id" api:"required" format:"uuid"`
	// Folder ID from the source provider
	ProviderFolderID string `json:"provider_folder_id" api:"required"`
	// Sync mode for this folder
	//
	// Any of "sync", "skip", "manual".
	SyncMode string `json:"sync_mode" api:"required"`
	// Connection ID (null for integration defaults)
	ConnectionID string `json:"connection_id" api:"nullable" format:"uuid"`
	// Display name of the folder
	FolderName string `json:"folder_name" api:"nullable"`
	// Display path of the folder
	FolderPath string `json:"folder_path" api:"nullable"`
	// Parent folder's provider ID
	ParentFolderID string `json:"parent_folder_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ProviderFolderID respjson.Field
		SyncMode         respjson.Field
		ConnectionID     respjson.Field
		FolderName       respjson.Field
		FolderPath       respjson.Field
		ParentFolderID   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderListPoliciesResponsePolicy) RawJSON() string { return r.JSON.raw }
func (r *FolderListPoliciesResponsePolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FolderSetPoliciesResponse struct {
	// Unique policy ID
	ID string `json:"id" api:"required" format:"uuid"`
	// Folder ID from the source provider
	ProviderFolderID string `json:"provider_folder_id" api:"required"`
	// Sync mode for this folder
	//
	// Any of "sync", "skip", "manual".
	SyncMode FolderSetPoliciesResponseSyncMode `json:"sync_mode" api:"required"`
	// Connection ID (null for integration defaults)
	ConnectionID string `json:"connection_id" api:"nullable" format:"uuid"`
	// Display name of the folder
	FolderName string `json:"folder_name" api:"nullable"`
	// Display path of the folder
	FolderPath string `json:"folder_path" api:"nullable"`
	// Parent folder's provider ID
	ParentFolderID string `json:"parent_folder_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ProviderFolderID respjson.Field
		SyncMode         respjson.Field
		ConnectionID     respjson.Field
		FolderName       respjson.Field
		FolderPath       respjson.Field
		ParentFolderID   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FolderSetPoliciesResponse) RawJSON() string { return r.JSON.raw }
func (r *FolderSetPoliciesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sync mode for this folder
type FolderSetPoliciesResponseSyncMode string

const (
	FolderSetPoliciesResponseSyncModeSync   FolderSetPoliciesResponseSyncMode = "sync"
	FolderSetPoliciesResponseSyncModeSkip   FolderSetPoliciesResponseSyncMode = "skip"
	FolderSetPoliciesResponseSyncModeManual FolderSetPoliciesResponseSyncMode = "manual"
)

type FolderListParams struct {
	// Parent folder ID. Omit for root-level folders.
	ParentID param.Opt[string] `query:"parent_id,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FolderListParams]'s query parameters as `url.Values`.
func (r FolderListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FolderDeletePolicyParams struct {
	ConnectionID string `path:"connection_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type FolderSetPoliciesParams struct {
	// Folder ID from the source provider
	ProviderFolderID string `json:"provider_folder_id" api:"required"`
	// Sync mode for this folder
	//
	// Any of "sync", "skip", "manual".
	SyncMode FolderSetPoliciesParamsSyncMode `json:"sync_mode,omitzero" api:"required"`
	// Display name of the folder
	FolderName param.Opt[string] `json:"folder_name,omitzero"`
	// Display path of the folder
	FolderPath param.Opt[string] `json:"folder_path,omitzero"`
	// Parent folder's provider ID for inheritance resolution
	ParentFolderID param.Opt[string] `json:"parent_folder_id,omitzero"`
	paramObj
}

func (r FolderSetPoliciesParams) MarshalJSON() (data []byte, err error) {
	type shadow FolderSetPoliciesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FolderSetPoliciesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sync mode for this folder
type FolderSetPoliciesParamsSyncMode string

const (
	FolderSetPoliciesParamsSyncModeSync   FolderSetPoliciesParamsSyncMode = "sync"
	FolderSetPoliciesParamsSyncModeSkip   FolderSetPoliciesParamsSyncMode = "skip"
	FolderSetPoliciesParamsSyncModeManual FolderSetPoliciesParamsSyncMode = "manual"
)
