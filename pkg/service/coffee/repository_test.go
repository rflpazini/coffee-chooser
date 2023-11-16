package coffee

import (
	"context"
	"errors"
	"testing"
	"time"

	"coffee-choose/pkg/utils"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func TestBrewingRepository_GetByName(t *testing.T) {
	type input struct {
		name      string
		input     string
		result    bson.M
		resultErr error
		want      BrewingResponse
		wantErr   error
	}

	date, err := time.Parse(time.Now().String(), "2023-11-16T15:35:43.511Z")
	if err != nil {
		return
	}

	tests := []input{
		{
			name:  "Should return a success response when search for a valid name",
			input: "test-name",
			result: bson.M{
				"id":          "1",
				"name":        "test-name",
				"description": "test-description",
				"updatedAt":   "2023-11-16T15:35:43.511Z",
			},
			resultErr: nil,
			want: BrewingResponse{
				ID:          "1",
				Name:        "test-name",
				Description: "test-description",
				UpdatedAt:   date,
			},
			wantErr: nil,
		},
		{
			name:   "Should return an error response when search for an invalid name",
			input:  "invalid-name",
			result: bson.M{},
			resultErr: &mongo.CommandError{
				Code:    1,
				Message: "failed to find the brewing method with name: invalid-name",
			},
			want:    BrewingResponse{},
			wantErr: errors.New("failed to find the brewing method with name: invalid-name"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			collection := new(utils.MockCollection)
			result := mongo.NewSingleResultFromDocument(tt.result, tt.resultErr, nil)
			filter := bson.M{"name": tt.input}
			collection.On("FindOne", ctx, filter).Return(result).Once()

			target := makeGetBrewingMethodByName(collection)
			actualResult, actualErr := target(ctx, tt.input)

			if tt.wantErr != nil {
				assert.Error(t, actualErr)
				assert.Equal(t, tt.wantErr.Error(), actualErr.Error())
				return
			}

			assert.NoError(t, actualErr)
			assert.Equal(t, tt.want, *actualResult)
		})
	}

}
