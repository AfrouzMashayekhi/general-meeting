package main

import (
	"github.com/gin-gonic/gin"
)

var (
	port = ":8080"
)

func main() {
	// Creates a router without any middleware by default
	app := gin.Default()
	app.LoadHTMLGlob("templates/*")
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())
	// Serve static files
	app.Static("/assets", "./assets")
	app.GET("/home", GetHome)
	app.POST("/home", PostHome)
	app.GET("/view/t", GetViewTrader)
	app.GET("/view/c", GetViewCompany)
	app.GET("/register", GetRegister)
	app.GET("/view/t/:user", GetTrader)
	app.GET("/view/c/:company", GetComapny)
	// Listen and serve on 0.0.0.0:8080
	app.Run(port)
}
