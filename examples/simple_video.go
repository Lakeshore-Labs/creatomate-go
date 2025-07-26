package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	creatomate "github.com/Lakeshore-Labs/creatomate-go"
	"github.com/Lakeshore-Labs/creatomate-go/elements"
	"github.com/Lakeshore-Labs/creatomate-go/properties"
)

func main() {
	// Get API key from environment
	apiKey := os.Getenv("CREATOMATE_API_KEY")
	if apiKey == "" {
		log.Fatal("Please set CREATOMATE_API_KEY environment variable")
	}

	// Create a simple video with text overlay
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		Width:        1920,
		Height:       1080,
		Duration:     10,
		FrameRate:    30,
		Elements: []interface{}{
			// Background video
			elements.NewVideo(elements.VideoProperties{
				ElementProperties: elements.ElementProperties{
					Track: intPtr(1),
				},
				Source: "https://creatomate-static.s3.amazonaws.com/demo/video4.mp4",
				Fit:    properties.FitCover,
			}),
			// Text overlay
			elements.NewText(elements.TextProperties{
				ElementProperties: elements.ElementProperties{
					Track:      intPtr(2),
					Y:          "75%",
					Width:      "100%",
					Height:     "20%",
					XPadding:   "5 vw",
					YPadding:   "5 vh",
					YAlignment: "100%",
				},
				Text:       "Hello from Creatomate Go!",
				Font:       creatomate.NewFont("Open Sans", 700),
				FillColor:  "#FFFFFF",
				Background: creatomate.NewTextBackground("rgba(0,0,0,0.7)", "20%", "10%", "5%", "0%"),
			}),
		},
	})

	// Create client
	client := creatomate.NewClient(apiKey)

	// Start render
	fmt.Println("Starting render...")
	ctx := context.Background()
	renders, err := client.Render(ctx, creatomate.RenderOptions{
		Source: source,
	}, 5*time.Minute)

	if err != nil {
		log.Fatalf("Render failed: %v", err)
	}

	// Print results
	for _, render := range renders {
		if render.Status == creatomate.RenderStatusSucceeded {
			fmt.Printf("Render completed successfully!\n")
			fmt.Printf("URL: %s\n", render.URL)
			if render.SnapshotURL != "" {
				fmt.Printf("Snapshot: %s\n", render.SnapshotURL)
			}
		} else {
			fmt.Printf("Render failed with status: %s\n", render.Status)
			if render.ErrorMessage != "" {
				fmt.Printf("Error: %s\n", render.ErrorMessage)
			}
		}
	}
}

func intPtr(i int) *int {
	return &i
}