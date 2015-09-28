package errors

import (
	"fmt"
)

type AtgoError interface {
	GetMessage() string
	GetInner() string
	Error() string
}

type AtgoBaseError struct {
	Msg   string
	inner error
}

func GetMessage(err interface{}) string {
	switch e := err.(type) {
	case AtgoError:

	case runtime.error:

	default:

	}
}
