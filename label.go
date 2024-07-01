package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

// Custom widget that acts like label with binding, but can change text size
type CustomLabel struct {
	widget.BaseWidget
	text     *canvas.Text
	data     binding.String
	fontSize float32
}

// Create CustomLabel widget with data binding
func NewCustomLabel(data binding.String) *CustomLabel {
	label := &CustomLabel{
		text:     canvas.NewText("", color.Black),
		data:     data,
		fontSize: 14,
	}
	label.ExtendBaseWidget(label)

	// Add listener to data binding that updates text
	label.data.AddListener(binding.NewDataListener(func() {
		value, _ := label.data.Get()
		label.text.Text = value
		label.Refresh()
	}))
	return label
}

// Set font size
func (l *CustomLabel) SetFontSize(size float32) {
	l.fontSize = size
	l.text.TextSize = size
	l.Refresh()
}

// Set text alignment
func (l *CustomLabel) SetAlignment(align fyne.TextAlign) {
	l.text.Alignment = align
	l.Refresh()
}

// Returns new renderer for CustomLabel
func (l *CustomLabel) CreateRenderer() fyne.WidgetRenderer {
	return &customLabelRenderer{label: l}
}

// Renderer for CustomLabel widget
type customLabelRenderer struct {
	label *CustomLabel
}

// Returns minimum size of CustomLabel widget
func (r *customLabelRenderer) MinSize() fyne.Size {
	return r.label.text.MinSize()
}

// Resizes widget to fit available space
func (r *customLabelRenderer) Layout(size fyne.Size) {
	r.label.text.Resize(size)
}

// Refreshes canvas on which text displayed
func (r *customLabelRenderer) Refresh() {
	canvas.Refresh(r.label)
}

// Returns child widgets of CustomLabel
func (r *customLabelRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.label.text}
}

// Does nothing as Custom label doesn't hold any resources
func (r *customLabelRenderer) Destroy() {}
