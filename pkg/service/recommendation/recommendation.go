package recommendation

import (
	"context"

	"coffee-choose/pkg/service/coffeeTypes"
	"coffee-choose/pkg/service/openaiClient"
	"coffee-choose/pkg/service/preferences"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type Service interface {
	RecommendCoffee(ctx context.Context, userPreferences preferences.UserPreferences) (string, error)
}

type recommendationServiceImpl struct {
	openAIService      openaiClient.OpenAIService
	getCoffeeVarieties coffeeTypes.GetAllCoffeeVarieties
}

type ServiceParams struct {
	dig.In

	OpenAIService      openaiClient.OpenAIService
	GetCoffeeVarieties coffeeTypes.GetAllCoffeeVarieties
}

// Constructor function for the recommendation service
func makeRecommendationService(p ServiceParams) Service {
	return &recommendationServiceImpl{
		openAIService:      p.OpenAIService,
		getCoffeeVarieties: p.GetCoffeeVarieties,
	}
}

func (r *recommendationServiceImpl) RecommendCoffee(ctx context.Context, userPref preferences.UserPreferences) (string, error) {
	// Fetch all coffee varieties
	coffeeVarieties, err := r.getCoffeeVarieties(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to retrieve coffee varieties for recommendation")
		return "", err
	}

	// Get the recommendation from OpenAI service
	recommendation, err := r.openAIService.GetCoffeeRecommendation(ctx, userPref, coffeeVarieties)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get coffee recommendation from OpenAI service")
		return "", err
	}

	return recommendation, nil
}
