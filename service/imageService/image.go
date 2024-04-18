package imageservice

import (
	"image"
	"image/color"
	"math"
)

type JPG struct{}
type WEBP struct{}

type ImageService struct {
	JPG  JPG
	WEBP WEBP
}

// func LoadImage(filePath string) (image.Image, error) {
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

// func SaveImage(img image.Image, filePath string) error {
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}

// 	defer file.Close()
// 	jpeg.Encode(file, img, nil)

// 	return nil
// }

func (i *ImageService) GetNewBounds(img image.Image, pixels int) (image.Point, error) {
	X := img.Bounds().Size().X
	Y := img.Bounds().Size().Y
	currentPixels := X * Y

	if currentPixels <= pixels {
		return image.Point{X, Y}, nil
	}

	aspectRatio := float64(X) / float64(Y)

	/*
		The Formulas below are bassed on these:

		aspectRatio == newX/newY
		newX == aspectRatio * newY
		newY == newX / aspectRatio


		newX * newY == NewPixels
		newX == NewPixels / newY

		newX == NewPixels * aspectRatio / newX		| *newX
		newX^2 == NewPixels * aspectRatio
		newX == sqrt(NewPixels*aspectRatio)
	*/

	newX := math.Sqrt(float64(float64(pixels) * aspectRatio))
	newY := newX / aspectRatio

	return image.Point{
		X: int(newX),
		Y: int(newY),
	}, nil

}

func (i *ImageService) Resize(img image.Image, newPoint image.Point) (image.Image, error) {
	scalefactor := float64(img.Bounds().Size().X) / float64(newPoint.X)
	ResizedImg := image.NewRGBA(image.Rect(0, 0, newPoint.X, newPoint.Y))

	// The following algorithm uses Bilinear Interpolation
	// to resize the image. The algorithm is based on the
	//
	//
	// 	Visual repesantation of values used below
	//  v1----------q2---------v2
	// 	|			|			|
	// 	|			q			|
	// 	|			|			|
	//  v3----------q1---------v4
	//
	// 	Where v1, v2, v3, v4 are the pixels of the original image
	// 	and q is the pixel of the new image
	//
	// 	For each pixel in the new image, the algorithm calculates
	// 	the pixel value by interpolating the values of the pixels
	// 	around it.
	//
	// 	For example, to calculate the value of pixel q, the algorithm
	// 	reads the values of the pixels v1, v2, v3, v4 and then calculates
	// 	the value of q1 and q2 by interpolating the values of v1, v2 and v3, v4.
	// 	Then it calculates the value of q by interpolating the values of q1 and q2.
	//
	for i := 0; i < newPoint.X; i++ {
		for j := 0; j < newPoint.Y; j++ {
			x := float64(i) * scalefactor
			y := float64(j) * scalefactor

			xFloor := math.Floor(x)
			xCeil := min(math.Ceil(x), float64(img.Bounds().Size().X-1))
			yFloor := math.Floor(y)
			yCeil := min(math.Ceil(y), float64(img.Bounds().Size().Y-1))

			var q color.RGBA
			var r, g, b, a uint32

			if (xCeil == xFloor) && (yCeil == yFloor) {
				r, g, b, a := img.At(int(x), int(y)).RGBA()

				q = color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: uint8(a),
				}

			} else if xCeil == xFloor { //
				q1 := img.At(int(x), int(yFloor))
				q2 := img.At(int(x), int(yCeil))

				r1, g1, b1, a1 := q1.RGBA()
				r2, g2, b2, a2 := q2.RGBA()

				r = uint32(float64(r1)*(yCeil-y) + float64(r2)*(y-yFloor))
				g = uint32(float64(g1)*(yCeil-y) + float64(g2)*(y-yFloor))
				b = uint32(float64(b1)*(yCeil-y) + float64(b2)*(y-yFloor))
				a = uint32(float64(a1)*(yCeil-y) + float64(a2)*(y-yFloor))

				q = color.RGBA{
					R: uint8(r / 256),
					G: uint8(g / 256),
					B: uint8(b / 256),
					A: uint8(a / 256),
				}
			} else if yCeil == yFloor {
				q1 := img.At(int(xFloor), int(y))
				q2 := img.At(int(xCeil), int(y))

				r1, g1, b1, a1 := q1.RGBA()
				r2, g2, b2, a2 := q2.RGBA()

				r = uint32(float64(r1)*(xCeil-x) + float64(r2)*(x-xFloor))
				g = uint32(float64(g1)*(xCeil-x) + float64(g2)*(x-xFloor))
				b = uint32(float64(b1)*(xCeil-x) + float64(b2)*(x-xFloor))
				a = uint32(float64(a1)*(xCeil-x) + float64(a2)*(x-xFloor))

				q = color.RGBA{
					R: uint8(r / 256),
					G: uint8(g / 256),
					B: uint8(b / 256),
					A: uint8(a / 256),
				}
			} else {
				v1 := img.At(int(xFloor), int(yFloor))
				v2 := img.At(int(xCeil), int(yFloor))
				v3 := img.At(int(xFloor), int(yCeil))
				v4 := img.At(int(xCeil), int(yCeil))

				rv1, gv1, bv1, av1 := v1.RGBA()
				rv2, gv2, bv2, av2 := v2.RGBA()
				rv3, gv3, bv3, av3 := v3.RGBA()
				rv4, gv4, bv4, av4 := v4.RGBA()

				r1 := uint32(float64(rv1)*(xCeil-x) + float64(rv2)*(x-xFloor))
				g1 := uint32(float64(gv1)*(xCeil-x) + float64(gv2)*(x-xFloor))
				b1 := uint32(float64(bv1)*(xCeil-x) + float64(bv2)*(x-xFloor))
				a1 := uint32(float64(av1)*(xCeil-x) + float64(av2)*(x-xFloor))

				r2 := uint32(float64(rv3)*(xCeil-x) + float64(rv4)*(x-xFloor))
				g2 := uint32(float64(gv3)*(xCeil-x) + float64(gv4)*(x-xFloor))
				b2 := uint32(float64(bv3)*(xCeil-x) + float64(bv4)*(x-xFloor))
				a2 := uint32(float64(av3)*(xCeil-x) + float64(av4)*(x-xFloor))

				r := uint32(float64(r1)*(yCeil-y) + float64(r2)*(y-yFloor))
				g := uint32(float64(g1)*(yCeil-y) + float64(g2)*(y-yFloor))
				b := uint32(float64(b1)*(yCeil-y) + float64(b2)*(y-yFloor))
				a := uint32(float64(a1)*(yCeil-y) + float64(a2)*(y-yFloor))

				q = color.RGBA{
					R: uint8(r / 256),
					G: uint8(g / 256),
					B: uint8(b / 256),
					A: uint8(a / 256),
				}
			}

			ResizedImg.SetRGBA(i, j, q)
		}
	}
	return ResizedImg, nil
}

/*
func main() {
	log.Println("Start")
	fmt.Println("Loading img...")
	originaImg, err := LoadImage("/home/el/Documents/go_gallery/image-1.jpg")
	if err != nil {
		fmt.Println("Couldn't load image")
		return
	}

	fmt.Println("Calculating new bounds")
	point, err := GetNewBounds(originaImg, 1_000_000)
	if err != nil {
		fmt.Println("Error while calculating new bounds")
		return
	}

	fmt.Println("Resizing image")
	newImg, err := Resize(originaImg, point)
	if err != nil {
		fmt.Println("Error while resizing image: ", err)
		return
	}

	fmt.Println("Saving image")
	SaveImage(newImg, "/home/el/Documents/go_gallery/test.jpg")
	fmt.Println("Done!")
	log.Println("End")
}
*/
