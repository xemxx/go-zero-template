package types

type ErrorCode int

const (
	Success ErrorCode = 0
	Failed  ErrorCode = -1
)

const (
	SuccessMessage string = "success"
	FailedMessage  string = "failed"
)
