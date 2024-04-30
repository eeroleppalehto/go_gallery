package imageservice

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"os"
)

func (i *JPG) LoadImageFromFile(filePath string) (image.Image, error) {
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

func (i *JPG) LoadImageFromReader(r io.Reader) (image.Image, error) {
	img, format, err := image.Decode(r)
	if err != nil {
		return nil, err
	}

	fmt.Println("Format: ", format)

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
