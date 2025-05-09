package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-file-server-api/internal/file"
	"go-file-server-api/internal/folder"
	mw "go-file-server-api/middleware"
	"log"
	"os"
)

func init() {
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	// Only load the godotenv when running in debug mode
	// But in release mode the .env will be supplied dynamically
	if ginMode == "debug" {
		err := godotenv.Load(".env")
		if err != nil {
			panic("Error loading .env file")
		}
	}

	uploadDir, err := folder.UseUploadDir()
	if err != nil {
		return
	}

	// if file not exists it will create automatically
	if _, err = os.Stat(uploadDir); os.IsNotExist(err) {
		err = os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			return
		}

		log.Println("Created directory", uploadDir)
	}
}

func main() {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// Root endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return
	}

	// Initialize middlewares
	r.Use(mw.SecurityHeaders)

	// Initialize folder module
	folderService := folder.NewService()
	folderController := folder.NewController(folderService)
	folderController.RegisterRoutes(r)

	// Initialize file module
	fileService := file.NewService()
	fileController := file.NewController(fileService)
	fileController.RegisterRoutes(r)

	port := os.Getenv("PORT")
	err = r.Run(port)
	if err != nil {
		return
	}

	log.Println("Server running on port", port)
}
