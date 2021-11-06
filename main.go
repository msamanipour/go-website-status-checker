package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://vadaad.com",
		"http://vadaad.ir",
	}

	c := make(chan string)
	for _, link := range links {
		go checkStatus(link, c)
	}
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkStatus(link, c)
		}(l)
	}
}

func checkStatus(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println("Site", l, "maybe down !")
		c <- l
		return
	}
	fmt.Println("Site", l, "is up !")
	c <- l
}
