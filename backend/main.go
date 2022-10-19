package main

import (
	// "net/http"
	"log"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
	// "path/filepath"
)

const (
	defaultPort = "8080"
)

func main() {
	log.Println("Show Tracker App v0.01")
	var buildPath string
	if _, prod := os.LookupEnv("PROD"); prod {
		buildPath = "/app/build"
	} else {
		buildPath = "../frontend/build"
	}
	log.Println("buildpath:", buildPath)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile(buildPath, true)))
	
	router.GET("/test", func(ctx *gin.Context) {
		ctx.Writer.Write([]byte("Test"))
	})

	// catch all appraoch to work with React Router
	router.NoRoute(func(c *gin.Context) {
		c.File(buildPath + "/index.html")
	})
	router.Run(":" + port)
}