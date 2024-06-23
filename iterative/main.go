package main

import (
	"fmt"
	"sync"
	"time"
)

func FetchParallel(urls []string) []string {

	fmt.Println("got called")
	fmt.Println(urls)

	ch := make(chan []string)

	var wg sync.WaitGroup

	for _, city := range urls {

		wg.Add(1)
		go fetch(city, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	resultUrls := []string{}

	for result := range ch {

		fmt.Println(result)
		resultUrls = append(resultUrls, result...)
	}

	return resultUrls

}

type fakeResult struct {
	body string
	urls []string
}

type fakeFetcher map[string]*fakeResult

func fetch(url string, ch chan []string, wg *sync.WaitGroup) {

	fetcher := fakeFetcher{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/", // a miss
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",     // already traversed
				"https://golang.org/cmd/", // a miss
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
		}}

	if res, ok := fetcher[url]; ok {
		ch <- res.urls
	}

	defer wg.Done()

}

func main() {

	now := time.Now()

	var (
		firstUrl = "https://golang.org/"
		urls     = []string{}
		//visited  = []string{}
	)

	urls = append(urls, firstUrl)

	for j := 0; j < 3; j++ {

		//	fmt.Println()

		// newUrls := []string{}

		// for _, url := range urls {

		// 	if slices.Contains(visited, url) {

		// 		continue
		// 	}

		// 	newUrls = append(newUrls, url)
		// }

		// urls = newUrls

		urlsNew := FetchParallel(urls)

		urls = append(urls, urlsNew...)

		//visited = append(visited, urls...)

	}

	// fmt.Printf("\nfetched URLS: %v\n", urls)
	// fmt.Printf("\ncached URLS: %v\n", urlsCache)

	// if len(urlsCache) != 5 {
	// 	fmt.Println("ERROR--------------urlsCache not 5----------------ERROR------------------------------ERROR------------------------------ERROR------------------------------")

	// }

	// if len(urls) != 0 {
	// 	fmt.Println("ERROR--------------urls not 0----------------ERROR------------------------------ERROR------------------------------ERROR------------------------------")

	// }

	fmt.Println(time.Since(now))

}
