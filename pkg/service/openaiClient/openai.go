package openaiClient

import (
	"context"
	"fmt"
	"strings"

	"coffee-choose/pkg/config"
	"coffee-choose/pkg/service/coffeeTypes"
	"coffee-choose/pkg/service/preferences"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/dig"
)

type OpenAIService interface {
	GetCoffeeRecommendation(ctx context.Context, userPreferences preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) (string, error)
}

type openAIServiceImpl struct {
	client *openai.Client
}

type OpenAIParams struct {
	dig.In

	*config.OpenAIConfig
}

func makeOpenAIService(p OpenAIParams) (OpenAIService, error) {
	client := openai.NewClient(p.OpenAIConfig.Key)
	return &openAIServiceImpl{client: client}, nil
}

func (s *openAIServiceImpl) GetCoffeeRecommendation(ctx context.Context, userPreferences preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) (string, error) {
	prompt := CreatePrompt(userPreferences, varieties)
	resp, err := s.client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:       openai.GPT3TextDavinci003,
		Prompt:      prompt,
		MaxTokens:   50,
		Temperature: 0.5,
	})

	if err != nil {
		return "", err
	}

	return ParseResponse(resp.Choices[0].Text), nil
}

func CreatePrompt(p preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) string {
	// Build the prompt
	prompt := fmt.Sprintf(
		"Given the following preferences: sweetness: %s, strength: %s, flavor notes: %s, and body: %s.\nThe available coffee varieties are:\n",
		p.Sweetness, p.Strength, p.FlavorNotes, p.Body,
	)

	for _, variety := range varieties {
		// Convert coffee variety flavor notes to a comma-separated string
		varietyFlavorNotes := strings.Join(variety.FlavorNotes, ", ")
		prompt += fmt.Sprintf(
			"- %s: sweetness: %s, strength: %s, flavor notes: %s, body: %s\n",
			variety.Variety, variety.Sweetness, variety.Strength, varietyFlavorNotes, variety.Body,
		)
	}

	prompt += "Please respond with only the variety name that best matches these preferences."

	return prompt
}

func ParseResponse(response string) string {
	// Simple parsing, assuming response is just the variety name.
	return strings.TrimSpace(response)
}
