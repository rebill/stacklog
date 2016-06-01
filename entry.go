package stacklog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"
)

type Entry struct {
	Logger *Logger

	// Time at which the log entry was created
	Time time.Time

	// Level the log entry was logged at: Debug, Info, Warn, Error, Fatal or Panic
	Level Level

	// Message passed to Debug, Info, Warn, Error, Fatal or Panic
	Message string

	// Server ip
	Ip string

	// Stack returns a formatted stack trace of the goroutine that calls it.
	Stack []byte
}

func NewEntry(logger *Logger) *Entry {
	return &Entry{
		Logger: logger,
	}
}

// Returns a reader for the entry, which is a proxy to the formatter.
func (entry *Entry) Reader() (*bytes.Buffer, error) {
	serialized, err := entry.Logger.Formatter.Format(entry)
	return bytes.NewBuffer(serialized), err
}

func (entry *Entry) log(level Level, msg string) {
	entry.Level = level
	entry.Message = msg
	if level <= ErrorLevel {
		entry.Stack = Stack()
	} else {
		buf := new(bytes.Buffer)
		_, file, line, ok := runtime.Caller(4)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		fmt.Fprintf(buf, "%s:%d", filename, line)
		entry.Stack = buf.Bytes()
	}

	reader, err := entry.Reader()
	if err != nil {
		entry.Logger.mu.Lock()
		fmt.Fprintf(os.Stderr, "Failed to obtain reader, %v\n", err)
		entry.Logger.mu.Unlock()
	}

	entry.Logger.mu.Lock()
	defer entry.Logger.mu.Unlock()

	_, err = io.Copy(entry.Logger.Out, reader)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to write to log, %v\n", err)
	}

	if level <= PanicLevel {
		panic(entry)
	} else if level <= FatalLevel {
		os.Exit(1)
	}
}

func (entry *Entry) Debug(args ...interface{}) {
	if entry.Logger.Level >= DebugLevel {
		entry.log(DebugLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Debugf(format string, args ...interface{}) {
	if entry.Logger.Level >= DebugLevel {
		entry.log(DebugLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Info(args ...interface{}) {
	if entry.Logger.Level >= InfoLevel {
		entry.log(InfoLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Infof(format string, args ...interface{}) {
	if entry.Logger.Level >= InfoLevel {
		entry.log(InfoLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Warn(args ...interface{}) {
	if entry.Logger.Level >= WarnLevel {
		entry.log(WarnLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Warnf(format string, args ...interface{}) {
	if entry.Logger.Level >= WarnLevel {
		entry.log(WarnLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Error(args ...interface{}) {
	if entry.Logger.Level >= ErrorLevel {
		entry.log(ErrorLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Errorf(format string, args ...interface{}) {
	if entry.Logger.Level >= ErrorLevel {
		entry.log(ErrorLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Fatal(args ...interface{}) {
	if entry.Logger.Level >= FatalLevel {
		entry.log(FatalLevel, fmt.Sprint(args...))
	}
}

func (entry *Entry) Fatalf(format string, args ...interface{}) {
	if entry.Logger.Level >= FatalLevel {
		entry.log(FatalLevel, fmt.Sprintf(format, args...))
	}
}

func (entry *Entry) Panic(args ...interface{}) {
	if entry.Logger.Level >= PanicLevel {
		entry.log(PanicLevel, fmt.Sprint(args...))
	}
}
func (entry *Entry) Panicf(format string, args ...interface{}) {
	if entry.Logger.Level >= PanicLevel {
		entry.log(PanicLevel, fmt.Sprintf(format, args...))
	}
}
