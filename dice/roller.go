package dice

import (
	"github.com/mascanio/disc-dice-go/messager"
)

type Result struct {
	ResultSum int
}

type Resulter interface {
	ResultStr() string
	ResultSum() int
	messager.Messager
}

type Roller interface {
	Roll() Resulter
}
