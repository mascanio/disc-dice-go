package dice

import (
	"fmt"
	"sort"
)

type MultiDice struct {
	Faces int
	Dices int
}

type MultiDiceResult struct {
	MultiDice
	Resulter
	sum         int
	resultChain string
}

func (d *MultiDiceResult) Message() string {
	return fmt.Sprintf("Rolling %dd%d: %v\nSum: %v\n", d.Dices, d.Faces, d.resultChain, d.sum)
}

func (d *MultiDiceResult) ResultSum() int {
	return d.sum
}

func (d *MultiDiceResult) ResultStr() string {
	return d.resultChain
}

func (d MultiDice) Roll() Resulter {
	sum := 0
	dice := GenericDice{Faces: d.Faces}
	results := make([]Resulter, d.Dices)
	for i := 0; i < d.Dices; i++ {
		results[i] = dice.Roll()
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].ResultSum() < results[j].ResultSum()
	})

	resultStr := "["
	for i := 0; i < d.Dices; i++ {
		r := results[i]
		sum += r.ResultSum()
		resultStr += r.ResultStr()
		if i < d.Dices-1 {
			resultStr += ", "
		}
	}
	resultStr += "]"

	return &MultiDiceResult{MultiDice: d, sum: sum, resultChain: resultStr}
}