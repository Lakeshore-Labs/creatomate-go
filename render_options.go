package creatomate

type RenderOutputFormat string

const (
	RenderOutputFormatJPG RenderOutputFormat = "jpg"
	RenderOutputFormatPNG RenderOutputFormat = "png"
	RenderOutputFormatGIF RenderOutputFormat = "gif"
	RenderOutputFormatMP4 RenderOutputFormat = "mp4"
)

type RenderOptions struct {
	OutputFormat  RenderOutputFormat     `json:"output_format,omitempty"`
	FrameRate     float64                `json:"frame_rate,omitempty"`
	RenderScale   float64                `json:"render_scale,omitempty"`
	MaxWidth      int                    `json:"max_width,omitempty"`
	MaxHeight     int                    `json:"max_height,omitempty"`
	TemplateID    string                 `json:"template_id,omitempty"`
	Tags          []string               `json:"tags,omitempty"`
	Source        interface{}            `json:"source,omitempty"` // Can be *Source or map[string]interface{}
	Modifications map[string]interface{} `json:"modifications,omitempty"`
	WebhookURL    string                 `json:"webhook_url,omitempty"`
	Metadata      string                 `json:"metadata,omitempty"`
}