package utils

import (
	"fmt"
	"sort"
)

func PrintReport(pages map[string]int, baseURL string){
	fmt.Printf(`
=============================
  REPORT for %v
=============================
`, baseURL)

	var keys []string
	for v := range pages{
			keys = append(keys, v)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if pages[keys[i]] == pages[keys[j]] {
			return sort.StringsAreSorted([]string{keys[i], keys[j]})
		}


		return pages[keys[i]] > pages[keys[j]]
	})

	for _, key := range keys{
		fmt.Printf("Found %v internal links to %v\n", pages[key], key)
	}
}