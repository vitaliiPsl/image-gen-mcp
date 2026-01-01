package docs

import (
	_ "embed"
	"encoding/base64"
	"fmt"
)

//go:embed icon.png
var iconPNG []byte

func IconDataURI() string {
	encoded := base64.StdEncoding.EncodeToString(iconPNG)
	return fmt.Sprintf("data:image/png;base64,%s", encoded)
}
