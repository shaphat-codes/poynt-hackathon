package consts

import (
	"fmt"
	"os"
)

func SetUpImagesDirectory() {
	dirname := "image"

	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		err := os.Mkdir(dirname, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}
		fmt.Println("Directory created:", dirname)
	} else if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Directory already exists:", dirname)
	}
}
