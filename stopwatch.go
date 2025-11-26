package main

import "time"

type State int

const (
	StateStopped State = iota
	StateRunning
	StatePaused
)

type Stopwatch struct {
	start time.Time
	saved time.Duration
	state State
}

func NewStopwatch() *Stopwatch {
	return &Stopwatch{
		state: StateStopped,
	}
}

func (s *Stopwatch) Start() {
	if s.state == StateRunning {
		return
	}
	s.start = time.Now()
	s.state = StateRunning
}

func (s *Stopwatch) Pause() {
	if s.state != StateRunning {
		return
	}

	s.saved += time.Since(s.start)
	s.state = StatePaused
}

func (s *Stopwatch) Reset() {
	s.saved = 0
	s.state = StateStopped
}

func (s *Stopwatch) Elapsed() time.Duration {
	switch s.state {
	case StateRunning:
		return s.saved + time.Since(s.start)
	default:
		return s.saved
	}
}
