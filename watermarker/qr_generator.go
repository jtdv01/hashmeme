package watermarker

import (
	qrcode "github.com/skip2/go-qrcode"
	"log"
	"image"
	"bytes"
)

func GenerateQr(textToEncode string) image.Image {
	var imgByteArray []byte
	imgByteArray, err := qrcode.Encode(textToEncode, qrcode.Medium, 128)
	if err != nil {
		log.Fatal(err)
	}
	img, _, _ := image.Decode(bytes.NewReader(imgByteArray))
	return img
}
