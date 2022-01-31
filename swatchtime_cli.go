package main

import (
	"fmt"
	"time"
)

func main() {
	utc_plusone := time.Now().UTC().Add(time.Hour)
	hours, minutes, seconds := utc_plusone.Clock()
	beats := (float64(seconds) + (float64(minutes) * 60) + (float64(hours) * 3600)) / float64(86.4)
	fmt.Printf("@%.2f Beats", beats)
}

//( UTC+1seconds + ( UTC+1minutes * 60 ) + ( UTC+1hours * 3600 ) ) / 86,4
// fmt.Sprintf("@%.2f Beats", beats)
