package dice

import (
	"fmt"
	"math/rand"
)

type GenericDice struct {
	Faces int
}

type genericDiceRoll struct {
	GenericDice
	diceResult int
}

func (d genericDiceRoll) Message() string {
	return fmt.Sprintf("Rolling 1d%d: %v\n", d.Faces, d.diceResult)
}

func (d genericDiceRoll) RollSum() int {
	return d.diceResult
}

func (d genericDiceRoll) RollStr() string {
	return fmt.Sprintf("%v", d.diceResult)
}

func (d GenericDice) Roll() Roll {
	result := rand.Intn(d.Faces) + 1
	return genericDiceRoll{GenericDice: d, diceResult: result}
}
