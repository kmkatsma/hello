package os_test

import (
	"github.com/kmkatsma/hello/os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewWrapper(t *testing.T) {
	instance := os.NewWrapper()
	assert.True(t, len(instance.OS) > 0)
}

func TestRunsOnWindows(t *testing.T) {
	instance := os.NewWrapper()
	assert.True(t, instance.RunsOnWindows())
}
