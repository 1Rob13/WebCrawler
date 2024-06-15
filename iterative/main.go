package main

import (
	"fmt"
	"slices"
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

	for j := 0; j < 1; j++ {

		fmt.Printf("iteration: ( %v )\n", j)
		fmt.Printf("urls: ( %v )\n", urls)
		fmt.Printf("cached urls: ( %v )\n", urlsCache)

		if len(urls) == 0 {
			break
		}

		for _, url := range urls {

			fmt.Println("reached loop")
			//cache check
			if slices.Contains(urlsCache, url) {
				continue
			}

			go func(msg string) {
				fmt.Println("reached concurrent")

				body, urls, err := fetcher.Fetch(urls[0])
				fmt.Printf("fetched body: (%s), found URLs: (%v)\n", body, urls)
				urlsCache = append(urlsCache, urls[0])
				if err != nil {
					return
				}
			}("going")

		}

	}

	fmt.Println(time.Since(now))

}
