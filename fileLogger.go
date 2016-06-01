package stacklog

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	DATEFORMAT        = "2006-01-02"
	DEFAULT_FILE_NAME = "application.log"
	DEFAULT_LOG_SCAN  = 60
)

type FileLogger struct {
	mu          *sync.RWMutex
	logPath     string
	logFileName string
	logFile     *os.File
	date        *time.Time
}

// NewDailyLogger return a logger split by daily
func NewDailyLogger(logPath string, fileName string) *FileLogger {
	if fileName == "" {
		fileName = DEFAULT_FILE_NAME
	}
	dailyLogger := &FileLogger{
		mu:          new(sync.RWMutex),
		logPath:     logPath,
		logFileName: fileName,
	}
	dailyLogger.initLogger()
	return dailyLogger
}

// init fileLogger split by daily
func (f *FileLogger) initLogger() {
	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))

	f.date = &t
	f.mu.Lock()
	defer f.mu.Unlock()

	if !f.isMustSplit() {
		var err error
		if !isExist(f.logPath) {
			err := os.MkdirAll(f.logPath, 0755)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to mkdir, %v\n", err)
			}
		}
		logFile := filepath.Join(f.logPath, f.logFileName)
		f.logFile, err = os.OpenFile(logFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to openfile, %v\n", err)
		}
	} else {
		f.split()
	}

	go f.fileMonitor()
}

// used for determine the fileLogger f is time to split.
// daily: once the current fileLogger stands for yesterday need to split
func (f *FileLogger) isMustSplit() bool {
	t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
	if t.After(*f.date) {
		return true
	}
	return false
}

// Split fileLogger
func (f *FileLogger) split() {
	logFile := filepath.Join(f.logPath, f.logFileName)
	logFileBak := logFile + "." + f.date.Format(DATEFORMAT)
	if !isExist(logFileBak) && f.isMustSplit() {
		r, err := os.Open(logFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open file, %v\n", err)
		}
		defer r.Close()

		w, err := os.Create(logFileBak)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create file, %v\n", err)
		}
		defer w.Close()

		_, err = io.Copy(w, r)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to copy file, %v\n", err)
		}

		err = os.Truncate(logFile, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to truncate file, %v\n", err)
		}

		t, _ := time.Parse(DATEFORMAT, time.Now().Format(DATEFORMAT))
		f.date = &t
	}
}

// After some interval time, goto check the current fileLogger's size or date
func (f *FileLogger) fileMonitor() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("FileLogger's FileMonitor() catch panic: %v\n", err)
		}
	}()

	timer := time.NewTicker(time.Duration(DEFAULT_LOG_SCAN) * time.Second)
	for {
		select {
		case <-timer.C:
			f.fileCheck()
		}
	}
}

// If the current fileLogger need to split, just split
func (f *FileLogger) fileCheck() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("FileLogger's FileCheck() catch panic: %v\n", err)
		}
	}()

	if f.isMustSplit() {
		f.mu.Lock()
		defer f.mu.Unlock()

		f.split()
	}
}

// passive to close fileLogger
func (f *FileLogger) Close() error {
	return f.logFile.Close()
}
