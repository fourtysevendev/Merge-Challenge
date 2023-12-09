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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
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
	fittingInterval := Interval{Start: intervals[0].Start, End: intervals[0].End}
	result := make([]Interval, 0, len(intervals))

	for i := 1; i < len(intervals); i++ {
		//compareInterval := &result[len(result)-1]

		//
		if fittingInterval.End > intervals[i].End && fittingInterval.Start < intervals[i].End {
			// if current
			// check which end value is greater than
			// and set it in intersection
			fittingInterval = Interval{Start: fittingInterval.Start, End: max(fittingInterval.End, intervals[i].End)}

		} else {
			// No overlap, add the current interval
			result = append(result, intervals[i])
		}
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
