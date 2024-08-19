package recommend

import (
	"context"
	"errors"
	"time"

	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type SaveUserPreferences func(ctx context.Context, p UserPreferences) (string, error)

func makeSaveUserPreferences(coll database.PreferencesCollection) SaveUserPreferences {
	return func(ctx context.Context, p UserPreferences) (string, error) {
		filter := bson.M{
			"sweetness":    p.Sweetness,
			"strength":     p.Strength,
			"flavor_notes": p.FlavorNotes,
			"body":         p.Body,
		}

		var existing UserPreferences
		err := coll.FindOne(ctx, filter).Decode(&existing)
		if !errors.Is(err, mongo.ErrNoDocuments) {
			return "", errors.New("preferences already exist")
		}

		p.CreatedAt = time.Now()

		res, err := coll.InsertOne(ctx, p)
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
