package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	_estime "github.com/echoframe/common/estime"
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
	var d interface{}
	if data != nil {
		d = data[0]
	} else {
		d = data
	}
	return &Result{
		Lasting: _estime.FormatTime(time.Now()),
		Status:  status,
		Code:    code,
		Message: message,
		Data:    d,
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
