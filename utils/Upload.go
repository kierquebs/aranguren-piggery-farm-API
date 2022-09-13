package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kierquebs/aranguren-piggery-farm-API/model"
)

func UploadQR(filename string) (model.ImageUpload, error) {

	url := "https://api.imgbb.com/1/upload?key=2a80b2922634eb48e8ac0171f4db4429&image="
	method := "POST"

	payload := &bytes.Buffer{}

	writer := multipart.NewWriter(payload)

	_ = writer.WriteField("key", "2a80b2922634eb48e8ac0171f4db4429")

	file, errFile2 := os.Open(filename)

	defer file.Close()

	var result model.ImageUpload

	part2, errFile2 := writer.CreateFormFile("image", filepath.Base(filename))
	_, errFile2 = io.Copy(part2, file)
	if errFile2 != nil {
		fmt.Println(errFile2)
		return result, errFile2
	}
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return result, err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return result, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if err := json.Unmarshal(body, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	os.Remove(filename)

	return result, nil

}
