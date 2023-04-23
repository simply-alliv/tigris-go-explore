package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	seedData, err := strconv.ParseBool(seedDataStr)
	if err != nil {
		log.Fatal("Unable to parse SEED_DATA string to boolean: ", err)
	}

	if seedDataStr != "" && seedDataBreedsFile != "" && seedData {
		seed.SeedData(ctx, seedDataBreedsFile, c)
		os.Exit(0)
	} else {
		// Initialise the servive
		r := breed.NewBreedRepository(c)
		s := breed.NewBreedService(r)

		// Create the routes
		router := mux.NewRouter()

		router.HandleFunc("/breeds", breed.GetAllBreeds(s)).Methods("GET")
		router.HandleFunc("/breeds/{id}", breed.GetSingleBreed(s)).Methods("GET")
		router.HandleFunc("/breeds", breed.CreateSingleBreed(s)).Methods("POST")
		router.HandleFunc("/breeds/{id}", breed.UpdateSingleBreed(s)).Methods("PATCH")
		router.HandleFunc("/breeds/{id}", breed.DeleteeSingleBreed(s)).Methods("DELETE")

		// Start the server
		port := os.Getenv("PORT")
		if port == "" {
			port = "8000"
		}
		server := &http.Server{Addr: ":" + port, Handler: router}

		go func() {
			if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("Could not listen on port %s: %v\n", port, err)
			}
		}()

		fmt.Printf("Listening on port %s...\n", port)

		// Wait for an interrupt signal to gracefully shutdown the server
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt)
		<-stop

		fmt.Println("Shutting down the server...")
		if err := server.Shutdown(context.Background()); err != nil {
			log.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		fmt.Println("Server stopped")
	}
}
