package main

import (
	"fmt"
	"time"

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
	//router.Run(servePort) // 0.0.0.0:8080
	R.Run(servePort) // 0.0.0.0:8080
}

// AppEngine contains all routes for the app
func AppEngine() *gin.Engine {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	mode := gin.Mode()
	fmt.Println(mode)
	if gin.Mode() == "release" {
		// See '/utils/log.go' for production logging settings
		u.Environment = "production"
		router.StaticFile("/favicon.ico", "/static/img/icon.ico")
		router.Static("/static", "/static")
		router.LoadHTMLGlob("/view/*")
	} else if gin.Mode() == "debug" {
		u.Log = zerolog.New(&lumberjack.Logger{
			Filename:   "../log/error.log",
			MaxSize:    50, // megabytes
			MaxBackups: 30,
			MaxAge:     90, //days
		}).With().Timestamp().Logger()
		router.StaticFile("/favicon.ico", "../static/img/icon.ico")
		router.Static("/static", "../static")
		router.LoadHTMLGlob("../view/*")
	}

	router.GET("/", controller.GetIndex)
	t := router.Group("/token")
	{
		t.GET("/new", controller.GetNewToken)
		t.GET("/view/:token", controller.GetListTokenRequests)
	}
	j := router.Group("/json")
	{
		j.GET("/get", controller.GetJSON)
		j.Any("/any", controller.AnyJSON)
		j.GET("/cookies", controller.GetCookiesJSON)
		j.GET("/cookies/set", controller.SetCookieJSON)
		j.GET("/cookies/delete", controller.DeleteCookieJSON)
		j.GET("/datetime", controller.GetDateTimeJSON)
		j.GET("/headers", controller.GetHeadersJSON)
		j.GET("/ip", controller.GetIPJSON)
		j.GET("/status", controller.GetStatusJSON)
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
	router.GET("/redirect/:num", controller.GetRedirects)
	router.GET("/status/:status", controller.GetStatus)

	u.Log.Info().Str("Current Time", u.StartTime.Format(time.UnixDate)).Msg("App start timestamp")
	u.Log.Info().Str("Server Port", servePort).Msg("Starting Server")

	return router
}
