package animations

type TextAnimationProperties struct {
	AnimationProperties
	
	// How the text animation is split.
	Split string `json:"split,omitempty"` // "letter", "word", "line"
	
	// Delay between animated parts.
	Stagger float64 `json:"stagger,omitempty"`
	
	// Whether to randomize the order.
	Random bool `json:"random,omitempty"`
}

type TextAppearProperties struct {
	TextAnimationProperties
	
	// Whether to highlight text as it appears.
	Highlighting bool `json:"highlighting,omitempty"`
}

type TextAppear struct {
	BaseAnimation
}

func NewTextAppear(props TextAppearProperties) *TextAppear {
	return &TextAppear{
		BaseAnimation: BaseAnimation{
			Type:       "text-appear",
			Properties: props,
		},
	}
}

type TextSlideProperties struct {
	TextAnimationProperties
	
	// The direction of the slide.
	Direction string `json:"direction,omitempty"` // "left", "right", "up", "down"
	
	// Whether text is clipped during animation.
	Clipped bool `json:"clipped,omitempty"`
}

type TextSlide struct {
	BaseAnimation
}

func NewTextSlide(props TextSlideProperties) *TextSlide {
	return &TextSlide{
		BaseAnimation: BaseAnimation{
			Type:       "text-slide",
			Properties: props,
		},
	}
}

type TextTypewriterProperties struct {
	AnimationProperties
	
	// Characters per second.
	Speed float64 `json:"speed,omitempty"`
}

type TextTypewriter struct {
	BaseAnimation
}

func NewTextTypewriter(props TextTypewriterProperties) *TextTypewriter {
	return &TextTypewriter{
		BaseAnimation: BaseAnimation{
			Type:       "text-typewriter",
			Properties: props,
		},
	}
}