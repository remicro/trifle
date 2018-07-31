package trifle

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestLastName(t *testing.T) {
	name := LastName()
	require.NotEmpty(t, name)
}

func TestFemaleFirstName(t *testing.T) {
	name := FemaleFirstName()
	assert.Contains(t, females, name)
}

func TestMaleFirstName(t *testing.T) {
	name := MaleFirstName()
	assert.Contains(t, males, name)
}

func TestName(t *testing.T) {
	t.Run("generate female name", func(t *testing.T) {
		nameSexGeneratror = func() Sex {
			return SexFemale
		}
		name := Name()
		require.NotEmpty(t, Name())
		t.Log(name)
		names := strings.SplitN(name, " ", 2)
		t.Log(names)
		t.Log(len(names))
		require.Len(t, names, 2)
		assert.NotEmpty(t, names[0])
		assert.NotEmpty(t, names[1])
		assert.Contains(t, females, names[0])
	})
	t.Run("generate male name", func(t *testing.T) {
		nameSexGeneratror = func() Sex {
			return SexMale
		}
		name := Name()
		require.NotEmpty(t, Name())
		t.Log(name)
		names := strings.SplitN(name, " ", 2)
		t.Log(names)
		t.Log(len(names))
		require.Len(t, names, 2)
		assert.NotEmpty(t, names[0])
		assert.NotEmpty(t, names[1])
		assert.Contains(t, males, names[0])
	})

}
