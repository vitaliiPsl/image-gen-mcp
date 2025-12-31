package gemini

import (
	"context"
	"encoding/base64"
	"fmt"

	"google.golang.org/genai"

	"github.com/vitaliipsl/image-gen-mcp/internal/dto"
)

const (
	modelGeminiNanoBanana = "gemini-2.5-flash-image"
)

type Client struct {
	client *genai.Client
}

func NewClient(ctx context.Context, apiKey string) (*Client, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &Client{client: client}, nil
}

func (c *Client) GenerateImage(ctx context.Context, params *dto.GenerateImageParams) (*dto.GenerateImageResult, error) {
	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"IMAGE", "TEXT"},
		ImageConfig: &genai.ImageConfig{
			AspectRatio: params.AspectRatio,
		},
	}

	result, err := c.client.Models.GenerateContent(
		ctx,
		modelGeminiNanoBanana,
		genai.Text(params.Prompt),
		config,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate content: %w", err)
	}

	genResult := &dto.GenerateImageResult{}
	for _, candidate := range result.Candidates {
		if candidate.Content == nil {
			continue
		}
		for _, part := range candidate.Content.Parts {
			if part.InlineData != nil {
				base64Data := base64.StdEncoding.EncodeToString(part.InlineData.Data)
				genResult.Images = append(genResult.Images, dto.GeneratedImage{
					Data:     base64Data,
					MIMEType: part.InlineData.MIMEType,
				})
			}
		}
	}

	return genResult, nil
}
