package creatomate

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	ClientVersion = "1.0.0"
	BaseURL       = "https://api.creatomate.com/v1"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Render starts a new render and awaits its completion.
func (c *Client) Render(ctx context.Context, options RenderOptions, timeout time.Duration) ([]Render, error) {
	if timeout == 0 {
		timeout = 15 * time.Minute
	}
	if timeout > 60*time.Minute {
		timeout = 60 * time.Minute
	}

	renders, err := c.StartRender(ctx, options)
	if err != nil {
		return nil, err
	}

	startTime := time.Now()
	unfinishedRenders := make([]Render, len(renders))
	copy(unfinishedRenders, renders)
	finishedRenders := make([]Render, 0, len(renders))

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for len(unfinishedRenders) > 0 {
		select {
		case <-ctx.Done():
			return finishedRenders, ctx.Err()
		case <-ticker.C:
			if time.Since(startTime) >= timeout {
				return finishedRenders, NewTimeoutError()
			}

			for i := len(unfinishedRenders) - 1; i >= 0; i-- {
				render := unfinishedRenders[i]
				updatedRender, err := c.FetchRender(ctx, render.ID)
				if err != nil {
					return finishedRenders, err
				}

				if !isRenderInProgress(updatedRender.Status) {
					unfinishedRenders = append(unfinishedRenders[:i], unfinishedRenders[i+1:]...)
					finishedRenders = append(finishedRenders, *updatedRender)
				}
			}
		}
	}

	return finishedRenders, nil
}

// StartRender starts a render, but doesn't wait for it to finish.
func (c *Client) StartRender(ctx context.Context, options RenderOptions) ([]Render, error) {
	payload := transformObjectKeysToSnake(options)
	
	if options.Source != nil {
		if source, ok := options.Source.(*Source); ok {
			payload["source"] = source.ToMap()
		} else {
			payload["source"] = options.Source
		}
	}

	var renders []Render
	err := c.httpRequest(ctx, "POST", "/renders", payload, &renders)
	return renders, err
}

// FetchRender fetches the status of the render.
func (c *Client) FetchRender(ctx context.Context, id string) (*Render, error) {
	var render Render
	err := c.httpRequest(ctx, "GET", fmt.Sprintf("/renders/%s", id), nil, &render)
	if err != nil {
		return nil, err
	}
	return &render, nil
}

func (c *Client) httpRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	req, err := http.NewRequestWithContext(ctx, method, BaseURL+path, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("User-Agent", "Creatomate-Go/"+ClientVersion)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return NewConnectionError()
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode >= 400 {
		return c.transformError(resp.StatusCode, respBody)
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w", err)
		}
	}

	return nil
}

func (c *Client) transformError(statusCode int, body []byte) error {
	var errorData struct {
		Hint string `json:"hint"`
	}
	json.Unmarshal(body, &errorData)

	switch statusCode {
	case 400:
		return NewBadRequestError(errorData.Hint)
	case 401:
		return NewInvalidApiKeyError()
	case 402:
		return NewInsufficientCreditsError()
	case 429:
		return NewRateLimitExceededError()
	default:
		return NewCreatomateError(errorData.Hint)
	}
}

func isRenderInProgress(status RenderStatus) bool {
	return status == RenderStatusPlanned ||
		status == RenderStatusWaiting ||
		status == RenderStatusTranscribing ||
		status == RenderStatusRendering
}

// transformObjectKeysToSnake converts map keys from camelCase to snake_case
func transformObjectKeysToSnake(obj interface{}) map[string]interface{} {
	// Simple implementation - in production would use reflection
	result := make(map[string]interface{})
	
	// Convert struct to JSON then back to map
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return result
	}
	
	if err := json.Unmarshal(jsonBytes, &result); err != nil {
		return result
	}
	
	return result
}