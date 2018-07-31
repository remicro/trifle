package trifle

import "math/rand"

type Sex int8

const (
	SexFemale  = Sex(0)
	SexMale    = Sex(1)
	SexUnknown = Sex(2)
)

func (s Sex) String() string {
	switch s % 3 {
	case 0:
		return "female"
	case 1:
		return "male"
	}
	return "unknown"
}

func TraditionalSex() Sex {
	return Sex(rand.Int() % 2)
}

func NonTraditionalSex() Sex {
	return Sex(rand.Int() % 3)
}
