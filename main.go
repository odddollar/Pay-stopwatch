package main

import (
	"fmt"
	"image/color"
	"time"

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

	// Custom binding for clock widget
	clockString := binding.NewString()
	seconds.AddListener(binding.NewDataListener(func() {
		// Get elapsed seconds and parse duration
		s, _ := seconds.Get()
		d, _ := time.ParseDuration(fmt.Sprintf("%ds", s))

		// Format string
		clockString.Set(fmt.Sprintf(
			"%02d:%02d:%02d",
			int(d.Hours()),
			int(d.Minutes())%60,
			int(d.Seconds())%60,
		))
	}))

	// Custom binding for pay widget
	payString := binding.NewString()
	seconds.AddListener(binding.NewDataListener(func() {
		// Get elapsed seconds and parse duration
		s, _ := seconds.Get()
		d, _ := time.ParseDuration(fmt.Sprintf("%ds", s))

		// Calculate pay
		pay := d.Hours() * a.Preferences().FloatWithFallback("payRate", 25.0)
		payString.Set(fmt.Sprintf("%.2f", pay))
	}))

	// Clock widget
	clock := NewCustomLabel(clockString)
	clock.SetAlignment(fyne.TextAlignCenter)
	clock.SetFontSize(22)

	// Pay widget
	pay := NewCustomLabel(payString)
	pay.SetAlignment(fyne.TextAlignCenter)
	pay.SetFontSize(22)

	// Clock widget labels
	clockLabel := canvas.NewText("Time", color.Black)
	clockLabel.Alignment = fyne.TextAlignCenter
	clockLabel.TextStyle.Bold = true
	clockLabel.TextSize = 24
	payLabel := canvas.NewText("Pay", color.Black)
	payLabel.Alignment = fyne.TextAlignCenter
	payLabel.TextStyle.Bold = true
	payLabel.TextSize = 24

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
		payLabel,
		pay,
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
