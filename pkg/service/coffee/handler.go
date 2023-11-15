package coffee

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

const (
	// BrewingCollection is the name of the collection in the database
	BrewingCollection = "brewing"
)

type SaveBrewingMethod func(ctx context.Context, brewing BrewingRequest) error

type saveParams struct {
	dig.In

	*mongo.Database
}

func makeSaveBrewingMethod(p saveParams) SaveBrewingMethod {
	return func(ctx context.Context, req BrewingRequest) error {
		coll := p.Collection(BrewingCollection)

		brewingData := bson.M{
			"name":        req.Name,
			"description": req.Description,
			"updatedAt":   req.UpdatedAt,
		}

		res, err := coll.InsertOne(ctx, brewingData)
		if err != nil {
			log.Error().Err(err).Msg("Failed to insert into the database")
			return err
		}

		log.Info().Msgf("Inserted %v into the database", res.InsertedID)

		return nil
	}
}

type GetBrewingMethod func(ctx context.Context) ([]BrewingResponse, error)

type getParams struct {
	dig.In

	*mongo.Database
}

func makeGetBrewingMethod(p getParams) GetBrewingMethod {
	return func(ctx context.Context) ([]BrewingResponse, error) {
		coll := p.Collection(BrewingCollection)

		res, err := coll.Find(ctx, bson.D{})
		if err != nil {
			return nil, err
		}

		defer res.Close(ctx)

		brewingResponse := make([]BrewingResponse, 0)
		for res.Next(ctx) {
			var row BrewingResponse
			err := res.Decode(&row)
			if err != nil {
				log.Error().Err(err).Msg("Failed to decode the database response")
			}
			brewingResponse = append(brewingResponse, row)
		}

		return brewingResponse, nil
	}
}
