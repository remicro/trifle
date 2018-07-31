package trifle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSex_String(t *testing.T) {
	assert.Equal(t, "male", SexMale.String())
	assert.Equal(t, "female", SexFemale.String())
	assert.Equal(t, "unknown", SexUnknown.String())
}

func TestTraditionalSex(t *testing.T) {
	sexs := map[string]struct{}{}
	for i := 0; i < 1024; i++ {
		sexs[TraditionalSex().String()] = struct{}{}
	}
	assert.Len(t, sexs, 2)
	assert.Contains(t, sexs, SexMale.String())
	assert.Contains(t, sexs, SexFemale.String())
}

func TestNonTraditionalSex(t *testing.T) {
	sexs := map[string]struct{}{}
	for i := 0; i < 1024; i++ {
		sexs[NonTraditionalSex().String()] = struct{}{}
	}
	assert.Len(t, sexs, 3)
	assert.Contains(t, sexs, SexMale.String())
	assert.Contains(t, sexs, SexFemale.String())
	assert.Contains(t, sexs, SexUnknown.String())
}
