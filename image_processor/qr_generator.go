package image_processor

import (
	"bytes"
	qrcode "github.com/skip2/go-qrcode"
	"image"
	"log"
)

func GenerateQr(textToEncode string) image.Image {
	var imgByteArray []byte
	imgByteArray, err := qrcode.Encode(textToEncode, qrcode.Medium, 2)
	if err != nil {
		log.Fatal(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(imgByteArray))
	return img
}
