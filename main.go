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
	fmt.Println(intervals)
	return intervals

}

func main() {
	intervals := []Interval{
		{1, 5},
		{2, 8},
		{6, 9},
		{10, 15},
		{13, 18},
		{20, 40},
	}

	overlappingIntervals := mergeIntervals(intervals)

	fmt.Println(overlappingIntervals)
}
