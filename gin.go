package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/", hello)
	r.Run(":8001")
}

func hello(c *gin.Context)  {
	c.String(http.StatusOK, "Hello, World!")
}
