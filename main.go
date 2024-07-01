package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Main app objects
var (
	a          fyne.App
	mainWindow fyne.Window
)

// Button widget variables
var (
	startButton *widget.Button
	resetButton *widget.Button
)

// Keeps track of current elapsed seconds
var seconds binding.Int

// Hold global running state
var running bool

func main() {
	// Create app
	a = app.New()
	mainWindow = a.NewWindow("Pay Stopwatch")

	// Set initial time
	seconds = binding.NewInt()
	seconds.Set(0)

	// Clock widgets
	clock := widget.NewLabelWithData(
		// Automatically format time from seconds
		binding.NewSprintf(
			"Time: %d",
			seconds,
		),
	)
	clock.Alignment = fyne.TextAlignCenter
	payClock := widget.NewLabelWithData(
		// Automatically calculate pay from seconds
		binding.NewSprintf(
			"Pay: %d",
			seconds,
		),
	)
	payClock.Alignment = fyne.TextAlignCenter

	// Start/stop and reset widgets
	startButton = widget.NewButton("Start", startButtonCallback)
	startButton.Importance = widget.HighImportance
	resetButton = widget.NewButton("Reset", resetButtonCallback)

	// Create main layout
	content := container.NewVBox(
		clock,
		payClock,
		container.NewBorder(
			nil,
			nil,
			nil,
			resetButton,
			startButton,
		),
	)

	// Create menu bar
	menu := fyne.NewMainMenu(
		fyne.NewMenu(
			"File",
			fyne.NewMenuItem(
				"Change pay rate",
				changePayRate,
			),
		),
		fyne.NewMenu(
			"Help",
			fyne.NewMenuItem(
				"About",
				showAbout,
			),
		),
	)

	// Set window properties
	mainWindow.SetContent(content)
	mainWindow.SetMainMenu(menu)
	mainWindow.Resize(fyne.NewSize(480, 290))
	mainWindow.SetFixedSize(true)

	// Run app
	mainWindow.Show()
	a.Run()
}
