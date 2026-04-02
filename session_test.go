// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/hyperspell-go"
	"github.com/stainless-sdks/hyperspell-go/internal/testutil"
	"github.com/stainless-sdks/hyperspell-go/option"
)

func TestSessionAddWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := hyperspell.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
		option.WithUserID("My User ID"),
	)
	_, err := client.Sessions.Add(context.TODO(), hyperspell.SessionAddParams{
		History: "history",
		Date:    hyperspell.Time(time.Now()),
		Extract: []string{"procedure"},
		Format:  hyperspell.SessionAddParamsFormatVercel,
		Metadata: map[string]hyperspell.SessionAddParamsMetadataUnion{
			"foo": {
				OfString: hyperspell.String("string"),
			},
		},
		SessionID: hyperspell.String("session_id"),
		Title:     hyperspell.String("title"),
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
