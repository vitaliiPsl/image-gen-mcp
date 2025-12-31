package mcp

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"

	"github.com/vitaliipsl/image-gen-mcp/internal/gemini"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp/tools"
)

const (
	serverName    = "image-gen-mcp"
	serverVersion = "1.0.0"
)

type Server struct {
	mcpServer *mcp.Server

	generateImageTool *tools.GenerateImageTool
}

func New(geminiClient *gemini.Client, generateImageTool *tools.GenerateImageTool) *Server {
	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:    serverName,
		Version: serverVersion,
	}, nil)

	s := &Server{
		mcpServer:         mcpServer,
		generateImageTool: generateImageTool,
	}

	s.registerTools()

	return s
}

func (s *Server) Run(ctx context.Context) error {
	return s.mcpServer.Run(ctx, &mcp.StdioTransport{})
}

func (s *Server) registerTools() {
	mcp.AddTool(s.mcpServer, s.generateImageTool.Definition(), s.generateImageTool.Handler)
}
