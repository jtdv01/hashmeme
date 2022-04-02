package watermarker

import (
	"image"
	"log"
	"os"
	"image/png"
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

// func BlendImageWithWaterMark(image.Image, ) {
//
// }