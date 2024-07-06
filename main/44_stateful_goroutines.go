package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	resp  chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		state := map[int]int{}
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.value
				write.resp <- true
			}
		}
	}()

	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				readResult := <-read.resp
				fmt.Println("read result = ", readResult)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}

		}()
	}

	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					resp:  make(chan bool),
				}
				fmt.Println("write to channel: ", write)
				writes <- write
				if success := <-write.resp; success {
					atomic.AddUint64(&writeOps, 1)
				} else {
					fmt.Println("write failed: ", write)
				}
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("read ops: ", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("write ops: ", writeOpsFinal)
}
