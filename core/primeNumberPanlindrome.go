package core

import (
	"sync"
)

// isPrime checks if a number is prime.
func IsPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

// isPalindrome checks if a number is a palindrome.
func isPalindrome(num int) bool {
	original := num
	reversed := 0
	for num > 0 {
		reversed = reversed*10 + num%10
		num /= 10
	}
	return original == reversed
}

// findPrimePandlidrome based on the numbers that was inouted
func FindPrimePalindromes(N int) int {
	results := make(chan int)
	stop := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		num := 2
		for {
			select {
			case <-stop:
				return
			default:
				if IsPrime(num) && isPalindrome(num) {
					select {
					case results <- num:
					case <-stop:
						return
					}
				}
				num++
			}
		}
	}()

	// Start a Goroutine to collect the first N prime palindromic numbers
	primePalindromes := make([]int, 0, N)
	go func() {
		for num := range results {
			primePalindromes = append(primePalindromes, num)
			if len(primePalindromes) == N {
				close(stop)
				break
			}
		}
	}()

	wg.Wait()
	close(results)

	// Calculate and return the sum
	return calculateSum(primePalindromes)
}

// calculateSum calculates the sum of a slice of integers.
func calculateSum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
