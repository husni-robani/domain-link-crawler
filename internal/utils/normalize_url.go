package utils

import (
	"errors"
	"net/url"
	"strings"
)

func NormalizeURL(raw_url string) (string, error){

	url, err := url.Parse(raw_url)
	if err != nil {
		return "", errors.New("error while parsing url")
	}

	// remove "/" in the last of url if it exist
	normalize_url := url.Host + url.Path
	last_slash := strings.LastIndex(normalize_url, "/")
	if last_slash == len(normalize_url) - 1{
		normalize_url, _ = strings.CutSuffix(normalize_url, "/")
	}


	return normalize_url, nil
}