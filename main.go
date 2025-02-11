package main

import (
	"CONCURRENTLOGPROCESSOR/core"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// run process log file
	var filePath string
	fmt.Print("Enter FilePath: ")
	fmt.Scan(&filePath)

	fmt.Print("Enter keywords (separated by spaces): ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	keywords := strings.Fields(scanner.Text())
	log.Println(keywords)

	keywordCounts, err := core.ProcessLogFile(filePath, keywords)
	if err != nil {
		fmt.Println("Error Processing log file:", err)
		return
	}
	log.Println(keywordCounts)

	// run prime panlidrome
	var N int
	fmt.Print("Enter N: ")
	fmt.Scan(&N)

	primePalindromes := core.FindPrimePalindromes(N)
	fmt.Printf("Sum of the first %d prime palindromic numbers: %d\n", N, primePalindromes)
}
