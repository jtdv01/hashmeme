package image_processor

import (
	"fmt"
	"testing"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	encodedImage := ReadImageFile(testImagePath)
	WriteImageToFile("./tmp/base.png", encodedImage)
}

func TestReadTextFromImage(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	text := ReadTextFromImage(testImagePath)
	fmt.Println(text)
}

func TestBlendImages(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	baseImage := ReadImageFile(testImagePath)
	text := ReadTextFromImage(testImagePath)
	qrImage := GenerateQr(text)
	outImage := BlendImageWithWaterMark(baseImage, qrImage)
	WriteImageToFile("./tmp/blended.png", outImage)
}

func TestHashImageSha256(t *testing.T) {
    testImagePath := "./resources/hashmeme.png"
    hash := HashImageSha256(testImagePath)
    fmt.Println(hash)
}