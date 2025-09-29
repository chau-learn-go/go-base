package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Bài 1:")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine")
	}()
	wg.Wait()
	fmt.Println()

	fmt.Println("Bài 2:")
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	val := <-ch
	fmt.Println("Received:", val)
	fmt.Println()

	fmt.Println("Bài 3:")
	chBuf := make(chan int, 3)
	chBuf <- 1
	chBuf <- 2
	chBuf <- 3
	fmt.Println(<-chBuf, <-chBuf, <-chBuf)
	fmt.Println()

	fmt.Println("Bài 4:")
	nums := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			nums <- i
		}
		close(nums)
	}()

	worker := func(id int, ch <-chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for n := range ch {
			fmt.Printf("Worker %d received %d\n", id, n)
		}
	}

	wg.Add(2)
	go worker(1, nums, &wg)
	go worker(2, nums, &wg)
	wg.Wait()
	fmt.Println()

	fmt.Println("Bài 5:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "ping"
			time.Sleep(500 * time.Millisecond)
		}
	}()
	go func() {
		for {
			ch2 <- "pong"
			time.Sleep(700 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			}
		}
	}()
	time.Sleep(4 * time.Second)
	fmt.Println()

	fmt.Println("Bài 6:")
	a := make(chan int)
	b := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			a <- i
		}
		close(a)
	}()

	go func() {
		for n := range a {
			b <- n * n
		}
		close(b)
	}()

	for res := range b {
		fmt.Print(res, " ")
	}
	fmt.Println("\n")

	fmt.Println("Bài 7:")
	tasks := make(chan int, 10)
	results := make(chan int, 10)

	// tạo 10 task
	for i := 1; i <= 10; i++ {
		tasks <- i
	}
	close(tasks)

	workerPool := func(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
		defer wg.Done()
		for t := range tasks {
			fmt.Printf("Worker %d processing task %d\n", id, t)
			time.Sleep(200 * time.Millisecond) // giả lập xử lý
			results <- t * 2
		}
	}

	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go workerPool(i, tasks, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result:", r)
	}
}
