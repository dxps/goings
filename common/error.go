package common

//
// Error is a common error, used in API responses.
//
type ApiError struct {

	// The text of the error.
	ErrorText string `json:"error_text"`

	// The code of the error.
	ErrorCode string `json:"error_code"`
}

//
// NewApiError is creating a new ApiError instance.
//
func NewApiError(errorText string) ApiError {
	return ApiError{ErrorText: errorText, ErrorCode: "n/a"}
}
