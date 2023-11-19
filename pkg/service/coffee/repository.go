package coffee

import (
	"context"
	"errors"

	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
)

type SaveBrewingMethod func(ctx context.Context, brewing BrewingRequest) error

func makeSaveBrewingMethod(coll database.BrewingCollection) SaveBrewingMethod {
	return func(ctx context.Context, req BrewingRequest) error {
		filter := bson.M{"name": req.Name}
		var result BrewingResponse

		err := coll.FindOne(ctx, filter).Decode(&result)
		if &result != nil {
			return errors.New("brewing method already exists: " + req.Name)
		}

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

type GetBrewingMethod func(ctx context.Context) ([]*BrewingResponse, error)

func makeGetBrewingMethod(coll database.BrewingCollection) GetBrewingMethod {
	return func(ctx context.Context) ([]*BrewingResponse, error) {
		res, err := coll.Find(ctx, bson.D{})
		if err != nil {
			return nil, err
		}

		defer res.Close(ctx)

		brewingResponse := make([]*BrewingResponse, 0)
		for res.Next(ctx) {
			var row *BrewingResponse
			err := res.Decode(&row)
			if err != nil {
				log.Error().Err(err).Msg("failed to decode the database response")
			}
			brewingResponse = append(brewingResponse, row)
		}

		return brewingResponse, nil
	}
}

type GetBrewingMethodByName func(ctx context.Context, name string) (*BrewingResponse, error)

func makeGetBrewingMethodByName(coll database.BrewingCollection) GetBrewingMethodByName {
	return func(ctx context.Context, name string) (*BrewingResponse, error) {
		filter := bson.M{"name": name}
		var result BrewingResponse

		err := coll.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Error().Err(err).Msgf("failed to find the brewing method with name: %s", name)
			return &BrewingResponse{}, err
		}

		return &result, nil
	}
}

type DeleteBrewingMethod func(ctx context.Context, name string) error

func makeDeleteBrewingMethod(coll database.BrewingCollection) DeleteBrewingMethod {
	return func(ctx context.Context, name string) error {
		methodName := bson.M{"name": name}

		_, err := coll.DeleteOne(ctx, methodName)
		if err != nil {
			log.Error().Err(err).Msgf("failed to delete: %s", name)
		}

		return err
	}
}
