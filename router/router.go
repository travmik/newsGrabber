package router

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.New()

	r.GET("/", indexHandler)
	r.GET("/search/:category", collectHandler)
	r.GET("/result/:category", resultHandler)

	return r
}
