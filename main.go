package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
	file, _ := os.ReadFile("recipes.json")
	_ = json.Unmarshal([]byte(file), &recipes)
}

type Response struct {
	code    int    `json:"code"`
	message string `json:"message"`
	data    any    `json:"data"`
}
type Recipe struct {
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublishedAt  time.Time `json:"published_at"`
}

func main() {
	r := gin.Default()
	// r.GET("/", index)
	// r.GET("/:name", show)
	r.GET("/recipes", RecipeIndex)
	r.POST("/recipes", RecipeStore)
	r.Run(":8080")
}

func RecipeIndex(c *gin.Context) {
	responseWithSuccess(c, Response{http.StatusOK, "", recipes})
}
func RecipeStore(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		responseWithError(c, Response{http.StatusBadRequest, "error occurs", err.Error()})
		return
	}
	recipe.PublishedAt = time.Now()
	recipes = append(recipes, recipe)
	responseWithSuccess(c, Response{http.StatusCreated, "", recipe})
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
func responseWithError(c *gin.Context, res Response) {
	response := gin.H{
		"status":  "error",
		"code":    res.code,
		"message": res.message,
	}
	if res.data != nil {
		response["data"] = res.data
	}
	// c.XML(res.code, response)
	c.JSON(res.code, response)
}

// func index(c *gin.Context) {
// 	responseWithSuccess(c, Response{code: 200, message: "hello"})
// }
// func show(c *gin.Context) {
// 	name := c.Params.ByName("name")
// 	responseWithSuccess(c, Response{code: 200, data: name})
// }
