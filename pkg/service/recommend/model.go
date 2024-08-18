//go:generate easyjson -lower_camel_case $GOFILE
package recommend

import "time"

//easyjson:json
type UserPreferences struct {
	Sweetness         string    `json:"sweetness" bson:"sweetness"`
	Strength          string    `json:"strength" bson:"strength"`
	FlavorNotes       string    `json:"flavor_notes" bson:"flavor_notes"`
	Body              string    `json:"body" bson:"body"`
	RecommendedCoffee string    `json:"recommended_coffee,omitempty" bson:"recommended_coffee"`
	RecommendedBeans  string    `json:"recommended_beans,omitempty" bson:"recommended_beans"`
	ExtractionMethod  string    `json:"extraction_method,omitempty" bson:"extraction_method"`
	Justification     string    `json:"justification,omitempty" bson:"justification"`
	CreatedAt         time.Time `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" bson:"updated_at"`
}

//easyjson:json
type SaveResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"err,omitempty"`
}
