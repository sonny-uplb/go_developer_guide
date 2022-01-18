package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://stackoverflow.com",
		"http://facebook.com",
		"http://google.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(15 * time.Second)
			checkLink(link, c)
		}(l)
	}

}

func checkLink(link string, c chan string) {

	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down"
		return
	}

	fmt.Println(link, "is up!")
	c <- "it is up"
}
