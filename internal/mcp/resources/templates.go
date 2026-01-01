package resources

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/docs/content"
)

type TemplatesResource struct{}

func NewTemplatesResource() *TemplatesResource {
	return &TemplatesResource{}
}

func (r *TemplatesResource) Definition() *mcp.Resource {
	return &mcp.Resource{
		URI:         "prompt://templates/list",
		Name:        "Prompt Template Library",
		Description: "Collection of 15 ready-to-use prompt templates across categories: photography, illustration, marketing, educational, and professional styles.",
		MIMEType:    "text/markdown",
	}
}

func (r *TemplatesResource) Handler(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      "prompt://templates/list",
				MIMEType: "text/markdown",
				Text:     content.FormatTemplateList(),
			},
		},
	}, nil
}
