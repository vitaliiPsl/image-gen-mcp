package resources

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/docs/content"
)

type PromptingGuideResource struct{}

func NewPromptingGuideResource() *PromptingGuideResource {
	return &PromptingGuideResource{}
}

func (r *PromptingGuideResource) Definition() *mcp.Resource {
	return &mcp.Resource{
		URI:         "prompt://guide/prompting",
		Name:        "Image Generation Prompting Guide",
		Description: "Comprehensive guide for creating effective image generation prompts. Covers core elements, aspect ratios, advanced techniques, photography terms, text rendering, and style-specific guidance.",
		MIMEType:    "text/markdown",
	}
}

func (r *PromptingGuideResource) Handler(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      "prompt://guide/prompting",
				MIMEType: "text/markdown",
				Text:     content.PromptingGuide,
			},
		},
	}, nil
}
