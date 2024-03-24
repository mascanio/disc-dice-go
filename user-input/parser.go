package userinput

import (
	"fmt"
	"strconv"

	"github.com/mascanio/disc-dice-go/dice"
	"github.com/mascanio/disc-dice-go/messager"
	nre "github.com/mascanio/regexp-named"
)

var RE_DICE = nre.MustCompile(`(?P<n>\d+)?d(?P<d>\d+)`)
var MAX_N_DICES = 100
var MAX_FACES_DICE = 100

func getNDicesAndFaces(s string) (int, int, error) {
	_, mm := RE_DICE.FindStringNamed(s)
	if mm == nil {
		return 0, 0, fmt.Errorf("no dice found")
	}
	nDices, err := strconv.Atoi(mm["n"])
	if err != nil {
		nDices = 1
	}
	nFaces, err := strconv.Atoi(mm["d"])
	if err != nil {
		return 0, 0, fmt.Errorf("error converting d to int")
	}
	if nDices > MAX_N_DICES || nFaces > MAX_FACES_DICE {
		return 0, 0, fmt.Errorf("too many dices or incorrect dice faces")
	}
	if nFaces < 2 {
		return 0, 0, fmt.Errorf("dice faces must be at least 2")
	}
	return nDices, nFaces, nil
}

func InputToMessager(s string) (messager.Messager, error) {
	nDices, nFaces, err := getNDicesAndFaces(s)
	if err != nil {
		return nil, err
	}
	switch nDices {
	case 1:
		return dice.GenericDice{Faces: nFaces}.Roll(), nil
	default:
		return dice.MultiDice{Faces: nFaces, Dices: nDices}.Roll(), nil
	}
}
