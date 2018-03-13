package callstats

import (
	"bufio"
	"os"
	"strconv"
)

// Read reads file and return slice of numbers
func Read(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		val, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, int(val))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return numbers, nil
}

// ReadIntoChan reads file and return chan of numbers
func ReadIntoChan(filename string) (chan int, chan error) {
	numbers := make(chan int, chanCap)
	errs := make(chan error)
	go func() {
		defer close(numbers)
		file, err := os.Open(filename)
		if err != nil {
			errs <- err
			return
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				continue
			}
			val, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				errs <- err
				return
			}
			numbers <- int(val)
		}
		if err := scanner.Err(); err != nil {
			errs <- err
			return
		}
	}()
	return numbers, errs
}
