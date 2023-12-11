package main

import (
	"fmt"
	"os"
	"time"

	"github.com/code_challenge/merge"
)

func main() {
	startTime := time.Now()
	intervalInput := os.Getenv("INTERVALS")
	if intervalInput == "" {
		fmt.Println("Intervals must not be empty")
		return
	}

	intervals, err := merge.ParseInterval(intervalInput)
	if err != nil {
		fmt.Printf("Failure while parsing Intervals: %v\n", err)
		return
	}

	overlappingIntervals := merge.Merge(intervals)
	endTime := time.Now()
	diffTime := endTime.Sub(startTime)
	fmt.Println("Runtime: ", diffTime)
	fmt.Println("OverlappingIntervalls:", overlappingIntervals)
}
