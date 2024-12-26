package parser

import (
	"github.com/mascanio/disc-dice-go/internal/roller"
)

func ParseAnima(s string) roller.Roller {
	if s == "a" {
		return roller.AnimaD100{StdOpen: true, AditionalOpen: true, CriticalFailThreshold: 3}
	}
	return nil
}
