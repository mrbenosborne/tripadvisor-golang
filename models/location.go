package models

// LocationResponse a response for a Location search.
type LocationResponse struct {
	Name               string            `json:"name,omitempty"`
	NumReviews         string            `json:"num_reviews,omitempty"`
	Category           *Category         `json:"category,omitempty"`
	SubCategory        *Category         `json:"sub_category,omitempty"`
	Address            *Address          `json:"address_obj,omitempty"`
	Latitude           string            `json:"latitude,omitempty"`
	Longitude          string            `json:"longitude,omitempty"`
	Rating             string            `json:"rating,omitempty"`
	LocationID         string            `json:"location_id,omitempty"`
	TripTypes          []*TripType       `json:"trip_types,omitempty"`
	Reviews            []*Review         `json:"reviews,omitempty"`
	WriteReviewURL     string            `json:"write_review,omitempty"`
	Ancestors          []*Ancestor       `json:"ancestors,omitempty"`
	PercentRecommended int               `json:"percent_recommended,omitempty"`
	ReviewRatingCount  map[string]string `json:"review_rating_count,omitempty"`
	PhotoCount         string            `json:"photo_count,omitempty"`
	LocationString     string            `json:"location_string,omitempty"`
	WebURL             string            `json:"web_url,omitempty"`
	PriceLevel         string            `json:"price_level,omitempty"`
	RatingImageURL     string            `json:"rating_image_url,omitempty"`
	Awards             []*Award          `json:"awards,omitempty"`
	SeeAllPhotos       string            `json:"see_all_photos,omitempty"`
}

// Address provides address details for the relevant location
type Address struct {
	Street1     string `json:"street1,omitempty"`
	Street2     string `json:"street2,omitempty"`
	City        string `json:"city,omitempty"`
	State       string `json:"state,omitempty"`
	Country     string `json:"country,omitempty"`
	PostalCode  string `json:"postalcode,omitempty"`
	FullAddress string `json:"address_string,omitempty"`
}

// TripType some information about a trip type
type TripType struct {
	Name          string `json:"name,omitempty"`
	Value         string `json:"value,omitempty"`
	LocalizedName string `json:"localized_name,omitempty"`
}

// User information about the user who wrote the review
type User struct {
	Username     string       `json:"username,omitempty"`
	UserLocation UserLocation `json:"user_location,omitempty"`
}

// UserLocation information about a user's location.
type UserLocation struct {
	Name string `json:"name,omitempty"`
	ID   string `json:"id,omitempty"`
}

// Ancestor .
type Ancestor struct {
	Abbrv      string `json:"abbrv,omitempty"`
	Level      string `json:"level,omitempty"`
	Name       string `json:"name,omitempty"`
	LocationID string `json:"location_id,omitempty"`
}

// Award information about an award
type Award struct {
	AwardType   string            `json:"award_type,omitempty"`
	Year        string            `json:"year,omitempty"`
	Images      map[string]string `json:"images,omitempty"`
	Categories  []string          `json:"categories,omitempty"`
	DisplayName string            `json:"display_name,omitempty"`
}

// Category information about a category
type Category struct {
	Name          string `json:"name,omitempty"`
	LocalizedName string `json:"localized_name,omitempty"`
}
