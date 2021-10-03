package main

import (
    "fmt"
    //"io/ioutil"
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

		//body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("[%v]\tRequested %v\n", i, url)

	}

	wg.Wait()
}