package report

import (
	"fmt"
	"github.com/vinicius-maker/stress-test/internal/worker"
)

func GenerateReport(results <-chan worker.Result) map[int]int {
	statusCounts := make(map[int]int)

	for result := range results {
		statusCounts[result.StatusCode]++
	}

	return statusCounts
}

func PrintStatus(statusCounts map[int]int) {
	for status, count := range statusCounts {
		if status != 200 {
			fmt.Printf("Outros Status HTTP %d: %d\n", status, count)
		} else {
			fmt.Printf("Outros Status HTTP: %d\n", 0)
		}
	}
}
