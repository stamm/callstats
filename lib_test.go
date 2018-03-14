package callstats_test

import (
	"testing"

	"github.com/stamm/callstats"
	"github.com/stretchr/testify/require"
)

func TestMediansImpr(t *testing.T) {
	table := []struct {
		name   string
		window int
		expect []int
		values []int
	}{
		{"Window = 3", 3, []int{-1, 101, 101, 102, 110, 115}, []int{100, 102, 101, 110, 120, 115}},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expect, callstats.GetMediansImpr(test.window, test.values))
		})
	}
}

func TestMediansChanImpr(t *testing.T) {
	table := []struct {
		name   string
		window int
		expect []int
		values []int
	}{
		{"Window = 3", 3, []int{-1, 101, 101, 102, 110, 115}, []int{100, 102, 101, 110, 120, 115}},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			ch := make(chan int)
			go func(vals []int, ch chan int) {
				for _, val := range vals {
					ch <- val
				}
				close(ch)
			}(test.values, ch)
			res := callstats.GetMediansChanImpr(test.window, ch)
			i := 0
			for val := range res {
				require.Equal(t, test.expect[i], val, "on step %d", i)
				i++
			}
		})
	}
}
