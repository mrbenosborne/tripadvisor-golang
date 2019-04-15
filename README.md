# TripAdvisor Golang
A TripAdvisor API wrapper for Golang.

## Installation
```
go get "github.com/mrbenosborne/tripadvisor-golang"
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
