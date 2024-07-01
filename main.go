package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

func main() {
	// Create app
	app := app.New()
	mainWindow := app.NewWindow("Pay Stopwatch")

	content := container.NewVBox(nil)

	// Run app
	mainWindow.SetContent(content)
	mainWindow.Show()
	app.Run()
}
