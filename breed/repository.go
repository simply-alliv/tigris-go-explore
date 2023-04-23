package breed

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/simply-alliv/tigris-go-explore/pkg/shared/pagination"
	"github.com/simply-alliv/tigris-go-explore/pkg/shared/params"
	"github.com/tigrisdata/tigris-client-go/fields"
	"github.com/tigrisdata/tigris-client-go/filter"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

// Repository is an implementation of the Service CRUD interface for organization's breeds.
//
// Repository is responsible for managing the persistence layer. (e.g. database operations)
type Repository interface {
	IService
}

type breedRepository struct {
	collection *tigris.Collection[Breed]
}

// NewBreedRepository returns a concrete implementation of the Repository interface.
func NewBreedRepository(collection *tigris.Collection[Breed]) Repository {
	return &breedRepository{collection: collection}
}

func (r breedRepository) GetAllBreeds(ctx context.Context, qp params.PaginationQueryParams, bqp params.BreedQueryParams) ([]Breed, *pagination.PaginationData, error) {
	var breeds []Breed = []Breed{}
	var f filter.Expr
	creationType := bqp.CreationType
	if creationType != nil {
		f = filter.Eq("creationType", *creationType)
	} else {
		// This is only added because if I do not initialise f, it automatically initialises
		// itself to an empty map (e.g. map[]).
		// TODO: Ask Tigris engineers if this is a bug because I would expect everything
		// to be returned when the map/filter is empty.
		f = filter.Or(filter.Eq("creationType", "original"), filter.Eq("creationType", "custom"))
	}
	it, err := r.collection.Read(ctx, f)
	var breed Breed
	for it.Next(&breed) {
		breeds = append(breeds, breed)
	}
	return breeds, &pagination.PaginationData{}, err
}

func (r breedRepository) GetSingleBreed(ctx context.Context, id string) (Breed, error) {
	breed, err := r.collection.ReadOne(ctx, filter.Eq("uniqueName", id))
	if err != nil {
		return *breed, err
	}
	return *breed, err
}

func (r breedRepository) CreateSingleBreed(ctx context.Context, dto CreateBreed) (Breed, error) {
	var breed Breed
	docs := &Breed{
		Name:         dto.Name,
		UniqeName:    dto.UniqeName,
		CreationType: dto.CreationType,
		URL:          dto.URL,
		CreatedAt:    dto.CreatedAt,
		UpdatedAt:    dto.UpdatedAt,
	}
	resp, err := r.collection.Insert(ctx, docs)
	if err != nil {
		return breed, err
	}
	if len(resp.Keys) == 1 {
		k := resp.Keys[0]
		data := map[string]interface{}{}
		err := json.Unmarshal(k, &data)
		fmt.Println(data)
		if err != nil {
			return breed, err
		}
	} else {
		return breed, fmt.Errorf("error counting length of response keys, only 1 is expected, %v received", len(resp.Keys))
	}
	return breed, nil
}

func (r breedRepository) UpdateSingleBreed(ctx context.Context, id string, dto UpdateBreed) (Breed, error) {
	update := fields.Update{}
	set := map[string]interface{}{}
	set["updatedAt"] = dto.UpdatedAt
	if dto.Name != "" {
		set["name"] = dto.Name
	}
	if dto.URL != "" {
		set["url"] = dto.URL
	}
	update.SetF = set
	_, err := r.collection.UpdateOne(ctx, filter.Eq("uniqueName", id), &update)
	if err != nil {
		return Breed{}, err
	}
	return r.GetSingleBreed(ctx, id)
}

func (r breedRepository) DeleteSingleBreed(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, filter.Eq("uniqueName", id))
	return err
}
