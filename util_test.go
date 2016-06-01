package stacklog

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUtilGetServerIp(t *testing.T) {
	assert.NotEqual(t, "127.0.0.1", getServerIp())
}

func TestUtilFolderIsExist(t *testing.T) {
	assert.True(t, isExist("/tmp"))
}
