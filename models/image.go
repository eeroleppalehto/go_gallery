package models

type Image struct {
	ID          int
	Title       string
	Description string
	Filename    string
}

func GetImages() []Image {
	images := []Image{
		{ID: 1, Title: "Image 1", Description: "The first image", Filename: "image-1.jpg"},
		{ID: 2, Title: "Image 2", Description: "The second image", Filename: "image-2.jpg"},
		{ID: 3, Title: "Image 3", Description: "The third image", Filename: "image-3.jpg"},
		{ID: 4, Title: "Image 4", Description: "The fourth image", Filename: "image-4.jpg"},
		{ID: 5, Title: "Image 5", Description: "The fifth image", Filename: "image-5.jpg"},
		{ID: 6, Title: "Image 6", Description: "The sixth image", Filename: "image-6.jpg"},
		{ID: 7, Title: "Image 7", Description: "The seventh image", Filename: "image-7.jpg"},
		{ID: 8, Title: "Image 8", Description: "The eighth image", Filename: "image-8.jpg"},
		{ID: 9, Title: "Image 9", Description: "The ninth image", Filename: "image-9.jpg"},
		{ID: 10, Title: "Image 10", Description: "The tenth image", Filename: "image-10.jpg"},
		{ID: 11, Title: "Image 11", Description: "The eleventh image", Filename: "image-11.jpg"},
		{ID: 12, Title: "Image 12", Description: "The twelfth image", Filename: "image-12.jpg"},
		{ID: 13, Title: "Image 13", Description: "The thirteenth image", Filename: "image-13.jpg"},
		{ID: 14, Title: "Image 14", Description: "The fourteenth image", Filename: "image-14.jpg"},
		{ID: 15, Title: "Image 15", Description: "The fifteenth image", Filename: "image-15.jpg"},
		{ID: 16, Title: "Image 16", Description: "The sixteenth image", Filename: "image-16.jpg"},
	}

	return images
}
