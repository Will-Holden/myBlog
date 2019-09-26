package main

import (
	"net/http"

	"github.com/rembern/myblog/api/controllers"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	router := gin.Default()
	// Listen and Server in 0.0.0.0:8080

	blogStruct := controllers.BlogStruct{}

	router.LoadHTMLGlob("ui-dist/*.html")
	router.Static("/static", "./ui-dist/static")
	dataUI := gin.H{}
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", dataUI)
	})
	router.GET("/hello", blogStruct.Hello)

	router.Run(":8080")
}
