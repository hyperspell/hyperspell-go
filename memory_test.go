// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"os"
	"testing"
	"time"

	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/internal/testutil"
	"github.com/hyperspell/hyperspell-go/option"
)

func TestMemoryUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Memories.Update(
		context.TODO(),
		"resource_id",
		hyperspell.MemoryUpdateParams{
			Source:     hyperspell.MemoryUpdateParamsSourceReddit,
			Collection: map[string]any{},
			Date:       map[string]any{},
			Metadata: map[string]hyperspell.MemoryUpdateParamsMetadataUnion{
				"foo": {
					OfString: hyperspell.String("string"),
				},
			},
			Text:  map[string]any{},
			Title: map[string]any{},
		},
	)
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryListWithOptionalParams(t *testing.T) {
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
	_, err := client.Memories.List(context.TODO(), hyperspell.MemoryListParams{
		Collection: hyperspell.String("collection"),
		Cursor:     hyperspell.String("cursor"),
		Filter:     hyperspell.String("filter"),
		Size:       hyperspell.Int(0),
		Source:     hyperspell.MemoryListParamsSourceReddit,
		Status:     hyperspell.MemoryListParamsStatusPending,
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryDelete(t *testing.T) {
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
	_, err := client.Memories.Delete(
		context.TODO(),
		"resource_id",
		hyperspell.MemoryDeleteParams{
			Source: hyperspell.MemoryDeleteParamsSourceReddit,
		},
	)
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryAddWithOptionalParams(t *testing.T) {
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
	_, err := client.Memories.Add(context.TODO(), hyperspell.MemoryAddParams{
		Text:       "...",
		Collection: hyperspell.String("my-collection"),
		Date:       hyperspell.Time(time.Now()),
		Metadata: map[string]hyperspell.MemoryAddParamsMetadataUnion{
			"author": {
				OfString: hyperspell.String("John Doe"),
			},
			"date": {
				OfString: hyperspell.String("2025-05-20T02:31:00Z"),
			},
			"rating": {
				OfFloat: hyperspell.Float(3),
			},
		},
		ResourceID: hyperspell.String("resource_id"),
		Title:      hyperspell.String("My Document"),
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryAddBulk(t *testing.T) {
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
	_, err := client.Memories.AddBulk(context.TODO(), hyperspell.MemoryAddBulkParams{
		Items: []hyperspell.MemoryAddBulkParamsItem{{
			Text:       "...",
			Collection: hyperspell.String("my-collection"),
			Date:       hyperspell.Time(time.Now()),
			Metadata: map[string]hyperspell.MemoryAddBulkParamsItemMetadataUnion{
				"author": {
					OfString: hyperspell.String("John Doe"),
				},
				"date": {
					OfString: hyperspell.String("2025-05-20T02:31:00Z"),
				},
				"rating": {
					OfFloat: hyperspell.Float(3),
				},
			},
			ResourceID: hyperspell.String("resource_id"),
			Title:      hyperspell.String("My Document"),
		}},
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryGet(t *testing.T) {
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
	_, err := client.Memories.Get(
		context.TODO(),
		"resource_id",
		hyperspell.MemoryGetParams{
			Source: hyperspell.MemoryGetParamsSourceReddit,
		},
	)
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemorySearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Memories.Search(context.TODO(), hyperspell.MemorySearchParams{
		Query:      "What does Hyperspell do?",
		Answer:     hyperspell.Bool(true),
		Effort:     hyperspell.MemorySearchParamsEffortMinimal,
		MaxResults: hyperspell.Int(1),
		Options: hyperspell.MemorySearchParamsOptions{
			After:       hyperspell.Time(time.Now()),
			AnswerModel: "llama-3.1",
			Before:      hyperspell.Time(time.Now()),
			Box: hyperspell.MemorySearchParamsOptionsBox{
				Weight: hyperspell.Float(0),
			},
			Filter: map[string]any{},
			GoogleCalendar: hyperspell.MemorySearchParamsOptionsGoogleCalendar{
				CalendarID: hyperspell.String("calendar_id"),
				Weight:     hyperspell.Float(0),
			},
			GoogleDrive: hyperspell.MemorySearchParamsOptionsGoogleDrive{
				Weight: hyperspell.Float(0),
			},
			GoogleMail: hyperspell.MemorySearchParamsOptionsGoogleMail{
				LabelIDs: []string{"string"},
				Weight:   hyperspell.Float(0),
			},
			MaxResults:  hyperspell.Int(1),
			MemoryTypes: []string{"procedure"},
			Notion: hyperspell.MemorySearchParamsOptionsNotion{
				NotionPageIDs: []string{"string"},
				Weight:        hyperspell.Float(0),
			},
			RecencyHalfLifeDays: hyperspell.Float(1),
			ResourceIDs:         []string{"string"},
			Slack: hyperspell.MemorySearchParamsOptionsSlack{
				Channels:        []string{"string"},
				ExcludeArchived: hyperspell.Bool(true),
				IncludeDms:      hyperspell.Bool(true),
				IncludeGroupDms: hyperspell.Bool(true),
				IncludePrivate:  hyperspell.Bool(true),
				Weight:          hyperspell.Float(0),
			},
			Vault: hyperspell.MemorySearchParamsOptionsVault{
				Weight: hyperspell.Float(0),
			},
			WebCrawler: hyperspell.MemorySearchParamsOptionsWebCrawler{
				MaxDepth: hyperspell.Int(0),
				URL:      hyperspell.String("url"),
				Weight:   hyperspell.Float(0),
			},
		},
		Provenance: hyperspell.Bool(true),
		Sources:    []string{"vault"},
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryStatus(t *testing.T) {
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
	_, err := client.Memories.Status(context.TODO())
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMemoryUploadWithOptionalParams(t *testing.T) {
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
	_, err := client.Memories.Upload(context.TODO(), hyperspell.MemoryUploadParams{
		File:       io.Reader(bytes.NewBuffer([]byte("Example data"))),
		Collection: hyperspell.String("collection"),
		Metadata:   hyperspell.String("metadata"),
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
