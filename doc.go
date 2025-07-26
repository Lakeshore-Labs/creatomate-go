// Package creatomate provides a Go SDK for the Creatomate API.
//
// Creatomate is a video generation API that allows you to create videos programmatically.
// This SDK provides a type-safe interface to the Creatomate API with full support for
// all elements, animations, and properties.
//
// # Basic Usage
//
//	client := creatomate.NewClient("your-api-key")
//	
//	source := creatomate.NewSource(creatomate.SourceProperties{
//	    OutputFormat: properties.OutputFormatMP4,
//	    Width:        1920,
//	    Height:       1080,
//	    Duration:     10,
//	    Elements: []interface{}{
//	        elements.NewVideo(elements.VideoProperties{
//	            Source: "https://example.com/video.mp4",
//	        }),
//	    },
//	})
//	
//	renders, err := client.Render(context.Background(), creatomate.RenderOptions{
//	    Source: source,
//	}, 5*time.Minute)
//
// # JSON Compatibility
//
// This package is designed to produce JSON output that is 100% compatible with the
// official Creatomate Node.js SDK. All property names are automatically converted
// to snake_case, and special types like Font and TextBackground are expanded into
// their constituent properties.
//
// # Elements
//
// The package supports all Creatomate elements:
//   - Video
//   - Image
//   - Text
//   - Audio
//   - Shape (Rectangle, Ellipse)
//   - Composition
//
// Each element type has its own properties and can be animated using keyframes
// or animation presets.
//
// # Animations
//
// Animations can be applied to elements using:
//   - Enter animations (played when element appears)
//   - Exit animations (played when element disappears)
//   - Keyframe animations (custom property animations over time)
//   - Text-specific animations (typewriter, slide, reveal, etc.)
//
// # Error Handling
//
// The SDK provides typed errors for all API error conditions:
//   - BadRequestError
//   - InvalidApiKeyError
//   - InsufficientCreditsError
//   - RateLimitExceededError
//   - ConnectionError
//   - TimeoutError
package creatomate