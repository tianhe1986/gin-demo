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

	// 强制转换
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

func MulResult(context *gin.Context) {
	// 创建结构体，用于接收参数
	type MulParam struct {
		A int `json:"a"`
		B int `json:"b"`
	}

	// 接收并绑定参数
	json := MulParam{}
	err := context.ShouldBindJSON(&json)

	// 绑定失败
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"result": "error",
		})
		return
	}

	// 返回结果
	context.JSON(http.StatusOK, gin.H{
		"result": action.Mul(json.A, json.B),
	})
}
