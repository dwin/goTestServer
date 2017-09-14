package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"

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
	mode := gin.Mode()
	fmt.Println(mode)
	if gin.Mode() == "release" {
		router.StaticFile("/favicon.ico", "/static/icon.ico")
		router.Static("/static", "/static")
		router.LoadHTMLGlob("/static/view/*")
	} else if gin.Mode() == "debug" {
		u.Log = zerolog.New(&lumberjack.Logger{
			Filename:   "../log/error.log",
			MaxSize:    50, // megabytes
			MaxBackups: 30,
			MaxAge:     90, //days
		}).With().Timestamp().Logger()
		router.StaticFile("/favicon.ico", "../static/icon.ico")
		router.Static("/static", "../static")
		router.LoadHTMLGlob("view/*")
	}
	//router.StaticFile("/favicon.ico", "/static/icon.ico")
	//router.Static("/static", "/static")

	router.GET("/", controller.GetIndex)

	j := router.Group("/json")
	{
		j.GET("/cookies", controller.GetCookiesJSON)
		j.GET("", controller.GetJSON)
		j.GET("/ip", controller.GetIPJSON)
		j.GET("/headers", controller.GetHeadersJSON)
		j.GET("/cookies/set", controller.SetCookieJSON)
		j.GET("/cookies/delete", controller.DeleteCookieJSON)
		j.GET("/user-agent", controller.GetUserAgentJSON)
		j.GET("/uuid/:version", controller.GetUUIDJSON)
		j.POST("/base32", controller.PostBase32JSON)
		j.POST("/base64", controller.PostBase64JSON)
		j.POST("/blake2b/:size", controller.PostBlake2bJSON)
		j.POST("/md5", controller.PostMD5JSON)
		j.POST("/sha1", controller.PostSHA1JSON)
		j.POST("/sha256", controller.PostSHA256JSON)
		j.POST("/sha3/:size", controller.PostSHA3JSON)
	}
	router.GET("/image/:type", controller.GetImage)
	router.GET("/video/:type", controller.GetVideo)
	router.GET("/redirect/:redirects") //TODO: Not working
	router.GET("/status/:status", controller.GetStatus)

	u.Log.Info().Str("Server Port", servePort).Msg("Starting Server")
	//router.Run(servePort) // 0.0.0.0:8080

	return router
}
