package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const version string = "v1.0.1"
const payRateFile string = "payrate.txt"
const defaultPayRate string = "10"

var payRate float64

func main() {
	// create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	seconds := 0

	// get pay rate
	payRate = getPayRate()

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
					payClock.SetText(calcPay(seconds))
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
		fyne.NewMenuItem("Change pay rate", func() {
			changePayrateWindow(app)
		}),
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
	mainWindow.SetIcon(resourceIconPng)
	mainWindow.Resize(fyne.NewSize(250, 160))
	mainWindow.SetFixedSize(true)
	mainWindow.Show()
	app.Run()
}

func getPayRate() float64 {
	// if unable to read file create file with default value
	f, err := ioutil.ReadFile(payRateFile)
	if err != nil {
		_ = fmt.Errorf("%v", err)
		writePayRate(defaultPayRate)
		f, _ = ioutil.ReadFile(payRateFile)
	}

	payRate, err := strconv.ParseFloat(string(f), 64)
	if err != nil {
		panic(err)
	}

	return payRate
}

func writePayRate(text string) {
	// write text parameter to file
	f, err := os.Create(payRateFile)
	if err != nil {
		panic(err)
	}

	_, err = f.WriteString(text)
	if err != nil {
		panic(err)
	}

	f.Close()
}

func calcPay(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	pay := duration.Hours() * payRate
	return fmt.Sprintf("Pay: %.2f", pay)
}

func formatDuration(seconds int) string {
	duration, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	return fmt.Sprintf("Time: %02d:%02d:%02d", int64(duration.Hours())%24, int64(duration.Minutes())%60, int64(duration.Seconds())%60)
}

func changePayrateWindow(app fyne.App) {
	payRateWindow := app.NewWindow("Change pay rate")

	// widgets
	entry := widget.NewEntry()
	entry.Validator = validation.NewRegexp(`^[0-9]+\.?[0-9]{0,3}$`, "Not valid hourly rate")

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "New hourly pay rate", Widget: entry},
		},
		OnSubmit: func() {
			// write text from entry field into file
			writePayRate(entry.Text)

			payRateWindow.Close()
		},
	}

	// set window to reload pay rate on close
	payRateWindow.SetOnClosed(func() {
		payRate = getPayRate()
	})

	// set window content and run
	payRateWindow.SetContent(form)
	payRateWindow.SetIcon(resourceIconPng)
	payRateWindow.Resize(fyne.NewSize(320, 80))
	payRateWindow.SetFixedSize(true)
	payRateWindow.Show()
}

func showAbout(app fyne.App) {
	appWindow := app.NewWindow("About")

	// widgets
	title := widget.NewLabel("Pay Stopwatch")
	title.Alignment = fyne.TextAlignCenter
	title.TextStyle = fyne.TextStyle{Bold: true}

	appVersion := widget.NewLabel(version)
	appVersion.Alignment = fyne.TextAlignCenter

	rich := widget.NewRichTextFromMarkdown(`
Created by: [Simon Eason (odddollar)](https://github.com/odddollar).

Source available: [github.com/odddollar/Pay-stopwatch](https://github.com/odddollar/Pay-stopwatch).
`)

	closeButton := widget.NewButton("OK", func() {
		appWindow.Close()
	})

	// layout
	content := container.NewVBox(
		title,
		appVersion,
		rich,
		closeButton,
	)

	// set window content and run
	appWindow.SetContent(content)
	appWindow.SetIcon(resourceIconPng)
	appWindow.SetFixedSize(true)
	appWindow.Show()
}
