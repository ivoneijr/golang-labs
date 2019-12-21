package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// <-chan = read only channel

func title(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func forward(from <-chan string, to chan string) {
	for {
		to <- <-from
	}
}

func group(first, second <-chan string) <-chan string {
	c := make(chan string)

	go forward(first, c)
	go forward(second, c)

	return c
}

func main() {
	c := group(
		title("https://www.google.com", "http://ivoneijr.fun/"),
		title("https://www.google.com", "https://www.youtube.com"),
	)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
