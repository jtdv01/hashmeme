package watermarker

import (
    "github.com/nfnt/resize"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
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

func WriteImageFile(outImagePath string, outImage image.Image) {
	outFile, err := os.Create(outImagePath)
	defer outFile.Close()
	if err != nil {
		log.Fatalf("Failed to create: %s", err)
	}
	png.Encode(outFile, outImage)
}

func PrintImage(fullImg image.Image) {
    // Resizes the image so it's easier to read the test logs
    // used only for preview
	resizedImg := resize.Resize(75, 0, fullImg, resize.Lanczos3)
	shades := []string{" ", "░", "▒", "▓", "█"}
	yStart := resizedImg.Bounds().Min.Y
	yEnd := resizedImg.Bounds().Max.Y
	for y := yStart; y < yEnd; y++ {
	    xStart := resizedImg.Bounds().Min.X
	    xEnd := resizedImg.Bounds().Max.X
		for x := xStart; x < xEnd; x++ {
			// Convert to greyscale
			col := color.GrayModel.Convert(resizedImg.At(x, y)).(color.Gray)
			shade := col.Y / 51
			if shade >= 5 {
				shade = shade - 1
			}
			shadeToPrint := shades[shade]
			fmt.Print(shadeToPrint)
		}
		// Start a new y
		fmt.Print("\n")
	}
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
