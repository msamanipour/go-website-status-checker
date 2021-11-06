package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http:vadaa.com",
		"http://vadaad.com",
	}

	for _, link := range links {
		checkStatus(link)
	}
}

func checkStatus(l string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println("Site", l, "maybe down !")
	}
	fmt.Println("Site", l, "is up !")
}
