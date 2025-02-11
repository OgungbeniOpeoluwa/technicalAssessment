package core

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
)

type ProcessedLogResponse struct {
	Key   string
	Value int
}

// ProcessLogFile reads the log file in chunks and processes it concurrently.
func ProcessLogFile(filePath string, keywords []string) ([]ProcessedLogResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("file not found: %s", filePath)
		}
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	linesCh := make(chan string)

	resultsCh := make(chan map[string]int)

	var wg sync.WaitGroup

	numWorkers := 4
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			CountKeywords(linesCh, resultsCh, keywords)
		}()
	}

	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			linesCh <- scanner.Text()
		}
		close(linesCh)

		wg.Wait()
		close(resultsCh)
	}()

	keywordCounts := make(map[string]int)
	for counts := range resultsCh {
		for keyword, count := range counts {
			keywordCounts[keyword] += count
		}
	}

	return SortKeywordByFrequencyInDecendingOrder(keywordCounts), nil
}

// CountKeywords processes lines and counts keyword occurrences.
func CountKeywords(linesCh <-chan string, resultsCh chan<- map[string]int, keywords []string) {
	localCounts := make(map[string]int)
	for line := range linesCh {
		for _, keyword := range keywords {
			if strings.Contains(strings.ToUpper(line), strings.ToUpper(keyword)) {
				localCounts[keyword]++
			}
		}
	}
	resultsCh <- localCounts
}

func SortKeywordByFrequencyInDecendingOrder(keywordCounts map[string]int) []ProcessedLogResponse {
	var sortedCounts []ProcessedLogResponse
	for k, v := range keywordCounts {
		sortedCounts = append(sortedCounts, ProcessedLogResponse{k, v})
	}
	sort.Slice(sortedCounts, func(i, j int) bool {
		return sortedCounts[i].Value > sortedCounts[j].Value
	})
	log.Println(sortedCounts)
	return sortedCounts

}
