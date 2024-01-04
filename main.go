package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
	"poynt/controllers"
)

func main() {
	SetUpImagesDirectory()
	serverApplication()
}

func SetUpImagesDirectory() {
	path := "images"

	dirname := filepath.Join(".", path)

	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			panic("Could not create directory to hold the image.")
		}
		fmt.Println("Directory created:", dirname)
	} else if err != nil {
		fmt.Println("Error:", err)
		panic("Could not check if image directory exists")
	} else {
		fmt.Println("Directory already exists:", dirname)
	}
}

func serverApplication() {
	router := gin.Default()
	publicRoutes := router.Group("/image/")
	publicRoutes.POST("upload", controllers.AddImage)
	publicRoutes.GET("*path", controllers.ServeImage)

	router.Run(":8000")
	fmt.Println("server running on port 8000")
}
