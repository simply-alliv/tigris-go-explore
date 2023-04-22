package breed

import (
	"context"
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
	}

	it, err := r.collection.Read(ctx, f)
	var breed Breed
	for it.Next(&breed) {
		breeds = append(breeds, breed)
	}
	return breeds, &pagination.PaginationData{}, err
}

func (r breedRepository) GetSingleBreed(ctx context.Context, id string) (Breed, error) {
	breed, err := r.collection.ReadOne(ctx, filter.Eq("id", id))
	if err != nil {
		return *breed, err
	}
	return *breed, err
}

func (r breedRepository) GetSingleBreedByUniqueName(ctx context.Context, uniqueName string) (Breed, error) {
	breed, err := r.collection.ReadOne(ctx, filter.Eq("uniqueName", uniqueName))
	if err != nil {
		return *breed, err
	}
	return *breed, err
}

func (r breedRepository) CreateSingleBreed(ctx context.Context, dto CreateBreed) (Breed, error) {
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
		return Breed{}, err
	}
	fmt.Println(resp)
	return Breed{}, nil
}

func (r breedRepository) UpdateSingleBreed(ctx context.Context, id string, dto UpdateBreed) (Breed, error) {
	var update *fields.Update
	if dto.Name != "" {
		update.Set("name", dto.Name)
	}
	if dto.URL != "" {
		update.Set("url", dto.URL)
	}
	resp, err := r.collection.UpdateOne(ctx, filter.Eq("id", id), update)
	if err != nil {
		return Breed{}, err
	}
	fmt.Println(resp)
	return Breed{}, err
}

func (r breedRepository) DeleteSingleBreed(ctx context.Context, id string) error {
	_, err := r.collection.DeleteOne(ctx, filter.Eq("id", id))
	return err
}
