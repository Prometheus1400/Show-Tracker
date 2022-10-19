package main

import (
	// "net/http"
	"log"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"os"
	// "path/filepath"
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

	router := gin.Default()

	// Serve frontend static files
	router.Use(static.Serve("/", static.LocalFile(buildPath, true)))
	router.Run(":8080")
}