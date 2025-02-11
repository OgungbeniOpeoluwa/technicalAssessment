package core

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessFilesThrowsErrorIfFileDoesntExists(t *testing.T) {
	filePath := "log.txt"
	keywords := []string{"INFO", "ERROR", "DEBUG"}

	_, err := ProcessLogFile(filePath, keywords)
	log.Print(err)
	assert.Error(t, err)

}

func TestProcessFilesReturnsEmptyIfFileIsEmpty(t *testing.T) {
	filePath := "/Users/tife/ConcurrentLogProcessor/log.txt"
	keywords := []string{"INFO", "ERROR", "DEBUG"}

	keywordCounts, err := ProcessLogFile(filePath, keywords)
	if err != nil {
		t.Errorf("expected %v got %v", keywordCounts, err)
	}
	log.Println(keywordCounts)
	assert.Empty(t, keywordCounts)

}

func TestProcessFiles(t *testing.T) {
	filePath := "/Users/tife/ConcurrentLogProcessor/log.txt"
	keywords := []string{"INFO", "ERROR", "DEBUG"}

	keywordCounts, err := ProcessLogFile(filePath, keywords)
	if err != nil {
		t.Errorf("expected %v got %v", keywordCounts, err)
	}
	log.Println(keywordCounts)
	assert.NotEmpty(t, keywordCounts)

}

func TestProcessFilesReturnResultInSortedForm(t *testing.T) {
	filePath := "/Users/tife/ConcurrentLogProcessor/log.txt"
	keywords := []string{"INFO", "ERROR", "DEBUG"}

	keywordCounts, err := ProcessLogFile(filePath, keywords)
	if err != nil {
		t.Errorf("expected %v got %v", keywordCounts, err)
	}
	log.Println(keywordCounts)
	assert.NotEmpty(t, keywordCounts)
	var isTrue = keywordCounts[0].Value > keywordCounts[1].Value
	assert.True(t, isTrue)

}
