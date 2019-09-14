package model

import (
	"math/rand"

	"github.com/masakurapa/botmeshi/app/domain/repository"
)

type randomizer struct {
}

// NewRandomizer return Randomizer instance
func NewRandomizer() repository.Randomizer {
	return &randomizer{}
}

func (*randomizer) Intn(i int) int {
	return rand.Intn(i)
}
