package callstats_test

import (
	"testing"

	"github.com/stamm/callstats"
)

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		filename string
		window   int
	}{
		{"1.csv", 3},
		{"1_1.csv", 3},
		{"2.csv", 100},
		{"3.csv", 1000},
		{"4.csv", 10000},
	}
	for _, bm := range benchmarks {
		b.Run("Simple;"+bm.filename, func(b *testing.B) {
			numbers, err := callstats.Read("files/" + bm.filename)
			if err != nil {
				panic(err)
			}
			for i := 0; i < b.N; i++ {
				callstats.GetMedians(bm.window, numbers)
			}
		})
		b.Run("Improved;"+bm.filename, func(b *testing.B) {
			numbers, err := callstats.Read("files/" + bm.filename)
			if err != nil {
				panic(err)
			}
			// b.ResetTimer()
			for i := 0; i < b.N; i++ {
				callstats.GetMediansImpr(bm.window, numbers)
			}
		})
		b.Run("ImprovedChan;"+bm.filename, func(b *testing.B) {
			numbers, _ := callstats.ReadIntoChan("files/" + bm.filename)
			// b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ch := callstats.GetMediansChanImpr(bm.window, numbers)
				for range ch {
				}
			}
		})
	}
}
