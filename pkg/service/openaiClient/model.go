//go:generate easyjson -lower_camel_case $GOFILE
package openaiClient

//easyjson:json
type CoffeeRecommendationsStruct struct {
	Recommendation            *RecommendationStruct   `json:"recommendation"`
	AdditionalRecommendations []*RecommendationStruct `json:"additional_recommendations,omitempty"`
}

//easyjson:json
type RecommendationStruct struct {
	Variety     string `json:"variety"`
	Description string `json:"description"`
}
