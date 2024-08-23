package coffeeTypes

import (
	"context"
	"errors"
	"time"

	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type GetAllCoffeeVarieties func(ctx context.Context) ([]CoffeeVariety, error)
type GetCoffeeVarietyByName func(ctx context.Context, name string) (*CoffeeVariety, error)
type PostCoffeeVariety func(ctx context.Context, input CoffeeVariety) (primitive.ObjectID, error)

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

func makeGetCoffeeVarietyByName(coll database.CoffeeVarietyCollection) GetCoffeeVarietyByName {
	return func(ctx context.Context, name string) (*CoffeeVariety, error) {
		filter := bson.M{"variety": name}
		var coffeeVariety CoffeeVariety
		err := coll.FindOne(ctx, filter).Decode(&coffeeVariety)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return nil, nil // Return nil if no document is found
			}
			log.Error().Err(err).Msg("Failed to find coffee variety by name")
			return nil, err
		}
		return &coffeeVariety, nil
	}
}

func makePostCoffeeVariety(coll database.CoffeeVarietyCollection) PostCoffeeVariety {
	return func(ctx context.Context, input CoffeeVariety) (primitive.ObjectID, error) {
		filter := bson.M{"variety": input.Variety}
		update := bson.M{
			"$set": bson.M{
				"sweetness":    input.Sweetness,
				"strength":     input.Strength,
				"flavor_notes": input.FlavorNotes,
				"body":         input.Body,
				"description":  input.Description,
				"link":         input.Link,
				"updated_at":   time.Now().Format(time.RFC3339),
			},
			"$setOnInsert": bson.M{
				"created_at": time.Now().Format(time.RFC3339),
			},
		}
		opts := options.Update().SetUpsert(true)

		result, err := coll.UpdateOne(ctx, filter, update, opts)
		if err != nil {
			return primitive.NilObjectID, err
		}

		if result.UpsertedID != nil {
			return result.UpsertedID.(primitive.ObjectID), nil
		}

		var updatedDoc CoffeeVariety
		err = coll.FindOne(ctx, filter).Decode(&updatedDoc)
		if err != nil {
			return primitive.NilObjectID, err
		}

		return updatedDoc.ID, nil
	}
}
