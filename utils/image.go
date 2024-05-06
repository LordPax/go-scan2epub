package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path"

	"golang.org/x/image/webp"
)

// Returns the dimensions of an image
func DimensionsImage(img image.Image) (int, int) {
	return img.Bounds().Dx(), img.Bounds().Dy()
}

// Decodes an image from a file
func DecodeImage(file string) (image.Image, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	if ext := path.Ext(file); ext == ".webp" {
		img, err := webp.Decode(f)
		if err != nil {
			return nil, "", err
		}
		return img, "webp", nil
	}

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

// Rotates an image 90 degrees
func RotateImage(img image.Image) *image.NRGBA {
	width, height := DimensionsImage(img)

	rotated := image.NewNRGBA(image.Rect(0, 0, height, width))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rotated.Set(height-y-1, x, img.At(x, y))
		}
	}

	return rotated
}

// Encodes an image to a file
func EncodeImage(file string, img image.Image, format string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	switch format {
	case "jpeg", "jpg":
		return jpeg.Encode(f, img, nil)
	case "png":
		return png.Encode(f, img)
	}

	return fmt.Errorf("Unsupported format %s", format)
}
