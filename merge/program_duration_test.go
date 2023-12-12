package merge

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

func generateRandomInterval(startIntervalSize, endIntervalSize int) Interval {
	start := rand.Intn(startIntervalSize)
	//Make sure that End is greater than Start
	end := start + rand.Intn(endIntervalSize) + 1
	return Interval{Start: start, End: end}
}

func TestRunTime(t *testing.T) {
	// measurement run time
	startTime := time.Now()

	// generate random interval
	var intervals []Interval
	for i := 0; i < 100; i++ {
		interval := generateRandomInterval(100, 100)
		intervals = append(intervals, interval)
	}

	// print intervals
	log.Printf("RunTime test for : [%v] ", intervals)

	overlappingIntervals := Merge(intervals)
	endTime := time.Now()
	diffTime := endTime.Sub(startTime)
	log.Printf("Runtime: %v", diffTime)
	log.Printf("OverlappingIntervalls: %v", overlappingIntervals)
}
