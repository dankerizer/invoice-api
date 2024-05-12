package helper

import "fmt"

type RequestError struct {
	StatusCode int    `json:"statusCode"`
	Err        error  `json:"err,omitempty"`
	Message    string `json:"message"`
}

type ResponseError struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: %v", r.StatusCode, r.Err)
}

func (r *RequestError) ToJson() ResponseError {
	return ResponseError{
		StatusCode: r.StatusCode,
		Message:    r.Err.Error(),
	}
}
