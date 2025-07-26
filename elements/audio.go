package elements

type AudioProperties struct {
	ElementProperties

	// Identical to duration from the common properties, but can be set to "media".
	Duration interface{} `json:"duration,omitempty"` // number, string, or "media"

	// The URL of an audio file you want to play.
	Source string `json:"source"`

	// Trims the source audio to begin at the specified time (in seconds).
	TrimStart interface{} `json:"trim_start,omitempty"` // number or string

	// Trims the source audio so that it stops playing after the specified duration.
	TrimDuration interface{} `json:"trim_duration,omitempty"` // number or string

	// When set to true, the audio starts over when it reaches the end.
	Loop bool `json:"loop,omitempty"`

	// Adjusts the volume from 0% to 100%.
	Volume interface{} `json:"volume,omitempty"` // number or string

	// Fades in the volume for the specified duration (in seconds) at the beginning.
	AudioFadeIn interface{} `json:"audio_fade_in,omitempty"` // number or string

	// Fades out the volume for the specified duration (in seconds) at the end.
	AudioFadeOut interface{} `json:"audio_fade_out,omitempty"` // number or string
}

type Audio struct {
	BaseElement
}

func NewAudio(props AudioProperties) *Audio {
	return &Audio{
		BaseElement: BaseElement{
			Type:       "audio",
			Properties: props,
		},
	}
}