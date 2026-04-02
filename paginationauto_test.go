// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell_test

import (
	"context"
	"os"
	"testing"

	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/internal/testutil"
	"github.com/hyperspell/hyperspell-go/option"
)

func TestAutoPagination(t *testing.T) {
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
	iter := client.Memories.ListAutoPaging(context.TODO(), hyperspell.MemoryListParams{})
	// The mock server isn't going to give us real pagination
	for i := 0; i < 3 && iter.Next(); i++ {
		memory := iter.Current()
		t.Logf("%+v\n", memory.ResourceID)
	}
	if err := iter.Err(); err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
