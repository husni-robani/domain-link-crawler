package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"

	"github.com/husni-robani/domain-link-crawler.git/internal/crawler"
)

func main(){		
	if len(os.Args) > 4 {
		log.Println("too many arguments provided")
		return
	}else if len(os.Args) < 4 {
		switch len(os.Args) {
		case 1: 
			log.Println("url, goroutine size, and max pages argument needed!")
		case 2: 
			log.Println("goroutine size and max pages argument needed!")
		case 3: 
			log.Println("max pages argument needed!")
		}
		return
	}

	raw_url := os.Args[1]
	base_url, err := url.Parse(raw_url)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	goroutine_size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Println("invalid goroutine size argument ! It should be int")
		return
	}

	max_pages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Println("invalid max pages argument ! It should be int")
		return
	}

	crawler := crawler.NewCrawl(base_url, max_pages, goroutine_size)
	crawler.RunCrawl()
}