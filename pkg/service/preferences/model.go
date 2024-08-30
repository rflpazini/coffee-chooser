//go:generate easyjson -lower_camel_case $GOFILE
package preferences

import "coffee-choose/pkg/service/geo"

//easyjson:json
type UserPreferences struct {
	CreatedAt   string       `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt   string       `json:"updated_at,omitempty" bson:"updated_at"`
	Sweetness   string       `json:"sweetness" bson:"sweetness" validate:"required"`
	Strength    string       `json:"strength" bson:"strength" validate:"required"`
	FlavorNotes string       `json:"flavor_notes" bson:"flavor_notes" validate:"required"`
	Body        string       `json:"body" bson:"body" validate:"required"`
	IPAddress   string       `json:"ip_address,omitempty" bson:"ip_address"`
	Location    geo.Location `json:"location,omitempty" bson:"location"`
}

//easyjson:json
type SaveResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"err,omitempty"`
}
