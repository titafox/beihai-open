package utils

import (
	"coltstail-server/ecodes"
	"encoding/json"
)

type Error struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *Error) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}

func NewError(code int) *Error {
	return &Error{Code: code, Msg: ecodes.ErrorMsg(code)}
}
