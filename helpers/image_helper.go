package helpers

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GetImageFromUrl(url string) image.Image {
	res, err := http.Get(url)
	if err != nil || res.StatusCode != 200 {
		fmt.Println("Error: ", err)
	}

	defer res.Body.Close()
	m, _, err := image.Decode(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	return m
}

func SaveImage(i image.Image) (path string) {
	path = string(time.Now().UnixNano())

	f, err := os.Create(path)
	LogError(err)

	err = jpeg.Encode(f, i, nil)
	LogError(err)

	err = f.Close()
	LogError(err)

	return
}

func GetReaderFromImage(url string) (out io.Reader) {
	b, err := ioutil.ReadFile(url)
	LogError(err)

	out = bytes.NewReader(b)

	return
}
