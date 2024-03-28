package parser

import (
	"github.com/mascanio/disc-dice-go/roller"
)

func ParseAbility(s string) roller.Roller {
	if s == "ability" {
		return roller.Ability{Base: 100, Name: "ability", Roller: roller.AnimaD100{StdOpen: true, AditionalOpen: true, CriticalFailThreshold: 3}}
	}
	return nil
}
