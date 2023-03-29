package golog_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
)

func testZerologStringWithTime(t *testing.T, log string, buf *bytes.Buffer) {
	assert.Regexp(t, fmt.Sprintf(log+"\n", "[a-zA-Z0-9:\\+\\-]*"), buf.String())
	buf.Reset()
}

func TestZerologLogger(t *testing.T) {
	var buf bytes.Buffer

	logger := golog.NewZerologLogger()
	logger.SetOutput(&buf)
	logger.SetLevel(golog.DebugLevel)

	logger.Debug().WithField("key", "value").Write("test")
	testZerologStringWithTime(t, "{\"level\":\"debug\",\"key\":\"value\",\"time\":\"%s\",\"message\":\"test\"}", &buf)

	logger.Debug().WithField("key2", "value2").Write("test")
	testZerologStringWithTime(t, "{\"level\":\"debug\",\"key2\":\"value2\",\"time\":\"%s\",\"message\":\"test\"}", &buf)

	logger.Debug().WithField("key", "value").WithField("key2", "value2").Write("test2")
	testZerologStringWithTime(t, "{\"level\":\"debug\",\"key\":\"value\",\"key2\":\"value2\",\"time\":\"%s\",\"message\":\"test2\"}", &buf)

	logger.Debug().WithError(fmt.Errorf("test err")).Write("err")
	testZerologStringWithTime(t, "{\"level\":\"debug\",\"error\":\"test err\",\"time\":\"%s\",\"message\":\"err\"}", &buf)

	logger.Debug().WithFields(map[string]any{"key": "value"}).Write("test")
	testZerologStringWithTime(t, "{\"level\":\"debug\",\"key\":\"value\",\"time\":\"%s\",\"message\":\"test\"}", &buf)

	logger.SetLevel(golog.InfoLevel)
	logger.Debug().WithFields(map[string]any{"key": "value"}).Write("test")
	assert.Empty(t, buf.String())
	buf.Reset()
}
