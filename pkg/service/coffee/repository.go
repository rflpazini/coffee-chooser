package coffee

import (
	"context"
	"errors"

	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SaveBrewingMethod func(ctx context.Context, brewing BrewingRequest) (string, error)

func makeSaveBrewingMethod(coll database.BrewingCollection) SaveBrewingMethod {
	return func(ctx context.Context, req BrewingRequest) (string, error) {
		filter := bson.M{"name": bson.M{"$eq": req.Name}}
		var result BrewingResponse

		err := coll.FindOne(ctx, filter).Decode(&result)
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("brewing method already exists: " + req.Name)
		}

		brewingData := bson.M{
			"name":        req.Name,
			"description": req.Description,
			"createdAt":   req.CreatedAt,
			"updatedAt":   req.UpdatedAt,
		}

		res, err := coll.InsertOne(ctx, brewingData)
		if err != nil {
			log.Error().Err(err).Msg("Failed to insert into the database")
			return "", err
		}

		var insertedId string
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			insertedId = oid.Hex()
			log.Info().Msgf("Inserted %v into the database", insertedId)
		}

		return insertedId, nil
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
		filter := bson.M{"name": bson.M{"$eq": name}}
		var result BrewingResponse

		err := coll.FindOne(ctx, filter).Decode(&result)
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Error().Err(err).Msgf("failed to find the brewing method with name: %s", name)
			return &BrewingResponse{}, err
		}

		return &result, nil
	}
}

type UpdateBrewingMethod func(ctx context.Context, brewing BrewingRequest, id string) error

func makeUpdateBrewingMethod(coll database.BrewingCollection) UpdateBrewingMethod {
	return func(ctx context.Context, req BrewingRequest, id string) error {
		objID, _ := primitive.ObjectIDFromHex(id)
		filter := bson.M{"_id": bson.M{"$eq": objID}}

		update := bson.D{{"$set", req}}

		_, err := coll.UpdateOne(ctx, filter, update)
		if err != nil {
			log.Error().Err(err).Msg("Failed to update the database")
			return err
		}

		return nil
	}
}

type DeleteBrewingMethod func(ctx context.Context, id string) error

func makeDeleteBrewingMethod(coll database.BrewingCollection) DeleteBrewingMethod {
	return func(ctx context.Context, id string) error {
		objID, _ := primitive.ObjectIDFromHex(id)
		methodId := bson.M{"_id": bson.M{"$eq": objID}}

		_, err := coll.DeleteOne(ctx, methodId)
		if err != nil {
			log.Error().Err(err).Msgf("failed to delete: %v", id)
		}

		return err
	}
}
