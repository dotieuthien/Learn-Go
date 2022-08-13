package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"http://stackoverflow.com",
	}
	// create a string channel
	c := make(chan string)

	for _, link := range links {
		// call go routine to execute checkLink func
		go checkLinks(link, c)
	}
	// listen the mesage from channel
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLinks(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		// send message to main routine
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, " is up!")
	c <- "Yep its up"
}