package recommendations

import (
	"context"
	"fmt"
	"time"

	"coffee-choose/pkg/auth"
	"coffee-choose/pkg/database"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SaveRecommendation func(ctx context.Context, r *Recommendation) (string, error)

func makeSaveRecommendation(coll database.RecommendationCollection) SaveRecommendation {
	return func(ctx context.Context, r *Recommendation) (string, error) {

		session, ok := ctx.Value("session").(*auth.SessionToken)
		if !ok {
			log.Error().Msg("Failed to retrieve session from context")
			return "", fmt.Errorf("failed to retrieve session from context")
		}

		r.UserID = session.UserID
		r.CreatedAt = time.Now().Format(time.RFC3339)

		res, err := coll.InsertOne(ctx, r)
		if err != nil {
			log.Error().Err(err).Msg("Failed to insert recommendation into the database")
			return "", err
		}

		var insertedId string
		if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
			insertedId = oid.Hex()
			log.Info().Msgf("Inserted recommendation %v into the database", insertedId)
		}

		return insertedId, nil
	}
}
