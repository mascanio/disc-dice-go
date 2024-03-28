package parser

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mascanio/disc-dice-go/dice"
	"github.com/mascanio/disc-dice-go/messager"
)

var parsers = []func(string) dice.Roller{
	ParseAnima,
	ParseDice,
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

func parseBaseRoller(s string) dice.Roller {
	for _, parser := range parsers {
		roller := parser(s)
		if roller != nil {
			return roller
		}
	}
	return nil
}

func buildRollerFromInput(s string) dice.Roller {
	s, adder, err := splitAdder(s)
	if err != nil {
		return nil
	}
	s = strings.TrimSpace(s)
	roller := parseBaseRoller(s)
	if roller == nil {
		return nil
	}
	if adder != 0 {
		roller = dice.RollAdder{Base: roller, Adder: adder}
	}
	return roller
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
