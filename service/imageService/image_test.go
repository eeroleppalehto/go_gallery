package imageservice_test

import (
	"image"
	"image/color"
	"math"
	"testing"

	imageservice "github.com/eeroleppalehto/go_gallery/service/imageService"
)

func TestGetNewBounds(t *testing.T) {
	expected := 500

	imgServ := imageservice.ImageService{}

	img := GenerateTestImage(t)

	p, err := imgServ.GetNewBounds(img, 250_000)
	if err != nil {
		t.Fatalf("Error while calculating new bounds: %s", err)
	}

	if p.X != expected {
		t.Fatalf("Expected %d, got %d", expected, p.X)
	}
}

func GenerateTestImage(t testing.TB) image.Image {
	size := 1_000
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	for i := range size {
		for j := range size {
			c := uint8(math.Floor(float64(256) * (float64(j) / float64(size))))

			rgba := color.RGBA{
				R: c,
				B: c,
				G: c,
				A: 0,
			}
			img.SetRGBA(i, j, rgba)
		}
	}
	return img
}
