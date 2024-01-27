package chessImager

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func compareImages(t *testing.T, filename string, img *image.Image) {
	ok, err := compareFiles(img, "test/valid/"+filename)
	if err != nil {
		t.Errorf("error during compare : %v", err)
	}
	if !ok {
		t.Errorf("failed to compare, images differ!")
	}
}

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

func saveImage(filename string, img image.Image) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
