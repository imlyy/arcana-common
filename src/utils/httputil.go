package utils

import (
	"bytes"
	"net/http"
	"io/ioutil"
)

type HttpUtil struct {
}

func (httpUtil *HttpUtil) CustomPost(url string, contentType string, content string) (string, error) {
	body := bytes.NewBuffer([]byte(content))
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return "data error", err
	}

	defer resp.Body.Close()

	respStr, err := ioutil.ReadAll(resp.Body)
	return string(respStr), err
}
