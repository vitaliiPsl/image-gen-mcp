# Image Generation Prompting Guide

## Core Principle

**Describe the scene, don't just list keywords.**

Effective prompts follow this structure:
[Subject] + [Composition] + [Action] + [Location] + [Style]

## Essential Elements

### 1. Subject
Be specific and descriptive about what you want to see.

❌ Avoid: "a robot"
✅ Better: "a stoic robot barista with glowing blue optics and brushed metal finish"

### 2. Composition
Frame your shot with camera terminology.

Examples:
- "extreme close-up shot"
- "wide panoramic view"
- "low angle perspective"
- "bird's eye view"
- "centered composition"

### 3. Action
Describe what's happening in the scene.

Examples:
- "pouring steaming coffee into a ceramic mug"
- "running through a field of wildflowers"
- "sitting contemplatively by a window"

### 4. Location
Specify the setting and environment.

Examples:
- "in a cozy minimalist café"
- "on a misty mountain peak at dawn"
- "inside a futuristic laboratory"

### 5. Style
Define the aesthetic and visual treatment.

Examples:
- "3D Pixar-style animation with warm lighting"
- "photorealistic with cinematic color grading"
- "watercolor illustration with soft edges"
- "vintage film photography aesthetic"

## Complete Example

Simple: "a sunset"

Enhanced: "A vibrant sunset over a calm ocean, wide panoramic shot, with waves gently lapping at the shore, from a sandy beach, photorealistic style with golden hour lighting and dramatic cloud formations"

## Aspect Ratios

Choose the right aspect ratio for your use case:

**Square:**
- **1:1** - Social media (Instagram), profile pictures, icons

**Portrait (Vertical):**
- **2:3** - Traditional portraits, posters
- **3:4** - Standard photo prints
- **4:5** - Instagram portrait posts
- **9:16** - Stories, mobile-first content, TikTok

**Landscape (Horizontal):**
- **3:2** - Classic camera/DSLR format
- **4:3** - Presentations, standard displays
- **16:9** - Desktop wallpapers, YouTube, widescreen
- **21:9** - Cinematic scenes, ultra-wide banners

**Usage:** Specify in your request as "aspect_ratio": "16:9". Default is 1:1 if not specified.

## Quick Tips

1. **Be specific, not vague** - The more detail, the better
2. **Use descriptive language** - Paint a picture with words
3. **Think visually** - Imagine the final image as you write
4. **Include atmosphere** - Mention mood, lighting, weather
5. **Specify quality** - Add terms like "highly detailed", "sharp focus", "4K"
6. **Choose appropriate aspect ratio** - Match your final use case

## Common Mistakes

❌ Too generic: "a cat"
✅ Specific: "a fluffy orange tabby cat with green eyes, close-up portrait, sitting on a windowsill, soft natural lighting, photorealistic"

❌ Just keywords: "sunset, beach, palm tree"
✅ Full description: "A tropical sunset scene with silhouetted palm trees, wide shot from the beach, with orange and pink sky reflecting on wet sand, photorealistic style"

❌ Conflicting styles: "realistic cartoon sketch"
✅ Clear style: "3D rendered cartoon style" OR "photorealistic"

## Advanced Techniques

### Photography Terms for Photorealistic Images

#### Camera Settings
- **Aperture**: "shot at f/1.4 for shallow depth of field" or "f/8 for sharp product focus"
- **Lens Types**: "50mm prime lens", "wide-angle lens", "telephoto lens", "macro lens"
- **Focus**: "tack sharp focus on subject", "bokeh background", "deep focus"

#### Lighting Techniques
- **Golden hour**: Warm, soft light during sunrise/sunset
- **Blue hour**: Cool twilight tones
- **Rim lighting**: Backlit with glowing edges
- **Dramatic shadows**: Strong contrast, film noir style
- **Soft diffused light**: Even, flattering illumination
- **Studio lighting**: Professional setup with key, fill, and rim lights

#### Camera Angles
- **Low angle**: Shot from below, makes subject imposing
- **High angle**: Shot from above, diminutive perspective
- **Dutch angle**: Tilted for dynamic tension
- **Eye level**: Natural, neutral perspective
- **Worm's eye view**: Extreme low angle
- **Bird's eye view**: Directly overhead

#### Quality Terms
- "4K resolution", "8K ultra HD"
- "sharp focus", "crystal clear"
- "highly detailed", "intricate details"
- "professional photography"
- "award-winning photograph"

### Text in Images

For images containing text (posters, mockups, signage):

**Be extremely specific:**
- Exact wording in quotes
- Font style (bold, serif, sans-serif, script)
- Font color
- Position (centered, top-third, bottom)
- Size relative to image
- Text effects (shadow, outline, 3D)

Example:
"Modern tech startup poster with headline reading 'INNOVATE TOMORROW' in bold white sans-serif font, centered at the top third, and smaller subtext 'Join the Revolution' in light gray below"

### Character Consistency

For maintaining consistent characters across images:

1. **Detailed initial description**: Capture all distinctive features
2. **Reference previous outputs**: "Same robot character from previous image"
3. **Iterate progressively**: Make small changes between generations
4. **Specify angle changes**: "Same character, now from side profile"

### Multi-Image References

When using reference images:
- Clearly state each image's purpose
- "Use Image 1 for composition, Image 2 for color palette"
- "Apply the style from Image 1 to the subject from Image 2"
- Maximum effectiveness: 3 images (Flash) or 5 images (Pro)

### Style-Specific Guidance

#### For Illustrations
- Specify art medium: "digital painting", "pencil sketch", "pen and ink"
- Art movement: "impressionist", "art nouveau", "minimalist"
- Artist inspiration: "in the style of Studio Ghibli", "Moebius-inspired"

#### For Product Photography
- Background: "clean white background", "natural wood surface"
- Lighting: "studio lighting", "soft box lighting"
- Angle: "3/4 product view", "flat lay composition"
- Context: "lifestyle product shot with complementary props"

#### For Diagrams & Infographics
- Specify "scientifically accurate" or "technically precise"
- Request "clear labels", "color-coded sections"
- State "minimalist design" or "detailed technical illustration"

### Complex Scene Construction

Break down complex scenes step by step:

1. Establish the environment
2. Place primary subjects
3. Add secondary elements
4. Specify interactions
5. Define lighting and atmosphere
6. Choose final style

Example:
"A bustling cyberpunk street market at night, wide establishing shot. Primary subject: a neon-lit ramen stand in the foreground. Background: towering skyscrapers with holographic advertisements. Street level: diverse crowd of people and androids browsing stalls. Lighting: neon blues and pinks reflecting on wet pavement. Atmosphere: light rain creating atmospheric fog. Style: cinematic sci-fi realism with blade runner aesthetic"

### Iterative Refinement

Start broad, then refine:

1. **First generation**: Basic description
2. **Analyze output**: What's missing or wrong?
3. **Second generation**: Add specific corrections
4. **Repeat**: Until satisfied

Example progression:
- V1: "A coffee shop"
- V2: "A modern minimalist coffee shop with large windows"
- V3: "A modern minimalist coffee shop with floor-to-ceiling windows, showing city street outside, warm interior lighting, mid-afternoon, Scandinavian design aesthetic"
