package golog_test

import (
	"testing"

	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
)

func TestMockLogger(t *testing.T) {
	assert.Implements(t, new(golog.Logger), new(golog.MockLogger))
}
