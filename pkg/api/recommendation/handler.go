package recommendation

import (
	"net/http"

	"coffee-choose/pkg/service/coffeeTypes"
	"coffee-choose/pkg/service/preferences"
	"coffee-choose/pkg/service/recommendation"
	"github.com/labstack/echo/v4"
	"go.uber.org/dig"
)

type makeGetParams struct {
	dig.In

	//RecommendationService recommendation.Service
	CoffeeVarietyService coffeeTypes.GetAllCoffeeVarieties
}

func makeRecommendationGet(p makeGetParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		//r := c.Request()

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
		//coffeeVarieties, err := p.CoffeeVarietyService(r.Context())
		//if err != nil {
		//	log.Error().Err(err).Msg("Failed to retrieve coffee varieties")
		//	return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve coffee varieties")
		//}

		// Call the recommendation service to get the best match
		//_, err = p.RecommendationService.RecommendCoffee(r.Context(), userPreferences)
		//if err != nil {
		//	log.Error().Err(err).Msg("Failed to get coffee recommendation")
		//	return echo.NewHTTPError(http.StatusInternalServerError, "Failed to get coffee recommendation")
		//}
		//
		//// Find the description for the recommended variety
		//var description string
		//for _, variety := range coffeeVarieties {
		//	if variety.Variety == recommendedVariety {
		//		description = variety.Description
		//		break
		//	}
		//}

		// Prepare the response
		response := &recommendation.Response{
			Variety:     "arara",
			Description: "arara is a brazilian coffee variety known for its vibrant fruity and floral profile, with a balanced sweetness and a full-bodied taste.",
		}

		// Return the response as JSON
		return c.JSON(http.StatusOK, response)
	}
}
