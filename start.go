package main

import "time"

// Starts and pauses timer counting
func startButtonCallback() {
	// Change button text based on running state
	running = !running
	if running {
		startButton.SetText("Pause")
	} else {
		startButton.SetText("Start")
	}

	go func() {
		// Run every second
		for range time.Tick(time.Second) {
			if running {
				// Increment second counter
				s, _ := seconds.Get()
				s++
				seconds.Set(s)
			} else {
				return
			}
		}
	}()
}
