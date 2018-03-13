package callstats

import "sort"

func GetMedians(window int, numbers []int) []int {
	medians := make([]int, 0, len(numbers))
	// medians := make([]int, 0)
	// fmt.Printf("cap(medians) = %+v\n", cap(medians))
	for i := range numbers {
		// fmt.Println(i)
		start := i - window + 1
		if start < 0 {
			start = 0
		}
		// fmt.Printf("start = %d, i = %d\n", start, i)
		// fmt.Printf("numbers[start : i+1] = %+v\n", numbers[start:i+1])
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
