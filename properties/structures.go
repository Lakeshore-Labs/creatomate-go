package properties

// Fill represents fill properties
type Fill struct {
	Mode   FillMode        `json:"mode,omitempty"`
	Color  interface{}     `json:"color,omitempty"` // string or []FillColorStop
	X0     interface{}     `json:"x0,omitempty"`     // number or string
	Y0     interface{}     `json:"y0,omitempty"`     // number or string
	X1     interface{}     `json:"x1,omitempty"`     // number or string
	Y1     interface{}     `json:"y1,omitempty"`     // number or string
	Radius interface{}     `json:"radius,omitempty"` // number or string
}

// FillColorStop represents a color stop in gradient
type FillColorStop struct {
	Offset float64 `json:"offset"`
	Color  string  `json:"color"`
}

// Shadow represents shadow properties
type Shadow struct {
	Color     string  `json:"color,omitempty"`
	OffsetX   float64 `json:"offset_x,omitempty"`
	OffsetY   float64 `json:"offset_y,omitempty"`
	Blur      float64 `json:"blur,omitempty"`
}

// Stroke represents stroke properties
type Stroke struct {
	Color string     `json:"color,omitempty"`
	Width float64    `json:"width,omitempty"`
	Cap   StrokeCap  `json:"cap,omitempty"`
	Join  StrokeJoin `json:"join,omitempty"`
}

// Font represents font properties
type Font struct {
	Family    string  `json:"family,omitempty"`
	Size      float64 `json:"size,omitempty"`
	Weight    int     `json:"weight,omitempty"`
	Style     string  `json:"style,omitempty"`
	Transform TextTransform `json:"transform,omitempty"`
}

// FontDefinition represents custom font definition
type FontDefinition struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// Blur represents blur properties
type Blur struct {
	Mode   BlurMode `json:"mode,omitempty"`
	Amount float64  `json:"amount,omitempty"`
	Angle  float64  `json:"angle,omitempty"` // For motion blur
}

// ColorFilter represents color filter properties
type ColorFilter struct {
	Type  ColorFilterType `json:"type"`
	Value float64         `json:"value,omitempty"`
}

// TextBackground represents text background properties
type TextBackground struct {
	Color   string  `json:"color,omitempty"`
	Padding float64 `json:"padding,omitempty"`
}

// Warp represents warp properties
type Warp struct {
	Mode   WarpMode    `json:"mode"`
	Amount float64     `json:"amount,omitempty"`
	Points []WarpPoint `json:"points,omitempty"`
}

// WarpPoint represents a control point for warp
type WarpPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Keyframe represents an animation keyframe
type Keyframe[T any] struct {
	Time   float64 `json:"time"`
	Value  T       `json:"value"`
	Easing Easing  `json:"easing,omitempty"`
}