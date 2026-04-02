// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell

import (
	"context"
	"net/http"
	"slices"

	"github.com/hyperspell/hyperspell-go/internal/apijson"
	"github.com/hyperspell/hyperspell-go/internal/requestconfig"
	"github.com/hyperspell/hyperspell-go/option"
	"github.com/hyperspell/hyperspell-go/packages/respjson"
)

// IntegrationGoogleCalendarService contains methods and other services that help
// with interacting with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationGoogleCalendarService] method instead.
type IntegrationGoogleCalendarService struct {
	options []option.RequestOption
}

// NewIntegrationGoogleCalendarService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIntegrationGoogleCalendarService(opts ...option.RequestOption) (r IntegrationGoogleCalendarService) {
	r = IntegrationGoogleCalendarService{}
	r.options = opts
	return
}

// List available calendars for a user. This can be used to ie. populate a dropdown
// for the user to select a calendar.
func (r *IntegrationGoogleCalendarService) List(ctx context.Context, opts ...option.RequestOption) (res *Calendar, err error) {
	opts = slices.Concat(r.options, opts)
	path := "integrations/google_calendar/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type Calendar struct {
	Items []CalendarItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Calendar) RawJSON() string { return r.JSON.raw }
func (r *Calendar) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CalendarItem struct {
	// The ID of the calendar
	ID string `json:"id" api:"required"`
	// The name of the calendar
	Name string `json:"name" api:"required"`
	// Whether the calendar is the primary calendar of the user
	Primary bool `json:"primary" api:"required"`
	// Default timezone of the calendar
	Timezone string `json:"timezone" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		Primary     respjson.Field
		Timezone    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CalendarItem) RawJSON() string { return r.JSON.raw }
func (r *CalendarItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
