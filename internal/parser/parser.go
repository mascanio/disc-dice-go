package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mascanio/disc-dice-go/internal/messager"
	"github.com/mascanio/disc-dice-go/internal/roller"
)

var parsers = []func(string) roller.Roller{
	ParseAnima,
	ParseDice,
	ParseAbility,
}

func splitAdder(s string) (string, int, error) {
	for i, r := range s {
		switch r {
		case '+', '-':
			adder, err := strconv.Atoi(s[i+1:])
			if err != nil {
				return "", 0, fmt.Errorf("error converting adder to int")
			}
			if r == '-' {
				adder = -adder
			}
			return s[:i], adder, nil
		}
	}
	return s, 0, nil
}

func parseBaseRoller(s string) roller.Roller {
	for _, parser := range parsers {
		roller := parser(s)
		if roller != nil {
			return roller
		}
	}
	return nil
}

func buildRollerFromInput(s string) roller.Roller {
	s, adder, err := splitAdder(s)
	if err != nil {
		return nil
	}
	s = strings.TrimSpace(s)
	baseRoller := parseBaseRoller(s)
	if baseRoller == nil {
		return nil
	}
	if adder != 0 {
		baseRoller = roller.RollAdder{Base: baseRoller, Adder: adder}
	}
	return baseRoller
}

func InputToMessager(s string) messager.Messager {
	defer func(timeStart time.Time) {
		fmt.Println("Time elapsed input: ", time.Since(timeStart))
	}(time.Now())
	s = strings.TrimSpace(s)
	diceRoller := buildRollerFromInput(s)
	if diceRoller == nil {
		return nil
	}
	return diceRoller.Roll()
}
