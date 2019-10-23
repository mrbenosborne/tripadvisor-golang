package models

// ReviewResponse represents a response from the review endpoint.
type ReviewResponse struct {
	Data []*Review `json:"data,omitempty"`
}

// Review information about a single review
type Review struct {
	ID             string `json:"id,omitempty"`
	LanguageCode   string `json:"lang,omitempty"`
	LocationID     string `json:"location_id,omitempty"`
	PublishedDate  string `json:"published_date,omitempty"`
	Rating         int    `json:"rating,omitempty"`
	HelpfulVotes   string `json:"helpful_votes,omitempty"`
	RatingImageURL string `json:"rating_image_url,omitempty"`
	URL            string `json:"url,omitempty"`
	TripType       string `json:"trip_type,omitempty"`
	TravelDate     string `json:"travel_date,omitempty"`
	Text           string `json:"text,omitempty"`
	User           User  `json:"user,omitempty"`
	Title          string `json:"title,omitempty"`
}
