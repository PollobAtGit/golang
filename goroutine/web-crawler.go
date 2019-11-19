package main

import (
	"fmt"
)

type fetcher interface {
	fetch(url string) (body string, urls []string, err error)
}

func crawl(url string, depth int, f fetcher) {
	if depth <= 0 {
		return
	}
	body, urls, err := f.fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		crawl(u, depth-1, f)
	}
	return
}

func main() {
	crawl("https://golang.org/", 4, fFetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

var fFetcher = fakeFetcher{

	"https://golang.org/": &fakeResult{
		"The Go Programming Language", []string{"https://golang.org/pkg/", "https://golang.org/cmd/"},
	},

	"https://golang.org/pkg/": &fakeResult{
		"Packages", []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"},
	},

	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt", []string{"https://golang.org/", "https://golang.org/pkg/"},
	},

	"https://golang.org/pkg/os/": &fakeResult{
		"Package os", []string{"https://golang.org/", "https://golang.org/pkg/"},
	},
}
