package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	// Create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	// Create main layout
	content := container.NewVBox(nil)

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

	// Run app
	mainWindow.SetContent(content)
	mainWindow.SetMainMenu(menu)
	mainWindow.Show()
	app.Run()
}
