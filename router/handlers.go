package router

import (
	"net/http"
	"github.com/gin-gonic/gin"

	"go-grabber/news"
)

func indexHandler(c *gin.Context){
	c.String(http.StatusOK, "Hello world!")
}

func collectHandler(c *gin.Context){
	category := c.Param("category")

	news.Collect(category)
	c.String(http.StatusOK, "Search is initialized")
}

func resultHandler(c *gin.Context){
	category := c.Param("category")

	articles := news.Result(category)
	c.JSON(http.StatusOK, articles)
}