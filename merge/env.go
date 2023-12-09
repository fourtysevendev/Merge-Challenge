package merge

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseInterval convert and checks the input of INTERVAL
func ParseInterval(input string) ([]Interval, error) {
	var intervals []Interval

	pairs := strings.Split(input, " ")
	for _, pair := range pairs {
		limits := strings.Split(strings.Trim(pair, "[]"), ",")
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
