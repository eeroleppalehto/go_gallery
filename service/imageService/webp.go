package imageservice

// import (
// 	"image"
// 	"image/webp"
// 	"github.com/kolesa-team/go-webp/encoder"
// 	"github.com/kolesa-team/go-webp/webp"
// 	"os"
// )

// func (i *WEBP) LoadImage(filePath string) (image.Image, error) {
// 	// Test if file can be opened
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer file.Close()

// 	img, _, err := image.Decode(file)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return img, err
// }

// func (i *WEBP) SaveImage(img image.Image, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()
// 	webp.Encode(file, img, nil)

// 	return nil
// }
