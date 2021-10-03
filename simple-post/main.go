package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "log"
	"net/http"
	"time"
	"math/rand"
	"sync"
)

const (
	url = "https://httpbin.org/post"
	contentType = "application/json"
)

type Event struct {
	Datetime	int64	`json:"datetime"`
    Username	string  `json:"username"`
	Text     	string  `json:"text"`
	SimNumber 	int 	`json:"simnumber"`
	GameMode	string  `json:"gamemode"`
}

func main() {

	var wg sync.WaitGroup

	for i:=0; i<5; i++ {

		json_data := generatePayload()

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

func generatePayload () []uint8 {

	modes := []string{
		"capture the flag",
		"battleroyale",
		"attack and defend",
		"creative"}

	values := Event{
		int64(time.Now().Unix()),
		"Alice", 
		"Test text message",
		rand.Intn(100),
		modes[rand.Intn(len(modes))]}

	json_data, err := json.Marshal(values)

    if err != nil {
        log.Fatal(err)
    }

	return json_data
}