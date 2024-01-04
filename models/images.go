package models

import (
	"encoding/base64"
	
	"fmt"
	
	"poynt/database"
	
	
)

type Image struct {
	ID   uint   `gorm:"primary_key" json:"id"`
	Data string `json:"data"`
	ByteData []byte `json:"data"`
	ImageUrl string `json:"imageUrl"`
}

func (image *Image) Save() (*Image, error) {
	fmt.Println("received image", image.Data)
	decodedData, err := base64.StdEncoding.DecodeString(image.Data)
	if err != nil {
		return nil, err
	}

	newImage := Image{
		ByteData: decodedData,
	}

	database.Database.Create(&newImage)
	
	return image, nil
	
}



