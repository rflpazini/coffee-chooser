//go:generate easyjson -lower_camel_case $GOFILE
package geo

//easyjson:json
type Location struct {
	Country   string  `json:"country" bson:"country"`
	City      string  `json:"city" bson:"city"`
	Timezone  string  `json:"timezone" bson:"timezone"`
	Latitude  float64 `json:"latitude" bson:"latitude"`
	Longitude float64 `json:"longitude" bson:"longitude"`
}
