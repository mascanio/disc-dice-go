package parser

import (
	"github.com/mascanio/disc-dice-go/dice"
)

func ParseAnima(s string) dice.Roller {
	if s == "a" {
		return dice.AnimaD100{StdOpen: true, AditionalOpen: true, CriticalFailThreshold: 3}
	}
	return nil
}
