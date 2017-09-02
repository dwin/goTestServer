package main

import (
	"github.com/gin-gonic/gin"

	"github.com/dwin/goTestServer/app/controller"
	u "github.com/dwin/goTestServer/app/utils"
)

var servePort = ":8080"

// TODO: Check all request length, kill connection if size unreasonable for type
func main() {
	R := AppEngine()
	R.Run(servePort) // 0.0.0.0:8080
}
func AppEngine() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.StaticFile("/favicon.ico", "../static/icon.ico")
	router.Static("/static", "../static")
	router.LoadHTMLGlob("view/*")
	router.GET("/", controller.GetIndex)

	j := router.Group("/json")
	{
		j.GET("/cookies", controller.GetCookiesJSON)
		j.GET("get", controller.GetJSON)
		j.GET("/ip", controller.GetIPJSON)
		j.GET("/headers", controller.GetHeadersJSON)
		j.GET("/cookies/set", controller.SetCookieJSON)
		j.GET("/cookies/delete", controller.DeleteCookieJSON)
		j.GET("/user-agent", controller.GetUserAgentJSON)
		j.GET("/uuid/:version", controller.GetUUIDJSON)
	}
	router.GET("/image/:type", controller.GetImage)
	router.GET("/video/:type", controller.GetVideo)
	router.GET("/redirect/:redirects") //TODO: Not working
	router.GET("/status/:status", controller.GetStatus)

	u.Log.Info().Str("Server Port", servePort).Msg("Starting Server")
	//router.Run(servePort) // 0.0.0.0:8080

	return router
}
