package breed

import (
	"context"
	"time"

	"github.com/simply-alliv/tigris-go-explore/pkg/shared/pagination"
	"github.com/simply-alliv/tigris-go-explore/pkg/shared/params"
)

// IService is a simple CRUD interface for organization's breeds.
type IService interface {
	GetAllBreeds(ctx context.Context, qp params.PaginationQueryParams, bqp params.BreedQueryParams) ([]Breed, *pagination.PaginationData, error)
	CreateSingleBreed(ctx context.Context, dto CreateBreed) (Breed, error)
	GetSingleBreed(ctx context.Context, id string) (Breed, error)
	GetSingleBreedByUniqueName(ctx context.Context, uniqueName string) (Breed, error)
	UpdateSingleBreed(ctx context.Context, id string, dto UpdateBreed) (Breed, error)
	DeleteSingleBreed(ctx context.Context, id string) error
}

type Service struct {
	r Repository
}

// NewBreedService returns a service
func NewBreedService(r Repository) *Service {
	return &Service{r: r}
}

// GetAllBreeds godoc
// @Summary Get all breed resources
// @Description Get all the breed resources
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Success 200 {object} JSONResultSuccess{data=[]Breed} "OK"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds [get]
func (s *Service) GetAllBreeds(ctx context.Context, qp params.PaginationQueryParams, bqp params.BreedQueryParams) ([]Breed, *pagination.PaginationData, error) {
	data, metadata, err := s.r.GetAllBreeds(ctx, qp, bqp)
	if err != nil {
		// TODO: Handle error
		return data, metadata, err
	} else {
		return data, metadata, nil
	}
}

// CreateSingleBreed godoc
// @Summary Create single breed resource
// @Description Create a single breed resource
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Param body body CreateBreed true "JSON body to create a breed resource"
// @Success 201 {object} JSONResultSuccess{data=string} "Created"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 409 {object} JSONResultFailure "Error: Conflict"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds [post]
func (s *Service) CreateSingleBreed(ctx context.Context, dto CreateBreed) (Breed, error) {
	// Set default timestamps
	dto.CreatedAt = time.Now()
	dto.UpdatedAt = time.Now()

	// Create the breed record with all the defaults set
	data, err := s.r.CreateSingleBreed(ctx, dto)
	if err != nil {
		// TODO: Handle error
		return Breed{}, err
	} else {
		return data, nil
	}
}

// GetSingleBreed godoc
// @Summary Get single breed resource
// @Description Get a single breed resource
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Param id path string true "ID of the breed resource"
// @Success 200 {object} JSONResultSuccess{data=Breed} "OK"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 404 {object} JSONResultFailure "Error: Not Found"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds/{id} [get]
func (s *Service) GetSingleBreed(ctx context.Context, id string) (Breed, error) {
	data, err := s.r.GetSingleBreed(ctx, id)
	if err != nil {
		// TODO: Handle error
		return data, err
	} else {
		return data, nil
	}
}

// GetSingleBreedByUniqueName godoc
// @Summary Get single breed resource by its unique name
// @Description Get a single breed resource by its unique name
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Param uniqueName path string true "Unique name of the breed resource"
// @Success 200 {object} JSONResultSuccess{data=Breed} "OK"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 404 {object} JSONResultFailure "Error: Not Found"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds/uniquename/{uniqueName} [get]
func (s *Service) GetSingleBreedByUniqueName(ctx context.Context, uniqueName string) (Breed, error) {
	data, err := s.r.GetSingleBreedByUniqueName(ctx, uniqueName)
	if err != nil {
		// TODO: Handle error
		return data, err
	} else {
		return data, nil
	}
}

// UpdateSingleBreed godoc
// @Summary Update single breed
// @Description Update a single breed
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Param id path string true "ID of the breed"
// @Param body body UpdateBreed true "JSON body to update a breed"
// @Success 200 {object} JSONResultSuccess{} "OK"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 404 {object} JSONResultFailure "Error: Not Found"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds/{id} [patch]
func (s *Service) UpdateSingleBreed(ctx context.Context, id string, dto UpdateBreed) (Breed, error) {
	// Set default timestamps
	dto.UpdatedAt = time.Now()

	c, err := s.r.UpdateSingleBreed(ctx, id, dto)
	if err != nil {
		// TODO: Handle error
		return c, err
	} else {
		return c, nil
	}
}

// DeleteSingleBreed godoc
// @Summary Delete single breed
// @Description Delete a single breed
// @Security Bearer
// @Tags Breed
// @Accept json
// @Produce json
// @Param id path string true "ID of the breed"
// @Success 200 {object} JSONResultSuccess{} "OK"
// @Failure 401 {object} JSONResultFailure "Error: Unauthorized"
// @Failure 404 {object} JSONResultFailure "Error: Not Found"
// @Failure 422 {object} JSONResultFailure "Error: Unprocessable Entity"
// @Failure 500 {object} JSONResultFailure "Error: Internal Server Error"
// @Router /breeds/{id} [delete]
func (s *Service) DeleteSingleBreed(ctx context.Context, id string) error {
	err := s.r.DeleteSingleBreed(ctx, id)
	if err != nil {
		// TODO: Handle error
		return err
	} else {
		return nil
	}
}
