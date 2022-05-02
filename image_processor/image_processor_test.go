package image_processor

import (
	"fmt"
	"testing"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	encodedImage, _, _ := ReadImageFile(testImagePath)
	WriteImageToFile("./tmp/base.png", encodedImage)
}

func TestReadTextFromImage(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	text, _ := ReadTextFromImage(testImagePath)
	fmt.Println(text)
}

func TestBlendImages(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	baseImage, _, _ := ReadImageFile(testImagePath)
	text, _ := ReadTextFromImage(testImagePath)
	qrImage := GenerateQr(text)
	outImage := BlendImageWithWaterMark(baseImage, qrImage)
	WriteImageToFile("./tmp/blended.png", outImage)
}

func TestHashImageSha256(t *testing.T) {
	testImagePath := "./resources/hashmeme.png"
	hash, _ := HashImageSha256(testImagePath)
	fmt.Println(fmt.Sprintf("The hash is %v", hash))
}
