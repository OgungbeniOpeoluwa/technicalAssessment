package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatNumberIsNotAPrimeNumber(t *testing.T) {
	num := 4
	assert.False(t, IsPrime(num))
}

func TestThatNumberIsAPrimeNumber(t *testing.T) {
	num := 2
	assert.True(t, IsPrime(num))
}

func TestThatNumber0IsNotAPrimeNumber(t *testing.T) {
	num := 0
	assert.False(t, IsPrime(num))
}

func TestThatNumberIsAPanlidrome(t *testing.T) {
	num := 121
	assert.True(t, isPalindrome(num))
}

func TestThatNumberIsNotAPanlidrome(t *testing.T) {
	num := 122
	assert.False(t, isPalindrome(num))
}

func TestFind5AnNPrimePanlidrom(t *testing.T) {
	num := 5
	assert.Equal(t, FindPrimePalindromes(num), 28)
}

func TestFind10AsNPrimePanlidrom(t *testing.T) {
	num := 10
	assert.Equal(t, FindPrimePalindromes(num), 783)
}
