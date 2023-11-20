//go:generate easyjson -lower_camel_case $GOFILE

package coffee

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//easyjson:json
type BrewingRequest struct {
	Name        string    `json:"name" bson:"name,omitempty" example:"V60"`
	Description string    `json:"description"  bson:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt" bson:"createdAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
}

//easyjson:json
type BrewingResponse struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Name        string             `json:"name" bson:"name,omitempty"`
	Description string             `json:"description"  bson:"description,omitempty"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
}

//easyjson:json
type BrewingUpdateRoutineResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}
