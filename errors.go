package creatomate

import "fmt"

type CreatomateError struct {
	Message string
}

func (e *CreatomateError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return "Creatomate API error"
}

func NewCreatomateError(message string) *CreatomateError {
	return &CreatomateError{Message: message}
}

type BadRequestError struct {
	CreatomateError
}

func NewBadRequestError(hint string) *BadRequestError {
	return &BadRequestError{
		CreatomateError: CreatomateError{Message: fmt.Sprintf("Bad request: %s", hint)},
	}
}

type InvalidApiKeyError struct {
	CreatomateError
}

func NewInvalidApiKeyError() *InvalidApiKeyError {
	return &InvalidApiKeyError{
		CreatomateError: CreatomateError{Message: "Invalid API key"},
	}
}

type InsufficientCreditsError struct {
	CreatomateError
}

func NewInsufficientCreditsError() *InsufficientCreditsError {
	return &InsufficientCreditsError{
		CreatomateError: CreatomateError{Message: "Insufficient credits"},
	}
}

type RateLimitExceededError struct {
	CreatomateError
}

func NewRateLimitExceededError() *RateLimitExceededError {
	return &RateLimitExceededError{
		CreatomateError: CreatomateError{Message: "Rate limit exceeded"},
	}
}

type ConnectionError struct {
	CreatomateError
}

func NewConnectionError() *ConnectionError {
	return &ConnectionError{
		CreatomateError: CreatomateError{Message: "Connection error"},
	}
}

type TimeoutError struct {
	CreatomateError
}

func NewTimeoutError() *TimeoutError {
	return &TimeoutError{
		CreatomateError: CreatomateError{Message: "Timeout waiting for render to complete"},
	}
}