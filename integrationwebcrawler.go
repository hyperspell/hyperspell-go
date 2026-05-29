// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/apiquery"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// IntegrationWebCrawlerService contains methods and other services that help with
// interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationWebCrawlerService] method instead.
type IntegrationWebCrawlerService struct {
	options []option.RequestOption
}

// NewIntegrationWebCrawlerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationWebCrawlerService(opts ...option.RequestOption) (r IntegrationWebCrawlerService) {
	r = IntegrationWebCrawlerService{}
	r.options = opts
	return
}

// Recursively crawl a website to make it available for indexed search.
func (r *IntegrationWebCrawlerService) Index(ctx context.Context, query IntegrationWebCrawlerIndexParams, opts ...option.RequestOption) (res *IntegrationWebCrawlerIndexResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations/web_crawler/index"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type IntegrationWebCrawlerIndexResponse struct {
	ResourceID string `json:"resource_id" api:"required"`
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "github", "google_drive", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions", "granola", "fathom", "fireflies", "linear",
	// "hubspot", "salesforce", "coda", "lightfield".
	Source IntegrationWebCrawlerIndexResponseSource `json:"source" api:"required"`
	// Any of "pending", "processing", "completed", "failed", "pending_review",
	// "skipped".
	Status IntegrationWebCrawlerIndexResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ResourceID  respjson.Field
		Source      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationWebCrawlerIndexResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationWebCrawlerIndexResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationWebCrawlerIndexResponseSource string

const (
	IntegrationWebCrawlerIndexResponseSourceReddit         IntegrationWebCrawlerIndexResponseSource = "reddit"
	IntegrationWebCrawlerIndexResponseSourceNotion         IntegrationWebCrawlerIndexResponseSource = "notion"
	IntegrationWebCrawlerIndexResponseSourceSlack          IntegrationWebCrawlerIndexResponseSource = "slack"
	IntegrationWebCrawlerIndexResponseSourceGoogleCalendar IntegrationWebCrawlerIndexResponseSource = "google_calendar"
	IntegrationWebCrawlerIndexResponseSourceGoogleMail     IntegrationWebCrawlerIndexResponseSource = "google_mail"
	IntegrationWebCrawlerIndexResponseSourceBox            IntegrationWebCrawlerIndexResponseSource = "box"
	IntegrationWebCrawlerIndexResponseSourceDropbox        IntegrationWebCrawlerIndexResponseSource = "dropbox"
	IntegrationWebCrawlerIndexResponseSourceGitHub         IntegrationWebCrawlerIndexResponseSource = "github"
	IntegrationWebCrawlerIndexResponseSourceGoogleDrive    IntegrationWebCrawlerIndexResponseSource = "google_drive"
	IntegrationWebCrawlerIndexResponseSourceVault          IntegrationWebCrawlerIndexResponseSource = "vault"
	IntegrationWebCrawlerIndexResponseSourceWebCrawler     IntegrationWebCrawlerIndexResponseSource = "web_crawler"
	IntegrationWebCrawlerIndexResponseSourceTrace          IntegrationWebCrawlerIndexResponseSource = "trace"
	IntegrationWebCrawlerIndexResponseSourceMicrosoftTeams IntegrationWebCrawlerIndexResponseSource = "microsoft_teams"
	IntegrationWebCrawlerIndexResponseSourceGmailActions   IntegrationWebCrawlerIndexResponseSource = "gmail_actions"
	IntegrationWebCrawlerIndexResponseSourceGranola        IntegrationWebCrawlerIndexResponseSource = "granola"
	IntegrationWebCrawlerIndexResponseSourceFathom         IntegrationWebCrawlerIndexResponseSource = "fathom"
	IntegrationWebCrawlerIndexResponseSourceFireflies      IntegrationWebCrawlerIndexResponseSource = "fireflies"
	IntegrationWebCrawlerIndexResponseSourceLinear         IntegrationWebCrawlerIndexResponseSource = "linear"
	IntegrationWebCrawlerIndexResponseSourceHubspot        IntegrationWebCrawlerIndexResponseSource = "hubspot"
	IntegrationWebCrawlerIndexResponseSourceSalesforce     IntegrationWebCrawlerIndexResponseSource = "salesforce"
	IntegrationWebCrawlerIndexResponseSourceCoda           IntegrationWebCrawlerIndexResponseSource = "coda"
	IntegrationWebCrawlerIndexResponseSourceLightfield     IntegrationWebCrawlerIndexResponseSource = "lightfield"
)

type IntegrationWebCrawlerIndexResponseStatus string

const (
	IntegrationWebCrawlerIndexResponseStatusPending       IntegrationWebCrawlerIndexResponseStatus = "pending"
	IntegrationWebCrawlerIndexResponseStatusProcessing    IntegrationWebCrawlerIndexResponseStatus = "processing"
	IntegrationWebCrawlerIndexResponseStatusCompleted     IntegrationWebCrawlerIndexResponseStatus = "completed"
	IntegrationWebCrawlerIndexResponseStatusFailed        IntegrationWebCrawlerIndexResponseStatus = "failed"
	IntegrationWebCrawlerIndexResponseStatusPendingReview IntegrationWebCrawlerIndexResponseStatus = "pending_review"
	IntegrationWebCrawlerIndexResponseStatusSkipped       IntegrationWebCrawlerIndexResponseStatus = "skipped"
)

type IntegrationWebCrawlerIndexParams struct {
	// The base URL of the website to crawl
	URL string `query:"url" api:"required" json:"-"`
	// Maximum number of pages to crawl in total
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Maximum depth of links to follow during crawling
	MaxDepth param.Opt[int64] `query:"max_depth,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [IntegrationWebCrawlerIndexParams]'s query parameters as
// `url.Values`.
func (r IntegrationWebCrawlerIndexParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
