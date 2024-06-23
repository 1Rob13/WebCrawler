package main

import (
	"fmt"
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
		fetcher = fakeFetcher{
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

	body, routinesUrls, err := fetcher.Fetch(firstUrl)
	if err != nil {
		fmt.Printf("routine fetched body: %v:", err)
	}
	fmt.Printf("routine fetched body: (%s), found URLs: (%v)\n\n\n\n", body, routinesUrls)

	// this sleep needs asserts that the command above can be printed
	//the main program does not wait for the goroutines to finish

	// fmt.Println("-------------------------")

	// fmt.Printf("iteration: ( %v )\n", j)
	// fmt.Printf("urls: ( %v )\n", urls)
	// fmt.Printf("cached urls: ( %v )\n", urlsCache)

	fmt.Println(time.Since(now))
	fmt.Println(time.Since(now) * 3)

}
