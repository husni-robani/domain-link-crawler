package main

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/husni-robani/domain-link-crawler.git/internal/crawler"
	"github.com/husni-robani/domain-link-crawler.git/internal/utils/logger"
)

type userInput struct {
	url url.URL
	goroutine_size int
	max_pages int
	is_export bool
}

func main(){		
	inputs, err := getInputs()
	if err != nil {
		logger.FatalDefaultLogger.Fatal(err.Error())
		return
	}

	base_url, err := url.Parse(inputs.url.String())
	if err != nil {
		logger.ErrDefaultLogger.Error(err.Error())
	}

	crawler := crawler.NewCrawl(base_url, inputs.max_pages, inputs.goroutine_size)
	crawler.RunCrawl()
}

func getInputs() (userInput, error){
	user_input := userInput{}

	// url
	fmt.Printf("URL: ")
	var url_input string
	fmt.Scanln(&url_input)
	
	url_parsed, err := url.Parse(url_input)
	// fmt.Println("err: ", err)
	if err != nil {
		return userInput{}, fmt.Errorf("failed to parse url [value: %v] [error: %v]", url_input, err)
	}
	
	user_input.url = *url_parsed

	// goroutine size
	fmt.Printf("Goroutine size: ")
	var goroutine_size_input string
	fmt.Scanln(&goroutine_size_input)

	goroutine_size, err := strconv.Atoi(goroutine_size_input)
	if err != nil {
		return userInput{}, fmt.Errorf("failed to convert goroutine size to int [value: %v] [error: %v]", goroutine_size_input, err)
	}
	user_input.goroutine_size = goroutine_size

	// max page
	fmt.Printf("Max pages: ")
	var max_page_input string
	fmt.Scanln(&max_page_input)

	max_pages, err := strconv.Atoi(max_page_input)
	if err != nil {
		return userInput{}, fmt.Errorf("failed to convert max page to int [value: %v] [error: %v]", max_page_input, err)
	}
	user_input.max_pages = max_pages

	// is export input
	fmt.Printf("Export to csv (y/n): ")	
	var is_export string
	fmt.Scanln(&is_export)

	if is_export == "Y" || is_export == "y" {
		user_input.is_export = true
	}else{
		user_input.is_export = false
	}

	return user_input, nil
}