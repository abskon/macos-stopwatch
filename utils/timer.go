package utils

import (
	"fmt"
	"time"
)

type Timer struct {
	start   time.Time
	elapsed time.Duration
	running bool
}

func NewTimer() *Timer {
	return &Timer{}
}

func (t *Timer) Start() {
	if !t.running {
		t.start = time.Now()
		t.running = true
	}
}

func (t *Timer) Stop() {
	if t.running {
		t.elapsed += time.Since(t.start)
		t.running = false
	}
}

func (t *Timer) Reset() {
	t.elapsed = 0
	t.running = false
}

func (t *Timer) Str() string {
	elapsed := t.elapsed
	if t.running {
		elapsed += time.Since(t.start)
	}

	seconds := int(elapsed.Seconds())
	milliseconds := int(elapsed.Milliseconds()) % 1000
	return fmt.Sprintf("%d.%03d", seconds, milliseconds)
}

func (t *Timer) IsRunning() bool {
	return t.running
}
