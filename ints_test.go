package trifle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInt(t *testing.T) {
	assert.NotPanics(t, func() {
		Int()
	})
}
