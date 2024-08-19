package geo

import (
	"context"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

type IPService interface {
	GetLocation(ctx context.Context, ipAddress string) (*geoip2.City, error)
	Close() error
}

type geoIPServiceImpl struct {
	db *geoip2.Reader
}

func makeGeoIPService() (IPService, error) {
	db, err := geoip2.Open("./scripts/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalf("Could not open GeoIP database: %v", err)
		return nil, err
	}
	return &geoIPServiceImpl{db: db}, nil
}

func (g *geoIPServiceImpl) GetLocation(ctx context.Context, ipAddress string) (*geoip2.City, error) {
	ip := net.ParseIP(ipAddress)
	record, err := g.db.City(ip)
	if err != nil {
		return nil, err
	}
	return record, nil
}

func (g *geoIPServiceImpl) Close() error {
	return g.db.Close()
}
