// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package hyperspell_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/hyperspell/hyperspell-go"
	"github.com/hyperspell/hyperspell-go/internal/testutil"
	"github.com/hyperspell/hyperspell-go/option"
)

func TestContextDocumentConfigUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.ContextDocuments.Config.Update(context.TODO(), hyperspell.ContextDocumentConfigUpdateParams{
		CompanyPrompts: map[string]string{
			"foo": "string",
		},
		DetectionPrompt: hyperspell.String("detection_prompt"),
		Domain:          hyperspell.String("domain"),
		PersonalPrompt:  hyperspell.String("personal_prompt"),
		SourceWeights: map[string]string{
			"foo": "string",
		},
		Structure: hyperspell.ContextDocumentConfigUpdateParamsStructure{
			Company: []hyperspell.ContextDocumentConfigUpdateParamsStructureCompany{{
				Filename:      "filename",
				Key:           "key",
				Prompt:        "prompt",
				SearchQueries: []string{"string"},
			}},
			Workstream: []hyperspell.ContextDocumentConfigUpdateParamsStructureWorkstream{{
				Filename:      "filename",
				Key:           "key",
				Prompt:        "prompt",
				SearchQueries: []string{"string"},
			}},
		},
		WorkstreamPrompts: map[string]string{
			"foo": "string",
		},
	})
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContextDocumentConfigGet(t *testing.T) {
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
	_, err := client.ContextDocuments.Config.Get(context.TODO())
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestContextDocumentConfigReset(t *testing.T) {
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
	_, err := client.ContextDocuments.Config.Reset(context.TODO())
	if err != nil {
		var apierr *hyperspell.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
