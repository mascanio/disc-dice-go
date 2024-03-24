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

func (d *genericDiceRoll) RollPrefix() string {
	return "Rolling 1d" + fmt.Sprintf("%d: ", d.Faces)
}

func (d *genericDiceRoll) Message() string {
	return fmt.Sprintf("%v %v\n", d.RollPrefix(), d.RollSum())
}

func (d *genericDiceRoll) RollSum() int {
	return d.diceResult
}

func (d *genericDiceRoll) RollStr() string {
	return fmt.Sprintf("%v", d.diceResult)
}

func (d GenericDice) Roll() Roll {
	result := rand.Intn(d.Faces) + 1
	return &genericDiceRoll{GenericDice: d, diceResult: result}
}
