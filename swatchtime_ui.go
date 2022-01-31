package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	utc_plusone := time.Now().UTC().Add(time.Hour)
	hours, minutes, seconds := utc_plusone.Clock()
	beats := (float64(seconds) + (float64(minutes) * 60) + (float64(hours) * 3600)) / float64(86.4)
	formatted := fmt.Sprintf("@%.2f Beats", beats)
	clock.SetText(formatted)
}

func main() {
	a := app.New()
	w := a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateTime(clock)

	w.SetContent(clock)
	go func() {
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()
	w.ShowAndRun()
}
