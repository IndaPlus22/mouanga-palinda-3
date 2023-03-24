package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"regexp"
	"strings"
	"sync"
	"time"
)

const DataFile = "loremipsum.txt"
const threads = 100

// "inspiration" from https://gosamples.dev/remove-non-alphanumeric/
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func count_partial(words []string, low int, high int, ch chan map[string]int, wg *sync.WaitGroup) {
	freqs := make(map[string]int)
	for i := low; i < high; i++ {
		words[i] = clearString(words[i])
		words[i] = strings.ToLower(words[i])
		freqs[words[i]] += 1
	}

	ch <- freqs
	wg.Done()
}

// Return the word frequencies of the text argument.
//
// Split load optimally across processor cores.
func WordCount(text string) map[string]int {
	ch := make(chan map[string]int)
	wg := new(sync.WaitGroup)
	wg.Add(threads)

	// partly copy pasted from my palinda-1 "maps" file
	freqs := make(map[string]int)

	words := strings.Fields(text)

	words_len := len(words)

	// calculate which thread counts what, by using offsets in the `words` array!
	offsets := make([]int, threads)
	for i := 0; i < threads; i++ {
		offsets[i] = int(math.Round(float64(i) / float64(threads) * float64(words_len)))
	}

	// send the workers to work!
	for i := 0; i < threads-1; i++ {
		go count_partial(words, offsets[i], offsets[i+1], ch, wg)
	}

	// and send the last one too, he's special!
	go func() {
		go count_partial(words, offsets[threads-1], words_len, ch, wg)
	}()

	for i := 0; i < threads; i++ {
		select {
		case new_map := <-ch:
			for word, count := range new_map {
				freqs[word] += count
			}
		}
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
	fmt.Printf("amount of runs: %d\n", numRuns)
	fmt.Printf("total time: %d ms\n", runtimeMillis)
	average := float64(runtimeMillis) / float64(numRuns)
	fmt.Printf("average time/run: %.2f ms\n", average)
}

func main() {
	data, _ := ioutil.ReadFile(DataFile)

	numRuns := 100
	runtimeMillis := benchmark(string(data), numRuns)
	printResults(runtimeMillis, numRuns)
}
