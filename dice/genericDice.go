package dice

import (
	"fmt"
	"math/rand"

	"github.com/mascanio/disc-dice-go/roller"
)

type GenericDice struct {
	Faces int
}

type GenericDiceResult struct {
	GenericDice
	roller.Resulter
	diceResult int
}

func (d GenericDiceResult) Message() string {
	return fmt.Sprintf("Rolling 1d%d: %v\n", d.Faces, d.diceResult)
}

func (d GenericDiceResult) ResultSum() int {
	return d.diceResult
}

func (d GenericDiceResult) ResultStr() string {
	return fmt.Sprintf("%v", d.diceResult)
}

func (d *GenericDice) Roll() roller.Resulter {
	result := rand.Intn(d.Faces) + 1
	return GenericDiceResult{GenericDice: *d, diceResult: result}
}
