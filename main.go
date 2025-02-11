package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/husni-robani/domain-link-crawler.git/app"
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

	runCrawl(raw_url, *base_url, goroutine_size, max_pages)
}

func runCrawl(raw_url string, base_url url.URL, goroutine_size int, maxPages int){
	if goroutine_size < 1 {
		goroutine_size = 2
	}
	crawler_config := app.CrawlConfig{
		Pages: map[string]*app.DataLink{},
		BaseURL: &base_url,
		Mu: &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, goroutine_size),
		Wg: &sync.WaitGroup{},
		MaxPages: maxPages,
	}

	start := time.Now()

	crawler_config.Wg.Add(1)
	go crawler_config.CrawlPage(raw_url)

	crawler_config.ConcurrencyControl <- struct{}{}
	
	crawler_config.Wg.Wait()

	close(crawler_config.ConcurrencyControl)

	crawler_config.PrintReport(crawler_config.BaseURL.String())
	fmt.Println("Total Pages: ", len(crawler_config.Pages))
	fmt.Println("Execution Time: ", time.Since(start))
}