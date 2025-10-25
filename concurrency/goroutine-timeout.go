package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	go func() {
		fmt.Println("Job started")
		time.Sleep(3 * time.Second)
		done <- "Job finished"
	}()

	select {
	case result := <-done:
		fmt.Println(result)
	case <-timer.C:
		fmt.Println("Job timeout! Canceled")
	}
}
