package seed

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/simply-alliv/tigris-go-explore/breed"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func SeedData(ctx context.Context, breedsFile string, collection *tigris.Collection[breed.Breed]) {
	fmt.Println("Seeding of the data will now start")
	// Read the JSON file
	data, err := os.ReadFile(breedsFile)
	if err != nil {
		log.Fatal("Error reading file: ", err)
		return
	}

	// Unmarshal the JSON data into a slice of Breed structs
	var originalBreeds []breed.Breed
	err = json.Unmarshal(data, &originalBreeds)
	if err != nil {
		log.Fatal("Error unmarshalling JSON: ", err)
		return
	}

	// Set the createdAt and updatedAt fields to now
	now := time.Now().UTC()
	for i := range originalBreeds {
		originalBreeds[i].CreatedAt = now
		originalBreeds[i].UpdatedAt = now
	}

	// Create a pointer for each original breed and append it to the new slice
	var breeds []*breed.Breed
	for i := range originalBreeds {
		b := &originalBreeds[i]
		breeds = append(breeds, b)
	}

	// Insert all the data into the database's breeds collection
	resp, err := collection.Insert(ctx, breeds...)
	if err != nil {
		log.Fatal("Error inserting seed data for breeds collection: ", err)
	}

	fmt.Printf("Seeded breeds data from breeds.json: Response: %+v\n", resp)
}
