package main

import (
    "fmt"
	"sync"
	"time"

	"github.com/zaratsian/go-simulator/simulator"
)

const (
	iterations = 20
)

func main() {

	start_time := time.Now().Unix()

	var wg sync.WaitGroup

	for i:=0; i<=iterations; i++ {
		payload := simulator.GeneratePayload()
		fmt.Printf("%v\n",string(payload))
	}

	wg.Wait()

	runtime := time.Now().Unix() - start_time
	fmt.Printf("Runtime: %v\t seconds\n", runtime)
	if runtime > 0 {
		fmt.Printf("Speed:   %v\t events/second\n", iterations/runtime)
	} else {
		fmt.Printf("Speed:   %v\t events/second\n", iterations)
	}
}