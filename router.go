package main

import (
	"github.com/EricWvi/atto/handler/login"
	"github.com/EricWvi/atto/handler/ping"
	"github.com/EricWvi/atto/handler/songs"
	"github.com/EricWvi/atto/middleware"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	g.Use(gzip.Gzip(gzip.DefaultCompression))

	g.NoRoute(func(c *gin.Context) {
		c.File(viper.GetString("route.front.index"))
	})

	dir := viper.GetString("route.front.dir")
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				p := strings.TrimPrefix(path, dir)
				g.StaticFile(p, path)
			}
			return nil
		})
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	g.Use(middleware.Logging)
	g.GET("/ping", ping.DefaultHandler)
	g.POST("/login", login.DefaultHandler)
	auth := g.Group(viper.GetString("route.back.base"))
	auth.Use(middleware.JWT)
	{
		auth.POST("/songs", songs.DefaultHandler)
	}

	return g
}
