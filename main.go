package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Response struct {
	code    int
	message string
	data    any
}

func responseWithSuccess(c *gin.Context, res Response) {
	response := gin.H{
		"status":  "success",
		"code":    res.code,
		"message": res.message,
	}
	if res.data != nil {
		response["data"] = res.data
	}
	c.JSON(res.code, response)
}
func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		responseWithSuccess(c, Response{code: 200})
	})

	router.Run(":8080")
}
