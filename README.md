# Internal Link Crawler

A Golang-base web crawler designed to crawl and count all internal links within a specified domain (e.g., `https://example.com`). This tool help you map and anlyze the structure of a website by discovering all pages under the same domain.

---

## Features

- **Internal Link Discovery**: Crawls and identifies all internal links within a domain.
- **Page Counting**: Reports the number of links found for each pages.
- **Efficient Crawling**: Uses concurrent goroutines for fast and efficient crawling.
- **Configurable**: Allows customization of goroutines size and max pages to control the resource used.
- **Lightweight**: Built with Go for better performance and lower resource usage.

## Installation

You can either clone the repository or download the pre-built executable file (domain-link-crawler.exe).

### Option 1: Clone the repository

1. Ensure you have [Go installed](https://golang.org/doc/install) on your system.
2. Clone this repository:
   ```bash
   git clone https://github.com/your-username/domain-link-crawler.git
   ```
3. Navigate to the project directory:
   ```bash
   cd domain-link-crawler
   ```
4. Build the project:
   ```bash
   go build -o domain-link-crawler
   ```

### Option 2: Download the executable

download the executable file in this repository

## Usage & Output Example

#### running the program

You can run the program using either of the following commands:

1. Using `go run`:

```bash
$ go run cmd/main.go
```

2. Using the compiled executable:

```bash
$ ./web-crawler
```

#### input prompts

Whaen the program start, it will prompt you for the following inputs:

1. **URL**: The starting URL for the crawler
2. **Goroutines Size**: The number of goroutines to use for concurrenct crawling
3. **Max Pages**: The maximum number of pages to crawl
4. **Export to CSV**: Whether to export the results to a CSV file.

#### Example Interaction

```bash
URL: https://www.kompas.com/#google_vignette
Goroutine size: 5
Max pages: 25
Export to csv (y/n): y
Directory name for save exported csv data: example_002
```

#### CSV Export

If you chose to export the results to CSV, the program will generate 3 CSV file in the specified directory. The files is:

- `page_data.csv`
  It will generate csv file with column `URL`, `Total External Links`, `Total Internal Links`, and `Total Appearence on Other Page`. The purpose of this data is to provide the overview each page has been crawled.

Example CSV content:

```
URL,Total External Links,Total Internal Links,Total Appearence on Other Page
https://www.kompas.com/food/tips-kuliner,245,95,12
https://www.kompas.com/tag/efisiensi-anggaran?source=trending,251,50,1
https://www.kompas.com/cekfakta/read/2025/02/14/131600582/infografik--hoaks-meninggalnya-aktor-stefan-william-simak-bantahannya,262,114,1
https://www.kompas.com/cekfakta/read/2024/12/13/132600382/mohammad-ahsan-pensiun-menutup-kisah-the-daddies-di-bulu-tangkis-,263,122,1
http://www.kompas.com/global/read/2025/01/06/215716170/ikan-tuna-segemuk-sapi-di-jepang-laku-rp-21-miliar,293,93,1
```

- `page_external_links.csv`
  It will generate csv file with column `Page` and `External Link Founded`. It provide the detail what are external links has been founded while crawling a page

Example CSV content:

```
Page,External Link Founded
https://www.kompas.com/food/tips-kuliner,https://play.kompas.com/loyalty
,https://account.kompas.com/login/a29tcGFz/aHR0cHM6Ly93d3cua29tcGFzLmNvbS9mb29kL3RpcHMta3VsaW5lcg==
,https://activity.kompas.com/voucher?source=navbar
,https://plus.kompas.com/detail?source=profile
```

- `page_internal_links.csv`
  It will generate csv file with column `Page` and `Internal Link Founded`. It provide the detail what are internal links has been founded while crawling a page

Example CSV content:

```
Page,Internal Link Founded
https://www.kompas.com/food/tips-kuliner,https://www.kompas.com
,https://www.kompas.com/feedback
,https://www.kompas.com
,https://www.kompas.com/global
```

#### Example Without CSV Export

```
$ go run cmd/main.go

URL: https://www.kompas.com/#google_vignette
Goroutine size: 1
Max pages: 2
Export to csv (y/n): n

[INFO] Starting crawler of: https://www.kompas.com/#google_vignette
...
-------------------------------------------------------------------------
[INFO] Starting crawler of: https://www.kompas.com/feedback
...
-------------------------------------------------------------------------

==========================================================
  REPORT for https://www.kompas.com/#google_vignette
==========================================================
---https://www.kompas.com/#google_vignette---
Appearence in other page: 3
External links: 355
Internal links: 85
---https://www.kompas.com/feedback---
Appearence in other page: 1
External links: 232
Internal links: 39

[INFO] Total Pages: 2
[INFO] Execution Time: 422.1585ms
```
