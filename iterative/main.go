package main

import (
	"fmt"
	"reflect"
	"slices"
	"sync"
	"time"
)

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

func main() {

	now := time.Now()

	var (
		urls      []string
		urlsCache []string
		wg        sync.WaitGroup
		fetcher   = fakeFetcher{
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
			}}
	)

	firstUrl := "https://golang.org/"

	urls = append(urls, firstUrl)

	for j := 0; j < 3; j++ {
		//for {
		fmt.Println("-------------------------")

		fmt.Printf("iteration: ( %v )\n", j)
		fmt.Printf("urls: ( %v )\n", urls)
		fmt.Printf("cached urls: ( %v )\n\n\n\n", urlsCache)

		if len(urls) == 0 {
			break
		}

		for _, url := range urls {

			//cache check
			if slices.Contains(urlsCache, url) {
				continue
			}
			fmt.Printf("url selected in urls: ( %s)\n", url)

			wg.Add(1)
			go func(msg string) {

				fmt.Printf("attempt fetch of %s\n", url)

				body, routinesUrls, err := fetcher.Fetch(url)
				fmt.Printf("routine fetched body: (%s), found URLs: (%v)\n\n\n\n", body, routinesUrls)
				urlsCache = append(urlsCache, url)

				urls = append(urls, routinesUrls...)

				if err != nil {
					return
				}

			}("going")

			wg.Done()

			// this sleep needs asserts that the command above can be printed
			//the main program does not wait for the goroutines to finish

			// fmt.Println("-------------------------")

			// fmt.Printf("iteration: ( %v )\n", j)
			// fmt.Printf("urls: ( %v )\n", urls)
			// fmt.Printf("cached urls: ( %v )\n", urlsCache)

		}

		wg.Wait()

		if reflect.DeepEqual(urls, urlsCache) {
			fmt.Println("reached BREAK")
			break
		}

	}

	fmt.Println(time.Since(now))

}
