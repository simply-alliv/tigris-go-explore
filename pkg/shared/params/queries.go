package params

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PaginationQueryParams struct {
	Page     int  `validate:"omitempty,number"`
	Limit    int  `validate:"omitempty,number"`
	Paginate bool `validate:"omitempty,boolean"`
}

type CustomerQueryParams struct {
	CustomerID *primitive.ObjectID `validate:"omitempty,min=1"`
}

type BreedQueryParams struct {
	CreationType *string `json:"creationType" bson:"creationType" enums:"original,custom"  validate:"omitempty,oneof=original custom"`
}

type BookingQueryParams struct {
	BookingState       *string    `json:"bookingState" bson:"bookingState" enums:"created,approved,paid,started,completed,cancelled,deleted,rescheduled"  validate:"omitempty,oneof=created approved paid started completed cancelled deleted rescheduled"`
	FromDate           *time.Time `validate:"omitempty,min=1"`
	ToDate             *time.Time `validate:"omitempty,min=1"`
	Approved           *bool      `validate:"omitempty"`
	Paid               *bool      `validate:"omitempty"`
	BookedForOrder     *int       `validate:"omitempty"`
	BookedForTimeOrder *int       `validate:"omitempty"`
	CreatedAtOrder     *int       `validate:"omitempty"`
}

type PaystackQueryParams struct {
	IsTest bool `validate:"omitempty,boolean"`
}

type XeroQueryParams struct {
	XeroTenantId string `validate:"required,min=1,uuid"`
}

type SortQueryParams struct {
	Key   string
	Value interface{}
}
