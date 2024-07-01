package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Show change pay rate dialog
func changePayRate() {
	// Entry widget
	entry := widget.NewEntry()
	entry.Validator = validation.NewRegexp(`^[0-9]+(\.?[0-9]{1,2})?$`, "Not valid hourly rate")

	// Set entry text to current rate
	currentRate := a.Preferences().FloatWithFallback("payRate", 25.0)
	entry.SetText(strconv.FormatFloat(currentRate, 'f', 2, 64))

	// Create form layout
	options := []*widget.FormItem{
		{Text: "Hourly pay rate", Widget: entry, HintText: "Decimal number"},
	}

	// Create dialog using form items
	d := dialog.NewForm(
		"Pay rate",
		"Save",
		"Cancel",
		options,
		func(b bool) {
			if b {
				// Update pay rate
				f, _ := strconv.ParseFloat(entry.Text, 64)
				a.Preferences().SetFloat("payRate", f)
			}
		},
		mainWindow,
	)
	d.Resize(fyne.NewSize(290, 175))
	d.Show()
}
