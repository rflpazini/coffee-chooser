//go:generate easyjson -lower_camel_case $GOFILE
package coffeeTypes

import "time"

//easyjson:json
type CoffeeVariety struct {
	Variety     string    `bson:"variety" json:"variety"`
	Sweetness   string    `bson:"sweetness" json:"sweetness"`
	Strength    string    `bson:"strength" json:"strength"`
	FlavorNotes []string  `bson:"flavor_notes" json:"flavor_notes"`
	Body        string    `bson:"body" json:"body"`
	Description string    `bson:"description" json:"description"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
}
