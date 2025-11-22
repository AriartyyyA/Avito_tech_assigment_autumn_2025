package models

type ErrorCode string

const (
	ErrorCodeTeamExists  ErrorCode = "TEAM_EXISTS"
	ErrorCodePRExists    ErrorCode = "PR_EXISTS"
	ErrorCodePRMerged    ErrorCode = "PR_MERGED"
	ErrorCodeNotAssigned ErrorCode = "NOT_ASSIGNED"
	ErrorCodeNoCandidate ErrorCode = "NO_CANDIDATE"

	ErrorCodeNotFound       ErrorCode = "NOT_FOUND"
	ErrorCodeInvalidRequest ErrorCode = "INVALID_REQUEST"
	ErrorCodePRNotFound     ErrorCode = "PR_NOT_FOUND"
	ErrorCodeUserNotFound   ErrorCode = "USER_NOT_FOUND"
	ErrorCodeTeamNotFound   ErrorCode = "TEAM_NOT_FOUND"
	ErrorCodeInternal       ErrorCode = "INTERNAL"
)

func (e ErrorCode) Error() string {
	return string(e)
}

type ErrorDetail struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

func NewErrorDetail(code ErrorCode, message string) ErrorDetail {
	return ErrorDetail{
		Code:    code,
		Message: message,
	}
}
