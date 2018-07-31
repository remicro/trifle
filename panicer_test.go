package trifle

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicer(t *testing.T) {
	t.Run("expect panic", func(t *testing.T) {
		err := errors.New("expect error")
		assert.Panics(t, func() {
			Panicer(err)
		})
	})
	t.Run("expect no panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			Panicer(nil)
		})
	})
}
