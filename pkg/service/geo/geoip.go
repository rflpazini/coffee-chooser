package geo

import (
	"context"
	"net"

	"github.com/oschwald/geoip2-golang"
	"github.com/rs/zerolog/log"
)

type IPService interface {
	GetLocation(ctx context.Context, ipAddress string) (*Location, error)
	Close() error
}

type geoIPServiceImpl struct {
	db *geoip2.Reader
}

func makeGeoIPService() (IPService, error) {
	db, err := geoip2.Open("./scripts/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal().Msgf("Could not open GeoIP database: %v", err)
		return nil, err
	}
	return &geoIPServiceImpl{db: db}, nil
}

func (g *geoIPServiceImpl) GetLocation(ctx context.Context, ipAddress string) (*Location, error) {
	ip := net.ParseIP(ipAddress)
	record, err := g.db.City(ip)

	if err != nil {
		log.Error().Err(err).Msgf("Failed to get location for IP %s", ipAddress)
		return nil, err
	}

	location := &Location{}
	if record.City.Names["en"] != "" {
		location.City = record.City.Names["en"]
	}

	location.Country = record.Country.IsoCode
	location.Latitude = record.Location.Latitude
	location.Longitude = record.Location.Longitude
	location.Timezone = record.Location.TimeZone

	return location, nil
}

func (g *geoIPServiceImpl) Close() error {
	return g.db.Close()
}
