package breed

import (
	"time"
)

// Breed struct
type Breed struct {
	UniqeName    string    `json:"uniqueName" tigris:"primaryKey:1,searchIndex" example:"affenpinscher"`
	Name         string    `json:"name" tigris:"searchIndex,sort" example:"Affenpinscher"`
	URL          string    `json:"url" example:"https://en.wikipedia.org/wiki/Affenpinscher"`
	CreationType string    `json:"creationType" tigris:"searchIndex" example:"original"`
	CreatedAt    time.Time `json:"createdAt" example:"2023-01-05T00:00:00.000Z"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2023-01-05T00:00:00.000Z"`
}

// CreateBreed struct
type CreateBreed struct {
	Name         string    `json:"name" validate:"required" example:"Affenpinscher"`
	UniqeName    string    `json:"uniqueName" validate:"required,alpha_underscore" example:"affenpinscher"`
	URL          string    `json:"url" validate:"required,url" example:"https://en.wikipedia.org/wiki/Affenpinscher"`
	CreationType string    `json:"creationType" validate:"required,oneof=original custom" example:"original"`
	CreatedAt    time.Time `json:"createdAt" example:"2023-01-05T00:00:00.000Z"`
	UpdatedAt    time.Time `json:"updatedAt" example:"2023-01-05T00:00:00.000Z"`
}

// UpdateBreed struct
type UpdateBreed struct {
	Name      string    `json:"name" validate:"omitempty" example:"Affenpinscher"`
	URL       string    `json:"url" validate:"omitempty,url" example:"https://en.wikipedia.org/wiki/Affenpinscher"`
	UpdatedAt time.Time `json:"updatedAt" example:"2023-01-05T00:00:00.000Z"`
}
