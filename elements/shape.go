package elements

import "github.com/Lakeshore-Labs/creatomate-go/properties"

type ShapeProperties struct {
	ElementProperties

	// The fill.
	Fill ValueOrKeyframes[*properties.Fill] `json:"fill,omitempty"`

	// The fill color.
	FillColor ValueOrKeyframes[interface{}] `json:"fill_color,omitempty"`

	// The fill mode.
	FillMode ValueOrKeyframes[string] `json:"fill_mode,omitempty"`

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

	// The stroke start.
	StrokeStart ValueOrKeyframes[interface{}] `json:"stroke_start,omitempty"`

	// The stroke offset.
	StrokeOffset ValueOrKeyframes[interface{}] `json:"stroke_offset,omitempty"`

	// The path for the shape.
	Path ValueOrKeyframes[string] `json:"path,omitempty"`
}

type Shape struct {
	BaseElement
}

func NewShape(props ShapeProperties) *Shape {
	return &Shape{
		BaseElement: BaseElement{
			Type:       "shape",
			Properties: props,
		},
	}
}

type RectangleProperties struct {
	ShapeProperties

	// The border radius of the rectangle.
	BorderRadius ValueOrKeyframes[interface{}] `json:"border_radius,omitempty"`
}

type Rectangle struct {
	BaseElement
}

func NewRectangle(props RectangleProperties) *Rectangle {
	return &Rectangle{
		BaseElement: BaseElement{
			Type:       "rectangle",
			Properties: props,
		},
	}
}

type EllipseProperties struct {
	ShapeProperties
}

type Ellipse struct {
	BaseElement
}

func NewEllipse(props EllipseProperties) *Ellipse {
	return &Ellipse{
		BaseElement: BaseElement{
			Type:       "ellipse",
			Properties: props,
		},
	}
}