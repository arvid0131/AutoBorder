package main

import (
	"bytes"
	"image"

	"github.com/skip2/go-qrcode"
)

func generateQrImg(msg string) image.Image {
	qrByte, _ := qrcode.Encode(msg, qrcode.Highest, 256)
	img, _, _ := image.Decode(bytes.NewReader(qrByte))

	return img
}

func generateQRBytes(msg string) []byte {
	qrByte, _ := qrcode.Encode(msg, qrcode.Highest, 256)
	return qrByte
}
