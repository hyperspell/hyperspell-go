// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/param"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// AuthService contains methods and other services that help with interacting with
// the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuthService] method instead.
type AuthService struct {
	options []option.RequestOption
}

// NewAuthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAuthService(opts ...option.RequestOption) (r AuthService) {
	r = AuthService{}
	r.options = opts
	return
}

// Endpoint to delete user.
func (r *AuthService) DeleteUser(ctx context.Context, opts ...option.RequestOption) (res *AuthDeleteUserResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "auth/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Endpoint to get basic user data.
func (r *AuthService) Me(ctx context.Context, opts ...option.RequestOption) (res *AuthMeResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "auth/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Use this endpoint to create a user token for a specific user. This token can be
// safely passed to your user-facing front-end.
func (r *AuthService) UserToken(ctx context.Context, body AuthUserTokenParams, opts ...option.RequestOption) (res *Token, err error) {
	opts = slices.Concat(r.options, opts)
	path := "auth/user_token"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type Token struct {
	Token     string    `json:"token" api:"required"`
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Token) RawJSON() string { return r.JSON.raw }
func (r *Token) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthDeleteUserResponse struct {
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
func (r AuthDeleteUserResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthDeleteUserResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthMeResponse struct {
	// The user's id
	ID string `json:"id" api:"required"`
	// The Hyperspell app's id this user belongs to
	App AuthMeResponseApp `json:"app" api:"required"`
	// All integrations available for the app
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "google_drive", "github", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
	AvailableIntegrations []string `json:"available_integrations" api:"required"`
	// All integrations installed for the user
	//
	// Any of "reddit", "notion", "slack", "google_calendar", "google_mail", "box",
	// "dropbox", "google_drive", "github", "vault", "web_crawler", "trace",
	// "microsoft_teams", "gmail_actions".
	InstalledIntegrations []string `json:"installed_integrations" api:"required"`
	// The expiration time of the user token used to make the request
	TokenExpiration time.Time `json:"token_expiration" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		App                   respjson.Field
		AvailableIntegrations respjson.Field
		InstalledIntegrations respjson.Field
		TokenExpiration       respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthMeResponse) RawJSON() string { return r.JSON.raw }
func (r *AuthMeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Hyperspell app's id this user belongs to
type AuthMeResponseApp struct {
	// The Hyperspell app's id this user belongs to
	ID string `json:"id" api:"required"`
	// The app's icon
	IconURL string `json:"icon_url" api:"required"`
	// The app's name
	Name string `json:"name" api:"required"`
	// The app's redirect URL
	RedirectURL string `json:"redirect_url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IconURL     respjson.Field
		Name        respjson.Field
		RedirectURL respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuthMeResponseApp) RawJSON() string { return r.JSON.raw }
func (r *AuthMeResponseApp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuthUserTokenParams struct {
	UserID string `json:"user_id" api:"required"`
	// Token lifetime, e.g., '30m', '2h', '1d'. Defaults to 24 hours if not provided.
	ExpiresIn param.Opt[string] `json:"expires_in,omitzero"`
	// Origin of the request, used for CSRF protection. If set, the token will only be
	// valid for requests originating from this origin.
	Origin param.Opt[string] `json:"origin,omitzero"`
	paramObj
}

func (r AuthUserTokenParams) MarshalJSON() (data []byte, err error) {
	type shadow AuthUserTokenParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AuthUserTokenParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
