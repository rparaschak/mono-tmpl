package apperror

import "errors"

type AppError struct {
	status  int
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func New(status int, code, message string) *AppError {
	return &AppError{
		status:  status,
		Code:    code,
		Message: message,
	}
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) GetStatus() int {
	return e.status
}

func (e *AppError) WithDetails(details any) *AppError {
	clone := *e
	clone.Details = details
	return &clone
}

func (e *AppError) Is(target error) bool {
	var targetAppError *AppError
	if !errors.As(target, &targetAppError) {
		return false
	}
	return e.Code == targetAppError.Code
}

func WithDetails(err *AppError, details any) *AppError {
	if err == nil {
		return nil
	}
	return err.WithDetails(details)
}

func ErrorIsOneOf(err error, errs ...error) bool {
	for _, target := range errs {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}
