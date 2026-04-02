// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/hyperspell-go/internal/apijson"
	"github.com/stainless-sdks/hyperspell-go/internal/requestconfig"
	"github.com/stainless-sdks/hyperspell-go/option"
	"github.com/stainless-sdks/hyperspell-go/packages/param"
	"github.com/stainless-sdks/hyperspell-go/packages/respjson"
)

// ActionService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	options []option.RequestOption
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.options = opts
	return
}

// Add an emoji reaction to a message on a connected integration.
func (r *ActionService) AddReaction(ctx context.Context, body ActionAddReactionParams, opts ...option.RequestOption) (res *ActionAddReactionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "actions/add_reaction"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Send a message to a channel or conversation on a connected integration.
func (r *ActionService) SendMessage(ctx context.Context, body ActionSendMessageParams, opts ...option.RequestOption) (res *ActionSendMessageResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "actions/send_message"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Result from executing an integration action.
type ActionAddReactionResponse struct {
	Success          bool           `json:"success" api:"required"`
	Error            string         `json:"error" api:"nullable"`
	ProviderResponse map[string]any `json:"provider_response" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success          respjson.Field
		Error            respjson.Field
		ProviderResponse respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionAddReactionResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionAddReactionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result from executing an integration action.
type ActionSendMessageResponse struct {
	Success          bool           `json:"success" api:"required"`
	Error            string         `json:"error" api:"nullable"`
	ProviderResponse map[string]any `json:"provider_response" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success          respjson.Field
		Error            respjson.Field
		ProviderResponse respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionSendMessageResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionSendMessageResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionAddReactionParams struct {
	// Channel ID containing the message
	Channel string `json:"channel" api:"required"`
	// Emoji name without colons (e.g., thumbsup)
	Name string `json:"name" api:"required"`
	// Integration provider (e.g., slack)
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "google_drive", "github", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
	Provider ActionAddReactionParamsProvider `json:"provider,omitzero" api:"required"`
	// Message timestamp to react to
	Timestamp string `json:"timestamp" api:"required"`
	// Connection ID. If omitted, auto-resolved from provider + user.
	Connection param.Opt[string] `json:"connection,omitzero"`
	paramObj
}

func (r ActionAddReactionParams) MarshalJSON() (data []byte, err error) {
	type shadow ActionAddReactionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionAddReactionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration provider (e.g., slack)
type ActionAddReactionParamsProvider string

const (
	ActionAddReactionParamsProviderReddit         ActionAddReactionParamsProvider = "reddit"
	ActionAddReactionParamsProviderNotion         ActionAddReactionParamsProvider = "notion"
	ActionAddReactionParamsProviderSlack          ActionAddReactionParamsProvider = "slack"
	ActionAddReactionParamsProviderGoogleCalendar ActionAddReactionParamsProvider = "google_calendar"
	ActionAddReactionParamsProviderGoogleMail     ActionAddReactionParamsProvider = "google_mail"
	ActionAddReactionParamsProviderBox            ActionAddReactionParamsProvider = "box"
	ActionAddReactionParamsProviderDropbox        ActionAddReactionParamsProvider = "dropbox"
	ActionAddReactionParamsProviderGoogleDrive    ActionAddReactionParamsProvider = "google_drive"
	ActionAddReactionParamsProviderGitHub         ActionAddReactionParamsProvider = "github"
	ActionAddReactionParamsProviderVault          ActionAddReactionParamsProvider = "vault"
	ActionAddReactionParamsProviderWebCrawler     ActionAddReactionParamsProvider = "web_crawler"
	ActionAddReactionParamsProviderTrace          ActionAddReactionParamsProvider = "trace"
	ActionAddReactionParamsProviderMicrosoftTeams ActionAddReactionParamsProvider = "microsoft_teams"
	ActionAddReactionParamsProviderGmailActions   ActionAddReactionParamsProvider = "gmail_actions"
)

type ActionSendMessageParams struct {
	// Integration provider (e.g., slack)
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "google_drive", "github", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
	Provider ActionSendMessageParamsProvider `json:"provider,omitzero" api:"required"`
	// Message text
	Text string `json:"text" api:"required"`
	// Channel ID (required for Slack)
	Channel param.Opt[string] `json:"channel,omitzero"`
	// Connection ID. If omitted, auto-resolved from provider + user.
	Connection param.Opt[string] `json:"connection,omitzero"`
	// Parent message ID for threading (thread_ts for Slack)
	Parent param.Opt[string] `json:"parent,omitzero"`
	paramObj
}

func (r ActionSendMessageParams) MarshalJSON() (data []byte, err error) {
	type shadow ActionSendMessageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionSendMessageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration provider (e.g., slack)
type ActionSendMessageParamsProvider string

const (
	ActionSendMessageParamsProviderReddit         ActionSendMessageParamsProvider = "reddit"
	ActionSendMessageParamsProviderNotion         ActionSendMessageParamsProvider = "notion"
	ActionSendMessageParamsProviderSlack          ActionSendMessageParamsProvider = "slack"
	ActionSendMessageParamsProviderGoogleCalendar ActionSendMessageParamsProvider = "google_calendar"
	ActionSendMessageParamsProviderGoogleMail     ActionSendMessageParamsProvider = "google_mail"
	ActionSendMessageParamsProviderBox            ActionSendMessageParamsProvider = "box"
	ActionSendMessageParamsProviderDropbox        ActionSendMessageParamsProvider = "dropbox"
	ActionSendMessageParamsProviderGoogleDrive    ActionSendMessageParamsProvider = "google_drive"
	ActionSendMessageParamsProviderGitHub         ActionSendMessageParamsProvider = "github"
	ActionSendMessageParamsProviderVault          ActionSendMessageParamsProvider = "vault"
	ActionSendMessageParamsProviderWebCrawler     ActionSendMessageParamsProvider = "web_crawler"
	ActionSendMessageParamsProviderTrace          ActionSendMessageParamsProvider = "trace"
	ActionSendMessageParamsProviderMicrosoftTeams ActionSendMessageParamsProvider = "microsoft_teams"
	ActionSendMessageParamsProviderGmailActions   ActionSendMessageParamsProvider = "gmail_actions"
)
