package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func UploadQR(filename string) []byte {

	url := "https://api.imgbb.com/1/upload?key=2a80b2922634eb48e8ac0171f4db4429&image="
	method := "POST"

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	_ = writer.WriteField("key", "2a80b2922634eb48e8ac0171f4db4429")

	file, errFile2 := os.Open(filename)

	defer file.Close()

	part2, errFile2 := writer.CreateFormFile("image", filepath.Base(filename))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return nil
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(filename)
	os.Remove(filename)

	return body

}
