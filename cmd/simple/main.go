package main

import (
	"flag"
	"fmt"

	"github.com/stamm/callstats"
)

func main() {
	window := flag.Int("window", 3, "Lenght of the window")
	filename := flag.String("file", "", "File to read")
	flag.Parse()
	if *filename == "" {
		panic("No file")
	}

	numbers, err := callstats.Read(*filename)
	if err != nil {
		panic(err)
	}

	medians := callstats.GetMediansImpr(*window, numbers)
	for _, mediana := range medians {
		fmt.Printf("%d\n", mediana)
	}
}
