package models

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"os"
)

type ImageRequest struct {
	Data string `json:"data"`
}

type Image struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Data     string `json:"data"`
	ByteData []byte `json:"data"`
	ImageUrl string `json:"imageUrl"`
}

func (image *Image) Save() (string, error) {
	decodedData, err := base64.StdEncoding.DecodeString(image.Data)
	if err != nil {
		return "", err
	}

	uniqueFileName := uuid.New().String() + ".png"

	filePath := "image/" + uniqueFileName

	if err := os.WriteFile("./"+filePath, decodedData, 0644); err != nil {
		fmt.Println("Error writing to file:", err)
		return "", err
	}

	return filePath, nil

}
