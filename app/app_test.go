package main

import (
	"testing"

	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
)

func TestGetIndex(t *testing.T) {
	r := gofight.New()
	r.GET("/").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}

func TestFavicon(t *testing.T) {
	r := gofight.New()
	r.GET("/favicon.ico").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}

func TestStaticText(t *testing.T) {
	r := gofight.New()
	r.GET("/static/test.txt").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}
func TestImageJPG(t *testing.T) {
	r := gofight.New()
	r.GET("/image/jpg").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}
func TestVideoMP4(t *testing.T) {
	r := gofight.New()
	r.GET("/video/mp4").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 200, resp.Code)
	})
}
func TestRedirects(t *testing.T) {
	r := gofight.New()
	r.GET("/redirect/15").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 302, resp.Code)
	})
}
func TestStatus(t *testing.T) {
	r := gofight.New()
	r.GET("/status/333").SetDebug(true).Run(AppEngine(), func(resp gofight.HTTPResponse, req gofight.HTTPRequest) {
		assert.Equal(t, 333, resp.Code)
	})
}
