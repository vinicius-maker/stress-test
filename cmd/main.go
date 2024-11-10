package main

import (
	"flag"
	"fmt"
	"github.com/vinicius-maker/stress-test/internal/report"
	"github.com/vinicius-maker/stress-test/internal/worker"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 10, "Número total de requests")
	concurrency := flag.Int("concurrency", 1, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: A URL é obrigatória.")
		flag.Usage()
		return
	}

	results := make(chan worker.Result, *requests)
	var wg sync.WaitGroup

	start := time.Now()

	requestsPerWorker := *requests / *concurrency
	remainder := *requests % *concurrency

	for i := 0; i < *concurrency; i++ {
		reqs := requestsPerWorker
		if i < remainder {
			reqs++
		}
		wg.Add(1)
		go worker.Execute(*url, reqs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	statusCounts := report.GenerateReport(results)

	elapsed := time.Since(start)
	report.PrintStatus(statusCounts, elapsed, requests)
}
