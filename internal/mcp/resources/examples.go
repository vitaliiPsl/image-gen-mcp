package resources

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/docs/content"
)

// ExamplesResource provides the prompt examples gallery
type ExamplesResource struct{}

func NewExamplesResource() *ExamplesResource {
	return &ExamplesResource{}
}

func (r *ExamplesResource) Definition() *mcp.Resource {
	return &mcp.Resource{
		URI:         "prompt://guide/examples",
		Name:        "Prompt Examples Gallery",
		Description: "Curated collection of example prompts demonstrating best practices across different categories: photography, illustrations, marketing, and more.",
		MIMEType:    "text/markdown",
	}
}

func (r *ExamplesResource) Handler(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      "prompt://guide/examples",
				MIMEType: "text/markdown",
				Text:     content.PromptingExamples,
			},
		},
	}, nil
}
