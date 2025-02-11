package configs

const (
	TimeOutSeconds = 10
	MaxWorkers     = 100
)

// Database response error codes.
const (
	ErrUserNotFound = iota + 1
	ErrUserExists
	ErrInvalidValue
)
