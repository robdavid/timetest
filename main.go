package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	var continuous bool
	var driftThreshold time.Duration
	var err error
	flag.BoolVar(&continuous, "continuous", false, "Run continuously until drift detected")
	flag.DurationVar(&driftThreshold, "drift", 100*time.Millisecond, "Threshold of drift, e.g. 500ms")
	flag.Parse()
	durations := make([]time.Duration, len(flag.Args()))
	for i, arg := range flag.Args() {
		var seconds int
		seconds, err = strconv.Atoi(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s", os.Args[0], err)
			os.Exit(1)
		}
		durations[i] = time.Duration(seconds) * time.Second
	}
	for {
		for _, tm := range durations {
			start := time.Now()
			fmt.Printf("Start at %s...\n", start.Format(time.StampMilli))
			time.Sleep(tm)
			end := time.Now()
			elapsed := end.Sub(start)
			fmt.Printf("End   at %s, elapsed %s\n", end.Format(time.StampMilli), elapsed)
			wallStart := readWall(start)
			wallEnd := readWall(end)
			drift := wallEnd.Sub(wallStart) - elapsed
			if drift != 0 {
				fmt.Printf("Drift of %s detected\n", drift)
				if continuous && (drift > driftThreshold || drift < -driftThreshold) {
					os.Exit(2)
				}
			}
		}
		if !continuous {
			break
		}
	}
}

func readWall(in time.Time) (out time.Time) {
	out = time.Date(in.Year(), in.Month(), in.Day(), in.Hour(), in.Minute(), in.Second(), in.Nanosecond(), in.Location())
	return
}
