package model

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/**
 * gin 的上下文
 */
type Context struct {
	Router *gin.Engine
	Logger *logrus.Logger
}

/**
 * 响应体
 */
type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
