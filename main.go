package main

import (
	"fmt"
	"os"

	"github.com/code_challenge/merge"
)

func main() {

	intervalInput := os.Getenv("INTERVALS")
	if intervalInput == "" {
		fmt.Println("Die Umgebungsvariable INTERVALS ist nicht gesetzt.")
		return
	}

	intervals, err := merge.ParseInterval(intervalInput)
	if err != nil {
		fmt.Printf("Fehler beim Parsen der Intervalle: %v\n", err)
		return
	}

	overlappingIntervals := merge.Merge(intervals)

	fmt.Println("OverlappingIntervalls:", overlappingIntervals)
}
