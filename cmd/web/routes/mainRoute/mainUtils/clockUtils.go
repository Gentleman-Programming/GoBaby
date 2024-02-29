package mainUtils

import (
	"fmt"
	"time"
)

type Clock struct {
	Stop      chan struct{}
	CountDown string
}

var currentDuration = 0

func SetDuration(duration int) {
	currentDuration = duration
}

func FormatDuration(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

var ticker = time.NewTicker(1 * time.Second)

func StartCountdown(clock *Clock, duration int) {
	defer ticker.Stop()

	clock.Stop = make(chan struct{})

	for range ticker.C {
		clock.CountDown = FormatDuration(currentDuration)
		fmt.Println(clock.CountDown)
		currentDuration--

		if currentDuration < 0 {
			currentDuration = duration // Restart countdown
		}
	}
}

func StopCountdown(clock *Clock) {
	time.Sleep(1 * time.Second)
	close(clock.Stop)
}
