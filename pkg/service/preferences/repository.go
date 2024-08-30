package preferences

import (
	"context"
	"time"

	"coffee-choose/pkg/database"
	"coffee-choose/pkg/service/geo"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveUserPreferences func(ctx context.Context, p UserPreferences) (string, error)

func makeSaveUserPreferences(coll database.PreferencesCollection, geoService geo.IPService) SaveUserPreferences {
	return func(ctx context.Context, p UserPreferences) (string, error) {
		uLocation, err := geoService.GetLocation(ctx, p.IPAddress)
		if err != nil {
			log.Error().Err(err).Msg("failed to get location")
		}

		p.Location = *uLocation
		p.CreatedAt = time.Now().Format(time.RFC3339)

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
