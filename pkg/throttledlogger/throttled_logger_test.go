package throttledlogger

import (
	"testing"
	"time"

	"github.com/devsy-org/devsy/pkg/log"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
)

func TestThrottledLogger(t *testing.T) {
	logs := log.InitTestObserved(t, zapcore.InfoLevel)

	now := time.Now()
	interval := time.Millisecond * 50
	timer := &Timer{
		nextMessage:  now.Add(interval),
		tickInterval: interval,
	}

	tLogger := &ThrottledLogger{
		timer: timer,
	}

	// Test: When timer indicates that interval hasn't passed
	tLogger.Infof("This is a test %s", "message")
	assert.Equal(t, 0, logs.Len()) // No log should be recorded

	// Test: When timer indicates that interval has passed
	time.Sleep(interval + time.Millisecond*1)
	tLogger.Infof("This is another test %s", "message")
	assert.Equal(t, 1, logs.Len())
	assert.Equal(t, "This is another test message", logs.All()[0].Message)
}
