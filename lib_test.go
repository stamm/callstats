package callstats_test

import (
	"callstats"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullImpr(t *testing.T) {
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

func TestFullChanImpr(t *testing.T) {
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
				require.Equal(t, test.expect[i], val)
				i++
			}
		})
	}
}

func BenchmarkAllImpr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callstats.GetMediansImpr(3, []int{100, 101, 102, 110, 115, 120, 110, 110, 110})
	}
}

func BenchmarkFilesImpr(b *testing.B) {
	benchmarks := []struct {
		filename string
		window   int
	}{
		{"2.csv", 100},
		{"3.csv", 1000},
		{"4.csv", 10000},
	}
	for _, bm := range benchmarks {
		b.Run(bm.filename, func(b *testing.B) {
			numbers, _ := callstats.ReadIntoChan("files/" + bm.filename)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ch := callstats.GetMediansChanImpr(bm.window, numbers)
				for range ch {
				}
			}
		})
	}
}
