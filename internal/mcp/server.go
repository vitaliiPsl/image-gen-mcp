package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/docs"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp/resources"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp/tools"
)

type Server struct {
	mcpServer *mcp.Server

	generateImageTool *tools.GenerateImageTool

	examplesResource       *resources.ExamplesResource
	templatesResource      *resources.TemplatesResource
	promptingGuideResource *resources.PromptingGuideResource
}

func New(generateImageTool *tools.GenerateImageTool,
	examplesResource *resources.ExamplesResource,
	templatesResource *resources.TemplatesResource,
	promptingGuideResource *resources.PromptingGuideResource,
) *Server {
	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:       "image-gen-mcp",
		Title:      "Image Generation MCP Server",
		Version:    "1.0.0",
		WebsiteURL: "https://github.com/vitaliipsl/image-gen-mcp",
		Icons: []mcp.Icon{
			{
				Source:   docs.IconDataURI(),
				MIMEType: "image/png",
			},
		},
	}, nil)

	s := &Server{
		mcpServer:              mcpServer,
		generateImageTool:      generateImageTool,
		examplesResource:       examplesResource,
		templatesResource:      templatesResource,
		promptingGuideResource: promptingGuideResource,
	}

	s.registerTools()
	s.registerResources()

	return s
}

func (s *Server) Run(ctx context.Context) error {
	return s.mcpServer.Run(ctx, &mcp.StdioTransport{})
}

func (s *Server) registerTools() {
	mcp.AddTool(s.mcpServer, s.generateImageTool.Definition(), s.generateImageTool.Handler)
}

func (s *Server) registerResources() {
	s.mcpServer.AddResource(s.examplesResource.Definition(), s.examplesResource.Handler)
	s.mcpServer.AddResource(s.templatesResource.Definition(), s.templatesResource.Handler)
	s.mcpServer.AddResource(s.promptingGuideResource.Definition(), s.promptingGuideResource.Handler)
}
