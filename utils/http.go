package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func PostForm(url string, header map[string]string, param string) (string, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(param)))
	if err != nil {
		return "", err
	}
	// req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
