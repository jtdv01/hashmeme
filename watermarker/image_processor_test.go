package watermarker

import (
	"testing"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	encodedImage := ReadImageFile(testImagePath)
	PrintImage(encodedImage)

}

func TestBlendImages(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	baseImage := ReadImageFile(testImagePath)
	qrImage := GenerateQr("Some text")
	outImage := BlendImageWithWaterMark(baseImage, qrImage)
	PrintImage(outImage)
}
