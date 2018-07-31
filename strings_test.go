package trifle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringN(t *testing.T) {
	t.Run("expect empty string", func(t *testing.T) {
		assert.Empty(t, StringN(0))
	})
	t.Run("expect string", func(t *testing.T) {
		exp := 42
		assert.Len(t, StringN(exp), exp)
	})
}

func TestString(t *testing.T) {
	t.Run("expect non empty string", func(t *testing.T) {
		assert.NotEmpty(t, String())
	})
}
