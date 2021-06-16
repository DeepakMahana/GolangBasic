package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.in",
	}

	// Create a channel
	c := make(chan string)

	for _, link := range links {
		// Creating child goroutines for all separate urls and pass the channel
		go checkLink(link, c)
	}

	// Loop when a channel recieves a msg
	// for l := range c {
	// 	go checkLink(l, c)
	// }

	// Using Fucntion Literal (Anonymous Func)
	for l := range c {
		go func(link string) {
			// Pause for 2 sec
			time.Sleep(2 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down !")
		c <- link
		return
	}
	fmt.Println("Working Url ", link)
	c <- link
}
