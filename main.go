package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	app.Static("template/css", "template/css")
	app.LoadHTMLFiles("template/index.tmpl.html")

	app.GET("/index", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.tmpl.html", gin.H{
			"msg": "Success",
		})
	})
	app.Run(":8000")
}
