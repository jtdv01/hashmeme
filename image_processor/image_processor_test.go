package image_processor

import (
	"fmt"
	"testing"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	encodedImage := ReadImageFile(testImagePath)
	WriteImageToFile("./tmp/base.png", encodedImage)
}

func TestReadTextFromImage(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	text := ReadTextFromImage(testImagePath)
	fmt.Println(text)
}

func TestBlendImages(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	baseImage := ReadImageFile(testImagePath)
	text := ReadTextFromImage(testImagePath)
	qrImage := GenerateQr(text)
	outImage := BlendImageWithWaterMark(baseImage, qrImage)
	WriteImageToFile("./tmp/blended.png", outImage)
}

