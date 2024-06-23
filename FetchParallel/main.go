package main

import (
	"fmt"
	"sync"
	"time"
)

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

	// if url == "https://golang.org/" {

	// 	ch <- []string{"20"}
	// }

	// if url == "berlin.com" {

	// 	ch <- []string{"30"}
	// }

	defer wg.Done()

}

func main() {

	start := time.Now()

	urls := []string{"https://golang.org/", "https://golang.org/pkg/"}

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

	for result := range ch {

		fmt.Println(result)
	}

	fmt.Println(time.Since(start))

}
