package utils

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	_estime "github.com/echo-scaffolding/common/estime"
)

var Res *Result

type Result struct {
	Lasting string      `json:"lasting"`
	Status  bool        `json:"status"`
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
		Data:    data,
	}
}

func ToJson(c echo.Context, res *Result) error {
	return c.JSON(http.StatusOK, res)
}

func ToXml(c echo.Context, res *Result) error {
	return c.XML(http.StatusOK, res)
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
