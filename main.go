package main

import (
	 "poynt/database"
		
	 "poynt/models"
	
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"github.com/gin-gonic/gin"
	"poynt/controllers"
)

func main () {
	loadEnv()
	loadDatabase()
	serverApplication()
}

 func loadDatabase() {
	
 	database.Connect()
	 database.Database.AutoMigrate(&models.Image{})
	
 }

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
}

func serverApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/image")
	publicRoutes.POST("/upload", controllers.AddImage)
	
	router.Run(":8000")
	fmt.Println("server running on port 8000")
}

