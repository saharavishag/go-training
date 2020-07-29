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
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		// use go routine (like async way)
		go checkLink(link, c)
		// similar to fork
		// go routine in go cannot work by itself
		// we'll use channels instead
	}

	// for {
	// 	go checkLink(<-c, c)
	// 	// fmt.Println(<-c)
	// }

	for l := range c {
		// lambda
		go func(link string) {
			time.Sleep(time.Second)
			checkLink(link, c)
		}(l)

		// fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		c <- "Might be down"
		return
	}
	fmt.Println(link, "is up!")
	c <- link
}
