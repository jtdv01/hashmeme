package watermarker

import (
    "fmt"
	"testing"
)

func TestImageProcessor(t *testing.T) {
	testImagePath := "./resources/whos-the-original.png"
	encodedImage := ReadImageFile(testImagePath)
	fmt.Println(encodedImage)

}
