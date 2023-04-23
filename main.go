package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/simply-alliv/tigris-go-explore/breed"
	"github.com/simply-alliv/tigris-go-explore/seed"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	// Initialise and configure the Tigris SDK client.
	cfg := &tigris.Config{
		URL:          os.Getenv("TIGRIS_URL"),
		ClientID:     os.Getenv("TIGRIS_CLIENT_ID"),
		ClientSecret: os.Getenv("TIGRIS_CLIENT_SECRET"),
		Project:      os.Getenv("TIGRIS_PROJECT"),
	}
	client, err := tigris.NewClient(ctx, cfg)
	if err != nil {
		panic(err)
	}

	// Create or update the collections and their schemas
	db, err := client.OpenDatabase(ctx, &breed.Breed{})
	if err != nil {
		panic(err)
	}
	c := tigris.GetCollection[breed.Breed](db)

	// Seed data if required flags are defined and valid,
	// instead of starting the server.
	seedDataStr := os.Getenv("SEED_DATA")
	seedDataBreedsFile := os.Getenv("SEED_DATA_BREEDS_FILE")
	if seedDataStr != "" && seedDataBreedsFile != "" {
		seedData, err := strconv.ParseBool(seedDataStr)
		if err != nil {
			panic(err)
		} else if seedData {
			seed.SeedData(ctx, seedDataBreedsFile, c)
		}
	} else {
		// Initialise the servive
		r := breed.NewBreedRepository(c)
		s := breed.NewBreedService(r)

		// Create and start the server
		router := mux.NewRouter()

		router.HandleFunc("/breeds", breed.GetAllBreeds(s)).Methods("GET")
		router.HandleFunc("/breeds/{uniqueName}", breed.GetSingleBreedByUniqueName(s)).Methods("GET")
		router.HandleFunc("/breeds", breed.CreateSingleBreed(s)).Methods("POST")
		router.HandleFunc("/breeds/{uniqueName}", breed.UpdateSingleBreed(s)).Methods("PATCH")
		router.HandleFunc("/breeds/{uniqueName}", breed.DeleteeSingleBreed(s)).Methods("DELETE")

		log.Fatal(http.ListenAndServe(":8000", router))
	}
}
