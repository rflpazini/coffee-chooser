package preferences

import (
	"context"
	"fmt"
	"time"

	"coffee-choose/pkg/auth"
	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveUserPreferences func(ctx context.Context, p UserPreferences) (string, error)

func makeSaveUserPreferences(coll database.PreferencesCollection) SaveUserPreferences {
	return func(ctx context.Context, p UserPreferences) (string, error) {
		// Retrieve the session claims from the context
		session, ok := ctx.Value("session").(*auth.SessionToken)
		if !ok {
			log.Error().Msg("Failed to retrieve session from context")
			return "", fmt.Errorf("failed to retrieve session from context")
		}

		p.Location = *session.Geolocation
		p.UserID = session.UserID
		p.CreatedAt = time.Now().Format(time.RFC3339)

		// Insert the preferences into the database
		res, err := coll.InsertOne(ctx, p)
		if err != nil {
			log.Error().Err(err).Msg("Failed to insert into the database")
			return "", err
		}

		// Extract the inserted ID
		var insertedId string
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			insertedId = oid.Hex()
			log.Info().Msgf("Inserted %v into the database", insertedId)
		}

		return insertedId, nil
	}
}
