package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	for _, arg := range os.Args[1:] {
		tm, err := strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
		} else {
			start := time.Now()
			fmt.Printf("Start at %s...\n", start)
			time.Sleep(time.Duration(tm) * time.Second)
			end := time.Now()
			fmt.Printf("End   at %s, elapsed %s\n", end, end.Sub(start))
		}
	}
}
