package stacklog

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stdout`. You can also set this to
	// something more adventorous, such as logging to Kafka.
	Out io.Writer
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` for which
	// TextFormatter is the default. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter
	// The logging level the logger should log at. This is typically (and defaults
	// to) `log.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged. `log.Debug` is useful in
	Level Level
	// Used to sync writing to the log.
	mu sync.Mutex
	// Set to true to bypass checking for a TTY before outputting colors.
	IsColored bool
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.Level >= DebugLevel {
		NewEntry(logger).Debug(args...)
	}
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.Level >= DebugLevel {
		NewEntry(logger).Debugf(format, args...)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.Level >= InfoLevel {
		NewEntry(logger).Info(args...)
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.Level >= InfoLevel {
		NewEntry(logger).Infof(format, args...)
	}
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.Level >= WarnLevel {
		NewEntry(logger).Warn(args...)
	}
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.Level >= WarnLevel {
		NewEntry(logger).Warnf(format, args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= ErrorLevel {
		NewEntry(logger).Error(args...)
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.Level >= ErrorLevel {
		NewEntry(logger).Errorf(format, args...)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.Level >= FatalLevel {
		NewEntry(logger).Fatal(args...)
	}
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.Level >= FatalLevel {
		NewEntry(logger).Fatalf(format, args...)
	}
}

func (logger *Logger) Panic(args ...interface{}) {
	if logger.Level >= PanicLevel {
		NewEntry(logger).Panic(args...)
	}
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	if logger.Level >= PanicLevel {
		NewEntry(logger).Panicf(format, args...)
	}
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(TextFormatter),
//      Level: log.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stderr,
		Formatter: new(TextFormatter),
		Level:     InfoLevel,
		IsColored: false,
	}
}
