package currencystack

import (
	"fmt"
	"net/http"
)

// HTTPErrorList -
type HTTPErrorList struct {
	Type   string      `json:"type"`
	Errors []HTTPError `json:"errors"`
}

// HTTPError -
type HTTPError struct {
	StatusCode int
	Code       string `json:"code"`
	Message    string `json:"message"`
}

// NewUnknownHTTPError -
func NewUnknownHTTPError(statusCode int) HTTPError {
	message := http.StatusText(statusCode)
	if message == "" {
		message = "Unknown Error"
	}
	return HTTPError{Code: "Unknown", Message: message, StatusCode: statusCode}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf("%d: %s, %s", e.StatusCode, e.Code, e.Message)
}

// GetStatusCode -
func (e HTTPError) GetStatusCode() int {
	return e.StatusCode
}

// GetCode -
func (e HTTPError) GetCode() string {
	return e.Code
}

// GetMessage -
func (e HTTPError) GetMessage() string {
	return e.Message
}
