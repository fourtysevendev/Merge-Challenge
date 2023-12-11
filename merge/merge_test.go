package merge

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_Merge(t *testing.T) {
	//randomIntervals := generateRandomIntervals(10, 100)
	tests := []struct {
		name             string
		inputInterval    []Interval
		expectedInterval []Interval
	}{{
		name:             "Succesfully merged",
		inputInterval:    []Interval{{1, 10}, {25, 30}, {-2, 8}, {30, 2}},
		expectedInterval: []Interval{{25, 30}, {-2, 10}},
	}, {
		name:             "Incompatible Intervals",
		inputInterval:    []Interval{{2, 1}, {30, 0}, {4, -8}, {35, 7}},
		expectedInterval: nil,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			megredIntervals := Merge(test.inputInterval)
			if !cmp.Equal(megredIntervals, test.expectedInterval) {
				t.Errorf("Do not merged correctly: expected %v, got %v", test.expectedInterval, megredIntervals)
			}
		})
	}
}

func Test_findSmallestIntervals(t *testing.T) {
	tests := []struct {
		name                     string
		actual                   []Interval
		expectedSmallestInterval Interval
		expectedSortedIntervals  []Interval
	}{
		{
			name:                     "successful find smallest Interval",
			actual:                   []Interval{{1, 10}, {1, 30}, {23, 87}, {3, 9}, {0, -1}, {33, -2}},
			expectedSmallestInterval: Interval{Start: 1, End: 30},
			expectedSortedIntervals:  []Interval{{1, 10}, {1, 30}, {3, 9}, {23, 87}, {0, -1}, {33, -2}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualSmallestInterval, actualSortIntervals := findSmallestIntervals(test.actual)
			if !cmp.Equal(actualSmallestInterval, test.expectedSmallestInterval) {
				t.Errorf(" smallestInterval : expected %v, got %v", test.expectedSmallestInterval, actualSmallestInterval)
			}
			if !cmp.Equal(actualSortIntervals, test.expectedSortedIntervals) {
				t.Errorf(" sortedIntervals : expected %v, got %v", test.expectedSortedIntervals, actualSortIntervals)
			}
		})
	}
}
