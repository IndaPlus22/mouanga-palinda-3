package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
)

const DataFile = "loremipsum.txt"

// "inspiration" from https://gosamples.dev/remove-non-alphanumeric/
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// Return the word frequencies of the text argument.
func WordCount(text string) map[string]int {
	// partly copy pasted from my palinda-1 "maps" file
	freqs := make(map[string]int)

	words := strings.Fields(text)

	words_len := len(words)

	// update the map
	for i := 0; i < words_len; i++ {
		words[i] = clearString(words[i])
		words[i] = strings.ToLower(words[i])
		freqs[words[i]] += 1
	}
	return freqs
}

// Benchmark how long it takes to count word frequencies in text numRuns times.
//
// Return the total time elapsed.
func benchmark(text string, numRuns int) int64 {
	start := time.Now()
	for i := 0; i < numRuns; i++ {
		WordCount(text)
	}
	runtimeMillis := time.Since(start).Nanoseconds() / 1e6

	return runtimeMillis
}

// Print the results of a benchmark
func printResults(runtimeMillis int64, numRuns int) {
	fmt.Printf("\namount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	// read in DataFile as a string called data

	data, _ := ioutil.ReadFile(DataFile)

	//fmt.Printf("%#v", WordCount(string(data)))

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
