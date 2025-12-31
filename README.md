# image-gen-mcp

A Model Context Protocol (MCP) server that provides AI-powered image generation using Google's Gemini API.

## Features

- Generate images from text prompts using Gemini's image generation model
- MCP-compliant server with stdio transport
- Lightweight Docker image

## Prerequisites

- Go 1.24+ (for local development)
- Docker (for containerized deployment)
- Google Gemini API key

## Getting a Gemini API Key

1. Go to [Google AI Studio](https://aistudio.google.com/apikey)
2. Create a new API key
3. Save it for use in configuration

## Installation

### Using Docker (Recommended)

Pull from Docker Hub:

```bash
docker pull vitaliipsl/image-gen-mcp
```

Or build locally:

```bash
docker build -t image-gen-mcp .
```

### From Source

```bash
go build -o bin/image-gen-mcp ./cmd/server
```

## Configuration

The server requires the following environment variable:

| Variable | Required | Description |
|----------|----------|-------------|
| `GEMINI_API_KEY` | Yes | Your Google Gemini API key |

You can also create a `.env` file in the project root:

```env
GEMINI_API_KEY=your_api_key_here
```

## Usage

### With Claude Desktop

Add the following to your Claude Desktop configuration file:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

#### Using Docker

```json
{
  "mcpServers": {
    "image-gen": {
      "command": "docker",
      "args": [
        "run",
        "-i",
        "--rm",
        "-e", "GEMINI_API_KEY=your_api_key_here",
        "vitaliipsl/image-gen-mcp"
      ]
    }
  }
}
```

#### Using Binary

```json
{
  "mcpServers": {
    "image-gen": {
      "command": "/path/to/image-gen-mcp",
      "env": {
        "GEMINI_API_KEY": "your_api_key_here"
      }
    }
  }
}
```

### With Claude Code

Add to your Claude Code MCP settings:

```json
{
  "mcpServers": {
    "image-gen": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "GEMINI_API_KEY=your_api_key_here", "vitaliipsl/image-gen-mcp"]
    }
  }
}
```

## Available Tools

### generate_image

Generate an image based on a text prompt.

**Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `prompt` | string | Yes | The text prompt describing the image to generate |
| `aspect_ratio` | string | No | Aspect ratio in `x:y` format (e.g., `16:9`, `1:1`, `4:3`) |

**Example:**

```
Generate an image of a sunset over mountains with aspect ratio 16:9
```

## License

MIT
