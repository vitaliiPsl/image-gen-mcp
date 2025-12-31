package dto

type GenerateImageParams struct {
	Prompt      string `json:"prompt" validate:"required" jsonschema:"The text prompt describing the image to generate"`
	AspectRatio string `json:"aspect_ratio" jsonschema:"Aspect ratio of the generated image in x:y format."`
}

type GeneratedImage struct {
	Data     string
	MIMEType string
}

type GenerateImageResult struct {
	Images []GeneratedImage
}
