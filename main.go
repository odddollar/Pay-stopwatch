package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
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
	clock := NewCustomLabel(binding.IntToString(seconds))
	clock.SetAlignment(fyne.TextAlignCenter)
	clock.SetFontSize(22)
	payClock := NewCustomLabel(binding.IntToString(seconds))
	payClock.SetAlignment(fyne.TextAlignCenter)
	payClock.SetFontSize(22)

	// Clock widget labels
	clockLabel := canvas.NewText("Time", color.Black)
	clockLabel.Alignment = fyne.TextAlignCenter
	clockLabel.TextStyle.Bold = true
	clockLabel.TextSize = 24
	payClockLabel := canvas.NewText("Pay", color.Black)
	payClockLabel.Alignment = fyne.TextAlignCenter
	payClockLabel.TextStyle.Bold = true
	payClockLabel.TextSize = 24

	// Start/stop and reset widgets
	startButton = widget.NewButton("Start", startButtonCallback)
	startButton.Importance = widget.HighImportance
	resetButton = widget.NewButton("Reset", resetButtonCallback)

	// Create main layout
	content := container.NewVBox(
		layout.NewSpacer(),
		clockLabel,
		clock,
		layout.NewSpacer(),
		payClockLabel,
		payClock,
		layout.NewSpacer(),
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
	mainWindow.Resize(fyne.NewSize(425, 290))
	mainWindow.SetFixedSize(true)

	// Run app
	mainWindow.Show()
	a.Run()
}
