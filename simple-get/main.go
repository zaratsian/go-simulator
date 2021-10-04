package main

import (
    "fmt"
    "log"
	"net/http"
	"sync"
)

const (
	url = "https://httpbin.org/get"
)

func main() {

	var wg sync.WaitGroup

	for i:=0; i<5; i++ {

		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		defer resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%v]\tRequested %v\n", i, url)

	}

	wg.Wait()
}