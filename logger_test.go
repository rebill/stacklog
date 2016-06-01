package stacklog

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func LogAndAssertText(t *testing.T, log func(*Logger), assertions func(fields map[string]string)) {
	var buffer bytes.Buffer

	logger := New()
	logger.Out = &buffer
	logger.Formatter = &TextFormatter{}

	log(logger)

	reg := regexp.MustCompile(`\[(?P<datetime>.*)\] \[(?P<level>.*)\] \[(?P<message>.*)\] \[(?P<ip>.*)\] (.+?)`)
	names := reg.SubexpNames()
	match := reg.FindStringSubmatch(buffer.String())
	fields := make(map[string]string)
	for i, n := range match {
		if i == 0 {
			continue
		}
		fields[names[i]] = n
	}
	assertions(fields)
}

func TestDebug(t *testing.T) {
	message := "Debug message"
	LogAndAssertText(t, func(log *Logger) {
		log.Level = DebugLevel
		log.Debug(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "debug", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}

func TestInfo(t *testing.T) {
	message := "Info message"
	LogAndAssertText(t, func(log *Logger) {
		log.Info(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "info", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}

func TestWarn(t *testing.T) {
	message := "Warning message"
	LogAndAssertText(t, func(log *Logger) {
		log.Warn(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "warning", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}

func TestError(t *testing.T) {
	message := "Error message"
	LogAndAssertText(t, func(log *Logger) {
		log.Error(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "error", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}

func TestFatal(t *testing.T) {
	message := "Fatal message"
	LogAndAssertText(t, func(log *Logger) {
		log.Fatal(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "fatal", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}

func TestPanic(t *testing.T) {
	message := "Panic message"
	LogAndAssertText(t, func(log *Logger) {
		log.Panic(message)
	}, func(fields map[string]string) {
		assert.Equal(t, "panic", fields["level"])
		assert.Equal(t, message, fields["message"])
	})
}
