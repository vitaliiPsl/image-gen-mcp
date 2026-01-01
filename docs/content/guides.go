package content

import _ "embed"

//go:embed prompting.md
var PromptingGuide string

//go:embed examples.md
var PromptingExamples string


type Template struct {
	Name        string
	Category    string
	Description string
	Template    string
	Example     string
}

var TemplateLibrary = []Template{
	{
		Name:        "photorealistic_portrait",
		Category:    "photography",
		Description: "Professional portrait photography with detailed lighting and composition",
		Template:    "Professional portrait of [subject description], [camera angle], [lighting setup], shot with [lens] at [aperture] for [depth of field effect], [color grading], [photography style] style",
		Example:     "Professional portrait of a woman in her 30s with auburn hair, three-quarter face angle, golden hour window lighting creating rim light, shot with 85mm lens at f/1.8 for soft background blur, warm color grading, fashion photography style",
	},
	{
		Name:        "photorealistic_landscape",
		Category:    "photography",
		Description: "Stunning landscape photography with atmospheric details",
		Template:    "[Landscape type] at [time of day], [composition type], [foreground elements], [midground elements], [background elements], [weather/atmosphere], shot with [lens type], [lighting description], [photography style] style",
		Example:     "Majestic mountain landscape at sunrise, wide panoramic shot, pristine alpine lake in foreground, scattered pine trees in midground, snow-capped peaks in background, morning mist in valleys, shot with wide-angle lens, dramatic golden hour lighting, National Geographic style photography",
	},
	{
		Name:        "product_photography",
		Category:    "commercial",
		Description: "Clean commercial product photography for e-commerce or marketing",
		Template:    "[Product name] in [color/finish], [camera angle] on [background type], [lighting setup], [aperture] for [focus description], highlighting [key features], [photography type], [quality level]",
		Example:     "Premium wireless headphones in matte black finish, 3/4 angle product view on clean white seamless background, studio lighting with soft shadows, f/8 for complete product focus, highlighting metallic accents and leather ear cups, commercial product photography, 4K sharp detail",
	},
	{
		Name:        "character_design",
		Category:    "illustration",
		Description: "Detailed character illustration for games, books, or concept art",
		Template:    "[Character description] with [distinctive features], [pose/position], wearing [clothing/armor details], [holding/interacting with objects], [background setting], [art style] with [texture/rendering description], [genre] character illustration",
		Example:     "A fierce elven warrior with silver braided hair and glowing emerald eyes, standing in confident pose, wearing intricate leather armor with Celtic knotwork details, holding an ornate bow, magical forest background, concept art style with painterly textures, fantasy RPG character illustration",
	},
	{
		Name:        "stylized_illustration",
		Category:    "illustration",
		Description: "Modern stylized illustration with specific art direction",
		Template:    "[Subject] illustration, [art style] style, [color palette description], [composition type], [design approach] with [texture/detail level], modern [category] illustration",
		Example:     "Nighttime city skyline illustration, geometric art deco style, vibrant gradient from deep purple to orange, simplified composition, flat design with subtle texture, modern digital illustration",
	},
	{
		Name:        "childrens_book",
		Category:    "illustration",
		Description: "Warm, friendly illustrations perfect for children's content",
		Template:    "[Adorable subject] with [cute features], [action/pose], in [whimsical setting] with [magical elements], [art medium] style with [characteristic details], children's book illustration, [mood/atmosphere]",
		Example:     "Adorable cartoon fox cub with oversized eyes and fluffy orange fur, sitting curiously with head tilted, in a whimsical enchanted forest clearing with mushrooms and fireflies, watercolor illustration style with soft edges and gentle colors, children's book illustration, warm and inviting atmosphere",
	},
	{
		Name:        "social_media_post",
		Category:    "marketing",
		Description: "Eye-catching social media graphics optimized for engagement",
		Template:    "[Campaign theme] social media graphic, [aspect ratio] format, [background description], bold text reading '[main headline]' in [font style], '[secondary text]' [positioned where], [decorative elements], [design style], [platform]-ready",
		Example:     "Summer sale social media graphic, 1:1 square format, bright gradient background from coral pink to sunny yellow, bold text reading 'SUMMER SALE' in white modern sans-serif font at top, '50% OFF' in large numbers centered, tropical leaf decorative elements in corners, minimal and clean design, Instagram-ready",
	},
	{
		Name:        "product_mockup",
		Category:    "marketing",
		Description: "Lifestyle product mockup in real-world context",
		Template:    "[Product type] on [surface/setting], [product details and colors], [camera angle], [environmental props], [lighting source and direction], lifestyle product photography, [depth of field], [background setting], [color tone]",
		Example:     "Premium coffee bag mockup on rustic wooden table, matte black packaging with gold foil logo, 3/4 angle view, scattered coffee beans around base, soft natural window lighting from left creating gentle shadows, lifestyle product photography, shallow depth of field with blurred café background, warm earth tones",
	},
	{
		Name:        "event_poster",
		Category:    "marketing",
		Description: "Vibrant promotional poster for events and announcements",
		Template:    "[Event type] poster in [aspect ratio] format, [artistic style], headline '[event name]' in [typography style], [date/details placement], [visual elements], [color scheme/aesthetic], high contrast",
		Example:     "Music festival poster in portrait 2:3 format, psychedelic art style with flowing organic shapes, headline 'SOUND WAVE FESTIVAL' in custom groovy typography at top, June 15-17 date below, illustrated crowd silhouettes at bottom, retro 1970s aesthetic with modern twist, high contrast and saturated colors",
	},
	{
		Name:        "infographic",
		Category:    "educational",
		Description: "Clear educational diagrams and infographics",
		Template:    "[Subject] diagram with [illustration style], [color scheme], clearly labeled [key elements], [visual structure type], [design approach], readable [font type] labels, suitable for [use case]",
		Example:     "Water cycle diagram with clean scientific illustration style, light blue color scheme, clearly labeled stages (evaporation, condensation, precipitation, collection), arrows showing flow direction, cross-section view, minimal design with readable sans-serif labels, suitable for textbook",
	},
	{
		Name:        "technical_illustration",
		Category:    "educational",
		Description: "Precise technical drawings and engineering diagrams",
		Template:    "Detailed [view type] technical illustration of [subject], [perspective type], color-coded [components with colors], precise [detail level], [background type], [line style], [discipline] diagram aesthetic, clearly showing [specific features]",
		Example:     "Detailed cutaway technical illustration of a car engine, isometric perspective, color-coded components (blue for cooling, red for fuel, gray for mechanical), precise mechanical details, clean white background, technical drawing style with thin line weights, engineering diagram aesthetic, clearly showing internal components",
	},
	{
		Name:        "anime_character",
		Category:    "artistic_styles",
		Description: "Anime and manga style character illustrations",
		Template:    "Anime character design of [character description] with [distinctive anime features], [expression/pose], wearing [outfit details], [background elements], modern anime style with [rendering technique], [color approach], studio anime quality",
		Example:     "Anime character design of a cheerful high school girl with long pink hair in twin tails, large expressive blue eyes, peace sign pose, wearing school uniform with red bow tie, cherry blossoms in background, modern anime style with clean cel-shading, vibrant colors, studio anime quality",
	},
	{
		Name:        "oil_painting",
		Category:    "artistic_styles",
		Description: "Classical fine art oil painting style",
		Template:    "Classical oil painting of [subject], [art movement] style, [lighting technique], [color palette], [background description], visible brush strokes and [texture technique], museum-quality fine art",
		Example:     "Classical oil painting of roses in ornate ceramic vase, Dutch Golden Age still life style, dramatic chiaroscuro lighting from single window, rich burgundy and cream colored roses, dark atmospheric background, visible brush strokes and impasto texture, museum-quality fine art",
	},
	{
		Name:        "architectural_visualization",
		Category:    "professional",
		Description: "Professional architectural rendering and photography",
		Template:    "[Building type] exterior, [camera angle] architectural photography, [materials and construction], [foreground elements], [time of day] lighting creating [effect], [location/setting], shot with [lens type] for architectural precision, professional [photography type], [quality level]",
		Example:     "Modern minimalist beach house exterior, low angle architectural photography, white concrete and floor-to-ceiling glass construction, infinity pool in foreground, golden hour lighting creating long shadows, Pacific coast setting with ocean in background, shot with tilt-shift lens for architectural precision, professional real estate photography, 4K detail",
	},
	{
		Name:        "food_photography",
		Category:    "professional",
		Description: "Appetizing professional food photography",
		Template:    "[Style] food photography of [dish/food item], [fresh/preparation detail], [plating/presentation], on [surface/setting], [props and styling elements], [lighting direction and quality], [depth of field], [color tone], [camera angle], professional [publication type] style",
		Example:     "Rustic food photography of artisan sourdough bread, fresh from oven with steam rising, broken open showing airy interior crumb, on dark moody wooden cutting board, scattered flour and wheat stalks as props, dramatic side lighting creating texture shadows, shallow depth of field, warm earthy tones, overhead 45-degree angle, professional culinary magazine style",
	},
}

func GetTemplateByName(name string) *Template {
	for i := range TemplateLibrary {
		if TemplateLibrary[i].Name == name {
			return &TemplateLibrary[i]
		}
	}
	return nil
}

func GetTemplatesByCategory(category string) []Template {
	var templates []Template
	for _, template := range TemplateLibrary {
		if template.Category == category {
			templates = append(templates, template)
		}
	}
	return templates
}

func GetAllCategories() []string {
	categoryMap := make(map[string]bool)
	for _, template := range TemplateLibrary {
		categoryMap[template.Category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for category := range categoryMap {
		categories = append(categories, category)
	}
	return categories
}

func FormatTemplateList() string {
	result := "# Available Prompt Templates\n\n"

	categories := GetAllCategories()
	for _, category := range categories {
		result += "## " + category + "\n\n"
		templates := GetTemplatesByCategory(category)
		for _, template := range templates {
			result += "### " + template.Name + "\n"
			result += template.Description + "\n\n"
			result += "**Template:**\n" + template.Template + "\n\n"
			result += "**Example:**\n" + template.Example + "\n\n"
		}
	}

	return result
}
