package utils

import (
	"fmt"
	"io"
	"net/http"
)

func GetHTML(rawURL string) (string, error) {
	response, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch html. error: %v", err)
	}

	body, err := io.ReadAll(response.Body)
	defer response.Body.Close()

	if response.StatusCode > 299 {
		return "", fmt.Errorf("response failed with status code: %d, while get response body", response.StatusCode)
	}
	if err != nil {
		return "", err
	}
	if (response.Header.Get("Content-Type") != "text/html; charset=utf-8") && (response.Header.Get("Content-Type") != "text/html" && response.Header.Get("Content-Type") != "text/html;charset=utf-8") && (response.Header.Get("Content-Type") != "text/html; charset=UTF-8"){
		return "", fmt.Errorf("invalid content type: %#v", response.Header.Get("Content-Type"))
	}

	return string(body), nil
}