package util

import (
	"github.com/go-resty/resty/v2"
	"healthy/config"
	"healthy/constant"
)

type GPTRequest struct {
	Model         string
	SystemMessage string
	UserMessage   string
	ImageURL      string
	MaxTokens     int
}

func (g *GPTRequest) BuildRequestBody() map[string]interface{} {
	return map[string]interface{}{
		"model": g.Model,
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": g.SystemMessage,
			},
			{
				"role":    "user",
				"content": g.UserMessage,
			},
			{
				"role": "user",
				"content": []map[string]interface{}{
					{
						"type": "image_url",
						"image_url": map[string]string{
							"url": g.ImageURL,
						},
					},
				},
			},
		},
		"max_tokens": g.MaxTokens,
	}
}

func (g *GPTRequest) SendRequest() (string, error) {
	requestBody := g.BuildRequestBody()

	resp, err := resty.New().R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+config.AppConfig.GPTApiKey).
		SetBody(requestBody).
		Post(constant.ChatApiUrl)

	if err != nil {
		return "", err
	}

	return resp.String(), nil
}
