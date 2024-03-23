package dice

import (
	"math/rand"
	"strconv"
)

type GenericDice struct {
	Faces int
}

func (d *GenericDice) Roll() DiceResult {
	result := rand.Intn(d.Faces) + 1
	return DiceResult{result, strconv.Itoa(result)}
}
