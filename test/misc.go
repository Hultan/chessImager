package test

import (
	"image"
	"image/png"
	"os"
)

func compareFiles(i1 *image.Image, f2 string) (bool, error) {
	i2, err := loadImage(f2)
	if err != nil {
		return false, err
	}

	if !(*i1).Bounds().Eq(i2.Bounds()) {
		return false, nil
	}

	for y := 0; y < (*i1).Bounds().Size().Y; y++ {
		for x := 0; x < (*i1).Bounds().Size().X; x++ {
			if (*i1).At(x, y) != i2.At(x, y) {
				return false, nil
			}
		}
	}

	return true, nil
}

func loadImage(f string) (image.Image, error) {
	file, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
