package main

import (
	"fmt"

	"github.com/eeroleppalehto/go_gallery/models"
	imageservice "github.com/eeroleppalehto/go_gallery/service/imageService"
)

func main() {
	MAX_PIXEL_SIZE := 500_000

	imageHandler := imageservice.ImageService{}

	images := models.GetImages()

	for _, jpgImage := range images {
		img, err := imageHandler.JPG.LoadImage(fmt.Sprintf("static/images/%s", jpgImage.Filename))
		if err != nil {
			fmt.Println("Failed to load image: ", jpgImage.Filename)
			return
		}

		pixel, err := imageHandler.GetNewBounds(img, MAX_PIXEL_SIZE)
		if err != nil {
			fmt.Println("Failed to calculate new bounds: ", jpgImage.Filename)
			return
		}

		resizedImg, err := imageHandler.Resize(img, pixel)
		if err != nil {
			fmt.Println("Failed to conver to new image: ", jpgImage.Filename)
			return
		}

		newFileName := fmt.Sprintf("static/images-lq/image-%d-lq.jpg", jpgImage.ID)

		err = imageHandler.JPG.SaveImage(resizedImg, newFileName)

		if err != nil {
			fmt.Println("Failed to save new image: ", jpgImage.Filename)
			return
		}

		fmt.Printf("Saved %s succesfully!\n", newFileName)
	}

	fmt.Println("----------------------------")
	fmt.Println("Exited succesfully")
}
