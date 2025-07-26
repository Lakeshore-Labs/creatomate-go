package elements

import "github.com/Lakeshore-Labs/creatomate-go/properties"

type TextProperties struct {
	ElementProperties

	// The text content to display.
	Text string `json:"text"`

	// Font properties.
	Font interface{} `json:"font,omitempty"` // *creatomate.Font

	// The font family.
	FontFamily string `json:"font_family,omitempty"`

	// The font size.
	FontSize ValueOrKeyframes[interface{}] `json:"font_size,omitempty"`

	// The font weight (100-900).
	FontWeight ValueOrKeyframes[int] `json:"font_weight,omitempty"`

	// The font style (normal or italic).
	FontStyle string `json:"font_style,omitempty"`

	// Text transform.
	TextTransform properties.TextTransform `json:"text_transform,omitempty"`

	// Line height as a percentage of the font size.
	LineHeight ValueOrKeyframes[interface{}] `json:"line_height,omitempty"`

	// Letter spacing.
	LetterSpacing ValueOrKeyframes[interface{}] `json:"letter_spacing,omitempty"`

	// The text color.
	Color ValueOrKeyframes[string] `json:"color,omitempty"`

	// Text background.
	TextBackground interface{} `json:"text_background,omitempty"` // *creatomate.TextBackground

	// Background properties (expanded from TextBackground)
	Background interface{} `json:"background,omitempty"` // *creatomate.TextBackground

	// The fill.
	Fill ValueOrKeyframes[*properties.Fill] `json:"fill,omitempty"`

	// The fill color.
	FillColor ValueOrKeyframes[interface{}] `json:"fill_color,omitempty"`

	// The fill mode.
	FillMode ValueOrKeyframes[string] `json:"fill_mode,omitempty"`

	// The stroke.
	Stroke ValueOrKeyframes[*properties.Stroke] `json:"stroke,omitempty"`

	// The stroke color of the text.
	StrokeColor ValueOrKeyframes[string] `json:"stroke_color,omitempty"`

	// The size of the stroke.
	StrokeWidth ValueOrKeyframes[interface{}] `json:"stroke_width,omitempty"`

	// The stroke cap.
	StrokeCap ValueOrKeyframes[properties.StrokeCap] `json:"stroke_cap,omitempty"`

	// The stroke join.
	StrokeJoin ValueOrKeyframes[properties.StrokeJoin] `json:"stroke_join,omitempty"`

	// Text flow direction.
	FlowDirection properties.FlowDirection `json:"flow_direction,omitempty"`

	// Whether text wraps to multiple lines.
	Wrap bool `json:"wrap,omitempty"`

	// Transcript settings for the text.
	TranscriptEffect properties.TranscriptEffect `json:"transcript_effect,omitempty"`
	TranscriptPlacement properties.TranscriptPlacement `json:"transcript_placement,omitempty"`
	TranscriptSplit properties.TranscriptSplit `json:"transcript_split,omitempty"`
	TranscriptMaximumLength int `json:"transcript_maximum_length,omitempty"`
}

type Text struct {
	BaseElement
}

func NewText(props TextProperties) *Text {
	return &Text{
		BaseElement: BaseElement{
			Type:       "text",
			Properties: props,
		},
	}
}