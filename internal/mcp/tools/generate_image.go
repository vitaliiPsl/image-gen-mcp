package tools

import (
	"context"
	"fmt"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/internal/client"
	"github.com/vitaliipsl/image-gen-mcp/internal/dto"
	"github.com/vitaliipsl/image-gen-mcp/internal/storage"
)

type GenerateImageTool struct {
	geminiClient client.Client
	storage      storage.ImageStorage
}

func NewGenerateImageTool(geminiClient client.Client, imageStorage storage.ImageStorage) *GenerateImageTool {
	return &GenerateImageTool{
		geminiClient: geminiClient,
		storage:      imageStorage,
	}
}

func (t *GenerateImageTool) Definition() *mcp.Tool {
	return &mcp.Tool{
		Name:        "generate_image",
		Description: "Generate an image based on a text prompt. Optionally accepts reference images for character consistency, style matching, or incorporating specific objects/people into the generated image.",
	}
}

func (t *GenerateImageTool) Handler(
	ctx context.Context,
	req *mcp.CallToolRequest,
	params dto.GenerateImageParams,
) (*mcp.CallToolResult, any, error) {
	referenceImages, err := t.loadReferenceImages(params.ReferenceImages)
	if err != nil {
		return errorResult(fmt.Sprintf("failed to load reference images: %v", err))
	}

	input := &dto.GenerateImageInput{
		Prompt:          params.Prompt,
		AspectRatio:     params.AspectRatio,
		ReferenceImages: referenceImages,
	}

	result, err := t.geminiClient.GenerateImage(ctx, input)
	if err != nil {
		return errorResult(fmt.Sprintf("failed to generate image: %v", err))
	}

	if len(result.Images) == 0 {
		return errorResult("no image was generated; the model may have declined the request")
	}

	var filePaths []string
	for _, img := range result.Images {
		filePath, err := t.storage.Save(img.Data, img.MIMEType)
		if err != nil {
			return errorResult(fmt.Sprintf("failed to save image: %v", err))
		}

		filePaths = append(filePaths, filePath)
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{
				Text: fmt.Sprintf("Generated %d image(s):\n%s", len(filePaths), strings.Join(filePaths, "\n")),
			},
		},
	}, nil, nil
}

func (t *GenerateImageTool) loadReferenceImages(paths []string) ([]dto.ReferenceImage, error) {
	var referenceImages []dto.ReferenceImage

	for i, refPath := range paths {
		data, mimeType, err := t.storage.Load(refPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load reference image %d (%s): %w", i+1, refPath, err)
		}

		referenceImages = append(referenceImages, dto.ReferenceImage{
			Data:     data,
			MIMEType: mimeType,
		})
	}

	return referenceImages, nil
}
