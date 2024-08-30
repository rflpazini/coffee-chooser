package recommendation

import (
	"net/http"
	"strings"

	"coffee-choose/pkg/service/coffeeTypes"
	opService "coffee-choose/pkg/service/openaiClient"
	"coffee-choose/pkg/service/preferences"
	"coffee-choose/pkg/service/recommendations"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"go.uber.org/dig"
)

type getParams struct {
	dig.In

	CoffeeVarietyService coffeeTypes.GetAllCoffeeVarieties
	OpenAi               opService.OpenAIService
	SaveRecommendation   recommendations.SaveRecommendation
}

func makeRecommendationGet(p getParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		r := c.Request()
		// Parse query parameters
		userPreferences := preferences.UserPreferences{
			Sweetness:   c.QueryParam("sweetness"),
			Strength:    c.QueryParam("strength"),
			FlavorNotes: c.QueryParam("flavor_notes"),
			Body:        c.QueryParam("body"),
		}

		// Validate required fields
		if userPreferences.Sweetness == "" || userPreferences.Strength == "" || len(userPreferences.FlavorNotes) == 0 || userPreferences.Body == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "All fields (sweetness, strength, flavor_notes, body) must be provided")
		}

		// Fetch all coffee varieties
		coffeeVarieties, err := p.CoffeeVarietyService(r.Context())
		if err != nil {
			log.Error().Err(err).Msg("Failed to retrieve coffee varieties")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve coffee varieties")
		}

		// Call the OpenAI service to get the best match
		recommendedVariety, err := p.OpenAi.GetCoffeeRecommendation(r.Context(), userPreferences, coffeeVarieties)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get coffee recommendation")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get coffee recommendation")
		}

		if recommendedVariety == "acai" {
			recommendedVariety = "acaia"
		}

		// Find the description for the recommended variety
		v := searchVariety(coffeeVarieties, recommendedVariety)

		rsp := &opService.CoffeeRecommendationsStruct{
			Recommendation: &opService.RecommendationStruct{
				Variety:     recommendedVariety,
				Description: v.Description,
				Vendors:     v.Vendors,
			},
		}

		_, err = p.SaveRecommendation(r.Context(), &recommendations.Recommendation{CoffeeVariety: rsp.Recommendation.Variety})
		if err != nil {
			log.Error().Err(err).Msg("Failed to save recommendation")
			return err
		}

		adRec, err := p.OpenAi.SuggestAdditionalVarieties(r.Context(), userPreferences, recommendedVariety, coffeeVarieties)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get additional coffee suggestions")
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get additional coffee suggestions")
		}

		var ar []*opService.RecommendationStruct
		for _, ad := range adRec {
			v := searchVariety(coffeeVarieties, strings.TrimSpace(ad))
			ar = append(ar, &opService.RecommendationStruct{
				Variety:     ad,
				Description: v.Description,
				Vendors:     v.Vendors,
			})
		}

		rsp.AdditionalRecommendations = ar
		return c.JSON(http.StatusOK, rsp)
	}
}

func searchVariety(coffeeVarieties []coffeeTypes.CoffeeVariety, recommendedVariety string) coffeeTypes.CoffeeVariety {
	var variety coffeeTypes.CoffeeVariety
	for _, v := range coffeeVarieties {
		if v.Variety == recommendedVariety {
			variety = v
			break
		}
	}

	return variety
}
