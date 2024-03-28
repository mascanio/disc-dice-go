package parser

import (
	"fmt"
	"strconv"
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

func buildRollerFromInput(s string) (dice.Roller, error) {
	s, adder, err := splitAdder(s)
	if err != nil {
		return nil, err
	}

	var roller dice.Roller
	for _, parser := range parsers {
		roller = parser(s)
		if roller != nil {
			break
		}
	}

	if adder != 0 {
		roller = dice.RollAdder{Base: roller, Adder: adder}
	}
	return roller, nil
}

func InputToMessager(s string) (messager.Messager, error) {
	defer func(timeStart time.Time) {
		fmt.Println("Time elapsed input: ", time.Since(timeStart))
	}(time.Now())
	diceRoller, err := buildRollerFromInput(s)
	if err != nil {
		return nil, err
	}
	return diceRoller.Roll(), nil
}
