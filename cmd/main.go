package main

import (
	"context"
	"log"

	"github.com/mrbenosborne/tripadvisor-golang/pkg/tripadvisor"
)

func main() {
	tClient := tripadvisor.New(
		tripadvisor.SetKey("XXXXXXXXXXXXXXXXXXXXXX"),
		tripadvisor.SetLanguageCode("en_UK"),
	)
	response, err := tClient.Location(context.Background(), 3539289)
	if err != nil {
		panic(err)
	}

	log.Printf("%s - Total number of reviews: %s\n", response.Name, response.NumReviews)
	for _, review := range response.Reviews {
		log.Printf("Review:\n\t%s\n", review.Content)
	}
}
