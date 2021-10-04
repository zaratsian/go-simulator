package simulator

import (
	"math/rand"
	"log"
	"time"
	"encoding/json"
)

type Event struct {
	Datetime	int64	`json:"datetime"`
    Username	string  `json:"username"`
	Text     	string  `json:"text"`
	SimNumber 	int 	`json:"simnumber"`
	GameMode	string  `json:"gamemode"`
}

func GeneratePayload () []uint8 {

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