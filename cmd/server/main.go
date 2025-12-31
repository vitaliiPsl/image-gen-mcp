package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/vitaliipsl/image-gen-mcp/internal/gemini"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp"
	"github.com/vitaliipsl/image-gen-mcp/internal/mcp/tools"
)

func main() {
	_ = godotenv.Load()

	ctx := context.Background()

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		log.Fatal("GEMINI_API_KEY environment variable is required")
	}

	// Clients
	geminiClient, err := gemini.NewClient(ctx, apiKey)
	if err != nil {
		log.Fatalf("Failed to create Gemini client: %v", err)
	}

	// MCP Tools
	generateImageTool := tools.NewGenerateImageTool(geminiClient)

	// MCP Server
	srv := mcp.New(geminiClient, generateImageTool)
	if err := srv.Run(ctx); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
