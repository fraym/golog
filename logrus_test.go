package golog_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
)

func testLogrusStringWithTime(t *testing.T, log string, buf *bytes.Buffer) {
	assert.Regexp(t, fmt.Sprintf("time=\"[a-zA-Z0-9:\\+\\- ]*\" %s\n", log), buf.String())
	buf.Reset()
}

func TestLogrusLogger(t *testing.T) {
	var buf bytes.Buffer

	logger := golog.NewLogrusLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(golog.DebugLevel)

	logger.Debug().WithField("key", "value").Write("test")
	testLogrusStringWithTime(t, "level=debug msg=test key=value", &buf)

	logger.Debug().WithField("key2", "value2").Write("test")
	testLogrusStringWithTime(t, "level=debug msg=test key2=value2", &buf)

	logger.Debug().WithField("key", "value").WithField("key2", "value2").Write("test2")
	testLogrusStringWithTime(t, "level=debug msg=test2 key=value key2=value2", &buf)

	logger.Debug().WithError(fmt.Errorf("test err")).Write("err")
	testLogrusStringWithTime(t, "level=debug msg=err error=\"test err\"", &buf)
}
