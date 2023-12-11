package merge

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

// findSmallestIntervals find the smallest Interval of the list
// and returns an ascending list of intervals
func findSmallestIntervals(intervals []Interval) (smallestInterval Interval, sortIntervals []Interval) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})
	for i := 0; i < len(intervals); i++ {
		if i == len(intervals)-1 {
			break
		}
		if checkInterval(intervals[i]) {
			if intervals[i].Start == intervals[i+1].Start {
				if intervals[i].End < intervals[i+1].End {
					smallestInterval = intervals[i+1]
					return smallestInterval, intervals
				}

			}
			smallestInterval = intervals[i]
			return smallestInterval, intervals
		}
	}
	return smallestInterval, intervals
}

func checkInterval(interval Interval) bool {
	if interval.Start > interval.End {
		return false
	}
	return true
}

// mergeInterval merges intervals
// like blocks that are plugged into each other
func Merge(intervals []Interval) []Interval {
	// if the list of intervals is empty or has length 1
	// then return the list of intervals
	if len(intervals) <= 1 {
		return intervals
	}

	fittingInterval, sortedIntervals := findSmallestIntervals(intervals)

	// result is the first interval in the list
	//fittingInterval := Interval{Start: intervals[0].Start, End: intervals[0].End}
	result := make([]Interval, 0, len(sortedIntervals))

	for i := 1; i < len(sortedIntervals); i++ {

		// check if the interval is valid
		if !checkInterval(intervals[i]) {
			fmt.Println("Invalid interval, because start is greater than end")
			continue
		}

		// if start from currnt interval is greater than end
		// from merged interval then the intervals do not overlapped
		// if they no overlapped then append the current interval to result
		if fittingInterval.End < sortedIntervals[i].Start {
			result = append(result, sortedIntervals[i])
			continue
		}
		// if they overlap then merge the intervals
		// that means the merged interval has the smallest start point
		// and the biggest end point
		fittingInterval = Interval{Start: fittingInterval.Start, End: greatestPointOf(fittingInterval.End, sortedIntervals[i].End)}
	}
	result = append(result, fittingInterval)
	return result

}
