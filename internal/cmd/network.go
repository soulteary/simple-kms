package cmd

import (
	"bytes"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/text/encoding/htmlindex"
)

func GetApi(primary string, fallback string) string {
	data, err := Get(primary)
	if err != nil && fallback != "" {
		data, err = Get(fallback)
		if err != nil {
			os.Exit(0)
		}
	}

	return strings.TrimSpace(strings.Split(string(data), "\n")[0])
}

func Get(url string) ([]byte, error) {

	client := &http.Client{Timeout: 3 * time.Second}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("http code error: " + res.Status)
	}

	decoder, err := htmlindex.Get("utf-8")
	if err != nil {
		return nil, err
	}

	buffer := new(bytes.Buffer)
	buffer.ReadFrom(decoder.NewDecoder().Reader(res.Body))

	return buffer.Bytes(), nil
}
