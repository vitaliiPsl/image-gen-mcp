package dto

type GenerateImageParams struct {
	Prompt          string   `json:"prompt" validate:"required" jsonschema:"The text prompt describing the image to generate"`
	AspectRatio     string   `json:"aspect_ratio" jsonschema:"Aspect ratio of the generated image in x:y format."`
	ReferenceImages []string `json:"reference_images" jsonschema:"Optional array of file paths to reference images (up 3). Use for character consistency, style matching, or incorporating specific objects/people."`
}

type GenerateImageInput struct {
	Prompt          string
	AspectRatio     string
	ReferenceImages []ReferenceImage
}

type ReferenceImage struct {
	Data     []byte
	MIMEType string
}

type GenerateImageResult struct {
	Images []GeneratedImage
}

type GeneratedImage struct {
	Data     []byte
	MIMEType string
}
