package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string, c chan FetchResult)
}

type FetchResult struct {
	Body string
	Urls []string
	Err  error
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// 1. for that to work i need to run the function as a go routine,
	// 2.  when they are in outside goroutines i cant get the response directly
	// 3. 1 channel with a struct

	c := make(chan FetchResult)

	//the channel needs to be passed to go routine

	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	go fetcher.Fetch(url, c)

	//add w8 group TODO

	//when w8 group is ended, close channel and read?!

	for v := range c {
		if v.Err != nil {
			fmt.Println(v.Err)
			return
		}
		fmt.Printf("found: %s %q\n", v.Urls, v.Body)
		for _, u := range v.Urls {
			Crawl(u, depth-1, fetcher)
		}

	}

	return
}

func main() {

	start := time.Now()
	Crawl("https://golang.org/", 4, fetcher)

	fmt.Printf("time since start %v", time.Since(start))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string, c chan FetchResult) {

	if res, ok := f[url]; ok {
		c <- FetchResult{Body: res.body, Urls: res.urls, Err: nil}
	}
	c <- FetchResult{Body: "", Urls: nil, Err: fmt.Errorf("not found: %s", url)}
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
