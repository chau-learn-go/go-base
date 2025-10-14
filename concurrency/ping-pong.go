package main

import (
	"fmt"
	"sync"
)

func main() {
	pingCh := make(chan bool)
	pongCh := make(chan bool)
	n := 5

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-pingCh
			fmt.Println("Ping")
			pongCh <- true
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-pongCh
			fmt.Println("Pong")
			if i < n-1 {
				pingCh <- true
			}
		}
	}()

	pingCh <- true
	wg.Wait()

	fmt.Println("Done")
}
