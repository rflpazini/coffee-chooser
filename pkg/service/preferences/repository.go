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
			log.Error().Err(err).Msgf("Failed to get location for IP %s", p.IPAddress)
		} else {
			log.Info().Msgf("Location for IP %s: %+v", p.IPAddress, uLocation)

			if uLocation.City.Names["en"] != "" {
				p.Location.City = uLocation.City.Names["en"]
			}

			p.Location.Country = uLocation.Country.IsoCode
			p.Location.Latitude = uLocation.Location.Latitude
			p.Location.Longitude = uLocation.Location.Longitude
			p.Location.Timezone = uLocation.Location.TimeZone
		}

		//filter := bson.M{
		//	"sweetness":    p.Sweetness,
		//	"strength":     p.Strength,
		//	"flavor_notes": p.FlavorNotes,
		//	"body":         p.Body,
		//}

		//var existing UserPreferences
		//err = coll.FindOne(ctx, filter).Decode(&existing)
		//if !errors.Is(err, mongo.ErrNoDocuments) {
		//	return "", errors.New("preferences already exist")
		//}

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
