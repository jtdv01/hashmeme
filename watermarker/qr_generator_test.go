package watermarker

import (
	"testing"
)

func TestQrGenerator(t *testing.T) {
	text := "Hello Future!"
	got := GenerateQr(text)
    PrintImage(got)
}
