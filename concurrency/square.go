package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)

	for _, num := range nums {
		go func() {
			ch <- num * num
		}()
	}

	for range nums {
		fmt.Println(<-ch)
	}
}
