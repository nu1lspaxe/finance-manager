package configs

import "fmt"

type ManagerError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e ManagerError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}

func NewManagerError(code int, message string) ManagerError {
	return ManagerError{
		Code:    code,
		Message: message,
	}
}
