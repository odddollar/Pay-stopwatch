package main

import (
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show about information in dialog
func showAbout() {
	// Create layout
	// Separate markdown widget for better spacing
	d := container.NewVBox(
		widget.NewRichTextFromMarkdown(fmt.Sprintf("Version: **%s**", a.Metadata().Version)),
		widget.NewRichTextFromMarkdown("Created by: [Simon Eason (odddollar)](https://github.com/odddollar)"),
		widget.NewRichTextFromMarkdown("Source available: [github.com/odddollar/Pay-stopwatch](https://github.com/odddollar/Pay-stopwatch)"),
	)

	// Show dialog with layout
	dialog.ShowCustom(
		"About",
		"OK",
		d,
		mainWindow,
	)
}
