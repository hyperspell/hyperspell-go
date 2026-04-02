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

func TestManualPagination(t *testing.T) {
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
	page, err := client.Memories.List(context.TODO(), hyperspell.MemoryListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	for _, memory := range page.Items {
		t.Logf("%+v\n", memory.ResourceID)
	}
	// The mock server isn't going to give us real pagination
	page, err = page.GetNextPage()
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if page != nil {
		for _, memory := range page.Items {
			t.Logf("%+v\n", memory.ResourceID)
		}
	}
}
