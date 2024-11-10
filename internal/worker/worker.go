package worker

import (
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
}

func Execute(url string, requests int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	for i := 0; i < requests; i++ {
		resp, err := client.Get(url)
		if err != nil {
			results <- Result{StatusCode: 0}
			continue
		}

		results <- Result{StatusCode: resp.StatusCode}
		resp.Body.Close()
	}
}
