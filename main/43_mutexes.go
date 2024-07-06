package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mutex    sync.Mutex
	counters map[string]int
}

func (c *Container) Inc(name string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.counters[name]++
}

func main() {
	c := Container{
		counters: map[string]int{},
	}
	c.counters["a"] = 0
	c.counters["b"] = 0

	var wg sync.WaitGroup

	doInc := func(name string, n int) {
		for range n {
			c.Inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	doInc("a", 5)
	doInc("b", 5)
	doInc("a", 10)
	wg.Wait()
	fmt.Println(c.counters)
}
