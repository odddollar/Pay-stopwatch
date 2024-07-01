package main

// Resets timer and running state
func resetButtonCallback() {
	running = false
	seconds.Set(0)
	startButton.SetText("Start")
}
