package api

import (
	"fmt"
	"net/http"
)

type Error struct {
	Message    string
	StatusCode int
}

func (e Error) Error() string {
	return fmt.Sprintf(`{"status": %d, "message": "%s"}`, e.StatusCode, e.Message)
}

func (e Error) Response(w http.ResponseWriter) {
	Response(w, e.StatusCode, StdResponse{
		Message: e.Message,
	})
}

func NewError(code int, message string, a ...any) Error {
	return Error{StatusCode: code, Message: fmt.Sprintf(message, a...)}
}
