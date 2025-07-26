package creatomate

// TextSlide represents a text slide animation
type TextSlide struct {
	Duration         float64 `json:"duration,omitempty"`
	Easing           string  `json:"easing,omitempty"`
	Split            string  `json:"split,omitempty"`            // "letter", "word", "line"
	Scope            string  `json:"scope,omitempty"`            // "element", "text"
	BackgroundEffect string  `json:"background_effect,omitempty"` // "scaling-clip", etc.
}

// NewTextSlide creates a new text slide animation
func NewTextSlide(config map[string]interface{}) *TextSlide {
	ts := &TextSlide{}
	
	if duration, ok := config["duration"].(float64); ok {
		ts.Duration = duration
	}
	if duration, ok := config["duration"].(int); ok {
		ts.Duration = float64(duration)
	}
	if easing, ok := config["easing"].(string); ok {
		ts.Easing = easing
	}
	if split, ok := config["split"].(string); ok {
		ts.Split = split
	}
	if scope, ok := config["scope"].(string); ok {
		ts.Scope = scope
	}
	if effect, ok := config["backgroundEffect"].(string); ok {
		ts.BackgroundEffect = effect
	}
	
	return ts
}

// ToMap converts the text slide to a map for JSON serialization
func (ts *TextSlide) ToMap() map[string]interface{} {
	result := map[string]interface{}{
		"type": "text-slide",
	}
	
	if ts.Duration > 0 {
		result["duration"] = ts.Duration
	}
	if ts.Easing != "" {
		result["easing"] = ts.Easing
	}
	if ts.Split != "" {
		result["split"] = ts.Split
	}
	if ts.Scope != "" {
		result["scope"] = ts.Scope
	}
	if ts.BackgroundEffect != "" {
		result["background_effect"] = ts.BackgroundEffect
	}
	
	return result
}