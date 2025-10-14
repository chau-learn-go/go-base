package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 5

	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			fmt.Println(num)
		}(i + 1)
	}
	wg.Wait()
}
