package callstats_test

import (
	"testing"

	"github.com/stamm/callstats"
	"github.com/stretchr/testify/require"
)

func TestMedianSimple(t *testing.T) {
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

func TestMediansSimple(t *testing.T) {
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
