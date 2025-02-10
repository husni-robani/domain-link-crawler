package utils

import (
	"fmt"
	"slices"
	"testing"
)

func TestGetURLs(t *testing.T){
	tests := []struct{
		name string
		inputURL string
		inputBody string
		expected []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no links present",
			inputURL: "https://nolinks.com",
			inputBody: `
		<html>
			<body>
				<p>No links in this HTML!</p>
			</body>
		</html>
			`,
			expected: []string{},
		},
		{
			name:     "link with query parameters",
			inputURL: "https://queries.com",
			inputBody: `
		<html>
			<body>
				<a href="/search?q=golang">Search Golang</a>
				<a href="https://docs.com/doc?id=1234">Documentation</a>
			</body>
		</html>
			`,
			expected: []string{
				"https://queries.com/search?q=golang",
				"https://docs.com/doc?id=1234",
			},
		},
		{
			name:     "link with hash fragment",
			inputURL: "https://fragments.com",
			inputBody: `
		<html>
			<body>
				<a href="/section#part1">Part 1</a>
				<a href="https://external.com/page#anchor">External Anchor</a>
			</body>
		</html>
			`,
			expected: []string{
				"https://fragments.com/section#part1",
				"https://external.com/page#anchor",
			},
		},
	}

	passCount := 0
	failCount := 0
	for i, tc := range tests {
		urls, err := GetURLs(tc.inputBody, tc.inputURL)
		if err != nil {
			t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			return
		}
		if !slices.Equal(urls, tc.expected) {
			failCount ++
			t.Errorf(`
----------------------------------
Test Fail
Expecting: %v
Actual: %v
`, tc.expected, urls)
		} else {
			passCount ++
			fmt.Printf(`
----------------------------------
Test Pass
Expecting: %v
Actual: %v
`, tc.expected, urls)
		}
	}
	fmt.Println("----------------------------------")
	fmt.Printf("%v Passed, %v Failed\n", passCount, failCount)
}