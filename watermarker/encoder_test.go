package watermarker

import (
	"fmt"
	"testing"
)

func TestEncoder(t *testing.T) {
	text := "Hello Future!"
	got := Encode(text)

	fmt.Println(got)
}
