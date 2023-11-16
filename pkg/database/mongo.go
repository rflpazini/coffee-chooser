package database

import (
	"context"
	"time"

	"coffee-choose/pkg/config"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/dig"
)

const (
	BrewingCollectionName = "brewing"
)

type MongoClient interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

type MongoDB interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

type mongoClientParams struct {
	dig.In

	*config.MongoConfig
}

func makeMongoClient(p mongoClientParams) (MongoClient, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(p.URI))
	if err != nil {
		panic(err)
	}

	return client, nil
}

type mongodbParams struct {
	dig.In

	MongoClient
	*config.MongoConfig
}

func makeMongoDB(p mongodbParams) (MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(p.Timeout)*time.Second)
	defer cancel()

	if err := p.Connect(ctx); err != nil {
		return nil, err
	}

	log.Ctx(ctx).Debug().Msgf("Pinging MongoDB database %s", p.Name)
	if err := p.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	return p.Database(p.Name), nil
}

type Ping func() error

func makeMongoPing(mongo *mongo.Database) Ping {
	return func() error {
		return mongo.Client().Ping(context.Background(), readpref.PrimaryPreferred())
	}
}

type Disconnect func() error

func makeMongoDisconnect(mongo *mongo.Database) Disconnect {
	return func() error {
		return mongo.Client().Disconnect(context.Background())
	}
}

type Decode interface {
	Decode(val interface{}) error
}

type Cursor interface {
	Next(ctx context.Context) bool
	Decode
	Err() error
	Close(ctx context.Context) error
}

type MongoCollection interface {
	Find(ctx context.Context, filter interface{},
		opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult
	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	BulkWrite(ctx context.Context, models []mongo.WriteModel,
		opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error)
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	DeleteOne(ctx context.Context, filter interface{},
		opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}

type BrewingCollection MongoCollection

func makeBrewingCollection(mongo MongoDB) BrewingCollection {
	return mongo.Collection(BrewingCollectionName)
}
