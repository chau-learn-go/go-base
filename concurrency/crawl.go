package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	Index int
	URL   string
}

type Result struct {
	Index      int
	URL        string
	StatusCode int
	Err        error
}

func worker(ctx context.Context, id int, client *http.Client, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		reqCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, job.URL, nil)
		if err != nil {
			results <- Result{Index: job.Index, URL: job.URL, Err: err}
			cancel()
			continue
		}

		resp, err := client.Do(req)
		if err != nil {
			results <- Result{Index: job.Index, URL: job.URL, Err: err}
			cancel()
			continue
		}

		_ = resp.Body.Close()
		results <- Result{Index: job.Index, URL: job.URL, StatusCode: resp.StatusCode}
		cancel()
	}
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.github.com",
		"https://golang.org",
		"https://httpbin.org/status/418",
		"https://invalid.local",
	}

	numWorkers := 4
	jobs := make(chan Job, len(urls))
	results := make(chan Result, len(urls))

	client := &http.Client{
		Timeout: 0,
	}

	ctx := context.Background()
	var wg sync.WaitGroup

	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, client, jobs, results, &wg)
	}

	go func() {
		for i, u := range urls {
			jobs <- Job{Index: i, URL: u}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		if res.Err != nil {
			fmt.Printf("[%d] %s -> ERROR: %v\n", res.Index, res.URL, res.Err)
		} else {
			fmt.Printf("[%d] %s -> %d\n", res.Index, res.URL, res.StatusCode)
		}
	}
}
