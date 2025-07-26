package creatomate

// Font represents font properties that get expanded in JSON
type Font struct {
	Family      string      `json:"family"`
	Weight      *int        `json:"weight,omitempty"`
	Style       *string     `json:"style,omitempty"`
	Size        interface{} `json:"size,omitempty"`        // number or string
	Minimum     interface{} `json:"minimum,omitempty"`     // number or string
	Maximum     interface{} `json:"maximum,omitempty"`     // number or string
}

// NewFont creates a new Font with family and weight
func NewFont(family string, weight int) *Font {
	return &Font{
		Family: family,
		Weight: &weight,
	}
}

// ToMap expands font properties for JSON serialization
func (f *Font) ToMap() map[string]interface{} {
	result := make(map[string]interface{})
	result["font_family"] = f.Family
	
	if f.Weight != nil {
		result["font_weight"] = *f.Weight
	}
	if f.Style != nil {
		result["font_style"] = *f.Style
	}
	if f.Size != nil {
		result["font_size"] = f.Size
	}
	if f.Minimum != nil {
		result["font_size_minimum"] = f.Minimum
	}
	if f.Maximum != nil {
		result["font_size_maximum"] = f.Maximum
	}
	
	return result
}