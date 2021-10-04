package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
	"net/http"
	"sync"

	"github.com/zaratsian/go-simulator/simulator"
)

const (
	url = "https://httpbin.org/post"
	contentType = "application/json"
)

func main() {

	var wg sync.WaitGroup

	for i:=0; i<5; i++ {

		json_data := simulator.GeneratePayload()

		resp, err := http.Post(url, contentType,
			bytes.NewBuffer(json_data))

		if err != nil {
			log.Fatal(err)
		}

		var res map[string]interface{}

		json.NewDecoder(resp.Body).Decode(&res)

		fmt.Println(res["json"])
	}

	wg.Wait()
}