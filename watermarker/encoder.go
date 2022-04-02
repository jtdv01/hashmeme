package watermarker

import (
	qrcode "github.com/skip2/go-qrcode"
	"log"
)

func Encode(textToEncode string) []byte {
	var png []byte
	png, err := qrcode.Encode(textToEncode, qrcode.Medium, 128)
	if err != nil {
		log.Fatal(err)
	}
	return png
}
