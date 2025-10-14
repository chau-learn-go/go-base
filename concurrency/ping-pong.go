package main

import (
	"sync"
)

func display(message string, wg *sync.WaitGroup) {
	println(message)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go display("Ping", &wg)
	}
	for i := 0; i < 10; i++ {
		go display("Pong", &wg)
	}
	wg.Wait()
}
