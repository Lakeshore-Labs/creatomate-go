package properties

// OutputFormat represents the output format of the render
type OutputFormat string

const (
	OutputFormatJPG OutputFormat = "jpg"
	OutputFormatPNG OutputFormat = "png"
	OutputFormatGIF OutputFormat = "gif"
	OutputFormatMP4 OutputFormat = "mp4"
)

// GifQuality represents the quality setting for GIF renders
type GifQuality string

const (
	GifQualityFast GifQuality = "fast"
	GifQualityBest GifQuality = "best"
)

// EmojiStyle represents the emoji style
type EmojiStyle string

const (
	EmojiStyleFacebook EmojiStyle = "facebook"
	EmojiStyleGoogle   EmojiStyle = "google"
	EmojiStyleTwitter  EmojiStyle = "twitter"
	EmojiStyleApple    EmojiStyle = "apple"
)

// BlendMode represents how layers blend
type BlendMode string

const (
	BlendModeNormal      BlendMode = "normal"
	BlendModeMultiply    BlendMode = "multiply"
	BlendModeScreen      BlendMode = "screen"
	BlendModeOverlay     BlendMode = "overlay"
	BlendModeDarken      BlendMode = "darken"
	BlendModeLighten     BlendMode = "lighten"
	BlendModeColorDodge  BlendMode = "color-dodge"
	BlendModeColorBurn   BlendMode = "color-burn"
	BlendModeHardLight   BlendMode = "hard-light"
	BlendModeSoftLight   BlendMode = "soft-light"
	BlendModeDifference  BlendMode = "difference"
	BlendModeExclusion   BlendMode = "exclusion"
	BlendModeHue         BlendMode = "hue"
	BlendModeSaturation  BlendMode = "saturation"
	BlendModeColor       BlendMode = "color"
	BlendModeLuminosity  BlendMode = "luminosity"
)

// FillMode represents the fill method
type FillMode string

const (
	FillModeSolid  FillMode = "solid"
	FillModeLinear FillMode = "linear"
	FillModeRadial FillMode = "radial"
)

// Fit represents how content fits within bounds
type Fit string

const (
	FitFill    Fit = "fill"
	FitContain Fit = "contain"
	FitCover   Fit = "cover"
	FitNone    Fit = "none"
)

// TextTransform represents text transformation
type TextTransform string

const (
	TextTransformNone       TextTransform = "none"
	TextTransformUppercase  TextTransform = "uppercase"
	TextTransformLowercase  TextTransform = "lowercase"
	TextTransformCapitalize TextTransform = "capitalize"
)

// StrokeCap represents stroke line cap style
type StrokeCap string

const (
	StrokeCapButt   StrokeCap = "butt"
	StrokeCapRound  StrokeCap = "round"
	StrokeCapSquare StrokeCap = "square"
)

// StrokeJoin represents stroke line join style
type StrokeJoin string

const (
	StrokeJoinMiter StrokeJoin = "miter"
	StrokeJoinRound StrokeJoin = "round"
	StrokeJoinBevel StrokeJoin = "bevel"
)

// FlowDirection represents text flow direction
type FlowDirection string

const (
	FlowDirectionLeftToRight FlowDirection = "left-to-right"
	FlowDirectionRightToLeft FlowDirection = "right-to-left"
)

// MaskMode represents mask mode
type MaskMode string

const (
	MaskModeAlpha      MaskMode = "alpha"
	MaskModeLuminance  MaskMode = "luminance"
	MaskModeInverted   MaskMode = "inverted"
)

// BlurMode represents blur mode
type BlurMode string

const (
	BlurModeGaussian BlurMode = "gaussian"
	BlurModeMotion   BlurMode = "motion"
)

// ColorFilterType represents color filter type
type ColorFilterType string

const (
	ColorFilterTypeGrayscale   ColorFilterType = "grayscale"
	ColorFilterTypeSepia       ColorFilterType = "sepia"
	ColorFilterTypeInvert      ColorFilterType = "invert"
	ColorFilterTypeBrightness  ColorFilterType = "brightness"
	ColorFilterTypeContrast    ColorFilterType = "contrast"
	ColorFilterTypeSaturate    ColorFilterType = "saturate"
	ColorFilterTypeHueRotate   ColorFilterType = "hue-rotate"
)

// TranscriptEffect represents transcript effect
type TranscriptEffect string

const (
	TranscriptEffectHighlight TranscriptEffect = "highlight"
	TranscriptEffectKaraoke   TranscriptEffect = "karaoke"
)

// TranscriptPlacement represents transcript placement
type TranscriptPlacement string

const (
	TranscriptPlacementCenter TranscriptPlacement = "center"
	TranscriptPlacementBottom TranscriptPlacement = "bottom"
)

// TranscriptSplit represents how transcript is split
type TranscriptSplit string

const (
	TranscriptSplitWord  TranscriptSplit = "word"
	TranscriptSplitLine  TranscriptSplit = "line"
	TranscriptSplitNone  TranscriptSplit = "none"
)

// WarpMode represents warp mode
type WarpMode string

const (
	WarpModeArc       WarpMode = "arc"
	WarpModeArch      WarpMode = "arch"
	WarpModeBulge     WarpMode = "bulge"
	WarpModeFlag      WarpMode = "flag"
	WarpModeWave      WarpMode = "wave"
	WarpModeFish      WarpMode = "fish"
	WarpModeRise      WarpMode = "rise"
	WarpModeFisheye   WarpMode = "fisheye"
)

// Easing represents animation easing
type Easing string

const (
	EasingLinear      Easing = "linear"
	EasingEaseIn      Easing = "ease-in"
	EasingEaseOut     Easing = "ease-out"
	EasingEaseInOut   Easing = "ease-in-out"
	EasingEaseInQuad  Easing = "ease-in-quad"
	EasingEaseOutQuad Easing = "ease-out-quad"
	EasingEaseInOutQuad Easing = "ease-in-out-quad"
	// ... add more as needed
)