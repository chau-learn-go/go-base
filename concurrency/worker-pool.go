package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	jobs := make(chan int)
	results := make(chan int)
	workerLimit := 4
	wg := sync.WaitGroup{}

	for i := 0; i < workerLimit; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case job, ok := <-jobs:
					if !ok {
						return
					}
					results <- job * job
				}
			}
		}()
	}

	go func() {
		for _, num := range nums {
			jobs <- num
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	sum := 0
	for result := range results {
		sum += result
	}

	fmt.Println(sum)
}
