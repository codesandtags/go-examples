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
		"http://localhost",
		"http://localhost",
		"http://localhost",
		"http://localhost",
		"http://localhost",
		"http://amazon.com",
		"http://localhost",
		"http://zalando.com",
	}

	c := make(chan string)

	fmt.Println("\n Testing the http in serial approach")
	for _, link := range links {
		// Serial approach
		checkLink(link)
	}

	fmt.Println("\n Testing the http in parallel approach")
	for _, link := range links {
		// Paralell approach
		go checkLinkRoutine(link, c)
		fmt.Println(<- c)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return 
	}

	fmt.Println(link, "is up!")	
} 

func checkLinkRoutine(link string, c chan string) {
	_, err := http.Get(link)
	time.Sleep(time.Second)

	if err != nil {
		c <- link + " Might be down I think"
		return 
	}

	c <- link + " Yep its up!"
} 