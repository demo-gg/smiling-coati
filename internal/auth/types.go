package types

type SlackResponse struct {
	OK      bool   `json:"ok"`
	Error   string `json:"error,omitempty"`
	Channel string `json:"channel,omitempty"`
	TS      string `json:"ts,omitempty"`
}

type ErrorResponse struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func NewErrorResponse(code int, message string) *ErrorResponse {
	return &ErrorResponse{Code: code, Message: message}
}
