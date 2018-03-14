package callstats

import "sort"

// GetMedians with sorting on each step
func GetMedians(window int, numbers []int) []int {
	medians := make([]int, 0, len(numbers))
	for i := range numbers {
		start := i - window + 1
		if start < 0 {
			start = 0
		}
		medians = append(medians, GetMedian(numbers[start:i+1]))
	}
	return medians
}

// GetMedian return median for slice
func GetMedian(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return -1
	}
	sort.Ints(numbers)
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
