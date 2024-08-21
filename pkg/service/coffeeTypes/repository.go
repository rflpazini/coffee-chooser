package coffeeTypes

import (
	"context"

	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type GetAllCoffeeVarieties func(ctx context.Context) ([]CoffeeVariety, error)

func makeGetAllCoffeeVarieties(coll database.CoffeeVarietyCollection) GetAllCoffeeVarieties {
	return func(ctx context.Context) ([]CoffeeVariety, error) {
		cur, err := coll.Find(ctx, bson.M{})
		if err != nil {
			log.Error().Err(err).Msg("Failed to retrieve coffee varieties")
			return nil, err
		}
		defer cur.Close(ctx)

		var coffeeVarieties []CoffeeVariety
		for cur.Next(ctx) {
			var coffeeVariety CoffeeVariety
			if err := cur.Decode(&coffeeVariety); err != nil {
				log.Error().Err(err).Msg("Failed to decode coffee variety")
				return nil, err
			}
			coffeeVarieties = append(coffeeVarieties, coffeeVariety)
		}

		if err := cur.Err(); err != nil {
			log.Error().Err(err).Msg("Cursor error while retrieving coffee varieties")
			return nil, err
		}

		return coffeeVarieties, nil
	}
}
