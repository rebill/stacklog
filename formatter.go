package stacklog

import (
	"bytes"
	"fmt"
	"runtime"
	"time"
)

const (
	DATETIME_FORMAT = "2006/01/02 15:04:05"
)

// The Formatter interface is used to implement a custom Formatter. It takes an `Entry`.
type Formatter interface {
	Format(*Entry) ([]byte, error)
}

type TextFormatter struct {
	// TimestampFormat to use for display when a full timestamp is printed
	TimestampFormat string
}

func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	if f.TimestampFormat == "" {
		f.TimestampFormat = DATETIME_FORMAT
	}
	// log format : [datetime] [level] [message] [ip] [stack]
	msg := fmt.Sprintf("[%s] [%s] [%s] [%s] [%s]", time.Now().Format(f.TimestampFormat), entry.Level,
		entry.Message, getServerIp(), entry.Stack)

	// checking for a TTY before outputting colors.
	if entry.Logger.IsColored {
		if goos := runtime.GOOS; goos != "windows" {
			msg = colors[entry.Level](msg)
		}
	}

	fmt.Fprint(b, msg)
	b.WriteByte('\n')
	return b.Bytes(), nil
}
