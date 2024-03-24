package dice

import (
	"github.com/mascanio/disc-dice-go/messager"
)

type Roll interface {
	RollStr() string
	RollSum() int
	messager.Messager
}

type Roller interface {
	Roll() Roll
}
