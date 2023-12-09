package main

import (
	"fmt"
	"sort"
)

// Interval has a start and end point
type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) []Interval {
	// if the list of intervals is empty or has length 1
	// then return the list of intervals
	if len(intervals) <= 1 {
		return intervals
	}

	// sort the intervals by start point
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	// result is the first interval in the list
	result := make([]Interval, 0, len(intervals))

	for i := 0; i < len(intervals)-1; i++ {
		current := intervals[i]
		compareIntervall := intervals[i+1]

		//
		if current.End > compareIntervall.Start {
			// if current
			// check which end value is greater than
			// and set it in intersection

			result = append(result, intervals[i])

		} else {
			// Keine Überlappung, füge das aktuelle Intervall hinzu
			result = append(result, intervals[i])
		}
	}
	return result

}

func main() {
	intervals := []Interval{
		{25, 30},
		{2, 19},
		{14, 23},
		{4, 8},
	}

	overlappingIntervals := mergeIntervals(intervals)

	fmt.Println("OverlappingIntervalls:", overlappingIntervals)
}
