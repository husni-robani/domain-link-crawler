package report

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/husni-robani/domain-link-crawler.git/internal/crawler"
	"github.com/husni-robani/domain-link-crawler.git/internal/utils/logger"
)

type reportCsv struct {
	dataPages []crawler.DataLink
	dirName string
}

func NewReportCsv (data []crawler.DataLink, dirName string) reportCsv{
	return reportCsv{
		dataPages: data,
		dirName: "report_output/" + dirName + "/",
	}
}

func (cfg reportCsv) processedLinkData() [][]string {
	data := [][]string{
		{"URL", "Total External Links", "Total Internal Links", "Total Appearence on Other Page"},
	}

	for _, page := range cfg.dataPages {
		data = append(data, []string{page.URL.String(), strconv.Itoa(len(page.ExternalLinksFound)), strconv.Itoa(len(page.InternalLinksFound)), strconv.Itoa(page.TotalURLAppearence)})
	}

	return data
}

func (cfg reportCsv) processedPageInternalLinks() [][]string {
	data := [][]string{
		{"Page", "Internal Link Founded"},
	}

	for _, page := range cfg.dataPages{
		for i, link := range page.InternalLinksFound{
			if i == 0{
				data = append(data, []string{page.URL.String(), link.String()})
			}else{
				data = append(data, []string{"", link.String()})
			}
		}
	}
	return data
}

func (cfg reportCsv) processedPageExternalLinks() [][]string {
	data := [][]string{
		{"Page", "External Link Founded"},
	}

	for _, page := range cfg.dataPages{
		for i, link := range page.ExternalLinksFound{
			if i == 0{
				data = append(data, []string{page.URL.String(), link.String()})
			}else{
				data = append(data, []string{"", link.String()})
			}
		}
	}
	return data
}

func (cfg reportCsv) Generate() {
	// type function func()
	type generateCsv struct {
		fileName string
		processDataFunc func() [][]string
	}
	generate_files := []generateCsv{
		{"page_data.csv", cfg.processedLinkData},
		{"page_internal_links.csv", cfg.processedPageInternalLinks},
		{"page_external_links.csv", cfg.processedPageExternalLinks},
	}

	wg := sync.WaitGroup{}

	for _, v := range generate_files {
		wg.Add(1)
		go func(){
			defer wg.Done()
			os.Mkdir(cfg.dirName, 0755)
			file, err := os.Create(cfg.dirName + v.fileName)
			if err != nil {
				logger.FatalDefaultLogger.Fatal("failed to create new directory", cfg.dirName, err)
				return
			}
			w := csv.NewWriter(file)
			w.WriteAll(v.processDataFunc())
			w.Flush()
		
			if err := w.Error(); err != nil {
				logger.FatalDefaultLogger.Fatal("Failed exporting to CSV", err)
				return
			}

			logger.InfoDefaultLogger.Info(fmt.Sprintf("%v created!", v.fileName))
		}()
	}

	wg.Wait()

	logger.InfoDefaultLogger.Info(fmt.Sprintf("Data exported successfully to %s", cfg.dirName))
}

