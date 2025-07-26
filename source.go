package creatomate

import (
	"encoding/json"
	
	"github.com/Lakeshore-Labs/creatomate-go/properties"
)

type ValueOrKeyframes[T any] interface{}

type SourceProperties struct {
	// The output format of the render, which can be jpg, png, gif, or mp4.
	OutputFormat properties.OutputFormat `json:"output_format"`

	// Only for MP4 renders. Sets the constant rate factor (CRF) ranging from 17 to 51.
	CRF int `json:"crf,omitempty"`

	// Only for GIF renders. With 'best', the GIF generation takes much longer.
	GifQuality properties.GifQuality `json:"gif_quality,omitempty"`

	// Only for GIF renders. A number ranging from 0 to 200 indicating the compression level.
	GifCompression int `json:"gif_compression,omitempty"`

	// The width of the output in pixels.
	Width int `json:"width,omitempty"`

	// The height of the output in pixels.
	Height int `json:"height,omitempty"`

	// The frame rate of the rendered video.
	FrameRate float64 `json:"frame_rate,omitempty"`

	// The duration of the output in seconds.
	Duration float64 `json:"duration,omitempty"`

	// Only for GIF renders. Set to true to make the GIF repeat.
	Loop bool `json:"loop,omitempty"`

	// If a snapshot image is desired, specify the time in seconds.
	SnapshotTime float64 `json:"snapshot_time,omitempty"`

	// The style of the Emojis used in text elements.
	EmojiStyle properties.EmojiStyle `json:"emoji_style,omitempty"`

	// The background fill.
	Fill ValueOrKeyframes[*properties.Fill] `json:"fill,omitempty"`

	// The background fill color.
	FillColor ValueOrKeyframes[string] `json:"fill_color,omitempty"`

	// The fill method used: solid, linear, and radial.
	FillMode ValueOrKeyframes[string] `json:"fill_mode,omitempty"`

	// The start position of the gradient on the x-axis.
	FillX0 ValueOrKeyframes[interface{}] `json:"fill_x0,omitempty"`

	// The start position of the gradient on the y-axis.
	FillY0 ValueOrKeyframes[interface{}] `json:"fill_y0,omitempty"`

	// The end position of the gradient on the x-axis.
	FillX1 ValueOrKeyframes[interface{}] `json:"fill_x1,omitempty"`

	// The end position of the gradient on the y-axis.
	FillY1 ValueOrKeyframes[interface{}] `json:"fill_y1,omitempty"`

	// The radius of the radial gradient.
	FillRadius ValueOrKeyframes[interface{}] `json:"fill_radius,omitempty"`

	// Custom fonts array.
	Fonts []properties.FontDefinition `json:"fonts,omitempty"`

	// Elements that make up the render.
	Elements []interface{} `json:"elements,omitempty"`
}

type Source struct {
	Properties SourceProperties
}

func NewSource(properties SourceProperties) *Source {
	return &Source{Properties: properties}
}

func (s *Source) ToMap() map[string]interface{} {
	// Convert struct to JSON then back to map to handle tags
	propsJSON, err := json.Marshal(s.Properties)
	if err != nil {
		return make(map[string]interface{})
	}
	
	var result map[string]interface{}
	if err := json.Unmarshal(propsJSON, &result); err != nil {
		return make(map[string]interface{})
	}
	
	// Handle elements conversion
	if s.Properties.Elements != nil {
		elements := make([]interface{}, len(s.Properties.Elements))
		for i, elem := range s.Properties.Elements {
			if elemWithMap, ok := elem.(interface{ ToMap() map[string]interface{} }); ok {
				elements[i] = elemWithMap.ToMap()
			} else {
				elements[i] = elem
			}
		}
		result["elements"] = elements
	}
	
	// Remove nil values
	for k, v := range result {
		if v == nil {
			delete(result, k)
		}
	}
	
	return result
}