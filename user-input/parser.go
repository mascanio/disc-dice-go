package userinput

import (
	"fmt"
	"strconv"

	nre "github.com/mascanio/regexp-named"
)

var RE_DICE = nre.MustCompile(`(?P<n>\d+)?d(?P<d>\d+)`)
var MAX_DICE = 100
var MAX_DICE_TYPE = 100

func GetNDiceType(s string) (int, int, error) {
	_, mm := RE_DICE.FindStringNamed(s)
	if mm == nil {
		return 0, 0, fmt.Errorf("no dice found")
	}
	nDices, err := strconv.Atoi(mm["n"])
	if err != nil {
		nDices = 1
	}
	diceType, err := strconv.Atoi(mm["d"])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting d to int")
	}
	if nDices > MAX_DICE || diceType > MAX_DICE_TYPE {
		return 0, 0, fmt.Errorf("too many dices or incorrect dice type")
	}
	return nDices, diceType, nil
}
