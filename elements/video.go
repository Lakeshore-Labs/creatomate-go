package elements

import "github.com/Lakeshore-Labs/creatomate-go/properties"

type VideoProperties struct {
	ElementProperties

	// Identical to duration from the common properties, but can be set to "media".
	Duration interface{} `json:"duration,omitempty"` // number, string, or "media"

	// The URL of a video (an mp4) you want to display.
	Source string `json:"source"`

	// This optional parameter indicates whether to generate the video using a third-party AI platform.
	Provider string `json:"provider,omitempty"`

	// Trims the source video to begin at the specified time (in seconds).
	TrimStart interface{} `json:"trim_start,omitempty"` // number or string

	// Trims the source video so that it stops playing after the specified duration.
	TrimDuration interface{} `json:"trim_duration,omitempty"` // number or string

	// When set to true, the video starts over when it reaches the end.
	Loop bool `json:"loop,omitempty"`

	// Adjusts the volume from 0% to 100%.
	Volume interface{} `json:"volume,omitempty"` // number or string

	// Fades in the volume for the specified duration (in seconds) at the beginning.
	AudioFadeIn interface{} `json:"audio_fade_in,omitempty"` // number or string

	// Fades out the volume for the specified duration (in seconds) at the end.
	AudioFadeOut interface{} `json:"audio_fade_out,omitempty"` // number or string

	// This property specifies how the video should be resized to fit the element.
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

type Video struct {
	BaseElement
}

func NewVideo(props VideoProperties) *Video {
	return &Video{
		BaseElement: BaseElement{
			Type:       "video",
			Properties: props,
		},
	}
}