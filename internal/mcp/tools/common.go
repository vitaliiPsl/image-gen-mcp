package tools

import (
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func errorResult(msg string) (*mcp.CallToolResult, any, error) {
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "Error: " + msg},
		},
		IsError: true,
	}, nil, nil
}
