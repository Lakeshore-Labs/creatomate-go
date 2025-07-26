package creatomate_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	creatomate "github.com/Lakeshore-Labs/creatomate-go"
	"github.com/Lakeshore-Labs/creatomate-go/elements"
	"github.com/Lakeshore-Labs/creatomate-go/properties"
)

func TestSimpleVideoJSON(t *testing.T) {
	// Create Go version
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		Width:        1920,
		Height:       1080,
		Duration:     10,
		FrameRate:    30,
		Elements: []interface{}{
			elements.NewVideo(elements.VideoProperties{
				ElementProperties: elements.ElementProperties{},
				Source:           "https://example.com/video.mp4",
				Duration:         "media",
				Volume:           "50%",
				Fit:              properties.FitCover,
			}),
		},
	})

	// Convert to JSON
	sourceMap := source.ToMap()
	goJSON, err := json.MarshalIndent(sourceMap, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Go JSON: %v", err)
	}

	// Load expected JSON from Node.js
	expectedJSON, err := ioutil.ReadFile("testdata/json-outputs/simple-video.json")
	if err != nil {
		t.Fatalf("Failed to read expected JSON: %v", err)
	}

	// Compare by unmarshaling both and comparing objects
	var expectedObj, actualObj map[string]interface{}
	if err := json.Unmarshal(expectedJSON, &expectedObj); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}
	if err := json.Unmarshal(goJSON, &actualObj); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Deep compare
	if !deepEqual(expectedObj, actualObj) {
		t.Errorf("JSON objects don't match.\nExpected:\n%s\n\nGot:\n%s", expectedJSON, goJSON)
	}
}

func TestTextOverlayJSON(t *testing.T) {
	// Create Go version
	font := creatomate.NewFont("Open Sans", 700)
	font.Maximum = "10.4 vmin"
	
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		FrameRate:    60,
		EmojiStyle:   properties.EmojiStyleApple,
		Elements: []interface{}{
			elements.NewVideo(elements.VideoProperties{
				ElementProperties: elements.ElementProperties{},
				Source:           "https://creatomate-static.s3.amazonaws.com/demo/video4.mp4",
			}),
			elements.NewText(elements.TextProperties{
				ElementProperties: elements.ElementProperties{
					Y:          "75%",
					Width:      "100%",
					Height:     "50%",
					XPadding:   "5 vw",
					YPadding:   "5 vh",
					YAlignment: "100%",
					Enter: creatomate.NewTextSlide(map[string]interface{}{
						"duration":         2,
						"easing":           "quadratic-out",
						"split":            "line",
						"scope":            "element",
						"backgroundEffect": "scaling-clip",
					}),
				},
				Text:       "This text adjusts automatically to the size of the video. 🔥",
				Font:       font,
				Background: creatomate.NewTextBackground("rgba(255,255,255,0.69)", "23%", "8%", "0%", "0%"),
				FillColor:  "#333333",
			}),
		},
	})

	// Convert to JSON
	sourceMap := source.ToMap()
	goJSON, err := json.MarshalIndent(sourceMap, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Go JSON: %v", err)
	}

	// Load expected JSON from Node.js
	expectedJSON, err := ioutil.ReadFile("testdata/json-outputs/text-overlay.json")
	if err != nil {
		t.Fatalf("Failed to read expected JSON: %v", err)
	}

	// Compare by unmarshaling both and comparing objects
	var expectedObj, actualObj map[string]interface{}
	if err := json.Unmarshal(expectedJSON, &expectedObj); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}
	if err := json.Unmarshal(goJSON, &actualObj); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Deep compare
	if !deepEqual(expectedObj, actualObj) {
		t.Errorf("JSON objects don't match.\nExpected:\n%s\n\nGot:\n%s", expectedJSON, goJSON)
	}
}

// Helper to compare JSON objects
func compareJSON(t *testing.T, name string, expected, actual interface{}) {
	expectedJSON, _ := json.MarshalIndent(expected, "", "  ")
	actualJSON, _ := json.MarshalIndent(actual, "", "  ")
	
	if string(expectedJSON) != string(actualJSON) {
		t.Errorf("%s JSON mismatch.\nExpected:\n%s\n\nGot:\n%s", name, expectedJSON, actualJSON)
	}
}

// Test the ToMap output directly
func TestSourceToMap(t *testing.T) {
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		Width:        1920,
		Height:       1080,
		Duration:     10,
		FrameRate:    30,
	})

	result := source.ToMap()

	// Check field names are snake_case
	if _, ok := result["output_format"]; !ok {
		t.Error("Expected output_format field")
	}
	if _, ok := result["frame_rate"]; !ok {
		t.Error("Expected frame_rate field")
	}
}

func TestVideoElementToMap(t *testing.T) {
	video := elements.NewVideo(elements.VideoProperties{
		ElementProperties: elements.ElementProperties{},
		Source:           "test.mp4",
		Duration:         "media",
		Volume:           "50%",
		Fit:              properties.FitCover,
	})

	result := video.ToMap()
	
	// Check type field
	if result["type"] != "video" {
		t.Errorf("Expected type to be 'video', got %v", result["type"])
	}
	
	// Print for debugging
	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	fmt.Printf("Video ToMap result:\n%s\n", jsonBytes)
}

// deepEqual compares two interface{} values deeply
func deepEqual(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func TestCompositionsJSON(t *testing.T) {
	// Create Go version
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		Width:        1280,
		Height:       720,
		Duration:     5,
		Elements: []interface{}{
			elements.NewComposition(elements.CompositionProperties{
				ElementProperties: elements.ElementProperties{
					Width: []interface{}{
						creatomate.NewKeyframe("100%", 1),
						creatomate.NewKeyframe("50%", 3),
					},
					Height: []interface{}{
						creatomate.NewKeyframe("100%", 3),
						creatomate.NewKeyframe("50%", 4),
					},
					YRotation: []interface{}{
						creatomate.NewKeyframe(0, 4),
						creatomate.NewKeyframe(360, 5),
					},
				},
				Elements: []interface{}{
					elements.NewImage(elements.ImageProperties{
						ElementProperties: elements.ElementProperties{},
						Source:           "https://creatomate-static.s3.amazonaws.com/demo/image1.jpg",
					}),
					elements.NewText(elements.TextProperties{
						ElementProperties: elements.ElementProperties{
							Width:      "100%",
							Height:     "10%",
							XAlignment: "50%",
						},
						Text:       "Place elements in the same composition to group them",
						Background: creatomate.NewTextBackground("#fff", "25%", "25%", "20%", "0%"),
					}),
				},
			}),
		},
	})

	// Convert to JSON
	sourceMap := source.ToMap()
	goJSON, err := json.MarshalIndent(sourceMap, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Go JSON: %v", err)
	}

	// Load expected JSON from Node.js
	expectedJSON, err := ioutil.ReadFile("testdata/json-outputs/compositions.json")
	if err != nil {
		t.Fatalf("Failed to read expected JSON: %v", err)
	}

	// Compare by unmarshaling both and comparing objects
	var expectedObj, actualObj map[string]interface{}
	if err := json.Unmarshal(expectedJSON, &expectedObj); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}
	if err := json.Unmarshal(goJSON, &actualObj); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Deep compare
	if !deepEqual(expectedObj, actualObj) {
		t.Errorf("JSON objects don't match.\nExpected:\n%s\n\nGot:\n%s", expectedJSON, goJSON)
	}
}

func TestKeyframesJSON(t *testing.T) {
	// Create Go version with keyframes
	source := creatomate.NewSource(creatomate.SourceProperties{
		OutputFormat: properties.OutputFormatMP4,
		Width:        1280,
		Height:       720,
		Duration:     4,
		Elements: []interface{}{
			elements.NewShape(elements.ShapeProperties{
				ElementProperties: elements.ElementProperties{
					Width:  "23.5227%",
					Height: "41.8179%",
					XScale: []interface{}{
						creatomate.NewKeyframe("20%", 0),
						creatomate.NewKeyframeWithEasing("100%", 2, "elastic-out"),
					},
					YScale: []interface{}{
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
					creatomate.NewKeyframe("#0079ff", 0.94),
					creatomate.NewKeyframe("#0079ff", 2),
					creatomate.NewKeyframe("rgba(0,121,255,0)", 2.5),
				},
				StrokeColor: "rgba(0,121,255,1)",
				StrokeWidth: []interface{}{
					creatomate.NewKeyframe("0 vmin", 2),
					creatomate.NewKeyframe("4.3 vmin", 2.5),
					creatomate.NewKeyframe("0 vmin", 3.5),
				},
				StrokeStart: []interface{}{
					creatomate.NewKeyframe("0%", 2.5),
					creatomate.NewKeyframe("100%", 3.5),
				},
				StrokeOffset: []interface{}{
					creatomate.NewKeyframe("0%", 2.5),
					creatomate.NewKeyframe("50%", 3.5),
				},
				Path: []interface{}{
					creatomate.NewKeyframe("M 0 0 L 100 0 L 100 100 L 0 100 L 0 0 Z", 0.94),
					creatomate.NewKeyframeWithEasing("M -20 -20 C 15 -55 85 -55 120 -20 C 155 15 155 85 120 120 C 85 155 15 155 -20 120 C -55 85 -55 15 -20 -20 Z", 2.5, "elastic-out"),
				},
			}),
		},
	})

	// Convert to JSON
	sourceMap := source.ToMap()
	goJSON, err := json.MarshalIndent(sourceMap, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal Go JSON: %v", err)
	}

	// Load expected JSON from Node.js
	expectedJSON, err := ioutil.ReadFile("testdata/json-outputs/keyframes.json")
	if err != nil {
		t.Fatalf("Failed to read expected JSON: %v", err)
	}

	// Compare by unmarshaling both and comparing objects
	var expectedObj, actualObj map[string]interface{}
	if err := json.Unmarshal(expectedJSON, &expectedObj); err != nil {
		t.Fatalf("Failed to unmarshal expected JSON: %v", err)
	}
	if err := json.Unmarshal(goJSON, &actualObj); err != nil {
		t.Fatalf("Failed to unmarshal actual JSON: %v", err)
	}

	// Deep compare
	if !deepEqual(expectedObj, actualObj) {
		t.Errorf("JSON objects don't match.\nExpected:\n%s\n\nGot:\n%s", expectedJSON, goJSON)
	}
}