package main

import (
	"net/url"
	"os"
	"strconv"

	"github.com/husni-robani/domain-link-crawler.git/internal/crawler"
	"github.com/husni-robani/domain-link-crawler.git/internal/utils/logger"
)

func main(){		
	if len(os.Args) > 4 {
		logger.ErrDefaultLogger.Error("too many arguments provided")
		return
	}else if len(os.Args) < 4 {
		switch len(os.Args) {
		case 1: 
			logger.ErrDefaultLogger.Error("url, goroutine size, and max pages argument needed!")
		case 2: 
			logger.ErrDefaultLogger.Error("goroutine size and max pages argument needed!")
		case 3: 
			logger.ErrDefaultLogger.Error("max pages argument needed!")
		}
		return
	}

	raw_url := os.Args[1]
	base_url, err := url.Parse(raw_url)
	if err != nil {
		logger.ErrDefaultLogger.Error(err.Error())
	}

	goroutine_size, err := strconv.Atoi(os.Args[2])
	if err != nil {
		logger.ErrDefaultLogger.Error(err.Error())
		return
	}

	max_pages, err := strconv.Atoi(os.Args[3])
	if err != nil {
		logger.ErrDefaultLogger.Error("invalid max pages argument ! It should be int")
		return
	}

	crawler := crawler.NewCrawl(base_url, max_pages, goroutine_size)
	crawler.RunCrawl()
}