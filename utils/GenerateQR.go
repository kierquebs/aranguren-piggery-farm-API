package utils

import (
	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQR(content string, filename string) error {
	err := qrcode.WriteFile(content, qrcode.Medium, 250, filename+".png")
	if err != nil {
		return err
	}
	return nil
}
