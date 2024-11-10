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

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go worker.Execute(*url, *requests / *concurrency, results, &wg)
	}

	wg.Wait()
	close(results)

	statusCounts := report.GenerateReport(results)

	elapsed := time.Since(start)
	fmt.Printf("\nTempo total: %v\n", elapsed)
	fmt.Printf("Requests totais: %d\n", *requests)
	fmt.Printf("Status HTTP 200: %d\n", statusCounts[200])
	report.PrintStatus(statusCounts)
}
