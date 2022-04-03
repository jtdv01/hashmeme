package image_processor

import (
	"testing"
	"fmt"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	encodedImage := ReadImageFile(testImagePath)
	WriteImageToFile("./tmp/base.png", encodedImage)
}

func TestBlendImages(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	baseImage := ReadImageFile(testImagePath)
	qrImage := GenerateQr("Hello future")
	outImage := BlendImageWithWaterMark(baseImage, qrImage)
	WriteImageToFile("./tmp/blended.png", outImage)
}

func TestReadTextFromImage(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
    text := ReadTextFromImage(testImagePath)
    fmt.Println(text)
}