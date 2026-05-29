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
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// ConnectionService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectionService] method instead.
type ConnectionService struct {
	options []option.RequestOption
}

// NewConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectionService(opts ...option.RequestOption) (r ConnectionService) {
	r = ConnectionService{}
	r.options = opts
	return
}

// List all connections for the user.
func (r *ConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *ConnectionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "connections/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revokes Hyperspell's access the given provider and deletes all stored
// credentials and indexed data.
func (r *ConnectionService) Revoke(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *ConnectionRevokeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("connections/%s/revoke", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ConnectionListResponse struct {
	Connections []ConnectionListResponseConnection `json:"connections" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Connections respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionListResponseConnection struct {
	// The connection's id
	ID string `json:"id" api:"required"`
	// The connection's integration id
	IntegrationID string `json:"integration_id" api:"required"`
	// The connection's label
	Label string `json:"label" api:"required"`
	// The connection's provider
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield".
	Provider string `json:"provider" api:"required"`
	// Count of items in user_options.channels (Teams: workspaces selected; 0 means
	// nothing is being indexed for integrations that require selection).
	SelectedCount int64 `json:"selected_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		IntegrationID respjson.Field
		Label         respjson.Field
		Provider      respjson.Field
		SelectedCount respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionListResponseConnection) RawJSON() string { return r.JSON.raw }
func (r *ConnectionListResponseConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionRevokeResponse struct {
	Message string `json:"message" api:"required"`
	Success bool   `json:"success" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionRevokeResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectionRevokeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
