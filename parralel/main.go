package main

import (
	"fmt"
	"sync"
	"time"
)

func fetchWeather(city string, ch chan string, wg *sync.WaitGroup) {

	if city == "berlin" {

		ch <- "20"
	}

	if city == "toronto" {

		ch <- "30"
	}

	defer wg.Done()

}

func main() {

	start := time.Now()

	citites := []string{"toronto", "berlin"}

	ch := make(chan string)

	var wg sync.WaitGroup

	for _, city := range citites {

		wg.Add(1)
		go fetchWeather(city, ch, &wg)
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
