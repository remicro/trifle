package trifle

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUnexpectedError(t *testing.T) {
	err := UnexpectedError()
	require.Error(t, err)
	assert.Contains(t, err.Error(), "unexpected error: ")
}
