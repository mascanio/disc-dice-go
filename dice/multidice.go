package dice

import (
	"fmt"
	"slices"
)

type MultiDice struct {
	Faces int
	Dices int
}

type multiDiceRoll struct {
	MultiDice
	sum         int
	resultChain string
}

func (d *multiDiceRoll) RollPrefix() string {
	return fmt.Sprintf("Rolling %dd%d: ", d.Dices, d.Faces)
}

func (d *multiDiceRoll) Message() string {
	return fmt.Sprintf("%v: %v\nTotal: %v\n", d.RollPrefix(), d.RollStr(), d.RollSum())
}

func (d *multiDiceRoll) RollSum() int {
	return d.sum
}

func (d *multiDiceRoll) RollStr() string {
	return d.resultChain
}

func (d MultiDice) Roll() Roll {
	sum := 0
	dice := GenericDice{Faces: d.Faces}
	results := make([]Roll, d.Dices)
	for i := 0; i < d.Dices; i++ {
		results[i] = dice.Roll()
	}

	slices.SortFunc(results, func(lhs, rhs Roll) int {
		return lhs.RollSum() - rhs.RollSum()
	})

	resultStr := "["
	for i := 0; i < d.Dices; i++ {
		r := results[i]
		sum += r.RollSum()
		resultStr += r.RollStr()
		if i < d.Dices-1 {
			resultStr += ", "
		}
	}
	resultStr += "]"

	return &multiDiceRoll{MultiDice: d, sum: sum, resultChain: resultStr}
}
