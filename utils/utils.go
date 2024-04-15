package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func FileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func RmTmpDir() error {
	return os.RemoveAll(os.Getenv("TMP_DIR"))
}

func DimensionsImage(img image.Image) (int, int) {
	return img.Bounds().Dx(), img.Bounds().Dy()
}

func DecodeImage(file string) (image.Image, string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, "", err
	}
	defer f.Close()

	img, format, err := image.Decode(f)
	if err != nil {
		return nil, "", err
	}

	return img, format, nil
}

func RotateImage(img image.Image) *image.NRGBA {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	rotated := image.NewNRGBA(image.Rect(0, 0, height, width))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rotated.Set(height-y-1, x, img.At(x, y))
		}
	}

	return rotated
}

func EncodeImage(file string, img image.Image, format string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	switch format {
	case "jpeg":
		return jpeg.Encode(f, img, nil)
	case "png":
		return png.Encode(f, img)
	}

	return fmt.Errorf("Unsupported format %s", format)
}
