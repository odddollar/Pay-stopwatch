package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	seconds := 0

	// widgets
	var start, reset *widget.Button
	running := false

	clock := widget.NewLabel("Time: 00:00:00")
	start = widget.NewButton("Start", func() {
		// change button text based on running state
		running = !running
		if running {
			start.SetText("Pause")
		} else {
			start.SetText("Start")
		}

		// increment timer
		go func() {
			for range time.Tick(time.Second) {
				if running {
					seconds++
					clock.SetText(formatDuration(seconds))
				} else {
					return
				}
			}
		}()
	})
	reset = widget.NewButton("Reset", func() {
		// stop timer and reset duration
		running = false
		seconds = 0

		clock.SetText(formatDuration(seconds))
	})

	// layout
	buttons := container.NewHBox(
		start,
		reset,
	)

	content := container.NewVBox(
		clock,
		buttons,
	)

	mainWindow.SetContent(content)

	// run the window
	mainWindow.Show()
	app.Run()
}

func formatDuration(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Time: %02d:%02d:%02d", int64(duration.Hours())%24, int64(duration.Minutes())%60, int64(duration.Seconds())%60)
}
