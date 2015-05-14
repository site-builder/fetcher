package randomizer

import (
	"github.com/satori/go.uuid"
)

type Randomizer interface {
	GenerateUUID() string
}

type randomizer struct{}

func CreateRandomizer() Randomizer {
	return &randomizer{}
}

func (_ *randomizer) GenerateUUID() string {
	return uuid.NewV4().String()
}
