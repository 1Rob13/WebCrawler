package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c chan string, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel. results need to hit channel
	// TODO: Don't fetch the same URL twice. -> with context
	// This implementation doesn't do either:
	//	now := time.Now()
	if depth <= 0 {
		return
	}

	wg.Add(1)
	defer wg.Done()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("found: %s %q\n", body, urls)
	c <- fmt.Sprintf("found: %s %q\n", body, urls)

	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, c, wg)
	}

	//	fmt.Println(time.Since(now))
	return
}

func main() {

	now := time.Now()

	ch := make(chan string, 10)
	wg := sync.WaitGroup{}

	Crawl("https://golang.org/", 4, fetcher, ch, &wg)
	wg.Wait()

	//close(ch)

	for i := range ch {
		fmt.Println(i)

	}

	fmt.Println(time.Since(now))
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
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
