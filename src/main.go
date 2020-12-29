package main

import (
	"gogin/src/httpd/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getDefault(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Status": http.StatusOK,
	})
}

func main() {
	r := gin.Default()

	r.GET("/", getDefault)
	r.GET("/person", handler.GetPersonDefault)
	r.GET("/search", handler.GetPersonByID)

	r.Run()
}
