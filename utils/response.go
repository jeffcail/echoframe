package utils

import (
	"time"

	_estime "github.com/echo-scaffolding/common/estime"
)

var Res *Result

type Result struct {
	Lasting string      `json:"lasting"`
	Status  bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r *Result) Response(status bool, message string, code int, data ...interface{}) *Result {
	return &Result{
		Lasting: time.Now().Format(_estime.LAYOUT),
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data[0],
	}
}

var ResP *PageList

type PageList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func (p *PageList) Pagination(count int64, list interface{}) *PageList {
	return &PageList{
		Total: count,
		List:  list,
	}
}
