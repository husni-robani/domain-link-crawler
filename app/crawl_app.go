package app

import (
	"domain-link-crawler/app/utils"
	"fmt"
	"log"
	"net/url"
	"sync"
)

type CrawlConfig struct {
	Pages map[string]int
	BaseURL *url.URL
	Mu *sync.Mutex
	ConcurrencyControl chan struct{}	
	Wg *sync.WaitGroup
	MaxPages int
}

func (cfg *CrawlConfig) CrawlPage(rawCurrentURL string) {
	defer cfg.Wg.Done()

	if len(cfg.Pages) >= cfg.MaxPages {
		// Not working as expected yet
		<- cfg.ConcurrencyControl
		return
	}

	// CHECK IS THE DOMAIN OF rawCurrentURL SAME AS rawBaseURL
	current_url_parsed, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Fatalf("Error parsing url. error: %v", err)
	}

	if current_url_parsed.Host != cfg.BaseURL.Host{
		<- cfg.ConcurrencyControl
		return
	}
	
	// get a normalized of raw current url
	normalized_current_url, err := utils.NormalizeURL(rawCurrentURL)
	// fmt.Printf("Normalized current URL: %v\n", normalized_current_url)
	if err != nil {
		log.Fatal(err.Error())
	}
	// check is normalized_current_url already crawled
	isCrawled := cfg.AddPageVisit(normalized_current_url)
	if isCrawled {
		<- cfg.ConcurrencyControl
		return
	}
	
	// crawling the page
	cfg.Mu.Lock()
	cfg.Pages[normalized_current_url] = 1
	cfg.Mu.Unlock()
	html, err := utils.GetHTML(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("\nStarting crawler of: %s\n...\n", normalized_current_url)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	urls, err := utils.GetURLs(html, cfg.BaseURL.String())
	if err != nil {
		log.Fatal(err)
	}

	<- cfg.ConcurrencyControl

	for _, url := range urls{
		cfg.Wg.Add(1)
		go cfg.CrawlPage(url)
		
		cfg.ConcurrencyControl <- struct{}{}
	}
}

func (cfg *CrawlConfig) AddPageVisit(normalizedURL string) (isFirst bool){
	defer cfg.Mu.Unlock()

	cfg.Mu.Lock()
	_, ok := cfg.Pages[normalizedURL]
	if ok {
		cfg.Pages[normalizedURL] = cfg.Pages[normalizedURL] + 1
		return true
	}

	return false
}