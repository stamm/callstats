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

	numbers, errs := callstats.ReadIntoChan(*filename)
	ch := callstats.GetMediansChanImpr(*window, numbers)
	for {
		select {
		case mediana, ok := <-ch:
			if !ok {
				return
			}
			fmt.Printf("%d\n", mediana)
		case err := <-errs:
			fmt.Printf("erros: %s\n", err)
			return
		}
	}
}
