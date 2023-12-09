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

// greatestPointOf return greates endpoint
func greatestPointOf(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// mergeInterval merges intervals
// like blocks that are plugged into each other
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
	fittingInterval := Interval{Start: intervals[0].Start, End: intervals[0].End}
	result := make([]Interval, 0, len(intervals))

	for i := 1; i < len(intervals); i++ {
		//compareInterval := &result[len(result)-1]

		// if start from currnt interval is greater than end
		// from merged interval then the intervals do not overlapped
		// if they no overlapped then append the current interval to result
		if fittingInterval.End < intervals[i].Start {
			result = append(result, intervals[i])
			break
		}
		// if they overlap then merge the intervals
		// that means the merged interval has the smallest start point
		// and the biggest end point
		fittingInterval = Interval{Start: fittingInterval.Start, End: greatestPointOf(fittingInterval.End, intervals[i].End)}
	}
	result = append(result, fittingInterval)
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
