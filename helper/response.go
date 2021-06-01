package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	resp := Response{
		Status:  status,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
	return resp
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	resp := Response{Status: false, Message: message,
		Errors: splittedError,
		Data:   data,
	}
	return resp
}
