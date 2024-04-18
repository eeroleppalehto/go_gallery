package imageservice

import (
	"image"
	"image/jpeg"
	"os"
)

func (i *JPG) LoadImage(filePath string) (image.Image, error) {
	// Test if file can be opened
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, err
}

func (i *JPG) SaveImage(img image.Image, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	jpeg.Encode(file, img, nil)

	return nil
}
