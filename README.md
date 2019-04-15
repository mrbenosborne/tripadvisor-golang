# TripAdvisor Golang
A TripAdvisor API wrapper for Golang.

## Installation
```
go get "github.com/mrbenosborne/tripadvisor-golang"
```

## Options

### SetKey
Set the TripAdvisor API key to use for all HTTP requests.

```go
SetKey(string)
```

### SetEndpoint
Set the TripAdvisor endpoint, only version 2 of the TripAdvisor API is currently supported so use at your own risk.

```go
SetEndpoint(string)
```

### SetLanguageCode
Set the language code for TripAdvisor responses, a full list of supported codes can be below:

[TripAdvisor API documentation page](https://developer-tripadvisor.com/content-api/supported-languages/).

```go
SetLanguageCode(string)
```

### SetTimeout
Set a timeout duration for all HTTP requests, the default is 30 seconds.

```go
SetTimeout(time.Duration)
```

## Usage
An example use of the library is below.

```go
package main

import (
	"context"
	"log"

	"github.com/mrbenosborne/tripadvisor-golang/pkg/tripadvisor"
)

func main() {

    // Create a new client
	tClient := tripadvisor.New(
		tripadvisor.SetKey("XXXXXXXXXXXXXXXXXXXXXX"),
		tripadvisor.SetLanguageCode("en_UK"),
    )

    // Get reviews for a location (The View from the Shard)
	response, err := tClient.Location(context.Background(), 3539289)
	if err != nil {
		panic(err)
	}

    // Print location data
	log.Printf("%s - Total number of reviews: %s\n", response.Name, response.NumReviews)
	for _, review := range response.Reviews {
		log.Printf("Review:\n\t%s\n", review.Content)
	}
}
```

## Example Response
The data below is an example output of the Location response in JSON format:

```json
{
  "name": "The View from The Shard",
  "num_reviews": "17349",
  "category": {
    "name": "attraction",
    "localized_name": "Attraction"
  },
  "address_obj": {
    "street1": "Joiner Street",
    "city": "London",
    "country": "United Kingdom",
    "postalcode": "SE1",
    "address_string": "Joiner Street, London SE1 England"
  },
  "latitude": "51.5045",
  "longitude": "-0.0865",
  "rating": "4.5",
  "location_id": "3539289",
  "trip_types": [
    {
      "name": "business",
      "value": "318",
      "localized_name": "Business"
    },
    {
      "name": "couples",
      "value": "7458",
      "localized_name": "Couples"
    },
    {
      "name": "solo",
      "value": "960",
      "localized_name": "Solo travel"
    },
    {
      "name": "family",
      "value": "3426",
      "localized_name": "Family"
    },
    {
      "name": "friends",
      "value": "3138",
      "localized_name": "Friends getaway"
    }
  ],
  "reviews": [
    {
      "id": "666623404",
      "lang": "en",
      "location_id": "3539289",
      "published_date": "2019-04-15T03:54:07-0400",
      "rating": 5,
      "helpful_votes": "1",
      "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/s5.0-20236-5.svg",
      "url": "https://www.tripadvisor.com/ShowUserReviews-g186338-d3539289-r666623404-The_View_from_The_Shard-London_England.html?m=20236#review666623404",
      "travel_date": "2019-03",
      "text": "Visiting The Shard (top level) is a must do in London, especially on not cloudy days.   Wonderful view all around (attractions even seem to be small from there).  No cheap (+- 30 GBP per person, but...",
      "user": {
        "username": "WimD94",
        "user_location": {}
      },
      "title": "Beautiful view on sunny days"
    },
    {
      "id": "666436956",
      "lang": "en",
      "location_id": "3539289",
      "published_date": "2019-04-14T11:13:08-0400",
      "rating": 5,
      "helpful_votes": "1",
      "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/s5.0-20236-5.svg",
      "url": "https://www.tripadvisor.com/ShowUserReviews-g186338-d3539289-r666436956-The_View_from_The_Shard-London_England.html?m=20236#review666436956",
      "trip_type": "Business",
      "travel_date": "2019-02",
      "text": "I visited IntechOpen Office at The Shard and they could not have chosen a better workplace.  The views from the top floor were spectacular.",
      "user": {
        "username": "MariaLorna1",
        "user_location": {
          "name": "Davis, California",
          "id": "32283"
        }
      },
      "title": "Beaautiful"
    },
    {
      "id": "666406491",
      "lang": "en",
      "location_id": "3539289",
      "published_date": "2019-04-14T09:31:24-0400",
      "rating": 5,
      "helpful_votes": "1",
      "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/s5.0-20236-5.svg",
      "url": "https://www.tripadvisor.com/ShowUserReviews-g186338-d3539289-r666406491-The_View_from_The_Shard-London_England.html?m=20236#review666406491",
      "trip_type": "Family",
      "travel_date": "2019-04",
      "text": "The partially open roof and turf flooring make this incredible view even better by giving a feeling of being outside. The Views are incredible of the entire city! Watching the sun sparkle on the...",
      "user": {
        "username": "WiseFun",
        "user_location": {
          "name": "Folsom, California",
          "id": "32389"
        }
      },
      "title": "Great place for a view of the city"
    },
    {
      "id": "666374293",
      "lang": "en",
      "location_id": "3539289",
      "published_date": "2019-04-14T07:55:39-0400",
      "rating": 5,
      "helpful_votes": "0",
      "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/s5.0-20236-5.svg",
      "url": "https://www.tripadvisor.com/ShowUserReviews-g186338-d3539289-r666374293-The_View_from_The_Shard-London_England.html?m=20236#review666374293",
      "trip_type": "Friends getaway",
      "travel_date": "2019-04",
      "text": "This is an excellent location from which to view central London. It is easy to access. On a clear day the view is worthwhile.",
      "user": {
        "username": "JennyNIreland",
        "user_location": {
          "name": "N Ireland"
        }
      },
      "title": "View from the Shard"
    },
    {
      "id": "666244282",
      "lang": "en",
      "location_id": "3539289",
      "published_date": "2019-04-13T23:26:20-0400",
      "rating": 5,
      "helpful_votes": "0",
      "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/s5.0-20236-5.svg",
      "url": "https://www.tripadvisor.com/ShowUserReviews-g186338-d3539289-r666244282-The_View_from_The_Shard-London_England.html?m=20236#review666244282",
      "travel_date": "2019-01",
      "text": "One of the main attractions in London that you can't miss if you are there for the first time. There are variety of restaurants and bars. You can book online your visit.",
      "user": {
        "username": "amr_farag2000",
        "user_location": {
          "name": "Dubai, United Arab Emirates",
          "id": "295424"
        }
      },
      "title": "Amazing skyline view of London"
    }
  ],
  "write_review": "https://www.tripadvisor.com/UserReview-g186338-d3539289-m20236-The_View_from_The_Shard-London_England.html",
  "ancestors": [
    {
      "level": "City",
      "name": "London",
      "location_id": "186338"
    },
    {
      "level": "Nation",
      "name": "England",
      "location_id": "186217"
    },
    {
      "abbrv": "UK",
      "level": "Country",
      "name": "United Kingdom",
      "location_id": "186216"
    }
  ],
  "percent_recommended": 82,
  "review_rating_count": {
    "1": "318",
    "2": "477",
    "3": "1802",
    "4": "4798",
    "5": "9680"
  },
  "photo_count": "11780",
  "location_string": "London, England",
  "web_url": "https://www.tripadvisor.com/Attraction_Review-g186338-d3539289-Reviews-The_View_from_The_Shard-London_England.html?m=20236",
  "rating_image_url": "https://www.tripadvisor.com/img/cdsi/img2/ratings/traveler/4.5-20236-5.svg",
  "awards": [
    {
      "award_type": "Certificate of Excellence",
      "year": "2018",
      "images": {
        "large": "https://www.tripadvisor.com/img/cdsi/img2/awards/CERTIFICATE_OF_EXCELLENCE_2018_en_US_large-20236-5.jpg",
        "small": "https://www.tripadvisor.com/img/cdsi/img2/awards/CERTIFICATE_OF_EXCELLENCE_v2_small-20236-5.jpg"
      },
      "display_name": "Certificate of Excellence 2018"
    }
  ],
  "see_all_photos": "https://www.tripadvisor.com/Attraction_Review-g186338-d3539289-m20236-Reviews-The_View_from_The_Shard-London_England.html#photos"
}
```
