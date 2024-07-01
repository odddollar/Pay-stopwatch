package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Button widget variables
var (
	startButton *widget.Button
	resetButton *widget.Button
)

// Keeps track of current elapsed seconds
var time binding.Int

func main() {
	// Create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	// Set initial time
	time = binding.NewInt()
	time.Set(0)

	// Clock widgets
	clock := widget.NewLabelWithData(
		// Automatically format time from seconds
		binding.NewSprintf(
			"Time: %d",
			time,
		),
	)
	clock.Alignment = fyne.TextAlignCenter
	payClock := widget.NewLabelWithData(
		// Automatically calculate pay from seconds
		binding.NewSprintf(
			"Pay: %d",
			time,
		),
	)
	payClock.Alignment = fyne.TextAlignCenter

	// Start/stop and reset widgets
	startButton = widget.NewButton("Start", func() {})
	resetButton = widget.NewButton("Reset", func() {})

	// Create main layout
	content := container.NewVBox(
		clock,
		payClock,
		container.NewHBox(
			layout.NewSpacer(),
			startButton,
			layout.NewSpacer(),
			resetButton,
			layout.NewSpacer(),
		),
	)

	// Create menu bar
	menu := fyne.NewMainMenu(
		fyne.NewMenu(
			"File",
			fyne.NewMenuItem(
				"Change pay rate",
				func() {},
			),
		),
		fyne.NewMenu(
			"Help",
			fyne.NewMenuItem(
				"About",
				func() {},
			),
		),
	)

	// Set window properties
	mainWindow.SetContent(content)
	mainWindow.SetMainMenu(menu)
	mainWindow.Resize(fyne.NewSize(250, 160))
	mainWindow.SetFixedSize(true)

	// Run app
	mainWindow.Show()
	app.Run()
}
