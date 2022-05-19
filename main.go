package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const version string = "v0.1"

func main() {
	// create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	seconds := 0

	// get payrate
	f, err := ioutil.ReadFile("payrate.txt")
	if err != nil {
		panic(err)
	}

	payrate, err := strconv.ParseFloat(string(f), 64)
	if err != nil {
		panic(err)
	}

	// widgets
	var start, reset *widget.Button
	running := false

	clock := widget.NewLabel("Time: 00:00:00")
	clock.Alignment = fyne.TextAlignCenter

	payClock := widget.NewLabel("Pay: 0.00")
	payClock.Alignment = fyne.TextAlignCenter

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

					// update seconds display
					clock.SetText(formatDuration(seconds))

					// update pay display
					payClock.SetText(calcPay(seconds, payrate))
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

		start.SetText("Start")
		payClock.SetText("Pay: 0.00")
		clock.SetText(formatDuration(seconds))
	})

	// layout
	buttons := container.New(
		layout.NewHBoxLayout(),
		layout.NewSpacer(),
		start,
		layout.NewSpacer(),
		reset,
		layout.NewSpacer(),
	)

	content := container.NewVBox(
		clock,
		payClock,
		buttons,
	)

	// create menu bar
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Change payrate", func() {}),
	)
	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			showAbout(app)
		}),
	)

	mainWindow.SetMainMenu(fyne.NewMainMenu(fileMenu, helpMenu))

	// set window content
	mainWindow.SetContent(content)

	// run the window
	mainWindow.SetMaster()
	mainWindow.Resize(fyne.NewSize(250, 160))
	//mainWindow.SetFixedSize(true)
	mainWindow.Show()
	app.Run()
}

func calcPay(seconds int, payrate float64) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	pay := duration.Hours() * payrate
	return fmt.Sprintf("Pay: %.2f", pay)
}

func formatDuration(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Time: %02d:%02d:%02d", int64(duration.Hours())%24, int64(duration.Minutes())%60, int64(duration.Seconds())%60)
}

func showAbout(app fyne.App) {
	appWindow := app.NewWindow("About")

	// widgets
	title := widget.NewLabel("Pay Stopwatch")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	appVersion := widget.NewLabel(version)
	appVersion.Alignment = fyne.TextAlignCenter

	// layout
	content := container.NewVBox(
		title,
		appVersion,
	)

	// set window content and run
	appWindow.SetContent(content)
	appWindow.Show()
}
