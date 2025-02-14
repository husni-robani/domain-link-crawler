package crawler

import (
	"net/url"
	"sync"
)

func NewCrawl(baseURL *url.URL, maxPages int, goroutineSize int) *CrawlerCofig{
	if goroutineSize <= 0 {
		goroutineSize = 1
	}
	return &CrawlerCofig{
		Pages: map[string]*DataLink{},
		BaseURL: baseURL,
		Mu: &sync.Mutex{},
		ConcurrencyControl: make(chan struct{}, goroutineSize),
		Wg: &sync.WaitGroup{},
		MaxPages: maxPages,
	}
}