package roller

import (
	"github.com/mascanio/disc-dice-go/internal/messager"
)

type Roll interface {
	RollPrefix() string
	RollStr() string
	RollSum() int
	messager.Messager
}

type Roller interface {
	Roll() Roll
}
