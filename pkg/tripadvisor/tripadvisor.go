package tripadvisor

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/mrbenosborne/tripadvisor-golang/pkg/tripadvisor/models"
)

var _ TripAdvisor = (*tripadvisor)(nil)

type TripAdvisor interface {
}

type (
	Location interface {
		Location(ctx context.Context, locationID int) (*models.LocationResponse, error)
	}
	LocationFunc func(ctx context.Context, locationID int) (*models.LocationResponse, error)
)

func (f LocationFunc) Location(ctx context.Context, locationID int) (*models.LocationResponse, error) {
	return f(ctx, locationID)
}

type tripadvisor struct {
	key          string
	languageCode string
	endpoint     string
	timeout      time.Duration
	client       *http.Client
}

type Option func(*tripadvisor)

// New Create a new tripadvisor instance.
//
// Default Endpoint is: https://api.tripadvisor.com/api/partner/2.0/", you can change this vie SetEndpoint()
func New(opts ...Option) *tripadvisor {
	t := tripadvisor{
		languageCode: "en_UK",
		endpoint:     "https://api.tripadvisor.com/api/partner/2.0/",
		timeout:      time.Second * 30,
	}
	for _, opt := range opts {
		opt(&t)
	}
	t.client = &http.Client{
		Timeout: t.timeout,
	}
	return &t
}

// SetKey Set your TripAdvisor key.
//
// Request a key if you do not have one at the TripAdvisor website
// found here: https://developer-tripadvisor.com/content-api/request-api-access/
func SetKey(key string) Option {
	return func(t *tripadvisor) {
		t.key = key
	}
}

// SetLanguageCode Set the language code, ie: en_UK.
// Default: en_UK
// A full list can be found here: https://developer-tripadvisor.com/content-api/supported-languages/
func SetLanguageCode(code string) Option {
	return func(t *tripadvisor) {
		t.languageCode = code
	}
}

// SetEndpoint Set the endpoint for the TripAdvisor API.
//
// An example and the default endpoint is: https://api.tripadvisor.com/api/partner/2.0/".
// Change at your own risk, newer versions or any version older than 2.0 is not currently
// supported.
func SetEndpoint(endpoint string) Option {
	return func(t *tripadvisor) {
		t.endpoint = endpoint
	}
}

// SetTimeout Set the timeout of HTTP requests made to the
// tripadvisor API.
func SetTimeout(timeout time.Duration) Option {
	return func(t *tripadvisor) {
		t.timeout = timeout
	}
}

// Location Search for a location based on a LocationID
func (t *tripadvisor) Location(ctx context.Context, locationID int) (*models.LocationResponse, error) {
	req, err := http.NewRequest(http.MethodGet, t.endpoint+"location/"+strconv.Itoa(locationID)+"?key="+t.key, nil)
	if err != nil {
		return nil, err
	}
	httpResponse, err := t.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	// decode response
	decoder := json.NewDecoder(httpResponse.Body)
	response := models.LocationResponse{}
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
