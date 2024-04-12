package main

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
)

type Response struct {
	code    int    `json:"code"`
	message string `json:"message"`
	data    any    `json:"data"`
}

func main() {
	router := gin.Default()
	router.GET("/", index)
	router.GET("/:name", show)
	router.Run(":8080")
}
func index(c *gin.Context) {
	responseWithSuccess(c, Response{code: 200, message: "hello"})
}
func show(c *gin.Context) {
	name := c.Params.ByName("name")
	responseWithSuccess(c, Response{code: 200, data: name})
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
	// c.XML(res.code, response)
	c.JSON(res.code, response)
}
