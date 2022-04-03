package image_processor

import (
	"bytes"
	"fmt"
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"log"
)

func GenerateQr(inputText string) image.Image {
	var imgByteArray []byte
	textToEncode := fmt.Sprintf("Transaction receipt: %s\nHashMeme Content: %s", "TODO", inputText)
	imgByteArray, err := qrcode.Encode(textToEncode, qrcode.Medium, 2)
	if err != nil {
		log.Fatal(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(imgByteArray))
	return img
}
