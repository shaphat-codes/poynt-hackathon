package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"poynt/models"
)

func AddImage(context *gin.Context) {
	var input models.ImageRequest
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image := models.Image{
		Data: input.Data,
	}

	filePath, err := image.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"imagePath": filePath})
}

func ServeImage(context *gin.Context) {
	path := context.Param("path")

	_, err := os.Stat("." + path)
	if os.IsNotExist(err) {
		context.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	context.Header("Content-Type", "image/png")

	context.File("." + path)
}
