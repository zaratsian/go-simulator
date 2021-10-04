package main

import (
        "context"
		"fmt"
		"time"
		"math/rand"
		"encoding/json"
		"log"
		"sync"

        "cloud.google.com/go/pubsub"
)

const (
	gcp_project_id  = "globalgame"
	pubsub_topic	= "ingest-text"
)

type Event struct {
	Datetime	int64	`json:"datetime"`
    Username	string  `json:"username"`
	Text     	string  `json:"text"`
	SimNumber 	int 	`json:"simnumber"`
	GameMode	string  `json:"gamemode"`
}

func main () {

	var wg sync.WaitGroup

	for i:=0; i<20; i++ {
		msg := generatePayload()
		publish(gcp_project_id, pubsub_topic, msg)
	}

	wg.Wait()

}

func publish(projectID string, topicID string, msg []uint8) error {

        ctx := context.Background()
        client, err := pubsub.NewClient(ctx, projectID)
        if err != nil {
                return fmt.Errorf("pubsub.NewClient: %v", err)
        }
        defer client.Close()

        t := client.Topic(topicID)
        result := t.Publish(ctx, &pubsub.Message{
                Data: []byte(msg),
        })
        // Block until the result is returned and a server-generated
        // ID is returned for the published message.
        id, err := result.Get(ctx)
        if err != nil {
                return fmt.Errorf("Get: %v", err)
        }
        fmt.Printf("Published a message; msg ID: %v\n", id)
        return nil
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