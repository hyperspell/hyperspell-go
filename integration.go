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

// IntegrationService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationService] method instead.
type IntegrationService struct {
	options        []option.RequestOption
	GoogleCalendar IntegrationGoogleCalendarService
	WebCrawler     IntegrationWebCrawlerService
	Slack          IntegrationSlackService
}

// NewIntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntegrationService(opts ...option.RequestOption) (r IntegrationService) {
	r = IntegrationService{}
	r.options = opts
	r.GoogleCalendar = NewIntegrationGoogleCalendarService(opts...)
	r.WebCrawler = NewIntegrationWebCrawlerService(opts...)
	r.Slack = NewIntegrationSlackService(opts...)
	return
}

// List all integrations for the user.
func (r *IntegrationService) List(ctx context.Context, opts ...option.RequestOption) (res *IntegrationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Redirects to the connect URL to link an integration.
func (r *IntegrationService) Connect(ctx context.Context, integrationID string, query IntegrationConnectParams, opts ...option.RequestOption) (res *IntegrationConnectResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if integrationID == "" {
		err = errors.New("missing required integration_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("integrations/%s/connect", url.PathEscape(integrationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type IntegrationListResponse struct {
	Integrations []IntegrationListResponseIntegration `json:"integrations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Integrations respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationListResponseIntegration struct {
	// The integration's id
	ID string `json:"id" api:"required"`
	// Whether the integration allows multiple connections
	AllowMultipleConnections bool `json:"allow_multiple_connections" api:"required"`
	// The integration's auth provider
	//
	// Any of "nango", "unified", "whitelabel".
	AuthProvider string `json:"auth_provider" api:"required"`
	// URL to the integration's icon
	Icon string `json:"icon" api:"required"`
	// The integration's display name
	Name string `json:"name" api:"required"`
	// The integration's provider
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "linear", "hubspot",
	// "salesforce", "coda".
	Provider string `json:"provider" api:"required"`
	// Whether this integration only supports write actions (no sync)
	ActionsOnly bool `json:"actions_only"`
	// Whether the user must select channels before indexing starts
	RequiresChannelSelection bool `json:"requires_channel_selection"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                       respjson.Field
		AllowMultipleConnections respjson.Field
		AuthProvider             respjson.Field
		Icon                     respjson.Field
		Name                     respjson.Field
		Provider                 respjson.Field
		ActionsOnly              respjson.Field
		RequiresChannelSelection respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationListResponseIntegration) RawJSON() string { return r.JSON.raw }
func (r *IntegrationListResponseIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationConnectResponse struct {
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	URL       string    `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationConnectResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationConnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationConnectParams struct {
	RedirectURL param.Opt[string] `query:"redirect_url,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationConnectParams]'s query parameters as
// `url.Values`.
func (r IntegrationConnectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
