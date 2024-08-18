package mindsdb

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type PredictionRequest struct {
	Model string                 `json:"model"`
	Data  map[string]interface{} `json:"data"`
}

type PredictionResponse struct {
	Prediction []map[string]interface{} `json:"prediction"`
}

func GetCoffeeRecommendation(apiKey, url string, data map[string]interface{}) (map[string]interface{}, error) {
	payload := PredictionRequest{
		Model: "coffee_recommender",
		Data:  data,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	if apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+apiKey)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get prediction from MindsDB")
	}

	var predictionResponse PredictionResponse
	if err := json.NewDecoder(resp.Body).Decode(&predictionResponse); err != nil {
		return nil, err
	}

	if len(predictionResponse.Prediction) == 0 {
		return nil, errors.New("no prediction returned from MindsDB")
	}

	return predictionResponse.Prediction[0], nil
}
