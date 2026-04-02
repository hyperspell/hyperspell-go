// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/hyperspell/hyperspell-go/internal/apiquery"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/param"
)

// IntegrationSlackService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationSlackService] method instead.
type IntegrationSlackService struct {
	options []option.RequestOption
}

// NewIntegrationSlackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationSlackService(opts ...option.RequestOption) (r IntegrationSlackService) {
	r = IntegrationSlackService{}
	r.options = opts
	return
}

// List Slack conversations accessible to the user via the live Nango connection.
//
// Returns minimal channel metadata suitable for selection UIs. If required scopes
// are missing, Slack's error is propagated with details.
//
// Supports filtering by channels, including/excluding private channels, DMs, group
// DMs, and archived channels based on the provided search options.
func (r *IntegrationSlackService) List(ctx context.Context, query IntegrationSlackListParams, opts ...option.RequestOption) (res *IntegrationSlackListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations/slack/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type IntegrationSlackListResponse = any

type IntegrationSlackListParams struct {
	// If set, pass 'exclude_archived' to Slack. If None, omit the param.
	ExcludeArchived param.Opt[bool] `query:"exclude_archived,omitzero" json:"-"`
	// Include direct messages (im) when listing conversations.
	IncludeDms param.Opt[bool] `query:"include_dms,omitzero" json:"-"`
	// Include group DMs (mpim) when listing conversations.
	IncludeGroupDms param.Opt[bool] `query:"include_group_dms,omitzero" json:"-"`
	// Include private channels when constructing Slack 'types'.
	IncludePrivate param.Opt[bool] `query:"include_private,omitzero" json:"-"`
	// List of Slack channels to include (by id, name, or #name).
	Channels []string `query:"channels,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationSlackListParams]'s query parameters as
// `url.Values`.
func (r IntegrationSlackListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
