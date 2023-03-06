package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huanghe314/gin-html/internal"
)

func main() {

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		internal.R.JSON(c.Writer, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
	})

	router.GET("/html", func(c *gin.Context) {
		internal.R.HTML(c.Writer, http.StatusOK, "example", "World")
	})

	router.Run("127.0.0.1:8080")
}
