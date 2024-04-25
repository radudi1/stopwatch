package stopwatch

import "time"

type StopWatch struct {
	firstStartTime  time.Time
	lastStartTime   time.Time
	runningDuration time.Duration
	running         bool
}

func Start() StopWatch {
	newSw := StopWatch{
		firstStartTime:  time.Now(),
		runningDuration: 0,
		running:         true,
	}
	newSw.lastStartTime = newSw.firstStartTime
	return newSw
}

func (sw *StopWatch) Stop() {
	sw.runningDuration += time.Since(sw.lastStartTime)
	sw.running = false
}

func (sw *StopWatch) Continue() {
	sw.running = true
	sw.lastStartTime = time.Now()
}

func (sw *StopWatch) GetRunningDuration() time.Duration {
	if sw.running {
		sw.Stop()
	}
	return sw.runningDuration
}

func (sw *StopWatch) GetDurationSinceStart() time.Duration {
	return time.Since(sw.firstStartTime)
}
