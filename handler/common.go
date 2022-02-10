package handler

import (
	"net/http"
	"strconv"

	"gin-demo/action"

	"github.com/gin-gonic/gin"
)

// 获取加法结果
func AddResult(context *gin.Context) {
	a := context.Param("a")
	b := context.Param("b")
	aint, erra := strconv.Atoi(a)
	if erra != nil {
		context.String(http.StatusInternalServerError, "error")
		return
	}
	bint, errb := strconv.Atoi(b)
	if errb != nil {
		context.String(http.StatusInternalServerError, "error")
		return
	}
	context.String(http.StatusOK, "%d", action.Add(aint, bint))
}