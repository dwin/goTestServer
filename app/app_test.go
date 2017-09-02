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
