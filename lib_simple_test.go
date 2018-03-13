package callstats_test

import (
	"callstats"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTree(t *testing.T) {
	table := []struct {
		name   string
		expect int
		values []int
	}{
		{"One element", -1, []int{1}},
		{"Two elements", 101, []int{100, 102}},
		{"Tree elements", 101, []int{100, 102, 101}},
		{"Tree elements", 102, []int{102, 101, 110}},
		{"Tree elements", 110, []int{101, 110, 120}},
		{"Tree elements", 115, []int{110, 120, 115}},
	}
	for _, test := range table {
		t.Run(test.name, func(t *testing.T) {
			require.Equal(t, test.expect, callstats.GetMedian(test.values))
		})
	}
}

func TestFull(t *testing.T) {
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
			require.Equal(t, test.expect, callstats.GetMedians(test.window, test.values))
		})
	}
}

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		callstats.GetMedians(3, []int{100, 101, 102, 110, 115, 120, 110, 110, 110})
	}
}

func BenchmarkFile2(b *testing.B) {
	numbers, _ := callstats.Read("./files/2.csv")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		callstats.GetMedians(100, numbers)
	}
}

func BenchmarkFiles(b *testing.B) {
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
			numbers, err := callstats.Read("files/" + bm.filename)
			if err != nil {
				panic(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				callstats.GetMedians(bm.window, numbers)
			}
		})
	}
}
