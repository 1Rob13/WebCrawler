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

type PubSub[T any] struct {
	subscribers []chan T
	mu          sync.RWMutex
	closed      bool
}

func NewPubSub[T any]() *PubSub[T] {
	return &PubSub[T]{
		mu: sync.RWMutex{},
	}
}

func (s *PubSub[T]) Subscribe() <-chan T {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return nil
	}

	r := make(chan T)

	s.subscribers = append(s.subscribers, r)

	return r
}

func (s *PubSub[T]) Publish(value T) {

	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.closed {
		return
	}

	for _, ch := range s.subscribers {
		ch <- value
	}
}

func (s *PubSub[T]) Close() {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.closed {
		return
	}

	for _, ch := range s.subscribers {
		close(ch)
	}
	s.closed = true
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel. results need to hit channel
	// TODO: Don't fetch the same URL twice. -> with context
	// This implementation doesn't do either:
	//now := time.Now()
	if depth <= 0 {
		return
	}

	//i think he panics because he does not control the recoursive calls and they send after the chan recieve is down

	//collect channels up?

	_, urls, err := fetcher.Fetch(url)
	if err != nil {
		return
	}

	//fmt.Printf("found: %s %q\n", body, urls)

	// for _, u := range urls {
	// 	go Crawl(u, depth-1, fetcher, c)

	// }

	for _, u := range urls {

		Crawl(u, depth-1, fetcher)
	}

	//	fmt.Println(time.Since(now))
	return
}

func main() {

	now := time.Now()

	ps := NewPubSub[string]()

	wg := sync.WaitGroup{}

	s1 := ps.Subscribe()

	go func() {

		wg.Add(1)

		for {

			select {

			case val, ok := <-s1:

				if !ok {

					fmt.Print("sub 1 , exiting")
					wg.Done()
					return
				}
				fmt.Println("sub 1, value,", val)
			}
		}
	}()

	// ps.Publish("one")
	// ps.Publish("two")
	// ps.Publish("three")
	// ps.Publish("four")

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
