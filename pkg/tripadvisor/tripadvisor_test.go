package tripadvisor

import (
	"context"
	"testing"
)

func TestLocation(t *testing.T) {
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
				SetKey("XXXXXXXXXXXXXXXXXXXXXX"),
				SetLanguageCode(tc.langCode),
			)
			_, err := tripAdvisor.Location(context.Background(), tc.locationID)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
