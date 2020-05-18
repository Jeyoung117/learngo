package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFailed = errors.New("Request failed")

func main() {
	results := make(map[string]string)
	c := make(chan result)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}

func hitURL(url string, c chan<- result) {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "Failed"
	}
	c <- result{url: url, status: status}
}

// func main() {
// 	c := make(chan string)
// 	people := []string{"nico", "flynn"}
// 	for _, person := range people {
// 		go isSexy(person, c)
// 	}

// 	fmt.Println("Waiting for messages")
// 	for i := 0; i < len(people); i++ {
// 		fmt.Println("Waiting for ", i)
// 		fmt.Println(<-c)
// 	}
// }

// func isSexy(person string, c chan string) {
// 	time.Sleep(time.Second * 3)
// 	c <- person + "is Sexy"

// }
