package crawler

import (
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/husni-robani/domain-link-crawler.git/internal/utils"
)

type DataLink struct {
	URL url.URL
	InternalLinksFound []url.URL
	ExternalLinksFound []url.URL
	TotalURLAppearence int
}

type CrawlerCofig struct {
	Pages map[string]*DataLink
	BaseURL *url.URL
	Mu *sync.Mutex
	ConcurrencyControl chan struct{}	
	Wg *sync.WaitGroup
	MaxPages int
}

func (cfg *CrawlerCofig) RunCrawl(){
	start := time.Now()

	cfg.Wg.Add(1)
	go cfg.crawlPage(cfg.BaseURL.String())

	cfg.ConcurrencyControl <- struct{}{}
	
	cfg.Wg.Wait()

	close(cfg.ConcurrencyControl)
	
	cfg.printReport(cfg.BaseURL.String())

	// TODO:  Masukan kedalam printReport:
	fmt.Println("Total Pages: ", len(cfg.Pages))
	fmt.Println("Execution Time: ", time.Since(start))
}

func (cfg *CrawlerCofig) crawlPage(rawCurrentURL string) {
	defer cfg.Wg.Done()

	if len(cfg.Pages) >= cfg.MaxPages {
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
	if err != nil {
		log.Fatal(err.Error())
	}

	// check is normalized_current_url already crawled
	isCrawled := cfg.addPageVisit(normalized_current_url)
	if isCrawled {
		<- cfg.ConcurrencyControl
		return
	}
	
	// crawling the page
	cfg.Mu.Lock()
	cfg.Pages[normalized_current_url] = &DataLink{
		URL: *current_url_parsed,
		TotalURLAppearence: 1,
	}
	cfg.Mu.Unlock()
	html, err := utils.GetHTML(rawCurrentURL)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Printf("\nStarting crawler of: %s\n...\n", rawCurrentURL)
	fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

	urls, err := utils.GetURLs(html, cfg.BaseURL.String())
	if err != nil {
		log.Fatal(err)
	}

	<- cfg.ConcurrencyControl

	for _, url_item := range urls{
		parsed_url, _ := url.Parse(url_item)
		if parsed_url.Host == cfg.BaseURL.Host{
			cfg.Mu.Lock()
			cfg.Pages[normalized_current_url].InternalLinksFound = append(cfg.Pages[normalized_current_url].InternalLinksFound, *parsed_url)
			cfg.Mu.Unlock()
		}else {
			cfg.Mu.Lock()
			cfg.Pages[normalized_current_url].ExternalLinksFound = append(cfg.Pages[normalized_current_url].ExternalLinksFound, *parsed_url)
			cfg.Mu.Unlock()
		}

		cfg.Wg.Add(1)
		go cfg.crawlPage(url_item)
		
		cfg.ConcurrencyControl <- struct{}{}
	}
}

func (cfg *CrawlerCofig) printReport(baseURL string){

	// TODO: Ubah menggunakan logger
	fmt.Printf(`
==========================================================
  REPORT for %v
==========================================================
`, baseURL)

	for _, page := range cfg.Pages{
		fmt.Printf("---%v---\n", page.URL.String())
		fmt.Printf("Appearence in other page: %v\n", page.TotalURLAppearence)
		fmt.Printf("External links: %v\n", len(page.ExternalLinksFound))
		fmt.Printf("Internal links: %v\n", len(page.InternalLinksFound))
	}
}

func (cfg *CrawlerCofig) addPageVisit(normalizedURL string) (isFirst bool){
	defer cfg.Mu.Unlock()

	cfg.Mu.Lock()
	_, ok := cfg.Pages[normalizedURL]
	if ok {
		cfg.Pages[normalizedURL].TotalURLAppearence += 1
		return true
	}

	return false
}