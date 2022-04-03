package image_processor

import (
	"errors"
	"github.com/otiai10/gosseract/v2"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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
	point1 := image.Point{bounds.Max.X - qrBounds.Max.X, bounds.Max.Y - qrBounds.Max.Y}
	point2 := bounds.Max
	// Put qr down at the bottom right
	qrBounds = image.Rectangle{point1, point2}
	outImage := image.NewRGBA(bounds)

	// Draw with bounds of baseImage
	draw.Draw(outImage, bounds, baseImage, image.ZP, draw.Src)
	draw.Draw(outImage, qrBounds, qrImage, image.ZP, draw.Src)
	return outImage
}

func FilterOutNonText(img image.Image) *image.Gray {
	// Filters out non text pixels from an image so OCR can read the text
	bounds := img.Bounds()
	gray := image.NewGray(bounds)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var pixel = img.At(x, y)
			col := color.RGBAModel.Convert(pixel).(color.RGBA)
			gray.Set(x, y, col)
			grayPixel := gray.At(x, y)
			grayness := color.GrayModel.Convert(grayPixel).(color.Gray)
			black := color.Gray{1}
			white := color.Gray{255}
			if grayness.Y < 230 {
				gray.Set(x, y, white)
			} else {
				gray.Set(x, y, black)
			}
		}
	}
	return gray
}

func ReadTextFromImage(inputImagePath string) string {
	// Convert to greyscale
	img := ReadImageFile(inputImagePath)
	grey := FilterOutNonText(img)
	tmpOutputForFilteredText := "./tmp/text_filtered.png"
	WriteImageToFile(tmpOutputForFilteredText, grey)

	// Tesseract for reading
	client := gosseract.NewClient()
	err := client.SetConfigFile("./tesseract.ini")
	if err != nil {
		log.Fatalf("Failed to load config %s", err)
	}
	defer client.Close()

	client.SetImage(tmpOutputForFilteredText)
	text, err := client.Text()
	if err != nil {
		log.Fatalf("Failed to read text: %s", err)
	}

	// Normalize spaces
	regexp, err := regexp.Compile(`[\s]+`)
	match := regexp.ReplaceAllString(text, " ")
	// Set to lower case
	out := strings.ToLower(match)
	return out
}
