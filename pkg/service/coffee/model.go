//go:generate easyjson -lower_camel_case $GOFILE

package coffee

import "time"

//easyjson:json
type BrewingRequest struct {
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description"  bson:"description,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
}

//easyjson:json
type BrewingResponse struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Description string    `json:"description"  bson:"description,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt" bson:"updatedAt,omitempty" example:"2020-07-12T18:17:43.511Z"`
}
