package utils

import (
	"fmt"
	"sync"
	"time"
)

type Clock struct {
	Stop      chan struct{}
	CountDown string
	mu        sync.Mutex
	running   bool
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

func NewClock() *Clock {
	return &Clock{
		Stop:      make(chan struct{}, 1), // Channel for stop signal (buffered to prevent blocking)
		CountDown: "04:00:00",
		mu:        sync.Mutex{},
	}
}

func StartCountdown(clock *Clock, duration int) {
	clock.mu.Lock()
	defer clock.mu.Unlock()

	if clock.running {
		return
	}

	clock.running = true

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	clock.Stop = make(chan struct{})

	for range ticker.C {
		clock.CountDown = FormatDuration(currentDuration)

		fmt.Println(clock.CountDown)
		currentDuration--

		if currentDuration < 0 {
			currentDuration = duration // Restart countdown
		}

		select {
		case <-clock.Stop:
			clock.running = false
			return
		default:
		}
	}
}

func StopCountdown(clock *Clock) {
	clock.Stop <- struct{}{}
}
