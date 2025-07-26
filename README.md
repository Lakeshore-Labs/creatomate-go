# Creatomate Go SDK

A Go SDK for the [Creatomate](https://creatomate.com) API, providing 100% JSON-compatible output with the official Node.js package.

## Installation

```bash
go get github.com/Lakeshore-Labs/creatomate-go
```

## Features

- ✅ 100% JSON compatibility with the Node.js package
- ✅ Full element support (Video, Image, Text, Audio, Shapes, Compositions)
- ✅ Keyframe animations
- ✅ Text animations (Slide, Typewriter, Appear, etc.)
- ✅ Property expansion (Font, TextBackground)
- ✅ Automatic snake_case JSON field conversion
- ✅ Type-safe API with Go generics
- ✅ Context support for cancellation
- ✅ Comprehensive error handling

## Quick Start

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"
    
    creatomate "github.com/Lakeshore-Labs/creatomate-go"
    "github.com/Lakeshore-Labs/creatomate-go/elements"
    "github.com/Lakeshore-Labs/creatomate-go/properties"
)

func main() {
    // Create client
    client := creatomate.NewClient("your-api-key")
    
    // Create a simple video
    source := creatomate.NewSource(creatomate.SourceProperties{
        OutputFormat: properties.OutputFormatMP4,
        Width:        1920,
        Height:       1080,
        Duration:     10,
        Elements: []interface{}{
            elements.NewVideo(elements.VideoProperties{
                Source:   "https://example.com/video.mp4",
                Duration: "media",
            }),
        },
    })
    
    // Render
    ctx := context.Background()
    renders, err := client.Render(ctx, creatomate.RenderOptions{
        Source: source,
    }, 5*time.Minute)
    
    if err != nil {
        log.Fatal(err)
    }
    
    for _, render := range renders {
        fmt.Printf("Render completed: %s\n", render.URL)
    }
}
```

## Examples

### Text with Animations

```go
font := creatomate.NewFont("Open Sans", 700)
font.Maximum = "10.4 vmin"

text := elements.NewText(elements.TextProperties{
    ElementProperties: elements.ElementProperties{
        Y:          "75%",
        Width:      "100%",
        Height:     "50%",
        Enter: creatomate.NewTextSlide(map[string]interface{}{
            "duration": 2,
            "easing":   "quadratic-out",
            "split":    "line",
        }),
    },
    Text:       "Hello Creatomate! 🔥",
    Font:       font,
    Background: creatomate.NewTextBackground("rgba(255,255,255,0.69)", "23%", "8%", "0%", "0%"),
    FillColor:  "#333333",
})
```

### Keyframe Animations

```go
shape := elements.NewShape(elements.ShapeProperties{
    ElementProperties: elements.ElementProperties{
        Width:  "50%",
        Height: "50%",
        XScale: []interface{}{
            creatomate.NewKeyframe("20%", 0),
            creatomate.NewKeyframeWithEasing("100%", 2, "elastic-out"),
        },
        ZRotation: []interface{}{
            creatomate.NewKeyframe(-90, 0),
            creatomate.NewKeyframeWithEasing(0, 2, "elastic-out"),
        },
    },
    FillColor: []interface{}{
        creatomate.NewKeyframe("#333333", 0),
        creatomate.NewKeyframe("#0079ff", 2),
    },
})
```

### Compositions

```go
composition := elements.NewComposition(elements.CompositionProperties{
    ElementProperties: elements.ElementProperties{
        Width: []interface{}{
            creatomate.NewKeyframe("100%", 1),
            creatomate.NewKeyframe("50%", 3),
        },
    },
    Elements: []interface{}{
        elements.NewImage(elements.ImageProperties{
            Source: "https://example.com/image.jpg",
        }),
        elements.NewText(elements.TextProperties{
            Text: "Grouped elements",
        }),
    },
})
```

## API Reference

### Client

```go
client := creatomate.NewClient(apiKey)

// Render and wait for completion
renders, err := client.Render(ctx, options, timeout)

// Start render without waiting
renders, err := client.StartRender(ctx, options)

// Check render status
render, err := client.FetchRender(ctx, renderID)
```

### Elements

- `elements.NewVideo()` - Video element
- `elements.NewImage()` - Image element
- `elements.NewText()` - Text element
- `elements.NewAudio()` - Audio element
- `elements.NewShape()` - Shape element
- `elements.NewRectangle()` - Rectangle shape
- `elements.NewEllipse()` - Ellipse shape
- `elements.NewComposition()` - Composition for grouping elements

### Properties

All property types from the Node.js SDK are available:
- `properties.OutputFormat` - Output formats (MP4, GIF, PNG, JPG)
- `properties.Fit` - Content fitting modes
- `properties.BlendMode` - Layer blend modes
- `properties.TextTransform` - Text transformations
- `properties.Easing` - Animation easing functions
- And many more...

## Testing

The package includes comprehensive tests that verify JSON output matches the Node.js package exactly:

```bash
cd test
go test -v ./...
```

## JSON Compatibility

This package produces identical JSON output to the official Node.js package. All property names are automatically converted to snake_case, and special types like Font and TextBackground are expanded into their constituent properties to match the API expectations.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.