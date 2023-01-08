package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type OkRsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg" example:"成功"`
	Data interface{} `json:"data"`
}

type errResponse struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Cause string `json:"cause"`
}

type Page struct {
	Index     int   `json:"page_index" form:"page_index" example:"1"`
	Size      int   `json:"page_size" form:"page_size" example:"1"`
	TotalPage int   `json:"total_page" example:"1"`
	Total     int64 `json:"total" example:"1"`
}

func ErrJson(c *gin.Context, err error) {
	response := new(errResponse)
	response.Code, response.Msg = 1_00_01, "系统异常"
	response.Cause = err.Error()
	c.JSON(http.StatusOK, response)
	return
}

func OkJson(c *gin.Context, data interface{}) {
	response := &OkRsp{
		Code: 0,
		Msg:  "成功",
		Data: struct{}{},
	}
	if data != nil {
		response.Data = data
	}
	c.JSON(http.StatusOK, response)
}
