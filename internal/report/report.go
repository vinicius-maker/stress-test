package report

import (
	"fmt"
	"github.com/vinicius-maker/stress-test/internal/worker"
	"time"
)

func GenerateReport(results <-chan worker.Result) map[int]int {
	statusCounts := make(map[int]int)

	for result := range results {
		statusCounts[result.StatusCode]++
	}

	return statusCounts
}

func PrintStatus(statusCounts map[int]int, elapsed time.Duration, requests *int) {
	fmt.Printf("\nTempo total: %v\n", elapsed)
	fmt.Printf("Requests totais: %d\n", *requests)
	fmt.Printf("Status HTTP 200: %d\n", statusCounts[200])

	for status, count := range statusCounts {
		if status != 200 {
			fmt.Printf("Outros Status HTTP (%d): %d\n", status, count)
		}
	}
}
