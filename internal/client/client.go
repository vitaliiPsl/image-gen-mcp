package client

import (
	"context"
	"fmt"

	"google.golang.org/genai"

	"github.com/vitaliipsl/image-gen-mcp/internal/dto"
)

const (
	modelGeminiNanoBanana = "gemini-2.5-flash-image"
)

type Client interface {
	GenerateImage(ctx context.Context, input *dto.GenerateImageInput) (*dto.GenerateImageResult, error)
}

type GeminiClient struct {
	client *genai.Client
}

func NewGeminiClient(ctx context.Context, apiKey string) (*GeminiClient, error) {
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create Gemini client: %w", err)
	}

	return &GeminiClient{client: client}, nil
}

func (c *GeminiClient) GenerateImage(ctx context.Context, input *dto.GenerateImageInput) (*dto.GenerateImageResult, error) {
	config := &genai.GenerateContentConfig{
		ResponseModalities: []string{"IMAGE", "TEXT"},
		ImageConfig: &genai.ImageConfig{
			AspectRatio: input.AspectRatio,
		},
	}

	parts := []*genai.Part{
		{Text: input.Prompt},
	}

	for _, refImg := range input.ReferenceImages {
		parts = append(parts, &genai.Part{
			InlineData: &genai.Blob{
				Data:     refImg.Data,
				MIMEType: refImg.MIMEType,
			},
		})
	}

	content := []*genai.Content{
		{Parts: parts},
	}

	result, err := c.client.Models.GenerateContent(
		ctx,
		modelGeminiNanoBanana,
		content,
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
				genResult.Images = append(genResult.Images, dto.GeneratedImage{
					Data:     part.InlineData.Data,
					MIMEType: part.InlineData.MIMEType,
				})
			}
		}
	}

	return genResult, nil
}
