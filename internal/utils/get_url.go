package utils

import (
	"strings"

	"golang.org/x/net/html"
)

func GetURLs(htmlBody, rawBaseURL string) ([]string, error){

	tokenizer := html.NewTokenizer(strings.NewReader(htmlBody))
	var urls []string

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken{
			if tokenizer.Err().Error() == "EOF"{
				break
			}
			return urls, tokenizer.Err()
		}

		if tokenType == html.StartTagToken{
			token := tokenizer.Token()

			if token.Data == "a"{
				for _, attribute := range token.Attr{
					if attribute.Key == "href"{
						// check is URL is a relative href or absolute href
						url := attribute.Val
						if !strings.HasPrefix(attribute.Val, "https://") && !strings.HasPrefix(attribute.Val, "http://"){
							url = rawBaseURL + attribute.Val
						}
						// store in []URLs
						urls = append(urls, url)
					}
				}
			}
		}
	}

	return urls, nil
}