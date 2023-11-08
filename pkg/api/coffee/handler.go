package coffee

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

type makeTestParams struct {
	dig.In

	*mongo.Database
}

func makeTestHandler(p makeTestParams) echo.HandlerFunc {
	return func(c echo.Context) error {
		coll := p.Collection("users")
		user := bson.D{{"fullName", "User 1"}, {"age", 30}}

		result, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return c.JSON(http.StatusInternalServerError, err)
		}

		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		return c.JSON(http.StatusOK, result)
	}
}
