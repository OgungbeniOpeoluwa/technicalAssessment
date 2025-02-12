# Technical Assessment: Concurrent Log Processor and Prime Palindrome Finder

This project is a Go-based tool that performs two main tasks:

1. **Concurrent Log Processing**: Analyzes a log file to count occurrences of specific keywords using concurrency.
2. **Prime Palindrome Finder**: Finds the first `N` prime palindromic numbers and calculates their sum.

## Features

### 1. Concurrent Log Processing

- Reads a log file in chunks using Goroutines.
- Counts occurrences of specified keywords concurrently.
- Outputs keyword frequencies in descending order.

### 2. Prime Palindrome Finder

- Finds the first `N` prime palindromic numbers using concurrency.
- Calculates the sum of these numbers.
- Efficiently handles large inputs using Goroutines and channels.

## Prerequisites

Before running the project, ensure you have the following installed:

- **Go**: Download and install Go.
- **Log File**: A log file to analyze (e.g., `log.txt`).

## Installation

1. Clone the repository:

   ```
   git clone https://github.com/OgungbeniOpeoluwa/technicalAssessment.git
   ```

2. Navigate to the project directory:
   ```
   cd technicalAssessment
   ```

## Usage

### 1. Concurrent Log Processing

To analyze a log file and count keyword occurrences:

1. Run the program:
   - go run main.go
2. Enter the file path when prompted:
   - Enter FilePath: log.txt
3. Enter the keywords (separated by comma):
   - Enter keywords (separated by comma): INFO,ERROR,DEBUG
4. The program will output the keyword counts:

   - {INFO 3}
   - {ERROR 2}
   - {DEBUG 1}

### 2. Prime Palindrome Finder

To find the first N prime palindromic numbers and calculate their sum:

1. Run the program:
   - go run main.go
2. Enter the value of N:
   - Enter N: 5
3. The program will output the sum:
   - Sum of the first 5 prime palindromic numbers: 28

## Example Input and Output

### Example 1: Log Processing

#### Input:

- Enter FilePath: log.txt

- Enter keywords (separated by comma): INFO,ERROR,DEBUG

#### Output:

- {DEBUG 1}
- {ERROR 2}
- {INFO 3}

## Example 2: Prime Palindrome Finder

### Input:

- Enter N: 5

### Output:

- Sum of the first 5 prime palindromic numbers: 28

## Project Structure

```
technicalAssessment/
├── core/
│ ├── file_processor.go
│ ├── prime_palindrome.go
│ ├── file_processor_test.go
│ ├── prime_palindrome_test.go
│
├── main.go # Main program entry point
├── README.md
├── log.txt # Sample log file
├── logs.txt # Empty Sample log file for test
├── go.mod
├── go.sum
```

## Concurrency Details

### Log Processing

**Uses Goroutines to read the log file in chunks**.

**Uses channels to communicate keyword counts between Goroutines**.

### Prime Palindrome Finder

**Uses Goroutines to check numbers for prime and palindromic properties concurrently**.

**Uses channels to collect results and signal completion**.

# Error Handling

The program handles the following errors gracefully:

**File not found**

**Error opening the log file**.

# License

This project is licensed under the MIT License. See the LICENSE file for details.
