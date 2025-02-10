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

1. download the domain-link-crawler file in this repository
2. Run the crawler with the following command:
   ```bash
   ./domain-link-crawler <url> <goroutine-size> <max-depth>
   ```
   Arguments:
   - `<url>`: The starting URL of the domain to crawl (required).
   - `<goroutine-size>`: The size of the goroutine pool for concurrent crawling (e.g., 3).
   - `<max-pages>`: The maximum depth of crawling (e.g., 5).

### Output

```bash
Starting crawler of: medium.com
...
------------------------------------------------------------------------------------------------------------------------------

Starting crawler of: medium.com/membership
...
------------------------------------------------------------------------------------------------------------------------------

=============================
  REPORT for https://medium.com
=============================
Found 2 internal links to medium.com
Found 1 internal links to medium.com/membership
Total Pages:  2
Execution Time:  683.382ms
```
