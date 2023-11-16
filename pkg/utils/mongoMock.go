package utils

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockCollection struct {
	mock.Mock
}

func (m *MockCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) BulkWrite(ctx context.Context, models []mongo.WriteModel, opts ...*options.BulkWriteOptions) (*mongo.BulkWriteResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockCollection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	var args mock.Arguments
	if len(opts) > 0 {
		args = m.Called(ctx, filter, opts[0])
	} else {
		args = m.Called(ctx, filter)
	}

	return args.Get(0).(*mongo.SingleResult)
}
