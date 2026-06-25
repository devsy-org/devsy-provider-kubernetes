package throttledlogger

import (
	"time"

	"github.com/devsy-org/devsy/pkg/log"
)

// ThrottledLogger is a logger that throttles the output,
// i.e. it only logs a message if a certain amount of time has passed since the last log message.
type ThrottledLogger struct {
	timer *Timer
}

func NewThrottledLogger(throttlingInterval time.Duration) *ThrottledLogger {
	return &ThrottledLogger{
		timer: NewTimer(throttlingInterval),
	}
}

func (t *ThrottledLogger) Infof(format string, args ...any) {
	t.log(log.Infof, format, args...)
}

func (t *ThrottledLogger) Debugf(format string, args ...any) {
	t.log(log.Debugf, format, args...)
}

type LoggingFunc func(string, ...any)

func (t *ThrottledLogger) log(loggingFunc LoggingFunc, format string, args ...any) {
	now := time.Now()
	if t.timer.IntervalPassed(now) {
		loggingFunc(format, args...)
		t.timer.Tick(now)
	}
}
