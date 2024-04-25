package stopwatch

import "time"

type StopWatch struct {
	lastStartTime time.Time
	totalDuration time.Duration
}

func Start() *StopWatch {
	return &StopWatch{
		lastStartTime: time.Now(),
	}
}

func (sw *StopWatch) Stop() {
	sw.totalDuration += time.Since(sw.lastStartTime)
}

func (sw *StopWatch) GetTotalDuration() time.Duration {
	return sw.totalDuration
}
