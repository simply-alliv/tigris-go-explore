package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/simply-alliv/tigris-go-explore/breed"
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

	// Create or update the collections' schemas
	db, err := client.OpenDatabase(ctx, &breed.Breed{})
	if err != nil {
		panic(err)
	}

	// Initialise the servive
	c := tigris.GetCollection[breed.Breed](db)
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
