package utils

import (
	"fmt"
	"time"
)

type Stopwatch struct {
	start   time.Time
	elapsed time.Duration
	running bool
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{}
}

func (sw *Stopwatch) Start() {
	if !sw.running {
		sw.start = time.Now()
		sw.running = true
	}
}

func (sw *Stopwatch) Stop() {
	if sw.running {
		sw.elapsed += time.Since(sw.start)
		sw.running = false
	}
}

func (sw *Stopwatch) Reset() {
	sw.elapsed = 0
	sw.running = false
}

func (sw *Stopwatch) Str() string {
	elapsed := sw.elapsed
	if sw.running {
		elapsed += time.Since(sw.start)
	}

	minutes := int(elapsed.Minutes())
	seconds := int(elapsed.Seconds()) % 60
	milliseconds := int(elapsed.Milliseconds()) % 1000
	return fmt.Sprintf("%02d:%02d.%02d", minutes, seconds, milliseconds/10)
}

func (sw *Stopwatch) IsRunning() bool {
	return sw.running
}
