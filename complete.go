package complete

import (
	"context"

	"github.com/sashabaranov/go-openai"
)

type Completion struct {
	client *openai.Client
}

func NewCompletion(token, baseUrl string) *Completion {
	clientConfig := openai.DefaultConfig(token)
	clientConfig.BaseURL = baseUrl
	return &Completion{openai.NewClientWithConfig(clientConfig)}
}

func (c Completion) Completion(prompt, model string) string {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := c.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		panic(err)
	}

	return resp.Choices[0].Message.Content
}
