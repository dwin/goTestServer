package controller

import (
	"io/ioutil"

	u "github.com/dwin/goTestServer/app/utils"
	"github.com/gin-gonic/gin"
)

// GetImage returns image of selected type ex. /image/jpg returns '.jpg' image
func GetImage(c *gin.Context) {
	switch imgType := c.Param("type"); imgType {
	case "jpeg", "jpg":
		img, err := ioutil.ReadFile(u.StaticBasePath + "/static/img/markus-spiske-221494.jpg")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open jpg for GetImage")
			c.AbortWithError(500, err)
		}
		c.Data(200, "image/jpeg", img)
		return
	case "png":
		img, err := ioutil.ReadFile(u.StaticBasePath + "/static/img/markus-spiske-221494.png")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open jpg for GetImage")
			c.AbortWithError(500, err)
		}
		c.Data(200, "image/png", img)
		return
	case "svg":
		img, err := ioutil.ReadFile(u.StaticBasePath + "/static/img/test.svg")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open svg for GetImage")
			c.AbortWithError(500, err)
		}
		c.Data(200, "image/svg+xml", img)
		return
	case "jp2":
		img, err := ioutil.ReadFile(u.StaticBasePath + "/static/img/markus-spiske-221494.jp2")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open jp2 for GetImage")
			c.AbortWithError(500, err)
		}
		c.Data(200, "image/jpeg", img)
		return
	case "webp":
		img, err := ioutil.ReadFile(u.StaticBasePath + "/static/img/markus-spiske-221494.webp")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open jpg for GetImage")
			c.AbortWithError(500, err)
		}
		c.Data(200, "image/webp", img)
		return
	default:
		u.Log.Warn().Str("Request param:", imgType).Msg("Client requested unsupported image type")
		c.String(400, "Unsupported image type requested")
		return
	}
}

// GetVideo returns video of selected type ex. /image/mp4 returns '.mp4' image
func GetVideo(c *gin.Context) {
	switch vidType := c.Param("type"); vidType {
	case "mp4", "mpeg4", "h264":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_h264.mp4")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open mp4 for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/mp4", vid)
		return
	case "webm", "vp8":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_vp8.webm")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open vp8/webm for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/webm", vid)
		return
	case "vp9":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_vp9.webm")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open vp9/webm for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/webm", vid)
		return
	case "h265":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_h265.mp4")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open h265/mp4 for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/mp4", vid)
		return
	case "theora", "ogv", "ogg":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_theora.ogv")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open theora/ogv for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/ogv", vid)
		return
	case "3gp":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_320x240_2mb.3gp")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open 3gp for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/3gpp", vid)
		return
	case "flv", "flash":
		vid, err := ioutil.ReadFile(u.StaticBasePath + "/static/vid/SampleVideo_640x360_2mb.flv")
		if err != nil {
			u.Log.Error().Err(err).Msg("Could not open flv for GetVideo")
			c.AbortWithError(500, err)
		}
		c.Data(200, "video/mp4", vid)
		return
	default:
		u.Log.Warn().Str("Request param:", vidType).Msg("Client requested unsupported video type")
		c.String(400, "Unsupported video type requested")
		return
	}
}
