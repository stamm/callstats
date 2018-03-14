package callstats

import "sort"

const (
	chanCap = 100
)

// GetMediansImpr get medians with sorted window slice
func GetMediansImpr(window int, numbers []int) []int {
	medians := make([]int, 0, len(numbers))
	windowSlice := make([]int, 0, window)
	for i, number := range numbers {
		lastValue := 0
		if i > window {
			lastValue = numbers[i-window]
		}
		windowSlice = getNewWindows(window, number, windowSlice, lastValue)
		medians = append(medians, GetMedianImpr(windowSlice))
	}
	return medians
}

// GetMediansChanImpr get medians with sorted window slice
func GetMediansChanImpr(window int, numbers <-chan int) <-chan int {
	medians := make(chan int, chanCap)
	go func() {
		lastInsertedValues := make([]int, 0, window)
		windowSlice := make([]int, 0, window)
		for number := range numbers {
			lastValue := 0
			if len(lastInsertedValues) > 0 {
				lastValue = lastInsertedValues[0]
			}
			windowSlice = getNewWindows(window, number, windowSlice, lastValue)

			if len(lastInsertedValues) == window {
				lastInsertedValues = lastInsertedValues[1:]
			}
			lastInsertedValues = append(lastInsertedValues, number)

			medians <- GetMedianImpr(windowSlice)
		}
		close(medians)
	}()
	return medians
}

func getNewWindows(window int, number int, windowSlice []int, lastValue int) []int {
	if len(windowSlice) <= 0 {
		windowSlice = append(windowSlice, number)
	} else {
		var index int
		if len(windowSlice) >= window {
			// delete the oldest element
			index = sort.SearchInts(windowSlice, lastValue)
			windowSlice = append(windowSlice[:index], windowSlice[index+1:]...)
		}
		index = sort.SearchInts(windowSlice, number)
		if index < len(windowSlice) {
			windowSlice = append(windowSlice[:index], append([]int{number}, windowSlice[index:]...)...)
		} else {
			windowSlice = append(windowSlice, number)
		}
	}
	return windowSlice
}

// GetMedianImpr get medians by sorted slice
func GetMedianImpr(numbers []int) int {
	n := len(numbers)
	if n == 1 {
		return -1
	}
	if n%2 == 0 {
		return (numbers[n/2-1] + numbers[n/2]) / 2
	}
	return numbers[n/2]
}
