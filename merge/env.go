package merge

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseInterval convert and checks the input of INTERVAL
func ParseInterval(input string) ([]Interval, error) {
	var intervals []Interval

	splitInterval := strings.ReplaceAll(input, " ", "")
	pairs := strings.Split(splitInterval, "][")

	if len(pairs) > 0 {
		pairs[0] = strings.TrimLeft(pairs[0], "[")
	}
	if len(pairs) > 1 {
		pairs[len(pairs)-1] = strings.TrimRight(pairs[len(pairs)-1], "]")
	}
	for _, pair := range pairs {
		limits := strings.Split(pair, ",")
		if len(limits) != 2 {
			return nil, fmt.Errorf("inavlid Interval: %s", pair)
		}

		start, err := strconv.Atoi(strings.TrimSpace(limits[0]))
		if err != nil {
			return nil, err
		}

		end, err := strconv.Atoi(strings.TrimSpace(limits[1]))
		if err != nil {
			return nil, err
		}

		intervals = append(intervals, Interval{Start: start, End: end})
	}

	return intervals, nil
}
