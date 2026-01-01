package main

import (
	"context"
	"log"

	"github.com/vitaliipsl/image-gen-mcp/internal/client"
	"github.com/vitaliipsl/image-gen-mcp/internal/config"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp/tools"
	"github.com/vitaliipsl/image-gen-mcp/internal/storage"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	imageStorage, err := storage.NewFileStorage(cfg.OutputDir)
	if err != nil {
		log.Fatalf("Failed to initialize storage: %v", err)
	}

	// Clients
	geminiClient, err := client.NewGeminiClient(ctx, cfg.GeminiAPIKey)
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// MCP Tools
	generateImageTool := tools.NewGenerateImageTool(geminiClient, imageStorage)

	// MCP Server
	srv := mcp.New(generateImageTool)
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
