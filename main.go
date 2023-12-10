package main

import (
	"fmt"
	"os"

	"github.com/code_challenge/merge"
)

func main() {

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

	fmt.Println("OverlappingIntervalls:", overlappingIntervals)
}
