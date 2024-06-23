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
		fetcher = fakeFetcher{
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

		firstUrl  = "https://golang.org/"
		urls      = []string{}
		urlsCache = []string{}
	)

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

		//parallisable from here------------------------------------------------
		//
		//
		// -- cant be done inside the loop because that would be done double for each url
		//  it will be reset during next iteration
		// chURLS := make(chan []string)

		//is this for loop taking more than normal? it should be at init fixed how many urls
		for _, url := range urls {

			urlsCache = append(urlsCache, url) // this should be doable after because the barrier of the semaphore happens later anyway

			fmt.Printf("url selected in urls: ( %s)\n", url)

			fmt.Printf("attempt fetch of %s\n", url)

			body, routinesUrls, err := fetcher.Fetch(url)
			if err != nil {
				fmt.Printf("routine fetched body: %v:", err)
			}
			fmt.Printf("routine fetched body: (%s), found URLs: (%v)\n\n\n\n", body, routinesUrls)

			//if new urls was already visited do not add it to urls
			for _, url := range routinesUrls {

				if slices.Contains(urlsCache, url) {
					continue
				}
				urls = append(urls, url)
			}

			//send filtered urls to channel
			//chURLS <- urls

			//done channels do not work because there are 2 routines that will both end the waiting
			// done <- 0

			//
			//
			//
			//parallisable to here------------------------------------------------

		}

		// for value := range chURLS {

		// 	fmt.Println(value) // Process the received values

		// }

		//cant pull this in for now, will end up new fetched, it kinda works semaphore like because it will wait for other processes to be done
		//this also saves extra lookups this way
		newUrls := []string{}
		//assert that urls is only new urls
		for _, url := range urls {

			if slices.Contains(urlsCache, url) {
				continue
			}

			newUrls = append(newUrls, url)

		}

		urls = newUrls

	}

	fmt.Printf("\nfetched URLS: %v\n", urls)
	fmt.Printf("\ncached URLS: %v\n", urlsCache)

	if len(urlsCache) != 5 {
		fmt.Println("ERROR--------------urlsCache not 5----------------ERROR------------------------------ERROR------------------------------ERROR------------------------------")

	}

	if len(urls) != 0 {
		fmt.Println("ERROR--------------urls not 0----------------ERROR------------------------------ERROR------------------------------ERROR------------------------------")

	}

	fmt.Println(time.Since(now))

}
