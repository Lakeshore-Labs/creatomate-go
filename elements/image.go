package elements

import "github.com/Lakeshore-Labs/creatomate-go/properties"

type ImageProperties struct {
	ElementProperties

	// The URL of an image you want to display.
	Source string `json:"source"`

	// This optional parameter indicates whether to generate the image using a third-party AI platform.
	Provider string `json:"provider,omitempty"`

	// This property specifies how the image should be resized to fit the element.
	Fit properties.Fit `json:"fit,omitempty"`

	// The stroke.
	Stroke ValueOrKeyframes[*properties.Stroke] `json:"stroke,omitempty"`

	// The stroke color of the element.
	StrokeColor ValueOrKeyframes[string] `json:"stroke_color,omitempty"`

	// The size of the stroke.
	StrokeWidth ValueOrKeyframes[interface{}] `json:"stroke_width,omitempty"`

	// The stroke cap.
	StrokeCap ValueOrKeyframes[properties.StrokeCap] `json:"stroke_cap,omitempty"`

	// The stroke join.
	StrokeJoin ValueOrKeyframes[properties.StrokeJoin] `json:"stroke_join,omitempty"`

	// The border radius of the element.
	BorderRadius ValueOrKeyframes[interface{}] `json:"border_radius,omitempty"`
}

type Image struct {
	BaseElement
}

func NewImage(props ImageProperties) *Image {
	return &Image{
		BaseElement: BaseElement{
			Type:       "image",
			Properties: props,
		},
	}
}