package watermarker

import (
    "testing"
    "fmt"
)

func TestEncoder(t *testing.T) {
    text := "Hello Future!"
    got := Encode(text)

    fmt.Println(got)
}