package common

import (
	"fmt"
	"time"
)

type ResponseType string

const (
	JsonType ResponseType = "json"
	FileType ResponseType = "file"
)

type ResponseMessage struct {
	Code     int           `json:"code"`
	Message  string        `json:"message"`
	Data     interface{}   `json:"data"`
	TimeUsed time.Duration `json:"time_used"`
	Type     ResponseType  `json:"type"`
}

func NewResponse(code int, message string, data interface{}, timeUsed time.Duration) error {
	r := &ResponseMessage{
		Code:     code,
		Data:     data,
		Message:  message,
		TimeUsed: timeUsed,
		Type:     JsonType,
	}
	panic(r)
}

func NewFileResponse(code int, message string, data interface{}, timeUsed time.Duration) error {
	r := &ResponseMessage{
		Code:     code,
		Data:     data,
		Message:  message,
		TimeUsed: timeUsed,
		Type:     FileType,
	}
	panic(r)
}

func (r *ResponseMessage) Error() string {
	if r.Type == FileType {
		return fmt.Sprintf("code: %d, timeUsed: %v, Message: %s", r.Code, r.TimeUsed, r.Message)
	}
	return fmt.Sprintf("Data: %v, code: %d, timeUsed: %v, Message: %s", r.Data, r.Code, r.TimeUsed, r.Message)
}
