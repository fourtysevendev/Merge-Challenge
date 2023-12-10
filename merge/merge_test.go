package merge

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func generateRandomIntervals(count, maxLength int) []Interval {

	var intervals []Interval

	for i := 0; i < count; i++ {
		start := rand.Intn(maxLength)
		end := start + rand.Intn(maxLength-start) + 1

		intervals = append(intervals, Interval{Start: start, End: end})
	}

	return intervals
}

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
		expectedInterval: []Interval{{0, 0}},
	},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			megredIntervals := Merge(test.inputInterval)
			if !cmp.Equal(megredIntervals, test.expectedInterval) {
				t.Errorf("Do not merged correctly: expected %v, got %v", test.expectedInterval, megredIntervals)
			}
		})
	}
}
