package openaiClient

import (
	"context"
	"fmt"
	"strings"

	"coffee-choose/pkg/config"
	"coffee-choose/pkg/service/coffeeTypes"
	"coffee-choose/pkg/service/preferences"
	"github.com/rs/zerolog/log"
	"github.com/sashabaranov/go-openai"
	"go.uber.org/dig"
)

type OpenAIService interface {
	GetCoffeeRecommendation(ctx context.Context, p preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) (string, error)
	SuggestAdditionalVarieties(ctx context.Context, p preferences.UserPreferences, initialRecommendation string, varieties []coffeeTypes.CoffeeVariety) ([]string, error)
}

type openAIServiceImpl struct {
	client *openai.Client
}

type OpenAIServiceParams struct {
	dig.In

	*config.OpenAIConfig
}

func makeOpenAIService(p OpenAIServiceParams) (OpenAIService, error) {
	client := openai.NewClient(p.OpenAIConfig.Key)
	if client == nil {
		log.Error().Msg("Failed to create OpenAI client")
		return nil, fmt.Errorf("failed to create OpenAI client")
	}
	return &openAIServiceImpl{client: client}, nil
}

func (s *openAIServiceImpl) GetCoffeeRecommendation(ctx context.Context, p preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) (string, error) {
	prompt := createPrompt(p, varieties)

	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to get chat completion from OpenAI")
		return "", err
	}

	return parseResponse(resp.Choices[0].Message.Content), nil
}

func (s *openAIServiceImpl) SuggestAdditionalVarieties(ctx context.Context, p preferences.UserPreferences, initialRecommendation string, varieties []coffeeTypes.CoffeeVariety) ([]string, error) {
	prompt := createSuggestionPrompt(p, initialRecommendation, varieties)

	resp, err := s.client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	})

	if err != nil {
		log.Error().Err(err).Msg("Failed to get chat completion for additional suggestions from OpenAI")
		return nil, err
	}

	return parseSuggestionResponse(resp.Choices[0].Message.Content), nil
}

func createPrompt(p preferences.UserPreferences, varieties []coffeeTypes.CoffeeVariety) string {
	prompt := fmt.Sprintf(
		"Given the following preferences: sweetness: %s, strength: %s, flavor notes: %s, and body: %s.\nThe available coffee varieties are:\n",
		p.Sweetness, p.Strength, p.FlavorNotes, p.Body,
	)

	for _, variety := range varieties {
		varietyFlavorNotes := strings.Join(variety.FlavorNotes, ", ")
		prompt += fmt.Sprintf(
			"- %s: sweetness: %s, strength: %s, flavor notes: %s, body: %s\n",
			variety.Variety, variety.Sweetness, variety.Strength, varietyFlavorNotes, variety.Body,
		)
	}

	prompt += "Please respond with only the variety name that best matches these preferences. i.e.: arara, catuai, bourbon, acaia."

	return prompt
}

func createSuggestionPrompt(p preferences.UserPreferences, initialRecommendation string, varieties []coffeeTypes.CoffeeVariety) string {
	prompt := fmt.Sprintf(
		"Given the coffee recommendation '%s' for the following preferences: sweetness: %s, strength: %s, flavor notes: %s, and body: %s.\nSuggest two additional coffee varieties from the following list that the user may like:\n",
		initialRecommendation, p.Sweetness, p.Strength, p.FlavorNotes, p.Body,
	)

	for _, variety := range varieties {
		if variety.Variety != initialRecommendation {
			varietyFlavorNotes := strings.Join(variety.FlavorNotes, ", ")
			prompt += fmt.Sprintf(
				"- %s: sweetness: %s, strength: %s, flavor notes: %s, body: %s\n",
				variety.Variety, variety.Sweetness, variety.Strength, varietyFlavorNotes, variety.Body,
			)
		}
	}

	prompt += "Please respond with the names of two varieties that the user may like, separated by a comma. i.e.: arara, bourbon"

	return prompt
}

func parseResponse(response string) string {
	return strings.TrimSpace(response)
}

func parseSuggestionResponse(response string) []string {
	return strings.Split(strings.TrimSpace(response), ",")
}
