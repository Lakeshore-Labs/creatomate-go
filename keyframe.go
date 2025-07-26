package creatomate

import "github.com/Lakeshore-Labs/creatomate-go/properties"

// Keyframe represents an animation keyframe
type Keyframe[T any] struct {
	Time   float64           `json:"time"`
	Value  T                 `json:"value"`
	Easing properties.Easing `json:"easing,omitempty"`
}

// NewKeyframe creates a new keyframe with time and value
func NewKeyframe[T any](value T, time float64) *Keyframe[T] {
	return &Keyframe[T]{
		Time:  time,
		Value: value,
	}
}

// NewKeyframeWithEasing creates a new keyframe with time, value, and easing
func NewKeyframeWithEasing[T any](value T, time float64, easing properties.Easing) *Keyframe[T] {
	return &Keyframe[T]{
		Time:   time,
		Value:  value,
		Easing: easing,
	}
}