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
)

// SessionService contains methods and other services that help with interacting
// with the hyperspell API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSessionService] method instead.
type SessionService struct {
	options []option.RequestOption
}

// NewSessionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSessionService(opts ...option.RequestOption) (r SessionService) {
	r = SessionService{}
	r.options = opts
	return
}

// Add an agent trace/transcript to the index.
//
// Accepts traces as a string in Hyperdoc format (native), Vercel AI SDK format, or
// OpenClaw JSONL format. The format is auto-detected if not specified.
//
// **Hyperdoc format** (JSON array, snake_case with type discriminators):
//
// ```json
//
//	{
//	  "history": "[{\"type\": \"trace_message\", \"role\": \"user\", \"text\": \"Hello\"}]"
//	}
//
// ```
//
// **Vercel AI SDK format** (JSON array, camelCase):
//
// ```json
// { "history": "[{\"role\": \"user\", \"content\": \"Hello\"}]" }
// ```
//
// **OpenClaw JSONL format** (newline-delimited JSON):
//
// ```json
//
//	{
//	  "history": "{\"type\":\"session\",\"id\":\"abc\"}\n{\"type\":\"message\",\"message\":{\"role\":\"user\",...}}"
//	}
//
// ```
func (r *SessionService) Add(ctx context.Context, body SessionAddParams, opts ...option.RequestOption) (res *MemoryStatus, err error) {
	opts = slices.Concat(r.options, opts)
	path := "trace/add"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SessionAddParams struct {
	// The trace history as a string. Can be a JSON array of Hyperdoc steps, a JSON
	// array of Vercel AI SDK steps, or OpenClaw JSONL.
	History string `json:"history" api:"required"`
	// Title of the trace
	Title param.Opt[string] `json:"title,omitzero"`
	// Date of the trace
	Date param.Opt[time.Time] `json:"date,omitzero" format:"date-time"`
	// Resource identifier for the trace.
	SessionID param.Opt[string] `json:"session_id,omitzero"`
	// Trace format: 'vercel', 'hyperdoc', or 'openclaw'. Auto-detected if not set.
	//
	// Any of "vercel", "hyperdoc", "openclaw".
	Format SessionAddParamsFormat `json:"format,omitzero"`
	// Custom metadata for filtering. Keys must be alphanumeric with underscores, max
	// 64 chars.
	Metadata map[string]SessionAddParamsMetadataUnion `json:"metadata,omitzero"`
	// What kind of memories to extract from the trace
	//
	// Any of "procedure", "memory", "mood".
	Extract []string `json:"extract,omitzero"`
	paramObj
}

func (r SessionAddParams) MarshalJSON() (data []byte, err error) {
	type shadow SessionAddParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SessionAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trace format: 'vercel', 'hyperdoc', or 'openclaw'. Auto-detected if not set.
type SessionAddParamsFormat string

const (
	SessionAddParamsFormatVercel   SessionAddParamsFormat = "vercel"
	SessionAddParamsFormatHyperdoc SessionAddParamsFormat = "hyperdoc"
	SessionAddParamsFormatOpenclaw SessionAddParamsFormat = "openclaw"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SessionAddParamsMetadataUnion struct {
	OfString param.Opt[string]  `json:",omitzero,inline"`
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]    `json:",omitzero,inline"`
	paramUnion
}

func (u SessionAddParamsMetadataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfFloat, u.OfBool)
}
func (u *SessionAddParamsMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}
