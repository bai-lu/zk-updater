package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"unsafe"
)

func main() {
	data := make(map[string]string)
	newFileUploadRequest("https://open.f.***.cn/open-apis/image/v4/put/",
		data, "image", "error.webp")
}
func newFileUploadRequest(uri string, params map[string]string, paramName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, path)
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.WriteField("image_type", "message")
	err = writer.Close()
	if err != nil {
		return err
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set("Authorization", "Bearer t-2ede9cf42706f6c41f6552d7144aafb2c067bb23")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println(*str)
	return err
}
