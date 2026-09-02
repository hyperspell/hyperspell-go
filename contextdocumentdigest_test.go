// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/internal/testutil"
	"github.com/hyperspell/hyperspell-go/option"
)

func TestContextDocumentDigestListWithOptionalParams(t *testing.T) {
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
	_, err := client.ContextDocuments.Digests.List(context.TODO(), hyperspell.ContextDocumentDigestListParams{
		Limit:  hyperspell.Int(0),
		Period: hyperspell.String("period"),
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContextDocumentDigestGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.ContextDocuments.Digests.Generate(context.TODO(), hyperspell.ContextDocumentDigestGenerateParams{
		Period:      hyperspell.String("period"),
		Sources:     []string{"string"},
		WindowEnd:   hyperspell.Time(time.Now()),
		WindowStart: hyperspell.Time(time.Now()),
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
