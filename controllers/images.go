package controllers

import (
	"net/http"
	"poynt/models"
	"github.com/gin-gonic/gin"
	
)

func AddImage(context *gin.Context) {
	var input models.Image
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	
	savedImage, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedImage})
}