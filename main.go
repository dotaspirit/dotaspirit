package main

import (
	"fmt"
	"time"

	"github.com/dotaspirit/dotaspirit/queue"
)

func main() {
	fmt.Println("Starting...")
	ret := make(map[int64]int)
	for {
		fmt.Println("Cycle")
		queue.QueueMatches(ret)
		time.Sleep(time.Minute)
	}
}
