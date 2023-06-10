package goRoutines

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func WebCralwer() {
	pages := make(chan Page)
	//Refactored into a loop
	urls := []string{"https://example.com", "https://golang.org/", "https://golang.org/doc"}

	for _, url := range urls {
		go responseSize(url, pages)
	}

	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

	// Difficult to tell which size goes with which url. So we need to refactor in structs.
	//Running these as Go routines.
	//Go routines do NOT return values. So we gotta use channels
	// Passing in the channel as an argument so we can load it up functioning as a return of sorts for a go routine.
	//go responseSize("https://example.com", sizesChan)
	//go responseSize("https://golang.org/", sizesChan)
	//go responseSize("https://golang.org/doc", sizesChan)
	//Because there are three sends, there will be three receives.
	//fmt.Println(<-sizesChan)
	//fmt.Println(<-sizesChan)
	//fmt.Println(<-sizesChan)
}

func responseSize(url string, channel chan Page) {

	fmt.Println("Getting", url)
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	//Going to close the connection after the main function finishes
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	//Basically loading the receive or "return" into a channel
	channel <- Page{URL: url, Size: len(body)}
}
