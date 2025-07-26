package creatomate

// TextBackground represents text background properties that get expanded in JSON
type TextBackground struct {
	Color          string      `json:"color"`
	XPadding       interface{} `json:"x_padding,omitempty"`       // number or string
	YPadding       interface{} `json:"y_padding,omitempty"`       // number or string
	BorderRadius   interface{} `json:"border_radius,omitempty"`   // number or string
	AlignThreshold interface{} `json:"align_threshold,omitempty"` // number or string
}

// NewTextBackground creates a new TextBackground
func NewTextBackground(color string, xPadding, yPadding, borderRadius, alignThreshold interface{}) *TextBackground {
	return &TextBackground{
		Color:          color,
		XPadding:       xPadding,
		YPadding:       yPadding,
		BorderRadius:   borderRadius,
		AlignThreshold: alignThreshold,
	}
}

// ToMap expands text background properties for JSON serialization
func (tb *TextBackground) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["background_color"] = tb.Color
	
	if tb.XPadding != nil {
		result["background_x_padding"] = tb.XPadding
	}
	if tb.YPadding != nil {
		result["background_y_padding"] = tb.YPadding
	}
	if tb.BorderRadius != nil {
		result["background_border_radius"] = tb.BorderRadius
	}
	if tb.AlignThreshold != nil {
		result["background_align_threshold"] = tb.AlignThreshold
	}
	
	return result
}