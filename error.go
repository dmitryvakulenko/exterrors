package exterrors

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	prev error
	msg  string
	data map[string]interface{}
}

func New(msg string, data map[string]interface{}) *Error {
	return &Error{
		msg:  msg,
		data: data,
	}
}

func Wrap(err error, msg string, data map[string]interface{}) *Error {
	return &Error{
		prev: err,
		msg:  msg,
		data: data,
	}
}

func (e Error) Error() string {
	data, err := json.Marshal(e.data)
	if err != nil {
		return fmt.Sprintf("error %q data marshalling error %v", e.msg, err)
	}

	return e.msg + " " + string(data)
}

func (e Error) Unwrap() error {
	return e.prev
}

func (e Error) GetData() map[string]interface{} {
	return e.data
}
