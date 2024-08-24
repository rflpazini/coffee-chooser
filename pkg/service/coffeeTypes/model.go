//go:generate easyjson -lower_camel_case $GOFILE
package coffeeTypes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//easyjson:json
type CoffeeVariety struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Variety     string             `bson:"variety" json:"variety"`
	Sweetness   string             `bson:"sweetness" json:"sweetness"`
	Strength    string             `bson:"strength" json:"strength"`
	Body        string             `bson:"body" json:"body"`
	FlavorNotes []string           `bson:"flavor_notes" json:"flavor_notes"`
	Vendors     []Vendor           `bson:"vendors,omitempty" json:"vendors,omitempty"`
	Description Description        `bson:"description,omitempty" json:"description,omitempty"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type Description struct {
	EN string `bson:"en" json:"en"`
	PT string `bson:"pt" json:"pt"`
}

type Vendor struct {
	URL   string `bson:"url" json:"URL"`
	Brand string `bson:"brand" json:"brand"`
}
