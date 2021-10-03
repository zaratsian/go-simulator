package main

import (
    "fmt"
	"sync"
	"time"
)

const (
	iterations = 200000
)

func main() {

	start_time := time.Now().Unix()

	var wg sync.WaitGroup

	for i:=0; i<=iterations; i++ {
		j := ((i + 1000000) * 10) / 5
		fmt.Printf("%v\n",j)
	}

	wg.Wait()

	runtime := time.Now().Unix() - start_time
	fmt.Printf("Runtime: %v\n seconds", runtime)
	fmt.Printf("Speed:   %v\n events/second", iterations/runtime)
}