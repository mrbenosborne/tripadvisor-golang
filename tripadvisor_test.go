package tripadvisor

import (
	"context"
	"os"
	"testing"
)

func TestLocation_valid(t *testing.T) {
	apiKey := os.Getenv("TRIPADVISOR_APIKEY")
	if apiKey == "" {
		t.Skipf("No TRIPADVISOR_APIKEY set")
	}
	t.Parallel()
	for _, tc := range []struct {
		name       string
		locationID int
		langCode   string
	}{
		{
			name:       "Test London, UK",
			locationID: 186338,
			langCode:   "en_UK",
		},
		{
			name:       "The Rubens at the Palace",
			locationID: 199868,
			langCode:   "ar",
		},
		{
			name:       "The View from the Shard",
			locationID: 3539289,
			langCode:   "ko",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Create a new instance
			tripAdvisor := New(
				SetKey(apiKey),
				SetLanguageCode(tc.langCode),
			)
			_, err := tripAdvisor.Location(context.Background(), tc.locationID)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestLocation_invalidKey(t *testing.T) {
	tripAdvisor := New(
		SetKey("invalid_key"),
		SetLanguageCode("en_UK"),
	)
	_, err := tripAdvisor.Location(context.Background(), 3539289)
	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err, ok := err.(*ErrorResponse); !ok {
		t.Errorf("Expected an ErrorResponse error; got %#v.", err)
	}
}

func TestReviews_valid(t *testing.T) {
	apiKey := os.Getenv("TRIPADVISOR_APIKEY")
	if apiKey == "" {
		t.Skipf("No TRIPADVISOR_APIKEY set")
	}
	t.Parallel()
	for _, tc := range []struct {
		name       string
		locationID int
		langCode   string
	}{
		{
			name:       "Valid Hotel",
			locationID: 302546,
			langCode:   "en_UK",
		},
	} {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			tripAdvisor := New(
				SetKey(apiKey),
				SetLanguageCode(tc.langCode),
			)
			got, err := tripAdvisor.Reviews(context.Background(), tc.locationID)
			if err != nil {
				t.Fatal(err)
			}
			if len(got.Data) != 5 {
				t.Errorf("Expected 5 reviews; got %#v.", got)
			}
		})
	}
}
