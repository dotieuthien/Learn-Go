package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string {
		"http://google.com",
		"http://facebook.com",
		"http://amazon.com",
		"htp://stackoverflow.com",
	}
	start_time := time.Now()
	for _, link := range links {
		checkLink(link)
	}
	fmt.Printf("Execution time %s\n", time.Since(start_time))

	start_time = time.Now()
	// create a string channel
	c := make(chan string)

	for _, link := range links {
		// call go routine to execute checkLink func
		// go routine need channel for communication with main routine
		go checkLinks(link, c)
	}
	// listen the mesage from channel
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
	fmt.Printf("Execution time %s\n", time.Since(start_time))
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}
	fmt.Println(link, " is up!")
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