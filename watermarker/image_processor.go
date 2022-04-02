package watermarker

import (
	"errors"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
)

func ReadImageFile(inputFile string) image.Image {
	image, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Failed to open: %s", err)
	}
	defer image.Close()

	decoded, err := png.Decode(image)
	if err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return decoded
}

func WriteImageToFile(outImagePath string, outImage image.Image) {
	if _, err := os.Stat(outImagePath); errors.Is(err, os.ErrNotExist) {
		dir := filepath.Dir(outImagePath)
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
		}
	}
	outFile, err := os.Create(outImagePath)
	defer outFile.Close()
	if err != nil {
		log.Fatalf("Failed to create: %s", err)
	}
	png.Encode(outFile, outImage)
}

func BlendImageWithWaterMark(baseImage image.Image, qrImage image.Image) image.Image {
	// Get x y of baseImage
	bounds := baseImage.Bounds()
	qrBounds := qrImage.Bounds()
	outImage := image.NewRGBA(bounds)

	// Draw with bounds of baseImage
	draw.Draw(outImage, bounds, baseImage, image.ZP, draw.Src)
	draw.Draw(outImage, qrBounds, qrImage, image.ZP, draw.Src)
	return outImage
}
