package creatomate

type RenderStatus string

const (
	RenderStatusPlanned      RenderStatus = "planned"
	RenderStatusWaiting      RenderStatus = "waiting"
	RenderStatusTranscribing RenderStatus = "transcribing"
	RenderStatusRendering    RenderStatus = "rendering"
	RenderStatusSucceeded    RenderStatus = "succeeded"
	RenderStatusFailed       RenderStatus = "failed"
)

type Render struct {
	ID            string                 `json:"id"`
	Status        RenderStatus           `json:"status"`
	ErrorMessage  string                 `json:"error_message,omitempty"`
	URL           string                 `json:"url"`
	SnapshotURL   string                 `json:"snapshot_url,omitempty"`
	TemplateID    string                 `json:"template_id,omitempty"`
	TemplateName  string                 `json:"template_name,omitempty"`
	TemplateTags  []string               `json:"template_tags,omitempty"`
	OutputFormat  string                 `json:"output_format"`
	RenderScale   float64                `json:"render_scale"`
	Width         int                    `json:"width,omitempty"`
	Height        int                    `json:"height,omitempty"`
	FrameRate     float64                `json:"frame_rate,omitempty"`
	Duration      float64                `json:"duration,omitempty"`
	FileSize      int64                  `json:"file_size,omitempty"`
	Modifications map[string]interface{} `json:"modifications,omitempty"`
	WebhookURL    string                 `json:"webhook_url,omitempty"`
	Metadata      string                 `json:"metadata,omitempty"`
}