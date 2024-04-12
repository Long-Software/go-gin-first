package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Response struct {
	code    int    `json:"code"`
	message string `json:"message"`
	data    any    `json:"data"`
}

type Recipe struct {
	name         string    `json:"name"`
	tags         []string  `json:"tags"`
	ingredients  []string  `json:"ingredients"`
	instructions []string  `json:"instructions"`
	published_at time.Time `json:"published_at"`
}

func main() {
	r := gin.Default()
	r.GET("/", index)
	r.GET("/:name", show)
	r.Run(":8080")
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
