package image_processor

import (
	"testing"
)

func TestQrGenerator(t *testing.T) {
	text := "Hello Future!"
	got := GenerateQr(text)
	WriteImageToFile("./tmp/qr.png", got)
}
