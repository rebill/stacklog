package stacklog

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestConvertLevelToString(t *testing.T) {
	assert.Equal(t, "debug", DebugLevel.String())
	assert.Equal(t, "info", InfoLevel.String())
	assert.Equal(t, "warning", WarnLevel.String())
	assert.Equal(t, "error", ErrorLevel.String())
	assert.Equal(t, "fatal", FatalLevel.String())
	assert.Equal(t, "panic", PanicLevel.String())
}

func TestGetSetLevelRace(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if i%2 == 0 {
				SetLevel(InfoLevel)
			} else {
				GetLevel()
			}
		}(i)

	}
	wg.Wait()
}
