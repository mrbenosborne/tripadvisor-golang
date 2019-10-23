package tripadvisor

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/mercury-holidays/tripadvisor-golang/models"
)

var _ TripAdvisorAPI = (*TripAdvisor)(nil)

// TripAdvisorAPI is an interface for the TripAdvisor API.
type TripAdvisorAPI interface {
	Location
}

type (
	// Location An interface to access location functions
	Location interface {
		Location(ctx context.Context, locationID int) (*models.LocationResponse, error)
		Reviews(ctx context.Context, locationID int) (*models.ReviewResponse, error)
	}

	LocationFunc func(ctx context.Context, locationID int) (*models.LocationResponse, error)
	ReviewsFunc  func(ctx context.Context, locationID int) (*models.ReviewResponse, error)
)

// Location Search for a location based on a LocationID
func (f LocationFunc) Location(ctx context.Context, locationID int) (*models.LocationResponse, error) {
	return f(ctx, locationID)
}

// Reviews for a location based on a LocationID
func (r ReviewsFunc) Reviews(ctx context.Context, locationID int) (*models.ReviewResponse, error) {
	return r(ctx, locationID)
}

type TripAdvisor struct {
	key          string
	languageCode string
	endpoint     string
	timeout      time.Duration
	client       *http.Client
}

// Option Set an option such as your TripAdvisor key.
type Option func(*TripAdvisor)

// New Create a new tripadvisor instance.
//
// Default Endpoint is: https://api.tripadvisor.com/api/partner/2.0/", you can change this vie SetEndpoint()
func New(opts ...Option) *TripAdvisor {
	t := TripAdvisor{
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
	return func(t *TripAdvisor) {
		t.key = key
	}
}

// SetLanguageCode Set the language code, ie: en_UK.
// Default: en_UK
// A full list can be found here: https://developer-tripadvisor.com/content-api/supported-languages/
func SetLanguageCode(code string) Option {
	return func(t *TripAdvisor) {
		t.languageCode = code
	}
}

// SetEndpoint Set the endpoint for the TripAdvisor API.
//
// An example and the default endpoint is: https://api.tripadvisor.com/api/partner/2.0/".
// Change at your own risk, newer versions or any version older than 2.0 is not currently
// supported.
func SetEndpoint(endpoint string) Option {
	return func(t *TripAdvisor) {
		t.endpoint = endpoint
	}
}

// SetTimeout Set the timeout of HTTP requests made to the
// tripadvisor API.
func SetTimeout(timeout time.Duration) Option {
	return func(t *TripAdvisor) {
		t.timeout = timeout
	}
}

// Location Search for a location based on a LocationID
func (t *TripAdvisor) Location(ctx context.Context, locationID int) (*models.LocationResponse, error) {
	var resp models.LocationResponse
	url := t.endpoint + "location/" + strconv.Itoa(locationID) + "?key=" + t.key
	if err := t.callAPI(ctx, url, locationID, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Reviews returns the 'Reviews' for the given location.
func (t *TripAdvisor) Reviews(ctx context.Context, locationID int) (*models.ReviewResponse, error) {
	var resp models.ReviewResponse
	url := t.endpoint + "location/" + strconv.Itoa(locationID) + "/reviews?key=" + t.key
	if err := t.callAPI(ctx, url, locationID, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ErrorResponse represents an error returned by the API, i.e. a non-HTTP 200.
type ErrorResponse struct {
	ErrorType struct {
		Code    string `json:"code"`
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
	Response *http.Response `json:"-"` // HTTP response that caused this error
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("code: %v, type: %v, message: %v,  http status code: %v", r.ErrorType.Code, r.ErrorType.Type, r.ErrorType.Message, r.Response.StatusCode)
}

func (t *TripAdvisor) callAPI(ctx context.Context, url string, locationID int, resp interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	httpResponse, err := t.client.Do(req)
	if err != nil {
		return err
	}
	defer httpResponse.Body.Close()

	if c := httpResponse.StatusCode; c < 200 || c > 299 {
		errorResponse := &ErrorResponse{Response: httpResponse}
		data, err := ioutil.ReadAll(httpResponse.Body)
		if err == nil && data != nil {
			json.Unmarshal(data, errorResponse)
		}
		return errorResponse
	}

	// decode response
	decoder := json.NewDecoder(httpResponse.Body)
	err = decoder.Decode(&resp)
	if err != nil {
		return err
	}
	return nil
}
