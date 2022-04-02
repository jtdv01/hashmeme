package watermarker

import (
	"fmt"
	"testing"
)

func TestQrGenerator(t *testing.T) {
	text := "Hello Future!"
	got := GenerateQr(text)

	fmt.Println(got)
}
