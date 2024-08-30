//go:generate easyjson -lower_camel_case $GOFILE
package recommendations

import "go.mongodb.org/mongo-driver/bson/primitive"

//easyjson:json
type Recommendation struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserID        string             `bson:"user_id"`
	CoffeeVariety string             `bson:"coffee_variety"`
	CreatedAt     string             `bson:"created_at"`
}
