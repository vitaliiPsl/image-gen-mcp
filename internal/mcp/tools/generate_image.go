package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/internal/dto"
	"github.com/vitaliipsl/image-gen-mcp/internal/gemini"
)

type GenerateImageTool struct {
	geminiClient *gemini.Client
}

func NewGenerateImageTool(geminiClient *gemini.Client) *GenerateImageTool {
	return &GenerateImageTool{
		geminiClient: geminiClient,
	}
}

func (t *GenerateImageTool) Definition() *mcp.Tool {
	return &mcp.Tool{
		Name:        "generate_image",
		Description: "Generate an image based on a text prompt",
	}
}

func (t *GenerateImageTool) Handler(
	ctx context.Context,
	req *mcp.CallToolRequest,
	params dto.GenerateImageParams,
) (*mcp.CallToolResult, any, error) {
	result, err := t.geminiClient.GenerateImage(ctx, &params)
	if err != nil {
		return errorResult(fmt.Sprintf("failed to generate image: %v", err))
	}

	if len(result.Images) == 0 {
		return errorResult("no image was generated; the model may have declined the request")
	}

	var content []mcp.Content
	for _, img := range result.Images {
		content = append(content, &mcp.ImageContent{
			Data:     []byte(img.Data),
			MIMEType: img.MIMEType,
		})
	}

	return &mcp.CallToolResult{Content: content}, nil, nil
}
