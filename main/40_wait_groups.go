package main

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("worker %d start\n", id)
	time.Sleep(time.Second)
	fmt.Printf("workder %d end\n", id)
}

func main() {
	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker3(i)
		}()
	}

	wg.Wait()
}
