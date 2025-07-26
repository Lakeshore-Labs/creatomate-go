package animations

import (
	"github.com/Lakeshore-Labs/creatomate-go/properties"
	"github.com/Lakeshore-Labs/creatomate-go/utility"
)

type AnimationBase interface {
	ToMap() map[string]interface{}
}

type BaseAnimation struct {
	Type       string
	Properties interface{}
}

func (a *BaseAnimation) ToMap() map[string]interface{} {
	result := utility.TransformObjectKeysToSnake(a.Properties)
	result["type"] = a.Type
	return result
}

type AnimationProperties struct {
	// The time at which the animation starts, relative to the element's timeline.
	Time interface{} `json:"time,omitempty"` // number or string

	// The duration of the animation.
	Duration interface{} `json:"duration,omitempty"` // number or string

	// The easing function.
	Easing properties.Easing `json:"easing,omitempty"`

	// Whether the animation is reversed.
	Reversed bool `json:"reversed,omitempty"`

	// Whether this is a transition animation.
	Transition bool `json:"transition,omitempty"`
}

// Common animation types

type FadeProperties struct {
	AnimationProperties
	
	// The starting opacity (0-100).
	From float64 `json:"from,omitempty"`
	
	// The ending opacity (0-100).
	To float64 `json:"to,omitempty"`
}

type Fade struct {
	BaseAnimation
}

func NewFade(props FadeProperties) *Fade {
	return &Fade{
		BaseAnimation: BaseAnimation{
			Type:       "fade",
			Properties: props,
		},
	}
}

type SlideProperties struct {
	AnimationProperties
	
	// The direction of the slide.
	Direction string `json:"direction,omitempty"` // "left", "right", "up", "down"
	
	// The distance to slide.
	Distance interface{} `json:"distance,omitempty"` // number or string
}

type Slide struct {
	BaseAnimation
}

func NewSlide(props SlideProperties) *Slide {
	return &Slide{
		BaseAnimation: BaseAnimation{
			Type:       "slide",
			Properties: props,
		},
	}
}

type ScaleProperties struct {
	AnimationProperties
	
	// The starting scale.
	From float64 `json:"from,omitempty"`
	
	// The ending scale.
	To float64 `json:"to,omitempty"`
}

type Scale struct {
	BaseAnimation
}

func NewScale(props ScaleProperties) *Scale {
	return &Scale{
		BaseAnimation: BaseAnimation{
			Type:       "scale",
			Properties: props,
		},
	}
}

type SpinProperties struct {
	AnimationProperties
	
	// The number of rotations.
	Rotations float64 `json:"rotations,omitempty"`
	
	// The direction of spin ("clockwise" or "counterclockwise").
	Direction string `json:"direction,omitempty"`
}

type Spin struct {
	BaseAnimation
}

func NewSpin(props SpinProperties) *Spin {
	return &Spin{
		BaseAnimation: BaseAnimation{
			Type:       "spin",
			Properties: props,
		},
	}
}